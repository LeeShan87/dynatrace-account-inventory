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

// checks if the AssetInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetInfoDto{}

// AssetInfoDto Metadata for an extension asset.
type AssetInfoDto struct {
	AssetSchemaDetails *AssetSchemaDetailsDto `json:"assetSchemaDetails,omitempty"`
	// User-friendly name of the asset.
	DisplayName *string `json:"displayName,omitempty"`
	// ID of the asset. Identifies the asset in REST API and/or UI (where applicable).
	Id *string `json:"id,omitempty"`
	// The type of the asset.
	Type *string `json:"type,omitempty"`
}

// NewAssetInfoDto instantiates a new AssetInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetInfoDto() *AssetInfoDto {
	this := AssetInfoDto{}
	return &this
}

// NewAssetInfoDtoWithDefaults instantiates a new AssetInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetInfoDtoWithDefaults() *AssetInfoDto {
	this := AssetInfoDto{}
	return &this
}

// GetAssetSchemaDetails returns the AssetSchemaDetails field value if set, zero value otherwise.
func (o *AssetInfoDto) GetAssetSchemaDetails() AssetSchemaDetailsDto {
	if o == nil || IsNil(o.AssetSchemaDetails) {
		var ret AssetSchemaDetailsDto
		return ret
	}
	return *o.AssetSchemaDetails
}

// GetAssetSchemaDetailsOk returns a tuple with the AssetSchemaDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfoDto) GetAssetSchemaDetailsOk() (*AssetSchemaDetailsDto, bool) {
	if o == nil || IsNil(o.AssetSchemaDetails) {
		return nil, false
	}
	return o.AssetSchemaDetails, true
}

// HasAssetSchemaDetails returns a boolean if a field has been set.
func (o *AssetInfoDto) HasAssetSchemaDetails() bool {
	if o != nil && !IsNil(o.AssetSchemaDetails) {
		return true
	}

	return false
}

// SetAssetSchemaDetails gets a reference to the given AssetSchemaDetailsDto and assigns it to the AssetSchemaDetails field.
func (o *AssetInfoDto) SetAssetSchemaDetails(v AssetSchemaDetailsDto) {
	o.AssetSchemaDetails = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *AssetInfoDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfoDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *AssetInfoDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *AssetInfoDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetInfoDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfoDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetInfoDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetInfoDto) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssetInfoDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetInfoDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssetInfoDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssetInfoDto) SetType(v string) {
	o.Type = &v
}

func (o AssetInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetSchemaDetails) {
		toSerialize["assetSchemaDetails"] = o.AssetSchemaDetails
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableAssetInfoDto struct {
	value *AssetInfoDto
	isSet bool
}

func (v NullableAssetInfoDto) Get() *AssetInfoDto {
	return v.value
}

func (v *NullableAssetInfoDto) Set(val *AssetInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetInfoDto(val *AssetInfoDto) *NullableAssetInfoDto {
	return &NullableAssetInfoDto{value: val, isSet: true}
}

func (v NullableAssetInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


