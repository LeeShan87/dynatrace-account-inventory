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

// checks if the RemoteConfigurationManagementJobPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteConfigurationManagementJobPreview{}

// RemoteConfigurationManagementJobPreview A preview of remote configuration management job.
type RemoteConfigurationManagementJobPreview struct {
	// The number of entities that are currently configured as defined by remote configuration management operation.
	AlreadyConfiguredEntitiesCount *int32 `json:"alreadyConfiguredEntitiesCount,omitempty"`
	// The attribute which is affected by the operation.
	Attribute *string `json:"attribute,omitempty"`
	// The operation performed on given attribute.
	Operation *string `json:"operation,omitempty"`
	// The number of entities that will be configured as defined by remote configuration management after it is completed.
	TargetEntitiesCount *int32 `json:"targetEntitiesCount,omitempty"`
	// The value which should be assigned to given attribute.
	Value *string `json:"value,omitempty"`
}

// NewRemoteConfigurationManagementJobPreview instantiates a new RemoteConfigurationManagementJobPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteConfigurationManagementJobPreview() *RemoteConfigurationManagementJobPreview {
	this := RemoteConfigurationManagementJobPreview{}
	return &this
}

// NewRemoteConfigurationManagementJobPreviewWithDefaults instantiates a new RemoteConfigurationManagementJobPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteConfigurationManagementJobPreviewWithDefaults() *RemoteConfigurationManagementJobPreview {
	this := RemoteConfigurationManagementJobPreview{}
	return &this
}

// GetAlreadyConfiguredEntitiesCount returns the AlreadyConfiguredEntitiesCount field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementJobPreview) GetAlreadyConfiguredEntitiesCount() int32 {
	if o == nil || IsNil(o.AlreadyConfiguredEntitiesCount) {
		var ret int32
		return ret
	}
	return *o.AlreadyConfiguredEntitiesCount
}

// GetAlreadyConfiguredEntitiesCountOk returns a tuple with the AlreadyConfiguredEntitiesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementJobPreview) GetAlreadyConfiguredEntitiesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AlreadyConfiguredEntitiesCount) {
		return nil, false
	}
	return o.AlreadyConfiguredEntitiesCount, true
}

// HasAlreadyConfiguredEntitiesCount returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementJobPreview) HasAlreadyConfiguredEntitiesCount() bool {
	if o != nil && !IsNil(o.AlreadyConfiguredEntitiesCount) {
		return true
	}

	return false
}

// SetAlreadyConfiguredEntitiesCount gets a reference to the given int32 and assigns it to the AlreadyConfiguredEntitiesCount field.
func (o *RemoteConfigurationManagementJobPreview) SetAlreadyConfiguredEntitiesCount(v int32) {
	o.AlreadyConfiguredEntitiesCount = &v
}

// GetAttribute returns the Attribute field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementJobPreview) GetAttribute() string {
	if o == nil || IsNil(o.Attribute) {
		var ret string
		return ret
	}
	return *o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementJobPreview) GetAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.Attribute) {
		return nil, false
	}
	return o.Attribute, true
}

// HasAttribute returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementJobPreview) HasAttribute() bool {
	if o != nil && !IsNil(o.Attribute) {
		return true
	}

	return false
}

// SetAttribute gets a reference to the given string and assigns it to the Attribute field.
func (o *RemoteConfigurationManagementJobPreview) SetAttribute(v string) {
	o.Attribute = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementJobPreview) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementJobPreview) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementJobPreview) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *RemoteConfigurationManagementJobPreview) SetOperation(v string) {
	o.Operation = &v
}

// GetTargetEntitiesCount returns the TargetEntitiesCount field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementJobPreview) GetTargetEntitiesCount() int32 {
	if o == nil || IsNil(o.TargetEntitiesCount) {
		var ret int32
		return ret
	}
	return *o.TargetEntitiesCount
}

// GetTargetEntitiesCountOk returns a tuple with the TargetEntitiesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementJobPreview) GetTargetEntitiesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TargetEntitiesCount) {
		return nil, false
	}
	return o.TargetEntitiesCount, true
}

// HasTargetEntitiesCount returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementJobPreview) HasTargetEntitiesCount() bool {
	if o != nil && !IsNil(o.TargetEntitiesCount) {
		return true
	}

	return false
}

// SetTargetEntitiesCount gets a reference to the given int32 and assigns it to the TargetEntitiesCount field.
func (o *RemoteConfigurationManagementJobPreview) SetTargetEntitiesCount(v int32) {
	o.TargetEntitiesCount = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementJobPreview) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementJobPreview) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementJobPreview) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RemoteConfigurationManagementJobPreview) SetValue(v string) {
	o.Value = &v
}

func (o RemoteConfigurationManagementJobPreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteConfigurationManagementJobPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AlreadyConfiguredEntitiesCount) {
		toSerialize["alreadyConfiguredEntitiesCount"] = o.AlreadyConfiguredEntitiesCount
	}
	if !IsNil(o.Attribute) {
		toSerialize["attribute"] = o.Attribute
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.TargetEntitiesCount) {
		toSerialize["targetEntitiesCount"] = o.TargetEntitiesCount
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRemoteConfigurationManagementJobPreview struct {
	value *RemoteConfigurationManagementJobPreview
	isSet bool
}

func (v NullableRemoteConfigurationManagementJobPreview) Get() *RemoteConfigurationManagementJobPreview {
	return v.value
}

func (v *NullableRemoteConfigurationManagementJobPreview) Set(val *RemoteConfigurationManagementJobPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteConfigurationManagementJobPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteConfigurationManagementJobPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteConfigurationManagementJobPreview(val *RemoteConfigurationManagementJobPreview) *NullableRemoteConfigurationManagementJobPreview {
	return &NullableRemoteConfigurationManagementJobPreview{value: val, isSet: true}
}

func (v NullableRemoteConfigurationManagementJobPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteConfigurationManagementJobPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


