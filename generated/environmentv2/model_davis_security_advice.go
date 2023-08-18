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

// checks if the DavisSecurityAdvice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DavisSecurityAdvice{}

// DavisSecurityAdvice Security advice from the Davis security advisor.
type DavisSecurityAdvice struct {
	// The type of the advice.
	AdviceType *string `json:"adviceType,omitempty"`
	// IDs of `critical` level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component.
	Critical []string `json:"critical,omitempty"`
	// IDs of `high` level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component.
	High []string `json:"high,omitempty"`
	// IDs of `low` level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component.
	Low []string `json:"low,omitempty"`
	// IDs of `medium` level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component.
	Medium []string `json:"medium,omitempty"`
	// The name of the advice.
	Name *string `json:"name,omitempty"`
	// IDs of `none` level [security problems](https://dt-url.net/p103u1h) caused by vulnerable component.
	None []string `json:"none,omitempty"`
	// The technology of the vulnerable component.
	Technology *string `json:"technology,omitempty"`
	// The vulnerable component to which advice applies.
	VulnerableComponent *string `json:"vulnerableComponent,omitempty"`
}

// NewDavisSecurityAdvice instantiates a new DavisSecurityAdvice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDavisSecurityAdvice() *DavisSecurityAdvice {
	this := DavisSecurityAdvice{}
	return &this
}

// NewDavisSecurityAdviceWithDefaults instantiates a new DavisSecurityAdvice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDavisSecurityAdviceWithDefaults() *DavisSecurityAdvice {
	this := DavisSecurityAdvice{}
	return &this
}

// GetAdviceType returns the AdviceType field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetAdviceType() string {
	if o == nil || IsNil(o.AdviceType) {
		var ret string
		return ret
	}
	return *o.AdviceType
}

// GetAdviceTypeOk returns a tuple with the AdviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetAdviceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AdviceType) {
		return nil, false
	}
	return o.AdviceType, true
}

// HasAdviceType returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasAdviceType() bool {
	if o != nil && !IsNil(o.AdviceType) {
		return true
	}

	return false
}

// SetAdviceType gets a reference to the given string and assigns it to the AdviceType field.
func (o *DavisSecurityAdvice) SetAdviceType(v string) {
	o.AdviceType = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetCritical() []string {
	if o == nil || IsNil(o.Critical) {
		var ret []string
		return ret
	}
	return o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetCriticalOk() ([]string, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given []string and assigns it to the Critical field.
func (o *DavisSecurityAdvice) SetCritical(v []string) {
	o.Critical = v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetHigh() []string {
	if o == nil || IsNil(o.High) {
		var ret []string
		return ret
	}
	return o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetHighOk() ([]string, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given []string and assigns it to the High field.
func (o *DavisSecurityAdvice) SetHigh(v []string) {
	o.High = v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetLow() []string {
	if o == nil || IsNil(o.Low) {
		var ret []string
		return ret
	}
	return o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetLowOk() ([]string, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given []string and assigns it to the Low field.
func (o *DavisSecurityAdvice) SetLow(v []string) {
	o.Low = v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetMedium() []string {
	if o == nil || IsNil(o.Medium) {
		var ret []string
		return ret
	}
	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetMediumOk() ([]string, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given []string and assigns it to the Medium field.
func (o *DavisSecurityAdvice) SetMedium(v []string) {
	o.Medium = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DavisSecurityAdvice) SetName(v string) {
	o.Name = &v
}

// GetNone returns the None field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetNone() []string {
	if o == nil || IsNil(o.None) {
		var ret []string
		return ret
	}
	return o.None
}

// GetNoneOk returns a tuple with the None field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetNoneOk() ([]string, bool) {
	if o == nil || IsNil(o.None) {
		return nil, false
	}
	return o.None, true
}

// HasNone returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasNone() bool {
	if o != nil && !IsNil(o.None) {
		return true
	}

	return false
}

// SetNone gets a reference to the given []string and assigns it to the None field.
func (o *DavisSecurityAdvice) SetNone(v []string) {
	o.None = v
}

// GetTechnology returns the Technology field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetTechnology() string {
	if o == nil || IsNil(o.Technology) {
		var ret string
		return ret
	}
	return *o.Technology
}

// GetTechnologyOk returns a tuple with the Technology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetTechnologyOk() (*string, bool) {
	if o == nil || IsNil(o.Technology) {
		return nil, false
	}
	return o.Technology, true
}

// HasTechnology returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasTechnology() bool {
	if o != nil && !IsNil(o.Technology) {
		return true
	}

	return false
}

// SetTechnology gets a reference to the given string and assigns it to the Technology field.
func (o *DavisSecurityAdvice) SetTechnology(v string) {
	o.Technology = &v
}

// GetVulnerableComponent returns the VulnerableComponent field value if set, zero value otherwise.
func (o *DavisSecurityAdvice) GetVulnerableComponent() string {
	if o == nil || IsNil(o.VulnerableComponent) {
		var ret string
		return ret
	}
	return *o.VulnerableComponent
}

// GetVulnerableComponentOk returns a tuple with the VulnerableComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DavisSecurityAdvice) GetVulnerableComponentOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerableComponent) {
		return nil, false
	}
	return o.VulnerableComponent, true
}

// HasVulnerableComponent returns a boolean if a field has been set.
func (o *DavisSecurityAdvice) HasVulnerableComponent() bool {
	if o != nil && !IsNil(o.VulnerableComponent) {
		return true
	}

	return false
}

// SetVulnerableComponent gets a reference to the given string and assigns it to the VulnerableComponent field.
func (o *DavisSecurityAdvice) SetVulnerableComponent(v string) {
	o.VulnerableComponent = &v
}

func (o DavisSecurityAdvice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DavisSecurityAdvice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdviceType) {
		toSerialize["adviceType"] = o.AdviceType
	}
	if !IsNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.None) {
		toSerialize["none"] = o.None
	}
	if !IsNil(o.Technology) {
		toSerialize["technology"] = o.Technology
	}
	if !IsNil(o.VulnerableComponent) {
		toSerialize["vulnerableComponent"] = o.VulnerableComponent
	}
	return toSerialize, nil
}

type NullableDavisSecurityAdvice struct {
	value *DavisSecurityAdvice
	isSet bool
}

func (v NullableDavisSecurityAdvice) Get() *DavisSecurityAdvice {
	return v.value
}

func (v *NullableDavisSecurityAdvice) Set(val *DavisSecurityAdvice) {
	v.value = val
	v.isSet = true
}

func (v NullableDavisSecurityAdvice) IsSet() bool {
	return v.isSet
}

func (v *NullableDavisSecurityAdvice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDavisSecurityAdvice(val *DavisSecurityAdvice) *NullableDavisSecurityAdvice {
	return &NullableDavisSecurityAdvice{value: val, isSet: true}
}

func (v NullableDavisSecurityAdvice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDavisSecurityAdvice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

