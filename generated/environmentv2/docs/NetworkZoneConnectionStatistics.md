# NetworkZoneConnectionStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostsConfiguredButNotConnected** | Pointer to **[]string** | Hosts from the network zone that use other zones. | [optional] [readonly] 
**HostsConnectedAsAlternative** | Pointer to **[]string** | Hosts that use the network zone as an alternative. | [optional] [readonly] 
**HostsConnectedAsFailover** | Pointer to **[]string** | Hosts from other zones that use the zone (not configured as an alternative) even though ActiveGates of higher priority are available. | [optional] [readonly] 
**HostsConnectedAsFailoverWithoutActiveGates** | Pointer to **[]string** | Hosts from other zones that use the zone (not configured as an alternative) and **no** ActiveGates of higher priority are available. | [optional] [readonly] 

## Methods

### NewNetworkZoneConnectionStatistics

`func NewNetworkZoneConnectionStatistics() *NetworkZoneConnectionStatistics`

NewNetworkZoneConnectionStatistics instantiates a new NetworkZoneConnectionStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkZoneConnectionStatisticsWithDefaults

`func NewNetworkZoneConnectionStatisticsWithDefaults() *NetworkZoneConnectionStatistics`

NewNetworkZoneConnectionStatisticsWithDefaults instantiates a new NetworkZoneConnectionStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostsConfiguredButNotConnected

`func (o *NetworkZoneConnectionStatistics) GetHostsConfiguredButNotConnected() []string`

GetHostsConfiguredButNotConnected returns the HostsConfiguredButNotConnected field if non-nil, zero value otherwise.

### GetHostsConfiguredButNotConnectedOk

`func (o *NetworkZoneConnectionStatistics) GetHostsConfiguredButNotConnectedOk() (*[]string, bool)`

GetHostsConfiguredButNotConnectedOk returns a tuple with the HostsConfiguredButNotConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsConfiguredButNotConnected

`func (o *NetworkZoneConnectionStatistics) SetHostsConfiguredButNotConnected(v []string)`

SetHostsConfiguredButNotConnected sets HostsConfiguredButNotConnected field to given value.

### HasHostsConfiguredButNotConnected

`func (o *NetworkZoneConnectionStatistics) HasHostsConfiguredButNotConnected() bool`

HasHostsConfiguredButNotConnected returns a boolean if a field has been set.

### GetHostsConnectedAsAlternative

`func (o *NetworkZoneConnectionStatistics) GetHostsConnectedAsAlternative() []string`

GetHostsConnectedAsAlternative returns the HostsConnectedAsAlternative field if non-nil, zero value otherwise.

### GetHostsConnectedAsAlternativeOk

`func (o *NetworkZoneConnectionStatistics) GetHostsConnectedAsAlternativeOk() (*[]string, bool)`

GetHostsConnectedAsAlternativeOk returns a tuple with the HostsConnectedAsAlternative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsConnectedAsAlternative

`func (o *NetworkZoneConnectionStatistics) SetHostsConnectedAsAlternative(v []string)`

SetHostsConnectedAsAlternative sets HostsConnectedAsAlternative field to given value.

### HasHostsConnectedAsAlternative

`func (o *NetworkZoneConnectionStatistics) HasHostsConnectedAsAlternative() bool`

HasHostsConnectedAsAlternative returns a boolean if a field has been set.

### GetHostsConnectedAsFailover

`func (o *NetworkZoneConnectionStatistics) GetHostsConnectedAsFailover() []string`

GetHostsConnectedAsFailover returns the HostsConnectedAsFailover field if non-nil, zero value otherwise.

### GetHostsConnectedAsFailoverOk

`func (o *NetworkZoneConnectionStatistics) GetHostsConnectedAsFailoverOk() (*[]string, bool)`

GetHostsConnectedAsFailoverOk returns a tuple with the HostsConnectedAsFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsConnectedAsFailover

`func (o *NetworkZoneConnectionStatistics) SetHostsConnectedAsFailover(v []string)`

SetHostsConnectedAsFailover sets HostsConnectedAsFailover field to given value.

### HasHostsConnectedAsFailover

`func (o *NetworkZoneConnectionStatistics) HasHostsConnectedAsFailover() bool`

HasHostsConnectedAsFailover returns a boolean if a field has been set.

### GetHostsConnectedAsFailoverWithoutActiveGates

`func (o *NetworkZoneConnectionStatistics) GetHostsConnectedAsFailoverWithoutActiveGates() []string`

GetHostsConnectedAsFailoverWithoutActiveGates returns the HostsConnectedAsFailoverWithoutActiveGates field if non-nil, zero value otherwise.

### GetHostsConnectedAsFailoverWithoutActiveGatesOk

`func (o *NetworkZoneConnectionStatistics) GetHostsConnectedAsFailoverWithoutActiveGatesOk() (*[]string, bool)`

GetHostsConnectedAsFailoverWithoutActiveGatesOk returns a tuple with the HostsConnectedAsFailoverWithoutActiveGates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsConnectedAsFailoverWithoutActiveGates

`func (o *NetworkZoneConnectionStatistics) SetHostsConnectedAsFailoverWithoutActiveGates(v []string)`

SetHostsConnectedAsFailoverWithoutActiveGates sets HostsConnectedAsFailoverWithoutActiveGates field to given value.

### HasHostsConnectedAsFailoverWithoutActiveGates

`func (o *NetworkZoneConnectionStatistics) HasHostsConnectedAsFailoverWithoutActiveGates() bool`

HasHostsConnectedAsFailoverWithoutActiveGates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


