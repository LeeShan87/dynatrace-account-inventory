# AnonymizationIdResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterRequestIds** | Pointer to [**[]AnonymizationClusterRequestID**](AnonymizationClusterRequestID.md) | A list of tuples of request ID and cluster name | [optional] 
**RequestId** | Pointer to **string** | The ID of the newly created anonymization job. If multiple datacenters are involved a list separated by \&quot;|\&quot; will be returned | [optional] 

## Methods

### NewAnonymizationIdResult

`func NewAnonymizationIdResult() *AnonymizationIdResult`

NewAnonymizationIdResult instantiates a new AnonymizationIdResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnonymizationIdResultWithDefaults

`func NewAnonymizationIdResultWithDefaults() *AnonymizationIdResult`

NewAnonymizationIdResultWithDefaults instantiates a new AnonymizationIdResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterRequestIds

`func (o *AnonymizationIdResult) GetClusterRequestIds() []AnonymizationClusterRequestID`

GetClusterRequestIds returns the ClusterRequestIds field if non-nil, zero value otherwise.

### GetClusterRequestIdsOk

`func (o *AnonymizationIdResult) GetClusterRequestIdsOk() (*[]AnonymizationClusterRequestID, bool)`

GetClusterRequestIdsOk returns a tuple with the ClusterRequestIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRequestIds

`func (o *AnonymizationIdResult) SetClusterRequestIds(v []AnonymizationClusterRequestID)`

SetClusterRequestIds sets ClusterRequestIds field to given value.

### HasClusterRequestIds

`func (o *AnonymizationIdResult) HasClusterRequestIds() bool`

HasClusterRequestIds returns a boolean if a field has been set.

### GetRequestId

`func (o *AnonymizationIdResult) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AnonymizationIdResult) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AnonymizationIdResult) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AnonymizationIdResult) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


