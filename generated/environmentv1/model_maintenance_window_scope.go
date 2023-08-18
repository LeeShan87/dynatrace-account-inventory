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

// checks if the MaintenanceWindowScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindowScope{}

// MaintenanceWindowScope An object defining the scope of your maintenance window.   You can specify particular Dynatrace entities or matching rules for dynamic formation of the scope.    If no scope is specified, the maintenance applies to the entire environment.    To specify the scope at least one entity or matching rule must be specified.
type MaintenanceWindowScope struct {
	// Defines Dynatrace entities to be included in scope, for example hosts, services, process groups.   Allowed values are Dynatrace entity IDs.
	Entities []string `json:"entities,omitempty"`
	// An object defining a matching rule for dynamic scope formation. An empty rule matches for all entities.
	Matches []MonitoredEntityFilter `json:"matches,omitempty"`
}

// NewMaintenanceWindowScope instantiates a new MaintenanceWindowScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindowScope() *MaintenanceWindowScope {
	this := MaintenanceWindowScope{}
	return &this
}

// NewMaintenanceWindowScopeWithDefaults instantiates a new MaintenanceWindowScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowScopeWithDefaults() *MaintenanceWindowScope {
	this := MaintenanceWindowScope{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *MaintenanceWindowScope) GetEntities() []string {
	if o == nil || IsNil(o.Entities) {
		var ret []string
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowScope) GetEntitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *MaintenanceWindowScope) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []string and assigns it to the Entities field.
func (o *MaintenanceWindowScope) SetEntities(v []string) {
	o.Entities = v
}

// GetMatches returns the Matches field value if set, zero value otherwise.
func (o *MaintenanceWindowScope) GetMatches() []MonitoredEntityFilter {
	if o == nil || IsNil(o.Matches) {
		var ret []MonitoredEntityFilter
		return ret
	}
	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowScope) GetMatchesOk() ([]MonitoredEntityFilter, bool) {
	if o == nil || IsNil(o.Matches) {
		return nil, false
	}
	return o.Matches, true
}

// HasMatches returns a boolean if a field has been set.
func (o *MaintenanceWindowScope) HasMatches() bool {
	if o != nil && !IsNil(o.Matches) {
		return true
	}

	return false
}

// SetMatches gets a reference to the given []MonitoredEntityFilter and assigns it to the Matches field.
func (o *MaintenanceWindowScope) SetMatches(v []MonitoredEntityFilter) {
	o.Matches = v
}

func (o MaintenanceWindowScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceWindowScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.Matches) {
		toSerialize["matches"] = o.Matches
	}
	return toSerialize, nil
}

type NullableMaintenanceWindowScope struct {
	value *MaintenanceWindowScope
	isSet bool
}

func (v NullableMaintenanceWindowScope) Get() *MaintenanceWindowScope {
	return v.value
}

func (v *NullableMaintenanceWindowScope) Set(val *MaintenanceWindowScope) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindowScope) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindowScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindowScope(val *MaintenanceWindowScope) *NullableMaintenanceWindowScope {
	return &NullableMaintenanceWindowScope{value: val, isSet: true}
}

func (v NullableMaintenanceWindowScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindowScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

