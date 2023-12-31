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

// checks if the MetricMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricMetadataDto{}

// MetricMetadataDto Metric metadata
type MetricMetadataDto struct {
	// A short description of the metric
	Description *string `json:"description,omitempty"`
	// The name of the metric in the user interface
	DisplayName *string `json:"displayName,omitempty"`
	// The unit of the metric
	Unit *string `json:"unit,omitempty"`
}

// NewMetricMetadataDto instantiates a new MetricMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricMetadataDto() *MetricMetadataDto {
	this := MetricMetadataDto{}
	return &this
}

// NewMetricMetadataDtoWithDefaults instantiates a new MetricMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricMetadataDtoWithDefaults() *MetricMetadataDto {
	this := MetricMetadataDto{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricMetadataDto) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadataDto) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricMetadataDto) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricMetadataDto) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MetricMetadataDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadataDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MetricMetadataDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MetricMetadataDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricMetadataDto) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadataDto) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricMetadataDto) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricMetadataDto) SetUnit(v string) {
	o.Unit = &v
}

func (o MetricMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableMetricMetadataDto struct {
	value *MetricMetadataDto
	isSet bool
}

func (v NullableMetricMetadataDto) Get() *MetricMetadataDto {
	return v.value
}

func (v *NullableMetricMetadataDto) Set(val *MetricMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricMetadataDto(val *MetricMetadataDto) *NullableMetricMetadataDto {
	return &NullableMetricMetadataDto{value: val, isSet: true}
}

func (v NullableMetricMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


