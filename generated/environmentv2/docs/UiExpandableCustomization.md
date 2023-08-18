# UiExpandableCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name | [optional] 
**Expanded** | Pointer to **bool** | Defines if the item should be expanded by default | [optional] 
**Sections** | Pointer to [**[]UiExpandableSectionCustomization**](UiExpandableSectionCustomization.md) | A list of sections | [optional] 

## Methods

### NewUiExpandableCustomization

`func NewUiExpandableCustomization() *UiExpandableCustomization`

NewUiExpandableCustomization instantiates a new UiExpandableCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiExpandableCustomizationWithDefaults

`func NewUiExpandableCustomizationWithDefaults() *UiExpandableCustomization`

NewUiExpandableCustomizationWithDefaults instantiates a new UiExpandableCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UiExpandableCustomization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UiExpandableCustomization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UiExpandableCustomization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UiExpandableCustomization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpanded

`func (o *UiExpandableCustomization) GetExpanded() bool`

GetExpanded returns the Expanded field if non-nil, zero value otherwise.

### GetExpandedOk

`func (o *UiExpandableCustomization) GetExpandedOk() (*bool, bool)`

GetExpandedOk returns a tuple with the Expanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpanded

`func (o *UiExpandableCustomization) SetExpanded(v bool)`

SetExpanded sets Expanded field to given value.

### HasExpanded

`func (o *UiExpandableCustomization) HasExpanded() bool`

HasExpanded returns a boolean if a field has been set.

### GetSections

`func (o *UiExpandableCustomization) GetSections() []UiExpandableSectionCustomization`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *UiExpandableCustomization) GetSectionsOk() (*[]UiExpandableSectionCustomization, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *UiExpandableCustomization) SetSections(v []UiExpandableSectionCustomization)`

SetSections sets Sections field to given value.

### HasSections

`func (o *UiExpandableCustomization) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


