# \BookingBookmarksApi

All URIs are relative to *https://api.portcast.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EtaBookingPost**](BookingBookmarksApi.md#EtaBookingPost) | **Post** /eta/booking | Create new bookmarks for booking



## EtaBookingPost

> BookingBookmarks EtaBookingPost(ctx).Body(body).Execute()

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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BookingBookmarksApi.EtaBookingPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BookingBookmarksApi.EtaBookingPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaBookingPost`: BookingBookmarks
    fmt.Fprintf(os.Stdout, "Response from `BookingBookmarksApi.EtaBookingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEtaBookingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**BookingBookmarks**](BookingBookmarks.md)

### Authorization

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

