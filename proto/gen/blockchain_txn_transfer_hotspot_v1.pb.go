// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: blockchain_txn_transfer_hotspot_v1.proto

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

type BlockchainTxnTransferHotspotV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gateway         []byte `protobuf:"bytes,1,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Seller          []byte `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	Buyer           []byte `protobuf:"bytes,3,opt,name=buyer,proto3" json:"buyer,omitempty"`
	SellerSignature []byte `protobuf:"bytes,4,opt,name=seller_signature,json=sellerSignature,proto3" json:"seller_signature,omitempty"`
	BuyerSignature  []byte `protobuf:"bytes,5,opt,name=buyer_signature,json=buyerSignature,proto3" json:"buyer_signature,omitempty"`
	BuyerNonce      uint64 `protobuf:"varint,6,opt,name=buyer_nonce,json=buyerNonce,proto3" json:"buyer_nonce,omitempty"`               // buyer's next payment nonce, required even if payment is 0
	AmountToSeller  uint64 `protobuf:"varint,7,opt,name=amount_to_seller,json=amountToSeller,proto3" json:"amount_to_seller,omitempty"` // in bones not raw HNT
	Fee             uint64 `protobuf:"varint,8,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (x *BlockchainTxnTransferHotspotV1) Reset() {
	*x = BlockchainTxnTransferHotspotV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_txn_transfer_hotspot_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainTxnTransferHotspotV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainTxnTransferHotspotV1) ProtoMessage() {}

func (x *BlockchainTxnTransferHotspotV1) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_txn_transfer_hotspot_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainTxnTransferHotspotV1.ProtoReflect.Descriptor instead.
func (*BlockchainTxnTransferHotspotV1) Descriptor() ([]byte, []int) {
	return file_blockchain_txn_transfer_hotspot_v1_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainTxnTransferHotspotV1) GetGateway() []byte {
	if x != nil {
		return x.Gateway
	}
	return nil
}

func (x *BlockchainTxnTransferHotspotV1) GetSeller() []byte {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *BlockchainTxnTransferHotspotV1) GetBuyer() []byte {
	if x != nil {
		return x.Buyer
	}
	return nil
}

func (x *BlockchainTxnTransferHotspotV1) GetSellerSignature() []byte {
	if x != nil {
		return x.SellerSignature
	}
	return nil
}

func (x *BlockchainTxnTransferHotspotV1) GetBuyerSignature() []byte {
	if x != nil {
		return x.BuyerSignature
	}
	return nil
}

func (x *BlockchainTxnTransferHotspotV1) GetBuyerNonce() uint64 {
	if x != nil {
		return x.BuyerNonce
	}
	return 0
}

func (x *BlockchainTxnTransferHotspotV1) GetAmountToSeller() uint64 {
	if x != nil {
		return x.AmountToSeller
	}
	return 0
}

func (x *BlockchainTxnTransferHotspotV1) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

var File_blockchain_txn_transfer_hotspot_v1_proto protoreflect.FileDescriptor

var file_blockchain_txn_transfer_hotspot_v1_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x78, 0x6e,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x74, 0x73, 0x70, 0x6f,
	0x74, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x65, 0x6c, 0x69,
	0x75, 0x6d, 0x22, 0x9d, 0x02, 0x0a, 0x22, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x74, 0x78, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x68,
	0x6f, 0x74, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x76, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x75, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x62, 0x75, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x75, 0x79, 0x65,
	0x72, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x6f, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66,
	0x65, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x69, 0x6e, 0x64, 0x75, 0x6e, 0x67, 0x57, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x2d, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_blockchain_txn_transfer_hotspot_v1_proto_rawDescOnce sync.Once
	file_blockchain_txn_transfer_hotspot_v1_proto_rawDescData = file_blockchain_txn_transfer_hotspot_v1_proto_rawDesc
)

func file_blockchain_txn_transfer_hotspot_v1_proto_rawDescGZIP() []byte {
	file_blockchain_txn_transfer_hotspot_v1_proto_rawDescOnce.Do(func() {
		file_blockchain_txn_transfer_hotspot_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_txn_transfer_hotspot_v1_proto_rawDescData)
	})
	return file_blockchain_txn_transfer_hotspot_v1_proto_rawDescData
}

var file_blockchain_txn_transfer_hotspot_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_txn_transfer_hotspot_v1_proto_goTypes = []interface{}{
	(*BlockchainTxnTransferHotspotV1)(nil), // 0: helium.blockchain_txn_transfer_hotspot_v1
}
var file_blockchain_txn_transfer_hotspot_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_txn_transfer_hotspot_v1_proto_init() }
func file_blockchain_txn_transfer_hotspot_v1_proto_init() {
	if File_blockchain_txn_transfer_hotspot_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_txn_transfer_hotspot_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainTxnTransferHotspotV1); i {
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
			RawDescriptor: file_blockchain_txn_transfer_hotspot_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_txn_transfer_hotspot_v1_proto_goTypes,
		DependencyIndexes: file_blockchain_txn_transfer_hotspot_v1_proto_depIdxs,
		MessageInfos:      file_blockchain_txn_transfer_hotspot_v1_proto_msgTypes,
	}.Build()
	File_blockchain_txn_transfer_hotspot_v1_proto = out.File
	file_blockchain_txn_transfer_hotspot_v1_proto_rawDesc = nil
	file_blockchain_txn_transfer_hotspot_v1_proto_goTypes = nil
	file_blockchain_txn_transfer_hotspot_v1_proto_depIdxs = nil
}