// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package container

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerDeviceReadBpsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceReadBpsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceReadBpsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadBpsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadBpsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadBpsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceReadBpsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerDeviceReadBpsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

