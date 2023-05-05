// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: envoy/extensions/access_loggers/grpc/v3/als.proto

package grpcv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/lugerinformatica/envoy-control-plane/envoy/config/core/v3"
	v31 "github.com/lugerinformatica/envoy-control-plane/envoy/type/tracing/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for the built-in ``envoy.access_loggers.http_grpc``
// :ref:`AccessLog <envoy_v3_api_msg_config.accesslog.v3.AccessLog>`. This configuration will
// populate :ref:`StreamAccessLogsMessage.http_logs
// <envoy_v3_api_field_service.accesslog.v3.StreamAccessLogsMessage.http_logs>`.
// [#extension: envoy.access_loggers.http_grpc]
type HttpGrpcAccessLogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
	// Additional request headers to log in :ref:`HTTPRequestProperties.request_headers
	// <envoy_v3_api_field_data.accesslog.v3.HTTPRequestProperties.request_headers>`.
	AdditionalRequestHeadersToLog []string `protobuf:"bytes,2,rep,name=additional_request_headers_to_log,json=additionalRequestHeadersToLog,proto3" json:"additional_request_headers_to_log,omitempty"`
	// Additional response headers to log in :ref:`HTTPResponseProperties.response_headers
	// <envoy_v3_api_field_data.accesslog.v3.HTTPResponseProperties.response_headers>`.
	AdditionalResponseHeadersToLog []string `protobuf:"bytes,3,rep,name=additional_response_headers_to_log,json=additionalResponseHeadersToLog,proto3" json:"additional_response_headers_to_log,omitempty"`
	// Additional response trailers to log in :ref:`HTTPResponseProperties.response_trailers
	// <envoy_v3_api_field_data.accesslog.v3.HTTPResponseProperties.response_trailers>`.
	AdditionalResponseTrailersToLog []string `protobuf:"bytes,4,rep,name=additional_response_trailers_to_log,json=additionalResponseTrailersToLog,proto3" json:"additional_response_trailers_to_log,omitempty"`
}

func (x *HttpGrpcAccessLogConfig) Reset() {
	*x = HttpGrpcAccessLogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpGrpcAccessLogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpGrpcAccessLogConfig) ProtoMessage() {}

func (x *HttpGrpcAccessLogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpGrpcAccessLogConfig.ProtoReflect.Descriptor instead.
func (*HttpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescGZIP(), []int{0}
}

func (x *HttpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if x != nil {
		return x.CommonConfig
	}
	return nil
}

func (x *HttpGrpcAccessLogConfig) GetAdditionalRequestHeadersToLog() []string {
	if x != nil {
		return x.AdditionalRequestHeadersToLog
	}
	return nil
}

func (x *HttpGrpcAccessLogConfig) GetAdditionalResponseHeadersToLog() []string {
	if x != nil {
		return x.AdditionalResponseHeadersToLog
	}
	return nil
}

func (x *HttpGrpcAccessLogConfig) GetAdditionalResponseTrailersToLog() []string {
	if x != nil {
		return x.AdditionalResponseTrailersToLog
	}
	return nil
}

// Configuration for the built-in ``envoy.access_loggers.tcp_grpc`` type. This configuration will
// populate ``StreamAccessLogsMessage.tcp_logs``.
// [#extension: envoy.access_loggers.tcp_grpc]
type TcpGrpcAccessLogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonConfig *CommonGrpcAccessLogConfig `protobuf:"bytes,1,opt,name=common_config,json=commonConfig,proto3" json:"common_config,omitempty"`
}

func (x *TcpGrpcAccessLogConfig) Reset() {
	*x = TcpGrpcAccessLogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TcpGrpcAccessLogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TcpGrpcAccessLogConfig) ProtoMessage() {}

