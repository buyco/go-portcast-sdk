# \TrackingApi

All URIs are relative to *https://api.portcast.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet**](TrackingApi.md#ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet) | **Get** /api/v1/eta/tracking/bill-of-lading-bookmarks/{bill_of_lading_bookmark_id} | Get bill of lading bookmark tracking results for a particular bill of lading bookmark by ID
[**ApiV1EtaTrackingBillOfLadingBookmarksGet**](TrackingApi.md#ApiV1EtaTrackingBillOfLadingBookmarksGet) | **Get** /api/v1/eta/tracking/bill-of-lading-bookmarks | Get a list of bill of lading bookmark tracking results based on filters



## ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet

> InlineResponse2002 ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet(ctx, billOfLadingBookmarkId).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Execute()

Get bill of lading bookmark tracking results for a particular bill of lading bookmark by ID



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
    resp, r, err := api_client.TrackingApi.ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet(context.Background(), billOfLadingBookmarkId).Accept(accept).ContentType(contentType).XApiKey(xApiKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingApi.ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `TrackingApi.ApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingBookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1EtaTrackingBillOfLadingBookmarksGet

> InlineResponse2001 ApiV1EtaTrackingBillOfLadingBookmarksGet(ctx).Accept(accept).ContentType(contentType).XApiKey(xApiKey).StartAfter(startAfter).CustomerId(customerId).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()

Get a list of bill of lading bookmark tracking results based on filters



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
    customerId := "customerId_example" // string | get last prediction result made before this timestamp, must be in ISO8601 format with timezone, defaults to now (optional)
    carrierNo := "carrierNo_example" // string | filter on the carrier scac code (optional)
    blNo := "blNo_example" // string | filter on the imo number of the vessel (optional)
    cntrNo := "cntrNo_example" // string | filter on the target port by port code (optional)
    status := "status_example" // string | filter on the target port eta (optional)
    statusCode := "statusCode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrackingApi.ApiV1EtaTrackingBillOfLadingBookmarksGet(context.Background()).Accept(accept).ContentType(contentType).XApiKey(xApiKey).StartAfter(startAfter).CustomerId(customerId).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingApi.ApiV1EtaTrackingBillOfLadingBookmarksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiV1EtaTrackingBillOfLadingBookmarksGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TrackingApi.ApiV1EtaTrackingBillOfLadingBookmarksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1EtaTrackingBillOfLadingBookmarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accept** | **string** |  | 
 **contentType** | **string** |  | 
 **xApiKey** | **string** |  | 
 **startAfter** | **string** | start query from after this voyage bookmark id | 
 **customerId** | **string** | get last prediction result made before this timestamp, must be in ISO8601 format with timezone, defaults to now | 
 **carrierNo** | **string** | filter on the carrier scac code | 
 **blNo** | **string** | filter on the imo number of the vessel | 
 **cntrNo** | **string** | filter on the target port by port code | 
 **status** | **string** | filter on the target port eta | 
 **statusCode** | **string** |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

