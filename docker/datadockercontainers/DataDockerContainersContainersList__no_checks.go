// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datadockercontainers

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataDockerContainersContainersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataDockerContainersContainersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataDockerContainersContainersList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataDockerContainersContainersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDockerContainersContainersList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataDockerContainersContainersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataDockerContainersContainersListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

