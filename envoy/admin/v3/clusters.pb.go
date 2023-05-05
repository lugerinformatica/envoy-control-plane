// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: envoy/admin/v3/clusters.proto

package adminv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
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

// Admin endpoint uses this wrapper for ``/clusters`` to display cluster status information.
// See :ref:`/clusters <operations_admin_interface_clusters>` for more information.
type Clusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mapping from cluster name to each cluster's status.
	ClusterStatuses []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
}

func (x *Clusters) Reset() {
	*x = Clusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_clusters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clusters) ProtoMessage() {}

func (x *Clusters) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_clusters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clusters.ProtoReflect.Descriptor instead.
func (*Clusters) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_clusters_proto_rawDescGZIP(), []int{0}
}

func (x *Clusters) GetClusterStatuses() []*ClusterStatus {
	if x != nil {
		return x.ClusterStatuses
	}
	return nil
}

// Details an individual cluster's current status.
// [#next-free-field: 9]
type ClusterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Denotes whether this cluster was added via API or configured statically.
	AddedViaApi bool `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	// The success rate threshold used in the last interval.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``false``, all errors: externally and locally generated were used to calculate the threshold.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``true``, only externally generated errors were used to calculate the threshold.
	// The threshold is used to eject hosts based on their success rate. See
	// :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for details.
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	SuccessRateEjectionThreshold *v3.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	// Mapping from host address to the host's current status.
	HostStatuses []*HostStatus `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	// The success rate threshold used in the last interval when only locally originated failures were
	// taken into account and externally originated errors were treated as success.
	// This field should be interpreted only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``true``. The threshold is used to eject hosts based on their success rate.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	LocalOriginSuccessRateEjectionThreshold *v3.Percent `protobuf:"bytes,5,opt,name=local_origin_success_rate_ejection_threshold,json=localOriginSuccessRateEjectionThreshold,proto3" json:"local_origin_success_rate_ejection_threshold,omitempty"`
	// :ref:`Circuit breaking <arch_overview_circuit_break>` settings of the cluster.
	CircuitBreakers *v31.CircuitBreakers `protobuf:"bytes,6,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	// Observability name of the cluster.
	ObservabilityName string `protobuf:"bytes,7,opt,name=observability_name,json=observabilityName,proto3" json:"observability_name,omitempty"`
	// The :ref:`EDS service name <envoy_v3_api_field_config.cluster.v3.Cluster.EdsClusterConfig.service_name>` if the cluster is an EDS cluster.
	EdsServiceName string `protobuf:"bytes,8,opt,name=eds_service_name,json=edsServiceName,proto3" json:"eds_service_name,omitempty"`
}

func (x *ClusterStatus) Reset() {
	*x = ClusterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_clusters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatus) ProtoMessage() {}

func (x *ClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_clusters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatus.ProtoReflect.Descriptor instead.
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_clusters_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterStatus) GetAddedViaApi() bool {
	if x != nil {
		return x.AddedViaApi
	}
	return false
}

func (x *ClusterStatus) GetSuccessRateEjectionThreshold() *v3.Percent {
	if x != nil {
		return x.SuccessRateEjectionThreshold
	}
	return nil
}

func (x *ClusterStatus) GetHostStatuses() []*HostStatus {
	if x != nil {
		return x.HostStatuses
	}
	return nil
}

func (x *ClusterStatus) GetLocalOriginSuccessRateEjectionThreshold() *v3.Percent {
	if x != nil {
		return x.LocalOriginSuccessRateEjectionThreshold
	}
	return nil
}

func (x *ClusterStatus) GetCircuitBreakers() *v31.CircuitBreakers {
	if x != nil {
		return x.CircuitBreakers
	}
	return nil
}

func (x *ClusterStatus) GetObservabilityName() string {
	if x != nil {
		return x.ObservabilityName
	}
	return ""
}

func (x *ClusterStatus) GetEdsServiceName() string {
	if x != nil {
		return x.EdsServiceName
	}
	return ""
}

// Current state of a particular host.
// [#next-free-field: 10]
type HostStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address of this host.
	Address *v32.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// List of stats specific to this host.
	Stats []*SimpleMetric `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	// The host's current health status.
	HealthStatus *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	// Request success rate for this host over the last calculated interval.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``false``, all errors: externally and locally generated were used in success rate
	// calculation. If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``true``, only externally generated errors were used in success rate calculation.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	SuccessRate *v3.Percent `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	// The host's weight. If not configured, the value defaults to 1.
	Weight uint32 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// The hostname of the host, if applicable.
	Hostname string `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// The host's priority. If not configured, the value defaults to 0 (highest priority).
	Priority uint32 `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	// Request success rate for this host over the last calculated
	// interval when only locally originated errors are taken into account and externally originated
	// errors were treated as success.
	// This field should be interpreted only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_v3_api_field_config.cluster.v3.OutlierDetection.split_external_local_origin_errors>`
	// is ``true``.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	LocalOriginSuccessRate *v3.Percent `protobuf:"bytes,8,opt,name=local_origin_success_rate,json=localOriginSuccessRate,proto3" json:"local_origin_success_rate,omitempty"`
	// locality of the host.
	Locality *v32.Locality `protobuf:"bytes,9,opt,name=locality,proto3" json:"locality,omitempty"`
}

