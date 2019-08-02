// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotSchedule) DeepCopyInto(out *SnapshotSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotSchedule.
func (in *SnapshotSchedule) DeepCopy() *SnapshotSchedule {
	if in == nil {
		return nil
	}
	out := new(SnapshotSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotScheduleList) DeepCopyInto(out *SnapshotScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SnapshotSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotScheduleList.
func (in *SnapshotScheduleList) DeepCopy() *SnapshotScheduleList {
	if in == nil {
		return nil
	}
	out := new(SnapshotScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnapshotScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotScheduleSpec) DeepCopyInto(out *SnapshotScheduleSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotScheduleSpec.
func (in *SnapshotScheduleSpec) DeepCopy() *SnapshotScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(SnapshotScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotScheduleStatus) DeepCopyInto(out *SnapshotScheduleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotScheduleStatus.
func (in *SnapshotScheduleStatus) DeepCopy() *SnapshotScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(SnapshotScheduleStatus)
	in.DeepCopyInto(out)
	return out
}
