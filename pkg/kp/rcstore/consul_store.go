package rcstore

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/square/p2/Godeps/_workspace/src/github.com/hashicorp/consul/api"
	"github.com/square/p2/Godeps/_workspace/src/github.com/pborman/uuid"

	"github.com/square/p2/pkg/labels"
	"github.com/square/p2/pkg/logging"
	"github.com/square/p2/pkg/pods"
	"github.com/square/p2/pkg/rc/fields"
)

const kvPrefix = "replication_controllers"

type kvPair struct {
	key   string
	value []byte
}

type consulKV interface {
	Get(key string, opts *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)
	Keys(prefix string, separator string, opts *api.QueryOptions) ([]string, *api.QueryMeta, error)
	List(prefix string, opts *api.QueryOptions) (api.KVPairs, *api.QueryMeta, error)
	Put(pair *api.KVPair, opts *api.WriteOptions) (*api.WriteMeta, error)
	DeleteTree(prefix string, opts *api.WriteOptions) (*api.WriteMeta, error)
}

type consulStore struct {
	kv     consulKV
	logger logging.Logger
}

var _ Store = &consulStore{}

func NewConsul(client *api.Client) *consulStore {
	return &consulStore{
		kv:     client.KV(),
		logger: logging.DefaultLogger,
	}
}

func (s *consulStore) Create(manifest pods.Manifest, nodeSelector labels.Selector, podLabels labels.Set) (fields.RC, error) {
	id := fields.ID(uuid.New())

	buf := bytes.Buffer{}
	err := manifest.Write(&buf)
	if err != nil {
		return fields.RC{}, err
	}

	err = s.put(id, []kvPair{
		kvPair{key: "disabled", value: []byte("false")},
		kvPair{key: "node_selector", value: []byte(nodeSelector.String())},
		kvPair{key: "pod_labels", value: []byte(podLabels.String())},
		kvPair{key: "pod_manifest", value: buf.Bytes()},
		kvPair{key: "replicas_desired", value: []byte("0")},
	})

	if err != nil {
		return fields.RC{}, err
	}

	return fields.RC{
		Id:              id,
		Manifest:        manifest,
		NodeSelector:    nodeSelector,
		PodLabels:       podLabels,
		ReplicasDesired: 0,
		Disabled:        false,
	}, nil
}

func (s *consulStore) Get(id fields.ID) (fields.RC, error) {
	listed, _, err := s.kv.List(idPrefix(id), nil)
	if err != nil {
		return fields.RC{}, nil
	}

	rcMap := s.kvpsToRcs(listed)

	rc, ok := rcMap[id]
	if !ok {
		return fields.RC{}, fmt.Errorf("No such replication controller %s", id)
	}

	return *rc, nil
}

func (s *consulStore) List() ([]fields.RC, error) {
	listed, _, err := s.kv.List(kvPrefix, nil)
	if err != nil {
		return []fields.RC{}, nil
	}

	rcMap := s.kvpsToRcs(listed)
	rcs := make([]fields.RC, 0)

	for _, rc := range rcMap {
		rcs = append(rcs, *rc)
	}

	return rcs, nil
}

func (s *consulStore) Disable(id fields.ID) error {
	if err := s.verifyExists(id); err != nil {
		return err
	}
	return s.putOne(id, "disabled", []byte("true"))
}

func (s *consulStore) SetDesiredReplicas(id fields.ID, n int) error {
	if err := s.verifyExists(id); err != nil {
		return err
	}
	return s.putOne(id, "replicas_desired", []byte(strconv.Itoa(n)))
}

func (s *consulStore) Delete(id fields.ID) error {
	rc, err := s.Get(id)
	if err != nil {
		return err
	}

	if rc.ReplicasDesired != 0 {
		return fmt.Errorf("Replication controller %s has %d desired replicas, must be 0 before can be deleted", id, rc.ReplicasDesired)
	}

	_, err = s.kv.DeleteTree(idPrefix(id), nil)
	if err != nil {
		return err
	}

	return err
}

