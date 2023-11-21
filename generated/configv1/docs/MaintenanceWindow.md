# MaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A short description of the maintenance purpose. | 
**Id** | Pointer to **string** | The ID of the maintenance window. | [optional] 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Name** | **string** | The name of the maintenance window, displayed in the UI. | 
**Schedule** | [**Schedule**](Schedule.md) |  | 
**Scope** | Pointer to [**Scope**](Scope.md) |  | [optional] 
**SuppressSyntheticMonitorsExecution** | Pointer to **bool** | Suppress execution of synthetic monitors during the maintenance. | [optional] 
**Suppression** | **string** | The type of suppression of alerting and problem detection during the maintenance. | 
**Type** | **string** | The type of the maintenance: planned or unplanned. | 

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow(description string, name string, schedule Schedule, suppression string, type_ string, ) *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MaintenanceWindow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceWindow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceWindow) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *MaintenanceWindow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceWindow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceWindow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceWindow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *MaintenanceWindow) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MaintenanceWindow) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MaintenanceWindow) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *MaintenanceWindow) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *MaintenanceWindow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MaintenanceWindow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MaintenanceWindow) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *MaintenanceWindow) GetSchedule() Schedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MaintenanceWindow) GetScheduleOk() (*Schedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MaintenanceWindow) SetSchedule(v Schedule)`

SetSchedule sets Schedule field to given value.


### GetScope

`func (o *MaintenanceWindow) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MaintenanceWindow) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MaintenanceWindow) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MaintenanceWindow) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSuppressSyntheticMonitorsExecution

`func (o *MaintenanceWindow) GetSuppressSyntheticMonitorsExecution() bool`

GetSuppressSyntheticMonitorsExecution returns the SuppressSyntheticMonitorsExecution field if non-nil, zero value otherwise.

### GetSuppressSyntheticMonitorsExecutionOk

`func (o *MaintenanceWindow) GetSuppressSyntheticMonitorsExecutionOk() (*bool, bool)`

GetSuppressSyntheticMonitorsExecutionOk returns a tuple with the SuppressSyntheticMonitorsExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressSyntheticMonitorsExecution

`func (o *MaintenanceWindow) SetSuppressSyntheticMonitorsExecution(v bool)`

SetSuppressSyntheticMonitorsExecution sets SuppressSyntheticMonitorsExecution field to given value.

### HasSuppressSyntheticMonitorsExecution

`func (o *MaintenanceWindow) HasSuppressSyntheticMonitorsExecution() bool`

HasSuppressSyntheticMonitorsExecution returns a boolean if a field has been set.

### GetSuppression

`func (o *MaintenanceWindow) GetSuppression() string`

GetSuppression returns the Suppression field if non-nil, zero value otherwise.

### GetSuppressionOk

`func (o *MaintenanceWindow) GetSuppressionOk() (*string, bool)`

GetSuppressionOk returns a tuple with the Suppression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppression

`func (o *MaintenanceWindow) SetSuppression(v string)`

SetSuppression sets Suppression field to given value.


### GetType

`func (o *MaintenanceWindow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MaintenanceWindow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MaintenanceWindow) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


