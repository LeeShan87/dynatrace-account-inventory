# DexpFilterRelationship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the relationship. e.g runsOn, isStepOf, etc | [optional] 
**TargetEntity** | Pointer to **string** | The target entity of the relationship. e.g HOST, VCENTER, SERVICE etc | [optional] 
**Type** | Pointer to **string** | The type of the relationship | [optional] 

## Methods

### NewDexpFilterRelationship

`func NewDexpFilterRelationship() *DexpFilterRelationship`

NewDexpFilterRelationship instantiates a new DexpFilterRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDexpFilterRelationshipWithDefaults

`func NewDexpFilterRelationshipWithDefaults() *DexpFilterRelationship`

NewDexpFilterRelationshipWithDefaults instantiates a new DexpFilterRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DexpFilterRelationship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DexpFilterRelationship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DexpFilterRelationship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DexpFilterRelationship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetEntity

`func (o *DexpFilterRelationship) GetTargetEntity() string`

GetTargetEntity returns the TargetEntity field if non-nil, zero value otherwise.

### GetTargetEntityOk

`func (o *DexpFilterRelationship) GetTargetEntityOk() (*string, bool)`

GetTargetEntityOk returns a tuple with the TargetEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntity

`func (o *DexpFilterRelationship) SetTargetEntity(v string)`

SetTargetEntity sets TargetEntity field to given value.

### HasTargetEntity

`func (o *DexpFilterRelationship) HasTargetEntity() bool`

HasTargetEntity returns a boolean if a field has been set.

### GetType

`func (o *DexpFilterRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DexpFilterRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DexpFilterRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DexpFilterRelationship) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


