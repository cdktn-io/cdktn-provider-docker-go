// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package container


type ContainerDeviceRequests struct {
	// List of device capabilities. Only used with `nvidia` driver (e.g., `gpu`, `compute`, `utility`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/container#capabilities Container#capabilities}
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Number of devices to request. Use -1 for all devices. Only used with `nvidia` driver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/container#count Container#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// List of device IDs or CDI device identifiers (e.g., `nvidia.com/gpu=all`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/container#device_ids Container#device_ids}
	DeviceIds *[]*string `field:"optional" json:"deviceIds" yaml:"deviceIds"`
	// The device driver to use. Common values: `cdi` for CDI devices, `nvidia` for NVIDIA GPU requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/container#driver Container#driver}
	Driver *string `field:"optional" json:"driver" yaml:"driver"`
	// Driver-specific options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.2.0/docs/resources/container#options Container#options}
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
}

