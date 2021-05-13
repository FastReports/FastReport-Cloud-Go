# \SubscriptionGroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionGroupsGetGroupsList**](SubscriptionGroupsApi.md#SubscriptionGroupsGetGroupsList) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/groups | returns groups of the subscription or subscription user



## SubscriptionGroupsGetGroupsList

> GroupsVM SubscriptionGroupsGetGroupsList(ctx, subscriptionId).UserId(userId).Execute()

returns groups of the subscription or subscription user

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
    subscriptionId := "subscriptionId_example" // string | subscripiton id
    userId := "userId_example" // string | user Id (optional) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionGroupsApi.SubscriptionGroupsGetGroupsList(context.Background(), subscriptionId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsApi.SubscriptionGroupsGetGroupsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionGroupsGetGroupsList`: GroupsVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsApi.SubscriptionGroupsGetGroupsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscripiton id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsGetGroupsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | user Id (optional) | 

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

