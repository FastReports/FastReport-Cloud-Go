# \AdminSubscriptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubscriptionsCreateSubscription**](AdminSubscriptionsApi.md#AdminSubscriptionsCreateSubscription) | **Post** /api/admin/v1/Subscriptions | Create a new subscription based on some plan
[**AdminSubscriptionsDeleteSubscription**](AdminSubscriptionsApi.md#AdminSubscriptionsDeleteSubscription) | **Delete** /api/admin/v1/Subscriptions/{id} | Delete the subscription by id
[**AdminSubscriptionsGetNewSibscriptionsPerMonth**](AdminSubscriptionsApi.md#AdminSubscriptionsGetNewSibscriptionsPerMonth) | **Get** /api/admin/v1/Subscriptions/stat/new/{from}/{to} | Returns a key-value pair of new(renew) subscriptions count per month for a specified time span: (month, number of new subscriptions)
[**AdminSubscriptionsGetPermissions**](AdminSubscriptionsApi.md#AdminSubscriptionsGetPermissions) | **Get** /api/admin/v1/Subscriptions/{id}/permissions | Get all subscription permissions
[**AdminSubscriptionsGetSubscription**](AdminSubscriptionsApi.md#AdminSubscriptionsGetSubscription) | **Get** /api/admin/v1/Subscriptions/{id} | Returns the subscription by id
[**AdminSubscriptionsGetSubscriptions**](AdminSubscriptionsApi.md#AdminSubscriptionsGetSubscriptions) | **Get** /api/admin/v1/Subscriptions | Returns a list of all subscriptions
[**AdminSubscriptionsReCountSubscription**](AdminSubscriptionsApi.md#AdminSubscriptionsReCountSubscription) | **Get** /api/admin/v1/Subscriptions/{id}/recount | Recount subscription&#39;s files and folders sizes.
[**AdminSubscriptionsUpdatePermissions**](AdminSubscriptionsApi.md#AdminSubscriptionsUpdatePermissions) | **Post** /api/admin/v1/Subscriptions/{id}/permissions | Update permissions to subscription
[**AdminSubscriptionsUpdateSubscription**](AdminSubscriptionsApi.md#AdminSubscriptionsUpdateSubscription) | **Put** /api/admin/v1/Subscriptions/{id} | Update the subscription by id



## AdminSubscriptionsCreateSubscription

> AdminSubscriptionVM AdminSubscriptionsCreateSubscription(ctx).ViewModel(viewModel).Execute()

Create a new subscription based on some plan

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
    viewModel := *openapiclient.NewCreateSubscriptionVM() // CreateSubscriptionVM | View model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsCreateSubscription(context.Background()).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsCreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionsCreateSubscription`: AdminSubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionsApi.AdminSubscriptionsCreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewModel** | [**CreateSubscriptionVM**](CreateSubscriptionVM.md) | View model | 

### Return type

[**AdminSubscriptionVM**](AdminSubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionsDeleteSubscription

> AdminSubscriptionsDeleteSubscription(ctx, id).Execute()

Delete the subscription by id

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
    id := "id_example" // string | Identifier of subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsDeleteSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsDeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsDeleteSubscriptionRequest struct via the builder pattern


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


## AdminSubscriptionsGetNewSibscriptionsPerMonth

> map[string]int32 AdminSubscriptionsGetNewSibscriptionsPerMonth(ctx, from, to).Execute()

Returns a key-value pair of new(renew) subscriptions count per month for a specified time span: (month, number of new subscriptions)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    from := time.Now() // time.Time | A starting date for stats calculation
    to := time.Now() // time.Time | An ending date for stats calculation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsGetNewSibscriptionsPerMonth(context.Background(), from, to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsGetNewSibscriptionsPerMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionsGetNewSibscriptionsPerMonth`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionsApi.AdminSubscriptionsGetNewSibscriptionsPerMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **time.Time** | A starting date for stats calculation | 
**to** | **time.Time** | An ending date for stats calculation | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsGetNewSibscriptionsPerMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]int32**

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionsGetPermissions

> SubscriptionPermissionsVM AdminSubscriptionsGetPermissions(ctx, id).Execute()

Get all subscription permissions

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionsGetPermissions`: SubscriptionPermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionsApi.AdminSubscriptionsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionPermissionsVM**](SubscriptionPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionsGetSubscription

> AdminSubscriptionVM AdminSubscriptionsGetSubscription(ctx, id).Execute()

Returns the subscription by id

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
    id := "id_example" // string | Identifier of subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsGetSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsGetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionsGetSubscription`: AdminSubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionsApi.AdminSubscriptionsGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdminSubscriptionVM**](AdminSubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionsGetSubscriptions

> AdminSubscriptionsVM AdminSubscriptionsGetSubscriptions(ctx).Skip(skip).Take(take).Execute()

Returns a list of all subscriptions

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
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsGetSubscriptions(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsGetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionsGetSubscriptions`: AdminSubscriptionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionsApi.AdminSubscriptionsGetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]

### Return type

[**AdminSubscriptionsVM**](AdminSubscriptionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionsReCountSubscription

> AdminSubscriptionsReCountSubscription(ctx, id).Execute()

Recount subscription's files and folders sizes.

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
    id := "id_example" // string | Identifier of subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsReCountSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsReCountSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsReCountSubscriptionRequest struct via the builder pattern


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


## AdminSubscriptionsUpdatePermissions

> AdminSubscriptionsUpdatePermissions(ctx, id).PermissionsVM(permissionsVM).Execute()

Update permissions to subscription

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
    id := "id_example" // string | subscription id
    permissionsVM := *openapiclient.NewUpdateSubscriptionPermissionsVM(*openapiclient.NewSubscriptionPermissions(), int32(123)) // UpdateSubscriptionPermissionsVM | permissions VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsUpdatePermissions(context.Background(), id).PermissionsVM(permissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionsVM** | [**UpdateSubscriptionPermissionsVM**](UpdateSubscriptionPermissionsVM.md) | permissions VM | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionsUpdateSubscription

> AdminSubscriptionVM AdminSubscriptionsUpdateSubscription(ctx, id).ViewModel(viewModel).Execute()

Update the subscription by id

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
    id := "id_example" // string | Identifier of subscription
    viewModel := *openapiclient.NewUpdateSubscriptionVM() // UpdateSubscriptionVM | View model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionsApi.AdminSubscriptionsUpdateSubscription(context.Background(), id).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionsApi.AdminSubscriptionsUpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionsUpdateSubscription`: AdminSubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionsApi.AdminSubscriptionsUpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionsUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewModel** | [**UpdateSubscriptionVM**](UpdateSubscriptionVM.md) | View model | 

### Return type

[**AdminSubscriptionVM**](AdminSubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

