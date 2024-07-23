// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: proto/bank/type/transfer.proto

package bank

import (
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

type TransferStatus int32

const (
	TransferStatus_TRANSFER_STATUS_UNSPECIFIED TransferStatus = 0
	TransferStatus_TRANSFER_STATUS_SUCCESS     TransferStatus = 1
	TransferStatus_TRANSFER_STATUS_FAILED      TransferStatus = 2
)

// Enum value maps for TransferStatus.
var (
	TransferStatus_name = map[int32]string{
		0: "TRANSFER_STATUS_UNSPECIFIED",
		1: "TRANSFER_STATUS_SUCCESS",
		2: "TRANSFER_STATUS_FAILED",
	}
	TransferStatus_value = map[string]int32{
		"TRANSFER_STATUS_UNSPECIFIED": 0,
		"TRANSFER_STATUS_SUCCESS":     1,
		"TRANSFER_STATUS_FAILED":      2,
	}
)

func (x TransferStatus) Enum() *TransferStatus {
	p := new(TransferStatus)
	*p = x
	return p
}

func (x TransferStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bank_type_transfer_proto_enumTypes[0].Descriptor()
}

func (TransferStatus) Type() protoreflect.EnumType {
	return &file_proto_bank_type_transfer_proto_enumTypes[0]
}

func (x TransferStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferStatus.Descriptor instead.
func (TransferStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_bank_type_transfer_proto_rawDescGZIP(), []int{0}
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountNumber string  `protobuf:"bytes,1,opt,name=from_account_number,proto3" json:"from_account_number,omitempty"`
	ToAccountNumber   string  `protobuf:"bytes,2,opt,name=to_account_number,proto3" json:"to_account_number,omitempty"`
	Currency          string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount            float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_type_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetFromAccountNumber() string {
	if x != nil {
		return x.FromAccountNumber
	}
	return ""
}

func (x *TransferRequest) GetToAccountNumber() string {
	if x != nil {
		return x.ToAccountNumber
	}
	return ""
}

func (x *TransferRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccountNumber string             `protobuf:"bytes,1,opt,name=from_account_number,proto3" json:"from_account_number,omitempty"`
	ToAccountNumber   string             `protobuf:"bytes,2,opt,name=to_account_number,proto3" json:"to_account_number,omitempty"`
	Currency          string             `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount            float64            `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Status            TransferStatus     `protobuf:"varint,5,opt,name=status,proto3,enum=bank.TransferStatus" json:"status,omitempty"`
	Timestamp         *datetime.DateTime `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_type_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TransferResponse) GetFromAccountNumber() string {
	if x != nil {
		return x.FromAccountNumber
	}
	return ""
}

func (x *TransferResponse) GetToAccountNumber() string {
	if x != nil {
		return x.ToAccountNumber
	}
	return ""
}

func (x *TransferResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferResponse) GetStatus() TransferStatus {
	if x != nil {
		return x.Status
	}
	return TransferStatus_TRANSFER_STATUS_UNSPECIFIED
}

func (x *TransferResponse) GetTimestamp() *datetime.DateTime {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_proto_bank_type_transfer_proto protoreflect.FileDescriptor

var file_proto_bank_type_transfer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x89, 0x02, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x6f, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x74, 0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x6e, 0x6b,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x6a, 0x0a, 0x0e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x1b, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6d, 0x70, 0x61, 0x6d, 0x75, 0x6e, 0x67,
	0x6b, 0x61, 0x73, 0x2f, 0x6d, 0x79, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bank_type_transfer_proto_rawDescOnce sync.Once
	file_proto_bank_type_transfer_proto_rawDescData = file_proto_bank_type_transfer_proto_rawDesc
)

func file_proto_bank_type_transfer_proto_rawDescGZIP() []byte {
	file_proto_bank_type_transfer_proto_rawDescOnce.Do(func() {
		file_proto_bank_type_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bank_type_transfer_proto_rawDescData)
	})
	return file_proto_bank_type_transfer_proto_rawDescData
}

var file_proto_bank_type_transfer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_bank_type_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_bank_type_transfer_proto_goTypes = []interface{}{
	(TransferStatus)(0),       // 0: bank.TransferStatus
	(*TransferRequest)(nil),   // 1: bank.TransferRequest
	(*TransferResponse)(nil),  // 2: bank.TransferResponse
	(*datetime.DateTime)(nil), // 3: google.type.DateTime
}
var file_proto_bank_type_transfer_proto_depIdxs = []int32{
	0, // 0: bank.TransferResponse.status:type_name -> bank.TransferStatus
	3, // 1: bank.TransferResponse.timestamp:type_name -> google.type.DateTime
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_bank_type_transfer_proto_init() }
func file_proto_bank_type_transfer_proto_init() {
	if File_proto_bank_type_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bank_type_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferRequest); i {
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
		file_proto_bank_type_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResponse); i {
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
			RawDescriptor: file_proto_bank_type_transfer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bank_type_transfer_proto_goTypes,
		DependencyIndexes: file_proto_bank_type_transfer_proto_depIdxs,
		EnumInfos:         file_proto_bank_type_transfer_proto_enumTypes,
		MessageInfos:      file_proto_bank_type_transfer_proto_msgTypes,
	}.Build()
	File_proto_bank_type_transfer_proto = out.File
	file_proto_bank_type_transfer_proto_rawDesc = nil
	file_proto_bank_type_transfer_proto_goTypes = nil
	file_proto_bank_type_transfer_proto_depIdxs = nil
}
