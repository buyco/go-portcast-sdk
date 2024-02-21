# \UploadContainersForTrackingAPI

All URIs are relative to *https://api.portcast.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostBooking**](UploadContainersForTrackingAPI.md#PostBooking) | **Post** /eta/booking | Upload using Booking/Bill of Lading
[**PostEtaBillOfLadingBookmarks**](UploadContainersForTrackingAPI.md#PostEtaBillOfLadingBookmarks) | **Post** /eta/bill-of-lading-bookmarks | Upload using Container Number



## PostBooking

> BookingAPI PostBooking(ctx).XCustomer(xCustomer).PostBookingRequest(postBookingRequest).Execute()

Upload using Booking/Bill of Lading



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
	xCustomer := "customer-ABC" // string | [Enterprise Customers] Customer Name to indicate which org to upload too (optional)
	postBookingRequest := *openapiclient.NewPostBookingRequest("CarrierNo_example", "DocNo_example", "DocType_example") // PostBookingRequest | Upload Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadContainersForTrackingAPI.PostBooking(context.Background()).XCustomer(xCustomer).PostBookingRequest(postBookingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadContainersForTrackingAPI.PostBooking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBooking`: BookingAPI
	fmt.Fprintf(os.Stdout, "Response from `UploadContainersForTrackingAPI.PostBooking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostBookingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCustomer** | **string** | [Enterprise Customers] Customer Name to indicate which org to upload too | 
 **postBookingRequest** | [**PostBookingRequest**](PostBookingRequest.md) | Upload Parameters | 

### Return type

[**BookingAPI**](BookingAPI.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEtaBillOfLadingBookmarks

> PostEtaBillOfLadingBookmarks200Response PostEtaBillOfLadingBookmarks(ctx).XCustomer(xCustomer).PostEtaBillOfLadingBookmarksRequest(postEtaBillOfLadingBookmarksRequest).Execute()

Upload using Container Number



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
	xCustomer := "customer-ABC" // string | [Enterprise Customers] Customer Name to indicate which org to upload too (optional)
	postEtaBillOfLadingBookmarksRequest := *openapiclient.NewPostEtaBillOfLadingBookmarksRequest("CarrierNo_example", "CntrNo_example") // PostEtaBillOfLadingBookmarksRequest | Upload Parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadContainersForTrackingAPI.PostEtaBillOfLadingBookmarks(context.Background()).XCustomer(xCustomer).PostEtaBillOfLadingBookmarksRequest(postEtaBillOfLadingBookmarksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadContainersForTrackingAPI.PostEtaBillOfLadingBookmarks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEtaBillOfLadingBookmarks`: PostEtaBillOfLadingBookmarks200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadContainersForTrackingAPI.PostEtaBillOfLadingBookmarks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEtaBillOfLadingBookmarksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCustomer** | **string** | [Enterprise Customers] Customer Name to indicate which org to upload too | 
 **postEtaBillOfLadingBookmarksRequest** | [**PostEtaBillOfLadingBookmarksRequest**](PostEtaBillOfLadingBookmarksRequest.md) | Upload Parameters | 

### Return type

[**PostEtaBillOfLadingBookmarks200Response**](PostEtaBillOfLadingBookmarks200Response.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

