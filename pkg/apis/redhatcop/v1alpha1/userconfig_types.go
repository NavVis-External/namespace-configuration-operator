package v1alpha1

import (
	"github.com/redhat-cop/operator-utils/pkg/util/apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// UserConfigSpec defines the desired state of UserConfig
// There are four selectors: "labelSelector", "annotationSelector", "identityExtraFieldSelector" and "providerName".
// labelSelector and annoationSelector are matches against the User object
// identityExtraFieldSelector and providerName are matched against any of the Identities associated with User
// Selectors are considered in AND, so if multiple are defined tthey must all be true for a User to be selected.
type UserConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// LabelSelector selects Users by label.
	// +kubebuilder:validation:Optional
	LabelSelector metav1.LabelSelector `json:"labelSelector,omitempty"`

	// AnnotationSelector selects Users by annotation.
	// +kubebuilder:validation:Optional
	AnnotationSelector metav1.LabelSelector `json:"annotationSelector,omitempty"`

	//IdentityExtraSelector allows you to specify a selector for the extra fields of the User's idenitities.
	//If one of the user identities matches the selector the User is selected
	//This condition is in OR with ProviderName
	// +kubebuilder:validation:Optional
	IdentityExtraFieldSelector metav1.LabelSelector `json:"identityExtraFieldSelector,omitempty"`

	//ProviderName allows you to specify an idenitity provider. If a user logged in with that provider it is selected.
	//This condition is in OR with IdentityExtraSelector
	// +kubebuilder:validation:Optional
	ProviderName string `json:"providerName,omitempty"`

	// +kubebuilder:validation:Optional
	Templates []apis.LockedResourceTemplate `json:"templates,omitempty"`
}

// UserConfigStatus defines the observed state of UserConfig
type UserConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	apis.EnforcingReconcileStatus `json:",inline"`
}

func (m *UserConfig) GetEnforcingReconcileStatus() apis.EnforcingReconcileStatus {
	return m.Status.EnforcingReconcileStatus
}

func (m *UserConfig) SetEnforcingReconcileStatus(reconcileStatus apis.EnforcingReconcileStatus) {
	m.Status.EnforcingReconcileStatus = reconcileStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserConfig is the Schema for the userconfigs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=userconfigs,scope=Cluster
type UserConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserConfigSpec   `json:"spec,omitempty"`
	Status UserConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserConfigList contains a list of UserConfig
type UserConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UserConfig{}, &UserConfigList{})
}