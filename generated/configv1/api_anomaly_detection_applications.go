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


// AnomalyDetectionApplicationsAPIService AnomalyDetectionApplicationsAPI service
type AnomalyDetectionApplicationsAPIService service

type ApiGetApplicationAnomalyDetectionConfigRequest struct {
	ctx context.Context
	ApiService *AnomalyDetectionApplicationsAPIService
}

func (r ApiGetApplicationAnomalyDetectionConfigRequest) Execute() (*ApplicationAnomalyDetectionConfig, *http.Response, error) {
	return r.ApiService.GetApplicationAnomalyDetectionConfigExecute(r)
}

/*
GetApplicationAnomalyDetectionConfig Gets the configuration of anomaly detection for applications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetApplicationAnomalyDetectionConfigRequest
*/
func (a *AnomalyDetectionApplicationsAPIService) GetApplicationAnomalyDetectionConfig(ctx context.Context) ApiGetApplicationAnomalyDetectionConfigRequest {
	return ApiGetApplicationAnomalyDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApplicationAnomalyDetectionConfig
func (a *AnomalyDetectionApplicationsAPIService) GetApplicationAnomalyDetectionConfigExecute(r ApiGetApplicationAnomalyDetectionConfigRequest) (*ApplicationAnomalyDetectionConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationAnomalyDetectionConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnomalyDetectionApplicationsAPIService.GetApplicationAnomalyDetectionConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalyDetection/applications"

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

type ApiUpdateApplicationAnomalyDetectionConfigRequest struct {
	ctx context.Context
	ApiService *AnomalyDetectionApplicationsAPIService
	applicationAnomalyDetectionConfig *ApplicationAnomalyDetectionConfig
}

// The JSON body of the request, containing parameters of the application anomaly detection configuration.
func (r ApiUpdateApplicationAnomalyDetectionConfigRequest) ApplicationAnomalyDetectionConfig(applicationAnomalyDetectionConfig ApplicationAnomalyDetectionConfig) ApiUpdateApplicationAnomalyDetectionConfigRequest {
	r.applicationAnomalyDetectionConfig = &applicationAnomalyDetectionConfig
	return r
}

func (r ApiUpdateApplicationAnomalyDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateApplicationAnomalyDetectionConfigExecute(r)
}

/*
UpdateApplicationAnomalyDetectionConfig Updates the configuration of anomaly detection for applications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateApplicationAnomalyDetectionConfigRequest
*/
func (a *AnomalyDetectionApplicationsAPIService) UpdateApplicationAnomalyDetectionConfig(ctx context.Context) ApiUpdateApplicationAnomalyDetectionConfigRequest {
	return ApiUpdateApplicationAnomalyDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnomalyDetectionApplicationsAPIService) UpdateApplicationAnomalyDetectionConfigExecute(r ApiUpdateApplicationAnomalyDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnomalyDetectionApplicationsAPIService.UpdateApplicationAnomalyDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalyDetection/applications"

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
	localVarPostBody = r.applicationAnomalyDetectionConfig
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

type ApiValidateApplicationAnomalyDetectionConfigRequest struct {
	ctx context.Context
	ApiService *AnomalyDetectionApplicationsAPIService
	applicationAnomalyDetectionConfig *ApplicationAnomalyDetectionConfig
}

// The JSON body of the request, containing parameters of the application anomaly detection configuration.
func (r ApiValidateApplicationAnomalyDetectionConfigRequest) ApplicationAnomalyDetectionConfig(applicationAnomalyDetectionConfig ApplicationAnomalyDetectionConfig) ApiValidateApplicationAnomalyDetectionConfigRequest {
	r.applicationAnomalyDetectionConfig = &applicationAnomalyDetectionConfig
	return r
}

func (r ApiValidateApplicationAnomalyDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateApplicationAnomalyDetectionConfigExecute(r)
}

/*
ValidateApplicationAnomalyDetectionConfig Validates the configuration of anomaly detection for applications for the `PUT /anomalyDetection/applications` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateApplicationAnomalyDetectionConfigRequest
*/
func (a *AnomalyDetectionApplicationsAPIService) ValidateApplicationAnomalyDetectionConfig(ctx context.Context) ApiValidateApplicationAnomalyDetectionConfigRequest {
	return ApiValidateApplicationAnomalyDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AnomalyDetectionApplicationsAPIService) ValidateApplicationAnomalyDetectionConfigExecute(r ApiValidateApplicationAnomalyDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnomalyDetectionApplicationsAPIService.ValidateApplicationAnomalyDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/anomalyDetection/applications/validator"

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
	localVarPostBody = r.applicationAnomalyDetectionConfig
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
