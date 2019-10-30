// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/agent.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NetworkProtocol int32

const (
	NetworkProtocol_TCP NetworkProtocol = 0
	NetworkProtocol_UDP NetworkProtocol = 1
)

var NetworkProtocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var NetworkProtocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x NetworkProtocol) String() string {
	return proto.EnumName(NetworkProtocol_name, int32(x))
}

func (NetworkProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{0}
}

type GameserverStatus int32

const (
	GameserverStatus_RUNNING GameserverStatus = 0
	GameserverStatus_PENDING GameserverStatus = 1
	GameserverStatus_ERROR   GameserverStatus = 2
)

var GameserverStatus_name = map[int32]string{
	0: "RUNNING",
	1: "PENDING",
	2: "ERROR",
}

var GameserverStatus_value = map[string]int32{
	"RUNNING": 0,
	"PENDING": 1,
	"ERROR":   2,
}

func (x GameserverStatus) String() string {
	return proto.EnumName(GameserverStatus_name, int32(x))
}

func (GameserverStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{1}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type AgentResources struct {
	Cpus                 int32    `protobuf:"varint,1,opt,name=cpus,proto3" json:"cpus,omitempty"`
	Memory               int32    `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentResources) Reset()         { *m = AgentResources{} }
func (m *AgentResources) String() string { return proto.CompactTextString(m) }
func (*AgentResources) ProtoMessage()    {}
func (*AgentResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{1}
}

func (m *AgentResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentResources.Unmarshal(m, b)
}
func (m *AgentResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentResources.Marshal(b, m, deterministic)
}
func (m *AgentResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentResources.Merge(m, src)
}
func (m *AgentResources) XXX_Size() int {
	return xxx_messageInfo_AgentResources.Size(m)
}
func (m *AgentResources) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentResources.DiscardUnknown(m)
}

var xxx_messageInfo_AgentResources proto.InternalMessageInfo

func (m *AgentResources) GetCpus() int32 {
	if m != nil {
		return m.Cpus
	}
	return 0
}

func (m *AgentResources) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

type AgentResourceUsage struct {
	Memory               int32    `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentResourceUsage) Reset()         { *m = AgentResourceUsage{} }
func (m *AgentResourceUsage) String() string { return proto.CompactTextString(m) }
func (*AgentResourceUsage) ProtoMessage()    {}
func (*AgentResourceUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{2}
}

func (m *AgentResourceUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentResourceUsage.Unmarshal(m, b)
}
func (m *AgentResourceUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentResourceUsage.Marshal(b, m, deterministic)
}
func (m *AgentResourceUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentResourceUsage.Merge(m, src)
}
func (m *AgentResourceUsage) XXX_Size() int {
	return xxx_messageInfo_AgentResourceUsage.Size(m)
}
func (m *AgentResourceUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentResourceUsage.DiscardUnknown(m)
}

var xxx_messageInfo_AgentResourceUsage proto.InternalMessageInfo

func (m *AgentResourceUsage) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

type AgentState struct {
	Hostname             string              `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Resources            *AgentResources     `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	ResourceUsage        *AgentResourceUsage `protobuf:"bytes,3,opt,name=resourceUsage,proto3" json:"resourceUsage,omitempty"`
	RunningGameservers   []*Gameserver       `protobuf:"bytes,4,rep,name=runningGameservers,proto3" json:"runningGameservers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AgentState) Reset()         { *m = AgentState{} }
func (m *AgentState) String() string { return proto.CompactTextString(m) }
func (*AgentState) ProtoMessage()    {}
func (*AgentState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{3}
}

func (m *AgentState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentState.Unmarshal(m, b)
}
func (m *AgentState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentState.Marshal(b, m, deterministic)
}
func (m *AgentState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentState.Merge(m, src)
}
func (m *AgentState) XXX_Size() int {
	return xxx_messageInfo_AgentState.Size(m)
}
func (m *AgentState) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentState.DiscardUnknown(m)
}

var xxx_messageInfo_AgentState proto.InternalMessageInfo

func (m *AgentState) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *AgentState) GetResources() *AgentResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *AgentState) GetResourceUsage() *AgentResourceUsage {
	if m != nil {
		return m.ResourceUsage
	}
	return nil
}

func (m *AgentState) GetRunningGameservers() []*Gameserver {
	if m != nil {
		return m.RunningGameservers
	}
	return nil
}

type AgentResourceReservation struct {
	Memory               int32    `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentResourceReservation) Reset()         { *m = AgentResourceReservation{} }
func (m *AgentResourceReservation) String() string { return proto.CompactTextString(m) }
func (*AgentResourceReservation) ProtoMessage()    {}
func (*AgentResourceReservation) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{4}
}

