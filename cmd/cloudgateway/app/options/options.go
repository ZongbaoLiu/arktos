package options

import (
	"io/ioutil"
	"k8s.io/klog"
	"path"
	"sigs.k8s.io/yaml"

	cliflag "k8s.io/component-base/cli/flag"
	v1 "k8s.io/kubernetes/pkg/apis/cloudgateway/v1"
	"k8s.io/kubernetes/pkg/cloudgateway/common/constants"
)

// Options runs a cloudgateway
type Options struct {
	Master     string
	Kubeconfig string
	ConfigFile string
}

// NewOptions creates a new Options object with default parameters
func NewOptions() *Options {
	o := Options{
		ConfigFile: path.Join(constants.DefaultConfigDir, constants.DefaultConfigFile),
	}

	return &o
}

// Flags returns flags for a specific CloudGateway by section name
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("cloudgateway")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Arktos API server.")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig, "Path to kubeconfig file with authorization"+
		" and master location information.")

	return fss
}

func (o *Options) Config() (*v1.CloudGatewayConfig, error) {
	cfg := v1.NewCloudGatewayConfig()
	data, err := ioutil.ReadFile(o.ConfigFile)
	if err != nil {
		klog.Errorf("Failed to read config file %s: %v", o.ConfigFile, err)
		return nil, err
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		klog.Errorf("Failed to unmarshal config file %s: %v", o.ConfigFile, err)
		return nil, err
	}
	return cfg, nil
}
