# \GroupUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupUsersAddUserToGroup**](GroupUsersApi.md#GroupUsersAddUserToGroup) | **Put** /api/manage/v1/Groups/{id}/Users/{userId} | Add user to the group by identifier
[**GroupUsersGetUsersInGroup**](GroupUsersApi.md#GroupUsersGetUsersInGroup) | **Get** /api/manage/v1/Groups/{id}/Users | Returns users in the group by identifier
[**GroupUsersLeaveFromGroup**](GroupUsersApi.md#GroupUsersLeaveFromGroup) | **Delete** /api/manage/v1/Groups/{id}/leave | Leave from the group
[**GroupUsersRemoveFromGroup**](GroupUsersApi.md#GroupUsersRemoveFromGroup) | **Delete** /api/manage/v1/Groups/{id}/Users/{userId} | Remove user from the group by identifier



## GroupUsersAddUserToGroup

> GroupUsersAddUserToGroup(ctx, id, userId).Execute()

Add user to the group by identifier

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
    userId := "userId_example" // string | Identifier of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupUsersApi.GroupUsersAddUserToGroup(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupUsersApi.GroupUsersAddUserToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 
**userId** | **string** | Identifier of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupUsersAddUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupUsersGetUsersInGroup

> GroupUsersVM GroupUsersGetUsersInGroup(ctx, id).Skip(skip).Take(take).Execute()

Returns users in the group by identifier

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
    skip := int32(56) // int32 | how many to skip (optional) (default to 0)
    take := int32(56) // int32 | how many to take (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupUsersApi.GroupUsersGetUsersInGroup(context.Background(), id).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupUsersApi.GroupUsersGetUsersInGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupUsersGetUsersInGroup`: GroupUsersVM
    fmt.Fprintf(os.Stdout, "Response from `GroupUsersApi.GroupUsersGetUsersInGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupUsersGetUsersInGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | how many to skip | [default to 0]
 **take** | **int32** | how many to take | [default to 10]

### Return type

[**GroupUsersVM**](GroupUsersVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupUsersLeaveFromGroup

> GroupUsersLeaveFromGroup(ctx, id).Execute()

Leave from the group

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
    resp, r, err := api_client.GroupUsersApi.GroupUsersLeaveFromGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupUsersApi.GroupUsersLeaveFromGroup``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGroupUsersLeaveFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupUsersRemoveFromGroup

> GroupUsersRemoveFromGroup(ctx, id, userId).Execute()

Remove user from the group by identifier

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
    userId := "userId_example" // string | Identifier of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GroupUsersApi.GroupUsersRemoveFromGroup(context.Background(), id, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupUsersApi.GroupUsersRemoveFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of group | 
**userId** | **string** | Identifier of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupUsersRemoveFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

