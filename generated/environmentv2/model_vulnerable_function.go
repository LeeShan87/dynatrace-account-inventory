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

// checks if the VulnerableFunction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VulnerableFunction{}

// VulnerableFunction Defines an vulnerable function.
type VulnerableFunction struct {
	// The class name of the vulnerable function.
	ClassName *string `json:"className,omitempty"`
	// The file path of the vulnerable function.
	FilePath *string `json:"filePath,omitempty"`
	// The function name of the vulnerable function.
	FunctionName *string `json:"functionName,omitempty"`
}

// NewVulnerableFunction instantiates a new VulnerableFunction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerableFunction() *VulnerableFunction {
	this := VulnerableFunction{}
	return &this
}

// NewVulnerableFunctionWithDefaults instantiates a new VulnerableFunction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerableFunctionWithDefaults() *VulnerableFunction {
	this := VulnerableFunction{}
	return &this
}

// GetClassName returns the ClassName field value if set, zero value otherwise.
func (o *VulnerableFunction) GetClassName() string {
	if o == nil || IsNil(o.ClassName) {
		var ret string
		return ret
	}
	return *o.ClassName
}

// GetClassNameOk returns a tuple with the ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerableFunction) GetClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClassName) {
		return nil, false
	}
	return o.ClassName, true
}

// HasClassName returns a boolean if a field has been set.
func (o *VulnerableFunction) HasClassName() bool {
	if o != nil && !IsNil(o.ClassName) {
		return true
	}

	return false
}

// SetClassName gets a reference to the given string and assigns it to the ClassName field.
func (o *VulnerableFunction) SetClassName(v string) {
	o.ClassName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *VulnerableFunction) GetFilePath() string {
	if o == nil || IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerableFunction) GetFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *VulnerableFunction) HasFilePath() bool {
	if o != nil && !IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *VulnerableFunction) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *VulnerableFunction) GetFunctionName() string {
	if o == nil || IsNil(o.FunctionName) {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VulnerableFunction) GetFunctionNameOk() (*string, bool) {
	if o == nil || IsNil(o.FunctionName) {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *VulnerableFunction) HasFunctionName() bool {
	if o != nil && !IsNil(o.FunctionName) {
		return true
	}

	return false
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *VulnerableFunction) SetFunctionName(v string) {
	o.FunctionName = &v
}

func (o VulnerableFunction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VulnerableFunction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClassName) {
		toSerialize["className"] = o.ClassName
	}
	if !IsNil(o.FilePath) {
		toSerialize["filePath"] = o.FilePath
	}
	if !IsNil(o.FunctionName) {
		toSerialize["functionName"] = o.FunctionName
	}
	return toSerialize, nil
}

type NullableVulnerableFunction struct {
	value *VulnerableFunction
	isSet bool
}

func (v NullableVulnerableFunction) Get() *VulnerableFunction {
	return v.value
}

func (v *NullableVulnerableFunction) Set(val *VulnerableFunction) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerableFunction) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerableFunction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerableFunction(val *VulnerableFunction) *NullableVulnerableFunction {
	return &NullableVulnerableFunction{value: val, isSet: true}
}

func (v NullableVulnerableFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerableFunction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


