# WebHookNotificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptAnyCertificate** | **bool** | Accept any, including self-signed and invalid, SSL certificate (&#x60;true&#x60;) or only trusted (&#x60;false&#x60;) certificates. | 
**Headers** | Pointer to [**[]HttpHeader**](HttpHeader.md) | A list of the additional HTTP headers. | [optional] 
**NotifyEventMergesEnabled** | Pointer to **bool** | Call webhook if new events merge into existing problems. | [optional] 
**Payload** | **string** | The content of the notification message.   You can use the following placeholders:  * &#x60;{ImpactedEntities}&#x60;: Details about the entities impacted by the problem in form of a JSON array.  * &#x60;{ImpactedEntity}&#x60;: The entity impacted by the problem or *X* impacted entities.  * &#x60;{PID}&#x60;: The ID of the reported problem.  * &#x60;{ProblemDetailsHTML}&#x60;: All problem event details, including root cause, as an HTML-formatted string.  * &#x60;{ProblemDetailsJSON}&#x60;: All problem event details, including root cause, as a JSON object.  * &#x60;{ProblemDetailsMarkdown}&#x60;: All problem event details, including root cause, as a [Markdown-formatted](https://dt-url.net/1yk3kkq) string.  * &#x60;{ProblemDetailsText}&#x60;: All problem event details, including root cause, as a text-formatted string.  * &#x60;{ProblemID}&#x60;: The display number of the reported problem.  * &#x60;{ProblemImpact}&#x60;: The [impact level](https://dt-url.net/klg3k4q) of the problem. Possible values are &#x60;APPLICATION&#x60;, &#x60;SERVICE&#x60;, and &#x60;INFRASTRUCTURE&#x60;.  * &#x60;{ProblemSeverity}&#x60;: The [severity level](https://dt-url.net/f1i3k5b) of the problem. Possible values are &#x60;AVAILABILITY&#x60;, &#x60;ERROR&#x60;, &#x60;PERFORMANCE&#x60;, &#x60;RESOURCE_CONTENTION&#x60;, and &#x60;CUSTOM_ALERT&#x60;.  * &#x60;{ProblemTitle}&#x60;: A short description of the problem.  * &#x60;{ProblemURL}&#x60;: The URL of the problem within Dynatrace.  * &#x60;{State}&#x60;: The state of the problem. Possible values are &#x60;OPEN&#x60; and &#x60;RESOLVED&#x60;.  * &#x60;{Tags}&#x60;: The list of tags that are defined for all impacted entities, separated by commas.   | 
**Url** | **string** | The URL of the WebHook endpoint. | 

## Methods

### NewWebHookNotificationConfig

`func NewWebHookNotificationConfig(acceptAnyCertificate bool, payload string, url string, ) *WebHookNotificationConfig`

NewWebHookNotificationConfig instantiates a new WebHookNotificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebHookNotificationConfigWithDefaults

`func NewWebHookNotificationConfigWithDefaults() *WebHookNotificationConfig`

NewWebHookNotificationConfigWithDefaults instantiates a new WebHookNotificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptAnyCertificate

`func (o *WebHookNotificationConfig) GetAcceptAnyCertificate() bool`

GetAcceptAnyCertificate returns the AcceptAnyCertificate field if non-nil, zero value otherwise.

### GetAcceptAnyCertificateOk

`func (o *WebHookNotificationConfig) GetAcceptAnyCertificateOk() (*bool, bool)`

GetAcceptAnyCertificateOk returns a tuple with the AcceptAnyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAnyCertificate

`func (o *WebHookNotificationConfig) SetAcceptAnyCertificate(v bool)`

SetAcceptAnyCertificate sets AcceptAnyCertificate field to given value.


### GetHeaders

`func (o *WebHookNotificationConfig) GetHeaders() []HttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebHookNotificationConfig) GetHeadersOk() (*[]HttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebHookNotificationConfig) SetHeaders(v []HttpHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebHookNotificationConfig) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetNotifyEventMergesEnabled

`func (o *WebHookNotificationConfig) GetNotifyEventMergesEnabled() bool`

GetNotifyEventMergesEnabled returns the NotifyEventMergesEnabled field if non-nil, zero value otherwise.

### GetNotifyEventMergesEnabledOk

`func (o *WebHookNotificationConfig) GetNotifyEventMergesEnabledOk() (*bool, bool)`

GetNotifyEventMergesEnabledOk returns a tuple with the NotifyEventMergesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEventMergesEnabled

`func (o *WebHookNotificationConfig) SetNotifyEventMergesEnabled(v bool)`

SetNotifyEventMergesEnabled sets NotifyEventMergesEnabled field to given value.

### HasNotifyEventMergesEnabled

`func (o *WebHookNotificationConfig) HasNotifyEventMergesEnabled() bool`

HasNotifyEventMergesEnabled returns a boolean if a field has been set.

### GetPayload

`func (o *WebHookNotificationConfig) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebHookNotificationConfig) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebHookNotificationConfig) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetUrl

`func (o *WebHookNotificationConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebHookNotificationConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebHookNotificationConfig) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


