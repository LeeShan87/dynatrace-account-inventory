# Warnings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangedMetricKeys** | Pointer to [**[]WarningLine**](WarningLine.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewWarnings

`func NewWarnings() *Warnings`

NewWarnings instantiates a new Warnings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningsWithDefaults

`func NewWarningsWithDefaults() *Warnings`

NewWarningsWithDefaults instantiates a new Warnings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangedMetricKeys

`func (o *Warnings) GetChangedMetricKeys() []WarningLine`

GetChangedMetricKeys returns the ChangedMetricKeys field if non-nil, zero value otherwise.

### GetChangedMetricKeysOk

`func (o *Warnings) GetChangedMetricKeysOk() (*[]WarningLine, bool)`

GetChangedMetricKeysOk returns a tuple with the ChangedMetricKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedMetricKeys

`func (o *Warnings) SetChangedMetricKeys(v []WarningLine)`

SetChangedMetricKeys sets ChangedMetricKeys field to given value.

### HasChangedMetricKeys

`func (o *Warnings) HasChangedMetricKeys() bool`

HasChangedMetricKeys returns a boolean if a field has been set.

### GetMessage

`func (o *Warnings) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Warnings) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Warnings) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Warnings) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


