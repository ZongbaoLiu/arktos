package cloudgateway

import (
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/kubernetes/pkg/client/clientset/versioned"
)

type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset 	kubernetes.Interface
	clientset 		clientset.Interface
}
