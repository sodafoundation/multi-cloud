// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamover.proto

package datamover

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type KV struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KV) Reset()         { *m = KV{} }
func (m *KV) String() string { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()    {}
func (*KV) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{0}
}

func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (m *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(m, src)
}
func (m *KV) XXX_Size() int {
	return xxx_messageInfo_KV.Size(m)
}
func (m *KV) XXX_DiscardUnknown() {
	xxx_messageInfo_KV.DiscardUnknown(m)
}

var xxx_messageInfo_KV proto.InternalMessageInfo

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Filter struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Tag                  []*KV    `protobuf:"bytes,2,rep,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{1}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Filter) GetTag() []*KV {
	if m != nil {
		return m.Tag
	}
	return nil
}

type Connector struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	BucketName           string   `protobuf:"bytes,2,opt,name=BucketName,proto3" json:"BucketName,omitempty"`
	ConnConfig           []*KV    `protobuf:"bytes,3,rep,name=ConnConfig,proto3" json:"ConnConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connector) Reset()         { *m = Connector{} }
func (m *Connector) String() string { return proto.CompactTextString(m) }
func (*Connector) ProtoMessage()    {}
func (*Connector) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{2}
}

func (m *Connector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connector.Unmarshal(m, b)
}
func (m *Connector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connector.Marshal(b, m, deterministic)
}
func (m *Connector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connector.Merge(m, src)
}
func (m *Connector) XXX_Size() int {
	return xxx_messageInfo_Connector.Size(m)
}
func (m *Connector) XXX_DiscardUnknown() {
	xxx_messageInfo_Connector.DiscardUnknown(m)
}

var xxx_messageInfo_Connector proto.InternalMessageInfo

func (m *Connector) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Connector) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *Connector) GetConnConfig() []*KV {
	if m != nil {
		return m.ConnConfig
	}
	return nil
}

type RunJobRequest struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceConn           *Connector `protobuf:"bytes,2,opt,name=sourceConn,proto3" json:"sourceConn,omitempty"`
	DestConn             *Connector `protobuf:"bytes,3,opt,name=destConn,proto3" json:"destConn,omitempty"`
	Filt                 *Filter    `protobuf:"bytes,4,opt,name=filt,proto3" json:"filt,omitempty"`
	RemainSource         bool       `protobuf:"varint,5,opt,name=remainSource,proto3" json:"remainSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RunJobRequest) Reset()         { *m = RunJobRequest{} }
func (m *RunJobRequest) String() string { return proto.CompactTextString(m) }
func (*RunJobRequest) ProtoMessage()    {}
func (*RunJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{3}
}

func (m *RunJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobRequest.Unmarshal(m, b)
}
func (m *RunJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobRequest.Marshal(b, m, deterministic)
}
func (m *RunJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobRequest.Merge(m, src)
}
func (m *RunJobRequest) XXX_Size() int {
	return xxx_messageInfo_RunJobRequest.Size(m)
}
func (m *RunJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobRequest proto.InternalMessageInfo

func (m *RunJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RunJobRequest) GetSourceConn() *Connector {
	if m != nil {
		return m.SourceConn
	}
	return nil
}

func (m *RunJobRequest) GetDestConn() *Connector {
	if m != nil {
		return m.DestConn
	}
	return nil
}

func (m *RunJobRequest) GetFilt() *Filter {
	if m != nil {
		return m.Filt
	}
	return nil
}

func (m *RunJobRequest) GetRemainSource() bool {
	if m != nil {
		return m.RemainSource
	}
	return false
}

type AbortJobRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AbortJobRequest) Reset()         { *m = AbortJobRequest{} }
func (m *AbortJobRequest) String() string { return proto.CompactTextString(m) }
func (*AbortJobRequest) ProtoMessage()    {}
func (*AbortJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{4}
}

func (m *AbortJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AbortJobRequest.Unmarshal(m, b)
}
func (m *AbortJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AbortJobRequest.Marshal(b, m, deterministic)
}
func (m *AbortJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AbortJobRequest.Merge(m, src)
}
func (m *AbortJobRequest) XXX_Size() int {
	return xxx_messageInfo_AbortJobRequest.Size(m)
}
func (m *AbortJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AbortJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AbortJobRequest proto.InternalMessageInfo

func (m *AbortJobRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RunJobResponse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunJobResponse) Reset()         { *m = RunJobResponse{} }
func (m *RunJobResponse) String() string { return proto.CompactTextString(m) }
func (*RunJobResponse) ProtoMessage()    {}
func (*RunJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{5}
}

func (m *RunJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobResponse.Unmarshal(m, b)
}
func (m *RunJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobResponse.Marshal(b, m, deterministic)
}
func (m *RunJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobResponse.Merge(m, src)
}
func (m *RunJobResponse) XXX_Size() int {
	return xxx_messageInfo_RunJobResponse.Size(m)
}
func (m *RunJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunJobResponse proto.InternalMessageInfo

func (m *RunJobResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type LifecycleActionRequest struct {
	ObjKey               string   `protobuf:"bytes,1,opt,name=objKey,proto3" json:"objKey,omitempty"`
	BucketName           string   `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Action               int32    `protobuf:"varint,3,opt,name=action,proto3" json:"action,omitempty"`
	SourceTier           int32    `protobuf:"varint,4,opt,name=sourceTier,proto3" json:"sourceTier,omitempty"`
	TargetTier           int32    `protobuf:"varint,5,opt,name=targetTier,proto3" json:"targetTier,omitempty"`
	SourceBackend        string   `protobuf:"bytes,6,opt,name=sourceBackend,proto3" json:"sourceBackend,omitempty"`
	TargetBackend        string   `protobuf:"bytes,7,opt,name=targetBackend,proto3" json:"targetBackend,omitempty"`
	ObjSize              int64    `protobuf:"varint,8,opt,name=objSize,proto3" json:"objSize,omitempty"`
	LastModified         int64    `protobuf:"varint,9,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	UploadId             string   `protobuf:"bytes,10,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleActionRequest) Reset()         { *m = LifecycleActionRequest{} }
