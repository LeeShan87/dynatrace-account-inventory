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

// checks if the TileFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TileFilter{}

// TileFilter A filter applied to a tile.    It overrides dashboard's filter.
type TileFilter struct {
	ManagementZone *EntityShortRepresentation `json:"managementZone,omitempty"`
	// The default timeframe of the tile.
	Timeframe *string `json:"timeframe,omitempty"`
}

// NewTileFilter instantiates a new TileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTileFilter() *TileFilter {
	this := TileFilter{}
	return &this
}

// NewTileFilterWithDefaults instantiates a new TileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTileFilterWithDefaults() *TileFilter {
	this := TileFilter{}
	return &this
}

// GetManagementZone returns the ManagementZone field value if set, zero value otherwise.
func (o *TileFilter) GetManagementZone() EntityShortRepresentation {
	if o == nil || IsNil(o.ManagementZone) {
		var ret EntityShortRepresentation
		return ret
	}
	return *o.ManagementZone
}

// GetManagementZoneOk returns a tuple with the ManagementZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TileFilter) GetManagementZoneOk() (*EntityShortRepresentation, bool) {
	if o == nil || IsNil(o.ManagementZone) {
		return nil, false
	}
	return o.ManagementZone, true
}

// HasManagementZone returns a boolean if a field has been set.
func (o *TileFilter) HasManagementZone() bool {
	if o != nil && !IsNil(o.ManagementZone) {
		return true
	}

	return false
}

// SetManagementZone gets a reference to the given EntityShortRepresentation and assigns it to the ManagementZone field.
func (o *TileFilter) SetManagementZone(v EntityShortRepresentation) {
	o.ManagementZone = &v
}

// GetTimeframe returns the Timeframe field value if set, zero value otherwise.
func (o *TileFilter) GetTimeframe() string {
	if o == nil || IsNil(o.Timeframe) {
		var ret string
		return ret
	}
	return *o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TileFilter) GetTimeframeOk() (*string, bool) {
	if o == nil || IsNil(o.Timeframe) {
		return nil, false
	}
	return o.Timeframe, true
}

// HasTimeframe returns a boolean if a field has been set.
func (o *TileFilter) HasTimeframe() bool {
	if o != nil && !IsNil(o.Timeframe) {
		return true
	}

	return false
}

// SetTimeframe gets a reference to the given string and assigns it to the Timeframe field.
func (o *TileFilter) SetTimeframe(v string) {
	o.Timeframe = &v
}

func (o TileFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TileFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagementZone) {
		toSerialize["managementZone"] = o.ManagementZone
	}
	if !IsNil(o.Timeframe) {
		toSerialize["timeframe"] = o.Timeframe
	}
	return toSerialize, nil
}

type NullableTileFilter struct {
	value *TileFilter
	isSet bool
}

func (v NullableTileFilter) Get() *TileFilter {
	return v.value
}

func (v *NullableTileFilter) Set(val *TileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableTileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableTileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTileFilter(val *TileFilter) *NullableTileFilter {
	return &NullableTileFilter{value: val, isSet: true}
}

func (v NullableTileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


