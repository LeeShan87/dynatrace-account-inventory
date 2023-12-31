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

// checks if the RemoteConfigurationManagementPreviewList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteConfigurationManagementPreviewList{}

// RemoteConfigurationManagementPreviewList A list of remote configuration management jobs previews.
type RemoteConfigurationManagementPreviewList struct {
	// A list of remote configuration management jobs previews.
	Previews []RemoteConfigurationManagementJobPreview `json:"previews,omitempty"`
}

// NewRemoteConfigurationManagementPreviewList instantiates a new RemoteConfigurationManagementPreviewList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteConfigurationManagementPreviewList() *RemoteConfigurationManagementPreviewList {
	this := RemoteConfigurationManagementPreviewList{}
	return &this
}

// NewRemoteConfigurationManagementPreviewListWithDefaults instantiates a new RemoteConfigurationManagementPreviewList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteConfigurationManagementPreviewListWithDefaults() *RemoteConfigurationManagementPreviewList {
	this := RemoteConfigurationManagementPreviewList{}
	return &this
}

// GetPreviews returns the Previews field value if set, zero value otherwise.
func (o *RemoteConfigurationManagementPreviewList) GetPreviews() []RemoteConfigurationManagementJobPreview {
	if o == nil || IsNil(o.Previews) {
		var ret []RemoteConfigurationManagementJobPreview
		return ret
	}
	return o.Previews
}

// GetPreviewsOk returns a tuple with the Previews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteConfigurationManagementPreviewList) GetPreviewsOk() ([]RemoteConfigurationManagementJobPreview, bool) {
	if o == nil || IsNil(o.Previews) {
		return nil, false
	}
	return o.Previews, true
}

// HasPreviews returns a boolean if a field has been set.
func (o *RemoteConfigurationManagementPreviewList) HasPreviews() bool {
	if o != nil && !IsNil(o.Previews) {
		return true
	}

	return false
}

// SetPreviews gets a reference to the given []RemoteConfigurationManagementJobPreview and assigns it to the Previews field.
func (o *RemoteConfigurationManagementPreviewList) SetPreviews(v []RemoteConfigurationManagementJobPreview) {
	o.Previews = v
}

func (o RemoteConfigurationManagementPreviewList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteConfigurationManagementPreviewList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Previews) {
		toSerialize["previews"] = o.Previews
	}
	return toSerialize, nil
}

type NullableRemoteConfigurationManagementPreviewList struct {
	value *RemoteConfigurationManagementPreviewList
	isSet bool
}

func (v NullableRemoteConfigurationManagementPreviewList) Get() *RemoteConfigurationManagementPreviewList {
	return v.value
}

func (v *NullableRemoteConfigurationManagementPreviewList) Set(val *RemoteConfigurationManagementPreviewList) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteConfigurationManagementPreviewList) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteConfigurationManagementPreviewList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteConfigurationManagementPreviewList(val *RemoteConfigurationManagementPreviewList) *NullableRemoteConfigurationManagementPreviewList {
	return &NullableRemoteConfigurationManagementPreviewList{value: val, isSet: true}
}

func (v NullableRemoteConfigurationManagementPreviewList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteConfigurationManagementPreviewList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


