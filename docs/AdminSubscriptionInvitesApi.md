# \AdminSubscriptionInvitesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubscriptionInvitesCreateInvite**](AdminSubscriptionInvitesApi.md#AdminSubscriptionInvitesCreateInvite) | **Post** /api/admin/v1/Subscriptions/{subscriptionId}/invite | Create invite to subscription
[**AdminSubscriptionInvitesDeleteInvite**](AdminSubscriptionInvitesApi.md#AdminSubscriptionInvitesDeleteInvite) | **Delete** /api/admin/v1/Subscriptions/{subscriptionId}/invite/{accesstoken} | Rename subscription
[**AdminSubscriptionInvitesGetInvites**](AdminSubscriptionInvitesApi.md#AdminSubscriptionInvitesGetInvites) | **Get** /api/admin/v1/Subscriptions/{subscriptionId}/invites | Get list of invites in a subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.



## AdminSubscriptionInvitesCreateInvite

> SubscriptionInviteVM AdminSubscriptionInvitesCreateInvite(ctx, subscriptionId).CreateInviteVM(createInviteVM).Execute()

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
    resp, r, err := api_client.AdminSubscriptionInvitesApi.AdminSubscriptionInvitesCreateInvite(context.Background(), subscriptionId).CreateInviteVM(createInviteVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionInvitesApi.AdminSubscriptionInvitesCreateInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionInvitesCreateInvite`: SubscriptionInviteVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionInvitesApi.AdminSubscriptionInvitesCreateInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionInvitesCreateInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInviteVM** | [**CreateSubscriptionInviteVM**](CreateSubscriptionInviteVM.md) | create VM | 

### Return type

[**SubscriptionInviteVM**](SubscriptionInviteVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionInvitesDeleteInvite

> AdminSubscriptionInvitesDeleteInvite(ctx, subscriptionId, accesstoken).Execute()

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
    resp, r, err := api_client.AdminSubscriptionInvitesApi.AdminSubscriptionInvitesDeleteInvite(context.Background(), subscriptionId, accesstoken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionInvitesApi.AdminSubscriptionInvitesDeleteInvite``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminSubscriptionInvitesDeleteInviteRequest struct via the builder pattern


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


## AdminSubscriptionInvitesGetInvites

> SubscriptionInvitesVM AdminSubscriptionInvitesGetInvites(ctx, subscriptionId).Execute()

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
    resp, r, err := api_client.AdminSubscriptionInvitesApi.AdminSubscriptionInvitesGetInvites(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionInvitesApi.AdminSubscriptionInvitesGetInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionInvitesGetInvites`: SubscriptionInvitesVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionInvitesApi.AdminSubscriptionInvitesGetInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Idenitifier of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionInvitesGetInvitesRequest struct via the builder pattern


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

