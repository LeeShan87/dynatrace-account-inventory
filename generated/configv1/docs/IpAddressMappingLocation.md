# IpAddressMappingLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The city name of the location. | [optional] 
**CountryCode** | **string** | The country code of the location.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request. | 
**Latitude** | Pointer to **float64** | The latitude of the location in &#x60;DDD.dddd&#x60; format. | [optional] 
**Longitude** | Pointer to **float64** | The longitude of the location in &#x60;DDD.dddd&#x60; format. | [optional] 
**RegionCode** | Pointer to **string** | The region code of the location.    To fetch the list of available region codes, use the [GET regions of the country](https://dt-url.net/az230x0) request. | [optional] 

## Methods

### NewIpAddressMappingLocation

`func NewIpAddressMappingLocation(countryCode string, ) *IpAddressMappingLocation`

NewIpAddressMappingLocation instantiates a new IpAddressMappingLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressMappingLocationWithDefaults

`func NewIpAddressMappingLocationWithDefaults() *IpAddressMappingLocation`

NewIpAddressMappingLocationWithDefaults instantiates a new IpAddressMappingLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *IpAddressMappingLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *IpAddressMappingLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *IpAddressMappingLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *IpAddressMappingLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *IpAddressMappingLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *IpAddressMappingLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *IpAddressMappingLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetLatitude

`func (o *IpAddressMappingLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *IpAddressMappingLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *IpAddressMappingLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *IpAddressMappingLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *IpAddressMappingLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *IpAddressMappingLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *IpAddressMappingLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *IpAddressMappingLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetRegionCode

`func (o *IpAddressMappingLocation) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *IpAddressMappingLocation) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *IpAddressMappingLocation) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *IpAddressMappingLocation) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


