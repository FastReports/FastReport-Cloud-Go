# \SubscriptionGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionGroupsCountGroupsAsync**](SubscriptionGroupsAPI.md#SubscriptionGroupsCountGroupsAsync) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/count | Returns a number of groups in subscription
[**SubscriptionGroupsGetGroupsList**](SubscriptionGroupsAPI.md#SubscriptionGroupsGetGroupsList) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/groups | returns groups of the subscription or subscription user



## SubscriptionGroupsCountGroupsAsync

> int64 SubscriptionGroupsCountGroupsAsync(ctx, subscriptionId).Execute()

Returns a number of groups in subscription

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
	subscriptionId := "subscriptionId_example" // string | subscripiton id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsCountGroupsAsync(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsCountGroupsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsCountGroupsAsync`: int64
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsCountGroupsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscripiton id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupsCountGroupsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int64**

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | subscripiton id
	userId := "userId_example" // string | user Id (optional) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionGroupsAPI.SubscriptionGroupsGetGroupsList(context.Background(), subscriptionId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupsAPI.SubscriptionGroupsGetGroupsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionGroupsGetGroupsList`: GroupsVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupsAPI.SubscriptionGroupsGetGroupsList`: %v\n", resp)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

