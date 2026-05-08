// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package config

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigLabelsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_ConfigLabelsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigLabelsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigLabelsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigLabelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigLabelsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigLabelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigLabelsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

