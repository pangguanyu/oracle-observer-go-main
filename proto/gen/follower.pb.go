// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: follower.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GatewayNotFound struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GatewayNotFound) Reset() {
	*x = GatewayNotFound{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayNotFound) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayNotFound) ProtoMessage() {}

func (x *GatewayNotFound) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayNotFound.ProtoReflect.Descriptor instead.
func (*GatewayNotFound) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayNotFound) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type FollowerError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*FollowerError_NotFound
	Type isFollowerError_Type `protobuf_oneof:"type"`
}

func (x *FollowerError) Reset() {
	*x = FollowerError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerError) ProtoMessage() {}

func (x *FollowerError) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerError.ProtoReflect.Descriptor instead.
func (*FollowerError) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{1}
}

func (m *FollowerError) GetType() isFollowerError_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *FollowerError) GetNotFound() *GatewayNotFound {
	if x, ok := x.GetType().(*FollowerError_NotFound); ok {
		return x.NotFound
	}
	return nil
}

type isFollowerError_Type interface {
	isFollowerError_Type()
}

type FollowerError_NotFound struct {
	NotFound *GatewayNotFound `protobuf:"bytes,1,opt,name=not_found,json=notFound,proto3,oneof"`
}

func (*FollowerError_NotFound) isFollowerError_Type() {}

type FollowerTxnStreamReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height   uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	TxnHash  []byte   `protobuf:"bytes,2,opt,name=txn_hash,json=txnHash,proto3" json:"txn_hash,omitempty"`
	TxnTypes []string `protobuf:"bytes,3,rep,name=txn_types,json=txnTypes,proto3" json:"txn_types,omitempty"`
}

func (x *FollowerTxnStreamReqV1) Reset() {
	*x = FollowerTxnStreamReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerTxnStreamReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerTxnStreamReqV1) ProtoMessage() {}

func (x *FollowerTxnStreamReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerTxnStreamReqV1.ProtoReflect.Descriptor instead.
func (*FollowerTxnStreamReqV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{2}
}

func (x *FollowerTxnStreamReqV1) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *FollowerTxnStreamReqV1) GetTxnHash() []byte {
	if x != nil {
		return x.TxnHash
	}
	return nil
}

func (x *FollowerTxnStreamReqV1) GetTxnTypes() []string {
	if x != nil {
		return x.TxnTypes
	}
	return nil
}

type FollowerTxnStreamRespV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    uint64         `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	TxnHash   []byte         `protobuf:"bytes,2,opt,name=txn_hash,json=txnHash,proto3" json:"txn_hash,omitempty"`
	Txn       *BlockchainTxn `protobuf:"bytes,3,opt,name=txn,proto3" json:"txn,omitempty"`
	Timestamp uint64         `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *FollowerTxnStreamRespV1) Reset() {
	*x = FollowerTxnStreamRespV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerTxnStreamRespV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerTxnStreamRespV1) ProtoMessage() {}

func (x *FollowerTxnStreamRespV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerTxnStreamRespV1.ProtoReflect.Descriptor instead.
func (*FollowerTxnStreamRespV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{3}
}

func (x *FollowerTxnStreamRespV1) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *FollowerTxnStreamRespV1) GetTxnHash() []byte {
	if x != nil {
		return x.TxnHash
	}
	return nil
}

func (x *FollowerTxnStreamRespV1) GetTxn() *BlockchainTxn {
	if x != nil {
		return x.Txn
	}
	return nil
}

func (x *FollowerTxnStreamRespV1) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}



// / Look up the owner of a given hotspot public key
type FollowerGatewayReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / The pubkey_bin address of the gateway to look up
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *FollowerGatewayReqV1) Reset() {
	*x = FollowerGatewayReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerGatewayReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerGatewayReqV1) ProtoMessage() {}

func (x *FollowerGatewayReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerGatewayReqV1.ProtoReflect.Descriptor instead.
func (*FollowerGatewayReqV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{5}
}

func (x *FollowerGatewayReqV1) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type FollowerGatewayRespV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / The height for at which the ownership was looked up
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// Types that are assignable to Result:
	//
	//	*FollowerGatewayRespV1_Info
	//	*FollowerGatewayRespV1_Error
	Result isFollowerGatewayRespV1_Result `protobuf_oneof:"result"`
}

