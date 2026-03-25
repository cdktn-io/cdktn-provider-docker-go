// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package registryimage

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegistryImageBuildSecretsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildSecretsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RegistryImageBuildSecretsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildSecretsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildSecretsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildSecretsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RegistryImageBuildSecretsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRegistryImageBuildSecretsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

