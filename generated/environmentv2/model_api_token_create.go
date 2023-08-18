/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the ApiTokenCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTokenCreate{}

// ApiTokenCreate Parameters of a new API token.
type ApiTokenCreate struct {
	// The expiration date of the token.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the token never expires.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// The name of the token.
	Name string `json:"name"`
	// The token is a personal access token (`true`) or an API token (`false`).    Personal access tokens are tied to the permissions of their owner.
	PersonalAccessToken *bool `json:"personalAccessToken,omitempty"`
	// A list of the scopes to be assigned to the token.  * `InstallerDownload`: PaaS integration - Installer download.  * `DataExport`: Access problem and event feed, metrics, and topology.  * `PluginUpload`: Upload Extension.  * `SupportAlert`: PaaS integration - Support alert.  * `AdvancedSyntheticIntegration`: Dynatrace module integration - Synthetic Classic.  * `ExternalSyntheticIntegration`: Create and read synthetic monitors, locations, and nodes.  * `RumBrowserExtension`: RUM Browser Extension.  * `LogExport`: Read logs.  * `ReadConfig`: Read configuration.  * `WriteConfig`: Write configuration.  * `DTAQLAccess`: User sessions.  * `UserSessionAnonymization`: Anonymize user session data for data privacy reasons.  * `DataPrivacy`: Change data privacy settings.  * `CaptureRequestData`: Capture request data.  * `Davis`: Dynatrace module integration - Davis.  * `DssFileManagement`: Mobile symbolication file management.  * `RumJavaScriptTagManagement`: Real user monitoring JavaScript tag management.  * `TenantTokenManagement`: Token management.  * `ActiveGateCertManagement`: ActiveGate certificate management.  * `RestRequestForwarding`: Fetch data from a remote environment.  * `ReadSyntheticData`: Read synthetic monitors, locations, and nodes.  * `DataImport`: Data ingest, e.g.: metrics and events.  * `syntheticExecutions.write`: Write synthetic monitor executions.  * `syntheticExecutions.read`: Read synthetic monitor execution results.  * `auditLogs.read`: Read audit logs.  * `metrics.read`: Read metrics.  * `metrics.write`: Write metrics.  * `entities.read`: Read entities.  * `entities.write`: Write entities.  * `problems.read`: Read problems.  * `problems.write`: Write problems.  * `events.read`: Read events.  * `events.ingest`: Ingest events.  * `analyzers.read`: Read analyzers.  * `analyzers.write`: Write & execute analyzers.  * `networkZones.read`: Read network zones.  * `networkZones.write`: Write network zones.  * `activeGates.read`: Read ActiveGates.  * `activeGates.write`: Write ActiveGates.  * `activeGateTokenManagement.read`: Read ActiveGate tokens.  * `activeGateTokenManagement.create`: Create ActiveGate tokens.  * `activeGateTokenManagement.write`: Write ActiveGate tokens.  * `credentialVault.read`: Read credential vault entries.  * `credentialVault.write`: Write credential vault entries.  * `extensions.read`: Read extensions.  * `extensions.write`: Write extensions.  * `extensionConfigurations.read`: Read extension monitoring configurations.  * `extensionConfigurations.write`: Write extension monitoring configurations.  * `extensionEnvironment.read`: Read extension environment configurations.  * `extensionEnvironment.write`: Write extension environment configurations.  * `metrics.ingest`: Ingest metrics.  * `attacks.read`: Read attacks.  * `attacks.write`: Write Application Protection settings.  * `securityProblems.read`: Read security problems.  * `securityProblems.write`: Write security problems.  * `syntheticLocations.read`: Read synthetic locations.  * `syntheticLocations.write`: Write synthetic locations.  * `settings.read`: Read settings.  * `settings.write`: Write settings.  * `tenantTokenRotation.write`: Tenant token rotation.  * `slo.read`: Read SLO.  * `slo.write`: Write SLO.  * `releases.read`: Read releases.  * `apiTokens.read`: Read API tokens.  * `apiTokens.write`: Write API tokens.  * `openTelemetryTrace.ingest`: Ingest OpenTelemetry traces.  * `logs.read`: Read logs.  * `logs.ingest`: Ingest logs.  * `geographicRegions.read`: Read Geographic regions.  * `oneAgents.read`: Read OneAgents.  * `oneAgents.write`: Write OneAgents.  * `traces.lookup`: Look up a single trace.  * `hub.read`: Read Hub related data.  * `hub.write`: Manage metadata of Hub items.  * `hub.install`: Install and update Hub items.  * `javaScriptMappingFiles.read`: Read JavaScript mapping files.  * `javaScriptMappingFiles.write`: Write JavaScript mapping files.  
	Scopes []string `json:"scopes"`
}

// NewApiTokenCreate instantiates a new ApiTokenCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTokenCreate(name string, scopes []string) *ApiTokenCreate {
	this := ApiTokenCreate{}
	this.Name = name
	this.Scopes = scopes
	return &this
}

// NewApiTokenCreateWithDefaults instantiates a new ApiTokenCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenCreateWithDefaults() *ApiTokenCreate {
	this := ApiTokenCreate{}
	return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ApiTokenCreate) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenCreate) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ApiTokenCreate) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ApiTokenCreate) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetName returns the Name field value
func (o *ApiTokenCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiTokenCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiTokenCreate) SetName(v string) {
	o.Name = v
}

// GetPersonalAccessToken returns the PersonalAccessToken field value if set, zero value otherwise.
func (o *ApiTokenCreate) GetPersonalAccessToken() bool {
	if o == nil || IsNil(o.PersonalAccessToken) {
		var ret bool
		return ret
	}
	return *o.PersonalAccessToken
}

// GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTokenCreate) GetPersonalAccessTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.PersonalAccessToken) {
		return nil, false
	}
	return o.PersonalAccessToken, true
}

// HasPersonalAccessToken returns a boolean if a field has been set.
func (o *ApiTokenCreate) HasPersonalAccessToken() bool {
	if o != nil && !IsNil(o.PersonalAccessToken) {
		return true
	}

	return false
}

// SetPersonalAccessToken gets a reference to the given bool and assigns it to the PersonalAccessToken field.
func (o *ApiTokenCreate) SetPersonalAccessToken(v bool) {
	o.PersonalAccessToken = &v
}

// GetScopes returns the Scopes field value
func (o *ApiTokenCreate) GetScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *ApiTokenCreate) GetScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *ApiTokenCreate) SetScopes(v []string) {
	o.Scopes = v
}

func (o ApiTokenCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTokenCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PersonalAccessToken) {
		toSerialize["personalAccessToken"] = o.PersonalAccessToken
	}
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

type NullableApiTokenCreate struct {
	value *ApiTokenCreate
	isSet bool
}

func (v NullableApiTokenCreate) Get() *ApiTokenCreate {
	return v.value
}

func (v *NullableApiTokenCreate) Set(val *ApiTokenCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTokenCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTokenCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTokenCreate(val *ApiTokenCreate) *NullableApiTokenCreate {
	return &NullableApiTokenCreate{value: val, isSet: true}
}

func (v NullableApiTokenCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTokenCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

