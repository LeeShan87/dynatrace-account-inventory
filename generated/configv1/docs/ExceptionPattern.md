# ExceptionPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassPattern** | Pointer to **string** |  | [optional] 
**MessagePattern** | Pointer to **string** |  | [optional] 

## Methods

### NewExceptionPattern

`func NewExceptionPattern() *ExceptionPattern`

NewExceptionPattern instantiates a new ExceptionPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionPatternWithDefaults

`func NewExceptionPatternWithDefaults() *ExceptionPattern`

NewExceptionPatternWithDefaults instantiates a new ExceptionPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassPattern

`func (o *ExceptionPattern) GetClassPattern() string`

GetClassPattern returns the ClassPattern field if non-nil, zero value otherwise.

### GetClassPatternOk

`func (o *ExceptionPattern) GetClassPatternOk() (*string, bool)`

GetClassPatternOk returns a tuple with the ClassPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassPattern

`func (o *ExceptionPattern) SetClassPattern(v string)`

SetClassPattern sets ClassPattern field to given value.

### HasClassPattern

`func (o *ExceptionPattern) HasClassPattern() bool`

HasClassPattern returns a boolean if a field has been set.

### GetMessagePattern

`func (o *ExceptionPattern) GetMessagePattern() string`

GetMessagePattern returns the MessagePattern field if non-nil, zero value otherwise.

### GetMessagePatternOk

`func (o *ExceptionPattern) GetMessagePatternOk() (*string, bool)`

GetMessagePatternOk returns a tuple with the MessagePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePattern

`func (o *ExceptionPattern) SetMessagePattern(v string)`

SetMessagePattern sets MessagePattern field to given value.

### HasMessagePattern

`func (o *ExceptionPattern) HasMessagePattern() bool`

HasMessagePattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