func (m *AgentResourceReservation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentResourceReservation.Unmarshal(m, b)
}
func (m *AgentResourceReservation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentResourceReservation.Marshal(b, m, deterministic)
}
func (m *AgentResourceReservation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentResourceReservation.Merge(m, src)
}
func (m *AgentResourceReservation) XXX_Size() int {
	return xxx_messageInfo_AgentResourceReservation.Size(m)
}
func (m *AgentResourceReservation) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentResourceReservation.DiscardUnknown(m)
}

var xxx_messageInfo_AgentResourceReservation proto.InternalMessageInfo

func (m *AgentResourceReservation) GetMemory() int32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

type ResourceRequirements struct {
	CpuReservation       int32    `protobuf:"varint,1,opt,name=cpuReservation,proto3" json:"cpuReservation,omitempty"`
	CpuLimit             int32    `protobuf:"varint,2,opt,name=cpuLimit,proto3" json:"cpuLimit,omitempty"`
	MemoryReservation    int32    `protobuf:"varint,3,opt,name=memoryReservation,proto3" json:"memoryReservation,omitempty"`
	MemoryLimit          int32    `protobuf:"varint,4,opt,name=memoryLimit,proto3" json:"memoryLimit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{5}
}

func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (m *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(m, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetCpuReservation() int32 {
	if m != nil {
		return m.CpuReservation
	}
	return 0
}

func (m *ResourceRequirements) GetCpuLimit() int32 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *ResourceRequirements) GetMemoryReservation() int32 {
	if m != nil {
		return m.MemoryReservation
	}
	return 0
}

func (m *ResourceRequirements) GetMemoryLimit() int32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

type NetworkPort struct {
	Protocol             NetworkProtocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=proto.NetworkProtocol" json:"protocol,omitempty"`
	ContainerPort        int32           `protobuf:"varint,2,opt,name=containerPort,proto3" json:"containerPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NetworkPort) Reset()         { *m = NetworkPort{} }
func (m *NetworkPort) String() string { return proto.CompactTextString(m) }
func (*NetworkPort) ProtoMessage()    {}
func (*NetworkPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{6}
}

func (m *NetworkPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPort.Unmarshal(m, b)
}
func (m *NetworkPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPort.Marshal(b, m, deterministic)
}
func (m *NetworkPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPort.Merge(m, src)
}
func (m *NetworkPort) XXX_Size() int {
	return xxx_messageInfo_NetworkPort.Size(m)
}
func (m *NetworkPort) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPort.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPort proto.InternalMessageInfo

func (m *NetworkPort) GetProtocol() NetworkProtocol {
	if m != nil {
		return m.Protocol
	}
	return NetworkProtocol_TCP
}

func (m *NetworkPort) GetContainerPort() int32 {
	if m != nil {
		return m.ContainerPort
	}
	return 0
}

type NetworkPortMapping struct {
	Protocol             NetworkProtocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=proto.NetworkProtocol" json:"protocol,omitempty"`
	ContainerPort        int32           `protobuf:"varint,2,opt,name=containerPort,proto3" json:"containerPort,omitempty"`
	HostPort             int32           `protobuf:"varint,3,opt,name=hostPort,proto3" json:"hostPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *NetworkPortMapping) Reset()         { *m = NetworkPortMapping{} }
func (m *NetworkPortMapping) String() string { return proto.CompactTextString(m) }
func (*NetworkPortMapping) ProtoMessage()    {}
func (*NetworkPortMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{7}
}

func (m *NetworkPortMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPortMapping.Unmarshal(m, b)
}
func (m *NetworkPortMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPortMapping.Marshal(b, m, deterministic)
}
func (m *NetworkPortMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPortMapping.Merge(m, src)
}
func (m *NetworkPortMapping) XXX_Size() int {
	return xxx_messageInfo_NetworkPortMapping.Size(m)
}
func (m *NetworkPortMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPortMapping.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPortMapping proto.InternalMessageInfo

func (m *NetworkPortMapping) GetProtocol() NetworkProtocol {
	if m != nil {
		return m.Protocol
	}
	return NetworkProtocol_TCP
}

func (m *NetworkPortMapping) GetContainerPort() int32 {
	if m != nil {
		return m.ContainerPort
	}
	return 0
}

func (m *NetworkPortMapping) GetHostPort() int32 {
	if m != nil {
		return m.HostPort
	}
	return 0
}

type EnvironmentVariable struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentVariable) Reset()         { *m = EnvironmentVariable{} }
func (m *EnvironmentVariable) String() string { return proto.CompactTextString(m) }
func (*EnvironmentVariable) ProtoMessage()    {}
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{8}
}

func (m *EnvironmentVariable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentVariable.Unmarshal(m, b)
}
func (m *EnvironmentVariable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentVariable.Marshal(b, m, deterministic)
}
func (m *EnvironmentVariable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentVariable.Merge(m, src)
}
func (m *EnvironmentVariable) XXX_Size() int {
	return xxx_messageInfo_EnvironmentVariable.Size(m)
}
func (m *EnvironmentVariable) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentVariable.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentVariable proto.InternalMessageInfo

func (m *EnvironmentVariable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvironmentVariable) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetGameserversRequest struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGameserversRequest) Reset()         { *m = GetGameserversRequest{} }
func (m *GetGameserversRequest) String() string { return proto.CompactTextString(m) }
func (*GetGameserversRequest) ProtoMessage()    {}
func (*GetGameserversRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{9}
}

func (m *GetGameserversRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGameserversRequest.Unmarshal(m, b)
}
func (m *GetGameserversRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGameserversRequest.Marshal(b, m, deterministic)
}
func (m *GetGameserversRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGameserversRequest.Merge(m, src)
}
func (m *GetGameserversRequest) XXX_Size() int {
	return xxx_messageInfo_GetGameserversRequest.Size(m)
}
func (m *GetGameserversRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGameserversRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGameserversRequest proto.InternalMessageInfo

func (m *GetGameserversRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type GameserverRunConfigurationList struct {
	Servers              []*GameserverRunConfiguration `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GameserverRunConfigurationList) Reset()         { *m = GameserverRunConfigurationList{} }
