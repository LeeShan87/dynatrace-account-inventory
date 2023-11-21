# ActiveGateTokenInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | Pointer to **string** | The environment ID to which the token belongs.   Only available if more than one environment is supported. | [optional] [readonly] 
**Id** | Pointer to **string** | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token. | [optional] [readonly] 
**State** | Pointer to **string** | State of the ActiveGate token. | [optional] [readonly] 

## Methods

### NewActiveGateTokenInfoDto

`func NewActiveGateTokenInfoDto() *ActiveGateTokenInfoDto`

NewActiveGateTokenInfoDto instantiates a new ActiveGateTokenInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateTokenInfoDtoWithDefaults

`func NewActiveGateTokenInfoDtoWithDefaults() *ActiveGateTokenInfoDto`

NewActiveGateTokenInfoDtoWithDefaults instantiates a new ActiveGateTokenInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *ActiveGateTokenInfoDto) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ActiveGateTokenInfoDto) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ActiveGateTokenInfoDto) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ActiveGateTokenInfoDto) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetId

`func (o *ActiveGateTokenInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActiveGateTokenInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActiveGateTokenInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActiveGateTokenInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *ActiveGateTokenInfoDto) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActiveGateTokenInfoDto) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActiveGateTokenInfoDto) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ActiveGateTokenInfoDto) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


