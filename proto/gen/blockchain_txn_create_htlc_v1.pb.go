// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: blockchain_txn_create_htlc_v1.proto

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

type BlockchainTxnCreateHtlcV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payer     []byte `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Payee     []byte `protobuf:"bytes,2,opt,name=payee,proto3" json:"payee,omitempty"`
	Address   []byte `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Hashlock  []byte `protobuf:"bytes,4,opt,name=hashlock,proto3" json:"hashlock,omitempty"`
	Timelock  uint64 `protobuf:"varint,5,opt,name=timelock,proto3" json:"timelock,omitempty"`
	Amount    uint64 `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee       uint64 `protobuf:"varint,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Signature []byte `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce     uint64 `protobuf:"varint,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *BlockchainTxnCreateHtlcV1) Reset() {
	*x = BlockchainTxnCreateHtlcV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_create_htlc_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnCreateHtlcV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnCreateHtlcV1) ProtoMessage() {}

func (x *BlockchainTxnCreateHtlcV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_create_htlc_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnCreateHtlcV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnCreateHtlcV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_create_htlc_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnCreateHtlcV1) GetPayer() []byte {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *BlockchainTxnCreateHtlcV1) GetPayee() []byte {
	if x != nil {
		return x.Payee
	}
	return nil
}

func (x *BlockchainTxnCreateHtlcV1) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *BlockchainTxnCreateHtlcV1) GetHashlock() []byte {
	if x != nil {
		return x.Hashlock
	}
	return nil
}

func (x *BlockchainTxnCreateHtlcV1) GetTimelock() uint64 {
	if x != nil {
		return x.Timelock
	}
	return 0
}

func (x *BlockchainTxnCreateHtlcV1) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BlockchainTxnCreateHtlcV1) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *BlockchainTxnCreateHtlcV1) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BlockchainTxnCreateHtlcV1) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

var File_blockchain_txn_create_htlc_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_create_htlc_v1_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x74, 0x6c, 0x63, 0x5f, 0x76, 0x31, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69, 0x75, 0x6d, 0x22, 0xfb, 0x01,
	0x0a, 0x1d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x74, 0x6c, 0x63, 0x5f, 0x76, 0x31, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6e, 0x64, 0x75, 0x6e,
	0x67, 0x57, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2d, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_txn_create_htlc_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_create_htlc_v1_proto_rawDescData = file_blockchain_txn_create_htlc_v1_proto_rawDesc
)

func file_blockchain_txn_create_htlc_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_create_htlc_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_create_htlc_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_create_htlc_v1_proto_rawDescData)
	})
	return file_blockchain_txn_create_htlc_v1_proto_rawDescData
}

var file_blockchain_txn_create_htlc_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_txn_create_htlc_v1_proto_goTypes = []interface{}{
	(*BlockchainTxnCreateHtlcV1)(nil), // 0: helium.blockchain_txn_create_htlc_v1
}
var file_blockchain_txn_create_htlc_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_txn_create_htlc_v1_proto_init() }
func file_blockchain_txn_create_htlc_v1_proto_init() {
	if File_blockchain_txn_create_htlc_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_create_htlc_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnCreateHtlcV1); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blockchain_txn_create_htlc_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_create_htlc_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_create_htlc_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_create_htlc_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_create_htlc_v1_proto = out.File
	file_blockchain_txn_create_htlc_v1_proto_rawDesc = nil
	file_blockchain_txn_create_htlc_v1_proto_goTypes = nil
	file_blockchain_txn_create_htlc_v1_proto_depIdxs = nil
}
