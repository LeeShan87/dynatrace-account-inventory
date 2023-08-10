# UserListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **float32** | The number of entries in the list. | 
**Items** | [**[]UsersDto**](UsersDto.md) | A list of the account&#39;s users. | 

## Methods

### NewUserListDto

`func NewUserListDto(count float32, items []UsersDto, ) *UserListDto`

NewUserListDto instantiates a new UserListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListDtoWithDefaults

`func NewUserListDtoWithDefaults() *UserListDto`

NewUserListDtoWithDefaults instantiates a new UserListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *UserListDto) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserListDto) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserListDto) SetCount(v float32)`

SetCount sets Count field to given value.


### GetItems

`func (o *UserListDto) GetItems() []UsersDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserListDto) GetItemsOk() (*[]UsersDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserListDto) SetItems(v []UsersDto)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


