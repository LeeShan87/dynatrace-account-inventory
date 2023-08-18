# RegisteredExtensionResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionName** | Pointer to **string** | FQN of the extension registered in the tenant. | [optional] 
**ExtensionVersion** | Pointer to **string** | Version number of the extension. | [optional] 

## Methods

### NewRegisteredExtensionResultDto

`func NewRegisteredExtensionResultDto() *RegisteredExtensionResultDto`

NewRegisteredExtensionResultDto instantiates a new RegisteredExtensionResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredExtensionResultDtoWithDefaults

`func NewRegisteredExtensionResultDtoWithDefaults() *RegisteredExtensionResultDto`

NewRegisteredExtensionResultDtoWithDefaults instantiates a new RegisteredExtensionResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionName

`func (o *RegisteredExtensionResultDto) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *RegisteredExtensionResultDto) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *RegisteredExtensionResultDto) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *RegisteredExtensionResultDto) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetExtensionVersion

`func (o *RegisteredExtensionResultDto) GetExtensionVersion() string`

GetExtensionVersion returns the ExtensionVersion field if non-nil, zero value otherwise.

### GetExtensionVersionOk

`func (o *RegisteredExtensionResultDto) GetExtensionVersionOk() (*string, bool)`

GetExtensionVersionOk returns a tuple with the ExtensionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionVersion

`func (o *RegisteredExtensionResultDto) SetExtensionVersion(v string)`

SetExtensionVersion sets ExtensionVersion field to given value.

### HasExtensionVersion

`func (o *RegisteredExtensionResultDto) HasExtensionVersion() bool`

HasExtensionVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


