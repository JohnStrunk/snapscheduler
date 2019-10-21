package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

// ReconcileState is an enumeration describing the Object's reconcile state
type ReconcileState string

const (
	// StateIdle indicates the controller is idle
	StateIdle ReconcileState = "Idle"
	// StateSnapshotting indicates the controller is creating snapshots
	StateSnapshotting ReconcileState = "Snapshotting"
	// StateUnknown indicates the controller is in an unknown state
	StateUnknown ReconcileState = ""
)

// SnapshotRetentionSpec defines how long snapshots should be kept.
// +k8s:openapi-gen=true
type SnapshotRetentionSpec struct {
	// Expires is the length of time (time.Duration) after which a given
	// Snapshot will be deleted.
	// +kubebuilder:validation:Pattern=^\d+(h|m|s)$
	Expires string `json:"expires,omitempty"`
	// +kubebuilder:validation:Minimum=1
	MaxCount *int32 `json:"maxCount,omitempty"`
}

// SnapshotTemplateSpec defines the template for Snapshot objects
// +k8s:openapi-gen=true
type SnapshotTemplateSpec struct {
	// Labels is a list of labels that should be added to each Snapshot
	// created by this schedule.
	Labels map[string]string `json:"labels,omitempty"`
	// SnapshotClassName is the name of the VSC to be used when creating
	// Snapshots.
	SnapshotClassName *string `json:"snapshotClassName,omitempty"`
}

// SnapshotScheduleSpec defines the desired state of SnapshotSchedule
// +k8s:openapi-gen=true
type SnapshotScheduleSpec struct {
	// ClaimSelector selects which PVCs will be snapshotted according to
	// this schedule.
	ClaimSelector metav1.LabelSelector `json:"claimSelector,omitempty"`
	// Retention determines how long this schedule's snapshots will be kept.
	Retention SnapshotRetentionSpec `json:"retention,omitempty"`
	// Schedule is a Cronspec specifying when snapshots should be taken. See
	// https://en.wikipedia.org/wiki/Cron for a description of the format.
	// +kubebuilder:validation:Pattern=^((\d+|\*)(/\d+)?(\s+(\d+|\*)(/\d+)?){4}|@\w+)$
	Schedule string `json:"schedule,omitempty"`
	// Disabled determines whether this schedule is currently disabled.
	Disabled bool `json:"disabled,omitempty"`
	// SnapshotTemplate is a template description of the Snapshots to be created.
	SnapshotTemplate *SnapshotTemplateSpec `json:"snapshotTemplate,omitempty"`
	// VolumeClaimTemplate is a template description of PVC clones to be created
	// VolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
}

// SnapshotScheduleStatus defines the observed state of SnapshotSchedule
// +k8s:openapi-gen=true
type SnapshotScheduleStatus struct {
	// State is the current reconciliation state of this resource
	State ReconcileState `json:"state,omitempty"`
	// ReconcileResult describes the result of the last reconcile
	ReconcileResult string `json:"reconcileResult,omitempty"`
	// LastSnapshotTime is the time of the most recent set of snapshots
	// generated by this schedule.
	LastSnapshotTime *metav1.Time `json:"lastSnapshotTime,omitempty"`
	// NextSnapshotTime is the time when this schedule should create the
	// next set of snapshots.
	NextSnapshotTime *metav1.Time `json:"nextSnapshotTime,omitempty"`
}

// SnapshotSchedule is the Schema for the snapshotschedules API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type SnapshotSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnapshotScheduleSpec   `json:"spec,omitempty"`
	Status SnapshotScheduleStatus `json:"status,omitempty"`
}

// SnapshotScheduleList contains a list of SnapshotSchedule
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type SnapshotScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotSchedule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SnapshotSchedule{}, &SnapshotScheduleList{})
}
