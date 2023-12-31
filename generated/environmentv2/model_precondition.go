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

// checks if the Precondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Precondition{}

// Precondition A precondition for visibility of a property.
type Precondition struct {
	// The expected value of the property.   Only applicable to properties of the `EQUALS` type.
	ExpectedValue map[string]interface{} `json:"expectedValue,omitempty"`
	// A list of valid values of the property.   Only applicable to properties of the `IN` type.
	ExpectedValues []map[string]interface{} `json:"expectedValues,omitempty"`
	// The Regular expression which is matched against the property.   Only applicable to properties of the `REGEX_MATCH` type.
	Pattern *string `json:"pattern,omitempty"`
	Precondition *Precondition `json:"precondition,omitempty"`
	// A list of child preconditions to be evaluated.   Only applicable to properties of the `AND` and `OR` types.
	Preconditions []Precondition `json:"preconditions,omitempty"`
	// The property to be evaluated.
	Property *string `json:"property,omitempty"`
	// The type of the precondition.
	Type string `json:"type"`
}

// NewPrecondition instantiates a new Precondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrecondition(type_ string) *Precondition {
	this := Precondition{}
	this.Type = type_
	return &this
}

// NewPreconditionWithDefaults instantiates a new Precondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreconditionWithDefaults() *Precondition {
	this := Precondition{}
	return &this
}

// GetExpectedValue returns the ExpectedValue field value if set, zero value otherwise.
func (o *Precondition) GetExpectedValue() map[string]interface{} {
	if o == nil || IsNil(o.ExpectedValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExpectedValue
}

// GetExpectedValueOk returns a tuple with the ExpectedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Precondition) GetExpectedValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExpectedValue) {
		return map[string]interface{}{}, false
	}
	return o.ExpectedValue, true
}

// HasExpectedValue returns a boolean if a field has been set.
func (o *Precondition) HasExpectedValue() bool {
	if o != nil && !IsNil(o.ExpectedValue) {
		return true
	}

	return false
}

// SetExpectedValue gets a reference to the given map[string]interface{} and assigns it to the ExpectedValue field.
func (o *Precondition) SetExpectedValue(v map[string]interface{}) {
	o.ExpectedValue = v
}

// GetExpectedValues returns the ExpectedValues field value if set, zero value otherwise.
func (o *Precondition) GetExpectedValues() []map[string]interface{} {
	if o == nil || IsNil(o.ExpectedValues) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ExpectedValues
}

// GetExpectedValuesOk returns a tuple with the ExpectedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Precondition) GetExpectedValuesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExpectedValues) {
		return nil, false
	}
	return o.ExpectedValues, true
}

// HasExpectedValues returns a boolean if a field has been set.
func (o *Precondition) HasExpectedValues() bool {
	if o != nil && !IsNil(o.ExpectedValues) {
		return true
	}

	return false
}

// SetExpectedValues gets a reference to the given []map[string]interface{} and assigns it to the ExpectedValues field.
func (o *Precondition) SetExpectedValues(v []map[string]interface{}) {
	o.ExpectedValues = v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *Precondition) GetPattern() string {
	if o == nil || IsNil(o.Pattern) {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Precondition) GetPatternOk() (*string, bool) {
	if o == nil || IsNil(o.Pattern) {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *Precondition) HasPattern() bool {
	if o != nil && !IsNil(o.Pattern) {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *Precondition) SetPattern(v string) {
	o.Pattern = &v
}

// GetPrecondition returns the Precondition field value if set, zero value otherwise.
func (o *Precondition) GetPrecondition() Precondition {
	if o == nil || IsNil(o.Precondition) {
		var ret Precondition
		return ret
	}
	return *o.Precondition
}

// GetPreconditionOk returns a tuple with the Precondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Precondition) GetPreconditionOk() (*Precondition, bool) {
	if o == nil || IsNil(o.Precondition) {
		return nil, false
	}
	return o.Precondition, true
}

// HasPrecondition returns a boolean if a field has been set.
func (o *Precondition) HasPrecondition() bool {
	if o != nil && !IsNil(o.Precondition) {
		return true
	}

	return false
}

// SetPrecondition gets a reference to the given Precondition and assigns it to the Precondition field.
func (o *Precondition) SetPrecondition(v Precondition) {
	o.Precondition = &v
}

// GetPreconditions returns the Preconditions field value if set, zero value otherwise.
func (o *Precondition) GetPreconditions() []Precondition {
	if o == nil || IsNil(o.Preconditions) {
		var ret []Precondition
		return ret
	}
	return o.Preconditions
}

// GetPreconditionsOk returns a tuple with the Preconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Precondition) GetPreconditionsOk() ([]Precondition, bool) {
	if o == nil || IsNil(o.Preconditions) {
		return nil, false
	}
	return o.Preconditions, true
}

// HasPreconditions returns a boolean if a field has been set.
func (o *Precondition) HasPreconditions() bool {
	if o != nil && !IsNil(o.Preconditions) {
		return true
	}

	return false
}

// SetPreconditions gets a reference to the given []Precondition and assigns it to the Preconditions field.
func (o *Precondition) SetPreconditions(v []Precondition) {
	o.Preconditions = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *Precondition) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Precondition) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *Precondition) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *Precondition) SetProperty(v string) {
	o.Property = &v
}

// GetType returns the Type field value
func (o *Precondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Precondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Precondition) SetType(v string) {
	o.Type = v
}

func (o Precondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Precondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpectedValue) {
		toSerialize["expectedValue"] = o.ExpectedValue
	}
	if !IsNil(o.ExpectedValues) {
		toSerialize["expectedValues"] = o.ExpectedValues
	}
	if !IsNil(o.Pattern) {
		toSerialize["pattern"] = o.Pattern
	}
	if !IsNil(o.Precondition) {
		toSerialize["precondition"] = o.Precondition
	}
	if !IsNil(o.Preconditions) {
		toSerialize["preconditions"] = o.Preconditions
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePrecondition struct {
	value *Precondition
	isSet bool
}

func (v NullablePrecondition) Get() *Precondition {
	return v.value
}

func (v *NullablePrecondition) Set(val *Precondition) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecondition) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecondition(val *Precondition) *NullablePrecondition {
	return &NullablePrecondition{value: val, isSet: true}
}

func (v NullablePrecondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


