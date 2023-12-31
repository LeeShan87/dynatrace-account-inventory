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

// checks if the UserActionAndSessionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActionAndSessionProperties{}

// UserActionAndSessionProperties Defines userAction and session custom defined properties settings of an application.
type UserActionAndSessionProperties struct {
	// The aggregation type of the property.     It defines how multiple values of the property are aggregated.
	Aggregation *string `json:"aggregation,omitempty"`
	// The cleanup rule of the property.   Defines how to extract the data you need from a string value. Specify the [regular expression](https://dt-url.net/k9e0iaq) for the data you need there.
	CleanupRule *string `json:"cleanupRule,omitempty"`
	// The display name of the property.
	DisplayName *string `json:"displayName,omitempty"`
	// If true, the value of this property will always be stored in lower case. Defaults to false.
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// Key of the property
	Key string `json:"key"`
	// If the type is LONG_STRING, the max length for this property. Must be a multiple of 100. Defaults to 200.
	LongStringLength *int32 `json:"longStringLength,omitempty"`
	// If the origin is META_DATA, metaData id of the property
	MetadataId *int32 `json:"metadataId,omitempty"`
	// The origin of the property
	Origin string `json:"origin"`
	// The ID of the request attribute.   Only applicable when the **origin** is set to `SERVER_SIDE_REQUEST_ATTRIBUTE`.
	ServerSideRequestAttribute *string `json:"serverSideRequestAttribute,omitempty"`
	// If `true`, the property is stored as a session property
	StoreAsSessionProperty *bool `json:"storeAsSessionProperty,omitempty"`
	// If `true`, the property is stored as a user action property
	StoreAsUserActionProperty *bool `json:"storeAsUserActionProperty,omitempty"`
	// The data type of the property.
	Type string `json:"type"`
	// Unique id among all userTags and properties of this application
	UniqueId int32 `json:"uniqueId"`
}

// NewUserActionAndSessionProperties instantiates a new UserActionAndSessionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActionAndSessionProperties(key string, origin string, type_ string, uniqueId int32) *UserActionAndSessionProperties {
	this := UserActionAndSessionProperties{}
	this.Key = key
	this.Origin = origin
	this.Type = type_
	this.UniqueId = uniqueId
	return &this
}

// NewUserActionAndSessionPropertiesWithDefaults instantiates a new UserActionAndSessionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActionAndSessionPropertiesWithDefaults() *UserActionAndSessionProperties {
	this := UserActionAndSessionProperties{}
	return &this
}

// GetAggregation returns the Aggregation field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetAggregation() string {
	if o == nil || IsNil(o.Aggregation) {
		var ret string
		return ret
	}
	return *o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetAggregationOk() (*string, bool) {
	if o == nil || IsNil(o.Aggregation) {
		return nil, false
	}
	return o.Aggregation, true
}

// HasAggregation returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasAggregation() bool {
	if o != nil && !IsNil(o.Aggregation) {
		return true
	}

	return false
}

// SetAggregation gets a reference to the given string and assigns it to the Aggregation field.
func (o *UserActionAndSessionProperties) SetAggregation(v string) {
	o.Aggregation = &v
}

// GetCleanupRule returns the CleanupRule field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetCleanupRule() string {
	if o == nil || IsNil(o.CleanupRule) {
		var ret string
		return ret
	}
	return *o.CleanupRule
}

// GetCleanupRuleOk returns a tuple with the CleanupRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetCleanupRuleOk() (*string, bool) {
	if o == nil || IsNil(o.CleanupRule) {
		return nil, false
	}
	return o.CleanupRule, true
}

// HasCleanupRule returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasCleanupRule() bool {
	if o != nil && !IsNil(o.CleanupRule) {
		return true
	}

	return false
}

