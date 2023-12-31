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

// checks if the RemediationDetailsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemediationDetailsItem{}

// RemediationDetailsItem Detailed information of a remediation item for a security problem.
type RemediationDetailsItem struct {
	Assessment *RemediationAssessment `json:"assessment,omitempty"`
	EntityIds []string `json:"entityIds,omitempty"`
	FirstAffectedTimestamp *int64 `json:"firstAffectedTimestamp,omitempty"`
	Id *string `json:"id,omitempty"`
	MuteState *RemediationItemMuteState `json:"muteState,omitempty"`
	Name *string `json:"name,omitempty"`
	RemediationProgress *RemediationProgress `json:"remediationProgress,omitempty"`
	ResolvedTimestamp *int64 `json:"resolvedTimestamp,omitempty"`
	VulnerabilityState *string `json:"vulnerabilityState,omitempty"`
	// A list of vulnerable components of the remediation item.   A vulnerable component is what causes the security problem.
	VulnerableComponents []RemediationItemDetailsVulnerableComponent `json:"vulnerableComponents,omitempty"`
}

// NewRemediationDetailsItem instantiates a new RemediationDetailsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationDetailsItem() *RemediationDetailsItem {
	this := RemediationDetailsItem{}
	return &this
}

// NewRemediationDetailsItemWithDefaults instantiates a new RemediationDetailsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationDetailsItemWithDefaults() *RemediationDetailsItem {
	this := RemediationDetailsItem{}
	return &this
}

// GetAssessment returns the Assessment field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetAssessment() RemediationAssessment {
	if o == nil || IsNil(o.Assessment) {
		var ret RemediationAssessment
		return ret
	}
	return *o.Assessment
}

// GetAssessmentOk returns a tuple with the Assessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetAssessmentOk() (*RemediationAssessment, bool) {
	if o == nil || IsNil(o.Assessment) {
		return nil, false
	}
	return o.Assessment, true
}

// HasAssessment returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasAssessment() bool {
	if o != nil && !IsNil(o.Assessment) {
		return true
	}

	return false
}

// SetAssessment gets a reference to the given RemediationAssessment and assigns it to the Assessment field.
func (o *RemediationDetailsItem) SetAssessment(v RemediationAssessment) {
	o.Assessment = &v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetEntityIds() []string {
	if o == nil || IsNil(o.EntityIds) {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetEntityIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityIds) {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasEntityIds() bool {
	if o != nil && !IsNil(o.EntityIds) {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *RemediationDetailsItem) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetFirstAffectedTimestamp returns the FirstAffectedTimestamp field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetFirstAffectedTimestamp() int64 {
	if o == nil || IsNil(o.FirstAffectedTimestamp) {
		var ret int64
		return ret
	}
	return *o.FirstAffectedTimestamp
}

// GetFirstAffectedTimestampOk returns a tuple with the FirstAffectedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetFirstAffectedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.FirstAffectedTimestamp) {
		return nil, false
	}
	return o.FirstAffectedTimestamp, true
}

// HasFirstAffectedTimestamp returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasFirstAffectedTimestamp() bool {
	if o != nil && !IsNil(o.FirstAffectedTimestamp) {
		return true
	}

	return false
}

// SetFirstAffectedTimestamp gets a reference to the given int64 and assigns it to the FirstAffectedTimestamp field.
func (o *RemediationDetailsItem) SetFirstAffectedTimestamp(v int64) {
	o.FirstAffectedTimestamp = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemediationDetailsItem) SetId(v string) {
	o.Id = &v
}

// GetMuteState returns the MuteState field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetMuteState() RemediationItemMuteState {
	if o == nil || IsNil(o.MuteState) {
		var ret RemediationItemMuteState
		return ret
	}
	return *o.MuteState
}

// GetMuteStateOk returns a tuple with the MuteState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetMuteStateOk() (*RemediationItemMuteState, bool) {
	if o == nil || IsNil(o.MuteState) {
		return nil, false
	}
	return o.MuteState, true
}

// HasMuteState returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasMuteState() bool {
	if o != nil && !IsNil(o.MuteState) {
		return true
	}

	return false
}

// SetMuteState gets a reference to the given RemediationItemMuteState and assigns it to the MuteState field.
func (o *RemediationDetailsItem) SetMuteState(v RemediationItemMuteState) {
	o.MuteState = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemediationDetailsItem) SetName(v string) {
	o.Name = &v
}

// GetRemediationProgress returns the RemediationProgress field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetRemediationProgress() RemediationProgress {
	if o == nil || IsNil(o.RemediationProgress) {
		var ret RemediationProgress
		return ret
	}
	return *o.RemediationProgress
}

// GetRemediationProgressOk returns a tuple with the RemediationProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetRemediationProgressOk() (*RemediationProgress, bool) {
	if o == nil || IsNil(o.RemediationProgress) {
		return nil, false
	}
	return o.RemediationProgress, true
}

