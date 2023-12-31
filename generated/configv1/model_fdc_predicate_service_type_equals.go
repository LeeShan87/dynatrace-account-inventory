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

// checks if the FdcPredicateServiceTypeEquals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateServiceTypeEquals{}

// FdcPredicateServiceTypeEquals The predicate of the `SERVICE_TYPE_EQUALS` type. It checks whether the attribute (which is the type of the service) equals one of the reference values.
type FdcPredicateServiceTypeEquals struct {
	// A list of reference values. The condition is fulfilled when the attribute (which is the type of the service) equals *any* of these.The possible values are: WebRequest, WebService, Database, Method, WebSite, Messaging, Mobile, Process, Rmi, External, QueueListener, QueueInteraction, RemoteCall, SaasVendor, AMP, CustomApplication, Cics, Ims, CicsInteraction, ImsInteraction, EnterpriseServiceBus, ZosConnect.
	Values []string `json:"values,omitempty"`
}

// NewFdcPredicateServiceTypeEquals instantiates a new FdcPredicateServiceTypeEquals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateServiceTypeEquals(type_ string) *FdcPredicateServiceTypeEquals {
	this := FdcPredicateServiceTypeEquals{}
	this.Type = type_
	return &this
}

// NewFdcPredicateServiceTypeEqualsWithDefaults instantiates a new FdcPredicateServiceTypeEquals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateServiceTypeEqualsWithDefaults() *FdcPredicateServiceTypeEquals {
	this := FdcPredicateServiceTypeEquals{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FdcPredicateServiceTypeEquals) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdcPredicateServiceTypeEquals) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FdcPredicateServiceTypeEquals) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *FdcPredicateServiceTypeEquals) SetValues(v []string) {
	o.Values = v
}

func (o FdcPredicateServiceTypeEquals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateServiceTypeEquals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableFdcPredicateServiceTypeEquals struct {
	value *FdcPredicateServiceTypeEquals
	isSet bool
}

func (v NullableFdcPredicateServiceTypeEquals) Get() *FdcPredicateServiceTypeEquals {
	return v.value
}

func (v *NullableFdcPredicateServiceTypeEquals) Set(val *FdcPredicateServiceTypeEquals) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateServiceTypeEquals) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateServiceTypeEquals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateServiceTypeEquals(val *FdcPredicateServiceTypeEquals) *NullableFdcPredicateServiceTypeEquals {
	return &NullableFdcPredicateServiceTypeEquals{value: val, isSet: true}
}

func (v NullableFdcPredicateServiceTypeEquals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateServiceTypeEquals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


