# NotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | The configuration is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**AlertingProfile** | **string** | The ID of the associated alerting profile. | 
**Id** | Pointer to **string** | The ID of the notification configuration. | [optional] 
**Name** | **string** | The name of the notification configuration. | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;EMAIL&#x60; -&gt; EmailNotificationConfig  * &#x60;PAGER_DUTY&#x60; -&gt; PagerDutyNotificationConfig  * &#x60;WEBHOOK&#x60; -&gt; WebHookNotificationConfig  * &#x60;SLACK&#x60; -&gt; SlackNotificationConfig  * &#x60;VICTOROPS&#x60; -&gt; VictorOpsNotificationConfig  * &#x60;SERVICE_NOW&#x60; -&gt; ServiceNowNotificationConfig  * &#x60;XMATTERS&#x60; -&gt; XMattersNotificationConfig  * &#x60;ANSIBLETOWER&#x60; -&gt; AnsibleTowerNotificationConfig  * &#x60;OPS_GENIE&#x60; -&gt; OpsGenieNotificationConfig  * &#x60;JIRA&#x60; -&gt; JiraNotificationConfig  * &#x60;TRELLO&#x60; -&gt; TrelloNotificationConfig   | 

## Methods

### NewNotificationConfig

`func NewNotificationConfig(active bool, alertingProfile string, name string, type_ string, ) *NotificationConfig`

NewNotificationConfig instantiates a new NotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationConfigWithDefaults

`func NewNotificationConfigWithDefaults() *NotificationConfig`

NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *NotificationConfig) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *NotificationConfig) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *NotificationConfig) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAlertingProfile

`func (o *NotificationConfig) GetAlertingProfile() string`

GetAlertingProfile returns the AlertingProfile field if non-nil, zero value otherwise.

### GetAlertingProfileOk

`func (o *NotificationConfig) GetAlertingProfileOk() (*string, bool)`

GetAlertingProfileOk returns a tuple with the AlertingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingProfile

`func (o *NotificationConfig) SetAlertingProfile(v string)`

SetAlertingProfile sets AlertingProfile field to given value.


### GetId

`func (o *NotificationConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationConfig) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *NotificationConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


