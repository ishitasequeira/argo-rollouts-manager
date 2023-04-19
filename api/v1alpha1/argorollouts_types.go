/*
Copyright 2023.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ArgoRolloutsSpec defines the desired state of ArgoRollouts
type ArgoRolloutsSpec struct {

	// Env lets you specify environment for Rollouts pods
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Extra Command arguments that would append to the Rollouts
	// ExtraCommandArgs will not be added, if one of these commands is already part of the Rollouts command
	// with same or different value.
	ExtraCommandArgs []string `json:"extraCommandArgs,omitempty"`

	// Image defines Argo Rollouts controller image (optional)
	Image string `json:"image,omitempty"`

	// NodePlacement defines NodeSelectors and Taints for Rollouts workloads
	NodePlacement *ArgoRolloutsNodePlacementSpec `json:"nodePlacement,omitempty"`

	// Version defines Argo Rollouts controller tag (optional)
	Version string `json:"version,omitempty"`

	// Resources defines the Compute Resources required by the container for Rollouts.
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// ArgoRolloutsNodePlacementSpec is used to specify NodeSelector and Tolerations for Rollouts workloads
type ArgoRolloutsNodePlacementSpec struct {
	// NodeSelector is a field of PodSpec, it is a map of key value pairs used for node selection
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Tolerations allow the pods to schedule onto nodes with matching taints
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
}

// ArgoRolloutsStatus defines the observed state of ArgoRollouts
type ArgoRolloutsStatus struct {
	// Phase is a simple, high-level summary of where the ArgoRollouts is in its lifecycle.
	// There are five possible phase values:
	// Pending: The ArgoRollouts has been accepted by the Kubernetes system, but one or more of the required resources have not been created.
	// Running: All of the containers for the ArgoRollouts are still running, or in the process of starting or restarting.
	// Succeeded: All containers for the ArgoRollouts have terminated in success, and will not be restarted.
	// Failed: At least one container has terminated in failure, either exited with non-zero status or was terminated by the system.
	// Unknown: For some reason the state of the ArgoRollouts could not be obtained.
	Phase string `json:"phase"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ArgoRollouts is the Schema for the argorollouts API
type ArgoRollouts struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArgoRolloutsSpec   `json:"spec,omitempty"`
	Status ArgoRolloutsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ArgoRolloutsList contains a list of ArgoRollouts
type ArgoRolloutsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoRollouts `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArgoRollouts{}, &ArgoRolloutsList{})
}