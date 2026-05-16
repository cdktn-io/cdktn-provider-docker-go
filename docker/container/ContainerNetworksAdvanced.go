// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package container


type ContainerNetworksAdvanced struct {
	// The name or id of the network to use.
	//
	// You can use `name` or `id` attribute from a `docker_network` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#name Container#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network aliases of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#aliases Container#aliases}
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// An array of driver options for the network endpoint, e.g. `opts1=value`. This is the equivalent to repeating `--driver-opt` for `docker run`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#driver_opts Container#driver_opts}
	DriverOpts *[]*string `field:"optional" json:"driverOpts" yaml:"driverOpts"`
	// Gateway priority for this endpoint.
	//
	// The endpoint with the highest priority will provide the default gateway for the container. This is the equivalent to `--gw-priority` for `docker run`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#gw_priority Container#gw_priority}
	GwPriority *float64 `field:"optional" json:"gwPriority" yaml:"gwPriority"`
	// The IPV4 address of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#ipv4_address Container#ipv4_address}
	Ipv4Address *string `field:"optional" json:"ipv4Address" yaml:"ipv4Address"`
	// The IPV6 address of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#ipv6_address Container#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The link-local IPs of the container in the specific network.
	//
	// This is the equivalent to repeating `--link-local-ip` for `docker run`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#link_local_ips Container#link_local_ips}
	LinkLocalIps *[]*string `field:"optional" json:"linkLocalIps" yaml:"linkLocalIps"`
	// The MAC address of the container in the specific network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#mac_address Container#mac_address}
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
}

