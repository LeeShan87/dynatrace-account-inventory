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

// checks if the ExtensionState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionState{}

// ExtensionState The state of the extension.
type ExtensionState struct {
	// The ID of the endpoint where the state is detected - Active Gate only.
	EndpointId *string `json:"endpointId,omitempty"`
	// The ID of the extension.
	ExtensionId *string `json:"extensionId,omitempty"`
	// The ID of the host on which the extension runs.
	HostId *string `json:"hostId,omitempty"`
	// The ID of the entity on which the extension is active.
	ProcessId *string `json:"processId,omitempty"`
	// The state of the extension.
	State *string `json:"state,omitempty"`
	// A short description of the state.
	StateDescription *string `json:"stateDescription,omitempty"`
	// The timestamp when the state was detected, in UTC milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// The version of the extension (for example `1.0.0`).
	Version *string `json:"version,omitempty"`
}

// NewExtensionState instantiates a new ExtensionState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionState() *ExtensionState {
	this := ExtensionState{}
	return &this
}

// NewExtensionStateWithDefaults instantiates a new ExtensionState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionStateWithDefaults() *ExtensionState {
	this := ExtensionState{}
	return &this
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *ExtensionState) GetEndpointId() string {
	if o == nil || IsNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *ExtensionState) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *ExtensionState) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetExtensionId returns the ExtensionId field value if set, zero value otherwise.
func (o *ExtensionState) GetExtensionId() string {
	if o == nil || IsNil(o.ExtensionId) {
		var ret string
		return ret
	}
	return *o.ExtensionId
}

// GetExtensionIdOk returns a tuple with the ExtensionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetExtensionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExtensionId) {
		return nil, false
	}
	return o.ExtensionId, true
}

// HasExtensionId returns a boolean if a field has been set.
func (o *ExtensionState) HasExtensionId() bool {
	if o != nil && !IsNil(o.ExtensionId) {
		return true
	}

	return false
}

// SetExtensionId gets a reference to the given string and assigns it to the ExtensionId field.
func (o *ExtensionState) SetExtensionId(v string) {
	o.ExtensionId = &v
}

// GetHostId returns the HostId field value if set, zero value otherwise.
func (o *ExtensionState) GetHostId() string {
	if o == nil || IsNil(o.HostId) {
		var ret string
		return ret
	}
	return *o.HostId
}

// GetHostIdOk returns a tuple with the HostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetHostIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostId) {
		return nil, false
	}
	return o.HostId, true
}

// HasHostId returns a boolean if a field has been set.
func (o *ExtensionState) HasHostId() bool {
	if o != nil && !IsNil(o.HostId) {
		return true
	}

	return false
}

// SetHostId gets a reference to the given string and assigns it to the HostId field.
func (o *ExtensionState) SetHostId(v string) {
	o.HostId = &v
}

// GetProcessId returns the ProcessId field value if set, zero value otherwise.
func (o *ExtensionState) GetProcessId() string {
	if o == nil || IsNil(o.ProcessId) {
		var ret string
		return ret
	}
	return *o.ProcessId
}

// GetProcessIdOk returns a tuple with the ProcessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetProcessIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessId) {
		return nil, false
	}
	return o.ProcessId, true
}

// HasProcessId returns a boolean if a field has been set.
func (o *ExtensionState) HasProcessId() bool {
	if o != nil && !IsNil(o.ProcessId) {
		return true
	}

	return false
}

// SetProcessId gets a reference to the given string and assigns it to the ProcessId field.
func (o *ExtensionState) SetProcessId(v string) {
	o.ProcessId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ExtensionState) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExtensionState) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ExtensionState) SetState(v string) {
	o.State = &v
}

// GetStateDescription returns the StateDescription field value if set, zero value otherwise.
func (o *ExtensionState) GetStateDescription() string {
	if o == nil || IsNil(o.StateDescription) {
		var ret string
		return ret
	}
	return *o.StateDescription
}

// GetStateDescriptionOk returns a tuple with the StateDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetStateDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.StateDescription) {
		return nil, false
	}
	return o.StateDescription, true
}

// HasStateDescription returns a boolean if a field has been set.
func (o *ExtensionState) HasStateDescription() bool {
	if o != nil && !IsNil(o.StateDescription) {
		return true
	}

	return false
}

// SetStateDescription gets a reference to the given string and assigns it to the StateDescription field.
func (o *ExtensionState) SetStateDescription(v string) {
	o.StateDescription = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ExtensionState) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ExtensionState) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ExtensionState) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ExtensionState) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionState) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ExtensionState) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ExtensionState) SetVersion(v string) {
	o.Version = &v
}

func (o ExtensionState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointId) {
		toSerialize["endpointId"] = o.EndpointId
	}
	if !IsNil(o.ExtensionId) {
		toSerialize["extensionId"] = o.ExtensionId
	}
	if !IsNil(o.HostId) {
		toSerialize["hostId"] = o.HostId
	}
	if !IsNil(o.ProcessId) {
		toSerialize["processId"] = o.ProcessId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StateDescription) {
		toSerialize["stateDescription"] = o.StateDescription
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableExtensionState struct {
	value *ExtensionState
	isSet bool
}

func (v NullableExtensionState) Get() *ExtensionState {
	return v.value
}

func (v *NullableExtensionState) Set(val *ExtensionState) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionState) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionState(val *ExtensionState) *NullableExtensionState {
	return &NullableExtensionState{value: val, isSet: true}
}

func (v NullableExtensionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


