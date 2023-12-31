/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FeaturesObservation struct {

	// A boolean for enabling FUSE mounts.
	Fuse *bool `json:"fuse,omitempty" tf:"fuse,omitempty"`

	// A boolean for enabling the keyctl() system call.
	Keyctl *bool `json:"keyctl,omitempty" tf:"keyctl,omitempty"`

	Mknod *bool `json:"mknod,omitempty" tf:"mknod,omitempty"`

	// Defines the filesystem types (separated by semicolons) that are allowed to be mounted.
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// A boolean to allow nested virtualization.
	Nesting *bool `json:"nesting,omitempty" tf:"nesting,omitempty"`
}

type FeaturesParameters struct {

	// A boolean for enabling FUSE mounts.
	// +kubebuilder:validation:Optional
	Fuse *bool `json:"fuse,omitempty" tf:"fuse,omitempty"`

	// A boolean for enabling the keyctl() system call.
	// +kubebuilder:validation:Optional
	Keyctl *bool `json:"keyctl,omitempty" tf:"keyctl,omitempty"`

	// +kubebuilder:validation:Optional
	Mknod *bool `json:"mknod,omitempty" tf:"mknod,omitempty"`

	// Defines the filesystem types (separated by semicolons) that are allowed to be mounted.
	// +kubebuilder:validation:Optional
	Mount *string `json:"mount,omitempty" tf:"mount,omitempty"`

	// A boolean to allow nested virtualization.
	// +kubebuilder:validation:Optional
	Nesting *bool `json:"nesting,omitempty" tf:"nesting,omitempty"`
}

