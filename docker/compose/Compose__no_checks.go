// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package compose

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Compose) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateImportFromParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (c *jsiiProxy_Compose) validateMoveToIdParameters(id *string) error {
	return nil
}

func (c *jsiiProxy_Compose) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateCompose_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateCompose_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCompose_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateCompose_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetConfigPathsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetEnvFilesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetProfilesParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetProjectDirectoryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetProjectNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetRemoveOrphansParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetWaitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Compose) validateSetWaitTimeoutParameters(val *string) error {
	return nil
}

func validateNewComposeParameters(scope constructs.Construct, id *string, config *ComposeConfig) error {
	return nil
}

