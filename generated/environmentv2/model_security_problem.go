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

// checks if the SecurityProblem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityProblem{}

// SecurityProblem Parameters of a security problem
type SecurityProblem struct {
	CodeLevelVulnerabilityDetails *CodeLevelVulnerabilityDetails `json:"codeLevelVulnerabilityDetails,omitempty"`
	// A list of CVE IDs of the security problem.
	CveIds []string `json:"cveIds,omitempty"`
	// The display ID of the security problem.
	DisplayId *string `json:"displayId,omitempty"`
	// The external vulnerability ID of the security problem.
	ExternalVulnerabilityId *string `json:"externalVulnerabilityId,omitempty"`
	// The timestamp of the first occurrence of the security problem.
	FirstSeenTimestamp *int64 `json:"firstSeenTimestamp,omitempty"`
	GlobalCounts *GlobalCountsDto `json:"globalCounts,omitempty"`
	// The timestamp when the security problem was last opened.
	LastOpenedTimestamp *int64 `json:"lastOpenedTimestamp,omitempty"`
	// The timestamp when the security problem was last resolved.
	LastResolvedTimestamp *int64 `json:"lastResolvedTimestamp,omitempty"`
	// The timestamp of the most recent security problem change.
	LastUpdatedTimestamp *int64 `json:"lastUpdatedTimestamp,omitempty"`
	// A list of management zones which the affected entities belong to.
	ManagementZones []ManagementZone `json:"managementZones,omitempty"`
	// The security problem is (`true`) or is not (`false`) muted.
	Muted *bool `json:"muted,omitempty"`
	// The package name of the security problem.
	PackageName *string `json:"packageName,omitempty"`
	RiskAssessment *RiskAssessment `json:"riskAssessment,omitempty"`
	// The ID of the security problem.
	SecurityProblemId *string `json:"securityProblemId,omitempty"`
	// The status of the security problem.
	Status *string `json:"status,omitempty"`
	// The technology of the security problem.
	Technology *string `json:"technology,omitempty"`
	// The title of the security problem.
	Title *string `json:"title,omitempty"`
	// The URL to the security problem details page.
	Url *string `json:"url,omitempty"`
	// The type of the vulnerability.
	VulnerabilityType *string `json:"vulnerabilityType,omitempty"`
}

// NewSecurityProblem instantiates a new SecurityProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityProblem() *SecurityProblem {
	this := SecurityProblem{}
	return &this
}

// NewSecurityProblemWithDefaults instantiates a new SecurityProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityProblemWithDefaults() *SecurityProblem {
	this := SecurityProblem{}
	return &this
}

// GetCodeLevelVulnerabilityDetails returns the CodeLevelVulnerabilityDetails field value if set, zero value otherwise.
func (o *SecurityProblem) GetCodeLevelVulnerabilityDetails() CodeLevelVulnerabilityDetails {
	if o == nil || IsNil(o.CodeLevelVulnerabilityDetails) {
		var ret CodeLevelVulnerabilityDetails
		return ret
	}
	return *o.CodeLevelVulnerabilityDetails
}

// GetCodeLevelVulnerabilityDetailsOk returns a tuple with the CodeLevelVulnerabilityDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetCodeLevelVulnerabilityDetailsOk() (*CodeLevelVulnerabilityDetails, bool) {
	if o == nil || IsNil(o.CodeLevelVulnerabilityDetails) {
		return nil, false
	}
	return o.CodeLevelVulnerabilityDetails, true
}

// HasCodeLevelVulnerabilityDetails returns a boolean if a field has been set.
func (o *SecurityProblem) HasCodeLevelVulnerabilityDetails() bool {
	if o != nil && !IsNil(o.CodeLevelVulnerabilityDetails) {
		return true
	}

	return false
}

// SetCodeLevelVulnerabilityDetails gets a reference to the given CodeLevelVulnerabilityDetails and assigns it to the CodeLevelVulnerabilityDetails field.
func (o *SecurityProblem) SetCodeLevelVulnerabilityDetails(v CodeLevelVulnerabilityDetails) {
	o.CodeLevelVulnerabilityDetails = &v
}

// GetCveIds returns the CveIds field value if set, zero value otherwise.
func (o *SecurityProblem) GetCveIds() []string {
	if o == nil || IsNil(o.CveIds) {
		var ret []string
		return ret
	}
	return o.CveIds
}

// GetCveIdsOk returns a tuple with the CveIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetCveIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CveIds) {
		return nil, false
	}
	return o.CveIds, true
}

// HasCveIds returns a boolean if a field has been set.
func (o *SecurityProblem) HasCveIds() bool {
	if o != nil && !IsNil(o.CveIds) {
		return true
	}

	return false
}

