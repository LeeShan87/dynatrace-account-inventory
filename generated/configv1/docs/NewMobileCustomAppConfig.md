# NewMobileCustomAppConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApdexSettings** | Pointer to [**MobileCustomApdex**](MobileCustomApdex.md) |  | [optional] 
**ApplicationId** | Pointer to **string** | The UUID of the application.  It is used only by OneAgent to send data to Dynatrace. | [optional] 
**ApplicationType** | **string** | The type of the application. | 
**BeaconEndpointType** | Pointer to **string** | The type of the beacon endpoint. | [optional] 
**BeaconEndpointUrl** | Pointer to **string** | The URL of the beacon endpoint.  Only applicable when the **beaconEndpointType** is set to &#x60;ENVIRONMENT_ACTIVE_GATE&#x60; or &#x60;INSTRUMENTED_WEB_SERVER&#x60;. | [optional] 
**CostControlUserSessionPercentage** | Pointer to **int32** | The percentage of user sessions to be analyzed. | [optional] 
**IconType** | Pointer to **string** | Custom application icon.  Mobile apps always use the mobile device icon, so this icon can only be set for custom apps. | [optional] 
**Name** | **string** | The name of the application. | 
**OptInModeEnabled** | Pointer to **bool** | The opt-in mode is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).  This value is only applicable to mobile and not to custom apps. | [optional] 
**SessionReplayEnabled** | Pointer to **bool** | The session replay is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). This value is only applicable to mobile and not to custom apps. | [optional] 
**SessionReplayOnCrashEnabled** | Pointer to **bool** | **Deprecated**. Please use &#x60;sessionReplayEnabled&#x60; to enable or disable session replay for mobile apps. | [optional] 

## Methods

### NewNewMobileCustomAppConfig

`func NewNewMobileCustomAppConfig(applicationType string, name string, ) *NewMobileCustomAppConfig`

NewNewMobileCustomAppConfig instantiates a new NewMobileCustomAppConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMobileCustomAppConfigWithDefaults

`func NewNewMobileCustomAppConfigWithDefaults() *NewMobileCustomAppConfig`

NewNewMobileCustomAppConfigWithDefaults instantiates a new NewMobileCustomAppConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApdexSettings

`func (o *NewMobileCustomAppConfig) GetApdexSettings() MobileCustomApdex`

GetApdexSettings returns the ApdexSettings field if non-nil, zero value otherwise.

### GetApdexSettingsOk

`func (o *NewMobileCustomAppConfig) GetApdexSettingsOk() (*MobileCustomApdex, bool)`

GetApdexSettingsOk returns a tuple with the ApdexSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdexSettings

`func (o *NewMobileCustomAppConfig) SetApdexSettings(v MobileCustomApdex)`

SetApdexSettings sets ApdexSettings field to given value.

### HasApdexSettings

`func (o *NewMobileCustomAppConfig) HasApdexSettings() bool`

HasApdexSettings returns a boolean if a field has been set.

### GetApplicationId

