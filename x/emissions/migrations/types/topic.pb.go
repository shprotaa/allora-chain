// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/v1/topic.proto

package types

import (
	fmt "fmt"
	github_com_allora_network_allora_chain_math "github.com/allora-network/allora-chain/math"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Topic struct {
	Id              uint64                                          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator         string                                          `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Metadata        string                                          `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LossLogic       string                                          `protobuf:"bytes,4,opt,name=loss_logic,json=lossLogic,proto3" json:"loss_logic,omitempty"`
	LossMethod      string                                          `protobuf:"bytes,5,opt,name=loss_method,json=lossMethod,proto3" json:"loss_method,omitempty"`
	InferenceLogic  string                                          `protobuf:"bytes,6,opt,name=inference_logic,json=inferenceLogic,proto3" json:"inference_logic,omitempty"`
	InferenceMethod string                                          `protobuf:"bytes,7,opt,name=inference_method,json=inferenceMethod,proto3" json:"inference_method,omitempty"`
	EpochLastEnded  int64                                           `protobuf:"varint,8,opt,name=epoch_last_ended,json=epochLastEnded,proto3" json:"epoch_last_ended,omitempty"`
	EpochLength     int64                                           `protobuf:"varint,9,opt,name=epoch_length,json=epochLength,proto3" json:"epoch_length,omitempty"`
	GroundTruthLag  int64                                           `protobuf:"varint,10,opt,name=ground_truth_lag,json=groundTruthLag,proto3" json:"ground_truth_lag,omitempty"`
	DefaultArg      string                                          `protobuf:"bytes,11,opt,name=default_arg,json=defaultArg,proto3" json:"default_arg,omitempty"`
	PNorm           github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,12,opt,name=p_norm,json=pNorm,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"p_norm"`
	AlphaRegret     github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,13,opt,name=alpha_regret,json=alphaRegret,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"alpha_regret"`
	AllowNegative   bool                                            `protobuf:"varint,14,opt,name=allow_negative,json=allowNegative,proto3" json:"allow_negative,omitempty"`
	Epsilon         github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,15,opt,name=epsilon,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"epsilon"`
	InitialRegret   github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,16,opt,name=initial_regret,json=initialRegret,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"initial_regret"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5610c9d5deb158, []int{0}
}
func (m *Topic) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return m.Size()
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Topic) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Topic) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *Topic) GetLossLogic() string {
	if m != nil {
		return m.LossLogic
	}
	return ""
}

func (m *Topic) GetLossMethod() string {
	if m != nil {
		return m.LossMethod
	}
	return ""
}

func (m *Topic) GetInferenceLogic() string {
	if m != nil {
		return m.InferenceLogic
	}
	return ""
}

func (m *Topic) GetInferenceMethod() string {
	if m != nil {
		return m.InferenceMethod
	}
	return ""
}

func (m *Topic) GetEpochLastEnded() int64 {
	if m != nil {
		return m.EpochLastEnded
	}
	return 0
}

func (m *Topic) GetEpochLength() int64 {
	if m != nil {
		return m.EpochLength
	}
	return 0
}

func (m *Topic) GetGroundTruthLag() int64 {
	if m != nil {
		return m.GroundTruthLag
	}
	return 0
}

func (m *Topic) GetDefaultArg() string {
	if m != nil {
		return m.DefaultArg
	}
	return ""
}

func (m *Topic) GetAllowNegative() bool {
	if m != nil {
		return m.AllowNegative
	}
	return false
}

type TopicList struct {
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (m *TopicList) Reset()         { *m = TopicList{} }
func (m *TopicList) String() string { return proto.CompactTextString(m) }
func (*TopicList) ProtoMessage()    {}
func (*TopicList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5610c9d5deb158, []int{1}
}
func (m *TopicList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopicList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopicList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopicList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicList.Merge(m, src)
}
func (m *TopicList) XXX_Size() int {
	return m.Size()
}
func (m *TopicList) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicList.DiscardUnknown(m)
}

var xxx_messageInfo_TopicList proto.InternalMessageInfo

func (m *TopicList) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type TimestampedActorNonce struct {
	BlockHeight int64  `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Actor       string `protobuf:"bytes,2,opt,name=actor,proto3" json:"actor,omitempty"`
	Nonce       *Nonce `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *TimestampedActorNonce) Reset()         { *m = TimestampedActorNonce{} }
func (m *TimestampedActorNonce) String() string { return proto.CompactTextString(m) }
func (*TimestampedActorNonce) ProtoMessage()    {}
func (*TimestampedActorNonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae5610c9d5deb158, []int{2}
}
func (m *TimestampedActorNonce) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimestampedActorNonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimestampedActorNonce.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimestampedActorNonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimestampedActorNonce.Merge(m, src)
}
func (m *TimestampedActorNonce) XXX_Size() int {
	return m.Size()
}
func (m *TimestampedActorNonce) XXX_DiscardUnknown() {
	xxx_messageInfo_TimestampedActorNonce.DiscardUnknown(m)
}

var xxx_messageInfo_TimestampedActorNonce proto.InternalMessageInfo

func (m *TimestampedActorNonce) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *TimestampedActorNonce) GetActor() string {
	if m != nil {
		return m.Actor
	}
	return ""
}

func (m *TimestampedActorNonce) GetNonce() *Nonce {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func init() {
	proto.RegisterType((*Topic)(nil), "emissions.v1.Topic")
	proto.RegisterType((*TopicList)(nil), "emissions.v1.TopicList")
	proto.RegisterType((*TimestampedActorNonce)(nil), "emissions.v1.TimestampedActorNonce")
}

func init() { proto.RegisterFile("emissions/v1/topic.proto", fileDescriptor_ae5610c9d5deb158) }

var fileDescriptor_ae5610c9d5deb158 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x18, 0x86, 0xb3, 0x4d, 0x93, 0x36, 0x93, 0x34, 0xad, 0x63, 0x85, 0xb1, 0x60, 0x12, 0x0b, 0x62,
	0x8a, 0x98, 0xa5, 0x7a, 0x50, 0xbc, 0xb5, 0x2a, 0x78, 0x88, 0x01, 0x97, 0x9e, 0x0a, 0xba, 0x4c,
	0x67, 0xa7, 0xb3, 0x43, 0x77, 0x67, 0x96, 0x99, 0x49, 0xab, 0x47, 0xff, 0x81, 0x3f, 0xc1, 0x9f,
	0xd3, 0x63, 0x8f, 0xe2, 0xa1, 0x48, 0x7b, 0xf1, 0x5f, 0x28, 0xf3, 0xed, 0xa6, 0x51, 0xbc, 0x08,
	0xb9, 0x84, 0xfd, 0x9e, 0xf7, 0xcd, 0xfb, 0xf1, 0x85, 0xbc, 0x8b, 0x08, 0xcf, 0xa5, 0xb5, 0x52,
	0x2b, 0x1b, 0x9e, 0xee, 0x86, 0x4e, 0x17, 0x92, 0x8d, 0x0a, 0xa3, 0x9d, 0xc6, 0x9d, 0x1b, 0x65,
	0x74, 0xba, 0xbb, 0x75, 0x97, 0x69, 0x9b, 0x6b, 0x1b, 0x83, 0x16, 0x96, 0x43, 0x69, 0xdc, 0xba,
	0x45, 0x73, 0xa9, 0x74, 0x08, 0x9f, 0x15, 0xda, 0x14, 0x5a, 0xe8, 0xd2, 0xea, 0x9f, 0x2a, 0xfa,
	0xf7, 0x2e, 0xa5, 0x15, 0xe3, 0xa5, 0xb2, 0xfd, 0xab, 0x81, 0x1a, 0x07, 0x7e, 0x37, 0xee, 0xa2,
	0x25, 0x99, 0x90, 0x60, 0x10, 0x0c, 0x97, 0xa3, 0x25, 0x99, 0x60, 0x82, 0x56, 0x98, 0xe1, 0xd4,
	0x69, 0x43, 0x96, 0x06, 0xc1, 0xb0, 0x15, 0xcd, 0x46, 0xbc, 0x85, 0x56, 0x73, 0xee, 0x68, 0x42,
	0x1d, 0x25, 0x75, 0x90, 0x6e, 0x66, 0x7c, 0x0f, 0xa1, 0x4c, 0x5b, 0x1b, 0x67, 0x5a, 0x48, 0x46,
	0x96, 0x41, 0x6d, 0x79, 0x32, 0xf6, 0x00, 0xf7, 0x51, 0x1b, 0xe4, 0x9c, 0xbb, 0x54, 0x27, 0xa4,
	0x01, 0x3a, 0x7c, 0xe3, 0x2d, 0x10, 0xfc, 0x10, 0xad, 0x4b, 0x75, 0xcc, 0x0d, 0x57, 0x8c, 0x57,
	0x21, 0x4d, 0x30, 0x75, 0x6f, 0x70, 0x99, 0xb4, 0x83, 0x36, 0xe6, 0xc6, 0x2a, 0x6e, 0x05, 0x9c,
	0xf3, 0x80, 0x2a, 0x73, 0x88, 0x36, 0x78, 0xa1, 0x59, 0x1a, 0x67, 0xd4, 0xba, 0x98, 0xab, 0x84,
	0x27, 0x64, 0x75, 0x10, 0x0c, 0xeb, 0x51, 0x17, 0xf8, 0x98, 0x5a, 0xf7, 0xda, 0x53, 0x7c, 0x1f,
	0x75, 0x2a, 0x27, 0x57, 0xc2, 0xa5, 0xa4, 0x05, 0xae, 0x76, 0xe9, 0x02, 0xe4, 0xc3, 0x84, 0xd1,
	0x53, 0x95, 0xc4, 0xce, 0x4c, 0x9d, 0xcf, 0x14, 0x04, 0x95, 0x61, 0x25, 0x3f, 0xf0, 0x78, 0x4c,
	0x85, 0xbf, 0x35, 0xe1, 0xc7, 0x74, 0x9a, 0xb9, 0x98, 0x1a, 0x41, 0xda, 0xe5, 0xad, 0x15, 0xda,
	0x33, 0x02, 0x4f, 0x50, 0xb3, 0x88, 0x95, 0x36, 0x39, 0xe9, 0x78, 0x6d, 0xff, 0xd9, 0xf9, 0x65,
	0xbf, 0xf6, 0xfd, 0xb2, 0x1f, 0x0a, 0xe9, 0xd2, 0xe9, 0xd1, 0x88, 0xe9, 0x3c, 0xa4, 0x59, 0xa6,
	0x0d, 0x7d, 0xac, 0xb8, 0x3b, 0xd3, 0xe6, 0x64, 0x36, 0xb2, 0x94, 0x4a, 0x15, 0xe6, 0xd4, 0xa5,
	0xa3, 0x57, 0x9c, 0x45, 0x8d, 0x62, 0xa2, 0x4d, 0x8e, 0x0f, 0x51, 0x87, 0x66, 0x45, 0x4a, 0x63,
	0xc3, 0x85, 0xe1, 0x8e, 0xac, 0x2d, 0x96, 0xda, 0x86, 0xb0, 0x08, 0xb2, 0xf0, 0x03, 0xd4, 0xf5,
	0xae, 0xb3, 0x58, 0x71, 0x41, 0x9d, 0x3c, 0xe5, 0xa4, 0x3b, 0x08, 0x86, 0xab, 0xd1, 0x1a, 0xd0,
	0x49, 0x05, 0xf1, 0x3b, 0xb4, 0xc2, 0x0b, 0x2b, 0x33, 0xad, 0xc8, 0xfa, 0x62, 0xdb, 0x67, 0x39,
	0xf8, 0x03, 0xea, 0x4a, 0x25, 0x9d, 0xa4, 0xd9, 0xec, 0xae, 0x8d, 0xc5, 0x92, 0xd7, 0xaa, 0xb8,
	0xf2, 0xb2, 0xed, 0xe7, 0xa8, 0x05, 0x05, 0x18, 0x4b, 0xeb, 0xf0, 0x23, 0xd4, 0x84, 0x26, 0x5a,
	0x12, 0x0c, 0xea, 0xc3, 0xf6, 0x93, 0xdb, 0xa3, 0x3f, 0xbb, 0x38, 0x02, 0x63, 0x54, 0x59, 0xb6,
	0x3f, 0x07, 0xe8, 0xce, 0x81, 0xcc, 0xb9, 0x75, 0x34, 0x2f, 0x78, 0xb2, 0xc7, 0x9c, 0x36, 0x13,
	0xdf, 0x2d, 0xff, 0x3f, 0x3a, 0xca, 0x34, 0x3b, 0x89, 0x53, 0x2e, 0x45, 0xea, 0xa0, 0x55, 0xf5,
	0xa8, 0x0d, 0xec, 0x0d, 0x20, 0xbc, 0x89, 0x1a, 0x94, 0xcd, 0xcb, 0x55, 0x0e, 0x78, 0x07, 0x35,
	0xa0, 0x9d, 0xd0, 0xab, 0x7f, 0xd6, 0x43, 0x78, 0x54, 0x3a, 0x5e, 0x2c, 0xff, 0xfc, 0xda, 0x0f,
	0xf6, 0xdf, 0x9f, 0x5f, 0xf5, 0x82, 0x8b, 0xab, 0x5e, 0xf0, 0xe3, 0xaa, 0x17, 0x7c, 0xb9, 0xee,
	0xd5, 0x2e, 0xae, 0x7b, 0xb5, 0x6f, 0xd7, 0xbd, 0xda, 0xe1, 0xcb, 0xff, 0xfc, 0x5d, 0x3e, 0x86,
	0xf3, 0x97, 0x43, 0x2e, 0x85, 0xa1, 0x0e, 0x1e, 0xdd, 0xa7, 0x82, 0xdb, 0xa3, 0x26, 0xbc, 0x25,
	0x9e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x50, 0x0c, 0xc1, 0xad, 0x04, 0x00, 0x00,
}

func (this *TimestampedActorNonce) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TimestampedActorNonce)
	if !ok {
		that2, ok := that.(TimestampedActorNonce)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.BlockHeight != that1.BlockHeight {
		return false
	}
	if this.Actor != that1.Actor {
		return false
	}
	if !this.Nonce.Equal(that1.Nonce) {
		return false
	}
	return true
}
func (m *Topic) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Topic) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Topic) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InitialRegret.Size()
		i -= size
		if _, err := m.InitialRegret.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	{
		size := m.Epsilon.Size()
		i -= size
		if _, err := m.Epsilon.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	if m.AllowNegative {
		i--
		if m.AllowNegative {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	{
		size := m.AlphaRegret.Size()
		i -= size
		if _, err := m.AlphaRegret.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.PNorm.Size()
		i -= size
		if _, err := m.PNorm.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTopic(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.DefaultArg) > 0 {
		i -= len(m.DefaultArg)
		copy(dAtA[i:], m.DefaultArg)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.DefaultArg)))
		i--
		dAtA[i] = 0x5a
	}
	if m.GroundTruthLag != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.GroundTruthLag))
		i--
		dAtA[i] = 0x50
	}
	if m.EpochLength != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.EpochLength))
		i--
		dAtA[i] = 0x48
	}
	if m.EpochLastEnded != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.EpochLastEnded))
		i--
		dAtA[i] = 0x40
	}
	if len(m.InferenceMethod) > 0 {
		i -= len(m.InferenceMethod)
		copy(dAtA[i:], m.InferenceMethod)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.InferenceMethod)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.InferenceLogic) > 0 {
		i -= len(m.InferenceLogic)
		copy(dAtA[i:], m.InferenceLogic)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.InferenceLogic)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.LossMethod) > 0 {
		i -= len(m.LossMethod)
		copy(dAtA[i:], m.LossMethod)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.LossMethod)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.LossLogic) > 0 {
		i -= len(m.LossLogic)
		copy(dAtA[i:], m.LossLogic)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.LossLogic)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TopicList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopicList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopicList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Topics) > 0 {
		for iNdEx := len(m.Topics) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Topics[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTopic(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TimestampedActorNonce) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimestampedActorNonce) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimestampedActorNonce) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != nil {
		{
			size, err := m.Nonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTopic(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Actor) > 0 {
		i -= len(m.Actor)
		copy(dAtA[i:], m.Actor)
		i = encodeVarintTopic(dAtA, i, uint64(len(m.Actor)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTopic(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTopic(dAtA []byte, offset int, v uint64) int {
	offset -= sovTopic(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Topic) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTopic(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.LossLogic)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.LossMethod)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.InferenceLogic)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = len(m.InferenceMethod)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	if m.EpochLastEnded != 0 {
		n += 1 + sovTopic(uint64(m.EpochLastEnded))
	}
	if m.EpochLength != 0 {
		n += 1 + sovTopic(uint64(m.EpochLength))
	}
	if m.GroundTruthLag != 0 {
		n += 1 + sovTopic(uint64(m.GroundTruthLag))
	}
	l = len(m.DefaultArg)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	l = m.PNorm.Size()
	n += 1 + l + sovTopic(uint64(l))
	l = m.AlphaRegret.Size()
	n += 1 + l + sovTopic(uint64(l))
	if m.AllowNegative {
		n += 2
	}
	l = m.Epsilon.Size()
	n += 1 + l + sovTopic(uint64(l))
	l = m.InitialRegret.Size()
	n += 2 + l + sovTopic(uint64(l))
	return n
}

func (m *TopicList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Topics) > 0 {
		for _, e := range m.Topics {
			l = e.Size()
			n += 1 + l + sovTopic(uint64(l))
		}
	}
	return n
}

func (m *TimestampedActorNonce) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovTopic(uint64(m.BlockHeight))
	}
	l = len(m.Actor)
	if l > 0 {
		n += 1 + l + sovTopic(uint64(l))
	}
	if m.Nonce != nil {
		l = m.Nonce.Size()
		n += 1 + l + sovTopic(uint64(l))
	}
	return n
}

func sovTopic(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTopic(x uint64) (n int) {
	return sovTopic(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Topic) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: Topic: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Topic: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LossLogic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LossLogic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LossMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LossMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InferenceLogic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InferenceLogic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InferenceMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InferenceMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochLastEnded", wireType)
			}
			m.EpochLastEnded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochLastEnded |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochLength", wireType)
			}
			m.EpochLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochLength |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroundTruthLag", wireType)
			}
			m.GroundTruthLag = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroundTruthLag |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultArg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultArg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PNorm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PNorm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AlphaRegret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AlphaRegret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowNegative", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
			m.AllowNegative = bool(v != 0)
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epsilon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Epsilon.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialRegret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialRegret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopic
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
func (m *TopicList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TopicList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopicList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topics = append(m.Topics, &Topic{})
			if err := m.Topics[len(m.Topics)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopic
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
func (m *TimestampedActorNonce) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopic
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
			return fmt.Errorf("proto: TimestampedActorNonce: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimestampedActorNonce: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopic
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
				return ErrInvalidLengthTopic
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopic
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nonce == nil {
				m.Nonce = &Nonce{}
			}
			if err := m.Nonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopic(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopic
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
func skipTopic(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopic
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
					return 0, ErrIntOverflowTopic
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
					return 0, ErrIntOverflowTopic
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
				return 0, ErrInvalidLengthTopic
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTopic
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTopic
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTopic        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopic          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTopic = fmt.Errorf("proto: unexpected end of group")
)
