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

// checks if the SubscriptionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDto{}

// SubscriptionDto struct for SubscriptionDto
type SubscriptionDto struct {
	// The UUID of the Dynatrace Platform Subscription.
	Uuid string `json:"uuid"`
	// The type of the Dynatrace Platform Subscription.
	Type string `json:"type"`
	// The name of the Dynatrace Platform Subscription.
	Name string `json:"name"`
	// The status of the Dynatrace Platform Subscription.
	Status string `json:"status"`
	// The start date of the subscription in `2021-05-01` format.
	StartTime string `json:"startTime"`
	// The end date of the subscription in `2021-05-01` format.
	EndTime string `json:"endTime"`
	Account SubscriptionAccountDto `json:"account"`
	Budget SubscriptionBudgetDto `json:"budget"`
	CurrentPeriod SubscriptionCurrentPeriodDto `json:"currentPeriod"`
	// A list of period data of the subscription.
	Periods []SubscriptionPeriodDto `json:"periods"`
	// A list of subscription capabilities.
	Capabilities []SubscriptionCapabilityDto `json:"capabilities"`
}

// NewSubscriptionDto instantiates a new SubscriptionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDto(uuid string, type_ string, name string, status string, startTime string, endTime string, account SubscriptionAccountDto, budget SubscriptionBudgetDto, currentPeriod SubscriptionCurrentPeriodDto, periods []SubscriptionPeriodDto, capabilities []SubscriptionCapabilityDto) *SubscriptionDto {
	this := SubscriptionDto{}
	this.Uuid = uuid
	this.Type = type_
	this.Name = name
	this.Status = status
	this.StartTime = startTime
	this.EndTime = endTime
	this.Account = account
	this.Budget = budget
	this.CurrentPeriod = currentPeriod
	this.Periods = periods
	this.Capabilities = capabilities
	return &this
}

// NewSubscriptionDtoWithDefaults instantiates a new SubscriptionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDtoWithDefaults() *SubscriptionDto {
	this := SubscriptionDto{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *SubscriptionDto) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SubscriptionDto) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *SubscriptionDto) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionDto) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SubscriptionDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubscriptionDto) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *SubscriptionDto) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubscriptionDto) SetStatus(v string) {
	o.Status = v
}

// GetStartTime returns the StartTime field value
func (o *SubscriptionDto) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *SubscriptionDto) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *SubscriptionDto) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *SubscriptionDto) SetEndTime(v string) {
	o.EndTime = v
}

// GetAccount returns the Account field value
func (o *SubscriptionDto) GetAccount() SubscriptionAccountDto {
	if o == nil {
		var ret SubscriptionAccountDto
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetAccountOk() (*SubscriptionAccountDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *SubscriptionDto) SetAccount(v SubscriptionAccountDto) {
	o.Account = v
}

// GetBudget returns the Budget field value
func (o *SubscriptionDto) GetBudget() SubscriptionBudgetDto {
	if o == nil {
		var ret SubscriptionBudgetDto
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetBudgetOk() (*SubscriptionBudgetDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SubscriptionDto) SetBudget(v SubscriptionBudgetDto) {
	o.Budget = v
}

// GetCurrentPeriod returns the CurrentPeriod field value
func (o *SubscriptionDto) GetCurrentPeriod() SubscriptionCurrentPeriodDto {
	if o == nil {
		var ret SubscriptionCurrentPeriodDto
		return ret
	}

	return o.CurrentPeriod
}

// GetCurrentPeriodOk returns a tuple with the CurrentPeriod field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetCurrentPeriodOk() (*SubscriptionCurrentPeriodDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriod, true
}

// SetCurrentPeriod sets field value
func (o *SubscriptionDto) SetCurrentPeriod(v SubscriptionCurrentPeriodDto) {
	o.CurrentPeriod = v
}

// GetPeriods returns the Periods field value
func (o *SubscriptionDto) GetPeriods() []SubscriptionPeriodDto {
	if o == nil {
		var ret []SubscriptionPeriodDto
		return ret
	}

	return o.Periods
}

// GetPeriodsOk returns a tuple with the Periods field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetPeriodsOk() ([]SubscriptionPeriodDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Periods, true
}

// SetPeriods sets field value
func (o *SubscriptionDto) SetPeriods(v []SubscriptionPeriodDto) {
	o.Periods = v
}

// GetCapabilities returns the Capabilities field value
func (o *SubscriptionDto) GetCapabilities() []SubscriptionCapabilityDto {
	if o == nil {
		var ret []SubscriptionCapabilityDto
		return ret
	}

	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDto) GetCapabilitiesOk() ([]SubscriptionCapabilityDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities sets field value
func (o *SubscriptionDto) SetCapabilities(v []SubscriptionCapabilityDto) {
	o.Capabilities = v
}

func (o SubscriptionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["startTime"] = o.StartTime
	toSerialize["endTime"] = o.EndTime
	toSerialize["account"] = o.Account
	toSerialize["budget"] = o.Budget
	toSerialize["currentPeriod"] = o.CurrentPeriod
	toSerialize["periods"] = o.Periods
	toSerialize["capabilities"] = o.Capabilities
	return toSerialize, nil
}

type NullableSubscriptionDto struct {
	value *SubscriptionDto
	isSet bool
}

func (v NullableSubscriptionDto) Get() *SubscriptionDto {
	return v.value
}

func (v *NullableSubscriptionDto) Set(val *SubscriptionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDto(val *SubscriptionDto) *NullableSubscriptionDto {
	return &NullableSubscriptionDto{value: val, isSet: true}
}

func (v NullableSubscriptionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