func (m *LifecycleActionRequest) String() string { return proto.CompactTextString(m) }
func (*LifecycleActionRequest) ProtoMessage()    {}
func (*LifecycleActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{6}
}

func (m *LifecycleActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionRequest.Unmarshal(m, b)
}
func (m *LifecycleActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionRequest.Marshal(b, m, deterministic)
}
func (m *LifecycleActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionRequest.Merge(m, src)
}
func (m *LifecycleActionRequest) XXX_Size() int {
	return xxx_messageInfo_LifecycleActionRequest.Size(m)
}
func (m *LifecycleActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleActionRequest proto.InternalMessageInfo

func (m *LifecycleActionRequest) GetObjKey() string {
	if m != nil {
		return m.ObjKey
	}
	return ""
}

func (m *LifecycleActionRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *LifecycleActionRequest) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *LifecycleActionRequest) GetSourceTier() int32 {
	if m != nil {
		return m.SourceTier
	}
	return 0
}

func (m *LifecycleActionRequest) GetTargetTier() int32 {
	if m != nil {
		return m.TargetTier
	}
	return 0
}

func (m *LifecycleActionRequest) GetSourceBackend() string {
	if m != nil {
		return m.SourceBackend
	}
	return ""
}

func (m *LifecycleActionRequest) GetTargetBackend() string {
	if m != nil {
		return m.TargetBackend
	}
	return ""
}

func (m *LifecycleActionRequest) GetObjSize() int64 {
	if m != nil {
		return m.ObjSize
	}
	return 0
}

func (m *LifecycleActionRequest) GetLastModified() int64 {
	if m != nil {
		return m.LastModified
	}
	return 0
}

func (m *LifecycleActionRequest) GetUploadId() string {
	if m != nil {
		return m.UploadId
	}
	return ""
}

type LifecycleActionResonse struct {
	Err                  string   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LifecycleActionResonse) Reset()         { *m = LifecycleActionResonse{} }
func (m *LifecycleActionResonse) String() string { return proto.CompactTextString(m) }
func (*LifecycleActionResonse) ProtoMessage()    {}
func (*LifecycleActionResonse) Descriptor() ([]byte, []int) {
	return fileDescriptor_302d3374d95ec8eb, []int{7}
}

