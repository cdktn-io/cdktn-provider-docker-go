// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package volume


type VolumeCluster struct {
	// Availability of the volume. Can be `active` (default), `pause`, or `drain`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#availability Volume#availability}
	Availability *string `field:"optional" json:"availability" yaml:"availability"`
	// Cluster Volume group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#group Volume#group}
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Minimum size of the Cluster Volume in human readable memory bytes (like 128MiB, 2GiB, etc).
	//
	// Must be in format of KiB, MiB, Gib, Tib or PiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#limit_bytes Volume#limit_bytes}
	LimitBytes *string `field:"optional" json:"limitBytes" yaml:"limitBytes"`
	// Maximum size of the Cluster Volume in human readable memory bytes (like 128MiB, 2GiB, etc).
	//
	// Must be in format of KiB, MiB, Gib, Tib or PiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#required_bytes Volume#required_bytes}
	RequiredBytes *string `field:"optional" json:"requiredBytes" yaml:"requiredBytes"`
	// The scope of the volume. Can be `single` (default) or `multi`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#scope Volume#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The sharing mode. Can be `none` (default), `readonly`, `onewriter` or `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#sharing Volume#sharing}
	Sharing *string `field:"optional" json:"sharing" yaml:"sharing"`
	// A topology that the Cluster Volume would be preferred in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#topology_preferred Volume#topology_preferred}
	TopologyPreferred *string `field:"optional" json:"topologyPreferred" yaml:"topologyPreferred"`
	// A topology that the Cluster Volume must be accessible from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#topology_required Volume#topology_required}
	TopologyRequired *string `field:"optional" json:"topologyRequired" yaml:"topologyRequired"`
	// Cluster Volume access type. Can be `mount` or `block` (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/volume#type Volume#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

