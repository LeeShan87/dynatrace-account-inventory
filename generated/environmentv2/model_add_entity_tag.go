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

// checks if the AddEntityTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddEntityTag{}

// AddEntityTag The custom tag to be added to monitored entities.
type AddEntityTag struct {
	// The key of the custom tag to be added to monitored entities.
	Key string `json:"key"`
	// The value of the custom tag to be added to monitored entities. May be null
	Value *string `json:"value,omitempty"`
}

// NewAddEntityTag instantiates a new AddEntityTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddEntityTag(key string) *AddEntityTag {
	this := AddEntityTag{}
	this.Key = key
	return &this
}

// NewAddEntityTagWithDefaults instantiates a new AddEntityTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddEntityTagWithDefaults() *AddEntityTag {
	this := AddEntityTag{}
	return &this
}

// GetKey returns the Key field value
func (o *AddEntityTag) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AddEntityTag) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AddEntityTag) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AddEntityTag) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddEntityTag) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AddEntityTag) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AddEntityTag) SetValue(v string) {
	o.Value = &v
}

func (o AddEntityTag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddEntityTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAddEntityTag struct {
	value *AddEntityTag
	isSet bool
}

func (v NullableAddEntityTag) Get() *AddEntityTag {
	return v.value
}

func (v *NullableAddEntityTag) Set(val *AddEntityTag) {
	v.value = val
	v.isSet = true
}

func (v NullableAddEntityTag) IsSet() bool {
	return v.isSet
}

func (v *NullableAddEntityTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddEntityTag(val *AddEntityTag) *NullableAddEntityTag {
	return &NullableAddEntityTag{value: val, isSet: true}
}

func (v NullableAddEntityTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddEntityTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


