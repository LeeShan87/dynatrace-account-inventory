# HttpStatusClassComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**Values** | Pointer to  | The values to compare to. | [optional] 

## Methods

### NewHttpStatusClassComparisonInfo

`func NewHttpStatusClassComparisonInfo(comparison string, ) *HttpStatusClassComparisonInfo`

NewHttpStatusClassComparisonInfo instantiates a new HttpStatusClassComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpStatusClassComparisonInfoWithDefaults

`func NewHttpStatusClassComparisonInfoWithDefaults() *HttpStatusClassComparisonInfo`

NewHttpStatusClassComparisonInfoWithDefaults instantiates a new HttpStatusClassComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparison

`func (o *HttpStatusClassComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *HttpStatusClassComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *HttpStatusClassComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *HttpStatusClassComparisonInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HttpStatusClassComparisonInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HttpStatusClassComparisonInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HttpStatusClassComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *HttpStatusClassComparisonInfo) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *HttpStatusClassComparisonInfo) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *HttpStatusClassComparisonInfo) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *HttpStatusClassComparisonInfo) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


