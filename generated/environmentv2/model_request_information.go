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

// checks if the RequestInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestInformation{}

// RequestInformation Describes the complete request information of an attack.
type RequestInformation struct {
	// The target host of the request.
	Host *string `json:"host,omitempty"`
	// The request path.
	Path *string `json:"path,omitempty"`
	ProtocolDetails *ProtocolDetails `json:"protocolDetails,omitempty"`
	// The requested URL.
	Url *string `json:"url,omitempty"`
}

// NewRequestInformation instantiates a new RequestInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestInformation() *RequestInformation {
	this := RequestInformation{}
	return &this
}

// NewRequestInformationWithDefaults instantiates a new RequestInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestInformationWithDefaults() *RequestInformation {
	this := RequestInformation{}
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *RequestInformation) GetHost() string {
	if o == nil || IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestInformation) GetHostOk() (*string, bool) {
	if o == nil || IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *RequestInformation) HasHost() bool {
	if o != nil && !IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *RequestInformation) SetHost(v string) {
	o.Host = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RequestInformation) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestInformation) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RequestInformation) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RequestInformation) SetPath(v string) {
	o.Path = &v
}

// GetProtocolDetails returns the ProtocolDetails field value if set, zero value otherwise.
func (o *RequestInformation) GetProtocolDetails() ProtocolDetails {
	if o == nil || IsNil(o.ProtocolDetails) {
		var ret ProtocolDetails
		return ret
	}
	return *o.ProtocolDetails
}

// GetProtocolDetailsOk returns a tuple with the ProtocolDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestInformation) GetProtocolDetailsOk() (*ProtocolDetails, bool) {
	if o == nil || IsNil(o.ProtocolDetails) {
		return nil, false
	}
	return o.ProtocolDetails, true
}

// HasProtocolDetails returns a boolean if a field has been set.
func (o *RequestInformation) HasProtocolDetails() bool {
	if o != nil && !IsNil(o.ProtocolDetails) {
		return true
	}

	return false
}

// SetProtocolDetails gets a reference to the given ProtocolDetails and assigns it to the ProtocolDetails field.
func (o *RequestInformation) SetProtocolDetails(v ProtocolDetails) {
	o.ProtocolDetails = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *RequestInformation) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestInformation) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *RequestInformation) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *RequestInformation) SetUrl(v string) {
	o.Url = &v
}

func (o RequestInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.ProtocolDetails) {
		toSerialize["protocolDetails"] = o.ProtocolDetails
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableRequestInformation struct {
	value *RequestInformation
	isSet bool
}

func (v NullableRequestInformation) Get() *RequestInformation {
	return v.value
}

func (v *NullableRequestInformation) Set(val *RequestInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestInformation(val *RequestInformation) *NullableRequestInformation {
	return &NullableRequestInformation{value: val, isSet: true}
}

func (v NullableRequestInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


