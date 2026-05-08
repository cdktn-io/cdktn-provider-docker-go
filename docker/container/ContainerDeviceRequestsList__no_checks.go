// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package container

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ContainerDeviceRequestsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceRequestsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ContainerDeviceRequestsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceRequestsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceRequestsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceRequestsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ContainerDeviceRequestsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewContainerDeviceRequestsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

