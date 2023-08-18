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

// checks if the CloudTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudTag{}

// CloudTag A cloud tag.
type CloudTag struct {
	// The name of the tag.
	Name *string `json:"name,omitempty"`
	// The value of the tag.    If set to `null`, then resources with any value of the tag are monitored.
	Value *string `json:"value,omitempty"`
}

// NewCloudTag instantiates a new CloudTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudTag() *CloudTag {
	this := CloudTag{}
	return &this
}

// NewCloudTagWithDefaults instantiates a new CloudTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudTagWithDefaults() *CloudTag {
	this := CloudTag{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudTag) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTag) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudTag) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudTag) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CloudTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CloudTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CloudTag) SetValue(v string) {
	o.Value = &v
}

func (o CloudTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCloudTag struct {
	value *CloudTag
	isSet bool
}

func (v NullableCloudTag) Get() *CloudTag {
	return v.value
}

func (v *NullableCloudTag) Set(val *CloudTag) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudTag) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudTag(val *CloudTag) *NullableCloudTag {
	return &NullableCloudTag{value: val, isSet: true}
}

func (v NullableCloudTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