func (x *FollowerGatewayRespV1) Reset() {
	*x = FollowerGatewayRespV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerGatewayRespV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerGatewayRespV1) ProtoMessage() {}

func (x *FollowerGatewayRespV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerGatewayRespV1.ProtoReflect.Descriptor instead.
func (*FollowerGatewayRespV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{6}
}

func (x *FollowerGatewayRespV1) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (m *FollowerGatewayRespV1) GetResult() isFollowerGatewayRespV1_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *FollowerGatewayRespV1) GetInfo() *GatewayInfo {
	if x, ok := x.GetResult().(*FollowerGatewayRespV1_Info); ok {
		return x.Info
	}
	return nil
}

func (x *FollowerGatewayRespV1) GetError() *FollowerError {
	if x, ok := x.GetResult().(*FollowerGatewayRespV1_Error); ok {
		return x.Error
	}
	return nil
}

type isFollowerGatewayRespV1_Result interface {
	isFollowerGatewayRespV1_Result()
}

type FollowerGatewayRespV1_Info struct {
	Info *GatewayInfo `protobuf:"bytes,2,opt,name=info,proto3,oneof"`
}

type FollowerGatewayRespV1_Error struct {
	Error *FollowerError `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*FollowerGatewayRespV1_Info) isFollowerGatewayRespV1_Result() {}

func (*FollowerGatewayRespV1_Error) isFollowerGatewayRespV1_Result() {}

// / Request a stream of all active gateways from the on-chain metadata
type FollowerGatewayStreamReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchSize uint32 `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *FollowerGatewayStreamReqV1) Reset() {
	*x = FollowerGatewayStreamReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerGatewayStreamReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerGatewayStreamReqV1) ProtoMessage() {}

func (x *FollowerGatewayStreamReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerGatewayStreamReqV1.ProtoReflect.Descriptor instead.
func (*FollowerGatewayStreamReqV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{7}
}

func (x *FollowerGatewayStreamReqV1) GetBatchSize() uint32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

// / Active gateway info streaming response containing a batch of gateways
type FollowerGatewayStreamRespV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateways []*FollowerGatewayRespV1 `protobuf:"bytes,1,rep,name=gateways,proto3" json:"gateways,omitempty"`
}

func (x *FollowerGatewayStreamRespV1) Reset() {
	*x = FollowerGatewayStreamRespV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerGatewayStreamRespV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerGatewayStreamRespV1) ProtoMessage() {}

func (x *FollowerGatewayStreamRespV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerGatewayStreamRespV1.ProtoReflect.Descriptor instead.
func (*FollowerGatewayStreamRespV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{8}
}

func (x *FollowerGatewayStreamRespV1) GetGateways() []*FollowerGatewayRespV1 {
	if x != nil {
		return x.Gateways
	}
	return nil
}

// / Query the last reward block for the given subnetwork by token type
type FollowerSubnetworkLastRewardHeightReqV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / The token type of the subnetwork to query
	TokenType BlockchainTokenTypeV1 `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=helium.BlockchainTokenTypeV1" json:"token_type,omitempty"`
}

func (x *FollowerSubnetworkLastRewardHeightReqV1) Reset() {
	*x = FollowerSubnetworkLastRewardHeightReqV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerSubnetworkLastRewardHeightReqV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerSubnetworkLastRewardHeightReqV1) ProtoMessage() {}

func (x *FollowerSubnetworkLastRewardHeightReqV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerSubnetworkLastRewardHeightReqV1.ProtoReflect.Descriptor instead.
func (*FollowerSubnetworkLastRewardHeightReqV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{9}
}

func (x *FollowerSubnetworkLastRewardHeightReqV1) GetTokenType() BlockchainTokenTypeV1 {
	if x != nil {
		return x.TokenType
	}
	return BlockchainTokenTypeV1_hnt
}

type FollowerSubnetworkLastRewardHeightRespV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// / The current height at the time of the request
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// / The height of the reward block
	RewardHeight uint64 `protobuf:"varint,2,opt,name=reward_height,json=rewardHeight,proto3" json:"reward_height,omitempty"`
}

