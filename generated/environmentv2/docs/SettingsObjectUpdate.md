# SettingsObjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsertAfter** | Pointer to **string** | The position of the updated object. The new object will be moved behind the specified one.   **insertAfter** and **insertBefore** are evaluated together and only one of both can be set.   If &#x60;null&#x60; and **insertBefore** &#39;null&#39;, the existing object keeps the current position.   If set to empty string, the updated object will be placed in the first position.   Only applicable for objects based on schemas with ordered objects (schema&#39;s **ordered** parameter is set to &#x60;true&#x60;). | [optional] 
**InsertBefore** | Pointer to **string** | The position of the updated object. The new object will be moved in front of the specified one.   **insertAfter** and **insertBefore** are evaluated together and only one of both can be set.   If &#x60;null&#x60; and **insertAfter** &#39;null&#39;, the existing object keeps the current position.   If set to empty string, the updated object will be placed in the last position.   Only applicable for objects based on schemas with ordered objects (schema&#39;s **ordered** parameter is set to &#x60;true&#x60;). | [optional] 
**SchemaVersion** | Pointer to **string** | The version of the schema on which the object is based. | [optional] 
**UpdateToken** | Pointer to **string** | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn&#39;t any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. | [optional] 
**Value** | **map[string]interface{}** | The value of the setting.    It defines the actual values of settings&#39; parameters.   The actual content depends on the object&#39;s schema. | 

## Methods

### NewSettingsObjectUpdate

`func NewSettingsObjectUpdate(value map[string]interface{}, ) *SettingsObjectUpdate`

NewSettingsObjectUpdate instantiates a new SettingsObjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsObjectUpdateWithDefaults

`func NewSettingsObjectUpdateWithDefaults() *SettingsObjectUpdate`

NewSettingsObjectUpdateWithDefaults instantiates a new SettingsObjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsertAfter

`func (o *SettingsObjectUpdate) GetInsertAfter() string`

GetInsertAfter returns the InsertAfter field if non-nil, zero value otherwise.

### GetInsertAfterOk

`func (o *SettingsObjectUpdate) GetInsertAfterOk() (*string, bool)`

GetInsertAfterOk returns a tuple with the InsertAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertAfter

`func (o *SettingsObjectUpdate) SetInsertAfter(v string)`

SetInsertAfter sets InsertAfter field to given value.

### HasInsertAfter

`func (o *SettingsObjectUpdate) HasInsertAfter() bool`

HasInsertAfter returns a boolean if a field has been set.

### GetInsertBefore

`func (o *SettingsObjectUpdate) GetInsertBefore() string`

GetInsertBefore returns the InsertBefore field if non-nil, zero value otherwise.

### GetInsertBeforeOk

`func (o *SettingsObjectUpdate) GetInsertBeforeOk() (*string, bool)`

GetInsertBeforeOk returns a tuple with the InsertBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertBefore

`func (o *SettingsObjectUpdate) SetInsertBefore(v string)`

SetInsertBefore sets InsertBefore field to given value.

### HasInsertBefore

`func (o *SettingsObjectUpdate) HasInsertBefore() bool`

HasInsertBefore returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *SettingsObjectUpdate) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SettingsObjectUpdate) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SettingsObjectUpdate) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SettingsObjectUpdate) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetUpdateToken

`func (o *SettingsObjectUpdate) GetUpdateToken() string`

GetUpdateToken returns the UpdateToken field if non-nil, zero value otherwise.

### GetUpdateTokenOk

`func (o *SettingsObjectUpdate) GetUpdateTokenOk() (*string, bool)`

GetUpdateTokenOk returns a tuple with the UpdateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateToken

`func (o *SettingsObjectUpdate) SetUpdateToken(v string)`

SetUpdateToken sets UpdateToken field to given value.

### HasUpdateToken

`func (o *SettingsObjectUpdate) HasUpdateToken() bool`

HasUpdateToken returns a boolean if a field has been set.

### GetValue

`func (o *SettingsObjectUpdate) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingsObjectUpdate) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingsObjectUpdate) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