func (m *LifecycleActionResonse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionResonse.Unmarshal(m, b)
}
func (m *LifecycleActionResonse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionResonse.Marshal(b, m, deterministic)
}
func (m *LifecycleActionResonse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionResonse.Merge(m, src)
}
func (m *LifecycleActionResonse) XXX_Size() int {
	return xxx_messageInfo_LifecycleActionResonse.Size(m)
}
func (m *LifecycleActionResonse) XXX_DiscardUnknown() {
	xxx_messageInfo_LifecycleActionResonse.DiscardUnknown(m)
}

var xxx_messageInfo_LifecycleActionResonse proto.InternalMessageInfo

func (m *LifecycleActionResonse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*KV)(nil), "KV")
	proto.RegisterType((*Filter)(nil), "Filter")
	proto.RegisterType((*Connector)(nil), "Connector")
	proto.RegisterType((*RunJobRequest)(nil), "RunJobRequest")
	proto.RegisterType((*AbortJobRequest)(nil), "AbortJobRequest")
	proto.RegisterType((*RunJobResponse)(nil), "RunJobResponse")
	proto.RegisterType((*LifecycleActionRequest)(nil), "LifecycleActionRequest")
	proto.RegisterType((*LifecycleActionResonse)(nil), "LifecycleActionResonse")
}

func init() { proto.RegisterFile("datamover.proto", fileDescriptor_302d3374d95ec8eb) }

