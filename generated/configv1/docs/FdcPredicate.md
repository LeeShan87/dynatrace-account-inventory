# FdcPredicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;STRING_EQUALS&#x60; -&gt; FdcPredicateStringEquals  * &#x60;STRING_STARTS_WITH&#x60; -&gt; FdcPredicateStringStartsWith  * &#x60;STRING_ENDS_WITH&#x60; -&gt; FdcPredicateStringEndsWith  * &#x60;STRING_CONTAINS_SUBSTRING&#x60; -&gt; FdcPredicateStringContains  * &#x60;INTEGER_EQUALS&#x60; -&gt; FdcPredicateIntegerEquals  * &#x60;INTEGER_LESS_THAN&#x60; -&gt; FdcPredicateIntegerLessThan  * &#x60;INTEGER_LESS_THAN_OR_EQUAL&#x60; -&gt; FdcPredicateIntegerLessThanOrEqual  * &#x60;INTEGER_GREATER_THAN&#x60; -&gt; FdcPredicateIntegerGreaterThan  * &#x60;INTEGER_GREATER_THAN_OR_EQUAL&#x60; -&gt; FdcPredicateIntegerGreaterThanOrEqual  * &#x60;LONG_EQUALS&#x60; -&gt; FdcPredicateLongEquals  * &#x60;LONG_LESS_THAN&#x60; -&gt; FdcPredicateLongLessThan  * &#x60;LONG_LESS_THAN_OR_EQUAL&#x60; -&gt; FdcPredicateLongLessThanOrEqual  * &#x60;LONG_GREATER_THAN&#x60; -&gt; FdcPredicateLongGreaterThan  * &#x60;LONG_GREATER_THAN_OR_EQUAL&#x60; -&gt; FdcPredicateLongGreaterThanOrEqual  * &#x60;TAG_KEY_EQUALS&#x60; -&gt; FdcPredicateTagKeyEquals  * &#x60;TAG_EQUALS&#x60; -&gt; FdcPredicateTagEquals  * &#x60;SERVICE_TYPE_EQUALS&#x60; -&gt; FdcPredicateServiceTypeEquals  * &#x60;MANAGEMENT_ZONES_CONTAINS_ALL&#x60; -&gt; FdcPredicateManagementZonesContainsAll  * &#x60;SET_OF_INTEGERS_CONTAINS_ANY&#x60; -&gt; FdcPredicateSetOfIntegersContainsAny  * &#x60;SET_OF_INTEGERS_CONTAINS_ALL&#x60; -&gt; FdcPredicateSetOfIntegersContainsAll   | 

## Methods

### NewFdcPredicate

`func NewFdcPredicate(type_ string, ) *FdcPredicate`

NewFdcPredicate instantiates a new FdcPredicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdcPredicateWithDefaults

`func NewFdcPredicateWithDefaults() *FdcPredicate`

NewFdcPredicateWithDefaults instantiates a new FdcPredicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FdcPredicate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FdcPredicate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FdcPredicate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


