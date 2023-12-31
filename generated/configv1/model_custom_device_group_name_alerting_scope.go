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

// checks if the CustomDeviceGroupNameAlertingScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDeviceGroupNameAlertingScope{}

// CustomDeviceGroupNameAlertingScope A scope filter for the related custom device group name.
type CustomDeviceGroupNameAlertingScope struct {
	NameFilter MetricEventTextFilterMetricEventTextFilterOperatorDto `json:"nameFilter"`
}

// NewCustomDeviceGroupNameAlertingScope instantiates a new CustomDeviceGroupNameAlertingScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDeviceGroupNameAlertingScope(nameFilter MetricEventTextFilterMetricEventTextFilterOperatorDto, filterType string) *CustomDeviceGroupNameAlertingScope {
	this := CustomDeviceGroupNameAlertingScope{}
	this.FilterType = filterType
	this.NameFilter = nameFilter
	return &this
}

// NewCustomDeviceGroupNameAlertingScopeWithDefaults instantiates a new CustomDeviceGroupNameAlertingScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDeviceGroupNameAlertingScopeWithDefaults() *CustomDeviceGroupNameAlertingScope {
	this := CustomDeviceGroupNameAlertingScope{}
	return &this
}

// GetNameFilter returns the NameFilter field value
func (o *CustomDeviceGroupNameAlertingScope) GetNameFilter() MetricEventTextFilterMetricEventTextFilterOperatorDto {
	if o == nil {
		var ret MetricEventTextFilterMetricEventTextFilterOperatorDto
		return ret
	}

	return o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value
// and a boolean to check if the value has been set.
func (o *CustomDeviceGroupNameAlertingScope) GetNameFilterOk() (*MetricEventTextFilterMetricEventTextFilterOperatorDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameFilter, true
}

// SetNameFilter sets field value
func (o *CustomDeviceGroupNameAlertingScope) SetNameFilter(v MetricEventTextFilterMetricEventTextFilterOperatorDto) {
	o.NameFilter = v
}

func (o CustomDeviceGroupNameAlertingScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDeviceGroupNameAlertingScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nameFilter"] = o.NameFilter
	return toSerialize, nil
}

type NullableCustomDeviceGroupNameAlertingScope struct {
	value *CustomDeviceGroupNameAlertingScope
	isSet bool
}

func (v NullableCustomDeviceGroupNameAlertingScope) Get() *CustomDeviceGroupNameAlertingScope {
	return v.value
}

func (v *NullableCustomDeviceGroupNameAlertingScope) Set(val *CustomDeviceGroupNameAlertingScope) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDeviceGroupNameAlertingScope) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDeviceGroupNameAlertingScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDeviceGroupNameAlertingScope(val *CustomDeviceGroupNameAlertingScope) *NullableCustomDeviceGroupNameAlertingScope {
	return &NullableCustomDeviceGroupNameAlertingScope{value: val, isSet: true}
}

func (v NullableCustomDeviceGroupNameAlertingScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDeviceGroupNameAlertingScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


