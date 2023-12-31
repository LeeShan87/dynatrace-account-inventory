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

// checks if the AlertingProfileStub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertingProfileStub{}

// AlertingProfileStub Short representation of the alerting profile.
type AlertingProfileStub struct {
	// The ID of the alerting profile.
	Id string `json:"id"`
	// The name of the alerting profile.
	Name *string `json:"name,omitempty"`
}

// NewAlertingProfileStub instantiates a new AlertingProfileStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingProfileStub(id string) *AlertingProfileStub {
	this := AlertingProfileStub{}
	this.Id = id
	return &this
}

// NewAlertingProfileStubWithDefaults instantiates a new AlertingProfileStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingProfileStubWithDefaults() *AlertingProfileStub {
	this := AlertingProfileStub{}
	return &this
}

// GetId returns the Id field value
func (o *AlertingProfileStub) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AlertingProfileStub) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AlertingProfileStub) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AlertingProfileStub) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertingProfileStub) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AlertingProfileStub) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AlertingProfileStub) SetName(v string) {
	o.Name = &v
}

func (o AlertingProfileStub) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertingProfileStub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAlertingProfileStub struct {
	value *AlertingProfileStub
	isSet bool
}

func (v NullableAlertingProfileStub) Get() *AlertingProfileStub {
	return v.value
}

func (v *NullableAlertingProfileStub) Set(val *AlertingProfileStub) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingProfileStub) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingProfileStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingProfileStub(val *AlertingProfileStub) *NullableAlertingProfileStub {
	return &NullableAlertingProfileStub{value: val, isSet: true}
}

func (v NullableAlertingProfileStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingProfileStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


