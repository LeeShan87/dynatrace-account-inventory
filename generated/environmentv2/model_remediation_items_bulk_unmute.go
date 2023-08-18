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

// checks if the RemediationItemsBulkUnmute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemediationItemsBulkUnmute{}

// RemediationItemsBulkUnmute Information on un-muting several remediation items.
type RemediationItemsBulkUnmute struct {
	// A comment about the un-muting reason.
	Comment *string `json:"comment,omitempty"`
	// The reason for un-muting the remediation items.
	Reason string `json:"reason"`
	// The ids of the remediation items to be un-muted.
	RemediationItemIds []string `json:"remediationItemIds"`
}

// NewRemediationItemsBulkUnmute instantiates a new RemediationItemsBulkUnmute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationItemsBulkUnmute(reason string, remediationItemIds []string) *RemediationItemsBulkUnmute {
	this := RemediationItemsBulkUnmute{}
	this.Reason = reason
	this.RemediationItemIds = remediationItemIds
	return &this
}

// NewRemediationItemsBulkUnmuteWithDefaults instantiates a new RemediationItemsBulkUnmute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationItemsBulkUnmuteWithDefaults() *RemediationItemsBulkUnmute {
	this := RemediationItemsBulkUnmute{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RemediationItemsBulkUnmute) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemsBulkUnmute) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RemediationItemsBulkUnmute) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RemediationItemsBulkUnmute) SetComment(v string) {
	o.Comment = &v
}

// GetReason returns the Reason field value
func (o *RemediationItemsBulkUnmute) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RemediationItemsBulkUnmute) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RemediationItemsBulkUnmute) SetReason(v string) {
	o.Reason = v
}

// GetRemediationItemIds returns the RemediationItemIds field value
func (o *RemediationItemsBulkUnmute) GetRemediationItemIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemediationItemIds
}

// GetRemediationItemIdsOk returns a tuple with the RemediationItemIds field value
// and a boolean to check if the value has been set.
func (o *RemediationItemsBulkUnmute) GetRemediationItemIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemediationItemIds, true
}

// SetRemediationItemIds sets field value
func (o *RemediationItemsBulkUnmute) SetRemediationItemIds(v []string) {
	o.RemediationItemIds = v
}

func (o RemediationItemsBulkUnmute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediationItemsBulkUnmute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["reason"] = o.Reason
	toSerialize["remediationItemIds"] = o.RemediationItemIds
	return toSerialize, nil
}

type NullableRemediationItemsBulkUnmute struct {
	value *RemediationItemsBulkUnmute
	isSet bool
}

func (v NullableRemediationItemsBulkUnmute) Get() *RemediationItemsBulkUnmute {
	return v.value
}

func (v *NullableRemediationItemsBulkUnmute) Set(val *RemediationItemsBulkUnmute) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationItemsBulkUnmute) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationItemsBulkUnmute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationItemsBulkUnmute(val *RemediationItemsBulkUnmute) *NullableRemediationItemsBulkUnmute {
	return &NullableRemediationItemsBulkUnmute{value: val, isSet: true}
}

func (v NullableRemediationItemsBulkUnmute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationItemsBulkUnmute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

