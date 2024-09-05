# \SubscriptionUsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionUsersAddUser**](SubscriptionUsersAPI.md#SubscriptionUsersAddUser) | **Put** /api/manage/v1/Subscriptions/{subscriptionId}/users/{userId} | Add a user to the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.
[**SubscriptionUsersCountUsersAsync**](SubscriptionUsersAPI.md#SubscriptionUsersCountUsersAsync) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/UsersCount | Returns a number of users in subscription
[**SubscriptionUsersGetUsers**](SubscriptionUsersAPI.md#SubscriptionUsersGetUsers) | **Get** /api/manage/v1/Subscriptions/{subscriptionId}/users | Returns all users of subscription
[**SubscriptionUsersLeaveSubscripiton**](SubscriptionUsersAPI.md#SubscriptionUsersLeaveSubscripiton) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/leave | Allows user to leave subscription,.
[**SubscriptionUsersRemoveUser**](SubscriptionUsersAPI.md#SubscriptionUsersRemoveUser) | **Delete** /api/manage/v1/Subscriptions/{subscriptionId}/users/{userId} | Delete a user from the subscription,  the added users will be displayed in the list of users of the subscription,  and these users will also have an active subscription.



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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
	userId := "userId_example" // string | Idenitifier of user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionUsersAPI.SubscriptionUsersAddUser(context.Background(), subscriptionId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersAPI.SubscriptionUsersAddUser``: %v\n", err)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUsersCountUsersAsync

> int64 SubscriptionUsersCountUsersAsync(ctx, subscriptionId).Execute()

Returns a number of users in subscription

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
	subscriptionId := "subscriptionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUsersAPI.SubscriptionUsersCountUsersAsync(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersAPI.SubscriptionUsersCountUsersAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUsersCountUsersAsync`: int64
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUsersAPI.SubscriptionUsersCountUsersAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUsersCountUsersAsyncRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
	skip := int32(56) // int32 | How many entities skip (optional) (default to 0)
	take := int32(56) // int32 | How many entities take (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionUsersAPI.SubscriptionUsersGetUsers(context.Background(), subscriptionId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersAPI.SubscriptionUsersGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUsersGetUsers`: SubscriptionUsersVM
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionUsersAPI.SubscriptionUsersGetUsers`: %v\n", resp)
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
- **Accept**: application/json

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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionUsersAPI.SubscriptionUsersLeaveSubscripiton(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersAPI.SubscriptionUsersLeaveSubscripiton``: %v\n", err)
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
- **Accept**: application/json

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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string | Idenitifier of subscription
	userId := "userId_example" // string | Idenitifier of user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubscriptionUsersAPI.SubscriptionUsersRemoveUser(context.Background(), subscriptionId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionUsersAPI.SubscriptionUsersRemoveUser``: %v\n", err)
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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

