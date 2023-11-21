# BMAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApdexType** | Pointer to **string** | The user experience index of the action. | [optional] 
**CdnBusyTime** | Pointer to **int32** | The time spent waiting for CDN resources for the action, in milliseconds. | [optional] 
**CdnResources** | Pointer to **int32** | The number of resources fetched from a CDN for the action. | [optional] 
**ClientTime** | Pointer to **int64** | The event startTime in client time, in milliseconds. | [optional] 
**CumulativeLayoutShift** | Pointer to **float64** | Cumulative layout shift: Available for Chromium-based browsers. Measured using Google-provided APIs. | [optional] 
**CustomErrorCount** | Pointer to **int32** | The total number of custom errors during the action. | [optional] 
**DocumentInteractiveTime** | Pointer to **int32** | The amount of time spent until the document for the action became interactive, in milliseconds. | [optional] 
**DomCompleteTime** | Pointer to **int32** | The amount of time until the DOM tree is completed, in milliseconds. | [optional] 
**DomContentLoadedTime** | Pointer to **int32** | The amount of time until the DOM tree is loaded, in milliseconds. | [optional] 
**Domain** | Pointer to **string** | The DNS domain where the action has been recorded | [optional] 
**Duration** | Pointer to **int64** | The duration of the action, in milliseconds | [optional] 
**EndTime** | Pointer to **int64** | The stop time of the action on the server, in UTC milliseconds | [optional] 
**EntryAction** | Pointer to **bool** |  | [optional] 
**ExitAction** | Pointer to **bool** |  | [optional] 
**FirstInputDelay** | Pointer to **int32** | The first input delay (FID) is the time (in milliseconds) that the browser took to respond to the first user input. | [optional] 
**FirstPartyBusyTime** | Pointer to **int32** | The time spent waiting for resources from the originating server for the action, in milliseconds. | [optional] 
**FirstPartyResources** | Pointer to **int32** | The number of resources fetched from the originating server for the action. | [optional] 
**FrontendTime** | Pointer to **int32** | The amount of time spent on the frontend rendering for the action, in milliseconds. | [optional] 
**JavascriptErrorCount** | Pointer to **int32** | The total number of Javascript errors during the action. | [optional] 
**LargestContentfulPaint** | Pointer to **int32** | The largest contentful paint (LCP) is the time (in milliseconds) that the largest element on the page took to render. | [optional] 
**LoadEventEnd** | Pointer to **int32** | The amount of time until the load event ended, in milliseconds. | [optional] 
**LoadEventStart** | Pointer to **int32** | The amount of time until the load event started, in milliseconds. | [optional] 
**Name** | Pointer to **string** | The name of the action. | [optional] 
**NavigationStartTime** | Pointer to **int64** | The timestamp of the navigation start, in UTC milliseconds. | [optional] 
**NetworkTime** | Pointer to **int32** | The amount of time spent on the data transfer for the action, in milliseconds. | [optional] 
**Referrer** | Pointer to **string** | The referrer. | [optional] 
**RequestErrorCount** | Pointer to **int32** | The total number of request errors during the action. | [optional] 
**RequestStart** | Pointer to **int32** | The amount of time until the request started, in milliseconds. | [optional] 
**ResponseEnd** | Pointer to **int32** | The amount of time until the response ended, in milliseconds. | [optional] 
**ResponseStart** | Pointer to **int32** | The amount of time until the response started, in milliseconds. | [optional] 
**ServerTime** | Pointer to **int32** | The amount of time spent on the server-side processing for the action, in milliseconds. | [optional] 
**SpeedIndex** | Pointer to **int32** | A score indicating how quickly the page content is visually populated. A low speed index means that most parts of a page are rendering quickly. | [optional] 
**StartSequenceNumber** | Pointer to **int32** | The sequence number of the action (to get a kind of order). | [optional] 
**StartTime** | Pointer to **int64** | The start time of the action on the server, in in UTC milliseconds. | [optional] 
**TargetUrl** | Pointer to **string** | The URL of the action. | [optional] 
**ThirdPartyBusyTime** | Pointer to **int32** | The time spent waiting for third party resources for the action, in milliseconds. | [optional] 
**ThirdPartyResources** | Pointer to **int32** | The number of third party resources loaded for the action. | [optional] 
**TotalBlockingTime** | Pointer to **int32** | The time between the moment when the browser receives a request to download a resource and the time that it actually starts downloading the resource in ms. | [optional] 
**Type** | Pointer to **string** | The type of the action. | [optional] 
**UserActionPropertyCount** | Pointer to **int32** | The total number of properties in the action. | [optional] 
**VisuallyCompleteTime** | Pointer to **int32** | The amount of time until the page is visually complete, in milliseconds. | [optional] 

