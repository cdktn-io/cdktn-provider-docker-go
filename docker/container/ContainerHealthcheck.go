// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package container


type ContainerHealthcheck struct {
	// Time between running the check (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#interval Container#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Consecutive failures needed to report unhealthy. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#retries Container#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// Interval before the healthcheck starts (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#start_interval Container#start_interval}
	StartInterval *string `field:"optional" json:"startInterval" yaml:"startInterval"`
	// Start period for the container to initialize before counting retries towards unstable (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#start_period Container#start_period}
	StartPeriod *string `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// Command to run to check health.
	//
	// For example, to run `curl -f localhost/health` set the command to be `["CMD", "curl", "-f", "localhost/health"]`. It works in the same way, and has the same default values, as the HEALTHCHECK Dockerfile instruction set by the service's Docker image. Your Compose file can override the values set in the Dockerfile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#test Container#test}
	Test *[]*string `field:"optional" json:"test" yaml:"test"`
	// Maximum time to allow one check to run (ms|s|m|h). Defaults to `0s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/container#timeout Container#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

