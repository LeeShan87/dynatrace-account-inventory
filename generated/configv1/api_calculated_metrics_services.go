/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CalculatedMetricsServicesAPIService CalculatedMetricsServicesAPI service
type CalculatedMetricsServicesAPIService service

type ApiCreateServiceMetricRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
	calculatedServiceMetric *CalculatedServiceMetric
}

// The JSON body of the request. Contains parameters of the new calculated service metric.
func (r ApiCreateServiceMetricRequest) CalculatedServiceMetric(calculatedServiceMetric CalculatedServiceMetric) ApiCreateServiceMetricRequest {
	r.calculatedServiceMetric = &calculatedServiceMetric
	return r
}

func (r ApiCreateServiceMetricRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.CreateServiceMetricExecute(r)
}

/*
CreateServiceMetric Creates a new calculated service metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateServiceMetricRequest
*/
func (a *CalculatedMetricsServicesAPIService) CreateServiceMetric(ctx context.Context) ApiCreateServiceMetricRequest {
	return ApiCreateServiceMetricRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *CalculatedMetricsServicesAPIService) CreateServiceMetricExecute(r ApiCreateServiceMetricRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.CreateServiceMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.calculatedServiceMetric == nil {
		return localVarReturnValue, nil, reportError("calculatedServiceMetric is required and must be specified")
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
	localVarPostBody = r.calculatedServiceMetric
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

type ApiDeleteServiceMetricRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
	metricKey string
}

func (r ApiDeleteServiceMetricRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteServiceMetricExecute(r)
}

/*
DeleteServiceMetric Deletes the specified calculated service metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metricKey The key of the calculated service metric to be deleted.
 @return ApiDeleteServiceMetricRequest
*/
func (a *CalculatedMetricsServicesAPIService) DeleteServiceMetric(ctx context.Context, metricKey string) ApiDeleteServiceMetricRequest {
	return ApiDeleteServiceMetricRequest{
		ApiService: a,
		ctx: ctx,
		metricKey: metricKey,
	}
}

// Execute executes the request
func (a *CalculatedMetricsServicesAPIService) DeleteServiceMetricExecute(r ApiDeleteServiceMetricRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.DeleteServiceMetric")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service/{metricKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"metricKey"+"}", url.PathEscape(parameterValueToString(r.metricKey, "metricKey")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetServiceMetricRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
	metricKey string
}

func (r ApiGetServiceMetricRequest) Execute() (*CalculatedServiceMetric, *http.Response, error) {
	return r.ApiService.GetServiceMetricExecute(r)
}

/*
GetServiceMetric Gets the descriptor of the specified calculated service metric

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metricKey The key of the required calculated service metric.
 @return ApiGetServiceMetricRequest
*/
func (a *CalculatedMetricsServicesAPIService) GetServiceMetric(ctx context.Context, metricKey string) ApiGetServiceMetricRequest {
	return ApiGetServiceMetricRequest{
		ApiService: a,
		ctx: ctx,
		metricKey: metricKey,
	}
}

// Execute executes the request
//  @return CalculatedServiceMetric
func (a *CalculatedMetricsServicesAPIService) GetServiceMetricExecute(r ApiGetServiceMetricRequest) (*CalculatedServiceMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CalculatedServiceMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.GetServiceMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service/{metricKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"metricKey"+"}", url.PathEscape(parameterValueToString(r.metricKey, "metricKey")), -1)

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

type ApiListServiceMetricsRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
}

func (r ApiListServiceMetricsRequest) Execute() (*StubList, *http.Response, error) {
	return r.ApiService.ListServiceMetricsExecute(r)
}