## Methods

### NewBMAction

`func NewBMAction() *BMAction`

NewBMAction instantiates a new BMAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBMActionWithDefaults

`func NewBMActionWithDefaults() *BMAction`

NewBMActionWithDefaults instantiates a new BMAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApdexType

`func (o *BMAction) GetApdexType() string`

GetApdexType returns the ApdexType field if non-nil, zero value otherwise.

### GetApdexTypeOk

`func (o *BMAction) GetApdexTypeOk() (*string, bool)`

GetApdexTypeOk returns a tuple with the ApdexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApdexType

`func (o *BMAction) SetApdexType(v string)`

SetApdexType sets ApdexType field to given value.

### HasApdexType

`func (o *BMAction) HasApdexType() bool`

HasApdexType returns a boolean if a field has been set.

### GetCdnBusyTime

`func (o *BMAction) GetCdnBusyTime() int32`

GetCdnBusyTime returns the CdnBusyTime field if non-nil, zero value otherwise.

### GetCdnBusyTimeOk

`func (o *BMAction) GetCdnBusyTimeOk() (*int32, bool)`

GetCdnBusyTimeOk returns a tuple with the CdnBusyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnBusyTime

`func (o *BMAction) SetCdnBusyTime(v int32)`

SetCdnBusyTime sets CdnBusyTime field to given value.

### HasCdnBusyTime

`func (o *BMAction) HasCdnBusyTime() bool`

HasCdnBusyTime returns a boolean if a field has been set.

### GetCdnResources

`func (o *BMAction) GetCdnResources() int32`

GetCdnResources returns the CdnResources field if non-nil, zero value otherwise.

### GetCdnResourcesOk

`func (o *BMAction) GetCdnResourcesOk() (*int32, bool)`

GetCdnResourcesOk returns a tuple with the CdnResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnResources

`func (o *BMAction) SetCdnResources(v int32)`

SetCdnResources sets CdnResources field to given value.

### HasCdnResources

`func (o *BMAction) HasCdnResources() bool`

HasCdnResources returns a boolean if a field has been set.

### GetClientTime

`func (o *BMAction) GetClientTime() int64`

GetClientTime returns the ClientTime field if non-nil, zero value otherwise.

### GetClientTimeOk

`func (o *BMAction) GetClientTimeOk() (*int64, bool)`

GetClientTimeOk returns a tuple with the ClientTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTime

`func (o *BMAction) SetClientTime(v int64)`

SetClientTime sets ClientTime field to given value.

### HasClientTime

`func (o *BMAction) HasClientTime() bool`

HasClientTime returns a boolean if a field has been set.

### GetCumulativeLayoutShift

`func (o *BMAction) GetCumulativeLayoutShift() float64`

GetCumulativeLayoutShift returns the CumulativeLayoutShift field if non-nil, zero value otherwise.

### GetCumulativeLayoutShiftOk

`func (o *BMAction) GetCumulativeLayoutShiftOk() (*float64, bool)`

GetCumulativeLayoutShiftOk returns a tuple with the CumulativeLayoutShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeLayoutShift

`func (o *BMAction) SetCumulativeLayoutShift(v float64)`

