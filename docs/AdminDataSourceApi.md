# \AdminDataSourceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDataSourceCreateDataSource**](AdminDataSourceApi.md#AdminDataSourceCreateDataSource) | **Post** /api/admin/v1/DataSource | Create new data source
[**AdminDataSourceDeleteDataSource**](AdminDataSourceApi.md#AdminDataSourceDeleteDataSource) | **Delete** /api/admin/v1/DataSource/{id} | Delete datasource by id
[**AdminDataSourceFetchData**](AdminDataSourceApi.md#AdminDataSourceFetchData) | **Get** /api/admin/v1/DataSource/{id}/fetch | This should connect to a database and set data structure
[**AdminDataSourceGetDataSource**](AdminDataSourceApi.md#AdminDataSourceGetDataSource) | **Get** /api/admin/v1/DataSource/{id} | Get datasource by id
[**AdminDataSourceGetDataSources**](AdminDataSourceApi.md#AdminDataSourceGetDataSources) | **Get** /api/admin/v1/DataSource | Get list of datasources from database
[**AdminDataSourceGetPermissions**](AdminDataSourceApi.md#AdminDataSourceGetPermissions) | **Get** /api/admin/v1/DataSource/{id}/permissions | Get all Data source permissions
[**AdminDataSourceUpdateDataSource**](AdminDataSourceApi.md#AdminDataSourceUpdateDataSource) | **Put** /api/admin/v1/DataSource/{id} | Update datasource with update VM
[**AdminDataSourceUpdatePermissions**](AdminDataSourceApi.md#AdminDataSourceUpdatePermissions) | **Post** /api/admin/v1/DataSource/{dataSourceId}/permissions | Update permissions to datasource



## AdminDataSourceCreateDataSource

> DataSourceVM AdminDataSourceCreateDataSource(ctx).CreateVM(createVM).Execute()

Create new data source

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
    createVM := *openapiclient.NewCreateDataSourceAdminVM() // CreateDataSourceAdminVM | create VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceCreateDataSource(context.Background()).CreateVM(createVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceCreateDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminDataSourceCreateDataSource`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `AdminDataSourceApi.AdminDataSourceCreateDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceCreateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVM** | [**CreateDataSourceAdminVM**](CreateDataSourceAdminVM.md) | create VM | 

### Return type

[**DataSourceVM**](DataSourceVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDataSourceDeleteDataSource

> AdminDataSourceDeleteDataSource(ctx, id).Execute()

Delete datasource by id

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
    id := "id_example" // string | datasource's id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceDeleteDataSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceDeleteDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | datasource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceDeleteDataSourceRequest struct via the builder pattern


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


## AdminDataSourceFetchData

> AdminDataSourceFetchData(ctx, id).Execute()

This should connect to a database and set data structure

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
    id := "id_example" // string | datasource's id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceFetchData(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceFetchData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | datasource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceFetchDataRequest struct via the builder pattern


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


## AdminDataSourceGetDataSource

> DataSourceVM AdminDataSourceGetDataSource(ctx, id).Execute()

Get datasource by id

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
    id := "id_example" // string | datasource's id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceGetDataSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceGetDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminDataSourceGetDataSource`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `AdminDataSourceApi.AdminDataSourceGetDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | datasource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceGetDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSourceVM**](DataSourceVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDataSourceGetDataSources

> DataSourcesVM AdminDataSourceGetDataSources(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

Get list of datasources from database

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
    skip := int32(56) // int32 | how many datasources will be skiped (optional)
    take := int32(56) // int32 | how many datasources will be taken (optional)
    subscriptionId := "subscriptionId_example" // string | Allow filters by subscription ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceGetDataSources(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceGetDataSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminDataSourceGetDataSources`: DataSourcesVM
    fmt.Fprintf(os.Stdout, "Response from `AdminDataSourceApi.AdminDataSourceGetDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceGetDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | how many datasources will be skiped | 
 **take** | **int32** | how many datasources will be taken | 
 **subscriptionId** | **string** | Allow filters by subscription ID | 

### Return type

[**DataSourcesVM**](DataSourcesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDataSourceGetPermissions

> DataSourcePermissionsVM AdminDataSourceGetPermissions(ctx, id).Execute()

Get all Data source permissions

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
    id := "id_example" // string | data source id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminDataSourceGetPermissions`: DataSourcePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminDataSourceApi.AdminDataSourceGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DataSourcePermissionsVM**](DataSourcePermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDataSourceUpdateDataSource

> DataSourceVM AdminDataSourceUpdateDataSource(ctx, id).UpdateVM(updateVM).Execute()

Update datasource with update VM

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
    id := "id_example" // string | datasource's id
    updateVM := *openapiclient.NewUpdateDataSourceVM() // UpdateDataSourceVM | update VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceUpdateDataSource(context.Background(), id).UpdateVM(updateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceUpdateDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminDataSourceUpdateDataSource`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `AdminDataSourceApi.AdminDataSourceUpdateDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | datasource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceUpdateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVM** | [**UpdateDataSourceVM**](UpdateDataSourceVM.md) | update VM | 

### Return type

[**DataSourceVM**](DataSourceVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminDataSourceUpdatePermissions

> AdminDataSourceUpdatePermissions(ctx, dataSourceId).NewPermissions(newPermissions).Execute()

Update permissions to datasource

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
    dataSourceId := "dataSourceId_example" // string | 
    newPermissions := *openapiclient.NewUpdateDataSourcePermissionsVM(*openapiclient.NewDataSourcePermissions(), int32(123)) // UpdateDataSourcePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminDataSourceApi.AdminDataSourceUpdatePermissions(context.Background(), dataSourceId).NewPermissions(newPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminDataSourceApi.AdminDataSourceUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dataSourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminDataSourceUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newPermissions** | [**UpdateDataSourcePermissionsVM**](UpdateDataSourcePermissionsVM.md) |  | 

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

