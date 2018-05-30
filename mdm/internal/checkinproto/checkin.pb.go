// Code generated by protoc-gen-go. DO NOT EDIT.
// source: checkin.proto

package checkinproto

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

type Event struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Time                 int64             `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Command              *Command          `protobuf:"bytes,3,opt,name=command" json:"command,omitempty"`
	Raw                  []byte            `protobuf:"bytes,4,opt,name=raw,proto3" json:"raw,omitempty"`
	Params               map[string]string `protobuf:"bytes,5,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_checkin_7e1d11c144d2ce7a, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Event) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *Event) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Event) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Command struct {
	MessageType          string        `protobuf:"bytes,1,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	Topic                string        `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	Udid                 string        `protobuf:"bytes,3,opt,name=udid" json:"udid,omitempty"`
	Authenticate         *Authenticate `protobuf:"bytes,4,opt,name=authenticate" json:"authenticate,omitempty"`
	TokenUpdate          *TokenUpdate  `protobuf:"bytes,5,opt,name=token_update,json=tokenUpdate" json:"token_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_checkin_7e1d11c144d2ce7a, []int{1}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *Command) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Command) GetUdid() string {
	if m != nil {
		return m.Udid
	}
	return ""
}

func (m *Command) GetAuthenticate() *Authenticate {
	if m != nil {
		return m.Authenticate
	}
	return nil
}

func (m *Command) GetTokenUpdate() *TokenUpdate {
	if m != nil {
		return m.TokenUpdate
	}
	return nil
}

type Authenticate struct {
	OsVersion            string   `protobuf:"bytes,1,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	BuildVersion         string   `protobuf:"bytes,2,opt,name=build_version,json=buildVersion" json:"build_version,omitempty"`
	ProductName          string   `protobuf:"bytes,3,opt,name=product_name,json=productName" json:"product_name,omitempty"`
	SerialNumber         string   `protobuf:"bytes,4,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	Imei                 string   `protobuf:"bytes,5,opt,name=imei" json:"imei,omitempty"`
	Meid                 string   `protobuf:"bytes,6,opt,name=meid" json:"meid,omitempty"`
	DeviceName           string   `protobuf:"bytes,7,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
	Challenge            []byte   `protobuf:"bytes,8,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Model                string   `protobuf:"bytes,9,opt,name=model" json:"model,omitempty"`
	ModelName            string   `protobuf:"bytes,10,opt,name=model_name,json=modelName" json:"model_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Authenticate) Reset()         { *m = Authenticate{} }
func (m *Authenticate) String() string { return proto.CompactTextString(m) }
func (*Authenticate) ProtoMessage()    {}
func (*Authenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_checkin_7e1d11c144d2ce7a, []int{2}
}
func (m *Authenticate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Authenticate.Unmarshal(m, b)
}
func (m *Authenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Authenticate.Marshal(b, m, deterministic)
}
func (dst *Authenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authenticate.Merge(dst, src)
}
func (m *Authenticate) XXX_Size() int {
	return xxx_messageInfo_Authenticate.Size(m)
}
func (m *Authenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_Authenticate.DiscardUnknown(m)
}

var xxx_messageInfo_Authenticate proto.InternalMessageInfo

func (m *Authenticate) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *Authenticate) GetBuildVersion() string {
	if m != nil {
		return m.BuildVersion
	}
	return ""
}

func (m *Authenticate) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *Authenticate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Authenticate) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *Authenticate) GetMeid() string {
	if m != nil {
		return m.Meid
	}
	return ""
}

func (m *Authenticate) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *Authenticate) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *Authenticate) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *Authenticate) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

type TokenUpdate struct {
	Token                 []byte   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	PushMagic             string   `protobuf:"bytes,2,opt,name=push_magic,json=pushMagic" json:"push_magic,omitempty"`
	UnlockToken           []byte   `protobuf:"bytes,3,opt,name=unlock_token,json=unlockToken,proto3" json:"unlock_token,omitempty"`
	AwaitingConfiguration bool     `protobuf:"varint,4,opt,name=awaiting_configuration,json=awaitingConfiguration" json:"awaiting_configuration,omitempty"`
	UserId                string   `protobuf:"bytes,5,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserLongName          string   `protobuf:"bytes,6,opt,name=user_long_name,json=userLongName" json:"user_long_name,omitempty"`
	UserShortName         string   `protobuf:"bytes,7,opt,name=user_short_name,json=userShortName" json:"user_short_name,omitempty"`
	NotOnConsole          bool     `protobuf:"varint,8,opt,name=not_on_console,json=notOnConsole" json:"not_on_console,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *TokenUpdate) Reset()         { *m = TokenUpdate{} }
func (m *TokenUpdate) String() string { return proto.CompactTextString(m) }
func (*TokenUpdate) ProtoMessage()    {}
func (*TokenUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_checkin_7e1d11c144d2ce7a, []int{3}
}
func (m *TokenUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenUpdate.Unmarshal(m, b)
}
func (m *TokenUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenUpdate.Marshal(b, m, deterministic)
}
func (dst *TokenUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenUpdate.Merge(dst, src)
}
func (m *TokenUpdate) XXX_Size() int {
	return xxx_messageInfo_TokenUpdate.Size(m)
}
func (m *TokenUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_TokenUpdate proto.InternalMessageInfo

func (m *TokenUpdate) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *TokenUpdate) GetPushMagic() string {
	if m != nil {
		return m.PushMagic
	}
	return ""
}

func (m *TokenUpdate) GetUnlockToken() []byte {
	if m != nil {
		return m.UnlockToken
	}
	return nil
}

func (m *TokenUpdate) GetAwaitingConfiguration() bool {
	if m != nil {
		return m.AwaitingConfiguration
	}
	return false
}

func (m *TokenUpdate) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TokenUpdate) GetUserLongName() string {
	if m != nil {
		return m.UserLongName
	}
	return ""
}

