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

// checks if the SettingsObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsObject{}

// SettingsObject A settings object.
type SettingsObject struct {
	// The user (identified by a user ID or a public token ID) who performed that most recent modification.
	Author *string `json:"author,omitempty"`
	// The timestamp of the creation.
	Created *int64 `json:"created,omitempty"`
	// The unique identifier of the user who created the settings object.
	CreatedBy *string `json:"createdBy,omitempty"`
	// The external identifier of the settings object.
	ExternalId *string `json:"externalId,omitempty"`
	ModificationInfo *ModificationInfo `json:"modificationInfo,omitempty"`
	// The timestamp of the last modification.
	Modified *int64 `json:"modified,omitempty"`
	// The unique identifier of the user who performed the most recent modification.
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// The ID of the settings object.
	ObjectId *string `json:"objectId,omitempty"`
	// The schema on which the object is based.
	SchemaId *string `json:"schemaId,omitempty"`
	// The version of the schema on which the object is based.
	SchemaVersion *string `json:"schemaVersion,omitempty"`
	// The scope that the object targets. For more details, please see [Dynatrace Documentation](https://dt-url.net/ky03459). 
	Scope *string `json:"scope,omitempty"`
	// A searchable summary string of the setting value. Plain text without Markdown.
	SearchSummary *string `json:"searchSummary,omitempty"`
	// A short summary of settings. This can contain Markdown and will be escaped accordingly.
	Summary *string `json:"summary,omitempty"`
	// The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn't any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks.
	UpdateToken *string `json:"updateToken,omitempty"`
	// The value of the setting.    It defines the actual values of settings' parameters.   The actual content depends on the object's schema.
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewSettingsObject instantiates a new SettingsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsObject() *SettingsObject {
	this := SettingsObject{}
	return &this
}

// NewSettingsObjectWithDefaults instantiates a new SettingsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsObjectWithDefaults() *SettingsObject {
	this := SettingsObject{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *SettingsObject) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *SettingsObject) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *SettingsObject) SetAuthor(v string) {
	o.Author = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *SettingsObject) GetCreated() int64 {
	if o == nil || IsNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *SettingsObject) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *SettingsObject) SetCreated(v int64) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SettingsObject) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SettingsObject) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *SettingsObject) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SettingsObject) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *SettingsObject) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SettingsObject) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetModificationInfo returns the ModificationInfo field value if set, zero value otherwise.
func (o *SettingsObject) GetModificationInfo() ModificationInfo {
	if o == nil || IsNil(o.ModificationInfo) {
		var ret ModificationInfo
		return ret
	}
	return *o.ModificationInfo
}

// GetModificationInfoOk returns a tuple with the ModificationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetModificationInfoOk() (*ModificationInfo, bool) {
	if o == nil || IsNil(o.ModificationInfo) {
		return nil, false
	}
	return o.ModificationInfo, true
}

// HasModificationInfo returns a boolean if a field has been set.
func (o *SettingsObject) HasModificationInfo() bool {
	if o != nil && !IsNil(o.ModificationInfo) {
		return true
	}

	return false
}

