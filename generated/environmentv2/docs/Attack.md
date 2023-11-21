# Attack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedEntities** | Pointer to [**AffectedEntities**](AffectedEntities.md) |  | [optional] 
**AttackId** | Pointer to **string** | The ID of the attack. | [optional] [readonly] 
**AttackTarget** | Pointer to [**AttackTarget**](AttackTarget.md) |  | [optional] 
**AttackType** | Pointer to **string** | The type of the attack. | [optional] [readonly] 
**Attacker** | Pointer to [**Attacker**](Attacker.md) |  | [optional] 
**DisplayId** | Pointer to **string** | The display ID of the attack. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the attack. | [optional] [readonly] 
**Entrypoint** | Pointer to [**AttackEntrypoint**](AttackEntrypoint.md) |  | [optional] 
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | A list of management zones which the affected entities belong to. | [optional] [readonly] 
**Request** | Pointer to [**RequestInformation**](RequestInformation.md) |  | [optional] 
**SecurityProblem** | Pointer to [**AttackSecurityProblem**](AttackSecurityProblem.md) |  | [optional] 
**State** | Pointer to **string** | The state of the attack. | [optional] [readonly] 
**Technology** | Pointer to **string** | The technology of the attack. | [optional] [readonly] 
**Timestamp** | Pointer to **int64** | The timestamp when the attack occurred. | [optional] [readonly] 
**Vulnerability** | Pointer to [**Vulnerability**](Vulnerability.md) |  | [optional] 

## Methods

### NewAttack

`func NewAttack() *Attack`

NewAttack instantiates a new Attack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackWithDefaults

`func NewAttackWithDefaults() *Attack`

NewAttackWithDefaults instantiates a new Attack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedEntities

`func (o *Attack) GetAffectedEntities() AffectedEntities`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *Attack) GetAffectedEntitiesOk() (*AffectedEntities, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *Attack) SetAffectedEntities(v AffectedEntities)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *Attack) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.

### GetAttackId

`func (o *Attack) GetAttackId() string`

GetAttackId returns the AttackId field if non-nil, zero value otherwise.

### GetAttackIdOk

`func (o *Attack) GetAttackIdOk() (*string, bool)`

GetAttackIdOk returns a tuple with the AttackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackId

`func (o *Attack) SetAttackId(v string)`

SetAttackId sets AttackId field to given value.

### HasAttackId

`func (o *Attack) HasAttackId() bool`

HasAttackId returns a boolean if a field has been set.

### GetAttackTarget

`func (o *Attack) GetAttackTarget() AttackTarget`

GetAttackTarget returns the AttackTarget field if non-nil, zero value otherwise.

### GetAttackTargetOk

`func (o *Attack) GetAttackTargetOk() (*AttackTarget, bool)`

GetAttackTargetOk returns a tuple with the AttackTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackTarget

`func (o *Attack) SetAttackTarget(v AttackTarget)`

SetAttackTarget sets AttackTarget field to given value.

### HasAttackTarget

`func (o *Attack) HasAttackTarget() bool`

HasAttackTarget returns a boolean if a field has been set.

### GetAttackType

`func (o *Attack) GetAttackType() string`

GetAttackType returns the AttackType field if non-nil, zero value otherwise.

### GetAttackTypeOk

`func (o *Attack) GetAttackTypeOk() (*string, bool)`

GetAttackTypeOk returns a tuple with the AttackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackType

`func (o *Attack) SetAttackType(v string)`

SetAttackType sets AttackType field to given value.

### HasAttackType

`func (o *Attack) HasAttackType() bool`

HasAttackType returns a boolean if a field has been set.

### GetAttacker

`func (o *Attack) GetAttacker() Attacker`

GetAttacker returns the Attacker field if non-nil, zero value otherwise.

### GetAttackerOk

`func (o *Attack) GetAttackerOk() (*Attacker, bool)`

GetAttackerOk returns a tuple with the Attacker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttacker

`func (o *Attack) SetAttacker(v Attacker)`

SetAttacker sets Attacker field to given value.

### HasAttacker

`func (o *Attack) HasAttacker() bool`

HasAttacker returns a boolean if a field has been set.

### GetDisplayId

`func (o *Attack) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *Attack) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *Attack) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *Attack) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetDisplayName

`func (o *Attack) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Attack) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Attack) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Attack) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntrypoint

`func (o *Attack) GetEntrypoint() AttackEntrypoint`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *Attack) GetEntrypointOk() (*AttackEntrypoint, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *Attack) SetEntrypoint(v AttackEntrypoint)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *Attack) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetManagementZones

`func (o *Attack) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Attack) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Attack) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Attack) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetRequest

`func (o *Attack) GetRequest() RequestInformation`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Attack) GetRequestOk() (*RequestInformation, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Attack) SetRequest(v RequestInformation)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *Attack) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetSecurityProblem

`func (o *Attack) GetSecurityProblem() AttackSecurityProblem`

GetSecurityProblem returns the SecurityProblem field if non-nil, zero value otherwise.

### GetSecurityProblemOk

`func (o *Attack) GetSecurityProblemOk() (*AttackSecurityProblem, bool)`

GetSecurityProblemOk returns a tuple with the SecurityProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblem

`func (o *Attack) SetSecurityProblem(v AttackSecurityProblem)`

SetSecurityProblem sets SecurityProblem field to given value.

### HasSecurityProblem

`func (o *Attack) HasSecurityProblem() bool`

HasSecurityProblem returns a boolean if a field has been set.

### GetState

`func (o *Attack) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Attack) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Attack) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Attack) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTechnology

`func (o *Attack) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *Attack) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *Attack) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *Attack) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetTimestamp

`func (o *Attack) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Attack) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Attack) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Attack) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetVulnerability

`func (o *Attack) GetVulnerability() Vulnerability`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *Attack) GetVulnerabilityOk() (*Vulnerability, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *Attack) SetVulnerability(v Vulnerability)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *Attack) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


