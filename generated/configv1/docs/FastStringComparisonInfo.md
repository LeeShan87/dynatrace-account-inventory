# FastStringComparisonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseSensitive** | Pointer to **bool** | The comparison is case-sensitive (&#x60;true&#x60;) or not case-sensitive (&#x60;false&#x60;). | [optional] 
**Comparison** | **string** | Operator of the comparision. You can reverse it by setting **negate** to &#x60;true&#x60;. | 
**Value** | Pointer to **string** | The value to compare to. | [optional] 
**Values** | Pointer to  | The values to compare to. | [optional] 

## Methods

### NewFastStringComparisonInfo

`func NewFastStringComparisonInfo(comparison string, ) *FastStringComparisonInfo`

NewFastStringComparisonInfo instantiates a new FastStringComparisonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFastStringComparisonInfoWithDefaults

`func NewFastStringComparisonInfoWithDefaults() *FastStringComparisonInfo`

NewFastStringComparisonInfoWithDefaults instantiates a new FastStringComparisonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseSensitive

`func (o *FastStringComparisonInfo) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *FastStringComparisonInfo) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *FastStringComparisonInfo) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *FastStringComparisonInfo) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetComparison

`func (o *FastStringComparisonInfo) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *FastStringComparisonInfo) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *FastStringComparisonInfo) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetValue

`func (o *FastStringComparisonInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FastStringComparisonInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FastStringComparisonInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FastStringComparisonInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValues

`func (o *FastStringComparisonInfo) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FastStringComparisonInfo) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FastStringComparisonInfo) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FastStringComparisonInfo) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


