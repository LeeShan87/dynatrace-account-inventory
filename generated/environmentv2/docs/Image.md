# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alt** | Pointer to **string** | Alternate text for the image. | [optional] 
**Src** | Pointer to **string** | Url of the image. | [optional] 
**Title** | Pointer to **string** | Title of the image. | [optional] 

## Methods

### NewImage

`func NewImage() *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlt

`func (o *Image) GetAlt() string`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *Image) GetAltOk() (*string, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *Image) SetAlt(v string)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *Image) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetSrc

`func (o *Image) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *Image) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *Image) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *Image) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetTitle

`func (o *Image) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Image) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Image) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Image) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


