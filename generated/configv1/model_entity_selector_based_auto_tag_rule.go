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

// checks if the EntitySelectorBasedAutoTagRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitySelectorBasedAutoTagRule{}

// EntitySelectorBasedAutoTagRule The entity-selector-based rule for auto tag usage. It allows tagging entities via an entity selector.
type EntitySelectorBasedAutoTagRule struct {
	// The rule is enabled (`true`) or disabled (`false`).
	Enabled *bool `json:"enabled,omitempty"`
	// The entity selector string, by which the entities are selected.
	EntitySelector string `json:"entitySelector"`
	// Changes applied to the value after applying the value format. Default is LEAVE_TEXT_AS_IS.
	Normalization NullableString `json:"normalization,omitempty"`
	// The value of the entity-selector-based auto-tag. If specified, the tag is used in the `name:valueFormat` format.   For example, you can extend the `Infrastructure` tag to `Infrastructure:Windows` and `Infrastructure:Linux`.   
	ValueFormat *string `json:"valueFormat,omitempty"`
}

// NewEntitySelectorBasedAutoTagRule instantiates a new EntitySelectorBasedAutoTagRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitySelectorBasedAutoTagRule(entitySelector string) *EntitySelectorBasedAutoTagRule {
	this := EntitySelectorBasedAutoTagRule{}
	this.EntitySelector = entitySelector
	return &this
}

// NewEntitySelectorBasedAutoTagRuleWithDefaults instantiates a new EntitySelectorBasedAutoTagRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitySelectorBasedAutoTagRuleWithDefaults() *EntitySelectorBasedAutoTagRule {
	this := EntitySelectorBasedAutoTagRule{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *EntitySelectorBasedAutoTagRule) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySelectorBasedAutoTagRule) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *EntitySelectorBasedAutoTagRule) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *EntitySelectorBasedAutoTagRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEntitySelector returns the EntitySelector field value
func (o *EntitySelectorBasedAutoTagRule) GetEntitySelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitySelector
}

// GetEntitySelectorOk returns a tuple with the EntitySelector field value
// and a boolean to check if the value has been set.
func (o *EntitySelectorBasedAutoTagRule) GetEntitySelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitySelector, true
}

// SetEntitySelector sets field value
func (o *EntitySelectorBasedAutoTagRule) SetEntitySelector(v string) {
	o.EntitySelector = v
}

// GetNormalization returns the Normalization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EntitySelectorBasedAutoTagRule) GetNormalization() string {
	if o == nil || IsNil(o.Normalization.Get()) {
		var ret string
		return ret
	}
	return *o.Normalization.Get()
}

// GetNormalizationOk returns a tuple with the Normalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EntitySelectorBasedAutoTagRule) GetNormalizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Normalization.Get(), o.Normalization.IsSet()
}

// HasNormalization returns a boolean if a field has been set.
func (o *EntitySelectorBasedAutoTagRule) HasNormalization() bool {
	if o != nil && o.Normalization.IsSet() {
		return true
	}

	return false
}

// SetNormalization gets a reference to the given NullableString and assigns it to the Normalization field.
func (o *EntitySelectorBasedAutoTagRule) SetNormalization(v string) {
	o.Normalization.Set(&v)
}
// SetNormalizationNil sets the value for Normalization to be an explicit nil
func (o *EntitySelectorBasedAutoTagRule) SetNormalizationNil() {
	o.Normalization.Set(nil)
}

// UnsetNormalization ensures that no value is present for Normalization, not even an explicit nil
func (o *EntitySelectorBasedAutoTagRule) UnsetNormalization() {
	o.Normalization.Unset()
}

// GetValueFormat returns the ValueFormat field value if set, zero value otherwise.
func (o *EntitySelectorBasedAutoTagRule) GetValueFormat() string {
	if o == nil || IsNil(o.ValueFormat) {
		var ret string
		return ret
	}
	return *o.ValueFormat
}

// GetValueFormatOk returns a tuple with the ValueFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitySelectorBasedAutoTagRule) GetValueFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ValueFormat) {
		return nil, false
	}
	return o.ValueFormat, true
}

// HasValueFormat returns a boolean if a field has been set.
func (o *EntitySelectorBasedAutoTagRule) HasValueFormat() bool {
	if o != nil && !IsNil(o.ValueFormat) {
		return true
	}

	return false
}

// SetValueFormat gets a reference to the given string and assigns it to the ValueFormat field.
func (o *EntitySelectorBasedAutoTagRule) SetValueFormat(v string) {
	o.ValueFormat = &v
}

func (o EntitySelectorBasedAutoTagRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitySelectorBasedAutoTagRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["entitySelector"] = o.EntitySelector
	if o.Normalization.IsSet() {
		toSerialize["normalization"] = o.Normalization.Get()
	}
	if !IsNil(o.ValueFormat) {
		toSerialize["valueFormat"] = o.ValueFormat
	}
	return toSerialize, nil
}

type NullableEntitySelectorBasedAutoTagRule struct {
	value *EntitySelectorBasedAutoTagRule
	isSet bool
}

func (v NullableEntitySelectorBasedAutoTagRule) Get() *EntitySelectorBasedAutoTagRule {
	return v.value
}

func (v *NullableEntitySelectorBasedAutoTagRule) Set(val *EntitySelectorBasedAutoTagRule) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitySelectorBasedAutoTagRule) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitySelectorBasedAutoTagRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitySelectorBasedAutoTagRule(val *EntitySelectorBasedAutoTagRule) *NullableEntitySelectorBasedAutoTagRule {
	return &NullableEntitySelectorBasedAutoTagRule{value: val, isSet: true}
}

func (v NullableEntitySelectorBasedAutoTagRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitySelectorBasedAutoTagRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

