# \CustomersAPI

All URIs are relative to *https://api.portcast.io/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomer**](CustomersAPI.md#CreateCustomer) | **Post** /org/customers | Create a new customer
[**ListCustomers**](CustomersAPI.md#ListCustomers) | **Get** /org/customers | Get a list of customers



## CreateCustomer

> Customer CreateCustomer(ctx).CustomerRequest(customerRequest).Execute()

Create a new customer

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
	customerRequest := *openapiclient.NewCustomerRequest("Customer_example") // CustomerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomersAPI.CreateCustomer(context.Background()).CustomerRequest(customerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.CreateCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomer`: Customer
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.CreateCustomer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerRequest** | [**CustomerRequest**](CustomerRequest.md) |  | 

### Return type

[**Customer**](Customer.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomers

> []Customer ListCustomers(ctx).Execute()

Get a list of customers

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
	resp, r, err := apiClient.CustomersAPI.ListCustomers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomersAPI.ListCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCustomers`: []Customer
	fmt.Fprintf(os.Stdout, "Response from `CustomersAPI.ListCustomers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


### Return type

[**[]Customer**](Customer.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

