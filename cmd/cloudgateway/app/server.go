package app

import (
	"github.com/spf13/cobra"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/cloudgateway/app/options"
	utilflag "k8s.io/kubernetes/pkg/util/flag"
)

func NewCloudGatewayCommand() *cobra.Command {
	o := options.NewOptions()
	cmd := &cobra.Command{
		Use: "cloudgateway",
		Long: `TODO(nkaptx)`,
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
