# Unit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the unit. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the unit. | [optional] 
**DisplayNamePlural** | Pointer to **string** | The plural display name of the unit. | [optional] 
**Symbol** | Pointer to **string** | The symbol of the unit. | [optional] 
**UnitId** | **string** | The ID of the unit. | 

## Methods

### NewUnit

`func NewUnit(unitId string, ) *Unit`

NewUnit instantiates a new Unit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitWithDefaults

`func NewUnitWithDefaults() *Unit`

NewUnitWithDefaults instantiates a new Unit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Unit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Unit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Unit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Unit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Unit) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Unit) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Unit) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Unit) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNamePlural

`func (o *Unit) GetDisplayNamePlural() string`

GetDisplayNamePlural returns the DisplayNamePlural field if non-nil, zero value otherwise.

### GetDisplayNamePluralOk

`func (o *Unit) GetDisplayNamePluralOk() (*string, bool)`

GetDisplayNamePluralOk returns a tuple with the DisplayNamePlural field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNamePlural

`func (o *Unit) SetDisplayNamePlural(v string)`

SetDisplayNamePlural sets DisplayNamePlural field to given value.

### HasDisplayNamePlural

`func (o *Unit) HasDisplayNamePlural() bool`

HasDisplayNamePlural returns a boolean if a field has been set.

### GetSymbol

`func (o *Unit) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Unit) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Unit) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Unit) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetUnitId

`func (o *Unit) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *Unit) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *Unit) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


