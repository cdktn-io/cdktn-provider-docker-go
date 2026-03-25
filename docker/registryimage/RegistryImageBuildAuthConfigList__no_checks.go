// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package registryimage

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegistryImageBuildAuthConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildAuthConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildAuthConfigList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildAuthConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRegistryImageBuildAuthConfigListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

