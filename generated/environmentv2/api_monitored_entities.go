/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// MonitoredEntitiesAPIService MonitoredEntitiesAPI service
type MonitoredEntitiesAPIService service

type ApiGetEntitiesRequest struct {
	ctx context.Context
	ApiService *MonitoredEntitiesAPIService
	nextPageKey *string
	pageSize *int64
	entitySelector *string
	from *string
	to *string
	fields *string
	sort *string
}

// The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters. 
func (r ApiGetEntitiesRequest) NextPageKey(nextPageKey string) ApiGetEntitiesRequest {
	r.nextPageKey = &nextPageKey
	return r
}

// The amount of entities.   If not set, 50 is used.
func (r ApiGetEntitiesRequest) PageSize(pageSize int64) ApiGetEntitiesRequest {
	r.pageSize = &pageSize
	return r
}

// Defines the scope of the query. Only entities matching the specified criteria are included into response.   You must set one of these criteria:   * Entity type: &#x60;type(\&quot;TYPE\&quot;)&#x60;  * Dynatrace entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. You can specify several IDs, separated by a comma (&#x60;entityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;). All requested entities must be of the same type.   You can add one or more of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.   * Tag: &#x60;tag(\&quot;value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. Any colons (&#x60;:&#x60;) that are part of the key or value must be escaped with a backslash(&#x60;\\&#x60;). Otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: &#x60;mzId(123)&#x60;  * Management zone name: &#x60;mzName(\&quot;value\&quot;)&#x60; * Entity name:   * &#x60;entityName.equals&#x60;: performs a non-casesensitive &#x60;EQUALS&#x60; query.   * &#x60;entityName.startsWith&#x60;: changes the operator to &#x60;BEGINS WITH&#x60;.   * &#x60;entityName.in&#x60;: enables you to provide multiple values. The &#x60;EQUALS&#x60; operator applies.   * &#x60;caseSensitive(entityName.equals(\&quot;value\&quot;))&#x60;: takes any entity name criterion as an argument and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; * First seen timestamp: &#x60;firstSeenTms.&lt;operator&gt;(now-3h)&#x60;. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * &#x60;lte&#x60;: earlier than or at the specified time  * &#x60;lt&#x60;: earlier than the specified time  * &#x60;gte&#x60;: later than or at the specified time  * &#x60;gt&#x60;: later than the specified time * Entity attribute: &#x60;&lt;attribute&gt;(\&quot;value1\&quot;,\&quot;value2\&quot;)&#x60; and &#x60;&lt;attribute&gt;.exists()&#x60;. To fetch the list of available attributes, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **properties** field of the response.  * Relationships: &#x60;fromRelationships.&lt;relationshipName&gt;()&#x60; and &#x60;toRelationships.&lt;relationshipName&gt;()&#x60;. This criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **fromRelationships** and **toRelationships** fields. * Negation: &#x60;not(&lt;criterion&gt;)&#x60;. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;type(\&quot;HOST\&quot;),healthState(\&quot;HEALTHY\&quot;)&#x60;. Only results matching **all** criteria are included in the response.   The maximum string length is 2,000 characters.   The field is **required** when you&#39;re querying the first page of results.
func (r ApiGetEntitiesRequest) EntitySelector(entitySelector string) ApiGetEntitiesRequest {
	r.entitySelector = &entitySelector
	return r
}

// The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of three days is used (&#x60;now-3d&#x60;).
func (r ApiGetEntitiesRequest) From(from string) ApiGetEntitiesRequest {
	r.from = &from
	return r
}

// The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used.
func (r ApiGetEntitiesRequest) To(to string) ApiGetEntitiesRequest {
	r.to = &to
	return r
}

// Defines the list of entity properties included in the response. The ID and the name of an entity are **always** included to the response.   To add properties, list them with leading plus &#x60;+&#x60;. You can specify several properties, separated by a comma (for example &#x60;fields&#x3D;+lastSeenTms,+properties.BITNESS&#x60;).   Use the [GET entity type](https://dt-url.net/2ka3ivt) request to fetch the list of properties available for your entity type. Fields from the **properties** object must be specified in the &#x60;properties.FIELD&#x60; format (for example, &#x60;properties.BITNESS&#x60;).
func (r ApiGetEntitiesRequest) Fields(fields string) ApiGetEntitiesRequest {
	r.fields = &fields
	return r
}

