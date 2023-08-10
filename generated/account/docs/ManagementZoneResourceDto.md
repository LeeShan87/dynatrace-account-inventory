# ManagementZoneResourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | **string** | The ID of the environment to which the management zone belongs. | 
**Name** | **string** | The name of the management zone. | 
**Id** | **string** | The ID of the management zone. | 

## Methods

### NewManagementZoneResourceDto

`func NewManagementZoneResourceDto(parent string, name string, id string, ) *ManagementZoneResourceDto`

NewManagementZoneResourceDto instantiates a new ManagementZoneResourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementZoneResourceDtoWithDefaults

`func NewManagementZoneResourceDtoWithDefaults() *ManagementZoneResourceDto`

NewManagementZoneResourceDtoWithDefaults instantiates a new ManagementZoneResourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *ManagementZoneResourceDto) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ManagementZoneResourceDto) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ManagementZoneResourceDto) SetParent(v string)`

SetParent sets Parent field to given value.


### GetName

`func (o *ManagementZoneResourceDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementZoneResourceDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementZoneResourceDto) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *ManagementZoneResourceDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementZoneResourceDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementZoneResourceDto) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


