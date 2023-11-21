# ActiveGateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveGateType** | **string** | The type of the ActiveGate for which the token is valid. | 
**CreationDate** | **string** | The token creation date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;). | 
**ExpirationDate** | Pointer to **string** | The token expiration date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;).    If not set, the token never expires. | [optional] 
**Id** | **string** | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token. | 
**LastUsedDate** | Pointer to **string** | The token last used date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;). | [optional] 
**Name** | **string** | The name of the token. | 
**Owner** | **string** | The owner of the token. | 
**SeedToken** | Pointer to **bool** | The token is a seed token (&#x60;true&#x60;) or an individual token (&#x60;false&#x60;). | [optional] 

## Methods

### NewActiveGateToken

`func NewActiveGateToken(activeGateType string, creationDate string, id string, name string, owner string, ) *ActiveGateToken`

NewActiveGateToken instantiates a new ActiveGateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateTokenWithDefaults

`func NewActiveGateTokenWithDefaults() *ActiveGateToken`

NewActiveGateTokenWithDefaults instantiates a new ActiveGateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveGateType

`func (o *ActiveGateToken) GetActiveGateType() string`

GetActiveGateType returns the ActiveGateType field if non-nil, zero value otherwise.

### GetActiveGateTypeOk

`func (o *ActiveGateToken) GetActiveGateTypeOk() (*string, bool)`

GetActiveGateTypeOk returns a tuple with the ActiveGateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGateType

`func (o *ActiveGateToken) SetActiveGateType(v string)`

SetActiveGateType sets ActiveGateType field to given value.


### GetCreationDate

`func (o *ActiveGateToken) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ActiveGateToken) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ActiveGateToken) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.


### GetExpirationDate

`func (o *ActiveGateToken) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ActiveGateToken) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ActiveGateToken) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ActiveGateToken) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *ActiveGateToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveGateToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveGateToken) SetId(v string)`

SetId sets Id field to given value.


### GetLastUsedDate

`func (o *ActiveGateToken) GetLastUsedDate() string`

GetLastUsedDate returns the LastUsedDate field if non-nil, zero value otherwise.

### GetLastUsedDateOk

`func (o *ActiveGateToken) GetLastUsedDateOk() (*string, bool)`

GetLastUsedDateOk returns a tuple with the LastUsedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedDate

`func (o *ActiveGateToken) SetLastUsedDate(v string)`

SetLastUsedDate sets LastUsedDate field to given value.

### HasLastUsedDate

`func (o *ActiveGateToken) HasLastUsedDate() bool`

HasLastUsedDate returns a boolean if a field has been set.

### GetName

`func (o *ActiveGateToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActiveGateToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActiveGateToken) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *ActiveGateToken) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ActiveGateToken) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ActiveGateToken) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetSeedToken

`func (o *ActiveGateToken) GetSeedToken() bool`

GetSeedToken returns the SeedToken field if non-nil, zero value otherwise.

### GetSeedTokenOk

`func (o *ActiveGateToken) GetSeedTokenOk() (*bool, bool)`

GetSeedTokenOk returns a tuple with the SeedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedToken

`func (o *ActiveGateToken) SetSeedToken(v bool)`

SetSeedToken sets SeedToken field to given value.

### HasSeedToken

`func (o *ActiveGateToken) HasSeedToken() bool`

HasSeedToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