// Defines the ordering of the entities returned.   This field is **optional**, each field has a sign prefix (+/-), which corresponds to sorting order ( + for ascending and - for descending). If no sign prefix is set, then default ascending sorting order will be applied.   Currently ordering is only available for the display name (for example &#x60;sort&#x3D;name or sort &#x3D;+name for ascending, sort&#x3D;-name for descending&#x60;)
func (r ApiGetEntitiesRequest) Sort(sort string) ApiGetEntitiesRequest {
	r.sort = &sort
	return r
}

func (r ApiGetEntitiesRequest) Execute() (*EntitiesList, *http.Response, error) {
	return r.ApiService.GetEntitiesExecute(r)
}

/*
GetEntities Gets the information about monitored entities

Lists entities observed within the specified timeframe along with their properties. 

When you query entities of the `SERVICE_METHOD` type, only the following requests are returned: 

* [Key requests](https://dt-url.net/a903u9s) 
* Top X requests that are used for [baselining](https://dt-url.net/0j23uqb) 
* Requests that have caused a [problem](https://dt-url.net/pf43uqg) 

You can limit the output by using pagination: 
1. Specify the number of results per page in the **pageSize** query parameter. 
2. Use the cursor from the **nextPageKey** field of the previous response in the **nextPageKey** query parameter to obtain subsequent pages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEntitiesRequest
*/
func (a *MonitoredEntitiesAPIService) GetEntities(ctx context.Context) ApiGetEntitiesRequest {
	return ApiGetEntitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntitiesList
func (a *MonitoredEntitiesAPIService) GetEntitiesExecute(r ApiGetEntitiesRequest) (*EntitiesList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntitiesList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredEntitiesAPIService.GetEntities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextPageKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageKey", r.nextPageKey, "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "")
	}
	if r.entitySelector != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entitySelector", r.entitySelector, "")
	}
	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
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

type ApiGetEntityRequest struct {
	ctx context.Context
	ApiService *MonitoredEntitiesAPIService
	entityId string
	from *string
	to *string
	fields *string
}

// The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of three days is used (&#x60;now-3d&#x60;).
func (r ApiGetEntityRequest) From(from string) ApiGetEntityRequest {
	r.from = &from
	return r
}

// The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used.
func (r ApiGetEntityRequest) To(to string) ApiGetEntityRequest {
	r.to = &to
	return r
}

// Defines the list of entity properties included in the response. The ID and the name of an entity are **always** included to the response.   To add properties, list them with leading plus &#x60;+&#x60;. You can specify several properties, separated by a comma (for example &#x60;fields&#x3D;+lastSeenTms,+properties.BITNESS&#x60;).   Use the [GET entity type](https://dt-url.net/2ka3ivt) request to fetch the list of properties available for your entity type. Fields from the **properties** object must be specified in the &#x60;properties.FIELD&#x60; format (for example, &#x60;properties.BITNESS&#x60;).
func (r ApiGetEntityRequest) Fields(fields string) ApiGetEntityRequest {
	r.fields = &fields
	return r
}

func (r ApiGetEntityRequest) Execute() (*Entity, *http.Response, error) {
	return r.ApiService.GetEntityExecute(r)
}

/*
GetEntity Gets the properties of the specified monitored entity

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityId The ID of the required entity.
 @return ApiGetEntityRequest
*/
func (a *MonitoredEntitiesAPIService) GetEntity(ctx context.Context, entityId string) ApiGetEntityRequest {
	return ApiGetEntityRequest{
		ApiService: a,
		ctx: ctx,
		entityId: entityId,
	}
}

// Execute executes the request
//  @return Entity
func (a *MonitoredEntitiesAPIService) GetEntityExecute(r ApiGetEntityRequest) (*Entity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Entity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredEntitiesAPIService.GetEntity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entities/{entityId}"
	localVarPath = strings.Replace(localVarPath, "{"+"entityId"+"}", url.PathEscape(parameterValueToString(r.entityId, "entityId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.from != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from", r.from, "")
	}
	if r.to != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to", r.to, "")
	}
	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "")
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

