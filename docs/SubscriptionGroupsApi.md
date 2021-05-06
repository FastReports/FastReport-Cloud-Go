# \SubscriptionGroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionGroupsGetGroupList**](SubscriptionGroupsApi.md#SubscriptionGroupsGetGroupList) | **Get** /api/manage/v1/Subscriptions/{id}/groups | returns list of groups in the subscription



## SubscriptionGroupsGetGroupList

> GroupsVM SubscriptionGroupsGetGroupList(ctx, id).Execute()

returns list of groups in the subscription

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
    id := "id_example" // string | subscripiton id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SubscriptionGroupsApi.SubscriptionGroupsGetGroupList(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsApi.SubscriptionGroupsGetGroupList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionGroupsGetGroupList`: GroupsVM
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsApi.SubscriptionGroupsGetGroupList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | subscripiton id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsGetGroupListRequest struct via the builder pattern


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