// HasRemediationProgress returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasRemediationProgress() bool {
	if o != nil && !IsNil(o.RemediationProgress) {
		return true
	}

	return false
}

// SetRemediationProgress gets a reference to the given RemediationProgress and assigns it to the RemediationProgress field.
func (o *RemediationDetailsItem) SetRemediationProgress(v RemediationProgress) {
	o.RemediationProgress = &v
}

// GetResolvedTimestamp returns the ResolvedTimestamp field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetResolvedTimestamp() int64 {
	if o == nil || IsNil(o.ResolvedTimestamp) {
		var ret int64
		return ret
	}
	return *o.ResolvedTimestamp
}

// GetResolvedTimestampOk returns a tuple with the ResolvedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetResolvedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.ResolvedTimestamp) {
		return nil, false
	}
	return o.ResolvedTimestamp, true
}

// HasResolvedTimestamp returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasResolvedTimestamp() bool {
	if o != nil && !IsNil(o.ResolvedTimestamp) {
		return true
	}

	return false
}

// SetResolvedTimestamp gets a reference to the given int64 and assigns it to the ResolvedTimestamp field.
func (o *RemediationDetailsItem) SetResolvedTimestamp(v int64) {
	o.ResolvedTimestamp = &v
}

// GetVulnerabilityState returns the VulnerabilityState field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetVulnerabilityState() string {
	if o == nil || IsNil(o.VulnerabilityState) {
		var ret string
		return ret
	}
	return *o.VulnerabilityState
}

// GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetVulnerabilityStateOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityState) {
		return nil, false
	}
	return o.VulnerabilityState, true
}

// HasVulnerabilityState returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasVulnerabilityState() bool {
	if o != nil && !IsNil(o.VulnerabilityState) {
		return true
	}

	return false
}

// SetVulnerabilityState gets a reference to the given string and assigns it to the VulnerabilityState field.
func (o *RemediationDetailsItem) SetVulnerabilityState(v string) {
	o.VulnerabilityState = &v
}

// GetVulnerableComponents returns the VulnerableComponents field value if set, zero value otherwise.
func (o *RemediationDetailsItem) GetVulnerableComponents() []RemediationItemDetailsVulnerableComponent {
	if o == nil || IsNil(o.VulnerableComponents) {
		var ret []RemediationItemDetailsVulnerableComponent
		return ret
	}
	return o.VulnerableComponents
}

// GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationDetailsItem) GetVulnerableComponentsOk() ([]RemediationItemDetailsVulnerableComponent, bool) {
	if o == nil || IsNil(o.VulnerableComponents) {
		return nil, false
	}
	return o.VulnerableComponents, true
}

// HasVulnerableComponents returns a boolean if a field has been set.
func (o *RemediationDetailsItem) HasVulnerableComponents() bool {
	if o != nil && !IsNil(o.VulnerableComponents) {
		return true
	}

	return false
}

// SetVulnerableComponents gets a reference to the given []RemediationItemDetailsVulnerableComponent and assigns it to the VulnerableComponents field.
func (o *RemediationDetailsItem) SetVulnerableComponents(v []RemediationItemDetailsVulnerableComponent) {
	o.VulnerableComponents = v
}

func (o RemediationDetailsItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediationDetailsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assessment) {
		toSerialize["assessment"] = o.Assessment
	}
	if !IsNil(o.EntityIds) {
		toSerialize["entityIds"] = o.EntityIds
	}
	if !IsNil(o.FirstAffectedTimestamp) {
		toSerialize["firstAffectedTimestamp"] = o.FirstAffectedTimestamp
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MuteState) {
		toSerialize["muteState"] = o.MuteState
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RemediationProgress) {
		toSerialize["remediationProgress"] = o.RemediationProgress
	}
	if !IsNil(o.ResolvedTimestamp) {
		toSerialize["resolvedTimestamp"] = o.ResolvedTimestamp
	}
	if !IsNil(o.VulnerabilityState) {
		toSerialize["vulnerabilityState"] = o.VulnerabilityState
	}
	if !IsNil(o.VulnerableComponents) {
		toSerialize["vulnerableComponents"] = o.VulnerableComponents
	}
	return toSerialize, nil
}

type NullableRemediationDetailsItem struct {
	value *RemediationDetailsItem
	isSet bool
}

func (v NullableRemediationDetailsItem) Get() *RemediationDetailsItem {
	return v.value
}

func (v *NullableRemediationDetailsItem) Set(val *RemediationDetailsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationDetailsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationDetailsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationDetailsItem(val *RemediationDetailsItem) *NullableRemediationDetailsItem {
	return &NullableRemediationDetailsItem{value: val, isSet: true}
}

func (v NullableRemediationDetailsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationDetailsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


