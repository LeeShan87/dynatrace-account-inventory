# TimeZoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The UTC-based name of the time zone. | 
**Name** | **string** | The standard name of the time zone. | 

## Methods

### NewTimeZoneDto

`func NewTimeZoneDto(displayName string, name string, ) *TimeZoneDto`

NewTimeZoneDto instantiates a new TimeZoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneDtoWithDefaults

`func NewTimeZoneDtoWithDefaults() *TimeZoneDto`

NewTimeZoneDtoWithDefaults instantiates a new TimeZoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *TimeZoneDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimeZoneDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimeZoneDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetName

`func (o *TimeZoneDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeZoneDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeZoneDto) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


