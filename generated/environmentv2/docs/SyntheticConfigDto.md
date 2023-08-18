# SyntheticConfigDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BmMonitorTimeout** | **int64** | bmMonitorTimeout - browser monitor execution timeout (ms) | 
**BmStepTimeout** | **int64** | bmStepTimeout - browser monitor single step execution timeout (ms) | 

## Methods

### NewSyntheticConfigDto

`func NewSyntheticConfigDto(bmMonitorTimeout int64, bmStepTimeout int64, ) *SyntheticConfigDto`

NewSyntheticConfigDto instantiates a new SyntheticConfigDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticConfigDtoWithDefaults

`func NewSyntheticConfigDtoWithDefaults() *SyntheticConfigDto`

NewSyntheticConfigDtoWithDefaults instantiates a new SyntheticConfigDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBmMonitorTimeout

`func (o *SyntheticConfigDto) GetBmMonitorTimeout() int64`

GetBmMonitorTimeout returns the BmMonitorTimeout field if non-nil, zero value otherwise.

### GetBmMonitorTimeoutOk

`func (o *SyntheticConfigDto) GetBmMonitorTimeoutOk() (*int64, bool)`

GetBmMonitorTimeoutOk returns a tuple with the BmMonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmMonitorTimeout

`func (o *SyntheticConfigDto) SetBmMonitorTimeout(v int64)`

SetBmMonitorTimeout sets BmMonitorTimeout field to given value.


### GetBmStepTimeout

`func (o *SyntheticConfigDto) GetBmStepTimeout() int64`

GetBmStepTimeout returns the BmStepTimeout field if non-nil, zero value otherwise.

### GetBmStepTimeoutOk

`func (o *SyntheticConfigDto) GetBmStepTimeoutOk() (*int64, bool)`

GetBmStepTimeoutOk returns a tuple with the BmStepTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmStepTimeout

`func (o *SyntheticConfigDto) SetBmStepTimeout(v int64)`

SetBmStepTimeout sets BmStepTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