// SetCleanupRule gets a reference to the given string and assigns it to the CleanupRule field.
func (o *UserActionAndSessionProperties) SetCleanupRule(v string) {
	o.CleanupRule = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UserActionAndSessionProperties) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIgnoreCase returns the IgnoreCase field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetIgnoreCase() bool {
	if o == nil || IsNil(o.IgnoreCase) {
		var ret bool
		return ret
	}
	return *o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreCase) {
		return nil, false
	}
	return o.IgnoreCase, true
}

// HasIgnoreCase returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasIgnoreCase() bool {
	if o != nil && !IsNil(o.IgnoreCase) {
		return true
	}

	return false
}

// SetIgnoreCase gets a reference to the given bool and assigns it to the IgnoreCase field.
func (o *UserActionAndSessionProperties) SetIgnoreCase(v bool) {
	o.IgnoreCase = &v
}

// GetKey returns the Key field value
func (o *UserActionAndSessionProperties) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *UserActionAndSessionProperties) SetKey(v string) {
	o.Key = v
}

// GetLongStringLength returns the LongStringLength field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetLongStringLength() int32 {
	if o == nil || IsNil(o.LongStringLength) {
		var ret int32
		return ret
	}
	return *o.LongStringLength
}

// GetLongStringLengthOk returns a tuple with the LongStringLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetLongStringLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.LongStringLength) {
		return nil, false
	}
	return o.LongStringLength, true
}

// HasLongStringLength returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasLongStringLength() bool {
	if o != nil && !IsNil(o.LongStringLength) {
		return true
	}

	return false
}

// SetLongStringLength gets a reference to the given int32 and assigns it to the LongStringLength field.
func (o *UserActionAndSessionProperties) SetLongStringLength(v int32) {
	o.LongStringLength = &v
}

// GetMetadataId returns the MetadataId field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetMetadataId() int32 {
	if o == nil || IsNil(o.MetadataId) {
		var ret int32
		return ret
	}
	return *o.MetadataId
}

// GetMetadataIdOk returns a tuple with the MetadataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetMetadataIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MetadataId) {
		return nil, false
	}
	return o.MetadataId, true
}

// HasMetadataId returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasMetadataId() bool {
	if o != nil && !IsNil(o.MetadataId) {
		return true
	}

	return false
}

// SetMetadataId gets a reference to the given int32 and assigns it to the MetadataId field.
func (o *UserActionAndSessionProperties) SetMetadataId(v int32) {
	o.MetadataId = &v
}

// GetOrigin returns the Origin field value
func (o *UserActionAndSessionProperties) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *UserActionAndSessionProperties) SetOrigin(v string) {
	o.Origin = v
}

// GetServerSideRequestAttribute returns the ServerSideRequestAttribute field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetServerSideRequestAttribute() string {
	if o == nil || IsNil(o.ServerSideRequestAttribute) {
		var ret string
		return ret
	}
	return *o.ServerSideRequestAttribute
}

// GetServerSideRequestAttributeOk returns a tuple with the ServerSideRequestAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetServerSideRequestAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.ServerSideRequestAttribute) {
		return nil, false
	}
	return o.ServerSideRequestAttribute, true
}

// HasServerSideRequestAttribute returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasServerSideRequestAttribute() bool {
	if o != nil && !IsNil(o.ServerSideRequestAttribute) {
		return true
	}

	return false
}

// SetServerSideRequestAttribute gets a reference to the given string and assigns it to the ServerSideRequestAttribute field.
func (o *UserActionAndSessionProperties) SetServerSideRequestAttribute(v string) {
	o.ServerSideRequestAttribute = &v
}

// GetStoreAsSessionProperty returns the StoreAsSessionProperty field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetStoreAsSessionProperty() bool {
	if o == nil || IsNil(o.StoreAsSessionProperty) {
		var ret bool
		return ret
	}
	return *o.StoreAsSessionProperty
}

// GetStoreAsSessionPropertyOk returns a tuple with the StoreAsSessionProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetStoreAsSessionPropertyOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreAsSessionProperty) {
		return nil, false
	}
	return o.StoreAsSessionProperty, true
}

