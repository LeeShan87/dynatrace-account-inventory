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

// checks if the ExtensionEventDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionEventDto{}

// ExtensionEventDto A list of extension events.
type ExtensionEventDto struct {
	// Content of the event
	Content *string `json:"content,omitempty"`
	// Hexadecimal ID of Active Gate that uses this monitoring configuration.  Example: `0x1a2b3c4d`
	DtActiveGateId *string `json:"dt.active_gate.id,omitempty"`
	// Host that uses this monitoring configuration.  Example: `HOST-ABCDEF0123456789`
	DtEntityHost *string `json:"dt.entity.host,omitempty"`
	// Data source that uses this monitoring configuration.  Example: `snmp`
	DtExtensionDs *string `json:"dt.extension.ds,omitempty"`
	// Severity of the event
	Severity *string `json:"severity,omitempty"`
	// Status of the event
	Status *string `json:"status,omitempty"`
	// Timestamp of the event
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewExtensionEventDto instantiates a new ExtensionEventDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionEventDto() *ExtensionEventDto {
	this := ExtensionEventDto{}
	return &this
}

// NewExtensionEventDtoWithDefaults instantiates a new ExtensionEventDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionEventDtoWithDefaults() *ExtensionEventDto {
	this := ExtensionEventDto{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ExtensionEventDto) SetContent(v string) {
	o.Content = &v
}

// GetDtActiveGateId returns the DtActiveGateId field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetDtActiveGateId() string {
	if o == nil || IsNil(o.DtActiveGateId) {
		var ret string
		return ret
	}
	return *o.DtActiveGateId
}

// GetDtActiveGateIdOk returns a tuple with the DtActiveGateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetDtActiveGateIdOk() (*string, bool) {
	if o == nil || IsNil(o.DtActiveGateId) {
		return nil, false
	}
	return o.DtActiveGateId, true
}

// HasDtActiveGateId returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasDtActiveGateId() bool {
	if o != nil && !IsNil(o.DtActiveGateId) {
		return true
	}

	return false
}

// SetDtActiveGateId gets a reference to the given string and assigns it to the DtActiveGateId field.
func (o *ExtensionEventDto) SetDtActiveGateId(v string) {
	o.DtActiveGateId = &v
}

// GetDtEntityHost returns the DtEntityHost field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetDtEntityHost() string {
	if o == nil || IsNil(o.DtEntityHost) {
		var ret string
		return ret
	}
	return *o.DtEntityHost
}

// GetDtEntityHostOk returns a tuple with the DtEntityHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetDtEntityHostOk() (*string, bool) {
	if o == nil || IsNil(o.DtEntityHost) {
		return nil, false
	}
	return o.DtEntityHost, true
}

// HasDtEntityHost returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasDtEntityHost() bool {
	if o != nil && !IsNil(o.DtEntityHost) {
		return true
	}

	return false
}

// SetDtEntityHost gets a reference to the given string and assigns it to the DtEntityHost field.
func (o *ExtensionEventDto) SetDtEntityHost(v string) {
	o.DtEntityHost = &v
}

// GetDtExtensionDs returns the DtExtensionDs field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetDtExtensionDs() string {
	if o == nil || IsNil(o.DtExtensionDs) {
		var ret string
		return ret
	}
	return *o.DtExtensionDs
}

// GetDtExtensionDsOk returns a tuple with the DtExtensionDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetDtExtensionDsOk() (*string, bool) {
	if o == nil || IsNil(o.DtExtensionDs) {
		return nil, false
	}
	return o.DtExtensionDs, true
}

// HasDtExtensionDs returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasDtExtensionDs() bool {
	if o != nil && !IsNil(o.DtExtensionDs) {
		return true
	}

	return false
}

// SetDtExtensionDs gets a reference to the given string and assigns it to the DtExtensionDs field.
func (o *ExtensionEventDto) SetDtExtensionDs(v string) {
	o.DtExtensionDs = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *ExtensionEventDto) SetSeverity(v string) {
	o.Severity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ExtensionEventDto) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ExtensionEventDto) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionEventDto) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ExtensionEventDto) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ExtensionEventDto) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o ExtensionEventDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionEventDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.DtActiveGateId) {
		toSerialize["dt.active_gate.id"] = o.DtActiveGateId
	}
	if !IsNil(o.DtEntityHost) {
		toSerialize["dt.entity.host"] = o.DtEntityHost
	}
	if !IsNil(o.DtExtensionDs) {
		toSerialize["dt.extension.ds"] = o.DtExtensionDs
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableExtensionEventDto struct {
	value *ExtensionEventDto
	isSet bool
}

func (v NullableExtensionEventDto) Get() *ExtensionEventDto {
	return v.value
}

func (v *NullableExtensionEventDto) Set(val *ExtensionEventDto) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionEventDto) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionEventDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionEventDto(val *ExtensionEventDto) *NullableExtensionEventDto {
	return &NullableExtensionEventDto{value: val, isSet: true}
}

func (v NullableExtensionEventDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionEventDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

