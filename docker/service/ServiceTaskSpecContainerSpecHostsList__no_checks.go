// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package service

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ServiceTaskSpecContainerSpecHostsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewServiceTaskSpecContainerSpecHostsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

