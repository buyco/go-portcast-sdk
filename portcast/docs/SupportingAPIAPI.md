# \SupportingAPIAPI

All URIs are relative to *https://api.portcast.io/api/v2/eta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScac**](SupportingAPIAPI.md#GetScac) | **Get** /scac | Supported Carrier SCAC List



## GetScac

> []GetScac200ResponseInner GetScac(ctx).Execute()

Supported Carrier SCAC List



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SupportingAPIAPI.GetScac(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportingAPIAPI.GetScac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScac`: []GetScac200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `SupportingAPIAPI.GetScac`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetScacRequest struct via the builder pattern


### Return type

[**[]GetScac200ResponseInner**](GetScac200ResponseInner.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

