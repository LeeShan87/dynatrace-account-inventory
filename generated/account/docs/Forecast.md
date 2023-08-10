# Forecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForecastMedian** | **float32** | The median of the forecast snapshot. | 
**ForecastLower** | **float32** | The lower bound for the forecast snapshot. | 
**ForecastUpper** | **float32** | The upper bound for the forecast snapshot. | 
**Budget** | **float32** | The budget for the forecast snapshot. | 
**ForecastBudgetPct** | **float32** | The budget percent for the forecast snapshot. | 
**ForecastBudgetDate** | **time.Time** | The date the forecast snapshot&#39;s budget reached the quota amount. | 
**ForecastCreatedAt** | **time.Time** | The date the forecast snapshot was created. | 

## Methods

### NewForecast

`func NewForecast(forecastMedian float32, forecastLower float32, forecastUpper float32, budget float32, forecastBudgetPct float32, forecastBudgetDate time.Time, forecastCreatedAt time.Time, ) *Forecast`

NewForecast instantiates a new Forecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastWithDefaults

`func NewForecastWithDefaults() *Forecast`

NewForecastWithDefaults instantiates a new Forecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForecastMedian

`func (o *Forecast) GetForecastMedian() float32`

GetForecastMedian returns the ForecastMedian field if non-nil, zero value otherwise.

### GetForecastMedianOk

`func (o *Forecast) GetForecastMedianOk() (*float32, bool)`

GetForecastMedianOk returns a tuple with the ForecastMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastMedian

`func (o *Forecast) SetForecastMedian(v float32)`

SetForecastMedian sets ForecastMedian field to given value.


### GetForecastLower

`func (o *Forecast) GetForecastLower() float32`

GetForecastLower returns the ForecastLower field if non-nil, zero value otherwise.

### GetForecastLowerOk

`func (o *Forecast) GetForecastLowerOk() (*float32, bool)`

GetForecastLowerOk returns a tuple with the ForecastLower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastLower

`func (o *Forecast) SetForecastLower(v float32)`

SetForecastLower sets ForecastLower field to given value.


### GetForecastUpper

`func (o *Forecast) GetForecastUpper() float32`

GetForecastUpper returns the ForecastUpper field if non-nil, zero value otherwise.

### GetForecastUpperOk

`func (o *Forecast) GetForecastUpperOk() (*float32, bool)`

GetForecastUpperOk returns a tuple with the ForecastUpper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastUpper

`func (o *Forecast) SetForecastUpper(v float32)`

SetForecastUpper sets ForecastUpper field to given value.


### GetBudget

`func (o *Forecast) GetBudget() float32`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *Forecast) GetBudgetOk() (*float32, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *Forecast) SetBudget(v float32)`

SetBudget sets Budget field to given value.


### GetForecastBudgetPct

`func (o *Forecast) GetForecastBudgetPct() float32`

GetForecastBudgetPct returns the ForecastBudgetPct field if non-nil, zero value otherwise.

### GetForecastBudgetPctOk

`func (o *Forecast) GetForecastBudgetPctOk() (*float32, bool)`

GetForecastBudgetPctOk returns a tuple with the ForecastBudgetPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastBudgetPct

`func (o *Forecast) SetForecastBudgetPct(v float32)`

SetForecastBudgetPct sets ForecastBudgetPct field to given value.


### GetForecastBudgetDate

`func (o *Forecast) GetForecastBudgetDate() time.Time`

GetForecastBudgetDate returns the ForecastBudgetDate field if non-nil, zero value otherwise.

### GetForecastBudgetDateOk

`func (o *Forecast) GetForecastBudgetDateOk() (*time.Time, bool)`

GetForecastBudgetDateOk returns a tuple with the ForecastBudgetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastBudgetDate

`func (o *Forecast) SetForecastBudgetDate(v time.Time)`

SetForecastBudgetDate sets ForecastBudgetDate field to given value.


### GetForecastCreatedAt

`func (o *Forecast) GetForecastCreatedAt() time.Time`

GetForecastCreatedAt returns the ForecastCreatedAt field if non-nil, zero value otherwise.

### GetForecastCreatedAtOk

`func (o *Forecast) GetForecastCreatedAtOk() (*time.Time, bool)`

GetForecastCreatedAtOk returns a tuple with the ForecastCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastCreatedAt

`func (o *Forecast) SetForecastCreatedAt(v time.Time)`

SetForecastCreatedAt sets ForecastCreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


