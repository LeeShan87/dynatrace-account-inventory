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

// checks if the Plugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Plugin{}

// Plugin General configuration of a plugin.
type Plugin struct {
	// The ID of the plugin, for example `custom.remote.python.demo`.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The metric group of the plugin. All the metrics, captured by the plugin are children of this group.
	MetricGroup *string `json:"metricGroup,omitempty"`
	// The name of the plugin, displayed in Dynatrace.
	Name *string `json:"name,omitempty"`
	// A list of plugin properties.
	Properties []PluginProperty `json:"properties,omitempty"`
	// The type of the plugin. It indicates the runtime environment of the plugin (for example, ActiveGate).
	Type *string `json:"type,omitempty"`
	// The version of the plugin, displayed in Dynatrace.
	Version *string `json:"version,omitempty"`
}

// NewPlugin instantiates a new Plugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlugin() *Plugin {
	this := Plugin{}
	return &this
}

// NewPluginWithDefaults instantiates a new Plugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginWithDefaults() *Plugin {
	this := Plugin{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Plugin) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Plugin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Plugin) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Plugin) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Plugin) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *Plugin) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetMetricGroup returns the MetricGroup field value if set, zero value otherwise.
func (o *Plugin) GetMetricGroup() string {
	if o == nil || IsNil(o.MetricGroup) {
		var ret string
		return ret
	}
	return *o.MetricGroup
}

// GetMetricGroupOk returns a tuple with the MetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetMetricGroupOk() (*string, bool) {
	if o == nil || IsNil(o.MetricGroup) {
		return nil, false
	}
	return o.MetricGroup, true
}

// HasMetricGroup returns a boolean if a field has been set.
func (o *Plugin) HasMetricGroup() bool {
	if o != nil && !IsNil(o.MetricGroup) {
		return true
	}

	return false
}

// SetMetricGroup gets a reference to the given string and assigns it to the MetricGroup field.
func (o *Plugin) SetMetricGroup(v string) {
	o.MetricGroup = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Plugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Plugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Plugin) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Plugin) GetProperties() []PluginProperty {
	if o == nil || IsNil(o.Properties) {
		var ret []PluginProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetPropertiesOk() ([]PluginProperty, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Plugin) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []PluginProperty and assigns it to the Properties field.
func (o *Plugin) SetProperties(v []PluginProperty) {
	o.Properties = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Plugin) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Plugin) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Plugin) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Plugin) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plugin) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Plugin) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Plugin) SetVersion(v string) {
	o.Version = &v
}

func (o Plugin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Plugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MetricGroup) {
		toSerialize["metricGroup"] = o.MetricGroup
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePlugin struct {
	value *Plugin
	isSet bool
}

func (v NullablePlugin) Get() *Plugin {
	return v.value
}

func (v *NullablePlugin) Set(val *Plugin) {
	v.value = val
	v.isSet = true
}

func (v NullablePlugin) IsSet() bool {
	return v.isSet
}

func (v *NullablePlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlugin(val *Plugin) *NullablePlugin {
	return &NullablePlugin{value: val, isSet: true}
}

func (v NullablePlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


