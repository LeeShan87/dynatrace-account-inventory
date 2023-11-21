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

// checks if the ActiveGate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGate{}

// ActiveGate Parameters of the ActiveGate.
type ActiveGate struct {
	// A list of the ActiveGate tokens.
	ActiveGateTokens []ActiveGateTokenInfoDto `json:"activeGateTokens,omitempty"`
	AutoUpdateSettings *ActiveGateAutoUpdateConfig `json:"autoUpdateSettings,omitempty"`
	// The current status of auto-updates of the ActiveGate.
	AutoUpdateStatus *string `json:"autoUpdateStatus,omitempty"`
	ConnectedHosts *ActiveGateConnectedHosts `json:"connectedHosts,omitempty"`
	// ActiveGate is deployed in container (`true`) or not (`false`).
	Containerized *bool `json:"containerized,omitempty"`
	// A list of environments (specified by IDs) the ActiveGate can connect to.
	Environments []string `json:"environments,omitempty"`
	// The group of the ActiveGate.
	Group *string `json:"group,omitempty"`
	// The name of the host the ActiveGate is running on.
	Hostname *string `json:"hostname,omitempty"`
	// The ID of the ActiveGate.
	Id *string `json:"id,omitempty"`
	// A list of Load Balancer addresses of the ActiveGate.
	LoadBalancerAddresses []string `json:"loadBalancerAddresses,omitempty"`
	// The ID of the main environment for a multi-environment ActiveGate.
	MainEnvironment *string `json:"mainEnvironment,omitempty"`
	// A list of modules of the ActiveGate.
	Modules []ActiveGateModule `json:"modules,omitempty"`
	// A list of network addresses of the ActiveGate.
	NetworkAddresses []string `json:"networkAddresses,omitempty"`
	// The network zone of the ActiveGate.
	NetworkZone *string `json:"networkZone,omitempty"`
	// The timestamp since when the ActiveGate is offline.    The `null` value means the ActiveGate is online.
	OfflineSince *int64 `json:"offlineSince,omitempty"`
	// The OS architecture that the ActiveGate is running on.
	OsArchitecture *string `json:"osArchitecture,omitempty"`
	// The OS bitness that the ActiveGate is running on.
	OsBitness *string `json:"osBitness,omitempty"`
	// The OS type that the ActiveGate is running on.
	OsType *string `json:"osType,omitempty"`
	// The type of the ActiveGate.
	Type *string `json:"type,omitempty"`
	// The current version of the ActiveGate in the `<major>.<minor>.<revision>.<timestamp>` format.
	Version *string `json:"version,omitempty"`
}

// NewActiveGate instantiates a new ActiveGate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGate() *ActiveGate {
	this := ActiveGate{}
	return &this
}

// NewActiveGateWithDefaults instantiates a new ActiveGate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateWithDefaults() *ActiveGate {
	this := ActiveGate{}
	return &this
}

// GetActiveGateTokens returns the ActiveGateTokens field value if set, zero value otherwise.
func (o *ActiveGate) GetActiveGateTokens() []ActiveGateTokenInfoDto {
	if o == nil || IsNil(o.ActiveGateTokens) {
		var ret []ActiveGateTokenInfoDto
		return ret
	}
	return o.ActiveGateTokens
}

// GetActiveGateTokensOk returns a tuple with the ActiveGateTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetActiveGateTokensOk() ([]ActiveGateTokenInfoDto, bool) {
	if o == nil || IsNil(o.ActiveGateTokens) {
		return nil, false
	}
	return o.ActiveGateTokens, true
}

// HasActiveGateTokens returns a boolean if a field has been set.
func (o *ActiveGate) HasActiveGateTokens() bool {
	if o != nil && !IsNil(o.ActiveGateTokens) {
		return true
	}

	return false
}

// SetActiveGateTokens gets a reference to the given []ActiveGateTokenInfoDto and assigns it to the ActiveGateTokens field.
func (o *ActiveGate) SetActiveGateTokens(v []ActiveGateTokenInfoDto) {
	o.ActiveGateTokens = v
}

// GetAutoUpdateSettings returns the AutoUpdateSettings field value if set, zero value otherwise.
func (o *ActiveGate) GetAutoUpdateSettings() ActiveGateAutoUpdateConfig {
	if o == nil || IsNil(o.AutoUpdateSettings) {
		var ret ActiveGateAutoUpdateConfig
		return ret
	}
	return *o.AutoUpdateSettings
}

