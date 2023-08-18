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

// checks if the TagInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagInfo{}

// TagInfo Tag of a Dynatrace entity.
type TagInfo struct {
	// The origin of the tag, such as AWS or Cloud Foundry.    Custom tags use the `CONTEXTLESS` value.
	Context string `json:"context"`
	// The key of the tag.    Custom tags have the tag value here.
	Key string `json:"key"`
	// The value of the tag.    Not applicable to custom tags.
	Value *string `json:"value,omitempty"`
}

// NewTagInfo instantiates a new TagInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagInfo(context string, key string) *TagInfo {
	this := TagInfo{}
	this.Context = context
	this.Key = key
	return &this
}

// NewTagInfoWithDefaults instantiates a new TagInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagInfoWithDefaults() *TagInfo {
	this := TagInfo{}
	return &this
}

// GetContext returns the Context field value
func (o *TagInfo) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *TagInfo) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *TagInfo) SetContext(v string) {
	o.Context = v
}

// GetKey returns the Key field value
func (o *TagInfo) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TagInfo) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *TagInfo) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TagInfo) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagInfo) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TagInfo) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TagInfo) SetValue(v string) {
	o.Value = &v
}

func (o TagInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["key"] = o.Key
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTagInfo struct {
	value *TagInfo
	isSet bool
}

func (v NullableTagInfo) Get() *TagInfo {
	return v.value
}

func (v *NullableTagInfo) Set(val *TagInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTagInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTagInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagInfo(val *TagInfo) *NullableTagInfo {
	return &NullableTagInfo{value: val, isSet: true}
}

func (v NullableTagInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