// SetModificationInfo gets a reference to the given ModificationInfo and assigns it to the ModificationInfo field.
func (o *SettingsObject) SetModificationInfo(v ModificationInfo) {
	o.ModificationInfo = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *SettingsObject) GetModified() int64 {
	if o == nil || IsNil(o.Modified) {
		var ret int64
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetModifiedOk() (*int64, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *SettingsObject) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given int64 and assigns it to the Modified field.
func (o *SettingsObject) SetModified(v int64) {
	o.Modified = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *SettingsObject) GetModifiedBy() string {
	if o == nil || IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetModifiedByOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *SettingsObject) HasModifiedBy() bool {
	if o != nil && !IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *SettingsObject) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *SettingsObject) GetObjectId() string {
	if o == nil || IsNil(o.ObjectId) {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *SettingsObject) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *SettingsObject) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *SettingsObject) GetSchemaId() string {
	if o == nil || IsNil(o.SchemaId) {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetSchemaIdOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaId) {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *SettingsObject) HasSchemaId() bool {
	if o != nil && !IsNil(o.SchemaId) {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *SettingsObject) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetSchemaVersion returns the SchemaVersion field value if set, zero value otherwise.
func (o *SettingsObject) GetSchemaVersion() string {
	if o == nil || IsNil(o.SchemaVersion) {
		var ret string
		return ret
	}
	return *o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetSchemaVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SchemaVersion) {
		return nil, false
	}
	return o.SchemaVersion, true
}

// HasSchemaVersion returns a boolean if a field has been set.
func (o *SettingsObject) HasSchemaVersion() bool {
	if o != nil && !IsNil(o.SchemaVersion) {
		return true
	}

	return false
}

// SetSchemaVersion gets a reference to the given string and assigns it to the SchemaVersion field.
func (o *SettingsObject) SetSchemaVersion(v string) {
	o.SchemaVersion = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SettingsObject) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SettingsObject) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *SettingsObject) SetScope(v string) {
	o.Scope = &v
}

// GetSearchSummary returns the SearchSummary field value if set, zero value otherwise.
func (o *SettingsObject) GetSearchSummary() string {
	if o == nil || IsNil(o.SearchSummary) {
		var ret string
		return ret
	}
	return *o.SearchSummary
}

// GetSearchSummaryOk returns a tuple with the SearchSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetSearchSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.SearchSummary) {
		return nil, false
	}
	return o.SearchSummary, true
}

// HasSearchSummary returns a boolean if a field has been set.
func (o *SettingsObject) HasSearchSummary() bool {
	if o != nil && !IsNil(o.SearchSummary) {
		return true
	}

	return false
}

// SetSearchSummary gets a reference to the given string and assigns it to the SearchSummary field.
func (o *SettingsObject) SetSearchSummary(v string) {
	o.SearchSummary = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *SettingsObject) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *SettingsObject) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *SettingsObject) SetSummary(v string) {
	o.Summary = &v
}

// GetUpdateToken returns the UpdateToken field value if set, zero value otherwise.
func (o *SettingsObject) GetUpdateToken() string {
	if o == nil || IsNil(o.UpdateToken) {
		var ret string
		return ret
	}
	return *o.UpdateToken
}

// GetUpdateTokenOk returns a tuple with the UpdateToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetUpdateTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateToken) {
		return nil, false
	}
	return o.UpdateToken, true
}

// HasUpdateToken returns a boolean if a field has been set.
func (o *SettingsObject) HasUpdateToken() bool {
	if o != nil && !IsNil(o.UpdateToken) {
		return true
	}

	return false
}

// SetUpdateToken gets a reference to the given string and assigns it to the UpdateToken field.
func (o *SettingsObject) SetUpdateToken(v string) {
	o.UpdateToken = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SettingsObject) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsObject) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SettingsObject) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *SettingsObject) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o SettingsObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.ModificationInfo) {
		toSerialize["modificationInfo"] = o.ModificationInfo
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !IsNil(o.ObjectId) {
		toSerialize["objectId"] = o.ObjectId
	}
	if !IsNil(o.SchemaId) {
		toSerialize["schemaId"] = o.SchemaId
	}
	if !IsNil(o.SchemaVersion) {
		toSerialize["schemaVersion"] = o.SchemaVersion
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.SearchSummary) {
		toSerialize["searchSummary"] = o.SearchSummary
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.UpdateToken) {
		toSerialize["updateToken"] = o.UpdateToken
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSettingsObject struct {
	value *SettingsObject
	isSet bool
}

func (v NullableSettingsObject) Get() *SettingsObject {
	return v.value
}

func (v *NullableSettingsObject) Set(val *SettingsObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsObject(val *SettingsObject) *NullableSettingsObject {
	return &NullableSettingsObject{value: val, isSet: true}
}

func (v NullableSettingsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


