// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
// nolint: maligned
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImagePullSecretSpec defines the desired state of ImagePullSecret
type ImagePullSecretSpec struct {
	// Target specifies what should be the name of the secret created in a
	// given namespace
	Target TargetConfig `json:"target"`

	// Registry contains the details of the secret to be created in each namespace
	Registry RegistryConfig `json:"registry"`
}

type NamespaceSelectorConfiguration struct {
	// Selectors specify the conditions, which are matched against the namespaces labels
	// to decide if this ImagePullSecret should be applied to the given namespace, if multiple
	// selectors are specified if one is matches the secret will be managed (OR)
	Selectors []metav1.LabelSelector `json:"selectors,omitempty"`
	// Namespaces specifies additional namespaces by name to generate the secret into
	Names []string `json:"names,omitempty"`
}

// TargetConfig describes the secret to be created and the selectors required to determine which namespaces should
// contain this secret
type TargetConfig struct {
	Secret TargetSecretConfig `json:"secret"`
	// Namespaces specify conditions on the namespaces that should have the TargetSecret generated
	Namespaces NamespaceSelectorConfiguration `json:"namespaces,omitempty"`
	// Pods specify the conditions, which are matched against the pods in each namespace
	// to decide if this ImagePullSecret should be applied to the given pod's namespace, if multiple
	// selectors are specified if one is matches the secret will be managed (OR)
	NamespacesWithPods []metav1.LabelSelector `json:"namespacesWithPods,omitempty"`
}

// TargetSecretConfig describes the properties of the secrets created in each selected namespace
type TargetSecretConfig struct {
	// Name specifies the name of the secret object inside all the selected namespace
	Name string `json:"name"`
	// Labels specifies additional labels to be put on the Secret object
	Labels LabelSet `json:"labels,omitempty"`
	// Annotations specifies additional annotations to be put on the Secret object
	Annotations LabelSet `json:"annotations,omitempty"`
}

type RegistryType string

const (
	RegistryECR      = RegistryType("ecr")
	RegistryPassthru = RegistryType("passthru")
)

// RegistryConfig specifies what secret to be used as the basis of the pull secets
type RegistryConfig struct {
	// Registry specifies which registry backend is used, if left empty the system will assume
	// passthru mode, in case of ECR the Credentials secret is expected to contain an ECR IAM user's
	// secrets.
	// +kubebuilder:validation:Enum=ecr;passthru
	Type RegistryType `json:"type,omitempty"`
	// Credentials specifies which secret to be used as the source for docker login credentials
	Credentials NamespacedName `json:"credentials"`
}

// ImagePullSecretStatus defines the observed state of ImagePullSecret
type ImagePullSecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

type NamespacedName struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ImagePullSecret is the Schema for the imagepullsecrets API
// +k8s:openapi-gen=true
type ImagePullSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImagePullSecretSpec   `json:"spec,omitempty"`
	Status ImagePullSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
// ImagePullSecretList contains a list of ImagePullSecret
type ImagePullSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePullSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImagePullSecret{}, &ImagePullSecretList{})
}
