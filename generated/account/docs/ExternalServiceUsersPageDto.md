# ExternalServiceUsersPageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ExternalServiceUserDto**](ExternalServiceUserDto.md) |  | 
**NextPageKey** | **string** | Next page key to be used in querying for next results page | 
**TotalCount** | **float32** | Total number of service users | 

## Methods

### NewExternalServiceUsersPageDto

`func NewExternalServiceUsersPageDto(results []ExternalServiceUserDto, nextPageKey string, totalCount float32, ) *ExternalServiceUsersPageDto`

NewExternalServiceUsersPageDto instantiates a new ExternalServiceUsersPageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalServiceUsersPageDtoWithDefaults

`func NewExternalServiceUsersPageDtoWithDefaults() *ExternalServiceUsersPageDto`

NewExternalServiceUsersPageDtoWithDefaults instantiates a new ExternalServiceUsersPageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ExternalServiceUsersPageDto) GetResults() []ExternalServiceUserDto`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ExternalServiceUsersPageDto) GetResultsOk() (*[]ExternalServiceUserDto, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ExternalServiceUsersPageDto) SetResults(v []ExternalServiceUserDto)`

SetResults sets Results field to given value.


### GetNextPageKey

`func (o *ExternalServiceUsersPageDto) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ExternalServiceUsersPageDto) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ExternalServiceUsersPageDto) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.


### GetTotalCount

`func (o *ExternalServiceUsersPageDto) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ExternalServiceUsersPageDto) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ExternalServiceUsersPageDto) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