SetCumulativeLayoutShift sets CumulativeLayoutShift field to given value.

### HasCumulativeLayoutShift

`func (o *BMAction) HasCumulativeLayoutShift() bool`

HasCumulativeLayoutShift returns a boolean if a field has been set.

### GetCustomErrorCount

`func (o *BMAction) GetCustomErrorCount() int32`

GetCustomErrorCount returns the CustomErrorCount field if non-nil, zero value otherwise.

### GetCustomErrorCountOk

`func (o *BMAction) GetCustomErrorCountOk() (*int32, bool)`

GetCustomErrorCountOk returns a tuple with the CustomErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomErrorCount

`func (o *BMAction) SetCustomErrorCount(v int32)`

SetCustomErrorCount sets CustomErrorCount field to given value.

### HasCustomErrorCount

`func (o *BMAction) HasCustomErrorCount() bool`

HasCustomErrorCount returns a boolean if a field has been set.

### GetDocumentInteractiveTime

`func (o *BMAction) GetDocumentInteractiveTime() int32`

GetDocumentInteractiveTime returns the DocumentInteractiveTime field if non-nil, zero value otherwise.

### GetDocumentInteractiveTimeOk

`func (o *BMAction) GetDocumentInteractiveTimeOk() (*int32, bool)`

GetDocumentInteractiveTimeOk returns a tuple with the DocumentInteractiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentInteractiveTime

`func (o *BMAction) SetDocumentInteractiveTime(v int32)`

SetDocumentInteractiveTime sets DocumentInteractiveTime field to given value.

### HasDocumentInteractiveTime

`func (o *BMAction) HasDocumentInteractiveTime() bool`

HasDocumentInteractiveTime returns a boolean if a field has been set.

### GetDomCompleteTime

`func (o *BMAction) GetDomCompleteTime() int32`

GetDomCompleteTime returns the DomCompleteTime field if non-nil, zero value otherwise.

### GetDomCompleteTimeOk

`func (o *BMAction) GetDomCompleteTimeOk() (*int32, bool)`

GetDomCompleteTimeOk returns a tuple with the DomCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomCompleteTime

`func (o *BMAction) SetDomCompleteTime(v int32)`

SetDomCompleteTime sets DomCompleteTime field to given value.

### HasDomCompleteTime

`func (o *BMAction) HasDomCompleteTime() bool`

HasDomCompleteTime returns a boolean if a field has been set.

### GetDomContentLoadedTime

`func (o *BMAction) GetDomContentLoadedTime() int32`

GetDomContentLoadedTime returns the DomContentLoadedTime field if non-nil, zero value otherwise.

### GetDomContentLoadedTimeOk

`func (o *BMAction) GetDomContentLoadedTimeOk() (*int32, bool)`

GetDomContentLoadedTimeOk returns a tuple with the DomContentLoadedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomContentLoadedTime

`func (o *BMAction) SetDomContentLoadedTime(v int32)`

SetDomContentLoadedTime sets DomContentLoadedTime field to given value.

### HasDomContentLoadedTime

`func (o *BMAction) HasDomContentLoadedTime() bool`

HasDomContentLoadedTime returns a boolean if a field has been set.

### GetDomain

`func (o *BMAction) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BMAction) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BMAction) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *BMAction) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDuration

`func (o *BMAction) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BMAction) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BMAction) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BMAction) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndTime

