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
)


// LogMonitoringProcessGroupsAPIService LogMonitoringProcessGroupsAPI service
type LogMonitoringProcessGroupsAPIService service

type ApiProcessGroupLogJobDeleteRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
	jobId string
}

func (r ApiProcessGroupLogJobDeleteRequest) Execute() (*LogJobDeleteResult, *http.Response, error) {
	return r.ApiService.ProcessGroupLogJobDeleteExecute(r)
}

/*
ProcessGroupLogJobDelete Deletes or cancels the specified log analysis job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The Dynatrace entity ID of the required process group.
 @param jobId The ID of the log analysis job to be deleted.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
 @return ApiProcessGroupLogJobDeleteRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobDelete(ctx context.Context, pgId string, jobId string) ApiProcessGroupLogJobDeleteRequest {
	return ApiProcessGroupLogJobDeleteRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return LogJobDeleteResult
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobDeleteExecute(r ApiProcessGroupLogJobDeleteRequest) (*LogJobDeleteResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogJobDeleteResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogJobDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorEnvelope
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiProcessGroupLogJobRecordsRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
	jobId string
	scrollToken *string
	pageSize *int32
}

// The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned.
func (r ApiProcessGroupLogJobRecordsRequest) ScrollToken(scrollToken string) ApiProcessGroupLogJobRecordsRequest {
	r.scrollToken = &scrollToken
	return r
}

// The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is &#x60;10000&#x60;.
func (r ApiProcessGroupLogJobRecordsRequest) PageSize(pageSize int32) ApiProcessGroupLogJobRecordsRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiProcessGroupLogJobRecordsRequest) Execute() (*LogJobRecordsResult, *http.Response, error) {
	return r.ApiService.ProcessGroupLogJobRecordsExecute(r)
}

/*
ProcessGroupLogJobRecords Gets the content of the analyzed log

Results are available only when the status of the analysis job for this log is `READY`. To check the job status, use the [GET analysis job status](https://dt-url.net/wve3r83) request. 

Long results split into several pages. By default, a page contains 100 results. You can change this value with the **pageSize** query parameter, up to 10,000.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The Dynatrace entity ID of the required process group.
 @param jobId The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
 @return ApiProcessGroupLogJobRecordsRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobRecords(ctx context.Context, pgId string, jobId string) ApiProcessGroupLogJobRecordsRequest {
	return ApiProcessGroupLogJobRecordsRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return LogJobRecordsResult
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobRecordsExecute(r ApiProcessGroupLogJobRecordsRequest) (*LogJobRecordsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogJobRecordsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogJobRecords")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.scrollToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scrollToken", r.scrollToken, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
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

type ApiProcessGroupLogJobRecordsFilteredRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
	jobId string
	scrollToken *string
	pageSize *int32
	filterLogContent *FilterLogContent
}

// The **scrollToken** value from the previous response.    You can use it to get the next page of results. Without it, the first page is always returned.
func (r ApiProcessGroupLogJobRecordsFilteredRequest) ScrollToken(scrollToken string) ApiProcessGroupLogJobRecordsFilteredRequest {
	r.scrollToken = &scrollToken
	return r
}

// The number of records per result page.    If not set, each page contains 100 results.    Maximum allowed value is &#x60;10000&#x60;.
func (r ApiProcessGroupLogJobRecordsFilteredRequest) PageSize(pageSize int32) ApiProcessGroupLogJobRecordsFilteredRequest {
	r.pageSize = &pageSize
	return r
}

// Filter the log content by the specified criteria.   See [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) in Dynatrace Documentation for the syntax definition and examples.
func (r ApiProcessGroupLogJobRecordsFilteredRequest) FilterLogContent(filterLogContent FilterLogContent) ApiProcessGroupLogJobRecordsFilteredRequest {
	r.filterLogContent = &filterLogContent
	return r
}

func (r ApiProcessGroupLogJobRecordsFilteredRequest) Execute() (*LogJobRecordsResult, *http.Response, error) {
	return r.ApiService.ProcessGroupLogJobRecordsFilteredExecute(r)
}

/*
ProcessGroupLogJobRecordsFiltered Gets the content of the analyzed log

Results are available only when the status of the analysis job for this log is `READY`. To check the job status, use the [GET analysis job status](https://dt-url.net/wve3r83) request. 

Long results split into several pages. By default, a page contains 100 results. You can change this value with the **pageSize** query parameter, up to 10,000.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The Dynatrace entity ID of the required process group.
 @param jobId The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
 @return ApiProcessGroupLogJobRecordsFilteredRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobRecordsFiltered(ctx context.Context, pgId string, jobId string) ApiProcessGroupLogJobRecordsFilteredRequest {
	return ApiProcessGroupLogJobRecordsFilteredRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return LogJobRecordsResult
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobRecordsFilteredExecute(r ApiProcessGroupLogJobRecordsFilteredRequest) (*LogJobRecordsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogJobRecordsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogJobRecordsFiltered")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.scrollToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scrollToken", r.scrollToken, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
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
	localVarPostBody = r.filterLogContent
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

type ApiProcessGroupLogJobRecordsTopRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
	jobId string
	filterTopLogRecords *FilterTopLogRecords
}

// Filter the log content by the specified criteria.   See [Search patterns in log data and parse results](https://dt-url.net/57a3rgv) in Dynatrace Documentation for the syntax definition and examples.
func (r ApiProcessGroupLogJobRecordsTopRequest) FilterTopLogRecords(filterTopLogRecords FilterTopLogRecords) ApiProcessGroupLogJobRecordsTopRequest {
	r.filterTopLogRecords = &filterTopLogRecords
	return r
}

func (r ApiProcessGroupLogJobRecordsTopRequest) Execute() (*LogJobRecordsTopValuesRestResult, *http.Response, error) {
	return r.ApiService.ProcessGroupLogJobRecordsTopExecute(r)
}

/*
ProcessGroupLogJobRecordsTop Gets the top values of fields present in the content of the analyzed log

Results are available only when the status of the analysis job for this log is `READY`. To check the job status, use the [GET analysis job status](https://dt-url.net/usg3rbv) request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The Dynatrace entity ID of the required process group.
 @param jobId The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
 @return ApiProcessGroupLogJobRecordsTopRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobRecordsTop(ctx context.Context, pgId string, jobId string) ApiProcessGroupLogJobRecordsTopRequest {
	return ApiProcessGroupLogJobRecordsTopRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return LogJobRecordsTopValuesRestResult
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobRecordsTopExecute(r ApiProcessGroupLogJobRecordsTopRequest) (*LogJobRecordsTopValuesRestResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogJobRecordsTopValuesRestResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogJobRecordsTop")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records/top"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

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
	localVarPostBody = r.filterTopLogRecords
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

type ApiProcessGroupLogJobStartRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
	logPath string
	hostFilter *string
	query *string
	startTimestamp *int64
	endTimestamp *int64
	extractFields *ExtractFields
}

// Narrows down the scope of the analysis to process groups, running at the specified hosts.   Specify the entity ID of the required host here. To specify several IDs, separate them with a comma.
func (r ApiProcessGroupLogJobStartRequest) HostFilter(hostFilter string) ApiProcessGroupLogJobStartRequest {
	r.hostFilter = &hostFilter
	return r
}

// Narrows down the scope of the analysis to the entries, matching the specified criteria.   The criteria must use the [text pattern query syntax](https://dt-url.net/vv83rhp).
func (r ApiProcessGroupLogJobStartRequest) Query(query string) ApiProcessGroupLogJobStartRequest {
	r.query = &query
	return r
}

// The start timestamp of the analysis range, in UTC milliseconds.    If not set, then 2 hours behind from current timestamp is used.
func (r ApiProcessGroupLogJobStartRequest) StartTimestamp(startTimestamp int64) ApiProcessGroupLogJobStartRequest {
	r.startTimestamp = &startTimestamp
	return r
}

// The end timestamp of the analysis range, in UTC milliseconds.    If not set, then the current timestamp is used.
func (r ApiProcessGroupLogJobStartRequest) EndTimestamp(endTimestamp int64) ApiProcessGroupLogJobStartRequest {
	r.endTimestamp = &endTimestamp
	return r
}

// Extract fields from the log content to form custom columns.    See [Search patterns in log data and parse results](https://dt-url.net/vv83rhp) in Dynatrace Documentation for the syntax definition and examples.   The special characters must be escaped.
func (r ApiProcessGroupLogJobStartRequest) ExtractFields(extractFields ExtractFields) ApiProcessGroupLogJobStartRequest {
	r.extractFields = &extractFields
	return r
}

func (r ApiProcessGroupLogJobStartRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ProcessGroupLogJobStartExecute(r)
}

/*
ProcessGroupLogJobStart Starts analysis job for the specified process group log

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The Dynatrace entity ID of the required process group.
 @param logPath The full pathname of the log.
 @return ApiProcessGroupLogJobStartRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobStart(ctx context.Context, pgId string, logPath string) ApiProcessGroupLogJobStartRequest {
	return ApiProcessGroupLogJobStartRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
		logPath: logPath,
	}
}

// Execute executes the request
//  @return string
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobStartExecute(r ApiProcessGroupLogJobStartRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogJobStart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs/{logPath}"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"logPath"+"}", url.PathEscape(parameterValueToString(r.logPath, "logPath")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.hostFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hostFilter", r.hostFilter, "")
	}
	if r.query != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "")
	}
	if r.startTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTimestamp", r.startTimestamp, "")
	}
	if r.endTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTimestamp", r.endTimestamp, "")
	}
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
	localVarPostBody = r.extractFields
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiProcessGroupLogJobStatusRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
	jobId string
}

func (r ApiProcessGroupLogJobStatusRequest) Execute() (*LogJobStatusResult, *http.Response, error) {
	return r.ApiService.ProcessGroupLogJobStatusExecute(r)
}

/*
ProcessGroupLogJobStatus Gets status of the specified log analysis job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The Dynatrace entity ID of the required process group.
 @param jobId The ID of the required log analysis job.    You can retrieve it from the response of the [POST analysis job](https://dt-url.net/c2m3rxl) request.
 @return ApiProcessGroupLogJobStatusRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobStatus(ctx context.Context, pgId string, jobId string) ApiProcessGroupLogJobStatusRequest {
	return ApiProcessGroupLogJobStatusRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
		jobId: jobId,
	}
}

// Execute executes the request
//  @return LogJobStatusResult
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogJobStatusExecute(r ApiProcessGroupLogJobStatusRequest) (*LogJobStatusResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogJobStatusResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogJobStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterValueToString(r.jobId, "jobId")), -1)

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

type ApiProcessGroupLogListRequest struct {
	ctx context.Context
	ApiService *LogMonitoringProcessGroupsAPIService
	pgId string
}

func (r ApiProcessGroupLogListRequest) Execute() (*LogList4pgResult, *http.Response, error) {
	return r.ApiService.ProcessGroupLogListExecute(r)
}

/*
ProcessGroupLogList Lists all the available logs of the specified process group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param pgId The entity ID of the process group.
 @return ApiProcessGroupLogListRequest
*/
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogList(ctx context.Context, pgId string) ApiProcessGroupLogListRequest {
	return ApiProcessGroupLogListRequest{
		ApiService: a,
		ctx: ctx,
		pgId: pgId,
	}
}

// Execute executes the request
//  @return LogList4pgResult
func (a *LogMonitoringProcessGroupsAPIService) ProcessGroupLogListExecute(r ApiProcessGroupLogListRequest) (*LogList4pgResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogList4pgResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogMonitoringProcessGroupsAPIService.ProcessGroupLogList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity/infrastructure/process-groups/{pgId}/logs"
	localVarPath = strings.Replace(localVarPath, "{"+"pgId"+"}", url.PathEscape(parameterValueToString(r.pgId, "pgId")), -1)

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