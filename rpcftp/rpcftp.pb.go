// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpcftp.proto

package rpcftp

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rpcftp_b9d9c67dd436e415, []int{0}
}

// File transfer message
type Chunk struct {
	FileName             string   `protobuf:"bytes,1,opt,name=FileName,proto3" json:"FileName,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcftp_b9d9c67dd436e415, []int{0}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (dst *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(dst, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadStatus struct {
	Message              string           `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,2,opt,name=code,proto3,enum=rpcftp.UploadStatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcftp_b9d9c67dd436e415, []int{1}
}
func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadStatus.Unmarshal(m, b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
}
func (dst *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(dst, src)
}
func (m *UploadStatus) XXX_Size() int {
	return xxx_messageInfo_UploadStatus.Size(m)
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UploadStatus) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

type DownloadRequest struct {
	Path                 string   `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadRequest) Reset()         { *m = DownloadRequest{} }
func (m *DownloadRequest) String() string { return proto.CompactTextString(m) }
func (*DownloadRequest) ProtoMessage()    {}
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_rpcftp_b9d9c67dd436e415, []int{2}
}
func (m *DownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadRequest.Unmarshal(m, b)
}
func (m *DownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadRequest.Marshal(b, m, deterministic)
}
func (dst *DownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadRequest.Merge(dst, src)
}
func (m *DownloadRequest) XXX_Size() int {
	return xxx_messageInfo_DownloadRequest.Size(m)
}
func (m *DownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadRequest proto.InternalMessageInfo

func (m *DownloadRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Chunk)(nil), "rpcftp.Chunk")
	proto.RegisterType((*UploadStatus)(nil), "rpcftp.UploadStatus")
	proto.RegisterType((*DownloadRequest)(nil), "rpcftp.DownloadRequest")
	proto.RegisterEnum("rpcftp.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcftpClient is the client API for Rpcftp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcftpClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Rpcftp_UploadClient, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Rpcftp_DownloadClient, error)
}

type rpcftpClient struct {
	cc *grpc.ClientConn
}

func NewRpcftpClient(cc *grpc.ClientConn) RpcftpClient {
	return &rpcftpClient{cc}
}

func (c *rpcftpClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Rpcftp_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Rpcftp_serviceDesc.Streams[0], "/rpcftp.Rpcftp/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcftpUploadClient{stream}
	return x, nil
}

type Rpcftp_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type rpcftpUploadClient struct {
	grpc.ClientStream
}

func (x *rpcftpUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rpcftpUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcftpClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Rpcftp_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Rpcftp_serviceDesc.Streams[1], "/rpcftp.Rpcftp/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcftpDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rpcftp_DownloadClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type rpcftpDownloadClient struct {
	grpc.ClientStream
}

func (x *rpcftpDownloadClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcftpServer is the server API for Rpcftp service.
type RpcftpServer interface {
	Upload(Rpcftp_UploadServer) error
	Download(*DownloadRequest, Rpcftp_DownloadServer) error
}

func RegisterRpcftpServer(s *grpc.Server, srv RpcftpServer) {
	s.RegisterService(&_Rpcftp_serviceDesc, srv)
}

func _Rpcftp_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RpcftpServer).Upload(&rpcftpUploadServer{stream})
}

type Rpcftp_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type rpcftpUploadServer struct {
	grpc.ServerStream
}

func (x *rpcftpUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rpcftpUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Rpcftp_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcftpServer).Download(m, &rpcftpDownloadServer{stream})
}

type Rpcftp_DownloadServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type rpcftpDownloadServer struct {
	grpc.ServerStream
}

func (x *rpcftpDownloadServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Rpcftp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcftp.Rpcftp",
	HandlerType: (*RpcftpServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Rpcftp_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _Rpcftp_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpcftp.proto",
}

func init() { proto.RegisterFile("rpcftp.proto", fileDescriptor_rpcftp_b9d9c67dd436e415) }

var fileDescriptor_rpcftp_b9d9c67dd436e415 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x9b, 0x32, 0xb3, 0xf9, 0xac, 0x5a, 0x1e, 0x82, 0xa5, 0xa7, 0x11, 0x10, 0x8a, 0xc8,
	0xd0, 0x0d, 0xbc, 0x79, 0xaa, 0xec, 0xe6, 0x0f, 0x22, 0xf3, 0x1e, 0xd7, 0xa7, 0x1b, 0xad, 0x49,
	0x5d, 0x52, 0xf6, 0xef, 0xcb, 0xb2, 0x45, 0x74, 0x78, 0xcb, 0x97, 0x7c, 0xdf, 0xe7, 0x7d, 0x78,
	0x90, 0xac, 0xda, 0xf9, 0xbb, 0x6b, 0x47, 0xed, 0xca, 0x38, 0x83, 0x7c, 0x9b, 0xc4, 0x1d, 0x1c,
	0x94, 0x8b, 0x4e, 0xd7, 0x98, 0xc3, 0x60, 0xba, 0x6c, 0xe8, 0x51, 0x7d, 0x52, 0xc6, 0x86, 0xac,
	0x38, 0x94, 0x3f, 0x19, 0x33, 0xe8, 0x97, 0x46, 0x3b, 0xd2, 0x2e, 0x8b, 0x87, 0xac, 0x48, 0x64,
	0x88, 0xe2, 0x15, 0x92, 0x59, 0xdb, 0x18, 0x55, 0xbd, 0x38, 0xe5, 0x3a, 0xbb, 0x69, 0x3e, 0x90,
	0xb5, 0xea, 0x23, 0x40, 0x42, 0xc4, 0x2b, 0xe8, 0xcd, 0x4d, 0x45, 0x1e, 0x70, 0x32, 0xce, 0x46,
	0x3b, 0x9b, 0xdf, 0xd3, 0xa5, 0xa9, 0x48, 0xfa, 0x96, 0xb8, 0x80, 0xd3, 0x7b, 0xb3, 0xd6, 0x9b,
	0x3f, 0x49, 0x5f, 0x1d, 0x59, 0x87, 0x08, 0xbd, 0x67, 0xe5, 0x16, 0x3b, 0xae, 0x7f, 0x5f, 0x4e,
	0x20, 0xdd, 0x07, 0xe0, 0x11, 0xf4, 0x67, 0xba, 0xd6, 0x66, 0xad, 0xd3, 0x08, 0x39, 0xc4, 0x4f,
	0x75, 0xca, 0x10, 0x80, 0x4f, 0xd5, 0xb2, 0xa1, 0x2a, 0x8d, 0xc7, 0x16, 0xb8, 0xf4, 0xcb, 0xf1,
	0x06, 0xf8, 0x76, 0x1c, 0x8f, 0x83, 0x8f, 0x3f, 0x46, 0x7e, 0xf6, 0x9f, 0x9e, 0x88, 0x0a, 0x86,
	0xb7, 0x30, 0x08, 0x62, 0x78, 0x1e, 0x5a, 0x7b, 0xaa, 0xf9, 0x5f, 0x9a, 0x88, 0xae, 0xd9, 0x1b,
	0xf7, 0x67, 0x9f, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x99, 0x61, 0x9a, 0x0a, 0x86, 0x01, 0x00,
	0x00,
}
