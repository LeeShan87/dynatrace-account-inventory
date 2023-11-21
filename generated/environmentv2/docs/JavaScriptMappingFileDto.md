# JavaScriptMappingFileDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | The name of the file. | [optional] 
**FileType** | Pointer to **string** | The type of the file. | [optional] 
**MinifiedJsFileUrl** | Pointer to **string** | The minified JavaScript file URL to which the mapping file belongs to. | [optional] 
**NumberOfFiles** | Pointer to **int32** | The number of files. | [optional] 
**Pinned** | Pointer to **bool** | Whether the file is pinned and therefore not automatically deleted. | [optional] 
**Size** | Pointer to **int32** | The size of the file, in KB. | [optional] 
**UploadTimestamp** | Pointer to **int64** | The timestamp of the file upload, in UTC milliseconds. | [optional] 
**Zipped** | Pointer to **bool** | Whether several files are zipped into one file. | [optional] 

## Methods

### NewJavaScriptMappingFileDto

`func NewJavaScriptMappingFileDto() *JavaScriptMappingFileDto`

NewJavaScriptMappingFileDto instantiates a new JavaScriptMappingFileDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJavaScriptMappingFileDtoWithDefaults

`func NewJavaScriptMappingFileDtoWithDefaults() *JavaScriptMappingFileDto`

NewJavaScriptMappingFileDtoWithDefaults instantiates a new JavaScriptMappingFileDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *JavaScriptMappingFileDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *JavaScriptMappingFileDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *JavaScriptMappingFileDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *JavaScriptMappingFileDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *JavaScriptMappingFileDto) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *JavaScriptMappingFileDto) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *JavaScriptMappingFileDto) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *JavaScriptMappingFileDto) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetMinifiedJsFileUrl

`func (o *JavaScriptMappingFileDto) GetMinifiedJsFileUrl() string`

GetMinifiedJsFileUrl returns the MinifiedJsFileUrl field if non-nil, zero value otherwise.

### GetMinifiedJsFileUrlOk

`func (o *JavaScriptMappingFileDto) GetMinifiedJsFileUrlOk() (*string, bool)`

GetMinifiedJsFileUrlOk returns a tuple with the MinifiedJsFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinifiedJsFileUrl

`func (o *JavaScriptMappingFileDto) SetMinifiedJsFileUrl(v string)`

SetMinifiedJsFileUrl sets MinifiedJsFileUrl field to given value.

### HasMinifiedJsFileUrl

`func (o *JavaScriptMappingFileDto) HasMinifiedJsFileUrl() bool`

HasMinifiedJsFileUrl returns a boolean if a field has been set.

### GetNumberOfFiles

`func (o *JavaScriptMappingFileDto) GetNumberOfFiles() int32`

GetNumberOfFiles returns the NumberOfFiles field if non-nil, zero value otherwise.

### GetNumberOfFilesOk

`func (o *JavaScriptMappingFileDto) GetNumberOfFilesOk() (*int32, bool)`

GetNumberOfFilesOk returns a tuple with the NumberOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFiles

`func (o *JavaScriptMappingFileDto) SetNumberOfFiles(v int32)`

SetNumberOfFiles sets NumberOfFiles field to given value.

### HasNumberOfFiles

`func (o *JavaScriptMappingFileDto) HasNumberOfFiles() bool`

HasNumberOfFiles returns a boolean if a field has been set.

### GetPinned

`func (o *JavaScriptMappingFileDto) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *JavaScriptMappingFileDto) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *JavaScriptMappingFileDto) SetPinned(v bool)`

SetPinned sets Pinned field to given value.

### HasPinned

`func (o *JavaScriptMappingFileDto) HasPinned() bool`

HasPinned returns a boolean if a field has been set.

### GetSize

`func (o *JavaScriptMappingFileDto) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *JavaScriptMappingFileDto) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *JavaScriptMappingFileDto) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *JavaScriptMappingFileDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetUploadTimestamp

`func (o *JavaScriptMappingFileDto) GetUploadTimestamp() int64`

GetUploadTimestamp returns the UploadTimestamp field if non-nil, zero value otherwise.

### GetUploadTimestampOk

`func (o *JavaScriptMappingFileDto) GetUploadTimestampOk() (*int64, bool)`

GetUploadTimestampOk returns a tuple with the UploadTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTimestamp

`func (o *JavaScriptMappingFileDto) SetUploadTimestamp(v int64)`

SetUploadTimestamp sets UploadTimestamp field to given value.

### HasUploadTimestamp

`func (o *JavaScriptMappingFileDto) HasUploadTimestamp() bool`

HasUploadTimestamp returns a boolean if a field has been set.

### GetZipped

`func (o *JavaScriptMappingFileDto) GetZipped() bool`

GetZipped returns the Zipped field if non-nil, zero value otherwise.

### GetZippedOk

`func (o *JavaScriptMappingFileDto) GetZippedOk() (*bool, bool)`

GetZippedOk returns a tuple with the Zipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipped

`func (o *JavaScriptMappingFileDto) SetZipped(v bool)`

SetZipped sets Zipped field to given value.

### HasZipped

`func (o *JavaScriptMappingFileDto) HasZipped() bool`

HasZipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