type LxcObservation struct {

	// Sets the container OS architecture type. Default is "amd64".
	Arch *string `json:"arch,omitempty" tf:"arch,omitempty"`

	// A number for setting the override I/O bandwidth limit (in KiB/s).
	Bwlimit *float64 `json:"bwlimit,omitempty" tf:"bwlimit,omitempty"`

	// The lxc vmid to clone
	Clone *string `json:"clone,omitempty" tf:"clone,omitempty"`

	// Target storage for full clone.
	CloneStorage *string `json:"cloneStorage,omitempty" tf:"clone_storage,omitempty"`

	// Configures console mode. "tty" tries to open a connection to one of the available tty devices. "console"
	// tries to attach to /dev/console instead. "shell" simply invokes a shell inside the container (no login). Default
	// is "tty".
	Cmode *string `json:"cmode,omitempty" tf:"cmode,omitempty"`

	// A boolean to attach a console device to the container. Default is true.
	Console *bool `json:"console,omitempty" tf:"console,omitempty"`

	// The number of cores assigned to the container. A container can use all available cores by default.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// A number to limit CPU usage by. Default is 0.
	Cpulimit *float64 `json:"cpulimit,omitempty" tf:"cpulimit,omitempty"`

	// A number of the CPU weight that the container possesses. Default is 1024.
	Cpuunits *float64 `json:"cpuunits,omitempty" tf:"cpuunits,omitempty"`

	// Sets the container description seen in the web interface.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An object for allowing the container to access advanced features.
	Features []FeaturesObservation `json:"features,omitempty" tf:"features,omitempty"`

	// A boolean that allows the overwriting of pre-existing containers.
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// When cloning, create a full copy of all disks. This is always done when you clone a normal CT. For CT
	// template it creates a linked clone by default.
	Full *bool `json:"full,omitempty" tf:"full,omitempty"`

	// The HA group identifier the resource belongs to (requires hastate to be set!). See
	// the docs about HA for more
	// info.
	Hagroup *string `json:"hagroup,omitempty" tf:"hagroup,omitempty"`

	// Requested HA state for the resource. One of "started", "stopped", "enabled", "disabled", or "ignored". See
	// the docs about HA for more
	// info.
	Hastate *string `json:"hastate,omitempty" tf:"hastate,omitempty"`

	// A string
	// containing a volume identifier to a script
	// that will be executed during various steps throughout the container's lifetime. The script must be an executable file.
	Hookscript *string `json:"hookscript,omitempty" tf:"hookscript,omitempty"`

	// Specifies the host name of the container.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A boolean that determines if template extraction errors are ignored during container
	// creation.
	IgnoreUnpackErrors *bool `json:"ignoreUnpackErrors,omitempty" tf:"ignore_unpack_errors,omitempty"`

	// A string for locking or unlocking the VM.
	Lock *string `json:"lock,omitempty" tf:"lock,omitempty"`

	// A number containing the amount of RAM to assign to the container (in MB).
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// An object for defining a volume to use as a container mount point. Can be specified multiple times.
	Mountpoint []MountpointObservation `json:"mountpoint,omitempty" tf:"mountpoint,omitempty"`

	// The DNS server IP address used by the container. If neither nameserver nor searchdomain are
	// specified, the values of the Proxmox host will be used by default.
	Nameserver *string `json:"nameserver,omitempty" tf:"nameserver,omitempty"`

	// An object defining a network interface for the container. Can be specified multiple times.
	Network []NetworkObservation `json:"network,omitempty" tf:"network,omitempty"`

	// A boolean that determines if the container will start on boot. Default is false.
	Onboot *bool `json:"onboot,omitempty" tf:"onboot,omitempty"`

	// The volume identifier that points to
	// the OS template or backup file.
	Ostemplate *string `json:"ostemplate,omitempty" tf:"ostemplate,omitempty"`

	// The operating system type, used by LXC to set up and configure the container. Automatically determined if
	// not set.
	Ostype *string `json:"ostype,omitempty" tf:"ostype,omitempty"`

	// The name of the Proxmox resource pool to add this container to.
	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// A boolean that enables the protection flag on this container. Stops the container and its disk from
	// being removed/updated. Default is false.
	Protection *bool `json:"protection,omitempty" tf:"protection,omitempty"`

	// A boolean to mark the container creation/update as a restore task.
	Restore *bool `json:"restore,omitempty" tf:"restore,omitempty"`

	// An object for configuring the root mount point of the container. Can only be specified once.
	Rootfs []RootfsObservation `json:"rootfs,omitempty" tf:"rootfs,omitempty"`

	// Multi-line string of SSH public keys that will be added to the container. Can be defined
	// using heredoc syntax.
	SSHPublicKeys *string `json:"sshPublicKeys,omitempty" tf:"ssh_public_keys,omitempty"`

	// Sets the DNS search domains for the container. If neither nameserver nor searchdomain are
	// specified, the values of the Proxmox host will be used by default.
	Searchdomain *string `json:"searchdomain,omitempty" tf:"searchdomain,omitempty"`

	// A boolean that determines if the container is started after creation. Default is false.
	Start *bool `json:"start,omitempty" tf:"start,omitempty"`

	// The startup and shutdown behaviour
	// of the container.
	Startup *string `json:"startup,omitempty" tf:"startup,omitempty"`

	// A number that sets the amount of swap memory available to the container. Default is 512.
	Swap *float64 `json:"swap,omitempty" tf:"swap,omitempty"`

	// Tags of the container. This is only meta information.
	Tags *string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A string containing the cluster node name.
	TargetNode *string `json:"targetNode,omitempty" tf:"target_node,omitempty"`

	// A boolean that determines if this container is a template.
	Template *bool `json:"template,omitempty" tf:"template,omitempty"`

	// A number that specifies the TTYs available to the container. Default is 2.
	Tty *float64 `json:"tty,omitempty" tf:"tty,omitempty"`

	// A boolean that determines if a unique random ethernet address is assigned to the container.
	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`

	// A boolean that makes the container run as an unprivileged user. Default is false.
	Unprivileged *bool `json:"unprivileged,omitempty" tf:"unprivileged,omitempty"`

	Unused []*string `json:"unused,omitempty" tf:"unused,omitempty"`

	// A number that sets the VMID of the container. If set to 0, the next available VMID is used. Default is 0.
	// The VM identifier in proxmox (100-999999999)
	Vmid *float64 `json:"vmid,omitempty" tf:"vmid,omitempty"`
}

type LxcParameters struct {

	// Sets the container OS architecture type. Default is "amd64".
	// +kubebuilder:validation:Optional
	Arch *string `json:"arch,omitempty" tf:"arch,omitempty"`

	// A number for setting the override I/O bandwidth limit (in KiB/s).
	// +kubebuilder:validation:Optional
	Bwlimit *float64 `json:"bwlimit,omitempty" tf:"bwlimit,omitempty"`

	// The lxc vmid to clone
	// +kubebuilder:validation:Optional
	Clone *string `json:"clone,omitempty" tf:"clone,omitempty"`

	// Target storage for full clone.
	// +kubebuilder:validation:Optional
	CloneStorage *string `json:"cloneStorage,omitempty" tf:"clone_storage,omitempty"`

	// Configures console mode. "tty" tries to open a connection to one of the available tty devices. "console"
	// tries to attach to /dev/console instead. "shell" simply invokes a shell inside the container (no login). Default
	// is "tty".
	// +kubebuilder:validation:Optional
	Cmode *string `json:"cmode,omitempty" tf:"cmode,omitempty"`

	// A boolean to attach a console device to the container. Default is true.
	// +kubebuilder:validation:Optional
	Console *bool `json:"console,omitempty" tf:"console,omitempty"`

	// The number of cores assigned to the container. A container can use all available cores by default.
	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// A number to limit CPU usage by. Default is 0.
	// +kubebuilder:validation:Optional
	Cpulimit *float64 `json:"cpulimit,omitempty" tf:"cpulimit,omitempty"`

	// A number of the CPU weight that the container possesses. Default is 1024.
	// +kubebuilder:validation:Optional
	Cpuunits *float64 `json:"cpuunits,omitempty" tf:"cpuunits,omitempty"`

	// Sets the container description seen in the web interface.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An object for allowing the container to access advanced features.
	// +kubebuilder:validation:Optional
	Features []FeaturesParameters `json:"features,omitempty" tf:"features,omitempty"`

	// A boolean that allows the overwriting of pre-existing containers.
	// +kubebuilder:validation:Optional
	Force *bool `json:"force,omitempty" tf:"force,omitempty"`

	// When cloning, create a full copy of all disks. This is always done when you clone a normal CT. For CT
	// template it creates a linked clone by default.
	// +kubebuilder:validation:Optional
	Full *bool `json:"full,omitempty" tf:"full,omitempty"`

	// The HA group identifier the resource belongs to (requires hastate to be set!). See
	// the docs about HA for more
	// info.
	// +kubebuilder:validation:Optional
	Hagroup *string `json:"hagroup,omitempty" tf:"hagroup,omitempty"`

	// Requested HA state for the resource. One of "started", "stopped", "enabled", "disabled", or "ignored". See
	// the docs about HA for more
	// info.
	// +kubebuilder:validation:Optional
	Hastate *string `json:"hastate,omitempty" tf:"hastate,omitempty"`

	// A string
	// containing a volume identifier to a script
	// that will be executed during various steps throughout the container's lifetime. The script must be an executable file.
	// +kubebuilder:validation:Optional
	Hookscript *string `json:"hookscript,omitempty" tf:"hookscript,omitempty"`

	// Specifies the host name of the container.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// A boolean that determines if template extraction errors are ignored during container
	// creation.
	// +kubebuilder:validation:Optional
	IgnoreUnpackErrors *bool `json:"ignoreUnpackErrors,omitempty" tf:"ignore_unpack_errors,omitempty"`

	// A string for locking or unlocking the VM.
	// +kubebuilder:validation:Optional
	Lock *string `json:"lock,omitempty" tf:"lock,omitempty"`

	// A number containing the amount of RAM to assign to the container (in MB).
	// +kubebuilder:validation:Optional
	Memory *float64 `json:"memory,omitempty" tf:"memory,omitempty"`

	// An object for defining a volume to use as a container mount point. Can be specified multiple times.
	// +kubebuilder:validation:Optional
	Mountpoint []MountpointParameters `json:"mountpoint,omitempty" tf:"mountpoint,omitempty"`

	// The DNS server IP address used by the container. If neither nameserver nor searchdomain are
	// specified, the values of the Proxmox host will be used by default.
	// +kubebuilder:validation:Optional
	Nameserver *string `json:"nameserver,omitempty" tf:"nameserver,omitempty"`

	// An object defining a network interface for the container. Can be specified multiple times.
	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network,omitempty"`

	// A boolean that determines if the container will start on boot. Default is false.
	// +kubebuilder:validation:Optional
	Onboot *bool `json:"onboot,omitempty" tf:"onboot,omitempty"`

	// The volume identifier that points to
	// the OS template or backup file.
	// +kubebuilder:validation:Optional
	Ostemplate *string `json:"ostemplate,omitempty" tf:"ostemplate,omitempty"`

	// The operating system type, used by LXC to set up and configure the container. Automatically determined if
	// not set.
	// +kubebuilder:validation:Optional
	Ostype *string `json:"ostype,omitempty" tf:"ostype,omitempty"`

	// Sets the root password inside the container.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The name of the Proxmox resource pool to add this container to.
	// +kubebuilder:validation:Optional
	Pool *string `json:"pool,omitempty" tf:"pool,omitempty"`

	// A boolean that enables the protection flag on this container. Stops the container and its disk from
	// being removed/updated. Default is false.
	// +kubebuilder:validation:Optional
	Protection *bool `json:"protection,omitempty" tf:"protection,omitempty"`

	// A boolean to mark the container creation/update as a restore task.
	// +kubebuilder:validation:Optional
	Restore *bool `json:"restore,omitempty" tf:"restore,omitempty"`

	// An object for configuring the root mount point of the container. Can only be specified once.
	// +kubebuilder:validation:Optional
	Rootfs []RootfsParameters `json:"rootfs,omitempty" tf:"rootfs,omitempty"`

	// Multi-line string of SSH public keys that will be added to the container. Can be defined
	// using heredoc syntax.
	// +kubebuilder:validation:Optional
	SSHPublicKeys *string `json:"sshPublicKeys,omitempty" tf:"ssh_public_keys,omitempty"`

	// Sets the DNS search domains for the container. If neither nameserver nor searchdomain are
	// specified, the values of the Proxmox host will be used by default.
	// +kubebuilder:validation:Optional
	Searchdomain *string `json:"searchdomain,omitempty" tf:"searchdomain,omitempty"`

	// A boolean that determines if the container is started after creation. Default is false.
	// +kubebuilder:validation:Optional
	Start *bool `json:"start,omitempty" tf:"start,omitempty"`

	// The startup and shutdown behaviour
	// of the container.
	// +kubebuilder:validation:Optional
	Startup *string `json:"startup,omitempty" tf:"startup,omitempty"`

	// A number that sets the amount of swap memory available to the container. Default is 512.
	// +kubebuilder:validation:Optional
	Swap *float64 `json:"swap,omitempty" tf:"swap,omitempty"`

	// Tags of the container. This is only meta information.
	// +kubebuilder:validation:Optional
	Tags *string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A string containing the cluster node name.
	// +kubebuilder:validation:Optional
	TargetNode *string `json:"targetNode,omitempty" tf:"target_node,omitempty"`

	// A boolean that determines if this container is a template.
	// +kubebuilder:validation:Optional
	Template *bool `json:"template,omitempty" tf:"template,omitempty"`

	// A number that specifies the TTYs available to the container. Default is 2.
	// +kubebuilder:validation:Optional
	Tty *float64 `json:"tty,omitempty" tf:"tty,omitempty"`

	// A boolean that determines if a unique random ethernet address is assigned to the container.
	// +kubebuilder:validation:Optional
	Unique *bool `json:"unique,omitempty" tf:"unique,omitempty"`

	// A boolean that makes the container run as an unprivileged user. Default is false.
	// +kubebuilder:validation:Optional
	Unprivileged *bool `json:"unprivileged,omitempty" tf:"unprivileged,omitempty"`

	// A number that sets the VMID of the container. If set to 0, the next available VMID is used. Default is 0.
	// The VM identifier in proxmox (100-999999999)
	// +kubebuilder:validation:Optional
	Vmid *float64 `json:"vmid,omitempty" tf:"vmid,omitempty"`
}

