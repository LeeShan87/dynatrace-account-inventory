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


// RUMApplicationDetectionRulesAPIService RUMApplicationDetectionRulesAPI service
type RUMApplicationDetectionRulesAPIService service

type ApiCreateApplicationDetectionConfigRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	position *string
	applicationDetectionRuleConfig *ApplicationDetectionRuleConfig
}

// The position of the new rule:    * &#x60;APPEND&#x60;: at the bottom of the rule list.   * &#x60;PREPEND&#x60;: at the top of the rule list.    If not set, the &#x60;APPEND&#x60; is used.
func (r ApiCreateApplicationDetectionConfigRequest) Position(position string) ApiCreateApplicationDetectionConfigRequest {
	r.position = &position
	return r
}

// The JSON body of the request. Contains configuration of the new application detection rule.    You must not specify the ID of the rule.   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /applicationDetectionRules/order&#x60; request.
func (r ApiCreateApplicationDetectionConfigRequest) ApplicationDetectionRuleConfig(applicationDetectionRuleConfig ApplicationDetectionRuleConfig) ApiCreateApplicationDetectionConfigRequest {
	r.applicationDetectionRuleConfig = &applicationDetectionRuleConfig
	return r
}

func (r ApiCreateApplicationDetectionConfigRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.CreateApplicationDetectionConfigExecute(r)
}

