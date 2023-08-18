# AffectedEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The monitored entity ID of the affected entity. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the affected entity. | [optional] [readonly] 

## Methods

### NewAffectedEntity

`func NewAffectedEntity() *AffectedEntity`

NewAffectedEntity instantiates a new AffectedEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffectedEntityWithDefaults

`func NewAffectedEntityWithDefaults() *AffectedEntity`

NewAffectedEntityWithDefaults instantiates a new AffectedEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AffectedEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffectedEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffectedEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffectedEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AffectedEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AffectedEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AffectedEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AffectedEntity) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


