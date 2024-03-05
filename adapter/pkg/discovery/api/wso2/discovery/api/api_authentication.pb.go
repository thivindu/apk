//  Copyright (c) 2021, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  WSO2 LLC. licenses this file to you under the Apache License,
//  Version 2.0 (the "License"); you may not use this file except
//  in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/api/api_authentication.proto

package api

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

// APIKey apikey related configurations
type APIKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                // name of the security scheme
	In                  string `protobuf:"bytes,2,opt,name=in,proto3" json:"in,omitempty"`                                    // location of the API key in request
	SendTokenToUpstream bool   `protobuf:"varint,3,opt,name=sendTokenToUpstream,proto3" json:"sendTokenToUpstream,omitempty"` // send the token to upstream
}

func (x *APIKey) Reset() {
	*x = APIKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIKey) ProtoMessage() {}

func (x *APIKey) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIKey.ProtoReflect.Descriptor instead.
func (*APIKey) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_api_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *APIKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *APIKey) GetIn() string {
	if x != nil {
		return x.In
	}
	return ""
}

func (x *APIKey) GetSendTokenToUpstream() bool {
	if x != nil {
		return x.SendTokenToUpstream
	}
	return false
}

type JWT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header              string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`                            // name of the header containing the JWT
	SendTokenToUpstream bool   `protobuf:"varint,2,opt,name=sendTokenToUpstream,proto3" json:"sendTokenToUpstream,omitempty"` // send the token to upstream
}

func (x *JWT) Reset() {
	*x = JWT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWT) ProtoMessage() {}

func (x *JWT) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWT.ProtoReflect.Descriptor instead.
func (*JWT) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_api_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *JWT) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *JWT) GetSendTokenToUpstream() bool {
	if x != nil {
		return x.SendTokenToUpstream
	}
	return false
}

type APIAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Disabled       bool            `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"` // disable authentication
	Jwt            *JWT            `protobuf:"bytes,2,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Apikey         []*APIKey       `protobuf:"bytes,3,rep,name=apikey,proto3" json:"apikey,omitempty"`
	TestConsoleKey *TestConsoleKey `protobuf:"bytes,4,opt,name=testConsoleKey,proto3" json:"testConsoleKey,omitempty"`
}

func (x *APIAuthentication) Reset() {
	*x = APIAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIAuthentication) ProtoMessage() {}

func (x *APIAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIAuthentication.ProtoReflect.Descriptor instead.
func (*APIAuthentication) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_api_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *APIAuthentication) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *APIAuthentication) GetJwt() *JWT {
	if x != nil {
		return x.Jwt
	}
	return nil
}

func (x *APIAuthentication) GetApikey() []*APIKey {
	if x != nil {
		return x.Apikey
	}
	return nil
}

func (x *APIAuthentication) GetTestConsoleKey() *TestConsoleKey {
	if x != nil {
		return x.TestConsoleKey
	}
	return nil
}

type TestConsoleKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header              string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`                            // name of the header containing the test key
	SendTokenToUpstream bool   `protobuf:"varint,2,opt,name=sendTokenToUpstream,proto3" json:"sendTokenToUpstream,omitempty"` // send the token to upstream
}

func (x *TestConsoleKey) Reset() {
	*x = TestConsoleKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConsoleKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConsoleKey) ProtoMessage() {}

func (x *TestConsoleKey) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_api_authentication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConsoleKey.ProtoReflect.Descriptor instead.
func (*TestConsoleKey) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_api_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *TestConsoleKey) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *TestConsoleKey) GetSendTokenToUpstream() bool {
	if x != nil {
		return x.SendTokenToUpstream
	}
	return false
}

var File_wso2_discovery_api_api_authentication_proto protoreflect.FileDescriptor

var file_wso2_discovery_api_api_authentication_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77,
	0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x22, 0x5e, 0x0a, 0x06, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x6e, 0x12,
	0x30, 0x0a, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x55, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x65,
	0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x22, 0x4f, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x30, 0x0a, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x55,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0xda, 0x01, 0x0a, 0x11, 0x41, 0x50, 0x49, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12,
	0x32, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x61, 0x70, 0x69,
	0x6b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x73,
	0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x52,
	0x0e, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x22,
	0x5a, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x65, 0x6e,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x6f, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x6f, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x7e, 0x0a, 0x23, 0x6f,
	0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x61, 0x70, 0x6b, 0x2e, 0x65, 0x6e, 0x66, 0x6f,
	0x72, 0x63, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x42, 0x16, 0x41, 0x50, 0x49, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_api_api_authentication_proto_rawDescOnce sync.Once
	file_wso2_discovery_api_api_authentication_proto_rawDescData = file_wso2_discovery_api_api_authentication_proto_rawDesc
)

func file_wso2_discovery_api_api_authentication_proto_rawDescGZIP() []byte {
	file_wso2_discovery_api_api_authentication_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_api_api_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_api_api_authentication_proto_rawDescData)
	})
	return file_wso2_discovery_api_api_authentication_proto_rawDescData
}

var file_wso2_discovery_api_api_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wso2_discovery_api_api_authentication_proto_goTypes = []interface{}{
	(*APIKey)(nil),            // 0: wso2.discovery.api.APIKey
	(*JWT)(nil),               // 1: wso2.discovery.api.JWT
	(*APIAuthentication)(nil), // 2: wso2.discovery.api.APIAuthentication
	(*TestConsoleKey)(nil),    // 3: wso2.discovery.api.TestConsoleKey
}
var file_wso2_discovery_api_api_authentication_proto_depIdxs = []int32{
	1, // 0: wso2.discovery.api.APIAuthentication.jwt:type_name -> wso2.discovery.api.JWT
	0, // 1: wso2.discovery.api.APIAuthentication.apikey:type_name -> wso2.discovery.api.APIKey
	3, // 2: wso2.discovery.api.APIAuthentication.testConsoleKey:type_name -> wso2.discovery.api.TestConsoleKey
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_wso2_discovery_api_api_authentication_proto_init() }
func file_wso2_discovery_api_api_authentication_proto_init() {
	if File_wso2_discovery_api_api_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_api_api_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIKey); i {
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
		file_wso2_discovery_api_api_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWT); i {
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
		file_wso2_discovery_api_api_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIAuthentication); i {
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
		file_wso2_discovery_api_api_authentication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConsoleKey); i {
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
			RawDescriptor: file_wso2_discovery_api_api_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_api_api_authentication_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_api_api_authentication_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_api_api_authentication_proto_msgTypes,
	}.Build()
	File_wso2_discovery_api_api_authentication_proto = out.File
	file_wso2_discovery_api_api_authentication_proto_rawDesc = nil
	file_wso2_discovery_api_api_authentication_proto_goTypes = nil
	file_wso2_discovery_api_api_authentication_proto_depIdxs = nil
}