// GetAutoUpdateSettingsOk returns a tuple with the AutoUpdateSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetAutoUpdateSettingsOk() (*ActiveGateAutoUpdateConfig, bool) {
	if o == nil || IsNil(o.AutoUpdateSettings) {
		return nil, false
	}
	return o.AutoUpdateSettings, true
}

// HasAutoUpdateSettings returns a boolean if a field has been set.
func (o *ActiveGate) HasAutoUpdateSettings() bool {
	if o != nil && !IsNil(o.AutoUpdateSettings) {
		return true
	}

	return false
}

// SetAutoUpdateSettings gets a reference to the given ActiveGateAutoUpdateConfig and assigns it to the AutoUpdateSettings field.
func (o *ActiveGate) SetAutoUpdateSettings(v ActiveGateAutoUpdateConfig) {
	o.AutoUpdateSettings = &v
}

// GetAutoUpdateStatus returns the AutoUpdateStatus field value if set, zero value otherwise.
func (o *ActiveGate) GetAutoUpdateStatus() string {
	if o == nil || IsNil(o.AutoUpdateStatus) {
		var ret string
		return ret
	}
	return *o.AutoUpdateStatus
}

// GetAutoUpdateStatusOk returns a tuple with the AutoUpdateStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetAutoUpdateStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AutoUpdateStatus) {
		return nil, false
	}
	return o.AutoUpdateStatus, true
}

// HasAutoUpdateStatus returns a boolean if a field has been set.
func (o *ActiveGate) HasAutoUpdateStatus() bool {
	if o != nil && !IsNil(o.AutoUpdateStatus) {
		return true
	}

	return false
}

// SetAutoUpdateStatus gets a reference to the given string and assigns it to the AutoUpdateStatus field.
func (o *ActiveGate) SetAutoUpdateStatus(v string) {
	o.AutoUpdateStatus = &v
}

// GetConnectedHosts returns the ConnectedHosts field value if set, zero value otherwise.
func (o *ActiveGate) GetConnectedHosts() ActiveGateConnectedHosts {
	if o == nil || IsNil(o.ConnectedHosts) {
		var ret ActiveGateConnectedHosts
		return ret
	}
	return *o.ConnectedHosts
}

// GetConnectedHostsOk returns a tuple with the ConnectedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetConnectedHostsOk() (*ActiveGateConnectedHosts, bool) {
	if o == nil || IsNil(o.ConnectedHosts) {
		return nil, false
	}
	return o.ConnectedHosts, true
}

// HasConnectedHosts returns a boolean if a field has been set.
func (o *ActiveGate) HasConnectedHosts() bool {
	if o != nil && !IsNil(o.ConnectedHosts) {
		return true
	}

	return false
}

// SetConnectedHosts gets a reference to the given ActiveGateConnectedHosts and assigns it to the ConnectedHosts field.
func (o *ActiveGate) SetConnectedHosts(v ActiveGateConnectedHosts) {
	o.ConnectedHosts = &v
}

// GetContainerized returns the Containerized field value if set, zero value otherwise.
func (o *ActiveGate) GetContainerized() bool {
	if o == nil || IsNil(o.Containerized) {
		var ret bool
		return ret
	}
	return *o.Containerized
}

// GetContainerizedOk returns a tuple with the Containerized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetContainerizedOk() (*bool, bool) {
	if o == nil || IsNil(o.Containerized) {
		return nil, false
	}
	return o.Containerized, true
}

// HasContainerized returns a boolean if a field has been set.
func (o *ActiveGate) HasContainerized() bool {
	if o != nil && !IsNil(o.Containerized) {
		return true
	}

	return false
}

// SetContainerized gets a reference to the given bool and assigns it to the Containerized field.
func (o *ActiveGate) SetContainerized(v bool) {
	o.Containerized = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *ActiveGate) GetEnvironments() []string {
	if o == nil || IsNil(o.Environments) {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *ActiveGate) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []string and assigns it to the Environments field.
func (o *ActiveGate) SetEnvironments(v []string) {
	o.Environments = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ActiveGate) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ActiveGate) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ActiveGate) SetGroup(v string) {
	o.Group = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ActiveGate) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ActiveGate) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ActiveGate) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActiveGate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActiveGate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ActiveGate) SetId(v string) {
	o.Id = &v
}

// GetLoadBalancerAddresses returns the LoadBalancerAddresses field value if set, zero value otherwise.
func (o *ActiveGate) GetLoadBalancerAddresses() []string {
	if o == nil || IsNil(o.LoadBalancerAddresses) {
		var ret []string
		return ret
	}
	return o.LoadBalancerAddresses
}

