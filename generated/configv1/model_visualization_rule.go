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

// checks if the VisualizationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizationRule{}

// VisualizationRule Rules for Visualization
type VisualizationRule struct {
	Matcher string `json:"matcher"`
	Properties VisualizationProperties `json:"properties"`
	SeriesOverrides []VisualizationSerieOverride `json:"seriesOverrides,omitempty"`
	UnitTransform *string `json:"unitTransform,omitempty"`
	ValueFormat *string `json:"valueFormat,omitempty"`
}

// NewVisualizationRule instantiates a new VisualizationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizationRule(matcher string, properties VisualizationProperties) *VisualizationRule {
	this := VisualizationRule{}
	this.Matcher = matcher
	this.Properties = properties
	return &this
}

// NewVisualizationRuleWithDefaults instantiates a new VisualizationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizationRuleWithDefaults() *VisualizationRule {
	this := VisualizationRule{}
	return &this
}

// GetMatcher returns the Matcher field value
func (o *VisualizationRule) GetMatcher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Matcher
}

// GetMatcherOk returns a tuple with the Matcher field value
// and a boolean to check if the value has been set.
func (o *VisualizationRule) GetMatcherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matcher, true
}

// SetMatcher sets field value
func (o *VisualizationRule) SetMatcher(v string) {
	o.Matcher = v
}

// GetProperties returns the Properties field value
func (o *VisualizationRule) GetProperties() VisualizationProperties {
	if o == nil {
		var ret VisualizationProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *VisualizationRule) GetPropertiesOk() (*VisualizationProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *VisualizationRule) SetProperties(v VisualizationProperties) {
	o.Properties = v
}

// GetSeriesOverrides returns the SeriesOverrides field value if set, zero value otherwise.
func (o *VisualizationRule) GetSeriesOverrides() []VisualizationSerieOverride {
	if o == nil || IsNil(o.SeriesOverrides) {
		var ret []VisualizationSerieOverride
		return ret
	}
	return o.SeriesOverrides
}

// GetSeriesOverridesOk returns a tuple with the SeriesOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizationRule) GetSeriesOverridesOk() ([]VisualizationSerieOverride, bool) {
	if o == nil || IsNil(o.SeriesOverrides) {
		return nil, false
	}
	return o.SeriesOverrides, true
}

// HasSeriesOverrides returns a boolean if a field has been set.
func (o *VisualizationRule) HasSeriesOverrides() bool {
	if o != nil && !IsNil(o.SeriesOverrides) {
		return true
	}

	return false
}

// SetSeriesOverrides gets a reference to the given []VisualizationSerieOverride and assigns it to the SeriesOverrides field.
func (o *VisualizationRule) SetSeriesOverrides(v []VisualizationSerieOverride) {
	o.SeriesOverrides = v
}

// GetUnitTransform returns the UnitTransform field value if set, zero value otherwise.
func (o *VisualizationRule) GetUnitTransform() string {
	if o == nil || IsNil(o.UnitTransform) {
		var ret string
		return ret
	}
	return *o.UnitTransform
}

// GetUnitTransformOk returns a tuple with the UnitTransform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizationRule) GetUnitTransformOk() (*string, bool) {
	if o == nil || IsNil(o.UnitTransform) {
		return nil, false
	}
	return o.UnitTransform, true
}

// HasUnitTransform returns a boolean if a field has been set.
func (o *VisualizationRule) HasUnitTransform() bool {
	if o != nil && !IsNil(o.UnitTransform) {
		return true
	}

	return false
}

// SetUnitTransform gets a reference to the given string and assigns it to the UnitTransform field.
func (o *VisualizationRule) SetUnitTransform(v string) {
	o.UnitTransform = &v
}

// GetValueFormat returns the ValueFormat field value if set, zero value otherwise.
func (o *VisualizationRule) GetValueFormat() string {
	if o == nil || IsNil(o.ValueFormat) {
		var ret string
		return ret
	}
	return *o.ValueFormat
}

// GetValueFormatOk returns a tuple with the ValueFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizationRule) GetValueFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ValueFormat) {
		return nil, false
	}
	return o.ValueFormat, true
}

// HasValueFormat returns a boolean if a field has been set.
func (o *VisualizationRule) HasValueFormat() bool {
	if o != nil && !IsNil(o.ValueFormat) {
		return true
	}

	return false
}

// SetValueFormat gets a reference to the given string and assigns it to the ValueFormat field.
func (o *VisualizationRule) SetValueFormat(v string) {
	o.ValueFormat = &v
}

func (o VisualizationRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matcher"] = o.Matcher
	toSerialize["properties"] = o.Properties
	if !IsNil(o.SeriesOverrides) {
		toSerialize["seriesOverrides"] = o.SeriesOverrides
	}
	if !IsNil(o.UnitTransform) {
		toSerialize["unitTransform"] = o.UnitTransform
	}
	if !IsNil(o.ValueFormat) {
		toSerialize["valueFormat"] = o.ValueFormat
	}
	return toSerialize, nil
}

type NullableVisualizationRule struct {
	value *VisualizationRule
	isSet bool
}

func (v NullableVisualizationRule) Get() *VisualizationRule {
	return v.value
}

func (v *NullableVisualizationRule) Set(val *VisualizationRule) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizationRule) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizationRule(val *VisualizationRule) *NullableVisualizationRule {
	return &NullableVisualizationRule{value: val, isSet: true}
}

func (v NullableVisualizationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


