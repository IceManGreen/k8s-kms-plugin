//
//Copyright 2018 The Kubernetes Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.0
// source: apis/k8s/v1/messages.proto

package k8s

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be decrypted.
	Cipher []byte `protobuf:"bytes,2,opt,name=cipher,proto3" json:"cipher,omitempty"`
	// Required. The Keyring that is holding the key to use
	KeyringId string `protobuf:"bytes,3,opt,name=keyring_id,json=keyringId,proto3" json:"keyring_id,omitempty"`
	// Required. The Key to use
	KeyId string `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_k8s_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_k8s_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_apis_k8s_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *DecryptRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *DecryptRequest) GetCipher() []byte {
	if x != nil {
		return x.Cipher
	}
	return nil
}

func (x *DecryptRequest) GetKeyringId() string {
	if x != nil {
		return x.KeyringId
	}
	return ""
}

func (x *DecryptRequest) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type DecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The decrypted data.
	Plain []byte `protobuf:"bytes,1,opt,name=plain,proto3" json:"plain,omitempty"`
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_k8s_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_k8s_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_apis_k8s_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *DecryptResponse) GetPlain() []byte {
	if x != nil {
		return x.Plain
	}
	return nil
}

type EncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of the KMS plugin API.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The data to be encrypted.
	Plain []byte `protobuf:"bytes,2,opt,name=plain,proto3" json:"plain,omitempty"`
	// Required. The Keyring that is holding the key to use
	KeyringId string `protobuf:"bytes,3,opt,name=keyring_id,json=keyringId,proto3" json:"keyring_id,omitempty"`
	// Required. The Key to use
	KeyId string `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
}

func (x *EncryptRequest) Reset() {
	*x = EncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_k8s_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptRequest) ProtoMessage() {}

func (x *EncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_k8s_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return file_apis_k8s_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *EncryptRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *EncryptRequest) GetPlain() []byte {
	if x != nil {
		return x.Plain
	}
	return nil
}

func (x *EncryptRequest) GetKeyringId() string {
	if x != nil {
		return x.KeyringId
	}
	return ""
}

func (x *EncryptRequest) GetKeyId() string {
	if x != nil {
		return x.KeyId
	}
	return ""
}

type EncryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The encrypted data.
	Cipher []byte `protobuf:"bytes,1,opt,name=cipher,proto3" json:"cipher,omitempty"`
}

func (x *EncryptResponse) Reset() {
	*x = EncryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_k8s_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptResponse) ProtoMessage() {}

func (x *EncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_k8s_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptResponse.ProtoReflect.Descriptor instead.
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return file_apis_k8s_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptResponse) GetCipher() []byte {
	if x != nil {
		return x.Cipher
	}
	return nil
}

var File_apis_k8s_v1_messages_proto protoreflect.FileDescriptor

var file_apis_k8s_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x68,
	0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b,
	0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64, 0x22,
	0x27, 0x0a, 0x0f, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x22, 0x76, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65,
	0x79, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6b, 0x65, 0x79, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6b, 0x65, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x65, 0x79, 0x49, 0x64,
	0x22, 0x29, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73,
	0x63, 0x70, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x6b, 0x38, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_k8s_v1_messages_proto_rawDescOnce sync.Once
	file_apis_k8s_v1_messages_proto_rawDescData = file_apis_k8s_v1_messages_proto_rawDesc
)

func file_apis_k8s_v1_messages_proto_rawDescGZIP() []byte {
	file_apis_k8s_v1_messages_proto_rawDescOnce.Do(func() {
		file_apis_k8s_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_k8s_v1_messages_proto_rawDescData)
	})
	return file_apis_k8s_v1_messages_proto_rawDescData
}

var file_apis_k8s_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apis_k8s_v1_messages_proto_goTypes = []interface{}{
	(*DecryptRequest)(nil),  // 0: thalescpl.io.ekms.k8s.v1.DecryptRequest
	(*DecryptResponse)(nil), // 1: thalescpl.io.ekms.k8s.v1.DecryptResponse
	(*EncryptRequest)(nil),  // 2: thalescpl.io.ekms.k8s.v1.EncryptRequest
	(*EncryptResponse)(nil), // 3: thalescpl.io.ekms.k8s.v1.EncryptResponse
}
var file_apis_k8s_v1_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apis_k8s_v1_messages_proto_init() }
func file_apis_k8s_v1_messages_proto_init() {
	if File_apis_k8s_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_k8s_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptRequest); i {
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
		file_apis_k8s_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptResponse); i {
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
		file_apis_k8s_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptRequest); i {
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
		file_apis_k8s_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptResponse); i {
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
			RawDescriptor: file_apis_k8s_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_k8s_v1_messages_proto_goTypes,
		DependencyIndexes: file_apis_k8s_v1_messages_proto_depIdxs,
		MessageInfos:      file_apis_k8s_v1_messages_proto_msgTypes,
	}.Build()
	File_apis_k8s_v1_messages_proto = out.File
	file_apis_k8s_v1_messages_proto_rawDesc = nil
	file_apis_k8s_v1_messages_proto_goTypes = nil
	file_apis_k8s_v1_messages_proto_depIdxs = nil
}
