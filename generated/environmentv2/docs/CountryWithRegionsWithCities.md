# CountryWithRegionsWithCities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The ISO code of the country. | [optional] 
**CountryName** | Pointer to **string** | The name of the country. | [optional] 
**RegionCount** | Pointer to **int32** | The number of regions in the country. | [optional] 
**Regions** | Pointer to [**[]RegionWithCities**](RegionWithCities.md) | The list of regions in the country. | [optional] 

## Methods

### NewCountryWithRegionsWithCities

`func NewCountryWithRegionsWithCities() *CountryWithRegionsWithCities`

NewCountryWithRegionsWithCities instantiates a new CountryWithRegionsWithCities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithRegionsWithCitiesWithDefaults

`func NewCountryWithRegionsWithCitiesWithDefaults() *CountryWithRegionsWithCities`

NewCountryWithRegionsWithCitiesWithDefaults instantiates a new CountryWithRegionsWithCities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CountryWithRegionsWithCities) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CountryWithRegionsWithCities) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CountryWithRegionsWithCities) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CountryWithRegionsWithCities) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *CountryWithRegionsWithCities) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *CountryWithRegionsWithCities) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *CountryWithRegionsWithCities) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *CountryWithRegionsWithCities) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetRegionCount

`func (o *CountryWithRegionsWithCities) GetRegionCount() int32`

GetRegionCount returns the RegionCount field if non-nil, zero value otherwise.

### GetRegionCountOk

`func (o *CountryWithRegionsWithCities) GetRegionCountOk() (*int32, bool)`

GetRegionCountOk returns a tuple with the RegionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCount

`func (o *CountryWithRegionsWithCities) SetRegionCount(v int32)`

SetRegionCount sets RegionCount field to given value.

### HasRegionCount

`func (o *CountryWithRegionsWithCities) HasRegionCount() bool`

HasRegionCount returns a boolean if a field has been set.

### GetRegions

`func (o *CountryWithRegionsWithCities) GetRegions() []RegionWithCities`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CountryWithRegionsWithCities) GetRegionsOk() (*[]RegionWithCities, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CountryWithRegionsWithCities) SetRegions(v []RegionWithCities)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CountryWithRegionsWithCities) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


