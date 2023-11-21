# DeletionConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMessage** | Pointer to **string** | A custom message for invalid values. | [optional] 
**CustomValidatorId** | Pointer to **string** | The ID of a custom validator. | [optional] 

## Methods

### NewDeletionConstraint

`func NewDeletionConstraint() *DeletionConstraint`

NewDeletionConstraint instantiates a new DeletionConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletionConstraintWithDefaults

`func NewDeletionConstraintWithDefaults() *DeletionConstraint`

NewDeletionConstraintWithDefaults instantiates a new DeletionConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMessage

`func (o *DeletionConstraint) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *DeletionConstraint) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *DeletionConstraint) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *DeletionConstraint) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetCustomValidatorId

`func (o *DeletionConstraint) GetCustomValidatorId() string`

GetCustomValidatorId returns the CustomValidatorId field if non-nil, zero value otherwise.

### GetCustomValidatorIdOk

`func (o *DeletionConstraint) GetCustomValidatorIdOk() (*string, bool)`

GetCustomValidatorIdOk returns a tuple with the CustomValidatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValidatorId

`func (o *DeletionConstraint) SetCustomValidatorId(v string)`

SetCustomValidatorId sets CustomValidatorId field to given value.

### HasCustomValidatorId

`func (o *DeletionConstraint) HasCustomValidatorId() bool`

HasCustomValidatorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