/*
CreateApplicationDetectionConfig Creates a new application detection rule

The body must not provide an ID. An ID is assigned automatically by the Dynatrace server. 

You can only create detection rules for an existing application. If you need to create a rule for an application that doesn't exist yet, [create an application first](https://dt-url.net/vt03khh) and then configure detection rules for it.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateApplicationDetectionConfigRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) CreateApplicationDetectionConfig(ctx context.Context) ApiCreateApplicationDetectionConfigRequest {
	return ApiCreateApplicationDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *RUMApplicationDetectionRulesAPIService) CreateApplicationDetectionConfigExecute(r ApiCreateApplicationDetectionConfigRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.CreateApplicationDetectionConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.position != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "position", r.position, "")
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
	localVarPostBody = r.applicationDetectionRuleConfig
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

type ApiDeleteApplicationDetectionConfigRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	id string
}

func (r ApiDeleteApplicationDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationDetectionConfigExecute(r)
}

/*
DeleteApplicationDetectionConfig Deletes the specified application detection rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the application detection rule to be deleted.
 @return ApiDeleteApplicationDetectionConfigRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) DeleteApplicationDetectionConfig(ctx context.Context, id string) ApiDeleteApplicationDetectionConfigRequest {
	return ApiDeleteApplicationDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *RUMApplicationDetectionRulesAPIService) DeleteApplicationDetectionConfigExecute(r ApiDeleteApplicationDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.DeleteApplicationDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules/{id}"
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

type ApiGetApplicationDetectionConfigRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	id string
}

func (r ApiGetApplicationDetectionConfigRequest) Execute() (*ApplicationDetectionRuleConfig, *http.Response, error) {
	return r.ApiService.GetApplicationDetectionConfigExecute(r)
}

/*
GetApplicationDetectionConfig Gets the parameters of the specified application detection rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the required application detection rule.
 @return ApiGetApplicationDetectionConfigRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) GetApplicationDetectionConfig(ctx context.Context, id string) ApiGetApplicationDetectionConfigRequest {
	return ApiGetApplicationDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ApplicationDetectionRuleConfig
func (a *RUMApplicationDetectionRulesAPIService) GetApplicationDetectionConfigExecute(r ApiGetApplicationDetectionConfigRequest) (*ApplicationDetectionRuleConfig, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicationDetectionRuleConfig
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.GetApplicationDetectionConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules/{id}"
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

type ApiListApplicationDetectionConfigsRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
}

func (r ApiListApplicationDetectionConfigsRequest) Execute() (*StubList, *http.Response, error) {
	return r.ApiService.ListApplicationDetectionConfigsExecute(r)
}

/*
ListApplicationDetectionConfigs Lists all available application detection rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListApplicationDetectionConfigsRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) ListApplicationDetectionConfigs(ctx context.Context) ApiListApplicationDetectionConfigsRequest {
	return ApiListApplicationDetectionConfigsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StubList
func (a *RUMApplicationDetectionRulesAPIService) ListApplicationDetectionConfigsExecute(r ApiListApplicationDetectionConfigsRequest) (*StubList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StubList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.ListApplicationDetectionConfigs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules"

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

type ApiOrderApplicationDetectionConfigsRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	stubList *StubList
}

// The JSON body of the request. Contains the IDs of the application detection rules in the desired order. Any further properties (**name**, **description**) are ignored.
func (r ApiOrderApplicationDetectionConfigsRequest) StubList(stubList StubList) ApiOrderApplicationDetectionConfigsRequest {
	r.stubList = &stubList
	return r
}

func (r ApiOrderApplicationDetectionConfigsRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrderApplicationDetectionConfigsExecute(r)
}

/*
OrderApplicationDetectionConfigs Reorders the application detection rules

This request reorders the application detection rules according to the submitted list of IDs. Application detection rules not present in the body of the request will retain their relative ordering but are placed *after* all those present in the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrderApplicationDetectionConfigsRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) OrderApplicationDetectionConfigs(ctx context.Context) ApiOrderApplicationDetectionConfigsRequest {
	return ApiOrderApplicationDetectionConfigsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RUMApplicationDetectionRulesAPIService) OrderApplicationDetectionConfigsExecute(r ApiOrderApplicationDetectionConfigsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.OrderApplicationDetectionConfigs")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules/order"

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
	localVarPostBody = r.stubList
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

type ApiUpdateApplicationDetectionConfigRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	id string
	applicationDetectionRuleConfig *ApplicationDetectionRuleConfig
}

// The JSON body of the request. Contains updated parameters of the application detection rule.    If the **order** parameter is set, the rule is placed to this position.
func (r ApiUpdateApplicationDetectionConfigRequest) ApplicationDetectionRuleConfig(applicationDetectionRuleConfig ApplicationDetectionRuleConfig) ApiUpdateApplicationDetectionConfigRequest {
	r.applicationDetectionRuleConfig = &applicationDetectionRuleConfig
	return r
}

func (r ApiUpdateApplicationDetectionConfigRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.UpdateApplicationDetectionConfigExecute(r)
}

/*
UpdateApplicationDetectionConfig Updates the specified application detection rule

If the application detection rule with the specified ID doesn't exist, a new application is created and appended to the end of the rule list. 

If the **order** parameter is set for an existing rule, the request uses this value. Otherwise it keeps the existing order of rules.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the application detection rule to be updated.    If you set the ID in the body as well, it must match this ID.
 @return ApiUpdateApplicationDetectionConfigRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) UpdateApplicationDetectionConfig(ctx context.Context, id string) ApiUpdateApplicationDetectionConfigRequest {
	return ApiUpdateApplicationDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *RUMApplicationDetectionRulesAPIService) UpdateApplicationDetectionConfigExecute(r ApiUpdateApplicationDetectionConfigRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.UpdateApplicationDetectionConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarPostBody = r.applicationDetectionRuleConfig
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

type ApiValidateCreateApplicationDetectionConfigRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	applicationDetectionRuleConfig *ApplicationDetectionRuleConfig
}

// The JSON body of the request. Contains the application detection rule to be validated.
func (r ApiValidateCreateApplicationDetectionConfigRequest) ApplicationDetectionRuleConfig(applicationDetectionRuleConfig ApplicationDetectionRuleConfig) ApiValidateCreateApplicationDetectionConfigRequest {
	r.applicationDetectionRuleConfig = &applicationDetectionRuleConfig
	return r
}

func (r ApiValidateCreateApplicationDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateCreateApplicationDetectionConfigExecute(r)
}

/*
ValidateCreateApplicationDetectionConfig Validates the payload for the `POST /applicationDetection` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateCreateApplicationDetectionConfigRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) ValidateCreateApplicationDetectionConfig(ctx context.Context) ApiValidateCreateApplicationDetectionConfigRequest {
	return ApiValidateCreateApplicationDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RUMApplicationDetectionRulesAPIService) ValidateCreateApplicationDetectionConfigExecute(r ApiValidateCreateApplicationDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.ValidateCreateApplicationDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules/validator"

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
	localVarPostBody = r.applicationDetectionRuleConfig
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

type ApiValidateUpdateApplicationDetectionConfigRequest struct {
	ctx context.Context
	ApiService *RUMApplicationDetectionRulesAPIService
	id string
	applicationDetectionRuleConfig *ApplicationDetectionRuleConfig
}

// The JSON body of the request. Contains the application detection rule to be validated.
func (r ApiValidateUpdateApplicationDetectionConfigRequest) ApplicationDetectionRuleConfig(applicationDetectionRuleConfig ApplicationDetectionRuleConfig) ApiValidateUpdateApplicationDetectionConfigRequest {
	r.applicationDetectionRuleConfig = &applicationDetectionRuleConfig
	return r
}

func (r ApiValidateUpdateApplicationDetectionConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateUpdateApplicationDetectionConfigExecute(r)
}

/*
ValidateUpdateApplicationDetectionConfig Validate the payload for the `PUT /applicationDetection/{id}` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the application detection rule to be validated.    If you set the ID in the body as well, it must match this ID.
 @return ApiValidateUpdateApplicationDetectionConfigRequest
*/
func (a *RUMApplicationDetectionRulesAPIService) ValidateUpdateApplicationDetectionConfig(ctx context.Context, id string) ApiValidateUpdateApplicationDetectionConfigRequest {
	return ApiValidateUpdateApplicationDetectionConfigRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *RUMApplicationDetectionRulesAPIService) ValidateUpdateApplicationDetectionConfigExecute(r ApiValidateUpdateApplicationDetectionConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RUMApplicationDetectionRulesAPIService.ValidateUpdateApplicationDetectionConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applicationDetectionRules/{id}/validator"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarPostBody = r.applicationDetectionRuleConfig
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
