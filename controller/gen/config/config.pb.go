// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/config.proto

package config // import "github.com/linkerd/linkerd2/controller/gen/config"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TLS int32

const (
	TLS_NONE     TLS = 0
	TLS_OPTIONAL TLS = 1
)

var TLS_name = map[int32]string{
	0: "NONE",
	1: "OPTIONAL",
}
var TLS_value = map[string]int32{
	"NONE":     0,
	"OPTIONAL": 1,
}

func (x TLS) String() string {
	return proto.EnumName(TLS_name, int32(x))
}
func (TLS) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{0}
}

type GlobalConfig struct {
	LinkerdNamespace     string   `protobuf:"bytes,1,opt,name=linkerd_namespace,json=linkerdNamespace,proto3" json:"linkerd_namespace,omitempty"`
	CniEnabled           bool     `protobuf:"varint,2,opt,name=cni_enabled,json=cniEnabled,proto3" json:"cni_enabled,omitempty"`
	TlsEnabled           bool     `protobuf:"varint,3,opt,name=tls_enabled,json=tlsEnabled,proto3" json:"tls_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlobalConfig) Reset()         { *m = GlobalConfig{} }
func (m *GlobalConfig) String() string { return proto.CompactTextString(m) }
func (*GlobalConfig) ProtoMessage()    {}
func (*GlobalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{0}
}
func (m *GlobalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlobalConfig.Unmarshal(m, b)
}
func (m *GlobalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlobalConfig.Marshal(b, m, deterministic)
}
func (dst *GlobalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlobalConfig.Merge(dst, src)
}
func (m *GlobalConfig) XXX_Size() int {
	return xxx_messageInfo_GlobalConfig.Size(m)
}
func (m *GlobalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GlobalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GlobalConfig proto.InternalMessageInfo

func (m *GlobalConfig) GetLinkerdNamespace() string {
	if m != nil {
		return m.LinkerdNamespace
	}
	return ""
}

func (m *GlobalConfig) GetCniEnabled() bool {
	if m != nil {
		return m.CniEnabled
	}
	return false
}

func (m *GlobalConfig) GetTlsEnabled() bool {
	if m != nil {
		return m.TlsEnabled
	}
	return false
}

type ProxyConfig struct {
	ProxyImage              *Image                `protobuf:"bytes,1,opt,name=proxy_image,json=proxyImage,proto3" json:"proxy_image,omitempty"`
	ProxyInitImage          *Image                `protobuf:"bytes,2,opt,name=proxy_init_image,json=proxyInitImage,proto3" json:"proxy_init_image,omitempty"`
	ApiPort                 *Port                 `protobuf:"bytes,3,opt,name=api_port,json=apiPort,proto3" json:"api_port,omitempty"`
	ControlPort             *Port                 `protobuf:"bytes,4,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	IgnoreInboundPorts      []*Port               `protobuf:"bytes,5,rep,name=ignore_inbound_ports,json=ignoreInboundPorts,proto3" json:"ignore_inbound_ports,omitempty"`
	IgnoreOutboundPorts     []*Port               `protobuf:"bytes,6,rep,name=ignore_outbound_ports,json=ignoreOutboundPorts,proto3" json:"ignore_outbound_ports,omitempty"`
	InboundPort             *Port                 `protobuf:"bytes,7,opt,name=inbound_port,json=inboundPort,proto3" json:"inbound_port,omitempty"`
	MetricsPort             *Port                 `protobuf:"bytes,8,opt,name=metrics_port,json=metricsPort,proto3" json:"metrics_port,omitempty"`
	OutboundPort            *Port                 `protobuf:"bytes,9,opt,name=outbound_port,json=outboundPort,proto3" json:"outbound_port,omitempty"`
	Resource                *ResourceRequirements `protobuf:"bytes,10,opt,name=resource,proto3" json:"resource,omitempty"`
	TlsEnabled              TLS                   `protobuf:"varint,11,opt,name=tls_enabled,json=tlsEnabled,proto3,enum=linkerd2.config.TLS" json:"tls_enabled,omitempty"`
	ProxyUid                int64                 `protobuf:"varint,12,opt,name=proxy_uid,json=proxyUid,proto3" json:"proxy_uid,omitempty"`
	LogLevel                *LogLevel             `protobuf:"bytes,13,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	DisableExternalProfiles bool                  `protobuf:"varint,14,opt,name=disable_external_profiles,json=disableExternalProfiles,proto3" json:"disable_external_profiles,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}              `json:"-"`
	XXX_unrecognized        []byte                `json:"-"`
	XXX_sizecache           int32                 `json:"-"`
}

