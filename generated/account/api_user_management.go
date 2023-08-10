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
	"reflect"
)


// UserManagementAPIService UserManagementAPI service
type UserManagementAPIService service

type ApiUsersControllerAddUserToGroupsRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	email string
	requestBody *[]string
}

// The body of the request. Contains a list of groups (specified by UUIDs) to which the user is to be added.    Any existing group membership remains unaffected.
func (r ApiUsersControllerAddUserToGroupsRequest) RequestBody(requestBody []string) ApiUsersControllerAddUserToGroupsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiUsersControllerAddUserToGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersControllerAddUserToGroupsExecute(r)
}

/*
UsersControllerAddUserToGroups Adds a user to groups. Any existing group membership remains unaffected.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param email The email address of the required user.
 @return ApiUsersControllerAddUserToGroupsRequest
*/
func (a *UserManagementAPIService) UsersControllerAddUserToGroups(ctx context.Context, accountUuid string, email string) ApiUsersControllerAddUserToGroupsRequest {
	return ApiUsersControllerAddUserToGroupsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		email: email,
	}
}

// Execute executes the request
func (a *UserManagementAPIService) UsersControllerAddUserToGroupsExecute(r ApiUsersControllerAddUserToGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerAddUserToGroups")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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

type ApiUsersControllerCreateUserForAccountRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	userEmailDto *UserEmailDto
}

// The JSON body of the request. Contains the email address of the new user.
func (r ApiUsersControllerCreateUserForAccountRequest) UserEmailDto(userEmailDto UserEmailDto) ApiUsersControllerCreateUserForAccountRequest {
	r.userEmailDto = &userEmailDto
	return r
}

func (r ApiUsersControllerCreateUserForAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersControllerCreateUserForAccountExecute(r)
}

/*
UsersControllerCreateUserForAccount Creates a new user in an account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @return ApiUsersControllerCreateUserForAccountRequest
*/
func (a *UserManagementAPIService) UsersControllerCreateUserForAccount(ctx context.Context, accountUuid string) ApiUsersControllerCreateUserForAccountRequest {
	return ApiUsersControllerCreateUserForAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
	}
}

// Execute executes the request
func (a *UserManagementAPIService) UsersControllerCreateUserForAccountExecute(r ApiUsersControllerCreateUserForAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerCreateUserForAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userEmailDto == nil {
		return nil, reportError("userEmailDto is required and must be specified")
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
	localVarPostBody = r.userEmailDto
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

type ApiUsersControllerGetUserGroupsRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	email string
}

func (r ApiUsersControllerGetUserGroupsRequest) Execute() (*GroupUserDto, *http.Response, error) {
	return r.ApiService.UsersControllerGetUserGroupsExecute(r)
}

/*
UsersControllerGetUserGroups Lists all groups of a user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param email The email address of the required user.
 @return ApiUsersControllerGetUserGroupsRequest
*/
func (a *UserManagementAPIService) UsersControllerGetUserGroups(ctx context.Context, accountUuid string, email string) ApiUsersControllerGetUserGroupsRequest {
	return ApiUsersControllerGetUserGroupsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		email: email,
	}
}

// Execute executes the request
//  @return GroupUserDto
func (a *UserManagementAPIService) UsersControllerGetUserGroupsExecute(r ApiUsersControllerGetUserGroupsRequest) (*GroupUserDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupUserDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerGetUserGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

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

type ApiUsersControllerGetUsersRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	serviceUsers *bool
}

// Specifies whether service users are included in results.
func (r ApiUsersControllerGetUsersRequest) ServiceUsers(serviceUsers bool) ApiUsersControllerGetUsersRequest {
	r.serviceUsers = &serviceUsers
	return r
}

func (r ApiUsersControllerGetUsersRequest) Execute() (*UserListDto, *http.Response, error) {
	return r.ApiService.UsersControllerGetUsersExecute(r)
}

/*
UsersControllerGetUsers Lists all users of an account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @return ApiUsersControllerGetUsersRequest
*/
func (a *UserManagementAPIService) UsersControllerGetUsers(ctx context.Context, accountUuid string) ApiUsersControllerGetUsersRequest {
	return ApiUsersControllerGetUsersRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
	}
}

