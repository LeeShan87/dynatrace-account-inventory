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

// checks if the ExtensionAssetsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionAssetsDto{}

// ExtensionAssetsDto List of assets imported with the active extension environment configuration.
type ExtensionAssetsDto struct {
	// The list of the imported assets.
	Assets []AssetInfoDto `json:"assets,omitempty"`
	// List of errors during asset import
	Errors []string `json:"errors,omitempty"`
	// The status of the assets list.
	Status *string `json:"status,omitempty"`
	// Extension version
	Version *string `json:"version,omitempty"`
}

// NewExtensionAssetsDto instantiates a new ExtensionAssetsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionAssetsDto() *ExtensionAssetsDto {
	this := ExtensionAssetsDto{}
	return &this
}

// NewExtensionAssetsDtoWithDefaults instantiates a new ExtensionAssetsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionAssetsDtoWithDefaults() *ExtensionAssetsDto {
	this := ExtensionAssetsDto{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *ExtensionAssetsDto) GetAssets() []AssetInfoDto {
	if o == nil || IsNil(o.Assets) {
		var ret []AssetInfoDto
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAssetsDto) GetAssetsOk() ([]AssetInfoDto, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *ExtensionAssetsDto) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []AssetInfoDto and assigns it to the Assets field.
func (o *ExtensionAssetsDto) SetAssets(v []AssetInfoDto) {
	o.Assets = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ExtensionAssetsDto) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAssetsDto) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ExtensionAssetsDto) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *ExtensionAssetsDto) SetErrors(v []string) {
	o.Errors = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtensionAssetsDto) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAssetsDto) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtensionAssetsDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExtensionAssetsDto) SetStatus(v string) {
	o.Status = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ExtensionAssetsDto) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionAssetsDto) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ExtensionAssetsDto) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ExtensionAssetsDto) SetVersion(v string) {
	o.Version = &v
}

func (o ExtensionAssetsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionAssetsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableExtensionAssetsDto struct {
	value *ExtensionAssetsDto
	isSet bool
}

func (v NullableExtensionAssetsDto) Get() *ExtensionAssetsDto {
	return v.value
}

func (v *NullableExtensionAssetsDto) Set(val *ExtensionAssetsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionAssetsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionAssetsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionAssetsDto(val *ExtensionAssetsDto) *NullableExtensionAssetsDto {
	return &NullableExtensionAssetsDto{value: val, isSet: true}
}

func (v NullableExtensionAssetsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionAssetsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

