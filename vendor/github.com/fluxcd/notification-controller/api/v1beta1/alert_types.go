/*
Copyright 2020 The Flux authors

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/fluxcd/pkg/apis/meta"
)

// AlertSpec defines an alerting rule for events involving a list of objects
type AlertSpec struct {
	// Send events using this provider
	// +required
	ProviderRef corev1.LocalObjectReference `json:"providerRef"`

	// Filter events based on severity, defaults to ('info').
	// If set to 'info' no events will be filtered.
	// +kubebuilder:validation:Enum=info;error
	// +kubebuilder:default:=info
	// +optional
	EventSeverity string `json:"eventSeverity,omitempty"`

	// Filter events based on the involved objects
	// +required
	EventSources []CrossNamespaceObjectReference `json:"eventSources"`

	// This flag tells the controller to suspend subsequent events dispatching.
	// Defaults to false.
	// +optional
	Suspend bool `json:"suspend,omitempty"`
}

// AlertStatus defines the observed state of Alert
type AlertStatus struct {
	// +optional
	Conditions []meta.Condition `json:"conditions,omitempty"`
}

// +genclient
// +genclient:Namespaced
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message",description=""
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description=""

// Alert is the Schema for the alerts API
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertSpec   `json:"spec,omitempty"`
	Status AlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alert
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}
