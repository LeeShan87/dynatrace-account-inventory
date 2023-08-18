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

// checks if the ProcessGroupNameAlertingScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupNameAlertingScope{}

// ProcessGroupNameAlertingScope A scope filter for the related process group name.
type ProcessGroupNameAlertingScope struct {
	NameFilter MetricEventTextFilterMetricEventTextFilterOperatorDto `json:"nameFilter"`
}

// NewProcessGroupNameAlertingScope instantiates a new ProcessGroupNameAlertingScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupNameAlertingScope(nameFilter MetricEventTextFilterMetricEventTextFilterOperatorDto, filterType string) *ProcessGroupNameAlertingScope {
	this := ProcessGroupNameAlertingScope{}
	this.FilterType = filterType
	this.NameFilter = nameFilter
	return &this
}

// NewProcessGroupNameAlertingScopeWithDefaults instantiates a new ProcessGroupNameAlertingScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupNameAlertingScopeWithDefaults() *ProcessGroupNameAlertingScope {
	this := ProcessGroupNameAlertingScope{}
	return &this
}

// GetNameFilter returns the NameFilter field value
func (o *ProcessGroupNameAlertingScope) GetNameFilter() MetricEventTextFilterMetricEventTextFilterOperatorDto {
	if o == nil {
		var ret MetricEventTextFilterMetricEventTextFilterOperatorDto
		return ret
	}

	return o.NameFilter
}

// GetNameFilterOk returns a tuple with the NameFilter field value
// and a boolean to check if the value has been set.
func (o *ProcessGroupNameAlertingScope) GetNameFilterOk() (*MetricEventTextFilterMetricEventTextFilterOperatorDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameFilter, true
}

// SetNameFilter sets field value
func (o *ProcessGroupNameAlertingScope) SetNameFilter(v MetricEventTextFilterMetricEventTextFilterOperatorDto) {
	o.NameFilter = v
}

func (o ProcessGroupNameAlertingScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupNameAlertingScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nameFilter"] = o.NameFilter
	return toSerialize, nil
}

type NullableProcessGroupNameAlertingScope struct {
	value *ProcessGroupNameAlertingScope
	isSet bool
}

func (v NullableProcessGroupNameAlertingScope) Get() *ProcessGroupNameAlertingScope {
	return v.value
}

func (v *NullableProcessGroupNameAlertingScope) Set(val *ProcessGroupNameAlertingScope) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupNameAlertingScope) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupNameAlertingScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupNameAlertingScope(val *ProcessGroupNameAlertingScope) *NullableProcessGroupNameAlertingScope {
	return &NullableProcessGroupNameAlertingScope{value: val, isSet: true}
}

func (v NullableProcessGroupNameAlertingScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupNameAlertingScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


