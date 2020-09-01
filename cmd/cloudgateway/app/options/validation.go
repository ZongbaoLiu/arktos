package options

import "fmt"

// Validate checks Options and return a slice of found errs
func (o *Options) Validate() []error {
	var errs []error
	if o.Kubeconfig == "" {
		errs = append(errs, fmt.Errorf("--kubeconfig should be empty"))
	}

	return errs
}
