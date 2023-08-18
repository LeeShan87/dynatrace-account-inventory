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

// checks if the AppIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppIdentifier{}

// AppIdentifier The identification info of the application to which the file belongs to.
type AppIdentifier struct {
	// The ID of the mobile app.
	Id *string `json:"id,omitempty"`
	// The operating system the file belongs to.
	Os *string `json:"os,omitempty"`
	// The bundleId (iOS) or package name (Android) the file belongs to.
	PackageName *string `json:"packageName,omitempty"`
	// The version code (Android) / bundle version (iOS) the file belongs to.
	VersionCode string `json:"versionCode"`
	// The version name (Android) / bundle versions string (iOS) the file belongs to.
	VersionName string `json:"versionName"`
}

// NewAppIdentifier instantiates a new AppIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppIdentifier(versionCode string, versionName string) *AppIdentifier {
	this := AppIdentifier{}
	this.VersionCode = versionCode
	this.VersionName = versionName
	return &this
}

// NewAppIdentifierWithDefaults instantiates a new AppIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppIdentifierWithDefaults() *AppIdentifier {
	this := AppIdentifier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppIdentifier) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppIdentifier) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppIdentifier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppIdentifier) SetId(v string) {
	o.Id = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *AppIdentifier) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppIdentifier) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *AppIdentifier) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *AppIdentifier) SetOs(v string) {
	o.Os = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *AppIdentifier) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppIdentifier) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *AppIdentifier) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *AppIdentifier) SetPackageName(v string) {
	o.PackageName = &v
}

// GetVersionCode returns the VersionCode field value
func (o *AppIdentifier) GetVersionCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionCode
}

// GetVersionCodeOk returns a tuple with the VersionCode field value
// and a boolean to check if the value has been set.
func (o *AppIdentifier) GetVersionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionCode, true
}

// SetVersionCode sets field value
func (o *AppIdentifier) SetVersionCode(v string) {
	o.VersionCode = v
}

// GetVersionName returns the VersionName field value
func (o *AppIdentifier) GetVersionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value
// and a boolean to check if the value has been set.
func (o *AppIdentifier) GetVersionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionName, true
}

// SetVersionName sets field value
func (o *AppIdentifier) SetVersionName(v string) {
	o.VersionName = v
}

func (o AppIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.PackageName) {
		toSerialize["packageName"] = o.PackageName
	}
	toSerialize["versionCode"] = o.VersionCode
	toSerialize["versionName"] = o.VersionName
	return toSerialize, nil
}

type NullableAppIdentifier struct {
	value *AppIdentifier
	isSet bool
}

func (v NullableAppIdentifier) Get() *AppIdentifier {
	return v.value
}

func (v *NullableAppIdentifier) Set(val *AppIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableAppIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableAppIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppIdentifier(val *AppIdentifier) *NullableAppIdentifier {
	return &NullableAppIdentifier{value: val, isSet: true}
}

func (v NullableAppIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

