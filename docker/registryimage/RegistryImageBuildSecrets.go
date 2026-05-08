// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package registryimage


type RegistryImageBuildSecrets struct {
	// ID of the secret. By default, secrets are mounted to /run/secrets/<id>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/registry_image#id RegistryImage#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Environment variable source of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/registry_image#env RegistryImage#env}
	Env *string `field:"optional" json:"env" yaml:"env"`
	// File source of the secret. Takes precedence over `env`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/registry_image#src RegistryImage#src}
	Src *string `field:"optional" json:"src" yaml:"src"`
}

