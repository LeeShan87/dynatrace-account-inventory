# AbstractSloAlertDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertName** | **string** | Name of the alert. | 
**AlertThreshold** | **float64** | Threshold of the alert. Status alerts trigger if they fall below this value, burn rate alerts trigger if they exceed the value. | 
**AlertType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BURN_RATE&#x60; -&gt; BurnRateAlert  * &#x60;STATUS&#x60; -&gt; StatusAlert   | 

## Methods

### NewAbstractSloAlertDto

`func NewAbstractSloAlertDto(alertName string, alertThreshold float64, alertType string, ) *AbstractSloAlertDto`

NewAbstractSloAlertDto instantiates a new AbstractSloAlertDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractSloAlertDtoWithDefaults

`func NewAbstractSloAlertDtoWithDefaults() *AbstractSloAlertDto`

NewAbstractSloAlertDtoWithDefaults instantiates a new AbstractSloAlertDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertName

`func (o *AbstractSloAlertDto) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AbstractSloAlertDto) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AbstractSloAlertDto) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.


### GetAlertThreshold

`func (o *AbstractSloAlertDto) GetAlertThreshold() float64`

GetAlertThreshold returns the AlertThreshold field if non-nil, zero value otherwise.

### GetAlertThresholdOk

`func (o *AbstractSloAlertDto) GetAlertThresholdOk() (*float64, bool)`

GetAlertThresholdOk returns a tuple with the AlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertThreshold

`func (o *AbstractSloAlertDto) SetAlertThreshold(v float64)`

SetAlertThreshold sets AlertThreshold field to given value.


### GetAlertType

`func (o *AbstractSloAlertDto) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AbstractSloAlertDto) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AbstractSloAlertDto) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


