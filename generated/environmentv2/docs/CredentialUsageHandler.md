# CredentialUsageHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of uses. | [optional] 
**Type** | Pointer to **string** | Type of usage. | [optional] 

## Methods

### NewCredentialUsageHandler

`func NewCredentialUsageHandler() *CredentialUsageHandler`

NewCredentialUsageHandler instantiates a new CredentialUsageHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialUsageHandlerWithDefaults

`func NewCredentialUsageHandlerWithDefaults() *CredentialUsageHandler`

NewCredentialUsageHandlerWithDefaults instantiates a new CredentialUsageHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CredentialUsageHandler) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CredentialUsageHandler) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CredentialUsageHandler) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CredentialUsageHandler) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *CredentialUsageHandler) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CredentialUsageHandler) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CredentialUsageHandler) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CredentialUsageHandler) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


