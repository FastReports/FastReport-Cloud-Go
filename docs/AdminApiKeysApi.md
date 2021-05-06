# \AdminApiKeysApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminApiKeysCreateApiKey**](AdminApiKeysApi.md#AdminApiKeysCreateApiKey) | **Post** /api/admin/v1/ApiKeys/{userId} | Create a new apikey, 5 apikeys for user. Hardcoded for ddos.
[**AdminApiKeysDeleteApiKey**](AdminApiKeysApi.md#AdminApiKeysDeleteApiKey) | **Delete** /api/admin/v1/ApiKeys/{userId} | Delete an apikey
[**AdminApiKeysGetApiKeys**](AdminApiKeysApi.md#AdminApiKeysGetApiKeys) | **Get** /api/admin/v1/ApiKeys/{userId} | Returns list with all api keys of a specified user



## AdminApiKeysCreateApiKey

> ApiKeyVM AdminApiKeysCreateApiKey(ctx, userId).Model(model).Execute()

Create a new apikey, 5 apikeys for user. Hardcoded for ddos.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    userId := "userId_example" // string | 
    model := *openapiclient.NewCreateApiKeyVM(time.Now()) // CreateApiKeyVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApiKeysApi.AdminApiKeysCreateApiKey(context.Background(), userId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApiKeysApi.AdminApiKeysCreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminApiKeysCreateApiKey`: ApiKeyVM
    fmt.Fprintf(os.Stdout, "Response from `AdminApiKeysApi.AdminApiKeysCreateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminApiKeysCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**CreateApiKeyVM**](CreateApiKeyVM.md) |  | 

### Return type

[**ApiKeyVM**](ApiKeyVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminApiKeysDeleteApiKey

> AdminApiKeysDeleteApiKey(ctx, userId).Model(model).Execute()

Delete an apikey

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
    userId := "userId_example" // string | 
    model := *openapiclient.NewDeleteApiKeyVM("ApiKey_example") // DeleteApiKeyVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApiKeysApi.AdminApiKeysDeleteApiKey(context.Background(), userId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApiKeysApi.AdminApiKeysDeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminApiKeysDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | [**DeleteApiKeyVM**](DeleteApiKeyVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminApiKeysGetApiKeys

> ApiKeysVM AdminApiKeysGetApiKeys(ctx, userId).Execute()

Returns list with all api keys of a specified user



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminApiKeysApi.AdminApiKeysGetApiKeys(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminApiKeysApi.AdminApiKeysGetApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminApiKeysGetApiKeys`: ApiKeysVM
    fmt.Fprintf(os.Stdout, "Response from `AdminApiKeysApi.AdminApiKeysGetApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminApiKeysGetApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKeysVM**](ApiKeysVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