func (m *GameserverRunConfigurationList) String() string { return proto.CompactTextString(m) }
func (*GameserverRunConfigurationList) ProtoMessage()    {}
func (*GameserverRunConfigurationList) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{10}
}

func (m *GameserverRunConfigurationList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameserverRunConfigurationList.Unmarshal(m, b)
}
func (m *GameserverRunConfigurationList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameserverRunConfigurationList.Marshal(b, m, deterministic)
}
func (m *GameserverRunConfigurationList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameserverRunConfigurationList.Merge(m, src)
}
func (m *GameserverRunConfigurationList) XXX_Size() int {
	return xxx_messageInfo_GameserverRunConfigurationList.Size(m)
}
func (m *GameserverRunConfigurationList) XXX_DiscardUnknown() {
	xxx_messageInfo_GameserverRunConfigurationList.DiscardUnknown(m)
}

var xxx_messageInfo_GameserverRunConfigurationList proto.InternalMessageInfo

func (m *GameserverRunConfigurationList) GetServers() []*GameserverRunConfiguration {
	if m != nil {
		return m.Servers
	}
	return nil
}

type GameserverRunConfiguration struct {
	UUID                 string                 `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name                 string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Agent                string                 `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	Image                string                 `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	ResourceRequirements *ResourceRequirements  `protobuf:"bytes,5,opt,name=resourceRequirements,proto3" json:"resourceRequirements,omitempty"`
	Ports                []*NetworkPort         `protobuf:"bytes,6,rep,name=ports,proto3" json:"ports,omitempty"`
	Environment          []*EnvironmentVariable `protobuf:"bytes,7,rep,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *GameserverRunConfiguration) Reset()         { *m = GameserverRunConfiguration{} }
func (m *GameserverRunConfiguration) String() string { return proto.CompactTextString(m) }
func (*GameserverRunConfiguration) ProtoMessage()    {}
func (*GameserverRunConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{11}
}

func (m *GameserverRunConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameserverRunConfiguration.Unmarshal(m, b)
}
func (m *GameserverRunConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameserverRunConfiguration.Marshal(b, m, deterministic)
}
func (m *GameserverRunConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameserverRunConfiguration.Merge(m, src)
}
func (m *GameserverRunConfiguration) XXX_Size() int {
	return xxx_messageInfo_GameserverRunConfiguration.Size(m)
}
func (m *GameserverRunConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_GameserverRunConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_GameserverRunConfiguration proto.InternalMessageInfo

func (m *GameserverRunConfiguration) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *GameserverRunConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameserverRunConfiguration) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *GameserverRunConfiguration) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *GameserverRunConfiguration) GetResourceRequirements() *ResourceRequirements {
	if m != nil {
		return m.ResourceRequirements
	}
	return nil
}

