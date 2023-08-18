# UiExpandableSectionCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description | [optional] 
**DisplayName** | **string** | The display name | 
**Expanded** | Pointer to **bool** | Defines if the section should be expanded by default | [optional] 
**Properties** | **[]string** | A list of properties | 

## Methods

### NewUiExpandableSectionCustomization

`func NewUiExpandableSectionCustomization(displayName string, properties []string, ) *UiExpandableSectionCustomization`

NewUiExpandableSectionCustomization instantiates a new UiExpandableSectionCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiExpandableSectionCustomizationWithDefaults

`func NewUiExpandableSectionCustomizationWithDefaults() *UiExpandableSectionCustomization`

NewUiExpandableSectionCustomizationWithDefaults instantiates a new UiExpandableSectionCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UiExpandableSectionCustomization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UiExpandableSectionCustomization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UiExpandableSectionCustomization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UiExpandableSectionCustomization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UiExpandableSectionCustomization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UiExpandableSectionCustomization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UiExpandableSectionCustomization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpanded

`func (o *UiExpandableSectionCustomization) GetExpanded() bool`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *UiExpandableSectionCustomization) GetExpandedOk() (*bool, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *UiExpandableSectionCustomization) SetExpanded(v bool)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *UiExpandableSectionCustomization) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### GetProperties

`func (o *UiExpandableSectionCustomization) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UiExpandableSectionCustomization) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UiExpandableSectionCustomization) SetProperties(v []string)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


