package app

import (
	"github.com/spf13/cobra"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/cloudgateway/app/options"
	clientset "k8s.io/kubernetes/pkg/client/clientset/versioned"
	"k8s.io/kubernetes/pkg/cloudgateway"
	utilflag "k8s.io/kubernetes/pkg/util/flag"
	informers "k8s.io/kubernetes/pkg/client/informers/externalversions"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// start parameters
var (
	onlyOneSignalHandler = make(chan struct{})
	shutdownSignals		 = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

func setupSignalHandler() (stopCh <-chan struct{}){
	close(onlyOneSignalHandler)
	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func(){
		<-c
		close(stop)
		<-c
		os.Exit(1)
	}()

	return stop
}

func NewCloudGatewayCommand() *cobra.Command {
	o := options.NewOptions()
	cmd := &cobra.Command{
		Use: "cloudgateway",
		Long: `As the proxy or gateway of the services or component in the edge site, cloudgateway provides secure
communication and access capabilities for services and components of the cloud and edge sites.`,
		RunE: func(cmd *cobra.Command, args []string) error{
			utilflag.PrintFlags(cmd.Flags())

			// Set default options
			completedOptions, err := Complete(o)
			if err != nil {
				return err
			}

			// validate options
			if errs := completedOptions.Validate(); len(errs) != 0 {
				return utilerrors.NewAggregate(errs)
			}

			return runCommand(completedOptions)
		},
	}

	fs := cmd.Flags()
	namedFlagSets := o.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return cmd
}

// runCommand runs the cloudgateway
func runCommand(options completedOptions) error{
	stopCh := setupSignalHandler()
	return Run(options, stopCh)
}

func Run(options completedOptions, stopCh <-chan struct{}) error{
	klog.V(4).Infof("Cloudgateway start to run")
	cfg, err := clientcmd.BuildConfigFromFlags(options.Master, options.Kubeconfig)
	if err != nil{
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil{
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	cloudgatewayClient, err := clientset.NewForConfig(cfg)
	if err != nil{
		klog.Fatalf("Error building cloudgateway clientset: %s", err.Error())
	}

	cloudgatewayInformerFactory := informers.NewSharedInformerFactory(cloudgatewayClient, time.Second*30)
	controller := cloudgateway.NewController(kubeClient, cloudgatewayClient,
		cloudgatewayInformerFactory.Cloudgateway().V1().ESites(), &cloudgateway.TestHandler{})
	go cloudgatewayInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil{
		klog.Fatalf("Error running controller: %s", err.Error())
	}

	return nil
}

// completeOptions is a private wrapper that enforces a call of Complete() before Run can be invoked
type completedOptions struct{
	*options.Options
}

// Complete set default Options
// Should be called after cloudgateway flags parsed
func Complete(o *options.Options) (completedOptions, error) {
	var options completedOptions

	if o.Master == "" {
		o.Master = "127.0.0.1:8080"
		klog.Infof("Set master to default value %v.", o.Master)
	}

	options.Options = o
	return options, nil
}
