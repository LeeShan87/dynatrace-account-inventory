# Attacker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**AttackerLocation**](AttackerLocation.md) |  | [optional] 
**SourceIp** | Pointer to **string** | The source IP of the attacker. | [optional] [readonly] 

## Methods

### NewAttacker

`func NewAttacker() *Attacker`

NewAttacker instantiates a new Attacker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackerWithDefaults

`func NewAttackerWithDefaults() *Attacker`

NewAttackerWithDefaults instantiates a new Attacker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *Attacker) GetLocation() AttackerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Attacker) GetLocationOk() (*AttackerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Attacker) SetLocation(v AttackerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Attacker) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetSourceIp

`func (o *Attacker) GetSourceIp() string`

GetSourceIp returns the SourceIp field if non-nil, zero value otherwise.

### GetSourceIpOk

`func (o *Attacker) GetSourceIpOk() (*string, bool)`

GetSourceIpOk returns a tuple with the SourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIp

`func (o *Attacker) SetSourceIp(v string)`

SetSourceIp sets SourceIp field to given value.

### HasSourceIp

`func (o *Attacker) HasSourceIp() bool`

HasSourceIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