// SetCveIds gets a reference to the given []string and assigns it to the CveIds field.
func (o *SecurityProblem) SetCveIds(v []string) {
	o.CveIds = v
}

// GetDisplayId returns the DisplayId field value if set, zero value otherwise.
func (o *SecurityProblem) GetDisplayId() string {
	if o == nil || IsNil(o.DisplayId) {
		var ret string
		return ret
	}
	return *o.DisplayId
}

// GetDisplayIdOk returns a tuple with the DisplayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetDisplayIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayId) {
		return nil, false
	}
	return o.DisplayId, true
}

// HasDisplayId returns a boolean if a field has been set.
func (o *SecurityProblem) HasDisplayId() bool {
	if o != nil && !IsNil(o.DisplayId) {
		return true
	}

	return false
}

// SetDisplayId gets a reference to the given string and assigns it to the DisplayId field.
func (o *SecurityProblem) SetDisplayId(v string) {
	o.DisplayId = &v
}

// GetExternalVulnerabilityId returns the ExternalVulnerabilityId field value if set, zero value otherwise.
func (o *SecurityProblem) GetExternalVulnerabilityId() string {
	if o == nil || IsNil(o.ExternalVulnerabilityId) {
		var ret string
		return ret
	}
	return *o.ExternalVulnerabilityId
}

// GetExternalVulnerabilityIdOk returns a tuple with the ExternalVulnerabilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetExternalVulnerabilityIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalVulnerabilityId) {
		return nil, false
	}
	return o.ExternalVulnerabilityId, true
}

// HasExternalVulnerabilityId returns a boolean if a field has been set.
func (o *SecurityProblem) HasExternalVulnerabilityId() bool {
	if o != nil && !IsNil(o.ExternalVulnerabilityId) {
		return true
	}

	return false
}

// SetExternalVulnerabilityId gets a reference to the given string and assigns it to the ExternalVulnerabilityId field.
func (o *SecurityProblem) SetExternalVulnerabilityId(v string) {
	o.ExternalVulnerabilityId = &v
}

// GetFirstSeenTimestamp returns the FirstSeenTimestamp field value if set, zero value otherwise.
func (o *SecurityProblem) GetFirstSeenTimestamp() int64 {
	if o == nil || IsNil(o.FirstSeenTimestamp) {
		var ret int64
		return ret
	}
	return *o.FirstSeenTimestamp
}

// GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetFirstSeenTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.FirstSeenTimestamp) {
		return nil, false
	}
	return o.FirstSeenTimestamp, true
}

// HasFirstSeenTimestamp returns a boolean if a field has been set.
func (o *SecurityProblem) HasFirstSeenTimestamp() bool {
	if o != nil && !IsNil(o.FirstSeenTimestamp) {
		return true
	}

	return false
}

// SetFirstSeenTimestamp gets a reference to the given int64 and assigns it to the FirstSeenTimestamp field.
func (o *SecurityProblem) SetFirstSeenTimestamp(v int64) {
	o.FirstSeenTimestamp = &v
}

// GetGlobalCounts returns the GlobalCounts field value if set, zero value otherwise.
func (o *SecurityProblem) GetGlobalCounts() GlobalCountsDto {
	if o == nil || IsNil(o.GlobalCounts) {
		var ret GlobalCountsDto
		return ret
	}
	return *o.GlobalCounts
}

// GetGlobalCountsOk returns a tuple with the GlobalCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetGlobalCountsOk() (*GlobalCountsDto, bool) {
	if o == nil || IsNil(o.GlobalCounts) {
		return nil, false
	}
	return o.GlobalCounts, true
}

// HasGlobalCounts returns a boolean if a field has been set.
func (o *SecurityProblem) HasGlobalCounts() bool {
	if o != nil && !IsNil(o.GlobalCounts) {
		return true
	}

	return false
}

// SetGlobalCounts gets a reference to the given GlobalCountsDto and assigns it to the GlobalCounts field.
func (o *SecurityProblem) SetGlobalCounts(v GlobalCountsDto) {
	o.GlobalCounts = &v
}

// GetLastOpenedTimestamp returns the LastOpenedTimestamp field value if set, zero value otherwise.
func (o *SecurityProblem) GetLastOpenedTimestamp() int64 {
	if o == nil || IsNil(o.LastOpenedTimestamp) {
		var ret int64
		return ret
	}
	return *o.LastOpenedTimestamp
}

// GetLastOpenedTimestampOk returns a tuple with the LastOpenedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetLastOpenedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.LastOpenedTimestamp) {
		return nil, false
	}
	return o.LastOpenedTimestamp, true
}

// HasLastOpenedTimestamp returns a boolean if a field has been set.
func (o *SecurityProblem) HasLastOpenedTimestamp() bool {
	if o != nil && !IsNil(o.LastOpenedTimestamp) {
		return true
	}

	return false
}

