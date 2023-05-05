// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto

package dynamic_forward_proxyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/lugerinformatica/envoy-control-plane/envoy/annotations"
	v3 "github.com/lugerinformatica/envoy-control-plane/envoy/config/cluster/v3"
	v32 "github.com/lugerinformatica/envoy-control-plane/envoy/config/common/key_value/v3"
	v31 "github.com/lugerinformatica/envoy-control-plane/envoy/config/core/v3"
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

// Configuration of circuit breakers for resolver.
type DnsCacheCircuitBreakers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of pending requests that Envoy will allow to the
	// resolver. If not specified, the default is 1024.
	MaxPendingRequests *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
}

func (x *DnsCacheCircuitBreakers) Reset() {
	*x = DnsCacheCircuitBreakers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsCacheCircuitBreakers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsCacheCircuitBreakers) ProtoMessage() {}

func (x *DnsCacheCircuitBreakers) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsCacheCircuitBreakers.ProtoReflect.Descriptor instead.
func (*DnsCacheCircuitBreakers) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescGZIP(), []int{0}
}

func (x *DnsCacheCircuitBreakers) GetMaxPendingRequests() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxPendingRequests
	}
	return nil
}

// Configuration for the dynamic forward proxy DNS cache. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#next-free-field: 15]
type DnsCacheConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cache. Multiple named caches allow independent dynamic forward proxy
	// configurations to operate within a single Envoy process using different configurations. All
	// configurations with the same name *must* otherwise have the same settings when referenced
	// from different configuration components. Configuration will fail to load if this is not
	// the case.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The DNS lookup family to use during resolution.
	//
	// [#comment:TODO(mattklein123): Figure out how to support IPv4/IPv6 "happy eyeballs" mode. The
	// way this might work is a new lookup family which returns both IPv4 and IPv6 addresses, and
	// then configures a host to have a primary and fall back address. With this, we could very
	// likely build a "happy eyeballs" connection pool which would race the primary / fall back
	// address and return the one that wins. This same method could potentially also be used for
	// QUIC to TCP fall back.]
	DnsLookupFamily v3.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.config.cluster.v3.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	// The DNS refresh rate for unresolved DNS hosts. If not specified defaults to 60s.
	//
	// The refresh rate is rounded to the closest millisecond, and must be at least 1ms.
	//
	// Once a host has been resolved, the refresh rate will be the DNS TTL, capped
	// at a minimum of ``dns_min_refresh_rate``.
	DnsRefreshRate *duration.Duration `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	// The minimum rate that DNS resolution will occur. Per ``dns_refresh_rate``, once a host is
	// resolved, the DNS TTL will be used, with a minimum set by ``dns_min_refresh_rate``.
	// ``dns_min_refresh_rate`` defaults to 5s and must also be >= 5s.
	DnsMinRefreshRate *duration.Duration `protobuf:"bytes,14,opt,name=dns_min_refresh_rate,json=dnsMinRefreshRate,proto3" json:"dns_min_refresh_rate,omitempty"`
	// The TTL for hosts that are unused. Hosts that have not been used in the configured time
	// interval will be purged. If not specified defaults to 5m.
	//
	// .. note:
	//
	//   The TTL is only checked at the time of DNS refresh, as specified by ``dns_refresh_rate``. This
	//   means that if the configured TTL is shorter than the refresh rate the host may not be removed
	//   immediately.
	//
	//  .. note:
	//
	//   The TTL has no relation to DNS TTL and is only used to control Envoy's resource usage.
	HostTtl *duration.Duration `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	// The maximum number of hosts that the cache will hold. If not specified defaults to 1024.
	//
	// .. note:
	//
	//   The implementation is approximate and enforced independently on each worker thread, thus
	//   it is possible for the maximum hosts in the cache to go slightly above the configured
	//   value depending on timing. This is similar to how other circuit breakers work.
	MaxHosts *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	// If the DNS failure refresh rate is specified,
	// this is used as the cache's DNS refresh rate when DNS requests are failing. If this setting is
	// not specified, the failure refresh rate defaults to the dns_refresh_rate.
	DnsFailureRefreshRate *v3.Cluster_RefreshRate `protobuf:"bytes,6,opt,name=dns_failure_refresh_rate,json=dnsFailureRefreshRate,proto3" json:"dns_failure_refresh_rate,omitempty"`
	// The config of circuit breakers for resolver. It provides a configurable threshold.
	// Envoy will use dns cache circuit breakers with default settings even if this value is not set.
	DnsCacheCircuitBreaker *DnsCacheCircuitBreakers `protobuf:"bytes,7,opt,name=dns_cache_circuit_breaker,json=dnsCacheCircuitBreaker,proto3" json:"dns_cache_circuit_breaker,omitempty"`
	// Always use TCP queries instead of UDP queries for DNS lookups.
	// This field is deprecated in favor of ``dns_resolution_config``
	// which aggregates all of the DNS resolver configuration in a single message.
	//
	// Deprecated: Do not use.
	UseTcpForDnsLookups bool `protobuf:"varint,8,opt,name=use_tcp_for_dns_lookups,json=useTcpForDnsLookups,proto3" json:"use_tcp_for_dns_lookups,omitempty"`
	// DNS resolution configuration which includes the underlying dns resolver addresses and options.
	// This field is deprecated in favor of
	// :ref:`typed_dns_resolver_config <envoy_v3_api_field_extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.typed_dns_resolver_config>`.
	//
	// Deprecated: Do not use.
	DnsResolutionConfig *v31.DnsResolutionConfig `protobuf:"bytes,9,opt,name=dns_resolution_config,json=dnsResolutionConfig,proto3" json:"dns_resolution_config,omitempty"`
	// DNS resolver type configuration extension. This extension can be used to configure c-ares, apple,
	// or any other DNS resolver types and the related parameters.
	// For example, an object of
	// :ref:`CaresDnsResolverConfig <envoy_v3_api_msg_extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig>`
	// can be packed into this ``typed_dns_resolver_config``. This configuration replaces the
	// :ref:`dns_resolution_config <envoy_v3_api_field_extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_resolution_config>`
	// configuration.
	// During the transition period when both ``dns_resolution_config`` and ``typed_dns_resolver_config`` exists,
	// when ``typed_dns_resolver_config`` is in place, Envoy will use it and ignore ``dns_resolution_config``.
	// When ``typed_dns_resolver_config`` is missing, the default behavior is in place.
	// [#extension-category: envoy.network.dns_resolver]
	TypedDnsResolverConfig *v31.TypedExtensionConfig `protobuf:"bytes,12,opt,name=typed_dns_resolver_config,json=typedDnsResolverConfig,proto3" json:"typed_dns_resolver_config,omitempty"`
	// Hostnames that should be preresolved into the cache upon creation. This might provide a
	// performance improvement, in the form of cache hits, for hostnames that are going to be
	// resolved during steady state and are known at config load time.
	PreresolveHostnames []*v31.SocketAddress `protobuf:"bytes,10,rep,name=preresolve_hostnames,json=preresolveHostnames,proto3" json:"preresolve_hostnames,omitempty"`
	// The timeout used for DNS queries. This timeout is independent of any timeout and retry policy
	// used by the underlying DNS implementation (e.g., c-areas and Apple DNS) which are opaque.
	// Setting this timeout will ensure that queries succeed or fail within the specified time frame
	// and are then retried using the standard refresh rates. Defaults to 5s if not set.
	DnsQueryTimeout *duration.Duration `protobuf:"bytes,11,opt,name=dns_query_timeout,json=dnsQueryTimeout,proto3" json:"dns_query_timeout,omitempty"`
	// Configuration to flush the DNS cache to long term storage.
	KeyValueConfig *v32.KeyValueStoreConfig `protobuf:"bytes,13,opt,name=key_value_config,json=keyValueConfig,proto3" json:"key_value_config,omitempty"`
}

