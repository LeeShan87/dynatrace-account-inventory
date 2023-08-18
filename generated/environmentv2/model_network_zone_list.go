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

// checks if the NetworkZoneList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkZoneList{}

// NetworkZoneList A list of network zones.
type NetworkZoneList struct {
	// A list of network zones.
	NetworkZones []NetworkZone `json:"networkZones"`
}

// NewNetworkZoneList instantiates a new NetworkZoneList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkZoneList(networkZones []NetworkZone) *NetworkZoneList {
	this := NetworkZoneList{}
	this.NetworkZones = networkZones
	return &this
}

// NewNetworkZoneListWithDefaults instantiates a new NetworkZoneList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkZoneListWithDefaults() *NetworkZoneList {
	this := NetworkZoneList{}
	return &this
}

// GetNetworkZones returns the NetworkZones field value
func (o *NetworkZoneList) GetNetworkZones() []NetworkZone {
	if o == nil {
		var ret []NetworkZone
		return ret
	}

	return o.NetworkZones
}

// GetNetworkZonesOk returns a tuple with the NetworkZones field value
// and a boolean to check if the value has been set.
func (o *NetworkZoneList) GetNetworkZonesOk() ([]NetworkZone, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkZones, true
}

// SetNetworkZones sets field value
func (o *NetworkZoneList) SetNetworkZones(v []NetworkZone) {
	o.NetworkZones = v
}

func (o NetworkZoneList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkZoneList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkZones"] = o.NetworkZones
	return toSerialize, nil
}

type NullableNetworkZoneList struct {
	value *NetworkZoneList
	isSet bool
}

func (v NullableNetworkZoneList) Get() *NetworkZoneList {
	return v.value
}

func (v *NullableNetworkZoneList) Set(val *NetworkZoneList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkZoneList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkZoneList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkZoneList(val *NetworkZoneList) *NullableNetworkZoneList {
	return &NullableNetworkZoneList{value: val, isSet: true}
}

func (v NullableNetworkZoneList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkZoneList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


