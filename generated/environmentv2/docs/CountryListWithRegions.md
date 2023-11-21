# CountryListWithRegions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to [**[]CountryRegions**](CountryRegions.md) | The list of countries. | [optional] 
**CountryCount** | Pointer to **int32** | The number of countries. | [optional] 

## Methods

### NewCountryListWithRegions

`func NewCountryListWithRegions() *CountryListWithRegions`

NewCountryListWithRegions instantiates a new CountryListWithRegions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryListWithRegionsWithDefaults

`func NewCountryListWithRegionsWithDefaults() *CountryListWithRegions`

NewCountryListWithRegionsWithDefaults instantiates a new CountryListWithRegions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *CountryListWithRegions) GetCountries() []CountryRegions`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CountryListWithRegions) GetCountriesOk() (*[]CountryRegions, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CountryListWithRegions) SetCountries(v []CountryRegions)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CountryListWithRegions) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCountryCount

`func (o *CountryListWithRegions) GetCountryCount() int32`

GetCountryCount returns the CountryCount field if non-nil, zero value otherwise.

### GetCountryCountOk

`func (o *CountryListWithRegions) GetCountryCountOk() (*int32, bool)`

GetCountryCountOk returns a tuple with the CountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCount

`func (o *CountryListWithRegions) SetCountryCount(v int32)`

SetCountryCount sets CountryCount field to given value.

### HasCountryCount

`func (o *CountryListWithRegions) HasCountryCount() bool`

HasCountryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


