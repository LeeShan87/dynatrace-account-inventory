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

// checks if the KubernetesEventPattern type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubernetesEventPattern{}

// KubernetesEventPattern Represents a single Kubernetes events field selector (=event filter based on the K8s field selector).
type KubernetesEventPattern struct {
	// Whether subscription to this events field selector is enabled (value set to `true`). If disabled (value set to `false`), Dynatrace will stop fetching events from the Kubernetes API for this events field selector
	Active bool `json:"active"`
	// The field selector string (url decoding is applied) when storing it.
	FieldSelector string `json:"fieldSelector"`
	// A label of the events field selector.
	Label string `json:"label"`
}

// NewKubernetesEventPattern instantiates a new KubernetesEventPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesEventPattern(active bool, fieldSelector string, label string) *KubernetesEventPattern {
	this := KubernetesEventPattern{}
	this.Active = active
	this.FieldSelector = fieldSelector
	this.Label = label
	return &this
}

// NewKubernetesEventPatternWithDefaults instantiates a new KubernetesEventPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesEventPatternWithDefaults() *KubernetesEventPattern {
	this := KubernetesEventPattern{}
	return &this
}

// GetActive returns the Active field value
func (o *KubernetesEventPattern) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventPattern) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *KubernetesEventPattern) SetActive(v bool) {
	o.Active = v
}

// GetFieldSelector returns the FieldSelector field value
func (o *KubernetesEventPattern) GetFieldSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldSelector
}

// GetFieldSelectorOk returns a tuple with the FieldSelector field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventPattern) GetFieldSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldSelector, true
}

// SetFieldSelector sets field value
func (o *KubernetesEventPattern) SetFieldSelector(v string) {
	o.FieldSelector = v
}

// GetLabel returns the Label field value
func (o *KubernetesEventPattern) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *KubernetesEventPattern) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *KubernetesEventPattern) SetLabel(v string) {
	o.Label = v
}

func (o KubernetesEventPattern) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubernetesEventPattern) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["fieldSelector"] = o.FieldSelector
	toSerialize["label"] = o.Label
	return toSerialize, nil
}

type NullableKubernetesEventPattern struct {
	value *KubernetesEventPattern
	isSet bool
}

func (v NullableKubernetesEventPattern) Get() *KubernetesEventPattern {
	return v.value
}

func (v *NullableKubernetesEventPattern) Set(val *KubernetesEventPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesEventPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesEventPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesEventPattern(val *KubernetesEventPattern) *NullableKubernetesEventPattern {
	return &NullableKubernetesEventPattern{value: val, isSet: true}
}

func (v NullableKubernetesEventPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesEventPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


