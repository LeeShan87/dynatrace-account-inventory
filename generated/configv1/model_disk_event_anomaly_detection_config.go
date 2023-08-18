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

// checks if the DiskEventAnomalyDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskEventAnomalyDetectionConfig{}

// DiskEventAnomalyDetectionConfig struct for DiskEventAnomalyDetectionConfig
type DiskEventAnomalyDetectionConfig struct {
	DiskNameFilter *DiskNameFilter `json:"diskNameFilter,omitempty"`
	// Disk event rule enabled/disabled.
	Enabled bool `json:"enabled"`
	// Narrows the rule usage down to disks that run on hosts that themselves run on the specified host group.
	HostGroupId *string `json:"hostGroupId,omitempty"`
	// The ID of the disk event rule.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The metric to monitor.
	Metric string `json:"metric"`
	// The name of the disk event rule.
	Name string `json:"name"`
	// The number of samples to evaluate.
	Samples int32 `json:"samples"`
	// Narrows the rule usage down to the hosts matching the specified tags.
	TagFilters []TagFilter `json:"tagFilters,omitempty"`
	// The threshold to trigger disk event.    * A percentage for `LowDiskSpace` or `LowInodes` metrics.   * In milliseconds for `ReadTimeExceeding` or `WriteTimeExceeding` metrics.
	Threshold float64 `json:"threshold"`
	// The number of samples that must violate the threshold to trigger an event. Must not exceed the number of evaluated samples.
	ViolatingSamples int32 `json:"violatingSamples"`
}

// NewDiskEventAnomalyDetectionConfig instantiates a new DiskEventAnomalyDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskEventAnomalyDetectionConfig(enabled bool, metric string, name string, samples int32, threshold float64, violatingSamples int32) *DiskEventAnomalyDetectionConfig {
	this := DiskEventAnomalyDetectionConfig{}
	this.Enabled = enabled
	this.Metric = metric
	this.Name = name
	this.Samples = samples
	this.Threshold = threshold
	this.ViolatingSamples = violatingSamples
	return &this
}

// NewDiskEventAnomalyDetectionConfigWithDefaults instantiates a new DiskEventAnomalyDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskEventAnomalyDetectionConfigWithDefaults() *DiskEventAnomalyDetectionConfig {
	this := DiskEventAnomalyDetectionConfig{}
	return &this
}

// GetDiskNameFilter returns the DiskNameFilter field value if set, zero value otherwise.
func (o *DiskEventAnomalyDetectionConfig) GetDiskNameFilter() DiskNameFilter {
	if o == nil || IsNil(o.DiskNameFilter) {
		var ret DiskNameFilter
		return ret
	}
	return *o.DiskNameFilter
}

// GetDiskNameFilterOk returns a tuple with the DiskNameFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetDiskNameFilterOk() (*DiskNameFilter, bool) {
	if o == nil || IsNil(o.DiskNameFilter) {
		return nil, false
	}
	return o.DiskNameFilter, true
}

// HasDiskNameFilter returns a boolean if a field has been set.
func (o *DiskEventAnomalyDetectionConfig) HasDiskNameFilter() bool {
	if o != nil && !IsNil(o.DiskNameFilter) {
		return true
	}

	return false
}

// SetDiskNameFilter gets a reference to the given DiskNameFilter and assigns it to the DiskNameFilter field.
func (o *DiskEventAnomalyDetectionConfig) SetDiskNameFilter(v DiskNameFilter) {
	o.DiskNameFilter = &v
}

// GetEnabled returns the Enabled field value
func (o *DiskEventAnomalyDetectionConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DiskEventAnomalyDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHostGroupId returns the HostGroupId field value if set, zero value otherwise.
func (o *DiskEventAnomalyDetectionConfig) GetHostGroupId() string {
	if o == nil || IsNil(o.HostGroupId) {
		var ret string
		return ret
	}
	return *o.HostGroupId
}

// GetHostGroupIdOk returns a tuple with the HostGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetHostGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostGroupId) {
		return nil, false
	}
	return o.HostGroupId, true
}

// HasHostGroupId returns a boolean if a field has been set.
func (o *DiskEventAnomalyDetectionConfig) HasHostGroupId() bool {
	if o != nil && !IsNil(o.HostGroupId) {
		return true
	}

	return false
}

