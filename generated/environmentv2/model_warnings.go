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

// checks if the Warnings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Warnings{}

// Warnings struct for Warnings
type Warnings struct {
	ChangedMetricKeys []WarningLine `json:"changedMetricKeys,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewWarnings instantiates a new Warnings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarnings() *Warnings {
	this := Warnings{}
	return &this
}

// NewWarningsWithDefaults instantiates a new Warnings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarningsWithDefaults() *Warnings {
	this := Warnings{}
	return &this
}

// GetChangedMetricKeys returns the ChangedMetricKeys field value if set, zero value otherwise.
func (o *Warnings) GetChangedMetricKeys() []WarningLine {
	if o == nil || IsNil(o.ChangedMetricKeys) {
		var ret []WarningLine
		return ret
	}
	return o.ChangedMetricKeys
}

// GetChangedMetricKeysOk returns a tuple with the ChangedMetricKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warnings) GetChangedMetricKeysOk() ([]WarningLine, bool) {
	if o == nil || IsNil(o.ChangedMetricKeys) {
		return nil, false
	}
	return o.ChangedMetricKeys, true
}

// HasChangedMetricKeys returns a boolean if a field has been set.
func (o *Warnings) HasChangedMetricKeys() bool {
	if o != nil && !IsNil(o.ChangedMetricKeys) {
		return true
	}

	return false
}

// SetChangedMetricKeys gets a reference to the given []WarningLine and assigns it to the ChangedMetricKeys field.
func (o *Warnings) SetChangedMetricKeys(v []WarningLine) {
	o.ChangedMetricKeys = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Warnings) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Warnings) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Warnings) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Warnings) SetMessage(v string) {
	o.Message = &v
}

func (o Warnings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Warnings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangedMetricKeys) {
		toSerialize["changedMetricKeys"] = o.ChangedMetricKeys
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableWarnings struct {
	value *Warnings
	isSet bool
}

func (v NullableWarnings) Get() *Warnings {
	return v.value
}

func (v *NullableWarnings) Set(val *Warnings) {
	v.value = val
	v.isSet = true
}

func (v NullableWarnings) IsSet() bool {
	return v.isSet
}

func (v *NullableWarnings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarnings(val *Warnings) *NullableWarnings {
	return &NullableWarnings{value: val, isSet: true}
}

func (v NullableWarnings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarnings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


