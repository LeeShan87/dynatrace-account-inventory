/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter A dimensional or series filter on a metric.
type Filter struct {
	// If the type is `not`, `and` or `or`, then holds the contained filters.
	Operands []Filter `json:"operands,omitempty"`
	ReferenceInvocation *Invocation `json:"referenceInvocation,omitempty"`
	// For filters that match a dimension against a valkue, such as `eq` or `ne`, holds the value to compare the dimension against.
	ReferenceString *string `json:"referenceString,omitempty"`
	// For the operands of `series` filters that match against a number, holds the number to compare against.
	ReferenceValue *float32 `json:"referenceValue,omitempty"`
	Rollup *Rollup `json:"rollup,omitempty"`
	// If the type applies to a dimension, then holds the target dimension.
	TargetDimension *string `json:"targetDimension,omitempty"`
	// If the type applies to n dimensions, then holds the target dimensions. Currently only used for the `remainder` filter.
	TargetDimensions []string `json:"targetDimensions,omitempty"`
	// Type of this filter, determines which other fields are present.Can be any of:  * `eq`, * `ne`, * `prefix`, * `in`, * `remainder`, * `suffix`, * `contains`, * `existsKey`, * `series`, * `or`, * `and`, * `not`, * `ge`, * `gt`, * `le`, * `lt`, * `otherwise`.
	Type *string `json:"type,omitempty"`
}

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter() *Filter {
	this := Filter{}
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	return &this
}

// GetOperands returns the Operands field value if set, zero value otherwise.
func (o *Filter) GetOperands() []Filter {
	if o == nil || IsNil(o.Operands) {
		var ret []Filter
		return ret
	}
	return o.Operands
}

// GetOperandsOk returns a tuple with the Operands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetOperandsOk() ([]Filter, bool) {
	if o == nil || IsNil(o.Operands) {
		return nil, false
	}
	return o.Operands, true
}

// HasOperands returns a boolean if a field has been set.
func (o *Filter) HasOperands() bool {
	if o != nil && !IsNil(o.Operands) {
		return true
	}

	return false
}

// SetOperands gets a reference to the given []Filter and assigns it to the Operands field.
func (o *Filter) SetOperands(v []Filter) {
	o.Operands = v
}

// GetReferenceInvocation returns the ReferenceInvocation field value if set, zero value otherwise.
func (o *Filter) GetReferenceInvocation() Invocation {
	if o == nil || IsNil(o.ReferenceInvocation) {
		var ret Invocation
		return ret
	}
	return *o.ReferenceInvocation
}

// GetReferenceInvocationOk returns a tuple with the ReferenceInvocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetReferenceInvocationOk() (*Invocation, bool) {
	if o == nil || IsNil(o.ReferenceInvocation) {
		return nil, false
	}
	return o.ReferenceInvocation, true
}

// HasReferenceInvocation returns a boolean if a field has been set.
func (o *Filter) HasReferenceInvocation() bool {
	if o != nil && !IsNil(o.ReferenceInvocation) {
		return true
	}

	return false
}

// SetReferenceInvocation gets a reference to the given Invocation and assigns it to the ReferenceInvocation field.
func (o *Filter) SetReferenceInvocation(v Invocation) {
	o.ReferenceInvocation = &v
}

// GetReferenceString returns the ReferenceString field value if set, zero value otherwise.
func (o *Filter) GetReferenceString() string {
	if o == nil || IsNil(o.ReferenceString) {
		var ret string
		return ret
	}
	return *o.ReferenceString
}

// GetReferenceStringOk returns a tuple with the ReferenceString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetReferenceStringOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceString) {
		return nil, false
	}
	return o.ReferenceString, true
}

// HasReferenceString returns a boolean if a field has been set.
func (o *Filter) HasReferenceString() bool {
	if o != nil && !IsNil(o.ReferenceString) {
		return true
	}

	return false
}

// SetReferenceString gets a reference to the given string and assigns it to the ReferenceString field.
func (o *Filter) SetReferenceString(v string) {
	o.ReferenceString = &v
}

