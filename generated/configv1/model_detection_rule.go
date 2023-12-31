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

// checks if the DetectionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetectionRule{}

// DetectionRule struct for DetectionRule
type DetectionRule struct {
	// Additional annotations filter of the rule.   Only classes where all listed annotations are available in the class itself or any of its superclasses are instrumented.   Not applicable to PHP.
	Annotations []string `json:"annotations,omitempty"`
	// The fully qualified class or interface to instrument.   Required for Java and .NET custom services.    Not applicable to PHP.
	ClassName *string `json:"className,omitempty"`
	// Rule enabled/disabled.
	Enabled bool `json:"enabled"`
	// The PHP file containing the class or methods to instrument.   Required for PHP custom service.    Not applicable to Java and .NET.
	FileName *string `json:"fileName,omitempty"`
	// Matcher applying to the file name. Default value is `ENDS_WITH` (if applicable).
	FileNameMatcher *string `json:"fileNameMatcher,omitempty"`
	// The ID of the detection rule.
	Id *string `json:"id,omitempty"`
	// Matcher applying to the class name. `STARTS_WITH` can only be used if there is at least one annotation defined. Default value is `EQUALS`.
	Matcher *string `json:"matcher,omitempty"`
	// List of methods to instrument.
	MethodRules []MethodRule `json:"methodRules"`
}

// NewDetectionRule instantiates a new DetectionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetectionRule(enabled bool, methodRules []MethodRule) *DetectionRule {
	this := DetectionRule{}
	this.Enabled = enabled
	this.MethodRules = methodRules
	return &this
}

// NewDetectionRuleWithDefaults instantiates a new DetectionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetectionRuleWithDefaults() *DetectionRule {
	this := DetectionRule{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *DetectionRule) GetAnnotations() []string {
	if o == nil || IsNil(o.Annotations) {
		var ret []string
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetAnnotationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *DetectionRule) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given []string and assigns it to the Annotations field.
func (o *DetectionRule) SetAnnotations(v []string) {
	o.Annotations = v
}

// GetClassName returns the ClassName field value if set, zero value otherwise.
func (o *DetectionRule) GetClassName() string {
	if o == nil || IsNil(o.ClassName) {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClassName) {
		return nil, false
	}
	return o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *DetectionRule) HasClassName() bool {
	if o != nil && !IsNil(o.ClassName) {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *DetectionRule) SetClassName(v string) {
	o.ClassName = &v
}

// GetEnabled returns the Enabled field value
func (o *DetectionRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DetectionRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DetectionRule) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DetectionRule) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DetectionRule) SetFileName(v string) {
	o.FileName = &v
}

// GetFileNameMatcher returns the FileNameMatcher field value if set, zero value otherwise.
func (o *DetectionRule) GetFileNameMatcher() string {
	if o == nil || IsNil(o.FileNameMatcher) {
		var ret string
		return ret
	}
	return *o.FileNameMatcher
}

// GetFileNameMatcherOk returns a tuple with the FileNameMatcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetFileNameMatcherOk() (*string, bool) {
	if o == nil || IsNil(o.FileNameMatcher) {
		return nil, false
	}
	return o.FileNameMatcher, true
}

// HasFileNameMatcher returns a boolean if a field has been set.
func (o *DetectionRule) HasFileNameMatcher() bool {
	if o != nil && !IsNil(o.FileNameMatcher) {
		return true
	}

	return false
}

// SetFileNameMatcher gets a reference to the given string and assigns it to the FileNameMatcher field.
func (o *DetectionRule) SetFileNameMatcher(v string) {
	o.FileNameMatcher = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetectionRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetectionRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DetectionRule) SetId(v string) {
	o.Id = &v
}

// GetMatcher returns the Matcher field value if set, zero value otherwise.
func (o *DetectionRule) GetMatcher() string {
	if o == nil || IsNil(o.Matcher) {
		var ret string
		return ret
	}
	return *o.Matcher
}

// GetMatcherOk returns a tuple with the Matcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetMatcherOk() (*string, bool) {
	if o == nil || IsNil(o.Matcher) {
		return nil, false
	}
	return o.Matcher, true
}

// HasMatcher returns a boolean if a field has been set.
func (o *DetectionRule) HasMatcher() bool {
	if o != nil && !IsNil(o.Matcher) {
		return true
	}

	return false
}

// SetMatcher gets a reference to the given string and assigns it to the Matcher field.
func (o *DetectionRule) SetMatcher(v string) {
	o.Matcher = &v
}

// GetMethodRules returns the MethodRules field value
func (o *DetectionRule) GetMethodRules() []MethodRule {
	if o == nil {
		var ret []MethodRule
		return ret
	}

	return o.MethodRules
}

// GetMethodRulesOk returns a tuple with the MethodRules field value
// and a boolean to check if the value has been set.
func (o *DetectionRule) GetMethodRulesOk() ([]MethodRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.MethodRules, true
}

// SetMethodRules sets field value
func (o *DetectionRule) SetMethodRules(v []MethodRule) {
	o.MethodRules = v
}

func (o DetectionRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetectionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.ClassName) {
		toSerialize["className"] = o.ClassName
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.FileName) {
		toSerialize["fileName"] = o.FileName
	}
	if !IsNil(o.FileNameMatcher) {
		toSerialize["fileNameMatcher"] = o.FileNameMatcher
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Matcher) {
		toSerialize["matcher"] = o.Matcher
	}
	toSerialize["methodRules"] = o.MethodRules
	return toSerialize, nil
}

type NullableDetectionRule struct {
	value *DetectionRule
	isSet bool
}

func (v NullableDetectionRule) Get() *DetectionRule {
	return v.value
}

func (v *NullableDetectionRule) Set(val *DetectionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectionRule(val *DetectionRule) *NullableDetectionRule {
	return &NullableDetectionRule{value: val, isSet: true}
}

func (v NullableDetectionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


