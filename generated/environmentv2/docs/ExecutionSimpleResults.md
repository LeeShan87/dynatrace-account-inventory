# ExecutionSimpleResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChromeError** | Pointer to **bool** | Informs whether is Chrome error. | [optional] 
**EngineId** | Pointer to **int64** | Synthetic engine id on which monitor was executed. | [optional] 
**ErrorCode** | Pointer to **string** | Error code. | [optional] 
**ExecutedSteps** | Pointer to **int32** | Number of the executed steps by Synthetic engine | [optional] 
**FailureMessage** | Pointer to **string** | Failure message. | [optional] 
**HostNameResolutionTime** | Pointer to **int64** | A hostname resolution time measured in milliseconds. | [optional] 
**Httperror** | Pointer to **bool** | Informs whether is HTTP error. | [optional] 
**PeerCertificateExpiryDate** | Pointer to **int64** | An expiry date of the first SSL certificate from the certificate chain. | [optional] 
**PublicLocation** | Pointer to **bool** | Flag informs whether request was executed on public location. | [optional] 
**RedirectionTime** | Pointer to **int64** | Total number of milliseconds spent on handling all redirect requests, measured in milliseconds. | [optional] 
**RedirectsCount** | Pointer to **int32** | Number of redirects. | [optional] 
**ResponseBodySizeLimitExceeded** | Pointer to **bool** | A flag indicating that the response payload size limit of 10MB has been exceeded. | [optional] 
**ResponseSize** | Pointer to **int64** | Request&#39;s response size in bytes. | [optional] 
**ResponseStatusCode** | Pointer to **int32** | Response status code. | [optional] 
**StartTimestamp** | Pointer to **int64** | Start timestamp. | [optional] 
**Status** | Pointer to **string** | Execution status. | [optional] 
**TcpConnectTime** | Pointer to **int64** | A TCP connect time measured in milliseconds. | [optional] 
**TimeToFirstByte** | Pointer to **int64** | A time to first byte measured in milliseconds. | [optional] 
**TlsHandshakeTime** | Pointer to **int64** | A TLS handshake time measured in milliseconds. | [optional] 
**TotalTime** | Pointer to **int64** | A total time measured in milliseconds. | [optional] 

## Methods

### NewExecutionSimpleResults

`func NewExecutionSimpleResults() *ExecutionSimpleResults`

NewExecutionSimpleResults instantiates a new ExecutionSimpleResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionSimpleResultsWithDefaults

`func NewExecutionSimpleResultsWithDefaults() *ExecutionSimpleResults`

NewExecutionSimpleResultsWithDefaults instantiates a new ExecutionSimpleResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChromeError

`func (o *ExecutionSimpleResults) GetChromeError() bool`

GetChromeError returns the ChromeError field if non-nil, zero value otherwise.

### GetChromeErrorOk

`func (o *ExecutionSimpleResults) GetChromeErrorOk() (*bool, bool)`

GetChromeErrorOk returns a tuple with the ChromeError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromeError

`func (o *ExecutionSimpleResults) SetChromeError(v bool)`

SetChromeError sets ChromeError field to given value.

### HasChromeError

`func (o *ExecutionSimpleResults) HasChromeError() bool`

HasChromeError returns a boolean if a field has been set.

### GetEngineId

`func (o *ExecutionSimpleResults) GetEngineId() int64`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *ExecutionSimpleResults) GetEngineIdOk() (*int64, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *ExecutionSimpleResults) SetEngineId(v int64)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *ExecutionSimpleResults) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetErrorCode

