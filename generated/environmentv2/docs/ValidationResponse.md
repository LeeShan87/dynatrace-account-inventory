# ValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**MetricIngestError**](MetricIngestError.md) |  | [optional] 
**LinesInvalid** | Pointer to **int32** |  | [optional] 
**LinesOk** | Pointer to **int32** |  | [optional] 
**Warnings** | Pointer to [**Warnings**](Warnings.md) |  | [optional] 

## Methods

### NewValidationResponse

`func NewValidationResponse() *ValidationResponse`

NewValidationResponse instantiates a new ValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResponseWithDefaults

`func NewValidationResponseWithDefaults() *ValidationResponse`

NewValidationResponseWithDefaults instantiates a new ValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ValidationResponse) GetError() MetricIngestError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidationResponse) GetErrorOk() (*MetricIngestError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidationResponse) SetError(v MetricIngestError)`

SetError sets Error field to given value.

### HasError

`func (o *ValidationResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLinesInvalid

`func (o *ValidationResponse) GetLinesInvalid() int32`

GetLinesInvalid returns the LinesInvalid field if non-nil, zero value otherwise.

### GetLinesInvalidOk

`func (o *ValidationResponse) GetLinesInvalidOk() (*int32, bool)`

GetLinesInvalidOk returns a tuple with the LinesInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesInvalid

`func (o *ValidationResponse) SetLinesInvalid(v int32)`

SetLinesInvalid sets LinesInvalid field to given value.

### HasLinesInvalid

`func (o *ValidationResponse) HasLinesInvalid() bool`

HasLinesInvalid returns a boolean if a field has been set.

### GetLinesOk

`func (o *ValidationResponse) GetLinesOk() int32`

GetLinesOk returns the LinesOk field if non-nil, zero value otherwise.

### GetLinesOkOk

`func (o *ValidationResponse) GetLinesOkOk() (*int32, bool)`

GetLinesOkOk returns a tuple with the LinesOk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesOk

`func (o *ValidationResponse) SetLinesOk(v int32)`

SetLinesOk sets LinesOk field to given value.

### HasLinesOk

`func (o *ValidationResponse) HasLinesOk() bool`

HasLinesOk returns a boolean if a field has been set.

### GetWarnings

`func (o *ValidationResponse) GetWarnings() Warnings`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidationResponse) GetWarningsOk() (*Warnings, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidationResponse) SetWarnings(v Warnings)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ValidationResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


