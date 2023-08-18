# CountryRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The ISO code of the country. | [optional] 
**Name** | Pointer to **string** | The name of the country. | [optional] 
**RegionCount** | Pointer to **int32** | The number of regions in the country. | [optional] 
**Regions** | Pointer to [**[]Region**](Region.md) | The list of regions in the country. | [optional] 

## Methods

### NewCountryRegions

`func NewCountryRegions() *CountryRegions`

NewCountryRegions instantiates a new CountryRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryRegionsWithDefaults

`func NewCountryRegionsWithDefaults() *CountryRegions`

NewCountryRegionsWithDefaults instantiates a new CountryRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CountryRegions) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CountryRegions) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CountryRegions) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CountryRegions) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *CountryRegions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CountryRegions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CountryRegions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CountryRegions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionCount

`func (o *CountryRegions) GetRegionCount() int32`

GetRegionCount returns the RegionCount field if non-nil, zero value otherwise.

### GetRegionCountOk

`func (o *CountryRegions) GetRegionCountOk() (*int32, bool)`

GetRegionCountOk returns a tuple with the RegionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCount

`func (o *CountryRegions) SetRegionCount(v int32)`

SetRegionCount sets RegionCount field to given value.

### HasRegionCount

`func (o *CountryRegions) HasRegionCount() bool`

HasRegionCount returns a boolean if a field has been set.

### GetRegions

`func (o *CountryRegions) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CountryRegions) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CountryRegions) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CountryRegions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


