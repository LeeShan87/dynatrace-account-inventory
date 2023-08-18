# ExportedLogRecordList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**Results** | Pointer to [**[]LogRecord**](LogRecord.md) | A list of retrieved log records. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewExportedLogRecordList

`func NewExportedLogRecordList(totalCount int64, ) *ExportedLogRecordList`

NewExportedLogRecordList instantiates a new ExportedLogRecordList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportedLogRecordListWithDefaults

`func NewExportedLogRecordListWithDefaults() *ExportedLogRecordList`

NewExportedLogRecordListWithDefaults instantiates a new ExportedLogRecordList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageKey

`func (o *ExportedLogRecordList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExportedLogRecordList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExportedLogRecordList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ExportedLogRecordList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ExportedLogRecordList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ExportedLogRecordList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ExportedLogRecordList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ExportedLogRecordList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetResults

`func (o *ExportedLogRecordList) GetResults() []LogRecord`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ExportedLogRecordList) GetResultsOk() (*[]LogRecord, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ExportedLogRecordList) SetResults(v []LogRecord)`

SetResults sets Results field to given value.

### HasResults

`func (o *ExportedLogRecordList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalCount

`func (o *ExportedLogRecordList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExportedLogRecordList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExportedLogRecordList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


