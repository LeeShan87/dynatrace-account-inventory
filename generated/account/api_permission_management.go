/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PermissionManagementAPIService PermissionManagementAPI service
type PermissionManagementAPIService service

type ApiPermissionsControllerAddGroupPermissionsRequest struct {
	ctx context.Context
	ApiService *PermissionManagementAPIService
	accountUuid string
	groupUuid string
	permissionsDto *[]PermissionsDto
}

// The body of the request. Contains a list of permissions to be assigned to the group.   Existing permissions remain unaffected.
func (r ApiPermissionsControllerAddGroupPermissionsRequest) PermissionsDto(permissionsDto []PermissionsDto) ApiPermissionsControllerAddGroupPermissionsRequest {
	r.permissionsDto = &permissionsDto
	return r
}

func (r ApiPermissionsControllerAddGroupPermissionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PermissionsControllerAddGroupPermissionsExecute(r)
}

/*
PermissionsControllerAddGroupPermissions Assigns permissions to a user group. Existing permissions remain unaffected.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param groupUuid The UUID of the required user group.
 @return ApiPermissionsControllerAddGroupPermissionsRequest
*/
func (a *PermissionManagementAPIService) PermissionsControllerAddGroupPermissions(ctx context.Context, accountUuid string, groupUuid string) ApiPermissionsControllerAddGroupPermissionsRequest {
	return ApiPermissionsControllerAddGroupPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		groupUuid: groupUuid,
	}
}

// Execute executes the request
func (a *PermissionManagementAPIService) PermissionsControllerAddGroupPermissionsExecute(r ApiPermissionsControllerAddGroupPermissionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementAPIService.PermissionsControllerAddGroupPermissions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupUuid"+"}", url.PathEscape(parameterValueToString(r.groupUuid, "groupUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.permissionsDto == nil {
		return nil, reportError("permissionsDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.permissionsDto
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

type ApiPermissionsControllerGetGroupPermissionsRequest struct {
	ctx context.Context
	ApiService *PermissionManagementAPIService
	accountUuid string
	groupUuid string
}

func (r ApiPermissionsControllerGetGroupPermissionsRequest) Execute() (*PermissionsGroupDto, *http.Response, error) {
	return r.ApiService.PermissionsControllerGetGroupPermissionsExecute(r)
}

/*
PermissionsControllerGetGroupPermissions Lists all permissions of a user group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param groupUuid The UUID of the required user group.
 @return ApiPermissionsControllerGetGroupPermissionsRequest
*/
func (a *PermissionManagementAPIService) PermissionsControllerGetGroupPermissions(ctx context.Context, accountUuid string, groupUuid string) ApiPermissionsControllerGetGroupPermissionsRequest {
	return ApiPermissionsControllerGetGroupPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		groupUuid: groupUuid,
	}
}

// Execute executes the request
//  @return PermissionsGroupDto
func (a *PermissionManagementAPIService) PermissionsControllerGetGroupPermissionsExecute(r ApiPermissionsControllerGetGroupPermissionsRequest) (*PermissionsGroupDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PermissionsGroupDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementAPIService.PermissionsControllerGetGroupPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupUuid"+"}", url.PathEscape(parameterValueToString(r.groupUuid, "groupUuid")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ApiPermissionsControllerOverwriteGroupPermissionsRequest struct {
	ctx context.Context
	ApiService *PermissionManagementAPIService
	accountUuid string
	groupUuid string
	permissionsDto *[]PermissionsDto
}

// The body of the request. Contains a list of permissions to be assigned to the group.    Existing permissions are overwritten.
func (r ApiPermissionsControllerOverwriteGroupPermissionsRequest) PermissionsDto(permissionsDto []PermissionsDto) ApiPermissionsControllerOverwriteGroupPermissionsRequest {
	r.permissionsDto = &permissionsDto
	return r
}

func (r ApiPermissionsControllerOverwriteGroupPermissionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PermissionsControllerOverwriteGroupPermissionsExecute(r)
}

/*
PermissionsControllerOverwriteGroupPermissions Sets permissions of a user group. Existing permissions are overwritten.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param groupUuid The UUID of the required user group.
 @return ApiPermissionsControllerOverwriteGroupPermissionsRequest
*/
func (a *PermissionManagementAPIService) PermissionsControllerOverwriteGroupPermissions(ctx context.Context, accountUuid string, groupUuid string) ApiPermissionsControllerOverwriteGroupPermissionsRequest {
	return ApiPermissionsControllerOverwriteGroupPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		groupUuid: groupUuid,
	}
}

// Execute executes the request
func (a *PermissionManagementAPIService) PermissionsControllerOverwriteGroupPermissionsExecute(r ApiPermissionsControllerOverwriteGroupPermissionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementAPIService.PermissionsControllerOverwriteGroupPermissions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupUuid"+"}", url.PathEscape(parameterValueToString(r.groupUuid, "groupUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.permissionsDto == nil {
		return nil, reportError("permissionsDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.permissionsDto
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

type ApiPermissionsControllerRemoveGroupPermissionsRequest struct {
	ctx context.Context
	ApiService *PermissionManagementAPIService
	accountUuid string
	groupUuid string
	scope *string
	permissionName *string
	scopeType *string
}

// The scope of the permission to be deleted. Depending on the type of the scope, specify one of the following:    * &#x60;account&#x60;: The UUID of the account.  * &#x60;tenant&#x60;: The ID of the environment.  * &#x60;management-zone&#x60;: The ID of the management zone from an environment in &#x60;{environment-id}:{management-zone-id}&#x60; format.
func (r ApiPermissionsControllerRemoveGroupPermissionsRequest) Scope(scope string) ApiPermissionsControllerRemoveGroupPermissionsRequest {
	r.scope = &scope
	return r
}

// The name of the permission to be deleted.
func (r ApiPermissionsControllerRemoveGroupPermissionsRequest) PermissionName(permissionName string) ApiPermissionsControllerRemoveGroupPermissionsRequest {
	r.permissionName = &permissionName
	return r
}

// The scope type of the permission to be deleted.
func (r ApiPermissionsControllerRemoveGroupPermissionsRequest) ScopeType(scopeType string) ApiPermissionsControllerRemoveGroupPermissionsRequest {
	r.scopeType = &scopeType
	return r
}

func (r ApiPermissionsControllerRemoveGroupPermissionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PermissionsControllerRemoveGroupPermissionsExecute(r)
}

/*
PermissionsControllerRemoveGroupPermissions Removes a permission from a user group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param groupUuid The UUID of the required user group.
 @return ApiPermissionsControllerRemoveGroupPermissionsRequest
*/
func (a *PermissionManagementAPIService) PermissionsControllerRemoveGroupPermissions(ctx context.Context, accountUuid string, groupUuid string) ApiPermissionsControllerRemoveGroupPermissionsRequest {
	return ApiPermissionsControllerRemoveGroupPermissionsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		groupUuid: groupUuid,
	}
}

// Execute executes the request
func (a *PermissionManagementAPIService) PermissionsControllerRemoveGroupPermissionsExecute(r ApiPermissionsControllerRemoveGroupPermissionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionManagementAPIService.PermissionsControllerRemoveGroupPermissions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupUuid"+"}", url.PathEscape(parameterValueToString(r.groupUuid, "groupUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.scope == nil {
		return nil, reportError("scope is required and must be specified")
	}
	if r.permissionName == nil {
		return nil, reportError("permissionName is required and must be specified")
	}
	if r.scopeType == nil {
		return nil, reportError("scopeType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "permission-name", r.permissionName, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "scope-type", r.scopeType, "")
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
