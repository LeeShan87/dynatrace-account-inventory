# KubernetesConfigShortRepresentation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the Dynatrace entity. | [optional] [readonly] 
**EndpointUrl** | **string** | The URL of the Kubernetes API server.   It must be unique within a Dynatrace environment.   The URL must valid according to RFC 2396. Leading or trailing whitespaces are not allowed. | 
**Id** | **string** | The ID of the Dynatrace entity. | 
**Name** | Pointer to **string** | The name of the Dynatrace entity. | [optional] [readonly] 

## Methods

### NewKubernetesConfigShortRepresentation

`func NewKubernetesConfigShortRepresentation(endpointUrl string, id string, ) *KubernetesConfigShortRepresentation`

NewKubernetesConfigShortRepresentation instantiates a new KubernetesConfigShortRepresentation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesConfigShortRepresentationWithDefaults

`func NewKubernetesConfigShortRepresentationWithDefaults() *KubernetesConfigShortRepresentation`

NewKubernetesConfigShortRepresentationWithDefaults instantiates a new KubernetesConfigShortRepresentation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KubernetesConfigShortRepresentation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesConfigShortRepresentation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesConfigShortRepresentation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesConfigShortRepresentation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointUrl

`func (o *KubernetesConfigShortRepresentation) GetEndpointUrl() string`

GetEndpointUrl returns the EndpointUrl field if non-nil, zero value otherwise.

### GetEndpointUrlOk

`func (o *KubernetesConfigShortRepresentation) GetEndpointUrlOk() (*string, bool)`

GetEndpointUrlOk returns a tuple with the EndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUrl

`func (o *KubernetesConfigShortRepresentation) SetEndpointUrl(v string)`

SetEndpointUrl sets EndpointUrl field to given value.


### GetId

`func (o *KubernetesConfigShortRepresentation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesConfigShortRepresentation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesConfigShortRepresentation) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *KubernetesConfigShortRepresentation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesConfigShortRepresentation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesConfigShortRepresentation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesConfigShortRepresentation) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


