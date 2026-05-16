// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package container


type ContainerDeviceWriteIops struct {
	// The device path on the host, e.g. `/dev/sda`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#path Container#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The write IOPS limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#rate Container#rate}
	Rate *float64 `field:"required" json:"rate" yaml:"rate"`
}

