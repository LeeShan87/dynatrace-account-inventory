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

// checks if the CountryListWithRegions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryListWithRegions{}

// CountryListWithRegions A list of countries with their regions.
type CountryListWithRegions struct {
	// The list of countries.
	Countries []CountryRegions `json:"countries,omitempty"`
	// The number of countries.
	CountryCount *int32 `json:"countryCount,omitempty"`
}

// NewCountryListWithRegions instantiates a new CountryListWithRegions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryListWithRegions() *CountryListWithRegions {
	this := CountryListWithRegions{}
	return &this
}

// NewCountryListWithRegionsWithDefaults instantiates a new CountryListWithRegions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryListWithRegionsWithDefaults() *CountryListWithRegions {
	this := CountryListWithRegions{}
	return &this
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *CountryListWithRegions) GetCountries() []CountryRegions {
	if o == nil || IsNil(o.Countries) {
		var ret []CountryRegions
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryListWithRegions) GetCountriesOk() ([]CountryRegions, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *CountryListWithRegions) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []CountryRegions and assigns it to the Countries field.
func (o *CountryListWithRegions) SetCountries(v []CountryRegions) {
	o.Countries = v
}

// GetCountryCount returns the CountryCount field value if set, zero value otherwise.
func (o *CountryListWithRegions) GetCountryCount() int32 {
	if o == nil || IsNil(o.CountryCount) {
		var ret int32
		return ret
	}
	return *o.CountryCount
}

// GetCountryCountOk returns a tuple with the CountryCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryListWithRegions) GetCountryCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CountryCount) {
		return nil, false
	}
	return o.CountryCount, true
}

// HasCountryCount returns a boolean if a field has been set.
func (o *CountryListWithRegions) HasCountryCount() bool {
	if o != nil && !IsNil(o.CountryCount) {
		return true
	}

	return false
}

// SetCountryCount gets a reference to the given int32 and assigns it to the CountryCount field.
func (o *CountryListWithRegions) SetCountryCount(v int32) {
	o.CountryCount = &v
}

func (o CountryListWithRegions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryListWithRegions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	if !IsNil(o.CountryCount) {
		toSerialize["countryCount"] = o.CountryCount
	}
	return toSerialize, nil
}

type NullableCountryListWithRegions struct {
	value *CountryListWithRegions
	isSet bool
}

func (v NullableCountryListWithRegions) Get() *CountryListWithRegions {
	return v.value
}

func (v *NullableCountryListWithRegions) Set(val *CountryListWithRegions) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryListWithRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryListWithRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryListWithRegions(val *CountryListWithRegions) *NullableCountryListWithRegions {
	return &NullableCountryListWithRegions{value: val, isSet: true}
}

func (v NullableCountryListWithRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryListWithRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


