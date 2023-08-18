# CountryWithRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** | The ISO code of the country. | [optional] 
**CountryName** | Pointer to **string** | The name of the country. | [optional] 
**RegionCount** | Pointer to **int32** | The number of regions in the country. | [optional] 
**Regions** | Pointer to [**[]Region**](Region.md) | The list of regions in the country. | [optional] 

## Methods

### NewCountryWithRegions

`func NewCountryWithRegions() *CountryWithRegions`

NewCountryWithRegions instantiates a new CountryWithRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryWithRegionsWithDefaults

`func NewCountryWithRegionsWithDefaults() *CountryWithRegions`

NewCountryWithRegionsWithDefaults instantiates a new CountryWithRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CountryWithRegions) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CountryWithRegions) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CountryWithRegions) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CountryWithRegions) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountryName

`func (o *CountryWithRegions) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *CountryWithRegions) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *CountryWithRegions) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.

### HasCountryName

`func (o *CountryWithRegions) HasCountryName() bool`

HasCountryName returns a boolean if a field has been set.

### GetRegionCount

`func (o *CountryWithRegions) GetRegionCount() int32`

GetRegionCount returns the RegionCount field if non-nil, zero value otherwise.

### GetRegionCountOk

`func (o *CountryWithRegions) GetRegionCountOk() (*int32, bool)`

GetRegionCountOk returns a tuple with the RegionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCount

`func (o *CountryWithRegions) SetRegionCount(v int32)`

SetRegionCount sets RegionCount field to given value.

### HasRegionCount

`func (o *CountryWithRegions) HasRegionCount() bool`

HasRegionCount returns a boolean if a field has been set.

### GetRegions

`func (o *CountryWithRegions) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CountryWithRegions) GetRegionsOk() (*[]Region, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CountryWithRegions) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CountryWithRegions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


