# \ApiKeysAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiKeysCreateApiKey**](ApiKeysAPI.md#ApiKeysCreateApiKey) | **Post** /api/manage/v1/ApiKeys | Create a new apikey, 5 apikeys for user. Hardcoded for ddos.
[**ApiKeysDeleteApiKey**](ApiKeysAPI.md#ApiKeysDeleteApiKey) | **Delete** /api/manage/v1/ApiKeys | Delete an apikey
[**ApiKeysGetApiKeys**](ApiKeysAPI.md#ApiKeysGetApiKeys) | **Get** /api/manage/v1/ApiKeys | Returns list with all apikeys of current user



## ApiKeysCreateApiKey

> ApiKeyVM ApiKeysCreateApiKey(ctx).CreateApiKeyVM(createApiKeyVM).Execute()

Create a new apikey, 5 apikeys for user. Hardcoded for ddos.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	createApiKeyVM := *openapiclient.NewCreateApiKeyVM(time.Now(), "T_example") // CreateApiKeyVM | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeysAPI.ApiKeysCreateApiKey(context.Background()).CreateApiKeyVM(createApiKeyVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysAPI.ApiKeysCreateApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeysCreateApiKey`: ApiKeyVM
	fmt.Fprintf(os.Stdout, "Response from `ApiKeysAPI.ApiKeysCreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeysCreateApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiKeyVM** | [**CreateApiKeyVM**](CreateApiKeyVM.md) |  | 

### Return type

[**ApiKeyVM**](ApiKeyVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeysDeleteApiKey

> ApiKeyVM ApiKeysDeleteApiKey(ctx).DeleteApiKeyVM(deleteApiKeyVM).Execute()

Delete an apikey

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
	deleteApiKeyVM := *openapiclient.NewDeleteApiKeyVM("ApiKey_example", "T_example") // DeleteApiKeyVM | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeysAPI.ApiKeysDeleteApiKey(context.Background()).DeleteApiKeyVM(deleteApiKeyVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysAPI.ApiKeysDeleteApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeysDeleteApiKey`: ApiKeyVM
	fmt.Fprintf(os.Stdout, "Response from `ApiKeysAPI.ApiKeysDeleteApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiKeysDeleteApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteApiKeyVM** | [**DeleteApiKeyVM**](DeleteApiKeyVM.md) |  | 

### Return type

[**ApiKeyVM**](ApiKeyVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiKeysGetApiKeys

> ApiKeysVM ApiKeysGetApiKeys(ctx).Execute()

Returns list with all apikeys of current user



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
	resp, r, err := apiClient.ApiKeysAPI.ApiKeysGetApiKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysAPI.ApiKeysGetApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiKeysGetApiKeys`: ApiKeysVM
	fmt.Fprintf(os.Stdout, "Response from `ApiKeysAPI.ApiKeysGetApiKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiKeysGetApiKeysRequest struct via the builder pattern


### Return type

[**ApiKeysVM**](ApiKeysVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