type ApiGetEntityTypeRequest struct {
	ctx context.Context
	ApiService *MonitoredEntitiesAPIService
	type_ string
}

func (r ApiGetEntityTypeRequest) Execute() (*EntityType, *http.Response, error) {
	return r.ApiService.GetEntityTypeExecute(r)
}

/*
GetEntityType Gets a list of properties for the specified entity type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param type_ The required entity type.
 @return ApiGetEntityTypeRequest
*/
func (a *MonitoredEntitiesAPIService) GetEntityType(ctx context.Context, type_ string) ApiGetEntityTypeRequest {
	return ApiGetEntityTypeRequest{
		ApiService: a,
		ctx: ctx,
		type_: type_,
	}
}

// Execute executes the request
//  @return EntityType
func (a *MonitoredEntitiesAPIService) GetEntityTypeExecute(r ApiGetEntityTypeRequest) (*EntityType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredEntitiesAPIService.GetEntityType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entityTypes/{type}"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", url.PathEscape(parameterValueToString(r.type_, "type_")), -1)

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

type ApiGetEntityTypesRequest struct {
	ctx context.Context
	ApiService *MonitoredEntitiesAPIService
	nextPageKey *string
	pageSize *int64
}

// The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters. 
func (r ApiGetEntityTypesRequest) NextPageKey(nextPageKey string) ApiGetEntityTypesRequest {
	r.nextPageKey = &nextPageKey
	return r
}

// The amount of entity types in a single response payload.   The maximal allowed page size is 500.   If not set, 50 is used.
func (r ApiGetEntityTypesRequest) PageSize(pageSize int64) ApiGetEntityTypesRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiGetEntityTypesRequest) Execute() (*EntityTypeList, *http.Response, error) {
	return r.ApiService.GetEntityTypesExecute(r)
}

/*
GetEntityTypes Gets a list of properties for all entity types

You can limit the output by using pagination: 
1. Specify the number of results per page in the **pageSize** query parameter. 
2. Use the cursor from the **nextPageKey** field of the previous response in the **nextPageKey** query parameter to obtain subsequent pages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetEntityTypesRequest
*/
func (a *MonitoredEntitiesAPIService) GetEntityTypes(ctx context.Context) ApiGetEntityTypesRequest {
	return ApiGetEntityTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntityTypeList
func (a *MonitoredEntitiesAPIService) GetEntityTypesExecute(r ApiGetEntityTypesRequest) (*EntityTypeList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityTypeList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredEntitiesAPIService.GetEntityTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entityTypes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextPageKey != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextPageKey", r.nextPageKey, "")
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

type ApiPushCustomDeviceRequest struct {
	ctx context.Context
	ApiService *MonitoredEntitiesAPIService
	customDeviceCreation *CustomDeviceCreation
}

// The JSON body of the request. Contains parameters of a custom device.
func (r ApiPushCustomDeviceRequest) CustomDeviceCreation(customDeviceCreation CustomDeviceCreation) ApiPushCustomDeviceRequest {
	r.customDeviceCreation = &customDeviceCreation
	return r
}

func (r ApiPushCustomDeviceRequest) Execute() (*CustomDeviceCreationResult, *http.Response, error) {
	return r.ApiService.PushCustomDeviceExecute(r)
}

/*
PushCustomDevice Creates or updates a custom device

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPushCustomDeviceRequest
*/
func (a *MonitoredEntitiesAPIService) PushCustomDevice(ctx context.Context) ApiPushCustomDeviceRequest {
	return ApiPushCustomDeviceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomDeviceCreationResult
func (a *MonitoredEntitiesAPIService) PushCustomDeviceExecute(r ApiPushCustomDeviceRequest) (*CustomDeviceCreationResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomDeviceCreationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MonitoredEntitiesAPIService.PushCustomDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entities/custom"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customDeviceCreation == nil {
		return localVarReturnValue, nil, reportError("customDeviceCreation is required and must be specified")
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
	localVarPostBody = r.customDeviceCreation
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