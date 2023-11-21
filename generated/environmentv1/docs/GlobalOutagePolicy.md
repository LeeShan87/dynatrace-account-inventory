# GlobalOutagePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsecutiveRuns** | **int32** | Alert if all locations are unable to access the web application *X* times consecutively. | 

## Methods

### NewGlobalOutagePolicy

`func NewGlobalOutagePolicy(consecutiveRuns int32, ) *GlobalOutagePolicy`

NewGlobalOutagePolicy instantiates a new GlobalOutagePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalOutagePolicyWithDefaults

`func NewGlobalOutagePolicyWithDefaults() *GlobalOutagePolicy`

NewGlobalOutagePolicyWithDefaults instantiates a new GlobalOutagePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsecutiveRuns

`func (o *GlobalOutagePolicy) GetConsecutiveRuns() int32`

GetConsecutiveRuns returns the ConsecutiveRuns field if non-nil, zero value otherwise.

### GetConsecutiveRunsOk

`func (o *GlobalOutagePolicy) GetConsecutiveRunsOk() (*int32, bool)`

GetConsecutiveRunsOk returns a tuple with the ConsecutiveRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveRuns

`func (o *GlobalOutagePolicy) SetConsecutiveRuns(v int32)`

SetConsecutiveRuns sets ConsecutiveRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


