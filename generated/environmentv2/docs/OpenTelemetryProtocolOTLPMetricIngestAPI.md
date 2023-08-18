# \OpenTelemetryProtocolOTLPMetricIngestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IngestMetric**](OpenTelemetryProtocolOTLPMetricIngestAPI.md#IngestMetric) | **Post** /otlp/v1/metrics | Implementation of the OTLP/HTTP protocol for metric ingest.



## IngestMetric

> IngestMetric(ctx).RequestBody(requestBody).Execute()

Implementation of the OTLP/HTTP protocol for metric ingest.



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
    requestBody := []string{string(123)} // []string | An ExportMetricServiceRequest message in binary protobuf format.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OpenTelemetryProtocolOTLPMetricIngestAPI.IngestMetric(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenTelemetryProtocolOTLPMetricIngestAPI.IngestMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIngestMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | An ExportMetricServiceRequest message in binary protobuf format. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/x-protobuf
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

