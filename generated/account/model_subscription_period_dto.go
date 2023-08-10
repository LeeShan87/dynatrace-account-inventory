/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// checks if the SubscriptionPeriodDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPeriodDto{}

// SubscriptionPeriodDto struct for SubscriptionPeriodDto
type SubscriptionPeriodDto struct {
	// The subscription period start time in `2021-05-01T15:11:00Z` format.
	StartTime string `json:"startTime"`
	// The subscription period end time in `2021-05-01T15:11:00Z` format.
	EndTime string `json:"endTime"`
}

// NewSubscriptionPeriodDto instantiates a new SubscriptionPeriodDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPeriodDto(startTime string, endTime string) *SubscriptionPeriodDto {
	this := SubscriptionPeriodDto{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewSubscriptionPeriodDtoWithDefaults instantiates a new SubscriptionPeriodDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPeriodDtoWithDefaults() *SubscriptionPeriodDto {
	this := SubscriptionPeriodDto{}
	return &this
}

// GetStartTime returns the StartTime field value
func (o *SubscriptionPeriodDto) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPeriodDto) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *SubscriptionPeriodDto) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *SubscriptionPeriodDto) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPeriodDto) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SubscriptionPeriodDto) SetEndTime(v string) {
	o.EndTime = v
}

func (o SubscriptionPeriodDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPeriodDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	return toSerialize, nil
}

type NullableSubscriptionPeriodDto struct {
	value *SubscriptionPeriodDto
	isSet bool
}

func (v NullableSubscriptionPeriodDto) Get() *SubscriptionPeriodDto {
	return v.value
}

func (v *NullableSubscriptionPeriodDto) Set(val *SubscriptionPeriodDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPeriodDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPeriodDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPeriodDto(val *SubscriptionPeriodDto) *NullableSubscriptionPeriodDto {
	return &NullableSubscriptionPeriodDto{value: val, isSet: true}
}

func (v NullableSubscriptionPeriodDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPeriodDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


