# FailureDetectionParameterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrokenLinkDomains** | Pointer to **[]string** | A list of domains for the special handling of the 404 HTTP response code.   If the top domain of the *Referer* is presented in the list OR equals the top domain of the request&#39;s host, the 404 code is considered a failure.   Only applicable when **isHttp404NotFoundFailureEnabled** is set to &#x60;true&#x60;. | [optional] 
**ClientSideMissingHttpCodeIsFailure** | Pointer to **bool** | The missing HTTP response code on the client side is a failure (&#x60;true&#x60;) or a success (&#x60;false&#x60;).   If not set, &#x60;false&#x60; is used. | [optional] 
**Description** | Pointer to **string** | A short description of the FDP set. | [optional] 
**ExceptionOnAnyNodeExceptionPatterns** | Pointer to [**[]ExceptionPattern**](ExceptionPattern.md) | A list of faulty exceptions.   If an exception on *any* node of the service matches *any* of these patterns it is considered a failure. | [optional] 
**FailingHttpCodesClientSide** | Pointer to **string** | A list of client side HTTP response codes that are considered a failure.   The format is a list of ranges and individual values (for example, &#x60;500-599, 400-499, 1008&#x60;).   If not set, the range of &#x60;400-599&#x60; is used. | [optional] 
**FailingHttpCodesServerSide** | Pointer to **string** | A list of server side HTTP response codes that are considered a failure.   The format is a list of ranges and individual values (for example, &#x60;500-599, 400-499, 1008&#x60;).If not set, the range of &#x60;500-599&#x60; is used. | [optional] 
**Http404NotFoundFailureEnabled** | Pointer to **bool** | Special handling of the 404 HTTP response code is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). See **brokenLinkDomains** for special handling details.   Only applicable when 404 is NOT in the list of failing HTTP response codes and only for the server side.   If not set, &#x60;false&#x60; is used. | [optional] 
**Id** | Pointer to **string** | The ID of the parameter set. | [optional] 
**IgnoreAllExceptions** | Pointer to **bool** | If set to true all exceptions will be ignored. Which means defined exception patterns will not have any effect. | [optional] 
**IgnoreSpanFailureDetection** | Pointer to **bool** | If set to true span failure detection will be ignored. | [optional] 
**IgnoredExceptionPatterns** | Pointer to [**[]ExceptionPattern**](ExceptionPattern.md) | A list of ignored exceptions.   If an exception on the entry node of the service matches *any* of these patterns it is ignored by failure detection. | [optional] 
**Name** | Pointer to **string** | The display name of the FDP set.   The length of the name is limited to 150 characters. | [optional] 
**ServerSideMissingHttpCodeIsFailure** | Pointer to **bool** | The missing HTTP response code on the server side is a failure (&#x60;true&#x60;) or a success (&#x60;false&#x60;).   If not set, &#x60;false&#x60; is used. | [optional] 
**SuccessForcingExceptionPatterns** | Pointer to [**[]ExceptionPattern**](ExceptionPattern.md) | A list of success exceptions.   If an exception on the entry node of the service matches *any* of these patterns it is considered a success. | [optional] 
**TagConditions** | Pointer to [**[]FdpTagCondition**](FdpTagCondition.md) | A list of tag-based conditions.   If *any* condition is fulfilled the request is considered a failure. | [optional] 

## Methods

### NewFailureDetectionParameterSet

`func NewFailureDetectionParameterSet() *FailureDetectionParameterSet`

NewFailureDetectionParameterSet instantiates a new FailureDetectionParameterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureDetectionParameterSetWithDefaults

`func NewFailureDetectionParameterSetWithDefaults() *FailureDetectionParameterSet`

NewFailureDetectionParameterSetWithDefaults instantiates a new FailureDetectionParameterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokenLinkDomains

`func (o *FailureDetectionParameterSet) GetBrokenLinkDomains() []string`

GetBrokenLinkDomains returns the BrokenLinkDomains field if non-nil, zero value otherwise.

### GetBrokenLinkDomainsOk

`func (o *FailureDetectionParameterSet) GetBrokenLinkDomainsOk() (*[]string, bool)`

GetBrokenLinkDomainsOk returns a tuple with the BrokenLinkDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokenLinkDomains

`func (o *FailureDetectionParameterSet) SetBrokenLinkDomains(v []string)`

SetBrokenLinkDomains sets BrokenLinkDomains field to given value.

### HasBrokenLinkDomains

`func (o *FailureDetectionParameterSet) HasBrokenLinkDomains() bool`

HasBrokenLinkDomains returns a boolean if a field has been set.

### GetClientSideMissingHttpCodeIsFailure

`func (o *FailureDetectionParameterSet) GetClientSideMissingHttpCodeIsFailure() bool`

GetClientSideMissingHttpCodeIsFailure returns the ClientSideMissingHttpCodeIsFailure field if non-nil, zero value otherwise.

### GetClientSideMissingHttpCodeIsFailureOk

`func (o *FailureDetectionParameterSet) GetClientSideMissingHttpCodeIsFailureOk() (*bool, bool)`

GetClientSideMissingHttpCodeIsFailureOk returns a tuple with the ClientSideMissingHttpCodeIsFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSideMissingHttpCodeIsFailure

`func (o *FailureDetectionParameterSet) SetClientSideMissingHttpCodeIsFailure(v bool)`

SetClientSideMissingHttpCodeIsFailure sets ClientSideMissingHttpCodeIsFailure field to given value.

### HasClientSideMissingHttpCodeIsFailure

`func (o *FailureDetectionParameterSet) HasClientSideMissingHttpCodeIsFailure() bool`

HasClientSideMissingHttpCodeIsFailure returns a boolean if a field has been set.

### GetDescription

`func (o *FailureDetectionParameterSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FailureDetectionParameterSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FailureDetectionParameterSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FailureDetectionParameterSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExceptionOnAnyNodeExceptionPatterns

`func (o *FailureDetectionParameterSet) GetExceptionOnAnyNodeExceptionPatterns() []ExceptionPattern`

GetExceptionOnAnyNodeExceptionPatterns returns the ExceptionOnAnyNodeExceptionPatterns field if non-nil, zero value otherwise.

### GetExceptionOnAnyNodeExceptionPatternsOk

`func (o *FailureDetectionParameterSet) GetExceptionOnAnyNodeExceptionPatternsOk() (*[]ExceptionPattern, bool)`

GetExceptionOnAnyNodeExceptionPatternsOk returns a tuple with the ExceptionOnAnyNodeExceptionPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionOnAnyNodeExceptionPatterns

`func (o *FailureDetectionParameterSet) SetExceptionOnAnyNodeExceptionPatterns(v []ExceptionPattern)`

SetExceptionOnAnyNodeExceptionPatterns sets ExceptionOnAnyNodeExceptionPatterns field to given value.

### HasExceptionOnAnyNodeExceptionPatterns

`func (o *FailureDetectionParameterSet) HasExceptionOnAnyNodeExceptionPatterns() bool`

HasExceptionOnAnyNodeExceptionPatterns returns a boolean if a field has been set.

### GetFailingHttpCodesClientSide

`func (o *FailureDetectionParameterSet) GetFailingHttpCodesClientSide() string`

GetFailingHttpCodesClientSide returns the FailingHttpCodesClientSide field if non-nil, zero value otherwise.

### GetFailingHttpCodesClientSideOk

`func (o *FailureDetectionParameterSet) GetFailingHttpCodesClientSideOk() (*string, bool)`

GetFailingHttpCodesClientSideOk returns a tuple with the FailingHttpCodesClientSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingHttpCodesClientSide

`func (o *FailureDetectionParameterSet) SetFailingHttpCodesClientSide(v string)`

SetFailingHttpCodesClientSide sets FailingHttpCodesClientSide field to given value.

### HasFailingHttpCodesClientSide

`func (o *FailureDetectionParameterSet) HasFailingHttpCodesClientSide() bool`

HasFailingHttpCodesClientSide returns a boolean if a field has been set.

### GetFailingHttpCodesServerSide

`func (o *FailureDetectionParameterSet) GetFailingHttpCodesServerSide() string`

GetFailingHttpCodesServerSide returns the FailingHttpCodesServerSide field if non-nil, zero value otherwise.

### GetFailingHttpCodesServerSideOk

`func (o *FailureDetectionParameterSet) GetFailingHttpCodesServerSideOk() (*string, bool)`

GetFailingHttpCodesServerSideOk returns a tuple with the FailingHttpCodesServerSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingHttpCodesServerSide

`func (o *FailureDetectionParameterSet) SetFailingHttpCodesServerSide(v string)`

SetFailingHttpCodesServerSide sets FailingHttpCodesServerSide field to given value.

### HasFailingHttpCodesServerSide

`func (o *FailureDetectionParameterSet) HasFailingHttpCodesServerSide() bool`

HasFailingHttpCodesServerSide returns a boolean if a field has been set.

### GetHttp404NotFoundFailureEnabled

`func (o *FailureDetectionParameterSet) GetHttp404NotFoundFailureEnabled() bool`

GetHttp404NotFoundFailureEnabled returns the Http404NotFoundFailureEnabled field if non-nil, zero value otherwise.

### GetHttp404NotFoundFailureEnabledOk

`func (o *FailureDetectionParameterSet) GetHttp404NotFoundFailureEnabledOk() (*bool, bool)`

GetHttp404NotFoundFailureEnabledOk returns a tuple with the Http404NotFoundFailureEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp404NotFoundFailureEnabled

`func (o *FailureDetectionParameterSet) SetHttp404NotFoundFailureEnabled(v bool)`

SetHttp404NotFoundFailureEnabled sets Http404NotFoundFailureEnabled field to given value.

### HasHttp404NotFoundFailureEnabled

`func (o *FailureDetectionParameterSet) HasHttp404NotFoundFailureEnabled() bool`

HasHttp404NotFoundFailureEnabled returns a boolean if a field has been set.

### GetId

`func (o *FailureDetectionParameterSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureDetectionParameterSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureDetectionParameterSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FailureDetectionParameterSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreAllExceptions

`func (o *FailureDetectionParameterSet) GetIgnoreAllExceptions() bool`

GetIgnoreAllExceptions returns the IgnoreAllExceptions field if non-nil, zero value otherwise.

### GetIgnoreAllExceptionsOk

`func (o *FailureDetectionParameterSet) GetIgnoreAllExceptionsOk() (*bool, bool)`

GetIgnoreAllExceptionsOk returns a tuple with the IgnoreAllExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreAllExceptions

`func (o *FailureDetectionParameterSet) SetIgnoreAllExceptions(v bool)`

SetIgnoreAllExceptions sets IgnoreAllExceptions field to given value.

### HasIgnoreAllExceptions

`func (o *FailureDetectionParameterSet) HasIgnoreAllExceptions() bool`

HasIgnoreAllExceptions returns a boolean if a field has been set.

### GetIgnoreSpanFailureDetection

`func (o *FailureDetectionParameterSet) GetIgnoreSpanFailureDetection() bool`

GetIgnoreSpanFailureDetection returns the IgnoreSpanFailureDetection field if non-nil, zero value otherwise.

### GetIgnoreSpanFailureDetectionOk

`func (o *FailureDetectionParameterSet) GetIgnoreSpanFailureDetectionOk() (*bool, bool)`

GetIgnoreSpanFailureDetectionOk returns a tuple with the IgnoreSpanFailureDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSpanFailureDetection

`func (o *FailureDetectionParameterSet) SetIgnoreSpanFailureDetection(v bool)`

SetIgnoreSpanFailureDetection sets IgnoreSpanFailureDetection field to given value.

### HasIgnoreSpanFailureDetection

`func (o *FailureDetectionParameterSet) HasIgnoreSpanFailureDetection() bool`

HasIgnoreSpanFailureDetection returns a boolean if a field has been set.

### GetIgnoredExceptionPatterns

`func (o *FailureDetectionParameterSet) GetIgnoredExceptionPatterns() []ExceptionPattern`

GetIgnoredExceptionPatterns returns the IgnoredExceptionPatterns field if non-nil, zero value otherwise.

### GetIgnoredExceptionPatternsOk

`func (o *FailureDetectionParameterSet) GetIgnoredExceptionPatternsOk() (*[]ExceptionPattern, bool)`

GetIgnoredExceptionPatternsOk returns a tuple with the IgnoredExceptionPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredExceptionPatterns

`func (o *FailureDetectionParameterSet) SetIgnoredExceptionPatterns(v []ExceptionPattern)`

SetIgnoredExceptionPatterns sets IgnoredExceptionPatterns field to given value.

### HasIgnoredExceptionPatterns

`func (o *FailureDetectionParameterSet) HasIgnoredExceptionPatterns() bool`

HasIgnoredExceptionPatterns returns a boolean if a field has been set.

### GetName

`func (o *FailureDetectionParameterSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FailureDetectionParameterSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FailureDetectionParameterSet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FailureDetectionParameterSet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerSideMissingHttpCodeIsFailure

`func (o *FailureDetectionParameterSet) GetServerSideMissingHttpCodeIsFailure() bool`

GetServerSideMissingHttpCodeIsFailure returns the ServerSideMissingHttpCodeIsFailure field if non-nil, zero value otherwise.

### GetServerSideMissingHttpCodeIsFailureOk

`func (o *FailureDetectionParameterSet) GetServerSideMissingHttpCodeIsFailureOk() (*bool, bool)`

GetServerSideMissingHttpCodeIsFailureOk returns a tuple with the ServerSideMissingHttpCodeIsFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideMissingHttpCodeIsFailure

`func (o *FailureDetectionParameterSet) SetServerSideMissingHttpCodeIsFailure(v bool)`

SetServerSideMissingHttpCodeIsFailure sets ServerSideMissingHttpCodeIsFailure field to given value.

### HasServerSideMissingHttpCodeIsFailure

`func (o *FailureDetectionParameterSet) HasServerSideMissingHttpCodeIsFailure() bool`

HasServerSideMissingHttpCodeIsFailure returns a boolean if a field has been set.

### GetSuccessForcingExceptionPatterns

`func (o *FailureDetectionParameterSet) GetSuccessForcingExceptionPatterns() []ExceptionPattern`

GetSuccessForcingExceptionPatterns returns the SuccessForcingExceptionPatterns field if non-nil, zero value otherwise.

### GetSuccessForcingExceptionPatternsOk

`func (o *FailureDetectionParameterSet) GetSuccessForcingExceptionPatternsOk() (*[]ExceptionPattern, bool)`

GetSuccessForcingExceptionPatternsOk returns a tuple with the SuccessForcingExceptionPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessForcingExceptionPatterns

`func (o *FailureDetectionParameterSet) SetSuccessForcingExceptionPatterns(v []ExceptionPattern)`

SetSuccessForcingExceptionPatterns sets SuccessForcingExceptionPatterns field to given value.

### HasSuccessForcingExceptionPatterns

`func (o *FailureDetectionParameterSet) HasSuccessForcingExceptionPatterns() bool`

HasSuccessForcingExceptionPatterns returns a boolean if a field has been set.

### GetTagConditions

`func (o *FailureDetectionParameterSet) GetTagConditions() []FdpTagCondition`

GetTagConditions returns the TagConditions field if non-nil, zero value otherwise.

### GetTagConditionsOk

`func (o *FailureDetectionParameterSet) GetTagConditionsOk() (*[]FdpTagCondition, bool)`

GetTagConditionsOk returns a tuple with the TagConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagConditions

`func (o *FailureDetectionParameterSet) SetTagConditions(v []FdpTagCondition)`

SetTagConditions sets TagConditions field to given value.

### HasTagConditions

`func (o *FailureDetectionParameterSet) HasTagConditions() bool`

HasTagConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


