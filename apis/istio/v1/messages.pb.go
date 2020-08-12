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
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: apis/istio/v1/messages.proto

package istio

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

type KeyKind int32

const (
	KeyKind_UNKNOWN KeyKind = 0
	KeyKind_AES     KeyKind = 1
	KeyKind_RSA     KeyKind = 2
	KeyKind_ECC     KeyKind = 3
)

// Enum value maps for KeyKind.
var (
	KeyKind_name = map[int32]string{
		0: "UNKNOWN",
		1: "AES",
		2: "RSA",
		3: "ECC",
	}
	KeyKind_value = map[string]int32{
		"UNKNOWN": 0,
		"AES":     1,
		"RSA":     2,
		"ECC":     3,
	}
)

func (x KeyKind) Enum() *KeyKind {
	p := new(KeyKind)
	*p = x
	return p
}

func (x KeyKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeyKind) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_istio_v1_messages_proto_enumTypes[0].Descriptor()
}

func (KeyKind) Type() protoreflect.EnumType {
	return &file_apis_istio_v1_messages_proto_enumTypes[0]
}

func (x KeyKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeyKind.Descriptor instead.
func (KeyKind) EnumDescriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{0}
}

type GenerateDEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key size in bits
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// What kind of key is it... only Symmetric kinds will be supported
	Kind KeyKind `protobuf:"varint,2,opt,name=kind,proto3,enum=thalescpl.io.ekms.istio.v1.KeyKind" json:"kind,omitempty"`
	// Parent KID of the KEK
	KekKid []byte `protobuf:"bytes,3,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
}

func (x *GenerateDEKRequest) Reset() {
	*x = GenerateDEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDEKRequest) ProtoMessage() {}

func (x *GenerateDEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDEKRequest.ProtoReflect.Descriptor instead.
func (*GenerateDEKRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateDEKRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GenerateDEKRequest) GetKind() KeyKind {
	if x != nil {
		return x.Kind
	}
	return KeyKind_UNKNOWN
}

func (x *GenerateDEKRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

type GenerateDEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encrypted key blob
	EncryptedDekBlob []byte `protobuf:"bytes,1,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
}

func (x *GenerateDEKResponse) Reset() {
	*x = GenerateDEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateDEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateDEKResponse) ProtoMessage() {}

func (x *GenerateDEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateDEKResponse.ProtoReflect.Descriptor instead.
func (*GenerateDEKResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateDEKResponse) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

type GenerateSEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key size in bits
	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// What kind of key is it... only Asymmetric kinds will be supported
	Kind KeyKind `protobuf:"varint,2,opt,name=kind,proto3,enum=thalescpl.io.ekms.istio.v1.KeyKind" json:"kind,omitempty"`
	// Encrypted blob of DEK
	EncryptedDekBlob []byte `protobuf:"bytes,3,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
	// Encrypted blob of SEK
	EncryptedSekBlob []byte `protobuf:"bytes,4,opt,name=encrypted_sek_blob,json=encryptedSekBlob,proto3" json:"encrypted_sek_blob,omitempty"`
	// Parent KID of the KEK
	KekKid []byte `protobuf:"bytes,5,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
}

func (x *GenerateSEKRequest) Reset() {
	*x = GenerateSEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSEKRequest) ProtoMessage() {}

func (x *GenerateSEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSEKRequest.ProtoReflect.Descriptor instead.
func (*GenerateSEKRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateSEKRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GenerateSEKRequest) GetKind() KeyKind {
	if x != nil {
		return x.Kind
	}
	return KeyKind_UNKNOWN
}

func (x *GenerateSEKRequest) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

func (x *GenerateSEKRequest) GetEncryptedSekBlob() []byte {
	if x != nil {
		return x.EncryptedSekBlob
	}
	return nil
}

func (x *GenerateSEKRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

type GenerateSEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encrypted blob of SEK encrypted by DEK
	EncryptedSekBlob []byte `protobuf:"bytes,1,opt,name=encrypted_sek_blob,json=encryptedSekBlob,proto3" json:"encrypted_sek_blob,omitempty"`
}

func (x *GenerateSEKResponse) Reset() {
	*x = GenerateSEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSEKResponse) ProtoMessage() {}

func (x *GenerateSEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSEKResponse.ProtoReflect.Descriptor instead.
func (*GenerateSEKResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateSEKResponse) GetEncryptedSekBlob() []byte {
	if x != nil {
		return x.EncryptedSekBlob
	}
	return nil
}

type LoadSEKRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encrypted blob of DEK
	EncryptedDekBlob []byte `protobuf:"bytes,1,opt,name=encrypted_dek_blob,json=encryptedDekBlob,proto3" json:"encrypted_dek_blob,omitempty"`
	// Encrypted blob of SEK
	EncryptedSekBlob []byte `protobuf:"bytes,2,opt,name=encrypted_sek_blob,json=encryptedSekBlob,proto3" json:"encrypted_sek_blob,omitempty"`
	// KEK
	KekKid []byte `protobuf:"bytes,3,opt,name=kek_kid,json=kekKid,proto3" json:"kek_kid,omitempty"`
	// Root CA Key ID
	RootCaKid []byte `protobuf:"bytes,4,opt,name=root_ca_kid,json=rootCaKid,proto3" json:"root_ca_kid,omitempty"`
}

func (x *LoadSEKRequest) Reset() {
	*x = LoadSEKRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadSEKRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadSEKRequest) ProtoMessage() {}

func (x *LoadSEKRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadSEKRequest.ProtoReflect.Descriptor instead.
func (*LoadSEKRequest) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *LoadSEKRequest) GetEncryptedDekBlob() []byte {
	if x != nil {
		return x.EncryptedDekBlob
	}
	return nil
}

func (x *LoadSEKRequest) GetEncryptedSekBlob() []byte {
	if x != nil {
		return x.EncryptedSekBlob
	}
	return nil
}

func (x *LoadSEKRequest) GetKekKid() []byte {
	if x != nil {
		return x.KekKid
	}
	return nil
}

func (x *LoadSEKRequest) GetRootCaKid() []byte {
	if x != nil {
		return x.RootCaKid
	}
	return nil
}

type LoadSEKResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Clear SEK
	ClearSek []byte `protobuf:"bytes,1,opt,name=clear_sek,json=clearSek,proto3" json:"clear_sek,omitempty"`
}

