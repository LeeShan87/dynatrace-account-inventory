# EntryPointUsageSegment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SegmentType** | Pointer to **string** | The type of this input segment. | [optional] [readonly] 
**SegmentValue** | Pointer to **string** | The value of this input segment. | [optional] [readonly] 
**SourceArgumentName** | Pointer to **string** | The name used in the source for this segment. | [optional] [readonly] 
**SourceType** | Pointer to **string** | The type of the HTTP request part that contains the value that was used in this segment. | [optional] [readonly] 

## Methods

### NewEntryPointUsageSegment

`func NewEntryPointUsageSegment() *EntryPointUsageSegment`

NewEntryPointUsageSegment instantiates a new EntryPointUsageSegment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryPointUsageSegmentWithDefaults

`func NewEntryPointUsageSegmentWithDefaults() *EntryPointUsageSegment`

NewEntryPointUsageSegmentWithDefaults instantiates a new EntryPointUsageSegment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegmentType

`func (o *EntryPointUsageSegment) GetSegmentType() string`

GetSegmentType returns the SegmentType field if non-nil, zero value otherwise.

### GetSegmentTypeOk

`func (o *EntryPointUsageSegment) GetSegmentTypeOk() (*string, bool)`

GetSegmentTypeOk returns a tuple with the SegmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentType

`func (o *EntryPointUsageSegment) SetSegmentType(v string)`

SetSegmentType sets SegmentType field to given value.

### HasSegmentType

`func (o *EntryPointUsageSegment) HasSegmentType() bool`

HasSegmentType returns a boolean if a field has been set.

### GetSegmentValue

`func (o *EntryPointUsageSegment) GetSegmentValue() string`

GetSegmentValue returns the SegmentValue field if non-nil, zero value otherwise.

### GetSegmentValueOk

`func (o *EntryPointUsageSegment) GetSegmentValueOk() (*string, bool)`

GetSegmentValueOk returns a tuple with the SegmentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentValue

`func (o *EntryPointUsageSegment) SetSegmentValue(v string)`

SetSegmentValue sets SegmentValue field to given value.

### HasSegmentValue

`func (o *EntryPointUsageSegment) HasSegmentValue() bool`

HasSegmentValue returns a boolean if a field has been set.

### GetSourceArgumentName

`func (o *EntryPointUsageSegment) GetSourceArgumentName() string`

GetSourceArgumentName returns the SourceArgumentName field if non-nil, zero value otherwise.

### GetSourceArgumentNameOk

`func (o *EntryPointUsageSegment) GetSourceArgumentNameOk() (*string, bool)`

GetSourceArgumentNameOk returns a tuple with the SourceArgumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceArgumentName

`func (o *EntryPointUsageSegment) SetSourceArgumentName(v string)`

SetSourceArgumentName sets SourceArgumentName field to given value.

### HasSourceArgumentName

`func (o *EntryPointUsageSegment) HasSourceArgumentName() bool`

HasSourceArgumentName returns a boolean if a field has been set.

### GetSourceType

`func (o *EntryPointUsageSegment) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *EntryPointUsageSegment) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *EntryPointUsageSegment) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *EntryPointUsageSegment) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


