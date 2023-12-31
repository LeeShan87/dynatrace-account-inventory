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

// checks if the DimensionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DimensionDefinition{}

// DimensionDefinition Parameters of a definition of a calculated service metric.
type DimensionDefinition struct {
	// The dimension value pattern.    You can define custom placeholders in the **placeholders** field and use them here.
	Dimension string `json:"dimension"`
	// The name of the dimension.
	Name string `json:"name"`
	// The list of custom placeholders to be used in a dimension value pattern.
	Placeholders []Placeholder `json:"placeholders,omitempty"`
	// The number of top values to be calculated.
	TopX int32 `json:"topX"`
	// The aggregation of the dimension.
	TopXAggregation string `json:"topXAggregation"`
	// How to calculate the **topX** values.
	TopXDirection string `json:"topXDirection"`
}

// NewDimensionDefinition instantiates a new DimensionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionDefinition(dimension string, name string, topX int32, topXAggregation string, topXDirection string) *DimensionDefinition {
	this := DimensionDefinition{}
	this.Dimension = dimension
	this.Name = name
	this.TopX = topX
	this.TopXAggregation = topXAggregation
	this.TopXDirection = topXDirection
	return &this
}

// NewDimensionDefinitionWithDefaults instantiates a new DimensionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionDefinitionWithDefaults() *DimensionDefinition {
	this := DimensionDefinition{}
	return &this
}

// GetDimension returns the Dimension field value
func (o *DimensionDefinition) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *DimensionDefinition) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *DimensionDefinition) SetDimension(v string) {
	o.Dimension = v
}

// GetName returns the Name field value
func (o *DimensionDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DimensionDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DimensionDefinition) SetName(v string) {
	o.Name = v
}

// GetPlaceholders returns the Placeholders field value if set, zero value otherwise.
func (o *DimensionDefinition) GetPlaceholders() []Placeholder {
	if o == nil || IsNil(o.Placeholders) {
		var ret []Placeholder
		return ret
	}
	return o.Placeholders
}

// GetPlaceholdersOk returns a tuple with the Placeholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionDefinition) GetPlaceholdersOk() ([]Placeholder, bool) {
	if o == nil || IsNil(o.Placeholders) {
		return nil, false
	}
	return o.Placeholders, true
}

// HasPlaceholders returns a boolean if a field has been set.
func (o *DimensionDefinition) HasPlaceholders() bool {
	if o != nil && !IsNil(o.Placeholders) {
		return true
	}

	return false
}

// SetPlaceholders gets a reference to the given []Placeholder and assigns it to the Placeholders field.
func (o *DimensionDefinition) SetPlaceholders(v []Placeholder) {
	o.Placeholders = v
}

// GetTopX returns the TopX field value
func (o *DimensionDefinition) GetTopX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TopX
}

// GetTopXOk returns a tuple with the TopX field value
// and a boolean to check if the value has been set.
func (o *DimensionDefinition) GetTopXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopX, true
}

// SetTopX sets field value
func (o *DimensionDefinition) SetTopX(v int32) {
	o.TopX = v
}

// GetTopXAggregation returns the TopXAggregation field value
func (o *DimensionDefinition) GetTopXAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopXAggregation
}

// GetTopXAggregationOk returns a tuple with the TopXAggregation field value
// and a boolean to check if the value has been set.
func (o *DimensionDefinition) GetTopXAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopXAggregation, true
}

// SetTopXAggregation sets field value
func (o *DimensionDefinition) SetTopXAggregation(v string) {
	o.TopXAggregation = v
}

// GetTopXDirection returns the TopXDirection field value
func (o *DimensionDefinition) GetTopXDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopXDirection
}

// GetTopXDirectionOk returns a tuple with the TopXDirection field value
// and a boolean to check if the value has been set.
func (o *DimensionDefinition) GetTopXDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopXDirection, true
}

// SetTopXDirection sets field value
func (o *DimensionDefinition) SetTopXDirection(v string) {
	o.TopXDirection = v
}

func (o DimensionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DimensionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dimension"] = o.Dimension
	toSerialize["name"] = o.Name
	if !IsNil(o.Placeholders) {
		toSerialize["placeholders"] = o.Placeholders
	}
	toSerialize["topX"] = o.TopX
	toSerialize["topXAggregation"] = o.TopXAggregation
	toSerialize["topXDirection"] = o.TopXDirection
	return toSerialize, nil
}

type NullableDimensionDefinition struct {
	value *DimensionDefinition
	isSet bool
}

func (v NullableDimensionDefinition) Get() *DimensionDefinition {
	return v.value
}

func (v *NullableDimensionDefinition) Set(val *DimensionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensionDefinition(val *DimensionDefinition) *NullableDimensionDefinition {
	return &NullableDimensionDefinition{value: val, isSet: true}
}

func (v NullableDimensionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDimensionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


