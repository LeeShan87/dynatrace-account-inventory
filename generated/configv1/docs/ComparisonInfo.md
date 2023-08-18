# ComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Negate** | **bool** | Reverse the comparison **operator**. For example, it turns **equals** into **does not equal**. | 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STRING&#x60; -&gt; StringComparisonInfo  * &#x60;NUMBER&#x60; -&gt; NumberComparisonInfo  * &#x60;BOOLEAN&#x60; -&gt; BooleanComparisonInfo  * &#x60;HTTP_METHOD&#x60; -&gt; HttpMethodComparisonInfo  * &#x60;STRING_REQUEST_ATTRIBUTE&#x60; -&gt; StringRequestAttributeComparisonInfo  * &#x60;NUMBER_REQUEST_ATTRIBUTE&#x60; -&gt; NumberRequestAttributeComparisonInfo  * &#x60;ZOS_CALL_TYPE&#x60; -&gt; ZosComparisonInfo  * &#x60;IIB_INPUT_NODE_TYPE&#x60; -&gt; IIBInputNodeTypeComparisonInfo  * &#x60;ESB_INPUT_NODE_TYPE&#x60; -&gt; ESBInputNodeTypeComparisonInfo  * &#x60;FAILED_STATE&#x60; -&gt; FailedStateComparisonInfo  * &#x60;FLAW_STATE&#x60; -&gt; FlawStateComparisonInfo  * &#x60;FAILURE_REASON&#x60; -&gt; FailureReasonComparisonInfo  * &#x60;HTTP_STATUS_CLASS&#x60; -&gt; HttpStatusClassComparisonInfo  * &#x60;TAG&#x60; -&gt; TagComparisonInfo  * &#x60;FAST_STRING&#x60; -&gt; FastStringComparisonInfo  * &#x60;SERVICE_TYPE&#x60; -&gt; ServiceTypeComparisonInfo   | 
**Value** | Pointer to **map[string]interface{}** | The value to compare to. | [optional] 
**Values** | Pointer to  | The values to compare to. | [optional] 

## Methods

### NewComparisonInfo

`func NewComparisonInfo(comparison string, negate bool, type_ string, ) *ComparisonInfo`

NewComparisonInfo instantiates a new ComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComparisonInfoWithDefaults

`func NewComparisonInfoWithDefaults() *ComparisonInfo`

NewComparisonInfoWithDefaults instantiates a new ComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *ComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *ComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *ComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetNegate

`func (o *ComparisonInfo) GetNegate() bool`

GetNegate returns the Negate field if non-nil, zero value otherwise.

### GetNegateOk

`func (o *ComparisonInfo) GetNegateOk() (*bool, bool)`

GetNegateOk returns a tuple with the Negate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegate

`func (o *ComparisonInfo) SetNegate(v bool)`

SetNegate sets Negate field to given value.


### GetType

`func (o *ComparisonInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComparisonInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComparisonInfo) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ComparisonInfo) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComparisonInfo) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComparisonInfo) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *ComparisonInfo) GetValues() []map[string]interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ComparisonInfo) GetValuesOk() (*[]map[string]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ComparisonInfo) SetValues(v []map[string]interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *ComparisonInfo) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


