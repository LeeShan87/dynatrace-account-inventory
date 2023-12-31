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

// checks if the SettingsObjectUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsObjectUpdate{}

// SettingsObjectUpdate An update of a settings object.
type SettingsObjectUpdate struct {
	// The position of the updated object. The new object will be moved behind the specified one.   **insertAfter** and **insertBefore** are evaluated together and only one of both can be set.   If `null` and **insertBefore** 'null', the existing object keeps the current position.   If set to empty string, the updated object will be placed in the first position.   Only applicable for objects based on schemas with ordered objects (schema's **ordered** parameter is set to `true`).
	InsertAfter *string `json:"insertAfter,omitempty"`
	// The position of the updated object. The new object will be moved in front of the specified one.   **insertAfter** and **insertBefore** are evaluated together and only one of both can be set.   If `null` and **insertAfter** 'null', the existing object keeps the current position.   If set to empty string, the updated object will be placed in the last position.   Only applicable for objects based on schemas with ordered objects (schema's **ordered** parameter is set to `true`).
	InsertBefore *string `json:"insertBefore,omitempty"`
	// The version of the schema on which the object is based.
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	// The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn't any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks.
	UpdateToken *string `json:"updateToken,omitempty"`
	// The value of the setting.    It defines the actual values of settings' parameters.   The actual content depends on the object's schema.
	Value map[string]interface{} `json:"value"`
}

// NewSettingsObjectUpdate instantiates a new SettingsObjectUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsObjectUpdate(value map[string]interface{}) *SettingsObjectUpdate {
	this := SettingsObjectUpdate{}
	this.Value = value
	return &this
}

// NewSettingsObjectUpdateWithDefaults instantiates a new SettingsObjectUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsObjectUpdateWithDefaults() *SettingsObjectUpdate {
	this := SettingsObjectUpdate{}
	return &this
}

// GetInsertAfter returns the InsertAfter field value if set, zero value otherwise.
func (o *SettingsObjectUpdate) GetInsertAfter() string {
	if o == nil || IsNil(o.InsertAfter) {
		var ret string
		return ret
	}
	return *o.InsertAfter
}

// GetInsertAfterOk returns a tuple with the InsertAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectUpdate) GetInsertAfterOk() (*string, bool) {
	if o == nil || IsNil(o.InsertAfter) {
		return nil, false
	}
	return o.InsertAfter, true
}

// HasInsertAfter returns a boolean if a field has been set.
func (o *SettingsObjectUpdate) HasInsertAfter() bool {
	if o != nil && !IsNil(o.InsertAfter) {
		return true
	}

	return false
}

// SetInsertAfter gets a reference to the given string and assigns it to the InsertAfter field.
func (o *SettingsObjectUpdate) SetInsertAfter(v string) {
	o.InsertAfter = &v
}

// GetInsertBefore returns the InsertBefore field value if set, zero value otherwise.
func (o *SettingsObjectUpdate) GetInsertBefore() string {
	if o == nil || IsNil(o.InsertBefore) {
		var ret string
		return ret
	}
	return *o.InsertBefore
}

// GetInsertBeforeOk returns a tuple with the InsertBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectUpdate) GetInsertBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.InsertBefore) {
		return nil, false
	}
	return o.InsertBefore, true
}

// HasInsertBefore returns a boolean if a field has been set.
func (o *SettingsObjectUpdate) HasInsertBefore() bool {
	if o != nil && !IsNil(o.InsertBefore) {
		return true
	}

	return false
}

// SetInsertBefore gets a reference to the given string and assigns it to the InsertBefore field.
func (o *SettingsObjectUpdate) SetInsertBefore(v string) {
	o.InsertBefore = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *SettingsObjectUpdate) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectUpdate) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *SettingsObjectUpdate) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *SettingsObjectUpdate) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetUpdateToken returns the UpdateToken field value if set, zero value otherwise.
func (o *SettingsObjectUpdate) GetUpdateToken() string {
	if o == nil || IsNil(o.UpdateToken) {
		var ret string
		return ret
	}
	return *o.UpdateToken
}

// GetUpdateTokenOk returns a tuple with the UpdateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObjectUpdate) GetUpdateTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateToken) {
		return nil, false
	}
	return o.UpdateToken, true
}

// HasUpdateToken returns a boolean if a field has been set.
func (o *SettingsObjectUpdate) HasUpdateToken() bool {
	if o != nil && !IsNil(o.UpdateToken) {
		return true
	}

	return false
}

// SetUpdateToken gets a reference to the given string and assigns it to the UpdateToken field.
func (o *SettingsObjectUpdate) SetUpdateToken(v string) {
	o.UpdateToken = &v
}

// GetValue returns the Value field value
func (o *SettingsObjectUpdate) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SettingsObjectUpdate) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *SettingsObjectUpdate) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o SettingsObjectUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsObjectUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InsertAfter) {
		toSerialize["insertAfter"] = o.InsertAfter
	}
	if !IsNil(o.InsertBefore) {
		toSerialize["insertBefore"] = o.InsertBefore
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	if !IsNil(o.UpdateToken) {
		toSerialize["updateToken"] = o.UpdateToken
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSettingsObjectUpdate struct {
	value *SettingsObjectUpdate
	isSet bool
}

func (v NullableSettingsObjectUpdate) Get() *SettingsObjectUpdate {
	return v.value
}

func (v *NullableSettingsObjectUpdate) Set(val *SettingsObjectUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsObjectUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsObjectUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsObjectUpdate(val *SettingsObjectUpdate) *NullableSettingsObjectUpdate {
	return &NullableSettingsObjectUpdate{value: val, isSet: true}
}

func (v NullableSettingsObjectUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsObjectUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


