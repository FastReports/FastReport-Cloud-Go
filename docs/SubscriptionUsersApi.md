# \SubscriptionUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionUsersAddUser**](SubscriptionUsersApi.md#SubscriptionUsersAddUser) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/users/{userId} | Add a user to the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
[**SubscriptionUsersGetUserGroups**](SubscriptionUsersApi.md#SubscriptionUsersGetUserGroups) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/user/{userId}/groups | Returns all users of subscription
[**SubscriptionUsersGetUsers**](SubscriptionUsersApi.md#SubscriptionUsersGetUsers) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/users | Returns all users of subscription
[**SubscriptionUsersLeaveSubscripiton**](SubscriptionUsersApi.md#SubscriptionUsersLeaveSubscripiton) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/leave | Allows user to leave subscription,.
[**SubscriptionUsersRemoveUser**](SubscriptionUsersApi.md#SubscriptionUsersRemoveUser) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/users/{userId} | Delete a user from the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.



## SubscriptionUsersAddUser

> SubscriptionUsersAddUser(ctx, subscriptionId, userId).Execute()

Add a user to the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.

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
    subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
    userId := "userId_example" // string | Idenitifier of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionUsersApi.SubscriptionUsersAddUser(context.Background(), subscriptionId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersApi.SubscriptionUsersAddUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 
**userId** | **string** | Idenitifier of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUsersAddUserRequest struct via the builder pattern


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


## SubscriptionUsersGetUserGroups

> GroupsVM SubscriptionUsersGetUserGroups(ctx, subscriptionId, userId).Execute()

Returns all users of subscription

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
    subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
    userId := "userId_example" // string | user id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionUsersApi.SubscriptionUsersGetUserGroups(context.Background(), subscriptionId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersApi.SubscriptionUsersGetUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionUsersGetUserGroups`: GroupsVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionUsersApi.SubscriptionUsersGetUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 
**userId** | **string** | user id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUsersGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupsVM**](GroupsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUsersGetUsers

> SubscriptionUsersVM SubscriptionUsersGetUsers(ctx, subscriptionId).Skip(skip).Take(take).Execute()

Returns all users of subscription

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
    subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
    skip := int32(56) // int32 | How many entities skip (optional) (default to 0)
    take := int32(56) // int32 | How many entities take (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionUsersApi.SubscriptionUsersGetUsers(context.Background(), subscriptionId).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersApi.SubscriptionUsersGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionUsersGetUsers`: SubscriptionUsersVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionUsersApi.SubscriptionUsersGetUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUsersGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | How many entities skip | [default to 0]
 **take** | **int32** | How many entities take | [default to 10]

### Return type

[**SubscriptionUsersVM**](SubscriptionUsersVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUsersLeaveSubscripiton

> SubscriptionUsersLeaveSubscripiton(ctx, subscriptionId).Execute()

Allows user to leave subscription,.

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
    subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionUsersApi.SubscriptionUsersLeaveSubscripiton(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersApi.SubscriptionUsersLeaveSubscripiton``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUsersLeaveSubscripitonRequest struct via the builder pattern


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


## SubscriptionUsersRemoveUser

> SubscriptionUsersRemoveUser(ctx, subscriptionId, userId).Execute()

Delete a user from the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.

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
    subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
    userId := "userId_example" // string | Idenitifier of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionUsersApi.SubscriptionUsersRemoveUser(context.Background(), subscriptionId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersApi.SubscriptionUsersRemoveUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 
**userId** | **string** | Idenitifier of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUsersRemoveUserRequest struct via the builder pattern


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

