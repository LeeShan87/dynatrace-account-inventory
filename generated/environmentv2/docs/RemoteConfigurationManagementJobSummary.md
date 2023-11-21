# RemoteConfigurationManagementJobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | Date (in ISO 8601 format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;) when the remote configuration management job was finished. This field is present only for finished jobs. | [optional] 
**EntityType** | Pointer to **string** | Type of entities modified by remote configuration management. | [optional] 
**Id** | Pointer to **string** | The ID of the remote configuration management job. | [optional] 
**StartTime** | Pointer to **string** | Date (in ISO 8601 format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;) when the remote configuration management job was started. | [optional] 

## Methods

### NewRemoteConfigurationManagementJobSummary

`func NewRemoteConfigurationManagementJobSummary() *RemoteConfigurationManagementJobSummary`

NewRemoteConfigurationManagementJobSummary instantiates a new RemoteConfigurationManagementJobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementJobSummaryWithDefaults

`func NewRemoteConfigurationManagementJobSummaryWithDefaults() *RemoteConfigurationManagementJobSummary`

NewRemoteConfigurationManagementJobSummaryWithDefaults instantiates a new RemoteConfigurationManagementJobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *RemoteConfigurationManagementJobSummary) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RemoteConfigurationManagementJobSummary) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RemoteConfigurationManagementJobSummary) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *RemoteConfigurationManagementJobSummary) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityType

`func (o *RemoteConfigurationManagementJobSummary) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *RemoteConfigurationManagementJobSummary) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *RemoteConfigurationManagementJobSummary) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *RemoteConfigurationManagementJobSummary) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetId

`func (o *RemoteConfigurationManagementJobSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteConfigurationManagementJobSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteConfigurationManagementJobSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteConfigurationManagementJobSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *RemoteConfigurationManagementJobSummary) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RemoteConfigurationManagementJobSummary) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RemoteConfigurationManagementJobSummary) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RemoteConfigurationManagementJobSummary) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


