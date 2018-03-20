// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationSubject) DeepCopyInto(out *DestinationSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationSubject.
func (in *DestinationSubject) DeepCopy() *DestinationSubject {
	if in == nil {
		return nil
	}
	out := new(DestinationSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTrigger) DeepCopyInto(out *EventTrigger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTrigger.
func (in *EventTrigger) DeepCopy() *EventTrigger {
	if in == nil {
		return nil
	}
	out := new(EventTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OriginSubject) DeepCopyInto(out *OriginSubject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OriginSubject.
func (in *OriginSubject) DeepCopy() *OriginSubject {
	if in == nil {
		return nil
	}
	out := new(OriginSubject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessDelegation) DeepCopyInto(out *SubjectAccessDelegation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessDelegation.
func (in *SubjectAccessDelegation) DeepCopy() *SubjectAccessDelegation {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessDelegation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubjectAccessDelegation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessDelegationList) DeepCopyInto(out *SubjectAccessDelegationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubjectAccessDelegation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessDelegationList.
func (in *SubjectAccessDelegationList) DeepCopy() *SubjectAccessDelegationList {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessDelegationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubjectAccessDelegationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessDelegationSpec) DeepCopyInto(out *SubjectAccessDelegationSpec) {
	*out = *in
	out.OriginSubject = in.OriginSubject
	if in.DestinationSubjects != nil {
		in, out := &in.DestinationSubjects, &out.DestinationSubjects
		*out = make([]DestinationSubject, len(*in))
		copy(*out, *in)
	}
	if in.EventTriggers != nil {
		in, out := &in.EventTriggers, &out.EventTriggers
		*out = make([]EventTrigger, len(*in))
		copy(*out, *in)
	}
	if in.DeletionTriggers != nil {
		in, out := &in.DeletionTriggers, &out.DeletionTriggers
		*out = make([]EventTrigger, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessDelegationSpec.
func (in *SubjectAccessDelegationSpec) DeepCopy() *SubjectAccessDelegationSpec {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessDelegationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessDelegationStatus) DeepCopyInto(out *SubjectAccessDelegationStatus) {
	*out = *in
	if in.RoleBindings != nil {
		in, out := &in.RoleBindings, &out.RoleBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterRoleBindings != nil {
		in, out := &in.ClusterRoleBindings, &out.ClusterRoleBindings
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessDelegationStatus.
func (in *SubjectAccessDelegationStatus) DeepCopy() *SubjectAccessDelegationStatus {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessDelegationStatus)
	in.DeepCopyInto(out)
	return out
}