func (m *TokenUpdate) GetUserShortName() string {
	if m != nil {
		return m.UserShortName
	}
	return ""
}

func (m *TokenUpdate) GetNotOnConsole() bool {
	if m != nil {
		return m.NotOnConsole
	}
	return false
}

func init() {
	proto.RegisterType((*Event)(nil), "checkinproto.Event")
	proto.RegisterMapType((map[string]string)(nil), "checkinproto.Event.ParamsEntry")
	proto.RegisterType((*Command)(nil), "checkinproto.Command")
	proto.RegisterType((*Authenticate)(nil), "checkinproto.Authenticate")
	proto.RegisterType((*TokenUpdate)(nil), "checkinproto.TokenUpdate")
}

func init() { proto.RegisterFile("checkin.proto", fileDescriptor_checkin_7e1d11c144d2ce7a) }

var fileDescriptor_checkin_7e1d11c144d2ce7a = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0x74, 0x6d, 0x97, 0x93, 0x6c, 0x20, 0x8b, 0x41, 0x98, 0x40, 0xeb, 0xca, 0x84,
	0x7a, 0x55, 0xa4, 0x21, 0xc4, 0x1f, 0x21, 0x24, 0x34, 0xed, 0x02, 0x09, 0x06, 0x32, 0x83, 0xdb,
	0xc8, 0x4d, 0x4c, 0x6a, 0x35, 0xb1, 0xa3, 0xc4, 0xe9, 0xd4, 0x47, 0xe2, 0x05, 0x78, 0x12, 0x9e,
	0x80, 0x27, 0x41, 0xe7, 0x38, 0xdd, 0xd2, 0xbb, 0x73, 0x7e, 0xfe, 0xfc, 0xe5, 0x9c, 0x2f, 0x86,
	0x83, 0x74, 0x29, 0xd3, 0x95, 0xd2, 0xf3, 0xaa, 0x36, 0xd6, 0xb0, 0xa8, 0x6b, 0xa9, 0x9b, 0xfe,
	0xf3, 0x60, 0x78, 0xb9, 0x96, 0xda, 0xb2, 0x43, 0xf0, 0x55, 0x16, 0x7b, 0x13, 0x6f, 0x16, 0x70,
	0x5f, 0x65, 0x8c, 0xc1, 0x9e, 0x55, 0xa5, 0x8c, 0xfd, 0x89, 0x37, 0x1b, 0x70, 0xaa, 0xd9, 0x0b,
	0x18, 0xa7, 0xa6, 0x2c, 0x85, 0xce, 0xe2, 0xc1, 0xc4, 0x9b, 0x85, 0xe7, 0x47, 0xf3, 0xbe, 0xdb,
	0xfc, 0xc2, 0x1d, 0xf2, 0xad, 0x8a, 0xdd, 0x87, 0x41, 0x2d, 0x6e, 0xe2, 0xbd, 0x89, 0x37, 0x8b,
	0x38, 0x96, 0xec, 0x35, 0x8c, 0x2a, 0x51, 0x8b, 0xb2, 0x89, 0x87, 0x93, 0xc1, 0x2c, 0x3c, 0x3f,
	0xd9, 0x75, 0xa0, 0x59, 0xe6, 0xdf, 0x48, 0x71, 0xa9, 0x6d, 0xbd, 0xe1, 0x9d, 0xfc, 0xf8, 0x2d,
	0x84, 0x3d, 0x8c, 0xce, 0x2b, 0xb9, 0xe9, 0xe6, 0xc5, 0x92, 0x3d, 0x80, 0xe1, 0x5a, 0x14, 0xad,
	0x9b, 0x38, 0xe0, 0xae, 0x79, 0xe7, 0xbf, 0xf1, 0xa6, 0x7f, 0x3d, 0x18, 0x77, 0xa3, 0xb1, 0x53,
	0x88, 0x4a, 0xd9, 0x34, 0x22, 0x97, 0x89, 0xdd, 0x54, 0xb2, 0x33, 0x08, 0x3b, 0x76, 0xbd, 0xa9,
	0x24, 0x1a, 0x59, 0x53, 0xa9, 0x74, 0x6b, 0x44, 0x0d, 0xe6, 0xd1, 0x66, 0xca, 0x2d, 0x1e, 0x70,
	0xaa, 0xd9, 0x07, 0x88, 0x44, 0x6b, 0x97, 0x52, 0x5b, 0x95, 0x0a, 0x2b, 0x69, 0xcf, 0xf0, 0xfc,
	0x78, 0x77, 0xa5, 0x8f, 0x3d, 0x05, 0xdf, 0xd1, 0xb3, 0xf7, 0x10, 0x59, 0xb3, 0x92, 0x3a, 0x69,
	0xab, 0x0c, 0xef, 0x0f, 0xe9, 0xfe, 0xe3, 0xdd, 0xfb, 0xd7, 0xa8, 0xf8, 0x41, 0x02, 0x1e, 0xda,
	0xbb, 0x66, 0xfa, 0xc7, 0x87, 0xa8, 0x6f, 0xce, 0x9e, 0x02, 0x98, 0x26, 0x59, 0xcb, 0xba, 0x51,
	0x46, 0x77, 0x9b, 0x05, 0xa6, 0xf9, 0xe9, 0x00, 0x7b, 0x06, 0x07, 0x8b, 0x56, 0x15, 0xd9, 0xad,
	0xc2, 0xed, 0x17, 0x11, 0xdc, 0x8a, 0x4e, 0x21, 0xaa, 0x6a, 0x93, 0xb5, 0xa9, 0x4d, 0xb4, 0x28,
	0x65, 0xb7, 0x6e, 0xd8, 0xb1, 0x2b, 0x51, 0x4a, 0xf4, 0x69, 0x64, 0xad, 0x44, 0x91, 0xe8, 0xb6,
	0x5c, 0xc8, 0x9a, 0xd6, 0x0e, 0x78, 0xe4, 0xe0, 0x15, 0x31, 0x8c, 0x4b, 0x95, 0x52, 0xd1, 0x4a,
	0x01, 0xa7, 0x1a, 0x59, 0x29, 0x55, 0x16, 0x8f, 0x1c, 0xc3, 0x9a, 0x9d, 0x40, 0x98, 0xc9, 0xb5,
	0x4a, 0xa5, 0xfb, 0xdc, 0x98, 0x8e, 0xc0, 0x21, 0xfa, 0xda, 0x13, 0x08, 0xd2, 0xa5, 0x28, 0x0a,
	0xa9, 0x73, 0x19, 0xef, 0xd3, 0x43, 0xba, 0x03, 0xf8, 0xaf, 0x4a, 0x93, 0xc9, 0x22, 0x0e, 0xdc,
	0xbf, 0xa2, 0x06, 0x83, 0xa0, 0xc2, 0x79, 0x82, 0x0b, 0x82, 0x08, 0x5a, 0x4e, 0x7f, 0xfb, 0x10,
	0xf6, 0x52, 0x75, 0x3f, 0x7c, 0x25, 0x5d, 0x64, 0x11, 0x77, 0x0d, 0x9a, 0x54, 0x6d, 0xb3, 0x4c,
	0x4a, 0x91, 0xdf, 0xbe, 0x85, 0x00, 0xc9, 0x17, 0x04, 0x18, 0x54, 0xab, 0x0b, 0x93, 0xae, 0x12,
	0x77, 0x77, 0x40, 0x77, 0x43, 0xc7, 0xc8, 0x9d, 0xbd, 0x82, 0x87, 0xe2, 0x46, 0x28, 0xab, 0x74,
	0x9e, 0xa4, 0x46, 0xff, 0x52, 0x79, 0x5b, 0x0b, 0x8b, 0xc9, 0x63, 0x62, 0xfb, 0xfc, 0x68, 0x7b,
	0x7a, 0xd1, 0x3f, 0x64, 0x8f, 0x60, 0xdc, 0x36, 0xb2, 0x4e, 0x54, 0xd6, 0xa5, 0x37, 0xc2, 0xf6,
	0x53, 0xc6, 0xce, 0xe0, 0x90, 0x0e, 0x0a, 0xa3, 0x73, 0xb7, 0x9a, 0x4b, 0x32, 0x42, 0xfa, 0xd9,
	0xe8, 0x9c, 0x02, 0x7b, 0x0e, 0xf7, 0x48, 0xd5, 0x2c, 0x4d, 0x6d, 0xfb, 0xa9, 0x1e, 0x20, 0xfe,
	0x8e, 0x94, 0x74, 0x67, 0x70, 0xa8, 0x8d, 0x4d, 0x8c, 0xc6, 0xd9, 0x1a, 0x53, 0xb8, 0x74, 0xf7,
	0x79, 0xa4, 0x8d, 0xfd, 0xaa, 0x2f, 0x1c, 0x5b, 0x8c, 0xe8, 0x11, 0xbe, 0xfc, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0x87, 0xd1, 0x6a, 0xe1, 0x46, 0x04, 0x00, 0x00,
}