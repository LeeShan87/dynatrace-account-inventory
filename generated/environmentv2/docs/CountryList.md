# CountryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to [**[]Country**](Country.md) | The list of countries. | [optional] 
**CountryCount** | Pointer to **int32** | The number of countries. | [optional] 

## Methods

### NewCountryList

`func NewCountryList() *CountryList`

NewCountryList instantiates a new CountryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryListWithDefaults

`func NewCountryListWithDefaults() *CountryList`

NewCountryListWithDefaults instantiates a new CountryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *CountryList) GetCountries() []Country`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CountryList) GetCountriesOk() (*[]Country, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CountryList) SetCountries(v []Country)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CountryList) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCountryCount

`func (o *CountryList) GetCountryCount() int32`

GetCountryCount returns the CountryCount field if non-nil, zero value otherwise.

### GetCountryCountOk

`func (o *CountryList) GetCountryCountOk() (*int32, bool)`

GetCountryCountOk returns a tuple with the CountryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCount

`func (o *CountryList) SetCountryCount(v int32)`

SetCountryCount sets CountryCount field to given value.

### HasCountryCount

`func (o *CountryList) HasCountryCount() bool`

HasCountryCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


