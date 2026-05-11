// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadockerregistryimagetags

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDockerRegistryImageTagsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Docker image repository, including any tag or digest. For example, `alpine:latest`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/data-sources/registry_image_tags#name DataDockerRegistryImageTags#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If `true`, the verification of TLS certificates of the server/registry is disabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/data-sources/registry_image_tags#insecure_skip_verify DataDockerRegistryImageTags#insecure_skip_verify}
	InsecureSkipVerify interface{} `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// If `true`, only stable semantic version tags are returned.
	//
	// Prerelease tags such as `1.2.3-rc.1` are excluded as well as any other tags that do not conform to the semantic versioning specification. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/data-sources/registry_image_tags#strict_semver DataDockerRegistryImageTags#strict_semver}
	StrictSemver interface{} `field:"optional" json:"strictSemver" yaml:"strictSemver"`
}

