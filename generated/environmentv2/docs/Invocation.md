# Invocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | Arguments to pass to the function, e.g. entity selector source code. | [optional] 
**Function** | Pointer to **string** | Function that is invoked, e.g. &#x60;entitySelector&#x60;. | [optional] 

## Methods

### NewInvocation

`func NewInvocation() *Invocation`

NewInvocation instantiates a new Invocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvocationWithDefaults

`func NewInvocationWithDefaults() *Invocation`

NewInvocationWithDefaults instantiates a new Invocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *Invocation) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Invocation) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Invocation) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Invocation) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetFunction

`func (o *Invocation) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Invocation) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Invocation) SetFunction(v string)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *Invocation) HasFunction() bool`

HasFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


