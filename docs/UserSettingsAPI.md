# \UserSettingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserSettingsAcceptAgreements**](UserSettingsAPI.md#UserSettingsAcceptAgreements) | **Post** /api/manage/v1/UserSettings/accept | Use this endpoint to accept current version of service license agreement
[**UserSettingsGetCurrentUserSettings**](UserSettingsAPI.md#UserSettingsGetCurrentUserSettings) | **Get** /api/manage/v1/UserSettings | Return current user settings.
[**UserSettingsUpdateMySettings**](UserSettingsAPI.md#UserSettingsUpdateMySettings) | **Put** /api/manage/v1/UserSettings | Update settings of the current user



## UserSettingsAcceptAgreements

> UserSettingsAcceptAgreements(ctx).AcceptAgreementsVM(acceptAgreementsVM).Execute()

Use this endpoint to accept current version of service license agreement

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
	acceptAgreementsVM := *openapiclient.NewAcceptAgreementsVM("T_example") // AcceptAgreementsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserSettingsAPI.UserSettingsAcceptAgreements(context.Background()).AcceptAgreementsVM(acceptAgreementsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.UserSettingsAcceptAgreements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserSettingsAcceptAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptAgreementsVM** | [**AcceptAgreementsVM**](AcceptAgreementsVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingsAPI.UserSettingsGetCurrentUserSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.UserSettingsGetCurrentUserSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsGetCurrentUserSettings`: UserSettingsVM
	fmt.Fprintf(os.Stdout, "Response from `UserSettingsAPI.UserSettingsGetCurrentUserSettings`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	updateUserSettingsVM := *openapiclient.NewUpdateUserSettingsVM("T_example") // UpdateUserSettingsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSettingsAPI.UserSettingsUpdateMySettings(context.Background()).UpdateUserSettingsVM(updateUserSettingsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSettingsAPI.UserSettingsUpdateMySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSettingsUpdateMySettings`: UserSettingsVM
	fmt.Fprintf(os.Stdout, "Response from `UserSettingsAPI.UserSettingsUpdateMySettings`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

