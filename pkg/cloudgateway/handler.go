package cloudgateway

import "k8s.io/klog"

// handler interface contains the methods that are required
type Handler interface{
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(obj interface{})
}

// TestHandler is a test sample
type TestHandler struct{}

func (t *TestHandler) Init() error{
	klog.V(4).Info("TestHandler.Init")
	return nil
}

func (t *TestHandler) ObjectCreated(obj interface{}) {
	klog.V(4).Info("TestHandler.ObjectCreated")
}

func (t *TestHandler) ObjectUpdated(obj interface{}) {
	klog.V(4).Info("TestHandler.ObjectUpdated")
}

func (t *TestHandler) ObjectDeleted(obj interface{}) {
	klog.V(4).Info("TestHandler.ObjectDeleted")
}
