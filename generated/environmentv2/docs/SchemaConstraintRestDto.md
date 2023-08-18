# SchemaConstraintRestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomMessage** | Pointer to **string** | A custom message for invalid values. | [optional] 
**CustomValidatorId** | Pointer to **string** | The ID of a custom validator. | [optional] 
**Type** | **string** | The type of the schema constraint. | 
**UniqueProperties** | Pointer to **[]string** | The list of properties for which the combination of values needs to be unique | [optional] 

## Methods

### NewSchemaConstraintRestDto

`func NewSchemaConstraintRestDto(type_ string, ) *SchemaConstraintRestDto`

NewSchemaConstraintRestDto instantiates a new SchemaConstraintRestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaConstraintRestDtoWithDefaults

`func NewSchemaConstraintRestDtoWithDefaults() *SchemaConstraintRestDto`

NewSchemaConstraintRestDtoWithDefaults instantiates a new SchemaConstraintRestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomMessage

`func (o *SchemaConstraintRestDto) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *SchemaConstraintRestDto) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *SchemaConstraintRestDto) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *SchemaConstraintRestDto) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetCustomValidatorId

`func (o *SchemaConstraintRestDto) GetCustomValidatorId() string`

GetCustomValidatorId returns the CustomValidatorId field if non-nil, zero value otherwise.

### GetCustomValidatorIdOk

`func (o *SchemaConstraintRestDto) GetCustomValidatorIdOk() (*string, bool)`

GetCustomValidatorIdOk returns a tuple with the CustomValidatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValidatorId

`func (o *SchemaConstraintRestDto) SetCustomValidatorId(v string)`

SetCustomValidatorId sets CustomValidatorId field to given value.

### HasCustomValidatorId

`func (o *SchemaConstraintRestDto) HasCustomValidatorId() bool`

HasCustomValidatorId returns a boolean if a field has been set.

### GetType

`func (o *SchemaConstraintRestDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaConstraintRestDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaConstraintRestDto) SetType(v string)`

SetType sets Type field to given value.


### GetUniqueProperties

`func (o *SchemaConstraintRestDto) GetUniqueProperties() []string`

GetUniqueProperties returns the UniqueProperties field if non-nil, zero value otherwise.

### GetUniquePropertiesOk

`func (o *SchemaConstraintRestDto) GetUniquePropertiesOk() (*[]string, bool)`

GetUniquePropertiesOk returns a tuple with the UniqueProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueProperties

`func (o *SchemaConstraintRestDto) SetUniqueProperties(v []string)`

SetUniqueProperties sets UniqueProperties field to given value.

### HasUniqueProperties

`func (o *SchemaConstraintRestDto) HasUniqueProperties() bool`

HasUniqueProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


