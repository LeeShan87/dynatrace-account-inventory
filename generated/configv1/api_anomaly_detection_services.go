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
)


// AnomalyDetectionServicesAPIService AnomalyDetectionServicesAPI service
type AnomalyDetectionServicesAPIService service

type ApiGetServiceAnomalyDetectionConfigRequest struct {
	ctx context.Context
	ApiService *AnomalyDetectionServicesAPIService
}

func (r ApiGetServiceAnomalyDetectionConfigRequest) Execute() (*ServiceAnomalyDetectionConfig, *http.Response, error) {
	return r.ApiService.GetServiceAnomalyDetectionConfigExecute(r)
}

/*
GetServiceAnomalyDetectionConfig Gets the service anomaly detection configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetServiceAnomalyDetectionConfigRequest
*/
func (a *AnomalyDetectionServicesAPIService) GetServiceAnomalyDetectionConfig(ctx context.Context) ApiGetServiceAnomalyDetectionConfigRequest {
	return ApiGetServiceAnomalyDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ServiceAnomalyDetectionConfig
func (a *AnomalyDetectionServicesAPIService) GetServiceAnomalyDetectionConfigExecute(r ApiGetServiceAnomalyDetectionConfigRequest) (*ServiceAnomalyDetectionConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceAnomalyDetectionConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnomalyDetectionServicesAPIService.GetServiceAnomalyDetectionConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalyDetection/services"

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

type ApiUpdateServiceAnomalyDetectionConfigRequest struct {
	ctx context.Context
	ApiService *AnomalyDetectionServicesAPIService
	serviceAnomalyDetectionConfig *ServiceAnomalyDetectionConfig
}

// The JSON body of the request. Contains parameters of the service anomaly detection configuration.
func (r ApiUpdateServiceAnomalyDetectionConfigRequest) ServiceAnomalyDetectionConfig(serviceAnomalyDetectionConfig ServiceAnomalyDetectionConfig) ApiUpdateServiceAnomalyDetectionConfigRequest {
	r.serviceAnomalyDetectionConfig = &serviceAnomalyDetectionConfig
	return r
}

func (r ApiUpdateServiceAnomalyDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateServiceAnomalyDetectionConfigExecute(r)
}

/*
UpdateServiceAnomalyDetectionConfig Updates the service anomaly detection configuration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateServiceAnomalyDetectionConfigRequest
*/
func (a *AnomalyDetectionServicesAPIService) UpdateServiceAnomalyDetectionConfig(ctx context.Context) ApiUpdateServiceAnomalyDetectionConfigRequest {
	return ApiUpdateServiceAnomalyDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnomalyDetectionServicesAPIService) UpdateServiceAnomalyDetectionConfigExecute(r ApiUpdateServiceAnomalyDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnomalyDetectionServicesAPIService.UpdateServiceAnomalyDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalyDetection/services"

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
	localVarPostBody = r.serviceAnomalyDetectionConfig
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

type ApiValidateServiceAnomalyDetectionConfigRequest struct {
	ctx context.Context
	ApiService *AnomalyDetectionServicesAPIService
	serviceAnomalyDetectionConfig *ServiceAnomalyDetectionConfig
}

// The JSON body of the request. Contains parameters of the service anomaly detection configuration.
func (r ApiValidateServiceAnomalyDetectionConfigRequest) ServiceAnomalyDetectionConfig(serviceAnomalyDetectionConfig ServiceAnomalyDetectionConfig) ApiValidateServiceAnomalyDetectionConfigRequest {
	r.serviceAnomalyDetectionConfig = &serviceAnomalyDetectionConfig
	return r
}

func (r ApiValidateServiceAnomalyDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateServiceAnomalyDetectionConfigExecute(r)
}

/*
ValidateServiceAnomalyDetectionConfig Validates the payload for the `PUT /anomalyDetection/services` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateServiceAnomalyDetectionConfigRequest
*/
func (a *AnomalyDetectionServicesAPIService) ValidateServiceAnomalyDetectionConfig(ctx context.Context) ApiValidateServiceAnomalyDetectionConfigRequest {
	return ApiValidateServiceAnomalyDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnomalyDetectionServicesAPIService) ValidateServiceAnomalyDetectionConfigExecute(r ApiValidateServiceAnomalyDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnomalyDetectionServicesAPIService.ValidateServiceAnomalyDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalyDetection/services/validator"

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
	localVarPostBody = r.serviceAnomalyDetectionConfig
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