func (x *DnsCacheConfig) Reset() {
	*x = DnsCacheConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsCacheConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsCacheConfig) ProtoMessage() {}

func (x *DnsCacheConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsCacheConfig.ProtoReflect.Descriptor instead.
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescGZIP(), []int{1}
}

func (x *DnsCacheConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DnsCacheConfig) GetDnsLookupFamily() v3.Cluster_DnsLookupFamily {
	if x != nil {
		return x.DnsLookupFamily
	}
	return v3.Cluster_DnsLookupFamily(0)
}

func (x *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if x != nil {
		return x.DnsRefreshRate
	}
	return nil
}

func (x *DnsCacheConfig) GetDnsMinRefreshRate() *duration.Duration {
	if x != nil {
		return x.DnsMinRefreshRate
	}
	return nil
}

func (x *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if x != nil {
		return x.HostTtl
	}
	return nil
}

func (x *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxHosts
	}
	return nil
}

func (x *DnsCacheConfig) GetDnsFailureRefreshRate() *v3.Cluster_RefreshRate {
	if x != nil {
		return x.DnsFailureRefreshRate
	}
	return nil
}

func (x *DnsCacheConfig) GetDnsCacheCircuitBreaker() *DnsCacheCircuitBreakers {
	if x != nil {
		return x.DnsCacheCircuitBreaker
	}
	return nil
}

