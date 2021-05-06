# \AdminSubscriptionPlansApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubscriptionPlansCreateSubscriptionPlan**](AdminSubscriptionPlansApi.md#AdminSubscriptionPlansCreateSubscriptionPlan) | **Post** /api/admin/v1/SubscriptionPlans | Create a new subscription plan, returns a new model
[**AdminSubscriptionPlansDeleteSubscriptionPlan**](AdminSubscriptionPlansApi.md#AdminSubscriptionPlansDeleteSubscriptionPlan) | **Delete** /api/admin/v1/SubscriptionPlans/{id} | Delete a subscription plan.
[**AdminSubscriptionPlansGetSubscriptionPlan**](AdminSubscriptionPlansApi.md#AdminSubscriptionPlansGetSubscriptionPlan) | **Get** /api/admin/v1/SubscriptionPlans/{id} | Returns a subscription plan. Not all subscriptions can be issued for customer.
[**AdminSubscriptionPlansGetSubscriptionPlans**](AdminSubscriptionPlansApi.md#AdminSubscriptionPlansGetSubscriptionPlans) | **Get** /api/admin/v1/SubscriptionPlans | Returns a list of active subscription plans that can be issued to the user.
[**AdminSubscriptionPlansUpdateSubscriptionPlan**](AdminSubscriptionPlansApi.md#AdminSubscriptionPlansUpdateSubscriptionPlan) | **Put** /api/admin/v1/SubscriptionPlans/{id} | Update a subscription plan.



## AdminSubscriptionPlansCreateSubscriptionPlan

> SubscriptionPlanVM AdminSubscriptionPlansCreateSubscriptionPlan(ctx).ViewModel(viewModel).Execute()

Create a new subscription plan, returns a new model

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
    viewModel := *openapiclient.NewCreateSubscriptionPlanVM() // CreateSubscriptionPlanVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionPlansApi.AdminSubscriptionPlansCreateSubscriptionPlan(context.Background()).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPlansApi.AdminSubscriptionPlansCreateSubscriptionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionPlansCreateSubscriptionPlan`: SubscriptionPlanVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionPlansApi.AdminSubscriptionPlansCreateSubscriptionPlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionPlansCreateSubscriptionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewModel** | [**CreateSubscriptionPlanVM**](CreateSubscriptionPlanVM.md) |  | 

### Return type

[**SubscriptionPlanVM**](SubscriptionPlanVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionPlansDeleteSubscriptionPlan

> AdminSubscriptionPlansDeleteSubscriptionPlan(ctx, id).Execute()

Delete a subscription plan.

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
    resp, r, err := api_client.AdminSubscriptionPlansApi.AdminSubscriptionPlansDeleteSubscriptionPlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPlansApi.AdminSubscriptionPlansDeleteSubscriptionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subsctiption plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionPlansDeleteSubscriptionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionPlansGetSubscriptionPlan

> SubscriptionPlanVM AdminSubscriptionPlansGetSubscriptionPlan(ctx, id).Execute()

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
    resp, r, err := api_client.AdminSubscriptionPlansApi.AdminSubscriptionPlansGetSubscriptionPlan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPlansApi.AdminSubscriptionPlansGetSubscriptionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionPlansGetSubscriptionPlan`: SubscriptionPlanVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionPlansApi.AdminSubscriptionPlansGetSubscriptionPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subsctiption plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionPlansGetSubscriptionPlanRequest struct via the builder pattern


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


## AdminSubscriptionPlansGetSubscriptionPlans

> SubscriptionPlansVM AdminSubscriptionPlansGetSubscriptionPlans(ctx).Skip(skip).Take(take).Execute()

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
    resp, r, err := api_client.AdminSubscriptionPlansApi.AdminSubscriptionPlansGetSubscriptionPlans(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPlansApi.AdminSubscriptionPlansGetSubscriptionPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionPlansGetSubscriptionPlans`: SubscriptionPlansVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionPlansApi.AdminSubscriptionPlansGetSubscriptionPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionPlansGetSubscriptionPlansRequest struct via the builder pattern


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


## AdminSubscriptionPlansUpdateSubscriptionPlan

> SubscriptionPlanVM AdminSubscriptionPlansUpdateSubscriptionPlan(ctx, id).ViewModel(viewModel).Execute()

Update a subscription plan.

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
    viewModel := *openapiclient.NewUpdateSubscriptionPlanVM() // UpdateSubscriptionPlanVM | Update VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionPlansApi.AdminSubscriptionPlansUpdateSubscriptionPlan(context.Background(), id).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPlansApi.AdminSubscriptionPlansUpdateSubscriptionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionPlansUpdateSubscriptionPlan`: SubscriptionPlanVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionPlansApi.AdminSubscriptionPlansUpdateSubscriptionPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subsctiption plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionPlansUpdateSubscriptionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewModel** | [**UpdateSubscriptionPlanVM**](UpdateSubscriptionPlanVM.md) | Update VM | 

### Return type

[**SubscriptionPlanVM**](SubscriptionPlanVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

