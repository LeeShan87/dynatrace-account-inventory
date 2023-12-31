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

// checks if the BurnRateAlert type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BurnRateAlert{}

// BurnRateAlert Parameters of an error budget burn rate alert.
type BurnRateAlert struct {
}

// NewBurnRateAlert instantiates a new BurnRateAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBurnRateAlert(alertName string, alertThreshold float64, alertType string) *BurnRateAlert {
	this := BurnRateAlert{}
	this.AlertName = alertName
	this.AlertThreshold = alertThreshold
	this.AlertType = alertType
	return &this
}

// NewBurnRateAlertWithDefaults instantiates a new BurnRateAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBurnRateAlertWithDefaults() *BurnRateAlert {
	this := BurnRateAlert{}
	return &this
}

func (o BurnRateAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BurnRateAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableBurnRateAlert struct {
	value *BurnRateAlert
	isSet bool
}

func (v NullableBurnRateAlert) Get() *BurnRateAlert {
	return v.value
}

func (v *NullableBurnRateAlert) Set(val *BurnRateAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableBurnRateAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableBurnRateAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBurnRateAlert(val *BurnRateAlert) *NullableBurnRateAlert {
	return &NullableBurnRateAlert{value: val, isSet: true}
}

func (v NullableBurnRateAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBurnRateAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


