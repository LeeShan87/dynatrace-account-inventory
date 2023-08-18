/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the WebServiceNameSpace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebServiceNameSpace{}

// WebServiceNameSpace The contribution to the service ID calculation from the detected web service name space.    You have two mutually exclusive options:  * Override the detected value with a specified static value. Specify the new value in the **valueOverride** field.  * Dynamically transform the detected value. Specify the transformation parameters in the **transformations** field.
type WebServiceNameSpace struct {
	// Transformations to be applied to the detected value.
	Transformations []TransformationBase `json:"transformations,omitempty"`
	// The value to be used instead of the detected value.
	ValueOverride *string `json:"valueOverride,omitempty"`
}

// NewWebServiceNameSpace instantiates a new WebServiceNameSpace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebServiceNameSpace() *WebServiceNameSpace {
	this := WebServiceNameSpace{}
	return &this
}

// NewWebServiceNameSpaceWithDefaults instantiates a new WebServiceNameSpace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebServiceNameSpaceWithDefaults() *WebServiceNameSpace {
	this := WebServiceNameSpace{}
	return &this
}

// GetTransformations returns the Transformations field value if set, zero value otherwise.
func (o *WebServiceNameSpace) GetTransformations() []TransformationBase {
	if o == nil || IsNil(o.Transformations) {
		var ret []TransformationBase
		return ret
	}
	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebServiceNameSpace) GetTransformationsOk() ([]TransformationBase, bool) {
	if o == nil || IsNil(o.Transformations) {
		return nil, false
	}
	return o.Transformations, true
}

// HasTransformations returns a boolean if a field has been set.
func (o *WebServiceNameSpace) HasTransformations() bool {
	if o != nil && !IsNil(o.Transformations) {
		return true
	}

	return false
}

// SetTransformations gets a reference to the given []TransformationBase and assigns it to the Transformations field.
func (o *WebServiceNameSpace) SetTransformations(v []TransformationBase) {
	o.Transformations = v
}

// GetValueOverride returns the ValueOverride field value if set, zero value otherwise.
func (o *WebServiceNameSpace) GetValueOverride() string {
	if o == nil || IsNil(o.ValueOverride) {
		var ret string
		return ret
	}
	return *o.ValueOverride
}

// GetValueOverrideOk returns a tuple with the ValueOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebServiceNameSpace) GetValueOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ValueOverride) {
		return nil, false
	}
	return o.ValueOverride, true
}

// HasValueOverride returns a boolean if a field has been set.
func (o *WebServiceNameSpace) HasValueOverride() bool {
	if o != nil && !IsNil(o.ValueOverride) {
		return true
	}

	return false
}

// SetValueOverride gets a reference to the given string and assigns it to the ValueOverride field.
func (o *WebServiceNameSpace) SetValueOverride(v string) {
	o.ValueOverride = &v
}

func (o WebServiceNameSpace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebServiceNameSpace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transformations) {
		toSerialize["transformations"] = o.Transformations
	}
	if !IsNil(o.ValueOverride) {
		toSerialize["valueOverride"] = o.ValueOverride
	}
	return toSerialize, nil
}

type NullableWebServiceNameSpace struct {
	value *WebServiceNameSpace
	isSet bool
}

func (v NullableWebServiceNameSpace) Get() *WebServiceNameSpace {
	return v.value
}

func (v *NullableWebServiceNameSpace) Set(val *WebServiceNameSpace) {
	v.value = val
	v.isSet = true
}

func (v NullableWebServiceNameSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableWebServiceNameSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebServiceNameSpace(val *WebServiceNameSpace) *NullableWebServiceNameSpace {
	return &NullableWebServiceNameSpace{value: val, isSet: true}
}

func (v NullableWebServiceNameSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebServiceNameSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


