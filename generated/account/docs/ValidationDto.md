# ValidationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | **[]string** | A list of validation warnings. | 

## Methods

### NewValidationDto

`func NewValidationDto(warnings []string, ) *ValidationDto`

NewValidationDto instantiates a new ValidationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationDtoWithDefaults

`func NewValidationDtoWithDefaults() *ValidationDto`

NewValidationDtoWithDefaults instantiates a new ValidationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *ValidationDto) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidationDto) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidationDto) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


