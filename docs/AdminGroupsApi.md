# \AdminGroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGroupsCreateGroup**](AdminGroupsApi.md#AdminGroupsCreateGroup) | **Post** /api/admin/v1/Groups | Create a new group, returns a new model
[**AdminGroupsDeleteGroup**](AdminGroupsApi.md#AdminGroupsDeleteGroup) | **Delete** /api/admin/v1/Groups/{id} | Delete a group by id
[**AdminGroupsGetGroup**](AdminGroupsApi.md#AdminGroupsGetGroup) | **Get** /api/admin/v1/Groups/{id} | Returns a group by id
[**AdminGroupsGetGroups**](AdminGroupsApi.md#AdminGroupsGetGroups) | **Get** /api/admin/v1/Groups | Returns a list of groups
[**AdminGroupsGetPermissions**](AdminGroupsApi.md#AdminGroupsGetPermissions) | **Get** /api/admin/v1/Groups/{id}/permissions | Gets group permissions by identifier
[**AdminGroupsUpdateGroup**](AdminGroupsApi.md#AdminGroupsUpdateGroup) | **Put** /api/admin/v1/Groups/{id} | Update a group by id
[**AdminGroupsUpdatePermissions**](AdminGroupsApi.md#AdminGroupsUpdatePermissions) | **Post** /api/admin/v1/Groups/{id}/permissions | Update permissions in user group by identifier



## AdminGroupsCreateGroup

> GroupVM AdminGroupsCreateGroup(ctx).ViewModel(viewModel).Execute()

Create a new group, returns a new model

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
    viewModel := *openapiclient.NewCreateGroupAdminVM("Name_example") // CreateGroupAdminVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsCreateGroup(context.Background()).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsCreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGroupsCreateGroup`: GroupVM
    fmt.Fprintf(os.Stdout, "Response from `AdminGroupsApi.AdminGroupsCreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewModel** | [**CreateGroupAdminVM**](CreateGroupAdminVM.md) |  | 

### Return type

[**GroupVM**](GroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGroupsDeleteGroup

> AdminGroupsDeleteGroup(ctx, id).Execute()

Delete a group by id

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
    id := "id_example" // string | Identifier of group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsDeleteGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsDeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGroupsGetGroup

> GroupVM AdminGroupsGetGroup(ctx, id).Execute()

Returns a group by id

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
    id := "id_example" // string | Identifier of group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsGetGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsGetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGroupsGetGroup`: GroupVM
    fmt.Fprintf(os.Stdout, "Response from `AdminGroupsApi.AdminGroupsGetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupVM**](GroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGroupsGetGroups

> GroupsVM AdminGroupsGetGroups(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

Returns a list of groups



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
    skip := int32(56) // int32 | Variable for pagination, defautl value is 0 (optional) (default to 0)
    take := int32(56) // int32 | Variable for pagination, default value is 10 (optional) (default to 10)
    subscriptionId := "subscriptionId_example" // string | Allow to filter by subscription id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsGetGroups(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGroupsGetGroups`: GroupsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminGroupsApi.AdminGroupsGetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]
 **subscriptionId** | **string** | Allow to filter by subscription id | 

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


## AdminGroupsGetPermissions

> GroupPermissionsVM AdminGroupsGetPermissions(ctx, id).Execute()

Gets group permissions by identifier

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
    id := "id_example" // string | Identifier of group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGroupsGetPermissions`: GroupPermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminGroupsApi.AdminGroupsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupPermissionsVM**](GroupPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGroupsUpdateGroup

> GroupVM AdminGroupsUpdateGroup(ctx, id).ViewModel(viewModel).Execute()

Update a group by id

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
    id := "id_example" // string | Identifier of group
    viewModel := *openapiclient.NewUpdateGroupVM() // UpdateGroupVM | update vm (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsUpdateGroup(context.Background(), id).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsUpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminGroupsUpdateGroup`: GroupVM
    fmt.Fprintf(os.Stdout, "Response from `AdminGroupsApi.AdminGroupsUpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewModel** | [**UpdateGroupVM**](UpdateGroupVM.md) | update vm | 

### Return type

[**GroupVM**](GroupVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminGroupsUpdatePermissions

> AdminGroupsUpdatePermissions(ctx, id).NewPermissions(newPermissions).Execute()

Update permissions in user group by identifier

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
    id := "id_example" // string | Identifier of group
    newPermissions := *openapiclient.NewUpdateGroupPermissionsVM(*openapiclient.NewGroupPermissions(), int32(123)) // UpdateGroupPermissionsVM | Model with new permissions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminGroupsApi.AdminGroupsUpdatePermissions(context.Background(), id).NewPermissions(newPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminGroupsApi.AdminGroupsUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminGroupsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newPermissions** | [**UpdateGroupPermissionsVM**](UpdateGroupPermissionsVM.md) | Model with new permissions | 

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

