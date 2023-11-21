# VulnerableFunctionInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSegments** | Pointer to [**[]VulnerableFunctionInputSegment**](VulnerableFunctionInputSegment.md) | A list of input segments. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the input. | [optional] [readonly] 

## Methods

### NewVulnerableFunctionInput

`func NewVulnerableFunctionInput() *VulnerableFunctionInput`

NewVulnerableFunctionInput instantiates a new VulnerableFunctionInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerableFunctionInputWithDefaults

`func NewVulnerableFunctionInputWithDefaults() *VulnerableFunctionInput`

NewVulnerableFunctionInputWithDefaults instantiates a new VulnerableFunctionInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputSegments

`func (o *VulnerableFunctionInput) GetInputSegments() []VulnerableFunctionInputSegment`

GetInputSegments returns the InputSegments field if non-nil, zero value otherwise.

### GetInputSegmentsOk

`func (o *VulnerableFunctionInput) GetInputSegmentsOk() (*[]VulnerableFunctionInputSegment, bool)`

GetInputSegmentsOk returns a tuple with the InputSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSegments

`func (o *VulnerableFunctionInput) SetInputSegments(v []VulnerableFunctionInputSegment)`

SetInputSegments sets InputSegments field to given value.

### HasInputSegments

`func (o *VulnerableFunctionInput) HasInputSegments() bool`

HasInputSegments returns a boolean if a field has been set.

### GetType

`func (o *VulnerableFunctionInput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VulnerableFunctionInput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VulnerableFunctionInput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VulnerableFunctionInput) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


