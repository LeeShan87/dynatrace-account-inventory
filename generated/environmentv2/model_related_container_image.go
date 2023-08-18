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

// checks if the RelatedContainerImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedContainerImage{}

// RelatedContainerImage Related container image of a security problem.
type RelatedContainerImage struct {
	// A list of affected entities.
	AffectedEntities []string `json:"affectedEntities,omitempty"`
	// The image ID of the related container image.
	ImageId *string `json:"imageId,omitempty"`
	// The image name of the related container image.
	ImageName *string `json:"imageName,omitempty"`
	// The number of affected entities.
	NumberOfAffectedEntities *int32 `json:"numberOfAffectedEntities,omitempty"`
}

// NewRelatedContainerImage instantiates a new RelatedContainerImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedContainerImage() *RelatedContainerImage {
	this := RelatedContainerImage{}
	return &this
}

// NewRelatedContainerImageWithDefaults instantiates a new RelatedContainerImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedContainerImageWithDefaults() *RelatedContainerImage {
	this := RelatedContainerImage{}
	return &this
}

// GetAffectedEntities returns the AffectedEntities field value if set, zero value otherwise.
func (o *RelatedContainerImage) GetAffectedEntities() []string {
	if o == nil || IsNil(o.AffectedEntities) {
		var ret []string
		return ret
	}
	return o.AffectedEntities
}

// GetAffectedEntitiesOk returns a tuple with the AffectedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedContainerImage) GetAffectedEntitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.AffectedEntities) {
		return nil, false
	}
	return o.AffectedEntities, true
}

// HasAffectedEntities returns a boolean if a field has been set.
func (o *RelatedContainerImage) HasAffectedEntities() bool {
	if o != nil && !IsNil(o.AffectedEntities) {
		return true
	}

	return false
}

// SetAffectedEntities gets a reference to the given []string and assigns it to the AffectedEntities field.
func (o *RelatedContainerImage) SetAffectedEntities(v []string) {
	o.AffectedEntities = v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *RelatedContainerImage) GetImageId() string {
	if o == nil || IsNil(o.ImageId) {
		var ret string
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedContainerImage) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *RelatedContainerImage) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *RelatedContainerImage) SetImageId(v string) {
	o.ImageId = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *RelatedContainerImage) GetImageName() string {
	if o == nil || IsNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedContainerImage) GetImageNameOk() (*string, bool) {
	if o == nil || IsNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *RelatedContainerImage) HasImageName() bool {
	if o != nil && !IsNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *RelatedContainerImage) SetImageName(v string) {
	o.ImageName = &v
}

// GetNumberOfAffectedEntities returns the NumberOfAffectedEntities field value if set, zero value otherwise.
func (o *RelatedContainerImage) GetNumberOfAffectedEntities() int32 {
	if o == nil || IsNil(o.NumberOfAffectedEntities) {
		var ret int32
		return ret
	}
	return *o.NumberOfAffectedEntities
}

// GetNumberOfAffectedEntitiesOk returns a tuple with the NumberOfAffectedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedContainerImage) GetNumberOfAffectedEntitiesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfAffectedEntities) {
		return nil, false
	}
	return o.NumberOfAffectedEntities, true
}

// HasNumberOfAffectedEntities returns a boolean if a field has been set.
func (o *RelatedContainerImage) HasNumberOfAffectedEntities() bool {
	if o != nil && !IsNil(o.NumberOfAffectedEntities) {
		return true
	}

	return false
}

// SetNumberOfAffectedEntities gets a reference to the given int32 and assigns it to the NumberOfAffectedEntities field.
func (o *RelatedContainerImage) SetNumberOfAffectedEntities(v int32) {
	o.NumberOfAffectedEntities = &v
}

func (o RelatedContainerImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedContainerImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedEntities) {
		toSerialize["affectedEntities"] = o.AffectedEntities
	}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if !IsNil(o.ImageName) {
		toSerialize["imageName"] = o.ImageName
	}
	if !IsNil(o.NumberOfAffectedEntities) {
		toSerialize["numberOfAffectedEntities"] = o.NumberOfAffectedEntities
	}
	return toSerialize, nil
}

type NullableRelatedContainerImage struct {
	value *RelatedContainerImage
	isSet bool
}

func (v NullableRelatedContainerImage) Get() *RelatedContainerImage {
	return v.value
}

func (v *NullableRelatedContainerImage) Set(val *RelatedContainerImage) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedContainerImage) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedContainerImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedContainerImage(val *RelatedContainerImage) *NullableRelatedContainerImage {
	return &NullableRelatedContainerImage{value: val, isSet: true}
}

func (v NullableRelatedContainerImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedContainerImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


