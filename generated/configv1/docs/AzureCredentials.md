# AzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | The monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).   If not set on creation, the &#x60;true&#x60; value is used.   If the field is omitted during an update, the old value remains unaffected. | [optional] 
**AppId** | Pointer to **string** | The application ID (also referred to as client ID).   The field is **required** when creating a new credentials configuration.    The field is ignored during an update, the old value remains unaffected. | [optional] 
**AutoTagging** | **bool** | The automatic capture of Azure tags is on (&#x60;true&#x60;) or off (&#x60;false&#x60;). | 
**DirectoryId** | Pointer to **string** | The directory ID (also referred to as tenant ID).   The field is **required** when creating a new credentials configuration.    The field is ignored during an update, the old value remains unaffected. | [optional] 
**Id** | Pointer to **string** | The Dynatrace entity ID of the Azure credentials configuration. | [optional] [readonly] 
**Key** | Pointer to **string** | The secret key associated with the application ID.   For security reasons, GET requests return this field as &#x60;null&#x60;.    Submit your key on creation or update of the configuration.    The field is **required** when creating a new credentials configuration. If the field is omitted during an update, the old value remains unaffected. | [optional] 
**Label** | **string** | The unique name of the Azure credentials configuration.   Allowed characters are letters, numbers, and spaces. Also the special characters &#x60;.+-_&#x60; are allowed. | 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**MonitorOnlyExcludingTagPairs** | Pointer to [**[]CloudTag**](CloudTag.md) | A list of Azure tags to be excluded from monitoring.   You can specify up to 20 tags. A resource tagged with *any* of the specified tags will not be monitored.   Only applicable when the **monitorOnlyTaggedEntities** parameter is set to &#x60;true&#x60;. | [optional] 
**MonitorOnlyTagPairs** | Pointer to [**[]CloudTag**](CloudTag.md) | A list of Azure tags to be monitored.   You can specify up to 20 tags. A resource tagged with *any* of the specified tags is monitored.   Only applicable when the **monitorOnlyTaggedEntities** parameter is set to &#x60;true&#x60;. | [optional] 
**MonitorOnlyTaggedEntities** | **bool** | Monitor only resources that have specified Azure tags (&#x60;true&#x60;) or all resources (&#x60;false&#x60;). | 
**SupportingServices** | Pointer to [**[]AzureSupportingService**](AzureSupportingService.md) | **Deprecated**. To manage services use [/azure/credentials/{id}/services](https://dt-url.net/1w62s27) operation. Built-in services are not supported here.  A list of Azure services to be monitored. Available services are listed by [/azure/supportedServices](https://dt-url.net/wt42sdq) operation.  For each service, a list of metrics and dimensions can be specified. A list of supported metrics and dimensions for a given service can be checked in [documentation](https://dt-url.net/kx2351b).  List of metrics can be skipped (set to null), resulting in recommended (default) set of metrics and dimensions being chosen for monitoring.  | [optional] 

## Methods

### NewAzureCredentials

`func NewAzureCredentials(autoTagging bool, label string, monitorOnlyTaggedEntities bool, ) *AzureCredentials`

NewAzureCredentials instantiates a new AzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureCredentialsWithDefaults

`func NewAzureCredentialsWithDefaults() *AzureCredentials`

NewAzureCredentialsWithDefaults instantiates a new AzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *AzureCredentials) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AzureCredentials) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AzureCredentials) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AzureCredentials) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAppId

`func (o *AzureCredentials) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AzureCredentials) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AzureCredentials) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AzureCredentials) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetAutoTagging

`func (o *AzureCredentials) GetAutoTagging() bool`

GetAutoTagging returns the AutoTagging field if non-nil, zero value otherwise.

### GetAutoTaggingOk

`func (o *AzureCredentials) GetAutoTaggingOk() (*bool, bool)`

GetAutoTaggingOk returns a tuple with the AutoTagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTagging

`func (o *AzureCredentials) SetAutoTagging(v bool)`

SetAutoTagging sets AutoTagging field to given value.


### GetDirectoryId

`func (o *AzureCredentials) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *AzureCredentials) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *AzureCredentials) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *AzureCredentials) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### GetId

