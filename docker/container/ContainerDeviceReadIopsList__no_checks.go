// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package container

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerDeviceReadIopsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceReadIopsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceReadIopsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadIopsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadIopsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadIopsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadIopsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerDeviceReadIopsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

