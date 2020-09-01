package cloudgateway

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	crdv1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
	clientset "k8s.io/kubernetes/pkg/client/clientset/versioned"
	informers "k8s.io/kubernetes/pkg/client/informers/externalversions/cloudgateway/v1"
	listers "k8s.io/kubernetes/pkg/client/listers/cloudgateway/v1"
)

type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset 	kubernetes.Interface
	clientset 		clientset.Interface
	siteInformer	listers.ESiteLister
	informerSynced  cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of preforming it as soon as a change happens.This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same itme
	// simultaneously in two different workers.
	workqueue 		workqueue.RateLimitingInterface

	// controller handler
	handler 		Handler
}

// NewController returns a new Controller
func NewController(
	kubecllientset kubernetes.Interface,
	clientset clientset.Interface,
	siteInformer informers.ESiteInformer,
	handler Handler) *Controller{
	controller := &Controller{
		kubeclientset: kubecllientset,
		clientset: clientset,
		siteInformer: siteInformer.Lister(),
		informerSynced: siteInformer.Informer().HasSynced,
		workqueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ESite"),
		handler: handler,
	}

	siteInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueue,
		UpdateFunc: func(old, new interface{}){
			oldv := old.(*crdv1.ESite)
			newv := new.(*crdv1.ESite)
			if oldv.ResourceVersion == newv.ResourceVersion{
				return
			}

			controller.enqueue(new)
		},
		DeleteFunc: controller.enqueueForDelete,
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error{
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	klog.Info("Starting cloudgateway control loop")
	if ok := cache.WaitForCacheSync(stopCh, c.informerSynced); !ok{
		return fmt.Errorf("failed to wait for cloudgateway caches to sync")
	}

	klog.Info("Starting cloudgateway workers")
	<-stopCh
	klog.Info("Shutting down cloudgateway workers")
	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue
func (c *Controller) runWorker(){
	klog.Info("CloudGateway controller,runWorker: starting")

	// invoke processNextItem to fetch and consume the next change
	// to a watched or listed resource

}

func (c *Controller) processNextItem() bool{
	klog.Info("CloudGateway controller.processNextWorkItem: start")

	// fetch the next item from the workqueue to process or
	// if a shutdown iss requested then return out of this to stop
	// processing
	obj, shutdown := c.workqueue.Get()
	if shutdown{
		return false
	}

	err := func(obj interface{}) error{
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		if key, ok = obj.(string); !ok{
			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}

		if err := c.syncHandler(key); err != nil{
			return fmt.Errorf("error syncing '%s': %s", key, err.Error())
		}

		c.workqueue.Forget(obj)
		klog.Infof("Successfully cloudgateway synced '%s'", key)
		return nil
	}(obj)

	if err != nil{
		runtime.HandleError(err)
		return true
	}

	return true
}

func (c *Controller) syncHandler(key string) error{
	// convert the tenant/namespace/name string into a distinct namespace and name
	tenant, namespace, name, err := cache.SplitMetaTenantNamespaceKey(key)
	if err != nil{
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	site, err := c.siteInformer.ESitesWithMultiTenancy(namespace, tenant).Get(name)
	if errors.IsNotFound(err){
		klog.V(4).Infof("%v has been deleted", key)
		c.handler.ObjectDeleted(site)
		return nil
	} else if err != nil{
		runtime.HandleError(fmt.Errorf("failed to list site by: %s/%s/%s", tenant, namespace, name))
		return err
	}

	// Add or update cases
	klog.V(4).Infof("%v has been added/updated", key)
	c.handler.ObjectCreated(site)
	return nil
}


func (c *Controller) enqueue(obj interface{}){
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil{
		runtime.HandleError(err)
		return
	}

	c.workqueue.AddRateLimited(key)
	klog.Infof("Try to enqueue key: %#v ...", key)
}

func (c *Controller) enqueueForDelete(obj interface{}){
	var key string
	var err error
	if key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj); err != nil{
		runtime.HandleError(err)
		return
	}

	c.workqueue.AddRateLimited(key)
	klog.Infof("Try to enqueueForDelete key: %#v ...", key)
}
