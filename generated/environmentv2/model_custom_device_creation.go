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

// checks if the CustomDeviceCreation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDeviceCreation{}

// CustomDeviceCreation Configuration of a custom device.
type CustomDeviceCreation struct {
	// The URL of a configuration web page for the custom device, such as a login page for a firewall or router.
	ConfigUrl *string `json:"configUrl,omitempty"`
	// The internal ID of the custom device.    If you use the ID of an existing device, the respective parameters will be updated.
	CustomDeviceId string `json:"customDeviceId"`
	// The name of the custom device to be displayed in the user interface.
	DisplayName string `json:"displayName"`
	// The list of DNS names related to the custom device.   These names are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value, the existing values will be overwritten.   If you send `null` or an empty value; or omit this field, the existing values will be kept.
	DnsNames []string `json:"dnsNames,omitempty"`
	// The icon to be displayed for your custom component within Smartscape. Provide the full URL of the icon file.
	FaviconUrl *string `json:"faviconUrl,omitempty"`
	// User defined group ID of entity.   The group ID helps to keep a consistent picture of device-group relations. One of many cases where a proper group is important is service detection: you can define which custom devices should lead to the same service by defining the same group ID for them.   If you set a group ID, it will be hashed into the Dynatrace entity ID of the custom device. In that case the custom device can only be part of one custom device group.   If you don't set the group ID, Dynatrace will create it based on the ID or type of the custom device. Also, the group will not be hashed into the device ID which means the device may switch groups.
	Group *string `json:"group,omitempty"`
	// The list of IP addresses that belong to the custom device.   These addresses are used to automatically discover the horizontal communication relationship between this component and all other observed components within Smartscape. Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If you send a value (including an empty value), the existing values will be overwritten.   If you send `null` or omit this field, the existing values will be kept.
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// The list of ports the custom devices listens to.   These ports are used to discover the horizontal communication relationship between this component and all other observed components within Smartscape.   Once a connection is discovered, it is automatically mapped and shown within Smartscape.   If ports are specified, you should also add at least one IP address or a DNS name for the custom device.   If you send a value, the existing values will be overwritten.   If you send `null`, or an empty value, or omit this field, the existing values will be kept.
	ListenPorts []int32 `json:"listenPorts,omitempty"`
	// The list of key-value pair properties that will be shown beneath the infographics of your custom device.
	Properties *map[string]string `json:"properties,omitempty"`
	// The technology type definition of the custom device.   It must be the same technology type of the metric you're reporting.   If you send a value, the existing value will be overwritten.   If you send `null`, empty or omit this field, the existing value will be kept.
	Type *string `json:"type,omitempty"`
}

// NewCustomDeviceCreation instantiates a new CustomDeviceCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDeviceCreation(customDeviceId string, displayName string) *CustomDeviceCreation {
	this := CustomDeviceCreation{}
	this.CustomDeviceId = customDeviceId
	this.DisplayName = displayName
	return &this
}

// NewCustomDeviceCreationWithDefaults instantiates a new CustomDeviceCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDeviceCreationWithDefaults() *CustomDeviceCreation {
	this := CustomDeviceCreation{}
	return &this
}

// GetConfigUrl returns the ConfigUrl field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetConfigUrl() string {
	if o == nil || IsNil(o.ConfigUrl) {
		var ret string
		return ret
	}
	return *o.ConfigUrl
}

// GetConfigUrlOk returns a tuple with the ConfigUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetConfigUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigUrl) {
		return nil, false
	}
	return o.ConfigUrl, true
}

// HasConfigUrl returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasConfigUrl() bool {
	if o != nil && !IsNil(o.ConfigUrl) {
		return true
	}

	return false
}

// SetConfigUrl gets a reference to the given string and assigns it to the ConfigUrl field.
func (o *CustomDeviceCreation) SetConfigUrl(v string) {
	o.ConfigUrl = &v
}

// GetCustomDeviceId returns the CustomDeviceId field value
func (o *CustomDeviceCreation) GetCustomDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomDeviceId
}

