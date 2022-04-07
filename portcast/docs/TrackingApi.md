# \TrackingApi

All URIs are relative to *https://api.portcast.io/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet**](TrackingApi.md#EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet) | **Get** /eta/tracking/bill-of-lading-bookmarks/{bill_of_lading_bookmark_id} | Get bill of lading bookmark tracking results for a particular bill of lading bookmark by ID
[**EtaTrackingBillOfLadingBookmarksGet**](TrackingApi.md#EtaTrackingBillOfLadingBookmarksGet) | **Get** /eta/tracking/bill-of-lading-bookmarks | Get a list of bill of lading bookmark tracking results based on filters



## EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet

> InlineResponse2002 EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet(ctx, billOfLadingBookmarkId).Execute()

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrackingApi.EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet(context.Background(), billOfLadingBookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingApi.EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `TrackingApi.EtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**billOfLadingBookmarkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEtaTrackingBillOfLadingBookmarksBillOfLadingBookmarkIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EtaTrackingBillOfLadingBookmarksGet

> InlineResponse2001 EtaTrackingBillOfLadingBookmarksGet(ctx).StartAfter(startAfter).CustomerId(customerId).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()

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
    startAfter := "startAfter_example" // string | start query from after this voyage bookmark id (optional)
    customerId := "customerId_example" // string | get last prediction result made before this timestamp, must be in ISO8601 format with timezone, defaults to now (optional)
    carrierNo := "carrierNo_example" // string | filter on the carrier scac code (optional)
    blNo := "blNo_example" // string | filter on the imo number of the vessel (optional)
    cntrNo := "cntrNo_example" // string | filter on the target port by port code (optional)
    status := "status_example" // string | filter on the target port eta (optional)
    statusCode := "statusCode_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TrackingApi.EtaTrackingBillOfLadingBookmarksGet(context.Background()).StartAfter(startAfter).CustomerId(customerId).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).Status(status).StatusCode(statusCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrackingApi.EtaTrackingBillOfLadingBookmarksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EtaTrackingBillOfLadingBookmarksGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TrackingApi.EtaTrackingBillOfLadingBookmarksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEtaTrackingBillOfLadingBookmarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

[authorization](../README.md#authorization), [x-api-key](../README.md#x-api-key), [x-org-api-key](../README.md#x-org-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

