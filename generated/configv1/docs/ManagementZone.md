# ManagementZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the management zone. | [optional] 
**DimensionalRules** | Pointer to [**[]MzDimensionalRule**](MzDimensionalRule.md) | A list of dimensional data rules for management zone usage.   If several rules are specified, the **OR** logic applies. | [optional] 
**EntitySelectorBasedRules** | Pointer to [**[]EntitySelectorBasedMzRule**](EntitySelectorBasedMzRule.md) | A list of entity-selector based rules for management zone usage.  If several rules are specified, the **OR** logic applies. | [optional] 
**Id** | Pointer to **string** | The ID of the management zone. | [optional] 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Name** | **string** | The name of the management zone. | 
**Rules** | Pointer to [**[]MzRule**](MzRule.md) | A list of rules for management zone usage.   If several rules are specified, the **OR** logic applies. | [optional] 

## Methods

### NewManagementZone

`func NewManagementZone(name string, ) *ManagementZone`

NewManagementZone instantiates a new ManagementZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementZoneWithDefaults

`func NewManagementZoneWithDefaults() *ManagementZone`

NewManagementZoneWithDefaults instantiates a new ManagementZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ManagementZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagementZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagementZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagementZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDimensionalRules

`func (o *ManagementZone) GetDimensionalRules() []MzDimensionalRule`

GetDimensionalRules returns the DimensionalRules field if non-nil, zero value otherwise.

### GetDimensionalRulesOk

`func (o *ManagementZone) GetDimensionalRulesOk() (*[]MzDimensionalRule, bool)`

GetDimensionalRulesOk returns a tuple with the DimensionalRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionalRules

`func (o *ManagementZone) SetDimensionalRules(v []MzDimensionalRule)`

SetDimensionalRules sets DimensionalRules field to given value.

### HasDimensionalRules

`func (o *ManagementZone) HasDimensionalRules() bool`

HasDimensionalRules returns a boolean if a field has been set.

### GetEntitySelectorBasedRules

`func (o *ManagementZone) GetEntitySelectorBasedRules() []EntitySelectorBasedMzRule`

GetEntitySelectorBasedRules returns the EntitySelectorBasedRules field if non-nil, zero value otherwise.

### GetEntitySelectorBasedRulesOk

`func (o *ManagementZone) GetEntitySelectorBasedRulesOk() (*[]EntitySelectorBasedMzRule, bool)`

GetEntitySelectorBasedRulesOk returns a tuple with the EntitySelectorBasedRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelectorBasedRules

`func (o *ManagementZone) SetEntitySelectorBasedRules(v []EntitySelectorBasedMzRule)`

SetEntitySelectorBasedRules sets EntitySelectorBasedRules field to given value.

### HasEntitySelectorBasedRules

`func (o *ManagementZone) HasEntitySelectorBasedRules() bool`

HasEntitySelectorBasedRules returns a boolean if a field has been set.

### GetId

`func (o *ManagementZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ManagementZone) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ManagementZone) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ManagementZone) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ManagementZone) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *ManagementZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementZone) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *ManagementZone) GetRules() []MzRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ManagementZone) GetRulesOk() (*[]MzRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ManagementZone) SetRules(v []MzRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ManagementZone) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


