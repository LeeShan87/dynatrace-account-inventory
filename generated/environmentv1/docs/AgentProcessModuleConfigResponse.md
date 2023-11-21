# AgentProcessModuleConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to [**[]SectionProperty**](SectionProperty.md) | The properties and their sections in this response. | [optional] 
**Revision** | Pointer to **int64** | The new revision associated with the config. | [optional] 

## Methods

### NewAgentProcessModuleConfigResponse

`func NewAgentProcessModuleConfigResponse() *AgentProcessModuleConfigResponse`

NewAgentProcessModuleConfigResponse instantiates a new AgentProcessModuleConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentProcessModuleConfigResponseWithDefaults

`func NewAgentProcessModuleConfigResponseWithDefaults() *AgentProcessModuleConfigResponse`

NewAgentProcessModuleConfigResponseWithDefaults instantiates a new AgentProcessModuleConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *AgentProcessModuleConfigResponse) GetProperties() []SectionProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AgentProcessModuleConfigResponse) GetPropertiesOk() (*[]SectionProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AgentProcessModuleConfigResponse) SetProperties(v []SectionProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AgentProcessModuleConfigResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetRevision

`func (o *AgentProcessModuleConfigResponse) GetRevision() int64`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AgentProcessModuleConfigResponse) GetRevisionOk() (*int64, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AgentProcessModuleConfigResponse) SetRevision(v int64)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AgentProcessModuleConfigResponse) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


