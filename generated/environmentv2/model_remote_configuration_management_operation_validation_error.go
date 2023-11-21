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

// checks if the RemoteConfigurationManagementOperationValidationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteConfigurationManagementOperationValidationError{}

// RemoteConfigurationManagementOperationValidationError Validation error of remote configuration management operation definition.
type RemoteConfigurationManagementOperationValidationError struct {
	// The attribute which is affected by the operation.
	Attribute *string `json:"attribute,omitempty"`
	// The operation performed on given attribute.
	Operation *string `json:"operation,omitempty"`
	// The reason of validation failure.
	Reason *string `json:"reason,omitempty"`
	// The value which should be assigned to given attribute.
	Value *string `json:"value,omitempty"`
}

// NewRemoteConfigurationManagementOperationValidationError instantiates a new RemoteConfigurationManagementOperationValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteConfigurationManagementOperationValidationError() *RemoteConfigurationManagementOperationValidationError {
	this := RemoteConfigurationManagementOperationValidationError{}
	return &this
}

// NewRemoteConfigurationManagementOperationValidationErrorWithDefaults instantiates a new RemoteConfigurationManagementOperationValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteConfigurationManagementOperationValidationErrorWithDefaults() *RemoteConfigurationManagementOperationValidationError {
	this := RemoteConfigurationManagementOperationValidationError{}
	return &this
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementOperationValidationError) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementOperationValidationError) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementOperationValidationError) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *RemoteConfigurationManagementOperationValidationError) SetAttribute(v string) {
	o.Attribute = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementOperationValidationError) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementOperationValidationError) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementOperationValidationError) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *RemoteConfigurationManagementOperationValidationError) SetOperation(v string) {
	o.Operation = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementOperationValidationError) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementOperationValidationError) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementOperationValidationError) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RemoteConfigurationManagementOperationValidationError) SetReason(v string) {
	o.Reason = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementOperationValidationError) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementOperationValidationError) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementOperationValidationError) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RemoteConfigurationManagementOperationValidationError) SetValue(v string) {
	o.Value = &v
}

func (o RemoteConfigurationManagementOperationValidationError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteConfigurationManagementOperationValidationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRemoteConfigurationManagementOperationValidationError struct {
	value *RemoteConfigurationManagementOperationValidationError
	isSet bool
}

func (v NullableRemoteConfigurationManagementOperationValidationError) Get() *RemoteConfigurationManagementOperationValidationError {
	return v.value
}

func (v *NullableRemoteConfigurationManagementOperationValidationError) Set(val *RemoteConfigurationManagementOperationValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteConfigurationManagementOperationValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteConfigurationManagementOperationValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteConfigurationManagementOperationValidationError(val *RemoteConfigurationManagementOperationValidationError) *NullableRemoteConfigurationManagementOperationValidationError {
	return &NullableRemoteConfigurationManagementOperationValidationError{value: val, isSet: true}
}

func (v NullableRemoteConfigurationManagementOperationValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteConfigurationManagementOperationValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


