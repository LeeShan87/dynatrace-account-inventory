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

// checks if the MaintenanceWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindow{}

// MaintenanceWindow Configuration of a maintenance window.
type MaintenanceWindow struct {
	// A short description of the maintenance purpose.
	Description string `json:"description"`
	// The ID of the maintenance window.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The name of the maintenance window, displayed in the UI.
	Name string `json:"name"`
	Schedule Schedule `json:"schedule"`
	Scope *Scope `json:"scope,omitempty"`
	// Suppress execution of synthetic monitors during the maintenance.
	SuppressSyntheticMonitorsExecution *bool `json:"suppressSyntheticMonitorsExecution,omitempty"`
	// The type of suppression of alerting and problem detection during the maintenance.
	Suppression string `json:"suppression"`
	// The type of the maintenance: planned or unplanned.
	Type string `json:"type"`
}

// NewMaintenanceWindow instantiates a new MaintenanceWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindow(description string, name string, schedule Schedule, suppression string, type_ string) *MaintenanceWindow {
	this := MaintenanceWindow{}
	this.Description = description
	this.Name = name
	this.Schedule = schedule
	this.Suppression = suppression
	this.Type = type_
	return &this
}

// NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowWithDefaults() *MaintenanceWindow {
	this := MaintenanceWindow{}
	return &this
}

// GetDescription returns the Description field value
func (o *MaintenanceWindow) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MaintenanceWindow) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MaintenanceWindow) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *MaintenanceWindow) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *MaintenanceWindow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MaintenanceWindow) SetName(v string) {
	o.Name = v
}

// GetSchedule returns the Schedule field value
func (o *MaintenanceWindow) GetSchedule() Schedule {
	if o == nil {
		var ret Schedule
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetScheduleOk() (*Schedule, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *MaintenanceWindow) SetSchedule(v Schedule) {
	o.Schedule = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetScope() Scope {
	if o == nil || IsNil(o.Scope) {
		var ret Scope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetScopeOk() (*Scope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given Scope and assigns it to the Scope field.
func (o *MaintenanceWindow) SetScope(v Scope) {
	o.Scope = &v
}

// GetSuppressSyntheticMonitorsExecution returns the SuppressSyntheticMonitorsExecution field value if set, zero value otherwise.
func (o *MaintenanceWindow) GetSuppressSyntheticMonitorsExecution() bool {
	if o == nil || IsNil(o.SuppressSyntheticMonitorsExecution) {
		var ret bool
		return ret
	}
	return *o.SuppressSyntheticMonitorsExecution
}

// GetSuppressSyntheticMonitorsExecutionOk returns a tuple with the SuppressSyntheticMonitorsExecution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetSuppressSyntheticMonitorsExecutionOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressSyntheticMonitorsExecution) {
		return nil, false
	}
	return o.SuppressSyntheticMonitorsExecution, true
}

// HasSuppressSyntheticMonitorsExecution returns a boolean if a field has been set.
func (o *MaintenanceWindow) HasSuppressSyntheticMonitorsExecution() bool {
	if o != nil && !IsNil(o.SuppressSyntheticMonitorsExecution) {
		return true
	}

	return false
}

// SetSuppressSyntheticMonitorsExecution gets a reference to the given bool and assigns it to the SuppressSyntheticMonitorsExecution field.
func (o *MaintenanceWindow) SetSuppressSyntheticMonitorsExecution(v bool) {
	o.SuppressSyntheticMonitorsExecution = &v
}

// GetSuppression returns the Suppression field value
func (o *MaintenanceWindow) GetSuppression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suppression
}

// GetSuppressionOk returns a tuple with the Suppression field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetSuppressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suppression, true
}

// SetSuppression sets field value
func (o *MaintenanceWindow) SetSuppression(v string) {
	o.Suppression = v
}

// GetType returns the Type field value
func (o *MaintenanceWindow) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindow) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MaintenanceWindow) SetType(v string) {
	o.Type = v
}

func (o MaintenanceWindow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.SuppressSyntheticMonitorsExecution) {
		toSerialize["suppressSyntheticMonitorsExecution"] = o.SuppressSyntheticMonitorsExecution
	}
	toSerialize["suppression"] = o.Suppression
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMaintenanceWindow struct {
	value *MaintenanceWindow
	isSet bool
}

func (v NullableMaintenanceWindow) Get() *MaintenanceWindow {
	return v.value
}

func (v *NullableMaintenanceWindow) Set(val *MaintenanceWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindow(val *MaintenanceWindow) *NullableMaintenanceWindow {
	return &NullableMaintenanceWindow{value: val, isSet: true}
}

func (v NullableMaintenanceWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

