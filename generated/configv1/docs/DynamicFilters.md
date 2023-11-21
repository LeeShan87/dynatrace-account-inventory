# DynamicFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | **[]string** | A set of all possible global dashboard filters that can be applied to a dashboard   Currently supported values are:    OS_TYPE,  SERVICE_TYPE,  DEPLOYMENT_TYPE,  APPLICATION_INJECTION_TYPE,  PAAS_VENDOR_TYPE,  DATABASE_VENDOR,  HOST_VIRTUALIZATION_TYPE,  HOST_MONITORING_MODE,  KUBERNETES_CLUSTER,  RELATED_CLOUD_APPLICATION,  RELATED_NAMESPACE,  SERVICE_TAG_KEY:&lt;tagname&gt;,  HOST_TAG_KEY:&lt;tagname&gt;,  APPLICATION_TAG_KEY:&lt;tagname&gt;,  CUSTOM_DIMENSION:&lt;key&gt;,  PROCESS_GROUP_TAG_KEY:&lt;tagname&gt;,  PROCESS_GROUP_INSTANCE_TAG_KEY:&lt;tagname&gt; | 
**GenericTagFilters** | [**[]DashboardGenericTagFilter**](DashboardGenericTagFilter.md) | A set of generic tag filters that can be applied to a dashboard | 

## Methods

### NewDynamicFilters

`func NewDynamicFilters(filters []string, genericTagFilters []DashboardGenericTagFilter, ) *DynamicFilters`

NewDynamicFilters instantiates a new DynamicFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicFiltersWithDefaults

`func NewDynamicFiltersWithDefaults() *DynamicFilters`

NewDynamicFiltersWithDefaults instantiates a new DynamicFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *DynamicFilters) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DynamicFilters) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DynamicFilters) SetFilters(v []string)`

SetFilters sets Filters field to given value.


### GetGenericTagFilters

`func (o *DynamicFilters) GetGenericTagFilters() []DashboardGenericTagFilter`

GetGenericTagFilters returns the GenericTagFilters field if non-nil, zero value otherwise.

### GetGenericTagFiltersOk

`func (o *DynamicFilters) GetGenericTagFiltersOk() (*[]DashboardGenericTagFilter, bool)`

GetGenericTagFiltersOk returns a tuple with the GenericTagFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericTagFilters

`func (o *DynamicFilters) SetGenericTagFilters(v []DashboardGenericTagFilter)`

SetGenericTagFilters sets GenericTagFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


