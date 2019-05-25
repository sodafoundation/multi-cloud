// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamover.proto

package datamover

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{0}
}
func (m *KV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KV.Unmarshal(m, b)
}
func (m *KV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KV.Marshal(b, m, deterministic)
}
func (dst *KV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KV.Merge(dst, src)
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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{1}
}
func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (dst *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(dst, src)
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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{2}
}
func (m *Connector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connector.Unmarshal(m, b)
}
func (m *Connector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connector.Marshal(b, m, deterministic)
}
func (dst *Connector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connector.Merge(dst, src)
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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{3}
}
func (m *RunJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobRequest.Unmarshal(m, b)
}
func (m *RunJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobRequest.Marshal(b, m, deterministic)
}
func (dst *RunJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobRequest.Merge(dst, src)
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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{4}
}
func (m *RunJobResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunJobResponse.Unmarshal(m, b)
}
func (m *RunJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunJobResponse.Marshal(b, m, deterministic)
}
func (dst *RunJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunJobResponse.Merge(dst, src)
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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{5}
}
func (m *LifecycleActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionRequest.Unmarshal(m, b)
}
func (m *LifecycleActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionRequest.Marshal(b, m, deterministic)
}
func (dst *LifecycleActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionRequest.Merge(dst, src)
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
	return fileDescriptor_datamover_59927d4c780b36c9, []int{6}
}
func (m *LifecycleActionResonse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LifecycleActionResonse.Unmarshal(m, b)
}
func (m *LifecycleActionResonse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LifecycleActionResonse.Marshal(b, m, deterministic)
}
func (dst *LifecycleActionResonse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LifecycleActionResonse.Merge(dst, src)
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
	proto.RegisterType((*RunJobResponse)(nil), "RunJobResponse")
	proto.RegisterType((*LifecycleActionRequest)(nil), "LifecycleActionRequest")
	proto.RegisterType((*LifecycleActionResonse)(nil), "LifecycleActionResonse")
}

func init() { proto.RegisterFile("datamover.proto", fileDescriptor_datamover_59927d4c780b36c9) }

