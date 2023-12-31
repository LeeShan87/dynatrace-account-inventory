/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
	"time"
)

// checks if the CloudEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudEvent{}

// CloudEvent CloudEvents is a [specification](https://github.com/cloudevents/spec/blob/v1.0/spec.json) for describing event data in common formats to provide interoperability across services, platforms and systems.
type CloudEvent struct {
	Data map[string]interface{} `json:"data,omitempty"`
	DataBase64 *string `json:"data_base64,omitempty"`
	Datacontenttype *string `json:"datacontenttype,omitempty"`
	Dataschema *string `json:"dataschema,omitempty"`
	// Dynatrace context
	Dtcontext *string `json:"dtcontext,omitempty"`
	Id string `json:"id"`
	Source string `json:"source"`
	Specversion string `json:"specversion"`
	Subject *string `json:"subject,omitempty"`
	Time *time.Time `json:"time,omitempty"`
	// Trace related to this event. See [distributed tracing](https://github.com/cloudevents/spec/blob/main/cloudevents/extensions/distributed-tracing.md) for further information.
	Traceparent *string `json:"traceparent,omitempty"`
	Type string `json:"type"`
}

// NewCloudEvent instantiates a new CloudEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudEvent(id string, source string, specversion string, type_ string) *CloudEvent {
	this := CloudEvent{}
	this.Id = id
	this.Source = source
	this.Specversion = specversion
	this.Type = type_
	return &this
}

// NewCloudEventWithDefaults instantiates a new CloudEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudEventWithDefaults() *CloudEvent {
	this := CloudEvent{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudEvent) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CloudEvent) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetDataBase64 returns the DataBase64 field value if set, zero value otherwise.
func (o *CloudEvent) GetDataBase64() string {
	if o == nil || IsNil(o.DataBase64) {
		var ret string
		return ret
	}
	return *o.DataBase64
}

// GetDataBase64Ok returns a tuple with the DataBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetDataBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.DataBase64) {
		return nil, false
	}
	return o.DataBase64, true
}

// HasDataBase64 returns a boolean if a field has been set.
func (o *CloudEvent) HasDataBase64() bool {
	if o != nil && !IsNil(o.DataBase64) {
		return true
	}

	return false
}

// SetDataBase64 gets a reference to the given string and assigns it to the DataBase64 field.
func (o *CloudEvent) SetDataBase64(v string) {
	o.DataBase64 = &v
}

// GetDatacontenttype returns the Datacontenttype field value if set, zero value otherwise.
func (o *CloudEvent) GetDatacontenttype() string {
	if o == nil || IsNil(o.Datacontenttype) {
		var ret string
		return ret
	}
	return *o.Datacontenttype
}

// GetDatacontenttypeOk returns a tuple with the Datacontenttype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetDatacontenttypeOk() (*string, bool) {
	if o == nil || IsNil(o.Datacontenttype) {
		return nil, false
	}
	return o.Datacontenttype, true
}

// HasDatacontenttype returns a boolean if a field has been set.
func (o *CloudEvent) HasDatacontenttype() bool {
	if o != nil && !IsNil(o.Datacontenttype) {
		return true
	}

	return false
}

// SetDatacontenttype gets a reference to the given string and assigns it to the Datacontenttype field.
func (o *CloudEvent) SetDatacontenttype(v string) {
	o.Datacontenttype = &v
}

// GetDataschema returns the Dataschema field value if set, zero value otherwise.
func (o *CloudEvent) GetDataschema() string {
	if o == nil || IsNil(o.Dataschema) {
		var ret string
		return ret
	}
	return *o.Dataschema
}

// GetDataschemaOk returns a tuple with the Dataschema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetDataschemaOk() (*string, bool) {
	if o == nil || IsNil(o.Dataschema) {
		return nil, false
	}
	return o.Dataschema, true
}

// HasDataschema returns a boolean if a field has been set.
func (o *CloudEvent) HasDataschema() bool {
	if o != nil && !IsNil(o.Dataschema) {
		return true
	}

	return false
}

