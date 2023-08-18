# ActiveGateTokenCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **string** | The token expiration date in ISO 8601 format (&#x60;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;&#x60;). | [optional] 
**Id** | **string** | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token. | 
**Token** | **string** | The secret of the token. | 

## Methods

### NewActiveGateTokenCreated

`func NewActiveGateTokenCreated(id string, token string, ) *ActiveGateTokenCreated`

NewActiveGateTokenCreated instantiates a new ActiveGateTokenCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateTokenCreatedWithDefaults

`func NewActiveGateTokenCreatedWithDefaults() *ActiveGateTokenCreated`

NewActiveGateTokenCreatedWithDefaults instantiates a new ActiveGateTokenCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *ActiveGateTokenCreated) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ActiveGateTokenCreated) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ActiveGateTokenCreated) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ActiveGateTokenCreated) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetId

`func (o *ActiveGateTokenCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveGateTokenCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveGateTokenCreated) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *ActiveGateTokenCreated) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ActiveGateTokenCreated) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ActiveGateTokenCreated) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


