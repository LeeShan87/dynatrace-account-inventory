# ActiveGateTokenCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveGateType** | **string** | The type of the ActiveGate for which the token is valid. | 
**ExpirationDate** | Pointer to **string** | The expiration date of the token.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the token never expires. | [optional] 
**Name** | **string** | The name of the token. | 
**SeedToken** | Pointer to **bool** | The token is a seed token (&#x60;true&#x60;) or an individual token (&#x60;false&#x60;).    We recommend the individual token option (false). | [optional] 

## Methods

### NewActiveGateTokenCreate

`func NewActiveGateTokenCreate(activeGateType string, name string, ) *ActiveGateTokenCreate`

NewActiveGateTokenCreate instantiates a new ActiveGateTokenCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateTokenCreateWithDefaults

`func NewActiveGateTokenCreateWithDefaults() *ActiveGateTokenCreate`

NewActiveGateTokenCreateWithDefaults instantiates a new ActiveGateTokenCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveGateType

`func (o *ActiveGateTokenCreate) GetActiveGateType() string`

GetActiveGateType returns the ActiveGateType field if non-nil, zero value otherwise.

### GetActiveGateTypeOk

`func (o *ActiveGateTokenCreate) GetActiveGateTypeOk() (*string, bool)`

GetActiveGateTypeOk returns a tuple with the ActiveGateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGateType

`func (o *ActiveGateTokenCreate) SetActiveGateType(v string)`

SetActiveGateType sets ActiveGateType field to given value.


### GetExpirationDate

`func (o *ActiveGateTokenCreate) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ActiveGateTokenCreate) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ActiveGateTokenCreate) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ActiveGateTokenCreate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetName

`func (o *ActiveGateTokenCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActiveGateTokenCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActiveGateTokenCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSeedToken

`func (o *ActiveGateTokenCreate) GetSeedToken() bool`

GetSeedToken returns the SeedToken field if non-nil, zero value otherwise.

### GetSeedTokenOk

`func (o *ActiveGateTokenCreate) GetSeedTokenOk() (*bool, bool)`

GetSeedTokenOk returns a tuple with the SeedToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedToken

`func (o *ActiveGateTokenCreate) SetSeedToken(v bool)`

SetSeedToken sets SeedToken field to given value.

### HasSeedToken

`func (o *ActiveGateTokenCreate) HasSeedToken() bool`

HasSeedToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


