# SectionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The property key. | [optional] 
**Section** | Pointer to **string** | The section this property belongs to. | [optional] 
**Value** | Pointer to **string** | The property value. | [optional] 

## Methods

### NewSectionProperty

`func NewSectionProperty() *SectionProperty`

NewSectionProperty instantiates a new SectionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionPropertyWithDefaults

`func NewSectionPropertyWithDefaults() *SectionProperty`

NewSectionPropertyWithDefaults instantiates a new SectionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SectionProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SectionProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SectionProperty) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SectionProperty) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSection

`func (o *SectionProperty) GetSection() string`

GetSection returns the Section field if non-nil, zero value otherwise.

### GetSectionOk

`func (o *SectionProperty) GetSectionOk() (*string, bool)`

GetSectionOk returns a tuple with the Section field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSection

`func (o *SectionProperty) SetSection(v string)`

SetSection sets Section field to given value.

### HasSection

`func (o *SectionProperty) HasSection() bool`

HasSection returns a boolean if a field has been set.

### GetValue

`func (o *SectionProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SectionProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SectionProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SectionProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


