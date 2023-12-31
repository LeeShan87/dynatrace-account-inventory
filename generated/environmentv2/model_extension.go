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

// checks if the Extension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Extension{}

// Extension struct for Extension
type Extension struct {
	Author *AuthorDto `json:"author,omitempty"`
	// Data sources that extension uses to gather data
	DataSources []string `json:"dataSources,omitempty"`
	// Extension name
	ExtensionName *string `json:"extensionName,omitempty"`
	// Available feature sets
	FeatureSets []string `json:"featureSets,omitempty"`
	// Details of feature sets
	FeatureSetsDetails *map[string]FeatureSetDetails `json:"featureSetsDetails,omitempty"`
	// SHA-256 hash of uploaded Extension file
	FileHash *string `json:"fileHash,omitempty"`
	// Minimal Dynatrace version that works with the extension
	MinDynatraceVersion *string `json:"minDynatraceVersion,omitempty"`
	// Minimal Extension Execution Controller version that works with the extension
	MinEECVersion *string `json:"minEECVersion,omitempty"`
	// Custom variables used in extension configuration
	Variables []string `json:"variables,omitempty"`
	// Extension version
	Version *string `json:"version,omitempty"`
}

// NewExtension instantiates a new Extension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtension() *Extension {
	this := Extension{}
	return &this
}

// NewExtensionWithDefaults instantiates a new Extension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionWithDefaults() *Extension {
	this := Extension{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Extension) GetAuthor() AuthorDto {
	if o == nil || IsNil(o.Author) {
		var ret AuthorDto
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetAuthorOk() (*AuthorDto, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Extension) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AuthorDto and assigns it to the Author field.
func (o *Extension) SetAuthor(v AuthorDto) {
	o.Author = &v
}

// GetDataSources returns the DataSources field value if set, zero value otherwise.
func (o *Extension) GetDataSources() []string {
	if o == nil || IsNil(o.DataSources) {
		var ret []string
		return ret
	}
	return o.DataSources
}

// GetDataSourcesOk returns a tuple with the DataSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetDataSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.DataSources) {
		return nil, false
	}
	return o.DataSources, true
}

// HasDataSources returns a boolean if a field has been set.
func (o *Extension) HasDataSources() bool {
	if o != nil && !IsNil(o.DataSources) {
		return true
	}

	return false
}

// SetDataSources gets a reference to the given []string and assigns it to the DataSources field.
func (o *Extension) SetDataSources(v []string) {
	o.DataSources = v
}

// GetExtensionName returns the ExtensionName field value if set, zero value otherwise.
func (o *Extension) GetExtensionName() string {
	if o == nil || IsNil(o.ExtensionName) {
		var ret string
		return ret
	}
	return *o.ExtensionName
}

// GetExtensionNameOk returns a tuple with the ExtensionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetExtensionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExtensionName) {
		return nil, false
	}
	return o.ExtensionName, true
}

// HasExtensionName returns a boolean if a field has been set.
func (o *Extension) HasExtensionName() bool {
	if o != nil && !IsNil(o.ExtensionName) {
		return true
	}

	return false
}

// SetExtensionName gets a reference to the given string and assigns it to the ExtensionName field.
func (o *Extension) SetExtensionName(v string) {
	o.ExtensionName = &v
}

// GetFeatureSets returns the FeatureSets field value if set, zero value otherwise.
func (o *Extension) GetFeatureSets() []string {
	if o == nil || IsNil(o.FeatureSets) {
		var ret []string
		return ret
	}
	return o.FeatureSets
}

// GetFeatureSetsOk returns a tuple with the FeatureSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetFeatureSetsOk() ([]string, bool) {
	if o == nil || IsNil(o.FeatureSets) {
		return nil, false
	}
	return o.FeatureSets, true
}

// HasFeatureSets returns a boolean if a field has been set.
func (o *Extension) HasFeatureSets() bool {
	if o != nil && !IsNil(o.FeatureSets) {
		return true
	}

	return false
}

// SetFeatureSets gets a reference to the given []string and assigns it to the FeatureSets field.
func (o *Extension) SetFeatureSets(v []string) {
	o.FeatureSets = v
}

// GetFeatureSetsDetails returns the FeatureSetsDetails field value if set, zero value otherwise.
func (o *Extension) GetFeatureSetsDetails() map[string]FeatureSetDetails {
	if o == nil || IsNil(o.FeatureSetsDetails) {
		var ret map[string]FeatureSetDetails
		return ret
	}
	return *o.FeatureSetsDetails
}

// GetFeatureSetsDetailsOk returns a tuple with the FeatureSetsDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetFeatureSetsDetailsOk() (*map[string]FeatureSetDetails, bool) {
	if o == nil || IsNil(o.FeatureSetsDetails) {
		return nil, false
	}
	return o.FeatureSetsDetails, true
}