// GetLoadBalancerAddressesOk returns a tuple with the LoadBalancerAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetLoadBalancerAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.LoadBalancerAddresses) {
		return nil, false
	}
	return o.LoadBalancerAddresses, true
}

// HasLoadBalancerAddresses returns a boolean if a field has been set.
func (o *ActiveGate) HasLoadBalancerAddresses() bool {
	if o != nil && !IsNil(o.LoadBalancerAddresses) {
		return true
	}

	return false
}

// SetLoadBalancerAddresses gets a reference to the given []string and assigns it to the LoadBalancerAddresses field.
func (o *ActiveGate) SetLoadBalancerAddresses(v []string) {
	o.LoadBalancerAddresses = v
}

// GetMainEnvironment returns the MainEnvironment field value if set, zero value otherwise.
func (o *ActiveGate) GetMainEnvironment() string {
	if o == nil || IsNil(o.MainEnvironment) {
		var ret string
		return ret
	}
	return *o.MainEnvironment
}

// GetMainEnvironmentOk returns a tuple with the MainEnvironment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetMainEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.MainEnvironment) {
		return nil, false
	}
	return o.MainEnvironment, true
}

// HasMainEnvironment returns a boolean if a field has been set.
func (o *ActiveGate) HasMainEnvironment() bool {
	if o != nil && !IsNil(o.MainEnvironment) {
		return true
	}

	return false
}

// SetMainEnvironment gets a reference to the given string and assigns it to the MainEnvironment field.
func (o *ActiveGate) SetMainEnvironment(v string) {
	o.MainEnvironment = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *ActiveGate) GetModules() []ActiveGateModule {
	if o == nil || IsNil(o.Modules) {
		var ret []ActiveGateModule
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetModulesOk() ([]ActiveGateModule, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ActiveGate) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given []ActiveGateModule and assigns it to the Modules field.
func (o *ActiveGate) SetModules(v []ActiveGateModule) {
	o.Modules = v
}

// GetNetworkAddresses returns the NetworkAddresses field value if set, zero value otherwise.
func (o *ActiveGate) GetNetworkAddresses() []string {
	if o == nil || IsNil(o.NetworkAddresses) {
		var ret []string
		return ret
	}
	return o.NetworkAddresses
}

// GetNetworkAddressesOk returns a tuple with the NetworkAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetNetworkAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkAddresses) {
		return nil, false
	}
	return o.NetworkAddresses, true
}

// HasNetworkAddresses returns a boolean if a field has been set.
func (o *ActiveGate) HasNetworkAddresses() bool {
	if o != nil && !IsNil(o.NetworkAddresses) {
		return true
	}

	return false
}

// SetNetworkAddresses gets a reference to the given []string and assigns it to the NetworkAddresses field.
func (o *ActiveGate) SetNetworkAddresses(v []string) {
	o.NetworkAddresses = v
}

// GetNetworkZone returns the NetworkZone field value if set, zero value otherwise.
func (o *ActiveGate) GetNetworkZone() string {
	if o == nil || IsNil(o.NetworkZone) {
		var ret string
		return ret
	}
	return *o.NetworkZone
}

// GetNetworkZoneOk returns a tuple with the NetworkZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetNetworkZoneOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkZone) {
		return nil, false
	}
	return o.NetworkZone, true
}

// HasNetworkZone returns a boolean if a field has been set.
func (o *ActiveGate) HasNetworkZone() bool {
	if o != nil && !IsNil(o.NetworkZone) {
		return true
	}

	return false
}

// SetNetworkZone gets a reference to the given string and assigns it to the NetworkZone field.
func (o *ActiveGate) SetNetworkZone(v string) {
	o.NetworkZone = &v
}

// GetOfflineSince returns the OfflineSince field value if set, zero value otherwise.
func (o *ActiveGate) GetOfflineSince() int64 {
	if o == nil || IsNil(o.OfflineSince) {
		var ret int64
		return ret
	}
	return *o.OfflineSince
}

// GetOfflineSinceOk returns a tuple with the OfflineSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetOfflineSinceOk() (*int64, bool) {
	if o == nil || IsNil(o.OfflineSince) {
		return nil, false
	}
	return o.OfflineSince, true
}

// HasOfflineSince returns a boolean if a field has been set.
func (o *ActiveGate) HasOfflineSince() bool {
	if o != nil && !IsNil(o.OfflineSince) {
		return true
	}

	return false
}

// SetOfflineSince gets a reference to the given int64 and assigns it to the OfflineSince field.
func (o *ActiveGate) SetOfflineSince(v int64) {
	o.OfflineSince = &v
}

