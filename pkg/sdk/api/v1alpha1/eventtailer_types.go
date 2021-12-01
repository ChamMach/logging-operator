// Copyright (c) 2020 Banzai Cloud Zrt. All Rights Reserved.

package v1alpha1

import (
	"github.com/banzaicloud/operator-tools/pkg/types"
	"github.com/banzaicloud/operator-tools/pkg/volume"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +name:"EventTailer"
// +weight:"200"
type _hugoEventTailer = interface{}

// +name:"EventTailer"
// +version:"v1alpha1"
// +description:"Eventtailer's main goal is to listen kubernetes events and transmit their changes to stdout. This way the logging-operator is able to process them."
type _metaEventTailer = interface{}

// EventTailerSpec defines the desired state of EventTailer
type EventTailerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//+kubebuilder:validation:Required
	// The resources of EventTailer will be placed into this namespace
	ControlNamespace string `json:"controlNamespace"`
	// Volume definition for tracking fluentbit file positions (optional)
	PositionVolume volume.KubernetesVolume `json:"positionVolume,omitempty"`
	// Override metadata of the created resources
	WorkloadMetaBase *types.MetaBase `json:"workloadMetaOverrides,omitempty"`
	// Override podSpec fields for the given statefulset
	WorkloadBase *types.PodSpecBase `json:"workloadOverrides,omitempty"`
	// Override container fields for the given statefulset
	ContainerBase *types.ContainerBase `json:"containerOverrides,omitempty"`
}

// EventTailerStatus defines the observed state of EventTailer
type EventTailerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=eventtailers,scope=Cluster

// EventTailer is the Schema for the eventtailers API
type EventTailer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventTailerSpec   `json:"spec,omitempty"`
	Status EventTailerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventTailerList contains a list of EventTailer
type EventTailerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventTailer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EventTailer{}, &EventTailerList{})
}
