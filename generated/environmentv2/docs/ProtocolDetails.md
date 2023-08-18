# ProtocolDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Http** | Pointer to [**HttpProtocolDetails**](HttpProtocolDetails.md) |  | [optional] 

## Methods

### NewProtocolDetails

`func NewProtocolDetails() *ProtocolDetails`

NewProtocolDetails instantiates a new ProtocolDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolDetailsWithDefaults

`func NewProtocolDetailsWithDefaults() *ProtocolDetails`

NewProtocolDetailsWithDefaults instantiates a new ProtocolDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttp

`func (o *ProtocolDetails) GetHttp() HttpProtocolDetails`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ProtocolDetails) GetHttpOk() (*HttpProtocolDetails, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ProtocolDetails) SetHttp(v HttpProtocolDetails)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ProtocolDetails) HasHttp() bool`

HasHttp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


