/*


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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RuntimeConfigSpec defines the desired state of RuntimeConfig
type RuntimeConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of RuntimeConfig. Edit RuntimeConfig_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// RuntimeConfigStatus defines the observed state of RuntimeConfig
type RuntimeConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RuntimeConfig is the Schema for the runtimeconfigs API
type RuntimeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RuntimeConfigSpec   `json:"spec,omitempty"`
	Status RuntimeConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuntimeConfigList contains a list of RuntimeConfig
type RuntimeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuntimeConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RuntimeConfig{}, &RuntimeConfigList{})
}