/*
ListServiceMetrics Lists all calculated service metrics configured in your environment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListServiceMetricsRequest
*/
func (a *CalculatedMetricsServicesAPIService) ListServiceMetrics(ctx context.Context) ApiListServiceMetricsRequest {
	return ApiListServiceMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StubList
func (a *CalculatedMetricsServicesAPIService) ListServiceMetricsExecute(r ApiListServiceMetricsRequest) (*StubList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StubList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.ListServiceMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service"

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

type ApiUpdateServiceMetricRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
	metricKey string
	calculatedServiceMetric *CalculatedServiceMetric
}

// The JSON body of the request. Contains updated parameters of the calculated service metric.
func (r ApiUpdateServiceMetricRequest) CalculatedServiceMetric(calculatedServiceMetric CalculatedServiceMetric) ApiUpdateServiceMetricRequest {
	r.calculatedServiceMetric = &calculatedServiceMetric
	return r
}

func (r ApiUpdateServiceMetricRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.UpdateServiceMetricExecute(r)
}

/*
UpdateServiceMetric Updates the specified calculated service metric

If the calculated service metric with the specified key doesn't exist, a new metric is created.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metricKey The key of the calculated service metric to be updated.    The key of the calculated service metric in the body of the request must match this key.
 @return ApiUpdateServiceMetricRequest
*/
func (a *CalculatedMetricsServicesAPIService) UpdateServiceMetric(ctx context.Context, metricKey string) ApiUpdateServiceMetricRequest {
	return ApiUpdateServiceMetricRequest{
		ApiService: a,
		ctx: ctx,
		metricKey: metricKey,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *CalculatedMetricsServicesAPIService) UpdateServiceMetricExecute(r ApiUpdateServiceMetricRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.UpdateServiceMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service/{metricKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"metricKey"+"}", url.PathEscape(parameterValueToString(r.metricKey, "metricKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.calculatedServiceMetric == nil {
		return localVarReturnValue, nil, reportError("calculatedServiceMetric is required and must be specified")
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
	localVarPostBody = r.calculatedServiceMetric
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

type ApiValidateCreateServiceMetricRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
	calculatedServiceMetric *CalculatedServiceMetric
}

// The JSON body of the request. Contains the calculated service metric to be validated.
func (r ApiValidateCreateServiceMetricRequest) CalculatedServiceMetric(calculatedServiceMetric CalculatedServiceMetric) ApiValidateCreateServiceMetricRequest {
	r.calculatedServiceMetric = &calculatedServiceMetric
	return r
}

func (r ApiValidateCreateServiceMetricRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateCreateServiceMetricExecute(r)
}

/*
ValidateCreateServiceMetric Validates the payload for the `POST /calculatedMetric/service` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateCreateServiceMetricRequest
*/
func (a *CalculatedMetricsServicesAPIService) ValidateCreateServiceMetric(ctx context.Context) ApiValidateCreateServiceMetricRequest {
	return ApiValidateCreateServiceMetricRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CalculatedMetricsServicesAPIService) ValidateCreateServiceMetricExecute(r ApiValidateCreateServiceMetricRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.ValidateCreateServiceMetric")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service/validator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.calculatedServiceMetric == nil {
		return nil, reportError("calculatedServiceMetric is required and must be specified")
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
	localVarPostBody = r.calculatedServiceMetric
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

type ApiValidateUpdateServiceMetricRequest struct {
	ctx context.Context
	ApiService *CalculatedMetricsServicesAPIService
	metricKey string
	calculatedServiceMetric *CalculatedServiceMetric
}

// The JSON body of the request. Contains the calculated service metric to be validated.
func (r ApiValidateUpdateServiceMetricRequest) CalculatedServiceMetric(calculatedServiceMetric CalculatedServiceMetric) ApiValidateUpdateServiceMetricRequest {
	r.calculatedServiceMetric = &calculatedServiceMetric
	return r
}

func (r ApiValidateUpdateServiceMetricRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateUpdateServiceMetricExecute(r)
}

/*
ValidateUpdateServiceMetric Validates the payload for the `PUT /calculatedMetric/service/{id}` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param metricKey The key of the calculated service metric to be validated .   The key of the metric in the body of the request must match this key.
 @return ApiValidateUpdateServiceMetricRequest
*/
func (a *CalculatedMetricsServicesAPIService) ValidateUpdateServiceMetric(ctx context.Context, metricKey string) ApiValidateUpdateServiceMetricRequest {
	return ApiValidateUpdateServiceMetricRequest{
		ApiService: a,
		ctx: ctx,
		metricKey: metricKey,
	}
}

// Execute executes the request
func (a *CalculatedMetricsServicesAPIService) ValidateUpdateServiceMetricExecute(r ApiValidateUpdateServiceMetricRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CalculatedMetricsServicesAPIService.ValidateUpdateServiceMetric")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/calculatedMetrics/service/{metricKey}/validator"
	localVarPath = strings.Replace(localVarPath, "{"+"metricKey"+"}", url.PathEscape(parameterValueToString(r.metricKey, "metricKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.calculatedServiceMetric == nil {
		return nil, reportError("calculatedServiceMetric is required and must be specified")
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
	localVarPostBody = r.calculatedServiceMetric
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
