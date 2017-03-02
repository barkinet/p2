// Code generated by protoc-gen-go.
// source: pkg/grpc/podstore/protos/podstore.proto
// DO NOT EDIT!

/*
Package podstore is a generated protocol buffer package.

It is generated from these files:
	pkg/grpc/podstore/protos/podstore.proto

It has these top-level messages:
	SchedulePodRequest
	SchedulePodResponse
	WatchPodStatusRequest
	PodStatusResponse
	ProcessStatus
	ExitStatus
	UnschedulePodRequest
	UnschedulePodResponse
*/
package podstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SchedulePodRequest struct {
	Manifest string `protobuf:"bytes,1,opt,name=manifest" json:"manifest,omitempty"`
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *SchedulePodRequest) Reset()                    { *m = SchedulePodRequest{} }
func (m *SchedulePodRequest) String() string            { return proto.CompactTextString(m) }
func (*SchedulePodRequest) ProtoMessage()               {}
func (*SchedulePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SchedulePodRequest) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *SchedulePodRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type SchedulePodResponse struct {
	PodUniqueKey string `protobuf:"bytes,1,opt,name=pod_unique_key,json=podUniqueKey" json:"pod_unique_key,omitempty"`
}

func (m *SchedulePodResponse) Reset()                    { *m = SchedulePodResponse{} }
func (m *SchedulePodResponse) String() string            { return proto.CompactTextString(m) }
func (*SchedulePodResponse) ProtoMessage()               {}
func (*SchedulePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SchedulePodResponse) GetPodUniqueKey() string {
	if m != nil {
		return m.PodUniqueKey
	}
	return ""
}

type WatchPodStatusRequest struct {
	PodUniqueKey    string `protobuf:"bytes,1,opt,name=pod_unique_key,json=podUniqueKey" json:"pod_unique_key,omitempty"`
	WaitIndex       uint64 `protobuf:"varint,2,opt,name=wait_index,json=waitIndex" json:"wait_index,omitempty"`
	StatusNamespace string `protobuf:"bytes,3,opt,name=status_namespace,json=statusNamespace" json:"status_namespace,omitempty"`
	WaitForExists   bool   `protobuf:"varint,4,opt,name=wait_for_exists,json=waitForExists" json:"wait_for_exists,omitempty"`
}

func (m *WatchPodStatusRequest) Reset()                    { *m = WatchPodStatusRequest{} }
func (m *WatchPodStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchPodStatusRequest) ProtoMessage()               {}
func (*WatchPodStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WatchPodStatusRequest) GetPodUniqueKey() string {
	if m != nil {
		return m.PodUniqueKey
	}
	return ""
}

func (m *WatchPodStatusRequest) GetWaitIndex() uint64 {
	if m != nil {
		return m.WaitIndex
	}
	return 0
}

func (m *WatchPodStatusRequest) GetStatusNamespace() string {
	if m != nil {
		return m.StatusNamespace
	}
	return ""
}

func (m *WatchPodStatusRequest) GetWaitForExists() bool {
	if m != nil {
		return m.WaitForExists
	}
	return false
}

type PodStatusResponse struct {
	Manifest        string           `protobuf:"bytes,1,opt,name=manifest" json:"manifest,omitempty"`
	PodState        string           `protobuf:"bytes,2,opt,name=pod_state,json=podState" json:"pod_state,omitempty"`
	ProcessStatuses []*ProcessStatus `protobuf:"bytes,3,rep,name=process_statuses,json=processStatuses" json:"process_statuses,omitempty"`
	LastIndex       uint64           `protobuf:"varint,4,opt,name=last_index,json=lastIndex" json:"last_index,omitempty"`
}

func (m *PodStatusResponse) Reset()                    { *m = PodStatusResponse{} }
func (m *PodStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*PodStatusResponse) ProtoMessage()               {}
func (*PodStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PodStatusResponse) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *PodStatusResponse) GetPodState() string {
	if m != nil {
		return m.PodState
	}
	return ""
}

func (m *PodStatusResponse) GetProcessStatuses() []*ProcessStatus {
	if m != nil {
		return m.ProcessStatuses
	}
	return nil
}

func (m *PodStatusResponse) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

type ProcessStatus struct {
	LaunchableId string      `protobuf:"bytes,1,opt,name=launchable_id,json=launchableId" json:"launchable_id,omitempty"`
	EntryPoint   string      `protobuf:"bytes,2,opt,name=entry_point,json=entryPoint" json:"entry_point,omitempty"`
	LastExit     *ExitStatus `protobuf:"bytes,3,opt,name=last_exit,json=lastExit" json:"last_exit,omitempty"`
}

