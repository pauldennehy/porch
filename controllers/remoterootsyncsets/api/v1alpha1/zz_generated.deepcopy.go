//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AggregatedStatus) DeepCopyInto(out *AggregatedStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AggregatedStatus.
func (in *AggregatedStatus) DeepCopy() *AggregatedStatus {
	if in == nil {
		return nil
	}
	out := new(AggregatedStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRef) DeepCopyInto(out *ClusterRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRef.
func (in *ClusterRef) DeepCopy() *ClusterRef {
	if in == nil {
		return nil
	}
	out := new(ClusterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCISpec) DeepCopyInto(out *OCISpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCISpec.
func (in *OCISpec) DeepCopy() *OCISpec {
	if in == nil {
		return nil
	}
	out := new(OCISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRef) DeepCopyInto(out *PackageRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRef.
func (in *PackageRef) DeepCopy() *PackageRef {
	if in == nil {
		return nil
	}
	out := new(PackageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncSet) DeepCopyInto(out *RemoteRootSyncSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncSet.
func (in *RemoteRootSyncSet) DeepCopy() *RemoteRootSyncSet {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteRootSyncSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncSetList) DeepCopyInto(out *RemoteRootSyncSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RemoteRootSyncSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncSetList.
func (in *RemoteRootSyncSetList) DeepCopy() *RemoteRootSyncSetList {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteRootSyncSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncSetSpec) DeepCopyInto(out *RemoteRootSyncSetSpec) {
	*out = *in
	if in.ClusterRefs != nil {
		in, out := &in.ClusterRefs, &out.ClusterRefs
		*out = make([]*ClusterRef, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClusterRef)
				**out = **in
			}
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(RootSyncTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncSetSpec.
func (in *RemoteRootSyncSetSpec) DeepCopy() *RemoteRootSyncSetSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteRootSyncSetStatus) DeepCopyInto(out *RemoteRootSyncSetStatus) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.AggregatedStatus.DeepCopyInto(&out.AggregatedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteRootSyncSetStatus.
func (in *RemoteRootSyncSetStatus) DeepCopy() *RemoteRootSyncSetStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteRootSyncSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootSyncTemplate) DeepCopyInto(out *RootSyncTemplate) {
	*out = *in
	if in.OCI != nil {
		in, out := &in.OCI, &out.OCI
		*out = new(OCISpec)
		**out = **in
	}
	if in.PackageRef != nil {
		in, out := &in.PackageRef, &out.PackageRef
		*out = new(PackageRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootSyncTemplate.
func (in *RootSyncTemplate) DeepCopy() *RootSyncTemplate {
	if in == nil {
		return nil
	}
	out := new(RootSyncTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetStatus) DeepCopyInto(out *TargetStatus) {
	*out = *in
	out.Ref = in.Ref
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetStatus.
func (in *TargetStatus) DeepCopy() *TargetStatus {
	if in == nil {
		return nil
	}
	out := new(TargetStatus)
	in.DeepCopyInto(out)
	return out
}