type MountpointObservation struct {

	// A boolean for enabling ACL support. Default is false.
	ACL *bool `json:"acl,omitempty" tf:"acl,omitempty"`

	// A boolean for including the mount point in backups. Default is false.
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// The number that identifies the mount point (i.e. the n
	// in mp[n]).
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The path to the mount point as seen from inside the container. The path must not contain
	// symlinks for security reasons.
	Mp *string `json:"mp,omitempty" tf:"mp,omitempty"`

	// A boolean for enabling user quotas inside the container for this mount point. Default is false.
	Quota *bool `json:"quota,omitempty" tf:"quota,omitempty"`

	// A boolean for including this volume in a storage replica job. Default is false.
	Replicate *bool `json:"replicate,omitempty" tf:"replicate,omitempty"`

	// A boolean for marking the volume as available on all nodes. Default is false.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// Size of the underlying volume. Must end in T, G, M, or K (e.g. "1T", "1G", "1024M"
	// , "1048576K"). Note that this is a read only value.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// A string containing the number that identifies the mount point (i.e. the n
	// in mp[n]).
	Slot *float64 `json:"slot,omitempty" tf:"slot,omitempty"`

	// A string containing
	// the volume
	// , directory,
	// or device to be mounted into the
	// container (at the path specified by mp). E.g. local-lvm, local-zfs, local etc.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type MountpointParameters struct {

	// A boolean for enabling ACL support. Default is false.
	// +kubebuilder:validation:Optional
	ACL *bool `json:"acl,omitempty" tf:"acl,omitempty"`

	// A boolean for including the mount point in backups. Default is false.
	// +kubebuilder:validation:Optional
	Backup *bool `json:"backup,omitempty" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// The number that identifies the mount point (i.e. the n
	// in mp[n]).
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// The path to the mount point as seen from inside the container. The path must not contain
	// symlinks for security reasons.
	// +kubebuilder:validation:Required
	Mp *string `json:"mp" tf:"mp,omitempty"`

	// A boolean for enabling user quotas inside the container for this mount point. Default is false.
	// +kubebuilder:validation:Optional
	Quota *bool `json:"quota,omitempty" tf:"quota,omitempty"`

	// A boolean for including this volume in a storage replica job. Default is false.
	// +kubebuilder:validation:Optional
	Replicate *bool `json:"replicate,omitempty" tf:"replicate,omitempty"`

	// A boolean for marking the volume as available on all nodes. Default is false.
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// Size of the underlying volume. Must end in T, G, M, or K (e.g. "1T", "1G", "1024M"
	// , "1048576K"). Note that this is a read only value.
	// +kubebuilder:validation:Required
	Size *string `json:"size" tf:"size,omitempty"`

	// A string containing the number that identifies the mount point (i.e. the n
	// in mp[n]).
	// +kubebuilder:validation:Required
	Slot *float64 `json:"slot" tf:"slot,omitempty"`

	// A string containing
	// the volume
	// , directory,
	// or device to be mounted into the
	// container (at the path specified by mp). E.g. local-lvm, local-zfs, local etc.
	// +kubebuilder:validation:Required
	Storage *string `json:"storage" tf:"storage,omitempty"`

	// +kubebuilder:validation:Optional
	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type NetworkObservation struct {

	// The bridge to attach the network interface to (e.g. "vmbr0").
	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	// A boolean to enable the firewall on the network interface.
	Firewall *bool `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// The IPv4 address belonging to the network interface's default gateway.
	Gw *string `json:"gw,omitempty" tf:"gw,omitempty"`

	// The IPv6 address of the network interface's default gateway.
	Gw6 *string `json:"gw6,omitempty" tf:"gw6,omitempty"`

	// A string to set a common MAC address with the I/G (Individual/Group) bit not set. Automatically
	// determined if not set.
	Hwaddr *string `json:"hwaddr,omitempty" tf:"hwaddr,omitempty"`

	// The IPv4 address of the network interface. Can be a static IPv4 address (in CIDR notation), "dhcp",
	// or "manual".
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The IPv6 address of the network interface. Can be a static IPv6 address (in CIDR notation), "auto"
	// , "dhcp", or "manual".
	Ip6 *string `json:"ip6,omitempty" tf:"ip6,omitempty"`

	// A string to set the MTU on the network interface.
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the network interface as seen from inside the container (e.g. "eth0").
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A number that sets rate limiting on the network interface (Mbps).
	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// A number that specifies the VLAN tag of the network interface. Automatically determined if not set.
	Tag *float64 `json:"tag,omitempty" tf:"tag,omitempty"`

	Trunks *string `json:"trunks,omitempty" tf:"trunks,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkParameters struct {

	// The bridge to attach the network interface to (e.g. "vmbr0").
	// +kubebuilder:validation:Optional
	Bridge *string `json:"bridge,omitempty" tf:"bridge,omitempty"`

	// A boolean to enable the firewall on the network interface.
	// +kubebuilder:validation:Optional
	Firewall *bool `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// The IPv4 address belonging to the network interface's default gateway.
	// +kubebuilder:validation:Optional
	Gw *string `json:"gw,omitempty" tf:"gw,omitempty"`

	// The IPv6 address of the network interface's default gateway.
	// +kubebuilder:validation:Optional
	Gw6 *string `json:"gw6,omitempty" tf:"gw6,omitempty"`

	// A string to set a common MAC address with the I/G (Individual/Group) bit not set. Automatically
	// determined if not set.
	// +kubebuilder:validation:Optional
	Hwaddr *string `json:"hwaddr,omitempty" tf:"hwaddr,omitempty"`

	// The IPv4 address of the network interface. Can be a static IPv4 address (in CIDR notation), "dhcp",
	// or "manual".
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// The IPv6 address of the network interface. Can be a static IPv6 address (in CIDR notation), "auto"
	// , "dhcp", or "manual".
	// +kubebuilder:validation:Optional
	Ip6 *string `json:"ip6,omitempty" tf:"ip6,omitempty"`

	// A string to set the MTU on the network interface.
	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// The name of the network interface as seen from inside the container (e.g. "eth0").
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A number that sets rate limiting on the network interface (Mbps).
	// +kubebuilder:validation:Optional
	Rate *float64 `json:"rate,omitempty" tf:"rate,omitempty"`

	// A number that specifies the VLAN tag of the network interface. Automatically determined if not set.
	// +kubebuilder:validation:Optional
	Tag *float64 `json:"tag,omitempty" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Trunks *string `json:"trunks,omitempty" tf:"trunks,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RootfsObservation struct {

	// A boolean for enabling ACL support. Default is false.
	ACL *bool `json:"acl,omitempty" tf:"acl,omitempty"`

	// A boolean for enabling user quotas inside the container for this mount point. Default is false.
	Quota *bool `json:"quota,omitempty" tf:"quota,omitempty"`

	// A boolean for including this volume in a storage replica job. Default is false.
	Replicate *bool `json:"replicate,omitempty" tf:"replicate,omitempty"`

	Ro *bool `json:"ro,omitempty" tf:"ro,omitempty"`

	// A boolean for marking the volume as available on all nodes. Default is false.
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// Size of the underlying volume. Must end in T, G, M, or K (e.g. "1T", "1G", "1024M"
	// , "1048576K"). Note that this is a read only value.
	Size *string `json:"size,omitempty" tf:"size,omitempty"`

	// A string containing
	// the volume
	// , directory,
	// or device to be mounted into the
	// container (at the path specified by mp). E.g. local-lvm, local-zfs, local etc.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	Volume *string `json:"volume,omitempty" tf:"volume,omitempty"`
}