func (x *HostStatus) Reset() {
	*x = HostStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_clusters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostStatus) ProtoMessage() {}

func (x *HostStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_clusters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostStatus.ProtoReflect.Descriptor instead.
func (*HostStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_clusters_proto_rawDescGZIP(), []int{2}
}

func (x *HostStatus) GetAddress() *v32.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *HostStatus) GetStats() []*SimpleMetric {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *HostStatus) GetHealthStatus() *HostHealthStatus {
	if x != nil {
		return x.HealthStatus
	}
	return nil
}

func (x *HostStatus) GetSuccessRate() *v3.Percent {
	if x != nil {
		return x.SuccessRate
	}
	return nil
}

func (x *HostStatus) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *HostStatus) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *HostStatus) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *HostStatus) GetLocalOriginSuccessRate() *v3.Percent {
	if x != nil {
		return x.LocalOriginSuccessRate
	}
	return nil
}

func (x *HostStatus) GetLocality() *v32.Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

// Health status for a host.
// [#next-free-field: 9]
type HostHealthStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The host is currently failing active health checks.
	FailedActiveHealthCheck bool `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	// The host is currently considered an outlier and has been ejected.
	FailedOutlierCheck bool `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	// The host is currently being marked as degraded through active health checking.
	FailedActiveDegradedCheck bool `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	// The host has been removed from service discovery, but is being stabilized due to active
	// health checking.
	PendingDynamicRemoval bool `protobuf:"varint,5,opt,name=pending_dynamic_removal,json=pendingDynamicRemoval,proto3" json:"pending_dynamic_removal,omitempty"`
	// The host has not yet been health checked.
	PendingActiveHc bool `protobuf:"varint,6,opt,name=pending_active_hc,json=pendingActiveHc,proto3" json:"pending_active_hc,omitempty"`
	// The host should be excluded from panic, spillover, etc. calculations because it was explicitly
	// taken out of rotation via protocol signal and is not meant to be routed to.
	ExcludedViaImmediateHcFail bool `protobuf:"varint,7,opt,name=excluded_via_immediate_hc_fail,json=excludedViaImmediateHcFail,proto3" json:"excluded_via_immediate_hc_fail,omitempty"`
	// The host failed active HC due to timeout.
	ActiveHcTimeout bool `protobuf:"varint,8,opt,name=active_hc_timeout,json=activeHcTimeout,proto3" json:"active_hc_timeout,omitempty"`
	// Health status as reported by EDS. Note: only HEALTHY and UNHEALTHY are currently supported
	// here.
	// [#comment:TODO(mrice32): pipe through remaining EDS health status possibilities.]
	EdsHealthStatus v32.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.config.core.v3.HealthStatus" json:"eds_health_status,omitempty"`
}