// SetLastOpenedTimestamp gets a reference to the given int64 and assigns it to the LastOpenedTimestamp field.
func (o *SecurityProblem) SetLastOpenedTimestamp(v int64) {
	o.LastOpenedTimestamp = &v
}

// GetLastResolvedTimestamp returns the LastResolvedTimestamp field value if set, zero value otherwise.
func (o *SecurityProblem) GetLastResolvedTimestamp() int64 {
	if o == nil || IsNil(o.LastResolvedTimestamp) {
		var ret int64
		return ret
	}
	return *o.LastResolvedTimestamp
}

// GetLastResolvedTimestampOk returns a tuple with the LastResolvedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetLastResolvedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.LastResolvedTimestamp) {
		return nil, false
	}
	return o.LastResolvedTimestamp, true
}

// HasLastResolvedTimestamp returns a boolean if a field has been set.
func (o *SecurityProblem) HasLastResolvedTimestamp() bool {
	if o != nil && !IsNil(o.LastResolvedTimestamp) {
		return true
	}

	return false
}

// SetLastResolvedTimestamp gets a reference to the given int64 and assigns it to the LastResolvedTimestamp field.
func (o *SecurityProblem) SetLastResolvedTimestamp(v int64) {
	o.LastResolvedTimestamp = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *SecurityProblem) GetLastUpdatedTimestamp() int64 {
	if o == nil || IsNil(o.LastUpdatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdatedTimestamp) {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *SecurityProblem) HasLastUpdatedTimestamp() bool {
	if o != nil && !IsNil(o.LastUpdatedTimestamp) {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *SecurityProblem) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetManagementZones returns the ManagementZones field value if set, zero value otherwise.
func (o *SecurityProblem) GetManagementZones() []ManagementZone {
	if o == nil || IsNil(o.ManagementZones) {
		var ret []ManagementZone
		return ret
	}
	return o.ManagementZones
}

// GetManagementZonesOk returns a tuple with the ManagementZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetManagementZonesOk() ([]ManagementZone, bool) {
	if o == nil || IsNil(o.ManagementZones) {
		return nil, false
	}
	return o.ManagementZones, true
}

// HasManagementZones returns a boolean if a field has been set.
func (o *SecurityProblem) HasManagementZones() bool {
	if o != nil && !IsNil(o.ManagementZones) {
		return true
	}

	return false
}

// SetManagementZones gets a reference to the given []ManagementZone and assigns it to the ManagementZones field.
func (o *SecurityProblem) SetManagementZones(v []ManagementZone) {
	o.ManagementZones = v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *SecurityProblem) GetMuted() bool {
	if o == nil || IsNil(o.Muted) {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.Muted) {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *SecurityProblem) HasMuted() bool {
	if o != nil && !IsNil(o.Muted) {
		return true
	}

	return false
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *SecurityProblem) SetMuted(v bool) {
	o.Muted = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *SecurityProblem) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *SecurityProblem) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *SecurityProblem) SetPackageName(v string) {
	o.PackageName = &v
}

// GetRiskAssessment returns the RiskAssessment field value if set, zero value otherwise.
func (o *SecurityProblem) GetRiskAssessment() RiskAssessment {
	if o == nil || IsNil(o.RiskAssessment) {
		var ret RiskAssessment
		return ret
	}
	return *o.RiskAssessment
}

// GetRiskAssessmentOk returns a tuple with the RiskAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetRiskAssessmentOk() (*RiskAssessment, bool) {
	if o == nil || IsNil(o.RiskAssessment) {
		return nil, false
	}
	return o.RiskAssessment, true
}

// HasRiskAssessment returns a boolean if a field has been set.
func (o *SecurityProblem) HasRiskAssessment() bool {
	if o != nil && !IsNil(o.RiskAssessment) {
		return true
	}

	return false
}

// SetRiskAssessment gets a reference to the given RiskAssessment and assigns it to the RiskAssessment field.
func (o *SecurityProblem) SetRiskAssessment(v RiskAssessment) {
	o.RiskAssessment = &v
}

// GetSecurityProblemId returns the SecurityProblemId field value if set, zero value otherwise.
func (o *SecurityProblem) GetSecurityProblemId() string {
	if o == nil || IsNil(o.SecurityProblemId) {
		var ret string
		return ret
	}
	return *o.SecurityProblemId
}

// GetSecurityProblemIdOk returns a tuple with the SecurityProblemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetSecurityProblemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityProblemId) {
		return nil, false
	}
	return o.SecurityProblemId, true
}

// HasSecurityProblemId returns a boolean if a field has been set.
func (o *SecurityProblem) HasSecurityProblemId() bool {
	if o != nil && !IsNil(o.SecurityProblemId) {
		return true
	}

	return false
}

