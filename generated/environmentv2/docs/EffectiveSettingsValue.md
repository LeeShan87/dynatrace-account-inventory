# EffectiveSettingsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | The user (identified by a user ID or a public token ID) who performed that most recent modification. | [optional] 
**Created** | Pointer to **int64** | The timestamp of the creation. | [optional] 
**ExternalId** | Pointer to **string** | The external identifier of the settings object. | [optional] 
**Modified** | Pointer to **int64** | The timestamp of the last modification. | [optional] 
**Origin** | Pointer to **string** | The origin of the settings value. | [optional] 
**SchemaId** | Pointer to **string** | The schema on which the object is based. | [optional] 
**SchemaVersion** | Pointer to **string** | The version of the schema on which the object is based. | [optional] 
**SearchSummary** | Pointer to **string** | A searchable summary string of the setting value. Plain text without Markdown. | [optional] 
**Summary** | Pointer to **string** | A short summary of settings. This can contain Markdown and will be escaped accordingly. | [optional] 
**Value** | Pointer to **map[string]interface{}** | The value of the setting.    It defines the actual values of settings&#39; parameters.   The actual content depends on the object&#39;s schema. | [optional] 

## Methods

### NewEffectiveSettingsValue

`func NewEffectiveSettingsValue() *EffectiveSettingsValue`

NewEffectiveSettingsValue instantiates a new EffectiveSettingsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectiveSettingsValueWithDefaults

`func NewEffectiveSettingsValueWithDefaults() *EffectiveSettingsValue`

NewEffectiveSettingsValueWithDefaults instantiates a new EffectiveSettingsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *EffectiveSettingsValue) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *EffectiveSettingsValue) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *EffectiveSettingsValue) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *EffectiveSettingsValue) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreated

`func (o *EffectiveSettingsValue) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EffectiveSettingsValue) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EffectiveSettingsValue) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EffectiveSettingsValue) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExternalId

`func (o *EffectiveSettingsValue) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EffectiveSettingsValue) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EffectiveSettingsValue) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EffectiveSettingsValue) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetModified

`func (o *EffectiveSettingsValue) GetModified() int64`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *EffectiveSettingsValue) GetModifiedOk() (*int64, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *EffectiveSettingsValue) SetModified(v int64)`

SetModified sets Modified field to given value.

### HasModified

`func (o *EffectiveSettingsValue) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetOrigin

`func (o *EffectiveSettingsValue) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *EffectiveSettingsValue) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *EffectiveSettingsValue) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *EffectiveSettingsValue) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSchemaId

`func (o *EffectiveSettingsValue) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *EffectiveSettingsValue) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *EffectiveSettingsValue) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *EffectiveSettingsValue) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *EffectiveSettingsValue) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *EffectiveSettingsValue) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *EffectiveSettingsValue) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *EffectiveSettingsValue) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSearchSummary

`func (o *EffectiveSettingsValue) GetSearchSummary() string`

GetSearchSummary returns the SearchSummary field if non-nil, zero value otherwise.

### GetSearchSummaryOk

`func (o *EffectiveSettingsValue) GetSearchSummaryOk() (*string, bool)`

GetSearchSummaryOk returns a tuple with the SearchSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSummary

`func (o *EffectiveSettingsValue) SetSearchSummary(v string)`

SetSearchSummary sets SearchSummary field to given value.

### HasSearchSummary

`func (o *EffectiveSettingsValue) HasSearchSummary() bool`

HasSearchSummary returns a boolean if a field has been set.

### GetSummary

`func (o *EffectiveSettingsValue) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *EffectiveSettingsValue) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *EffectiveSettingsValue) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *EffectiveSettingsValue) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetValue

`func (o *EffectiveSettingsValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EffectiveSettingsValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EffectiveSettingsValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *EffectiveSettingsValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


