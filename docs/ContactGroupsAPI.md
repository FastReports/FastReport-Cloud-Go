# \ContactGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactGroupsCreate**](ContactGroupsAPI.md#ContactGroupsCreate) | **Post** /api/v1/ContactGroups/group | Creates contact group
[**ContactGroupsDelete**](ContactGroupsAPI.md#ContactGroupsDelete) | **Delete** /api/v1/ContactGroups/{id} | Removes contact by id
[**ContactGroupsGet**](ContactGroupsAPI.md#ContactGroupsGet) | **Get** /api/v1/ContactGroups/{id} | Returns contact group by id
[**ContactGroupsGetList**](ContactGroupsAPI.md#ContactGroupsGetList) | **Get** /api/v1/ContactGroups/subscription/{subscriptionId}/groups | Returns contact groups by subscriptionId
[**ContactGroupsUpdate**](ContactGroupsAPI.md#ContactGroupsUpdate) | **Put** /api/v1/ContactGroups/{id} | Updates contact group by id



## ContactGroupsCreate

> ContactGroupVM ContactGroupsCreate(ctx).CreateContactGroupVM(createContactGroupVM).Execute()

Creates contact group

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
	createContactGroupVM := *openapiclient.NewCreateContactGroupVM("Name_example") // CreateContactGroupVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.ContactGroupsCreate(context.Background()).CreateContactGroupVM(createContactGroupVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.ContactGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactGroupsCreate`: ContactGroupVM
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.ContactGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactGroupVM** | [**CreateContactGroupVM**](CreateContactGroupVM.md) |  | 

### Return type

[**ContactGroupVM**](ContactGroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactGroupsDelete

> ContactGroupsDelete(ctx, id).Execute()

Removes contact by id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContactGroupsAPI.ContactGroupsDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.ContactGroupsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactGroupsDeleteRequest struct via the builder pattern


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


## ContactGroupsGet

> ContactGroupVM ContactGroupsGet(ctx, id).Execute()

Returns contact group by id

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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.ContactGroupsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.ContactGroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactGroupsGet`: ContactGroupVM
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.ContactGroupsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactGroupsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactGroupVM**](ContactGroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactGroupsGetList

> ContactGroupsVM ContactGroupsGetList(ctx, subscriptionId).Skip(skip).Take(take).Execute()

Returns contact groups by subscriptionId

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
	skip := int32(56) // int32 |  (optional) (default to 0)
	take := int32(56) // int32 |  (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.ContactGroupsGetList(context.Background(), subscriptionId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.ContactGroupsGetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactGroupsGetList`: ContactGroupsVM
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.ContactGroupsGetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactGroupsGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **take** | **int32** |  | [default to 10]

### Return type

[**ContactGroupsVM**](ContactGroupsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactGroupsUpdate

> ContactGroupVM ContactGroupsUpdate(ctx, id).UpdateContactGroupVM(updateContactGroupVM).Execute()

Updates contact group by id

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
	id := "id_example" // string | 
	updateContactGroupVM := *openapiclient.NewUpdateContactGroupVM("Name_example") // UpdateContactGroupVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactGroupsAPI.ContactGroupsUpdate(context.Background(), id).UpdateContactGroupVM(updateContactGroupVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactGroupsAPI.ContactGroupsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ContactGroupsUpdate`: ContactGroupVM
	fmt.Fprintf(os.Stdout, "Response from `ContactGroupsAPI.ContactGroupsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiContactGroupsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContactGroupVM** | [**UpdateContactGroupVM**](UpdateContactGroupVM.md) |  | 

### Return type

[**ContactGroupVM**](ContactGroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

