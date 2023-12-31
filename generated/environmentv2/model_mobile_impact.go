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

// checks if the MobileImpact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileImpact{}

// MobileImpact Analysis of problem impact to a mobile application.
type MobileImpact struct {
}

// NewMobileImpact instantiates a new MobileImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileImpact(estimatedAffectedUsers int64, impactType string, impactedEntity EntityStub) *MobileImpact {
	this := MobileImpact{}
	this.EstimatedAffectedUsers = estimatedAffectedUsers
	this.ImpactType = impactType
	this.ImpactedEntity = impactedEntity
	return &this
}

// NewMobileImpactWithDefaults instantiates a new MobileImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileImpactWithDefaults() *MobileImpact {
	this := MobileImpact{}
	return &this
}

func (o MobileImpact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileImpact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableMobileImpact struct {
	value *MobileImpact
	isSet bool
}

func (v NullableMobileImpact) Get() *MobileImpact {
	return v.value
}

func (v *NullableMobileImpact) Set(val *MobileImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileImpact(val *MobileImpact) *NullableMobileImpact {
	return &NullableMobileImpact{value: val, isSet: true}
}

func (v NullableMobileImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


