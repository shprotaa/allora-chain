// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: emissions/v1/events.proto

package types

import (
	fmt "fmt"
	github_com_allora_network_allora_chain_math "github.com/allora-network/allora-chain/math"
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

type ActorType int32

const (
	ActorType_INFERER    ActorType = 0
	ActorType_FORECASTER ActorType = 1
	ActorType_REPUTER    ActorType = 2
)

var ActorType_name = map[int32]string{
	0: "INFERER",
	1: "FORECASTER",
	2: "REPUTER",
}

var ActorType_value = map[string]int32{
	"INFERER":    0,
	"FORECASTER": 1,
	"REPUTER":    2,
}

func (x ActorType) String() string {
	return proto.EnumName(ActorType_name, int32(x))
}

func (ActorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5cc3b6a19d61d65b, []int{0}
}

type EventScoresSet struct {
	ActorType   ActorType                                         `protobuf:"varint,1,opt,name=actor_type,json=actorType,proto3,enum=emissions.v1.ActorType" json:"actor_type,omitempty"`
	TopicId     uint64                                            `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	BlockHeight int64                                             `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Addresses   []string                                          `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Scores      []github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,5,rep,name=scores,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"scores"`
}

func (m *EventScoresSet) Reset()         { *m = EventScoresSet{} }
func (m *EventScoresSet) String() string { return proto.CompactTextString(m) }
func (*EventScoresSet) ProtoMessage()    {}
func (*EventScoresSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc3b6a19d61d65b, []int{0}
}
func (m *EventScoresSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventScoresSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventScoresSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventScoresSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventScoresSet.Merge(m, src)
}
func (m *EventScoresSet) XXX_Size() int {
	return m.Size()
}
func (m *EventScoresSet) XXX_DiscardUnknown() {
	xxx_messageInfo_EventScoresSet.DiscardUnknown(m)
}

var xxx_messageInfo_EventScoresSet proto.InternalMessageInfo

func (m *EventScoresSet) GetActorType() ActorType {
	if m != nil {
		return m.ActorType
	}
	return ActorType_INFERER
}

func (m *EventScoresSet) GetTopicId() uint64 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *EventScoresSet) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *EventScoresSet) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type EventRewardsSettled struct {
	ActorType   ActorType                                         `protobuf:"varint,1,opt,name=actor_type,json=actorType,proto3,enum=emissions.v1.ActorType" json:"actor_type,omitempty"`
	TopicId     uint64                                            `protobuf:"varint,2,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	BlockHeight int64                                             `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Addresses   []string                                          `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Rewards     []github_com_allora_network_allora_chain_math.Dec `protobuf:"bytes,5,rep,name=rewards,proto3,customtype=github.com/allora-network/allora-chain/math.Dec" json:"rewards"`
}

func (m *EventRewardsSettled) Reset()         { *m = EventRewardsSettled{} }
func (m *EventRewardsSettled) String() string { return proto.CompactTextString(m) }
func (*EventRewardsSettled) ProtoMessage()    {}
func (*EventRewardsSettled) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc3b6a19d61d65b, []int{1}
}
func (m *EventRewardsSettled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardsSettled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardsSettled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardsSettled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardsSettled.Merge(m, src)
}
func (m *EventRewardsSettled) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardsSettled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardsSettled.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardsSettled proto.InternalMessageInfo

func (m *EventRewardsSettled) GetActorType() ActorType {
	if m != nil {
		return m.ActorType
	}
	return ActorType_INFERER
}

func (m *EventRewardsSettled) GetTopicId() uint64 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *EventRewardsSettled) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *EventRewardsSettled) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type EventNetworkLossSet struct {
	TopicId     uint64       `protobuf:"varint,1,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
	BlockHeight int64        `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	ValueBundle *ValueBundle `protobuf:"bytes,3,opt,name=value_bundle,json=valueBundle,proto3" json:"value_bundle,omitempty"`
}

func (m *EventNetworkLossSet) Reset()         { *m = EventNetworkLossSet{} }
func (m *EventNetworkLossSet) String() string { return proto.CompactTextString(m) }
func (*EventNetworkLossSet) ProtoMessage()    {}
func (*EventNetworkLossSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc3b6a19d61d65b, []int{2}
}
func (m *EventNetworkLossSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventNetworkLossSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventNetworkLossSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventNetworkLossSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventNetworkLossSet.Merge(m, src)
}
func (m *EventNetworkLossSet) XXX_Size() int {
	return m.Size()
}
func (m *EventNetworkLossSet) XXX_DiscardUnknown() {
	xxx_messageInfo_EventNetworkLossSet.DiscardUnknown(m)
}

var xxx_messageInfo_EventNetworkLossSet proto.InternalMessageInfo

func (m *EventNetworkLossSet) GetTopicId() uint64 {
	if m != nil {
		return m.TopicId
	}
	return 0
}

func (m *EventNetworkLossSet) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *EventNetworkLossSet) GetValueBundle() *ValueBundle {
	if m != nil {
		return m.ValueBundle
	}
	return nil
}

func init() {
	proto.RegisterEnum("emissions.v1.ActorType", ActorType_name, ActorType_value)
	proto.RegisterType((*EventScoresSet)(nil), "emissions.v1.EventScoresSet")
	proto.RegisterType((*EventRewardsSettled)(nil), "emissions.v1.EventRewardsSettled")
	proto.RegisterType((*EventNetworkLossSet)(nil), "emissions.v1.EventNetworkLossSet")
}

func init() { proto.RegisterFile("emissions/v1/events.proto", fileDescriptor_5cc3b6a19d61d65b) }

var fileDescriptor_5cc3b6a19d61d65b = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xe3, 0x76, 0xac, 0xd4, 0xad, 0xaa, 0xc9, 0x20, 0x91, 0x56, 0x28, 0x2b, 0x3d, 0x55,
	0x48, 0x24, 0xda, 0x10, 0x7f, 0x0e, 0x5c, 0x56, 0xc8, 0xc4, 0x24, 0xb4, 0x81, 0x3b, 0x38, 0x70,
	0xa9, 0xdc, 0xe4, 0x55, 0x13, 0x2d, 0x8d, 0x23, 0xdb, 0xc9, 0xd8, 0xb7, 0x80, 0x6f, 0xb5, 0xe3,
	0x8e, 0x88, 0xc3, 0x84, 0xda, 0xcf, 0xc0, 0x8d, 0x03, 0xb2, 0x43, 0xd6, 0x55, 0x48, 0x08, 0x89,
	0xcb, 0x6e, 0x7e, 0xde, 0xe7, 0xc9, 0x1b, 0xff, 0x5e, 0xeb, 0xc5, 0x5d, 0x98, 0xc7, 0x52, 0xc6,
	0x3c, 0x95, 0x5e, 0xb1, 0xe3, 0x41, 0x01, 0xa9, 0x92, 0x6e, 0x26, 0xb8, 0xe2, 0xa4, 0x7d, 0x65,
	0xb9, 0xc5, 0x4e, 0xef, 0xee, 0x8c, 0xcf, 0xb8, 0x31, 0x3c, 0x7d, 0x2a, 0x33, 0xbd, 0xde, 0xda,
	0xe7, 0x02, 0xb2, 0x5c, 0x81, 0x28, 0xbd, 0xc1, 0x0f, 0x84, 0x3b, 0xbe, 0x6e, 0x38, 0x0e, 0xb8,
	0x00, 0x39, 0x06, 0x45, 0x9e, 0x62, 0xcc, 0x02, 0xc5, 0xc5, 0x44, 0x9d, 0x65, 0x60, 0xa3, 0x3e,
	0x1a, 0x76, 0x76, 0xef, 0xb9, 0xd7, 0xff, 0xe3, 0xee, 0x69, 0xff, 0xf8, 0x2c, 0x03, 0xda, 0x64,
	0xd5, 0x91, 0x74, 0xf1, 0x6d, 0xc5, 0xb3, 0x38, 0x98, 0xc4, 0xa1, 0x5d, 0xeb, 0xa3, 0xe1, 0x06,
	0x6d, 0x18, 0x7d, 0x10, 0x92, 0x07, 0xb8, 0x3d, 0x4d, 0x78, 0x70, 0x32, 0x89, 0x20, 0x9e, 0x45,
	0xca, 0xae, 0xf7, 0xd1, 0xb0, 0x4e, 0x5b, 0xa6, 0xf6, 0xda, 0x94, 0xc8, 0x7d, 0xdc, 0x64, 0x61,
	0x28, 0x40, 0x4a, 0x90, 0xf6, 0x46, 0xbf, 0x3e, 0x6c, 0xd2, 0x55, 0x81, 0x1c, 0xe1, 0x4d, 0x69,
	0x2e, 0x68, 0xdf, 0xd2, 0xd6, 0xe8, 0xd9, 0xf9, 0xe5, 0xb6, 0xf5, 0xed, 0x72, 0xdb, 0x9b, 0xc5,
	0x2a, 0xca, 0xa7, 0x6e, 0xc0, 0xe7, 0x1e, 0x4b, 0x12, 0x2e, 0xd8, 0xa3, 0x14, 0xd4, 0x29, 0x17,
	0x27, 0x95, 0x0c, 0x22, 0x16, 0xa7, 0xde, 0x9c, 0xa9, 0xc8, 0x7d, 0x05, 0x01, 0xfd, 0xdd, 0x66,
	0xf0, 0x13, 0xe1, 0x3b, 0x86, 0x9b, 0xc2, 0x29, 0x13, 0xa1, 0x06, 0x57, 0x09, 0x84, 0x37, 0x12,
	0xfe, 0x1d, 0x6e, 0x88, 0xf2, 0x96, 0xff, 0x4b, 0x5f, 0xf5, 0x19, 0x7c, 0xa9, 0xf0, 0x0f, 0xcb,
	0xfc, 0x1b, 0x2e, 0xcd, 0xdb, 0x5f, 0xc7, 0x40, 0x7f, 0xc7, 0xa8, 0xfd, 0x89, 0xf1, 0x02, 0xb7,
	0x0b, 0x96, 0xe4, 0x30, 0x99, 0xe6, 0x69, 0x98, 0x80, 0x21, 0x6d, 0xed, 0x76, 0xd7, 0xc7, 0xf7,
	0x41, 0x27, 0x46, 0x26, 0x40, 0x5b, 0xc5, 0x4a, 0x3c, 0x7c, 0x82, 0x9b, 0x57, 0xa3, 0x25, 0x2d,
	0xdc, 0x38, 0x38, 0xdc, 0xf7, 0xa9, 0x4f, 0xb7, 0x2c, 0xd2, 0xc1, 0x78, 0xff, 0x88, 0xfa, 0x2f,
	0xf7, 0xc6, 0xc7, 0x3e, 0xdd, 0x42, 0xda, 0xa4, 0xfe, 0xdb, 0xf7, 0x5a, 0xd4, 0x46, 0xf4, 0x7c,
	0xe1, 0xa0, 0x8b, 0x85, 0x83, 0xbe, 0x2f, 0x1c, 0xf4, 0x79, 0xe9, 0x58, 0x17, 0x4b, 0xc7, 0xfa,
	0xba, 0x74, 0xac, 0x8f, 0xcf, 0xff, 0x71, 0x3c, 0x9f, 0xbc, 0xd5, 0x82, 0xe8, 0x67, 0x97, 0xd3,
	0x4d, 0xb3, 0x1c, 0x8f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf0, 0x0d, 0xd2, 0x79, 0x03,
	0x00, 0x00,
}

func (m *EventScoresSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventScoresSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventScoresSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Scores) > 0 {
		for iNdEx := len(m.Scores) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Scores[iNdEx].Size()
				i -= size
				if _, err := m.Scores[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.TopicId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TopicId))
		i--
		dAtA[i] = 0x10
	}
	if m.ActorType != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ActorType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRewardsSettled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardsSettled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardsSettled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Rewards[iNdEx].Size()
				i -= size
				if _, err := m.Rewards[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.TopicId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TopicId))
		i--
		dAtA[i] = 0x10
	}
	if m.ActorType != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ActorType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventNetworkLossSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventNetworkLossSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventNetworkLossSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ValueBundle != nil {
		{
			size, err := m.ValueBundle.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvents(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.TopicId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.TopicId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventScoresSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ActorType != 0 {
		n += 1 + sovEvents(uint64(m.ActorType))
	}
	if m.TopicId != 0 {
		n += 1 + sovEvents(uint64(m.TopicId))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovEvents(uint64(m.BlockHeight))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if len(m.Scores) > 0 {
		for _, e := range m.Scores {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventRewardsSettled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ActorType != 0 {
		n += 1 + sovEvents(uint64(m.ActorType))
	}
	if m.TopicId != 0 {
		n += 1 + sovEvents(uint64(m.TopicId))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovEvents(uint64(m.BlockHeight))
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventNetworkLossSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TopicId != 0 {
		n += 1 + sovEvents(uint64(m.TopicId))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovEvents(uint64(m.BlockHeight))
	}
	if m.ValueBundle != nil {
		l = m.ValueBundle.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventScoresSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventScoresSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventScoresSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActorType", wireType)
			}
			m.ActorType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActorType |= ActorType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicId", wireType)
			}
			m.TopicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scores", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_allora_network_allora_chain_math.Dec
			m.Scores = append(m.Scores, v)
			if err := m.Scores[len(m.Scores)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRewardsSettled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardsSettled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardsSettled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActorType", wireType)
			}
			m.ActorType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActorType |= ActorType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicId", wireType)
			}
			m.TopicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_allora_network_allora_chain_math.Dec
			m.Rewards = append(m.Rewards, v)
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventNetworkLossSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventNetworkLossSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventNetworkLossSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicId", wireType)
			}
			m.TopicId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopicId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValueBundle", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValueBundle == nil {
				m.ValueBundle = &ValueBundle{}
			}
			if err := m.ValueBundle.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
