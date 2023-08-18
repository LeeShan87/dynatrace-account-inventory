# MetricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Metric key | [optional] 
**Metadata** | Pointer to [**MetricMetadataDto**](MetricMetadataDto.md) |  | [optional] 

## Methods

### NewMetricDto

`func NewMetricDto() *MetricDto`

NewMetricDto instantiates a new MetricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDtoWithDefaults

`func NewMetricDtoWithDefaults() *MetricDto`

NewMetricDtoWithDefaults instantiates a new MetricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMetadata

`func (o *MetricDto) GetMetadata() MetricMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MetricDto) GetMetadataOk() (*MetricMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MetricDto) SetMetadata(v MetricMetadataDto)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MetricDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