func (m *ProcessStatus) Reset()                    { *m = ProcessStatus{} }
func (m *ProcessStatus) String() string            { return proto.CompactTextString(m) }
func (*ProcessStatus) ProtoMessage()               {}
func (*ProcessStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ProcessStatus) GetLaunchableId() string {
	if m != nil {
		return m.LaunchableId
	}
	return ""
}

func (m *ProcessStatus) GetEntryPoint() string {
	if m != nil {
		return m.EntryPoint
	}
	return ""
}

func (m *ProcessStatus) GetLastExit() *ExitStatus {
	if m != nil {
		return m.LastExit
	}
	return nil
}

type ExitStatus struct {
	ExitTime   int64 `protobuf:"varint,1,opt,name=exit_time,json=exitTime" json:"exit_time,omitempty"`
	ExitCode   int64 `protobuf:"varint,2,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
	ExitStatus int64 `protobuf:"varint,3,opt,name=exit_status,json=exitStatus" json:"exit_status,omitempty"`
}

func (m *ExitStatus) Reset()                    { *m = ExitStatus{} }
func (m *ExitStatus) String() string            { return proto.CompactTextString(m) }
func (*ExitStatus) ProtoMessage()               {}
func (*ExitStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExitStatus) GetExitTime() int64 {
	if m != nil {
		return m.ExitTime
	}
	return 0
}

func (m *ExitStatus) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *ExitStatus) GetExitStatus() int64 {
	if m != nil {
		return m.ExitStatus
	}
	return 0
}

type UnschedulePodRequest struct {
	PodUniqueKey string `protobuf:"bytes,1,opt,name=pod_unique_key,json=podUniqueKey" json:"pod_unique_key,omitempty"`
}

func (m *UnschedulePodRequest) Reset()                    { *m = UnschedulePodRequest{} }
func (m *UnschedulePodRequest) String() string            { return proto.CompactTextString(m) }
func (*UnschedulePodRequest) ProtoMessage()               {}
func (*UnschedulePodRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UnschedulePodRequest) GetPodUniqueKey() string {
	if m != nil {
		return m.PodUniqueKey
	}
	return ""
}

type UnschedulePodResponse struct {
}

func (m *UnschedulePodResponse) Reset()                    { *m = UnschedulePodResponse{} }
func (m *UnschedulePodResponse) String() string            { return proto.CompactTextString(m) }
func (*UnschedulePodResponse) ProtoMessage()               {}
func (*UnschedulePodResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*SchedulePodRequest)(nil), "podstore.SchedulePodRequest")
	proto.RegisterType((*SchedulePodResponse)(nil), "podstore.SchedulePodResponse")
	proto.RegisterType((*WatchPodStatusRequest)(nil), "podstore.WatchPodStatusRequest")
	proto.RegisterType((*PodStatusResponse)(nil), "podstore.PodStatusResponse")
	proto.RegisterType((*ProcessStatus)(nil), "podstore.ProcessStatus")
	proto.RegisterType((*ExitStatus)(nil), "podstore.ExitStatus")
	proto.RegisterType((*UnschedulePodRequest)(nil), "podstore.UnschedulePodRequest")
	proto.RegisterType((*UnschedulePodResponse)(nil), "podstore.UnschedulePodResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for P2PodStore service

type P2PodStoreClient interface {
	// Schedules a uuid pod on a host
	SchedulePod(ctx context.Context, in *SchedulePodRequest, opts ...grpc.CallOption) (*SchedulePodResponse, error)
	WatchPodStatus(ctx context.Context, in *WatchPodStatusRequest, opts ...grpc.CallOption) (P2PodStore_WatchPodStatusClient, error)
	UnschedulePod(ctx context.Context, in *UnschedulePodRequest, opts ...grpc.CallOption) (*UnschedulePodResponse, error)
}

type p2PodStoreClient struct {
	cc *grpc.ClientConn
}

func NewP2PodStoreClient(cc *grpc.ClientConn) P2PodStoreClient {
	return &p2PodStoreClient{cc}
}

func (c *p2PodStoreClient) SchedulePod(ctx context.Context, in *SchedulePodRequest, opts ...grpc.CallOption) (*SchedulePodResponse, error) {
	out := new(SchedulePodResponse)
	err := grpc.Invoke(ctx, "/podstore.P2PodStore/SchedulePod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *p2PodStoreClient) WatchPodStatus(ctx context.Context, in *WatchPodStatusRequest, opts ...grpc.CallOption) (P2PodStore_WatchPodStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_P2PodStore_serviceDesc.Streams[0], c.cc, "/podstore.P2PodStore/WatchPodStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PodStoreWatchPodStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P2PodStore_WatchPodStatusClient interface {
	Recv() (*PodStatusResponse, error)
	grpc.ClientStream
}

type p2PodStoreWatchPodStatusClient struct {
	grpc.ClientStream
}

func (x *p2PodStoreWatchPodStatusClient) Recv() (*PodStatusResponse, error) {
	m := new(PodStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *p2PodStoreClient) UnschedulePod(ctx context.Context, in *UnschedulePodRequest, opts ...grpc.CallOption) (*UnschedulePodResponse, error) {
	out := new(UnschedulePodResponse)
	err := grpc.Invoke(ctx, "/podstore.P2PodStore/UnschedulePod", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for P2PodStore service

type P2PodStoreServer interface {
	// Schedules a uuid pod on a host
	SchedulePod(context.Context, *SchedulePodRequest) (*SchedulePodResponse, error)
	WatchPodStatus(*WatchPodStatusRequest, P2PodStore_WatchPodStatusServer) error
	UnschedulePod(context.Context, *UnschedulePodRequest) (*UnschedulePodResponse, error)
}

func RegisterP2PodStoreServer(s *grpc.Server, srv P2PodStoreServer) {
	s.RegisterService(&_P2PodStore_serviceDesc, srv)
}

func _P2PodStore_SchedulePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PodStoreServer).SchedulePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podstore.P2PodStore/SchedulePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PodStoreServer).SchedulePod(ctx, req.(*SchedulePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _P2PodStore_WatchPodStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPodStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P2PodStoreServer).WatchPodStatus(m, &p2PodStoreWatchPodStatusServer{stream})
}

type P2PodStore_WatchPodStatusServer interface {
	Send(*PodStatusResponse) error
	grpc.ServerStream
}

type p2PodStoreWatchPodStatusServer struct {
	grpc.ServerStream
}

func (x *p2PodStoreWatchPodStatusServer) Send(m *PodStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _P2PodStore_UnschedulePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnschedulePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(P2PodStoreServer).UnschedulePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/podstore.P2PodStore/UnschedulePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(P2PodStoreServer).UnschedulePod(ctx, req.(*UnschedulePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _P2PodStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "podstore.P2PodStore",
	HandlerType: (*P2PodStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SchedulePod",
			Handler:    _P2PodStore_SchedulePod_Handler,
		},
		{
			MethodName: "UnschedulePod",
			Handler:    _P2PodStore_UnschedulePod_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPodStatus",
			Handler:       _P2PodStore_WatchPodStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/podstore/protos/podstore.proto",
}

func init() { proto.RegisterFile("pkg/grpc/podstore/protos/podstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xed, 0x36, 0x15, 0x4a, 0x26, 0xa4, 0x29, 0xa6, 0x55, 0xa3, 0x94, 0x92, 0x68, 0x41, 0x10,
	0x2e, 0x0d, 0x84, 0x23, 0x9c, 0x40, 0x45, 0xaa, 0x80, 0x2a, 0xda, 0x52, 0x71, 0x5c, 0xb9, 0xeb,
	0x69, 0x62, 0x35, 0xb1, 0xdd, 0xb5, 0x57, 0x24, 0x57, 0x4e, 0xfc, 0x1b, 0xfe, 0x22, 0xb2, 0xbd,
	0x1f, 0x49, 0x3f, 0x44, 0x8f, 0xf3, 0x66, 0xe7, 0xcd, 0x9b, 0xe7, 0x99, 0x85, 0xd7, 0xea, 0x6a,
	0x32, 0x9c, 0xa4, 0x2a, 0x19, 0x2a, 0xc9, 0xb4, 0x91, 0x29, 0x0e, 0x55, 0x2a, 0x8d, 0xd4, 0x65,
	0x7c, 0xe4, 0x62, 0x52, 0x2f, 0xe2, 0xf0, 0x3b, 0x90, 0xb3, 0x64, 0x8a, 0x2c, 0x9b, 0xe1, 0x58,
	0xb2, 0x08, 0xaf, 0x33, 0xd4, 0x86, 0x74, 0xa1, 0x3e, 0xa7, 0x82, 0x5f, 0xa2, 0x36, 0x9d, 0xa0,
	0x1f, 0x0c, 0x1a, 0x51, 0x19, 0x93, 0x03, 0x68, 0x08, 0xc9, 0x30, 0x16, 0x74, 0x8e, 0x9d, 0x4d,
	0x9f, 0xb4, 0xc0, 0x29, 0x9d, 0x63, 0xf8, 0x01, 0x9e, 0xae, 0xd1, 0x69, 0x25, 0x85, 0x46, 0xf2,
	0x12, 0xb6, 0x95, 0x64, 0x71, 0x26, 0xf8, 0x75, 0x86, 0xf1, 0x15, 0x2e, 0x73, 0xd6, 0xc7, 0x4a,
	0xb2, 0x73, 0x07, 0x7e, 0xc5, 0x65, 0xf8, 0x37, 0x80, 0xbd, 0x9f, 0xd4, 0x24, 0xd3, 0xb1, 0x64,
	0x67, 0x86, 0x9a, 0x4c, 0x17, 0x7a, 0x1e, 0x54, 0x4f, 0x0e, 0x01, 0x7e, 0x51, 0x6e, 0x62, 0x2e,
	0x18, 0x2e, 0x9c, 0xb4, 0xad, 0xa8, 0x61, 0x91, 0x13, 0x0b, 0x90, 0x37, 0xb0, 0xa3, 0x1d, 0xab,
	0x93, 0xae, 0x15, 0x4d, 0xb0, 0x53, 0x73, 0x34, 0x6d, 0x8f, 0x9f, 0x16, 0x30, 0x79, 0x05, 0x6d,
	0xc7, 0x74, 0x29, 0xd3, 0x18, 0x17, 0x5c, 0x1b, 0xdd, 0xd9, 0xea, 0x07, 0x83, 0x7a, 0xd4, 0xb2,
	0xf0, 0x17, 0x99, 0x1e, 0x3b, 0xd0, 0x2a, 0x7e, 0xb2, 0x22, 0x36, 0x9f, 0xf6, 0x3f, 0xee, 0xd9,
	0x49, 0x6c, 0xc3, 0xd2, 0x3d, 0xe5, 0x19, 0x90, 0x7c, 0x82, 0x1d, 0x95, 0xca, 0x04, 0xb5, 0x8e,
	0xbd, 0x22, 0xd4, 0x9d, 0x5a, 0xbf, 0x36, 0x68, 0x8e, 0xf6, 0x8f, 0xca, 0x17, 0x1c, 0xfb, 0x2f,
	0xf2, 0x9e, 0x6d, 0xb5, 0x1a, 0xa2, 0xb6, 0x26, 0xcc, 0xa8, 0x2e, 0x4c, 0xd8, 0xf2, 0x26, 0x58,
	0xc4, 0x99, 0x10, 0xfe, 0x09, 0xa0, 0xb5, 0xc6, 0x40, 0x5e, 0x40, 0x6b, 0x46, 0x33, 0x91, 0x4c,
	0xe9, 0xc5, 0x0c, 0x63, 0xce, 0x0a, 0x6b, 0x2b, 0xf0, 0x84, 0x91, 0x1e, 0x34, 0x51, 0x98, 0x74,
	0x19, 0x2b, 0xc9, 0x85, 0xc9, 0x85, 0x83, 0x83, 0xc6, 0x16, 0x21, 0xef, 0xc0, 0x35, 0xb1, 0x6e,
	0x19, 0xe7, 0x6a, 0x73, 0xb4, 0x5b, 0x69, 0x3e, 0x5e, 0x70, 0x93, 0x0b, 0xae, 0xdb, 0xcf, 0x6c,
	0x1c, 0x4e, 0x00, 0x2a, 0xdc, 0x1a, 0x63, 0x6b, 0x63, 0xc3, 0xe7, 0xe8, 0x24, 0xd4, 0xa2, 0xba,
	0x05, 0x7e, 0xf0, 0x39, 0x96, 0xc9, 0x44, 0x32, 0xef, 0x5a, 0x9e, 0xfc, 0x2c, 0x19, 0x3a, 0x6d,
	0x36, 0xe9, 0x2d, 0x73, 0xcd, 0x6b, 0x11, 0x60, 0x49, 0x1d, 0x7e, 0x84, 0xdd, 0x73, 0xa1, 0x6f,
	0x6f, 0xf9, 0xc3, 0xb6, 0x72, 0x1f, 0xf6, 0x6e, 0x54, 0xfb, 0x67, 0x1e, 0xfd, 0xde, 0x04, 0x18,
	0x8f, 0xdc, 0xf3, 0xcb, 0x14, 0xc9, 0x37, 0x68, 0xae, 0xac, 0x3e, 0x79, 0x56, 0x4d, 0x7f, 0xfb,
	0xc0, 0xba, 0x87, 0xf7, 0x64, 0x3d, 0x75, 0xb8, 0x41, 0x22, 0xd8, 0x5e, 0x3f, 0x05, 0xd2, 0xab,
	0x4a, 0xee, 0x3c, 0x92, 0xee, 0xc1, 0xca, 0x8e, 0xdc, 0xdc, 0xc9, 0x70, 0xe3, 0x6d, 0x40, 0x22,
	0x68, 0xad, 0x4d, 0x42, 0x9e, 0x57, 0x15, 0x77, 0x19, 0xd4, 0xed, 0xdd, 0x9b, 0x2f, 0x58, 0x2f,
	0x1e, 0xb9, 0x1f, 0xca, 0xfb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x97, 0xaf, 0x7d, 0x14, 0x7b,
	0x04, 0x00, 0x00,
}