func (s *consulStore) Watch(rc *fields.RC, quit <-chan struct{}) (<-chan struct{}, <-chan error) {
	updated := make(chan struct{})
	errors := make(chan error)
	var curIndex uint64 = 0

	go func() {
		defer close(updated)
		defer close(errors)

		for {
			select {
			case <-quit:
				return
			case <-time.After(1 * time.Second):
				pairs, meta, err := s.kv.List(idPrefix(rc.Id), &api.QueryOptions{
					WaitIndex: curIndex,
				})
				if err != nil {
					errors <- err
				} else {
					curIndex = meta.LastIndex

					rcMap := s.kvpsToRcs(pairs)

					if newRc, ok := rcMap[rc.Id]; ok {
						*rc = *newRc
						updated <- struct{}{}
					}
				}
			}
		}
	}()

	return updated, errors
}

func idPrefix(id fields.ID) string {
	return fmt.Sprintf("%s/%s", kvPrefix, id)
}

func (s *consulStore) putOne(id fields.ID, key string, value []byte) error {
	prefix := idPrefix(id)
	p := &api.KVPair{Key: prefix + "/" + key, Value: value}
	_, err := s.kv.Put(p, nil)
	return err
}

// consulPut attempts to put multiple KV Pairs into consul, returning an error if any fail.
// For now, this performs no CAS or other locking.
// It is assumed that only one process will be writing to an RC at a time.
func (s *consulStore) put(id fields.ID, pairs []kvPair) error {
	for _, pair := range pairs {
		if err := s.putOne(id, pair.key, pair.value); err != nil {
			return err
		}
	}
	return nil
}

func (s *consulStore) verifyExists(id fields.ID) error {
	keys, _, err := s.kv.Keys(idPrefix(id), "", nil)
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return fmt.Errorf("No such replication controller %s", id)
	}
	return nil
}

func (s *consulStore) kvpsToRcs(kvps api.KVPairs) map[fields.ID]*fields.RC {
	rcs := make(map[fields.ID]*fields.RC)

	for _, kvp := range kvps {
		// The key will be of the form replication_controllers/ID/rckey
		// (according to idPrefix)
		parts := strings.SplitN(kvp.Key, "/", 3)
		if len(parts) < 3 {
			s.logger.NoFields().Infof("Ignoring unexpected key %s", kvp.Key)
			continue
		}
		id := fields.ID(parts[1])
		if _, ok := rcs[id]; !ok {
			rcs[id] = &fields.RC{Id: id}
		}

		switch parts[2] {

		case "disabled":
			rcs[id].Disabled = string(kvp.Value) == "true"

		case "node_selector":
			nodeSelector, err := labels.Parse(string(kvp.Value))
			if err == nil {
				rcs[id].NodeSelector = nodeSelector
			} else {
				s.logger.WithError(err).Warnf("%s: Can't unmarshal %s, ignoring", parts[2], kvp.Value)
			}

		case "pod_labels":
			// I don't think there's a way to parse a selector string into a label set, so we have to do it ourselves?!
			labels := make(labels.Set)
			splits := strings.Split(string(kvp.Value), ",")
			for i, split := range splits {
				parts := strings.SplitN(split, "=", 2)
				if len(parts) < 2 {
					s.logger.NoFields().Warnf(
						"%s: Can't unmarshal part %d (%s) out of %d (%s), ignoring",
						parts[2], i, split, len(splits), kvp.Value,
					)
					continue
				}
				labels[parts[0]] = parts[1]
			}
			rcs[id].PodLabels = labels

		case "pod_manifest":
			manifest, err := pods.ManifestFromBytes(kvp.Value)
			if err == nil {
				rcs[id].Manifest = manifest
			} else {
				s.logger.WithError(err).Warnf("%s: Can't unmarshal %s, ignoring", parts[2], kvp.Value)
			}

		case "replicas_desired":
			i, err := strconv.Atoi(string(kvp.Value))
			if err == nil {
				rcs[id].ReplicasDesired = i
			} else {
				s.logger.WithError(err).Warnf("%s: Can't unmarshal %s, ignoring", parts[2], kvp.Value)
			}

		default:
			s.logger.NoFields().Infof("Ignoring unexpcted key %s", kvp.Key)
		}
	}

	return rcs
}
