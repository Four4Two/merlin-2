// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/downtime-detector/v1beta1/query.proto

package queryproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types1 "github.com/osmosis-labs/osmosis/v13/x/downtime-detector/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Query for has it been at least $RECOVERY_DURATION units of time,
// since the chain has been down for $DOWNTIME_DURATION.
// Note: $DOWNTIME_DURATION must be in set {SPECIFY_SET}
type RecoveredSinceDowntimeOfLengthRequest struct {
	Downtime time.Duration `protobuf:"bytes,1,opt,name=downtime,proto3,stdduration" json:"downtime" yaml:"downtime_duration"`
	Recovery time.Duration `protobuf:"bytes,2,opt,name=recovery,proto3,stdduration" json:"recovery" yaml:"recovery_duration"`
}

func (m *RecoveredSinceDowntimeOfLengthRequest) Reset()         { *m = RecoveredSinceDowntimeOfLengthRequest{} }
func (m *RecoveredSinceDowntimeOfLengthRequest) String() string { return proto.CompactTextString(m) }
func (*RecoveredSinceDowntimeOfLengthRequest) ProtoMessage()    {}
func (*RecoveredSinceDowntimeOfLengthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b748b3d07fa8b8cb, []int{0}
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest.Merge(m, src)
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Size() int {
	return m.Size()
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest proto.InternalMessageInfo

func (m *RecoveredSinceDowntimeOfLengthRequest) GetDowntime() time.Duration {
	if m != nil {
		return m.Downtime
	}
	return 0
}

func (m *RecoveredSinceDowntimeOfLengthRequest) GetRecovery() time.Duration {
	if m != nil {
		return m.Recovery
	}
	return 0
}

type RecoveredSinceDowntimeOfLengthResponse struct {
	SuccesfullyRecovered bool `protobuf:"varint,1,opt,name=succesfully_recovered,json=succesfullyRecovered,proto3" json:"succesfully_recovered,omitempty"`
}

func (m *RecoveredSinceDowntimeOfLengthResponse) Reset() {
	*m = RecoveredSinceDowntimeOfLengthResponse{}
}
func (m *RecoveredSinceDowntimeOfLengthResponse) String() string { return proto.CompactTextString(m) }
func (*RecoveredSinceDowntimeOfLengthResponse) ProtoMessage()    {}
func (*RecoveredSinceDowntimeOfLengthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b748b3d07fa8b8cb, []int{1}
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse.Merge(m, src)
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Size() int {
	return m.Size()
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse proto.InternalMessageInfo

func (m *RecoveredSinceDowntimeOfLengthResponse) GetSuccesfullyRecovered() bool {
	if m != nil {
		return m.SuccesfullyRecovered
	}
	return false
}

type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b748b3d07fa8b8cb, []int{2}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

type ParamsResponse struct {
	Params types1.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b748b3d07fa8b8cb, []int{3}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

func (m *ParamsResponse) GetParams() types1.Params {
	if m != nil {
		return m.Params
	}
	return types1.Params{}
}

func init() {
	proto.RegisterType((*RecoveredSinceDowntimeOfLengthRequest)(nil), "osmosis.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthRequest")
	proto.RegisterType((*RecoveredSinceDowntimeOfLengthResponse)(nil), "osmosis.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthResponse")
	proto.RegisterType((*ParamsRequest)(nil), "osmosis.downtimedetector.v1beta1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "osmosis.downtimedetector.v1beta1.ParamsResponse")
}

func init() {
	proto.RegisterFile("osmosis/downtime-detector/v1beta1/query.proto", fileDescriptor_b748b3d07fa8b8cb)
}

var fileDescriptor_b748b3d07fa8b8cb = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x8b, 0xd3, 0x4c,
	0x18, 0x6f, 0xf6, 0x7d, 0x2d, 0xcb, 0x88, 0x0a, 0x61, 0x85, 0x6e, 0x91, 0x74, 0x09, 0x2a, 0xab,
	0xd2, 0x8c, 0xdd, 0xde, 0xbc, 0x59, 0x17, 0x75, 0x41, 0x50, 0xe3, 0x45, 0x14, 0x29, 0x93, 0xe9,
	0xd3, 0x6c, 0x20, 0x99, 0xc9, 0x66, 0x26, 0xd5, 0x5c, 0xfd, 0x04, 0x82, 0x17, 0x4f, 0xde, 0xfd,
	0x26, 0x7b, 0x5c, 0xf0, 0xe2, 0x69, 0xd5, 0xd6, 0x4f, 0xe0, 0x07, 0x10, 0x49, 0x66, 0x26, 0x94,
	0x22, 0xdb, 0x88, 0xa7, 0x36, 0xf3, 0xfc, 0xfe, 0x3e, 0x4c, 0x82, 0xfa, 0x5c, 0x24, 0x5c, 0x44,
	0x02, 0x4f, 0xf8, 0x6b, 0x26, 0xa3, 0x04, 0xfa, 0x13, 0x90, 0x40, 0x25, 0xcf, 0xf0, 0x6c, 0x10,
	0x80, 0x24, 0x03, 0x7c, 0x94, 0x43, 0x56, 0x78, 0x69, 0xc6, 0x25, 0xb7, 0x77, 0x34, 0xdc, 0x33,
	0x70, 0x83, 0xf6, 0x34, 0xba, 0xbb, 0x15, 0xf2, 0x90, 0x57, 0x60, 0x5c, 0xfe, 0x53, 0xbc, 0x2e,
	0x5e, 0x6f, 0x13, 0x02, 0x83, 0x52, 0x59, 0x11, 0x1c, 0x5a, 0x31, 0x70, 0x40, 0x04, 0xd4, 0x10,
	0xca, 0x23, 0xa6, 0xe7, 0x37, 0x97, 0xe7, 0x55, 0xc2, 0x1a, 0x95, 0x92, 0x30, 0x62, 0x44, 0x46,
	0xdc, 0x60, 0xaf, 0x84, 0x9c, 0x87, 0x31, 0x60, 0x92, 0x46, 0x98, 0x30, 0xc6, 0x65, 0x35, 0x34,
	0x4e, 0xdb, 0x7a, 0x5a, 0x3d, 0x05, 0xf9, 0x14, 0x13, 0x56, 0x98, 0x91, 0x32, 0x19, 0xab, 0x3a,
	0xea, 0xc1, 0xe4, 0x5b, 0x65, 0x4d, 0xf2, 0x6c, 0xd9, 0xb3, 0xb7, 0x3a, 0x2f, 0x4b, 0x0b, 0x49,
	0x92, 0x54, 0x01, 0xdc, 0xef, 0x16, 0xba, 0xe6, 0x03, 0xe5, 0x33, 0xc8, 0x60, 0xf2, 0x2c, 0x62,
	0x14, 0xf6, 0xf5, 0x6a, 0x1e, 0x4f, 0x1f, 0x01, 0x0b, 0xe5, 0xa1, 0x0f, 0x47, 0x39, 0x08, 0x69,
	0xbf, 0x44, 0x9b, 0x66, 0x6b, 0x1d, 0x6b, 0xc7, 0xda, 0x3d, 0xbf, 0xb7, 0xed, 0x29, 0x75, 0xcf,
	0xa8, 0x7b, 0xfb, 0xda, 0x7d, 0x74, 0xf5, 0xf8, 0xb4, 0xd7, 0xfa, 0x79, 0xda, 0xeb, 0x14, 0x24,
	0x89, 0xef, 0xb8, 0x86, 0x38, 0x36, 0xf1, 0xdc, 0x0f, 0x5f, 0x7b, 0x96, 0x5f, 0x0b, 0x96, 0xe2,
	0x99, 0x4a, 0x51, 0x74, 0x36, 0xfe, 0x52, 0xdc, 0x10, 0x57, 0xc5, 0xcd, 0xb9, 0xfb, 0x0a, 0x5d,
	0x5f, 0x57, 0x51, 0xa4, 0x9c, 0x09, 0xb0, 0x87, 0xe8, 0xb2, 0xc8, 0x29, 0x05, 0x31, 0xcd, 0xe3,
	0xb8, 0x18, 0x67, 0x86, 0x55, 0x15, 0xde, 0xf4, 0xb7, 0x96, 0x86, 0xb5, 0xa2, 0x7b, 0x09, 0x5d,
	0x78, 0x42, 0x32, 0x92, 0x08, 0xbd, 0x29, 0xf7, 0x39, 0xba, 0x68, 0x0e, 0xb4, 0xee, 0x7d, 0xd4,
	0x4e, 0xab, 0x13, 0xbd, 0xb9, 0x5d, 0x6f, 0xdd, 0x05, 0xf6, 0x94, 0xc2, 0xe8, 0xff, 0xb2, 0xab,
	0xaf, 0xd9, 0x7b, 0x9f, 0xfe, 0x43, 0xe7, 0x9e, 0x96, 0xb7, 0xcc, 0xfe, 0x68, 0xa1, 0xb6, 0x82,
	0xd8, 0xb8, 0xa9, 0x98, 0xce, 0xd7, 0xbd, 0xdd, 0x9c, 0xa0, 0xf2, 0xbb, 0x83, 0xb7, 0x9f, 0x7f,
	0xbc, 0xdf, 0xb8, 0x65, 0xdf, 0x68, 0xf0, 0x02, 0xe9, 0x54, 0xbf, 0x2c, 0xe4, 0x9c, 0xbd, 0x75,
	0xfb, 0xc1, 0xfa, 0x1c, 0x8d, 0xae, 0x66, 0xf7, 0xe1, 0xbf, 0x0b, 0xe9, 0xa2, 0x07, 0x55, 0xd1,
	0x7b, 0xf6, 0xdd, 0x06, 0x45, 0xcf, 0x96, 0x1c, 0xd1, 0xe3, 0xb9, 0x63, 0x9d, 0xcc, 0x1d, 0xeb,
	0xdb, 0xdc, 0xb1, 0xde, 0x2d, 0x9c, 0xd6, 0xc9, 0xc2, 0x69, 0x7d, 0x59, 0x38, 0xad, 0x17, 0x07,
	0x61, 0x24, 0x0f, 0xf3, 0xc0, 0xa3, 0x3c, 0x31, 0x36, 0xfd, 0x98, 0x04, 0xa2, 0xf6, 0x9c, 0x0d,
	0x86, 0xf8, 0xcd, 0x1f, 0x9c, 0x69, 0x1c, 0x01, 0x93, 0xea, 0x3b, 0xa3, 0x5e, 0x8b, 0x76, 0xf5,
	0x33, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x09, 0x15, 0x2a, 0x40, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
	RecoveredSinceDowntimeOfLength(ctx context.Context, in *RecoveredSinceDowntimeOfLengthRequest, opts ...grpc.CallOption) (*RecoveredSinceDowntimeOfLengthResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.downtimedetector.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RecoveredSinceDowntimeOfLength(ctx context.Context, in *RecoveredSinceDowntimeOfLengthRequest, opts ...grpc.CallOption) (*RecoveredSinceDowntimeOfLengthResponse, error) {
	out := new(RecoveredSinceDowntimeOfLengthResponse)
	err := c.cc.Invoke(ctx, "/osmosis.downtimedetector.v1beta1.Query/RecoveredSinceDowntimeOfLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
	RecoveredSinceDowntimeOfLength(context.Context, *RecoveredSinceDowntimeOfLengthRequest) (*RecoveredSinceDowntimeOfLengthResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) RecoveredSinceDowntimeOfLength(ctx context.Context, req *RecoveredSinceDowntimeOfLengthRequest) (*RecoveredSinceDowntimeOfLengthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoveredSinceDowntimeOfLength not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.downtimedetector.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RecoveredSinceDowntimeOfLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoveredSinceDowntimeOfLengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RecoveredSinceDowntimeOfLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.downtimedetector.v1beta1.Query/RecoveredSinceDowntimeOfLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RecoveredSinceDowntimeOfLength(ctx, req.(*RecoveredSinceDowntimeOfLengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.downtimedetector.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "RecoveredSinceDowntimeOfLength",
			Handler:    _Query_RecoveredSinceDowntimeOfLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/downtime-detector/v1beta1/query.proto",
}

func (m *RecoveredSinceDowntimeOfLengthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveredSinceDowntimeOfLengthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveredSinceDowntimeOfLengthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Recovery, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Recovery):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintQuery(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Downtime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Downtime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintQuery(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RecoveredSinceDowntimeOfLengthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveredSinceDowntimeOfLengthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveredSinceDowntimeOfLengthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SuccesfullyRecovered {
		i--
		if m.SuccesfullyRecovered {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RecoveredSinceDowntimeOfLengthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Downtime)
	n += 1 + l + sovQuery(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Recovery)
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *RecoveredSinceDowntimeOfLengthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SuccesfullyRecovered {
		n += 2
	}
	return n
}

func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RecoveredSinceDowntimeOfLengthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downtime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Downtime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recovery", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Recovery, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RecoveredSinceDowntimeOfLengthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccesfullyRecovered", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SuccesfullyRecovered = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)