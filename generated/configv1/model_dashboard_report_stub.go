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

// checks if the DashboardReportStub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardReportStub{}

// DashboardReportStub A short representations of the report.
type DashboardReportStub struct {
	// The ID of the associated dashboard.
	DashboardId string `json:"dashboardId"`
	// The ID of the report.
	Id string `json:"id"`
	// The type of the report.
	Type string `json:"type"`
}

// NewDashboardReportStub instantiates a new DashboardReportStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardReportStub(dashboardId string, id string, type_ string) *DashboardReportStub {
	this := DashboardReportStub{}
	this.DashboardId = dashboardId
	this.Id = id
	this.Type = type_
	return &this
}

// NewDashboardReportStubWithDefaults instantiates a new DashboardReportStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardReportStubWithDefaults() *DashboardReportStub {
	this := DashboardReportStub{}
	return &this
}

// GetDashboardId returns the DashboardId field value
func (o *DashboardReportStub) GetDashboardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value
// and a boolean to check if the value has been set.
func (o *DashboardReportStub) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardId, true
}

// SetDashboardId sets field value
func (o *DashboardReportStub) SetDashboardId(v string) {
	o.DashboardId = v
}

// GetId returns the Id field value
func (o *DashboardReportStub) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardReportStub) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DashboardReportStub) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DashboardReportStub) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardReportStub) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DashboardReportStub) SetType(v string) {
	o.Type = v
}

func (o DashboardReportStub) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardReportStub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboardId"] = o.DashboardId
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableDashboardReportStub struct {
	value *DashboardReportStub
	isSet bool
}

func (v NullableDashboardReportStub) Get() *DashboardReportStub {
	return v.value
}

func (v *NullableDashboardReportStub) Set(val *DashboardReportStub) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardReportStub) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardReportStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardReportStub(val *DashboardReportStub) *NullableDashboardReportStub {
	return &NullableDashboardReportStub{value: val, isSet: true}
}

func (v NullableDashboardReportStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardReportStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


