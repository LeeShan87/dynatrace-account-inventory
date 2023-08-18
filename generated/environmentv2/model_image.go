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

// checks if the Image type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Image{}

// Image Information about the image details of a capability.
type Image struct {
	// Alternate text for the image.
	Alt *string `json:"alt,omitempty"`
	// Url of the image.
	Src *string `json:"src,omitempty"`
	// Title of the image.
	Title *string `json:"title,omitempty"`
}

// NewImage instantiates a new Image object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImage() *Image {
	this := Image{}
	return &this
}

// NewImageWithDefaults instantiates a new Image object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageWithDefaults() *Image {
	this := Image{}
	return &this
}

// GetAlt returns the Alt field value if set, zero value otherwise.
func (o *Image) GetAlt() string {
	if o == nil || IsNil(o.Alt) {
		var ret string
		return ret
	}
	return *o.Alt
}

// GetAltOk returns a tuple with the Alt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetAltOk() (*string, bool) {
	if o == nil || IsNil(o.Alt) {
		return nil, false
	}
	return o.Alt, true
}

// HasAlt returns a boolean if a field has been set.
func (o *Image) HasAlt() bool {
	if o != nil && !IsNil(o.Alt) {
		return true
	}

	return false
}

// SetAlt gets a reference to the given string and assigns it to the Alt field.
func (o *Image) SetAlt(v string) {
	o.Alt = &v
}

// GetSrc returns the Src field value if set, zero value otherwise.
func (o *Image) GetSrc() string {
	if o == nil || IsNil(o.Src) {
		var ret string
		return ret
	}
	return *o.Src
}

// GetSrcOk returns a tuple with the Src field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetSrcOk() (*string, bool) {
	if o == nil || IsNil(o.Src) {
		return nil, false
	}
	return o.Src, true
}

// HasSrc returns a boolean if a field has been set.
func (o *Image) HasSrc() bool {
	if o != nil && !IsNil(o.Src) {
		return true
	}

	return false
}

// SetSrc gets a reference to the given string and assigns it to the Src field.
func (o *Image) SetSrc(v string) {
	o.Src = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Image) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Image) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Image) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Image) SetTitle(v string) {
	o.Title = &v
}

func (o Image) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Image) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alt) {
		toSerialize["alt"] = o.Alt
	}
	if !IsNil(o.Src) {
		toSerialize["src"] = o.Src
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableImage struct {
	value *Image
	isSet bool
}

func (v NullableImage) Get() *Image {
	return v.value
}

func (v *NullableImage) Set(val *Image) {
	v.value = val
	v.isSet = true
}

func (v NullableImage) IsSet() bool {
	return v.isSet
}

func (v *NullableImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImage(val *Image) *NullableImage {
	return &NullableImage{value: val, isSet: true}
}

func (v NullableImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

