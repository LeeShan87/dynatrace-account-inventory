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

// checks if the ExtractFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtractFields{}

// ExtractFields A query to extract log content to a custom column.
type ExtractFields struct {
	// The query for content extraction.   See [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) in Dynatrace Documentation for the syntax definition and examples.
	CustomParsingPatterns *string `json:"customParsingPatterns,omitempty"`
	// The parsing mode for log analysis entries presentation. Possible values are: `json`, `disabled`, and `all`.
	ParsingMode *string `json:"parsingMode,omitempty"`
}

// NewExtractFields instantiates a new ExtractFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractFields() *ExtractFields {
	this := ExtractFields{}
	return &this
}

// NewExtractFieldsWithDefaults instantiates a new ExtractFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractFieldsWithDefaults() *ExtractFields {
	this := ExtractFields{}
	return &this
}

// GetCustomParsingPatterns returns the CustomParsingPatterns field value if set, zero value otherwise.
func (o *ExtractFields) GetCustomParsingPatterns() string {
	if o == nil || IsNil(o.CustomParsingPatterns) {
		var ret string
		return ret
	}
	return *o.CustomParsingPatterns
}

// GetCustomParsingPatternsOk returns a tuple with the CustomParsingPatterns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractFields) GetCustomParsingPatternsOk() (*string, bool) {
	if o == nil || IsNil(o.CustomParsingPatterns) {
		return nil, false
	}
	return o.CustomParsingPatterns, true
}

// HasCustomParsingPatterns returns a boolean if a field has been set.
func (o *ExtractFields) HasCustomParsingPatterns() bool {
	if o != nil && !IsNil(o.CustomParsingPatterns) {
		return true
	}

	return false
}

// SetCustomParsingPatterns gets a reference to the given string and assigns it to the CustomParsingPatterns field.
func (o *ExtractFields) SetCustomParsingPatterns(v string) {
	o.CustomParsingPatterns = &v
}

// GetParsingMode returns the ParsingMode field value if set, zero value otherwise.
func (o *ExtractFields) GetParsingMode() string {
	if o == nil || IsNil(o.ParsingMode) {
		var ret string
		return ret
	}
	return *o.ParsingMode
}

// GetParsingModeOk returns a tuple with the ParsingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtractFields) GetParsingModeOk() (*string, bool) {
	if o == nil || IsNil(o.ParsingMode) {
		return nil, false
	}
	return o.ParsingMode, true
}

// HasParsingMode returns a boolean if a field has been set.
func (o *ExtractFields) HasParsingMode() bool {
	if o != nil && !IsNil(o.ParsingMode) {
		return true
	}

	return false
}

// SetParsingMode gets a reference to the given string and assigns it to the ParsingMode field.
func (o *ExtractFields) SetParsingMode(v string) {
	o.ParsingMode = &v
}

func (o ExtractFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtractFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomParsingPatterns) {
		toSerialize["customParsingPatterns"] = o.CustomParsingPatterns
	}
	if !IsNil(o.ParsingMode) {
		toSerialize["parsingMode"] = o.ParsingMode
	}
	return toSerialize, nil
}

type NullableExtractFields struct {
	value *ExtractFields
	isSet bool
}

func (v NullableExtractFields) Get() *ExtractFields {
	return v.value
}

func (v *NullableExtractFields) Set(val *ExtractFields) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractFields) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractFields(val *ExtractFields) *NullableExtractFields {
	return &NullableExtractFields{value: val, isSet: true}
}

func (v NullableExtractFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


