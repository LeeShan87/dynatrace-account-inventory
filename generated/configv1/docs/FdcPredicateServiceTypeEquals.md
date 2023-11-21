# FdcPredicateServiceTypeEquals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **[]string** | A list of reference values. The condition is fulfilled when the attribute (which is the type of the service) equals *any* of these.The possible values are: WebRequest, WebService, Database, Method, WebSite, Messaging, Mobile, Process, Rmi, External, QueueListener, QueueInteraction, RemoteCall, SaasVendor, AMP, CustomApplication, Cics, Ims, CicsInteraction, ImsInteraction, EnterpriseServiceBus, ZosConnect. | [optional] 

## Methods

### NewFdcPredicateServiceTypeEquals

`func NewFdcPredicateServiceTypeEquals() *FdcPredicateServiceTypeEquals`

NewFdcPredicateServiceTypeEquals instantiates a new FdcPredicateServiceTypeEquals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdcPredicateServiceTypeEqualsWithDefaults

`func NewFdcPredicateServiceTypeEqualsWithDefaults() *FdcPredicateServiceTypeEquals`

NewFdcPredicateServiceTypeEqualsWithDefaults instantiates a new FdcPredicateServiceTypeEquals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *FdcPredicateServiceTypeEquals) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FdcPredicateServiceTypeEquals) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FdcPredicateServiceTypeEquals) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FdcPredicateServiceTypeEquals) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


