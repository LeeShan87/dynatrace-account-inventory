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


// TopologySmartscapeApplicationAPIService TopologySmartscapeApplicationAPI service
type TopologySmartscapeApplicationAPIService service

type ApiGetApplicationsRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeApplicationAPIService
	startTimestamp *int64
	endTimestamp *int64
	relativeTime *string
	tag *[]string
	entity *[]string
	managementZone *int64
	includeDetails *bool
	pageSize *int32
	nextPageKey *string
}

// The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used.
func (r ApiGetApplicationsRequest) StartTimestamp(startTimestamp int64) ApiGetApplicationsRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 3 days.
func (r ApiGetApplicationsRequest) EndTimestamp(endTimestamp int64) ApiGetApplicationsRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// The relative timeframe, back from now.
func (r ApiGetApplicationsRequest) RelativeTime(relativeTime string) ApiGetApplicationsRequest {
	r.relativeTime = &relativeTime
	return r
}

// Filters the resulting set of applications by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The application has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;.
func (r ApiGetApplicationsRequest) Tag(tag []string) ApiGetApplicationsRequest {
	r.tag = &tag
	return r
}

// Filters result to the specified applications only.    To specify several applications use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;.
func (r ApiGetApplicationsRequest) Entity(entity []string) ApiGetApplicationsRequest {
	r.entity = &entity
	return r
}

// Only return applications that are part of the specified management zone.
func (r ApiGetApplicationsRequest) ManagementZone(managementZone int64) ApiGetApplicationsRequest {
	r.managementZone = &managementZone
	return r
}

// Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used.
func (r ApiGetApplicationsRequest) IncludeDetails(includeDetails bool) ApiGetApplicationsRequest {
	r.includeDetails = &includeDetails
	return r
}

// The number of applications per result page.    If not set, pagination is not used and the result contains all applications fitting the specified filtering criteria.
func (r ApiGetApplicationsRequest) PageSize(pageSize int32) ApiGetApplicationsRequest {
	r.pageSize = &pageSize
	return r
}

// The cursor for the next page of results. You can find it in the **Next-Page-Key** header of the previous response.   If you&#39;re using pagination, the first page is always returned without this cursor.   You must keep all other query parameters as they were in the first request to obtain subsequent pages.
func (r ApiGetApplicationsRequest) NextPageKey(nextPageKey string) ApiGetApplicationsRequest {
	r.nextPageKey = &nextPageKey
	return r
}

func (r ApiGetApplicationsRequest) Execute() ([]Application, *http.Response, error) {
	return r.ApiService.GetApplicationsExecute(r)
}

/*
GetApplications Gets the list of all applications in your environment along with their parameters

You can narrow down the output by specifying filtering parameters for the request. 

You can additionally limit the output by using pagination: 
1. Specify the number of results per page in the **pageSize** query parameter. 
2. Then use the URL-encoded cursor from the **Next-Page-Key** response header in the **nextPageKey** query parameter to obtain subsequent pages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetApplicationsRequest

Deprecated
*/
func (a *TopologySmartscapeApplicationAPIService) GetApplications(ctx context.Context) ApiGetApplicationsRequest {
	return ApiGetApplicationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Application
// Deprecated
func (a *TopologySmartscapeApplicationAPIService) GetApplicationsExecute(r ApiGetApplicationsRequest) ([]Application, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Application
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeApplicationAPIService.GetApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/applications"

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

type ApiGetBaselineValuesForSingleApplicationRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeApplicationAPIService
	meIdentifier string
}

func (r ApiGetBaselineValuesForSingleApplicationRequest) Execute() (*ApplicationBaselineValues, *http.Response, error) {
	return r.ApiService.GetBaselineValuesForSingleApplicationExecute(r)
}

/*
GetBaselineValuesForSingleApplication Gets baseline data for the specified application | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meIdentifier The Dynatrace entity ID of the required application.
 @return ApiGetBaselineValuesForSingleApplicationRequest

Deprecated
*/
func (a *TopologySmartscapeApplicationAPIService) GetBaselineValuesForSingleApplication(ctx context.Context, meIdentifier string) ApiGetBaselineValuesForSingleApplicationRequest {
	return ApiGetBaselineValuesForSingleApplicationRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

// Execute executes the request
//  @return ApplicationBaselineValues
// Deprecated
func (a *TopologySmartscapeApplicationAPIService) GetBaselineValuesForSingleApplicationExecute(r ApiGetBaselineValuesForSingleApplicationRequest) (*ApplicationBaselineValues, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationBaselineValues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeApplicationAPIService.GetBaselineValuesForSingleApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/applications/{meIdentifier}/baseline"
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

type ApiGetSingleApplicationRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeApplicationAPIService
	meIdentifier string
}

func (r ApiGetSingleApplicationRequest) Execute() (*Application, *http.Response, error) {
	return r.ApiService.GetSingleApplicationExecute(r)
}

/*
GetSingleApplication Gets parameters of the specified application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meIdentifier The Dynatrace entity ID of the required application.
 @return ApiGetSingleApplicationRequest

Deprecated
*/
func (a *TopologySmartscapeApplicationAPIService) GetSingleApplication(ctx context.Context, meIdentifier string) ApiGetSingleApplicationRequest {
	return ApiGetSingleApplicationRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

// Execute executes the request
//  @return Application
// Deprecated
func (a *TopologySmartscapeApplicationAPIService) GetSingleApplicationExecute(r ApiGetSingleApplicationRequest) (*Application, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Application
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeApplicationAPIService.GetSingleApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/applications/{meIdentifier}"
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

type ApiUpdateApplicationRequest struct {
	ctx context.Context
	ApiService *TopologySmartscapeApplicationAPIService
	meIdentifier string
	updateEntity *UpdateEntity
}

func (r ApiUpdateApplicationRequest) UpdateEntity(updateEntity UpdateEntity) ApiUpdateApplicationRequest {
	r.updateEntity = &updateEntity
	return r
}

func (r ApiUpdateApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateApplicationExecute(r)
}

/*
UpdateApplication Updates parameters of the specified application

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param meIdentifier The Dynatrace entity ID of the application you want to update.
 @return ApiUpdateApplicationRequest

Deprecated
*/
func (a *TopologySmartscapeApplicationAPIService) UpdateApplication(ctx context.Context, meIdentifier string) ApiUpdateApplicationRequest {
	return ApiUpdateApplicationRequest{
		ApiService: a,
		ctx: ctx,
		meIdentifier: meIdentifier,
	}
}

// Execute executes the request
// Deprecated
func (a *TopologySmartscapeApplicationAPIService) UpdateApplicationExecute(r ApiUpdateApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TopologySmartscapeApplicationAPIService.UpdateApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/applications/{meIdentifier}"
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
