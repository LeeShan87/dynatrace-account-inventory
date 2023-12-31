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


// ServiceDetectionOpaqueAndExternalWebRequestAPIService ServiceDetectionOpaqueAndExternalWebRequestAPI service
type ServiceDetectionOpaqueAndExternalWebRequestAPIService service

type ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	position *string
	opaqueAndExternalWebRequestRule *OpaqueAndExternalWebRequestRule
}

// The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.   
func (r ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest) Position(position string) ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest {
	r.position = &position
	return r
}

// The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/reorder&#x60; request.
func (r ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest) OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule OpaqueAndExternalWebRequestRule) ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest {
	r.opaqueAndExternalWebRequestRule = &opaqueAndExternalWebRequestRule
	return r
}

func (r ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.CreateOpaqueAndExternalWebRequestDetectionRuleExecute(r)
}

/*
CreateOpaqueAndExternalWebRequestDetectionRule Creates a new service detection rule

The body must not provide an ID as it will be automatically assigned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) CreateOpaqueAndExternalWebRequestDetectionRule(ctx context.Context) ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest {
	return ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) CreateOpaqueAndExternalWebRequestDetectionRuleExecute(r ApiCreateOpaqueAndExternalWebRequestDetectionRuleRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.CreateOpaqueAndExternalWebRequestDetectionRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST"

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
	localVarPostBody = r.opaqueAndExternalWebRequestRule
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

type ApiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	id string
}

func (r ApiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOpaqueAndExternalWebRequestDetectionRuleExecute(r)
}

/*
DeleteOpaqueAndExternalWebRequestDetectionRule Deletes the specified service detection rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the service detection rule to be deleted.
 @return ApiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) DeleteOpaqueAndExternalWebRequestDetectionRule(ctx context.Context, id string) ApiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest {
	return ApiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) DeleteOpaqueAndExternalWebRequestDetectionRuleExecute(r ApiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.DeleteOpaqueAndExternalWebRequestDetectionRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}"
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

type ApiGetOpaqueAndExternalWebRequestDetectionRuleRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	id string
}

func (r ApiGetOpaqueAndExternalWebRequestDetectionRuleRequest) Execute() (*OpaqueAndExternalWebRequestRule, *http.Response, error) {
	return r.ApiService.GetOpaqueAndExternalWebRequestDetectionRuleExecute(r)
}

/*
GetOpaqueAndExternalWebRequestDetectionRule Shows the properties of the specified service detection rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the required service detection rule.
 @return ApiGetOpaqueAndExternalWebRequestDetectionRuleRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) GetOpaqueAndExternalWebRequestDetectionRule(ctx context.Context, id string) ApiGetOpaqueAndExternalWebRequestDetectionRuleRequest {
	return ApiGetOpaqueAndExternalWebRequestDetectionRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return OpaqueAndExternalWebRequestRule
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) GetOpaqueAndExternalWebRequestDetectionRuleExecute(r ApiGetOpaqueAndExternalWebRequestDetectionRuleRequest) (*OpaqueAndExternalWebRequestRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpaqueAndExternalWebRequestRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.GetOpaqueAndExternalWebRequestDetectionRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}"
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

type ApiListOpaqueAndExternalWebRequestDetectionRulesRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
}

func (r ApiListOpaqueAndExternalWebRequestDetectionRulesRequest) Execute() (*StubList, *http.Response, error) {
	return r.ApiService.ListOpaqueAndExternalWebRequestDetectionRulesExecute(r)
}

/*
ListOpaqueAndExternalWebRequestDetectionRules Lists all opaque and external web request service detection rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOpaqueAndExternalWebRequestDetectionRulesRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) ListOpaqueAndExternalWebRequestDetectionRules(ctx context.Context) ApiListOpaqueAndExternalWebRequestDetectionRulesRequest {
	return ApiListOpaqueAndExternalWebRequestDetectionRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StubList
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) ListOpaqueAndExternalWebRequestDetectionRulesExecute(r ApiListOpaqueAndExternalWebRequestDetectionRulesRequest) (*StubList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StubList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.ListOpaqueAndExternalWebRequestDetectionRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST"

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

type ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	stubList *StubList
}

// The JSON body of the request containing the service detection rules in the required order.
func (r ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest) StubList(stubList StubList) ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest {
	r.stubList = &stubList
	return r
}

func (r ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest) Execute() (*http.Response, error) {
	return r.ApiService.OrderOpaqueAndExternalWebRequestDetectionRulesExecute(r)
}

/*
OrderOpaqueAndExternalWebRequestDetectionRules Reorders the service detection rules of the specified type

The request reorders the rules of the specified type according to the order of the IDs in the body of the request. 

 Rules that are omitted in the body of the request will retain their relative order but will be placed *after* all those present in the request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) OrderOpaqueAndExternalWebRequestDetectionRules(ctx context.Context) ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest {
	return ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) OrderOpaqueAndExternalWebRequestDetectionRulesExecute(r ApiOrderOpaqueAndExternalWebRequestDetectionRulesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.OrderOpaqueAndExternalWebRequestDetectionRules")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/order"

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

type ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	id string
	opaqueAndExternalWebRequestRule *OpaqueAndExternalWebRequestRule
}

// The JSON body of the request containing updated parameters of the service detection rule.
func (r ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest) OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule OpaqueAndExternalWebRequestRule) ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest {
	r.opaqueAndExternalWebRequestRule = &opaqueAndExternalWebRequestRule
	return r
}

func (r ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest) Execute() (*EntityShortRepresentation, *http.Response, error) {
	return r.ApiService.UpdateOpaqueAndExternalWebRequestDetectionRuleExecute(r)
}

/*
UpdateOpaqueAndExternalWebRequestDetectionRule Updates an existing service detection rule

If the rule with the specified ID doesn't exist, a new rule will be created and appended to the end of the rule list. 

 The request keeps an existing order of rules, unless the **order** parameter is set.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the rule to be updated.
 @return ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) UpdateOpaqueAndExternalWebRequestDetectionRule(ctx context.Context, id string) ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest {
	return ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return EntityShortRepresentation
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) UpdateOpaqueAndExternalWebRequestDetectionRuleExecute(r ApiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest) (*EntityShortRepresentation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityShortRepresentation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.UpdateOpaqueAndExternalWebRequestDetectionRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}"
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
	localVarPostBody = r.opaqueAndExternalWebRequestRule
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

type ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	opaqueAndExternalWebRequestRule *OpaqueAndExternalWebRequestRule
}

func (r ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest) OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule OpaqueAndExternalWebRequestRule) ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest {
	r.opaqueAndExternalWebRequestRule = &opaqueAndExternalWebRequestRule
	return r
}

func (r ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateCreateOpaqueAndExternalWebRequestDetectionRuleExecute(r)
}

/*
ValidateCreateOpaqueAndExternalWebRequestDetectionRule Validates the payload for the `POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) ValidateCreateOpaqueAndExternalWebRequestDetectionRule(ctx context.Context) ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest {
	return ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) ValidateCreateOpaqueAndExternalWebRequestDetectionRuleExecute(r ApiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.ValidateCreateOpaqueAndExternalWebRequestDetectionRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/validator"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.opaqueAndExternalWebRequestRule == nil {
		return nil, reportError("opaqueAndExternalWebRequestRule is required and must be specified")
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
	localVarPostBody = r.opaqueAndExternalWebRequestRule
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

type ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest struct {
	ctx context.Context
	ApiService *ServiceDetectionOpaqueAndExternalWebRequestAPIService
	id string
	opaqueAndExternalWebRequestRule *OpaqueAndExternalWebRequestRule
}

func (r ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest) OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule OpaqueAndExternalWebRequestRule) ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest {
	r.opaqueAndExternalWebRequestRule = &opaqueAndExternalWebRequestRule
	return r
}

func (r ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest) Execute() (*http.Response, error) {
	return r.ApiService.ValidateUpdateOpaqueAndExternalWebRequestDetectionRuleExecute(r)
}

/*
ValidateUpdateOpaqueAndExternalWebRequestDetectionRule Validate the payload for the `PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}` request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the service detection rule to be validated.
 @return ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest
*/
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) ValidateUpdateOpaqueAndExternalWebRequestDetectionRule(ctx context.Context, id string) ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest {
	return ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ServiceDetectionOpaqueAndExternalWebRequestAPIService) ValidateUpdateOpaqueAndExternalWebRequestDetectionRuleExecute(r ApiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceDetectionOpaqueAndExternalWebRequestAPIService.ValidateUpdateOpaqueAndExternalWebRequestDetectionRule")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}/validator"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.opaqueAndExternalWebRequestRule == nil {
		return nil, reportError("opaqueAndExternalWebRequestRule is required and must be specified")
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
	localVarPostBody = r.opaqueAndExternalWebRequestRule
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
