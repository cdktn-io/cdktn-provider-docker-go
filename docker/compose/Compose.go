// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package compose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-docker-go/docker/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-docker-go/docker/v15/compose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/compose docker_compose}.
type Compose interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ConfigPaths() *[]*string
	SetConfigPaths(val *[]*string)
	ConfigPathsInput() *[]*string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnvFiles() *[]*string
	SetEnvFiles(val *[]*string)
	EnvFilesInput() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Profiles() *[]*string
	SetProfiles(val *[]*string)
	ProfilesInput() *[]*string
	ProjectDirectory() *string
	SetProjectDirectory(val *string)
	ProjectDirectoryInput() *string
	ProjectName() *string
	SetProjectName(val *string)
	ProjectNameInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoveOrphans() interface{}
	SetRemoveOrphans(val interface{})
	RemoveOrphansInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Wait() interface{}
	SetWait(val interface{})
	WaitInput() interface{}
	WaitTimeout() *string
	SetWaitTimeout(val *string)
	WaitTimeoutInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetEnvFiles()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProfiles()
	ResetProjectDirectory()
	ResetProjectName()
	ResetRemoveOrphans()
	ResetWait()
	ResetWaitTimeout()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for Compose
type jsiiProxy_Compose struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_Compose) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ConfigPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ConfigPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"configPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) EnvFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"envFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) EnvFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"envFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Profiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"profiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ProfilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"profilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ProjectDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ProjectDirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) ProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) RemoveOrphans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeOrphans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) RemoveOrphansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeOrphansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) Wait() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) WaitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) WaitTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Compose) WaitTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitTimeoutInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/compose docker_compose} Resource.
func NewCompose(scope constructs.Construct, id *string, config *ComposeConfig) Compose {
	_init_.Initialize()

	if err := validateNewComposeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Compose{}

	_jsii_.Create(
		"@cdktn/provider-docker.compose.Compose",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/kreuzwerker/docker/4.4.0/docs/resources/compose docker_compose} Resource.
func NewCompose_Override(c Compose, scope constructs.Construct, id *string, config *ComposeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-docker.compose.Compose",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Compose)SetConfigPaths(val *[]*string) {
	if err := j.validateSetConfigPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configPaths",
		val,
	)
}

func (j *jsiiProxy_Compose)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Compose)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Compose)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Compose)SetEnvFiles(val *[]*string) {
	if err := j.validateSetEnvFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envFiles",
		val,
	)
}

func (j *jsiiProxy_Compose)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Compose)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Compose)SetProfiles(val *[]*string) {
	if err := j.validateSetProfilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profiles",
		val,
	)
}

func (j *jsiiProxy_Compose)SetProjectDirectory(val *string) {
	if err := j.validateSetProjectDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectDirectory",
		val,
	)
}

func (j *jsiiProxy_Compose)SetProjectName(val *string) {
	if err := j.validateSetProjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectName",
		val,
	)
}

func (j *jsiiProxy_Compose)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Compose)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Compose)SetRemoveOrphans(val interface{}) {
	if err := j.validateSetRemoveOrphansParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeOrphans",
		val,
	)
}

func (j *jsiiProxy_Compose)SetWait(val interface{}) {
	if err := j.validateSetWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wait",
		val,
	)
}

func (j *jsiiProxy_Compose)SetWaitTimeout(val *string) {
	if err := j.validateSetWaitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitTimeout",
		val,
	)
}

// Generates CDKTN code for importing a Compose resource upon running "cdktn plan <stack-name>".
func Compose_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateCompose_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-docker.compose.Compose",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Compose_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCompose_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-docker.compose.Compose",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Compose_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCompose_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-docker.compose.Compose",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Compose_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCompose_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-docker.compose.Compose",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Compose_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-docker.compose.Compose",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Compose) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Compose) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Compose) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Compose) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Compose) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Compose) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Compose) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Compose) ResetEnvFiles() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvFiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetProfiles() {
	_jsii_.InvokeVoid(
		c,
		"resetProfiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetProjectDirectory() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectDirectory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetProjectName() {
	_jsii_.InvokeVoid(
		c,
		"resetProjectName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetRemoveOrphans() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoveOrphans",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetWait() {
	_jsii_.InvokeVoid(
		c,
		"resetWait",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) ResetWaitTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Compose) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Compose) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

