# MobileSessionUserActionPropertyList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionProperties** | Pointer to [**[]MobileSessionUserActionPropertyShort**](MobileSessionUserActionPropertyShort.md) | A list of short representations of mobile session properties. | [optional] 
**UserActionProperties** | Pointer to [**[]MobileSessionUserActionPropertyShort**](MobileSessionUserActionPropertyShort.md) | A list of short representations of mobile user action properties. | [optional] 

## Methods

### NewMobileSessionUserActionPropertyList

`func NewMobileSessionUserActionPropertyList() *MobileSessionUserActionPropertyList`

NewMobileSessionUserActionPropertyList instantiates a new MobileSessionUserActionPropertyList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileSessionUserActionPropertyListWithDefaults

`func NewMobileSessionUserActionPropertyListWithDefaults() *MobileSessionUserActionPropertyList`

NewMobileSessionUserActionPropertyListWithDefaults instantiates a new MobileSessionUserActionPropertyList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionProperties

`func (o *MobileSessionUserActionPropertyList) GetSessionProperties() []MobileSessionUserActionPropertyShort`

GetSessionProperties returns the SessionProperties field if non-nil, zero value otherwise.

### GetSessionPropertiesOk

`func (o *MobileSessionUserActionPropertyList) GetSessionPropertiesOk() (*[]MobileSessionUserActionPropertyShort, bool)`

GetSessionPropertiesOk returns a tuple with the SessionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProperties

`func (o *MobileSessionUserActionPropertyList) SetSessionProperties(v []MobileSessionUserActionPropertyShort)`

SetSessionProperties sets SessionProperties field to given value.

### HasSessionProperties

`func (o *MobileSessionUserActionPropertyList) HasSessionProperties() bool`

HasSessionProperties returns a boolean if a field has been set.

### GetUserActionProperties

`func (o *MobileSessionUserActionPropertyList) GetUserActionProperties() []MobileSessionUserActionPropertyShort`

GetUserActionProperties returns the UserActionProperties field if non-nil, zero value otherwise.

### GetUserActionPropertiesOk

`func (o *MobileSessionUserActionPropertyList) GetUserActionPropertiesOk() (*[]MobileSessionUserActionPropertyShort, bool)`

GetUserActionPropertiesOk returns a tuple with the UserActionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionProperties

`func (o *MobileSessionUserActionPropertyList) SetUserActionProperties(v []MobileSessionUserActionPropertyShort)`

SetUserActionProperties sets UserActionProperties field to given value.

### HasUserActionProperties

`func (o *MobileSessionUserActionPropertyList) HasUserActionProperties() bool`

HasUserActionProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


