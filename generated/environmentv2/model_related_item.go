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

// checks if the RelatedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedItem{}

// RelatedItem Related items.
type RelatedItem struct {
	Description *string `json:"description,omitempty"`
	// External link (marketing/documentation) that can provide with additional information.
	ExternalLink *string `json:"externalLink,omitempty"`
	// Indicates whether there is a page within the product to activate this item.
	HasClusterLink *bool `json:"hasClusterLink,omitempty"`
	// The logo of the item. Can be a URL or Base64 encoded. Intended for <image> html tags
	IconUrl *string `json:"iconUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Represents the type of item. It can be TECHNOLOGY, EXTENSION1 or EXTENSION2.
	Type *string `json:"type,omitempty"`
}

// NewRelatedItem instantiates a new RelatedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedItem() *RelatedItem {
	this := RelatedItem{}
	return &this
}

// NewRelatedItemWithDefaults instantiates a new RelatedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedItemWithDefaults() *RelatedItem {
	this := RelatedItem{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RelatedItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RelatedItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RelatedItem) SetDescription(v string) {
	o.Description = &v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *RelatedItem) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *RelatedItem) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *RelatedItem) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetHasClusterLink returns the HasClusterLink field value if set, zero value otherwise.
func (o *RelatedItem) GetHasClusterLink() bool {
	if o == nil || IsNil(o.HasClusterLink) {
		var ret bool
		return ret
	}
	return *o.HasClusterLink
}

// GetHasClusterLinkOk returns a tuple with the HasClusterLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetHasClusterLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.HasClusterLink) {
		return nil, false
	}
	return o.HasClusterLink, true
}

// HasHasClusterLink returns a boolean if a field has been set.
func (o *RelatedItem) HasHasClusterLink() bool {
	if o != nil && !IsNil(o.HasClusterLink) {
		return true
	}

	return false
}

// SetHasClusterLink gets a reference to the given bool and assigns it to the HasClusterLink field.
func (o *RelatedItem) SetHasClusterLink(v bool) {
	o.HasClusterLink = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *RelatedItem) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetIconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *RelatedItem) HasIconUrl() bool {
	if o != nil && !IsNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *RelatedItem) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelatedItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelatedItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelatedItem) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RelatedItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RelatedItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RelatedItem) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelatedItem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelatedItem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelatedItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RelatedItem) SetType(v string) {
	o.Type = &v
}

func (o RelatedItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	if !IsNil(o.HasClusterLink) {
		toSerialize["hasClusterLink"] = o.HasClusterLink
	}
	if !IsNil(o.IconUrl) {
		toSerialize["iconUrl"] = o.IconUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRelatedItem struct {
	value *RelatedItem
	isSet bool
}

func (v NullableRelatedItem) Get() *RelatedItem {
	return v.value
}

func (v *NullableRelatedItem) Set(val *RelatedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedItem(val *RelatedItem) *NullableRelatedItem {
	return &NullableRelatedItem{value: val, isSet: true}
}

func (v NullableRelatedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


