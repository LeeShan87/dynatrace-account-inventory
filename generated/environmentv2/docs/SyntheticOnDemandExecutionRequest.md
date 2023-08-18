# SyntheticOnDemandExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailOnPerformanceIssue** | Pointer to **bool** | If true, the execution will fail in case of performance issue. | [optional] [default to true]
**FailOnSslWarning** | Pointer to **bool** | Applies to HTTP monitors only. If true, the execution will fail in case of an SSL certificate expiration warning or if the certificate is missing. | [optional] [default to true]
**Group** | Pointer to [**SyntheticOnDemandExecutionRequestGroup**](SyntheticOnDemandExecutionRequestGroup.md) |  | [optional] 
**Metadata** | Pointer to **map[string]string** | String to string map of metadata properties for execution | [optional] 
**Monitors** | Pointer to [**[]SyntheticOnDemandExecutionRequestMonitor**](SyntheticOnDemandExecutionRequestMonitor.md) | List of monitors to be triggered. | [optional] 
**ProcessingMode** | Pointer to **string** | The execution&#39;s processing mode | [optional] [default to "STANDARD"]
**StopOnProblem** | Pointer to **bool** | If true, no executions will be scheduled if a problem occurs. | [optional] [default to false]
**TakeScreenshotsOnSuccess** | Pointer to **bool** | If true, the screenshots will be taken during the execution of a browser monitor. | [optional] [default to false]

## Methods

### NewSyntheticOnDemandExecutionRequest

`func NewSyntheticOnDemandExecutionRequest() *SyntheticOnDemandExecutionRequest`

NewSyntheticOnDemandExecutionRequest instantiates a new SyntheticOnDemandExecutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandExecutionRequestWithDefaults

`func NewSyntheticOnDemandExecutionRequestWithDefaults() *SyntheticOnDemandExecutionRequest`

NewSyntheticOnDemandExecutionRequestWithDefaults instantiates a new SyntheticOnDemandExecutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailOnPerformanceIssue

`func (o *SyntheticOnDemandExecutionRequest) GetFailOnPerformanceIssue() bool`

GetFailOnPerformanceIssue returns the FailOnPerformanceIssue field if non-nil, zero value otherwise.

### GetFailOnPerformanceIssueOk

`func (o *SyntheticOnDemandExecutionRequest) GetFailOnPerformanceIssueOk() (*bool, bool)`

GetFailOnPerformanceIssueOk returns a tuple with the FailOnPerformanceIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnPerformanceIssue

`func (o *SyntheticOnDemandExecutionRequest) SetFailOnPerformanceIssue(v bool)`

SetFailOnPerformanceIssue sets FailOnPerformanceIssue field to given value.

### HasFailOnPerformanceIssue

`func (o *SyntheticOnDemandExecutionRequest) HasFailOnPerformanceIssue() bool`

HasFailOnPerformanceIssue returns a boolean if a field has been set.

### GetFailOnSslWarning

`func (o *SyntheticOnDemandExecutionRequest) GetFailOnSslWarning() bool`

GetFailOnSslWarning returns the FailOnSslWarning field if non-nil, zero value otherwise.

### GetFailOnSslWarningOk

`func (o *SyntheticOnDemandExecutionRequest) GetFailOnSslWarningOk() (*bool, bool)`

GetFailOnSslWarningOk returns a tuple with the FailOnSslWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnSslWarning

`func (o *SyntheticOnDemandExecutionRequest) SetFailOnSslWarning(v bool)`

SetFailOnSslWarning sets FailOnSslWarning field to given value.

### HasFailOnSslWarning

`func (o *SyntheticOnDemandExecutionRequest) HasFailOnSslWarning() bool`

HasFailOnSslWarning returns a boolean if a field has been set.

### GetGroup

`func (o *SyntheticOnDemandExecutionRequest) GetGroup() SyntheticOnDemandExecutionRequestGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SyntheticOnDemandExecutionRequest) GetGroupOk() (*SyntheticOnDemandExecutionRequestGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SyntheticOnDemandExecutionRequest) SetGroup(v SyntheticOnDemandExecutionRequestGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SyntheticOnDemandExecutionRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMetadata

`func (o *SyntheticOnDemandExecutionRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SyntheticOnDemandExecutionRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SyntheticOnDemandExecutionRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SyntheticOnDemandExecutionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonitors

`func (o *SyntheticOnDemandExecutionRequest) GetMonitors() []SyntheticOnDemandExecutionRequestMonitor`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *SyntheticOnDemandExecutionRequest) GetMonitorsOk() (*[]SyntheticOnDemandExecutionRequestMonitor, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *SyntheticOnDemandExecutionRequest) SetMonitors(v []SyntheticOnDemandExecutionRequestMonitor)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *SyntheticOnDemandExecutionRequest) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetProcessingMode

`func (o *SyntheticOnDemandExecutionRequest) GetProcessingMode() string`

GetProcessingMode returns the ProcessingMode field if non-nil, zero value otherwise.

### GetProcessingModeOk

`func (o *SyntheticOnDemandExecutionRequest) GetProcessingModeOk() (*string, bool)`

GetProcessingModeOk returns a tuple with the ProcessingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMode

`func (o *SyntheticOnDemandExecutionRequest) SetProcessingMode(v string)`

SetProcessingMode sets ProcessingMode field to given value.

### HasProcessingMode

`func (o *SyntheticOnDemandExecutionRequest) HasProcessingMode() bool`

HasProcessingMode returns a boolean if a field has been set.

### GetStopOnProblem

`func (o *SyntheticOnDemandExecutionRequest) GetStopOnProblem() bool`

GetStopOnProblem returns the StopOnProblem field if non-nil, zero value otherwise.

### GetStopOnProblemOk

`func (o *SyntheticOnDemandExecutionRequest) GetStopOnProblemOk() (*bool, bool)`

GetStopOnProblemOk returns a tuple with the StopOnProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnProblem

`func (o *SyntheticOnDemandExecutionRequest) SetStopOnProblem(v bool)`

SetStopOnProblem sets StopOnProblem field to given value.

### HasStopOnProblem

`func (o *SyntheticOnDemandExecutionRequest) HasStopOnProblem() bool`

HasStopOnProblem returns a boolean if a field has been set.

### GetTakeScreenshotsOnSuccess

`func (o *SyntheticOnDemandExecutionRequest) GetTakeScreenshotsOnSuccess() bool`

GetTakeScreenshotsOnSuccess returns the TakeScreenshotsOnSuccess field if non-nil, zero value otherwise.

### GetTakeScreenshotsOnSuccessOk

`func (o *SyntheticOnDemandExecutionRequest) GetTakeScreenshotsOnSuccessOk() (*bool, bool)`

GetTakeScreenshotsOnSuccessOk returns a tuple with the TakeScreenshotsOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeScreenshotsOnSuccess

`func (o *SyntheticOnDemandExecutionRequest) SetTakeScreenshotsOnSuccess(v bool)`

SetTakeScreenshotsOnSuccess sets TakeScreenshotsOnSuccess field to given value.

### HasTakeScreenshotsOnSuccess

`func (o *SyntheticOnDemandExecutionRequest) HasTakeScreenshotsOnSuccess() bool`

HasTakeScreenshotsOnSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


