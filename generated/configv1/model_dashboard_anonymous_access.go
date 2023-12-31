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

// checks if the DashboardAnonymousAccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DashboardAnonymousAccess{}

// DashboardAnonymousAccess Configuration of the [anonymous access](https://dt-url.net/ov03sf1) to the dashboard.
type DashboardAnonymousAccess struct {
	// A list of management zones that can display data on the publicly shared dashboard.   Specify management zone IDs here. For each management zone you specify Dynatrace generates an access link. You can access them in the **urls** list.   To share the dashboard with its default management zone, use the `default` value.
	ManagementZoneIds []string `json:"managementZoneIds"`
	// A list of URLs for anonymous access to the dashboard.   Each link grants access to data from the specific management zone, listed in the in the **managementZoneIds** list.   These links are automatically generated by Dynatrace, you can't change them.
	Urls *map[string]string `json:"urls,omitempty"`
}

// NewDashboardAnonymousAccess instantiates a new DashboardAnonymousAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardAnonymousAccess(managementZoneIds []string) *DashboardAnonymousAccess {
	this := DashboardAnonymousAccess{}
	this.ManagementZoneIds = managementZoneIds
	return &this
}

// NewDashboardAnonymousAccessWithDefaults instantiates a new DashboardAnonymousAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardAnonymousAccessWithDefaults() *DashboardAnonymousAccess {
	this := DashboardAnonymousAccess{}
	return &this
}

// GetManagementZoneIds returns the ManagementZoneIds field value
func (o *DashboardAnonymousAccess) GetManagementZoneIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ManagementZoneIds
}

// GetManagementZoneIdsOk returns a tuple with the ManagementZoneIds field value
// and a boolean to check if the value has been set.
func (o *DashboardAnonymousAccess) GetManagementZoneIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagementZoneIds, true
}

// SetManagementZoneIds sets field value
func (o *DashboardAnonymousAccess) SetManagementZoneIds(v []string) {
	o.ManagementZoneIds = v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *DashboardAnonymousAccess) GetUrls() map[string]string {
	if o == nil || IsNil(o.Urls) {
		var ret map[string]string
		return ret
	}
	return *o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardAnonymousAccess) GetUrlsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *DashboardAnonymousAccess) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given map[string]string and assigns it to the Urls field.
func (o *DashboardAnonymousAccess) SetUrls(v map[string]string) {
	o.Urls = &v
}

func (o DashboardAnonymousAccess) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DashboardAnonymousAccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["managementZoneIds"] = o.ManagementZoneIds
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	return toSerialize, nil
}

type NullableDashboardAnonymousAccess struct {
	value *DashboardAnonymousAccess
	isSet bool
}

func (v NullableDashboardAnonymousAccess) Get() *DashboardAnonymousAccess {
	return v.value
}

func (v *NullableDashboardAnonymousAccess) Set(val *DashboardAnonymousAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardAnonymousAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardAnonymousAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardAnonymousAccess(val *DashboardAnonymousAccess) *NullableDashboardAnonymousAccess {
	return &NullableDashboardAnonymousAccess{value: val, isSet: true}
}

func (v NullableDashboardAnonymousAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardAnonymousAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