`func (o *BMAction) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BMAction) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BMAction) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BMAction) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntryAction

`func (o *BMAction) GetEntryAction() bool`

GetEntryAction returns the EntryAction field if non-nil, zero value otherwise.

### GetEntryActionOk

`func (o *BMAction) GetEntryActionOk() (*bool, bool)`

GetEntryActionOk returns a tuple with the EntryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryAction

`func (o *BMAction) SetEntryAction(v bool)`

SetEntryAction sets EntryAction field to given value.

### HasEntryAction

`func (o *BMAction) HasEntryAction() bool`

HasEntryAction returns a boolean if a field has been set.

### GetExitAction

`func (o *BMAction) GetExitAction() bool`

GetExitAction returns the ExitAction field if non-nil, zero value otherwise.

### GetExitActionOk

`func (o *BMAction) GetExitActionOk() (*bool, bool)`

GetExitActionOk returns a tuple with the ExitAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitAction

`func (o *BMAction) SetExitAction(v bool)`

SetExitAction sets ExitAction field to given value.

### HasExitAction

`func (o *BMAction) HasExitAction() bool`

HasExitAction returns a boolean if a field has been set.

### GetFirstInputDelay

`func (o *BMAction) GetFirstInputDelay() int32`

GetFirstInputDelay returns the FirstInputDelay field if non-nil, zero value otherwise.

### GetFirstInputDelayOk

`func (o *BMAction) GetFirstInputDelayOk() (*int32, bool)`

GetFirstInputDelayOk returns a tuple with the FirstInputDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstInputDelay

`func (o *BMAction) SetFirstInputDelay(v int32)`

SetFirstInputDelay sets FirstInputDelay field to given value.

### HasFirstInputDelay

`func (o *BMAction) HasFirstInputDelay() bool`

HasFirstInputDelay returns a boolean if a field has been set.

### GetFirstPartyBusyTime

`func (o *BMAction) GetFirstPartyBusyTime() int32`

GetFirstPartyBusyTime returns the FirstPartyBusyTime field if non-nil, zero value otherwise.

### GetFirstPartyBusyTimeOk

`func (o *BMAction) GetFirstPartyBusyTimeOk() (*int32, bool)`

GetFirstPartyBusyTimeOk returns a tuple with the FirstPartyBusyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyBusyTime

`func (o *BMAction) SetFirstPartyBusyTime(v int32)`

SetFirstPartyBusyTime sets FirstPartyBusyTime field to given value.

### HasFirstPartyBusyTime

`func (o *BMAction) HasFirstPartyBusyTime() bool`

HasFirstPartyBusyTime returns a boolean if a field has been set.

### GetFirstPartyResources

`func (o *BMAction) GetFirstPartyResources() int32`

GetFirstPartyResources returns the FirstPartyResources field if non-nil, zero value otherwise.

### GetFirstPartyResourcesOk

`func (o *BMAction) GetFirstPartyResourcesOk() (*int32, bool)`

GetFirstPartyResourcesOk returns a tuple with the FirstPartyResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPartyResources

`func (o *BMAction) SetFirstPartyResources(v int32)`

SetFirstPartyResources sets FirstPartyResources field to given value.

### HasFirstPartyResources

`func (o *BMAction) HasFirstPartyResources() bool`

HasFirstPartyResources returns a boolean if a field has been set.

### GetFrontendTime

`func (o *BMAction) GetFrontendTime() int32`

GetFrontendTime returns the FrontendTime field if non-nil, zero value otherwise.

### GetFrontendTimeOk

`func (o *BMAction) GetFrontendTimeOk() (*int32, bool)`

GetFrontendTimeOk returns a tuple with the FrontendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontendTime

`func (o *BMAction) SetFrontendTime(v int32)`

SetFrontendTime sets FrontendTime field to given value.

### HasFrontendTime

`func (o *BMAction) HasFrontendTime() bool`

HasFrontendTime returns a boolean if a field has been set.

### GetJavascriptErrorCount

`func (o *BMAction) GetJavascriptErrorCount() int32`

GetJavascriptErrorCount returns the JavascriptErrorCount field if non-nil, zero value otherwise.

### GetJavascriptErrorCountOk

`func (o *BMAction) GetJavascriptErrorCountOk() (*int32, bool)`

GetJavascriptErrorCountOk returns a tuple with the JavascriptErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptErrorCount

`func (o *BMAction) SetJavascriptErrorCount(v int32)`

SetJavascriptErrorCount sets JavascriptErrorCount field to given value.

### HasJavascriptErrorCount

`func (o *BMAction) HasJavascriptErrorCount() bool`

HasJavascriptErrorCount returns a boolean if a field has been set.

### GetLargestContentfulPaint

`func (o *BMAction) GetLargestContentfulPaint() int32`

GetLargestContentfulPaint returns the LargestContentfulPaint field if non-nil, zero value otherwise.

### GetLargestContentfulPaintOk

`func (o *BMAction) GetLargestContentfulPaintOk() (*int32, bool)`

GetLargestContentfulPaintOk returns a tuple with the LargestContentfulPaint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargestContentfulPaint

`func (o *BMAction) SetLargestContentfulPaint(v int32)`

SetLargestContentfulPaint sets LargestContentfulPaint field to given value.

### HasLargestContentfulPaint

`func (o *BMAction) HasLargestContentfulPaint() bool`

HasLargestContentfulPaint returns a boolean if a field has been set.

### GetLoadEventEnd

`func (o *BMAction) GetLoadEventEnd() int32`

GetLoadEventEnd returns the LoadEventEnd field if non-nil, zero value otherwise.

### GetLoadEventEndOk

`func (o *BMAction) GetLoadEventEndOk() (*int32, bool)`

GetLoadEventEndOk returns a tuple with the LoadEventEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadEventEnd

`func (o *BMAction) SetLoadEventEnd(v int32)`

SetLoadEventEnd sets LoadEventEnd field to given value.

### HasLoadEventEnd

`func (o *BMAction) HasLoadEventEnd() bool`

HasLoadEventEnd returns a boolean if a field has been set.

### GetLoadEventStart

`func (o *BMAction) GetLoadEventStart() int32`

GetLoadEventStart returns the LoadEventStart field if non-nil, zero value otherwise.

### GetLoadEventStartOk

`func (o *BMAction) GetLoadEventStartOk() (*int32, bool)`

GetLoadEventStartOk returns a tuple with the LoadEventStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadEventStart

`func (o *BMAction) SetLoadEventStart(v int32)`

SetLoadEventStart sets LoadEventStart field to given value.

### HasLoadEventStart

`func (o *BMAction) HasLoadEventStart() bool`

HasLoadEventStart returns a boolean if a field has been set.

### GetName

`func (o *BMAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BMAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BMAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BMAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavigationStartTime

`func (o *BMAction) GetNavigationStartTime() int64`

GetNavigationStartTime returns the NavigationStartTime field if non-nil, zero value otherwise.

### GetNavigationStartTimeOk

`func (o *BMAction) GetNavigationStartTimeOk() (*int64, bool)`

GetNavigationStartTimeOk returns a tuple with the NavigationStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationStartTime

`func (o *BMAction) SetNavigationStartTime(v int64)`

SetNavigationStartTime sets NavigationStartTime field to given value.

### HasNavigationStartTime

`func (o *BMAction) HasNavigationStartTime() bool`

HasNavigationStartTime returns a boolean if a field has been set.

### GetNetworkTime

`func (o *BMAction) GetNetworkTime() int32`

GetNetworkTime returns the NetworkTime field if non-nil, zero value otherwise.

### GetNetworkTimeOk

`func (o *BMAction) GetNetworkTimeOk() (*int32, bool)`

GetNetworkTimeOk returns a tuple with the NetworkTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTime

`func (o *BMAction) SetNetworkTime(v int32)`

SetNetworkTime sets NetworkTime field to given value.

### HasNetworkTime

`func (o *BMAction) HasNetworkTime() bool`

HasNetworkTime returns a boolean if a field has been set.

### GetReferrer

`func (o *BMAction) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *BMAction) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *BMAction) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *BMAction) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetRequestErrorCount

`func (o *BMAction) GetRequestErrorCount() int32`

GetRequestErrorCount returns the RequestErrorCount field if non-nil, zero value otherwise.

### GetRequestErrorCountOk

`func (o *BMAction) GetRequestErrorCountOk() (*int32, bool)`

GetRequestErrorCountOk returns a tuple with the RequestErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestErrorCount

`func (o *BMAction) SetRequestErrorCount(v int32)`

SetRequestErrorCount sets RequestErrorCount field to given value.

### HasRequestErrorCount

`func (o *BMAction) HasRequestErrorCount() bool`

HasRequestErrorCount returns a boolean if a field has been set.

### GetRequestStart

`func (o *BMAction) GetRequestStart() int32`

GetRequestStart returns the RequestStart field if non-nil, zero value otherwise.

### GetRequestStartOk

`func (o *BMAction) GetRequestStartOk() (*int32, bool)`

GetRequestStartOk returns a tuple with the RequestStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStart

`func (o *BMAction) SetRequestStart(v int32)`

SetRequestStart sets RequestStart field to given value.

### HasRequestStart

`func (o *BMAction) HasRequestStart() bool`

HasRequestStart returns a boolean if a field has been set.

### GetResponseEnd

`func (o *BMAction) GetResponseEnd() int32`

GetResponseEnd returns the ResponseEnd field if non-nil, zero value otherwise.

### GetResponseEndOk

`func (o *BMAction) GetResponseEndOk() (*int32, bool)`

GetResponseEndOk returns a tuple with the ResponseEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseEnd

`func (o *BMAction) SetResponseEnd(v int32)`

SetResponseEnd sets ResponseEnd field to given value.

### HasResponseEnd

`func (o *BMAction) HasResponseEnd() bool`

HasResponseEnd returns a boolean if a field has been set.

### GetResponseStart

`func (o *BMAction) GetResponseStart() int32`

GetResponseStart returns the ResponseStart field if non-nil, zero value otherwise.

### GetResponseStartOk

`func (o *BMAction) GetResponseStartOk() (*int32, bool)`

GetResponseStartOk returns a tuple with the ResponseStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStart

`func (o *BMAction) SetResponseStart(v int32)`

SetResponseStart sets ResponseStart field to given value.

### HasResponseStart

`func (o *BMAction) HasResponseStart() bool`

HasResponseStart returns a boolean if a field has been set.

### GetServerTime

`func (o *BMAction) GetServerTime() int32`

GetServerTime returns the ServerTime field if non-nil, zero value otherwise.

### GetServerTimeOk

`func (o *BMAction) GetServerTimeOk() (*int32, bool)`

GetServerTimeOk returns a tuple with the ServerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTime

`func (o *BMAction) SetServerTime(v int32)`

SetServerTime sets ServerTime field to given value.

### HasServerTime

`func (o *BMAction) HasServerTime() bool`

HasServerTime returns a boolean if a field has been set.

### GetSpeedIndex

`func (o *BMAction) GetSpeedIndex() int32`

GetSpeedIndex returns the SpeedIndex field if non-nil, zero value otherwise.

### GetSpeedIndexOk

`func (o *BMAction) GetSpeedIndexOk() (*int32, bool)`

GetSpeedIndexOk returns a tuple with the SpeedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedIndex

`func (o *BMAction) SetSpeedIndex(v int32)`

SetSpeedIndex sets SpeedIndex field to given value.

### HasSpeedIndex

`func (o *BMAction) HasSpeedIndex() bool`

HasSpeedIndex returns a boolean if a field has been set.

### GetStartSequenceNumber

`func (o *BMAction) GetStartSequenceNumber() int32`

GetStartSequenceNumber returns the StartSequenceNumber field if non-nil, zero value otherwise.

### GetStartSequenceNumberOk

`func (o *BMAction) GetStartSequenceNumberOk() (*int32, bool)`

GetStartSequenceNumberOk returns a tuple with the StartSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSequenceNumber

`func (o *BMAction) SetStartSequenceNumber(v int32)`

SetStartSequenceNumber sets StartSequenceNumber field to given value.

### HasStartSequenceNumber

`func (o *BMAction) HasStartSequenceNumber() bool`

HasStartSequenceNumber returns a boolean if a field has been set.

### GetStartTime

`func (o *BMAction) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BMAction) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BMAction) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BMAction) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTargetUrl

