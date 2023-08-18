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

// checks if the SchemaType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaType{}

// SchemaType A list of definitions of types.    A type is a complex property that contains its own set of subproperties.
type SchemaType struct {
	// A list of constraints limiting the values to be accepted.
	Constraints []ComplexConstraint `json:"constraints,omitempty"`
	// A short description of the property.
	Description string `json:"description"`
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
	// An extended description and/or links to documentation.
	Documentation string `json:"documentation"`
	// Definition of properties that can be persisted.
	Properties map[string]PropertyDefinition `json:"properties"`
	// The pattern for the summary search(for example, \"Alert after *X* minutes.\") of the configuration in the UI.
	SearchPattern *string `json:"searchPattern,omitempty"`
	// The pattern for the summary (for example, \"Alert after *X* minutes.\") of the configuration in the UI.
	SummaryPattern string `json:"summaryPattern"`
	// Type of the reference type.
	Type string `json:"type"`
	// The version of the type.
	Version string `json:"version"`
	// A short description of the version.
	VersionInfo *string `json:"versionInfo,omitempty"`
}

// NewSchemaType instantiates a new SchemaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaType(description string, documentation string, properties map[string]PropertyDefinition, summaryPattern string, type_ string, version string) *SchemaType {
	this := SchemaType{}
	this.Description = description
	this.Documentation = documentation
	this.Properties = properties
	this.SummaryPattern = summaryPattern
	this.Type = type_
	this.Version = version
	return &this
}

// NewSchemaTypeWithDefaults instantiates a new SchemaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaTypeWithDefaults() *SchemaType {
	this := SchemaType{}
	return &this
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *SchemaType) GetConstraints() []ComplexConstraint {
	if o == nil || IsNil(o.Constraints) {
		var ret []ComplexConstraint
		return ret
	}
	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetConstraintsOk() ([]ComplexConstraint, bool) {
	if o == nil || IsNil(o.Constraints) {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *SchemaType) HasConstraints() bool {
	if o != nil && !IsNil(o.Constraints) {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []ComplexConstraint and assigns it to the Constraints field.
func (o *SchemaType) SetConstraints(v []ComplexConstraint) {
	o.Constraints = v
}

// GetDescription returns the Description field value
func (o *SchemaType) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemaType) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SchemaType) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchemaType) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SchemaType) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDocumentation returns the Documentation field value
func (o *SchemaType) GetDocumentation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value
// and a boolean to check if the value has been set.
func (o *SchemaType) GetDocumentationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Documentation, true
}

// SetDocumentation sets field value
func (o *SchemaType) SetDocumentation(v string) {
	o.Documentation = v
}

// GetProperties returns the Properties field value
func (o *SchemaType) GetProperties() map[string]PropertyDefinition {
	if o == nil {
		var ret map[string]PropertyDefinition
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SchemaType) GetPropertiesOk() (*map[string]PropertyDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SchemaType) SetProperties(v map[string]PropertyDefinition) {
	o.Properties = v
}

// GetSearchPattern returns the SearchPattern field value if set, zero value otherwise.
func (o *SchemaType) GetSearchPattern() string {
	if o == nil || IsNil(o.SearchPattern) {
		var ret string
		return ret
	}
	return *o.SearchPattern
}

// GetSearchPatternOk returns a tuple with the SearchPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetSearchPatternOk() (*string, bool) {
	if o == nil || IsNil(o.SearchPattern) {
		return nil, false
	}
	return o.SearchPattern, true
}

// HasSearchPattern returns a boolean if a field has been set.
func (o *SchemaType) HasSearchPattern() bool {
	if o != nil && !IsNil(o.SearchPattern) {
		return true
	}

	return false
}

// SetSearchPattern gets a reference to the given string and assigns it to the SearchPattern field.
func (o *SchemaType) SetSearchPattern(v string) {
	o.SearchPattern = &v
}

// GetSummaryPattern returns the SummaryPattern field value
func (o *SchemaType) GetSummaryPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SummaryPattern
}

// GetSummaryPatternOk returns a tuple with the SummaryPattern field value
// and a boolean to check if the value has been set.
func (o *SchemaType) GetSummaryPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SummaryPattern, true
}

// SetSummaryPattern sets field value
func (o *SchemaType) SetSummaryPattern(v string) {
	o.SummaryPattern = v
}

// GetType returns the Type field value
func (o *SchemaType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SchemaType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SchemaType) SetType(v string) {
	o.Type = v
}

// GetVersion returns the Version field value
func (o *SchemaType) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SchemaType) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *SchemaType) SetVersion(v string) {
	o.Version = v
}

// GetVersionInfo returns the VersionInfo field value if set, zero value otherwise.
func (o *SchemaType) GetVersionInfo() string {
	if o == nil || IsNil(o.VersionInfo) {
		var ret string
		return ret
	}
	return *o.VersionInfo
}

// GetVersionInfoOk returns a tuple with the VersionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaType) GetVersionInfoOk() (*string, bool) {
	if o == nil || IsNil(o.VersionInfo) {
		return nil, false
	}
	return o.VersionInfo, true
}

// HasVersionInfo returns a boolean if a field has been set.
func (o *SchemaType) HasVersionInfo() bool {
	if o != nil && !IsNil(o.VersionInfo) {
		return true
	}

	return false
}

// SetVersionInfo gets a reference to the given string and assigns it to the VersionInfo field.
func (o *SchemaType) SetVersionInfo(v string) {
	o.VersionInfo = &v
}

func (o SchemaType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Constraints) {
		toSerialize["constraints"] = o.Constraints
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["documentation"] = o.Documentation
	toSerialize["properties"] = o.Properties
	if !IsNil(o.SearchPattern) {
		toSerialize["searchPattern"] = o.SearchPattern
	}
	toSerialize["summaryPattern"] = o.SummaryPattern
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version
	if !IsNil(o.VersionInfo) {
		toSerialize["versionInfo"] = o.VersionInfo
	}
	return toSerialize, nil
}

type NullableSchemaType struct {
	value *SchemaType
	isSet bool
}

func (v NullableSchemaType) Get() *SchemaType {
	return v.value
}

func (v *NullableSchemaType) Set(val *SchemaType) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaType) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaType(val *SchemaType) *NullableSchemaType {
	return &NullableSchemaType{value: val, isSet: true}
}

func (v NullableSchemaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

