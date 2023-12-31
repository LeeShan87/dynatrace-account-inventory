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


// ServiceUserManagementAPIService ServiceUserManagementAPI service
type ServiceUserManagementAPIService service

type ApiServiceUsersControllerCreateServiceUserForAccountRequest struct {
	ctx context.Context
	ApiService *ServiceUserManagementAPIService
	accountUuid string
	serviceUserNameDto *ServiceUserNameDto
}

// The JSON body of the request. Contains the name of the new service user.
func (r ApiServiceUsersControllerCreateServiceUserForAccountRequest) ServiceUserNameDto(serviceUserNameDto ServiceUserNameDto) ApiServiceUsersControllerCreateServiceUserForAccountRequest {
	r.serviceUserNameDto = &serviceUserNameDto
	return r
}

func (r ApiServiceUsersControllerCreateServiceUserForAccountRequest) Execute() (*ServiceUserUuidDto, *http.Response, error) {
	return r.ApiService.ServiceUsersControllerCreateServiceUserForAccountExecute(r)
}

/*
ServiceUsersControllerCreateServiceUserForAccount Creates a new service user in an account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @return ApiServiceUsersControllerCreateServiceUserForAccountRequest
*/
func (a *ServiceUserManagementAPIService) ServiceUsersControllerCreateServiceUserForAccount(ctx context.Context, accountUuid string) ApiServiceUsersControllerCreateServiceUserForAccountRequest {
	return ApiServiceUsersControllerCreateServiceUserForAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
	}
}

// Execute executes the request
//  @return ServiceUserUuidDto
func (a *ServiceUserManagementAPIService) ServiceUsersControllerCreateServiceUserForAccountExecute(r ApiServiceUsersControllerCreateServiceUserForAccountRequest) (*ServiceUserUuidDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ServiceUserUuidDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceUserManagementAPIService.ServiceUsersControllerCreateServiceUserForAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/service-users"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serviceUserNameDto == nil {
		return localVarReturnValue, nil, reportError("serviceUserNameDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.serviceUserNameDto
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

type ApiServiceUsersControllerDeleteUserRequest struct {
	ctx context.Context
	ApiService *ServiceUserManagementAPIService
	accountUuid string
	userUuid string
}

func (r ApiServiceUsersControllerDeleteUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.ServiceUsersControllerDeleteUserExecute(r)
}

/*
ServiceUsersControllerDeleteUser Removes service user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param userUuid The UUID of the required user.
 @return ApiServiceUsersControllerDeleteUserRequest
*/
func (a *ServiceUserManagementAPIService) ServiceUsersControllerDeleteUser(ctx context.Context, accountUuid string, userUuid string) ApiServiceUsersControllerDeleteUserRequest {
	return ApiServiceUsersControllerDeleteUserRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		userUuid: userUuid,
	}
}

// Execute executes the request
func (a *ServiceUserManagementAPIService) ServiceUsersControllerDeleteUserExecute(r ApiServiceUsersControllerDeleteUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceUserManagementAPIService.ServiceUsersControllerDeleteUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/service-users/{userUuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userUuid"+"}", url.PathEscape(parameterValueToString(r.userUuid, "userUuid")), -1)

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

type ApiServiceUsersControllerGetServiceUsersFromAccountRequest struct {
	ctx context.Context
	ApiService *ServiceUserManagementAPIService
	accountUuid string
}

func (r ApiServiceUsersControllerGetServiceUsersFromAccountRequest) Execute() (*ExternalServiceUsersPageDto, *http.Response, error) {
	return r.ApiService.ServiceUsersControllerGetServiceUsersFromAccountExecute(r)
}

/*
ServiceUsersControllerGetServiceUsersFromAccount Get service users assigned to account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @return ApiServiceUsersControllerGetServiceUsersFromAccountRequest
*/
func (a *ServiceUserManagementAPIService) ServiceUsersControllerGetServiceUsersFromAccount(ctx context.Context, accountUuid string) ApiServiceUsersControllerGetServiceUsersFromAccountRequest {
	return ApiServiceUsersControllerGetServiceUsersFromAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
	}
}

// Execute executes the request
//  @return ExternalServiceUsersPageDto
func (a *ServiceUserManagementAPIService) ServiceUsersControllerGetServiceUsersFromAccountExecute(r ApiServiceUsersControllerGetServiceUsersFromAccountRequest) (*ExternalServiceUsersPageDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExternalServiceUsersPageDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceUserManagementAPIService.ServiceUsersControllerGetServiceUsersFromAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/service-users"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)

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
