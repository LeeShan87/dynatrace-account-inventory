# PolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The ID of the policy. | 
**Name** | **string** | The display name of the policy. | 
**Description** | **string** | A short description of the policy. | 

## Methods

### NewPolicyDto

`func NewPolicyDto(uuid string, name string, description string, ) *PolicyDto`

NewPolicyDto instantiates a new PolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDtoWithDefaults

`func NewPolicyDtoWithDefaults() *PolicyDto`

NewPolicyDtoWithDefaults instantiates a new PolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PolicyDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PolicyDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PolicyDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *PolicyDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyDto) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


