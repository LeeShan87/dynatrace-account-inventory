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

// checks if the MaintenanceWindowEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaintenanceWindowEvidence{}

// MaintenanceWindowEvidence The  maintenance window evidence of the problem.   The maintenance window during which the problem occurred.
type MaintenanceWindowEvidence struct {
	// The end time of the evidence, in UTC milliseconds.
	EndTime int64 `json:"endTime"`
	// The ID of the related maintenance window.
	MaintenanceWindowConfigId string `json:"maintenanceWindowConfigId"`
}

// NewMaintenanceWindowEvidence instantiates a new MaintenanceWindowEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaintenanceWindowEvidence(endTime int64, maintenanceWindowConfigId string, displayName string, entity EntityStub, evidenceType string, rootCauseRelevant bool, startTime int64) *MaintenanceWindowEvidence {
	this := MaintenanceWindowEvidence{}
	this.DisplayName = displayName
	this.Entity = entity
	this.EvidenceType = evidenceType
	this.RootCauseRelevant = rootCauseRelevant
	this.StartTime = startTime
	this.EndTime = endTime
	this.MaintenanceWindowConfigId = maintenanceWindowConfigId
	return &this
}

// NewMaintenanceWindowEvidenceWithDefaults instantiates a new MaintenanceWindowEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaintenanceWindowEvidenceWithDefaults() *MaintenanceWindowEvidence {
	this := MaintenanceWindowEvidence{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *MaintenanceWindowEvidence) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowEvidence) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *MaintenanceWindowEvidence) SetEndTime(v int64) {
	o.EndTime = v
}

// GetMaintenanceWindowConfigId returns the MaintenanceWindowConfigId field value
func (o *MaintenanceWindowEvidence) GetMaintenanceWindowConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaintenanceWindowConfigId
}

// GetMaintenanceWindowConfigIdOk returns a tuple with the MaintenanceWindowConfigId field value
// and a boolean to check if the value has been set.
func (o *MaintenanceWindowEvidence) GetMaintenanceWindowConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenanceWindowConfigId, true
}

// SetMaintenanceWindowConfigId sets field value
func (o *MaintenanceWindowEvidence) SetMaintenanceWindowConfigId(v string) {
	o.MaintenanceWindowConfigId = v
}

func (o MaintenanceWindowEvidence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaintenanceWindowEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endTime"] = o.EndTime
	toSerialize["maintenanceWindowConfigId"] = o.MaintenanceWindowConfigId
	return toSerialize, nil
}

type NullableMaintenanceWindowEvidence struct {
	value *MaintenanceWindowEvidence
	isSet bool
}

func (v NullableMaintenanceWindowEvidence) Get() *MaintenanceWindowEvidence {
	return v.value
}

func (v *NullableMaintenanceWindowEvidence) Set(val *MaintenanceWindowEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableMaintenanceWindowEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableMaintenanceWindowEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaintenanceWindowEvidence(val *MaintenanceWindowEvidence) *NullableMaintenanceWindowEvidence {
	return &NullableMaintenanceWindowEvidence{value: val, isSet: true}
}

func (v NullableMaintenanceWindowEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaintenanceWindowEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


