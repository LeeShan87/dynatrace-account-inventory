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

// checks if the ResourceUrlCleanupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceUrlCleanupRule{}

// ResourceUrlCleanupRule A rule for the URL cleanup rule.
type ResourceUrlCleanupRule struct {
	// The pattern (regular expression) to look for.
	RegularExpression string `json:"regularExpression"`
	// The text to replace the found pattern with.
	ReplaceWith string `json:"replaceWith"`
	// The name of the rule.
	ResourceName string `json:"resourceName"`
}

// NewResourceUrlCleanupRule instantiates a new ResourceUrlCleanupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceUrlCleanupRule(regularExpression string, replaceWith string, resourceName string) *ResourceUrlCleanupRule {
	this := ResourceUrlCleanupRule{}
	this.RegularExpression = regularExpression
	this.ReplaceWith = replaceWith
	this.ResourceName = resourceName
	return &this
}

// NewResourceUrlCleanupRuleWithDefaults instantiates a new ResourceUrlCleanupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceUrlCleanupRuleWithDefaults() *ResourceUrlCleanupRule {
	this := ResourceUrlCleanupRule{}
	return &this
}

// GetRegularExpression returns the RegularExpression field value
func (o *ResourceUrlCleanupRule) GetRegularExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegularExpression
}

// GetRegularExpressionOk returns a tuple with the RegularExpression field value
// and a boolean to check if the value has been set.
func (o *ResourceUrlCleanupRule) GetRegularExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegularExpression, true
}

// SetRegularExpression sets field value
func (o *ResourceUrlCleanupRule) SetRegularExpression(v string) {
	o.RegularExpression = v
}

// GetReplaceWith returns the ReplaceWith field value
func (o *ResourceUrlCleanupRule) GetReplaceWith() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReplaceWith
}

// GetReplaceWithOk returns a tuple with the ReplaceWith field value
// and a boolean to check if the value has been set.
func (o *ResourceUrlCleanupRule) GetReplaceWithOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplaceWith, true
}

// SetReplaceWith sets field value
func (o *ResourceUrlCleanupRule) SetReplaceWith(v string) {
	o.ReplaceWith = v
}

// GetResourceName returns the ResourceName field value
func (o *ResourceUrlCleanupRule) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *ResourceUrlCleanupRule) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *ResourceUrlCleanupRule) SetResourceName(v string) {
	o.ResourceName = v
}

func (o ResourceUrlCleanupRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceUrlCleanupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["regularExpression"] = o.RegularExpression
	toSerialize["replaceWith"] = o.ReplaceWith
	toSerialize["resourceName"] = o.ResourceName
	return toSerialize, nil
}

type NullableResourceUrlCleanupRule struct {
	value *ResourceUrlCleanupRule
	isSet bool
}

func (v NullableResourceUrlCleanupRule) Get() *ResourceUrlCleanupRule {
	return v.value
}

func (v *NullableResourceUrlCleanupRule) Set(val *ResourceUrlCleanupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceUrlCleanupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceUrlCleanupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceUrlCleanupRule(val *ResourceUrlCleanupRule) *NullableResourceUrlCleanupRule {
	return &NullableResourceUrlCleanupRule{value: val, isSet: true}
}

func (v NullableResourceUrlCleanupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceUrlCleanupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


