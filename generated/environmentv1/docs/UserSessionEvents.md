# UserSessionEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | The name of the application, based on the configured detection rules. | [optional] 
**Domain** | Pointer to **string** | The DNS domain where the event has been recorded. | [optional] 
**InternalApplicationId** | Pointer to **string** | The Dynatrace entity ID of the application.    This information is useful when calling various REST APIs, for example, as a key for time series queries. | [optional] 
**Metadata** | Pointer to **string** | The metadata attached to the event. | [optional] 
**Name** | Pointer to **string** | The name of the event. | [optional] 
**Page** | Pointer to **string** | The name of the page the user navigated to during a page change event. | [optional] 
**PageGroup** | Pointer to **string** | The page group is automatically derived from the page. | [optional] 
**PageReferrer** | Pointer to **string** | The name of the previous page from which the user navigated from during a page change event. | [optional] 
**PageReferrerGroup** | Pointer to **string** | The page referrer group is automatically derived from the page referrer. | [optional] 
**StartTime** | Pointer to **int64** | The timestamp of the event, in UTC milliseconds. | [optional] 
**Type** | Pointer to **string** | The type of event. | [optional] 

## Methods

### NewUserSessionEvents

`func NewUserSessionEvents() *UserSessionEvents`

NewUserSessionEvents instantiates a new UserSessionEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionEventsWithDefaults

`func NewUserSessionEventsWithDefaults() *UserSessionEvents`

NewUserSessionEventsWithDefaults instantiates a new UserSessionEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *UserSessionEvents) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *UserSessionEvents) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *UserSessionEvents) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *UserSessionEvents) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDomain

`func (o *UserSessionEvents) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserSessionEvents) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserSessionEvents) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserSessionEvents) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetInternalApplicationId

`func (o *UserSessionEvents) GetInternalApplicationId() string`

GetInternalApplicationId returns the InternalApplicationId field if non-nil, zero value otherwise.

### GetInternalApplicationIdOk

`func (o *UserSessionEvents) GetInternalApplicationIdOk() (*string, bool)`

GetInternalApplicationIdOk returns a tuple with the InternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplicationId

`func (o *UserSessionEvents) SetInternalApplicationId(v string)`

SetInternalApplicationId sets InternalApplicationId field to given value.

### HasInternalApplicationId

`func (o *UserSessionEvents) HasInternalApplicationId() bool`

HasInternalApplicationId returns a boolean if a field has been set.

### GetMetadata

`func (o *UserSessionEvents) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserSessionEvents) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserSessionEvents) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UserSessionEvents) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UserSessionEvents) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSessionEvents) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSessionEvents) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserSessionEvents) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPage

`func (o *UserSessionEvents) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UserSessionEvents) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UserSessionEvents) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *UserSessionEvents) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageGroup

`func (o *UserSessionEvents) GetPageGroup() string`

GetPageGroup returns the PageGroup field if non-nil, zero value otherwise.

### GetPageGroupOk

`func (o *UserSessionEvents) GetPageGroupOk() (*string, bool)`

GetPageGroupOk returns a tuple with the PageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageGroup

`func (o *UserSessionEvents) SetPageGroup(v string)`

SetPageGroup sets PageGroup field to given value.

### HasPageGroup

`func (o *UserSessionEvents) HasPageGroup() bool`

HasPageGroup returns a boolean if a field has been set.

### GetPageReferrer

`func (o *UserSessionEvents) GetPageReferrer() string`

GetPageReferrer returns the PageReferrer field if non-nil, zero value otherwise.

### GetPageReferrerOk

`func (o *UserSessionEvents) GetPageReferrerOk() (*string, bool)`

GetPageReferrerOk returns a tuple with the PageReferrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReferrer

`func (o *UserSessionEvents) SetPageReferrer(v string)`

SetPageReferrer sets PageReferrer field to given value.

### HasPageReferrer

`func (o *UserSessionEvents) HasPageReferrer() bool`

HasPageReferrer returns a boolean if a field has been set.

### GetPageReferrerGroup

`func (o *UserSessionEvents) GetPageReferrerGroup() string`

GetPageReferrerGroup returns the PageReferrerGroup field if non-nil, zero value otherwise.

### GetPageReferrerGroupOk

`func (o *UserSessionEvents) GetPageReferrerGroupOk() (*string, bool)`

GetPageReferrerGroupOk returns a tuple with the PageReferrerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReferrerGroup

`func (o *UserSessionEvents) SetPageReferrerGroup(v string)`

SetPageReferrerGroup sets PageReferrerGroup field to given value.

### HasPageReferrerGroup

`func (o *UserSessionEvents) HasPageReferrerGroup() bool`

HasPageReferrerGroup returns a boolean if a field has been set.

### GetStartTime

`func (o *UserSessionEvents) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UserSessionEvents) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UserSessionEvents) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UserSessionEvents) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetType

`func (o *UserSessionEvents) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionEvents) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionEvents) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserSessionEvents) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