// SetDataschema gets a reference to the given string and assigns it to the Dataschema field.
func (o *CloudEvent) SetDataschema(v string) {
	o.Dataschema = &v
}

// GetDtcontext returns the Dtcontext field value if set, zero value otherwise.
func (o *CloudEvent) GetDtcontext() string {
	if o == nil || IsNil(o.Dtcontext) {
		var ret string
		return ret
	}
	return *o.Dtcontext
}

// GetDtcontextOk returns a tuple with the Dtcontext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetDtcontextOk() (*string, bool) {
	if o == nil || IsNil(o.Dtcontext) {
		return nil, false
	}
	return o.Dtcontext, true
}

// HasDtcontext returns a boolean if a field has been set.
func (o *CloudEvent) HasDtcontext() bool {
	if o != nil && !IsNil(o.Dtcontext) {
		return true
	}

	return false
}

// SetDtcontext gets a reference to the given string and assigns it to the Dtcontext field.
func (o *CloudEvent) SetDtcontext(v string) {
	o.Dtcontext = &v
}

// GetId returns the Id field value
func (o *CloudEvent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CloudEvent) SetId(v string) {
	o.Id = v
}

// GetSource returns the Source field value
func (o *CloudEvent) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CloudEvent) SetSource(v string) {
	o.Source = v
}

// GetSpecversion returns the Specversion field value
func (o *CloudEvent) GetSpecversion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Specversion
}

// GetSpecversionOk returns a tuple with the Specversion field value
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetSpecversionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Specversion, true
}

// SetSpecversion sets field value
func (o *CloudEvent) SetSpecversion(v string) {
	o.Specversion = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CloudEvent) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CloudEvent) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CloudEvent) SetSubject(v string) {
	o.Subject = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CloudEvent) GetTime() time.Time {
	if o == nil || IsNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CloudEvent) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *CloudEvent) SetTime(v time.Time) {
	o.Time = &v
}

// GetTraceparent returns the Traceparent field value if set, zero value otherwise.
func (o *CloudEvent) GetTraceparent() string {
	if o == nil || IsNil(o.Traceparent) {
		var ret string
		return ret
	}
	return *o.Traceparent
}

// GetTraceparentOk returns a tuple with the Traceparent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetTraceparentOk() (*string, bool) {
	if o == nil || IsNil(o.Traceparent) {
		return nil, false
	}
	return o.Traceparent, true
}

// HasTraceparent returns a boolean if a field has been set.
func (o *CloudEvent) HasTraceparent() bool {
	if o != nil && !IsNil(o.Traceparent) {
		return true
	}

	return false
}

// SetTraceparent gets a reference to the given string and assigns it to the Traceparent field.
func (o *CloudEvent) SetTraceparent(v string) {
	o.Traceparent = &v
}

// GetType returns the Type field value
func (o *CloudEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CloudEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CloudEvent) SetType(v string) {
	o.Type = v
}

func (o CloudEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.DataBase64) {
		toSerialize["data_base64"] = o.DataBase64
	}
	if !IsNil(o.Datacontenttype) {
		toSerialize["datacontenttype"] = o.Datacontenttype
	}
	if !IsNil(o.Dataschema) {
		toSerialize["dataschema"] = o.Dataschema
	}
	if !IsNil(o.Dtcontext) {
		toSerialize["dtcontext"] = o.Dtcontext
	}
	toSerialize["id"] = o.Id
	toSerialize["source"] = o.Source
	toSerialize["specversion"] = o.Specversion
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Traceparent) {
		toSerialize["traceparent"] = o.Traceparent
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableCloudEvent struct {
	value *CloudEvent
	isSet bool
}

func (v NullableCloudEvent) Get() *CloudEvent {
	return v.value
}

func (v *NullableCloudEvent) Set(val *CloudEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEvent(val *CloudEvent) *NullableCloudEvent {
	return &NullableCloudEvent{value: val, isSet: true}
}

func (v NullableCloudEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