func (m *ProxyConfig) Reset()         { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()    {}
func (*ProxyConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{1}
}
func (m *ProxyConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyConfig.Unmarshal(m, b)
}
func (m *ProxyConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyConfig.Marshal(b, m, deterministic)
}
func (dst *ProxyConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyConfig.Merge(dst, src)
}
func (m *ProxyConfig) XXX_Size() int {
	return xxx_messageInfo_ProxyConfig.Size(m)
}
func (m *ProxyConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyConfig proto.InternalMessageInfo

func (m *ProxyConfig) GetProxyImage() *Image {
	if m != nil {
		return m.ProxyImage
	}
	return nil
}

func (m *ProxyConfig) GetProxyInitImage() *Image {
	if m != nil {
		return m.ProxyInitImage
	}
	return nil
}

func (m *ProxyConfig) GetApiPort() *Port {
	if m != nil {
		return m.ApiPort
	}
	return nil
}

func (m *ProxyConfig) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *ProxyConfig) GetIgnoreInboundPorts() []*Port {
	if m != nil {
		return m.IgnoreInboundPorts
	}
	return nil
}

func (m *ProxyConfig) GetIgnoreOutboundPorts() []*Port {
	if m != nil {
		return m.IgnoreOutboundPorts
	}
	return nil
}

func (m *ProxyConfig) GetInboundPort() *Port {
	if m != nil {
		return m.InboundPort
	}
	return nil
}

func (m *ProxyConfig) GetMetricsPort() *Port {
	if m != nil {
		return m.MetricsPort
	}
	return nil
}

func (m *ProxyConfig) GetOutboundPort() *Port {
	if m != nil {
		return m.OutboundPort
	}
	return nil
}

func (m *ProxyConfig) GetResource() *ResourceRequirements {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ProxyConfig) GetTlsEnabled() TLS {
	if m != nil {
		return m.TlsEnabled
	}
	return TLS_NONE
}

func (m *ProxyConfig) GetProxyUid() int64 {
	if m != nil {
		return m.ProxyUid
	}
	return 0
}

func (m *ProxyConfig) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *ProxyConfig) GetDisableExternalProfiles() bool {
	if m != nil {
		return m.DisableExternalProfiles
	}
	return false
}

type Image struct {
	ImageName            string   `protobuf:"bytes,1,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	PullPolicy           string   `protobuf:"bytes,2,opt,name=pull_policy,json=pullPolicy,proto3" json:"pull_policy,omitempty"`
	Registry             string   `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{2}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *Image) GetPullPolicy() string {
	if m != nil {
		return m.PullPolicy
	}
	return ""
}

func (m *Image) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