// Deprecated: Do not use.
func (x *DnsCacheConfig) GetUseTcpForDnsLookups() bool {
	if x != nil {
		return x.UseTcpForDnsLookups
	}
	return false
}

// Deprecated: Do not use.
func (x *DnsCacheConfig) GetDnsResolutionConfig() *v31.DnsResolutionConfig {
	if x != nil {
		return x.DnsResolutionConfig
	}
	return nil
}

func (x *DnsCacheConfig) GetTypedDnsResolverConfig() *v31.TypedExtensionConfig {
	if x != nil {
		return x.TypedDnsResolverConfig
	}
	return nil
}

func (x *DnsCacheConfig) GetPreresolveHostnames() []*v31.SocketAddress {
	if x != nil {
		return x.PreresolveHostnames
	}
	return nil
}

func (x *DnsCacheConfig) GetDnsQueryTimeout() *duration.Duration {
	if x != nil {
		return x.DnsQueryTimeout
	}
	return nil
}

func (x *DnsCacheConfig) GetKeyValueConfig() *v32.KeyValueStoreConfig {
	if x != nil {
		return x.KeyValueConfig
	}
	return nil
}

var File_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto protoreflect.FileDescriptor

var file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x33, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x17, 0x44, 0x6e, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x12, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x22, 0x9b, 0x0a, 0x0a, 0x0e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x64, 0x6e, 0x73, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x44, 0x6e, 0x73, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x64, 0x6e, 0x73,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x51, 0x0a, 0x10,
	0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0xaa, 0x01, 0x06, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52,
	0x0e, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x56, 0x0a, 0x14, 0x64, 0x6e, 0x73, 0x5f, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04,
	0x32, 0x02, 0x08, 0x05, 0x52, 0x11, 0x64, 0x6e, 0x73, 0x4d, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x5f,
	0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x07,
	0x68, 0x6f, 0x73, 0x74, 0x54, 0x74, 0x6c, 0x12, 0x42, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x68,
	0x6f, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x65, 0x0a, 0x18, 0x64,
	0x6e, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x15, 0x64, 0x6e, 0x73,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x19, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x73, 0x52, 0x16, 0x64, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x17, 0x75, 0x73, 0x65,
	0x5f, 0x74, 0x63, 0x70, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x6c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0b, 0x18, 0x01, 0x92, 0xc7,
	0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x52, 0x13, 0x75, 0x73, 0x65, 0x54, 0x63, 0x70, 0x46,
	0x6f, 0x72, 0x44, 0x6e, 0x73, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x6a, 0x0a, 0x15,
	0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0b, 0x18, 0x01, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03,
	0x33, 0x2e, 0x30, 0x52, 0x13, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x65, 0x0a, 0x19, 0x74, 0x79, 0x70, 0x65,
	0x64, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x16, 0x74, 0x79, 0x70, 0x65, 0x64, 0x44, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x56, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x13, 0x70, 0x72, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x48, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x11, 0x64, 0x6e, 0x73, 0x5f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00, 0x52, 0x0f, 0x64, 0x6e, 0x73, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x5f, 0x0a, 0x10, 0x6b, 0x65, 0x79, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x6b, 0x65, 0x79, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x47, 0x9a, 0xc5, 0x88, 0x1e, 0x42,
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0xca, 0x01, 0x0a, 0x3e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescOnce sync.Once
	file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescData = file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDesc
)

