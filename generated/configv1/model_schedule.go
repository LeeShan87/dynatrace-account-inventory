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

// checks if the Schedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Schedule{}

// Schedule The schedule of the maintenance window.
type Schedule struct {
	// The end date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	End string `json:"end"`
	Recurrence *Recurrence `json:"recurrence,omitempty"`
	// The type of the schedule recurrence.
	RecurrenceType string `json:"recurrenceType"`
	// The start date and time of the maintenance window validity period in yyyy-mm-dd HH:mm format.
	Start string `json:"start"`
	// The time zone of the start and end time. Default time zone is UTC.   You can use either UTC offset `UTC+01:00` format or the IANA Time Zone Database format (for example, `Europe/Vienna`).
	ZoneId string `json:"zoneId"`
}

// NewSchedule instantiates a new Schedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedule(end string, recurrenceType string, start string, zoneId string) *Schedule {
	this := Schedule{}
	this.End = end
	this.RecurrenceType = recurrenceType
	this.Start = start
	this.ZoneId = zoneId
	return &this
}

// NewScheduleWithDefaults instantiates a new Schedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleWithDefaults() *Schedule {
	this := Schedule{}
	return &this
}

// GetEnd returns the End field value
func (o *Schedule) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *Schedule) SetEnd(v string) {
	o.End = v
}

// GetRecurrence returns the Recurrence field value if set, zero value otherwise.
func (o *Schedule) GetRecurrence() Recurrence {
	if o == nil || IsNil(o.Recurrence) {
		var ret Recurrence
		return ret
	}
	return *o.Recurrence
}

// GetRecurrenceOk returns a tuple with the Recurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Schedule) GetRecurrenceOk() (*Recurrence, bool) {
	if o == nil || IsNil(o.Recurrence) {
		return nil, false
	}
	return o.Recurrence, true
}

// HasRecurrence returns a boolean if a field has been set.
func (o *Schedule) HasRecurrence() bool {
	if o != nil && !IsNil(o.Recurrence) {
		return true
	}

	return false
}

// SetRecurrence gets a reference to the given Recurrence and assigns it to the Recurrence field.
func (o *Schedule) SetRecurrence(v Recurrence) {
	o.Recurrence = &v
}

// GetRecurrenceType returns the RecurrenceType field value
func (o *Schedule) GetRecurrenceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecurrenceType
}

// GetRecurrenceTypeOk returns a tuple with the RecurrenceType field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetRecurrenceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecurrenceType, true
}

// SetRecurrenceType sets field value
func (o *Schedule) SetRecurrenceType(v string) {
	o.RecurrenceType = v
}

// GetStart returns the Start field value
func (o *Schedule) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *Schedule) SetStart(v string) {
	o.Start = v
}

// GetZoneId returns the ZoneId field value
func (o *Schedule) GetZoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *Schedule) GetZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *Schedule) SetZoneId(v string) {
	o.ZoneId = v
}

func (o Schedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Schedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	if !IsNil(o.Recurrence) {
		toSerialize["recurrence"] = o.Recurrence
	}
	toSerialize["recurrenceType"] = o.RecurrenceType
	toSerialize["start"] = o.Start
	toSerialize["zoneId"] = o.ZoneId
	return toSerialize, nil
}

type NullableSchedule struct {
	value *Schedule
	isSet bool
}

func (v NullableSchedule) Get() *Schedule {
	return v.value
}

func (v *NullableSchedule) Set(val *Schedule) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedule(val *Schedule) *NullableSchedule {
	return &NullableSchedule{value: val, isSet: true}
}

func (v NullableSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


