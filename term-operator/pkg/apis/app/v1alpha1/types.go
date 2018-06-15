package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TerminatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Terminator `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Terminator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TerminatorSpec   `json:"spec"`
	Status            TerminatorStatus `json:"status,omitempty"`
}

type TerminatorSpec struct {
	Memcache bool `json:"memcache"`
	Redis    bool `json:"redis"`
}

type TerminatorStatus struct {
	MemcacheNode []string `json:"memcacheNode"`
	RedisNode    []string `json:"redisNode"`
}
