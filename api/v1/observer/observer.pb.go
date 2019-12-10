// Code generated by protoc-gen-go. DO NOT EDIT.
// source: observer/observer.proto

package observer

import (
	context "context"
	fmt "fmt"
	flow "github.com/cilium/hubble/api/v1/flow"
	types "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// Flow from public import flow/flow.proto
type Flow = flow.Flow

// Layer4 from public import flow/flow.proto
type Layer4 = flow.Layer4
type Layer4_TCP = flow.Layer4_TCP
type Layer4_UDP = flow.Layer4_UDP
type Layer4_ICMPv4 = flow.Layer4_ICMPv4
type Layer4_ICMPv6 = flow.Layer4_ICMPv6

// Layer7 from public import flow/flow.proto
type Layer7 = flow.Layer7
type Layer7_Dns = flow.Layer7_Dns
type Layer7_Http = flow.Layer7_Http
type Layer7_Kafka = flow.Layer7_Kafka

// Endpoint from public import flow/flow.proto
type Endpoint = flow.Endpoint

// TCP from public import flow/flow.proto
type TCP = flow.TCP

// IP from public import flow/flow.proto
type IP = flow.IP

// Ethernet from public import flow/flow.proto
type Ethernet = flow.Ethernet

// TCPFlags from public import flow/flow.proto
type TCPFlags = flow.TCPFlags

// UDP from public import flow/flow.proto
type UDP = flow.UDP

// ICMPv4 from public import flow/flow.proto
type ICMPv4 = flow.ICMPv4

// ICMPv6 from public import flow/flow.proto
type ICMPv6 = flow.ICMPv6

// EventTypeFilter from public import flow/flow.proto
type EventTypeFilter = flow.EventTypeFilter

// CiliumEventType from public import flow/flow.proto
type CiliumEventType = flow.CiliumEventType

// FlowFilter from public import flow/flow.proto
type FlowFilter = flow.FlowFilter

// Payload from public import flow/flow.proto
type Payload = flow.Payload

// DNS from public import flow/flow.proto
type DNS = flow.DNS

// HTTPHeader from public import flow/flow.proto
type HTTPHeader = flow.HTTPHeader

// HTTP from public import flow/flow.proto
type HTTP = flow.HTTP

// Kafka from public import flow/flow.proto
type Kafka = flow.Kafka

// Service from public import flow/flow.proto
type Service = flow.Service

// FlowType from public import flow/flow.proto
type FlowType = flow.FlowType

var FlowType_name = flow.FlowType_name
var FlowType_value = flow.FlowType_value

const FlowType_UNKNOWN_TYPE = FlowType(flow.FlowType_UNKNOWN_TYPE)
const FlowType_L3_L4 = FlowType(flow.FlowType_L3_L4)
const FlowType_L7 = FlowType(flow.FlowType_L7)

// L7FlowType from public import flow/flow.proto
type L7FlowType = flow.L7FlowType

var L7FlowType_name = flow.L7FlowType_name
var L7FlowType_value = flow.L7FlowType_value

const L7FlowType_UNKNOWN_L7_TYPE = L7FlowType(flow.L7FlowType_UNKNOWN_L7_TYPE)
const L7FlowType_REQUEST = L7FlowType(flow.L7FlowType_REQUEST)
const L7FlowType_RESPONSE = L7FlowType(flow.L7FlowType_RESPONSE)
const L7FlowType_SAMPLE = L7FlowType(flow.L7FlowType_SAMPLE)

// IPVersion from public import flow/flow.proto
type IPVersion = flow.IPVersion

var IPVersion_name = flow.IPVersion_name
var IPVersion_value = flow.IPVersion_value

const IPVersion_IP_NOT_USED = IPVersion(flow.IPVersion_IP_NOT_USED)
const IPVersion_IPv4 = IPVersion(flow.IPVersion_IPv4)
const IPVersion_IPv6 = IPVersion(flow.IPVersion_IPv6)

// Verdict from public import flow/flow.proto
type Verdict = flow.Verdict

var Verdict_name = flow.Verdict_name
var Verdict_value = flow.Verdict_value

const Verdict_VERDICT_UNKNOWN = Verdict(flow.Verdict_VERDICT_UNKNOWN)
const Verdict_FORWARDED = Verdict(flow.Verdict_FORWARDED)
const Verdict_DROPPED = Verdict(flow.Verdict_DROPPED)

// EventType from public import flow/flow.proto
type EventType = flow.EventType

var EventType_name = flow.EventType_name
var EventType_value = flow.EventType_value

const EventType_UNKNOWN = EventType(flow.EventType_UNKNOWN)
const EventType_EventSample = EventType(flow.EventType_EventSample)
const EventType_RecordLost = EventType(flow.EventType_RecordLost)

// ProtocolMessageType enumerates the type of messages the server can return
// to the client.
type ProtocolMessageType int32

const (
	ProtocolMessageType_UNKNOWN_PROTOCOL_MESSAGE_TYPE  ProtocolMessageType = 0
	ProtocolMessageType_PROGRESS_PROTOCOL_MESSAGE_TYPE ProtocolMessageType = 1
	ProtocolMessageType_ERROR_PROTOCOL_MESSAGE_TYPE    ProtocolMessageType = 2
)

var ProtocolMessageType_name = map[int32]string{
	0: "UNKNOWN_PROTOCOL_MESSAGE_TYPE",
	1: "PROGRESS_PROTOCOL_MESSAGE_TYPE",
	2: "ERROR_PROTOCOL_MESSAGE_TYPE",
}

var ProtocolMessageType_value = map[string]int32{
	"UNKNOWN_PROTOCOL_MESSAGE_TYPE":  0,
	"PROGRESS_PROTOCOL_MESSAGE_TYPE": 1,
	"ERROR_PROTOCOL_MESSAGE_TYPE":    2,
}

func (x ProtocolMessageType) String() string {
	return proto.EnumName(ProtocolMessageType_name, int32(x))
}

func (ProtocolMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{0}
}

type ServerStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStatusRequest) Reset()         { *m = ServerStatusRequest{} }
func (m *ServerStatusRequest) String() string { return proto.CompactTextString(m) }
func (*ServerStatusRequest) ProtoMessage()    {}
func (*ServerStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{0}
}