func (x *HostHealthStatus) Reset() {
	*x = HostHealthStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v3_clusters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostHealthStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostHealthStatus) ProtoMessage() {}

func (x *HostHealthStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v3_clusters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostHealthStatus.ProtoReflect.Descriptor instead.
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v3_clusters_proto_rawDescGZIP(), []int{3}
}

func (x *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if x != nil {
		return x.FailedActiveHealthCheck
	}
	return false
}

func (x *HostHealthStatus) GetFailedOutlierCheck() bool {
	if x != nil {
		return x.FailedOutlierCheck
	}
	return false
}

func (x *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if x != nil {
		return x.FailedActiveDegradedCheck
	}
	return false
}

func (x *HostHealthStatus) GetPendingDynamicRemoval() bool {
	if x != nil {
		return x.PendingDynamicRemoval
	}
	return false
}

func (x *HostHealthStatus) GetPendingActiveHc() bool {
	if x != nil {
		return x.PendingActiveHc
	}
	return false
}

func (x *HostHealthStatus) GetExcludedViaImmediateHcFail() bool {
	if x != nil {
		return x.ExcludedViaImmediateHcFail
	}
	return false
}

func (x *HostHealthStatus) GetActiveHcTimeout() bool {
	if x != nil {
		return x.ActiveHcTimeout
	}
	return false
}

func (x *HostHealthStatus) GetEdsHealthStatus() v32.HealthStatus {
	if x != nil {
		return x.EdsHealthStatus
	}
	return v32.HealthStatus(0)
}

var File_envoy_admin_v3_clusters_proto protoreflect.FileDescriptor

var file_envoy_admin_v3_clusters_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x1a,
	0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5f, 0x62,
	0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x08, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x48, 0x0a, 0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x3a,
	0x23, 0x9a, 0xc5, 0x88, 0x1e, 0x1e, 0x0a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x22, 0xb6, 0x04, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x61, 0x64,
	0x64, 0x65, 0x64, 0x5f, 0x76, 0x69, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x64, 0x56, 0x69, 0x61, 0x41, 0x70, 0x69, 0x12, 0x5d,
	0x0a, 0x1f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52,
	0x1c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x45, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x3f, 0x0a,
	0x0d, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x75,
	0x0a, 0x2c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x27, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x61, 0x74, 0x65, 0x45, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x53, 0x0a, 0x10, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x5f, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x52, 0x0f, 0x63, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x42, 0x72, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x64, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x3a, 0x28, 0x9a, 0xc5, 0x88, 0x1e, 0x23, 0x0a, 0x21, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x81, 0x04,
	0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x37, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0d, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x33, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x39, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x0b,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x19, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x16, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x3a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x25, 0x9a, 0xc5, 0x88, 0x1e,
	0x20, 0x0a, 0x1e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x93, 0x04, 0x0a, 0x10, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x12, 0x30, 0x0a, 0x14, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x6f, 0x75,
	0x74, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x6c, 0x69, 0x65, 0x72,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x3f, 0x0a, 0x1c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x2a,
	0x0a, 0x11, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x68, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x48, 0x63, 0x12, 0x42, 0x0a, 0x1e, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f, 0x76, 0x69, 0x61, 0x5f, 0x69, 0x6d, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x63, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1a, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x56, 0x69, 0x61, 0x49,
	0x6d, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x48, 0x63, 0x46, 0x61, 0x69, 0x6c, 0x12, 0x2a,
	0x0a, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x68, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x48, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x4e, 0x0a, 0x11, 0x65, 0x64,
	0x73, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x65, 0x64, 0x73, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x2b, 0x9a, 0xc5, 0x88, 0x1e,
	0x26, 0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x76, 0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x33, 0x3b,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v3_clusters_proto_rawDescOnce sync.Once
	file_envoy_admin_v3_clusters_proto_rawDescData = file_envoy_admin_v3_clusters_proto_rawDesc
)

