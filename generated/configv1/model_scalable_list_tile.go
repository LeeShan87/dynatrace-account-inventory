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

// checks if the ScalableListTile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScalableListTile{}

// ScalableListTile Configuration of a tile with the built-in custom filter id. This is only for internal usage.
type ScalableListTile struct {
	ChartVisible *bool `json:"chartVisible,omitempty"`
	// The ID of the custom filter.
	CustomFilterId *string `json:"customFilterId,omitempty"`
	// The type of the entity specific tile.
	EntitySpecificTileType *string `json:"entitySpecificTileType,omitempty"`
}

// NewScalableListTile instantiates a new ScalableListTile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalableListTile(bounds TileBounds, name string, tileType string) *ScalableListTile {
	this := ScalableListTile{}
	this.Bounds = bounds
	this.Name = name
	this.TileType = tileType
	return &this
}

// NewScalableListTileWithDefaults instantiates a new ScalableListTile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalableListTileWithDefaults() *ScalableListTile {
	this := ScalableListTile{}
	return &this
}

// GetChartVisible returns the ChartVisible field value if set, zero value otherwise.
func (o *ScalableListTile) GetChartVisible() bool {
	if o == nil || IsNil(o.ChartVisible) {
		var ret bool
		return ret
	}
	return *o.ChartVisible
}

// GetChartVisibleOk returns a tuple with the ChartVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalableListTile) GetChartVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.ChartVisible) {
		return nil, false
	}
	return o.ChartVisible, true
}

// HasChartVisible returns a boolean if a field has been set.
func (o *ScalableListTile) HasChartVisible() bool {
	if o != nil && !IsNil(o.ChartVisible) {
		return true
	}

	return false
}

// SetChartVisible gets a reference to the given bool and assigns it to the ChartVisible field.
func (o *ScalableListTile) SetChartVisible(v bool) {
	o.ChartVisible = &v
}

// GetCustomFilterId returns the CustomFilterId field value if set, zero value otherwise.
func (o *ScalableListTile) GetCustomFilterId() string {
	if o == nil || IsNil(o.CustomFilterId) {
		var ret string
		return ret
	}
	return *o.CustomFilterId
}

// GetCustomFilterIdOk returns a tuple with the CustomFilterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalableListTile) GetCustomFilterIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomFilterId) {
		return nil, false
	}
	return o.CustomFilterId, true
}

// HasCustomFilterId returns a boolean if a field has been set.
func (o *ScalableListTile) HasCustomFilterId() bool {
	if o != nil && !IsNil(o.CustomFilterId) {
		return true
	}

	return false
}

// SetCustomFilterId gets a reference to the given string and assigns it to the CustomFilterId field.
func (o *ScalableListTile) SetCustomFilterId(v string) {
	o.CustomFilterId = &v
}

// GetEntitySpecificTileType returns the EntitySpecificTileType field value if set, zero value otherwise.
func (o *ScalableListTile) GetEntitySpecificTileType() string {
	if o == nil || IsNil(o.EntitySpecificTileType) {
		var ret string
		return ret
	}
	return *o.EntitySpecificTileType
}

// GetEntitySpecificTileTypeOk returns a tuple with the EntitySpecificTileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalableListTile) GetEntitySpecificTileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntitySpecificTileType) {
		return nil, false
	}
	return o.EntitySpecificTileType, true
}

// HasEntitySpecificTileType returns a boolean if a field has been set.
func (o *ScalableListTile) HasEntitySpecificTileType() bool {
	if o != nil && !IsNil(o.EntitySpecificTileType) {
		return true
	}

	return false
}

// SetEntitySpecificTileType gets a reference to the given string and assigns it to the EntitySpecificTileType field.
func (o *ScalableListTile) SetEntitySpecificTileType(v string) {
	o.EntitySpecificTileType = &v
}

func (o ScalableListTile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScalableListTile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChartVisible) {
		toSerialize["chartVisible"] = o.ChartVisible
	}
	if !IsNil(o.CustomFilterId) {
		toSerialize["customFilterId"] = o.CustomFilterId
	}
	if !IsNil(o.EntitySpecificTileType) {
		toSerialize["entitySpecificTileType"] = o.EntitySpecificTileType
	}
	return toSerialize, nil
}

type NullableScalableListTile struct {
	value *ScalableListTile
	isSet bool
}

func (v NullableScalableListTile) Get() *ScalableListTile {
	return v.value
}

func (v *NullableScalableListTile) Set(val *ScalableListTile) {
	v.value = val
	v.isSet = true
}

func (v NullableScalableListTile) IsSet() bool {
	return v.isSet
}

func (v *NullableScalableListTile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalableListTile(val *ScalableListTile) *NullableScalableListTile {
	return &NullableScalableListTile{value: val, isSet: true}
}

func (v NullableScalableListTile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalableListTile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


