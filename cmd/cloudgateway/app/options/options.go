package options

import cliflag "k8s.io/component-base/cli/flag"

// Options runs a cloudgateway
type Options struct {
	Master      string
	Kubeconfig  string
}

// NewOptions creates a new Options object with default parameters
func NewOptions() *Options {
	o := Options{
	}

	return &o
}

// Flags returns flags for a specific CloudGateway by section name
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	fs := fss.FlagSet("cloudgateway")
	fs.StringVar(&o.Master, "master", o.Master, "The address of the Arktos API server.")
	fs.StringVar(&o.Kubeconfig, "kubeconfig", o.Kubeconfig,"Path to kubeconfig file with authorization" +
		" and master location information.")

	return fss
}

