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

// checks if the PublicSyntheticLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSyntheticLocation{}

// PublicSyntheticLocation Configuration of a public synthetic location.    Some fields are inherited from the base *SyntheticLocation* object.
type PublicSyntheticLocation struct {
	// The type of the browser the location is using to execute browser monitors.
	BrowserType string `json:"browserType"`
	// The version of the browser the location is using to execute browser monitors.
	BrowserVersion string `json:"browserVersion"`
	// A list of location capabilities.
	Capabilities []string `json:"capabilities,omitempty"`
	// The cloud provider where the location is hosted.
	CloudPlatform string `json:"cloudPlatform"`
	// The list of IP addresses assigned to the location.
	Ips []string `json:"ips"`
	// The stage of the location.
	Stage string `json:"stage"`
}

// NewPublicSyntheticLocation instantiates a new PublicSyntheticLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSyntheticLocation(browserType string, browserVersion string, cloudPlatform string, ips []string, stage string, entityId string, latitude float64, longitude float64, name string, type_ string) *PublicSyntheticLocation {
	this := PublicSyntheticLocation{}
	this.EntityId = entityId
	this.Latitude = latitude
	this.Longitude = longitude
	this.Name = name
	this.Type = type_
	this.BrowserType = browserType
	this.BrowserVersion = browserVersion
	this.CloudPlatform = cloudPlatform
	this.Ips = ips
	this.Stage = stage
	return &this
}

// NewPublicSyntheticLocationWithDefaults instantiates a new PublicSyntheticLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSyntheticLocationWithDefaults() *PublicSyntheticLocation {
	this := PublicSyntheticLocation{}
	return &this
}

// GetBrowserType returns the BrowserType field value
func (o *PublicSyntheticLocation) GetBrowserType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserType
}

// GetBrowserTypeOk returns a tuple with the BrowserType field value
// and a boolean to check if the value has been set.
func (o *PublicSyntheticLocation) GetBrowserTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserType, true
}

// SetBrowserType sets field value
func (o *PublicSyntheticLocation) SetBrowserType(v string) {
	o.BrowserType = v
}

// GetBrowserVersion returns the BrowserVersion field value
func (o *PublicSyntheticLocation) GetBrowserVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value
// and a boolean to check if the value has been set.
func (o *PublicSyntheticLocation) GetBrowserVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrowserVersion, true
}

// SetBrowserVersion sets field value
func (o *PublicSyntheticLocation) SetBrowserVersion(v string) {
	o.BrowserVersion = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *PublicSyntheticLocation) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSyntheticLocation) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *PublicSyntheticLocation) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *PublicSyntheticLocation) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetCloudPlatform returns the CloudPlatform field value
func (o *PublicSyntheticLocation) GetCloudPlatform() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudPlatform
}

// GetCloudPlatformOk returns a tuple with the CloudPlatform field value
// and a boolean to check if the value has been set.
func (o *PublicSyntheticLocation) GetCloudPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudPlatform, true
}

// SetCloudPlatform sets field value
func (o *PublicSyntheticLocation) SetCloudPlatform(v string) {
	o.CloudPlatform = v
}

// GetIps returns the Ips field value
func (o *PublicSyntheticLocation) GetIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ips
}

// GetIpsOk returns a tuple with the Ips field value
// and a boolean to check if the value has been set.
func (o *PublicSyntheticLocation) GetIpsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ips, true
}

// SetIps sets field value
func (o *PublicSyntheticLocation) SetIps(v []string) {
	o.Ips = v
}

// GetStage returns the Stage field value
func (o *PublicSyntheticLocation) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *PublicSyntheticLocation) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *PublicSyntheticLocation) SetStage(v string) {
	o.Stage = v
}

func (o PublicSyntheticLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSyntheticLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["browserType"] = o.BrowserType
	toSerialize["browserVersion"] = o.BrowserVersion
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	toSerialize["cloudPlatform"] = o.CloudPlatform
	toSerialize["ips"] = o.Ips
	toSerialize["stage"] = o.Stage
	return toSerialize, nil
}

type NullablePublicSyntheticLocation struct {
	value *PublicSyntheticLocation
	isSet bool
}

func (v NullablePublicSyntheticLocation) Get() *PublicSyntheticLocation {
	return v.value
}

func (v *NullablePublicSyntheticLocation) Set(val *PublicSyntheticLocation) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSyntheticLocation) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSyntheticLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSyntheticLocation(val *PublicSyntheticLocation) *NullablePublicSyntheticLocation {
	return &NullablePublicSyntheticLocation{value: val, isSet: true}
}

func (v NullablePublicSyntheticLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSyntheticLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

