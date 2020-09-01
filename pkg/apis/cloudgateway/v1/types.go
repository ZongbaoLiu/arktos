package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ESite describe the edge site resource definition
type ESite struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// list type
type ESiteList struct {
	metav1.TypeMeta
	metav1.ListMeta

	Items []ESite
}