func (x *FollowerSubnetworkLastRewardHeightRespV1) Reset() {
	*x = FollowerSubnetworkLastRewardHeightRespV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowerSubnetworkLastRewardHeightRespV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowerSubnetworkLastRewardHeightRespV1) ProtoMessage() {}

func (x *FollowerSubnetworkLastRewardHeightRespV1) ProtoReflect() protoreflect.Message {
	mi := &file_follower_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowerSubnetworkLastRewardHeightRespV1.ProtoReflect.Descriptor instead.
func (*FollowerSubnetworkLastRewardHeightRespV1) Descriptor() ([]byte, []int) {
	return file_follower_proto_rawDescGZIP(), []int{10}
}

func (x *FollowerSubnetworkLastRewardHeightRespV1) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *FollowerSubnetworkLastRewardHeightRespV1) GetRewardHeight() uint64 {
	if x != nil {
		return x.RewardHeight
	}
	return 0
}

var File_follower_proto protoreflect.FileDescriptor

var file_follower_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x1a, 0x1e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x11, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x5b, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x6e,
	0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x6c, 0x0a, 0x1a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x78, 0x6e, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x78, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x78, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x98, 0x01,
	0x0a, 0x1b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6e, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x28, 0x0a, 0x03, 0x74, 0x78, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x52, 0x03, 0x74, 0x78, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa1, 0x02, 0x0a, 0x0c, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x65,
	0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x67, 0x61, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x68, 0x65, 0x6c, 0x69,
	0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x5f, 0x76, 0x31, 0x52, 0x0c,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x33, 0x0a, 0x17,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0xaa, 0x01, 0x0a, 0x18, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3f,
	0x0a, 0x1e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x68, 0x0a, 0x1f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f,
	0x76, 0x31, 0x12, 0x45, 0x0a, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x52,
	0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x22, 0x70, 0x0a, 0x2d, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x12, 0x3f, 0x0a, 0x0a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x31,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6d, 0x0a, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x32, 0xf5, 0x03, 0x0a, 0x08, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x0a, 0x74, 0x78, 0x6e, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x2b, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f,
	0x76, 0x31, 0x1a, 0x2c, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x78,
	0x6e, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31,
	0x30, 0x01, 0x12, 0x63, 0x0a, 0x0c, 0x66, 0x69, 0x6e, 0x64, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x28, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x1a, 0x29, 0x2e, 0x68,
	0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x12, 0x76, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x73, 0x12, 0x2f, 0x2e, 0x68, 0x65, 0x6c,
	0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76, 0x31, 0x1a, 0x30, 0x2e, 0x68, 0x65,
	0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x76, 0x31, 0x30, 0x01, 0x12,
	0xa0, 0x01, 0x0a, 0x1d, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x3e, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x76,
	0x31, 0x1a, 0x3f, 0x2e, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f,
	0x76, 0x31, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x69, 0x6e, 0x64, 0x75, 0x6e, 0x67, 0x57, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_follower_proto_rawDescOnce sync.Once
	file_follower_proto_rawDescData = file_follower_proto_rawDesc
)

func file_follower_proto_rawDescGZIP() []byte {
	file_follower_proto_rawDescOnce.Do(func() {
		file_follower_proto_rawDescData = protoimpl.X.CompressGZIP(file_follower_proto_rawDescData)
	})
	return file_follower_proto_rawDescData
}