func file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescGZIP() []byte {
	file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescData)
	})
	return file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDescData
}

var file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_goTypes = []interface{}{
	(*DnsCacheCircuitBreakers)(nil),  // 0: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheCircuitBreakers
	(*DnsCacheConfig)(nil),           // 1: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	(*wrappers.UInt32Value)(nil),     // 2: google.protobuf.UInt32Value
	(v3.Cluster_DnsLookupFamily)(0),  // 3: envoy.config.cluster.v3.Cluster.DnsLookupFamily
	(*duration.Duration)(nil),        // 4: google.protobuf.Duration
	(*v3.Cluster_RefreshRate)(nil),   // 5: envoy.config.cluster.v3.Cluster.RefreshRate
	(*v31.DnsResolutionConfig)(nil),  // 6: envoy.config.core.v3.DnsResolutionConfig
	(*v31.TypedExtensionConfig)(nil), // 7: envoy.config.core.v3.TypedExtensionConfig
	(*v31.SocketAddress)(nil),        // 8: envoy.config.core.v3.SocketAddress
	(*v32.KeyValueStoreConfig)(nil),  // 9: envoy.config.common.key_value.v3.KeyValueStoreConfig
}
var file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_depIdxs = []int32{
	2,  // 0: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheCircuitBreakers.max_pending_requests:type_name -> google.protobuf.UInt32Value
	3,  // 1: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_lookup_family:type_name -> envoy.config.cluster.v3.Cluster.DnsLookupFamily
	4,  // 2: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_refresh_rate:type_name -> google.protobuf.Duration
	4,  // 3: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_min_refresh_rate:type_name -> google.protobuf.Duration
	4,  // 4: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.host_ttl:type_name -> google.protobuf.Duration
	2,  // 5: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.max_hosts:type_name -> google.protobuf.UInt32Value
	5,  // 6: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_failure_refresh_rate:type_name -> envoy.config.cluster.v3.Cluster.RefreshRate
	0,  // 7: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_cache_circuit_breaker:type_name -> envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheCircuitBreakers
	6,  // 8: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_resolution_config:type_name -> envoy.config.core.v3.DnsResolutionConfig
	7,  // 9: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.typed_dns_resolver_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	8,  // 10: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.preresolve_hostnames:type_name -> envoy.config.core.v3.SocketAddress
	4,  // 11: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.dns_query_timeout:type_name -> google.protobuf.Duration
	9,  // 12: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig.key_value_config:type_name -> envoy.config.common.key_value.v3.KeyValueStoreConfig
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_init() }
func file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_init() {
	if File_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsCacheCircuitBreakers); i {
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
		file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsCacheConfig); i {
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
			RawDescriptor: file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_msgTypes,
	}.Build()
	File_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto = out.File
	file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_rawDesc = nil
	file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_goTypes = nil
	file_envoy_extensions_common_dynamic_forward_proxy_v3_dns_cache_proto_depIdxs = nil
}
