// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userop/proto/message.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	MessageType          int32    `protobuf:"varint,3,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	File                 string   `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0b918726cc0b203, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MessageRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MessageRequest) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *MessageRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *MessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MessageRequest) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

type MessageResponse struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	MessageType          int32    `protobuf:"varint,3,opt,name=messageType,proto3" json:"messageType,omitempty"`
	Subject              string   `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	File                 string   `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0b918726cc0b203, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MessageResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MessageResponse) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *MessageResponse) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *MessageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MessageResponse) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

type MessageListResponse struct {
	Total                int32              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data                 []*MessageResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MessageListResponse) Reset()         { *m = MessageListResponse{} }
func (m *MessageListResponse) String() string { return proto.CompactTextString(m) }
func (*MessageListResponse) ProtoMessage()    {}
func (*MessageListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0b918726cc0b203, []int{2}
}

func (m *MessageListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageListResponse.Unmarshal(m, b)
}
func (m *MessageListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageListResponse.Marshal(b, m, deterministic)
}
func (m *MessageListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageListResponse.Merge(m, src)
}
func (m *MessageListResponse) XXX_Size() int {
	return xxx_messageInfo_MessageListResponse.Size(m)
}
func (m *MessageListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageListResponse proto.InternalMessageInfo

func (m *MessageListResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *MessageListResponse) GetData() []*MessageResponse {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "MessageResponse")
	proto.RegisterType((*MessageListResponse)(nil), "MessageListResponse")
}

func init() { proto.RegisterFile("userop/proto/message.proto", fileDescriptor_d0b918726cc0b203) }

var fileDescriptor_d0b918726cc0b203 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0x34, 0x7f, 0xd4, 0x8b, 0x68, 0xd1, 0x51, 0x21, 0xab, 0x53, 0x54, 0x31, 0x64,
	0x72, 0xa5, 0xc0, 0xc6, 0x06, 0x13, 0x12, 0x0c, 0x44, 0x4c, 0x6c, 0x2e, 0x39, 0x50, 0x50, 0xc1,
	0x21, 0xe7, 0x0c, 0x7c, 0x1d, 0xf8, 0xa2, 0xa8, 0x8e, 0x5b, 0x15, 0xc2, 0x07, 0xe8, 0x94, 0xbc,
	0xf7, 0xec, 0x77, 0x3f, 0xcb, 0x86, 0x79, 0xc7, 0xd4, 0x9a, 0x66, 0xd9, 0xb4, 0xc6, 0x9a, 0xe5,
	0x1b, 0x31, 0xeb, 0x17, 0x52, 0x4e, 0x2d, 0xbe, 0x04, 0x4c, 0xee, 0x7a, 0xa7, 0xa4, 0x8f, 0x8e,
	0xd8, 0xe2, 0x04, 0x82, 0xba, 0x92, 0x22, 0x13, 0x79, 0x54, 0x06, 0x75, 0x85, 0xa7, 0x10, 0x6f,
	0x0a, 0x6e, 0x2a, 0x19, 0x38, 0xcf, 0x2b, 0xcc, 0x20, 0xf5, 0x5d, 0x0f, 0x9f, 0x0d, 0xc9, 0x91,
	0x0b, 0xf7, 0x2d, 0x94, 0x90, 0x70, 0xb7, 0x7a, 0xa5, 0x27, 0x2b, 0xc3, 0x4c, 0xe4, 0xe3, 0x72,
	0x2b, 0x37, 0x89, 0x5f, 0x28, 0xa3, 0x3e, 0xf1, 0x12, 0x11, 0xc2, 0xe7, 0x7a, 0x4d, 0x32, 0x76,
	0xb6, 0xfb, 0x5f, 0x7c, 0x0b, 0x98, 0xee, 0x20, 0xb9, 0x31, 0xef, 0x4c, 0x07, 0x48, 0x79, 0x0f,
	0x27, 0x1e, 0xf2, 0xb6, 0x66, 0xbb, 0x03, 0x9d, 0x41, 0x64, 0x8d, 0xd5, 0x6b, 0xcf, 0xda, 0x0b,
	0x3c, 0x83, 0xb0, 0xd2, 0x56, 0xcb, 0x20, 0x1b, 0xe5, 0x69, 0x71, 0xac, 0xfe, 0x1c, 0xaf, 0x74,
	0x69, 0xc1, 0x90, 0xf8, 0x00, 0x2f, 0x20, 0xdd, 0x6b, 0xc7, 0xa9, 0xfa, 0x7d, 0x6b, 0xf3, 0x99,
	0xfa, 0x6f, 0x78, 0x01, 0x47, 0xd7, 0x2d, 0x69, 0x4b, 0xdb, 0x9a, 0xc1, 0xbe, 0xc1, 0xe8, 0xab,
	0xf1, 0x63, 0xa2, 0x2e, 0xdd, 0xeb, 0x58, 0xc5, 0xee, 0x73, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x4a, 0x84, 0xd0, 0x4b, 0x42, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageClient interface {
	MessageList(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageListResponse, error)
	CreateMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) MessageList(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageListResponse, error) {
	out := new(MessageListResponse)
	err := c.cc.Invoke(ctx, "/Message/MessageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) CreateMessage(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/Message/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
type MessageServer interface {
	MessageList(context.Context, *MessageRequest) (*MessageListResponse, error)
	CreateMessage(context.Context, *MessageRequest) (*MessageResponse, error)
}

// UnimplementedMessageServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (*UnimplementedMessageServer) MessageList(ctx context.Context, req *MessageRequest) (*MessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageList not implemented")
}
func (*UnimplementedMessageServer) CreateMessage(ctx context.Context, req *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_MessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).MessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Message/MessageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).MessageList(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Message/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateMessage(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageList",
			Handler:    _Message_MessageList_Handler,
		},
		{
			MethodName: "CreateMessage",
			Handler:    _Message_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userop/proto/message.proto",
}
