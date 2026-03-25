// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package registryimage

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegistryImageBuildUlimitList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildUlimitList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildUlimitList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildUlimitList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildUlimitList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildUlimitList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildUlimitList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRegistryImageBuildUlimitListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

