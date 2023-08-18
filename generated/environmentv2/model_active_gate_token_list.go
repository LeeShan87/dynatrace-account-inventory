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

// checks if the ActiveGateTokenList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGateTokenList{}

// ActiveGateTokenList A list of ActiveGate tokens.
type ActiveGateTokenList struct {
	// A list of ActiveGate tokens.
	ActiveGateTokens *[]ActiveGateToken `json:"activeGateTokens,omitempty"`
	// The cursor for the next page of results. Has the value of `null` on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result.
	NextPageKey *string `json:"nextPageKey,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"pageSize,omitempty"`
	// The total number of entries in the result.
	TotalCount int64 `json:"totalCount"`
}

// NewActiveGateTokenList instantiates a new ActiveGateTokenList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateTokenList(totalCount int64) *ActiveGateTokenList {
	this := ActiveGateTokenList{}
	this.TotalCount = totalCount
	return &this
}

// NewActiveGateTokenListWithDefaults instantiates a new ActiveGateTokenList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateTokenListWithDefaults() *ActiveGateTokenList {
	this := ActiveGateTokenList{}
	return &this
}

// GetActiveGateTokens returns the ActiveGateTokens field value if set, zero value otherwise.
func (o *ActiveGateTokenList) GetActiveGateTokens() []ActiveGateToken {
	if o == nil || IsNil(o.ActiveGateTokens) {
		var ret []ActiveGateToken
		return ret
	}
	return *o.ActiveGateTokens
}

// GetActiveGateTokensOk returns a tuple with the ActiveGateTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenList) GetActiveGateTokensOk() (*[]ActiveGateToken, bool) {
	if o == nil || IsNil(o.ActiveGateTokens) {
		return nil, false
	}
	return o.ActiveGateTokens, true
}

// HasActiveGateTokens returns a boolean if a field has been set.
func (o *ActiveGateTokenList) HasActiveGateTokens() bool {
	if o != nil && !IsNil(o.ActiveGateTokens) {
		return true
	}

	return false
}

// SetActiveGateTokens gets a reference to the given []ActiveGateToken and assigns it to the ActiveGateTokens field.
func (o *ActiveGateTokenList) SetActiveGateTokens(v []ActiveGateToken) {
	o.ActiveGateTokens = &v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *ActiveGateTokenList) GetNextPageKey() string {
	if o == nil || IsNil(o.NextPageKey) {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenList) GetNextPageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageKey) {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *ActiveGateTokenList) HasNextPageKey() bool {
	if o != nil && !IsNil(o.NextPageKey) {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *ActiveGateTokenList) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ActiveGateTokenList) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenList) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ActiveGateTokenList) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ActiveGateTokenList) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value
func (o *ActiveGateTokenList) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenList) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ActiveGateTokenList) SetTotalCount(v int64) {
	o.TotalCount = v
}

func (o ActiveGateTokenList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGateTokenList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveGateTokens) {
		toSerialize["activeGateTokens"] = o.ActiveGateTokens
	}
	if !IsNil(o.NextPageKey) {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

type NullableActiveGateTokenList struct {
	value *ActiveGateTokenList
	isSet bool
}

func (v NullableActiveGateTokenList) Get() *ActiveGateTokenList {
	return v.value
}

func (v *NullableActiveGateTokenList) Set(val *ActiveGateTokenList) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateTokenList) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateTokenList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateTokenList(val *ActiveGateTokenList) *NullableActiveGateTokenList {
	return &NullableActiveGateTokenList{value: val, isSet: true}
}

func (v NullableActiveGateTokenList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateTokenList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

