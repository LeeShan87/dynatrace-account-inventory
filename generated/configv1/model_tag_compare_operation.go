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

// checks if the TagCompareOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagCompareOperation{}

// TagCompareOperation The condition of the `TAG` type.   The condition checks if the process group of the potential service is tagged with a specific tag.
type TagCompareOperation struct {
	// If `true` ignores the tag values and only validates that the tag key is matching. Defaults to `false`.
	CompareKeyOnly *bool `json:"compareKeyOnly,omitempty"`
	// The value to compare to.   If several values are specified, the OR logic applies.
	Tags []TagInfo `json:"tags"`
}

// NewTagCompareOperation instantiates a new TagCompareOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagCompareOperation(tags []TagInfo, type_ string) *TagCompareOperation {
	this := TagCompareOperation{}
	this.Type = type_
	this.Tags = tags
	return &this
}

// NewTagCompareOperationWithDefaults instantiates a new TagCompareOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagCompareOperationWithDefaults() *TagCompareOperation {
	this := TagCompareOperation{}
	return &this
}

// GetCompareKeyOnly returns the CompareKeyOnly field value if set, zero value otherwise.
func (o *TagCompareOperation) GetCompareKeyOnly() bool {
	if o == nil || IsNil(o.CompareKeyOnly) {
		var ret bool
		return ret
	}
	return *o.CompareKeyOnly
}

// GetCompareKeyOnlyOk returns a tuple with the CompareKeyOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagCompareOperation) GetCompareKeyOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.CompareKeyOnly) {
		return nil, false
	}
	return o.CompareKeyOnly, true
}

// HasCompareKeyOnly returns a boolean if a field has been set.
func (o *TagCompareOperation) HasCompareKeyOnly() bool {
	if o != nil && !IsNil(o.CompareKeyOnly) {
		return true
	}

	return false
}

// SetCompareKeyOnly gets a reference to the given bool and assigns it to the CompareKeyOnly field.
func (o *TagCompareOperation) SetCompareKeyOnly(v bool) {
	o.CompareKeyOnly = &v
}

// GetTags returns the Tags field value
func (o *TagCompareOperation) GetTags() []TagInfo {
	if o == nil {
		var ret []TagInfo
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagCompareOperation) GetTagsOk() ([]TagInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagCompareOperation) SetTags(v []TagInfo) {
	o.Tags = v
}

func (o TagCompareOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagCompareOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompareKeyOnly) {
		toSerialize["compareKeyOnly"] = o.CompareKeyOnly
	}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableTagCompareOperation struct {
	value *TagCompareOperation
	isSet bool
}

func (v NullableTagCompareOperation) Get() *TagCompareOperation {
	return v.value
}

func (v *NullableTagCompareOperation) Set(val *TagCompareOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCompareOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCompareOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCompareOperation(val *TagCompareOperation) *NullableTagCompareOperation {
	return &NullableTagCompareOperation{value: val, isSet: true}
}

func (v NullableTagCompareOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCompareOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


