# CloudEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**DataBase64** | Pointer to **string** |  | [optional] 
**Datacontenttype** | Pointer to **string** |  | [optional] 
**Dataschema** | Pointer to **string** |  | [optional] 
**Dtcontext** | Pointer to **string** | Dynatrace context | [optional] 
**Id** | **string** |  | 
**Source** | **string** |  | 
**Specversion** | **string** |  | 
**Subject** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Traceparent** | Pointer to **string** | Trace related to this event. See [distributed tracing](https://github.com/cloudevents/spec/blob/main/cloudevents/extensions/distributed-tracing.md) for further information. | [optional] 
**Type** | **string** |  | 

## Methods

### NewCloudEvent

`func NewCloudEvent(id string, source string, specversion string, type_ string, ) *CloudEvent`

NewCloudEvent instantiates a new CloudEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEventWithDefaults

`func NewCloudEventWithDefaults() *CloudEvent`

NewCloudEventWithDefaults instantiates a new CloudEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudEvent) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudEvent) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudEvent) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *CloudEvent) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDataBase64

`func (o *CloudEvent) GetDataBase64() string`

GetDataBase64 returns the DataBase64 field if non-nil, zero value otherwise.

### GetDataBase64Ok

`func (o *CloudEvent) GetDataBase64Ok() (*string, bool)`

GetDataBase64Ok returns a tuple with the DataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataBase64

`func (o *CloudEvent) SetDataBase64(v string)`

SetDataBase64 sets DataBase64 field to given value.

### HasDataBase64

`func (o *CloudEvent) HasDataBase64() bool`

HasDataBase64 returns a boolean if a field has been set.

### GetDatacontenttype

`func (o *CloudEvent) GetDatacontenttype() string`

GetDatacontenttype returns the Datacontenttype field if non-nil, zero value otherwise.

### GetDatacontenttypeOk

`func (o *CloudEvent) GetDatacontenttypeOk() (*string, bool)`

GetDatacontenttypeOk returns a tuple with the Datacontenttype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacontenttype

`func (o *CloudEvent) SetDatacontenttype(v string)`

SetDatacontenttype sets Datacontenttype field to given value.

### HasDatacontenttype

`func (o *CloudEvent) HasDatacontenttype() bool`

HasDatacontenttype returns a boolean if a field has been set.

### GetDataschema

`func (o *CloudEvent) GetDataschema() string`

GetDataschema returns the Dataschema field if non-nil, zero value otherwise.

### GetDataschemaOk

`func (o *CloudEvent) GetDataschemaOk() (*string, bool)`

GetDataschemaOk returns a tuple with the Dataschema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataschema

`func (o *CloudEvent) SetDataschema(v string)`

SetDataschema sets Dataschema field to given value.

### HasDataschema

`func (o *CloudEvent) HasDataschema() bool`

HasDataschema returns a boolean if a field has been set.

### GetDtcontext

`func (o *CloudEvent) GetDtcontext() string`

GetDtcontext returns the Dtcontext field if non-nil, zero value otherwise.

### GetDtcontextOk

`func (o *CloudEvent) GetDtcontextOk() (*string, bool)`

GetDtcontextOk returns a tuple with the Dtcontext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcontext

`func (o *CloudEvent) SetDtcontext(v string)`

SetDtcontext sets Dtcontext field to given value.

### HasDtcontext

`func (o *CloudEvent) HasDtcontext() bool`

HasDtcontext returns a boolean if a field has been set.

### GetId

`func (o *CloudEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudEvent) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *CloudEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CloudEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CloudEvent) SetSource(v string)`

SetSource sets Source field to given value.


### GetSpecversion

`func (o *CloudEvent) GetSpecversion() string`

GetSpecversion returns the Specversion field if non-nil, zero value otherwise.

### GetSpecversionOk

`func (o *CloudEvent) GetSpecversionOk() (*string, bool)`

GetSpecversionOk returns a tuple with the Specversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecversion

`func (o *CloudEvent) SetSpecversion(v string)`

SetSpecversion sets Specversion field to given value.


### GetSubject

`func (o *CloudEvent) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CloudEvent) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CloudEvent) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CloudEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTime

`func (o *CloudEvent) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CloudEvent) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CloudEvent) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *CloudEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTraceparent

`func (o *CloudEvent) GetTraceparent() string`

GetTraceparent returns the Traceparent field if non-nil, zero value otherwise.

### GetTraceparentOk

`func (o *CloudEvent) GetTraceparentOk() (*string, bool)`

GetTraceparentOk returns a tuple with the Traceparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceparent

`func (o *CloudEvent) SetTraceparent(v string)`

SetTraceparent sets Traceparent field to given value.

### HasTraceparent

`func (o *CloudEvent) HasTraceparent() bool`

HasTraceparent returns a boolean if a field has been set.

### GetType

`func (o *CloudEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


