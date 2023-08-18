# RequestInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The target host of the request. | [optional] [readonly] 
**Path** | Pointer to **string** | The request path. | [optional] [readonly] 
**ProtocolDetails** | Pointer to [**ProtocolDetails**](ProtocolDetails.md) |  | [optional] 
**Url** | Pointer to **string** | The requested URL. | [optional] [readonly] 

## Methods

### NewRequestInformation

`func NewRequestInformation() *RequestInformation`

NewRequestInformation instantiates a new RequestInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestInformationWithDefaults

`func NewRequestInformationWithDefaults() *RequestInformation`

NewRequestInformationWithDefaults instantiates a new RequestInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *RequestInformation) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RequestInformation) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RequestInformation) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RequestInformation) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPath

`func (o *RequestInformation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestInformation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestInformation) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RequestInformation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocolDetails

`func (o *RequestInformation) GetProtocolDetails() ProtocolDetails`

GetProtocolDetails returns the ProtocolDetails field if non-nil, zero value otherwise.

### GetProtocolDetailsOk

`func (o *RequestInformation) GetProtocolDetailsOk() (*ProtocolDetails, bool)`

GetProtocolDetailsOk returns a tuple with the ProtocolDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolDetails

`func (o *RequestInformation) SetProtocolDetails(v ProtocolDetails)`

SetProtocolDetails sets ProtocolDetails field to given value.

### HasProtocolDetails

`func (o *RequestInformation) HasProtocolDetails() bool`

HasProtocolDetails returns a boolean if a field has been set.

### GetUrl

`func (o *RequestInformation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestInformation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestInformation) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequestInformation) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


