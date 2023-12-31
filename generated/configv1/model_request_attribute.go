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

// checks if the RequestAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestAttribute{}

// RequestAttribute struct for RequestAttribute
type RequestAttribute struct {
	// Aggregation type for the request values.
	Aggregation string `json:"aggregation"`
	// Confidential data flag. Set `true` to treat the captured data as confidential.
	Confidential bool `json:"confidential"`
	// The list of data sources.
	DataSources []DataSource `json:"dataSources"`
	// The data type of the request attribute.
	DataType string `json:"dataType"`
	// The request attribute is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The ID of the request attribute.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The name of the request attribute.
	Name string `json:"name"`
	// String values transformation.    If the **dataType** is not `string`, set the `Original` here.
	Normalization string `json:"normalization"`
	// Personal data masking flag. Set `true` to skip masking.    Warning: This will potentially access personalized data.
	SkipPersonalDataMasking bool `json:"skipPersonalDataMasking"`
}

// NewRequestAttribute instantiates a new RequestAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestAttribute(aggregation string, confidential bool, dataSources []DataSource, dataType string, enabled bool, name string, normalization string, skipPersonalDataMasking bool) *RequestAttribute {
	this := RequestAttribute{}
	this.Aggregation = aggregation
	this.Confidential = confidential
	this.DataSources = dataSources
	this.DataType = dataType
	this.Enabled = enabled
	this.Name = name
	this.Normalization = normalization
	this.SkipPersonalDataMasking = skipPersonalDataMasking
	return &this
}

// NewRequestAttributeWithDefaults instantiates a new RequestAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestAttributeWithDefaults() *RequestAttribute {
	this := RequestAttribute{}
	return &this
}

// GetAggregation returns the Aggregation field value
func (o *RequestAttribute) GetAggregation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetAggregationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *RequestAttribute) SetAggregation(v string) {
	o.Aggregation = v
}

// GetConfidential returns the Confidential field value
func (o *RequestAttribute) GetConfidential() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Confidential
}

// GetConfidentialOk returns a tuple with the Confidential field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetConfidentialOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidential, true
}

// SetConfidential sets field value
func (o *RequestAttribute) SetConfidential(v bool) {
	o.Confidential = v
}

// GetDataSources returns the DataSources field value
func (o *RequestAttribute) GetDataSources() []DataSource {
	if o == nil {
		var ret []DataSource
		return ret
	}

	return o.DataSources
}

// GetDataSourcesOk returns a tuple with the DataSources field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetDataSourcesOk() ([]DataSource, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataSources, true
}

// SetDataSources sets field value
func (o *RequestAttribute) SetDataSources(v []DataSource) {
	o.DataSources = v
}

// GetDataType returns the DataType field value
func (o *RequestAttribute) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *RequestAttribute) SetDataType(v string) {
	o.DataType = v
}

// GetEnabled returns the Enabled field value
func (o *RequestAttribute) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RequestAttribute) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequestAttribute) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequestAttribute) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RequestAttribute) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RequestAttribute) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RequestAttribute) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *RequestAttribute) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value
func (o *RequestAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RequestAttribute) SetName(v string) {
	o.Name = v
}

// GetNormalization returns the Normalization field value
func (o *RequestAttribute) GetNormalization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Normalization
}

// GetNormalizationOk returns a tuple with the Normalization field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetNormalizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Normalization, true
}

// SetNormalization sets field value
func (o *RequestAttribute) SetNormalization(v string) {
	o.Normalization = v
}

// GetSkipPersonalDataMasking returns the SkipPersonalDataMasking field value
func (o *RequestAttribute) GetSkipPersonalDataMasking() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SkipPersonalDataMasking
}

// GetSkipPersonalDataMaskingOk returns a tuple with the SkipPersonalDataMasking field value
// and a boolean to check if the value has been set.
func (o *RequestAttribute) GetSkipPersonalDataMaskingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SkipPersonalDataMasking, true
}

// SetSkipPersonalDataMasking sets field value
func (o *RequestAttribute) SetSkipPersonalDataMasking(v bool) {
	o.SkipPersonalDataMasking = v
}

func (o RequestAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["confidential"] = o.Confidential
	toSerialize["dataSources"] = o.DataSources
	toSerialize["dataType"] = o.DataType
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	toSerialize["normalization"] = o.Normalization
	toSerialize["skipPersonalDataMasking"] = o.SkipPersonalDataMasking
	return toSerialize, nil
}

type NullableRequestAttribute struct {
	value *RequestAttribute
	isSet bool
}

func (v NullableRequestAttribute) Get() *RequestAttribute {
	return v.value
}

func (v *NullableRequestAttribute) Set(val *RequestAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestAttribute(val *RequestAttribute) *NullableRequestAttribute {
	return &NullableRequestAttribute{value: val, isSet: true}
}

func (v NullableRequestAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


