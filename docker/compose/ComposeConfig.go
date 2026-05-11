// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package compose

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ComposeConfig struct {
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
	// One or more Compose file paths, equivalent to repeating the `-f` flag with `docker compose`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#config_paths Compose#config_paths}
	ConfigPaths *[]*string `field:"required" json:"configPaths" yaml:"configPaths"`
	// Optional list of env files to load before parsing the Compose configuration.
	//
	// If omitted, Compose uses the default `.env` behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#env_files Compose#env_files}
	EnvFiles *[]*string `field:"optional" json:"envFiles" yaml:"envFiles"`
	// Optional list of Compose profiles to enable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#profiles Compose#profiles}
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
	// Optional project directory used as the Compose working directory.
	//
	// If omitted, Compose uses the directory of the first file in `config_paths`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#project_directory Compose#project_directory}
	ProjectDirectory *string `field:"optional" json:"projectDirectory" yaml:"projectDirectory"`
	// Optional Compose project name. If omitted, Compose derives the project name the same way as the CLI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#project_name Compose#project_name}
	ProjectName *string `field:"optional" json:"projectName" yaml:"projectName"`
	// If `true`, remove containers for services that are no longer present in the Compose configuration during apply and destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#remove_orphans Compose#remove_orphans}
	RemoveOrphans interface{} `field:"optional" json:"removeOrphans" yaml:"removeOrphans"`
	// If `true`, wait until services reach the running or healthy state before returning from apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#wait Compose#wait}
	Wait interface{} `field:"optional" json:"wait" yaml:"wait"`
	// Optional duration for `wait`, for example `30s` or `2m`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/compose#wait_timeout Compose#wait_timeout}
	WaitTimeout *string `field:"optional" json:"waitTimeout" yaml:"waitTimeout"`
}

