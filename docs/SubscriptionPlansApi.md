# \SubscriptionPlansApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionPlansGetSubscriptionPlan**](SubscriptionPlansApi.md#SubscriptionPlansGetSubscriptionPlan) | **Get** /api/manage/v1/SubscriptionPlans/{id} | Returns a subscription plan. Not all subscriptions can be issued for customer.
[**SubscriptionPlansGetSubscriptionPlans**](SubscriptionPlansApi.md#SubscriptionPlansGetSubscriptionPlans) | **Get** /api/manage/v1/SubscriptionPlans | Returns a list of active subscription plans that can be issued to the user.



## SubscriptionPlansGetSubscriptionPlan

> SubscriptionPlanVM SubscriptionPlansGetSubscriptionPlan(ctx, id).Execute()

Returns a subscription plan. Not all subscriptions can be issued for customer.

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
    id := "id_example" // string | Identifier of subsctiption plan

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionPlansApi.SubscriptionPlansGetSubscriptionPlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPlansApi.SubscriptionPlansGetSubscriptionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionPlansGetSubscriptionPlan`: SubscriptionPlanVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionPlansApi.SubscriptionPlansGetSubscriptionPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subsctiption plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPlansGetSubscriptionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionPlanVM**](SubscriptionPlanVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPlansGetSubscriptionPlans

> SubscriptionPlansVM SubscriptionPlansGetSubscriptionPlans(ctx).Skip(skip).Take(take).Execute()

Returns a list of active subscription plans that can be issued to the user.



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
    skip := int32(56) // int32 | Variable for pagination, defautl value is 0 (optional) (default to 0)
    take := int32(56) // int32 | Variable for pagination, default value is 10 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionPlansApi.SubscriptionPlansGetSubscriptionPlans(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionPlansApi.SubscriptionPlansGetSubscriptionPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionPlansGetSubscriptionPlans`: SubscriptionPlansVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionPlansApi.SubscriptionPlansGetSubscriptionPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPlansGetSubscriptionPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]

### Return type

[**SubscriptionPlansVM**](SubscriptionPlansVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

