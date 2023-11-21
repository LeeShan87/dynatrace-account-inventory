# \BusinessEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Ingest1**](BusinessEventsAPI.md#Ingest1) | **Post** /bizevents/ingest | Ingests a business event



## Ingest1

> Ingest1(ctx).CloudEvent(cloudEvent).Execute()

Ingests a business event



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    cloudEvent := *openapiclient.NewCloudEvent("Id_example", "Source_example", "Specversion_example", "Type_example") // CloudEvent | The Business Event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BusinessEventsAPI.Ingest1(context.Background()).CloudEvent(cloudEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusinessEventsAPI.Ingest1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngest1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudEvent** | [**CloudEvent**](CloudEvent.md) | The Business Event | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/cloudevent+json, application/cloudevent-batch+json, application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