func (m *GameserverRunConfiguration) GetPorts() []*NetworkPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *GameserverRunConfiguration) GetEnvironment() []*EnvironmentVariable {
	if m != nil {
		return m.Environment
	}
	return nil
}

type Gameserver struct {
	UUID                 string                `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Status               GameserverStatus      `protobuf:"varint,2,opt,name=status,proto3,enum=proto.GameserverStatus" json:"status,omitempty"`
	Info                 string                `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	PortMappings         []*NetworkPortMapping `protobuf:"bytes,4,rep,name=portMappings,proto3" json:"portMappings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Gameserver) Reset()         { *m = Gameserver{} }
func (m *Gameserver) String() string { return proto.CompactTextString(m) }
func (*Gameserver) ProtoMessage()    {}
func (*Gameserver) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd830a99d5efef4e, []int{12}
}

func (m *Gameserver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gameserver.Unmarshal(m, b)
}
func (m *Gameserver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gameserver.Marshal(b, m, deterministic)
}
func (m *Gameserver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gameserver.Merge(m, src)
}
func (m *Gameserver) XXX_Size() int {
	return xxx_messageInfo_Gameserver.Size(m)
}
func (m *Gameserver) XXX_DiscardUnknown() {
	xxx_messageInfo_Gameserver.DiscardUnknown(m)
}

var xxx_messageInfo_Gameserver proto.InternalMessageInfo

func (m *Gameserver) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Gameserver) GetStatus() GameserverStatus {
	if m != nil {
		return m.Status
	}
	return GameserverStatus_RUNNING
}

func (m *Gameserver) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *Gameserver) GetPortMappings() []*NetworkPortMapping {
	if m != nil {
		return m.PortMappings
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.NetworkProtocol", NetworkProtocol_name, NetworkProtocol_value)
	proto.RegisterEnum("proto.GameserverStatus", GameserverStatus_name, GameserverStatus_value)
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*AgentResources)(nil), "proto.AgentResources")
	proto.RegisterType((*AgentResourceUsage)(nil), "proto.AgentResourceUsage")
	proto.RegisterType((*AgentState)(nil), "proto.AgentState")
	proto.RegisterType((*AgentResourceReservation)(nil), "proto.AgentResourceReservation")
	proto.RegisterType((*ResourceRequirements)(nil), "proto.ResourceRequirements")
	proto.RegisterType((*NetworkPort)(nil), "proto.NetworkPort")
	proto.RegisterType((*NetworkPortMapping)(nil), "proto.NetworkPortMapping")
	proto.RegisterType((*EnvironmentVariable)(nil), "proto.EnvironmentVariable")
	proto.RegisterType((*GetGameserversRequest)(nil), "proto.GetGameserversRequest")
	proto.RegisterType((*GameserverRunConfigurationList)(nil), "proto.GameserverRunConfigurationList")
	proto.RegisterType((*GameserverRunConfiguration)(nil), "proto.GameserverRunConfiguration")
	proto.RegisterType((*Gameserver)(nil), "proto.Gameserver")
}

func init() { proto.RegisterFile("proto/agent.proto", fileDescriptor_dd830a99d5efef4e) }

var fileDescriptor_dd830a99d5efef4e = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x4e, 0xdb, 0x4e,
	0x10, 0xc7, 0xf9, 0x24, 0x13, 0xc8, 0x3f, 0x99, 0x3f, 0x50, 0x93, 0x56, 0x15, 0x75, 0x3f, 0x84,
	0x28, 0x02, 0x29, 0x1c, 0x7a, 0x28, 0x15, 0x42, 0x10, 0x21, 0x24, 0x1a, 0xa2, 0xa5, 0xe9, 0xad,
	0x07, 0x13, 0x2d, 0xa9, 0x55, 0xbc, 0xeb, 0xae, 0xd7, 0xa9, 0x78, 0x85, 0x1e, 0xfb, 0x0c, 0xbd,
	0xb4, 0x6f, 0xd5, 0x37, 0xa9, 0x76, 0x6c, 0x27, 0x4e, 0x62, 0xda, 0x53, 0x4f, 0xde, 0x99, 0xfd,
	0xed, 0x7c, 0xfe, 0x66, 0x0c, 0xad, 0x40, 0x49, 0x2d, 0xf7, 0xdd, 0x11, 0x17, 0x7a, 0x8f, 0xce,
	0x58, 0xa6, 0x8f, 0x53, 0x85, 0x72, 0xd7, 0x0f, 0xf4, 0x9d, 0x73, 0x08, 0x8d, 0x63, 0x73, 0xcd,
	0x78, 0x28, 0x23, 0x35, 0xe4, 0x21, 0x22, 0x94, 0x86, 0x41, 0x14, 0xda, 0xd6, 0x96, 0xb5, 0x5d,
	0x66, 0x74, 0xc6, 0x0d, 0xa8, 0xf8, 0xdc, 0x97, 0xea, 0xce, 0x2e, 0x90, 0x36, 0x91, 0x9c, 0x5d,
	0xc0, 0x99, 0xd7, 0x83, 0xd0, 0x1d, 0xf1, 0x7b, 0xd1, 0xbf, 0x2c, 0x00, 0x82, 0x5f, 0x69, 0x57,
	0x73, 0x6c, 0xc3, 0xf2, 0x47, 0x19, 0x6a, 0xe1, 0xfa, 0x9c, 0x9c, 0xd5, 0xd8, 0x44, 0xc6, 0x03,
	0xa8, 0xa9, 0x34, 0x22, 0xb2, 0x52, 0xef, 0xac, 0xc7, 0x19, 0xec, 0xcd, 0x86, 0xcb, 0xa6, 0x38,
	0x3c, 0x82, 0x55, 0x95, 0x0d, 0xc4, 0x2e, 0xd2, 0xc3, 0xcd, 0xbc, 0x87, 0x04, 0x60, 0xb3, 0x78,
	0x3c, 0x06, 0x54, 0x91, 0x10, 0x9e, 0x18, 0x9d, 0xb9, 0x3e, 0x0f, 0xb9, 0x1a, 0x73, 0x15, 0xda,
	0xa5, 0xad, 0xe2, 0x76, 0xbd, 0xd3, 0x4a, 0xac, 0x4c, 0x6f, 0x58, 0x0e, 0xd8, 0xe9, 0x80, 0x3d,
	0xe3, 0x87, 0xd1, 0x8d, 0xab, 0x3d, 0x29, 0xee, 0xad, 0xcb, 0x4f, 0x0b, 0xd6, 0xa6, 0xf8, 0xcf,
	0x91, 0xa7, 0xb8, 0xcf, 0x85, 0x0e, 0xf1, 0x05, 0x34, 0x86, 0x41, 0x94, 0x31, 0x91, 0x34, 0x65,
	0x4e, 0x6b, 0x2a, 0x39, 0x0c, 0xa2, 0x0b, 0xcf, 0xf7, 0x74, 0x62, 0x7a, 0x22, 0xe3, 0x2e, 0xb4,
	0x62, 0x37, 0x59, 0x33, 0x45, 0x02, 0x2d, 0x5e, 0xe0, 0x16, 0xd4, 0x63, 0x65, 0x6c, 0xac, 0x44,
	0xb8, 0xac, 0xca, 0x19, 0x41, 0xbd, 0xc7, 0xf5, 0x17, 0xa9, 0x3e, 0xf5, 0xa5, 0xd2, 0xd8, 0x81,
	0x65, 0xaa, 0xcb, 0x50, 0xde, 0x52, 0x70, 0x8d, 0xce, 0x46, 0x52, 0xa8, 0x14, 0x95, 0xdc, 0xb2,
	0x09, 0x0e, 0x9f, 0xc1, 0xea, 0x50, 0x0a, 0xed, 0x7a, 0x82, 0x2b, 0x63, 0x24, 0x89, 0x79, 0x56,
	0xe9, 0x7c, 0xb5, 0x00, 0x33, 0x9e, 0xde, 0xba, 0x41, 0xe0, 0x89, 0xd1, 0xbf, 0x73, 0x98, 0xf2,
	0x91, 0x00, 0x71, 0x81, 0x26, 0xb2, 0x73, 0x04, 0xff, 0x77, 0xc5, 0xd8, 0x53, 0x52, 0x98, 0xce,
	0xbc, 0x77, 0x95, 0xe7, 0x5e, 0xdf, 0x72, 0x33, 0x2b, 0x19, 0xfa, 0xd2, 0x19, 0xd7, 0xa0, 0x3c,
	0x76, 0x6f, 0x23, 0x4e, 0x4e, 0x6a, 0x2c, 0x16, 0x9c, 0x03, 0x58, 0x3f, 0xe3, 0x3a, 0xc3, 0x14,
	0xd3, 0x68, 0x1e, 0xea, 0x3f, 0x4d, 0x81, 0xf3, 0x01, 0x1e, 0x67, 0xe8, 0x16, 0x89, 0x13, 0x29,
	0x6e, 0xbc, 0x51, 0xa4, 0xa8, 0x57, 0x17, 0x5e, 0xa8, 0xf1, 0x35, 0x54, 0x53, 0x9a, 0x5a, 0x44,
	0xd3, 0x27, 0x8b, 0x34, 0x9d, 0x7b, 0xc7, 0xd2, 0x17, 0xce, 0x8f, 0x02, 0xb4, 0xef, 0xc7, 0x99,
	0xe4, 0x06, 0x83, 0xf3, 0xd3, 0x34, 0x39, 0x73, 0x9e, 0x24, 0x5c, 0x98, 0x4d, 0x98, 0x36, 0x0c,
	0x15, 0xad, 0xc6, 0x62, 0xc1, 0x68, 0x3d, 0xdf, 0x0c, 0x61, 0x29, 0xd6, 0x92, 0x80, 0x97, 0xb0,
	0xa6, 0x72, 0x98, 0x6e, 0x97, 0x69, 0x52, 0x1f, 0x26, 0xc1, 0xe7, 0x0d, 0x03, 0xcb, 0x7d, 0x88,
	0xdb, 0x50, 0x0e, 0xa4, 0xd2, 0xa1, 0x5d, 0xa1, 0xf4, 0x71, 0x8e, 0x0b, 0x52, 0x69, 0x16, 0x03,
	0xf0, 0x10, 0xea, 0x7c, 0xda, 0x42, 0xbb, 0x4a, 0xf8, 0x76, 0x82, 0xcf, 0x69, 0x2e, 0xcb, 0xc2,
	0x9d, 0xef, 0x16, 0xc0, 0xb4, 0x56, 0xb9, 0xb5, 0xd9, 0x87, 0x4a, 0xa8, 0x5d, 0x1d, 0xc5, 0x0b,
	0xab, 0xd1, 0x79, 0xb0, 0xd0, 0x8a, 0x2b, 0xba, 0x66, 0x09, 0xcc, 0x18, 0xf1, 0xc4, 0x8d, 0x4c,
	0xea, 0x46, 0x67, 0x7c, 0x03, 0x2b, 0xc1, 0x94, 0xed, 0xe9, 0xf2, 0xd9, 0x5c, 0x4c, 0x2b, 0x41,
	0xb0, 0x19, 0xf8, 0xce, 0x53, 0xf8, 0x6f, 0x6e, 0x0c, 0xb0, 0x0a, 0xc5, 0x77, 0x27, 0xfd, 0xe6,
	0x92, 0x39, 0x0c, 0x4e, 0xfb, 0x4d, 0x6b, 0xe7, 0x15, 0x34, 0xe7, 0x63, 0xc2, 0x3a, 0x54, 0xd9,
	0xa0, 0xd7, 0x3b, 0xef, 0x9d, 0x35, 0x97, 0x8c, 0xd0, 0xef, 0xf6, 0x4e, 0x8d, 0x60, 0x61, 0x0d,
	0xca, 0x5d, 0xc6, 0x2e, 0x59, 0xb3, 0xd0, 0xf9, 0x66, 0xc1, 0x4a, 0xbc, 0xc0, 0xb9, 0x1a, 0x7b,
	0x43, 0x8e, 0x2f, 0x61, 0x99, 0xf1, 0x91, 0x17, 0x6a, 0xae, 0xb0, 0x95, 0x5d, 0xb3, 0xb4, 0xe1,
	0xdb, 0x2b, 0x69, 0x75, 0xcd, 0xaf, 0x06, 0x07, 0xd0, 0x48, 0x46, 0xe0, 0x2a, 0x26, 0x20, 0x3e,
	0x4a, 0x2b, 0x94, 0x37, 0x19, 0xed, 0xe7, 0x7f, 0xa5, 0xb2, 0x19, 0x81, 0xeb, 0x0a, 0xa1, 0x0e,
	0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x61, 0x7f, 0x72, 0xed, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	Register(ctx context.Context, in *AgentState, opts ...grpc.CallOption) (*Empty, error)
	GetGameServers(ctx context.Context, in *GetGameserversRequest, opts ...grpc.CallOption) (*GameserverRunConfigurationList, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Register(ctx context.Context, in *AgentState, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.AgentService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) GetGameServers(ctx context.Context, in *GetGameserversRequest, opts ...grpc.CallOption) (*GameserverRunConfigurationList, error) {
	out := new(GameserverRunConfigurationList)
	err := c.cc.Invoke(ctx, "/proto.AgentService/GetGameServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	Register(context.Context, *AgentState) (*Empty, error)
	GetGameServers(context.Context, *GetGameserversRequest) (*GameserverRunConfigurationList, error)
}

// UnimplementedAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) Register(ctx context.Context, req *AgentState) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedAgentServiceServer) GetGameServers(ctx context.Context, req *GetGameserversRequest) (*GameserverRunConfigurationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameServers not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Register(ctx, req.(*AgentState))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_GetGameServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameserversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).GetGameServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AgentService/GetGameServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).GetGameServers(ctx, req.(*GetGameserversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AgentService_Register_Handler,
		},
		{
			MethodName: "GetGameServers",
			Handler:    _AgentService_GetGameServers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/agent.proto",
}