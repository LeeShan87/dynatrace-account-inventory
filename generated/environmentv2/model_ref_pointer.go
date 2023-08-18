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

// checks if the RefPointer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefPointer{}

// RefPointer Object with a pointer to a JSON object
type RefPointer struct {
	// Pointer to a JSON object this object should be logically replaced with.
	Ref string `json:"$ref"`
}

// NewRefPointer instantiates a new RefPointer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefPointer(ref string) *RefPointer {
	this := RefPointer{}
	this.Ref = ref
	return &this
}

// NewRefPointerWithDefaults instantiates a new RefPointer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefPointerWithDefaults() *RefPointer {
	this := RefPointer{}
	return &this
}

// GetRef returns the Ref field value
func (o *RefPointer) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *RefPointer) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *RefPointer) SetRef(v string) {
	o.Ref = v
}

func (o RefPointer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefPointer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["$ref"] = o.Ref
	return toSerialize, nil
}

type NullableRefPointer struct {
	value *RefPointer
	isSet bool
}

func (v NullableRefPointer) Get() *RefPointer {
	return v.value
}

func (v *NullableRefPointer) Set(val *RefPointer) {
	v.value = val
	v.isSet = true
}

func (v NullableRefPointer) IsSet() bool {
	return v.isSet
}

func (v *NullableRefPointer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefPointer(val *RefPointer) *NullableRefPointer {
	return &NullableRefPointer{value: val, isSet: true}
}

func (v NullableRefPointer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefPointer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


