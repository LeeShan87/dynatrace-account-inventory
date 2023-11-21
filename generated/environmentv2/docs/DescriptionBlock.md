# DescriptionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]Image**](Image.md) | Collection of images (in case of gallery). | [optional] 
**Source** | Pointer to **NullableString** | Source of the description block (in case of markdown). | [optional] 
**SourceId** | Pointer to **NullableString** | Optional identifier of special description blocks. | [optional] 
**Title** | Pointer to **string** | Title of the description block. | [optional] 
**Type** | Pointer to **string** | Type of the data, either markdown or gallery. | [optional] 

## Methods

### NewDescriptionBlock

`func NewDescriptionBlock() *DescriptionBlock`

NewDescriptionBlock instantiates a new DescriptionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescriptionBlockWithDefaults

`func NewDescriptionBlockWithDefaults() *DescriptionBlock`

NewDescriptionBlockWithDefaults instantiates a new DescriptionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *DescriptionBlock) GetImages() []Image`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *DescriptionBlock) GetImagesOk() (*[]Image, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *DescriptionBlock) SetImages(v []Image)`

SetImages sets Images field to given value.

### HasImages

`func (o *DescriptionBlock) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *DescriptionBlock) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *DescriptionBlock) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetSource

`func (o *DescriptionBlock) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DescriptionBlock) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DescriptionBlock) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DescriptionBlock) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *DescriptionBlock) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *DescriptionBlock) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetSourceId

`func (o *DescriptionBlock) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DescriptionBlock) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DescriptionBlock) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *DescriptionBlock) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *DescriptionBlock) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *DescriptionBlock) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetTitle

`func (o *DescriptionBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DescriptionBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DescriptionBlock) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DescriptionBlock) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *DescriptionBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescriptionBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescriptionBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DescriptionBlock) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


