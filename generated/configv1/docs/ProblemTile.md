# ProblemTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitySelector** | Pointer to **string** | The entity scope of the problem tile. For further information please look at the Problems API v2 &#39;/problems&#39; endpoint. | [optional] 
**ProblemSelector** | Pointer to **string** | Defines the scope of the problem tile. Only problems matching the specified criteria are taken into account. For further information please look at the Problems API v2 &#39;/problems&#39; endpoint. | [optional] 

## Methods

### NewProblemTile

`func NewProblemTile() *ProblemTile`

NewProblemTile instantiates a new ProblemTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProblemTileWithDefaults

`func NewProblemTileWithDefaults() *ProblemTile`

NewProblemTileWithDefaults instantiates a new ProblemTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitySelector

`func (o *ProblemTile) GetEntitySelector() string`

GetEntitySelector returns the EntitySelector field if non-nil, zero value otherwise.

### GetEntitySelectorOk

`func (o *ProblemTile) GetEntitySelectorOk() (*string, bool)`

GetEntitySelectorOk returns a tuple with the EntitySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelector

`func (o *ProblemTile) SetEntitySelector(v string)`

SetEntitySelector sets EntitySelector field to given value.

### HasEntitySelector

`func (o *ProblemTile) HasEntitySelector() bool`

HasEntitySelector returns a boolean if a field has been set.

### GetProblemSelector

`func (o *ProblemTile) GetProblemSelector() string`

GetProblemSelector returns the ProblemSelector field if non-nil, zero value otherwise.

### GetProblemSelectorOk

`func (o *ProblemTile) GetProblemSelectorOk() (*string, bool)`

GetProblemSelectorOk returns a tuple with the ProblemSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemSelector

`func (o *ProblemTile) SetProblemSelector(v string)`

SetProblemSelector sets ProblemSelector field to given value.

### HasProblemSelector

`func (o *ProblemTile) HasProblemSelector() bool`

HasProblemSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


