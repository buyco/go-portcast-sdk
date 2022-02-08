# \BookingBookmarksApi

All URIs are relative to *https://api.portcast.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1EtaBookingPost**](BookingBookmarksApi.md#ApiV1EtaBookingPost) | **Post** /api/v1/eta/booking | Create new bookmarks for booking



## ApiV1EtaBookingPost

> BookingBookmarks ApiV1EtaBookingPost(ctx).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Body(body).Execute()

Create new bookmarks for booking



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
    accept := "application/json" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    xApiKey := "{{x-api-key}}" // string |  (optional)
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingBookmarksApi.ApiV1EtaBookingPost(context.Background()).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingBookmarksApi.ApiV1EtaBookingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaBookingPost`: BookingBookmarks
    fmt.Fprintf(os.Stdout, "Response from `BookingBookmarksApi.ApiV1EtaBookingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaBookingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**BookingBookmarks**](BookingBookmarks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

