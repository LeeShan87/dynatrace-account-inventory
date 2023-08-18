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

// checks if the DataPrivacyAndSecurity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPrivacyAndSecurity{}

// DataPrivacyAndSecurity Global configuration for data privacy and security.
type DataPrivacyAndSecurity struct {
	// The audit logging is enabled (`true`) or disabled (`false`).
	LogAuditEvents *bool `json:"logAuditEvents,omitempty"`
	// Masking of IP addresses and GPS coordinates is enabled (`true`) or disabled (`false`).
	MaskIpAddressesAndGpsCoordinates bool `json:"maskIpAddressesAndGpsCoordinates"`
	// Masking of personal data in URIs is enabled (`true`) or disabled (`false`).
	MaskPersonalDataInUris bool `json:"maskPersonalDataInUris"`
	// Masking of user action names is enabled (`true`) or disabled (`false`).   This masking is available only for web applications.
	MaskUserActionNames bool `json:"maskUserActionNames"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
}

// NewDataPrivacyAndSecurity instantiates a new DataPrivacyAndSecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPrivacyAndSecurity(maskIpAddressesAndGpsCoordinates bool, maskPersonalDataInUris bool, maskUserActionNames bool) *DataPrivacyAndSecurity {
	this := DataPrivacyAndSecurity{}
	this.MaskIpAddressesAndGpsCoordinates = maskIpAddressesAndGpsCoordinates
	this.MaskPersonalDataInUris = maskPersonalDataInUris
	this.MaskUserActionNames = maskUserActionNames
	return &this
}

// NewDataPrivacyAndSecurityWithDefaults instantiates a new DataPrivacyAndSecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPrivacyAndSecurityWithDefaults() *DataPrivacyAndSecurity {
	this := DataPrivacyAndSecurity{}
	return &this
}

// GetLogAuditEvents returns the LogAuditEvents field value if set, zero value otherwise.
func (o *DataPrivacyAndSecurity) GetLogAuditEvents() bool {
	if o == nil || IsNil(o.LogAuditEvents) {
		var ret bool
		return ret
	}
	return *o.LogAuditEvents
}

// GetLogAuditEventsOk returns a tuple with the LogAuditEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPrivacyAndSecurity) GetLogAuditEventsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogAuditEvents) {
		return nil, false
	}
	return o.LogAuditEvents, true
}

// HasLogAuditEvents returns a boolean if a field has been set.
func (o *DataPrivacyAndSecurity) HasLogAuditEvents() bool {
	if o != nil && !IsNil(o.LogAuditEvents) {
		return true
	}

	return false
}

// SetLogAuditEvents gets a reference to the given bool and assigns it to the LogAuditEvents field.
func (o *DataPrivacyAndSecurity) SetLogAuditEvents(v bool) {
	o.LogAuditEvents = &v
}

// GetMaskIpAddressesAndGpsCoordinates returns the MaskIpAddressesAndGpsCoordinates field value
func (o *DataPrivacyAndSecurity) GetMaskIpAddressesAndGpsCoordinates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MaskIpAddressesAndGpsCoordinates
}

// GetMaskIpAddressesAndGpsCoordinatesOk returns a tuple with the MaskIpAddressesAndGpsCoordinates field value
// and a boolean to check if the value has been set.
func (o *DataPrivacyAndSecurity) GetMaskIpAddressesAndGpsCoordinatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskIpAddressesAndGpsCoordinates, true
}

// SetMaskIpAddressesAndGpsCoordinates sets field value
func (o *DataPrivacyAndSecurity) SetMaskIpAddressesAndGpsCoordinates(v bool) {
	o.MaskIpAddressesAndGpsCoordinates = v
}

// GetMaskPersonalDataInUris returns the MaskPersonalDataInUris field value
func (o *DataPrivacyAndSecurity) GetMaskPersonalDataInUris() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MaskPersonalDataInUris
}

// GetMaskPersonalDataInUrisOk returns a tuple with the MaskPersonalDataInUris field value
// and a boolean to check if the value has been set.
func (o *DataPrivacyAndSecurity) GetMaskPersonalDataInUrisOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskPersonalDataInUris, true
}

// SetMaskPersonalDataInUris sets field value
func (o *DataPrivacyAndSecurity) SetMaskPersonalDataInUris(v bool) {
	o.MaskPersonalDataInUris = v
}

// GetMaskUserActionNames returns the MaskUserActionNames field value
func (o *DataPrivacyAndSecurity) GetMaskUserActionNames() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MaskUserActionNames
}

// GetMaskUserActionNamesOk returns a tuple with the MaskUserActionNames field value
// and a boolean to check if the value has been set.
func (o *DataPrivacyAndSecurity) GetMaskUserActionNamesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskUserActionNames, true
}

// SetMaskUserActionNames sets field value
func (o *DataPrivacyAndSecurity) SetMaskUserActionNames(v bool) {
	o.MaskUserActionNames = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DataPrivacyAndSecurity) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPrivacyAndSecurity) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DataPrivacyAndSecurity) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *DataPrivacyAndSecurity) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

func (o DataPrivacyAndSecurity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPrivacyAndSecurity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LogAuditEvents) {
		toSerialize["logAuditEvents"] = o.LogAuditEvents
	}
	toSerialize["maskIpAddressesAndGpsCoordinates"] = o.MaskIpAddressesAndGpsCoordinates
	toSerialize["maskPersonalDataInUris"] = o.MaskPersonalDataInUris
	toSerialize["maskUserActionNames"] = o.MaskUserActionNames
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableDataPrivacyAndSecurity struct {
	value *DataPrivacyAndSecurity
	isSet bool
}

func (v NullableDataPrivacyAndSecurity) Get() *DataPrivacyAndSecurity {
	return v.value
}

func (v *NullableDataPrivacyAndSecurity) Set(val *DataPrivacyAndSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPrivacyAndSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPrivacyAndSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPrivacyAndSecurity(val *DataPrivacyAndSecurity) *NullableDataPrivacyAndSecurity {
	return &NullableDataPrivacyAndSecurity{value: val, isSet: true}
}

func (v NullableDataPrivacyAndSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPrivacyAndSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


