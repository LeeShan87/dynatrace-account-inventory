# ErrorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **float32** | The code of the error. | 
**Message** | **string** | A short description of the error. | 
**ErrorsMap** | **map[string]string** |  | 

## Methods

### NewErrorDto

`func NewErrorDto(code float32, message string, errorsMap map[string]string, ) *ErrorDto`

NewErrorDto instantiates a new ErrorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDtoWithDefaults

`func NewErrorDtoWithDefaults() *ErrorDto`

NewErrorDtoWithDefaults instantiates a new ErrorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorDto) GetCode() float32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDto) GetCodeOk() (*float32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDto) SetCode(v float32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorsMap

`func (o *ErrorDto) GetErrorsMap() map[string]string`

GetErrorsMap returns the ErrorsMap field if non-nil, zero value otherwise.

### GetErrorsMapOk

`func (o *ErrorDto) GetErrorsMapOk() (*map[string]string, bool)`

GetErrorsMapOk returns a tuple with the ErrorsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsMap

`func (o *ErrorDto) SetErrorsMap(v map[string]string)`

SetErrorsMap sets ErrorsMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


