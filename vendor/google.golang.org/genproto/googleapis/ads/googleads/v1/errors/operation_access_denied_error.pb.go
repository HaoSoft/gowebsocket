// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/operation_access_denied_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum_OperationAccessDeniedError int32

const (
	// Enum unspecified.
	OperationAccessDeniedErrorEnum_UNSPECIFIED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 0
	// The received error code is not known in this version.
	OperationAccessDeniedErrorEnum_UNKNOWN OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 1
	// Unauthorized invocation of a service's method (get, mutate, etc.)
	OperationAccessDeniedErrorEnum_ACTION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 2
	// Unauthorized CREATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_CREATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 3
	// Unauthorized REMOVE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_REMOVE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 4
	// Unauthorized UPDATE operation in invoking a service's mutate method.
	OperationAccessDeniedErrorEnum_UPDATE_OPERATION_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 5
	// A mutate action is not allowed on this campaign, from this client.
	OperationAccessDeniedErrorEnum_MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 6
	// This operation is not permitted on this campaign type
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 7
	// A CREATE operation may not set status to REMOVED.
	OperationAccessDeniedErrorEnum_CREATE_AS_REMOVED_NOT_PERMITTED OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 8
	// This operation is not allowed because the campaign or adgroup is removed.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 9
	// This operation is not permitted on this ad group type.
	OperationAccessDeniedErrorEnum_OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 10
	// The mutate is not allowed for this customer.
	OperationAccessDeniedErrorEnum_MUTATE_NOT_PERMITTED_FOR_CUSTOMER OperationAccessDeniedErrorEnum_OperationAccessDeniedError = 11
)

var OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "ACTION_NOT_PERMITTED",
	3:  "CREATE_OPERATION_NOT_PERMITTED",
	4:  "REMOVE_OPERATION_NOT_PERMITTED",
	5:  "UPDATE_OPERATION_NOT_PERMITTED",
	6:  "MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT",
	7:  "OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE",
	8:  "CREATE_AS_REMOVED_NOT_PERMITTED",
	9:  "OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE",
	10: "OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE",
	11: "MUTATE_NOT_PERMITTED_FOR_CUSTOMER",
}

var OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":                                  0,
	"UNKNOWN":                                      1,
	"ACTION_NOT_PERMITTED":                         2,
	"CREATE_OPERATION_NOT_PERMITTED":               3,
	"REMOVE_OPERATION_NOT_PERMITTED":               4,
	"UPDATE_OPERATION_NOT_PERMITTED":               5,
	"MUTATE_ACTION_NOT_PERMITTED_FOR_CLIENT":       6,
	"OPERATION_NOT_PERMITTED_FOR_CAMPAIGN_TYPE":    7,
	"CREATE_AS_REMOVED_NOT_PERMITTED":              8,
	"OPERATION_NOT_PERMITTED_FOR_REMOVED_RESOURCE": 9,
	"OPERATION_NOT_PERMITTED_FOR_AD_GROUP_TYPE":    10,
	"MUTATE_NOT_PERMITTED_FOR_CUSTOMER":            11,
}

func (x OperationAccessDeniedErrorEnum_OperationAccessDeniedError) String() string {
	return proto.EnumName(OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name, int32(x))
}

func (OperationAccessDeniedErrorEnum_OperationAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f047890255dc139d, []int{0, 0}
}

// Container for enum describing possible operation access denied errors.
type OperationAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationAccessDeniedErrorEnum) Reset()         { *m = OperationAccessDeniedErrorEnum{} }
func (m *OperationAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*OperationAccessDeniedErrorEnum) ProtoMessage()    {}
func (*OperationAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f047890255dc139d, []int{0}
}

func (m *OperationAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationAccessDeniedErrorEnum.Merge(m, src)
}
func (m *OperationAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_OperationAccessDeniedErrorEnum.Size(m)
}
func (m *OperationAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_OperationAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.OperationAccessDeniedErrorEnum_OperationAccessDeniedError", OperationAccessDeniedErrorEnum_OperationAccessDeniedError_name, OperationAccessDeniedErrorEnum_OperationAccessDeniedError_value)
	proto.RegisterType((*OperationAccessDeniedErrorEnum)(nil), "google.ads.googleads.v1.errors.OperationAccessDeniedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/operation_access_denied_error.proto", fileDescriptor_f047890255dc139d)
}

