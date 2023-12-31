# SettingsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | The user (identified by a user ID or a public token ID) who performed that most recent modification. | [optional] 
**Created** | Pointer to **int64** | The timestamp of the creation. | [optional] 
**CreatedBy** | Pointer to **string** | The unique identifier of the user who created the settings object. | [optional] 
**ExternalId** | Pointer to **string** | The external identifier of the settings object. | [optional] 
**ModificationInfo** | Pointer to [**ModificationInfo**](ModificationInfo.md) |  | [optional] 
**Modified** | Pointer to **int64** | The timestamp of the last modification. | [optional] 
**ModifiedBy** | Pointer to **string** | The unique identifier of the user who performed the most recent modification. | [optional] 
**ObjectId** | Pointer to **string** | The ID of the settings object. | [optional] 
**SchemaId** | Pointer to **string** | The schema on which the object is based. | [optional] 
**SchemaVersion** | Pointer to **string** | The version of the schema on which the object is based. | [optional] 
**Scope** | Pointer to **string** | The scope that the object targets. For more details, please see [Dynatrace Documentation](https://dt-url.net/ky03459).  | [optional] 
**SearchSummary** | Pointer to **string** | A searchable summary string of the setting value. Plain text without Markdown. | [optional] 
**Summary** | Pointer to **string** | A short summary of settings. This can contain Markdown and will be escaped accordingly. | [optional] 
**UpdateToken** | Pointer to **string** | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn&#39;t any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value of the setting.    It defines the actual values of settings&#39; parameters.   The actual content depends on the object&#39;s schema. | [optional] 

## Methods

### NewSettingsObject

`func NewSettingsObject() *SettingsObject`

NewSettingsObject instantiates a new SettingsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsObjectWithDefaults

`func NewSettingsObjectWithDefaults() *SettingsObject`

NewSettingsObjectWithDefaults instantiates a new SettingsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *SettingsObject) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *SettingsObject) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *SettingsObject) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *SettingsObject) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *SettingsObject) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SettingsObject) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SettingsObject) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SettingsObject) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *SettingsObject) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SettingsObject) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SettingsObject) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SettingsObject) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetExternalId

`func (o *SettingsObject) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SettingsObject) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SettingsObject) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SettingsObject) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetModificationInfo

`func (o *SettingsObject) GetModificationInfo() ModificationInfo`

GetModificationInfo returns the ModificationInfo field if non-nil, zero value otherwise.

### GetModificationInfoOk

`func (o *SettingsObject) GetModificationInfoOk() (*ModificationInfo, bool)`

GetModificationInfoOk returns a tuple with the ModificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationInfo

`func (o *SettingsObject) SetModificationInfo(v ModificationInfo)`

SetModificationInfo sets ModificationInfo field to given value.

### HasModificationInfo

`func (o *SettingsObject) HasModificationInfo() bool`

HasModificationInfo returns a boolean if a field has been set.

### GetModified

`func (o *SettingsObject) GetModified() int64`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *SettingsObject) GetModifiedOk() (*int64, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *SettingsObject) SetModified(v int64)`

SetModified sets Modified field to given value.

### HasModified

`func (o *SettingsObject) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetModifiedBy

`func (o *SettingsObject) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SettingsObject) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SettingsObject) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *SettingsObject) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetObjectId

`func (o *SettingsObject) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *SettingsObject) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *SettingsObject) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *SettingsObject) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSchemaId

`func (o *SettingsObject) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SettingsObject) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SettingsObject) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *SettingsObject) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *SettingsObject) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SettingsObject) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SettingsObject) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SettingsObject) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetScope

`func (o *SettingsObject) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SettingsObject) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SettingsObject) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SettingsObject) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSearchSummary

`func (o *SettingsObject) GetSearchSummary() string`

GetSearchSummary returns the SearchSummary field if non-nil, zero value otherwise.

### GetSearchSummaryOk

`func (o *SettingsObject) GetSearchSummaryOk() (*string, bool)`

GetSearchSummaryOk returns a tuple with the SearchSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSummary

`func (o *SettingsObject) SetSearchSummary(v string)`

SetSearchSummary sets SearchSummary field to given value.

### HasSearchSummary

`func (o *SettingsObject) HasSearchSummary() bool`

HasSearchSummary returns a boolean if a field has been set.

### GetSummary

`func (o *SettingsObject) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SettingsObject) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SettingsObject) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *SettingsObject) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdateToken

`func (o *SettingsObject) GetUpdateToken() string`

GetUpdateToken returns the UpdateToken field if non-nil, zero value otherwise.

### GetUpdateTokenOk

`func (o *SettingsObject) GetUpdateTokenOk() (*string, bool)`

GetUpdateTokenOk returns a tuple with the UpdateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateToken

`func (o *SettingsObject) SetUpdateToken(v string)`

SetUpdateToken sets UpdateToken field to given value.

### HasUpdateToken

`func (o *SettingsObject) HasUpdateToken() bool`

HasUpdateToken returns a boolean if a field has been set.

### GetValue

`func (o *SettingsObject) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingsObject) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingsObject) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SettingsObject) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