`func (o *AzureCredentials) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzureCredentials) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzureCredentials) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AzureCredentials) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AzureCredentials) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AzureCredentials) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AzureCredentials) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AzureCredentials) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *AzureCredentials) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AzureCredentials) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AzureCredentials) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetadata

`func (o *AzureCredentials) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AzureCredentials) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AzureCredentials) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AzureCredentials) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonitorOnlyExcludingTagPairs

`func (o *AzureCredentials) GetMonitorOnlyExcludingTagPairs() []CloudTag`

GetMonitorOnlyExcludingTagPairs returns the MonitorOnlyExcludingTagPairs field if non-nil, zero value otherwise.

### GetMonitorOnlyExcludingTagPairsOk

`func (o *AzureCredentials) GetMonitorOnlyExcludingTagPairsOk() (*[]CloudTag, bool)`

GetMonitorOnlyExcludingTagPairsOk returns a tuple with the MonitorOnlyExcludingTagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorOnlyExcludingTagPairs

`func (o *AzureCredentials) SetMonitorOnlyExcludingTagPairs(v []CloudTag)`

SetMonitorOnlyExcludingTagPairs sets MonitorOnlyExcludingTagPairs field to given value.

### HasMonitorOnlyExcludingTagPairs

`func (o *AzureCredentials) HasMonitorOnlyExcludingTagPairs() bool`

HasMonitorOnlyExcludingTagPairs returns a boolean if a field has been set.

### GetMonitorOnlyTagPairs

`func (o *AzureCredentials) GetMonitorOnlyTagPairs() []CloudTag`

GetMonitorOnlyTagPairs returns the MonitorOnlyTagPairs field if non-nil, zero value otherwise.

### GetMonitorOnlyTagPairsOk

`func (o *AzureCredentials) GetMonitorOnlyTagPairsOk() (*[]CloudTag, bool)`

GetMonitorOnlyTagPairsOk returns a tuple with the MonitorOnlyTagPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorOnlyTagPairs

`func (o *AzureCredentials) SetMonitorOnlyTagPairs(v []CloudTag)`

SetMonitorOnlyTagPairs sets MonitorOnlyTagPairs field to given value.

### HasMonitorOnlyTagPairs

`func (o *AzureCredentials) HasMonitorOnlyTagPairs() bool`

HasMonitorOnlyTagPairs returns a boolean if a field has been set.

### GetMonitorOnlyTaggedEntities

`func (o *AzureCredentials) GetMonitorOnlyTaggedEntities() bool`

GetMonitorOnlyTaggedEntities returns the MonitorOnlyTaggedEntities field if non-nil, zero value otherwise.

### GetMonitorOnlyTaggedEntitiesOk

`func (o *AzureCredentials) GetMonitorOnlyTaggedEntitiesOk() (*bool, bool)`

GetMonitorOnlyTaggedEntitiesOk returns a tuple with the MonitorOnlyTaggedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorOnlyTaggedEntities

`func (o *AzureCredentials) SetMonitorOnlyTaggedEntities(v bool)`

SetMonitorOnlyTaggedEntities sets MonitorOnlyTaggedEntities field to given value.


### GetSupportingServices

`func (o *AzureCredentials) GetSupportingServices() []AzureSupportingService`

GetSupportingServices returns the SupportingServices field if non-nil, zero value otherwise.

### GetSupportingServicesOk

`func (o *AzureCredentials) GetSupportingServicesOk() (*[]AzureSupportingService, bool)`

GetSupportingServicesOk returns a tuple with the SupportingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportingServices

`func (o *AzureCredentials) SetSupportingServices(v []AzureSupportingService)`

SetSupportingServices sets SupportingServices field to given value.

### HasSupportingServices

`func (o *AzureCredentials) HasSupportingServices() bool`

HasSupportingServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


