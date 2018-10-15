// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/bundle/v1alpha1/object_meta.proto

package v1alpha1 // import "github.com/GoogleCloudPlatform/k8s-cluster-bundle/pkg/apis/bundle/v1alpha1"

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

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create.
type ObjectMeta struct {
	// Name must be unique within a namespace. Is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	// +optional
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// GenerateName is an optional prefix, used by the server, to generate a unique
	// name ONLY IF the Name field has not been provided.
	// If this field is used, the name returned to the client will be different
	// than the name passed. This value will also be combined with a unique suffix.
	// The provided value has the same validation rules as the Name field,
	// and may be truncated by the length of the suffix required to make the value
	// unique on the server.
	//
	// If this field is specified and the generated name exists, the server will
	// NOT return a 409 - instead, it will either return 201 Created or 500 with Reason
	// ServerTimeout indicating a unique name could not be found in the time allotted, and the client
	// should retry (optionally after the time indicated in the Retry-After header).
	//
	// Applied only if Name is not specified.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#idempotency
	// +optional
	GenerateName string `protobuf:"bytes,2,opt,name=generate_name,json=generateName,proto3" json:"generate_name,omitempty"`
	// Namespace defines the space within each name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	//
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/namespaces
	// +optional
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// SelfLink is a URL representing this object.
	// Populated by the system.
	// Read-only.
	// +optional
	SelfLink string `protobuf:"bytes,4,opt,name=self_link,json=selfLink,proto3" json:"self_link,omitempty"`
	// UID is the unique in time and space value for this object. It is typically generated by
	// the server on successful creation of a resource and is not allowed to change on PUT
	// operations.
	//
	// Populated by the system.
	// Read-only.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	// +optional
	Uid string `protobuf:"bytes,5,opt,name=uid,proto3" json:"uid,omitempty"`
	// An opaque value that represents the internal version of this object that can
	// be used by clients to determine when objects have changed. May be used for optimistic
	// concurrency, change detection, and the watch operation on a resource or set of resources.
	// Clients must treat these values as opaque and passed unmodified back to the server.
	// They may only be valid for a particular resource or set of resources.
	//
	// Populated by the system.
	// Read-only.
	// Value must be treated as opaque by clients and .
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	// +optional
	ResourceVersion string `protobuf:"bytes,6,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// A sequence number representing a specific generation of the desired state.
	// Populated by the system. Read-only.
	// +optional
	Generation int64 `protobuf:"varint,7,opt,name=generation,proto3" json:"generation,omitempty"`
	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system.
	// Read-only.
	// Null for lists.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	CreationTimestamp *Time `protobuf:"bytes,8,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	// DeletionTimestamp is RFC 3339 date and time at which this resource will be deleted. This
	// field is set by the server when a graceful deletion is requested by the user, and is not
	// directly settable by a client. The resource is expected to be deleted (no longer visible
	// from resource lists, and not reachable by name) after the time in this field, once the
	// finalizers list is empty. As long as the finalizers list contains items, deletion is blocked.
	// Once the deletionTimestamp is set, this value may not be unset or be set further into the
	// future, although it may be shortened or the resource may be deleted prior to this time.
	// For example, a user may request that a pod is deleted in 30 seconds. The Kubelet will react
	// by sending a graceful termination signal to the containers in the pod. After that 30 seconds,
	// the Kubelet will send a hard termination signal (SIGKILL) to the container and after cleanup,
	// remove the pod from the API. In the presence of network partitions, this object may still
	// exist after this timestamp, until an administrator or automated process can determine the
	// resource is fully terminated.
	// If not set, graceful deletion of the object has not been requested.
	//
	// Populated by the system when a graceful deletion is requested.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	DeletionTimestamp *Time `protobuf:"bytes,9,opt,name=deletion_timestamp,json=deletionTimestamp,proto3" json:"deletion_timestamp,omitempty"`
	// Number of seconds allowed for this object to gracefully terminate before
	// it will be removed from the system. Only set when deletionTimestamp is also set.
	// May only be shortened.
	// Read-only.
	// +optional
	DeletionGracePeriodSeconds int64 `protobuf:"varint,10,opt,name=deletion_grace_period_seconds,json=deletionGracePeriodSeconds,proto3" json:"deletion_grace_period_seconds,omitempty"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `protobuf:"bytes,12,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// List of objects depended by this object. If ALL objects in the list have
	// been deleted, this object will be garbage collected. If this object is managed by a controller,
	// then an entry in this list will point to this controller, with the controller field set to true.
	// There cannot be more than one managing controller.
	// +optional
	// +patchMergeKey=uid
	// +patchStrategy=merge
	OwnerReferences []*OwnerReference `protobuf:"bytes,13,rep,name=owner_references,json=ownerReferences,proto3" json:"owner_references,omitempty"`
	// Must be empty before the object is deleted from the registry. Each entry
	// is an identifier for the responsible component that will remove the entry
	// from the list. If the deletionTimestamp of the object is non-nil, entries
	// in this list can only be removed.
	// +optional
	// +patchStrategy=merge
	Finalizers []string `protobuf:"bytes,14,rep,name=finalizers,proto3" json:"finalizers,omitempty"`
	// The name of the cluster which the object belongs to.
	// This is used to distinguish resources with same name and namespace in different clusters.
	// This field is not set anywhere right now and apiserver is going to ignore it if set in create or update request.
	// +optional
	ClusterName          string   `protobuf:"bytes,15,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f6eff1922ec58c3, []int{0}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMeta.Unmarshal(m, b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return xxx_messageInfo_ObjectMeta.Size(m)
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMeta) GetGenerateName() string {
	if m != nil {
		return m.GenerateName
	}
	return ""
}

func (m *ObjectMeta) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ObjectMeta) GetSelfLink() string {
	if m != nil {
		return m.SelfLink
	}
	return ""
}

func (m *ObjectMeta) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ObjectMeta) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *ObjectMeta) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *ObjectMeta) GetCreationTimestamp() *Time {
	if m != nil {
		return m.CreationTimestamp
	}
	return nil
}

func (m *ObjectMeta) GetDeletionTimestamp() *Time {
	if m != nil {
		return m.DeletionTimestamp
	}
	return nil
}

func (m *ObjectMeta) GetDeletionGracePeriodSeconds() int64 {
	if m != nil {
		return m.DeletionGracePeriodSeconds
	}
	return 0
}

func (m *ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *ObjectMeta) GetOwnerReferences() []*OwnerReference {
	if m != nil {
		return m.OwnerReferences
	}
	return nil
}

func (m *ObjectMeta) GetFinalizers() []string {
	if m != nil {
		return m.Finalizers
	}
	return nil
}

func (m *ObjectMeta) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

// OwnerReference contains enough information to let you identify an owning
// object. Currently, an owning object must be in the same namespace, so there
// is no namespace field.
type OwnerReference struct {
	// API version of the referent.
	ApiVersion string `protobuf:"bytes,5,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Name of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// UID of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#uids
	Uid string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	// If true, this reference points to the managing controller.
	// +optional
	Controller bool `protobuf:"varint,6,opt,name=controller,proto3" json:"controller,omitempty"`
	// If true, AND if the owner has the "foregroundDeletion" finalizer, then
	// the owner cannot be deleted from the key-value store until this
	// reference is removed.
	// Defaults to false.
	// To set this field, a user needs "delete" permission of the owner,
	// otherwise 422 (Unprocessable Entity) will be returned.
	// +optional
	BlockOwnerDeletion   bool     `protobuf:"varint,7,opt,name=block_owner_deletion,json=blockOwnerDeletion,proto3" json:"block_owner_deletion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OwnerReference) Reset()         { *m = OwnerReference{} }
func (m *OwnerReference) String() string { return proto.CompactTextString(m) }
func (*OwnerReference) ProtoMessage()    {}
func (*OwnerReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f6eff1922ec58c3, []int{1}
}
func (m *OwnerReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerReference.Unmarshal(m, b)
}
func (m *OwnerReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerReference.Marshal(b, m, deterministic)
}
func (m *OwnerReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerReference.Merge(m, src)
}
func (m *OwnerReference) XXX_Size() int {
	return xxx_messageInfo_OwnerReference.Size(m)
}
func (m *OwnerReference) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerReference.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerReference proto.InternalMessageInfo

func (m *OwnerReference) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *OwnerReference) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *OwnerReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OwnerReference) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *OwnerReference) GetController() bool {
	if m != nil {
		return m.Controller
	}
	return false
}

func (m *OwnerReference) GetBlockOwnerDeletion() bool {
	if m != nil {
		return m.BlockOwnerDeletion
	}
	return false
}

// Time is a wrapper around time.Time which supports correct
// marshaling to YAML and JSON.  Wrappers are provided for many
// of the factory methods that the time package offers.
//
// +protobuf.options.marshal=false
// +protobuf.as=Timestamp
// +protobuf.options.(gogoproto.goproto_stringer)=false
type Time struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive. This field may be limited in precision depending on context.
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f6eff1922ec58c3, []int{2}
}
func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

func (m *Time) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Time) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*ObjectMeta)(nil), "pkg.apis.bundle.v1alpha1.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "pkg.apis.bundle.v1alpha1.ObjectMeta.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "pkg.apis.bundle.v1alpha1.ObjectMeta.LabelsEntry")
	proto.RegisterType((*OwnerReference)(nil), "pkg.apis.bundle.v1alpha1.OwnerReference")
	proto.RegisterType((*Time)(nil), "pkg.apis.bundle.v1alpha1.Time")
}

func init() {
	proto.RegisterFile("pkg/apis/bundle/v1alpha1/object_meta.proto", fileDescriptor_2f6eff1922ec58c3)
}

var fileDescriptor_2f6eff1922ec58c3 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x95, 0x49, 0x08, 0xc9, 0x35, 0x1f, 0x61, 0xc4, 0xc2, 0xe2, 0xbd, 0xc7, 0x4b, 0xe9, 0x26,
	0xad, 0x84, 0x03, 0x54, 0xad, 0x68, 0x17, 0x95, 0xe8, 0x87, 0xa8, 0x10, 0xb4, 0xc8, 0xa0, 0x56,
	0xea, 0xc6, 0x9a, 0xd8, 0x37, 0x61, 0xea, 0xf1, 0x8c, 0x35, 0x63, 0x53, 0xd1, 0xff, 0xd6, 0x3f,
	0xd6, 0x55, 0x35, 0x63, 0x0f, 0x09, 0x48, 0x54, 0x74, 0x77, 0xe7, 0xdc, 0x73, 0x8e, 0xe6, 0xe3,
	0xdc, 0x81, 0xa7, 0x45, 0x36, 0x1d, 0xd1, 0x82, 0xe9, 0xd1, 0xb8, 0x12, 0x29, 0xc7, 0xd1, 0xd5,
	0x1e, 0xe5, 0xc5, 0x25, 0xdd, 0x1b, 0xc9, 0xf1, 0x37, 0x4c, 0xca, 0x38, 0xc7, 0x92, 0x86, 0x85,
	0x92, 0xa5, 0x24, 0x41, 0x91, 0x4d, 0x43, 0xc3, 0x0d, 0x6b, 0x6e, 0xe8, 0xb8, 0xdb, 0xbf, 0x3a,
	0x00, 0x9f, 0x2c, 0xff, 0x14, 0x4b, 0x4a, 0x08, 0xb4, 0x05, 0xcd, 0x31, 0xf0, 0x06, 0xde, 0xb0,
	0x17, 0xd9, 0x9a, 0x3c, 0x86, 0x95, 0x29, 0x0a, 0x54, 0xb4, 0xc4, 0xd8, 0x36, 0x17, 0x6c, 0x73,
	0xd9, 0x81, 0x1f, 0x0d, 0xe9, 0x5f, 0xe8, 0x99, 0x9e, 0x2e, 0x68, 0x82, 0x41, 0xcb, 0x12, 0x66,
	0x00, 0xf9, 0x07, 0x7a, 0x1a, 0xf9, 0x24, 0xe6, 0x4c, 0x64, 0x41, 0xdb, 0x76, 0xbb, 0x06, 0x38,
	0x61, 0x22, 0x23, 0x7d, 0x68, 0x55, 0x2c, 0x0d, 0x16, 0x2d, 0x6c, 0x4a, 0xf2, 0x04, 0xfa, 0x0a,
	0xb5, 0xac, 0x54, 0x82, 0xf1, 0x15, 0x2a, 0xcd, 0xa4, 0x08, 0x3a, 0xb6, 0xbd, 0xe6, 0xf0, 0xcf,
	0x35, 0x4c, 0xb6, 0x00, 0x9a, 0x7d, 0x18, 0xd2, 0xd2, 0xc0, 0x1b, 0xb6, 0xa2, 0x39, 0x84, 0x9c,
	0x02, 0x49, 0x14, 0xda, 0x3a, 0x2e, 0x59, 0x8e, 0xba, 0xa4, 0x79, 0x11, 0x74, 0x07, 0xde, 0xd0,
	0xdf, 0xdf, 0x0a, 0xef, 0xbb, 0x96, 0xf0, 0x82, 0xe5, 0x18, 0xad, 0x3b, 0xe5, 0x85, 0x13, 0x1a,
	0xbb, 0x14, 0x39, 0xde, 0xb1, 0xeb, 0x3d, 0xcc, 0xce, 0x29, 0x67, 0x76, 0x87, 0xf0, 0xdf, 0x8d,
	0xdd, 0x54, 0xd1, 0x04, 0xe3, 0x02, 0x15, 0x93, 0x69, 0xac, 0x31, 0x91, 0x22, 0xd5, 0x01, 0xd8,
	0x03, 0x6d, 0x3a, 0xd2, 0x91, 0xe1, 0x9c, 0x59, 0xca, 0x79, 0xcd, 0x20, 0x1f, 0xa0, 0xc3, 0xe9,
	0x18, 0xb9, 0x0e, 0xfc, 0x41, 0x6b, 0xe8, 0xef, 0xef, 0xde, 0xbf, 0x8b, 0xd9, 0x3b, 0x87, 0x27,
	0x56, 0xf2, 0x5e, 0x94, 0xea, 0x3a, 0x6a, 0xf4, 0xe4, 0x0b, 0xf8, 0x54, 0x08, 0x59, 0xda, 0x23,
	0xeb, 0x60, 0xd9, 0xda, 0x3d, 0x7f, 0x90, 0xdd, 0xe1, 0x4c, 0x57, 0x7b, 0xce, 0x3b, 0x91, 0x73,
	0xe8, 0xcb, 0xef, 0x02, 0x55, 0xac, 0x70, 0x82, 0x0a, 0x45, 0x82, 0x3a, 0x58, 0xb1, 0xee, 0xc3,
	0x3f, 0xb8, 0x1b, 0x45, 0xe4, 0x04, 0xd1, 0x9a, 0xbc, 0xb5, 0xd6, 0xe6, 0xe1, 0x27, 0x4c, 0x50,
	0xce, 0x7e, 0xa0, 0xd2, 0xc1, 0xea, 0xa0, 0x35, 0xec, 0x45, 0x73, 0x08, 0x79, 0x04, 0xcb, 0x09,
	0xaf, 0x74, 0x89, 0xaa, 0x0e, 0xed, 0x9a, 0xcd, 0x8f, 0xdf, 0x60, 0x26, 0xb3, 0x9b, 0x2f, 0xc1,
	0x9f, 0xbb, 0x07, 0x93, 0xc3, 0x0c, 0xaf, 0x9b, 0xe8, 0x9b, 0x92, 0x6c, 0xc0, 0xe2, 0x15, 0xe5,
	0x95, 0x4b, 0x7c, 0xbd, 0x78, 0xb5, 0x70, 0xe0, 0x6d, 0xbe, 0x86, 0xfe, 0xdd, 0x33, 0xff, 0x8d,
	0xfe, 0xb8, 0xdd, 0xed, 0xf7, 0xd7, 0xb7, 0x7f, 0x7a, 0xb0, 0x7a, 0xfb, 0x9c, 0xe4, 0x7f, 0xf0,
	0x69, 0xc1, 0x6e, 0x52, 0x5f, 0x0f, 0x05, 0xd0, 0x82, 0xb9, 0xc0, 0x13, 0x68, 0x67, 0x4c, 0xa4,
	0x6e, 0x42, 0x4d, 0x7d, 0x33, 0xb5, 0xad, 0xb9, 0xa9, 0x6d, 0xa6, 0xaa, 0x3d, 0x9b, 0xaa, 0x2d,
	0x80, 0x44, 0x8a, 0x52, 0x49, 0xce, 0x51, 0xd9, 0x79, 0xea, 0x46, 0x73, 0x08, 0xd9, 0x85, 0x8d,
	0x31, 0x97, 0x49, 0x16, 0xd7, 0x8f, 0xe5, 0x32, 0x67, 0x87, 0xaa, 0x1b, 0x11, 0xdb, 0xb3, 0xbb,
	0x7d, 0xd7, 0x74, 0xb6, 0x5f, 0x40, 0xdb, 0x64, 0x99, 0x04, 0xb0, 0xe4, 0x02, 0xeb, 0xd9, 0xc0,
	0xba, 0xa5, 0xb9, 0x01, 0x41, 0x85, 0xd4, 0xf6, 0x06, 0x16, 0xa3, 0x7a, 0xf1, 0xe6, 0xe4, 0xeb,
	0xf1, 0x94, 0x95, 0x97, 0xd5, 0x38, 0x4c, 0x64, 0x3e, 0x3a, 0x92, 0x72, 0xca, 0xf1, 0x2d, 0x97,
	0x55, 0x7a, 0xc6, 0x69, 0x39, 0x91, 0x2a, 0x1f, 0x65, 0x07, 0x7a, 0xa7, 0x79, 0xaa, 0x9d, 0xe6,
	0x7b, 0xbb, 0xef, 0xbb, 0x1b, 0x77, 0xec, 0x1f, 0xf7, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xac, 0x58, 0x82, 0xd2, 0x11, 0x05, 0x00, 0x00,
}
