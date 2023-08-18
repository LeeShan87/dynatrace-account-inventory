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

// checks if the WarningLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarningLine{}

// WarningLine struct for WarningLine
type WarningLine struct {
	Line *int32 `json:"line,omitempty"`
	Warning *string `json:"warning,omitempty"`
}

// NewWarningLine instantiates a new WarningLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarningLine() *WarningLine {
	this := WarningLine{}
	return &this
}

// NewWarningLineWithDefaults instantiates a new WarningLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningLineWithDefaults() *WarningLine {
	this := WarningLine{}
	return &this
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *WarningLine) GetLine() int32 {
	if o == nil || IsNil(o.Line) {
		var ret int32
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningLine) GetLineOk() (*int32, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *WarningLine) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given int32 and assigns it to the Line field.
func (o *WarningLine) SetLine(v int32) {
	o.Line = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *WarningLine) GetWarning() string {
	if o == nil || IsNil(o.Warning) {
		var ret string
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarningLine) GetWarningOk() (*string, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *WarningLine) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given string and assigns it to the Warning field.
func (o *WarningLine) SetWarning(v string) {
	o.Warning = &v
}

func (o WarningLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarningLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if !IsNil(o.Warning) {
		toSerialize["warning"] = o.Warning
	}
	return toSerialize, nil
}

type NullableWarningLine struct {
	value *WarningLine
	isSet bool
}

func (v NullableWarningLine) Get() *WarningLine {
	return v.value
}

func (v *NullableWarningLine) Set(val *WarningLine) {
	v.value = val
	v.isSet = true
}

func (v NullableWarningLine) IsSet() bool {
	return v.isSet
}

func (v *NullableWarningLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarningLine(val *WarningLine) *NullableWarningLine {
	return &NullableWarningLine{value: val, isSet: true}
}

func (v NullableWarningLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarningLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