// GetCustomDeviceIdOk returns a tuple with the CustomDeviceId field value
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetCustomDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomDeviceId, true
}

// SetCustomDeviceId sets field value
func (o *CustomDeviceCreation) SetCustomDeviceId(v string) {
	o.CustomDeviceId = v
}

// GetDisplayName returns the DisplayName field value
func (o *CustomDeviceCreation) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CustomDeviceCreation) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDnsNames returns the DnsNames field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetDnsNames() []string {
	if o == nil || IsNil(o.DnsNames) {
		var ret []string
		return ret
	}
	return o.DnsNames
}

// GetDnsNamesOk returns a tuple with the DnsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetDnsNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DnsNames) {
		return nil, false
	}
	return o.DnsNames, true
}

// HasDnsNames returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasDnsNames() bool {
	if o != nil && !IsNil(o.DnsNames) {
		return true
	}

	return false
}

// SetDnsNames gets a reference to the given []string and assigns it to the DnsNames field.
func (o *CustomDeviceCreation) SetDnsNames(v []string) {
	o.DnsNames = v
}

// GetFaviconUrl returns the FaviconUrl field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetFaviconUrl() string {
	if o == nil || IsNil(o.FaviconUrl) {
		var ret string
		return ret
	}
	return *o.FaviconUrl
}

// GetFaviconUrlOk returns a tuple with the FaviconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetFaviconUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FaviconUrl) {
		return nil, false
	}
	return o.FaviconUrl, true
}

// HasFaviconUrl returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasFaviconUrl() bool {
	if o != nil && !IsNil(o.FaviconUrl) {
		return true
	}

	return false
}

// SetFaviconUrl gets a reference to the given string and assigns it to the FaviconUrl field.
func (o *CustomDeviceCreation) SetFaviconUrl(v string) {
	o.FaviconUrl = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *CustomDeviceCreation) SetGroup(v string) {
	o.Group = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetIpAddresses() []string {
	if o == nil || IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetIpAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasIpAddresses() bool {
	if o != nil && !IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *CustomDeviceCreation) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetListenPorts returns the ListenPorts field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetListenPorts() []int32 {
	if o == nil || IsNil(o.ListenPorts) {
		var ret []int32
		return ret
	}
	return o.ListenPorts
}

// GetListenPortsOk returns a tuple with the ListenPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetListenPortsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ListenPorts) {
		return nil, false
	}
	return o.ListenPorts, true
}

// HasListenPorts returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasListenPorts() bool {
	if o != nil && !IsNil(o.ListenPorts) {
		return true
	}

	return false
}

// SetListenPorts gets a reference to the given []int32 and assigns it to the ListenPorts field.
func (o *CustomDeviceCreation) SetListenPorts(v []int32) {
	o.ListenPorts = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *CustomDeviceCreation) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CustomDeviceCreation) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreation) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CustomDeviceCreation) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CustomDeviceCreation) SetType(v string) {
	o.Type = &v
}

func (o CustomDeviceCreation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDeviceCreation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigUrl) {
		toSerialize["configUrl"] = o.ConfigUrl
	}
	toSerialize["customDeviceId"] = o.CustomDeviceId
	toSerialize["displayName"] = o.DisplayName
	if !IsNil(o.DnsNames) {
		toSerialize["dnsNames"] = o.DnsNames
	}
	if !IsNil(o.FaviconUrl) {
		toSerialize["faviconUrl"] = o.FaviconUrl
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !IsNil(o.ListenPorts) {
		toSerialize["listenPorts"] = o.ListenPorts
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCustomDeviceCreation struct {
	value *CustomDeviceCreation
	isSet bool
}

func (v NullableCustomDeviceCreation) Get() *CustomDeviceCreation {
	return v.value
}

func (v *NullableCustomDeviceCreation) Set(val *CustomDeviceCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDeviceCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDeviceCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDeviceCreation(val *CustomDeviceCreation) *NullableCustomDeviceCreation {
	return &NullableCustomDeviceCreation{value: val, isSet: true}
}

func (v NullableCustomDeviceCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDeviceCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