// GetReferenceValue returns the ReferenceValue field value if set, zero value otherwise.
func (o *Filter) GetReferenceValue() float32 {
	if o == nil || IsNil(o.ReferenceValue) {
		var ret float32
		return ret
	}
	return *o.ReferenceValue
}

// GetReferenceValueOk returns a tuple with the ReferenceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetReferenceValueOk() (*float32, bool) {
	if o == nil || IsNil(o.ReferenceValue) {
		return nil, false
	}
	return o.ReferenceValue, true
}

// HasReferenceValue returns a boolean if a field has been set.
func (o *Filter) HasReferenceValue() bool {
	if o != nil && !IsNil(o.ReferenceValue) {
		return true
	}

	return false
}

// SetReferenceValue gets a reference to the given float32 and assigns it to the ReferenceValue field.
func (o *Filter) SetReferenceValue(v float32) {
	o.ReferenceValue = &v
}

// GetRollup returns the Rollup field value if set, zero value otherwise.
func (o *Filter) GetRollup() Rollup {
	if o == nil || IsNil(o.Rollup) {
		var ret Rollup
		return ret
	}
	return *o.Rollup
}

// GetRollupOk returns a tuple with the Rollup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetRollupOk() (*Rollup, bool) {
	if o == nil || IsNil(o.Rollup) {
		return nil, false
	}
	return o.Rollup, true
}

// HasRollup returns a boolean if a field has been set.
func (o *Filter) HasRollup() bool {
	if o != nil && !IsNil(o.Rollup) {
		return true
	}

	return false
}

// SetRollup gets a reference to the given Rollup and assigns it to the Rollup field.
func (o *Filter) SetRollup(v Rollup) {
	o.Rollup = &v
}

// GetTargetDimension returns the TargetDimension field value if set, zero value otherwise.
func (o *Filter) GetTargetDimension() string {
	if o == nil || IsNil(o.TargetDimension) {
		var ret string
		return ret
	}
	return *o.TargetDimension
}

// GetTargetDimensionOk returns a tuple with the TargetDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetTargetDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetDimension) {
		return nil, false
	}
	return o.TargetDimension, true
}

// HasTargetDimension returns a boolean if a field has been set.
func (o *Filter) HasTargetDimension() bool {
	if o != nil && !IsNil(o.TargetDimension) {
		return true
	}

	return false
}

// SetTargetDimension gets a reference to the given string and assigns it to the TargetDimension field.
func (o *Filter) SetTargetDimension(v string) {
	o.TargetDimension = &v
}

// GetTargetDimensions returns the TargetDimensions field value if set, zero value otherwise.
func (o *Filter) GetTargetDimensions() []string {
	if o == nil || IsNil(o.TargetDimensions) {
		var ret []string
		return ret
	}
	return o.TargetDimensions
}

// GetTargetDimensionsOk returns a tuple with the TargetDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetTargetDimensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetDimensions) {
		return nil, false
	}
	return o.TargetDimensions, true
}

// HasTargetDimensions returns a boolean if a field has been set.
func (o *Filter) HasTargetDimensions() bool {
	if o != nil && !IsNil(o.TargetDimensions) {
		return true
	}

	return false
}

// SetTargetDimensions gets a reference to the given []string and assigns it to the TargetDimensions field.
func (o *Filter) SetTargetDimensions(v []string) {
	o.TargetDimensions = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Filter) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filter) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Filter) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Filter) SetType(v string) {
	o.Type = &v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operands) {
		toSerialize["operands"] = o.Operands
	}
	if !IsNil(o.ReferenceInvocation) {
		toSerialize["referenceInvocation"] = o.ReferenceInvocation
	}
	if !IsNil(o.ReferenceString) {
		toSerialize["referenceString"] = o.ReferenceString
	}
	if !IsNil(o.ReferenceValue) {
		toSerialize["referenceValue"] = o.ReferenceValue
	}
	if !IsNil(o.Rollup) {
		toSerialize["rollup"] = o.Rollup
	}
	if !IsNil(o.TargetDimension) {
		toSerialize["targetDimension"] = o.TargetDimension
	}
	if !IsNil(o.TargetDimensions) {
		toSerialize["targetDimensions"] = o.TargetDimensions
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


