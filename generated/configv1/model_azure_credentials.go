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

// checks if the AzureCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureCredentials{}

// AzureCredentials Configuration of Azure app-level credentials.
type AzureCredentials struct {
	// The monitoring is enabled (`true`) or disabled (`false`).   If not set on creation, the `true` value is used.   If the field is omitted during an update, the old value remains unaffected.
	Active *bool `json:"active,omitempty"`
	// The application ID (also referred to as client ID).   The field is **required** when creating a new credentials configuration.    The field is ignored during an update, the old value remains unaffected.
	AppId *string `json:"appId,omitempty"`
	// The automatic capture of Azure tags is on (`true`) or off (`false`).
	AutoTagging bool `json:"autoTagging"`
	// The directory ID (also referred to as tenant ID).   The field is **required** when creating a new credentials configuration.    The field is ignored during an update, the old value remains unaffected.
	DirectoryId *string `json:"directoryId,omitempty"`
	// The Dynatrace entity ID of the Azure credentials configuration.
	Id *string `json:"id,omitempty"`
	// The secret key associated with the application ID.   For security reasons, GET requests return this field as `null`.    Submit your key on creation or update of the configuration.    The field is **required** when creating a new credentials configuration. If the field is omitted during an update, the old value remains unaffected.
	Key *string `json:"key,omitempty"`
	// The unique name of the Azure credentials configuration.   Allowed characters are letters, numbers, and spaces. Also the special characters `.+-_` are allowed.
	Label string `json:"label"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// A list of Azure tags to be excluded from monitoring.   You can specify up to 20 tags. A resource tagged with *any* of the specified tags will not be monitored.   Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`.
	MonitorOnlyExcludingTagPairs []CloudTag `json:"monitorOnlyExcludingTagPairs,omitempty"`
	// A list of Azure tags to be monitored.   You can specify up to 20 tags. A resource tagged with *any* of the specified tags is monitored.   Only applicable when the **monitorOnlyTaggedEntities** parameter is set to `true`.
	MonitorOnlyTagPairs []CloudTag `json:"monitorOnlyTagPairs,omitempty"`
	// Monitor only resources that have specified Azure tags (`true`) or all resources (`false`).
	MonitorOnlyTaggedEntities bool `json:"monitorOnlyTaggedEntities"`
	// **Deprecated**. To manage services use [/azure/credentials/{id}/services](https://dt-url.net/1w62s27) operation. Built-in services are not supported here.  A list of Azure services to be monitored. Available services are listed by [/azure/supportedServices](https://dt-url.net/wt42sdq) operation.  For each service, a list of metrics and dimensions can be specified. A list of supported metrics and dimensions for a given service can be checked in [documentation](https://dt-url.net/kx2351b).  List of metrics can be skipped (set to null), resulting in recommended (default) set of metrics and dimensions being chosen for monitoring. 
	SupportingServices []AzureSupportingService `json:"supportingServices,omitempty"`
}

// NewAzureCredentials instantiates a new AzureCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureCredentials(autoTagging bool, label string, monitorOnlyTaggedEntities bool) *AzureCredentials {
	this := AzureCredentials{}
	this.AutoTagging = autoTagging
	this.Label = label
	this.MonitorOnlyTaggedEntities = monitorOnlyTaggedEntities
	return &this
}

// NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureCredentialsWithDefaults() *AzureCredentials {
	this := AzureCredentials{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AzureCredentials) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AzureCredentials) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AzureCredentials) SetActive(v bool) {
	o.Active = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AzureCredentials) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AzureCredentials) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AzureCredentials) SetAppId(v string) {
	o.AppId = &v
}

// GetAutoTagging returns the AutoTagging field value
func (o *AzureCredentials) GetAutoTagging() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoTagging
}

// GetAutoTaggingOk returns a tuple with the AutoTagging field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetAutoTaggingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoTagging, true
}

// SetAutoTagging sets field value
func (o *AzureCredentials) SetAutoTagging(v bool) {
	o.AutoTagging = v
}

// GetDirectoryId returns the DirectoryId field value if set, zero value otherwise.
func (o *AzureCredentials) GetDirectoryId() string {
	if o == nil || IsNil(o.DirectoryId) {
		var ret string
		return ret
	}
	return *o.DirectoryId
}

// GetDirectoryIdOk returns a tuple with the DirectoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetDirectoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.DirectoryId) {
		return nil, false
	}
	return o.DirectoryId, true
}

// HasDirectoryId returns a boolean if a field has been set.
func (o *AzureCredentials) HasDirectoryId() bool {
	if o != nil && !IsNil(o.DirectoryId) {
		return true
	}

	return false
}

// SetDirectoryId gets a reference to the given string and assigns it to the DirectoryId field.
func (o *AzureCredentials) SetDirectoryId(v string) {
	o.DirectoryId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureCredentials) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureCredentials) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureCredentials) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AzureCredentials) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AzureCredentials) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AzureCredentials) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value
func (o *AzureCredentials) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *AzureCredentials) SetLabel(v string) {
	o.Label = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AzureCredentials) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AzureCredentials) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *AzureCredentials) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetMonitorOnlyExcludingTagPairs returns the MonitorOnlyExcludingTagPairs field value if set, zero value otherwise.