`func (o *BMAction) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *BMAction) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *BMAction) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *BMAction) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetThirdPartyBusyTime

`func (o *BMAction) GetThirdPartyBusyTime() int32`

GetThirdPartyBusyTime returns the ThirdPartyBusyTime field if non-nil, zero value otherwise.

### GetThirdPartyBusyTimeOk

`func (o *BMAction) GetThirdPartyBusyTimeOk() (*int32, bool)`

GetThirdPartyBusyTimeOk returns a tuple with the ThirdPartyBusyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyBusyTime

`func (o *BMAction) SetThirdPartyBusyTime(v int32)`

SetThirdPartyBusyTime sets ThirdPartyBusyTime field to given value.

### HasThirdPartyBusyTime

`func (o *BMAction) HasThirdPartyBusyTime() bool`

HasThirdPartyBusyTime returns a boolean if a field has been set.

### GetThirdPartyResources

`func (o *BMAction) GetThirdPartyResources() int32`

GetThirdPartyResources returns the ThirdPartyResources field if non-nil, zero value otherwise.

### GetThirdPartyResourcesOk

`func (o *BMAction) GetThirdPartyResourcesOk() (*int32, bool)`

GetThirdPartyResourcesOk returns a tuple with the ThirdPartyResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyResources

`func (o *BMAction) SetThirdPartyResources(v int32)`

SetThirdPartyResources sets ThirdPartyResources field to given value.

### HasThirdPartyResources

`func (o *BMAction) HasThirdPartyResources() bool`

HasThirdPartyResources returns a boolean if a field has been set.

### GetTotalBlockingTime

`func (o *BMAction) GetTotalBlockingTime() int32`

GetTotalBlockingTime returns the TotalBlockingTime field if non-nil, zero value otherwise.

### GetTotalBlockingTimeOk

`func (o *BMAction) GetTotalBlockingTimeOk() (*int32, bool)`

GetTotalBlockingTimeOk returns a tuple with the TotalBlockingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBlockingTime

`func (o *BMAction) SetTotalBlockingTime(v int32)`

SetTotalBlockingTime sets TotalBlockingTime field to given value.

### HasTotalBlockingTime

`func (o *BMAction) HasTotalBlockingTime() bool`

HasTotalBlockingTime returns a boolean if a field has been set.

### GetType

`func (o *BMAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BMAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BMAction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BMAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserActionPropertyCount

`func (o *BMAction) GetUserActionPropertyCount() int32`

GetUserActionPropertyCount returns the UserActionPropertyCount field if non-nil, zero value otherwise.

### GetUserActionPropertyCountOk

`func (o *BMAction) GetUserActionPropertyCountOk() (*int32, bool)`

GetUserActionPropertyCountOk returns a tuple with the UserActionPropertyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionPropertyCount

`func (o *BMAction) SetUserActionPropertyCount(v int32)`

SetUserActionPropertyCount sets UserActionPropertyCount field to given value.

### HasUserActionPropertyCount

`func (o *BMAction) HasUserActionPropertyCount() bool`

HasUserActionPropertyCount returns a boolean if a field has been set.

### GetVisuallyCompleteTime

`func (o *BMAction) GetVisuallyCompleteTime() int32`

GetVisuallyCompleteTime returns the VisuallyCompleteTime field if non-nil, zero value otherwise.

### GetVisuallyCompleteTimeOk

`func (o *BMAction) GetVisuallyCompleteTimeOk() (*int32, bool)`

GetVisuallyCompleteTimeOk returns a tuple with the VisuallyCompleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisuallyCompleteTime

`func (o *BMAction) SetVisuallyCompleteTime(v int32)`

SetVisuallyCompleteTime sets VisuallyCompleteTime field to given value.

### HasVisuallyCompleteTime

`func (o *BMAction) HasVisuallyCompleteTime() bool`

HasVisuallyCompleteTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


