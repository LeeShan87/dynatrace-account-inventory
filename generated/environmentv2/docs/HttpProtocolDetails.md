# HttpProtocolDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**TruncatableListAttackRequestHeader**](TruncatableListAttackRequestHeader.md) |  | [optional] 
**Parameters** | Pointer to [**TruncatableListHttpRequestParameter**](TruncatableListHttpRequestParameter.md) |  | [optional] 
**RequestMethod** | Pointer to **string** | The HTTP request method. | [optional] [readonly] 

## Methods

### NewHttpProtocolDetails

`func NewHttpProtocolDetails() *HttpProtocolDetails`

NewHttpProtocolDetails instantiates a new HttpProtocolDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpProtocolDetailsWithDefaults

`func NewHttpProtocolDetailsWithDefaults() *HttpProtocolDetails`

NewHttpProtocolDetailsWithDefaults instantiates a new HttpProtocolDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *HttpProtocolDetails) GetHeaders() TruncatableListAttackRequestHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpProtocolDetails) GetHeadersOk() (*TruncatableListAttackRequestHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpProtocolDetails) SetHeaders(v TruncatableListAttackRequestHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpProtocolDetails) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetParameters

`func (o *HttpProtocolDetails) GetParameters() TruncatableListHttpRequestParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *HttpProtocolDetails) GetParametersOk() (*TruncatableListHttpRequestParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *HttpProtocolDetails) SetParameters(v TruncatableListHttpRequestParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *HttpProtocolDetails) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRequestMethod

`func (o *HttpProtocolDetails) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *HttpProtocolDetails) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *HttpProtocolDetails) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *HttpProtocolDetails) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


