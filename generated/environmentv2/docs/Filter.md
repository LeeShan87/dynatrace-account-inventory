# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operands** | Pointer to [**[]Filter**](Filter.md) | If the type is &#x60;not&#x60;, &#x60;and&#x60; or &#x60;or&#x60;, then holds the contained filters. | [optional] 
**ReferenceInvocation** | Pointer to [**Invocation**](Invocation.md) |  | [optional] 
**ReferenceString** | Pointer to **string** | For filters that match a dimension against a valkue, such as &#x60;eq&#x60; or &#x60;ne&#x60;, holds the value to compare the dimension against. | [optional] 
**ReferenceValue** | Pointer to **float32** | For the operands of &#x60;series&#x60; filters that match against a number, holds the number to compare against. | [optional] 
**Rollup** | Pointer to [**Rollup**](Rollup.md) |  | [optional] 
**TargetDimension** | Pointer to **string** | If the type applies to a dimension, then holds the target dimension. | [optional] 
**TargetDimensions** | Pointer to **[]string** | If the type applies to n dimensions, then holds the target dimensions. Currently only used for the &#x60;remainder&#x60; filter. | [optional] 
**Type** | Pointer to **string** | Type of this filter, determines which other fields are present.Can be any of:  * &#x60;eq&#x60;, * &#x60;ne&#x60;, * &#x60;prefix&#x60;, * &#x60;in&#x60;, * &#x60;remainder&#x60;, * &#x60;suffix&#x60;, * &#x60;contains&#x60;, * &#x60;existsKey&#x60;, * &#x60;series&#x60;, * &#x60;or&#x60;, * &#x60;and&#x60;, * &#x60;not&#x60;, * &#x60;ge&#x60;, * &#x60;gt&#x60;, * &#x60;le&#x60;, * &#x60;lt&#x60;, * &#x60;otherwise&#x60;. | [optional] 

## Methods

### NewFilter

`func NewFilter() *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperands

`func (o *Filter) GetOperands() []Filter`

GetOperands returns the Operands field if non-nil, zero value otherwise.

### GetOperandsOk

`func (o *Filter) GetOperandsOk() (*[]Filter, bool)`

GetOperandsOk returns a tuple with the Operands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperands

`func (o *Filter) SetOperands(v []Filter)`

SetOperands sets Operands field to given value.

### HasOperands

`func (o *Filter) HasOperands() bool`

HasOperands returns a boolean if a field has been set.

### GetReferenceInvocation

`func (o *Filter) GetReferenceInvocation() Invocation`

GetReferenceInvocation returns the ReferenceInvocation field if non-nil, zero value otherwise.

### GetReferenceInvocationOk

`func (o *Filter) GetReferenceInvocationOk() (*Invocation, bool)`

GetReferenceInvocationOk returns a tuple with the ReferenceInvocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceInvocation

`func (o *Filter) SetReferenceInvocation(v Invocation)`

SetReferenceInvocation sets ReferenceInvocation field to given value.

### HasReferenceInvocation

`func (o *Filter) HasReferenceInvocation() bool`

HasReferenceInvocation returns a boolean if a field has been set.

### GetReferenceString

`func (o *Filter) GetReferenceString() string`

GetReferenceString returns the ReferenceString field if non-nil, zero value otherwise.

### GetReferenceStringOk

`func (o *Filter) GetReferenceStringOk() (*string, bool)`

GetReferenceStringOk returns a tuple with the ReferenceString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceString

`func (o *Filter) SetReferenceString(v string)`

SetReferenceString sets ReferenceString field to given value.

### HasReferenceString

`func (o *Filter) HasReferenceString() bool`

HasReferenceString returns a boolean if a field has been set.

### GetReferenceValue

`func (o *Filter) GetReferenceValue() float32`

GetReferenceValue returns the ReferenceValue field if non-nil, zero value otherwise.

### GetReferenceValueOk

`func (o *Filter) GetReferenceValueOk() (*float32, bool)`

GetReferenceValueOk returns a tuple with the ReferenceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceValue

`func (o *Filter) SetReferenceValue(v float32)`

SetReferenceValue sets ReferenceValue field to given value.

### HasReferenceValue

`func (o *Filter) HasReferenceValue() bool`

HasReferenceValue returns a boolean if a field has been set.

### GetRollup

`func (o *Filter) GetRollup() Rollup`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *Filter) GetRollupOk() (*Rollup, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *Filter) SetRollup(v Rollup)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *Filter) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetTargetDimension

`func (o *Filter) GetTargetDimension() string`

GetTargetDimension returns the TargetDimension field if non-nil, zero value otherwise.

### GetTargetDimensionOk

`func (o *Filter) GetTargetDimensionOk() (*string, bool)`

GetTargetDimensionOk returns a tuple with the TargetDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDimension

`func (o *Filter) SetTargetDimension(v string)`

SetTargetDimension sets TargetDimension field to given value.

### HasTargetDimension

`func (o *Filter) HasTargetDimension() bool`

HasTargetDimension returns a boolean if a field has been set.

### GetTargetDimensions

`func (o *Filter) GetTargetDimensions() []string`

GetTargetDimensions returns the TargetDimensions field if non-nil, zero value otherwise.

### GetTargetDimensionsOk

`func (o *Filter) GetTargetDimensionsOk() (*[]string, bool)`

GetTargetDimensionsOk returns a tuple with the TargetDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDimensions

`func (o *Filter) SetTargetDimensions(v []string)`

SetTargetDimensions sets TargetDimensions field to given value.

### HasTargetDimensions

`func (o *Filter) HasTargetDimensions() bool`

HasTargetDimensions returns a boolean if a field has been set.

### GetType

`func (o *Filter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Filter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Filter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Filter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


