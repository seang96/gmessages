// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: client.proto

package gmproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

import _ "embed"

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BugleMessageType int32

const (
	BugleMessageType_UNKNOWN_BUGLE_MESSAGE_TYPE BugleMessageType = 0
	BugleMessageType_SMS                        BugleMessageType = 1
	BugleMessageType_MMS                        BugleMessageType = 2
	BugleMessageType_RCS                        BugleMessageType = 3
	BugleMessageType_CLOUD_SYNC                 BugleMessageType = 4
	BugleMessageType_IMDN_DELIVERED             BugleMessageType = 5
	BugleMessageType_IMDN_DISPLAYED             BugleMessageType = 6
	BugleMessageType_IMDN_FALLBACK              BugleMessageType = 7
	BugleMessageType_RCS_GENERIC                BugleMessageType = 8
	BugleMessageType_FTD                        BugleMessageType = 9
	BugleMessageType_FT_E2EE_LEGACY             BugleMessageType = 10
	BugleMessageType_FT_E2EE_XML                BugleMessageType = 11
	BugleMessageType_LIGHTER_MESSAGE            BugleMessageType = 12
	BugleMessageType_RBM_SPAM_REPORT            BugleMessageType = 13
	BugleMessageType_SATELLITE                  BugleMessageType = 14
)

// Enum value maps for BugleMessageType.
var (
	BugleMessageType_name = map[int32]string{
		0:  "UNKNOWN_BUGLE_MESSAGE_TYPE",
		1:  "SMS",
		2:  "MMS",
		3:  "RCS",
		4:  "CLOUD_SYNC",
		5:  "IMDN_DELIVERED",
		6:  "IMDN_DISPLAYED",
		7:  "IMDN_FALLBACK",
		8:  "RCS_GENERIC",
		9:  "FTD",
		10: "FT_E2EE_LEGACY",
		11: "FT_E2EE_XML",
		12: "LIGHTER_MESSAGE",
		13: "RBM_SPAM_REPORT",
		14: "SATELLITE",
	}
	BugleMessageType_value = map[string]int32{
		"UNKNOWN_BUGLE_MESSAGE_TYPE": 0,
		"SMS":                        1,
		"MMS":                        2,
		"RCS":                        3,
		"CLOUD_SYNC":                 4,
		"IMDN_DELIVERED":             5,
		"IMDN_DISPLAYED":             6,
		"IMDN_FALLBACK":              7,
		"RCS_GENERIC":                8,
		"FTD":                        9,
		"FT_E2EE_LEGACY":             10,
		"FT_E2EE_XML":                11,
		"LIGHTER_MESSAGE":            12,
		"RBM_SPAM_REPORT":            13,
		"SATELLITE":                  14,
	}
)

func (x BugleMessageType) Enum() *BugleMessageType {
	p := new(BugleMessageType)
	*p = x
	return p
}

func (x BugleMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BugleMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_client_proto_enumTypes[0].Descriptor()
}

func (BugleMessageType) Type() protoreflect.EnumType {
	return &file_client_proto_enumTypes[0]
}

