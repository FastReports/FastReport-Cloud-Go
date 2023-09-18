# \UserProfileApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserProfileGetMyProfile**](UserProfileApi.md#UserProfileGetMyProfile) | **Get** /api/manage/v1/UserProfile | Return current profile of the current user
[**UserProfileGetUserProfile**](UserProfileApi.md#UserProfileGetUserProfile) | **Get** /api/manage/v1/UserProfile/{userId} | Return user profile by user identifier.  If the user did not provide information about himself or blocked, then the endpoint will return an empty model. (only id)
[**UserProfileUpdateMyProfile**](UserProfileApi.md#UserProfileUpdateMyProfile) | **Put** /api/manage/v1/UserProfile | Update profile of the current user



## UserProfileGetMyProfile

> UserProfileVM UserProfileGetMyProfile(ctx).Execute()

Return current profile of the current user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserProfileApi.UserProfileGetMyProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProfileApi.UserProfileGetMyProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserProfileGetMyProfile`: UserProfileVM
    fmt.Fprintf(os.Stdout, "Response from `UserProfileApi.UserProfileGetMyProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserProfileGetMyProfileRequest struct via the builder pattern


### Return type

[**UserProfileVM**](UserProfileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserProfileGetUserProfile

> UserProfileVM UserProfileGetUserProfile(ctx, userId).Execute()

Return user profile by user identifier.  If the user did not provide information about himself or blocked, then the endpoint will return an empty model. (only id)

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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserProfileApi.UserProfileGetUserProfile(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProfileApi.UserProfileGetUserProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserProfileGetUserProfile`: UserProfileVM
    fmt.Fprintf(os.Stdout, "Response from `UserProfileApi.UserProfileGetUserProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserProfileGetUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserProfileVM**](UserProfileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserProfileUpdateMyProfile

> UserProfileUpdateMyProfile(ctx).UpdateUserProfileVM(updateUserProfileVM).Execute()

Update profile of the current user



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
    updateUserProfileVM := *openapiclient.NewUpdateUserProfileVM() // UpdateUserProfileVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserProfileApi.UserProfileUpdateMyProfile(context.Background()).UpdateUserProfileVM(updateUserProfileVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserProfileApi.UserProfileUpdateMyProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserProfileUpdateMyProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserProfileVM** | [**UpdateUserProfileVM**](UpdateUserProfileVM.md) |  | 

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

