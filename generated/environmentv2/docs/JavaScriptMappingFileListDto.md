# JavaScriptMappingFileListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JsMappingFiles** | Pointer to [**[]JavaScriptMappingFileDto**](JavaScriptMappingFileDto.md) | A list of JavaScript mapping files. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewJavaScriptMappingFileListDto

`func NewJavaScriptMappingFileListDto(totalCount int64, ) *JavaScriptMappingFileListDto`

NewJavaScriptMappingFileListDto instantiates a new JavaScriptMappingFileListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJavaScriptMappingFileListDtoWithDefaults

`func NewJavaScriptMappingFileListDtoWithDefaults() *JavaScriptMappingFileListDto`

NewJavaScriptMappingFileListDtoWithDefaults instantiates a new JavaScriptMappingFileListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJsMappingFiles

`func (o *JavaScriptMappingFileListDto) GetJsMappingFiles() []JavaScriptMappingFileDto`

GetJsMappingFiles returns the JsMappingFiles field if non-nil, zero value otherwise.

### GetJsMappingFilesOk

`func (o *JavaScriptMappingFileListDto) GetJsMappingFilesOk() (*[]JavaScriptMappingFileDto, bool)`

GetJsMappingFilesOk returns a tuple with the JsMappingFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsMappingFiles

`func (o *JavaScriptMappingFileListDto) SetJsMappingFiles(v []JavaScriptMappingFileDto)`

SetJsMappingFiles sets JsMappingFiles field to given value.

### HasJsMappingFiles

`func (o *JavaScriptMappingFileListDto) HasJsMappingFiles() bool`

HasJsMappingFiles returns a boolean if a field has been set.

### GetNextPageKey

`func (o *JavaScriptMappingFileListDto) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *JavaScriptMappingFileListDto) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *JavaScriptMappingFileListDto) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *JavaScriptMappingFileListDto) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *JavaScriptMappingFileListDto) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *JavaScriptMappingFileListDto) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *JavaScriptMappingFileListDto) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *JavaScriptMappingFileListDto) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *JavaScriptMappingFileListDto) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *JavaScriptMappingFileListDto) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *JavaScriptMappingFileListDto) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


