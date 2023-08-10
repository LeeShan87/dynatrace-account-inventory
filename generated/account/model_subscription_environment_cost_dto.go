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

// checks if the SubscriptionEnvironmentCostDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionEnvironmentCostDto{}

// SubscriptionEnvironmentCostDto struct for SubscriptionEnvironmentCostDto
type SubscriptionEnvironmentCostDto struct {
	// The id of the environment
	EnvironmentId string `json:"environmentId"`
	// A list of subscription cost for the environment.
	Cost []SubscriptionCapabilityCostDto `json:"cost"`
}

// NewSubscriptionEnvironmentCostDto instantiates a new SubscriptionEnvironmentCostDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionEnvironmentCostDto(environmentId string, cost []SubscriptionCapabilityCostDto) *SubscriptionEnvironmentCostDto {
	this := SubscriptionEnvironmentCostDto{}
	this.EnvironmentId = environmentId
	this.Cost = cost
	return &this
}

// NewSubscriptionEnvironmentCostDtoWithDefaults instantiates a new SubscriptionEnvironmentCostDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionEnvironmentCostDtoWithDefaults() *SubscriptionEnvironmentCostDto {
	this := SubscriptionEnvironmentCostDto{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *SubscriptionEnvironmentCostDto) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *SubscriptionEnvironmentCostDto) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *SubscriptionEnvironmentCostDto) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetCost returns the Cost field value
func (o *SubscriptionEnvironmentCostDto) GetCost() []SubscriptionCapabilityCostDto {
	if o == nil {
		var ret []SubscriptionCapabilityCostDto
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *SubscriptionEnvironmentCostDto) GetCostOk() ([]SubscriptionCapabilityCostDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost, true
}

// SetCost sets field value
func (o *SubscriptionEnvironmentCostDto) SetCost(v []SubscriptionCapabilityCostDto) {
	o.Cost = v
}

func (o SubscriptionEnvironmentCostDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionEnvironmentCostDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environmentId"] = o.EnvironmentId
	toSerialize["cost"] = o.Cost
	return toSerialize, nil
}

type NullableSubscriptionEnvironmentCostDto struct {
	value *SubscriptionEnvironmentCostDto
	isSet bool
}

func (v NullableSubscriptionEnvironmentCostDto) Get() *SubscriptionEnvironmentCostDto {
	return v.value
}

func (v *NullableSubscriptionEnvironmentCostDto) Set(val *SubscriptionEnvironmentCostDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionEnvironmentCostDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionEnvironmentCostDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionEnvironmentCostDto(val *SubscriptionEnvironmentCostDto) *NullableSubscriptionEnvironmentCostDto {
	return &NullableSubscriptionEnvironmentCostDto{value: val, isSet: true}
}

func (v NullableSubscriptionEnvironmentCostDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionEnvironmentCostDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