type Port struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{3}
}
func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (dst *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(dst, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ResourceRequirements struct {
	RequestCpu           float32  `protobuf:"fixed32,1,opt,name=request_cpu,json=requestCpu,proto3" json:"request_cpu,omitempty"`
	RequestMemory        float32  `protobuf:"fixed32,2,opt,name=request_memory,json=requestMemory,proto3" json:"request_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceRequirements) Reset()         { *m = ResourceRequirements{} }
func (m *ResourceRequirements) String() string { return proto.CompactTextString(m) }
func (*ResourceRequirements) ProtoMessage()    {}
func (*ResourceRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{4}
}
func (m *ResourceRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceRequirements.Unmarshal(m, b)
}
func (m *ResourceRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceRequirements.Marshal(b, m, deterministic)
}
func (dst *ResourceRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceRequirements.Merge(dst, src)
}
func (m *ResourceRequirements) XXX_Size() int {
	return xxx_messageInfo_ResourceRequirements.Size(m)
}
func (m *ResourceRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceRequirements proto.InternalMessageInfo

func (m *ResourceRequirements) GetRequestCpu() float32 {
	if m != nil {
		return m.RequestCpu
	}
	return 0
}

func (m *ResourceRequirements) GetRequestMemory() float32 {
	if m != nil {
		return m.RequestMemory
	}
	return 0
}

type LogLevel struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{5}
}
func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (dst *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(dst, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

func (m *LogLevel) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type CNI struct {
	Image                *Image       `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	LogLevel             *LogLevel    `protobuf:"bytes,2,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	ControlPort          *Port        `protobuf:"bytes,3,opt,name=control_port,json=controlPort,proto3" json:"control_port,omitempty"`
	DestCniBinDir        string       `protobuf:"bytes,4,opt,name=dest_cni_bin_dir,json=destCniBinDir,proto3" json:"dest_cni_bin_dir,omitempty"`
	DestCniNetDir        string       `protobuf:"bytes,5,opt,name=dest_cni_net_dir,json=destCniNetDir,proto3" json:"dest_cni_net_dir,omitempty"`
	ProxyConfig          *ProxyConfig `protobuf:"bytes,6,opt,name=proxy_config,json=proxyConfig,proto3" json:"proxy_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CNI) Reset()         { *m = CNI{} }
func (m *CNI) String() string { return proto.CompactTextString(m) }
func (*CNI) ProtoMessage()    {}
func (*CNI) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{6}
}
func (m *CNI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CNI.Unmarshal(m, b)
}
func (m *CNI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CNI.Marshal(b, m, deterministic)
}
func (dst *CNI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CNI.Merge(dst, src)
}
func (m *CNI) XXX_Size() int {
	return xxx_messageInfo_CNI.Size(m)
}
func (m *CNI) XXX_DiscardUnknown() {
	xxx_messageInfo_CNI.DiscardUnknown(m)
}

var xxx_messageInfo_CNI proto.InternalMessageInfo

func (m *CNI) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *CNI) GetLogLevel() *LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return nil
}

func (m *CNI) GetControlPort() *Port {
	if m != nil {
		return m.ControlPort
	}
	return nil
}

func (m *CNI) GetDestCniBinDir() string {
	if m != nil {
		return m.DestCniBinDir
	}
	return ""
}

func (m *CNI) GetDestCniNetDir() string {
	if m != nil {
		return m.DestCniNetDir
	}
	return ""
}

func (m *CNI) GetProxyConfig() *ProxyConfig {
	if m != nil {
		return m.ProxyConfig
	}
	return nil
}

type GetParam struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParam) Reset()         { *m = GetParam{} }
func (m *GetParam) String() string { return proto.CompactTextString(m) }
func (*GetParam) ProtoMessage()    {}
func (*GetParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{7}
}
func (m *GetParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParam.Unmarshal(m, b)
}
func (m *GetParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParam.Marshal(b, m, deterministic)
}
func (dst *GetParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParam.Merge(dst, src)
}
func (m *GetParam) XXX_Size() int {
	return xxx_messageInfo_GetParam.Size(m)
}
func (m *GetParam) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParam.DiscardUnknown(m)
}

var xxx_messageInfo_GetParam proto.InternalMessageInfo

