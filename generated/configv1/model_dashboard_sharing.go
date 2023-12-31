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

// checks if the DashboardSharing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardSharing{}

// DashboardSharing Sharing configuration of the dashboard.
type DashboardSharing struct {
	// The dashboard is shared (`true`) or private (`false`).
	Enabled *bool `json:"enabled,omitempty"`
	// The Dynatrace entity ID of the dashboard.
	Id string `json:"id"`
	// A list of permissions to access the dashboard.
	Permissions []DashboardSharePermissions `json:"permissions"`
	// If `true` the dashboard will be marked as preset.
	Preset *bool `json:"preset,omitempty"`
	PublicAccess DashboardAnonymousAccess `json:"publicAccess"`
}

// NewDashboardSharing instantiates a new DashboardSharing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardSharing(id string, permissions []DashboardSharePermissions, publicAccess DashboardAnonymousAccess) *DashboardSharing {
	this := DashboardSharing{}
	this.Id = id
	this.Permissions = permissions
	this.PublicAccess = publicAccess
	return &this
}

// NewDashboardSharingWithDefaults instantiates a new DashboardSharing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardSharingWithDefaults() *DashboardSharing {
	this := DashboardSharing{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DashboardSharing) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSharing) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DashboardSharing) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DashboardSharing) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value
func (o *DashboardSharing) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardSharing) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DashboardSharing) SetId(v string) {
	o.Id = v
}

// GetPermissions returns the Permissions field value
func (o *DashboardSharing) GetPermissions() []DashboardSharePermissions {
	if o == nil {
		var ret []DashboardSharePermissions
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *DashboardSharing) GetPermissionsOk() ([]DashboardSharePermissions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *DashboardSharing) SetPermissions(v []DashboardSharePermissions) {
	o.Permissions = v
}

// GetPreset returns the Preset field value if set, zero value otherwise.
func (o *DashboardSharing) GetPreset() bool {
	if o == nil || IsNil(o.Preset) {
		var ret bool
		return ret
	}
	return *o.Preset
}

// GetPresetOk returns a tuple with the Preset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardSharing) GetPresetOk() (*bool, bool) {
	if o == nil || IsNil(o.Preset) {
		return nil, false
	}
	return o.Preset, true
}

// HasPreset returns a boolean if a field has been set.
func (o *DashboardSharing) HasPreset() bool {
	if o != nil && !IsNil(o.Preset) {
		return true
	}

	return false
}

// SetPreset gets a reference to the given bool and assigns it to the Preset field.
func (o *DashboardSharing) SetPreset(v bool) {
	o.Preset = &v
}

// GetPublicAccess returns the PublicAccess field value
func (o *DashboardSharing) GetPublicAccess() DashboardAnonymousAccess {
	if o == nil {
		var ret DashboardAnonymousAccess
		return ret
	}

	return o.PublicAccess
}

// GetPublicAccessOk returns a tuple with the PublicAccess field value
// and a boolean to check if the value has been set.
func (o *DashboardSharing) GetPublicAccessOk() (*DashboardAnonymousAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicAccess, true
}

// SetPublicAccess sets field value
func (o *DashboardSharing) SetPublicAccess(v DashboardAnonymousAccess) {
	o.PublicAccess = v
}

func (o DashboardSharing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardSharing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["id"] = o.Id
	toSerialize["permissions"] = o.Permissions
	if !IsNil(o.Preset) {
		toSerialize["preset"] = o.Preset
	}
	toSerialize["publicAccess"] = o.PublicAccess
	return toSerialize, nil
}

type NullableDashboardSharing struct {
	value *DashboardSharing
	isSet bool
}

func (v NullableDashboardSharing) Get() *DashboardSharing {
	return v.value
}

func (v *NullableDashboardSharing) Set(val *DashboardSharing) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardSharing) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardSharing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardSharing(val *DashboardSharing) *NullableDashboardSharing {
	return &NullableDashboardSharing{value: val, isSet: true}
}

func (v NullableDashboardSharing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardSharing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


