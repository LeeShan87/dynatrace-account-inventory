# ManualApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** | The Dynatrace entity ID of the application. | [optional] 
**DisplayName** | Pointer to **string** | The name of the application. | [optional] 
**MonitoringEnabled** | Pointer to **bool** | Monitoring is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**Revision** | Pointer to **string** | The application settings revision. | [optional] 

## Methods

### NewManualApplication

`func NewManualApplication() *ManualApplication`

NewManualApplication instantiates a new ManualApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualApplicationWithDefaults

`func NewManualApplicationWithDefaults() *ManualApplication`

NewManualApplicationWithDefaults instantiates a new ManualApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ManualApplication) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ManualApplication) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ManualApplication) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ManualApplication) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ManualApplication) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManualApplication) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManualApplication) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManualApplication) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMonitoringEnabled

`func (o *ManualApplication) GetMonitoringEnabled() bool`

GetMonitoringEnabled returns the MonitoringEnabled field if non-nil, zero value otherwise.

### GetMonitoringEnabledOk

`func (o *ManualApplication) GetMonitoringEnabledOk() (*bool, bool)`

GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringEnabled

`func (o *ManualApplication) SetMonitoringEnabled(v bool)`

SetMonitoringEnabled sets MonitoringEnabled field to given value.

### HasMonitoringEnabled

`func (o *ManualApplication) HasMonitoringEnabled() bool`

HasMonitoringEnabled returns a boolean if a field has been set.

### GetRevision

`func (o *ManualApplication) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ManualApplication) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ManualApplication) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ManualApplication) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


