// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package config


type ConfigLabels struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/config#label Config#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/config#value Config#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