var file_follower_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_follower_proto_goTypes = []interface{}{
	(*GatewayNotFound)(nil),                          // 0: helium.follower.gateway_not_found
	(*FollowerError)(nil),                            // 1: helium.follower.follower_error
	(*FollowerTxnStreamReqV1)(nil),                   // 2: helium.follower.follower_txn_stream_req_v1
	(*FollowerTxnStreamRespV1)(nil),                  // 3: helium.follower.follower_txn_stream_resp_v1
	(*GatewayInfo)(nil),                              // 4: helium.follower.gateway_info
	(*FollowerGatewayReqV1)(nil),                     // 5: helium.follower.follower_gateway_req_v1
	(*FollowerGatewayRespV1)(nil),                    // 6: helium.follower.follower_gateway_resp_v1
	(*FollowerGatewayStreamReqV1)(nil),               // 7: helium.follower.follower_gateway_stream_req_v1
	(*FollowerGatewayStreamRespV1)(nil),              // 8: helium.follower.follower_gateway_stream_resp_v1
	(*FollowerSubnetworkLastRewardHeightReqV1)(nil),  // 9: helium.follower.follower_subnetwork_last_reward_height_req_v1
	(*FollowerSubnetworkLastRewardHeightRespV1)(nil), // 10: helium.follower.follower_subnetwork_last_reward_height_resp_v1
	(*BlockchainTxn)(nil),                            // 11: helium.blockchain_txn
	(GatewayStakingMode)(0),                          // 12: helium.gateway_staking_mode
	(Region)(0),                                      // 13: helium.region
	(*BlockchainRegionParamsV1)(nil),                 // 14: helium.blockchain_region_params_v1
	(BlockchainTokenTypeV1)(0),                       // 15: helium.blockchain_token_type_v1
}
var file_follower_proto_depIdxs = []int32{
	0,  // 0: helium.follower.follower_error.not_found:type_name -> helium.follower.gateway_not_found
	11, // 1: helium.follower.follower_txn_stream_resp_v1.txn:type_name -> helium.blockchain_txn
	12, // 2: helium.follower.gateway_info.staking_mode:type_name -> helium.gateway_staking_mode
	13, // 3: helium.follower.gateway_info.region:type_name -> helium.region
	14, // 4: helium.follower.gateway_info.region_params:type_name -> helium.blockchain_region_params_v1
	4,  // 5: helium.follower.follower_gateway_resp_v1.info:type_name -> helium.follower.gateway_info
	1,  // 6: helium.follower.follower_gateway_resp_v1.error:type_name -> helium.follower.follower_error
	6,  // 7: helium.follower.follower_gateway_stream_resp_v1.gateways:type_name -> helium.follower.follower_gateway_resp_v1
	15, // 8: helium.follower.follower_subnetwork_last_reward_height_req_v1.token_type:type_name -> helium.blockchain_token_type_v1
	2,  // 9: helium.follower.follower.txn_stream:input_type -> helium.follower.follower_txn_stream_req_v1
	5,  // 10: helium.follower.follower.find_gateway:input_type -> helium.follower.follower_gateway_req_v1
	7,  // 11: helium.follower.follower.active_gateways:input_type -> helium.follower.follower_gateway_stream_req_v1
	9,  // 12: helium.follower.follower.subnetwork_last_reward_height:input_type -> helium.follower.follower_subnetwork_last_reward_height_req_v1
	3,  // 13: helium.follower.follower.txn_stream:output_type -> helium.follower.follower_txn_stream_resp_v1
	6,  // 14: helium.follower.follower.find_gateway:output_type -> helium.follower.follower_gateway_resp_v1
	8,  // 15: helium.follower.follower.active_gateways:output_type -> helium.follower.follower_gateway_stream_resp_v1
	10, // 16: helium.follower.follower.subnetwork_last_reward_height:output_type -> helium.follower.follower_subnetwork_last_reward_height_resp_v1
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_follower_proto_init() }
func file_follower_proto_init() {
	if File_follower_proto != nil {
		return
	}
	file_blockchain_token_type_v1_proto_init()
	file_blockchain_txn_proto_init()
	file_gateway_staking_mode_proto_init()
	file_region_proto_init()
	file_blockchain_region_param_v1_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_follower_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayNotFound); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerTxnStreamReqV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerTxnStreamRespV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerGatewayReqV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerGatewayRespV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerGatewayStreamReqV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerGatewayStreamRespV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerSubnetworkLastRewardHeightReqV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_follower_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowerSubnetworkLastRewardHeightRespV1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_follower_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FollowerError_NotFound)(nil),
	}
	file_follower_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*FollowerGatewayRespV1_Info)(nil),
		(*FollowerGatewayRespV1_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_follower_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follower_proto_goTypes,
		DependencyIndexes: file_follower_proto_depIdxs,
		MessageInfos:      file_follower_proto_msgTypes,
	}.Build()
	File_follower_proto = out.File
	file_follower_proto_rawDesc = nil
	file_follower_proto_goTypes = nil
	file_follower_proto_depIdxs = nil
}