// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package volume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-docker-go/docker/v14/jsii"

	"github.com/cdktn-io/cdktn-provider-docker-go/docker/v14/volume/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VolumeClusterOutputReference interface {
	cdktn.ComplexObject
	Availability() *string
	SetAvailability(val *string)
	AvailabilityInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	Id() *string
	InternalValue() *VolumeCluster
	SetInternalValue(val *VolumeCluster)
	LimitBytes() *string
	SetLimitBytes(val *string)
	LimitBytesInput() *string
	RequiredBytes() *string
	SetRequiredBytes(val *string)
	RequiredBytesInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Sharing() *string
	SetSharing(val *string)
	SharingInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TopologyPreferred() *string
	SetTopologyPreferred(val *string)
	TopologyPreferredInput() *string
	TopologyRequired() *string
	SetTopologyRequired(val *string)
	TopologyRequiredInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetAvailability()
	ResetGroup()
	ResetLimitBytes()
	ResetRequiredBytes()
	ResetScope()
	ResetSharing()
	ResetTopologyPreferred()
	ResetTopologyRequired()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VolumeClusterOutputReference
type jsiiProxy_VolumeClusterOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VolumeClusterOutputReference) Availability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) AvailabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) InternalValue() *VolumeCluster {
	var returns *VolumeCluster
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) LimitBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) LimitBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) RequiredBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) RequiredBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requiredBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) Sharing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) SharingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TopologyPreferred() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyPreferred",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TopologyPreferredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyPreferredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TopologyRequired() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TopologyRequiredInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VolumeClusterOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewVolumeClusterOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VolumeClusterOutputReference {
	_init_.Initialize()

	if err := validateNewVolumeClusterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VolumeClusterOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-docker.volume.VolumeClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVolumeClusterOutputReference_Override(v VolumeClusterOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-docker.volume.VolumeClusterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetAvailability(val *string) {
	if err := j.validateSetAvailabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availability",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetInternalValue(val *VolumeCluster) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetLimitBytes(val *string) {
	if err := j.validateSetLimitBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"limitBytes",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetRequiredBytes(val *string) {
	if err := j.validateSetRequiredBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredBytes",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetSharing(val *string) {
	if err := j.validateSetSharingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharing",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetTopologyPreferred(val *string) {
	if err := j.validateSetTopologyPreferredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyPreferred",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetTopologyRequired(val *string) {
	if err := j.validateSetTopologyRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyRequired",
		val,
	)
}

func (j *jsiiProxy_VolumeClusterOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		v,
		"resetAvailability",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		v,
		"resetGroup",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetLimitBytes() {
	_jsii_.InvokeVoid(
		v,
		"resetLimitBytes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetRequiredBytes() {
	_jsii_.InvokeVoid(
		v,
		"resetRequiredBytes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		v,
		"resetScope",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetSharing() {
	_jsii_.InvokeVoid(
		v,
		"resetSharing",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetTopologyPreferred() {
	_jsii_.InvokeVoid(
		v,
		"resetTopologyPreferred",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetTopologyRequired() {
	_jsii_.InvokeVoid(
		v,
		"resetTopologyRequired",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		v,
		"resetType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VolumeClusterOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VolumeClusterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

