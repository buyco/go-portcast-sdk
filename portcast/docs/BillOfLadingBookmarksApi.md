# \BillOfLadingBookmarksApi

All URIs are relative to *https://api.portcast.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete**](BillOfLadingBookmarksApi.md#EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete) | **Delete** /eta/bill-of-lading-bookmarks/{bill_of_lading_bookmark_id} | Delete a particular bill of lading bookmark by ID
[**EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet**](BillOfLadingBookmarksApi.md#EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet) | **Get** /eta/bill-of-lading-bookmarks/{bill_of_lading_bookmark_id} | Get a particular bill of lading bookmark by ID
[**EtaBillOfLadingBookmarksGet**](BillOfLadingBookmarksApi.md#EtaBillOfLadingBookmarksGet) | **Get** /eta/bill-of-lading-bookmarks | Get a list of bookmarked bill of ladings
[**EtaBillOfLadingBookmarksPost**](BillOfLadingBookmarksApi.md#EtaBillOfLadingBookmarksPost) | **Post** /eta/bill-of-lading-bookmarks | Create a new bill of lading bookmark



## EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete

> MessageModel EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete(ctx, billOfLadingBookmarkId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete(context.Background(), billOfLadingBookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete`: MessageModel
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksBillOfLadingBookmarkIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingBookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEtaBillOfLadingBookmarksBillOfLadingBookmarkIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageModel**](MessageModel.md)

### Authorization

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet

> BookmarkResponse EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet(ctx, billOfLadingBookmarkId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet(context.Background(), billOfLadingBookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: BookmarkResponse
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingBookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEtaBillOfLadingBookmarksBillOfLadingBookmarkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BookmarkResponse**](BookmarkResponse.md)

### Authorization

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EtaBillOfLadingBookmarksGet

> InlineResponse200 EtaBillOfLadingBookmarksGet(ctx).StartAfter(startAfter).Ascending(ascending).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()

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
    startAfter := "startAfter_example" // string | start query from after this voyage bookmark id (optional)
    ascending := "ascending_example" // string | if set to true, sort results based on creation time in ascending order. false by default (optional)
    carrierNo := "carrierNo_example" // string | filter on the carrier scac (optional)
    blNo := "blNo_example" // string | filter on the bill of lading number (optional)
    cntrNo := "cntrNo_example" // string | filter on the container number (optional)
    status := "status_example" // string | filter on the status (optional)
    statusCode := "statusCode_example" // string | filter on the status code (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksGet(context.Background()).StartAfter(startAfter).Ascending(ascending).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaBillOfLadingBookmarksGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEtaBillOfLadingBookmarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EtaBillOfLadingBookmarksPost

> Bookmark EtaBillOfLadingBookmarksPost(ctx).Body(body).Execute()

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
    body := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaBillOfLadingBookmarksPost`: Bookmark
    fmt.Fprintf(os.Stdout, "Response from `BillOfLadingBookmarksApi.EtaBillOfLadingBookmarksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEtaBillOfLadingBookmarksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**Bookmark**](Bookmark.md)

### Authorization

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

