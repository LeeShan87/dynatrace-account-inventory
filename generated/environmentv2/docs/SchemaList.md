# SchemaList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]SchemaStub**](SchemaStub.md) | A list of settings schemas. | 
**TotalCount** | **int64** | The number of schemas in the list. | 

## Methods

### NewSchemaList

`func NewSchemaList(items []SchemaStub, totalCount int64, ) *SchemaList`

NewSchemaList instantiates a new SchemaList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaListWithDefaults

`func NewSchemaListWithDefaults() *SchemaList`

NewSchemaListWithDefaults instantiates a new SchemaList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SchemaList) GetItems() []SchemaStub`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SchemaList) GetItemsOk() (*[]SchemaStub, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SchemaList) SetItems(v []SchemaStub)`

SetItems sets Items field to given value.


### GetTotalCount

`func (o *SchemaList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SchemaList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SchemaList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


