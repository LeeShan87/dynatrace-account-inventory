# AttackerLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City of the attacker. | [optional] [readonly] 
**Country** | Pointer to **string** | The country of the attacker. | [optional] [readonly] 
**CountryCode** | Pointer to **string** | The country code of the country of the attacker, according to the ISO 3166-1 Alpha-2 standard. | [optional] [readonly] 

## Methods

### NewAttackerLocation

`func NewAttackerLocation() *AttackerLocation`

NewAttackerLocation instantiates a new AttackerLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackerLocationWithDefaults

`func NewAttackerLocationWithDefaults() *AttackerLocation`

NewAttackerLocationWithDefaults instantiates a new AttackerLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *AttackerLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AttackerLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AttackerLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AttackerLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AttackerLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AttackerLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AttackerLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AttackerLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *AttackerLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AttackerLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AttackerLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AttackerLocation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


