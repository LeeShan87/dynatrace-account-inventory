# UpdateToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the token. | [optional] 
**Revoked** | Pointer to **bool** | The token is revoked (&#x60;true&#x60;) or active (&#x60;false&#x60;). | [optional] 
**Scopes** | Pointer to **[]string** | The list of permissions assigned to the token.   Apart from the new permissions, you need to submit the existing permissions you want to keep, too. Any existing permission, missing in the payload, is revoked.  * &#x60;InstallerDownload&#x60;: PaaS integration - Installer download.  * &#x60;DataExport&#x60;: Access problem and event feed, metrics, and topology.  * &#x60;PluginUpload&#x60;: Upload Extension.  * &#x60;SupportAlert&#x60;: PaaS integration - Support alert.  * &#x60;AdvancedSyntheticIntegration&#x60;: Dynatrace module integration - Synthetic Classic.  * &#x60;ExternalSyntheticIntegration&#x60;: Create and read synthetic monitors, locations, and nodes.  * &#x60;RumBrowserExtension&#x60;: RUM Browser Extension.  * &#x60;LogExport&#x60;: Read logs.  * &#x60;ReadConfig&#x60;: Read configuration.  * &#x60;WriteConfig&#x60;: Write configuration.  * &#x60;DTAQLAccess&#x60;: User sessions.  * &#x60;UserSessionAnonymization&#x60;: Anonymize user session data for data privacy reasons.  * &#x60;DataPrivacy&#x60;: Change data privacy settings.  * &#x60;CaptureRequestData&#x60;: Capture request data.  * &#x60;Davis&#x60;: Dynatrace module integration - Davis.  * &#x60;DssFileManagement&#x60;: Mobile symbolication file management.  * &#x60;RumJavaScriptTagManagement&#x60;: Real user monitoring JavaScript tag management.  * &#x60;TenantTokenManagement&#x60;: Token management.  * &#x60;ActiveGateCertManagement&#x60;: ActiveGate certificate management.  * &#x60;RestRequestForwarding&#x60;: Fetch data from a remote environment.  * &#x60;ReadSyntheticData&#x60;: Read synthetic monitors, locations, and nodes.  * &#x60;DataImport&#x60;: Data ingest, e.g.: metrics and events.  * &#x60;syntheticExecutions.write&#x60;: Write synthetic monitor executions.  * &#x60;syntheticExecutions.read&#x60;: Read synthetic monitor execution results.  * &#x60;auditLogs.read&#x60;: Read audit logs.  * &#x60;metrics.read&#x60;: Read metrics.  * &#x60;metrics.write&#x60;: Write metrics.  * &#x60;entities.read&#x60;: Read entities.  * &#x60;entities.write&#x60;: Write entities.  * &#x60;problems.read&#x60;: Read problems.  * &#x60;problems.write&#x60;: Write problems.  * &#x60;events.read&#x60;: Read events.  * &#x60;events.ingest&#x60;: Ingest events.  * &#x60;analyzers.read&#x60;: Read analyzers.  * &#x60;analyzers.write&#x60;: Write &amp; execute analyzers.  * &#x60;networkZones.read&#x60;: Read network zones.  * &#x60;networkZones.write&#x60;: Write network zones.  * &#x60;activeGates.read&#x60;: Read ActiveGates.  * &#x60;activeGates.write&#x60;: Write ActiveGates.  * &#x60;activeGateTokenManagement.read&#x60;: Read ActiveGate tokens.  * &#x60;activeGateTokenManagement.create&#x60;: Create ActiveGate tokens.  * &#x60;activeGateTokenManagement.write&#x60;: Write ActiveGate tokens.  * &#x60;credentialVault.read&#x60;: Read credential vault entries.  * &#x60;credentialVault.write&#x60;: Write credential vault entries.  * &#x60;extensions.read&#x60;: Read extensions.  * &#x60;extensions.write&#x60;: Write extensions.  * &#x60;extensionConfigurations.read&#x60;: Read extension monitoring configurations.  * &#x60;extensionConfigurations.write&#x60;: Write extension monitoring configurations.  * &#x60;extensionEnvironment.read&#x60;: Read extension environment configurations.  * &#x60;extensionEnvironment.write&#x60;: Write extension environment configurations.  * &#x60;metrics.ingest&#x60;: Ingest metrics.  * &#x60;attacks.read&#x60;: Read attacks.  * &#x60;attacks.write&#x60;: Write Application Protection settings.  * &#x60;securityProblems.read&#x60;: Read security problems.  * &#x60;securityProblems.write&#x60;: Write security problems.  * &#x60;syntheticLocations.read&#x60;: Read synthetic locations.  * &#x60;syntheticLocations.write&#x60;: Write synthetic locations.  * &#x60;settings.read&#x60;: Read settings.  * &#x60;settings.write&#x60;: Write settings.  * &#x60;tenantTokenRotation.write&#x60;: Tenant token rotation.  * &#x60;slo.read&#x60;: Read SLO.  * &#x60;slo.write&#x60;: Write SLO.  * &#x60;releases.read&#x60;: Read releases.  * &#x60;apiTokens.read&#x60;: Read API tokens.  * &#x60;apiTokens.write&#x60;: Write API tokens.  * &#x60;openTelemetryTrace.ingest&#x60;: Ingest OpenTelemetry traces.  * &#x60;logs.read&#x60;: Read logs.  * &#x60;logs.ingest&#x60;: Ingest logs.  * &#x60;geographicRegions.read&#x60;: Read Geographic regions.  * &#x60;oneAgents.read&#x60;: Read OneAgents.  * &#x60;oneAgents.write&#x60;: Write OneAgents.  * &#x60;traces.lookup&#x60;: Look up a single trace.  * &#x60;hub.read&#x60;: Read Hub related data.  * &#x60;hub.write&#x60;: Manage metadata of Hub items.  * &#x60;hub.install&#x60;: Install and update Hub items.  * &#x60;javaScriptMappingFiles.read&#x60;: Read JavaScript mapping files.  * &#x60;javaScriptMappingFiles.write&#x60;: Write JavaScript mapping files.   | [optional] 

## Methods

### NewUpdateToken

`func NewUpdateToken() *UpdateToken`

NewUpdateToken instantiates a new UpdateToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTokenWithDefaults

`func NewUpdateTokenWithDefaults() *UpdateToken`

NewUpdateTokenWithDefaults instantiates a new UpdateToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevoked

`func (o *UpdateToken) GetRevoked() bool`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *UpdateToken) GetRevokedOk() (*bool, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *UpdateToken) SetRevoked(v bool)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *UpdateToken) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### GetScopes

`func (o *UpdateToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateToken) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