// SetHostGroupId gets a reference to the given string and assigns it to the HostGroupId field.
func (o *DiskEventAnomalyDetectionConfig) SetHostGroupId(v string) {
	o.HostGroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DiskEventAnomalyDetectionConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskEventAnomalyDetectionConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiskEventAnomalyDetectionConfig) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DiskEventAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DiskEventAnomalyDetectionConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *DiskEventAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetMetric returns the Metric field value
func (o *DiskEventAnomalyDetectionConfig) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *DiskEventAnomalyDetectionConfig) SetMetric(v string) {
	o.Metric = v
}

// GetName returns the Name field value
func (o *DiskEventAnomalyDetectionConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DiskEventAnomalyDetectionConfig) SetName(v string) {
	o.Name = v
}

// GetSamples returns the Samples field value
func (o *DiskEventAnomalyDetectionConfig) GetSamples() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Samples
}

// GetSamplesOk returns a tuple with the Samples field value
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetSamplesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Samples, true
}

// SetSamples sets field value
func (o *DiskEventAnomalyDetectionConfig) SetSamples(v int32) {
	o.Samples = v
}

// GetTagFilters returns the TagFilters field value if set, zero value otherwise.
func (o *DiskEventAnomalyDetectionConfig) GetTagFilters() []TagFilter {
	if o == nil || IsNil(o.TagFilters) {
		var ret []TagFilter
		return ret
	}
	return o.TagFilters
}

// GetTagFiltersOk returns a tuple with the TagFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetTagFiltersOk() ([]TagFilter, bool) {
	if o == nil || IsNil(o.TagFilters) {
		return nil, false
	}
	return o.TagFilters, true
}

// HasTagFilters returns a boolean if a field has been set.
func (o *DiskEventAnomalyDetectionConfig) HasTagFilters() bool {
	if o != nil && !IsNil(o.TagFilters) {
		return true
	}

	return false
}

// SetTagFilters gets a reference to the given []TagFilter and assigns it to the TagFilters field.
func (o *DiskEventAnomalyDetectionConfig) SetTagFilters(v []TagFilter) {
	o.TagFilters = v
}

// GetThreshold returns the Threshold field value
func (o *DiskEventAnomalyDetectionConfig) GetThreshold() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetThresholdOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *DiskEventAnomalyDetectionConfig) SetThreshold(v float64) {
	o.Threshold = v
}

// GetViolatingSamples returns the ViolatingSamples field value
func (o *DiskEventAnomalyDetectionConfig) GetViolatingSamples() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ViolatingSamples
}

// GetViolatingSamplesOk returns a tuple with the ViolatingSamples field value
// and a boolean to check if the value has been set.
func (o *DiskEventAnomalyDetectionConfig) GetViolatingSamplesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViolatingSamples, true
}

// SetViolatingSamples sets field value
func (o *DiskEventAnomalyDetectionConfig) SetViolatingSamples(v int32) {
	o.ViolatingSamples = v
}

func (o DiskEventAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskEventAnomalyDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiskNameFilter) {
		toSerialize["diskNameFilter"] = o.DiskNameFilter
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.HostGroupId) {
		toSerialize["hostGroupId"] = o.HostGroupId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["metric"] = o.Metric
	toSerialize["name"] = o.Name
	toSerialize["samples"] = o.Samples
	if !IsNil(o.TagFilters) {
		toSerialize["tagFilters"] = o.TagFilters
	}
	toSerialize["threshold"] = o.Threshold
	toSerialize["violatingSamples"] = o.ViolatingSamples
	return toSerialize, nil
}

type NullableDiskEventAnomalyDetectionConfig struct {
	value *DiskEventAnomalyDetectionConfig
	isSet bool
}

func (v NullableDiskEventAnomalyDetectionConfig) Get() *DiskEventAnomalyDetectionConfig {
	return v.value
}

func (v *NullableDiskEventAnomalyDetectionConfig) Set(val *DiskEventAnomalyDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskEventAnomalyDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskEventAnomalyDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskEventAnomalyDetectionConfig(val *DiskEventAnomalyDetectionConfig) *NullableDiskEventAnomalyDetectionConfig {
	return &NullableDiskEventAnomalyDetectionConfig{value: val, isSet: true}
}

func (v NullableDiskEventAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskEventAnomalyDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

