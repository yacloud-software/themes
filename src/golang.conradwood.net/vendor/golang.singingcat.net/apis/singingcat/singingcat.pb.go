// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/singingcat/singingcat.proto
// DO NOT EDIT!

/*
Package singingcat is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/singingcat/singingcat.proto

It has these top-level messages:
	Void
	Command
	Status
	ModuleRef
*/
package singingcat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SensorType int32

const (
	SensorType_SENSORTYPE_UNDEFINED   SensorType = 0
	SensorType_SENSORTYPE_TEMPERATURE SensorType = 1
)

var SensorType_name = map[int32]string{
	0: "SENSORTYPE_UNDEFINED",
	1: "SENSORTYPE_TEMPERATURE",
}
var SensorType_value = map[string]int32{
	"SENSORTYPE_UNDEFINED":   0,
	"SENSORTYPE_TEMPERATURE": 1,
}

func (x SensorType) String() string {
	return proto.EnumName(SensorType_name, int32(x))
}
func (SensorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ComType int32

const (
	ComType_UNDEFINED             ComType = 0
	ComType_NOOP                  ComType = 1
	ComType_LOG                   ComType = 2
	ComType_ARE_YOU_A_SINGINGCAT  ComType = 3
	ComType_SET_WIRELESS          ComType = 4
	ComType_SET_PIN               ComType = 5
	ComType_RESET                 ComType = 6
	ComType_WIFIRESET             ComType = 7
	ComType_RADIOPING             ComType = 8
	ComType_RADIOPINGLOOP         ComType = 9
	ComType_FREEZE                ComType = 10
	ComType_LIST_MODULES          ComType = 11
	ComType_ROUTE_REQUEST         ComType = 12
	ComType_STROBE                ComType = 13
	ComType_GETPUBKEY             ComType = 14
	ComType_SETCLOUDTOKEN         ComType = 15
	ComType_SET_LOGGING           ComType = 16
	ComType_SET_DAC               ComType = 17
	ComType_SET_SERIAL_PORT       ComType = 18
	ComType_FLASH_APP             ComType = 19
	ComType_MICROSTROBE           ComType = 20
	ComType_SET_SERVER_NAME       ComType = 21
	ComType_SENSOR_VALUE          ComType = 22
	ComType_ANNOUNCE              ComType = 23
	ComType_FACTORY_DEFAULT       ComType = 24
	ComType_SOFTWARE_INFO         ComType = 25
	ComType_BUTTON_PRESSED        ComType = 26
	ComType_STREAMSETUP           ComType = 27
	ComType_STREAMDATA            ComType = 28
	ComType_STREAMCONTROL         ComType = 29
	ComType_BLINKLED              ComType = 30
	ComType_SENSOR_REQUEST        ComType = 31
	ComType_USER                  ComType = 32
	ComType_RADIO_GET_CONFIG      ComType = 33
	ComType_RADIO_SET_CONFIG      ComType = 34
	ComType_STARTAPP              ComType = 35
	ComType_INSTALLED_SOFTWARE    ComType = 36
	ComType_POWERUP               ComType = 37
	ComType_GETCLOUDTOKEN         ComType = 38
	ComType_REQUEST_RADIO_FORWARD ComType = 39
	ComType_RADIO_FORWARD         ComType = 40
	ComType_SET_CONFIG_FLAG       ComType = 41
	ComType_SET_CONFIG_FLAGS      ComType = 42
	ComType_SET_SENSOR_SERVER     ComType = 43
	ComType_WIFI_INFO             ComType = 44
	ComType_GET_CONFIG_FLAGS      ComType = 45
	ComType_WIFI_SCAN             ComType = 46
	ComType_ROUTE                 ComType = 47
	ComType_SENSOR_LIST           ComType = 48
	ComType_SENSOR_CONFIG         ComType = 49
	ComType_SENSOR_REPORT         ComType = 50
)

var ComType_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "NOOP",
	2:  "LOG",
	3:  "ARE_YOU_A_SINGINGCAT",
	4:  "SET_WIRELESS",
	5:  "SET_PIN",
	6:  "RESET",
	7:  "WIFIRESET",
	8:  "RADIOPING",
	9:  "RADIOPINGLOOP",
	10: "FREEZE",
	11: "LIST_MODULES",
	12: "ROUTE_REQUEST",
	13: "STROBE",
	14: "GETPUBKEY",
	15: "SETCLOUDTOKEN",
	16: "SET_LOGGING",
	17: "SET_DAC",
	18: "SET_SERIAL_PORT",
	19: "FLASH_APP",
	20: "MICROSTROBE",
	21: "SET_SERVER_NAME",
	22: "SENSOR_VALUE",
	23: "ANNOUNCE",
	24: "FACTORY_DEFAULT",
	25: "SOFTWARE_INFO",
	26: "BUTTON_PRESSED",
	27: "STREAMSETUP",
	28: "STREAMDATA",
	29: "STREAMCONTROL",
	30: "BLINKLED",
	31: "SENSOR_REQUEST",
	32: "USER",
	33: "RADIO_GET_CONFIG",
	34: "RADIO_SET_CONFIG",
	35: "STARTAPP",
	36: "INSTALLED_SOFTWARE",
	37: "POWERUP",
	38: "GETCLOUDTOKEN",
	39: "REQUEST_RADIO_FORWARD",
	40: "RADIO_FORWARD",
	41: "SET_CONFIG_FLAG",
	42: "SET_CONFIG_FLAGS",
	43: "SET_SENSOR_SERVER",
	44: "WIFI_INFO",
	45: "GET_CONFIG_FLAGS",
	46: "WIFI_SCAN",
	47: "ROUTE",
	48: "SENSOR_LIST",
	49: "SENSOR_CONFIG",
	50: "SENSOR_REPORT",
}
var ComType_value = map[string]int32{
	"UNDEFINED": 0,
	"NOOP":      1,
	"LOG":       2,
	"ARE_YOU_A_SINGINGCAT":  3,
	"SET_WIRELESS":          4,
	"SET_PIN":               5,
	"RESET":                 6,
	"WIFIRESET":             7,
	"RADIOPING":             8,
	"RADIOPINGLOOP":         9,
	"FREEZE":                10,
	"LIST_MODULES":          11,
	"ROUTE_REQUEST":         12,
	"STROBE":                13,
	"GETPUBKEY":             14,
	"SETCLOUDTOKEN":         15,
	"SET_LOGGING":           16,
	"SET_DAC":               17,
	"SET_SERIAL_PORT":       18,
	"FLASH_APP":             19,
	"MICROSTROBE":           20,
	"SET_SERVER_NAME":       21,
	"SENSOR_VALUE":          22,
	"ANNOUNCE":              23,
	"FACTORY_DEFAULT":       24,
	"SOFTWARE_INFO":         25,
	"BUTTON_PRESSED":        26,
	"STREAMSETUP":           27,
	"STREAMDATA":            28,
	"STREAMCONTROL":         29,
	"BLINKLED":              30,
	"SENSOR_REQUEST":        31,
	"USER":                  32,
	"RADIO_GET_CONFIG":      33,
	"RADIO_SET_CONFIG":      34,
	"STARTAPP":              35,
	"INSTALLED_SOFTWARE":    36,
	"POWERUP":               37,
	"GETCLOUDTOKEN":         38,
	"REQUEST_RADIO_FORWARD": 39,
	"RADIO_FORWARD":         40,
	"SET_CONFIG_FLAG":       41,
	"SET_CONFIG_FLAGS":      42,
	"SET_SENSOR_SERVER":     43,
	"WIFI_INFO":             44,
	"GET_CONFIG_FLAGS":      45,
	"WIFI_SCAN":             46,
	"ROUTE":                 47,
	"SENSOR_LIST":           48,
	"SENSOR_CONFIG":         49,
	"SENSOR_REPORT":         50,
}

