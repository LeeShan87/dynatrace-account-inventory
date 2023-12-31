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

// checks if the ExtensionListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionListDto{}

// ExtensionListDto struct for ExtensionListDto
type ExtensionListDto struct {
	// A list of extensions.
	Extensions []ExtensionDto `json:"extensions,omitempty"`
	// The cursor for the next page of results. Has the value of `null` on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result.
	NextPageKey *string `json:"nextPageKey,omitempty"`
	// The total number of entries in the result.
	TotalResults *int32 `json:"totalResults,omitempty"`
}

// NewExtensionListDto instantiates a new ExtensionListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionListDto() *ExtensionListDto {
	this := ExtensionListDto{}
	return &this
}

// NewExtensionListDtoWithDefaults instantiates a new ExtensionListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionListDtoWithDefaults() *ExtensionListDto {
	this := ExtensionListDto{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *ExtensionListDto) GetExtensions() []ExtensionDto {
	if o == nil || IsNil(o.Extensions) {
		var ret []ExtensionDto
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionListDto) GetExtensionsOk() ([]ExtensionDto, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *ExtensionListDto) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []ExtensionDto and assigns it to the Extensions field.
func (o *ExtensionListDto) SetExtensions(v []ExtensionDto) {
	o.Extensions = v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *ExtensionListDto) GetNextPageKey() string {
	if o == nil || IsNil(o.NextPageKey) {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionListDto) GetNextPageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageKey) {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *ExtensionListDto) HasNextPageKey() bool {
	if o != nil && !IsNil(o.NextPageKey) {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *ExtensionListDto) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ExtensionListDto) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionListDto) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ExtensionListDto) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ExtensionListDto) SetTotalResults(v int32) {
	o.TotalResults = &v
}

func (o ExtensionListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	if !IsNil(o.NextPageKey) {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	return toSerialize, nil
}

type NullableExtensionListDto struct {
	value *ExtensionListDto
	isSet bool
}

func (v NullableExtensionListDto) Get() *ExtensionListDto {
	return v.value
}

func (v *NullableExtensionListDto) Set(val *ExtensionListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionListDto(val *ExtensionListDto) *NullableExtensionListDto {
	return &NullableExtensionListDto{value: val, isSet: true}
}

func (v NullableExtensionListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


