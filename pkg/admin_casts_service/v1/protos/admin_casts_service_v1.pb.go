// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: admin_casts_service_v1.proto

package protos

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_admin_casts_service_v1_proto protoreflect.FileDescriptor

var file_admin_casts_service_v1_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x25, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb4, 0x0b, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x74, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x31, 0x12, 0x66, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x61, 0x73, 0x74, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x7b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x12, 0x7d, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x61, 0x73, 0x74, 0x42, 0x79,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61,
	0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x43, 0x61, 0x73, 0x74, 0x42, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73,
	0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x73,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x5e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x73, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x12,
	0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x74, 0x12, 0x26, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x61, 0x73, 0x74, 0x12, 0x2e, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x46, 0x6f, 0x72, 0x43, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x73, 0x74, 0x2f, 0x7b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x7d, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x89, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x54, 0x6f, 0x54, 0x68, 0x65, 0x43, 0x61, 0x73, 0x74, 0x12, 0x2f,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x54,
	0x6f, 0x54, 0x68, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22,
	0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x7b, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x49, 0x44, 0x7d, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x3a,
	0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x68, 0x65, 0x43, 0x61, 0x73, 0x74, 0x12,
	0x34, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x68, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x2f,
	0x7b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x7d, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x2f, 0x7b, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x63, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x20, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x8b, 0x01, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xc0, 0x02, 0x5a,
	0x1d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x73, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x92, 0x41,
	0x9d, 0x02, 0x12, 0x65, 0x0a, 0x1d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x20, 0x70, 0x61, 0x6e, 0x65,
	0x6c, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x63, 0x61, 0x73, 0x74, 0x73, 0x20, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x07, 0x46, 0x61, 0x6c, 0x6f, 0x6b, 0x75, 0x74, 0x12, 0x1a,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x46, 0x61, 0x6c, 0x6f, 0x6b, 0x75, 0x74, 0x1a, 0x18, 0x74, 0x69, 0x6d, 0x75,
	0x72, 0x2e, 0x73, 0x69, 0x6e, 0x65, 0x6c, 0x6e, 0x69, 0x6b, 0x40, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x72, 0x75, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x49, 0x0a, 0x2a, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65,
	0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x3b, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x34, 0x0a, 0x15, 0x53, 0x6f, 0x6d,
	0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x65, 0x6e, 0x74, 0x20, 0x77, 0x72, 0x6f, 0x6e,
	0x67, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_admin_casts_service_v1_proto_goTypes = []interface{}{
	(*GetCastRequest)(nil),                  // 0: admin_casts_service.GetCastRequest
	(*SearchCastByLabelRequest)(nil),        // 1: admin_casts_service.SearchCastByLabelRequest
	(*GetCastsRequest)(nil),                 // 2: admin_casts_service.GetCastsRequest
	(*CreateCastRequest)(nil),               // 3: admin_casts_service.CreateCastRequest
	(*UpdateLabelForCastRequest)(nil),       // 4: admin_casts_service.UpdateLabelForCastRequest
	(*AddPersonsToTheCastRequest)(nil),      // 5: admin_casts_service.AddPersonsToTheCastRequest
	(*RemovePersonsFromTheCastRequest)(nil), // 6: admin_casts_service.RemovePersonsFromTheCastRequest
	(*DeleteCastRequest)(nil),               // 7: admin_casts_service.DeleteCastRequest
	(*emptypb.Empty)(nil),                   // 8: google.protobuf.Empty
	(*CreateProfessionRequest)(nil),         // 9: admin_casts_service.CreateProfessionRequest
	(*UpdateProfessionRequest)(nil),         // 10: admin_casts_service.UpdateProfessionRequest
	(*DeleteProfessionRequest)(nil),         // 11: admin_casts_service.DeleteProfessionRequest
	(*Cast)(nil),                            // 12: admin_casts_service.Cast
	(*CastsLabels)(nil),                     // 13: admin_casts_service.CastsLabels
	(*Casts)(nil),                           // 14: admin_casts_service.Casts
	(*Professions)(nil),                     // 15: admin_casts_service.Professions
	(*CreateProfessionResponse)(nil),        // 16: admin_casts_service.CreateProfessionResponse
}
var file_admin_casts_service_v1_proto_depIdxs = []int32{
	0,  // 0: admin_casts_service.castsServiceV1.GetCast:input_type -> admin_casts_service.GetCastRequest
	1,  // 1: admin_casts_service.castsServiceV1.SearchCastByLabel:input_type -> admin_casts_service.SearchCastByLabelRequest
	2,  // 2: admin_casts_service.castsServiceV1.GetCasts:input_type -> admin_casts_service.GetCastsRequest
	3,  // 3: admin_casts_service.castsServiceV1.CreateCast:input_type -> admin_casts_service.CreateCastRequest
	4,  // 4: admin_casts_service.castsServiceV1.UpdateLabelForCast:input_type -> admin_casts_service.UpdateLabelForCastRequest
	5,  // 5: admin_casts_service.castsServiceV1.AddPersonsToTheCast:input_type -> admin_casts_service.AddPersonsToTheCastRequest
	6,  // 6: admin_casts_service.castsServiceV1.RemovePersonsFromTheCast:input_type -> admin_casts_service.RemovePersonsFromTheCastRequest
	7,  // 7: admin_casts_service.castsServiceV1.DeleteCast:input_type -> admin_casts_service.DeleteCastRequest
	8,  // 8: admin_casts_service.castsServiceV1.GetProfessions:input_type -> google.protobuf.Empty
	9,  // 9: admin_casts_service.castsServiceV1.CreateProfession:input_type -> admin_casts_service.CreateProfessionRequest
	10, // 10: admin_casts_service.castsServiceV1.UpdateProfession:input_type -> admin_casts_service.UpdateProfessionRequest
	11, // 11: admin_casts_service.castsServiceV1.DeleteProfession:input_type -> admin_casts_service.DeleteProfessionRequest
	12, // 12: admin_casts_service.castsServiceV1.GetCast:output_type -> admin_casts_service.Cast
	13, // 13: admin_casts_service.castsServiceV1.SearchCastByLabel:output_type -> admin_casts_service.CastsLabels
	14, // 14: admin_casts_service.castsServiceV1.GetCasts:output_type -> admin_casts_service.Casts
	8,  // 15: admin_casts_service.castsServiceV1.CreateCast:output_type -> google.protobuf.Empty
	8,  // 16: admin_casts_service.castsServiceV1.UpdateLabelForCast:output_type -> google.protobuf.Empty
	8,  // 17: admin_casts_service.castsServiceV1.AddPersonsToTheCast:output_type -> google.protobuf.Empty
	8,  // 18: admin_casts_service.castsServiceV1.RemovePersonsFromTheCast:output_type -> google.protobuf.Empty
	8,  // 19: admin_casts_service.castsServiceV1.DeleteCast:output_type -> google.protobuf.Empty
	15, // 20: admin_casts_service.castsServiceV1.GetProfessions:output_type -> admin_casts_service.Professions
	16, // 21: admin_casts_service.castsServiceV1.CreateProfession:output_type -> admin_casts_service.CreateProfessionResponse
	8,  // 22: admin_casts_service.castsServiceV1.UpdateProfession:output_type -> google.protobuf.Empty
	8,  // 23: admin_casts_service.castsServiceV1.DeleteProfession:output_type -> google.protobuf.Empty
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_admin_casts_service_v1_proto_init() }
func file_admin_casts_service_v1_proto_init() {
	if File_admin_casts_service_v1_proto != nil {
		return
	}
	file_admin_casts_service_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_casts_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_casts_service_v1_proto_goTypes,
		DependencyIndexes: file_admin_casts_service_v1_proto_depIdxs,
	}.Build()
	File_admin_casts_service_v1_proto = out.File
	file_admin_casts_service_v1_proto_rawDesc = nil
	file_admin_casts_service_v1_proto_goTypes = nil
	file_admin_casts_service_v1_proto_depIdxs = nil
}
