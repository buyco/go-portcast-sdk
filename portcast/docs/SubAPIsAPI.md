# \SubAPIsAPI

All URIs are relative to *https://api.portcast.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiV1EtaBillOfLadingBookmarksTerminals**](SubAPIsAPI.md#GetApiV1EtaBillOfLadingBookmarksTerminals) | **Get** /eta/tracking/bill-of-lading-bookmarks/{bookmark_id}/import_export_plan | Fetch terminal data by Bookmark ID



## GetApiV1EtaBillOfLadingBookmarksTerminals

> PortTerminalAddOn GetApiV1EtaBillOfLadingBookmarksTerminals(ctx, bookmarkId).XCustomer(xCustomer).Body(body).Execute()

Fetch terminal data by Bookmark ID



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
	resp, r, err := apiClient.SubAPIsAPI.GetApiV1EtaBillOfLadingBookmarksTerminals(context.Background(), bookmarkId).XCustomer(xCustomer).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubAPIsAPI.GetApiV1EtaBillOfLadingBookmarksTerminals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiV1EtaBillOfLadingBookmarksTerminals`: PortTerminalAddOn
	fmt.Fprintf(os.Stdout, "Response from `SubAPIsAPI.GetApiV1EtaBillOfLadingBookmarksTerminals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bookmarkId** | **string** | Bill of Lading Bookmark ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiV1EtaBillOfLadingBookmarksTerminalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCustomer** | **string** | [Enterprise Customers] Customer Name to indicate which org to get tracking data from | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**PortTerminalAddOn**](PortTerminalAddOn.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