`func (o *ExecutionSimpleResults) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ExecutionSimpleResults) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ExecutionSimpleResults) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ExecutionSimpleResults) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetExecutedSteps

`func (o *ExecutionSimpleResults) GetExecutedSteps() int32`

GetExecutedSteps returns the ExecutedSteps field if non-nil, zero value otherwise.

### GetExecutedStepsOk

`func (o *ExecutionSimpleResults) GetExecutedStepsOk() (*int32, bool)`

GetExecutedStepsOk returns a tuple with the ExecutedSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedSteps

`func (o *ExecutionSimpleResults) SetExecutedSteps(v int32)`

SetExecutedSteps sets ExecutedSteps field to given value.

### HasExecutedSteps

`func (o *ExecutionSimpleResults) HasExecutedSteps() bool`

HasExecutedSteps returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ExecutionSimpleResults) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ExecutionSimpleResults) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ExecutionSimpleResults) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ExecutionSimpleResults) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetHostNameResolutionTime

`func (o *ExecutionSimpleResults) GetHostNameResolutionTime() int64`

GetHostNameResolutionTime returns the HostNameResolutionTime field if non-nil, zero value otherwise.

### GetHostNameResolutionTimeOk

`func (o *ExecutionSimpleResults) GetHostNameResolutionTimeOk() (*int64, bool)`

GetHostNameResolutionTimeOk returns a tuple with the HostNameResolutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNameResolutionTime

`func (o *ExecutionSimpleResults) SetHostNameResolutionTime(v int64)`

SetHostNameResolutionTime sets HostNameResolutionTime field to given value.

### HasHostNameResolutionTime

`func (o *ExecutionSimpleResults) HasHostNameResolutionTime() bool`

HasHostNameResolutionTime returns a boolean if a field has been set.

### GetHttperror

`func (o *ExecutionSimpleResults) GetHttperror() bool`

GetHttperror returns the Httperror field if non-nil, zero value otherwise.

### GetHttperrorOk

`func (o *ExecutionSimpleResults) GetHttperrorOk() (*bool, bool)`

GetHttperrorOk returns a tuple with the Httperror field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttperror

`func (o *ExecutionSimpleResults) SetHttperror(v bool)`

SetHttperror sets Httperror field to given value.

### HasHttperror

`func (o *ExecutionSimpleResults) HasHttperror() bool`

HasHttperror returns a boolean if a field has been set.

### GetPeerCertificateExpiryDate

`func (o *ExecutionSimpleResults) GetPeerCertificateExpiryDate() int64`

GetPeerCertificateExpiryDate returns the PeerCertificateExpiryDate field if non-nil, zero value otherwise.

### GetPeerCertificateExpiryDateOk

`func (o *ExecutionSimpleResults) GetPeerCertificateExpiryDateOk() (*int64, bool)`

GetPeerCertificateExpiryDateOk returns a tuple with the PeerCertificateExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCertificateExpiryDate

`func (o *ExecutionSimpleResults) SetPeerCertificateExpiryDate(v int64)`

SetPeerCertificateExpiryDate sets PeerCertificateExpiryDate field to given value.

### HasPeerCertificateExpiryDate

`func (o *ExecutionSimpleResults) HasPeerCertificateExpiryDate() bool`

HasPeerCertificateExpiryDate returns a boolean if a field has been set.

### GetPublicLocation

`func (o *ExecutionSimpleResults) GetPublicLocation() bool`

GetPublicLocation returns the PublicLocation field if non-nil, zero value otherwise.

### GetPublicLocationOk

`func (o *ExecutionSimpleResults) GetPublicLocationOk() (*bool, bool)`

GetPublicLocationOk returns a tuple with the PublicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicLocation

`func (o *ExecutionSimpleResults) SetPublicLocation(v bool)`

SetPublicLocation sets PublicLocation field to given value.

### HasPublicLocation

`func (o *ExecutionSimpleResults) HasPublicLocation() bool`

HasPublicLocation returns a boolean if a field has been set.

### GetRedirectionTime

`func (o *ExecutionSimpleResults) GetRedirectionTime() int64`

GetRedirectionTime returns the RedirectionTime field if non-nil, zero value otherwise.

### GetRedirectionTimeOk

`func (o *ExecutionSimpleResults) GetRedirectionTimeOk() (*int64, bool)`

GetRedirectionTimeOk returns a tuple with the RedirectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionTime

`func (o *ExecutionSimpleResults) SetRedirectionTime(v int64)`

SetRedirectionTime sets RedirectionTime field to given value.

### HasRedirectionTime

`func (o *ExecutionSimpleResults) HasRedirectionTime() bool`

HasRedirectionTime returns a boolean if a field has been set.

### GetRedirectsCount

`func (o *ExecutionSimpleResults) GetRedirectsCount() int32`

GetRedirectsCount returns the RedirectsCount field if non-nil, zero value otherwise.

### GetRedirectsCountOk

`func (o *ExecutionSimpleResults) GetRedirectsCountOk() (*int32, bool)`

GetRedirectsCountOk returns a tuple with the RedirectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectsCount

`func (o *ExecutionSimpleResults) SetRedirectsCount(v int32)`

SetRedirectsCount sets RedirectsCount field to given value.

### HasRedirectsCount

`func (o *ExecutionSimpleResults) HasRedirectsCount() bool`

HasRedirectsCount returns a boolean if a field has been set.

### GetResponseBodySizeLimitExceeded

`func (o *ExecutionSimpleResults) GetResponseBodySizeLimitExceeded() bool`

GetResponseBodySizeLimitExceeded returns the ResponseBodySizeLimitExceeded field if non-nil, zero value otherwise.

### GetResponseBodySizeLimitExceededOk

`func (o *ExecutionSimpleResults) GetResponseBodySizeLimitExceededOk() (*bool, bool)`

GetResponseBodySizeLimitExceededOk returns a tuple with the ResponseBodySizeLimitExceeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBodySizeLimitExceeded

`func (o *ExecutionSimpleResults) SetResponseBodySizeLimitExceeded(v bool)`

SetResponseBodySizeLimitExceeded sets ResponseBodySizeLimitExceeded field to given value.

### HasResponseBodySizeLimitExceeded

`func (o *ExecutionSimpleResults) HasResponseBodySizeLimitExceeded() bool`

HasResponseBodySizeLimitExceeded returns a boolean if a field has been set.

### GetResponseSize

`func (o *ExecutionSimpleResults) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *ExecutionSimpleResults) GetResponseSizeOk() (*int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSize

`func (o *ExecutionSimpleResults) SetResponseSize(v int64)`

SetResponseSize sets ResponseSize field to given value.

### HasResponseSize

`func (o *ExecutionSimpleResults) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *ExecutionSimpleResults) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *ExecutionSimpleResults) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *ExecutionSimpleResults) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *ExecutionSimpleResults) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *ExecutionSimpleResults) GetStartTimestamp() int64`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *ExecutionSimpleResults) GetStartTimestampOk() (*int64, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *ExecutionSimpleResults) SetStartTimestamp(v int64)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *ExecutionSimpleResults) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionSimpleResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionSimpleResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionSimpleResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionSimpleResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTcpConnectTime

`func (o *ExecutionSimpleResults) GetTcpConnectTime() int64`

GetTcpConnectTime returns the TcpConnectTime field if non-nil, zero value otherwise.

### GetTcpConnectTimeOk

`func (o *ExecutionSimpleResults) GetTcpConnectTimeOk() (*int64, bool)`

GetTcpConnectTimeOk returns a tuple with the TcpConnectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpConnectTime

`func (o *ExecutionSimpleResults) SetTcpConnectTime(v int64)`

SetTcpConnectTime sets TcpConnectTime field to given value.

### HasTcpConnectTime

`func (o *ExecutionSimpleResults) HasTcpConnectTime() bool`

HasTcpConnectTime returns a boolean if a field has been set.

### GetTimeToFirstByte

`func (o *ExecutionSimpleResults) GetTimeToFirstByte() int64`

GetTimeToFirstByte returns the TimeToFirstByte field if non-nil, zero value otherwise.

### GetTimeToFirstByteOk

`func (o *ExecutionSimpleResults) GetTimeToFirstByteOk() (*int64, bool)`

GetTimeToFirstByteOk returns a tuple with the TimeToFirstByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToFirstByte

`func (o *ExecutionSimpleResults) SetTimeToFirstByte(v int64)`

SetTimeToFirstByte sets TimeToFirstByte field to given value.

### HasTimeToFirstByte

`func (o *ExecutionSimpleResults) HasTimeToFirstByte() bool`

HasTimeToFirstByte returns a boolean if a field has been set.

### GetTlsHandshakeTime

`func (o *ExecutionSimpleResults) GetTlsHandshakeTime() int64`

GetTlsHandshakeTime returns the TlsHandshakeTime field if non-nil, zero value otherwise.

### GetTlsHandshakeTimeOk

`func (o *ExecutionSimpleResults) GetTlsHandshakeTimeOk() (*int64, bool)`

GetTlsHandshakeTimeOk returns a tuple with the TlsHandshakeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsHandshakeTime

`func (o *ExecutionSimpleResults) SetTlsHandshakeTime(v int64)`

SetTlsHandshakeTime sets TlsHandshakeTime field to given value.

### HasTlsHandshakeTime

`func (o *ExecutionSimpleResults) HasTlsHandshakeTime() bool`

HasTlsHandshakeTime returns a boolean if a field has been set.

### GetTotalTime

`func (o *ExecutionSimpleResults) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *ExecutionSimpleResults) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *ExecutionSimpleResults) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *ExecutionSimpleResults) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