func (o *AzureCredentials) GetMonitorOnlyExcludingTagPairs() []CloudTag {
	if o == nil || IsNil(o.MonitorOnlyExcludingTagPairs) {
		var ret []CloudTag
		return ret
	}
	return o.MonitorOnlyExcludingTagPairs
}

// GetMonitorOnlyExcludingTagPairsOk returns a tuple with the MonitorOnlyExcludingTagPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetMonitorOnlyExcludingTagPairsOk() ([]CloudTag, bool) {
	if o == nil || IsNil(o.MonitorOnlyExcludingTagPairs) {
		return nil, false
	}
	return o.MonitorOnlyExcludingTagPairs, true
}

// HasMonitorOnlyExcludingTagPairs returns a boolean if a field has been set.
func (o *AzureCredentials) HasMonitorOnlyExcludingTagPairs() bool {
	if o != nil && !IsNil(o.MonitorOnlyExcludingTagPairs) {
		return true
	}

	return false
}

// SetMonitorOnlyExcludingTagPairs gets a reference to the given []CloudTag and assigns it to the MonitorOnlyExcludingTagPairs field.
func (o *AzureCredentials) SetMonitorOnlyExcludingTagPairs(v []CloudTag) {
	o.MonitorOnlyExcludingTagPairs = v
}

// GetMonitorOnlyTagPairs returns the MonitorOnlyTagPairs field value if set, zero value otherwise.
func (o *AzureCredentials) GetMonitorOnlyTagPairs() []CloudTag {
	if o == nil || IsNil(o.MonitorOnlyTagPairs) {
		var ret []CloudTag
		return ret
	}
	return o.MonitorOnlyTagPairs
}

// GetMonitorOnlyTagPairsOk returns a tuple with the MonitorOnlyTagPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetMonitorOnlyTagPairsOk() ([]CloudTag, bool) {
	if o == nil || IsNil(o.MonitorOnlyTagPairs) {
		return nil, false
	}
	return o.MonitorOnlyTagPairs, true
}

// HasMonitorOnlyTagPairs returns a boolean if a field has been set.
func (o *AzureCredentials) HasMonitorOnlyTagPairs() bool {
	if o != nil && !IsNil(o.MonitorOnlyTagPairs) {
		return true
	}

	return false
}

// SetMonitorOnlyTagPairs gets a reference to the given []CloudTag and assigns it to the MonitorOnlyTagPairs field.
func (o *AzureCredentials) SetMonitorOnlyTagPairs(v []CloudTag) {
	o.MonitorOnlyTagPairs = v
}

// GetMonitorOnlyTaggedEntities returns the MonitorOnlyTaggedEntities field value
func (o *AzureCredentials) GetMonitorOnlyTaggedEntities() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MonitorOnlyTaggedEntities
}

// GetMonitorOnlyTaggedEntitiesOk returns a tuple with the MonitorOnlyTaggedEntities field value
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetMonitorOnlyTaggedEntitiesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorOnlyTaggedEntities, true
}

// SetMonitorOnlyTaggedEntities sets field value
func (o *AzureCredentials) SetMonitorOnlyTaggedEntities(v bool) {
	o.MonitorOnlyTaggedEntities = v
}

// GetSupportingServices returns the SupportingServices field value if set, zero value otherwise.
func (o *AzureCredentials) GetSupportingServices() []AzureSupportingService {
	if o == nil || IsNil(o.SupportingServices) {
		var ret []AzureSupportingService
		return ret
	}
	return o.SupportingServices
}

// GetSupportingServicesOk returns a tuple with the SupportingServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureCredentials) GetSupportingServicesOk() ([]AzureSupportingService, bool) {
	if o == nil || IsNil(o.SupportingServices) {
		return nil, false
	}
	return o.SupportingServices, true
}

// HasSupportingServices returns a boolean if a field has been set.
func (o *AzureCredentials) HasSupportingServices() bool {
	if o != nil && !IsNil(o.SupportingServices) {
		return true
	}

	return false
}

// SetSupportingServices gets a reference to the given []AzureSupportingService and assigns it to the SupportingServices field.
func (o *AzureCredentials) SetSupportingServices(v []AzureSupportingService) {
	o.SupportingServices = v
}

func (o AzureCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	toSerialize["autoTagging"] = o.AutoTagging
	if !IsNil(o.DirectoryId) {
		toSerialize["directoryId"] = o.DirectoryId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MonitorOnlyExcludingTagPairs) {
		toSerialize["monitorOnlyExcludingTagPairs"] = o.MonitorOnlyExcludingTagPairs
	}
	if !IsNil(o.MonitorOnlyTagPairs) {
		toSerialize["monitorOnlyTagPairs"] = o.MonitorOnlyTagPairs
	}
	toSerialize["monitorOnlyTaggedEntities"] = o.MonitorOnlyTaggedEntities
	if !IsNil(o.SupportingServices) {
		toSerialize["supportingServices"] = o.SupportingServices
	}
	return toSerialize, nil
}

type NullableAzureCredentials struct {
	value *AzureCredentials
	isSet bool
}

func (v NullableAzureCredentials) Get() *AzureCredentials {
	return v.value
}

func (v *NullableAzureCredentials) Set(val *AzureCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureCredentials(val *AzureCredentials) *NullableAzureCredentials {
	return &NullableAzureCredentials{value: val, isSet: true}
}

func (v NullableAzureCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