var fileDescriptor_datamover_59927d4c780b36c9 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0xa5, 0xc9, 0x9a, 0xb6, 0x77, 0xac, 0x43, 0x16, 0x54, 0xd6, 0x90, 0x50, 0x65, 0x10, 0xaa,
	0x06, 0xca, 0x43, 0x79, 0x80, 0x57, 0x36, 0x09, 0x09, 0x0a, 0x3c, 0x78, 0xd3, 0xde, 0x9d, 0xf8,
	0xb6, 0x72, 0x9b, 0xc6, 0xc5, 0x71, 0x26, 0xca, 0xef, 0xe2, 0x77, 0xf0, 0x9b, 0x90, 0x9d, 0x8f,
	0x36, 0xda, 0xde, 0x72, 0xce, 0x3d, 0xe7, 0x5e, 0xfb, 0xdc, 0x18, 0xce, 0xa5, 0xb0, 0x62, 0xab,
	0xef, 0xd1, 0xc4, 0x3b, 0xa3, 0xad, 0x66, 0xef, 0x21, 0x58, 0xdc, 0x91, 0x67, 0x10, 0x6e, 0x70,
	0x4f, 0x7b, 0xd3, 0xde, 0x6c, 0xc4, 0xdd, 0x27, 0x79, 0x0e, 0xfd, 0x7b, 0x91, 0x95, 0x48, 0x03,
	0xcf, 0x55, 0x80, 0x7d, 0x84, 0xe8, 0x8b, 0xca, 0x2c, 0x1a, 0x32, 0x81, 0x68, 0x67, 0x70, 0xa9,
	0x7e, 0xd7, 0xa6, 0x1a, 0x91, 0x17, 0x10, 0x5a, 0xb1, 0xa2, 0xc1, 0x34, 0x9c, 0x9d, 0xce, 0xc3,
	0x78, 0x71, 0xc7, 0x1d, 0x66, 0x12, 0x46, 0xd7, 0x3a, 0xcf, 0x31, 0xb5, 0xda, 0x10, 0x02, 0x27,
	0xb7, 0xfb, 0x1d, 0xd6, 0x4e, 0xff, 0x4d, 0x5e, 0x01, 0x5c, 0x95, 0xe9, 0x06, 0xed, 0x4f, 0xb1,
	0x6d, 0x86, 0x1e, 0x31, 0xe4, 0x35, 0x80, 0x6b, 0x70, 0xad, 0xf3, 0xa5, 0x5a, 0xd1, 0xf0, 0xd0,
	0xfe, 0x88, 0x66, 0x7f, 0x7b, 0x70, 0xc6, 0xcb, 0xfc, 0x9b, 0x4e, 0x38, 0xfe, 0x2a, 0xb1, 0xb0,
	0x64, 0x0c, 0x81, 0x92, 0xf5, 0xa0, 0x40, 0x49, 0x72, 0x09, 0x50, 0xe8, 0xd2, 0xa4, 0xe8, 0x5c,
	0x7e, 0xcc, 0xe9, 0x1c, 0xe2, 0xf6, 0x68, 0xfc, 0xa8, 0x4a, 0xde, 0xc2, 0x50, 0x62, 0x61, 0xbd,
	0x32, 0x7c, 0xa0, 0x6c, 0x6b, 0xe4, 0x25, 0x9c, 0x2c, 0x55, 0x66, 0xe9, 0x89, 0xd7, 0x0c, 0xe2,
	0x2a, 0x21, 0xee, 0x49, 0xc2, 0xe0, 0xa9, 0xc1, 0xad, 0x50, 0xf9, 0x8d, 0x6f, 0x4c, 0xfb, 0xd3,
	0xde, 0x6c, 0xc8, 0x3b, 0x1c, 0x63, 0x30, 0x6e, 0x4e, 0x5d, 0xec, 0x74, 0x5e, 0xa0, 0xdb, 0x07,
	0x1a, 0xd3, 0xec, 0x03, 0x8d, 0x61, 0xff, 0x02, 0x98, 0x7c, 0x57, 0x4b, 0x4c, 0xf7, 0x69, 0x86,
	0x9f, 0x53, 0xab, 0x74, 0xde, 0xdc, 0x71, 0x02, 0x91, 0x4e, 0xd6, 0x8b, 0x76, 0x7f, 0x35, 0x72,
	0x91, 0x26, 0x0f, 0x22, 0x3d, 0x30, 0xce, 0x27, 0x7c, 0x23, 0x7f, 0xbb, 0x3e, 0xaf, 0x91, 0xf3,
	0x55, 0x29, 0xdc, 0x2a, 0x34, 0xfe, 0x56, 0x7d, 0x7e, 0xc4, 0xb8, 0xba, 0x15, 0x66, 0x85, 0xd6,
	0xd7, 0xfb, 0x55, 0xfd, 0xc0, 0x90, 0x37, 0x70, 0x56, 0xa9, 0xaf, 0x44, 0xba, 0xc1, 0x5c, 0xd2,
	0xc8, 0x8f, 0xee, 0x92, 0x4e, 0x55, 0x79, 0x1a, 0xd5, 0xa0, 0x52, 0x75, 0x48, 0x42, 0x61, 0xa0,
	0x93, 0xf5, 0x8d, 0xfa, 0x83, 0x74, 0x38, 0xed, 0xcd, 0x42, 0xde, 0x40, 0x17, 0x6c, 0x26, 0x0a,
	0xfb, 0x43, 0x4b, 0xb5, 0x54, 0x28, 0xe9, 0xc8, 0x97, 0x3b, 0x1c, 0xb9, 0x80, 0x61, 0xb9, 0xcb,
	0xb4, 0x90, 0x5f, 0x25, 0x05, 0xdf, 0xbe, 0xc5, 0xec, 0xf2, 0x91, 0x3c, 0x8b, 0xc7, 0xc3, 0x9f,
	0x7f, 0x82, 0x51, 0xfb, 0x6e, 0xc8, 0x3b, 0x88, 0x78, 0x99, 0xaf, 0x75, 0x42, 0xc6, 0x71, 0xe7,
	0x67, 0xbb, 0x38, 0x8f, 0xbb, 0x6b, 0x64, 0x4f, 0x92, 0xc8, 0xbf, 0xb2, 0x0f, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x94, 0x35, 0x1f, 0xb6, 0x78, 0x03, 0x00, 0x00,
}
