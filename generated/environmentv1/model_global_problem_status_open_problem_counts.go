/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the GlobalProblemStatusOpenProblemCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalProblemStatusOpenProblemCounts{}

// GlobalProblemStatusOpenProblemCounts Numbers of open problems per impact level.
type GlobalProblemStatusOpenProblemCounts struct {
	// The count of impacted entities for the `APPLICATION` impact level.
	APPLICATION *int32 `json:"APPLICATION,omitempty"`
	// The count of impacted entities for the `ENVIRONMENT` impact level.
	ENVIRONMENT *int32 `json:"ENVIRONMENT,omitempty"`
	// The count of impacted entities for the `INFRASTRUCTURE` impact level.
	INFRASTRUCTURE *int32 `json:"INFRASTRUCTURE,omitempty"`
	// The count of impacted entities for the `SERVICE` impact level.
	SERVICE *int32 `json:"SERVICE,omitempty"`
}

// NewGlobalProblemStatusOpenProblemCounts instantiates a new GlobalProblemStatusOpenProblemCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalProblemStatusOpenProblemCounts() *GlobalProblemStatusOpenProblemCounts {
	this := GlobalProblemStatusOpenProblemCounts{}
	return &this
}

// NewGlobalProblemStatusOpenProblemCountsWithDefaults instantiates a new GlobalProblemStatusOpenProblemCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalProblemStatusOpenProblemCountsWithDefaults() *GlobalProblemStatusOpenProblemCounts {
	this := GlobalProblemStatusOpenProblemCounts{}
	return &this
}

// GetAPPLICATION returns the APPLICATION field value if set, zero value otherwise.
func (o *GlobalProblemStatusOpenProblemCounts) GetAPPLICATION() int32 {
	if o == nil || IsNil(o.APPLICATION) {
		var ret int32
		return ret
	}
	return *o.APPLICATION
}

// GetAPPLICATIONOk returns a tuple with the APPLICATION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProblemStatusOpenProblemCounts) GetAPPLICATIONOk() (*int32, bool) {
	if o == nil || IsNil(o.APPLICATION) {
		return nil, false
	}
	return o.APPLICATION, true
}

// HasAPPLICATION returns a boolean if a field has been set.
func (o *GlobalProblemStatusOpenProblemCounts) HasAPPLICATION() bool {
	if o != nil && !IsNil(o.APPLICATION) {
		return true
	}

	return false
}

// SetAPPLICATION gets a reference to the given int32 and assigns it to the APPLICATION field.
func (o *GlobalProblemStatusOpenProblemCounts) SetAPPLICATION(v int32) {
	o.APPLICATION = &v
}

// GetENVIRONMENT returns the ENVIRONMENT field value if set, zero value otherwise.
func (o *GlobalProblemStatusOpenProblemCounts) GetENVIRONMENT() int32 {
	if o == nil || IsNil(o.ENVIRONMENT) {
		var ret int32
		return ret
	}
	return *o.ENVIRONMENT
}

// GetENVIRONMENTOk returns a tuple with the ENVIRONMENT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProblemStatusOpenProblemCounts) GetENVIRONMENTOk() (*int32, bool) {
	if o == nil || IsNil(o.ENVIRONMENT) {
		return nil, false
	}
	return o.ENVIRONMENT, true
}

// HasENVIRONMENT returns a boolean if a field has been set.
func (o *GlobalProblemStatusOpenProblemCounts) HasENVIRONMENT() bool {
	if o != nil && !IsNil(o.ENVIRONMENT) {
		return true
	}

	return false
}

// SetENVIRONMENT gets a reference to the given int32 and assigns it to the ENVIRONMENT field.
func (o *GlobalProblemStatusOpenProblemCounts) SetENVIRONMENT(v int32) {
	o.ENVIRONMENT = &v
}

// GetINFRASTRUCTURE returns the INFRASTRUCTURE field value if set, zero value otherwise.
func (o *GlobalProblemStatusOpenProblemCounts) GetINFRASTRUCTURE() int32 {
	if o == nil || IsNil(o.INFRASTRUCTURE) {
		var ret int32
		return ret
	}
	return *o.INFRASTRUCTURE
}

// GetINFRASTRUCTUREOk returns a tuple with the INFRASTRUCTURE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProblemStatusOpenProblemCounts) GetINFRASTRUCTUREOk() (*int32, bool) {
	if o == nil || IsNil(o.INFRASTRUCTURE) {
		return nil, false
	}
	return o.INFRASTRUCTURE, true
}

// HasINFRASTRUCTURE returns a boolean if a field has been set.
func (o *GlobalProblemStatusOpenProblemCounts) HasINFRASTRUCTURE() bool {
	if o != nil && !IsNil(o.INFRASTRUCTURE) {
		return true
	}

	return false
}

// SetINFRASTRUCTURE gets a reference to the given int32 and assigns it to the INFRASTRUCTURE field.
func (o *GlobalProblemStatusOpenProblemCounts) SetINFRASTRUCTURE(v int32) {
	o.INFRASTRUCTURE = &v
}

// GetSERVICE returns the SERVICE field value if set, zero value otherwise.
func (o *GlobalProblemStatusOpenProblemCounts) GetSERVICE() int32 {
	if o == nil || IsNil(o.SERVICE) {
		var ret int32
		return ret
	}
	return *o.SERVICE
}

// GetSERVICEOk returns a tuple with the SERVICE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalProblemStatusOpenProblemCounts) GetSERVICEOk() (*int32, bool) {
	if o == nil || IsNil(o.SERVICE) {
		return nil, false
	}
	return o.SERVICE, true
}

// HasSERVICE returns a boolean if a field has been set.
func (o *GlobalProblemStatusOpenProblemCounts) HasSERVICE() bool {
	if o != nil && !IsNil(o.SERVICE) {
		return true
	}

	return false
}

// SetSERVICE gets a reference to the given int32 and assigns it to the SERVICE field.
func (o *GlobalProblemStatusOpenProblemCounts) SetSERVICE(v int32) {
	o.SERVICE = &v
}

func (o GlobalProblemStatusOpenProblemCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalProblemStatusOpenProblemCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.APPLICATION) {
		toSerialize["APPLICATION"] = o.APPLICATION
	}
	if !IsNil(o.ENVIRONMENT) {
		toSerialize["ENVIRONMENT"] = o.ENVIRONMENT
	}
	if !IsNil(o.INFRASTRUCTURE) {
		toSerialize["INFRASTRUCTURE"] = o.INFRASTRUCTURE
	}
	if !IsNil(o.SERVICE) {
		toSerialize["SERVICE"] = o.SERVICE
	}
	return toSerialize, nil
}

type NullableGlobalProblemStatusOpenProblemCounts struct {
	value *GlobalProblemStatusOpenProblemCounts
	isSet bool
}

func (v NullableGlobalProblemStatusOpenProblemCounts) Get() *GlobalProblemStatusOpenProblemCounts {
	return v.value
}

func (v *NullableGlobalProblemStatusOpenProblemCounts) Set(val *GlobalProblemStatusOpenProblemCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalProblemStatusOpenProblemCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalProblemStatusOpenProblemCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalProblemStatusOpenProblemCounts(val *GlobalProblemStatusOpenProblemCounts) *NullableGlobalProblemStatusOpenProblemCounts {
	return &NullableGlobalProblemStatusOpenProblemCounts{value: val, isSet: true}
}

func (v NullableGlobalProblemStatusOpenProblemCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalProblemStatusOpenProblemCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


