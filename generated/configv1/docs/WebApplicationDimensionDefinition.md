# WebApplicationDimensionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | **string** | The dimension of the metric. | 
**PropertyKey** | Pointer to **string** | The key of the user action property.    Only applicable for the &#x60;StringProperty&#x60; dimension. | [optional] 
**TopX** | **int32** | The number of top values to be calculated. | 

## Methods

### NewWebApplicationDimensionDefinition

`func NewWebApplicationDimensionDefinition(dimension string, topX int32, ) *WebApplicationDimensionDefinition`

NewWebApplicationDimensionDefinition instantiates a new WebApplicationDimensionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationDimensionDefinitionWithDefaults

`func NewWebApplicationDimensionDefinitionWithDefaults() *WebApplicationDimensionDefinition`

NewWebApplicationDimensionDefinitionWithDefaults instantiates a new WebApplicationDimensionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *WebApplicationDimensionDefinition) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *WebApplicationDimensionDefinition) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *WebApplicationDimensionDefinition) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetPropertyKey

`func (o *WebApplicationDimensionDefinition) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *WebApplicationDimensionDefinition) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *WebApplicationDimensionDefinition) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.

### HasPropertyKey

`func (o *WebApplicationDimensionDefinition) HasPropertyKey() bool`

HasPropertyKey returns a boolean if a field has been set.

### GetTopX

`func (o *WebApplicationDimensionDefinition) GetTopX() int32`

GetTopX returns the TopX field if non-nil, zero value otherwise.

### GetTopXOk

`func (o *WebApplicationDimensionDefinition) GetTopXOk() (*int32, bool)`

GetTopXOk returns a tuple with the TopX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopX

`func (o *WebApplicationDimensionDefinition) SetTopX(v int32)`

SetTopX sets TopX field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