// HasStoreAsSessionProperty returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasStoreAsSessionProperty() bool {
	if o != nil && !IsNil(o.StoreAsSessionProperty) {
		return true
	}

	return false
}

// SetStoreAsSessionProperty gets a reference to the given bool and assigns it to the StoreAsSessionProperty field.
func (o *UserActionAndSessionProperties) SetStoreAsSessionProperty(v bool) {
	o.StoreAsSessionProperty = &v
}

// GetStoreAsUserActionProperty returns the StoreAsUserActionProperty field value if set, zero value otherwise.
func (o *UserActionAndSessionProperties) GetStoreAsUserActionProperty() bool {
	if o == nil || IsNil(o.StoreAsUserActionProperty) {
		var ret bool
		return ret
	}
	return *o.StoreAsUserActionProperty
}

// GetStoreAsUserActionPropertyOk returns a tuple with the StoreAsUserActionProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetStoreAsUserActionPropertyOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreAsUserActionProperty) {
		return nil, false
	}
	return o.StoreAsUserActionProperty, true
}

// HasStoreAsUserActionProperty returns a boolean if a field has been set.
func (o *UserActionAndSessionProperties) HasStoreAsUserActionProperty() bool {
	if o != nil && !IsNil(o.StoreAsUserActionProperty) {
		return true
	}

	return false
}

// SetStoreAsUserActionProperty gets a reference to the given bool and assigns it to the StoreAsUserActionProperty field.
func (o *UserActionAndSessionProperties) SetStoreAsUserActionProperty(v bool) {
	o.StoreAsUserActionProperty = &v
}

// GetType returns the Type field value
func (o *UserActionAndSessionProperties) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserActionAndSessionProperties) SetType(v string) {
	o.Type = v
}

// GetUniqueId returns the UniqueId field value
func (o *UserActionAndSessionProperties) GetUniqueId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UniqueId
}

// GetUniqueIdOk returns a tuple with the UniqueId field value
// and a boolean to check if the value has been set.
func (o *UserActionAndSessionProperties) GetUniqueIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueId, true
}

// SetUniqueId sets field value
func (o *UserActionAndSessionProperties) SetUniqueId(v int32) {
	o.UniqueId = v
}

func (o UserActionAndSessionProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActionAndSessionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aggregation) {
		toSerialize["aggregation"] = o.Aggregation
	}
	if !IsNil(o.CleanupRule) {
		toSerialize["cleanupRule"] = o.CleanupRule
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IgnoreCase) {
		toSerialize["ignoreCase"] = o.IgnoreCase
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.LongStringLength) {
		toSerialize["longStringLength"] = o.LongStringLength
	}
	if !IsNil(o.MetadataId) {
		toSerialize["metadataId"] = o.MetadataId
	}
	toSerialize["origin"] = o.Origin
	if !IsNil(o.ServerSideRequestAttribute) {
		toSerialize["serverSideRequestAttribute"] = o.ServerSideRequestAttribute
	}
	if !IsNil(o.StoreAsSessionProperty) {
		toSerialize["storeAsSessionProperty"] = o.StoreAsSessionProperty
	}
	if !IsNil(o.StoreAsUserActionProperty) {
		toSerialize["storeAsUserActionProperty"] = o.StoreAsUserActionProperty
	}
	toSerialize["type"] = o.Type
	toSerialize["uniqueId"] = o.UniqueId
	return toSerialize, nil
}

type NullableUserActionAndSessionProperties struct {
	value *UserActionAndSessionProperties
	isSet bool
}

func (v NullableUserActionAndSessionProperties) Get() *UserActionAndSessionProperties {
	return v.value
}

func (v *NullableUserActionAndSessionProperties) Set(val *UserActionAndSessionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActionAndSessionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActionAndSessionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActionAndSessionProperties(val *UserActionAndSessionProperties) *NullableUserActionAndSessionProperties {
	return &NullableUserActionAndSessionProperties{value: val, isSet: true}
}

func (v NullableUserActionAndSessionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActionAndSessionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


