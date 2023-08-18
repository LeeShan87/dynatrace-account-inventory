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

// checks if the ActiveGateTokenCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGateTokenCreate{}

// ActiveGateTokenCreate Parameters of a new ActiveGate token.
type ActiveGateTokenCreate struct {
	// The type of the ActiveGate for which the token is valid.
	ActiveGateType string `json:"activeGateType"`
	// The expiration date of the token.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the token never expires.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// The name of the token.
	Name string `json:"name"`
	// The token is a seed token (`true`) or an individual token (`false`).    We recommend the individual token option (false).
	SeedToken *bool `json:"seedToken,omitempty"`
}

// NewActiveGateTokenCreate instantiates a new ActiveGateTokenCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateTokenCreate(activeGateType string, name string) *ActiveGateTokenCreate {
	this := ActiveGateTokenCreate{}
	this.ActiveGateType = activeGateType
	this.Name = name
	return &this
}

// NewActiveGateTokenCreateWithDefaults instantiates a new ActiveGateTokenCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateTokenCreateWithDefaults() *ActiveGateTokenCreate {
	this := ActiveGateTokenCreate{}
	return &this
}

// GetActiveGateType returns the ActiveGateType field value
func (o *ActiveGateTokenCreate) GetActiveGateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveGateType
}

// GetActiveGateTypeOk returns a tuple with the ActiveGateType field value
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreate) GetActiveGateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveGateType, true
}

// SetActiveGateType sets field value
func (o *ActiveGateTokenCreate) SetActiveGateType(v string) {
	o.ActiveGateType = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ActiveGateTokenCreate) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreate) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ActiveGateTokenCreate) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ActiveGateTokenCreate) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetName returns the Name field value
func (o *ActiveGateTokenCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActiveGateTokenCreate) SetName(v string) {
	o.Name = v
}

// GetSeedToken returns the SeedToken field value if set, zero value otherwise.
func (o *ActiveGateTokenCreate) GetSeedToken() bool {
	if o == nil || IsNil(o.SeedToken) {
		var ret bool
		return ret
	}
	return *o.SeedToken
}

// GetSeedTokenOk returns a tuple with the SeedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreate) GetSeedTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.SeedToken) {
		return nil, false
	}
	return o.SeedToken, true
}

// HasSeedToken returns a boolean if a field has been set.
func (o *ActiveGateTokenCreate) HasSeedToken() bool {
	if o != nil && !IsNil(o.SeedToken) {
		return true
	}

	return false
}

// SetSeedToken gets a reference to the given bool and assigns it to the SeedToken field.
func (o *ActiveGateTokenCreate) SetSeedToken(v bool) {
	o.SeedToken = &v
}

func (o ActiveGateTokenCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGateTokenCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeGateType"] = o.ActiveGateType
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.SeedToken) {
		toSerialize["seedToken"] = o.SeedToken
	}
	return toSerialize, nil
}

type NullableActiveGateTokenCreate struct {
	value *ActiveGateTokenCreate
	isSet bool
}

func (v NullableActiveGateTokenCreate) Get() *ActiveGateTokenCreate {
	return v.value
}

func (v *NullableActiveGateTokenCreate) Set(val *ActiveGateTokenCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateTokenCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateTokenCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateTokenCreate(val *ActiveGateTokenCreate) *NullableActiveGateTokenCreate {
	return &NullableActiveGateTokenCreate{value: val, isSet: true}
}

func (v NullableActiveGateTokenCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateTokenCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

