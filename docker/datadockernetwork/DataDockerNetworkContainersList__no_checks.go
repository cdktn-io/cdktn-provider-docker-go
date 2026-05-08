// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datadockernetwork

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDockerNetworkContainersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataDockerNetworkContainersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDockerNetworkContainersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDockerNetworkContainersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDockerNetworkContainersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDockerNetworkContainersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDockerNetworkContainersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

