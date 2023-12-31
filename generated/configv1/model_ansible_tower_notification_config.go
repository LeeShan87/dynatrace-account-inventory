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

// checks if the AnsibleTowerNotificationConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleTowerNotificationConfig{}

// AnsibleTowerNotificationConfig Configuration of the Ansible Tower notification.
type AnsibleTowerNotificationConfig struct {
	// Accept any, including self-signed and invalid, SSL certificate (`true`) or only trusted (`false`) certificates.
	AcceptAnyCertificate bool `json:"acceptAnyCertificate"`
	// The custom message of the notification.    This message will be displayed in the extra variables **Message** field of your job template.   You can use the following placeholders:  * `{ImpactedEntities}`: Details about the entities impacted by the problem in form of a JSON array.  * `{ImpactedEntity}`: The entity impacted by the problem or *X* impacted entities.  * `{PID}`: The ID of the reported problem.  * `{ProblemDetailsText}`: All problem event details, including root cause, as a text-formatted string.  * `{ProblemID}`: The display number of the reported problem.  * `{ProblemImpact}`: The [impact level](https://dt-url.net/klg3k4q) of the problem. Possible values are `APPLICATION`, `SERVICE`, and `INFRASTRUCTURE`.  * `{ProblemSeverity}`: The [severity level](https://dt-url.net/f1i3k5b) of the problem. Possible values are `AVAILABILITY`, `ERROR`, `PERFORMANCE`, `RESOURCE_CONTENTION`, and `CUSTOM_ALERT`.  * `{ProblemTitle}`: A short description of the problem.  * `{ProblemURL}`: The URL of the problem within Dynatrace.  * `{State}`: The state of the problem. Possible values are `OPEN` and `RESOLVED`.  * `{Tags}`: The list of tags that are defined for all impacted entities, separated by commas.  
	CustomMessage string `json:"customMessage"`
	// The ID of the target Ansible Tower job template.
	JobTemplateID int32 `json:"jobTemplateID"`
	// The URL of the target Ansible Tower job template.
	JobTemplateURL string `json:"jobTemplateURL"`
	// The password for the Ansible Tower account.
	Password *string `json:"password,omitempty"`
	// The username of the Ansible Tower account.
	Username string `json:"username"`
}

// NewAnsibleTowerNotificationConfig instantiates a new AnsibleTowerNotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleTowerNotificationConfig(acceptAnyCertificate bool, customMessage string, jobTemplateID int32, jobTemplateURL string, username string, active bool, alertingProfile string, name string, type_ string) *AnsibleTowerNotificationConfig {
	this := AnsibleTowerNotificationConfig{}
	this.Active = active
	this.AlertingProfile = alertingProfile
	this.Name = name
	this.Type = type_
	this.AcceptAnyCertificate = acceptAnyCertificate
	this.CustomMessage = customMessage
	this.JobTemplateID = jobTemplateID
	this.JobTemplateURL = jobTemplateURL
	this.Username = username
	return &this
}

// NewAnsibleTowerNotificationConfigWithDefaults instantiates a new AnsibleTowerNotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleTowerNotificationConfigWithDefaults() *AnsibleTowerNotificationConfig {
	this := AnsibleTowerNotificationConfig{}
	return &this
}

// GetAcceptAnyCertificate returns the AcceptAnyCertificate field value
func (o *AnsibleTowerNotificationConfig) GetAcceptAnyCertificate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptAnyCertificate
}

// GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field value
// and a boolean to check if the value has been set.
func (o *AnsibleTowerNotificationConfig) GetAcceptAnyCertificateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptAnyCertificate, true
}

// SetAcceptAnyCertificate sets field value
func (o *AnsibleTowerNotificationConfig) SetAcceptAnyCertificate(v bool) {
	o.AcceptAnyCertificate = v
}

// GetCustomMessage returns the CustomMessage field value
func (o *AnsibleTowerNotificationConfig) GetCustomMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomMessage
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value
// and a boolean to check if the value has been set.
func (o *AnsibleTowerNotificationConfig) GetCustomMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomMessage, true
}

// SetCustomMessage sets field value
func (o *AnsibleTowerNotificationConfig) SetCustomMessage(v string) {
	o.CustomMessage = v
}

// GetJobTemplateID returns the JobTemplateID field value
func (o *AnsibleTowerNotificationConfig) GetJobTemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.JobTemplateID
}

// GetJobTemplateIDOk returns a tuple with the JobTemplateID field value
// and a boolean to check if the value has been set.
func (o *AnsibleTowerNotificationConfig) GetJobTemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobTemplateID, true
}

// SetJobTemplateID sets field value
func (o *AnsibleTowerNotificationConfig) SetJobTemplateID(v int32) {
	o.JobTemplateID = v
}

// GetJobTemplateURL returns the JobTemplateURL field value
func (o *AnsibleTowerNotificationConfig) GetJobTemplateURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobTemplateURL
}

// GetJobTemplateURLOk returns a tuple with the JobTemplateURL field value
// and a boolean to check if the value has been set.
func (o *AnsibleTowerNotificationConfig) GetJobTemplateURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobTemplateURL, true
}

// SetJobTemplateURL sets field value
func (o *AnsibleTowerNotificationConfig) SetJobTemplateURL(v string) {
	o.JobTemplateURL = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AnsibleTowerNotificationConfig) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleTowerNotificationConfig) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AnsibleTowerNotificationConfig) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AnsibleTowerNotificationConfig) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value
func (o *AnsibleTowerNotificationConfig) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AnsibleTowerNotificationConfig) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AnsibleTowerNotificationConfig) SetUsername(v string) {
	o.Username = v
}

func (o AnsibleTowerNotificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleTowerNotificationConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acceptAnyCertificate"] = o.AcceptAnyCertificate
	toSerialize["customMessage"] = o.CustomMessage
	toSerialize["jobTemplateID"] = o.JobTemplateID
	toSerialize["jobTemplateURL"] = o.JobTemplateURL
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableAnsibleTowerNotificationConfig struct {
	value *AnsibleTowerNotificationConfig
	isSet bool
}

func (v NullableAnsibleTowerNotificationConfig) Get() *AnsibleTowerNotificationConfig {
	return v.value
}

func (v *NullableAnsibleTowerNotificationConfig) Set(val *AnsibleTowerNotificationConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleTowerNotificationConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleTowerNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleTowerNotificationConfig(val *AnsibleTowerNotificationConfig) *NullableAnsibleTowerNotificationConfig {
	return &NullableAnsibleTowerNotificationConfig{value: val, isSet: true}
}

func (v NullableAnsibleTowerNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleTowerNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


