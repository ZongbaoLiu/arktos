package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ESite describe the edge site resource definition
type ESite struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// list type
type ESiteList struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ListMeta		`json:"metadata"`

	Items []ESite		`json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EGateway describe the edge gateway definition
type EGateway struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`

	// Ip address of the gateway
	Ip string

	// Virtual presence ip address cidr of this gateway
	VirtualPresenceIPcidr string

	// ESiteName associated to the gateway
	ESiteName string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// list type
type EGatewayList struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ListMeta		`json:"metadata"`

	Items []EGateway	`json:"items"`
}
