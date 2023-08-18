# Apdex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrustratingFallbackThreshold** | Pointer to **float32** | Fallback threshold of an XHR action, defining a tolerable user experience, when the configured KPM is not available. | [optional] 
**FrustratingThreshold** | Pointer to **float32** | Maximal value of apdex, which is considered as tolerable user experience. | [optional] 
**ToleratedFallbackThreshold** | Pointer to **float32** | Fallback threshold of an XHR action, defining a satisfied user experience, when the configured KPM is not available. | [optional] 
**ToleratedThreshold** | Pointer to **float32** | Maximal value of apdex, which is considered as satisfied user experience. | [optional] 

## Methods

### NewApdex

`func NewApdex() *Apdex`

NewApdex instantiates a new Apdex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApdexWithDefaults

`func NewApdexWithDefaults() *Apdex`

NewApdexWithDefaults instantiates a new Apdex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrustratingFallbackThreshold

`func (o *Apdex) GetFrustratingFallbackThreshold() float32`

GetFrustratingFallbackThreshold returns the FrustratingFallbackThreshold field if non-nil, zero value otherwise.

### GetFrustratingFallbackThresholdOk

`func (o *Apdex) GetFrustratingFallbackThresholdOk() (*float32, bool)`

GetFrustratingFallbackThresholdOk returns a tuple with the FrustratingFallbackThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrustratingFallbackThreshold

`func (o *Apdex) SetFrustratingFallbackThreshold(v float32)`

SetFrustratingFallbackThreshold sets FrustratingFallbackThreshold field to given value.

### HasFrustratingFallbackThreshold

`func (o *Apdex) HasFrustratingFallbackThreshold() bool`

HasFrustratingFallbackThreshold returns a boolean if a field has been set.

### GetFrustratingThreshold

`func (o *Apdex) GetFrustratingThreshold() float32`

GetFrustratingThreshold returns the FrustratingThreshold field if non-nil, zero value otherwise.

### GetFrustratingThresholdOk

`func (o *Apdex) GetFrustratingThresholdOk() (*float32, bool)`

GetFrustratingThresholdOk returns a tuple with the FrustratingThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrustratingThreshold

`func (o *Apdex) SetFrustratingThreshold(v float32)`

SetFrustratingThreshold sets FrustratingThreshold field to given value.

### HasFrustratingThreshold

`func (o *Apdex) HasFrustratingThreshold() bool`

HasFrustratingThreshold returns a boolean if a field has been set.

### GetToleratedFallbackThreshold

`func (o *Apdex) GetToleratedFallbackThreshold() float32`

GetToleratedFallbackThreshold returns the ToleratedFallbackThreshold field if non-nil, zero value otherwise.

### GetToleratedFallbackThresholdOk

`func (o *Apdex) GetToleratedFallbackThresholdOk() (*float32, bool)`

GetToleratedFallbackThresholdOk returns a tuple with the ToleratedFallbackThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToleratedFallbackThreshold

`func (o *Apdex) SetToleratedFallbackThreshold(v float32)`

SetToleratedFallbackThreshold sets ToleratedFallbackThreshold field to given value.

### HasToleratedFallbackThreshold

`func (o *Apdex) HasToleratedFallbackThreshold() bool`

HasToleratedFallbackThreshold returns a boolean if a field has been set.

### GetToleratedThreshold

`func (o *Apdex) GetToleratedThreshold() float32`

GetToleratedThreshold returns the ToleratedThreshold field if non-nil, zero value otherwise.

### GetToleratedThresholdOk

`func (o *Apdex) GetToleratedThresholdOk() (*float32, bool)`

GetToleratedThresholdOk returns a tuple with the ToleratedThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToleratedThreshold

`func (o *Apdex) SetToleratedThreshold(v float32)`

SetToleratedThreshold sets ToleratedThreshold field to given value.

### HasToleratedThreshold

`func (o *Apdex) HasToleratedThreshold() bool`

HasToleratedThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


