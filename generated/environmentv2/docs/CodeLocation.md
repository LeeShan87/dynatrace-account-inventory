# CodeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | Pointer to **string** | The fully qualified class name of the code location. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | A human readable string representation of the code location. | [optional] [readonly] 
**FunctionName** | Pointer to **string** | The function/method name of the code location. | [optional] [readonly] 
**LineNumber** | Pointer to **int32** | The line number of the code location. | [optional] [readonly] 
**ParameterTypes** | Pointer to [**TruncatableListString**](TruncatableListString.md) |  | [optional] 
**ReturnType** | Pointer to **string** | The return type of the function. | [optional] [readonly] 

## Methods

### NewCodeLocation

`func NewCodeLocation() *CodeLocation`

NewCodeLocation instantiates a new CodeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeLocationWithDefaults

`func NewCodeLocationWithDefaults() *CodeLocation`

NewCodeLocationWithDefaults instantiates a new CodeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *CodeLocation) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *CodeLocation) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *CodeLocation) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *CodeLocation) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CodeLocation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CodeLocation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CodeLocation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CodeLocation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFunctionName

`func (o *CodeLocation) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *CodeLocation) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *CodeLocation) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *CodeLocation) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetLineNumber

`func (o *CodeLocation) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *CodeLocation) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *CodeLocation) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *CodeLocation) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetParameterTypes

`func (o *CodeLocation) GetParameterTypes() TruncatableListString`

GetParameterTypes returns the ParameterTypes field if non-nil, zero value otherwise.

### GetParameterTypesOk

`func (o *CodeLocation) GetParameterTypesOk() (*TruncatableListString, bool)`

GetParameterTypesOk returns a tuple with the ParameterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterTypes

`func (o *CodeLocation) SetParameterTypes(v TruncatableListString)`

SetParameterTypes sets ParameterTypes field to given value.

### HasParameterTypes

`func (o *CodeLocation) HasParameterTypes() bool`

HasParameterTypes returns a boolean if a field has been set.

### GetReturnType

`func (o *CodeLocation) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *CodeLocation) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *CodeLocation) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *CodeLocation) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


