// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PodSecurityPolicySpec PodSecurityPolicySpec defines the policy enforced.
//
// swagger:model PodSecurityPolicySpec
type PodSecurityPolicySpec struct {

	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation. If unspecified, defaults to true.
	AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`

	// AllowedCSIDrivers is an allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec. An empty value indicates that any CSI driver can be used for inline ephemeral volumes. This is a beta field, and is only honored if the API server enables the CSIInlineVolume feature gate.
	AllowedCSIDrivers []*AllowedCSIDriver `json:"allowedCSIDrivers"`

	// allowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.
	AllowedCapabilities []string `json:"allowedCapabilities"`

	// allowedFlexVolumes is an allowlist of Flexvolumes.  Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the "volumes" field.
	AllowedFlexVolumes []*AllowedFlexVolume `json:"allowedFlexVolumes"`

	// allowedHostPaths is an allowlist of host paths. Empty indicates that all host paths may be used.
	AllowedHostPaths []*AllowedHostPath `json:"allowedHostPaths"`

	// AllowedProcMountTypes is an allowlist of allowed ProcMountTypes. Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
	AllowedProcMountTypes []string `json:"allowedProcMountTypes"`

	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.
	//
	// Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar", "foo.baz", etc.
	AllowedUnsafeSysctls []string `json:"allowedUnsafeSysctls"`

	// defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.
	DefaultAddCapabilities []string `json:"defaultAddCapabilities"`

	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	DefaultAllowPrivilegeEscalation bool `json:"defaultAllowPrivilegeEscalation,omitempty"`

	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none. Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.
	//
	// Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar", "foo.baz", etc.
	ForbiddenSysctls []string `json:"forbiddenSysctls"`

	// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	// Required: true
	FsGroup *FSGroupStrategyOptions `json:"fsGroup"`

	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	HostIPC bool `json:"hostIPC,omitempty"`

	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	HostNetwork bool `json:"hostNetwork,omitempty"`

	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	HostPID bool `json:"hostPID,omitempty"`

	// hostPorts determines which host port ranges are allowed to be exposed.
	HostPorts []*HostPortRange `json:"hostPorts"`

	// privileged determines if a pod can request to be run as privileged.
	Privileged bool `json:"privileged,omitempty"`

	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	ReadOnlyRootFilesystem bool `json:"readOnlyRootFilesystem,omitempty"`

	// requiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.
	RequiredDropCapabilities []string `json:"requiredDropCapabilities"`

	// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set. If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.
	RunAsGroup *RunAsGroupStrategyOptions `json:"runAsGroup,omitempty"`

	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	// Required: true
	RunAsUser *RunAsUserStrategyOptions `json:"runAsUser"`

	// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod. If this field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being enabled.
	RuntimeClass *RuntimeClassStrategyOptions `json:"runtimeClass,omitempty"`

	// seLinux is the strategy that will dictate the allowable labels that may be set.
	// Required: true
	SeLinux *SELinuxStrategyOptions `json:"seLinux"`

	// supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
	// Required: true
	SupplementalGroups *SupplementalGroupsStrategyOptions `json:"supplementalGroups"`

	// volumes is an allowlist of volume plugins. Empty indicates that no volumes may be used. To allow all volumes you may use '*'.
	Volumes []string `json:"volumes"`
}
