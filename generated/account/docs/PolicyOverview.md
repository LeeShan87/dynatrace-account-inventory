# PolicyOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The ID of the policy. | 
**Name** | **string** | The display name of the policy. | 
**Description** | **string** | A short description of the policy. | 
**LevelId** | **string** | The ID of the level to which the policy applies. | 
**LevelType** | **string** | The type of the level to which the policy applies. | 

## Methods

### NewPolicyOverview

`func NewPolicyOverview(uuid string, name string, description string, levelId string, levelType string, ) *PolicyOverview`

NewPolicyOverview instantiates a new PolicyOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyOverviewWithDefaults

`func NewPolicyOverviewWithDefaults() *PolicyOverview`

NewPolicyOverviewWithDefaults instantiates a new PolicyOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PolicyOverview) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PolicyOverview) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PolicyOverview) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *PolicyOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyOverview) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PolicyOverview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyOverview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyOverview) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLevelId

`func (o *PolicyOverview) GetLevelId() string`

GetLevelId returns the LevelId field if non-nil, zero value otherwise.

### GetLevelIdOk

`func (o *PolicyOverview) GetLevelIdOk() (*string, bool)`

GetLevelIdOk returns a tuple with the LevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelId

`func (o *PolicyOverview) SetLevelId(v string)`

SetLevelId sets LevelId field to given value.


### GetLevelType

`func (o *PolicyOverview) GetLevelType() string`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *PolicyOverview) GetLevelTypeOk() (*string, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *PolicyOverview) SetLevelType(v string)`

SetLevelType sets LevelType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