// HasFeatureSetsDetails returns a boolean if a field has been set.
func (o *Extension) HasFeatureSetsDetails() bool {
	if o != nil && !IsNil(o.FeatureSetsDetails) {
		return true
	}

	return false
}

// SetFeatureSetsDetails gets a reference to the given map[string]FeatureSetDetails and assigns it to the FeatureSetsDetails field.
func (o *Extension) SetFeatureSetsDetails(v map[string]FeatureSetDetails) {
	o.FeatureSetsDetails = &v
}

// GetFileHash returns the FileHash field value if set, zero value otherwise.
func (o *Extension) GetFileHash() string {
	if o == nil || IsNil(o.FileHash) {
		var ret string
		return ret
	}
	return *o.FileHash
}

// GetFileHashOk returns a tuple with the FileHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetFileHashOk() (*string, bool) {
	if o == nil || IsNil(o.FileHash) {
		return nil, false
	}
	return o.FileHash, true
}

// HasFileHash returns a boolean if a field has been set.
func (o *Extension) HasFileHash() bool {
	if o != nil && !IsNil(o.FileHash) {
		return true
	}

	return false
}

// SetFileHash gets a reference to the given string and assigns it to the FileHash field.
func (o *Extension) SetFileHash(v string) {
	o.FileHash = &v
}

// GetMinDynatraceVersion returns the MinDynatraceVersion field value if set, zero value otherwise.
func (o *Extension) GetMinDynatraceVersion() string {
	if o == nil || IsNil(o.MinDynatraceVersion) {
		var ret string
		return ret
	}
	return *o.MinDynatraceVersion
}

// GetMinDynatraceVersionOk returns a tuple with the MinDynatraceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetMinDynatraceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinDynatraceVersion) {
		return nil, false
	}
	return o.MinDynatraceVersion, true
}

// HasMinDynatraceVersion returns a boolean if a field has been set.
func (o *Extension) HasMinDynatraceVersion() bool {
	if o != nil && !IsNil(o.MinDynatraceVersion) {
		return true
	}

	return false
}

// SetMinDynatraceVersion gets a reference to the given string and assigns it to the MinDynatraceVersion field.
func (o *Extension) SetMinDynatraceVersion(v string) {
	o.MinDynatraceVersion = &v
}

// GetMinEECVersion returns the MinEECVersion field value if set, zero value otherwise.
func (o *Extension) GetMinEECVersion() string {
	if o == nil || IsNil(o.MinEECVersion) {
		var ret string
		return ret
	}
	return *o.MinEECVersion
}

// GetMinEECVersionOk returns a tuple with the MinEECVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetMinEECVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinEECVersion) {
		return nil, false
	}
	return o.MinEECVersion, true
}

// HasMinEECVersion returns a boolean if a field has been set.
func (o *Extension) HasMinEECVersion() bool {
	if o != nil && !IsNil(o.MinEECVersion) {
		return true
	}

	return false
}

// SetMinEECVersion gets a reference to the given string and assigns it to the MinEECVersion field.
func (o *Extension) SetMinEECVersion(v string) {
	o.MinEECVersion = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *Extension) GetVariables() []string {
	if o == nil || IsNil(o.Variables) {
		var ret []string
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetVariablesOk() ([]string, bool) {
	if o == nil || IsNil(o.Variables) {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *Extension) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given []string and assigns it to the Variables field.
func (o *Extension) SetVariables(v []string) {
	o.Variables = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Extension) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Extension) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Extension) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Extension) SetVersion(v string) {
	o.Version = &v
}

func (o Extension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Extension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.DataSources) {
		toSerialize["dataSources"] = o.DataSources
	}
	if !IsNil(o.ExtensionName) {
		toSerialize["extensionName"] = o.ExtensionName
	}
	if !IsNil(o.FeatureSets) {
		toSerialize["featureSets"] = o.FeatureSets
	}
	if !IsNil(o.FeatureSetsDetails) {
		toSerialize["featureSetsDetails"] = o.FeatureSetsDetails
	}
	if !IsNil(o.FileHash) {
		toSerialize["fileHash"] = o.FileHash
	}
	if !IsNil(o.MinDynatraceVersion) {
		toSerialize["minDynatraceVersion"] = o.MinDynatraceVersion
	}
	if !IsNil(o.MinEECVersion) {
		toSerialize["minEECVersion"] = o.MinEECVersion
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableExtension struct {
	value *Extension
	isSet bool
}

func (v NullableExtension) Get() *Extension {
	return v.value
}

func (v *NullableExtension) Set(val *Extension) {
	v.value = val
	v.isSet = true
}

func (v NullableExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtension(val *Extension) *NullableExtension {
	return &NullableExtension{value: val, isSet: true}
}

func (v NullableExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


