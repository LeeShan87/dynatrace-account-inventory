# MetricMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the metric | [optional] 
**DisplayName** | Pointer to **string** | The name of the metric in the user interface | [optional] 
**Unit** | Pointer to **string** | The unit of the metric | [optional] 

## Methods

### NewMetricMetadataDto

`func NewMetricMetadataDto() *MetricMetadataDto`

NewMetricMetadataDto instantiates a new MetricMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricMetadataDtoWithDefaults

`func NewMetricMetadataDtoWithDefaults() *MetricMetadataDto`

NewMetricMetadataDtoWithDefaults instantiates a new MetricMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricMetadataDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricMetadataDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricMetadataDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricMetadataDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *MetricMetadataDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricMetadataDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetricMetadataDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MetricMetadataDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUnit

`func (o *MetricMetadataDto) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricMetadataDto) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricMetadataDto) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricMetadataDto) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