type Result struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	IsError              bool     `protobuf:"varint,2,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_ba3c98eae87a8adc, []int{8}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Result) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func init() {
	proto.RegisterType((*GlobalConfig)(nil), "linkerd2.config.GlobalConfig")
	proto.RegisterType((*ProxyConfig)(nil), "linkerd2.config.ProxyConfig")
	proto.RegisterType((*Image)(nil), "linkerd2.config.Image")
	proto.RegisterType((*Port)(nil), "linkerd2.config.Port")
	proto.RegisterType((*ResourceRequirements)(nil), "linkerd2.config.ResourceRequirements")
	proto.RegisterType((*LogLevel)(nil), "linkerd2.config.LogLevel")
	proto.RegisterType((*CNI)(nil), "linkerd2.config.CNI")
	proto.RegisterType((*GetParam)(nil), "linkerd2.config.GetParam")
	proto.RegisterType((*Result)(nil), "linkerd2.config.Result")
	proto.RegisterEnum("linkerd2.config.TLS", TLS_name, TLS_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigStoreClient is the client API for ConfigStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigStoreClient interface {
	// Return the JSON-serialized global config.
	GetGlobalConfig(ctx context.Context, in *GetParam, opts ...grpc.CallOption) (*GlobalConfig, error)
	// Return the JSON-serialized proxy config.
	GetProxyConfig(ctx context.Context, in *GetParam, opts ...grpc.CallOption) (*ProxyConfig, error)
	// Set the global config. It overwrites the existing value.
	SetGlobalConfig(ctx context.Context, in *GlobalConfig, opts ...grpc.CallOption) (*Result, error)
	// Set the proxy config. It overwrites the existing value.
	SetProxyConfig(ctx context.Context, in *ProxyConfig, opts ...grpc.CallOption) (*Result, error)
}

type configStoreClient struct {
	cc *grpc.ClientConn
}

func NewConfigStoreClient(cc *grpc.ClientConn) ConfigStoreClient {
	return &configStoreClient{cc}
}

func (c *configStoreClient) GetGlobalConfig(ctx context.Context, in *GetParam, opts ...grpc.CallOption) (*GlobalConfig, error) {
	out := new(GlobalConfig)
	err := c.cc.Invoke(ctx, "/linkerd2.config.ConfigStore/GetGlobalConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) GetProxyConfig(ctx context.Context, in *GetParam, opts ...grpc.CallOption) (*ProxyConfig, error) {
	out := new(ProxyConfig)
	err := c.cc.Invoke(ctx, "/linkerd2.config.ConfigStore/GetProxyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) SetGlobalConfig(ctx context.Context, in *GlobalConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/linkerd2.config.ConfigStore/SetGlobalConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configStoreClient) SetProxyConfig(ctx context.Context, in *ProxyConfig, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/linkerd2.config.ConfigStore/SetProxyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigStoreServer is the server API for ConfigStore service.
type ConfigStoreServer interface {
	// Return the JSON-serialized global config.
	GetGlobalConfig(context.Context, *GetParam) (*GlobalConfig, error)
	// Return the JSON-serialized proxy config.
	GetProxyConfig(context.Context, *GetParam) (*ProxyConfig, error)
	// Set the global config. It overwrites the existing value.
	SetGlobalConfig(context.Context, *GlobalConfig) (*Result, error)
	// Set the proxy config. It overwrites the existing value.
	SetProxyConfig(context.Context, *ProxyConfig) (*Result, error)
}

func RegisterConfigStoreServer(s *grpc.Server, srv ConfigStoreServer) {
	s.RegisterService(&_ConfigStore_serviceDesc, srv)
}

func _ConfigStore_GetGlobalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).GetGlobalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.config.ConfigStore/GetGlobalConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).GetGlobalConfig(ctx, req.(*GetParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_GetProxyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).GetProxyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.config.ConfigStore/GetProxyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).GetProxyConfig(ctx, req.(*GetParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_SetGlobalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlobalConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).SetGlobalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.config.ConfigStore/SetGlobalConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).SetGlobalConfig(ctx, req.(*GlobalConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigStore_SetProxyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigStoreServer).SetProxyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/linkerd2.config.ConfigStore/SetProxyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigStoreServer).SetProxyConfig(ctx, req.(*ProxyConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "linkerd2.config.ConfigStore",
	HandlerType: (*ConfigStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGlobalConfig",
			Handler:    _ConfigStore_GetGlobalConfig_Handler,
		},
		{
			MethodName: "GetProxyConfig",
			Handler:    _ConfigStore_GetProxyConfig_Handler,
		},
		{
			MethodName: "SetGlobalConfig",
			Handler:    _ConfigStore_SetGlobalConfig_Handler,
		},
		{
			MethodName: "SetProxyConfig",
			Handler:    _ConfigStore_SetProxyConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/config.proto",
}

func init() { proto.RegisterFile("config/config.proto", fileDescriptor_config_ba3c98eae87a8adc) }

var fileDescriptor_config_ba3c98eae87a8adc = []byte{
	// 850 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0x8d, 0xed, 0xd8, 0x91, 0xaf, 0x3f, 0xe2, 0xb1, 0xe9, 0xaa, 0x64, 0x0b, 0x66, 0x08, 0x28,
	0x16, 0x6c, 0x83, 0xb3, 0xb9, 0xe8, 0x56, 0xf4, 0x65, 0x6b, 0xb3, 0x20, 0x30, 0xe6, 0xda, 0x06,
	0x9d, 0xbd, 0xec, 0x61, 0x82, 0x2c, 0xb1, 0x1a, 0x31, 0x8a, 0x54, 0x29, 0x6a, 0x68, 0x1e, 0xf6,
	0xbf, 0xf6, 0x3a, 0xec, 0x8f, 0x0d, 0xfc, 0x70, 0x6a, 0xc7, 0xa9, 0x90, 0x27, 0x93, 0x97, 0xe7,
	0x9c, 0x7b, 0x7c, 0x75, 0x2f, 0x09, 0x8f, 0x62, 0xc1, 0xdf, 0xd2, 0xf4, 0xdc, 0xfe, 0x8c, 0x72,
	0x29, 0x94, 0x40, 0x87, 0x8c, 0xf2, 0x3f, 0x89, 0x4c, 0xc6, 0x23, 0x1b, 0x0e, 0xfe, 0x86, 0xee,
	0x15, 0x13, 0xab, 0x88, 0x5d, 0x98, 0x3d, 0xfa, 0x1a, 0x3e, 0x71, 0x90, 0x90, 0x47, 0x19, 0x29,
	0xf2, 0x28, 0x26, 0x7e, 0x6d, 0x58, 0x3b, 0x6b, 0xe3, 0x81, 0x3b, 0x98, 0xad, 0xe3, 0xe8, 0x0b,
	0xe8, 0xc4, 0x9c, 0x86, 0x84, 0x47, 0x2b, 0x46, 0x12, 0xbf, 0x3e, 0xac, 0x9d, 0x79, 0x18, 0x62,
	0x4e, 0x2f, 0x6d, 0x44, 0x03, 0x14, 0x2b, 0x6e, 0x01, 0x0d, 0x0b, 0x50, 0xac, 0x70, 0x80, 0xe0,
	0xbf, 0x16, 0x74, 0x16, 0x52, 0xbc, 0xbf, 0x71, 0xe9, 0x7f, 0x80, 0x4e, 0xae, 0xb7, 0x21, 0xcd,
	0xa2, 0xd4, 0x26, 0xee, 0x8c, 0x3f, 0x1d, 0xdd, 0x71, 0x3d, 0x9a, 0xe8, 0x53, 0x0c, 0x06, 0x6a,
	0xd6, 0xe8, 0x27, 0x18, 0x38, 0x22, 0xa7, 0xca, 0xb1, 0xeb, 0x95, 0xec, 0xbe, 0x65, 0x73, 0xaa,
	0xac, 0xc2, 0xb7, 0xe0, 0x45, 0x39, 0x0d, 0x73, 0x21, 0x95, 0x31, 0xda, 0x19, 0x3f, 0xde, 0x61,
	0x2e, 0x84, 0x54, 0xf8, 0x20, 0xca, 0xa9, 0x5e, 0xa0, 0x17, 0xd0, 0x8d, 0x05, 0x57, 0x52, 0x30,
	0xcb, 0xda, 0xaf, 0x62, 0x75, 0x1c, 0xd4, 0x30, 0xaf, 0xe0, 0x88, 0xa6, 0x5c, 0x48, 0x12, 0x52,
	0xbe, 0x12, 0x25, 0x4f, 0x8c, 0x40, 0xe1, 0x37, 0x87, 0x8d, 0x8f, 0x2b, 0x20, 0x4b, 0x99, 0x58,
	0x86, 0x0e, 0x15, 0x68, 0x02, 0x8f, 0x9d, 0x90, 0x28, 0xd5, 0xa6, 0x52, 0xab, 0x4a, 0xe9, 0x91,
	0xe5, 0xcc, 0x1d, 0xc5, 0x4a, 0xbd, 0x80, 0xee, 0xa6, 0x19, 0xff, 0xa0, 0xf2, 0xdf, 0xd0, 0x0f,
	0x2e, 0x34, 0x33, 0x23, 0x4a, 0xd2, 0xb8, 0xb0, 0x4c, 0xaf, 0x92, 0xe9, 0xa0, 0x86, 0xf9, 0x12,
	0x7a, 0x5b, 0xbe, 0xfd, 0x76, 0x15, 0xb5, 0x2b, 0x36, 0x0c, 0xa3, 0x57, 0xe0, 0x49, 0x52, 0x88,
	0x52, 0xc6, 0xc4, 0x07, 0x43, 0x7b, 0xba, 0x43, 0xc3, 0x0e, 0x80, 0xc9, 0xbb, 0x92, 0x4a, 0x92,
	0x11, 0xae, 0x0a, 0x7c, 0x4b, 0x43, 0xcf, 0xb7, 0xdb, 0xb3, 0x33, 0xac, 0x9d, 0xf5, 0xc7, 0x47,
	0x3b, 0x2a, 0xd7, 0xd3, 0xe5, 0x66, 0xd3, 0xa2, 0xcf, 0xa0, 0x6d, 0x7b, 0xad, 0xa4, 0x89, 0xdf,
	0x1d, 0xd6, 0xce, 0x1a, 0xd8, 0x33, 0x81, 0x5f, 0x69, 0x82, 0xbe, 0x87, 0x36, 0x13, 0x69, 0xc8,
	0xc8, 0x5f, 0x84, 0xf9, 0x3d, 0xe3, 0xeb, 0x78, 0x47, 0x71, 0x2a, 0xd2, 0xa9, 0x06, 0x60, 0x8f,
	0xb9, 0x15, 0x7a, 0x09, 0xc7, 0x09, 0x2d, 0x74, 0x82, 0x90, 0xbc, 0x57, 0x44, 0xf2, 0x88, 0x85,
	0xb9, 0x14, 0x6f, 0x29, 0x23, 0x85, 0xdf, 0x37, 0x83, 0xf3, 0xc4, 0x01, 0x2e, 0xdd, 0xf9, 0xc2,
	0x1d, 0x07, 0x31, 0x34, 0x6d, 0x0f, 0x9f, 0x02, 0x98, 0xd6, 0x37, 0xb3, 0xeb, 0xc6, 0xb6, 0x6d,
	0x22, 0x7a, 0x68, 0xf5, 0x38, 0xe6, 0x25, 0xd3, 0xdd, 0xca, 0x68, 0x7c, 0x63, 0xe6, 0xa3, 0x8d,
	0x41, 0x87, 0x16, 0x26, 0x82, 0x4e, 0x74, 0x4d, 0x53, 0x5a, 0x28, 0x79, 0x63, 0x66, 0xa0, 0x8d,
	0x6f, 0xf7, 0xc1, 0x09, 0xec, 0x9b, 0xba, 0x23, 0xd8, 0x37, 0x9f, 0x4a, 0xab, 0xf7, 0xb0, 0x59,
	0x07, 0xbf, 0xc3, 0xd1, 0x7d, 0xa5, 0xd6, 0x09, 0x25, 0x79, 0x57, 0x92, 0x42, 0x85, 0x71, 0x5e,
	0x1a, 0x4a, 0x1d, 0x83, 0x0b, 0x5d, 0xe4, 0x25, 0x7a, 0x0a, 0xfd, 0x35, 0x20, 0x23, 0x99, 0x90,
	0xd6, 0x54, 0x1d, 0xf7, 0x5c, 0xf4, 0x8d, 0x09, 0x06, 0x43, 0xf0, 0xd6, 0x25, 0x43, 0x47, 0xd0,
	0xb4, 0xc5, 0xb5, 0x7f, 0xcf, 0x6e, 0x82, 0x7f, 0xea, 0xd0, 0xb8, 0x98, 0x4d, 0xd0, 0x37, 0xd0,
	0x7c, 0xc8, 0xd5, 0x61, 0x41, 0xdb, 0x1f, 0xab, 0xfe, 0xf0, 0x8f, 0x75, 0x77, 0xf2, 0x1b, 0x0f,
	0x9e, 0xfc, 0x2f, 0x61, 0x90, 0x98, 0x72, 0x70, 0x1a, 0xae, 0x28, 0x0f, 0x13, 0x2a, 0xcd, 0xbd,
	0xd1, 0xc6, 0x3d, 0x1d, 0xbf, 0xe0, 0xf4, 0x35, 0xe5, 0x3f, 0x53, 0xb9, 0x05, 0xe4, 0x44, 0x19,
	0x60, 0x73, 0x0b, 0x38, 0x23, 0x4a, 0x03, 0x7f, 0x84, 0xae, 0xed, 0x46, 0x9b, 0xd3, 0x6f, 0x19,
	0x2f, 0x9f, 0xef, 0x7a, 0xf9, 0x70, 0xcd, 0x62, 0x7b, 0xc9, 0xda, 0x4d, 0x00, 0xe0, 0x5d, 0x11,
	0xb5, 0x88, 0x64, 0x94, 0x05, 0xcf, 0xa1, 0x85, 0x49, 0x51, 0x32, 0x85, 0x06, 0xd0, 0xc8, 0x8a,
	0xd4, 0x15, 0x59, 0x2f, 0xd1, 0x31, 0x78, 0xb4, 0x08, 0x89, 0x94, 0x42, 0xba, 0xab, 0xfe, 0x80,
	0x16, 0x97, 0x7a, 0xfb, 0xd5, 0x29, 0x34, 0xae, 0xa7, 0x4b, 0xe4, 0xc1, 0xfe, 0x6c, 0x3e, 0xbb,
	0x1c, 0xec, 0xa1, 0x2e, 0x78, 0xf3, 0xc5, 0xf5, 0x64, 0x3e, 0x7b, 0x35, 0x1d, 0xd4, 0xc6, 0xff,
	0xd6, 0xa1, 0x63, 0x93, 0x2d, 0x95, 0x90, 0x04, 0xbd, 0x81, 0xc3, 0x2b, 0xa2, 0xb6, 0xde, 0x9d,
	0xdd, 0xb2, 0xaf, 0x3d, 0x9d, 0x9c, 0xee, 0x1e, 0x6d, 0x30, 0x83, 0x3d, 0xf4, 0x0b, 0xf4, 0x35,
	0x78, 0xe3, 0x19, 0xa9, 0x50, 0xab, 0x2c, 0x8c, 0x11, 0x3b, 0x5c, 0xde, 0xf1, 0x56, 0x6d, 0xe0,
	0xe4, 0xc9, 0x7d, 0xd7, 0x4e, 0xc9, 0x54, 0xb0, 0x87, 0x26, 0xd0, 0x5f, 0x6e, 0x3b, 0xab, 0x4c,
	0x5f, 0x21, 0xf5, 0xfa, 0xd9, 0x6f, 0xdf, 0xa5, 0x54, 0xfd, 0x51, 0xae, 0x46, 0xb1, 0xc8, 0xce,
	0x1d, 0x6c, 0xfd, 0x3b, 0x3e, 0x77, 0x3d, 0xc6, 0x88, 0x3c, 0x4f, 0x09, 0x77, 0x8f, 0xfe, 0xaa,
	0x65, 0x5e, 0xfd, 0x67, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x06, 0x32, 0xda, 0xa8, 0x0c, 0x08,
	0x00, 0x00,
}
