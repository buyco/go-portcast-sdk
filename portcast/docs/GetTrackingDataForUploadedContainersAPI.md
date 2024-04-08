# \GetTrackingDataForUploadedContainersAPI

All URIs are relative to *https://api.portcast.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTrackingBillOfLadingBookmarksBookmarkId**](GetTrackingDataForUploadedContainersAPI.md#DeleteTrackingBillOfLadingBookmarksBookmarkId) | **Delete** /eta/bill-of-lading-bookmarks/{bookmark_id} | Archive Bookmark ID
[**GetApiV1EtaBillOfLadingBookmarks**](GetTrackingDataForUploadedContainersAPI.md#GetApiV1EtaBillOfLadingBookmarks) | **Get** /eta/tracking/bill-of-lading-bookmarks/{bookmark_id} | Fetch tracking data by Bookmark ID
[**ListApiV1EtaBillOfLadingBookmarks**](GetTrackingDataForUploadedContainersAPI.md#ListApiV1EtaBillOfLadingBookmarks) | **Get** /eta/tracking/bill-of-lading-bookmarks | List Tracking data by Container, Bill of Lading or Carrier



## DeleteTrackingBillOfLadingBookmarksBookmarkId

> DeleteTrackingBillOfLadingBookmarksBookmarkId200Response DeleteTrackingBillOfLadingBookmarksBookmarkId(ctx, bookmarkId).XCustomer(xCustomer).Body(body).Execute()

Archive Bookmark ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/buyco/go-portcast-sdk/portcast"
)

func main() {
	bookmarkId := "4cbcd8ba-584e-4f76-aca7-ff388dc3e57b" // string | Bill of Lading Bookmark ID
	xCustomer := "Customer-ABC" // string | [Enterprise Customers] Customer Name to indicate which org the bookmark belongs to (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetTrackingDataForUploadedContainersAPI.DeleteTrackingBillOfLadingBookmarksBookmarkId(context.Background(), bookmarkId).XCustomer(xCustomer).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetTrackingDataForUploadedContainersAPI.DeleteTrackingBillOfLadingBookmarksBookmarkId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTrackingBillOfLadingBookmarksBookmarkId`: DeleteTrackingBillOfLadingBookmarksBookmarkId200Response
	fmt.Fprintf(os.Stdout, "Response from `GetTrackingDataForUploadedContainersAPI.DeleteTrackingBillOfLadingBookmarksBookmarkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** | Bill of Lading Bookmark ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTrackingBillOfLadingBookmarksBookmarkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCustomer** | **string** | [Enterprise Customers] Customer Name to indicate which org the bookmark belongs to | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**DeleteTrackingBillOfLadingBookmarksBookmarkId200Response**](DeleteTrackingBillOfLadingBookmarksBookmarkId200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiV1EtaBillOfLadingBookmarks

> GetApiV1EtaBillOfLadingBookmarks200Response GetApiV1EtaBillOfLadingBookmarks(ctx, bookmarkId).XCustomer(xCustomer).Body(body).Execute()

Fetch tracking data by Bookmark ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/buyco/go-portcast-sdk/portcast"
)

func main() {
	bookmarkId := "4cbcd8ba-584e-4f76-aca7-ff388dc3e57b" // string | Bill of Lading Bookmark ID
	xCustomer := "customer-ABC" // string | [Enterprise Customers] Customer Name to indicate which org to get tracking data from (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetTrackingDataForUploadedContainersAPI.GetApiV1EtaBillOfLadingBookmarks(context.Background(), bookmarkId).XCustomer(xCustomer).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetTrackingDataForUploadedContainersAPI.GetApiV1EtaBillOfLadingBookmarks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1EtaBillOfLadingBookmarks`: GetApiV1EtaBillOfLadingBookmarks200Response
	fmt.Fprintf(os.Stdout, "Response from `GetTrackingDataForUploadedContainersAPI.GetApiV1EtaBillOfLadingBookmarks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** | Bill of Lading Bookmark ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1EtaBillOfLadingBookmarksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCustomer** | **string** | [Enterprise Customers] Customer Name to indicate which org to get tracking data from | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetApiV1EtaBillOfLadingBookmarks200Response**](GetApiV1EtaBillOfLadingBookmarks200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiV1EtaBillOfLadingBookmarks

> ListApiV1EtaBillOfLadingBookmarks200Response ListApiV1EtaBillOfLadingBookmarks(ctx).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).StartAfter(startAfter).XCustomer(xCustomer).Body(body).Execute()

List Tracking data by Container, Bill of Lading or Carrier



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/buyco/go-portcast-sdk/portcast"
)

func main() {
	carrierNo := "CMDU" // string | Carrier SCAC Code (optional)
	blNo := "CNDN166131" // string | Bill of Lading or Booking Number (optional)
	cntrNo := "SCTU9017654" // string | Container Number (optional)
	startAfter := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Pagination Token (optional)
	xCustomer := "customer-ABC" // string | [Enterprise Customers] Customer Name to indicate which org to get tracking data from (optional)
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GetTrackingDataForUploadedContainersAPI.ListApiV1EtaBillOfLadingBookmarks(context.Background()).CarrierNo(carrierNo).BlNo(blNo).CntrNo(cntrNo).StartAfter(startAfter).XCustomer(xCustomer).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GetTrackingDataForUploadedContainersAPI.ListApiV1EtaBillOfLadingBookmarks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApiV1EtaBillOfLadingBookmarks`: ListApiV1EtaBillOfLadingBookmarks200Response
	fmt.Fprintf(os.Stdout, "Response from `GetTrackingDataForUploadedContainersAPI.ListApiV1EtaBillOfLadingBookmarks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiV1EtaBillOfLadingBookmarksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierNo** | **string** | Carrier SCAC Code | 
 **blNo** | **string** | Bill of Lading or Booking Number | 
 **cntrNo** | **string** | Container Number | 
 **startAfter** | **string** | Pagination Token | 
 **xCustomer** | **string** | [Enterprise Customers] Customer Name to indicate which org to get tracking data from | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListApiV1EtaBillOfLadingBookmarks200Response**](ListApiV1EtaBillOfLadingBookmarks200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

