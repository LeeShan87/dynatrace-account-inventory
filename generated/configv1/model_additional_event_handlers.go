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

// checks if the AdditionalEventHandlers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalEventHandlers{}

// AdditionalEventHandlers Additional event handlers and wrappers.
type AdditionalEventHandlers struct {
	// Blur event handler enabled/disabled.
	BlurEventHandler bool `json:"blurEventHandler"`
	// Change event handler enabled/disabled.
	ChangeEventHandler bool `json:"changeEventHandler"`
	// Click event handler enabled/disabled.
	ClickEventHandler bool `json:"clickEventHandler"`
	// Max. number of DOM nodes to instrument. Valid values range from 0 to 100000.
	MaxDomNodesToInstrument int32 `json:"maxDomNodesToInstrument"`
	// Mouseup event handler enabled/disabled.
	MouseupEventHandler bool `json:"mouseupEventHandler"`
	// toString method enabled/disabled.
	ToStringMethod bool `json:"toStringMethod"`
	// Use mouseup event for clicks enabled/disabled.
	UserMouseupEventForClicks bool `json:"userMouseupEventForClicks"`
}

// NewAdditionalEventHandlers instantiates a new AdditionalEventHandlers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalEventHandlers(blurEventHandler bool, changeEventHandler bool, clickEventHandler bool, maxDomNodesToInstrument int32, mouseupEventHandler bool, toStringMethod bool, userMouseupEventForClicks bool) *AdditionalEventHandlers {
	this := AdditionalEventHandlers{}
	this.BlurEventHandler = blurEventHandler
	this.ChangeEventHandler = changeEventHandler
	this.ClickEventHandler = clickEventHandler
	this.MaxDomNodesToInstrument = maxDomNodesToInstrument
	this.MouseupEventHandler = mouseupEventHandler
	this.ToStringMethod = toStringMethod
	this.UserMouseupEventForClicks = userMouseupEventForClicks
	return &this
}

// NewAdditionalEventHandlersWithDefaults instantiates a new AdditionalEventHandlers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalEventHandlersWithDefaults() *AdditionalEventHandlers {
	this := AdditionalEventHandlers{}
	return &this
}

// GetBlurEventHandler returns the BlurEventHandler field value
func (o *AdditionalEventHandlers) GetBlurEventHandler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BlurEventHandler
}

// GetBlurEventHandlerOk returns a tuple with the BlurEventHandler field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetBlurEventHandlerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlurEventHandler, true
}

// SetBlurEventHandler sets field value
func (o *AdditionalEventHandlers) SetBlurEventHandler(v bool) {
	o.BlurEventHandler = v
}

// GetChangeEventHandler returns the ChangeEventHandler field value
func (o *AdditionalEventHandlers) GetChangeEventHandler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ChangeEventHandler
}

// GetChangeEventHandlerOk returns a tuple with the ChangeEventHandler field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetChangeEventHandlerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeEventHandler, true
}

// SetChangeEventHandler sets field value
func (o *AdditionalEventHandlers) SetChangeEventHandler(v bool) {
	o.ChangeEventHandler = v
}

// GetClickEventHandler returns the ClickEventHandler field value
func (o *AdditionalEventHandlers) GetClickEventHandler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClickEventHandler
}

// GetClickEventHandlerOk returns a tuple with the ClickEventHandler field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetClickEventHandlerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickEventHandler, true
}

// SetClickEventHandler sets field value
func (o *AdditionalEventHandlers) SetClickEventHandler(v bool) {
	o.ClickEventHandler = v
}

// GetMaxDomNodesToInstrument returns the MaxDomNodesToInstrument field value
func (o *AdditionalEventHandlers) GetMaxDomNodesToInstrument() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxDomNodesToInstrument
}

// GetMaxDomNodesToInstrumentOk returns a tuple with the MaxDomNodesToInstrument field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetMaxDomNodesToInstrumentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxDomNodesToInstrument, true
}

// SetMaxDomNodesToInstrument sets field value
func (o *AdditionalEventHandlers) SetMaxDomNodesToInstrument(v int32) {
	o.MaxDomNodesToInstrument = v
}

// GetMouseupEventHandler returns the MouseupEventHandler field value
func (o *AdditionalEventHandlers) GetMouseupEventHandler() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MouseupEventHandler
}

// GetMouseupEventHandlerOk returns a tuple with the MouseupEventHandler field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetMouseupEventHandlerOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MouseupEventHandler, true
}

// SetMouseupEventHandler sets field value
func (o *AdditionalEventHandlers) SetMouseupEventHandler(v bool) {
	o.MouseupEventHandler = v
}

// GetToStringMethod returns the ToStringMethod field value
func (o *AdditionalEventHandlers) GetToStringMethod() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ToStringMethod
}

// GetToStringMethodOk returns a tuple with the ToStringMethod field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetToStringMethodOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToStringMethod, true
}

// SetToStringMethod sets field value
func (o *AdditionalEventHandlers) SetToStringMethod(v bool) {
	o.ToStringMethod = v
}

// GetUserMouseupEventForClicks returns the UserMouseupEventForClicks field value
func (o *AdditionalEventHandlers) GetUserMouseupEventForClicks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserMouseupEventForClicks
}

// GetUserMouseupEventForClicksOk returns a tuple with the UserMouseupEventForClicks field value
// and a boolean to check if the value has been set.
func (o *AdditionalEventHandlers) GetUserMouseupEventForClicksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserMouseupEventForClicks, true
}

// SetUserMouseupEventForClicks sets field value
func (o *AdditionalEventHandlers) SetUserMouseupEventForClicks(v bool) {
	o.UserMouseupEventForClicks = v
}

func (o AdditionalEventHandlers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalEventHandlers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["blurEventHandler"] = o.BlurEventHandler
	toSerialize["changeEventHandler"] = o.ChangeEventHandler
	toSerialize["clickEventHandler"] = o.ClickEventHandler
	toSerialize["maxDomNodesToInstrument"] = o.MaxDomNodesToInstrument
	toSerialize["mouseupEventHandler"] = o.MouseupEventHandler
	toSerialize["toStringMethod"] = o.ToStringMethod
	toSerialize["userMouseupEventForClicks"] = o.UserMouseupEventForClicks
	return toSerialize, nil
}

type NullableAdditionalEventHandlers struct {
	value *AdditionalEventHandlers
	isSet bool
}

func (v NullableAdditionalEventHandlers) Get() *AdditionalEventHandlers {
	return v.value
}

func (v *NullableAdditionalEventHandlers) Set(val *AdditionalEventHandlers) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalEventHandlers) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalEventHandlers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalEventHandlers(val *AdditionalEventHandlers) *NullableAdditionalEventHandlers {
	return &NullableAdditionalEventHandlers{value: val, isSet: true}
}

func (v NullableAdditionalEventHandlers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalEventHandlers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


