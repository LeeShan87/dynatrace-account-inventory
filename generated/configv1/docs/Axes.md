# Axes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XAxis** | [**Axis**](Axis.md) |  | 
**YAxes** | [**[]YAxis**](YAxis.md) | y Axes configuration | 

## Methods

### NewAxes

`func NewAxes(xAxis Axis, yAxes []YAxis, ) *Axes`

NewAxes instantiates a new Axes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxesWithDefaults

`func NewAxesWithDefaults() *Axes`

NewAxesWithDefaults instantiates a new Axes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXAxis

`func (o *Axes) GetXAxis() Axis`

GetXAxis returns the XAxis field if non-nil, zero value otherwise.

### GetXAxisOk

`func (o *Axes) GetXAxisOk() (*Axis, bool)`

GetXAxisOk returns a tuple with the XAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxis

`func (o *Axes) SetXAxis(v Axis)`

SetXAxis sets XAxis field to given value.


### GetYAxes

`func (o *Axes) GetYAxes() []YAxis`

GetYAxes returns the YAxes field if non-nil, zero value otherwise.

### GetYAxesOk

`func (o *Axes) GetYAxesOk() (*[]YAxis, bool)`

GetYAxesOk returns a tuple with the YAxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxes

`func (o *Axes) SetYAxes(v []YAxis)`

SetYAxes sets YAxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