func (x BugleMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BugleMessageType.Descriptor instead.
func (BugleMessageType) EnumDescriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

type BrowserTypes int32

const (
	BrowserTypes_UNKNOWN_BROWSER_TYPE BrowserTypes = 0
	BrowserTypes_OTHER                BrowserTypes = 1
	BrowserTypes_CHROME               BrowserTypes = 2
	BrowserTypes_FIREFOX              BrowserTypes = 3
	BrowserTypes_SAFARI               BrowserTypes = 4
	BrowserTypes_OPERA                BrowserTypes = 5
	BrowserTypes_IE                   BrowserTypes = 6
	BrowserTypes_EDGE                 BrowserTypes = 7
)

// Enum value maps for BrowserTypes.
var (
	BrowserTypes_name = map[int32]string{
		0: "UNKNOWN_BROWSER_TYPE",
		1: "OTHER",
		2: "CHROME",
		3: "FIREFOX",
		4: "SAFARI",
		5: "OPERA",
		6: "IE",
		7: "EDGE",
	}
	BrowserTypes_value = map[string]int32{
		"UNKNOWN_BROWSER_TYPE": 0,
		"OTHER":                1,
		"CHROME":               2,
		"FIREFOX":              3,
		"SAFARI":               4,
		"OPERA":                5,
		"IE":                   6,
		"EDGE":                 7,
	}
)

func (x BrowserTypes) Enum() *BrowserTypes {
	p := new(BrowserTypes)
	*p = x
	return p
}

func (x BrowserTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrowserTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_client_proto_enumTypes[1].Descriptor()
}

func (BrowserTypes) Type() protoreflect.EnumType {
	return &file_client_proto_enumTypes[1]
}

func (x BrowserTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrowserTypes.Descriptor instead.
func (BrowserTypes) EnumDescriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

type NotifyDittoActivityPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is not actually a boolean: after logging out, field 2 has value 2, and field 3 has value 1.
	Success bool `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NotifyDittoActivityPayload) Reset() {
	*x = NotifyDittoActivityPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyDittoActivityPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyDittoActivityPayload) ProtoMessage() {}

func (x *NotifyDittoActivityPayload) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyDittoActivityPayload.ProtoReflect.Descriptor instead.
func (*NotifyDittoActivityPayload) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyDittoActivityPayload) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MessageReadPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConversationID string `protobuf:"bytes,2,opt,name=conversationID,proto3" json:"conversationID,omitempty"`
	MessageID      string `protobuf:"bytes,3,opt,name=messageID,proto3" json:"messageID,omitempty"`
}

func (x *MessageReadPayload) Reset() {
	*x = MessageReadPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageReadPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageReadPayload) ProtoMessage() {}

func (x *MessageReadPayload) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageReadPayload.ProtoReflect.Descriptor instead.
func (*MessageReadPayload) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *MessageReadPayload) GetConversationID() string {
	if x != nil {
		return x.ConversationID
	}
	return ""
}

func (x *MessageReadPayload) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

type AckMessagePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthData *AuthMessage      `protobuf:"bytes,1,opt,name=authData,proto3" json:"authData,omitempty"`
	EmptyArr *EmptyArr         `protobuf:"bytes,2,opt,name=emptyArr,proto3" json:"emptyArr,omitempty"`
	Acks     []*AckMessageData `protobuf:"bytes,4,rep,name=acks,proto3" json:"acks,omitempty"`
}

func (x *AckMessagePayload) Reset() {
	*x = AckMessagePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMessagePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessagePayload) ProtoMessage() {}

func (x *AckMessagePayload) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessagePayload.ProtoReflect.Descriptor instead.
func (*AckMessagePayload) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *AckMessagePayload) GetAuthData() *AuthMessage {
	if x != nil {
		return x.AuthData
	}
	return nil
}

func (x *AckMessagePayload) GetEmptyArr() *EmptyArr {
	if x != nil {
		return x.EmptyArr
	}
	return nil
}

func (x *AckMessagePayload) GetAcks() []*AckMessageData {
	if x != nil {
		return x.Acks
	}
	return nil
}

type AckMessageData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestID string  `protobuf:"bytes,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Device    *Device `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *AckMessageData) Reset() {
	*x = AckMessageData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckMessageData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessageData) ProtoMessage() {}

func (x *AckMessageData) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessageData.ProtoReflect.Descriptor instead.
func (*AckMessageData) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{3}
}

