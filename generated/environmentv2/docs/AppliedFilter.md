# AppliedFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedTo** | **[]string** | The keys of all metrics that this filter has been applied to.   Can contain multiple metrics for complex expressions and always at least one key. | 
**Filter** | Pointer to [**Filter**](Filter.md) |  | [optional] 

## Methods

### NewAppliedFilter

`func NewAppliedFilter(appliedTo []string, ) *AppliedFilter`

NewAppliedFilter instantiates a new AppliedFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppliedFilterWithDefaults

`func NewAppliedFilterWithDefaults() *AppliedFilter`

NewAppliedFilterWithDefaults instantiates a new AppliedFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedTo

`func (o *AppliedFilter) GetAppliedTo() []string`

GetAppliedTo returns the AppliedTo field if non-nil, zero value otherwise.

### GetAppliedToOk

`func (o *AppliedFilter) GetAppliedToOk() (*[]string, bool)`

GetAppliedToOk returns a tuple with the AppliedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedTo

`func (o *AppliedFilter) SetAppliedTo(v []string)`

SetAppliedTo sets AppliedTo field to given value.


### GetFilter

`func (o *AppliedFilter) GetFilter() Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AppliedFilter) GetFilterOk() (*Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AppliedFilter) SetFilter(v Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AppliedFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


