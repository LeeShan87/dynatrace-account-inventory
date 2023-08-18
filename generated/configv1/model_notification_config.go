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

// checks if the NotificationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationConfig{}

// NotificationConfig Configuration of a notification. The actual set of fields depends on the `type` of the notification. See [Notifications API - JSON models](https://dt-url.net/9qm3k5u) in Dynatrace Documentation for example models of every notification type.
type NotificationConfig struct {
	// The configuration is enabled (`true`) or disabled (`false`).
	Active bool `json:"active"`
	// The ID of the associated alerting profile.
	AlertingProfile string `json:"alertingProfile"`
	// The ID of the notification configuration.
	Id *string `json:"id,omitempty"`
	// The name of the notification configuration.
	Name string `json:"name"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `EMAIL` -> EmailNotificationConfig  * `PAGER_DUTY` -> PagerDutyNotificationConfig  * `WEBHOOK` -> WebHookNotificationConfig  * `SLACK` -> SlackNotificationConfig  * `VICTOROPS` -> VictorOpsNotificationConfig  * `SERVICE_NOW` -> ServiceNowNotificationConfig  * `XMATTERS` -> XMattersNotificationConfig  * `ANSIBLETOWER` -> AnsibleTowerNotificationConfig  * `OPS_GENIE` -> OpsGenieNotificationConfig  * `JIRA` -> JiraNotificationConfig  * `TRELLO` -> TrelloNotificationConfig  
	Type string `json:"type"`
}

// NewNotificationConfig instantiates a new NotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationConfig(active bool, alertingProfile string, name string, type_ string) *NotificationConfig {
	this := NotificationConfig{}
	this.Active = active
	this.AlertingProfile = alertingProfile
	this.Name = name
	this.Type = type_
	return &this
}

// NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationConfigWithDefaults() *NotificationConfig {
	this := NotificationConfig{}
	return &this
}

// GetActive returns the Active field value
func (o *NotificationConfig) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *NotificationConfig) SetActive(v bool) {
	o.Active = v
}

// GetAlertingProfile returns the AlertingProfile field value
func (o *NotificationConfig) GetAlertingProfile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlertingProfile
}

// GetAlertingProfileOk returns a tuple with the AlertingProfile field value
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetAlertingProfileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertingProfile, true
}

// SetAlertingProfile sets field value
func (o *NotificationConfig) SetAlertingProfile(v string) {
	o.AlertingProfile = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NotificationConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NotificationConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NotificationConfig) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *NotificationConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NotificationConfig) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *NotificationConfig) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationConfig) SetType(v string) {
	o.Type = v
}

func (o NotificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["alertingProfile"] = o.AlertingProfile
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableNotificationConfig struct {
	value *NotificationConfig
	isSet bool
}

func (v NullableNotificationConfig) Get() *NotificationConfig {
	return v.value
}

func (v *NullableNotificationConfig) Set(val *NotificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationConfig(val *NotificationConfig) *NullableNotificationConfig {
	return &NullableNotificationConfig{value: val, isSet: true}
}

func (v NullableNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

