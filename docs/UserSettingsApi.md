# \UserSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserSettingsGetCurrentUserSettings**](UserSettingsApi.md#UserSettingsGetCurrentUserSettings) | **Get** /api/manage/v1/UserSettings | Return current user settings.
[**UserSettingsUpdateMySettings**](UserSettingsApi.md#UserSettingsUpdateMySettings) | **Put** /api/manage/v1/UserSettings | Update settings of the current user



## UserSettingsGetCurrentUserSettings

> UserSettingsVM UserSettingsGetCurrentUserSettings(ctx).Execute()

Return current user settings.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.UserSettingsGetCurrentUserSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.UserSettingsGetCurrentUserSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSettingsGetCurrentUserSettings`: UserSettingsVM
    fmt.Fprintf(os.Stdout, "Response from `UserSettingsApi.UserSettingsGetCurrentUserSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsGetCurrentUserSettingsRequest struct via the builder pattern


### Return type

[**UserSettingsVM**](UserSettingsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserSettingsUpdateMySettings

> UserSettingsVM UserSettingsUpdateMySettings(ctx).UpdateUserSettingsVM(updateUserSettingsVM).Execute()

Update settings of the current user

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
    updateUserSettingsVM := *openapiclient.NewUpdateUserSettingsVM() // UpdateUserSettingsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserSettingsApi.UserSettingsUpdateMySettings(context.Background()).UpdateUserSettingsVM(updateUserSettingsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsApi.UserSettingsUpdateMySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserSettingsUpdateMySettings`: UserSettingsVM
    fmt.Fprintf(os.Stdout, "Response from `UserSettingsApi.UserSettingsUpdateMySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsUpdateMySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserSettingsVM** | [**UpdateUserSettingsVM**](UpdateUserSettingsVM.md) |  | 

### Return type

[**UserSettingsVM**](UserSettingsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

