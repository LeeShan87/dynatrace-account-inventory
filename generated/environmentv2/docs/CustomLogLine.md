# CustomLogLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogLevel** | Pointer to **string** | Log level of the message | [optional] 
**Message** | Pointer to **string** | The message | [optional] 
**Timestamp** | Pointer to **int64** | A timestamp of this log message | [optional] 

## Methods

### NewCustomLogLine

`func NewCustomLogLine() *CustomLogLine`

NewCustomLogLine instantiates a new CustomLogLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLogLineWithDefaults

`func NewCustomLogLineWithDefaults() *CustomLogLine`

NewCustomLogLineWithDefaults instantiates a new CustomLogLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogLevel

`func (o *CustomLogLine) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *CustomLogLine) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *CustomLogLine) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *CustomLogLine) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetMessage

`func (o *CustomLogLine) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CustomLogLine) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CustomLogLine) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CustomLogLine) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *CustomLogLine) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CustomLogLine) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CustomLogLine) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CustomLogLine) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


