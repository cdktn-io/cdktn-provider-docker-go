// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package container

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerDeviceWriteIopsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceWriteIopsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceWriteIopsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceWriteIopsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceWriteIopsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceWriteIopsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceWriteIopsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerDeviceWriteIopsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

