# \AdminSubscriptionAnalyticsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubscriptionAnalyticsCheckAnonPermissions**](AdminSubscriptionAnalyticsApi.md#AdminSubscriptionAnalyticsCheckAnonPermissions) | **Get** /api/admin/v1/Analytics/Subscriptions/{subscriptionId}/AnonCheck | This will check if there are any files, related to subscription that available for anonymous users
[**AdminSubscriptionAnalyticsCheckOtherPermissions**](AdminSubscriptionAnalyticsApi.md#AdminSubscriptionAnalyticsCheckOtherPermissions) | **Get** /api/admin/v1/Analytics/Subscriptions/{subscriptionId}/OtherCheck | This will check if there are any files, related to subscription that not available for subscription users
[**AdminSubscriptionAnalyticsGetDeadSubsInUsers**](AdminSubscriptionAnalyticsApi.md#AdminSubscriptionAnalyticsGetDeadSubsInUsers) | **Get** /api/admin/v1/Analytics/Subscriptions/DeadSubsInUsers | This will check if there are any deleted subscriptions in users sub lists
[**AdminSubscriptionAnalyticsGetEmptySubs**](AdminSubscriptionAnalyticsApi.md#AdminSubscriptionAnalyticsGetEmptySubs) | **Get** /api/admin/v1/Analytics/Subscriptions/EmptySubs | This will check if there are any subscriptions without users
[**AdminSubscriptionAnalyticsGetLostFileChunks**](AdminSubscriptionAnalyticsApi.md#AdminSubscriptionAnalyticsGetLostFileChunks) | **Get** /api/admin/v1/Analytics/Subscriptions/LostFileChunks | This will check if there are any files in gridFS, which not related to any file model
[**AdminSubscriptionAnalyticsGetUnrelatedDocuments**](AdminSubscriptionAnalyticsApi.md#AdminSubscriptionAnalyticsGetUnrelatedDocuments) | **Get** /api/admin/v1/Analytics/Subscriptions/UnrelatedDocuments | This will check if there are any files, that not related to any existing subscription



## AdminSubscriptionAnalyticsCheckAnonPermissions

> AnalysisResultsVM AdminSubscriptionAnalyticsCheckAnonPermissions(ctx, subscriptionId).Execute()

This will check if there are any files, related to subscription that available for anonymous users

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
    subscriptionId := "subscriptionId_example" // string | subscription id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsCheckAnonPermissions(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsCheckAnonPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionAnalyticsCheckAnonPermissions`: AnalysisResultsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsCheckAnonPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionAnalyticsCheckAnonPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalysisResultsVM**](AnalysisResultsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionAnalyticsCheckOtherPermissions

> AnalysisResultsVM AdminSubscriptionAnalyticsCheckOtherPermissions(ctx, subscriptionId).Execute()

This will check if there are any files, related to subscription that not available for subscription users

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
    subscriptionId := "subscriptionId_example" // string | subscription id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsCheckOtherPermissions(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsCheckOtherPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionAnalyticsCheckOtherPermissions`: AnalysisResultsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsCheckOtherPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionAnalyticsCheckOtherPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnalysisResultsVM**](AnalysisResultsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionAnalyticsGetDeadSubsInUsers

> AnalysisResultsVM AdminSubscriptionAnalyticsGetDeadSubsInUsers(ctx).Execute()

This will check if there are any deleted subscriptions in users sub lists

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
    resp, r, err := api_client.AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetDeadSubsInUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetDeadSubsInUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionAnalyticsGetDeadSubsInUsers`: AnalysisResultsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetDeadSubsInUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionAnalyticsGetDeadSubsInUsersRequest struct via the builder pattern


### Return type

[**AnalysisResultsVM**](AnalysisResultsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionAnalyticsGetEmptySubs

> AnalysisResultsVM AdminSubscriptionAnalyticsGetEmptySubs(ctx).Execute()

This will check if there are any subscriptions without users

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
    resp, r, err := api_client.AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetEmptySubs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetEmptySubs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionAnalyticsGetEmptySubs`: AnalysisResultsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetEmptySubs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionAnalyticsGetEmptySubsRequest struct via the builder pattern


### Return type

[**AnalysisResultsVM**](AnalysisResultsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionAnalyticsGetLostFileChunks

> AnalysisResultsVM AdminSubscriptionAnalyticsGetLostFileChunks(ctx).Execute()

This will check if there are any files in gridFS, which not related to any file model

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
    resp, r, err := api_client.AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetLostFileChunks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetLostFileChunks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionAnalyticsGetLostFileChunks`: AnalysisResultsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetLostFileChunks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionAnalyticsGetLostFileChunksRequest struct via the builder pattern


### Return type

[**AnalysisResultsVM**](AnalysisResultsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminSubscriptionAnalyticsGetUnrelatedDocuments

> AnalysisResultsVM AdminSubscriptionAnalyticsGetUnrelatedDocuments(ctx).Execute()

This will check if there are any files, that not related to any existing subscription

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
    resp, r, err := api_client.AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetUnrelatedDocuments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetUnrelatedDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminSubscriptionAnalyticsGetUnrelatedDocuments`: AnalysisResultsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionAnalyticsApi.AdminSubscriptionAnalyticsGetUnrelatedDocuments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionAnalyticsGetUnrelatedDocumentsRequest struct via the builder pattern


### Return type

[**AnalysisResultsVM**](AnalysisResultsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

