// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package service


type ServiceTaskSpecNetworksAdvanced struct {
	// The network aliases of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/service#aliases Service#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// An array of driver options for the network, e.g. `opts1=value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/service#driver_opts Service#driver_opts}
	DriverOpts *[]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// The id of the docker network to use.
	//
	// Please use `docker_network.id`. Using the name attribute of the docker network will lead to constant replacements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/service#id Service#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Deprecated attribute. The name/id of the docker network. Conflicts with `id` attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/service#name Service#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

