# \DefaultApi

All URIs are relative to *https://api.portcast.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthcheckGet**](DefaultApi.md#HealthcheckGet) | **Get** /healthcheck | Health Check API



## HealthcheckGet

> MessageModel HealthcheckGet(ctx).Execute()

Health Check API

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.HealthcheckGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HealthcheckGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthcheckGet`: MessageModel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.HealthcheckGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthcheckGetRequest struct via the builder pattern


### Return type

[**MessageModel**](MessageModel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