var fileDescriptor_302d3374d95ec8eb = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x25, 0x49, 0x9b, 0xb6, 0xb3, 0x6c, 0x0b, 0x16, 0x94, 0xa8, 0x48, 0xa8, 0x18, 0x84, 0xaa,
	0x05, 0xe5, 0x50, 0x0e, 0x9c, 0x77, 0x17, 0x81, 0xa0, 0xc0, 0xc1, 0xbb, 0xda, 0xbb, 0x93, 0x4c,
	0x2b, 0xb7, 0x69, 0x5c, 0x1c, 0x67, 0x45, 0xb9, 0xf1, 0x4f, 0x7c, 0x07, 0xdf, 0x84, 0xec, 0x24,
	0x6d, 0xc3, 0x96, 0x5b, 0xe6, 0xbd, 0x37, 0x6f, 0xec, 0x37, 0x31, 0x0c, 0x12, 0xae, 0xf9, 0x5a,
	0xde, 0xa2, 0x0a, 0x37, 0x4a, 0x6a, 0x49, 0xdf, 0x80, 0x3b, 0xbb, 0x21, 0x0f, 0xc0, 0x5b, 0xe1,
	0x36, 0x70, 0xc6, 0xce, 0xa4, 0xc7, 0xcc, 0x27, 0x79, 0x04, 0xed, 0x5b, 0x9e, 0x16, 0x18, 0xb8,
	0x16, 0x2b, 0x0b, 0xfa, 0x0e, 0xfc, 0x0f, 0x22, 0xd5, 0xa8, 0xc8, 0x10, 0xfc, 0x8d, 0xc2, 0xb9,
	0xf8, 0x51, 0x35, 0x55, 0x15, 0x79, 0x0c, 0x9e, 0xe6, 0x8b, 0xc0, 0x1d, 0x7b, 0x93, 0x93, 0xa9,
	0x17, 0xce, 0x6e, 0x98, 0xa9, 0x69, 0x02, 0xbd, 0x4b, 0x99, 0x65, 0x18, 0x6b, 0xa9, 0x08, 0x81,
	0xd6, 0xf5, 0x76, 0x83, 0x55, 0xa7, 0xfd, 0x26, 0xcf, 0x00, 0x2e, 0x8a, 0x78, 0x85, 0xfa, 0x1b,
	0x5f, 0xd7, 0x43, 0x0f, 0x10, 0xf2, 0x02, 0xc0, 0x18, 0x5c, 0xca, 0x6c, 0x2e, 0x16, 0x81, 0xb7,
	0xb7, 0x3f, 0x80, 0xe9, 0x6f, 0x07, 0x4e, 0x59, 0x91, 0x7d, 0x96, 0x11, 0xc3, 0xef, 0x05, 0xe6,
	0x9a, 0xf4, 0xc1, 0x15, 0x49, 0x35, 0xc8, 0x15, 0x09, 0x39, 0x03, 0xc8, 0x65, 0xa1, 0x62, 0x34,
	0x5d, 0x76, 0xcc, 0xc9, 0x14, 0xc2, 0xdd, 0xd1, 0xd8, 0x01, 0x4b, 0x5e, 0x41, 0x37, 0xc1, 0x5c,
	0x5b, 0xa5, 0x77, 0x47, 0xb9, 0xe3, 0xc8, 0x53, 0x68, 0xcd, 0x45, 0xaa, 0x83, 0x96, 0xd5, 0x74,
	0xc2, 0x32, 0x21, 0x66, 0x41, 0x42, 0xe1, 0xbe, 0xc2, 0x35, 0x17, 0xd9, 0x95, 0x35, 0x0e, 0xda,
	0x63, 0x67, 0xd2, 0x65, 0x0d, 0x8c, 0x3e, 0x87, 0xc1, 0x79, 0x24, 0x95, 0xfe, 0xff, 0xb9, 0x29,
	0x85, 0x7e, 0x7d, 0xb1, 0x7c, 0x23, 0xb3, 0x1c, 0xcd, 0xca, 0x50, 0xa9, 0x7a, 0x65, 0xa8, 0x14,
	0xfd, 0xe3, 0xc2, 0xf0, 0x8b, 0x98, 0x63, 0xbc, 0x8d, 0x53, 0x3c, 0x8f, 0xb5, 0x90, 0x59, 0x6d,
	0x37, 0x04, 0x5f, 0x46, 0xcb, 0xd9, 0x6e, 0xc5, 0x55, 0x65, 0x52, 0x8f, 0xee, 0xa4, 0xbe, 0x47,
	0x4c, 0x1f, 0xb7, 0x46, 0x36, 0x80, 0x36, 0xab, 0x2a, 0xd3, 0x57, 0x06, 0x75, 0x2d, 0x50, 0xd9,
	0x8b, 0xb7, 0xd9, 0x01, 0x62, 0x78, 0xcd, 0xd5, 0x02, 0xb5, 0xe5, 0xdb, 0x25, 0xbf, 0x47, 0xc8,
	0x4b, 0x38, 0x2d, 0xd5, 0x17, 0x3c, 0x5e, 0x61, 0x96, 0x04, 0xbe, 0x1d, 0xdd, 0x04, 0x8d, 0xaa,
	0xec, 0xa9, 0x55, 0x9d, 0x52, 0xd5, 0x00, 0x49, 0x00, 0x1d, 0x19, 0x2d, 0xaf, 0xc4, 0x4f, 0x0c,
	0xba, 0x63, 0x67, 0xe2, 0xb1, 0xba, 0x34, 0xd9, 0xa7, 0x3c, 0xd7, 0x5f, 0x65, 0x22, 0xe6, 0x02,
	0x93, 0xa0, 0x67, 0xe9, 0x06, 0x46, 0x46, 0xd0, 0x2d, 0x36, 0xa9, 0xe4, 0xc9, 0xa7, 0x24, 0x00,
	0x6b, 0xbf, 0xab, 0xe9, 0xd9, 0x91, 0x3c, 0xf3, 0xe3, 0xe1, 0x4f, 0x7f, 0x39, 0xd0, 0xdb, 0xbd,
	0x2d, 0xf2, 0x1a, 0x7c, 0x56, 0x64, 0x4b, 0x19, 0x91, 0x7e, 0xd8, 0xf8, 0x21, 0x47, 0x83, 0xb0,
	0xb9, 0x47, 0x7a, 0x8f, 0x7c, 0x84, 0x87, 0xef, 0xe5, 0x3f, 0x83, 0xc8, 0x93, 0xf0, 0xf8, 0x2a,
	0x47, 0x47, 0x88, 0xbc, 0x34, 0x8a, 0x7c, 0xfb, 0xa4, 0xdf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xff, 0x71, 0x1a, 0x73, 0xe5, 0x03, 0x00, 0x00,
}
