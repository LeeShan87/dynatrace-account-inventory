# FdpTagPredicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STRING_EXISTS&#x60; -&gt; FdpTagStringExists  * &#x60;STRING_EQUALS&#x60; -&gt; FdpTagStringEquals  * &#x60;STRING_STARTS_WITH&#x60; -&gt; FdpTagStringStartsWith  * &#x60;STRING_ENDS_WITH&#x60; -&gt; FdpTagStringEndsWith  * &#x60;STRING_CONTAINS_SUBSTRING&#x60; -&gt; FdpTagStringContainsSubstring  * &#x60;INTEGER_EXISTS&#x60; -&gt; FdpTagIntegerExists  * &#x60;INTEGER_EQUALS&#x60; -&gt; FdpTagIntegerEquals  * &#x60;INTEGER_LESS_THAN&#x60; -&gt; FdpTagIntegerLessThan  * &#x60;INTEGER_LESS_THAN_OR_EQUAL&#x60; -&gt; FdpTagIntegerLessThanOrEqual  * &#x60;INTEGER_GREATER_THAN&#x60; -&gt; FdpTagIntegerGreaterThan  * &#x60;INTEGER_GREATER_THAN_OR_EQUAL&#x60; -&gt; FdpTagIntegerGreaterThanOrEqual  * &#x60;DOUBLE_EXISTS&#x60; -&gt; FdpTagDoubleExists  * &#x60;DOUBLE_EQUALS&#x60; -&gt; FdpTagDoubleEquals  * &#x60;DOUBLE_LESS_THAN&#x60; -&gt; FdpTagDoubleLessThan  * &#x60;DOUBLE_LESS_THAN_OR_EQUAL&#x60; -&gt; FdpTagDoubleLessThanOrEqual  * &#x60;DOUBLE_GREATER_THAN&#x60; -&gt; FdpTagDoubleGreaterThan  * &#x60;DOUBLE_GREATER_THAN_OR_EQUAL&#x60; -&gt; FdpTagDoubleGreaterThanOrEqual   | 

## Methods

### NewFdpTagPredicate

`func NewFdpTagPredicate(type_ string, ) *FdpTagPredicate`

NewFdpTagPredicate instantiates a new FdpTagPredicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdpTagPredicateWithDefaults

`func NewFdpTagPredicateWithDefaults() *FdpTagPredicate`

NewFdpTagPredicateWithDefaults instantiates a new FdpTagPredicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FdpTagPredicate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FdpTagPredicate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FdpTagPredicate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


