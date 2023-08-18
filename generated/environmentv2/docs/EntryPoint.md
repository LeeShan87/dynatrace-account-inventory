# EntryPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHttpPath** | Pointer to **string** | Source HTTP path of entry points. | [optional] [readonly] 
**UsageSegments** | Pointer to [**[]EntryPointUsageSegment**](EntryPointUsageSegment.md) | List of entry point usage segments. | [optional] [readonly] 

## Methods

### NewEntryPoint

`func NewEntryPoint() *EntryPoint`

NewEntryPoint instantiates a new EntryPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryPointWithDefaults

`func NewEntryPointWithDefaults() *EntryPoint`

NewEntryPointWithDefaults instantiates a new EntryPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceHttpPath

`func (o *EntryPoint) GetSourceHttpPath() string`

GetSourceHttpPath returns the SourceHttpPath field if non-nil, zero value otherwise.

### GetSourceHttpPathOk

`func (o *EntryPoint) GetSourceHttpPathOk() (*string, bool)`

GetSourceHttpPathOk returns a tuple with the SourceHttpPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHttpPath

`func (o *EntryPoint) SetSourceHttpPath(v string)`

SetSourceHttpPath sets SourceHttpPath field to given value.

### HasSourceHttpPath

`func (o *EntryPoint) HasSourceHttpPath() bool`

HasSourceHttpPath returns a boolean if a field has been set.

### GetUsageSegments

`func (o *EntryPoint) GetUsageSegments() []EntryPointUsageSegment`

GetUsageSegments returns the UsageSegments field if non-nil, zero value otherwise.

### GetUsageSegmentsOk

`func (o *EntryPoint) GetUsageSegmentsOk() (*[]EntryPointUsageSegment, bool)`

GetUsageSegmentsOk returns a tuple with the UsageSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageSegments

`func (o *EntryPoint) SetUsageSegments(v []EntryPointUsageSegment)`

SetUsageSegments sets UsageSegments field to given value.

### HasUsageSegments

`func (o *EntryPoint) HasUsageSegments() bool`

HasUsageSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


