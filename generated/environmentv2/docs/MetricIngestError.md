# MetricIngestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**InvalidLines** | Pointer to [**[]InvalidLine**](InvalidLine.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricIngestError

`func NewMetricIngestError() *MetricIngestError`

NewMetricIngestError instantiates a new MetricIngestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricIngestErrorWithDefaults

`func NewMetricIngestErrorWithDefaults() *MetricIngestError`

NewMetricIngestErrorWithDefaults instantiates a new MetricIngestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MetricIngestError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MetricIngestError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MetricIngestError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MetricIngestError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInvalidLines

`func (o *MetricIngestError) GetInvalidLines() []InvalidLine`

GetInvalidLines returns the InvalidLines field if non-nil, zero value otherwise.

### GetInvalidLinesOk

`func (o *MetricIngestError) GetInvalidLinesOk() (*[]InvalidLine, bool)`

GetInvalidLinesOk returns a tuple with the InvalidLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLines

`func (o *MetricIngestError) SetInvalidLines(v []InvalidLine)`

SetInvalidLines sets InvalidLines field to given value.

### HasInvalidLines

`func (o *MetricIngestError) HasInvalidLines() bool`

HasInvalidLines returns a boolean if a field has been set.

### GetMessage

`func (o *MetricIngestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetricIngestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetricIngestError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MetricIngestError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


