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

// checks if the OpaqueAndExternalWebRequestRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpaqueAndExternalWebRequestRule{}

// OpaqueAndExternalWebRequestRule The service detection rule of the `OPAQUE_AND_EXTERNAL_WEB_REQUEST` type.
type OpaqueAndExternalWebRequestRule struct {
	ApplicationId *ApplicationId `json:"applicationId,omitempty"`
	// A list of conditions of the rule.   If several conditions are specified, the AND logic applies.
	Conditions []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto `json:"conditions,omitempty"`
	ContextRoot *ContextRoot `json:"contextRoot,omitempty"`
	// A short description of the rule.
	Description *string `json:"description,omitempty"`
	// The rule is enabled(`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The ID of the service detection rule.
	Id *string `json:"id,omitempty"`
	// The management zone (specified by the ID) of the process group for which this service detection rule should be created.    You can specify only 1 management zone here.
	ManagementZones []string `json:"managementZones,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The name of the rule.
	Name string `json:"name"`
	// The order of the rule in the rules list.    The rules are evaluated from top to bottom. The first matching rule applies.
	Order *string `json:"order,omitempty"`
	Port *Port `json:"port,omitempty"`
	PublicDomainName *PublicDomainName `json:"publicDomainName,omitempty"`
	// The type of the service detection rule.
	Type string `json:"type"`
}

// NewOpaqueAndExternalWebRequestRule instantiates a new OpaqueAndExternalWebRequestRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpaqueAndExternalWebRequestRule(enabled bool, name string, type_ string) *OpaqueAndExternalWebRequestRule {
	this := OpaqueAndExternalWebRequestRule{}
	this.Enabled = enabled
	this.Name = name
	this.Type = type_
	return &this
}

// NewOpaqueAndExternalWebRequestRuleWithDefaults instantiates a new OpaqueAndExternalWebRequestRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpaqueAndExternalWebRequestRuleWithDefaults() *OpaqueAndExternalWebRequestRule {
	this := OpaqueAndExternalWebRequestRule{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetApplicationId() ApplicationId {
	if o == nil || IsNil(o.ApplicationId) {
		var ret ApplicationId
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetApplicationIdOk() (*ApplicationId, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given ApplicationId and assigns it to the ApplicationId field.
func (o *OpaqueAndExternalWebRequestRule) SetApplicationId(v ApplicationId) {
	o.ApplicationId = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetConditions() []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto {
	if o == nil || IsNil(o.Conditions) {
		var ret []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetConditionsOk() ([]ConditionsOpaqueAndExternalWebRequestAttributeTypeDto, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto and assigns it to the Conditions field.
func (o *OpaqueAndExternalWebRequestRule) SetConditions(v []ConditionsOpaqueAndExternalWebRequestAttributeTypeDto) {
	o.Conditions = v
}

// GetContextRoot returns the ContextRoot field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetContextRoot() ContextRoot {
	if o == nil || IsNil(o.ContextRoot) {
		var ret ContextRoot
		return ret
	}
	return *o.ContextRoot
}

// GetContextRootOk returns a tuple with the ContextRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetContextRootOk() (*ContextRoot, bool) {
	if o == nil || IsNil(o.ContextRoot) {
		return nil, false
	}
	return o.ContextRoot, true
}

// HasContextRoot returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasContextRoot() bool {
	if o != nil && !IsNil(o.ContextRoot) {
		return true
	}

	return false
}

// SetContextRoot gets a reference to the given ContextRoot and assigns it to the ContextRoot field.
func (o *OpaqueAndExternalWebRequestRule) SetContextRoot(v ContextRoot) {
	o.ContextRoot = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpaqueAndExternalWebRequestRule) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *OpaqueAndExternalWebRequestRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *OpaqueAndExternalWebRequestRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OpaqueAndExternalWebRequestRule) SetId(v string) {
	o.Id = &v
}

// GetManagementZones returns the ManagementZones field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetManagementZones() []string {
	if o == nil || IsNil(o.ManagementZones) {
		var ret []string
		return ret
	}
	return o.ManagementZones
}

// GetManagementZonesOk returns a tuple with the ManagementZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetManagementZonesOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagementZones) {
		return nil, false
	}
	return o.ManagementZones, true
}

// HasManagementZones returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasManagementZones() bool {
	if o != nil && !IsNil(o.ManagementZones) {
		return true
	}

	return false
}

// SetManagementZones gets a reference to the given []string and assigns it to the ManagementZones field.
func (o *OpaqueAndExternalWebRequestRule) SetManagementZones(v []string) {
	o.ManagementZones = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *OpaqueAndExternalWebRequestRule) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *OpaqueAndExternalWebRequestRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OpaqueAndExternalWebRequestRule) SetName(v string) {
	o.Name = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *OpaqueAndExternalWebRequestRule) SetOrder(v string) {
	o.Order = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetPort() Port {
	if o == nil || IsNil(o.Port) {
		var ret Port
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetPortOk() (*Port, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given Port and assigns it to the Port field.
func (o *OpaqueAndExternalWebRequestRule) SetPort(v Port) {
	o.Port = &v
}

// GetPublicDomainName returns the PublicDomainName field value if set, zero value otherwise.
func (o *OpaqueAndExternalWebRequestRule) GetPublicDomainName() PublicDomainName {
	if o == nil || IsNil(o.PublicDomainName) {
		var ret PublicDomainName
		return ret
	}
	return *o.PublicDomainName
}

// GetPublicDomainNameOk returns a tuple with the PublicDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetPublicDomainNameOk() (*PublicDomainName, bool) {
	if o == nil || IsNil(o.PublicDomainName) {
		return nil, false
	}
	return o.PublicDomainName, true
}

// HasPublicDomainName returns a boolean if a field has been set.
func (o *OpaqueAndExternalWebRequestRule) HasPublicDomainName() bool {
	if o != nil && !IsNil(o.PublicDomainName) {
		return true
	}

	return false
}

// SetPublicDomainName gets a reference to the given PublicDomainName and assigns it to the PublicDomainName field.
func (o *OpaqueAndExternalWebRequestRule) SetPublicDomainName(v PublicDomainName) {
	o.PublicDomainName = &v
}

// GetType returns the Type field value
func (o *OpaqueAndExternalWebRequestRule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OpaqueAndExternalWebRequestRule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OpaqueAndExternalWebRequestRule) SetType(v string) {
	o.Type = v
}

func (o OpaqueAndExternalWebRequestRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpaqueAndExternalWebRequestRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.ContextRoot) {
		toSerialize["contextRoot"] = o.ContextRoot
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ManagementZones) {
		toSerialize["managementZones"] = o.ManagementZones
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.PublicDomainName) {
		toSerialize["publicDomainName"] = o.PublicDomainName
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableOpaqueAndExternalWebRequestRule struct {
	value *OpaqueAndExternalWebRequestRule
	isSet bool
}

func (v NullableOpaqueAndExternalWebRequestRule) Get() *OpaqueAndExternalWebRequestRule {
	return v.value
}

func (v *NullableOpaqueAndExternalWebRequestRule) Set(val *OpaqueAndExternalWebRequestRule) {
	v.value = val
	v.isSet = true
}

func (v NullableOpaqueAndExternalWebRequestRule) IsSet() bool {
	return v.isSet
}

func (v *NullableOpaqueAndExternalWebRequestRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpaqueAndExternalWebRequestRule(val *OpaqueAndExternalWebRequestRule) *NullableOpaqueAndExternalWebRequestRule {
	return &NullableOpaqueAndExternalWebRequestRule{value: val, isSet: true}
}

func (v NullableOpaqueAndExternalWebRequestRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpaqueAndExternalWebRequestRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


