# UiButtonCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description to be shown in a tooltip when hovering over the button | [optional] 
**DisplayName** | **string** | The label of the button | 
**Identifier** | **string** | The identifier of the function to be called when the button is pressed | 
**Insert** | [**UiButtonCustomizationInsert**](UiButtonCustomizationInsert.md) |  | 

## Methods

### NewUiButtonCustomization

`func NewUiButtonCustomization(displayName string, identifier string, insert UiButtonCustomizationInsert, ) *UiButtonCustomization`

NewUiButtonCustomization instantiates a new UiButtonCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiButtonCustomizationWithDefaults

`func NewUiButtonCustomizationWithDefaults() *UiButtonCustomization`

NewUiButtonCustomizationWithDefaults instantiates a new UiButtonCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UiButtonCustomization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UiButtonCustomization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UiButtonCustomization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UiButtonCustomization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UiButtonCustomization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UiButtonCustomization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UiButtonCustomization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIdentifier

`func (o *UiButtonCustomization) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *UiButtonCustomization) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *UiButtonCustomization) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetInsert

`func (o *UiButtonCustomization) GetInsert() UiButtonCustomizationInsert`

GetInsert returns the Insert field if non-nil, zero value otherwise.

### GetInsertOk

`func (o *UiButtonCustomization) GetInsertOk() (*UiButtonCustomizationInsert, bool)`

GetInsertOk returns a tuple with the Insert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsert

`func (o *UiButtonCustomization) SetInsert(v UiButtonCustomizationInsert)`

SetInsert sets Insert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


