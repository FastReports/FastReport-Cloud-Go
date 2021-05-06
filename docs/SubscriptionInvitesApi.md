# \SubscriptionInvitesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionInvitesAcceptInvite**](SubscriptionInvitesApi.md#SubscriptionInvitesAcceptInvite) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/invite/{accessToken}/accept | Add a user to the subscription using invite,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
[**SubscriptionInvitesCreateInvite**](SubscriptionInvitesApi.md#SubscriptionInvitesCreateInvite) | **Post** /api/manage/v1/Subscriptions/{subscriptionId}/invite | Create invite to subscription
[**SubscriptionInvitesDeleteInvite**](SubscriptionInvitesApi.md#SubscriptionInvitesDeleteInvite) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/invite/{accesstoken} | Rename subscription
[**SubscriptionInvitesGetInvites**](SubscriptionInvitesApi.md#SubscriptionInvitesGetInvites) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/invites | Get list of invites in a subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.



## SubscriptionInvitesAcceptInvite

> SubscriptionInvitesAcceptInvite(ctx, subscriptionId, accessToken).Execute()

Add a user to the subscription using invite,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.

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
    accessToken := "accessToken_example" // string | access token of the subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionInvitesApi.SubscriptionInvitesAcceptInvite(context.Background(), subscriptionId, accessToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionInvitesApi.SubscriptionInvitesAcceptInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 
**accessToken** | **string** | access token of the subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionInvitesAcceptInviteRequest struct via the builder pattern


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


## SubscriptionInvitesCreateInvite

> SubscriptionVM SubscriptionInvitesCreateInvite(ctx, subscriptionId).CreateInviteVM(createInviteVM).Execute()

Create invite to subscription

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
    createInviteVM := *openapiclient.NewCreateSubscriptionInviteVM() // CreateSubscriptionInviteVM | create VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionInvitesApi.SubscriptionInvitesCreateInvite(context.Background(), subscriptionId).CreateInviteVM(createInviteVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionInvitesApi.SubscriptionInvitesCreateInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionInvitesCreateInvite`: SubscriptionVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionInvitesApi.SubscriptionInvitesCreateInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionInvitesCreateInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInviteVM** | [**CreateSubscriptionInviteVM**](CreateSubscriptionInviteVM.md) | create VM | 

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


## SubscriptionInvitesDeleteInvite

> SubscriptionInvitesDeleteInvite(ctx, subscriptionId, accesstoken).Execute()

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
    accesstoken := "accesstoken_example" // string | invite's token

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionInvitesApi.SubscriptionInvitesDeleteInvite(context.Background(), subscriptionId, accesstoken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionInvitesApi.SubscriptionInvitesDeleteInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 
**accesstoken** | **string** | invite&#39;s token | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionInvitesDeleteInviteRequest struct via the builder pattern


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


## SubscriptionInvitesGetInvites

> SubscriptionInvitesVM SubscriptionInvitesGetInvites(ctx, subscriptionId).Execute()

Get list of invites in a subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.

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
    resp, r, err := api_client.SubscriptionInvitesApi.SubscriptionInvitesGetInvites(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionInvitesApi.SubscriptionInvitesGetInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionInvitesGetInvites`: SubscriptionInvitesVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionInvitesApi.SubscriptionInvitesGetInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionInvitesGetInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionInvitesVM**](SubscriptionInvitesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

