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

// checks if the MetricDimensionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDimensionDefinition{}

// MetricDimensionDefinition The dimension of a metric.
type MetricDimensionDefinition struct {
	// The display name of the dimension.
	DisplayName string `json:"displayName"`
	// The unique 0-based index of the dimension.    Appending transformations such as :names or :parents may change the indexes of dimensions. `null` is used for the dimensions of a metric with flexible dimensions, which can be referenced with their dimension key, but do not have an intrinsic order that could be used for the index.
	Index int32 `json:"index"`
	// The key of the dimension.    It must be unique within the metric.
	Key string `json:"key"`
	// The name of the dimension.
	Name string `json:"name"`
	// The type of the dimension.
	Type string `json:"type"`
}

// NewMetricDimensionDefinition instantiates a new MetricDimensionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDimensionDefinition(displayName string, index int32, key string, name string, type_ string) *MetricDimensionDefinition {
	this := MetricDimensionDefinition{}
	this.DisplayName = displayName
	this.Index = index
	this.Key = key
	this.Name = name
	this.Type = type_
	return &this
}

// NewMetricDimensionDefinitionWithDefaults instantiates a new MetricDimensionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDimensionDefinitionWithDefaults() *MetricDimensionDefinition {
	this := MetricDimensionDefinition{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *MetricDimensionDefinition) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *MetricDimensionDefinition) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIndex returns the Index field value
func (o *MetricDimensionDefinition) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionDefinition) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *MetricDimensionDefinition) SetIndex(v int32) {
	o.Index = v
}

// GetKey returns the Key field value
func (o *MetricDimensionDefinition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionDefinition) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MetricDimensionDefinition) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *MetricDimensionDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricDimensionDefinition) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *MetricDimensionDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricDimensionDefinition) SetType(v string) {
	o.Type = v
}

func (o MetricDimensionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricDimensionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["index"] = o.Index
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMetricDimensionDefinition struct {
	value *MetricDimensionDefinition
	isSet bool
}

func (v NullableMetricDimensionDefinition) Get() *MetricDimensionDefinition {
	return v.value
}

func (v *NullableMetricDimensionDefinition) Set(val *MetricDimensionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDimensionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDimensionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDimensionDefinition(val *MetricDimensionDefinition) *NullableMetricDimensionDefinition {
	return &NullableMetricDimensionDefinition{value: val, isSet: true}
}

func (v NullableMetricDimensionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDimensionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


