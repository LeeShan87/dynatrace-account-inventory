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

// checks if the ContentResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentResources{}

// ContentResources The configuration of content resources.
type ContentResources struct {
	// An ordered list of manually added content providers.   Rules are evaluated from top to bottom; the first matching rules applies.
	ResourceProviders []ResourceProvider `json:"resourceProviders,omitempty"`
	// An ordered list of manually defined resource types.   Rules are evaluated from top to bottom; the first matching rules applies.
	ResourceTypes []ResourceType `json:"resourceTypes,omitempty"`
	// An ordered list of resource URL cleanup rules.   Rules are evaluated from top to bottom; the first matching rules applies.
	ResourceUrlCleanupRules []ResourceUrlCleanupRule `json:"resourceUrlCleanupRules,omitempty"`
}

// NewContentResources instantiates a new ContentResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentResources() *ContentResources {
	this := ContentResources{}
	return &this
}

// NewContentResourcesWithDefaults instantiates a new ContentResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentResourcesWithDefaults() *ContentResources {
	this := ContentResources{}
	return &this
}

// GetResourceProviders returns the ResourceProviders field value if set, zero value otherwise.
func (o *ContentResources) GetResourceProviders() []ResourceProvider {
	if o == nil || IsNil(o.ResourceProviders) {
		var ret []ResourceProvider
		return ret
	}
	return o.ResourceProviders
}

// GetResourceProvidersOk returns a tuple with the ResourceProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentResources) GetResourceProvidersOk() ([]ResourceProvider, bool) {
	if o == nil || IsNil(o.ResourceProviders) {
		return nil, false
	}
	return o.ResourceProviders, true
}

// HasResourceProviders returns a boolean if a field has been set.
func (o *ContentResources) HasResourceProviders() bool {
	if o != nil && !IsNil(o.ResourceProviders) {
		return true
	}

	return false
}

// SetResourceProviders gets a reference to the given []ResourceProvider and assigns it to the ResourceProviders field.
func (o *ContentResources) SetResourceProviders(v []ResourceProvider) {
	o.ResourceProviders = v
}

// GetResourceTypes returns the ResourceTypes field value if set, zero value otherwise.
func (o *ContentResources) GetResourceTypes() []ResourceType {
	if o == nil || IsNil(o.ResourceTypes) {
		var ret []ResourceType
		return ret
	}
	return o.ResourceTypes
}

// GetResourceTypesOk returns a tuple with the ResourceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentResources) GetResourceTypesOk() ([]ResourceType, bool) {
	if o == nil || IsNil(o.ResourceTypes) {
		return nil, false
	}
	return o.ResourceTypes, true
}

// HasResourceTypes returns a boolean if a field has been set.
func (o *ContentResources) HasResourceTypes() bool {
	if o != nil && !IsNil(o.ResourceTypes) {
		return true
	}

	return false
}

// SetResourceTypes gets a reference to the given []ResourceType and assigns it to the ResourceTypes field.
func (o *ContentResources) SetResourceTypes(v []ResourceType) {
	o.ResourceTypes = v
}

// GetResourceUrlCleanupRules returns the ResourceUrlCleanupRules field value if set, zero value otherwise.
func (o *ContentResources) GetResourceUrlCleanupRules() []ResourceUrlCleanupRule {
	if o == nil || IsNil(o.ResourceUrlCleanupRules) {
		var ret []ResourceUrlCleanupRule
		return ret
	}
	return o.ResourceUrlCleanupRules
}

// GetResourceUrlCleanupRulesOk returns a tuple with the ResourceUrlCleanupRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentResources) GetResourceUrlCleanupRulesOk() ([]ResourceUrlCleanupRule, bool) {
	if o == nil || IsNil(o.ResourceUrlCleanupRules) {
		return nil, false
	}
	return o.ResourceUrlCleanupRules, true
}

// HasResourceUrlCleanupRules returns a boolean if a field has been set.
func (o *ContentResources) HasResourceUrlCleanupRules() bool {
	if o != nil && !IsNil(o.ResourceUrlCleanupRules) {
		return true
	}

	return false
}

// SetResourceUrlCleanupRules gets a reference to the given []ResourceUrlCleanupRule and assigns it to the ResourceUrlCleanupRules field.
func (o *ContentResources) SetResourceUrlCleanupRules(v []ResourceUrlCleanupRule) {
	o.ResourceUrlCleanupRules = v
}

func (o ContentResources) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceProviders) {
		toSerialize["resourceProviders"] = o.ResourceProviders
	}
	if !IsNil(o.ResourceTypes) {
		toSerialize["resourceTypes"] = o.ResourceTypes
	}
	if !IsNil(o.ResourceUrlCleanupRules) {
		toSerialize["resourceUrlCleanupRules"] = o.ResourceUrlCleanupRules
	}
	return toSerialize, nil
}

type NullableContentResources struct {
	value *ContentResources
	isSet bool
}

func (v NullableContentResources) Get() *ContentResources {
	return v.value
}

func (v *NullableContentResources) Set(val *ContentResources) {
	v.value = val
	v.isSet = true
}

func (v NullableContentResources) IsSet() bool {
	return v.isSet
}

func (v *NullableContentResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentResources(val *ContentResources) *NullableContentResources {
	return &NullableContentResources{value: val, isSet: true}
}

func (v NullableContentResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


