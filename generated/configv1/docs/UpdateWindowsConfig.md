# UpdateWindowsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Windows** | [**[]UpdateWindow**](UpdateWindow.md) | List of maintenance windows when the OneAgent update can start. If there is no value and update should be performed, the update will start at earliest convenience. | 

## Methods

### NewUpdateWindowsConfig

`func NewUpdateWindowsConfig(windows []UpdateWindow, ) *UpdateWindowsConfig`

NewUpdateWindowsConfig instantiates a new UpdateWindowsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWindowsConfigWithDefaults

`func NewUpdateWindowsConfigWithDefaults() *UpdateWindowsConfig`

NewUpdateWindowsConfigWithDefaults instantiates a new UpdateWindowsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindows

`func (o *UpdateWindowsConfig) GetWindows() []UpdateWindow`

GetWindows returns the Windows field if non-nil, zero value otherwise.

### GetWindowsOk

`func (o *UpdateWindowsConfig) GetWindowsOk() (*[]UpdateWindow, bool)`

GetWindowsOk returns a tuple with the Windows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindows

`func (o *UpdateWindowsConfig) SetWindows(v []UpdateWindow)`

SetWindows sets Windows field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


