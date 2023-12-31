# SyntheticLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The city of the location. | [optional] 
**CountryCode** | Pointer to **string** | The country code of the location.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request. | [optional] 
**EntityId** | **string** | The Dynatrace entity ID of the location. | 
**Latitude** | **float64** | The latitude of the location in &#x60;DDD.dddd&#x60; format. | 
**Longitude** | **float64** | The longitude of the location in &#x60;DDD.dddd&#x60; format. | 
**Name** | **string** | The name of the location. | 
**RegionCode** | Pointer to **string** | The region code of the location.    To fetch the list of available region codes, use the [GET regions of the country](https://dt-url.net/az230x0) request. | [optional] 
**Status** | Pointer to **string** | The status of the location:   * &#x60;ENABLED&#x60;: The location is displayed as active in the UI. You can assign monitors to the location.  * &#x60;DISABLED&#x60;: The location is displayed as inactive in the UI. You can&#39;t assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * &#x60;HIDDEN&#x60;: The location is not displayed in the UI. You can&#39;t assign monitors to the location. You can only set location as &#x60;HIDDEN&#x60; when no monitor is assigned to it. | [optional] 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;PUBLIC&#x60; -&gt; PublicSyntheticLocation  * &#x60;PRIVATE&#x60; -&gt; PrivateSyntheticLocation  * &#x60;CLUSTER&#x60; -&gt; PrivateSyntheticLocation   | 

## Methods

### NewSyntheticLocation

`func NewSyntheticLocation(entityId string, latitude float64, longitude float64, name string, type_ string, ) *SyntheticLocation`

NewSyntheticLocation instantiates a new SyntheticLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticLocationWithDefaults

`func NewSyntheticLocationWithDefaults() *SyntheticLocation`

NewSyntheticLocationWithDefaults instantiates a new SyntheticLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *SyntheticLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SyntheticLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SyntheticLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SyntheticLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *SyntheticLocation) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SyntheticLocation) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SyntheticLocation) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SyntheticLocation) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetEntityId

`func (o *SyntheticLocation) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SyntheticLocation) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SyntheticLocation) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetLatitude

`func (o *SyntheticLocation) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SyntheticLocation) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SyntheticLocation) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SyntheticLocation) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SyntheticLocation) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SyntheticLocation) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetName

`func (o *SyntheticLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticLocation) SetName(v string)`

SetName sets Name field to given value.


### GetRegionCode

`func (o *SyntheticLocation) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *SyntheticLocation) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *SyntheticLocation) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *SyntheticLocation) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticLocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticLocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticLocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticLocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SyntheticLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticLocation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


