# UnitList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | The total number of units in the result. | 
**Units** | [**[]Unit**](Unit.md) | A list of units. | 

## Methods

### NewUnitList

`func NewUnitList(totalCount int64, units []Unit, ) *UnitList`

NewUnitList instantiates a new UnitList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnitListWithDefaults

`func NewUnitListWithDefaults() *UnitList`

NewUnitListWithDefaults instantiates a new UnitList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *UnitList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UnitList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UnitList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetUnits

`func (o *UnitList) GetUnits() []Unit`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *UnitList) GetUnitsOk() (*[]Unit, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *UnitList) SetUnits(v []Unit)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