// Execute executes the request
//  @return UserListDto
func (a *UserManagementAPIService) UsersControllerGetUsersExecute(r ApiUsersControllerGetUsersRequest) (*UserListDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserListDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerGetUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.serviceUsers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "service-users", r.serviceUsers, "")
	}
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

type ApiUsersControllerRemoveUserFromAccountRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	email string
}

func (r ApiUsersControllerRemoveUserFromAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersControllerRemoveUserFromAccountExecute(r)
}

/*
UsersControllerRemoveUserFromAccount Removes a user from an account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param email The email address of the required user.
 @return ApiUsersControllerRemoveUserFromAccountRequest
*/
func (a *UserManagementAPIService) UsersControllerRemoveUserFromAccount(ctx context.Context, accountUuid string, email string) ApiUsersControllerRemoveUserFromAccountRequest {
	return ApiUsersControllerRemoveUserFromAccountRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		email: email,
	}
}

// Execute executes the request
func (a *UserManagementAPIService) UsersControllerRemoveUserFromAccountExecute(r ApiUsersControllerRemoveUserFromAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerRemoveUserFromAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users/{email}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

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

type ApiUsersControllerRemoveUserFromGroupsRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	email string
	groupUuid *[]string
}

// A list of groups the user is no longer a member of.    To specify several groups, use the following format: &#x60;group-uuid&#x3D;aaaaaa&amp;group-uuid&#x3D;bbbb&#x60;.
func (r ApiUsersControllerRemoveUserFromGroupsRequest) GroupUuid(groupUuid []string) ApiUsersControllerRemoveUserFromGroupsRequest {
	r.groupUuid = &groupUuid
	return r
}

func (r ApiUsersControllerRemoveUserFromGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersControllerRemoveUserFromGroupsExecute(r)
}

/*
UsersControllerRemoveUserFromGroups Removes a user from groups

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param email The email address of the required user.
 @return ApiUsersControllerRemoveUserFromGroupsRequest
*/
func (a *UserManagementAPIService) UsersControllerRemoveUserFromGroups(ctx context.Context, accountUuid string, email string) ApiUsersControllerRemoveUserFromGroupsRequest {
	return ApiUsersControllerRemoveUserFromGroupsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		email: email,
	}
}

// Execute executes the request
func (a *UserManagementAPIService) UsersControllerRemoveUserFromGroupsExecute(r ApiUsersControllerRemoveUserFromGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerRemoveUserFromGroups")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users/{email}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupUuid == nil {
		return nil, reportError("groupUuid is required and must be specified")
	}

	{
		t := *r.groupUuid
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "group-uuid", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "group-uuid", t, "multi")
		}
	}
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

type ApiUsersControllerReplaceUserGroupsRequest struct {
	ctx context.Context
	ApiService *UserManagementAPIService
	accountUuid string
	email string
	requestBody *[]string
}

// The body of the request. Contains a list of groups (specified by UUIDs) where the user is to be a member.    The user will be removed from any group that is not specified here.
func (r ApiUsersControllerReplaceUserGroupsRequest) RequestBody(requestBody []string) ApiUsersControllerReplaceUserGroupsRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiUsersControllerReplaceUserGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UsersControllerReplaceUserGroupsExecute(r)
}

/*
UsersControllerReplaceUserGroups Sets group membership of a user. Any existing membership is overwritten.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountUuid The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
 @param email The email address of the required user.
 @return ApiUsersControllerReplaceUserGroupsRequest
*/
func (a *UserManagementAPIService) UsersControllerReplaceUserGroups(ctx context.Context, accountUuid string, email string) ApiUsersControllerReplaceUserGroupsRequest {
	return ApiUsersControllerReplaceUserGroupsRequest{
		ApiService: a,
		ctx: ctx,
		accountUuid: accountUuid,
		email: email,
	}
}

// Execute executes the request
func (a *UserManagementAPIService) UsersControllerReplaceUserGroupsExecute(r ApiUsersControllerReplaceUserGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserManagementAPIService.UsersControllerReplaceUserGroups")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iam/v1/accounts/{accountUuid}/users/{email}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"accountUuid"+"}", url.PathEscape(parameterValueToString(r.accountUuid, "accountUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"email"+"}", url.PathEscape(parameterValueToString(r.email, "email")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
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
	localVarPostBody = r.requestBody
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
