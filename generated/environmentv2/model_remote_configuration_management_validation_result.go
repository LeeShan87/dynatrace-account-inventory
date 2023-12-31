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

// checks if the RemoteConfigurationManagementValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteConfigurationManagementValidationResult{}

// RemoteConfigurationManagementValidationResult The result of remote configuration management validation.
type RemoteConfigurationManagementValidationResult struct {
	// A list of validation errors for entities.
	InvalidEntities []RemoteConfigurationManagementEntityValidationError `json:"invalidEntities,omitempty"`
	// A list of validation errors for operations.
	InvalidOperations []RemoteConfigurationManagementOperationValidationError `json:"invalidOperations,omitempty"`
}

// NewRemoteConfigurationManagementValidationResult instantiates a new RemoteConfigurationManagementValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteConfigurationManagementValidationResult() *RemoteConfigurationManagementValidationResult {
	this := RemoteConfigurationManagementValidationResult{}
	return &this
}

// NewRemoteConfigurationManagementValidationResultWithDefaults instantiates a new RemoteConfigurationManagementValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteConfigurationManagementValidationResultWithDefaults() *RemoteConfigurationManagementValidationResult {
	this := RemoteConfigurationManagementValidationResult{}
	return &this
}

// GetInvalidEntities returns the InvalidEntities field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementValidationResult) GetInvalidEntities() []RemoteConfigurationManagementEntityValidationError {
	if o == nil || IsNil(o.InvalidEntities) {
		var ret []RemoteConfigurationManagementEntityValidationError
		return ret
	}
	return o.InvalidEntities
}

// GetInvalidEntitiesOk returns a tuple with the InvalidEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementValidationResult) GetInvalidEntitiesOk() ([]RemoteConfigurationManagementEntityValidationError, bool) {
	if o == nil || IsNil(o.InvalidEntities) {
		return nil, false
	}
	return o.InvalidEntities, true
}

// HasInvalidEntities returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementValidationResult) HasInvalidEntities() bool {
	if o != nil && !IsNil(o.InvalidEntities) {
		return true
	}

	return false
}

// SetInvalidEntities gets a reference to the given []RemoteConfigurationManagementEntityValidationError and assigns it to the InvalidEntities field.
func (o *RemoteConfigurationManagementValidationResult) SetInvalidEntities(v []RemoteConfigurationManagementEntityValidationError) {
	o.InvalidEntities = v
}

// GetInvalidOperations returns the InvalidOperations field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementValidationResult) GetInvalidOperations() []RemoteConfigurationManagementOperationValidationError {
	if o == nil || IsNil(o.InvalidOperations) {
		var ret []RemoteConfigurationManagementOperationValidationError
		return ret
	}
	return o.InvalidOperations
}

// GetInvalidOperationsOk returns a tuple with the InvalidOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementValidationResult) GetInvalidOperationsOk() ([]RemoteConfigurationManagementOperationValidationError, bool) {
	if o == nil || IsNil(o.InvalidOperations) {
		return nil, false
	}
	return o.InvalidOperations, true
}

// HasInvalidOperations returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementValidationResult) HasInvalidOperations() bool {
	if o != nil && !IsNil(o.InvalidOperations) {
		return true
	}

	return false
}

// SetInvalidOperations gets a reference to the given []RemoteConfigurationManagementOperationValidationError and assigns it to the InvalidOperations field.
func (o *RemoteConfigurationManagementValidationResult) SetInvalidOperations(v []RemoteConfigurationManagementOperationValidationError) {
	o.InvalidOperations = v
}

func (o RemoteConfigurationManagementValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteConfigurationManagementValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvalidEntities) {
		toSerialize["invalidEntities"] = o.InvalidEntities
	}
	if !IsNil(o.InvalidOperations) {
		toSerialize["invalidOperations"] = o.InvalidOperations
	}
	return toSerialize, nil
}

type NullableRemoteConfigurationManagementValidationResult struct {
	value *RemoteConfigurationManagementValidationResult
	isSet bool
}

func (v NullableRemoteConfigurationManagementValidationResult) Get() *RemoteConfigurationManagementValidationResult {
	return v.value
}

func (v *NullableRemoteConfigurationManagementValidationResult) Set(val *RemoteConfigurationManagementValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteConfigurationManagementValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteConfigurationManagementValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteConfigurationManagementValidationResult(val *RemoteConfigurationManagementValidationResult) *NullableRemoteConfigurationManagementValidationResult {
	return &NullableRemoteConfigurationManagementValidationResult{value: val, isSet: true}
}

func (v NullableRemoteConfigurationManagementValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteConfigurationManagementValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


