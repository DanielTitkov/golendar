// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/grpc/pb/golendar.proto

package golendarpb

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

type Event struct {
	EventUUID            string   `protobuf:"bytes,1,opt,name=eventUUID,proto3" json:"eventUUID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Datetime             string   `protobuf:"bytes,3,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Duration             string   `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Desc                 string   `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	User                 string   `protobuf:"bytes,6,opt,name=user,proto3" json:"user,omitempty"`
	Notify               string   `protobuf:"bytes,7,opt,name=notify,proto3" json:"notify,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventUUID() string {
	if m != nil {
		return m.EventUUID
	}
	return ""
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetDatetime() string {
	if m != nil {
		return m.Datetime
	}
	return ""
}

func (m *Event) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Event) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetNotify() string {
	if m != nil {
		return m.Notify
	}
	return ""
}

type GetEventRequest struct {
	EventUUID            []string `protobuf:"bytes,1,rep,name=eventUUID,proto3" json:"eventUUID,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventRequest) Reset()         { *m = GetEventRequest{} }
func (m *GetEventRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventRequest) ProtoMessage()    {}
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{1}
}

func (m *GetEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventRequest.Unmarshal(m, b)
}
func (m *GetEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventRequest.Marshal(b, m, deterministic)
}
func (m *GetEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventRequest.Merge(m, src)
}
func (m *GetEventRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventRequest.Size(m)
}
func (m *GetEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventRequest proto.InternalMessageInfo

func (m *GetEventRequest) GetEventUUID() []string {
	if m != nil {
		return m.EventUUID
	}
	return nil
}

func (m *GetEventRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Event                []*Event `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventResponse) Reset()         { *m = GetEventResponse{} }
func (m *GetEventResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventResponse) ProtoMessage()    {}
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{2}
}

func (m *GetEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventResponse.Unmarshal(m, b)
}
func (m *GetEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventResponse.Marshal(b, m, deterministic)
}
func (m *GetEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventResponse.Merge(m, src)
}
func (m *GetEventResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventResponse.Size(m)
}
func (m *GetEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventResponse proto.InternalMessageInfo

func (m *GetEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetEventResponse) GetEvent() []*Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateEventRequest struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{3}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventResponse) Reset()         { *m = CreateEventResponse{} }
func (m *CreateEventResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEventResponse) ProtoMessage()    {}
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{4}
}

func (m *CreateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventResponse.Unmarshal(m, b)
}
func (m *CreateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventResponse.Marshal(b, m, deterministic)
}
func (m *CreateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventResponse.Merge(m, src)
}
func (m *CreateEventResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEventResponse.Size(m)
}
func (m *CreateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventResponse proto.InternalMessageInfo

func (m *CreateEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreateEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventRequest struct {
	EventUUID            string   `protobuf:"bytes,1,opt,name=eventUUID,proto3" json:"eventUUID,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventRequest) Reset()         { *m = UpdateEventRequest{} }
func (m *UpdateEventRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEventRequest) ProtoMessage()    {}
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{5}
}

func (m *UpdateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventRequest.Unmarshal(m, b)
}
func (m *UpdateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventRequest.Merge(m, src)
}
func (m *UpdateEventRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEventRequest.Size(m)
}
func (m *UpdateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventRequest proto.InternalMessageInfo

func (m *UpdateEventRequest) GetEventUUID() string {
	if m != nil {
		return m.EventUUID
	}
	return ""
}

func (m *UpdateEventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type UpdateEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventResponse) Reset()         { *m = UpdateEventResponse{} }
func (m *UpdateEventResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEventResponse) ProtoMessage()    {}
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{6}
}

func (m *UpdateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventResponse.Unmarshal(m, b)
}
func (m *UpdateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventResponse.Marshal(b, m, deterministic)
}
func (m *UpdateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventResponse.Merge(m, src)
}
func (m *UpdateEventResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateEventResponse.Size(m)
}
func (m *UpdateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventResponse proto.InternalMessageInfo

func (m *UpdateEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UpdateEventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type DeleteEventRequest struct {
	EventUUID            string   `protobuf:"bytes,1,opt,name=eventUUID,proto3" json:"eventUUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{7}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetEventUUID() string {
	if m != nil {
		return m.EventUUID
	}
	return ""
}

type DeleteEventResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventResponse) Reset()         { *m = DeleteEventResponse{} }
func (m *DeleteEventResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEventResponse) ProtoMessage()    {}
func (*DeleteEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08ae5bde3b2add88, []int{8}
}

func (m *DeleteEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventResponse.Unmarshal(m, b)
}
func (m *DeleteEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventResponse.Merge(m, src)
}
func (m *DeleteEventResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEventResponse.Size(m)
}
func (m *DeleteEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventResponse proto.InternalMessageInfo

func (m *DeleteEventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "golendar.Event")
	proto.RegisterType((*GetEventRequest)(nil), "golendar.GetEventRequest")
	proto.RegisterType((*GetEventResponse)(nil), "golendar.GetEventResponse")
	proto.RegisterType((*CreateEventRequest)(nil), "golendar.CreateEventRequest")
	proto.RegisterType((*CreateEventResponse)(nil), "golendar.CreateEventResponse")
	proto.RegisterType((*UpdateEventRequest)(nil), "golendar.UpdateEventRequest")
	proto.RegisterType((*UpdateEventResponse)(nil), "golendar.UpdateEventResponse")
	proto.RegisterType((*DeleteEventRequest)(nil), "golendar.DeleteEventRequest")
	proto.RegisterType((*DeleteEventResponse)(nil), "golendar.DeleteEventResponse")
}

func init() { proto.RegisterFile("api/grpc/pb/golendar.proto", fileDescriptor_08ae5bde3b2add88) }

var fileDescriptor_08ae5bde3b2add88 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6b, 0xea, 0x40,
	0x10, 0xc6, 0x31, 0x1a, 0x9f, 0x8e, 0x82, 0x8f, 0xf1, 0xf1, 0xd8, 0x06, 0x0b, 0x12, 0x28, 0x78,
	0xa9, 0x82, 0x3d, 0xf6, 0xd4, 0x6a, 0x29, 0xf5, 0x56, 0xab, 0x87, 0xf6, 0x16, 0xcd, 0x54, 0x02,
	0x36, 0x49, 0xb3, 0x1b, 0xa1, 0x7f, 0x57, 0xe9, 0xff, 0x57, 0xb2, 0x49, 0xcc, 0xc6, 0xa8, 0x58,
	0xe8, 0x6d, 0x67, 0xbe, 0xcc, 0x6f, 0x3f, 0xbe, 0x59, 0x02, 0x86, 0xe5, 0x3b, 0x83, 0x55, 0xe0,
	0x2f, 0x07, 0xfe, 0x62, 0xb0, 0xf2, 0xd6, 0xe4, 0xda, 0x56, 0xd0, 0xf7, 0x03, 0x4f, 0x78, 0x58,
	0x4b, 0x6b, 0xf3, 0xb3, 0x04, 0xfa, 0xdd, 0x86, 0x5c, 0x81, 0x1d, 0xa8, 0x53, 0x74, 0x98, 0xcf,
	0x1f, 0xc6, 0xac, 0xd4, 0x2d, 0xf5, 0xea, 0xd3, 0xac, 0x81, 0xff, 0x40, 0x17, 0x8e, 0x58, 0x13,
	0xd3, 0xa4, 0x12, 0x17, 0x68, 0x40, 0xcd, 0xb6, 0x04, 0x09, 0xe7, 0x8d, 0x58, 0x59, 0x0a, 0xdb,
	0x5a, 0x6a, 0x61, 0x60, 0x09, 0xc7, 0x73, 0x59, 0x25, 0xd1, 0x92, 0x1a, 0x11, 0x2a, 0x36, 0xf1,
	0x25, 0xd3, 0x65, 0x5f, 0x9e, 0xa3, 0x5e, 0xc8, 0x29, 0x60, 0xd5, 0xb8, 0x17, 0x9d, 0xf1, 0x3f,
	0x54, 0x5d, 0x4f, 0x38, 0xaf, 0x1f, 0xec, 0x8f, 0xec, 0x26, 0x95, 0x39, 0x82, 0xd6, 0x3d, 0x09,
	0xe9, 0x7b, 0x4a, 0xef, 0x21, 0xf1, 0x82, 0xfd, 0x72, 0xde, 0x7e, 0x0a, 0xd7, 0x32, 0xb8, 0xf9,
	0x08, 0x7f, 0x33, 0x08, 0xf7, 0x3d, 0x97, 0x53, 0x74, 0x21, 0x17, 0x96, 0x08, 0x79, 0x92, 0x40,
	0x52, 0xe1, 0x05, 0xe8, 0x12, 0xc6, 0xb4, 0x6e, 0xb9, 0xd7, 0x18, 0xb6, 0xfa, 0xdb, 0x40, 0xe3,
	0xf9, 0x58, 0x35, 0xaf, 0x01, 0x47, 0x01, 0x59, 0x82, 0x72, 0xd6, 0xb6, 0xc3, 0x11, 0xf3, 0xf0,
	0xf0, 0x0c, 0xda, 0xb9, 0xe1, 0xd3, 0x2d, 0x1d, 0xa3, 0x3e, 0x03, 0xce, 0x7d, 0x7b, 0xd7, 0xd2,
	0xf1, 0x65, 0x9f, 0x88, 0x9e, 0x41, 0x3b, 0x87, 0xfe, 0x1d, 0xc3, 0x43, 0xc0, 0x31, 0xad, 0xe9,
	0x27, 0x86, 0xcd, 0x4b, 0x68, 0xe7, 0x66, 0x8e, 0x3b, 0x19, 0x7e, 0x69, 0xd0, 0x94, 0x5f, 0x3e,
	0x51, 0xb0, 0x71, 0x96, 0x84, 0x37, 0x50, 0x4b, 0x9f, 0x02, 0x9e, 0x65, 0xbe, 0x76, 0xde, 0x98,
	0x61, 0xec, 0x93, 0x92, 0xbb, 0x26, 0xd0, 0x50, 0xb6, 0x87, 0x9d, 0xec, 0xd3, 0xe2, 0x8b, 0x30,
	0xce, 0x0f, 0xa8, 0x19, 0x4b, 0x09, 0x56, 0x65, 0x15, 0x57, 0xa9, 0xb2, 0xf6, 0x6d, 0x63, 0x02,
	0x0d, 0x25, 0x1a, 0x95, 0x55, 0x4c, 0x59, 0x65, 0xed, 0xc9, 0xf3, 0xb6, 0xf9, 0x02, 0xa9, 0xee,
	0x2f, 0x16, 0x55, 0xf9, 0x2f, 0xb9, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x69, 0xac, 0x02, 0x9b,
	0x69, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServiceClient interface {
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error)
}

type eventServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventServiceClient(cc *grpc.ClientConn) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/golendar.EventService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/golendar.EventService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	out := new(UpdateEventResponse)
	err := c.cc.Invoke(ctx, "/golendar.EventService/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServiceClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*DeleteEventResponse, error) {
	out := new(DeleteEventResponse)
	err := c.cc.Invoke(ctx, "/golendar.EventService/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
type EventServiceServer interface {
	GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error)
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	UpdateEvent(context.Context, *UpdateEventRequest) (*UpdateEventResponse, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*DeleteEventResponse, error)
}

// UnimplementedEventServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (*UnimplementedEventServiceServer) GetEvent(ctx context.Context, req *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedEventServiceServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedEventServiceServer) UpdateEvent(ctx context.Context, req *UpdateEventRequest) (*UpdateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedEventServiceServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*DeleteEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}

func RegisterEventServiceServer(s *grpc.Server, srv EventServiceServer) {
	s.RegisterService(&_EventService_serviceDesc, srv)
}

func _EventService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golendar.EventService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golendar.EventService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golendar.EventService/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).UpdateEvent(ctx, req.(*UpdateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventService_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/golendar.EventService/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "golendar.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _EventService_GetEvent_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _EventService_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _EventService_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _EventService_DeleteEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/pb/golendar.proto",
}