# \BillOfLadingBookmarksApi

All URIs are relative to *https://api.portcast.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete**](BillOfLadingBookmarksApi.md#ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete) | **Delete** /api/v1/eta/bill-of-lading-bookmarks/{bill_of_lading_bookmark_id} | Delete a particular bill of lading bookmark by ID
[**ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet**](BillOfLadingBookmarksApi.md#ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet) | **Get** /api/v1/eta/bill-of-lading-bookmarks/{bill_of_lading_bookmark_id} | Get a particular bill of lading bookmark by ID
[**ApiV1EtaBillOfLadingBookmarksGet**](BillOfLadingBookmarksApi.md#ApiV1EtaBillOfLadingBookmarksGet) | **Get** /api/v1/eta/bill-of-lading-bookmarks | Get a list of bookmarked bill of ladings
[**ApiV1EtaBillOfLadingBookmarksPost**](BillOfLadingBookmarksApi.md#ApiV1EtaBillOfLadingBookmarksPost) | **Post** /api/v1/eta/bill-of-lading-bookmarks | Create a new bill of lading bookmark



## ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete

> MessageModel ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete(ctx, billOfLadingBookmarkId).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Execute()

Delete a particular bill of lading bookmark by ID

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
    billOfLadingBookmarkId := "billOfLadingBookmarkId_example" // string | 
    accept := "application/json" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    xApiKey := "{{x-api-key}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete(context.Background(), billOfLadingBookmarkId).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete`: MessageModel
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingBookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 

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


## ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet

> BookmarkResponse ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet(ctx, billOfLadingBookmarkId).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Execute()

Get a particular bill of lading bookmark by ID

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
    billOfLadingBookmarkId := "billOfLadingBookmarkId_example" // string | 
    accept := "application/json" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    xApiKey := "{{x-api-key}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet(context.Background(), billOfLadingBookmarkId).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: BookmarkResponse
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingBookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 

### Return type

[**BookmarkResponse**](BookmarkResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1EtaBillOfLadingBookmarksGet

> InlineResponse200 ApiV1EtaBillOfLadingBookmarksGet(ctx).Accept(accept).ContentType(contentType).XApiKey(xApiKey).StartAfter(startAfter).Ascending(ascending).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()

Get a list of bookmarked bill of ladings



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
    startAfter := "startAfter_example" // string | start query from after this voyage bookmark id (optional)
    ascending := "ascending_example" // string | if set to true, sort results based on creation time in ascending order. false by default (optional)
    carrierNo := "carrierNo_example" // string | filter on the carrier scac (optional)
    blNo := "blNo_example" // string | filter on the bill of lading number (optional)
    cntrNo := "cntrNo_example" // string | filter on the container number (optional)
    status := "status_example" // string | filter on the status (optional)
    statusCode := "statusCode_example" // string | filter on the status code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksGet(context.Background()).Accept(accept).ContentType(contentType).XApiKey(xApiKey).StartAfter(startAfter).Ascending(ascending).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaBillOfLadingBookmarksGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaBillOfLadingBookmarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 
 **startAfter** | **string** | start query from after this voyage bookmark id | 
 **ascending** | **string** | if set to true, sort results based on creation time in ascending order. false by default | 
 **carrierNo** | **string** | filter on the carrier scac | 
 **blNo** | **string** | filter on the bill of lading number | 
 **cntrNo** | **string** | filter on the container number | 
 **status** | **string** | filter on the status | 
 **statusCode** | **string** | filter on the status code | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1EtaBillOfLadingBookmarksPost

> Bookmark ApiV1EtaBillOfLadingBookmarksPost(ctx).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Body(body).Execute()

Create a new bill of lading bookmark



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
    resp, r, err := api_client.BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksPost(context.Background()).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaBillOfLadingBookmarksPost`: Bookmark
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.ApiV1EtaBillOfLadingBookmarksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaBillOfLadingBookmarksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**Bookmark**](Bookmark.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

