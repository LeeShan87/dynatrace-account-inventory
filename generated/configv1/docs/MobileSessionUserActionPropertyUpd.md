# MobileSessionUserActionPropertyUpd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** | The aggregation type of the property.     It defines how multiple values of the property are aggregated. | [optional] 
**CleanupRule** | Pointer to **string** | The cleanup rule of the property.   Defines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 
**Name** | Pointer to **string** | The name of the reported value.   Only applicable when the **origin** is set to &#x60;API&#x60;. | [optional] 
**Origin** | **string** | The origin of the property | 
**ServerSideRequestAttribute** | Pointer to **string** | The ID of the request attribute.   Only applicable when the **origin** is set to &#x60;SERVER_SIDE_REQUEST_ATTRIBUTE&#x60;. | [optional] 
**StoreAsSessionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a session property | [optional] 
**StoreAsUserActionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a user action property | [optional] 
**Type** | **string** | The data type of the property. | 

## Methods

### NewMobileSessionUserActionPropertyUpd

`func NewMobileSessionUserActionPropertyUpd(origin string, type_ string, ) *MobileSessionUserActionPropertyUpd`

NewMobileSessionUserActionPropertyUpd instantiates a new MobileSessionUserActionPropertyUpd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileSessionUserActionPropertyUpdWithDefaults

`func NewMobileSessionUserActionPropertyUpdWithDefaults() *MobileSessionUserActionPropertyUpd`

NewMobileSessionUserActionPropertyUpdWithDefaults instantiates a new MobileSessionUserActionPropertyUpd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MobileSessionUserActionPropertyUpd) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MobileSessionUserActionPropertyUpd) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MobileSessionUserActionPropertyUpd) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MobileSessionUserActionPropertyUpd) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetCleanupRule

`func (o *MobileSessionUserActionPropertyUpd) GetCleanupRule() string`

GetCleanupRule returns the CleanupRule field if non-nil, zero value otherwise.

### GetCleanupRuleOk

`func (o *MobileSessionUserActionPropertyUpd) GetCleanupRuleOk() (*string, bool)`

GetCleanupRuleOk returns a tuple with the CleanupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRule

`func (o *MobileSessionUserActionPropertyUpd) SetCleanupRule(v string)`

SetCleanupRule sets CleanupRule field to given value.

### HasCleanupRule

`func (o *MobileSessionUserActionPropertyUpd) HasCleanupRule() bool`

HasCleanupRule returns a boolean if a field has been set.

### GetDisplayName

`func (o *MobileSessionUserActionPropertyUpd) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileSessionUserActionPropertyUpd) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileSessionUserActionPropertyUpd) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileSessionUserActionPropertyUpd) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *MobileSessionUserActionPropertyUpd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileSessionUserActionPropertyUpd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileSessionUserActionPropertyUpd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileSessionUserActionPropertyUpd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *MobileSessionUserActionPropertyUpd) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MobileSessionUserActionPropertyUpd) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MobileSessionUserActionPropertyUpd) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetServerSideRequestAttribute

`func (o *MobileSessionUserActionPropertyUpd) GetServerSideRequestAttribute() string`

GetServerSideRequestAttribute returns the ServerSideRequestAttribute field if non-nil, zero value otherwise.

### GetServerSideRequestAttributeOk

`func (o *MobileSessionUserActionPropertyUpd) GetServerSideRequestAttributeOk() (*string, bool)`

GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideRequestAttribute

`func (o *MobileSessionUserActionPropertyUpd) SetServerSideRequestAttribute(v string)`

SetServerSideRequestAttribute sets ServerSideRequestAttribute field to given value.

### HasServerSideRequestAttribute

`func (o *MobileSessionUserActionPropertyUpd) HasServerSideRequestAttribute() bool`

HasServerSideRequestAttribute returns a boolean if a field has been set.

### GetStoreAsSessionProperty

`func (o *MobileSessionUserActionPropertyUpd) GetStoreAsSessionProperty() bool`

GetStoreAsSessionProperty returns the StoreAsSessionProperty field if non-nil, zero value otherwise.

### GetStoreAsSessionPropertyOk

`func (o *MobileSessionUserActionPropertyUpd) GetStoreAsSessionPropertyOk() (*bool, bool)`

GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsSessionProperty

`func (o *MobileSessionUserActionPropertyUpd) SetStoreAsSessionProperty(v bool)`

SetStoreAsSessionProperty sets StoreAsSessionProperty field to given value.

### HasStoreAsSessionProperty

`func (o *MobileSessionUserActionPropertyUpd) HasStoreAsSessionProperty() bool`

HasStoreAsSessionProperty returns a boolean if a field has been set.

### GetStoreAsUserActionProperty

`func (o *MobileSessionUserActionPropertyUpd) GetStoreAsUserActionProperty() bool`

GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field if non-nil, zero value otherwise.

### GetStoreAsUserActionPropertyOk

`func (o *MobileSessionUserActionPropertyUpd) GetStoreAsUserActionPropertyOk() (*bool, bool)`

GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsUserActionProperty

`func (o *MobileSessionUserActionPropertyUpd) SetStoreAsUserActionProperty(v bool)`

SetStoreAsUserActionProperty sets StoreAsUserActionProperty field to given value.

### HasStoreAsUserActionProperty

`func (o *MobileSessionUserActionPropertyUpd) HasStoreAsUserActionProperty() bool`

HasStoreAsUserActionProperty returns a boolean if a field has been set.

### GetType

`func (o *MobileSessionUserActionPropertyUpd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileSessionUserActionPropertyUpd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileSessionUserActionPropertyUpd) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