func (x *LoadSEKResponse) Reset() {
	*x = LoadSEKResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_istio_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadSEKResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadSEKResponse) ProtoMessage() {}

func (x *LoadSEKResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_istio_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadSEKResponse.ProtoReflect.Descriptor instead.
func (*LoadSEKResponse) Descriptor() ([]byte, []int) {
	return file_apis_istio_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *LoadSEKResponse) GetClearSek() []byte {
	if x != nil {
		return x.ClearSek
	}
	return nil
}

var File_apis_istio_v1_messages_proto protoreflect.FileDescriptor

var file_apis_istio_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d,
	0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x12, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f,
	0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4b,
	0x65, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6b, 0x65, 0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6b,
	0x65, 0x6b, 0x4b, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b, 0x5f, 0x62, 0x6c,
	0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x44, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x22, 0xd6, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e,
	0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x65, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b, 0x5f,
	0x62, 0x6c, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x2c, 0x0a, 0x12,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x6b, 0x5f, 0x62, 0x6c,
	0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65,
	0x6b, 0x5f, 0x6b, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6b, 0x65, 0x6b,
	0x4b, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53,
	0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x53, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x22, 0xa5, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x61,
	0x64, 0x53, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x65, 0x6b, 0x5f, 0x62, 0x6c, 0x6f,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x44, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x53, 0x65, 0x6b, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65, 0x6b, 0x5f, 0x6b,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6b, 0x65, 0x6b, 0x4b, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0b, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x63, 0x61, 0x5f, 0x6b, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x74, 0x43, 0x61, 0x4b, 0x69, 0x64,
	0x22, 0x2e, 0x0a, 0x0f, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x53, 0x65, 0x6b,
	0x2a, 0x31, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x45, 0x53, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x41, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x43,
	0x43, 0x10, 0x03, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x6b,
	0x38, 0x73, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_istio_v1_messages_proto_rawDescOnce sync.Once
	file_apis_istio_v1_messages_proto_rawDescData = file_apis_istio_v1_messages_proto_rawDesc
)

func file_apis_istio_v1_messages_proto_rawDescGZIP() []byte {
	file_apis_istio_v1_messages_proto_rawDescOnce.Do(func() {
		file_apis_istio_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_istio_v1_messages_proto_rawDescData)
	})
	return file_apis_istio_v1_messages_proto_rawDescData
}

var file_apis_istio_v1_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apis_istio_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apis_istio_v1_messages_proto_goTypes = []interface{}{
	(KeyKind)(0),                // 0: thalescpl.io.ekms.istio.v1.KeyKind
	(*GenerateDEKRequest)(nil),  // 1: thalescpl.io.ekms.istio.v1.GenerateDEKRequest
	(*GenerateDEKResponse)(nil), // 2: thalescpl.io.ekms.istio.v1.GenerateDEKResponse
	(*GenerateSEKRequest)(nil),  // 3: thalescpl.io.ekms.istio.v1.GenerateSEKRequest
	(*GenerateSEKResponse)(nil), // 4: thalescpl.io.ekms.istio.v1.GenerateSEKResponse
	(*LoadSEKRequest)(nil),      // 5: thalescpl.io.ekms.istio.v1.LoadSEKRequest
	(*LoadSEKResponse)(nil),     // 6: thalescpl.io.ekms.istio.v1.LoadSEKResponse
}
var file_apis_istio_v1_messages_proto_depIdxs = []int32{
	0, // 0: thalescpl.io.ekms.istio.v1.GenerateDEKRequest.kind:type_name -> thalescpl.io.ekms.istio.v1.KeyKind
	0, // 1: thalescpl.io.ekms.istio.v1.GenerateSEKRequest.kind:type_name -> thalescpl.io.ekms.istio.v1.KeyKind
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apis_istio_v1_messages_proto_init() }
func file_apis_istio_v1_messages_proto_init() {
	if File_apis_istio_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_istio_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDEKRequest); i {
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
		file_apis_istio_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateDEKResponse); i {
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
		file_apis_istio_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateSEKRequest); i {
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
		file_apis_istio_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateSEKResponse); i {
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
		file_apis_istio_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadSEKRequest); i {
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
		file_apis_istio_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadSEKResponse); i {
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
			RawDescriptor: file_apis_istio_v1_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apis_istio_v1_messages_proto_goTypes,
		DependencyIndexes: file_apis_istio_v1_messages_proto_depIdxs,
		EnumInfos:         file_apis_istio_v1_messages_proto_enumTypes,
		MessageInfos:      file_apis_istio_v1_messages_proto_msgTypes,
	}.Build()
	File_apis_istio_v1_messages_proto = out.File
	file_apis_istio_v1_messages_proto_rawDesc = nil
	file_apis_istio_v1_messages_proto_goTypes = nil
	file_apis_istio_v1_messages_proto_depIdxs = nil
}