func (m *ServerStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStatusRequest.Unmarshal(m, b)
}
func (m *ServerStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStatusRequest.Marshal(b, m, deterministic)
}
func (m *ServerStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStatusRequest.Merge(m, src)
}
func (m *ServerStatusRequest) XXX_Size() int {
	return xxx_messageInfo_ServerStatusRequest.Size(m)
}
func (m *ServerStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStatusRequest proto.InternalMessageInfo

type ServerStatusResponse struct {
	// number of currently captured flows
	NumFlows uint64 `protobuf:"varint,1,opt,name=num_flows,json=numFlows,proto3" json:"num_flows,omitempty"`
	// maximum capacity of the ring buffer
	MaxFlows             uint64   `protobuf:"varint,2,opt,name=max_flows,json=maxFlows,proto3" json:"max_flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerStatusResponse) Reset()         { *m = ServerStatusResponse{} }
func (m *ServerStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ServerStatusResponse) ProtoMessage()    {}
func (*ServerStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{1}
}

func (m *ServerStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerStatusResponse.Unmarshal(m, b)
}
func (m *ServerStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerStatusResponse.Marshal(b, m, deterministic)
}
func (m *ServerStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerStatusResponse.Merge(m, src)
}
func (m *ServerStatusResponse) XXX_Size() int {
	return xxx_messageInfo_ServerStatusResponse.Size(m)
}
func (m *ServerStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerStatusResponse proto.InternalMessageInfo

func (m *ServerStatusResponse) GetNumFlows() uint64 {
	if m != nil {
		return m.NumFlows
	}
	return 0
}

func (m *ServerStatusResponse) GetMaxFlows() uint64 {
	if m != nil {
		return m.MaxFlows
	}
	return 0
}

type GetFlowsRequest struct {
	// Number of flows that should be returned. Incompatible with `since/until`.
	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// follow sets when the server should continue to stream flows after
	// printing the last N flows.
	Follow bool `protobuf:"varint,3,opt,name=follow,proto3" json:"follow,omitempty"`
	// blacklist defines a list of filters which have to match for a flow to be
	// excluded from the result.
	// If multiple blacklist filters are specified, only one of them has to
	// match for a flow to be excluded.
	Blacklist []*flow.FlowFilter `protobuf:"bytes,5,rep,name=blacklist,proto3" json:"blacklist,omitempty"`
	// whitelist defines a list of filters which have to match for a flow to be
	// included in the result.
	// If multiple whitelist filters are specified, only one of them has to
	// match for a flow to be included.
	// The whitelist and blacklist can both be specified. In such cases, the
	// set of the returned flows is the set difference `whitelist - blacklist`.
	// In other words, the result will contain all flows matched by the
	// whitelist that are not also simultaneously matched by the blacklist.
	Whitelist []*flow.FlowFilter `protobuf:"bytes,6,rep,name=whitelist,proto3" json:"whitelist,omitempty"`
	// Since this time for returned flows. Incompatible with `number`.
	Since *types.Timestamp `protobuf:"bytes,7,opt,name=since,proto3" json:"since,omitempty"`
	// Until this time for returned flows. Incompatible with `number`.
	Until *types.Timestamp `protobuf:"bytes,8,opt,name=until,proto3" json:"until,omitempty"`
	// Attach the original binary payload to the flow.
	IncludeBinaryPayload bool     `protobuf:"varint,9,opt,name=includeBinaryPayload,proto3" json:"includeBinaryPayload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlowsRequest) Reset()         { *m = GetFlowsRequest{} }
func (m *GetFlowsRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlowsRequest) ProtoMessage()    {}
func (*GetFlowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{2}
}

func (m *GetFlowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlowsRequest.Unmarshal(m, b)
}
func (m *GetFlowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlowsRequest.Marshal(b, m, deterministic)
}
func (m *GetFlowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlowsRequest.Merge(m, src)
}
func (m *GetFlowsRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlowsRequest.Size(m)
}
func (m *GetFlowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlowsRequest proto.InternalMessageInfo

func (m *GetFlowsRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GetFlowsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

func (m *GetFlowsRequest) GetBlacklist() []*flow.FlowFilter {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *GetFlowsRequest) GetWhitelist() []*flow.FlowFilter {
	if m != nil {
		return m.Whitelist
	}
	return nil
}

func (m *GetFlowsRequest) GetSince() *types.Timestamp {
	if m != nil {
		return m.Since
	}
	return nil
}

func (m *GetFlowsRequest) GetUntil() *types.Timestamp {
	if m != nil {
		return m.Until
	}
	return nil
}

func (m *GetFlowsRequest) GetIncludeBinaryPayload() bool {
	if m != nil {
		return m.IncludeBinaryPayload
	}
	return false
}

// GetFlowsResponse contains either a flow or a protocol message.
type GetFlowsResponse struct {
	// Types that are valid to be assigned to ResponseTypes:
	//	*GetFlowsResponse_Flow
	//	*GetFlowsResponse_ServerMsg
	ResponseTypes isGetFlowsResponse_ResponseTypes `protobuf_oneof:"response_types"`
	// Name of the node where this event was observed.
	NodeName string `protobuf:"bytes,1000,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Timestamp at which this event was observed.
	Time                 *types.Timestamp `protobuf:"bytes,1001,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetFlowsResponse) Reset()         { *m = GetFlowsResponse{} }
func (m *GetFlowsResponse) String() string { return proto.CompactTextString(m) }
func (*GetFlowsResponse) ProtoMessage()    {}
func (*GetFlowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{3}
}

func (m *GetFlowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlowsResponse.Unmarshal(m, b)
}
func (m *GetFlowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlowsResponse.Marshal(b, m, deterministic)
}
func (m *GetFlowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlowsResponse.Merge(m, src)
}
func (m *GetFlowsResponse) XXX_Size() int {
	return xxx_messageInfo_GetFlowsResponse.Size(m)
}
func (m *GetFlowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlowsResponse proto.InternalMessageInfo

type isGetFlowsResponse_ResponseTypes interface {
	isGetFlowsResponse_ResponseTypes()
}

type GetFlowsResponse_Flow struct {
	Flow *flow.Flow `protobuf:"bytes,1,opt,name=flow,proto3,oneof"`
}

type GetFlowsResponse_ServerMsg struct {
	ServerMsg *ProtocolMessage `protobuf:"bytes,2,opt,name=server_msg,json=serverMsg,proto3,oneof"`
}

func (*GetFlowsResponse_Flow) isGetFlowsResponse_ResponseTypes() {}

func (*GetFlowsResponse_ServerMsg) isGetFlowsResponse_ResponseTypes() {}

func (m *GetFlowsResponse) GetResponseTypes() isGetFlowsResponse_ResponseTypes {
	if m != nil {
		return m.ResponseTypes
	}
	return nil
}

func (m *GetFlowsResponse) GetFlow() *flow.Flow {
	if x, ok := m.GetResponseTypes().(*GetFlowsResponse_Flow); ok {
		return x.Flow
	}
	return nil
}

func (m *GetFlowsResponse) GetServerMsg() *ProtocolMessage {
	if x, ok := m.GetResponseTypes().(*GetFlowsResponse_ServerMsg); ok {
		return x.ServerMsg
	}
	return nil
}

func (m *GetFlowsResponse) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *GetFlowsResponse) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetFlowsResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetFlowsResponse_Flow)(nil),
		(*GetFlowsResponse_ServerMsg)(nil),
	}
}

// ProtocolMessage is a type of message type that can sent by the server to
// inform the client with its current status.
type ProtocolMessage struct {
	// Types that are valid to be assigned to Msg:
	//	*ProtocolMessage_Info
	Msg                  isProtocolMessage_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProtocolMessage) Reset()         { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()    {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{4}
}

func (m *ProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessage.Unmarshal(m, b)
}
func (m *ProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessage.Marshal(b, m, deterministic)
}
func (m *ProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessage.Merge(m, src)
}
func (m *ProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessage.Size(m)
}
func (m *ProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessage proto.InternalMessageInfo

type isProtocolMessage_Msg interface {
	isProtocolMessage_Msg()
}

type ProtocolMessage_Info struct {
	Info *ProtocolMessageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

func (*ProtocolMessage_Info) isProtocolMessage_Msg() {}

func (m *ProtocolMessage) GetMsg() isProtocolMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ProtocolMessage) GetInfo() *ProtocolMessageInfo {
	if x, ok := m.GetMsg().(*ProtocolMessage_Info); ok {
		return x.Info
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ProtocolMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ProtocolMessage_Info)(nil),
	}
}

// ProtocolMessageInfo is an informational message type that can sent by the server.
type ProtocolMessageInfo struct {
	Msg                  string              `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Type                 ProtocolMessageType `protobuf:"varint,2,opt,name=type,proto3,enum=observer.ProtocolMessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProtocolMessageInfo) Reset()         { *m = ProtocolMessageInfo{} }
func (m *ProtocolMessageInfo) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessageInfo) ProtoMessage()    {}
func (*ProtocolMessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3004233a4a5969ce, []int{5}
}

func (m *ProtocolMessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessageInfo.Unmarshal(m, b)
}
func (m *ProtocolMessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessageInfo.Marshal(b, m, deterministic)
}
func (m *ProtocolMessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessageInfo.Merge(m, src)
}
func (m *ProtocolMessageInfo) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessageInfo.Size(m)
}
func (m *ProtocolMessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessageInfo proto.InternalMessageInfo

func (m *ProtocolMessageInfo) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ProtocolMessageInfo) GetType() ProtocolMessageType {
	if m != nil {
		return m.Type
	}
	return ProtocolMessageType_UNKNOWN_PROTOCOL_MESSAGE_TYPE
}

func init() {
	proto.RegisterEnum("observer.ProtocolMessageType", ProtocolMessageType_name, ProtocolMessageType_value)
	proto.RegisterType((*ServerStatusRequest)(nil), "observer.ServerStatusRequest")
	proto.RegisterType((*ServerStatusResponse)(nil), "observer.ServerStatusResponse")
	proto.RegisterType((*GetFlowsRequest)(nil), "observer.GetFlowsRequest")
	proto.RegisterType((*GetFlowsResponse)(nil), "observer.GetFlowsResponse")
	proto.RegisterType((*ProtocolMessage)(nil), "observer.ProtocolMessage")
	proto.RegisterType((*ProtocolMessageInfo)(nil), "observer.ProtocolMessageInfo")
}

func init() { proto.RegisterFile("observer/observer.proto", fileDescriptor_3004233a4a5969ce) }

var fileDescriptor_3004233a4a5969ce = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x72, 0xd3, 0x30,
	0x10, 0x8d, 0xdb, 0xb4, 0xd8, 0x1b, 0xa6, 0xcd, 0xa8, 0x05, 0x4c, 0x4a, 0xdb, 0xe0, 0x53, 0x86,
	0x83, 0x53, 0xdc, 0x1b, 0x37, 0xca, 0xb8, 0x2d, 0x03, 0x89, 0x3d, 0x72, 0x18, 0x06, 0x2e, 0x1e,
	0x27, 0x55, 0x8c, 0x07, 0x5b, 0x0a, 0x96, 0x4c, 0x9b, 0x03, 0x3f, 0xc3, 0x5f, 0xf0, 0x19, 0xfc,
	0x05, 0xfc, 0x05, 0x23, 0xd9, 0x6e, 0x43, 0x49, 0xe8, 0xc5, 0xa3, 0xdd, 0xf7, 0x76, 0xb5, 0xfb,
	0x76, 0x65, 0x78, 0xc4, 0xc6, 0x9c, 0xe4, 0x5f, 0x49, 0xde, 0xaf, 0x0f, 0xf6, 0x2c, 0x67, 0x82,
	0x21, 0xbd, 0xb6, 0x3b, 0x87, 0x31, 0x63, 0x71, 0x4a, 0xfa, 0xca, 0x3f, 0x2e, 0xa6, 0x7d, 0x91,
	0x64, 0x84, 0x8b, 0x28, 0x9b, 0x95, 0xd4, 0xce, 0xf6, 0x34, 0x65, 0x97, 0x7d, 0xf9, 0x29, 0x1d,
	0xd6, 0x03, 0xd8, 0x09, 0x54, 0x6c, 0x20, 0x22, 0x51, 0x70, 0x4c, 0xbe, 0x14, 0x84, 0x0b, 0xcb,
	0x87, 0xdd, 0xbf, 0xdd, 0x7c, 0xc6, 0x28, 0x27, 0x68, 0x0f, 0x0c, 0x5a, 0x64, 0xa1, 0x4c, 0xc0,
	0x4d, 0xad, 0xab, 0xf5, 0x9a, 0x58, 0xa7, 0x45, 0x76, 0x2a, 0x6d, 0x09, 0x66, 0xd1, 0x55, 0x05,
	0xae, 0x95, 0x60, 0x16, 0x5d, 0x29, 0xd0, 0xfa, 0xb1, 0x06, 0xdb, 0x67, 0x44, 0x28, 0xa3, 0xba,
	0x05, 0x3d, 0x84, 0x4d, 0x5a, 0x64, 0x63, 0x92, 0x57, 0xa9, 0x2a, 0x4b, 0xfa, 0xa7, 0x2c, 0x4d,
	0xd9, 0xa5, 0xb9, 0xde, 0xd5, 0x7a, 0x3a, 0xae, 0x2c, 0x64, 0x83, 0x31, 0x4e, 0xa3, 0xc9, 0xe7,
	0x34, 0xe1, 0xc2, 0xdc, 0xe8, 0xae, 0xf7, 0x5a, 0x4e, 0xdb, 0x56, 0xcd, 0xc8, 0xb4, 0xa7, 0x49,
	0x2a, 0x48, 0x8e, 0x6f, 0x28, 0x92, 0x7f, 0xf9, 0x29, 0x11, 0x44, 0xf1, 0x37, 0x57, 0xf1, 0xaf,
	0x29, 0xe8, 0x08, 0x36, 0x78, 0x42, 0x27, 0xc4, 0xbc, 0xd7, 0xd5, 0x7a, 0x2d, 0xa7, 0x63, 0x97,
	0x72, 0xda, 0xb5, 0x9c, 0xf6, 0xa8, 0x96, 0x13, 0x97, 0x44, 0x19, 0x51, 0x50, 0x91, 0xa4, 0xa6,
	0x7e, 0x77, 0x84, 0x22, 0x22, 0x07, 0x76, 0x13, 0x3a, 0x49, 0x8b, 0x0b, 0x72, 0x92, 0xd0, 0x28,
	0x9f, 0xfb, 0xd1, 0x3c, 0x65, 0xd1, 0x85, 0x69, 0xa8, 0x4e, 0x97, 0x62, 0xd6, 0x4f, 0x0d, 0xda,
	0x37, 0xda, 0x55, 0xa3, 0xe8, 0x42, 0x53, 0xb6, 0xa2, 0xa4, 0x6b, 0x39, 0x70, 0xd3, 0xd7, 0x79,
	0x03, 0x2b, 0x04, 0xbd, 0x00, 0x28, 0xf7, 0x22, 0xcc, 0x78, 0xac, 0x06, 0xd2, 0x72, 0x1e, 0xdb,
	0xd7, 0xcb, 0xe3, 0xcb, 0x1a, 0x27, 0x2c, 0x1d, 0x10, 0xce, 0xa3, 0x98, 0x9c, 0x37, 0xb0, 0x51,
	0x22, 0x03, 0x1e, 0xa3, 0x27, 0x60, 0x50, 0x76, 0x41, 0x42, 0x1a, 0x65, 0xc4, 0xfc, 0x25, 0xf5,
	0x30, 0xb0, 0x2e, 0x3d, 0xc3, 0x28, 0x23, 0xa8, 0x0f, 0x4d, 0xb9, 0x59, 0xe6, 0xef, 0xbb, 0x85,
	0x52, 0xc4, 0x93, 0x36, 0x6c, 0xe5, 0x55, 0xe1, 0xa1, 0x98, 0xcf, 0x08, 0xb7, 0x06, 0xb0, 0x7d,
	0xab, 0x00, 0x74, 0x0c, 0xcd, 0x84, 0x4e, 0x59, 0xd5, 0xd1, 0xfe, 0xca, 0x4a, 0x5f, 0xd3, 0x29,
	0x93, 0x4d, 0x4a, 0xf2, 0xc9, 0x06, 0xac, 0x67, 0x3c, 0xb6, 0x3e, 0xc2, 0xce, 0x12, 0x16, 0x6a,
	0x2b, 0x54, 0x65, 0x34, 0xb0, 0x3c, 0xa2, 0xe7, 0xd0, 0x94, 0x05, 0x28, 0x39, 0xb6, 0xfe, 0x73,
	0xc9, 0x68, 0x3e, 0x23, 0x58, 0x51, 0x9f, 0x7d, 0xfb, 0x27, 0xb7, 0x04, 0xd1, 0x53, 0xd8, 0x7f,
	0x37, 0x7c, 0x33, 0xf4, 0xde, 0x0f, 0x43, 0x1f, 0x7b, 0x23, 0xef, 0x95, 0xf7, 0x36, 0x1c, 0xb8,
	0x41, 0xf0, 0xf2, 0xcc, 0x0d, 0x47, 0x1f, 0x7c, 0xb7, 0xdd, 0x40, 0x16, 0x1c, 0xf8, 0xd8, 0x3b,
	0xc3, 0x6e, 0x10, 0xac, 0xe0, 0x68, 0xe8, 0x10, 0xf6, 0x5c, 0x8c, 0x3d, 0xbc, 0x82, 0xb0, 0xe6,
	0x7c, 0xd7, 0x40, 0xf7, 0xaa, 0x2a, 0x91, 0x0b, 0x7a, 0xbd, 0x09, 0x68, 0x61, 0x96, 0xb7, 0x5e,
	0x56, 0xa7, 0xb3, 0x0c, 0x2a, 0xf5, 0xb7, 0x1a, 0x47, 0x1a, 0xf2, 0xe0, 0xfe, 0xe2, 0xfb, 0x46,
	0x0b, 0x3a, 0x2c, 0xf9, 0x1d, 0x74, 0x0e, 0x56, 0xc1, 0x75, 0x4a, 0x5f, 0x1b, 0x6f, 0xaa, 0xe1,
	0x1f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x03, 0xf2, 0x39, 0xa0, 0xa7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ObserverClient is the client API for Observer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ObserverClient interface {
	// GetFlows returning structured data, meant to eventually obsolete GetLastNFlows.
	GetFlows(ctx context.Context, in *GetFlowsRequest, opts ...grpc.CallOption) (Observer_GetFlowsClient, error)
	// ServerStatus returns some details about the running hubble server.
	ServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (*ServerStatusResponse, error)
}

type observerClient struct {
	cc *grpc.ClientConn
}

func NewObserverClient(cc *grpc.ClientConn) ObserverClient {
	return &observerClient{cc}
}

func (c *observerClient) GetFlows(ctx context.Context, in *GetFlowsRequest, opts ...grpc.CallOption) (Observer_GetFlowsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Observer_serviceDesc.Streams[0], "/observer.Observer/GetFlows", opts...)
	if err != nil {
		return nil, err
	}
	x := &observerGetFlowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Observer_GetFlowsClient interface {
	Recv() (*GetFlowsResponse, error)
	grpc.ClientStream
}

type observerGetFlowsClient struct {
	grpc.ClientStream
}

func (x *observerGetFlowsClient) Recv() (*GetFlowsResponse, error) {
	m := new(GetFlowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *observerClient) ServerStatus(ctx context.Context, in *ServerStatusRequest, opts ...grpc.CallOption) (*ServerStatusResponse, error) {
	out := new(ServerStatusResponse)
	err := c.cc.Invoke(ctx, "/observer.Observer/ServerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObserverServer is the server API for Observer service.
type ObserverServer interface {
	// GetFlows returning structured data, meant to eventually obsolete GetLastNFlows.
	GetFlows(*GetFlowsRequest, Observer_GetFlowsServer) error
	// ServerStatus returns some details about the running hubble server.
	ServerStatus(context.Context, *ServerStatusRequest) (*ServerStatusResponse, error)
}

// UnimplementedObserverServer can be embedded to have forward compatible implementations.
type UnimplementedObserverServer struct {
}

func (*UnimplementedObserverServer) GetFlows(req *GetFlowsRequest, srv Observer_GetFlowsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFlows not implemented")
}
func (*UnimplementedObserverServer) ServerStatus(ctx context.Context, req *ServerStatusRequest) (*ServerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerStatus not implemented")
}

func RegisterObserverServer(s *grpc.Server, srv ObserverServer) {
	s.RegisterService(&_Observer_serviceDesc, srv)
}

func _Observer_GetFlows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFlowsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObserverServer).GetFlows(m, &observerGetFlowsServer{stream})
}

type Observer_GetFlowsServer interface {
	Send(*GetFlowsResponse) error
	grpc.ServerStream
}

type observerGetFlowsServer struct {
	grpc.ServerStream
}

func (x *observerGetFlowsServer) Send(m *GetFlowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Observer_ServerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObserverServer).ServerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/observer.Observer/ServerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObserverServer).ServerStatus(ctx, req.(*ServerStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Observer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "observer.Observer",
	HandlerType: (*ObserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerStatus",
			Handler:    _Observer_ServerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFlows",
			Handler:       _Observer_GetFlows_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "observer/observer.proto",
}
