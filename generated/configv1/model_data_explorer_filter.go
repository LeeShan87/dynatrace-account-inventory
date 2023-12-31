/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the DataExplorerFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataExplorerFilter{}

// DataExplorerFilter Filter for data explorer queries.
type DataExplorerFilter struct {
	Criteria *[]DexpFilterCriterion `json:"criteria,omitempty"`
	EntityAttribute *string `json:"entityAttribute,omitempty"`
	Filter *string `json:"filter,omitempty"`
	FilterOperator *string `json:"filterOperator,omitempty"`
	FilterType *string `json:"filterType,omitempty"`
	GlobalEntity *string `json:"globalEntity,omitempty"`
	NestedFilters *[]DataExplorerFilter `json:"nestedFilters,omitempty"`
	Relationship *DexpFilterRelationship `json:"relationship,omitempty"`
}

// NewDataExplorerFilter instantiates a new DataExplorerFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataExplorerFilter() *DataExplorerFilter {
	this := DataExplorerFilter{}
	return &this
}

// NewDataExplorerFilterWithDefaults instantiates a new DataExplorerFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataExplorerFilterWithDefaults() *DataExplorerFilter {
	this := DataExplorerFilter{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetCriteria() []DexpFilterCriterion {
	if o == nil || IsNil(o.Criteria) {
		var ret []DexpFilterCriterion
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetCriteriaOk() (*[]DexpFilterCriterion, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []DexpFilterCriterion and assigns it to the Criteria field.
func (o *DataExplorerFilter) SetCriteria(v []DexpFilterCriterion) {
	o.Criteria = &v
}

// GetEntityAttribute returns the EntityAttribute field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetEntityAttribute() string {
	if o == nil || IsNil(o.EntityAttribute) {
		var ret string
		return ret
	}
	return *o.EntityAttribute
}

// GetEntityAttributeOk returns a tuple with the EntityAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetEntityAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityAttribute) {
		return nil, false
	}
	return o.EntityAttribute, true
}

// HasEntityAttribute returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasEntityAttribute() bool {
	if o != nil && !IsNil(o.EntityAttribute) {
		return true
	}

	return false
}

// SetEntityAttribute gets a reference to the given string and assigns it to the EntityAttribute field.
func (o *DataExplorerFilter) SetEntityAttribute(v string) {
	o.EntityAttribute = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *DataExplorerFilter) SetFilter(v string) {
	o.Filter = &v
}

// GetFilterOperator returns the FilterOperator field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetFilterOperator() string {
	if o == nil || IsNil(o.FilterOperator) {
		var ret string
		return ret
	}
	return *o.FilterOperator
}

// GetFilterOperatorOk returns a tuple with the FilterOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetFilterOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.FilterOperator) {
		return nil, false
	}
	return o.FilterOperator, true
}

// HasFilterOperator returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasFilterOperator() bool {
	if o != nil && !IsNil(o.FilterOperator) {
		return true
	}

	return false
}

// SetFilterOperator gets a reference to the given string and assigns it to the FilterOperator field.
func (o *DataExplorerFilter) SetFilterOperator(v string) {
	o.FilterOperator = &v
}

// GetFilterType returns the FilterType field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetFilterType() string {
	if o == nil || IsNil(o.FilterType) {
		var ret string
		return ret
	}
	return *o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetFilterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FilterType) {
		return nil, false
	}
	return o.FilterType, true
}

// HasFilterType returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasFilterType() bool {
	if o != nil && !IsNil(o.FilterType) {
		return true
	}

	return false
}

// SetFilterType gets a reference to the given string and assigns it to the FilterType field.
func (o *DataExplorerFilter) SetFilterType(v string) {
	o.FilterType = &v
}

// GetGlobalEntity returns the GlobalEntity field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetGlobalEntity() string {
	if o == nil || IsNil(o.GlobalEntity) {
		var ret string
		return ret
	}
	return *o.GlobalEntity
}

// GetGlobalEntityOk returns a tuple with the GlobalEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetGlobalEntityOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalEntity) {
		return nil, false
	}
	return o.GlobalEntity, true
}

// HasGlobalEntity returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasGlobalEntity() bool {
	if o != nil && !IsNil(o.GlobalEntity) {
		return true
	}

	return false
}

// SetGlobalEntity gets a reference to the given string and assigns it to the GlobalEntity field.
func (o *DataExplorerFilter) SetGlobalEntity(v string) {
	o.GlobalEntity = &v
}

// GetNestedFilters returns the NestedFilters field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetNestedFilters() []DataExplorerFilter {
	if o == nil || IsNil(o.NestedFilters) {
		var ret []DataExplorerFilter
		return ret
	}
	return *o.NestedFilters
}

// GetNestedFiltersOk returns a tuple with the NestedFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetNestedFiltersOk() (*[]DataExplorerFilter, bool) {
	if o == nil || IsNil(o.NestedFilters) {
		return nil, false
	}
	return o.NestedFilters, true
}

// HasNestedFilters returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasNestedFilters() bool {
	if o != nil && !IsNil(o.NestedFilters) {
		return true
	}

	return false
}

// SetNestedFilters gets a reference to the given []DataExplorerFilter and assigns it to the NestedFilters field.
func (o *DataExplorerFilter) SetNestedFilters(v []DataExplorerFilter) {
	o.NestedFilters = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *DataExplorerFilter) GetRelationship() DexpFilterRelationship {
	if o == nil || IsNil(o.Relationship) {
		var ret DexpFilterRelationship
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExplorerFilter) GetRelationshipOk() (*DexpFilterRelationship, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *DataExplorerFilter) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given DexpFilterRelationship and assigns it to the Relationship field.
func (o *DataExplorerFilter) SetRelationship(v DexpFilterRelationship) {
	o.Relationship = &v
}

func (o DataExplorerFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataExplorerFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.EntityAttribute) {
		toSerialize["entityAttribute"] = o.EntityAttribute
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.FilterOperator) {
		toSerialize["filterOperator"] = o.FilterOperator
	}
	if !IsNil(o.FilterType) {
		toSerialize["filterType"] = o.FilterType
	}
	if !IsNil(o.GlobalEntity) {
		toSerialize["globalEntity"] = o.GlobalEntity
	}
	if !IsNil(o.NestedFilters) {
		toSerialize["nestedFilters"] = o.NestedFilters
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	return toSerialize, nil
}

type NullableDataExplorerFilter struct {
	value *DataExplorerFilter
	isSet bool
}

func (v NullableDataExplorerFilter) Get() *DataExplorerFilter {
	return v.value
}

func (v *NullableDataExplorerFilter) Set(val *DataExplorerFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDataExplorerFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDataExplorerFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataExplorerFilter(val *DataExplorerFilter) *NullableDataExplorerFilter {
	return &NullableDataExplorerFilter{value: val, isSet: true}
}

func (v NullableDataExplorerFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataExplorerFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


