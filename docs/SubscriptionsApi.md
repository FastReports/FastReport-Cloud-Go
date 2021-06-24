# \SubscriptionsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsGetDefaultPermissions**](SubscriptionsApi.md#SubscriptionsGetDefaultPermissions) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/defaultPermissions | Get subscription&#39;s default permissions for new entities
[**SubscriptionsGetPermissions**](SubscriptionsApi.md#SubscriptionsGetPermissions) | **Get** /api/manage/v1/Subscriptions/{id}/permissions | Get permissions for a subscription by id
[**SubscriptionsGetSubscription**](SubscriptionsApi.md#SubscriptionsGetSubscription) | **Get** /api/manage/v1/Subscriptions/{id} | Returns the subscription by id
[**SubscriptionsGetSubscriptions**](SubscriptionsApi.md#SubscriptionsGetSubscriptions) | **Get** /api/manage/v1/Subscriptions | Returns a list of all subscriptions of current user
[**SubscriptionsRenameSubscription**](SubscriptionsApi.md#SubscriptionsRenameSubscription) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/rename | Rename subscription
[**SubscriptionsUpdateDefaultPermissions**](SubscriptionsApi.md#SubscriptionsUpdateDefaultPermissions) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/defaultPermissions | Change subscription&#39;s default permissions for new entities
[**SubscriptionsUpdateLocale**](SubscriptionsApi.md#SubscriptionsUpdateLocale) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/Locale | Update subscription&#39;s default locale
[**SubscriptionsUpdatePermissions**](SubscriptionsApi.md#SubscriptionsUpdatePermissions) | **Post** /api/manage/v1/Subscriptions/{id}/permissions | Update permissions



## SubscriptionsGetDefaultPermissions

> DefaultPermissions SubscriptionsGetDefaultPermissions(ctx, subscriptionId).Execute()

Get subscription's default permissions for new entities

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
    subscriptionId := "subscriptionId_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsGetDefaultPermissions(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsGetDefaultPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGetDefaultPermissions`: DefaultPermissions
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsGetDefaultPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetDefaultPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefaultPermissions**](DefaultPermissions.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetPermissions

> SubscriptionPermissionsVM SubscriptionsGetPermissions(ctx, id).Execute()

Get permissions for a subscription by id

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
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGetPermissions`: SubscriptionPermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionPermissionsVM**](SubscriptionPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetSubscription

> SubscriptionVM SubscriptionsGetSubscription(ctx, id).Execute()

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
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsGetSubscription(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsGetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGetSubscription`: SubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsGetSubscriptions

> SubscriptionsVM SubscriptionsGetSubscriptions(ctx).Skip(skip).Take(take).Execute()

Returns a list of all subscriptions of current user

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
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsGetSubscriptions(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsGetSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsGetSubscriptions`: SubscriptionsVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsGetSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsGetSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]

### Return type

[**SubscriptionsVM**](SubscriptionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsRenameSubscription

> SubscriptionVM SubscriptionsRenameSubscription(ctx, subscriptionId).RenameSubscriptionVM(renameSubscriptionVM).Execute()

Rename subscription

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
    subscriptionId := "subscriptionId_example" // string | id
    renameSubscriptionVM := *openapiclient.NewRenameSubscriptionVM("Name_example") // RenameSubscriptionVM | rename VM

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsRenameSubscription(context.Background(), subscriptionId).RenameSubscriptionVM(renameSubscriptionVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsRenameSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsRenameSubscription`: SubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsRenameSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsRenameSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameSubscriptionVM** | [**RenameSubscriptionVM**](RenameSubscriptionVM.md) | rename VM | 

### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdateDefaultPermissions

> DefaultPermissionsVM SubscriptionsUpdateDefaultPermissions(ctx, subscriptionId).PermissionsVM(permissionsVM).Execute()

Change subscription's default permissions for new entities

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
    subscriptionId := "subscriptionId_example" // string | id
    permissionsVM := *openapiclient.NewUpdateDefaultPermissionsVM() // UpdateDefaultPermissionsVM | update default permissions VM

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsUpdateDefaultPermissions(context.Background(), subscriptionId).PermissionsVM(permissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsUpdateDefaultPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsUpdateDefaultPermissions`: DefaultPermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsUpdateDefaultPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateDefaultPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionsVM** | [**UpdateDefaultPermissionsVM**](UpdateDefaultPermissionsVM.md) | update default permissions VM | 

### Return type

[**DefaultPermissionsVM**](DefaultPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdateLocale

> SubscriptionVM SubscriptionsUpdateLocale(ctx, subscriptionId).UpdateModel(updateModel).Execute()

Update subscription's default locale

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
    subscriptionId := "subscriptionId_example" // string | id
    updateModel := *openapiclient.NewUpdateSubscriptionLocaleVM("Locale_example") // UpdateSubscriptionLocaleVM | update VM

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsUpdateLocale(context.Background(), subscriptionId).UpdateModel(updateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsUpdateLocale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsUpdateLocale`: SubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.SubscriptionsUpdateLocale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdateLocaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateModel** | [**UpdateSubscriptionLocaleVM**](UpdateSubscriptionLocaleVM.md) | update VM | 

### Return type

[**SubscriptionVM**](SubscriptionVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsUpdatePermissions

> SubscriptionsUpdatePermissions(ctx, id).PermissionsVM(permissionsVM).Execute()

Update permissions

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
    permissionsVM := *openapiclient.NewUpdateSubscriptionPermissionsVM(*openapiclient.NewSubscriptionPermissions(), int32(123)) // UpdateSubscriptionPermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionsApi.SubscriptionsUpdatePermissions(context.Background(), id).PermissionsVM(permissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.SubscriptionsUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionsVM** | [**UpdateSubscriptionPermissionsVM**](UpdateSubscriptionPermissionsVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