// GetOsArchitecture returns the OsArchitecture field value if set, zero value otherwise.
func (o *ActiveGate) GetOsArchitecture() string {
	if o == nil || IsNil(o.OsArchitecture) {
		var ret string
		return ret
	}
	return *o.OsArchitecture
}

// GetOsArchitectureOk returns a tuple with the OsArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetOsArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.OsArchitecture) {
		return nil, false
	}
	return o.OsArchitecture, true
}

// HasOsArchitecture returns a boolean if a field has been set.
func (o *ActiveGate) HasOsArchitecture() bool {
	if o != nil && !IsNil(o.OsArchitecture) {
		return true
	}

	return false
}

// SetOsArchitecture gets a reference to the given string and assigns it to the OsArchitecture field.
func (o *ActiveGate) SetOsArchitecture(v string) {
	o.OsArchitecture = &v
}

// GetOsBitness returns the OsBitness field value if set, zero value otherwise.
func (o *ActiveGate) GetOsBitness() string {
	if o == nil || IsNil(o.OsBitness) {
		var ret string
		return ret
	}
	return *o.OsBitness
}

// GetOsBitnessOk returns a tuple with the OsBitness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetOsBitnessOk() (*string, bool) {
	if o == nil || IsNil(o.OsBitness) {
		return nil, false
	}
	return o.OsBitness, true
}

// HasOsBitness returns a boolean if a field has been set.
func (o *ActiveGate) HasOsBitness() bool {
	if o != nil && !IsNil(o.OsBitness) {
		return true
	}

	return false
}

// SetOsBitness gets a reference to the given string and assigns it to the OsBitness field.
func (o *ActiveGate) SetOsBitness(v string) {
	o.OsBitness = &v
}

// GetOsType returns the OsType field value if set, zero value otherwise.
func (o *ActiveGate) GetOsType() string {
	if o == nil || IsNil(o.OsType) {
		var ret string
		return ret
	}
	return *o.OsType
}

// GetOsTypeOk returns a tuple with the OsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetOsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OsType) {
		return nil, false
	}
	return o.OsType, true
}

// HasOsType returns a boolean if a field has been set.
func (o *ActiveGate) HasOsType() bool {
	if o != nil && !IsNil(o.OsType) {
		return true
	}

	return false
}

// SetOsType gets a reference to the given string and assigns it to the OsType field.
func (o *ActiveGate) SetOsType(v string) {
	o.OsType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActiveGate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActiveGate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActiveGate) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ActiveGate) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGate) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ActiveGate) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ActiveGate) SetVersion(v string) {
	o.Version = &v
}

func (o ActiveGate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActiveGateTokens) {
		toSerialize["activeGateTokens"] = o.ActiveGateTokens
	}
	if !IsNil(o.AutoUpdateSettings) {
		toSerialize["autoUpdateSettings"] = o.AutoUpdateSettings
	}
	if !IsNil(o.AutoUpdateStatus) {
		toSerialize["autoUpdateStatus"] = o.AutoUpdateStatus
	}
	if !IsNil(o.ConnectedHosts) {
		toSerialize["connectedHosts"] = o.ConnectedHosts
	}
	if !IsNil(o.Containerized) {
		toSerialize["containerized"] = o.Containerized
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LoadBalancerAddresses) {
		toSerialize["loadBalancerAddresses"] = o.LoadBalancerAddresses
	}
	if !IsNil(o.MainEnvironment) {
		toSerialize["mainEnvironment"] = o.MainEnvironment
	}
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.NetworkAddresses) {
		toSerialize["networkAddresses"] = o.NetworkAddresses
	}
	if !IsNil(o.NetworkZone) {
		toSerialize["networkZone"] = o.NetworkZone
	}
	if !IsNil(o.OfflineSince) {
		toSerialize["offlineSince"] = o.OfflineSince
	}
	if !IsNil(o.OsArchitecture) {
		toSerialize["osArchitecture"] = o.OsArchitecture
	}
	if !IsNil(o.OsBitness) {
		toSerialize["osBitness"] = o.OsBitness
	}
	if !IsNil(o.OsType) {
		toSerialize["osType"] = o.OsType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableActiveGate struct {
	value *ActiveGate
	isSet bool
}

func (v NullableActiveGate) Get() *ActiveGate {
	return v.value
}

func (v *NullableActiveGate) Set(val *ActiveGate) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGate) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGate(val *ActiveGate) *NullableActiveGate {
	return &NullableActiveGate{value: val, isSet: true}
}

func (v NullableActiveGate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


