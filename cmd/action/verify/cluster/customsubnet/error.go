package customsubnet

import "github.com/giantswarm/microerror"

var invalidConfigError = &microerror.Error{
	Kind: "invalidConfigError",
}

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return microerror.Cause(err) == invalidConfigError
}

var invalidFlagsError = &microerror.Error{
	Kind: "invalidFlagsError",
}

// IsInvalidFlags asserts invalidFlagsError.
func IsInvalidFlags(err error) bool {
	return microerror.Cause(err) == invalidFlagsError
}

var subnetMaskMasterUnusedError = &microerror.Error{
	Kind: "subnetMaskMasterUnusedError",
	Desc: "Custom Subnet mask is not being used on the master subnet.",
}

var subnetMaskWorkerUnusedError = &microerror.Error{
	Kind: "subnetMaskWorkerUnusedError",
	Desc: "Custom Subnet mask is not being used on the worker subnet.",
}

func IsSubnetMaskMasterUnusedError(err error) bool {
	return microerror.Cause(err) == subnetMaskMasterUnusedError
}

func IsSubnetMaskWorkerUnusedError(err error) bool {
	return microerror.Cause(err) == subnetMaskWorkerUnusedError
}
