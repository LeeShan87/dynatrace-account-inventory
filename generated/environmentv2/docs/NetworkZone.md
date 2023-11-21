# NetworkZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternativeZones** | Pointer to **[]string** | A list of alternative network zones. | [optional] 
**Description** | Pointer to **string** | A short description of the network zone. | [optional] 
**FallbackMode** | Pointer to **string** | The fallback mode of the network zone. | [optional] 
**Id** | Pointer to **string** | The ID of the network zone. | [optional] [readonly] 
**NumOfConfiguredActiveGates** | Pointer to **int32** | The number of ActiveGates in the network zone. | [optional] [readonly] 
**NumOfConfiguredOneAgents** | Pointer to **int32** | The number of OneAgents that are configured to use the network zone as primary. | [optional] [readonly] 
**NumOfOneAgentsFromOtherZones** | Pointer to **int32** | The number of OneAgents from other network zones that are using ActiveGates in the network zone.    This is a fraction of **numOfOneAgentsUsing**.   One possible reason for switching to another zone is that a firewall is preventing a OneAgent from connecting to any ActiveGate in the preferred network zone.   | [optional] [readonly] 
**NumOfOneAgentsUsing** | Pointer to **int32** | The number of OneAgents that are using ActiveGates in the network zone. | [optional] [readonly] 
**OverridesGlobal** | Pointer to **bool** | Indicates if a global network zone is overridden (managed only). | [optional] [readonly] 
**Scope** | Pointer to **string** | Specifies the scope of the network zone (managed only). | [optional] [readonly] 

## Methods

### NewNetworkZone

`func NewNetworkZone() *NetworkZone`

NewNetworkZone instantiates a new NetworkZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkZoneWithDefaults

`func NewNetworkZoneWithDefaults() *NetworkZone`

NewNetworkZoneWithDefaults instantiates a new NetworkZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternativeZones

`func (o *NetworkZone) GetAlternativeZones() []string`

GetAlternativeZones returns the AlternativeZones field if non-nil, zero value otherwise.

### GetAlternativeZonesOk

`func (o *NetworkZone) GetAlternativeZonesOk() (*[]string, bool)`

GetAlternativeZonesOk returns a tuple with the AlternativeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeZones

`func (o *NetworkZone) SetAlternativeZones(v []string)`

SetAlternativeZones sets AlternativeZones field to given value.

### HasAlternativeZones

`func (o *NetworkZone) HasAlternativeZones() bool`

HasAlternativeZones returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFallbackMode

`func (o *NetworkZone) GetFallbackMode() string`

GetFallbackMode returns the FallbackMode field if non-nil, zero value otherwise.

### GetFallbackModeOk

`func (o *NetworkZone) GetFallbackModeOk() (*string, bool)`

GetFallbackModeOk returns a tuple with the FallbackMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMode

`func (o *NetworkZone) SetFallbackMode(v string)`

SetFallbackMode sets FallbackMode field to given value.

### HasFallbackMode

`func (o *NetworkZone) HasFallbackMode() bool`

HasFallbackMode returns a boolean if a field has been set.

### GetId

`func (o *NetworkZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumOfConfiguredActiveGates

`func (o *NetworkZone) GetNumOfConfiguredActiveGates() int32`

GetNumOfConfiguredActiveGates returns the NumOfConfiguredActiveGates field if non-nil, zero value otherwise.

### GetNumOfConfiguredActiveGatesOk

`func (o *NetworkZone) GetNumOfConfiguredActiveGatesOk() (*int32, bool)`

GetNumOfConfiguredActiveGatesOk returns a tuple with the NumOfConfiguredActiveGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfConfiguredActiveGates

`func (o *NetworkZone) SetNumOfConfiguredActiveGates(v int32)`

SetNumOfConfiguredActiveGates sets NumOfConfiguredActiveGates field to given value.

### HasNumOfConfiguredActiveGates

`func (o *NetworkZone) HasNumOfConfiguredActiveGates() bool`

HasNumOfConfiguredActiveGates returns a boolean if a field has been set.

### GetNumOfConfiguredOneAgents

`func (o *NetworkZone) GetNumOfConfiguredOneAgents() int32`

GetNumOfConfiguredOneAgents returns the NumOfConfiguredOneAgents field if non-nil, zero value otherwise.

### GetNumOfConfiguredOneAgentsOk

`func (o *NetworkZone) GetNumOfConfiguredOneAgentsOk() (*int32, bool)`

GetNumOfConfiguredOneAgentsOk returns a tuple with the NumOfConfiguredOneAgents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfConfiguredOneAgents

`func (o *NetworkZone) SetNumOfConfiguredOneAgents(v int32)`

SetNumOfConfiguredOneAgents sets NumOfConfiguredOneAgents field to given value.

### HasNumOfConfiguredOneAgents

`func (o *NetworkZone) HasNumOfConfiguredOneAgents() bool`

HasNumOfConfiguredOneAgents returns a boolean if a field has been set.

### GetNumOfOneAgentsFromOtherZones

`func (o *NetworkZone) GetNumOfOneAgentsFromOtherZones() int32`

GetNumOfOneAgentsFromOtherZones returns the NumOfOneAgentsFromOtherZones field if non-nil, zero value otherwise.

### GetNumOfOneAgentsFromOtherZonesOk

`func (o *NetworkZone) GetNumOfOneAgentsFromOtherZonesOk() (*int32, bool)`

GetNumOfOneAgentsFromOtherZonesOk returns a tuple with the NumOfOneAgentsFromOtherZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfOneAgentsFromOtherZones

`func (o *NetworkZone) SetNumOfOneAgentsFromOtherZones(v int32)`

SetNumOfOneAgentsFromOtherZones sets NumOfOneAgentsFromOtherZones field to given value.

### HasNumOfOneAgentsFromOtherZones

`func (o *NetworkZone) HasNumOfOneAgentsFromOtherZones() bool`

HasNumOfOneAgentsFromOtherZones returns a boolean if a field has been set.

### GetNumOfOneAgentsUsing

`func (o *NetworkZone) GetNumOfOneAgentsUsing() int32`

GetNumOfOneAgentsUsing returns the NumOfOneAgentsUsing field if non-nil, zero value otherwise.

### GetNumOfOneAgentsUsingOk

`func (o *NetworkZone) GetNumOfOneAgentsUsingOk() (*int32, bool)`

GetNumOfOneAgentsUsingOk returns a tuple with the NumOfOneAgentsUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfOneAgentsUsing

`func (o *NetworkZone) SetNumOfOneAgentsUsing(v int32)`

SetNumOfOneAgentsUsing sets NumOfOneAgentsUsing field to given value.

### HasNumOfOneAgentsUsing

`func (o *NetworkZone) HasNumOfOneAgentsUsing() bool`

HasNumOfOneAgentsUsing returns a boolean if a field has been set.

### GetOverridesGlobal

`func (o *NetworkZone) GetOverridesGlobal() bool`

GetOverridesGlobal returns the OverridesGlobal field if non-nil, zero value otherwise.

### GetOverridesGlobalOk

`func (o *NetworkZone) GetOverridesGlobalOk() (*bool, bool)`

GetOverridesGlobalOk returns a tuple with the OverridesGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridesGlobal

`func (o *NetworkZone) SetOverridesGlobal(v bool)`

SetOverridesGlobal sets OverridesGlobal field to given value.

### HasOverridesGlobal

`func (o *NetworkZone) HasOverridesGlobal() bool`

HasOverridesGlobal returns a boolean if a field has been set.

### GetScope

`func (o *NetworkZone) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NetworkZone) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NetworkZone) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *NetworkZone) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