func file_envoy_admin_v3_clusters_proto_rawDescGZIP() []byte {
	file_envoy_admin_v3_clusters_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v3_clusters_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v3_clusters_proto_rawDescData)
	})
	return file_envoy_admin_v3_clusters_proto_rawDescData
}

var file_envoy_admin_v3_clusters_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_admin_v3_clusters_proto_goTypes = []interface{}{
	(*Clusters)(nil),            // 0: envoy.admin.v3.Clusters
	(*ClusterStatus)(nil),       // 1: envoy.admin.v3.ClusterStatus
	(*HostStatus)(nil),          // 2: envoy.admin.v3.HostStatus
	(*HostHealthStatus)(nil),    // 3: envoy.admin.v3.HostHealthStatus
	(*v3.Percent)(nil),          // 4: envoy.type.v3.Percent
	(*v31.CircuitBreakers)(nil), // 5: envoy.config.cluster.v3.CircuitBreakers
	(*v32.Address)(nil),         // 6: envoy.config.core.v3.Address
	(*SimpleMetric)(nil),        // 7: envoy.admin.v3.SimpleMetric
	(*v32.Locality)(nil),        // 8: envoy.config.core.v3.Locality
	(v32.HealthStatus)(0),       // 9: envoy.config.core.v3.HealthStatus
}
var file_envoy_admin_v3_clusters_proto_depIdxs = []int32{
	1,  // 0: envoy.admin.v3.Clusters.cluster_statuses:type_name -> envoy.admin.v3.ClusterStatus
	4,  // 1: envoy.admin.v3.ClusterStatus.success_rate_ejection_threshold:type_name -> envoy.type.v3.Percent
	2,  // 2: envoy.admin.v3.ClusterStatus.host_statuses:type_name -> envoy.admin.v3.HostStatus
	4,  // 3: envoy.admin.v3.ClusterStatus.local_origin_success_rate_ejection_threshold:type_name -> envoy.type.v3.Percent
	5,  // 4: envoy.admin.v3.ClusterStatus.circuit_breakers:type_name -> envoy.config.cluster.v3.CircuitBreakers
	6,  // 5: envoy.admin.v3.HostStatus.address:type_name -> envoy.config.core.v3.Address
	7,  // 6: envoy.admin.v3.HostStatus.stats:type_name -> envoy.admin.v3.SimpleMetric
	3,  // 7: envoy.admin.v3.HostStatus.health_status:type_name -> envoy.admin.v3.HostHealthStatus
	4,  // 8: envoy.admin.v3.HostStatus.success_rate:type_name -> envoy.type.v3.Percent
	4,  // 9: envoy.admin.v3.HostStatus.local_origin_success_rate:type_name -> envoy.type.v3.Percent
	8,  // 10: envoy.admin.v3.HostStatus.locality:type_name -> envoy.config.core.v3.Locality
	9,  // 11: envoy.admin.v3.HostHealthStatus.eds_health_status:type_name -> envoy.config.core.v3.HealthStatus
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_envoy_admin_v3_clusters_proto_init() }
func file_envoy_admin_v3_clusters_proto_init() {
	if File_envoy_admin_v3_clusters_proto != nil {
		return
	}
	file_envoy_admin_v3_metrics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v3_clusters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clusters); i {
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
		file_envoy_admin_v3_clusters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStatus); i {
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
		file_envoy_admin_v3_clusters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostStatus); i {
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
		file_envoy_admin_v3_clusters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostHealthStatus); i {
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
			RawDescriptor: file_envoy_admin_v3_clusters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v3_clusters_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v3_clusters_proto_depIdxs,
		MessageInfos:      file_envoy_admin_v3_clusters_proto_msgTypes,
	}.Build()
	File_envoy_admin_v3_clusters_proto = out.File
	file_envoy_admin_v3_clusters_proto_rawDesc = nil
	file_envoy_admin_v3_clusters_proto_goTypes = nil
	file_envoy_admin_v3_clusters_proto_depIdxs = nil
}
