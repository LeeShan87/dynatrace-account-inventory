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

// checks if the Dashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dashboard{}

// Dashboard Configuration of a dashboard.
type Dashboard struct {
	DashboardMetadata DashboardMetadata `json:"dashboardMetadata"`
	// The ID of the dashboard.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The list of tiles on the dashboard.
	Tiles []Tile `json:"tiles"`
}

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard(dashboardMetadata DashboardMetadata, tiles []Tile) *Dashboard {
	this := Dashboard{}
	this.DashboardMetadata = dashboardMetadata
	this.Tiles = tiles
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	return &this
}

// GetDashboardMetadata returns the DashboardMetadata field value
func (o *Dashboard) GetDashboardMetadata() DashboardMetadata {
	if o == nil {
		var ret DashboardMetadata
		return ret
	}

	return o.DashboardMetadata
}

// GetDashboardMetadataOk returns a tuple with the DashboardMetadata field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDashboardMetadataOk() (*DashboardMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardMetadata, true
}

// SetDashboardMetadata sets field value
func (o *Dashboard) SetDashboardMetadata(v DashboardMetadata) {
	o.DashboardMetadata = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Dashboard) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Dashboard) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Dashboard) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Dashboard) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Dashboard) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *Dashboard) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetTiles returns the Tiles field value
func (o *Dashboard) GetTiles() []Tile {
	if o == nil {
		var ret []Tile
		return ret
	}

	return o.Tiles
}

// GetTilesOk returns a tuple with the Tiles field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetTilesOk() ([]Tile, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tiles, true
}

// SetTiles sets field value
func (o *Dashboard) SetTiles(v []Tile) {
	o.Tiles = v
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboardMetadata"] = o.DashboardMetadata
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["tiles"] = o.Tiles
	return toSerialize, nil
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