// SetSecurityProblemId gets a reference to the given string and assigns it to the SecurityProblemId field.
func (o *SecurityProblem) SetSecurityProblemId(v string) {
	o.SecurityProblemId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityProblem) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityProblem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityProblem) SetStatus(v string) {
	o.Status = &v
}

// GetTechnology returns the Technology field value if set, zero value otherwise.
func (o *SecurityProblem) GetTechnology() string {
	if o == nil || IsNil(o.Technology) {
		var ret string
		return ret
	}
	return *o.Technology
}

// GetTechnologyOk returns a tuple with the Technology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetTechnologyOk() (*string, bool) {
	if o == nil || IsNil(o.Technology) {
		return nil, false
	}
	return o.Technology, true
}

// HasTechnology returns a boolean if a field has been set.
func (o *SecurityProblem) HasTechnology() bool {
	if o != nil && !IsNil(o.Technology) {
		return true
	}

	return false
}

// SetTechnology gets a reference to the given string and assigns it to the Technology field.
func (o *SecurityProblem) SetTechnology(v string) {
	o.Technology = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SecurityProblem) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SecurityProblem) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SecurityProblem) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SecurityProblem) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SecurityProblem) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SecurityProblem) SetUrl(v string) {
	o.Url = &v
}

// GetVulnerabilityType returns the VulnerabilityType field value if set, zero value otherwise.
func (o *SecurityProblem) GetVulnerabilityType() string {
	if o == nil || IsNil(o.VulnerabilityType) {
		var ret string
		return ret
	}
	return *o.VulnerabilityType
}

// GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblem) GetVulnerabilityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityType) {
		return nil, false
	}
	return o.VulnerabilityType, true
}

// HasVulnerabilityType returns a boolean if a field has been set.
func (o *SecurityProblem) HasVulnerabilityType() bool {
	if o != nil && !IsNil(o.VulnerabilityType) {
		return true
	}

	return false
}

// SetVulnerabilityType gets a reference to the given string and assigns it to the VulnerabilityType field.
func (o *SecurityProblem) SetVulnerabilityType(v string) {
	o.VulnerabilityType = &v
}

func (o SecurityProblem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityProblem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodeLevelVulnerabilityDetails) {
		toSerialize["codeLevelVulnerabilityDetails"] = o.CodeLevelVulnerabilityDetails
	}
	if !IsNil(o.CveIds) {
		toSerialize["cveIds"] = o.CveIds
	}
	if !IsNil(o.DisplayId) {
		toSerialize["displayId"] = o.DisplayId
	}
	if !IsNil(o.ExternalVulnerabilityId) {
		toSerialize["externalVulnerabilityId"] = o.ExternalVulnerabilityId
	}
	if !IsNil(o.FirstSeenTimestamp) {
		toSerialize["firstSeenTimestamp"] = o.FirstSeenTimestamp
	}
	if !IsNil(o.GlobalCounts) {
		toSerialize["globalCounts"] = o.GlobalCounts
	}
	if !IsNil(o.LastOpenedTimestamp) {
		toSerialize["lastOpenedTimestamp"] = o.LastOpenedTimestamp
	}
	if !IsNil(o.LastResolvedTimestamp) {
		toSerialize["lastResolvedTimestamp"] = o.LastResolvedTimestamp
	}
	if !IsNil(o.LastUpdatedTimestamp) {
		toSerialize["lastUpdatedTimestamp"] = o.LastUpdatedTimestamp
	}
	if !IsNil(o.ManagementZones) {
		toSerialize["managementZones"] = o.ManagementZones
	}
	if !IsNil(o.Muted) {
		toSerialize["muted"] = o.Muted
	}
	if !IsNil(o.PackageName) {
		toSerialize["packageName"] = o.PackageName
	}
	if !IsNil(o.RiskAssessment) {
		toSerialize["riskAssessment"] = o.RiskAssessment
	}
	if !IsNil(o.SecurityProblemId) {
		toSerialize["securityProblemId"] = o.SecurityProblemId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Technology) {
		toSerialize["technology"] = o.Technology
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.VulnerabilityType) {
		toSerialize["vulnerabilityType"] = o.VulnerabilityType
	}
	return toSerialize, nil
}

type NullableSecurityProblem struct {
	value *SecurityProblem
	isSet bool
}

func (v NullableSecurityProblem) Get() *SecurityProblem {
	return v.value
}

func (v *NullableSecurityProblem) Set(val *SecurityProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityProblem(val *SecurityProblem) *NullableSecurityProblem {
	return &NullableSecurityProblem{value: val, isSet: true}
}

func (v NullableSecurityProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


