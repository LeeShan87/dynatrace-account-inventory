# ESBInputNodeTypeComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**Values** | Pointer to  | The values to compare to. | [optional] 

## Methods

### NewESBInputNodeTypeComparisonInfo

`func NewESBInputNodeTypeComparisonInfo(comparison string, ) *ESBInputNodeTypeComparisonInfo`

NewESBInputNodeTypeComparisonInfo instantiates a new ESBInputNodeTypeComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewESBInputNodeTypeComparisonInfoWithDefaults

`func NewESBInputNodeTypeComparisonInfoWithDefaults() *ESBInputNodeTypeComparisonInfo`

NewESBInputNodeTypeComparisonInfoWithDefaults instantiates a new ESBInputNodeTypeComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *ESBInputNodeTypeComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *ESBInputNodeTypeComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *ESBInputNodeTypeComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *ESBInputNodeTypeComparisonInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ESBInputNodeTypeComparisonInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ESBInputNodeTypeComparisonInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ESBInputNodeTypeComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *ESBInputNodeTypeComparisonInfo) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ESBInputNodeTypeComparisonInfo) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ESBInputNodeTypeComparisonInfo) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *ESBInputNodeTypeComparisonInfo) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