type RootfsParameters struct {

	// A boolean for enabling ACL support. Default is false.
	// +kubebuilder:validation:Optional
	ACL *bool `json:"acl,omitempty" tf:"acl,omitempty"`

	// A boolean for enabling user quotas inside the container for this mount point. Default is false.
	// +kubebuilder:validation:Optional
	Quota *bool `json:"quota,omitempty" tf:"quota,omitempty"`

	// A boolean for including this volume in a storage replica job. Default is false.
	// +kubebuilder:validation:Optional
	Replicate *bool `json:"replicate,omitempty" tf:"replicate,omitempty"`

	// +kubebuilder:validation:Optional
	Ro *bool `json:"ro,omitempty" tf:"ro,omitempty"`

	// A boolean for marking the volume as available on all nodes. Default is false.
	// +kubebuilder:validation:Optional
	Shared *bool `json:"shared,omitempty" tf:"shared,omitempty"`

	// Size of the underlying volume. Must end in T, G, M, or K (e.g. "1T", "1G", "1024M"
	// , "1048576K"). Note that this is a read only value.
	// +kubebuilder:validation:Required
	Size *string `json:"size" tf:"size,omitempty"`

	// A string containing
	// the volume
	// , directory,
	// or device to be mounted into the
	// container (at the path specified by mp). E.g. local-lvm, local-zfs, local etc.
	// +kubebuilder:validation:Required
	Storage *string `json:"storage" tf:"storage,omitempty"`
}

// LxcSpec defines the desired state of Lxc
type LxcSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LxcParameters `json:"forProvider"`
}

// LxcStatus defines the observed state of Lxc.
type LxcStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LxcObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Lxc is the Schema for the Lxcs API. Provides a resource to manage LXC
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmoxve}
type Lxc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.targetNode)",message="targetNode is a required parameter"
	Spec   LxcSpec   `json:"spec"`
	Status LxcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LxcList contains a list of Lxcs
type LxcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lxc `json:"items"`
}

// Repository type metadata.
var (
	Lxc_Kind             = "Lxc"
	Lxc_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Lxc_Kind}.String()
	Lxc_KindAPIVersion   = Lxc_Kind + "." + CRDGroupVersion.String()
	Lxc_GroupVersionKind = CRDGroupVersion.WithKind(Lxc_Kind)
)

func init() {
	SchemeBuilder.Register(&Lxc{}, &LxcList{})
}
