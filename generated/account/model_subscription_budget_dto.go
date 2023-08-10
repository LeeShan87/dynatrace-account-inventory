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

// checks if the SubscriptionBudgetDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionBudgetDto{}

// SubscriptionBudgetDto struct for SubscriptionBudgetDto
type SubscriptionBudgetDto struct {
	// The total budget for the subscription
	Total float32 `json:"total"`
	// The total budget used for the subscription
	Used float32 `json:"used"`
	// The currency code for the subscription
	CurrencyCode string `json:"currencyCode"`
}

// NewSubscriptionBudgetDto instantiates a new SubscriptionBudgetDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionBudgetDto(total float32, used float32, currencyCode string) *SubscriptionBudgetDto {
	this := SubscriptionBudgetDto{}
	this.Total = total
	this.Used = used
	this.CurrencyCode = currencyCode
	return &this
}

// NewSubscriptionBudgetDtoWithDefaults instantiates a new SubscriptionBudgetDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionBudgetDtoWithDefaults() *SubscriptionBudgetDto {
	this := SubscriptionBudgetDto{}
	return &this
}

// GetTotal returns the Total field value
func (o *SubscriptionBudgetDto) GetTotal() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *SubscriptionBudgetDto) GetTotalOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *SubscriptionBudgetDto) SetTotal(v float32) {
	o.Total = v
}

// GetUsed returns the Used field value
func (o *SubscriptionBudgetDto) GetUsed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *SubscriptionBudgetDto) GetUsedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *SubscriptionBudgetDto) SetUsed(v float32) {
	o.Used = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *SubscriptionBudgetDto) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionBudgetDto) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *SubscriptionBudgetDto) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

func (o SubscriptionBudgetDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionBudgetDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["used"] = o.Used
	toSerialize["currencyCode"] = o.CurrencyCode
	return toSerialize, nil
}

type NullableSubscriptionBudgetDto struct {
	value *SubscriptionBudgetDto
	isSet bool
}

func (v NullableSubscriptionBudgetDto) Get() *SubscriptionBudgetDto {
	return v.value
}

func (v *NullableSubscriptionBudgetDto) Set(val *SubscriptionBudgetDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionBudgetDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionBudgetDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionBudgetDto(val *SubscriptionBudgetDto) *NullableSubscriptionBudgetDto {
	return &NullableSubscriptionBudgetDto{value: val, isSet: true}
}

func (v NullableSubscriptionBudgetDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionBudgetDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


