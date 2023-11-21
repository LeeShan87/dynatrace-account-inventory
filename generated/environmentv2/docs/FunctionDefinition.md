# FunctionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | Pointer to **string** | The fully qualified class name of the class that includes the function. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | A human readable string representation of the function definition. | [optional] [readonly] 
**FunctionName** | Pointer to **string** | The function/method name. | [optional] [readonly] 
**ParameterTypes** | Pointer to [**TruncatableListString**](TruncatableListString.md) |  | [optional] 
**ReturnType** | Pointer to **string** | The return type of the function. | [optional] [readonly] 

## Methods

### NewFunctionDefinition

`func NewFunctionDefinition() *FunctionDefinition`

NewFunctionDefinition instantiates a new FunctionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDefinitionWithDefaults

`func NewFunctionDefinitionWithDefaults() *FunctionDefinition`

NewFunctionDefinitionWithDefaults instantiates a new FunctionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *FunctionDefinition) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *FunctionDefinition) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *FunctionDefinition) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *FunctionDefinition) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetDisplayName

`func (o *FunctionDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FunctionDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FunctionDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FunctionDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFunctionName

`func (o *FunctionDefinition) GetFunctionName() string`

GetFunctionName returns the FunctionName field if non-nil, zero value otherwise.

### GetFunctionNameOk

`func (o *FunctionDefinition) GetFunctionNameOk() (*string, bool)`

GetFunctionNameOk returns a tuple with the FunctionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionName

`func (o *FunctionDefinition) SetFunctionName(v string)`

SetFunctionName sets FunctionName field to given value.

### HasFunctionName

`func (o *FunctionDefinition) HasFunctionName() bool`

HasFunctionName returns a boolean if a field has been set.

### GetParameterTypes

`func (o *FunctionDefinition) GetParameterTypes() TruncatableListString`

GetParameterTypes returns the ParameterTypes field if non-nil, zero value otherwise.

### GetParameterTypesOk

`func (o *FunctionDefinition) GetParameterTypesOk() (*TruncatableListString, bool)`

GetParameterTypesOk returns a tuple with the ParameterTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterTypes

`func (o *FunctionDefinition) SetParameterTypes(v TruncatableListString)`

SetParameterTypes sets ParameterTypes field to given value.

### HasParameterTypes

`func (o *FunctionDefinition) HasParameterTypes() bool`

HasParameterTypes returns a boolean if a field has been set.

### GetReturnType

`func (o *FunctionDefinition) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *FunctionDefinition) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *FunctionDefinition) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *FunctionDefinition) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


