// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1

import "istio.io/api/networking/v1alpha3"

// WorkloadEntry enables specifying the properties of a single non-Kubernetes workload such a VM or a bare metal services that can be referred to by service entries.
//
// <!-- crd generation tags
// +cue-gen:WorkloadEntry:groupName:networking.istio.io
// +cue-gen:WorkloadEntry:versions:v1beta1,v1alpha3,v1
// +cue-gen:WorkloadEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:WorkloadEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:WorkloadEntry:subresource:status
// +cue-gen:WorkloadEntry:scope:Namespaced
// +cue-gen:WorkloadEntry:resource:categories=istio-io,networking-istio-io,shortNames=we,plural=workloadentries
// +cue-gen:WorkloadEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:WorkloadEntry:printerColumn:name=Address,type=string,JSONPath=.spec.address,description="Address associated with the network endpoint."
// +cue-gen:WorkloadEntry:preserveUnknownFields:false
// +cue-gen:WorkloadEntry:spec:required
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
// +kubebuilder:validation:XValidation:message="Address is required",rule="has(self.address) || has(self.network)"
// +kubebuilder:validation:XValidation:message="UDS may not include ports",rule="(default(self.address, "").startsWith('unix://')) ? !has(self.ports) : true"
type WorkloadEntry = v1alpha3.WorkloadEntry
