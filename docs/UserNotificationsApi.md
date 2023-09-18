# \UserNotificationsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserNotificationsClearNotifications**](UserNotificationsApi.md#UserNotificationsClearNotifications) | **Delete** /api/manage/v1/notifications | Use this endpoint to \&quot;clear\&quot; your notifications
[**UserNotificationsGetNotifications**](UserNotificationsApi.md#UserNotificationsGetNotifications) | **Get** /api/manage/v1/notifications | Use this endpoint to recieve notifications



## UserNotificationsClearNotifications

> UserNotificationsClearNotifications(ctx).ClearNotificationsVM(clearNotificationsVM).Execute()

Use this endpoint to \"clear\" your notifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    clearNotificationsVM := *openapiclient.NewClearNotificationsVM() // ClearNotificationsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserNotificationsApi.UserNotificationsClearNotifications(context.Background()).ClearNotificationsVM(clearNotificationsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsApi.UserNotificationsClearNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserNotificationsClearNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clearNotificationsVM** | [**ClearNotificationsVM**](ClearNotificationsVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserNotificationsGetNotifications

> AuditActionsVM UserNotificationsGetNotifications(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

Use this endpoint to recieve notifications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    skip := int32(56) // int32 |  (optional) (default to 0)
    take := int32(56) // int32 |  (optional) (default to 5)
    subscriptionId := "subscriptionId_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserNotificationsApi.UserNotificationsGetNotifications(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotificationsApi.UserNotificationsGetNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserNotificationsGetNotifications`: AuditActionsVM
    fmt.Fprintf(os.Stdout, "Response from `UserNotificationsApi.UserNotificationsGetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserNotificationsGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | [default to 0]
 **take** | **int32** |  | [default to 5]
 **subscriptionId** | **string** |  | [default to &quot;&quot;]

### Return type

[**AuditActionsVM**](AuditActionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

