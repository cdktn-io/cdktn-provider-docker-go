// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package registryimage


type RegistryImageBuild struct {
	// Value to specify the build context.
	//
	// Currently, only a `PATH` context is supported. You can use the helper function '${path.cwd}/context-dir'. This always refers to the local working directory, even when building images on remote hosts. Please see https://docs.docker.com/build/building/context/ for more information about build contexts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#context RegistryImage#context}
	Context *string `field:"required" json:"context" yaml:"context"`
	// A list of additional build contexts.
	//
	// Only supported when using a buildx builder. Example: `["name=path", "src = https://example.org"}`. Please see https://docs.docker.com/reference/cli/docker/buildx/build/#build-context for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#additional_contexts RegistryImage#additional_contexts}
	AdditionalContexts *[]*string `field:"optional" json:"additionalContexts" yaml:"additionalContexts"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#auth_config RegistryImage#auth_config}
	AuthConfig interface{} `field:"optional" json:"authConfig" yaml:"authConfig"`
	// Pairs for build-time variables in the form of `ENDPOINT : "https://example.com"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#build_args RegistryImage#build_args}
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
	// The name of the buildx builder to use.
	//
	// If BUILDX_BUILDER environment variable is set, it will be used. If left empty, the provider tries to resolve to the default builder - which might not always work. If you are in Windows, the legacy builder is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#builder RegistryImage#builder}
	Builder *string `field:"optional" json:"builder" yaml:"builder"`
	// BuildID is an optional identifier that can be passed together with the build request.
	//
	// The same identifier can be used to gracefully cancel the build with the cancel request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#build_id RegistryImage#build_id}
	BuildId *string `field:"optional" json:"buildId" yaml:"buildId"`
	// Path to a file where the buildx log are written to.
	//
	// Only available when `builder` is set. If not set, no logs are available. The path is taken as is, so make sure to use a path that is available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#build_log_file RegistryImage#build_log_file}
	BuildLogFile *string `field:"optional" json:"buildLogFile" yaml:"buildLogFile"`
	// External cache sources (e.g., `user/app:cache`, `type=local,src=path/to/dir`). Only supported when using a buildx builder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cache_from RegistryImage#cache_from}
	CacheFrom *[]*string `field:"optional" json:"cacheFrom" yaml:"cacheFrom"`
	// Cache export destinations (e.g., `user/app:cache`, `type=local,dest=path/to/dir`). Only supported when using a buildx builder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cache_to RegistryImage#cache_to}
	CacheTo *[]*string `field:"optional" json:"cacheTo" yaml:"cacheTo"`
	// Optional parent cgroup for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cgroup_parent RegistryImage#cgroup_parent}
	CgroupParent *string `field:"optional" json:"cgroupParent" yaml:"cgroupParent"`
	// The length of a CPU period in microseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cpu_period RegistryImage#cpu_period}
	CpuPeriod *float64 `field:"optional" json:"cpuPeriod" yaml:"cpuPeriod"`
	// Microseconds of CPU time that the container can get in a CPU period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cpu_quota RegistryImage#cpu_quota}
	CpuQuota *float64 `field:"optional" json:"cpuQuota" yaml:"cpuQuota"`
	// CPUs in which to allow execution (e.g., `0-3`, `0`, `1`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cpu_set_cpus RegistryImage#cpu_set_cpus}
	CpuSetCpus *string `field:"optional" json:"cpuSetCpus" yaml:"cpuSetCpus"`
	// MEMs in which to allow execution (`0-3`, `0`, `1`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cpu_set_mems RegistryImage#cpu_set_mems}
	CpuSetMems *string `field:"optional" json:"cpuSetMems" yaml:"cpuSetMems"`
	// CPU shares (relative weight).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#cpu_shares RegistryImage#cpu_shares}
	CpuShares *float64 `field:"optional" json:"cpuShares" yaml:"cpuShares"`
	// Name of the Dockerfile. Defaults to `Dockerfile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#dockerfile RegistryImage#dockerfile}
	Dockerfile *string `field:"optional" json:"dockerfile" yaml:"dockerfile"`
	// A list of hostnames/IP mappings to add to the container’s /etc/hosts file. Specified in the form ["hostname:IP"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#extra_hosts RegistryImage#extra_hosts}
	ExtraHosts *[]*string `field:"optional" json:"extraHosts" yaml:"extraHosts"`
	// Always remove intermediate containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#force_remove RegistryImage#force_remove}
	ForceRemove interface{} `field:"optional" json:"forceRemove" yaml:"forceRemove"`
	// Isolation represents the isolation technology of a container. The supported values are.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#isolation RegistryImage#isolation}
	Isolation *string `field:"optional" json:"isolation" yaml:"isolation"`
	// Set metadata for an image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#label RegistryImage#label}
	Label *map[string]*string `field:"optional" json:"label" yaml:"label"`
	// User-defined key/value metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#labels RegistryImage#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Set memory limit for build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#memory RegistryImage#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// Total memory (memory + swap), -1 to enable unlimited swap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#memory_swap RegistryImage#memory_swap}
	MemorySwap *float64 `field:"optional" json:"memorySwap" yaml:"memorySwap"`
	// Set the networking mode for the RUN instructions during build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#network_mode RegistryImage#network_mode}
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Do not use the cache when building the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#no_cache RegistryImage#no_cache}
	NoCache interface{} `field:"optional" json:"noCache" yaml:"noCache"`
	// Set the target platform for the build. Defaults to `GOOS/GOARCH`. For more information see the [docker documentation](https://github.com/docker/buildx/blob/master/docs/reference/buildx.md#-set-the-target-platforms-for-the-build---platform).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#platform RegistryImage#platform}
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
	// Set provenance attestation for the build.
	//
	// BuildKit v0.11+ adds provenance attestations by default, which creates OCI image manifests that some registries (like AWS Lambda) don't support. Set to `false` to disable. Valid values: `false`, `true`, `min`, `max`, `mode=min`, `mode=max`, or a full provenance specification. Only available when using a buildx builder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#provenance RegistryImage#provenance}
	Provenance *string `field:"optional" json:"provenance" yaml:"provenance"`
	// Attempt to pull the image even if an older image exists locally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#pull_parent RegistryImage#pull_parent}
	PullParent interface{} `field:"optional" json:"pullParent" yaml:"pullParent"`
	// A Git repository URI or HTTP/HTTPS context URI. Will be ignored if `builder` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#remote_context RegistryImage#remote_context}
	RemoteContext *string `field:"optional" json:"remoteContext" yaml:"remoteContext"`
	// Remove intermediate containers after a successful build. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#remove RegistryImage#remove}
	Remove interface{} `field:"optional" json:"remove" yaml:"remove"`
	// Set SBOM (Software Bill of Materials) attestation for the build.
	//
	// Set to `false` to disable. Valid values: `false`, `true`, or a full SBOM specification. Only available when using a buildx builder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#sbom RegistryImage#sbom}
	Sbom *string `field:"optional" json:"sbom" yaml:"sbom"`
	// secrets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#secrets RegistryImage#secrets}
	Secrets interface{} `field:"optional" json:"secrets" yaml:"secrets"`
	// The security options.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#security_opt RegistryImage#security_opt}
	SecurityOpt *[]*string `field:"optional" json:"securityOpt" yaml:"securityOpt"`
	// Set an ID for the build session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#session_id RegistryImage#session_id}
	SessionId *string `field:"optional" json:"sessionId" yaml:"sessionId"`
	// Size of /dev/shm in bytes. The size must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#shm_size RegistryImage#shm_size}
	ShmSize *float64 `field:"optional" json:"shmSize" yaml:"shmSize"`
	// If true the new layers are squashed into a new image with a single new layer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#squash RegistryImage#squash}
	Squash interface{} `field:"optional" json:"squash" yaml:"squash"`
	// Suppress the build output and print image ID on success.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#suppress_output RegistryImage#suppress_output}
	SuppressOutput interface{} `field:"optional" json:"suppressOutput" yaml:"suppressOutput"`
	// Name and optionally a tag in the 'name:tag' format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#tag RegistryImage#tag}
	Tag *[]*string `field:"optional" json:"tag" yaml:"tag"`
	// Set the target build stage to build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#target RegistryImage#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// ulimit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#ulimit RegistryImage#ulimit}
	Ulimit interface{} `field:"optional" json:"ulimit" yaml:"ulimit"`
	// Force using the legacy Docker builder for image builds, even if buildx/buildkit would be available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#use_legacy_builder RegistryImage#use_legacy_builder}
	UseLegacyBuilder interface{} `field:"optional" json:"useLegacyBuilder" yaml:"useLegacyBuilder"`
	// Version of the underlying builder to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.3.0/docs/resources/registry_image#version RegistryImage#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