`func (o *NewMobileCustomAppConfig) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *NewMobileCustomAppConfig) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *NewMobileCustomAppConfig) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *NewMobileCustomAppConfig) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationType

`func (o *NewMobileCustomAppConfig) GetApplicationType() string`

GetApplicationType returns the ApplicationType field if non-nil, zero value otherwise.

### GetApplicationTypeOk

`func (o *NewMobileCustomAppConfig) GetApplicationTypeOk() (*string, bool)`

GetApplicationTypeOk returns a tuple with the ApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationType

`func (o *NewMobileCustomAppConfig) SetApplicationType(v string)`

SetApplicationType sets ApplicationType field to given value.


### GetBeaconEndpointType

`func (o *NewMobileCustomAppConfig) GetBeaconEndpointType() string`

GetBeaconEndpointType returns the BeaconEndpointType field if non-nil, zero value otherwise.

### GetBeaconEndpointTypeOk

`func (o *NewMobileCustomAppConfig) GetBeaconEndpointTypeOk() (*string, bool)`

GetBeaconEndpointTypeOk returns a tuple with the BeaconEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconEndpointType

`func (o *NewMobileCustomAppConfig) SetBeaconEndpointType(v string)`

SetBeaconEndpointType sets BeaconEndpointType field to given value.

### HasBeaconEndpointType

`func (o *NewMobileCustomAppConfig) HasBeaconEndpointType() bool`

HasBeaconEndpointType returns a boolean if a field has been set.

### GetBeaconEndpointUrl

`func (o *NewMobileCustomAppConfig) GetBeaconEndpointUrl() string`

GetBeaconEndpointUrl returns the BeaconEndpointUrl field if non-nil, zero value otherwise.

### GetBeaconEndpointUrlOk

`func (o *NewMobileCustomAppConfig) GetBeaconEndpointUrlOk() (*string, bool)`

GetBeaconEndpointUrlOk returns a tuple with the BeaconEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconEndpointUrl

`func (o *NewMobileCustomAppConfig) SetBeaconEndpointUrl(v string)`

SetBeaconEndpointUrl sets BeaconEndpointUrl field to given value.

### HasBeaconEndpointUrl

`func (o *NewMobileCustomAppConfig) HasBeaconEndpointUrl() bool`

HasBeaconEndpointUrl returns a boolean if a field has been set.

### GetCostControlUserSessionPercentage

`func (o *NewMobileCustomAppConfig) GetCostControlUserSessionPercentage() int32`

GetCostControlUserSessionPercentage returns the CostControlUserSessionPercentage field if non-nil, zero value otherwise.

### GetCostControlUserSessionPercentageOk

`func (o *NewMobileCustomAppConfig) GetCostControlUserSessionPercentageOk() (*int32, bool)`

GetCostControlUserSessionPercentageOk returns a tuple with the CostControlUserSessionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostControlUserSessionPercentage

`func (o *NewMobileCustomAppConfig) SetCostControlUserSessionPercentage(v int32)`

SetCostControlUserSessionPercentage sets CostControlUserSessionPercentage field to given value.

### HasCostControlUserSessionPercentage

`func (o *NewMobileCustomAppConfig) HasCostControlUserSessionPercentage() bool`

HasCostControlUserSessionPercentage returns a boolean if a field has been set.

### GetIconType

`func (o *NewMobileCustomAppConfig) GetIconType() string`

GetIconType returns the IconType field if non-nil, zero value otherwise.

### GetIconTypeOk

`func (o *NewMobileCustomAppConfig) GetIconTypeOk() (*string, bool)`

GetIconTypeOk returns a tuple with the IconType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconType

`func (o *NewMobileCustomAppConfig) SetIconType(v string)`

SetIconType sets IconType field to given value.

### HasIconType

`func (o *NewMobileCustomAppConfig) HasIconType() bool`

HasIconType returns a boolean if a field has been set.

### GetName

`func (o *NewMobileCustomAppConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewMobileCustomAppConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewMobileCustomAppConfig) SetName(v string)`

SetName sets Name field to given value.


### GetOptInModeEnabled

`func (o *NewMobileCustomAppConfig) GetOptInModeEnabled() bool`

GetOptInModeEnabled returns the OptInModeEnabled field if non-nil, zero value otherwise.

### GetOptInModeEnabledOk

`func (o *NewMobileCustomAppConfig) GetOptInModeEnabledOk() (*bool, bool)`

GetOptInModeEnabledOk returns a tuple with the OptInModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInModeEnabled

`func (o *NewMobileCustomAppConfig) SetOptInModeEnabled(v bool)`

SetOptInModeEnabled sets OptInModeEnabled field to given value.

### HasOptInModeEnabled

`func (o *NewMobileCustomAppConfig) HasOptInModeEnabled() bool`

HasOptInModeEnabled returns a boolean if a field has been set.

### GetSessionReplayEnabled

`func (o *NewMobileCustomAppConfig) GetSessionReplayEnabled() bool`

GetSessionReplayEnabled returns the SessionReplayEnabled field if non-nil, zero value otherwise.

### GetSessionReplayEnabledOk

`func (o *NewMobileCustomAppConfig) GetSessionReplayEnabledOk() (*bool, bool)`

GetSessionReplayEnabledOk returns a tuple with the SessionReplayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayEnabled

`func (o *NewMobileCustomAppConfig) SetSessionReplayEnabled(v bool)`

SetSessionReplayEnabled sets SessionReplayEnabled field to given value.

### HasSessionReplayEnabled

`func (o *NewMobileCustomAppConfig) HasSessionReplayEnabled() bool`

HasSessionReplayEnabled returns a boolean if a field has been set.

### GetSessionReplayOnCrashEnabled

`func (o *NewMobileCustomAppConfig) GetSessionReplayOnCrashEnabled() bool`

GetSessionReplayOnCrashEnabled returns the SessionReplayOnCrashEnabled field if non-nil, zero value otherwise.

### GetSessionReplayOnCrashEnabledOk

`func (o *NewMobileCustomAppConfig) GetSessionReplayOnCrashEnabledOk() (*bool, bool)`

GetSessionReplayOnCrashEnabledOk returns a tuple with the SessionReplayOnCrashEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionReplayOnCrashEnabled

`func (o *NewMobileCustomAppConfig) SetSessionReplayOnCrashEnabled(v bool)`

SetSessionReplayOnCrashEnabled sets SessionReplayOnCrashEnabled field to given value.

### HasSessionReplayOnCrashEnabled

`func (o *NewMobileCustomAppConfig) HasSessionReplayOnCrashEnabled() bool`

HasSessionReplayOnCrashEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


