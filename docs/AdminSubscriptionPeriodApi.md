# \AdminSubscriptionPeriodApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubscriptionPeriodRenewSubscription**](AdminSubscriptionPeriodApi.md#AdminSubscriptionPeriodRenewSubscription) | **Post** /api/admin/v1/Subscriptions/{id}/Renew | Create a new subscripiton period, move current period to old



## AdminSubscriptionPeriodRenewSubscription

> SubscriptionVM AdminSubscriptionPeriodRenewSubscription(ctx, id).ViewModel(viewModel).Execute()

Create a new subscripiton period, move current period to old

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
    id := "id_example" // string | 
    viewModel := *openapiclient.NewCreateSubscriptionPeriodVM() // CreateSubscriptionPeriodVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionPeriodApi.AdminSubscriptionPeriodRenewSubscription(context.Background(), id).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPeriodApi.AdminSubscriptionPeriodRenewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionPeriodRenewSubscription`: SubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionPeriodApi.AdminSubscriptionPeriodRenewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionPeriodRenewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewModel** | [**CreateSubscriptionPeriodVM**](CreateSubscriptionPeriodVM.md) |  | 

### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

