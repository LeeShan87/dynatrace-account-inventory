/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the UserSessionErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSessionErrors{}

// UserSessionErrors The error of a user session.
type UserSessionErrors struct {
	// The name of the application, based on the configured detection rules.
	Application *string `json:"application,omitempty"`
	// The DNS domain where the error has been recorded.
	Domain *string `json:"domain,omitempty"`
	// The Dynatrace entity ID of the application.    This information is useful when calling various REST APIs, for example, as a key for time series queries.
	InternalApplicationId *string `json:"internalApplicationId,omitempty"`
	// The name of the error.
	Name *string `json:"name,omitempty"`
	// The timestamp of the error, in UTC milliseconds.
	StartTime *int64 `json:"startTime,omitempty"`
	// The type of error.
	Type *string `json:"type,omitempty"`
}

// NewUserSessionErrors instantiates a new UserSessionErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSessionErrors() *UserSessionErrors {
	this := UserSessionErrors{}
	return &this
}

// NewUserSessionErrorsWithDefaults instantiates a new UserSessionErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSessionErrorsWithDefaults() *UserSessionErrors {
	this := UserSessionErrors{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *UserSessionErrors) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionErrors) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *UserSessionErrors) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *UserSessionErrors) SetApplication(v string) {
	o.Application = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *UserSessionErrors) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionErrors) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *UserSessionErrors) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *UserSessionErrors) SetDomain(v string) {
	o.Domain = &v
}

// GetInternalApplicationId returns the InternalApplicationId field value if set, zero value otherwise.
func (o *UserSessionErrors) GetInternalApplicationId() string {
	if o == nil || IsNil(o.InternalApplicationId) {
		var ret string
		return ret
	}
	return *o.InternalApplicationId
}

// GetInternalApplicationIdOk returns a tuple with the InternalApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionErrors) GetInternalApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalApplicationId) {
		return nil, false
	}
	return o.InternalApplicationId, true
}

// HasInternalApplicationId returns a boolean if a field has been set.
func (o *UserSessionErrors) HasInternalApplicationId() bool {
	if o != nil && !IsNil(o.InternalApplicationId) {
		return true
	}

	return false
}

// SetInternalApplicationId gets a reference to the given string and assigns it to the InternalApplicationId field.
func (o *UserSessionErrors) SetInternalApplicationId(v string) {
	o.InternalApplicationId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserSessionErrors) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionErrors) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserSessionErrors) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserSessionErrors) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UserSessionErrors) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionErrors) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UserSessionErrors) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *UserSessionErrors) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserSessionErrors) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSessionErrors) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserSessionErrors) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserSessionErrors) SetType(v string) {
	o.Type = &v
}

func (o UserSessionErrors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSessionErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.InternalApplicationId) {
		toSerialize["internalApplicationId"] = o.InternalApplicationId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableUserSessionErrors struct {
	value *UserSessionErrors
	isSet bool
}

func (v NullableUserSessionErrors) Get() *UserSessionErrors {
	return v.value
}

func (v *NullableUserSessionErrors) Set(val *UserSessionErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSessionErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSessionErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSessionErrors(val *UserSessionErrors) *NullableUserSessionErrors {
	return &NullableUserSessionErrors{value: val, isSet: true}
}

func (v NullableUserSessionErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSessionErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


