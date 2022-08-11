/*
Copyright 2022.

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
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RestoreMariaDBSpec defines the desired state of RestoreMariaDB
type RestoreMariaDBSpec struct {
	// +kubebuilder:validation:Required
	MariaDBRef corev1.LocalObjectReference `json:"mariaDbRef" webhook:"inmutable"`
	// +kubebuilder:validation:Required
	BackupRef corev1.LocalObjectReference `json:"backupRef" webhook:"inmutable"`
	// +kubebuilder:default=5
	BackoffLimit int32 `json:"backoffLimit,omitempty"`
	// +kubebuilder:default=OnFailure
	RestartPolicy corev1.RestartPolicy `json:"restartPolicy,omitempty" webhook:"inmutable"`

	Resources *corev1.ResourceRequirements `json:"resources,omitempty" webhook:"inmutable"`
}

// RestoreMariaDBStatus defines the observed state of RestoreMariaDB
type RestoreMariaDBStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func (r *RestoreMariaDBStatus) SetCondition(condition metav1.Condition) {
	if r.Conditions == nil {
		r.Conditions = make([]metav1.Condition, 0)
	}
	meta.SetStatusCondition(&r.Conditions, condition)
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=rmdb
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Complete",type="string",JSONPath=".status.conditions[?(@.type==\"Complete\")].status"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Complete\")].message"
// +kubebuilder:printcolumn:name="MariaDB",type="string",JSONPath=".spec.mariaDbRef.name"
// +kubebuilder:printcolumn:name="Backup",type="string",JSONPath=".spec.backupRef.name"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// RestoreMariaDB is the Schema for the restoremariadbs API
type RestoreMariaDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreMariaDBSpec   `json:"spec,omitempty"`
	Status RestoreMariaDBStatus `json:"status,omitempty"`
}

func (r *RestoreMariaDB) IsComplete() bool {
	return meta.IsStatusConditionTrue(r.Status.Conditions, ConditionTypeComplete)
}

// +kubebuilder:object:root=true

// RestoreMariaDBList contains a list of RestoreMariaDB
type RestoreMariaDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestoreMariaDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestoreMariaDB{}, &RestoreMariaDBList{})
}