func (x *TcpGrpcAccessLogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TcpGrpcAccessLogConfig.ProtoReflect.Descriptor instead.
func (*TcpGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescGZIP(), []int{1}
}

func (x *TcpGrpcAccessLogConfig) GetCommonConfig() *CommonGrpcAccessLogConfig {
	if x != nil {
		return x.CommonConfig
	}
	return nil
}

// Common configuration for gRPC access logs.
// [#next-free-field: 9]
type CommonGrpcAccessLogConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The friendly name of the access log to be returned in :ref:`StreamAccessLogsMessage.Identifier
	// <envoy_v3_api_msg_service.accesslog.v3.StreamAccessLogsMessage.Identifier>`. This allows the
	// access log server to differentiate between different access logs coming from the same Envoy.
	LogName string `protobuf:"bytes,1,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	// The gRPC service for the access log service.
	GrpcService *v3.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// API version for access logs service transport protocol. This describes the access logs service
	// gRPC endpoint and version of messages used on the wire.
	TransportApiVersion v3.ApiVersion `protobuf:"varint,6,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.config.core.v3.ApiVersion" json:"transport_api_version,omitempty"`
	// Interval for flushing access logs to the gRPC stream. Logger will flush requests every time
	// this interval is elapsed, or when batch size limit is hit, whichever comes first. Defaults to
	// 1 second.
	BufferFlushInterval *duration.Duration `protobuf:"bytes,3,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	// Soft size limit in bytes for access log entries buffer. Logger will buffer requests until
	// this limit it hit, or every time flush interval is elapsed, whichever comes first. Setting it
	// to zero effectively disables the batching. Defaults to 16384.
	BufferSizeBytes *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=buffer_size_bytes,json=bufferSizeBytes,proto3" json:"buffer_size_bytes,omitempty"`
	// Additional filter state objects to log in :ref:`filter_state_objects
	// <envoy_v3_api_field_data.accesslog.v3.AccessLogCommon.filter_state_objects>`.
	// Logger will call ``FilterState::Object::serializeAsProto`` to serialize the filter state object.
	FilterStateObjectsToLog []string `protobuf:"bytes,5,rep,name=filter_state_objects_to_log,json=filterStateObjectsToLog,proto3" json:"filter_state_objects_to_log,omitempty"`
	// Sets the retry policy when the establishment of a gRPC stream fails.
	// If the stream succeeds at least once in establishing itself,
	// no retry will be performed no matter what gRPC status is received.
	// Note that only :ref:`num_retries <envoy_v3_api_field_config.core.v3.RetryPolicy.num_retries>`
	// will be used in this configuration. This feature is used only when you are using
	// :ref:`Envoy gRPC client <envoy_v3_api_field_config.core.v3.GrpcService.envoy_grpc>`.
	GrpcStreamRetryPolicy *v3.RetryPolicy `protobuf:"bytes,7,opt,name=grpc_stream_retry_policy,json=grpcStreamRetryPolicy,proto3" json:"grpc_stream_retry_policy,omitempty"`
	// A list of custom tags with unique tag name to create tags for the logs.
	CustomTags []*v31.CustomTag `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
}

func (x *CommonGrpcAccessLogConfig) Reset() {
	*x = CommonGrpcAccessLogConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonGrpcAccessLogConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGrpcAccessLogConfig) ProtoMessage() {}

func (x *CommonGrpcAccessLogConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonGrpcAccessLogConfig.ProtoReflect.Descriptor instead.
func (*CommonGrpcAccessLogConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescGZIP(), []int{2}
}

func (x *CommonGrpcAccessLogConfig) GetLogName() string {
	if x != nil {
		return x.LogName
	}
	return ""
}

func (x *CommonGrpcAccessLogConfig) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *CommonGrpcAccessLogConfig) GetTransportApiVersion() v3.ApiVersion {
	if x != nil {
		return x.TransportApiVersion
	}
	return v3.ApiVersion(0)
}

func (x *CommonGrpcAccessLogConfig) GetBufferFlushInterval() *duration.Duration {
	if x != nil {
		return x.BufferFlushInterval
	}
	return nil
}

func (x *CommonGrpcAccessLogConfig) GetBufferSizeBytes() *wrappers.UInt32Value {
	if x != nil {
		return x.BufferSizeBytes
	}
	return nil
}

func (x *CommonGrpcAccessLogConfig) GetFilterStateObjectsToLog() []string {
	if x != nil {
		return x.FilterStateObjectsToLog
	}
	return nil
}

func (x *CommonGrpcAccessLogConfig) GetGrpcStreamRetryPolicy() *v3.RetryPolicy {
	if x != nil {
		return x.GrpcStreamRetryPolicy
	}
	return nil
}

func (x *CommonGrpcAccessLogConfig) GetCustomTags() []*v31.CustomTag {
	if x != nil {
		return x.CustomTags
	}
	return nil
}

var File_envoy_extensions_access_loggers_grpc_v3_als_proto protoreflect.FileDescriptor

var file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x17, 0x48, 0x74, 0x74, 0x70, 0x47, 0x72, 0x70, 0x63,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x71, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x48, 0x0a, 0x21, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1d, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x12, 0x4a, 0x0a, 0x22,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6c,
	0x6f, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x12, 0x4c, 0x0a, 0x23, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x1f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x73, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x3a, 0x38, 0x9a, 0xc5, 0x88, 0x1e, 0x33, 0x0a, 0x31, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x47, 0x72, 0x70,
	0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0xc4, 0x01, 0x0a, 0x16, 0x54, 0x63, 0x70, 0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x71, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x42, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01,
	0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x37,
	0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32,
	0x2e, 0x54, 0x63, 0x70, 0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xab, 0x05, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x67, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x15, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x41, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x41,
	0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x15, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x5f, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x13, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x48, 0x0a, 0x11, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x1b,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x17, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x12, 0x5a, 0x0a, 0x18, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x15, 0x67, 0x72, 0x70, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x41, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x61, 0x67, 0x73, 0x3a, 0x3a, 0x9a, 0xc5, 0x88, 0x1e, 0x35,
	0x0a, 0x33, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x47, 0x72, 0x70, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xa2, 0x01, 0x0a, 0x35, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x33, 0x42,
	0x08, 0x41, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x33, 0x3b, 0x67, 0x72, 0x70, 0x63,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescOnce sync.Once
	file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescData = file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDesc
)

func file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescGZIP() []byte {
	file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescData)
	})
	return file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDescData
}

var file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_access_loggers_grpc_v3_als_proto_goTypes = []interface{}{
	(*HttpGrpcAccessLogConfig)(nil),   // 0: envoy.extensions.access_loggers.grpc.v3.HttpGrpcAccessLogConfig
	(*TcpGrpcAccessLogConfig)(nil),    // 1: envoy.extensions.access_loggers.grpc.v3.TcpGrpcAccessLogConfig
	(*CommonGrpcAccessLogConfig)(nil), // 2: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig
	(*v3.GrpcService)(nil),            // 3: envoy.config.core.v3.GrpcService
	(v3.ApiVersion)(0),                // 4: envoy.config.core.v3.ApiVersion
	(*duration.Duration)(nil),         // 5: google.protobuf.Duration
	(*wrappers.UInt32Value)(nil),      // 6: google.protobuf.UInt32Value
	(*v3.RetryPolicy)(nil),            // 7: envoy.config.core.v3.RetryPolicy
	(*v31.CustomTag)(nil),             // 8: envoy.type.tracing.v3.CustomTag
}
var file_envoy_extensions_access_loggers_grpc_v3_als_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.access_loggers.grpc.v3.HttpGrpcAccessLogConfig.common_config:type_name -> envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig
	2, // 1: envoy.extensions.access_loggers.grpc.v3.TcpGrpcAccessLogConfig.common_config:type_name -> envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig
	3, // 2: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	4, // 3: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.transport_api_version:type_name -> envoy.config.core.v3.ApiVersion
	5, // 4: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.buffer_flush_interval:type_name -> google.protobuf.Duration
	6, // 5: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.buffer_size_bytes:type_name -> google.protobuf.UInt32Value
	7, // 6: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.grpc_stream_retry_policy:type_name -> envoy.config.core.v3.RetryPolicy
	8, // 7: envoy.extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.custom_tags:type_name -> envoy.type.tracing.v3.CustomTag
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_extensions_access_loggers_grpc_v3_als_proto_init() }
func file_envoy_extensions_access_loggers_grpc_v3_als_proto_init() {
	if File_envoy_extensions_access_loggers_grpc_v3_als_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpGrpcAccessLogConfig); i {
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
		file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TcpGrpcAccessLogConfig); i {
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
		file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonGrpcAccessLogConfig); i {
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
			RawDescriptor: file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_access_loggers_grpc_v3_als_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_access_loggers_grpc_v3_als_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_access_loggers_grpc_v3_als_proto_msgTypes,
	}.Build()
	File_envoy_extensions_access_loggers_grpc_v3_als_proto = out.File
	file_envoy_extensions_access_loggers_grpc_v3_als_proto_rawDesc = nil
	file_envoy_extensions_access_loggers_grpc_v3_als_proto_goTypes = nil
	file_envoy_extensions_access_loggers_grpc_v3_als_proto_depIdxs = nil
}
