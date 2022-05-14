// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Volume Volume represents a named volume in a pod that may be accessed by any container in the pod.
//
// swagger:model Volume
type Volume struct {

	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *AzureDiskVolumeSource `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *AzureFileVolumeSource `json:"azureFile,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	Cephfs *CephFSVolumeSource `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Cinder *CinderVolumeSource `json:"cinder,omitempty"`

	// ConfigMap represents a configMap that should populate this volume
	ConfigMap *ConfigMapVolumeSource `json:"configMap,omitempty"`

	// CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).
	Csi *CSIVolumeSource `json:"csi,omitempty"`

	// DownwardAPI represents downward API about the pod that should populate this volume
	DownwardAPI *DownwardAPIVolumeSource `json:"downwardAPI,omitempty"`

	// EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	EmptyDir *EmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.
	//
	// Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity
	//    tracking are needed,
	// c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through
	//    a PersistentVolumeClaim (see EphemeralVolumeSource for more
	//    information on the connection between this volume type
	//    and PersistentVolumeClaim).
	//
	// Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.
	//
	// Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.
	//
	// A pod can use both types of ephemeral volumes and persistent volumes at the same time.
	Ephemeral *EphemeralVolumeSource `json:"ephemeral,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	Fc *FCVolumeSource `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *FlexVolumeSource `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running
	Flocker *FlockerVolumeSource `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	GcePersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"`

	// GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.
	GitRepo *GitRepoVolumeSource `json:"gitRepo,omitempty"`

	// Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md
	Glusterfs *GlusterfsVolumeSource `json:"glusterfs,omitempty"`

	// HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md
	Iscsi *ISCSIVolumeSource `json:"iscsi,omitempty"`

	// Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Required: true
	Name *string `json:"name"`

	// NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Nfs *NFSVolumeSource `json:"nfs,omitempty"`

	// PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	PortworxVolume *PortworxVolumeSource `json:"portworxVolume,omitempty"`

	// Items for all in one resources secrets, configmaps, and downward API
	Projected *ProjectedVolumeSource `json:"projected,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	Quobyte *QuobyteVolumeSource `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
	Rbd *RBDVolumeSource `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	ScaleIO *ScaleIOVolumeSource `json:"scaleIO,omitempty"`

	// Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
	Secret *SecretVolumeSource `json:"secret,omitempty"`

	// StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
	Storageos *StorageOSVolumeSource `json:"storageos,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	VsphereVolume *VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"`
}
