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

// checks if the HttpRequestParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpRequestParameter{}

// HttpRequestParameter An HTTP request parameter.
type HttpRequestParameter struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty"`
	// The value of the parameter.
	Value *string `json:"value,omitempty"`
}

// NewHttpRequestParameter instantiates a new HttpRequestParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpRequestParameter() *HttpRequestParameter {
	this := HttpRequestParameter{}
	return &this
}

// NewHttpRequestParameterWithDefaults instantiates a new HttpRequestParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpRequestParameterWithDefaults() *HttpRequestParameter {
	this := HttpRequestParameter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HttpRequestParameter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequestParameter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HttpRequestParameter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HttpRequestParameter) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HttpRequestParameter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpRequestParameter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HttpRequestParameter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HttpRequestParameter) SetValue(v string) {
	o.Value = &v
}

func (o HttpRequestParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpRequestParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableHttpRequestParameter struct {
	value *HttpRequestParameter
	isSet bool
}

func (v NullableHttpRequestParameter) Get() *HttpRequestParameter {
	return v.value
}

func (v *NullableHttpRequestParameter) Set(val *HttpRequestParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpRequestParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpRequestParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpRequestParameter(val *HttpRequestParameter) *NullableHttpRequestParameter {
	return &NullableHttpRequestParameter{value: val, isSet: true}
}

func (v NullableHttpRequestParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpRequestParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


