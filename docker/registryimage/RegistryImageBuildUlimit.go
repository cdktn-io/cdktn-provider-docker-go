// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package registryimage


type RegistryImageBuildUlimit struct {
	// soft limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/registry_image#hard RegistryImage#hard}
	Hard *float64 `field:"required" json:"hard" yaml:"hard"`
	// type of ulimit, e.g. `nofile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/registry_image#name RegistryImage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// hard limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/registry_image#soft RegistryImage#soft}
	Soft *float64 `field:"required" json:"soft" yaml:"soft"`
}

