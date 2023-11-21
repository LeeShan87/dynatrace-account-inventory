/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the UpdateToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateToken{}

// UpdateToken struct for UpdateToken
type UpdateToken struct {
	// The name of the token.
	Name *string `json:"name,omitempty"`
	// The token is revoked (`true`) or active (`false`).
	Revoked *bool `json:"revoked,omitempty"`
	// The list of permissions assigned to the token.   Apart from the new permissions, you need to submit the existing permissions you want to keep, too. Any existing permission, missing in the payload, is revoked.  * `InstallerDownload`: PaaS integration - Installer download.  * `DataExport`: Access problem and event feed, metrics, and topology.  * `PluginUpload`: Upload Extension.  * `SupportAlert`: PaaS integration - Support alert.  * `AdvancedSyntheticIntegration`: Dynatrace module integration - Synthetic Classic.  * `ExternalSyntheticIntegration`: Create and read synthetic monitors, locations, and nodes.  * `RumBrowserExtension`: RUM Browser Extension.  * `LogExport`: Read logs.  * `ReadConfig`: Read configuration.  * `WriteConfig`: Write configuration.  * `DTAQLAccess`: User sessions.  * `UserSessionAnonymization`: Anonymize user session data for data privacy reasons.  * `DataPrivacy`: Change data privacy settings.  * `CaptureRequestData`: Capture request data.  * `Davis`: Dynatrace module integration - Davis.  * `DssFileManagement`: Mobile symbolication file management.  * `RumJavaScriptTagManagement`: Real user monitoring JavaScript tag management.  * `TenantTokenManagement`: Token management.  * `ActiveGateCertManagement`: ActiveGate certificate management.  * `RestRequestForwarding`: Fetch data from a remote environment.  * `ReadSyntheticData`: Read synthetic monitors, locations, and nodes.  * `DataImport`: Data ingest, e.g.: metrics and events.  * `syntheticExecutions.write`: Write synthetic monitor executions.  * `syntheticExecutions.read`: Read synthetic monitor execution results.  * `auditLogs.read`: Read audit logs.  * `metrics.read`: Read metrics.  * `metrics.write`: Write metrics.  * `entities.read`: Read entities.  * `entities.write`: Write entities.  * `problems.read`: Read problems.  * `problems.write`: Write problems.  * `events.read`: Read events.  * `events.ingest`: Ingest events.  * `analyzers.read`: Read analyzers.  * `analyzers.write`: Write & execute analyzers.  * `networkZones.read`: Read network zones.  * `networkZones.write`: Write network zones.  * `activeGates.read`: Read ActiveGates.  * `activeGates.write`: Write ActiveGates.  * `activeGateTokenManagement.read`: Read ActiveGate tokens.  * `activeGateTokenManagement.create`: Create ActiveGate tokens.  * `activeGateTokenManagement.write`: Write ActiveGate tokens.  * `credentialVault.read`: Read credential vault entries.  * `credentialVault.write`: Write credential vault entries.  * `extensions.read`: Read extensions.  * `extensions.write`: Write extensions.  * `extensionConfigurations.read`: Read extension monitoring configurations.  * `extensionConfigurations.write`: Write extension monitoring configurations.  * `extensionEnvironment.read`: Read extension environment configurations.  * `extensionEnvironment.write`: Write extension environment configurations.  * `metrics.ingest`: Ingest metrics.  * `attacks.read`: Read attacks.  * `attacks.write`: Write Application Protection settings.  * `securityProblems.read`: Read security problems.  * `securityProblems.write`: Write security problems.  * `syntheticLocations.read`: Read synthetic locations.  * `syntheticLocations.write`: Write synthetic locations.  * `settings.read`: Read settings.  * `settings.write`: Write settings.  * `tenantTokenRotation.write`: Tenant token rotation.  * `slo.read`: Read SLO.  * `slo.write`: Write SLO.  * `releases.read`: Read releases.  * `apiTokens.read`: Read API tokens.  * `apiTokens.write`: Write API tokens.  * `openTelemetryTrace.ingest`: Ingest OpenTelemetry traces.  * `logs.read`: Read logs.  * `logs.ingest`: Ingest logs.  * `geographicRegions.read`: Read Geographic regions.  * `oneAgents.read`: Read OneAgents.  * `oneAgents.write`: Write OneAgents.  * `traces.lookup`: Look up a single trace.  * `hub.read`: Read Hub related data.  * `hub.write`: Manage metadata of Hub items.  * `hub.install`: Install and update Hub items.  * `javaScriptMappingFiles.read`: Read JavaScript mapping files.  * `javaScriptMappingFiles.write`: Write JavaScript mapping files.  
	Scopes []string `json:"scopes,omitempty"`
}

// NewUpdateToken instantiates a new UpdateToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateToken() *UpdateToken {
	this := UpdateToken{}
	return &this
}

// NewUpdateTokenWithDefaults instantiates a new UpdateToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTokenWithDefaults() *UpdateToken {
	this := UpdateToken{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateToken) SetName(v string) {
	o.Name = &v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *UpdateToken) GetRevoked() bool {
	if o == nil || IsNil(o.Revoked) {
		var ret bool
		return ret
	}
	return *o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateToken) GetRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *UpdateToken) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given bool and assigns it to the Revoked field.
func (o *UpdateToken) SetRevoked(v bool) {
	o.Revoked = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UpdateToken) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateToken) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UpdateToken) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *UpdateToken) SetScopes(v []string) {
	o.Scopes = v
}

func (o UpdateToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableUpdateToken struct {
	value *UpdateToken
	isSet bool
}

func (v NullableUpdateToken) Get() *UpdateToken {
	return v.value
}

func (v *NullableUpdateToken) Set(val *UpdateToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateToken(val *UpdateToken) *NullableUpdateToken {
	return &NullableUpdateToken{value: val, isSet: true}
}

func (v NullableUpdateToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


