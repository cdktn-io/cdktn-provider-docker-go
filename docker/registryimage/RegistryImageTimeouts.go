// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package registryimage


type RegistryImageTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/registry_image#create RegistryImage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/registry_image#delete RegistryImage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/registry_image#update RegistryImage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

