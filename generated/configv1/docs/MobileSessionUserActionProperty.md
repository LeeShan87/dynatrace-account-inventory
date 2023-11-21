# MobileSessionUserActionProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregation** | Pointer to **string** | The aggregation type of the property.     It defines how multiple values of the property are aggregated. | [optional] 
**CleanupRule** | Pointer to **string** | The cleanup rule of the property.   Defines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 
**Key** | **string** | The unique key of the mobile session or user action property. | 
**Name** | Pointer to **string** | The name of the reported value.   Only applicable when the **origin** is set to &#x60;API&#x60;. | [optional] 
**Origin** | **string** | The origin of the property | 
**ServerSideRequestAttribute** | Pointer to **string** | The ID of the request attribute.   Only applicable when the **origin** is set to &#x60;SERVER_SIDE_REQUEST_ATTRIBUTE&#x60;. | [optional] 
**StoreAsSessionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a session property | [optional] 
**StoreAsUserActionProperty** | Pointer to **bool** | If &#x60;true&#x60;, the property is stored as a user action property | [optional] 
**Type** | **string** | The data type of the property. | 

## Methods

### NewMobileSessionUserActionProperty

`func NewMobileSessionUserActionProperty(key string, origin string, type_ string, ) *MobileSessionUserActionProperty`

NewMobileSessionUserActionProperty instantiates a new MobileSessionUserActionProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileSessionUserActionPropertyWithDefaults

`func NewMobileSessionUserActionPropertyWithDefaults() *MobileSessionUserActionProperty`

NewMobileSessionUserActionPropertyWithDefaults instantiates a new MobileSessionUserActionProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregation

`func (o *MobileSessionUserActionProperty) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *MobileSessionUserActionProperty) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *MobileSessionUserActionProperty) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *MobileSessionUserActionProperty) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetCleanupRule

`func (o *MobileSessionUserActionProperty) GetCleanupRule() string`

GetCleanupRule returns the CleanupRule field if non-nil, zero value otherwise.

### GetCleanupRuleOk

`func (o *MobileSessionUserActionProperty) GetCleanupRuleOk() (*string, bool)`

GetCleanupRuleOk returns a tuple with the CleanupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupRule

`func (o *MobileSessionUserActionProperty) SetCleanupRule(v string)`

SetCleanupRule sets CleanupRule field to given value.

### HasCleanupRule

`func (o *MobileSessionUserActionProperty) HasCleanupRule() bool`

HasCleanupRule returns a boolean if a field has been set.

### GetDisplayName

`func (o *MobileSessionUserActionProperty) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MobileSessionUserActionProperty) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MobileSessionUserActionProperty) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MobileSessionUserActionProperty) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *MobileSessionUserActionProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MobileSessionUserActionProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MobileSessionUserActionProperty) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *MobileSessionUserActionProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MobileSessionUserActionProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MobileSessionUserActionProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MobileSessionUserActionProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *MobileSessionUserActionProperty) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MobileSessionUserActionProperty) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MobileSessionUserActionProperty) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetServerSideRequestAttribute

`func (o *MobileSessionUserActionProperty) GetServerSideRequestAttribute() string`

GetServerSideRequestAttribute returns the ServerSideRequestAttribute field if non-nil, zero value otherwise.

### GetServerSideRequestAttributeOk

`func (o *MobileSessionUserActionProperty) GetServerSideRequestAttributeOk() (*string, bool)`

GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSideRequestAttribute

`func (o *MobileSessionUserActionProperty) SetServerSideRequestAttribute(v string)`

SetServerSideRequestAttribute sets ServerSideRequestAttribute field to given value.

### HasServerSideRequestAttribute

`func (o *MobileSessionUserActionProperty) HasServerSideRequestAttribute() bool`

HasServerSideRequestAttribute returns a boolean if a field has been set.

### GetStoreAsSessionProperty

`func (o *MobileSessionUserActionProperty) GetStoreAsSessionProperty() bool`

GetStoreAsSessionProperty returns the StoreAsSessionProperty field if non-nil, zero value otherwise.

### GetStoreAsSessionPropertyOk

`func (o *MobileSessionUserActionProperty) GetStoreAsSessionPropertyOk() (*bool, bool)`

GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsSessionProperty

`func (o *MobileSessionUserActionProperty) SetStoreAsSessionProperty(v bool)`

SetStoreAsSessionProperty sets StoreAsSessionProperty field to given value.

### HasStoreAsSessionProperty

`func (o *MobileSessionUserActionProperty) HasStoreAsSessionProperty() bool`

HasStoreAsSessionProperty returns a boolean if a field has been set.

### GetStoreAsUserActionProperty

`func (o *MobileSessionUserActionProperty) GetStoreAsUserActionProperty() bool`

GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field if non-nil, zero value otherwise.

### GetStoreAsUserActionPropertyOk

`func (o *MobileSessionUserActionProperty) GetStoreAsUserActionPropertyOk() (*bool, bool)`

GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAsUserActionProperty

`func (o *MobileSessionUserActionProperty) SetStoreAsUserActionProperty(v bool)`

SetStoreAsUserActionProperty sets StoreAsUserActionProperty field to given value.

### HasStoreAsUserActionProperty

`func (o *MobileSessionUserActionProperty) HasStoreAsUserActionProperty() bool`

HasStoreAsUserActionProperty returns a boolean if a field has been set.

### GetType

`func (o *MobileSessionUserActionProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MobileSessionUserActionProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MobileSessionUserActionProperty) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