func (x ComType) String() string {
	return proto.EnumName(ComType_name, int32(x))
}
func (ComType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ROUTE_COM_TYPE int32

const (
	ROUTE_COM_TYPE_ROUTE_NONE  ROUTE_COM_TYPE = 0
	ROUTE_COM_TYPE_ROUTE_LIST  ROUTE_COM_TYPE = 1
	ROUTE_COM_TYPE_ROUTE_ADD   ROUTE_COM_TYPE = 2
	ROUTE_COM_TYPE_ROUTE_DEL   ROUTE_COM_TYPE = 3
	ROUTE_COM_TYPE_ROUTE_CLEAR ROUTE_COM_TYPE = 4
)

var ROUTE_COM_TYPE_name = map[int32]string{
	0: "ROUTE_NONE",
	1: "ROUTE_LIST",
	2: "ROUTE_ADD",
	3: "ROUTE_DEL",
	4: "ROUTE_CLEAR",
}
var ROUTE_COM_TYPE_value = map[string]int32{
	"ROUTE_NONE":  0,
	"ROUTE_LIST":  1,
	"ROUTE_ADD":   2,
	"ROUTE_DEL":   3,
	"ROUTE_CLEAR": 4,
}

func (x ROUTE_COM_TYPE) String() string {
	return proto.EnumName(ROUTE_COM_TYPE_name, int32(x))
}
func (ROUTE_COM_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ErrorEnum int32

const (
	ErrorEnum_UNSPECIFIED              ErrorEnum = 0
	ErrorEnum_ACCOUNT_UPGRADE_REQUIRED ErrorEnum = 1
)

var ErrorEnum_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "ACCOUNT_UPGRADE_REQUIRED",
}
var ErrorEnum_value = map[string]int32{
	"UNSPECIFIED":              0,
	"ACCOUNT_UPGRADE_REQUIRED": 1,
}

func (x ErrorEnum) String() string {
	return proto.EnumName(ErrorEnum_name, int32(x))
}
func (ErrorEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// matches 'addressing.h' in firmware
type Device int32

const (
	Device_DEVICE_UNSPECIFIED Device = 0
	Device_WIFI               Device = 1
	Device_RADIO              Device = 2
	Device_SERIAL             Device = 3
	Device_USB                Device = 4
	Device_LORA               Device = 5
)

var Device_name = map[int32]string{
	0: "DEVICE_UNSPECIFIED",
	1: "WIFI",
	2: "RADIO",
	3: "SERIAL",
	4: "USB",
	5: "LORA",
}
var Device_value = map[string]int32{
	"DEVICE_UNSPECIFIED": 0,
	"WIFI":               1,
	"RADIO":              2,
	"SERIAL":             3,
	"USB":                4,
	"LORA":               5,
}

func (x Device) String() string {
	return proto.EnumName(Device_name, int32(x))
}
func (Device) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Void should be used when no explicit request or response is required.
type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Command struct {
	Type      ComType  `protobuf:"varint,1,opt,name=Type,enum=singingcat.ComType" json:"Type,omitempty"`
	Args      [][]byte `protobuf:"bytes,2,rep,name=Args,proto3" json:"Args,omitempty"`
	Sender    uint32   `protobuf:"varint,3,opt,name=Sender" json:"Sender,omitempty"`
	Recipient uint32   `protobuf:"varint,4,opt,name=Recipient" json:"Recipient,omitempty"`
	Target    uint32   `protobuf:"varint,5,opt,name=Target" json:"Target,omitempty"`
	Seq       uint32   `protobuf:"varint,6,opt,name=Seq" json:"Seq,omitempty"`
	Flags     uint32   `protobuf:"varint,7,opt,name=Flags" json:"Flags,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Command) GetType() ComType {
	if m != nil {
		return m.Type
	}
	return ComType_UNDEFINED
}

func (m *Command) GetArgs() [][]byte {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Command) GetSender() uint32 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *Command) GetRecipient() uint32 {
	if m != nil {
		return m.Recipient
	}
	return 0
}

func (m *Command) GetTarget() uint32 {
	if m != nil {
		return m.Target
	}
	return 0
}

func (m *Command) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Command) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

// Status represents whether or not the specified operation
// was successful, and what error occurred if it was not.
type Status struct {
	// Success is set to true or false depending on whether or
	// not the operation was successful.
	// In the event of an error, further details can be found
	// in the `ErrorCode` and `ErrorDescription` fields.
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	// ErrorCode is present if an error has occurred during the
	// operation. ErrorCode mappings will be listed in our
	// documentation.
	ErrorCode int32 `protobuf:"varint,2,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
	// ErrorDescription is present if an error has occurred during
	// the operation. This is intended to be human-readable (machines
	// should use the ErrorCode instead).
	ErrorDescription string `protobuf:"bytes,3,opt,name=ErrorDescription" json:"ErrorDescription,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Status) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type ModuleRef struct {
	ModuleID uint64 `protobuf:"varint,1,opt,name=ModuleID" json:"ModuleID,omitempty"`
}

func (m *ModuleRef) Reset()                    { *m = ModuleRef{} }
func (m *ModuleRef) String() string            { return proto.CompactTextString(m) }
func (*ModuleRef) ProtoMessage()               {}
func (*ModuleRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ModuleRef) GetModuleID() uint64 {
	if m != nil {
		return m.ModuleID
	}
	return 0
}

func init() {
	proto.RegisterType((*Void)(nil), "singingcat.Void")
	proto.RegisterType((*Command)(nil), "singingcat.Command")
	proto.RegisterType((*Status)(nil), "singingcat.Status")
	proto.RegisterType((*ModuleRef)(nil), "singingcat.ModuleRef")
	proto.RegisterEnum("singingcat.SensorType", SensorType_name, SensorType_value)
	proto.RegisterEnum("singingcat.ComType", ComType_name, ComType_value)
	proto.RegisterEnum("singingcat.ROUTE_COM_TYPE", ROUTE_COM_TYPE_name, ROUTE_COM_TYPE_value)
	proto.RegisterEnum("singingcat.ErrorEnum", ErrorEnum_name, ErrorEnum_value)
	proto.RegisterEnum("singingcat.Device", Device_name, Device_value)
}

func init() {
	proto.RegisterFile("golang.singingcat.net/apis/singingcat/singingcat.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x55, 0x5d, 0x73, 0xe3, 0x34,
	0x17, 0xde, 0x34, 0xdf, 0x67, 0xdb, 0xee, 0xa9, 0xfa, 0xf1, 0x7a, 0xfb, 0x16, 0x28, 0x85, 0xa5,
	0x25, 0x40, 0x0b, 0xcb, 0x0c, 0x17, 0xdc, 0x29, 0xf6, 0xb1, 0xf1, 0xd4, 0x91, 0x8c, 0x24, 0xb7,
	0x53, 0x6e, 0x3c, 0x21, 0x35, 0x99, 0xcc, 0xb4, 0x71, 0x49, 0x52, 0x66, 0xf8, 0x6f, 0xfc, 0x33,
	0x6e, 0x18, 0xc9, 0x49, 0x13, 0x96, 0x3b, 0x3d, 0x8f, 0x74, 0x3e, 0xf4, 0x9c, 0x73, 0x24, 0xf8,
	0x61, 0x5c, 0x3e, 0x0c, 0xa7, 0xe3, 0xcb, 0xf9, 0x64, 0x3a, 0x9e, 0x4c, 0xc7, 0xa3, 0xe1, 0xe2,
	0x72, 0x5a, 0x2c, 0xae, 0x86, 0x4f, 0x93, 0xf9, 0xd5, 0x9a, 0xdb, 0x58, 0x5e, 0x3e, 0xcd, 0xca,
	0x45, 0xc9, 0x60, 0xcd, 0x9c, 0xb5, 0xa0, 0x71, 0x53, 0x4e, 0xee, 0xcf, 0xfe, 0xaa, 0x41, 0xdb,
	0x2f, 0x1f, 0x1f, 0x87, 0xd3, 0x7b, 0x76, 0x0e, 0x0d, 0xf3, 0xe7, 0x53, 0xe1, 0xd5, 0x4e, 0x6b,
	0x17, 0xbb, 0xef, 0xf7, 0x37, 0xfd, 0xfb, 0xe5, 0xa3, 0xdd, 0x52, 0xee, 0x00, 0x63, 0xd0, 0xe0,
	0xb3, 0xf1, 0xdc, 0xdb, 0x3a, 0xad, 0x5f, 0x6c, 0x2b, 0xb7, 0x66, 0x47, 0xd0, 0xd2, 0xc5, 0xf4,
	0xbe, 0x98, 0x79, 0xf5, 0xd3, 0xda, 0xc5, 0x8e, 0x5a, 0x22, 0x76, 0x02, 0x5d, 0x55, 0x8c, 0x26,
	0x4f, 0x93, 0x62, 0xba, 0xf0, 0x1a, 0x6e, 0x6b, 0x4d, 0x58, 0x2b, 0x33, 0x9c, 0x8d, 0x8b, 0x85,
	0xd7, 0xac, 0xac, 0x2a, 0xc4, 0x10, 0xea, 0xba, 0xf8, 0xdd, 0x6b, 0x39, 0xd2, 0x2e, 0xd9, 0x01,
	0x34, 0xc3, 0x87, 0xe1, 0x78, 0xee, 0xb5, 0x1d, 0x57, 0x81, 0xb3, 0x07, 0x68, 0xe9, 0xc5, 0x70,
	0xf1, 0x3c, 0x67, 0x1e, 0xb4, 0xf5, 0xf3, 0x68, 0x54, 0xcc, 0xe7, 0x2e, 0xff, 0x8e, 0x5a, 0x41,
	0x9b, 0x01, 0xcd, 0x66, 0xe5, 0xcc, 0x2f, 0xef, 0x0b, 0x6f, 0xeb, 0xb4, 0x76, 0xd1, 0x54, 0x6b,
	0x82, 0xf5, 0x00, 0x1d, 0x08, 0x8a, 0xf9, 0x68, 0x36, 0x79, 0x5a, 0x4c, 0xca, 0xa9, 0xbb, 0x41,
	0x57, 0xfd, 0x87, 0x3f, 0x3b, 0x87, 0xee, 0xa0, 0xbc, 0x7f, 0x7e, 0x28, 0x54, 0xf1, 0x1b, 0x3b,
	0x86, 0x4e, 0x05, 0xe2, 0xc0, 0x45, 0x6c, 0xa8, 0x17, 0xdc, 0xeb, 0x03, 0xe8, 0x62, 0x3a, 0x2f,
	0x67, 0x4e, 0x2e, 0x0f, 0x0e, 0x34, 0x09, 0x2d, 0x95, 0xb9, 0x4b, 0x29, 0xcf, 0x44, 0x40, 0x61,
	0x2c, 0x28, 0xc0, 0x57, 0xec, 0x18, 0x8e, 0x36, 0x76, 0x0c, 0x0d, 0x52, 0x52, 0xdc, 0x64, 0x8a,
	0xb0, 0xd6, 0xfb, 0xbb, 0xe5, 0x2a, 0xe3, 0x3c, 0xec, 0x40, 0x77, 0xd3, 0xac, 0x03, 0x0d, 0x21,
	0x65, 0x8a, 0x35, 0xd6, 0x86, 0x7a, 0x22, 0x23, 0xdc, 0xb2, 0x31, 0xb8, 0xa2, 0xfc, 0x4e, 0x66,
	0x39, 0xcf, 0x75, 0x2c, 0xa2, 0x58, 0x44, 0x3e, 0x37, 0x58, 0x67, 0x08, 0xdb, 0x9a, 0x4c, 0x7e,
	0x1b, 0x2b, 0x4a, 0x48, 0x6b, 0x6c, 0xb0, 0xd7, 0xd0, 0xb6, 0x4c, 0x1a, 0x0b, 0x6c, 0xb2, 0x2e,
	0x34, 0x15, 0x69, 0x32, 0xd8, 0xb2, 0x51, 0x6e, 0xe3, 0x30, 0xae, 0x60, 0xdb, 0x42, 0xc5, 0x83,
	0x58, 0xa6, 0xb1, 0x88, 0xb0, 0xc3, 0xf6, 0x60, 0xe7, 0x05, 0x26, 0x36, 0x7a, 0x97, 0x01, 0xb4,
	0x42, 0x45, 0xf4, 0x0b, 0x21, 0xd8, 0x30, 0x49, 0xac, 0x4d, 0x3e, 0x90, 0x41, 0x96, 0x90, 0xc6,
	0xd7, 0xce, 0x40, 0x66, 0x86, 0x72, 0x45, 0x3f, 0x67, 0xa4, 0x0d, 0x6e, 0x5b, 0x03, 0x6d, 0x94,
	0xec, 0x13, 0xee, 0x58, 0xf7, 0x11, 0x99, 0x34, 0xeb, 0x5f, 0xd3, 0x1d, 0xee, 0xda, 0xd3, 0x9a,
	0x8c, 0x9f, 0xc8, 0x2c, 0x30, 0xf2, 0x9a, 0x04, 0xbe, 0x61, 0x6f, 0xe0, 0xb5, 0xcd, 0x33, 0x91,
	0x91, 0xbd, 0x0d, 0xe2, 0x2a, 0xf1, 0x80, 0xfb, 0xb8, 0xc7, 0xf6, 0xe1, 0x8d, 0x05, 0x9a, 0x54,
	0xcc, 0x93, 0x3c, 0x95, 0xca, 0x20, 0xb3, 0x4e, 0xc3, 0x84, 0xeb, 0x9f, 0x72, 0x9e, 0xa6, 0xb8,
	0x6f, 0x3d, 0x0c, 0x62, 0x5f, 0xc9, 0x65, 0xd0, 0x83, 0x0d, 0xa3, 0x1b, 0x52, 0xb9, 0xe0, 0x03,
	0xc2, 0xc3, 0x4a, 0x21, 0x5b, 0x85, 0xfc, 0x86, 0x27, 0x19, 0xe1, 0x11, 0xdb, 0x86, 0x0e, 0x17,
	0x42, 0x66, 0xc2, 0x27, 0xfc, 0x9f, 0x35, 0x0a, 0xb9, 0x6f, 0xa4, 0xba, 0xcb, 0x03, 0x0a, 0x79,
	0x96, 0x18, 0xf4, 0x5c, 0xbe, 0x32, 0x34, 0xb7, 0x56, 0xf5, 0x58, 0x84, 0x12, 0xdf, 0x32, 0x06,
	0xbb, 0xfd, 0xcc, 0x18, 0x29, 0xf2, 0x54, 0x91, 0xd6, 0x14, 0xe0, 0xb1, 0xbb, 0x83, 0x51, 0xc4,
	0x07, 0x9a, 0x4c, 0x96, 0xe2, 0xff, 0xd9, 0x2e, 0x40, 0x45, 0x04, 0xdc, 0x70, 0x3c, 0x71, 0x7e,
	0x1c, 0xf6, 0xa5, 0x30, 0x4a, 0x26, 0xf8, 0x91, 0x8d, 0xde, 0x4f, 0x62, 0x71, 0x9d, 0x50, 0x80,
	0x1f, 0x5b, 0xaf, 0xcb, 0xec, 0x56, 0x3a, 0x7e, 0x62, 0x1b, 0x20, 0xd3, 0xa4, 0xf0, 0x94, 0x1d,
	0x00, 0xba, 0xaa, 0xe4, 0x11, 0x99, 0xdc, 0x97, 0x22, 0x8c, 0x23, 0xfc, 0x74, 0xcd, 0xea, 0x35,
	0x7b, 0x66, 0xfd, 0x6a, 0xc3, 0x95, 0xb1, 0xda, 0x7c, 0xc6, 0x8e, 0x80, 0xc5, 0x42, 0x1b, 0x9e,
	0x24, 0x14, 0xe4, 0xab, 0xab, 0xe0, 0xe7, 0x56, 0xe4, 0x54, 0xde, 0x92, 0xca, 0x52, 0x7c, 0x67,
	0xb3, 0x8b, 0xfe, 0x55, 0x95, 0x2f, 0xd8, 0x5b, 0x38, 0x5c, 0x26, 0x92, 0x57, 0x31, 0x42, 0xa9,
	0x6e, 0xb9, 0x0a, 0xf0, 0xfc, 0xa5, 0x45, 0x5e, 0xa8, 0x8b, 0x95, 0xe0, 0x55, 0x0e, 0x79, 0x98,
	0xf0, 0x08, 0xbf, 0xb4, 0xe9, 0x7d, 0x40, 0x6a, 0xec, 0xb1, 0x43, 0xd8, 0xab, 0x6a, 0xe3, 0x2e,
	0x5b, 0x95, 0x08, 0xbf, 0x5a, 0x75, 0x65, 0x25, 0xf2, 0xd7, 0xd6, 0x36, 0xfa, 0xd0, 0xf6, 0x9b,
	0x97, 0x43, 0xda, 0xe7, 0x02, 0x2f, 0x5d, 0x53, 0xdb, 0xd6, 0xc3, 0xab, 0xaa, 0x89, 0x9c, 0x47,
	0xdb, 0x9e, 0xf8, 0x6d, 0xd5, 0x68, 0x8e, 0x58, 0x0a, 0xf3, 0xdd, 0x06, 0xa5, 0xc8, 0x35, 0xd2,
	0xfb, 0x5e, 0x0e, 0xbb, 0x55, 0xf3, 0xfa, 0x72, 0x90, 0xdb, 0xe9, 0xb4, 0x85, 0xab, 0x18, 0x21,
	0x05, 0xe1, 0xab, 0x35, 0x76, 0x7e, 0x6b, 0x6e, 0x5c, 0x1c, 0xe6, 0x41, 0x80, 0x5b, 0x6b, 0x18,
	0x50, 0x82, 0x75, 0x9b, 0xc6, 0xd2, 0x5f, 0x42, 0x5c, 0x61, 0xa3, 0xf7, 0xe3, 0xf2, 0x55, 0xa2,
	0xe9, 0xf3, 0xa3, 0xdd, 0xcd, 0x84, 0x4e, 0xc9, 0x8f, 0xc3, 0xd8, 0x4d, 0xf8, 0x09, 0x78, 0xdc,
	0xf7, 0x65, 0x26, 0x4c, 0x9e, 0xa5, 0x91, 0xe2, 0x41, 0x35, 0x45, 0xb1, 0xa2, 0x00, 0x6b, 0x3d,
	0x03, 0xad, 0xa0, 0xf8, 0x63, 0x32, 0x2a, 0x6c, 0x11, 0x03, 0xba, 0x89, 0x7d, 0xfb, 0xac, 0x6c,
	0xda, 0x77, 0xa0, 0x61, 0xf5, 0xc0, 0x9a, 0x93, 0xc2, 0xd6, 0x04, 0xb7, 0xdc, 0xf4, 0xb9, 0x69,
	0xc1, 0xba, 0x7d, 0x38, 0x32, 0xdd, 0xc7, 0x86, 0x3d, 0x99, 0x48, 0xc5, 0xb1, 0xd9, 0x3f, 0x87,
	0x77, 0xd3, 0x62, 0xb1, 0xf9, 0xea, 0x2f, 0xff, 0x19, 0xfb, 0xb1, 0x6c, 0xd0, 0xbf, 0xb6, 0xdc,
	0x77, 0xf2, 0xfd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0xc0, 0xb2, 0x17, 0x88, 0x06, 0x00,
	0x00,
}
