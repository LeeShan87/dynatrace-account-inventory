# RegionWithCities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cities** | Pointer to [**[]City**](City.md) | The list of cities in the region. | [optional] 
**CityCount** | Pointer to **int32** | The number of cities in a region of a country. | [optional] 
**Code** | Pointer to **string** | The code of the region. | [optional] 
**Name** | Pointer to **string** | The name of the region. | [optional] 

## Methods

### NewRegionWithCities

`func NewRegionWithCities() *RegionWithCities`

NewRegionWithCities instantiates a new RegionWithCities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithCitiesWithDefaults

`func NewRegionWithCitiesWithDefaults() *RegionWithCities`

NewRegionWithCitiesWithDefaults instantiates a new RegionWithCities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCities

`func (o *RegionWithCities) GetCities() []City`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *RegionWithCities) GetCitiesOk() (*[]City, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *RegionWithCities) SetCities(v []City)`

SetCities sets Cities field to given value.

### HasCities

`func (o *RegionWithCities) HasCities() bool`

HasCities returns a boolean if a field has been set.

### GetCityCount

`func (o *RegionWithCities) GetCityCount() int32`

GetCityCount returns the CityCount field if non-nil, zero value otherwise.

### GetCityCountOk

`func (o *RegionWithCities) GetCityCountOk() (*int32, bool)`

GetCityCountOk returns a tuple with the CityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityCount

`func (o *RegionWithCities) SetCityCount(v int32)`

SetCityCount sets CityCount field to given value.

### HasCityCount

`func (o *RegionWithCities) HasCityCount() bool`

HasCityCount returns a boolean if a field has been set.

### GetCode

`func (o *RegionWithCities) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RegionWithCities) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RegionWithCities) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RegionWithCities) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *RegionWithCities) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionWithCities) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionWithCities) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegionWithCities) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


