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


// CloudFoundryCredentialsConfigurationAPIService CloudFoundryCredentialsConfigurationAPI service
type CloudFoundryCredentialsConfigurationAPIService service

type ApiCreateCloudFoundryCredentialsConfigRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
	cloudFoundryCredentials *CloudFoundryCredentials
}

// &#x60;name&#x60;, &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be unique.
func (r ApiCreateCloudFoundryCredentialsConfigRequest) CloudFoundryCredentials(cloudFoundryCredentials CloudFoundryCredentials) ApiCreateCloudFoundryCredentialsConfigRequest {
	r.cloudFoundryCredentials = &cloudFoundryCredentials
	return r
}

func (r ApiCreateCloudFoundryCredentialsConfigRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.CreateCloudFoundryCredentialsConfigExecute(r)
}

/*
CreateCloudFoundryCredentialsConfig Create new credentials for a single foundation. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateCloudFoundryCredentialsConfigRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) CreateCloudFoundryCredentialsConfig(ctx context.Context) ApiCreateCloudFoundryCredentialsConfigRequest {
	return ApiCreateCloudFoundryCredentialsConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *CloudFoundryCredentialsConfigurationAPIService) CreateCloudFoundryCredentialsConfigExecute(r ApiCreateCloudFoundryCredentialsConfigRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.CreateCloudFoundryCredentialsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudFoundryCredentials == nil {
		return localVarReturnValue, nil, reportError("cloudFoundryCredentials is required and must be specified")
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
	localVarPostBody = r.cloudFoundryCredentials
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

type ApiDeleteCloudFoundryCredentialsConfigRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
	id string
}

func (r ApiDeleteCloudFoundryCredentialsConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCloudFoundryCredentialsConfigExecute(r)
}

/*
DeleteCloudFoundryCredentialsConfig Delete the specified Cloud Foundry foundation credentials. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the Cloud Foundry foundation credentials to be deleted.
 @return ApiDeleteCloudFoundryCredentialsConfigRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) DeleteCloudFoundryCredentialsConfig(ctx context.Context, id string) ApiDeleteCloudFoundryCredentialsConfigRequest {
	return ApiDeleteCloudFoundryCredentialsConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CloudFoundryCredentialsConfigurationAPIService) DeleteCloudFoundryCredentialsConfigExecute(r ApiDeleteCloudFoundryCredentialsConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.DeleteCloudFoundryCredentialsConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiGetCloudFoundryCredentialsConfigRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
	id string
}

func (r ApiGetCloudFoundryCredentialsConfigRequest) Execute() (*CloudFoundryCredentials, *http.Response, error) {
	return r.ApiService.GetCloudFoundryCredentialsConfigExecute(r)
}

/*
GetCloudFoundryCredentialsConfig Show the properties for the specified Cloud Foundry foundation credentials. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the required Cloud Foundry foundation credentials.
 @return ApiGetCloudFoundryCredentialsConfigRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) GetCloudFoundryCredentialsConfig(ctx context.Context, id string) ApiGetCloudFoundryCredentialsConfigRequest {
	return ApiGetCloudFoundryCredentialsConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CloudFoundryCredentials
func (a *CloudFoundryCredentialsConfigurationAPIService) GetCloudFoundryCredentialsConfigExecute(r ApiGetCloudFoundryCredentialsConfigRequest) (*CloudFoundryCredentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudFoundryCredentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.GetCloudFoundryCredentialsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ApiListCloudFoundryCredentialsConfigsRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
}

func (r ApiListCloudFoundryCredentialsConfigsRequest) Execute() (*StubList, *http.Response, error) {
	return r.ApiService.ListCloudFoundryCredentialsConfigsExecute(r)
}

/*
ListCloudFoundryCredentialsConfigs List all the Cloud Foundry foundations credentials. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCloudFoundryCredentialsConfigsRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) ListCloudFoundryCredentialsConfigs(ctx context.Context) ApiListCloudFoundryCredentialsConfigsRequest {
	return ApiListCloudFoundryCredentialsConfigsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StubList
func (a *CloudFoundryCredentialsConfigurationAPIService) ListCloudFoundryCredentialsConfigsExecute(r ApiListCloudFoundryCredentialsConfigsRequest) (*StubList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StubList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.ListCloudFoundryCredentialsConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials"

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

type ApiUpdateCloudFoundryCredentialsConfigRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
	id string
	cloudFoundryCredentials *CloudFoundryCredentials
}

// &#x60;name&#x60; must be unique. &#x60;password&#x60; can be omitted for updates, the existing one will be used. &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be set and may not differ from the existing config if it exists. Use this endpoint for copying credentials between environments while keeping their IDs and for updating existing credentials. You can *not* use this to create new credentials with an arbitrary ID, use POST instead.
func (r ApiUpdateCloudFoundryCredentialsConfigRequest) CloudFoundryCredentials(cloudFoundryCredentials CloudFoundryCredentials) ApiUpdateCloudFoundryCredentialsConfigRequest {
	r.cloudFoundryCredentials = &cloudFoundryCredentials
	return r
}

func (r ApiUpdateCloudFoundryCredentialsConfigRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.UpdateCloudFoundryCredentialsConfigExecute(r)
}

/*
UpdateCloudFoundryCredentialsConfig Create or update credentials for a single Cloud Foundry foundation. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the Cloud Foundry foundation credentials.
 @return ApiUpdateCloudFoundryCredentialsConfigRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) UpdateCloudFoundryCredentialsConfig(ctx context.Context, id string) ApiUpdateCloudFoundryCredentialsConfigRequest {
	return ApiUpdateCloudFoundryCredentialsConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *CloudFoundryCredentialsConfigurationAPIService) UpdateCloudFoundryCredentialsConfigExecute(r ApiUpdateCloudFoundryCredentialsConfigRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.UpdateCloudFoundryCredentialsConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudFoundryCredentials == nil {
		return localVarReturnValue, nil, reportError("cloudFoundryCredentials is required and must be specified")
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
	localVarPostBody = r.cloudFoundryCredentials
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

type ApiValidateCreateCloudFoundryCredentialsConfigRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
	cloudFoundryCredentials *CloudFoundryCredentials
}

// &#x60;name&#x60;, &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be unique.
func (r ApiValidateCreateCloudFoundryCredentialsConfigRequest) CloudFoundryCredentials(cloudFoundryCredentials CloudFoundryCredentials) ApiValidateCreateCloudFoundryCredentialsConfigRequest {
	r.cloudFoundryCredentials = &cloudFoundryCredentials
	return r
}

func (r ApiValidateCreateCloudFoundryCredentialsConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateCreateCloudFoundryCredentialsConfigExecute(r)
}

/*
ValidateCreateCloudFoundryCredentialsConfig Validate that creating credentials would be successful. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateCreateCloudFoundryCredentialsConfigRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) ValidateCreateCloudFoundryCredentialsConfig(ctx context.Context) ApiValidateCreateCloudFoundryCredentialsConfigRequest {
	return ApiValidateCreateCloudFoundryCredentialsConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *CloudFoundryCredentialsConfigurationAPIService) ValidateCreateCloudFoundryCredentialsConfigExecute(r ApiValidateCreateCloudFoundryCredentialsConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.ValidateCreateCloudFoundryCredentialsConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials/validator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudFoundryCredentials == nil {
		return nil, reportError("cloudFoundryCredentials is required and must be specified")
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
	localVarPostBody = r.cloudFoundryCredentials
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

type ApiValidateUpdateCloudFoundryCredentialsConfigRequest struct {
	ctx context.Context
	ApiService *CloudFoundryCredentialsConfigurationAPIService
	id string
	cloudFoundryCredentials *CloudFoundryCredentials
}

// &#x60;name&#x60; must be unique. &#x60;password&#x60; can be omitted for updates. See the constraints for the PUT /cloudFoundry/credentials/{id} request.
func (r ApiValidateUpdateCloudFoundryCredentialsConfigRequest) CloudFoundryCredentials(cloudFoundryCredentials CloudFoundryCredentials) ApiValidateUpdateCloudFoundryCredentialsConfigRequest {
	r.cloudFoundryCredentials = &cloudFoundryCredentials
	return r
}

func (r ApiValidateUpdateCloudFoundryCredentialsConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateUpdateCloudFoundryCredentialsConfigExecute(r)
}

/*
ValidateUpdateCloudFoundryCredentialsConfig Validate that updating or creating credentials would be successful. | maturity=EARLY_ADOPTER

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the Cloud Foundry foundation credentials.
 @return ApiValidateUpdateCloudFoundryCredentialsConfigRequest
*/
func (a *CloudFoundryCredentialsConfigurationAPIService) ValidateUpdateCloudFoundryCredentialsConfig(ctx context.Context, id string) ApiValidateUpdateCloudFoundryCredentialsConfigRequest {
	return ApiValidateUpdateCloudFoundryCredentialsConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CloudFoundryCredentialsConfigurationAPIService) ValidateUpdateCloudFoundryCredentialsConfigExecute(r ApiValidateUpdateCloudFoundryCredentialsConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudFoundryCredentialsConfigurationAPIService.ValidateUpdateCloudFoundryCredentialsConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cloudFoundry/credentials/{id}/validator"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudFoundryCredentials == nil {
		return nil, reportError("cloudFoundryCredentials is required and must be specified")
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
	localVarPostBody = r.cloudFoundryCredentials
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
