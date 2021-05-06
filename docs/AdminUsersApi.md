# \AdminUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminUsersDeleteUser**](AdminUsersApi.md#AdminUsersDeleteUser) | **Delete** /api/admin/v1/Users/{id} | Delete a user from cloud database by id
[**AdminUsersGetNewUsersPerMonth**](AdminUsersApi.md#AdminUsersGetNewUsersPerMonth) | **Get** /api/admin/v1/Users/stat/new/{from}/{to} | Returns a key-value pair of new users count per month for a specified time span: (month, number of new users)
[**AdminUsersGetUser**](AdminUsersApi.md#AdminUsersGetUser) | **Get** /api/admin/v1/Users/{id} | Returns a user view model by id
[**AdminUsersGetUsers**](AdminUsersApi.md#AdminUsersGetUsers) | **Get** /api/admin/v1/Users | Returns a list of users
[**AdminUsersRegisterUser**](AdminUsersApi.md#AdminUsersRegisterUser) | **Post** /api/admin/v1/Users | Register a new user in the cloud with specified id, returns a new model
[**AdminUsersUpdateUser**](AdminUsersApi.md#AdminUsersUpdateUser) | **Put** /api/admin/v1/Users/{id} | Update an user by id



## AdminUsersDeleteUser

> AdminUsersDeleteUser(ctx, id).Execute()

Delete a user from cloud database by id

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
    id := "id_example" // string | Identifier of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminUsersApi.AdminUsersDeleteUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersApi.AdminUsersDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersDeleteUserRequest struct via the builder pattern


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


## AdminUsersGetNewUsersPerMonth

> map[string]int32 AdminUsersGetNewUsersPerMonth(ctx, from, to).Execute()

Returns a key-value pair of new users count per month for a specified time span: (month, number of new users)



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
    from := time.Now() // time.Time | A starting date for stats calculation
    to := time.Now() // time.Time | An ending date for stats calculation

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminUsersApi.AdminUsersGetNewUsersPerMonth(context.Background(), from, to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersApi.AdminUsersGetNewUsersPerMonth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersGetNewUsersPerMonth`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `AdminUsersApi.AdminUsersGetNewUsersPerMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**from** | **time.Time** | A starting date for stats calculation | 
**to** | **time.Time** | An ending date for stats calculation | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersGetNewUsersPerMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]int32**

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsersGetUser

> UserVM AdminUsersGetUser(ctx, id).Execute()

Returns a user view model by id

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
    id := "id_example" // string | Identifier of user

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminUsersApi.AdminUsersGetUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersApi.AdminUsersGetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersGetUser`: UserVM
    fmt.Fprintf(os.Stdout, "Response from `AdminUsersApi.AdminUsersGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserVM**](UserVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsersGetUsers

> UsersVM AdminUsersGetUsers(ctx).Skip(skip).Take(take).Execute()

Returns a list of users



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminUsersApi.AdminUsersGetUsers(context.Background()).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersApi.AdminUsersGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersGetUsers`: UsersVM
    fmt.Fprintf(os.Stdout, "Response from `AdminUsersApi.AdminUsersGetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]

### Return type

[**UsersVM**](UsersVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsersRegisterUser

> UserVM AdminUsersRegisterUser(ctx).ViewModel(viewModel).Execute()

Register a new user in the cloud with specified id, returns a new model

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
    viewModel := *openapiclient.NewRegisterUserVM() // RegisterUserVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminUsersApi.AdminUsersRegisterUser(context.Background()).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersApi.AdminUsersRegisterUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersRegisterUser`: UserVM
    fmt.Fprintf(os.Stdout, "Response from `AdminUsersApi.AdminUsersRegisterUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersRegisterUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewModel** | [**RegisterUserVM**](RegisterUserVM.md) |  | 

### Return type

[**UserVM**](UserVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminUsersUpdateUser

> UserVM AdminUsersUpdateUser(ctx, id).ViewModel(viewModel).Execute()

Update an user by id

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
    id := "id_example" // string | Identifier of the user
    viewModel := *openapiclient.NewUpdateUserVM() // UpdateUserVM | update VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminUsersApi.AdminUsersUpdateUser(context.Background(), id).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersApi.AdminUsersUpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminUsersUpdateUser`: UserVM
    fmt.Fprintf(os.Stdout, "Response from `AdminUsersApi.AdminUsersUpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminUsersUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewModel** | [**UpdateUserVM**](UpdateUserVM.md) | update VM | 

### Return type

[**UserVM**](UserVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

