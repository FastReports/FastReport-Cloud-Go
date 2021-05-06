# \DataSourcesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataSourcesCreateDataSource**](DataSourcesApi.md#DataSourcesCreateDataSource) | **Post** /api/data/v1/DataSources | Create new data source
[**DataSourcesDeleteDataSource**](DataSourcesApi.md#DataSourcesDeleteDataSource) | **Delete** /api/data/v1/DataSources/{id} | Delete data source by id
[**DataSourcesFetchData**](DataSourcesApi.md#DataSourcesFetchData) | **Get** /api/data/v1/DataSources/{id}/fetch | This should connect to a database and set data structure
[**DataSourcesGetAvailableDataSources**](DataSourcesApi.md#DataSourcesGetAvailableDataSources) | **Get** /api/data/v1/DataSources | Returns all of the data sources, that current user have permission for in a subscription  if subscription id is null, returns all data sources, that current user have permission for
[**DataSourcesGetDataSource**](DataSourcesApi.md#DataSourcesGetDataSource) | **Get** /api/data/v1/DataSources/{id} | Get data source by id
[**DataSourcesGetPermissions**](DataSourcesApi.md#DataSourcesGetPermissions) | **Get** /api/data/v1/DataSources/{id}/permissions | Get all Data source permissions
[**DataSourcesRenameDataSource**](DataSourcesApi.md#DataSourcesRenameDataSource) | **Put** /api/data/v1/DataSources/{id}/rename | Rename data source by id
[**DataSourcesUpdateConnectionString**](DataSourcesApi.md#DataSourcesUpdateConnectionString) | **Put** /api/data/v1/DataSources/{id}/ConnectionString | Update data source&#39;s connection string by id
[**DataSourcesUpdatePermissions**](DataSourcesApi.md#DataSourcesUpdatePermissions) | **Post** /api/data/v1/DataSources/{id}/permissions | Update permissions
[**DataSourcesUpdateSubscriptionDataSource**](DataSourcesApi.md#DataSourcesUpdateSubscriptionDataSource) | **Put** /api/data/v1/DataSources/{id}/updateSubscription | Update data source&#39;s subscription



## DataSourcesCreateDataSource

> DataSourceVM DataSourcesCreateDataSource(ctx).ViewModel(viewModel).Execute()

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
    viewModel := *openapiclient.NewCreateDataSourceVM() // CreateDataSourceVM | create viewmodel (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DataSourcesCreateDataSource(context.Background()).ViewModel(viewModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesCreateDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataSourcesCreateDataSource`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DataSourcesCreateDataSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesCreateDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewModel** | [**CreateDataSourceVM**](CreateDataSourceVM.md) | create viewmodel | 

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


## DataSourcesDeleteDataSource

> DataSourcesDeleteDataSource(ctx, id).Execute()

Delete data source by id

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
    resp, r, err := api_client.DataSourcesApi.DataSourcesDeleteDataSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesDeleteDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesDeleteDataSourceRequest struct via the builder pattern


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


## DataSourcesFetchData

> DataSourcesFetchData(ctx, id).Execute()

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
    resp, r, err := api_client.DataSourcesApi.DataSourcesFetchData(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesFetchData``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDataSourcesFetchDataRequest struct via the builder pattern


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


## DataSourcesGetAvailableDataSources

> DataSourcesVM DataSourcesGetAvailableDataSources(ctx).SubscriptionId(subscriptionId).Skip(skip).Take(take).Execute()

Returns all of the data sources, that current user have permission for in a subscription  if subscription id is null, returns all data sources, that current user have permission for

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
    subscriptionId := "subscriptionId_example" // string | subscription id (optional)
    skip := int32(56) // int32 | how many data sources will be skipped (optional) (default to 0)
    take := int32(56) // int32 | how many data sources will be taken (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DataSourcesGetAvailableDataSources(context.Background()).SubscriptionId(subscriptionId).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesGetAvailableDataSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataSourcesGetAvailableDataSources`: DataSourcesVM
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DataSourcesGetAvailableDataSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesGetAvailableDataSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | subscription id | 
 **skip** | **int32** | how many data sources will be skipped | [default to 0]
 **take** | **int32** | how many data sources will be taken | [default to 10]

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


## DataSourcesGetDataSource

> DataSourceVM DataSourcesGetDataSource(ctx, id).Execute()

Get data source by id

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
    resp, r, err := api_client.DataSourcesApi.DataSourcesGetDataSource(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesGetDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataSourcesGetDataSource`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DataSourcesGetDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesGetDataSourceRequest struct via the builder pattern


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


## DataSourcesGetPermissions

> DataSourcePermissionsVM DataSourcesGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.DataSourcesApi.DataSourcesGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataSourcesGetPermissions`: DataSourcePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DataSourcesGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesGetPermissionsRequest struct via the builder pattern


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


## DataSourcesRenameDataSource

> DataSourceVM DataSourcesRenameDataSource(ctx, id).RenameModel(renameModel).Execute()

Rename data source by id

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
    renameModel := *openapiclient.NewRenameDataSourceVM() // RenameDataSourceVM | rename viewmodel (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DataSourcesRenameDataSource(context.Background(), id).RenameModel(renameModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesRenameDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataSourcesRenameDataSource`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DataSourcesRenameDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesRenameDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **renameModel** | [**RenameDataSourceVM**](RenameDataSourceVM.md) | rename viewmodel | 

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


## DataSourcesUpdateConnectionString

> DataSourceVM DataSourcesUpdateConnectionString(ctx, id).UpdateModel(updateModel).Execute()

Update data source's connection string by id

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
    updateModel := *openapiclient.NewUpdateDataSourceConnectionStringVM() // UpdateDataSourceConnectionStringVM | update viewmodel (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DataSourcesUpdateConnectionString(context.Background(), id).UpdateModel(updateModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesUpdateConnectionString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataSourcesUpdateConnectionString`: DataSourceVM
    fmt.Fprintf(os.Stdout, "Response from `DataSourcesApi.DataSourcesUpdateConnectionString`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesUpdateConnectionStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateModel** | [**UpdateDataSourceConnectionStringVM**](UpdateDataSourceConnectionStringVM.md) | update viewmodel | 

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


## DataSourcesUpdatePermissions

> DataSourcesUpdatePermissions(ctx, id).PermissionsVM(permissionsVM).Execute()

Update permissions

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
    id := "id_example" // string | 
    permissionsVM := *openapiclient.NewUpdateDataSourcePermissionsVM(*openapiclient.NewDataSourcePermissions(), int32(123)) // UpdateDataSourcePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DataSourcesUpdatePermissions(context.Background(), id).PermissionsVM(permissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDataSourcesUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionsVM** | [**UpdateDataSourcePermissionsVM**](UpdateDataSourcePermissionsVM.md) |  | 

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


## DataSourcesUpdateSubscriptionDataSource

> DataSourcesUpdateSubscriptionDataSource(ctx, id).UpdatesubscriptionModel(updatesubscriptionModel).Execute()

Update data source's subscription

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
    updatesubscriptionModel := *openapiclient.NewUpdateDataSourceSubscriptionVM() // UpdateDataSourceSubscriptionVM | update subscription viewmodel (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataSourcesApi.DataSourcesUpdateSubscriptionDataSource(context.Background(), id).UpdatesubscriptionModel(updatesubscriptionModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataSourcesApi.DataSourcesUpdateSubscriptionDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | data source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDataSourcesUpdateSubscriptionDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatesubscriptionModel** | [**UpdateDataSourceSubscriptionVM**](UpdateDataSourceSubscriptionVM.md) | update subscription viewmodel | 

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