var fileDescriptor_f047890255dc139d = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x8a, 0xd3, 0x40,
	0x1c, 0xc6, 0x6d, 0xab, 0xbb, 0x3a, 0x3d, 0x18, 0x06, 0x0f, 0xb2, 0x48, 0x17, 0x23, 0x0a, 0x8a,
	0x4e, 0x0c, 0xde, 0xe2, 0x69, 0x9a, 0xcc, 0x96, 0xa0, 0x99, 0x19, 0xa6, 0x93, 0x8a, 0x52, 0x18,
	0x62, 0x13, 0x42, 0x61, 0x37, 0x53, 0x32, 0xb5, 0x0f, 0xe4, 0xd1, 0x27, 0x11, 0x1f, 0x45, 0xdf,
	0xc0, 0x93, 0x24, 0xd3, 0xf4, 0x20, 0x4d, 0xf7, 0x94, 0x3f, 0xc9, 0xef, 0xff, 0x7d, 0xdf, 0x9f,
	0x7c, 0x60, 0x5a, 0x6a, 0x5d, 0x5e, 0x17, 0x5e, 0x96, 0x1b, 0xcf, 0x8e, 0xcd, 0xb4, 0xf3, 0xbd,
	0xa2, 0xae, 0x75, 0x6d, 0x3c, 0xbd, 0x29, 0xea, 0x6c, 0xbb, 0xd6, 0x95, 0xca, 0x56, 0xab, 0xc2,
	0x18, 0x95, 0x17, 0xd5, 0xba, 0xc8, 0x55, 0xfb, 0x19, 0x6d, 0x6a, 0xbd, 0xd5, 0x70, 0x62, 0x17,
	0x51, 0x96, 0x1b, 0x74, 0xd0, 0x40, 0x3b, 0x1f, 0x59, 0x8d, 0x8b, 0x27, 0x9d, 0xc7, 0x66, 0xed,
	0x65, 0x55, 0xa5, 0xb7, 0xad, 0xa0, 0xb1, 0xdb, 0xee, 0x9f, 0x11, 0x98, 0xb0, 0xce, 0x05, 0xb7,
	0x26, 0x51, 0xeb, 0x41, 0x9a, 0x6d, 0x52, 0x7d, 0xbb, 0x71, 0x7f, 0x8e, 0xc0, 0x45, 0x3f, 0x02,
	0x1f, 0x82, 0x71, 0x4a, 0xe7, 0x9c, 0x84, 0xf1, 0x55, 0x4c, 0x22, 0xe7, 0x0e, 0x1c, 0x83, 0xf3,
	0x94, 0x7e, 0xa0, 0xec, 0x13, 0x75, 0x06, 0xf0, 0x31, 0x78, 0x84, 0x43, 0x19, 0x33, 0xaa, 0x28,
	0x93, 0x8a, 0x13, 0x91, 0xc4, 0x52, 0x92, 0xc8, 0x19, 0x42, 0x17, 0x4c, 0x42, 0x41, 0xb0, 0x24,
	0x8a, 0x71, 0x22, 0xf0, 0x11, 0x66, 0xd4, 0x30, 0x82, 0x24, 0x6c, 0xd1, 0xcf, 0xdc, 0x6d, 0x98,
	0x94, 0x47, 0xa7, 0x74, 0xee, 0xc1, 0x57, 0xe0, 0x45, 0x92, 0xca, 0x86, 0x39, 0x16, 0x46, 0x5d,
	0x31, 0xa1, 0xc2, 0x8f, 0x31, 0xa1, 0xd2, 0x39, 0x83, 0x6f, 0xc0, 0xcb, 0x1e, 0x21, 0xcb, 0xe1,
	0x84, 0xe3, 0x78, 0x46, 0x95, 0xfc, 0xcc, 0x89, 0x73, 0x0e, 0x9f, 0x81, 0xcb, 0xfd, 0x19, 0x78,
	0xae, 0x6c, 0xd8, 0xe8, 0x3f, 0xff, 0xfb, 0xf0, 0x2d, 0x78, 0x7d, 0x4a, 0xb3, 0x5b, 0x13, 0x64,
	0xce, 0x52, 0x11, 0x12, 0xe7, 0xc1, 0x6d, 0x29, 0x70, 0xa4, 0x66, 0x82, 0xa5, 0xdc, 0xa6, 0x00,
	0xf0, 0x39, 0x78, 0xba, 0x3f, 0xf0, 0x48, 0xe2, 0x74, 0x2e, 0x59, 0x42, 0x84, 0x33, 0x9e, 0xfe,
	0x1d, 0x00, 0x77, 0xa5, 0x6f, 0xd0, 0xe9, 0xca, 0x4c, 0x2f, 0xfb, 0x7f, 0x37, 0x6f, 0x5a, 0xc3,
	0x07, 0x5f, 0xa2, 0xbd, 0x44, 0xa9, 0xaf, 0xb3, 0xaa, 0x44, 0xba, 0x2e, 0xbd, 0xb2, 0xa8, 0xda,
	0x4e, 0x75, 0x4d, 0xde, 0xac, 0x4d, 0x5f, 0xb1, 0xdf, 0xdb, 0xc7, 0xf7, 0xe1, 0x68, 0x86, 0xf1,
	0x8f, 0xe1, 0x64, 0x66, 0xc5, 0x70, 0x6e, 0x90, 0x1d, 0x9b, 0x69, 0xe1, 0xa3, 0xd6, 0xd2, 0xfc,
	0xea, 0x80, 0x25, 0xce, 0xcd, 0xf2, 0x00, 0x2c, 0x17, 0xfe, 0xd2, 0x02, 0xbf, 0x87, 0xae, 0x7d,
	0x1b, 0x04, 0x38, 0x37, 0x41, 0x70, 0x40, 0x82, 0x60, 0xe1, 0x07, 0x81, 0x85, 0xbe, 0x9e, 0xb5,
	0xe9, 0xde, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x05, 0x7d, 0xe7, 0xd4, 0x75, 0x03, 0x00, 0x00,
}
