/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// TopologySmartscapeHostAPIService TopologySmartscapeHostAPI service
type TopologySmartscapeHostAPIService service

type ApiGetHostsRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeHostAPIService
	startTimestamp *int64
	endTimestamp *int64
	relativeTime *string
	tag *[]string
	showMonitoringCandidates *bool
	entity *[]string
	managementZone *int64
	hostGroupId *string
	hostGroupName *string
	includeDetails *bool
	pageSize *int32
	nextPageKey *string
}

// The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used.
func (r ApiGetHostsRequest) StartTimestamp(startTimestamp int64) ApiGetHostsRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days.
func (r ApiGetHostsRequest) EndTimestamp(endTimestamp int64) ApiGetHostsRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// The relative timeframe, back from now.
func (r ApiGetHostsRequest) RelativeTime(relativeTime string) ApiGetHostsRequest {
	r.relativeTime = &relativeTime
	return r
}

// Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;.
func (r ApiGetHostsRequest) Tag(tag []string) ApiGetHostsRequest {
	r.tag = &tag
	return r
}

// Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) a monitoring candidate in the response.   Monitoring candidates are network entities that are detected but not monitored.
func (r ApiGetHostsRequest) ShowMonitoringCandidates(showMonitoringCandidates bool) ApiGetHostsRequest {
	r.showMonitoringCandidates = &showMonitoringCandidates
	return r
}

// Filters result to the specified hosts only.    To specify several hosts use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;.
func (r ApiGetHostsRequest) Entity(entity []string) ApiGetHostsRequest {
	r.entity = &entity
	return r
}

// Only return hosts that are part of the specified management zone.
func (r ApiGetHostsRequest) ManagementZone(managementZone int64) ApiGetHostsRequest {
	r.managementZone = &managementZone
	return r
}

// Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace IDs of the host group you&#39;re interested in.
func (r ApiGetHostsRequest) HostGroupId(hostGroupId string) ApiGetHostsRequest {
	r.hostGroupId = &hostGroupId
	return r
}

// Filters the resulting set of hosts by the specified host group.    Specify the name of the host group you&#39;re interested in.
func (r ApiGetHostsRequest) HostGroupName(hostGroupName string) ApiGetHostsRequest {
	r.hostGroupName = &hostGroupName
	return r
}

// Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used.
func (r ApiGetHostsRequest) IncludeDetails(includeDetails bool) ApiGetHostsRequest {
	r.includeDetails = &includeDetails
	return r
}

// The number of hosts per result page.    If not set, pagination is not used and the result contains all hosts fitting the specified filtering criteria.
func (r ApiGetHostsRequest) PageSize(pageSize int32) ApiGetHostsRequest {
	r.pageSize = &pageSize
	return r
}

// The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages.
func (r ApiGetHostsRequest) NextPageKey(nextPageKey string) ApiGetHostsRequest {
	r.nextPageKey = &nextPageKey
	return r
}

func (r ApiGetHostsRequest) Execute() ([]Host, *http.Response, error) {
	return r.ApiService.GetHostsExecute(r)
}

/*
GetHosts Lists all available hosts in your environment

You can narrow down the output by specifying filtering parameters for the request. 

You can additionally limit the output by using pagination: 
1. Specify the number of results per page in the **pageSize** query parameter. 
2. Then use the URL-encoded cursor from the **Next-Page-Key** response header in the **nextPageKey** query parameter to obtain subsequent pages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetHostsRequest

Deprecated
*/
func (a *TopologySmartscapeHostAPIService) GetHosts(ctx context.Context) ApiGetHostsRequest {
	return ApiGetHostsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Host
// Deprecated
func (a *TopologySmartscapeHostAPIService) GetHostsExecute(r ApiGetHostsRequest) ([]Host, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Host
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeHostAPIService.GetHosts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/hosts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "")
	}
	if r.endTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "")
	}
	if r.relativeTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relativeTime", r.relativeTime, "")
	}
	if r.tag != nil {
		t := *r.tag
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "tag", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "tag", t, "multi")
		}
	}
	if r.showMonitoringCandidates != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "showMonitoringCandidates", r.showMonitoringCandidates, "")
	}
	if r.entity != nil {
		t := *r.entity
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "entity", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "entity", t, "multi")
		}
	}
	if r.managementZone != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "managementZone", r.managementZone, "")
	}
	if r.hostGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostGroupId", r.hostGroupId, "")
	}
	if r.hostGroupName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostGroupName", r.hostGroupName, "")
	}
	if r.includeDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDetails", r.includeDetails, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.nextPageKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageKey", r.nextPageKey, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetSingleHostRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeHostAPIService
	meIdentifier string
}

func (r ApiGetSingleHostRequest) Execute() (*Host, *http.Response, error) {
	return r.ApiService.GetSingleHostExecute(r)
}

/*
GetSingleHost Gets parameters of the specified host

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meIdentifier The Dynatrace entity ID of the required host.
 @return ApiGetSingleHostRequest

Deprecated
*/
func (a *TopologySmartscapeHostAPIService) GetSingleHost(ctx context.Context, meIdentifier string) ApiGetSingleHostRequest {
	return ApiGetSingleHostRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

// Execute executes the request
//  @return Host
// Deprecated
func (a *TopologySmartscapeHostAPIService) GetSingleHostExecute(r ApiGetSingleHostRequest) (*Host, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Host
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeHostAPIService.GetSingleHost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/hosts/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", url.PathEscape(parameterValueToString(r.meIdentifier, "meIdentifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveTagsRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeHostAPIService
	meIdentifier string
	tag string
}

func (r ApiRemoveTagsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveTagsExecute(r)
}

/*
RemoveTags Remove tag of the specified host

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meIdentifier The Dynatrace entity ID of the host.
 @param tag The tag to be removed.
 @return ApiRemoveTagsRequest

Deprecated
*/
func (a *TopologySmartscapeHostAPIService) RemoveTags(ctx context.Context, meIdentifier string, tag string) ApiRemoveTagsRequest {
	return ApiRemoveTagsRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
		tag: tag,
	}
}

// Execute executes the request
// Deprecated
func (a *TopologySmartscapeHostAPIService) RemoveTagsExecute(r ApiRemoveTagsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeHostAPIService.RemoveTags")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/hosts/{meIdentifier}/tags/{tag}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", url.PathEscape(parameterValueToString(r.meIdentifier, "meIdentifier")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"tag"+"}", url.PathEscape(parameterValueToString(r.tag, "tag")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateHostRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeHostAPIService
	meIdentifier string
	updateEntity *UpdateEntity
}

func (r ApiUpdateHostRequest) UpdateEntity(updateEntity UpdateEntity) ApiUpdateHostRequest {
	r.updateEntity = &updateEntity
	return r
}

func (r ApiUpdateHostRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateHostExecute(r)
}

/*
UpdateHost Updates properties of the specified host

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meIdentifier The Dynatrace entity ID of the host to be updated.
 @return ApiUpdateHostRequest

Deprecated
*/
func (a *TopologySmartscapeHostAPIService) UpdateHost(ctx context.Context, meIdentifier string) ApiUpdateHostRequest {
	return ApiUpdateHostRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

// Execute executes the request
// Deprecated
func (a *TopologySmartscapeHostAPIService) UpdateHostExecute(r ApiUpdateHostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeHostAPIService.UpdateHost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/hosts/{meIdentifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"meIdentifier"+"}", url.PathEscape(parameterValueToString(r.meIdentifier, "meIdentifier")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json; charset=utf-8"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json; charset=utf-8"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateEntity
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