func (x *AckMessageData) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *AckMessageData) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type ImageMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageID   string `protobuf:"bytes,1,opt,name=imageID,proto3" json:"imageID,omitempty"`
	Encrypted bool   `protobuf:"varint,2,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *ImageMetaData) Reset() {
	*x = ImageMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageMetaData) ProtoMessage() {}

func (x *ImageMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageMetaData.ProtoReflect.Descriptor instead.
func (*ImageMetaData) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{4}
}

func (x *ImageMetaData) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

func (x *ImageMetaData) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

type UploadImagePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetaData *ImageMetaData `protobuf:"bytes,1,opt,name=metaData,proto3" json:"metaData,omitempty"`
	AuthData *AuthMessage   `protobuf:"bytes,2,opt,name=authData,proto3" json:"authData,omitempty"`
}

func (x *UploadImagePayload) Reset() {
	*x = UploadImagePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImagePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImagePayload) ProtoMessage() {}

func (x *UploadImagePayload) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImagePayload.ProtoReflect.Descriptor instead.
func (*UploadImagePayload) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{5}
}

func (x *UploadImagePayload) GetMetaData() *ImageMetaData {
	if x != nil {
		return x.MetaData
	}
	return nil
}

func (x *UploadImagePayload) GetAuthData() *AuthMessage {
	if x != nil {
		return x.AuthData
	}
	return nil
}

type BugleBackendService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *BugleCode `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BugleBackendService) Reset() {
	*x = BugleBackendService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BugleBackendService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BugleBackendService) ProtoMessage() {}

func (x *BugleBackendService) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BugleBackendService.ProtoReflect.Descriptor instead.
func (*BugleBackendService) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{6}
}

func (x *BugleBackendService) GetData() *BugleCode {
	if x != nil {
		return x.Data
	}
	return nil
}

type BugleCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int64 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *BugleCode) Reset() {
	*x = BugleCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BugleCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BugleCode) ProtoMessage() {}

func (x *BugleCode) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BugleCode.ProtoReflect.Descriptor instead.
func (*BugleCode) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{7}
}

func (x *BugleCode) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_client_proto protoreflect.FileDescriptor

//go:embed client.pb.raw
var file_client_proto_rawDesc []byte

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_client_proto_goTypes = []interface{}{
	(BugleMessageType)(0),              // 0: client.BugleMessageType
	(BrowserTypes)(0),                  // 1: client.BrowserTypes
	(*NotifyDittoActivityPayload)(nil), // 2: client.NotifyDittoActivityPayload
	(*MessageReadPayload)(nil),         // 3: client.MessageReadPayload
	(*AckMessagePayload)(nil),          // 4: client.AckMessagePayload
	(*AckMessageData)(nil),             // 5: client.AckMessageData
	(*ImageMetaData)(nil),              // 6: client.ImageMetaData
	(*UploadImagePayload)(nil),         // 7: client.UploadImagePayload
	(*BugleBackendService)(nil),        // 8: client.BugleBackendService
	(*BugleCode)(nil),                  // 9: client.BugleCode
	(*AuthMessage)(nil),                // 10: messages.AuthMessage
	(*EmptyArr)(nil),                   // 11: messages.EmptyArr
	(*Device)(nil),                     // 12: messages.Device
}
var file_client_proto_depIdxs = []int32{
	10, // 0: client.AckMessagePayload.authData:type_name -> messages.AuthMessage
	11, // 1: client.AckMessagePayload.emptyArr:type_name -> messages.EmptyArr
	5,  // 2: client.AckMessagePayload.acks:type_name -> client.AckMessageData
	12, // 3: client.AckMessageData.device:type_name -> messages.Device
	6,  // 4: client.UploadImagePayload.metaData:type_name -> client.ImageMetaData
	10, // 5: client.UploadImagePayload.authData:type_name -> messages.AuthMessage
	9,  // 6: client.BugleBackendService.data:type_name -> client.BugleCode
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	file_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyDittoActivityPayload); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageReadPayload); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckMessagePayload); i {
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
		file_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckMessageData); i {
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
		file_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageMetaData); i {
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
		file_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImagePayload); i {
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
		file_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BugleBackendService); i {
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
		file_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BugleCode); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		EnumInfos:         file_client_proto_enumTypes,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}