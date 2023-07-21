//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesObservation) DeepCopyInto(out *FeaturesObservation) {
	*out = *in
	if in.Fuse != nil {
		in, out := &in.Fuse, &out.Fuse
		*out = new(bool)
		**out = **in
	}
	if in.Keyctl != nil {
		in, out := &in.Keyctl, &out.Keyctl
		*out = new(bool)
		**out = **in
	}
	if in.Mknod != nil {
		in, out := &in.Mknod, &out.Mknod
		*out = new(bool)
		**out = **in
	}
	if in.Mount != nil {
		in, out := &in.Mount, &out.Mount
		*out = new(string)
		**out = **in
	}
	if in.Nesting != nil {
		in, out := &in.Nesting, &out.Nesting
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesObservation.
func (in *FeaturesObservation) DeepCopy() *FeaturesObservation {
	if in == nil {
		return nil
	}
	out := new(FeaturesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesParameters) DeepCopyInto(out *FeaturesParameters) {
	*out = *in
	if in.Fuse != nil {
		in, out := &in.Fuse, &out.Fuse
		*out = new(bool)
		**out = **in
	}
	if in.Keyctl != nil {
		in, out := &in.Keyctl, &out.Keyctl
		*out = new(bool)
		**out = **in
	}
	if in.Mknod != nil {
		in, out := &in.Mknod, &out.Mknod
		*out = new(bool)
		**out = **in
	}
	if in.Mount != nil {
		in, out := &in.Mount, &out.Mount
		*out = new(string)
		**out = **in
	}
	if in.Nesting != nil {
		in, out := &in.Nesting, &out.Nesting
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesParameters.
func (in *FeaturesParameters) DeepCopy() *FeaturesParameters {
	if in == nil {
		return nil
	}
	out := new(FeaturesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lxc) DeepCopyInto(out *Lxc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lxc.
func (in *Lxc) DeepCopy() *Lxc {
	if in == nil {
		return nil
	}
	out := new(Lxc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Lxc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LxcList) DeepCopyInto(out *LxcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Lxc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LxcList.
func (in *LxcList) DeepCopy() *LxcList {
	if in == nil {
		return nil
	}
	out := new(LxcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LxcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LxcObservation) DeepCopyInto(out *LxcObservation) {
	*out = *in
	if in.Arch != nil {
		in, out := &in.Arch, &out.Arch
		*out = new(string)
		**out = **in
	}
	if in.Bwlimit != nil {
		in, out := &in.Bwlimit, &out.Bwlimit
		*out = new(float64)
		**out = **in
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = new(string)
		**out = **in
	}
	if in.CloneStorage != nil {
		in, out := &in.CloneStorage, &out.CloneStorage
		*out = new(string)
		**out = **in
	}
	if in.Cmode != nil {
		in, out := &in.Cmode, &out.Cmode
		*out = new(string)
		**out = **in
	}
	if in.Console != nil {
		in, out := &in.Console, &out.Console
		*out = new(bool)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.Cpulimit != nil {
		in, out := &in.Cpulimit, &out.Cpulimit
		*out = new(float64)
		**out = **in
	}
	if in.Cpuunits != nil {
		in, out := &in.Cpuunits, &out.Cpuunits
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]FeaturesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.Full != nil {
		in, out := &in.Full, &out.Full
		*out = new(bool)
		**out = **in
	}
	if in.Hagroup != nil {
		in, out := &in.Hagroup, &out.Hagroup
		*out = new(string)
		**out = **in
	}
	if in.Hastate != nil {
		in, out := &in.Hastate, &out.Hastate
		*out = new(string)
		**out = **in
	}
	if in.Hookscript != nil {
		in, out := &in.Hookscript, &out.Hookscript
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IgnoreUnpackErrors != nil {
		in, out := &in.IgnoreUnpackErrors, &out.IgnoreUnpackErrors
		*out = new(bool)
		**out = **in
	}
	if in.Lock != nil {
		in, out := &in.Lock, &out.Lock
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = make([]MountpointObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Nameserver != nil {
		in, out := &in.Nameserver, &out.Nameserver
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Onboot != nil {
		in, out := &in.Onboot, &out.Onboot
		*out = new(bool)
		**out = **in
	}
	if in.Ostemplate != nil {
		in, out := &in.Ostemplate, &out.Ostemplate
		*out = new(string)
		**out = **in
	}
	if in.Ostype != nil {
		in, out := &in.Ostype, &out.Ostype
		*out = new(string)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.Protection != nil {
		in, out := &in.Protection, &out.Protection
		*out = new(bool)
		**out = **in
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		*out = new(bool)
		**out = **in
	}
	if in.Rootfs != nil {
		in, out := &in.Rootfs, &out.Rootfs
		*out = make([]RootfsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSHPublicKeys != nil {
		in, out := &in.SSHPublicKeys, &out.SSHPublicKeys
		*out = new(string)
		**out = **in
	}
	if in.Searchdomain != nil {
		in, out := &in.Searchdomain, &out.Searchdomain
		*out = new(string)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(bool)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = new(string)
		**out = **in
	}
	if in.Swap != nil {
		in, out := &in.Swap, &out.Swap
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.TargetNode != nil {
		in, out := &in.TargetNode, &out.TargetNode
		*out = new(string)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(bool)
		**out = **in
	}
	if in.Tty != nil {
		in, out := &in.Tty, &out.Tty
		*out = new(float64)
		**out = **in
	}
	if in.Unique != nil {
		in, out := &in.Unique, &out.Unique
		*out = new(bool)
		**out = **in
	}
	if in.Unprivileged != nil {
		in, out := &in.Unprivileged, &out.Unprivileged
		*out = new(bool)
		**out = **in
	}
	if in.Unused != nil {
		in, out := &in.Unused, &out.Unused
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Vmid != nil {
		in, out := &in.Vmid, &out.Vmid
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LxcObservation.
func (in *LxcObservation) DeepCopy() *LxcObservation {
	if in == nil {
		return nil
	}
	out := new(LxcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LxcParameters) DeepCopyInto(out *LxcParameters) {
	*out = *in
	if in.Arch != nil {
		in, out := &in.Arch, &out.Arch
		*out = new(string)
		**out = **in
	}
	if in.Bwlimit != nil {
		in, out := &in.Bwlimit, &out.Bwlimit
		*out = new(float64)
		**out = **in
	}
	if in.Clone != nil {
		in, out := &in.Clone, &out.Clone
		*out = new(string)
		**out = **in
	}
	if in.CloneStorage != nil {
		in, out := &in.CloneStorage, &out.CloneStorage
		*out = new(string)
		**out = **in
	}
	if in.Cmode != nil {
		in, out := &in.Cmode, &out.Cmode
		*out = new(string)
		**out = **in
	}
	if in.Console != nil {
		in, out := &in.Console, &out.Console
		*out = new(bool)
		**out = **in
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.Cpulimit != nil {
		in, out := &in.Cpulimit, &out.Cpulimit
		*out = new(float64)
		**out = **in
	}
	if in.Cpuunits != nil {
		in, out := &in.Cpuunits, &out.Cpuunits
		*out = new(float64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]FeaturesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(bool)
		**out = **in
	}
	if in.Full != nil {
		in, out := &in.Full, &out.Full
		*out = new(bool)
		**out = **in
	}
	if in.Hagroup != nil {
		in, out := &in.Hagroup, &out.Hagroup
		*out = new(string)
		**out = **in
	}
	if in.Hastate != nil {
		in, out := &in.Hastate, &out.Hastate
		*out = new(string)
		**out = **in
	}
	if in.Hookscript != nil {
		in, out := &in.Hookscript, &out.Hookscript
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.IgnoreUnpackErrors != nil {
		in, out := &in.IgnoreUnpackErrors, &out.IgnoreUnpackErrors
		*out = new(bool)
		**out = **in
	}
	if in.Lock != nil {
		in, out := &in.Lock, &out.Lock
		*out = new(string)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(float64)
		**out = **in
	}
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = make([]MountpointParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Nameserver != nil {
		in, out := &in.Nameserver, &out.Nameserver
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = make([]NetworkParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Onboot != nil {
		in, out := &in.Onboot, &out.Onboot
		*out = new(bool)
		**out = **in
	}
	if in.Ostemplate != nil {
		in, out := &in.Ostemplate, &out.Ostemplate
		*out = new(string)
		**out = **in
	}
	if in.Ostype != nil {
		in, out := &in.Ostype, &out.Ostype
		*out = new(string)
		**out = **in
	}
	if in.PasswordSecretRef != nil {
		in, out := &in.PasswordSecretRef, &out.PasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Pool != nil {
		in, out := &in.Pool, &out.Pool
		*out = new(string)
		**out = **in
	}
	if in.Protection != nil {
		in, out := &in.Protection, &out.Protection
		*out = new(bool)
		**out = **in
	}
	if in.Restore != nil {
		in, out := &in.Restore, &out.Restore
		*out = new(bool)
		**out = **in
	}
	if in.Rootfs != nil {
		in, out := &in.Rootfs, &out.Rootfs
		*out = make([]RootfsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSHPublicKeys != nil {
		in, out := &in.SSHPublicKeys, &out.SSHPublicKeys
		*out = new(string)
		**out = **in
	}
	if in.Searchdomain != nil {
		in, out := &in.Searchdomain, &out.Searchdomain
		*out = new(string)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = new(bool)
		**out = **in
	}
	if in.Startup != nil {
		in, out := &in.Startup, &out.Startup
		*out = new(string)
		**out = **in
	}
	if in.Swap != nil {
		in, out := &in.Swap, &out.Swap
		*out = new(float64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(string)
		**out = **in
	}
	if in.TargetNode != nil {
		in, out := &in.TargetNode, &out.TargetNode
		*out = new(string)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(bool)
		**out = **in
	}
	if in.Tty != nil {
		in, out := &in.Tty, &out.Tty
		*out = new(float64)
		**out = **in
	}
	if in.Unique != nil {
		in, out := &in.Unique, &out.Unique
		*out = new(bool)
		**out = **in
	}
	if in.Unprivileged != nil {
		in, out := &in.Unprivileged, &out.Unprivileged
		*out = new(bool)
		**out = **in
	}
	if in.Vmid != nil {
		in, out := &in.Vmid, &out.Vmid
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LxcParameters.
func (in *LxcParameters) DeepCopy() *LxcParameters {
	if in == nil {
		return nil
	}
	out := new(LxcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LxcSpec) DeepCopyInto(out *LxcSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LxcSpec.
func (in *LxcSpec) DeepCopy() *LxcSpec {
	if in == nil {
		return nil
	}
	out := new(LxcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LxcStatus) DeepCopyInto(out *LxcStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LxcStatus.
func (in *LxcStatus) DeepCopy() *LxcStatus {
	if in == nil {
		return nil
	}
	out := new(LxcStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountpointObservation) DeepCopyInto(out *MountpointObservation) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(bool)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Mp != nil {
		in, out := &in.Mp, &out.Mp
		*out = new(string)
		**out = **in
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(bool)
		**out = **in
	}
	if in.Replicate != nil {
		in, out := &in.Replicate, &out.Replicate
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Slot != nil {
		in, out := &in.Slot, &out.Slot
		*out = new(float64)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountpointObservation.
func (in *MountpointObservation) DeepCopy() *MountpointObservation {
	if in == nil {
		return nil
	}
	out := new(MountpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountpointParameters) DeepCopyInto(out *MountpointParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(bool)
		**out = **in
	}
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = new(bool)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Mp != nil {
		in, out := &in.Mp, &out.Mp
		*out = new(string)
		**out = **in
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(bool)
		**out = **in
	}
	if in.Replicate != nil {
		in, out := &in.Replicate, &out.Replicate
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Slot != nil {
		in, out := &in.Slot, &out.Slot
		*out = new(float64)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountpointParameters.
func (in *MountpointParameters) DeepCopy() *MountpointParameters {
	if in == nil {
		return nil
	}
	out := new(MountpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkObservation) DeepCopyInto(out *NetworkObservation) {
	*out = *in
	if in.Bridge != nil {
		in, out := &in.Bridge, &out.Bridge
		*out = new(string)
		**out = **in
	}
	if in.Firewall != nil {
		in, out := &in.Firewall, &out.Firewall
		*out = new(bool)
		**out = **in
	}
	if in.Gw != nil {
		in, out := &in.Gw, &out.Gw
		*out = new(string)
		**out = **in
	}
	if in.Gw6 != nil {
		in, out := &in.Gw6, &out.Gw6
		*out = new(string)
		**out = **in
	}
	if in.Hwaddr != nil {
		in, out := &in.Hwaddr, &out.Hwaddr
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Ip6 != nil {
		in, out := &in.Ip6, &out.Ip6
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(float64)
		**out = **in
	}
	if in.Trunks != nil {
		in, out := &in.Trunks, &out.Trunks
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkObservation.
func (in *NetworkObservation) DeepCopy() *NetworkObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParameters) DeepCopyInto(out *NetworkParameters) {
	*out = *in
	if in.Bridge != nil {
		in, out := &in.Bridge, &out.Bridge
		*out = new(string)
		**out = **in
	}
	if in.Firewall != nil {
		in, out := &in.Firewall, &out.Firewall
		*out = new(bool)
		**out = **in
	}
	if in.Gw != nil {
		in, out := &in.Gw, &out.Gw
		*out = new(string)
		**out = **in
	}
	if in.Gw6 != nil {
		in, out := &in.Gw6, &out.Gw6
		*out = new(string)
		**out = **in
	}
	if in.Hwaddr != nil {
		in, out := &in.Hwaddr, &out.Hwaddr
		*out = new(string)
		**out = **in
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Ip6 != nil {
		in, out := &in.Ip6, &out.Ip6
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Rate != nil {
		in, out := &in.Rate, &out.Rate
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(float64)
		**out = **in
	}
	if in.Trunks != nil {
		in, out := &in.Trunks, &out.Trunks
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParameters.
func (in *NetworkParameters) DeepCopy() *NetworkParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootfsObservation) DeepCopyInto(out *RootfsObservation) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(bool)
		**out = **in
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(bool)
		**out = **in
	}
	if in.Replicate != nil {
		in, out := &in.Replicate, &out.Replicate
		*out = new(bool)
		**out = **in
	}
	if in.Ro != nil {
		in, out := &in.Ro, &out.Ro
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootfsObservation.
func (in *RootfsObservation) DeepCopy() *RootfsObservation {
	if in == nil {
		return nil
	}
	out := new(RootfsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootfsParameters) DeepCopyInto(out *RootfsParameters) {
	*out = *in
	if in.ACL != nil {
		in, out := &in.ACL, &out.ACL
		*out = new(bool)
		**out = **in
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(bool)
		**out = **in
	}
	if in.Replicate != nil {
		in, out := &in.Replicate, &out.Replicate
		*out = new(bool)
		**out = **in
	}
	if in.Ro != nil {
		in, out := &in.Ro, &out.Ro
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootfsParameters.
func (in *RootfsParameters) DeepCopy() *RootfsParameters {
	if in == nil {
		return nil
	}
	out := new(RootfsParameters)
	in.DeepCopyInto(out)
	return out
}