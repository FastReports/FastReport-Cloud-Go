# \AdminReportsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminReportFoldersDeleteFolder**](AdminReportsApi.md#AdminReportFoldersDeleteFolder) | **Delete** /api/admin/v1/ReportFolders/{id} | Delete specified folder
[**AdminReportFoldersGetFolder**](AdminReportsApi.md#AdminReportFoldersGetFolder) | **Get** /api/admin/v1/ReportFolders/{id} | Returns a folder by id
[**AdminReportFoldersGetFolders**](AdminReportsApi.md#AdminReportFoldersGetFolders) | **Get** /api/admin/v1/ReportFolders | Returns a list of folders
[**AdminReportFoldersGetPermissions**](AdminReportsApi.md#AdminReportFoldersGetPermissions) | **Get** /api/admin/v1/ReportFolders/{id}/permissions | Get all folder permissions
[**AdminReportFoldersPostFolder**](AdminReportsApi.md#AdminReportFoldersPostFolder) | **Post** /api/admin/v1/ReportFolders | Create a folder
[**AdminReportFoldersUpdateFolder**](AdminReportsApi.md#AdminReportFoldersUpdateFolder) | **Put** /api/admin/v1/ReportFolders/{id} | Update a folder
[**AdminReportFoldersUpdatePermissions**](AdminReportsApi.md#AdminReportFoldersUpdatePermissions) | **Post** /api/admin/v1/ReportFolders/{id}/permissions | Revoke permission
[**AdminReportsDeleteFile**](AdminReportsApi.md#AdminReportsDeleteFile) | **Delete** /api/admin/v1/Reports/{id} | Delete specified file
[**AdminReportsGetFile**](AdminReportsApi.md#AdminReportsGetFile) | **Get** /api/admin/v1/Reports/{id} | Returns a file by id
[**AdminReportsGetFiles**](AdminReportsApi.md#AdminReportsGetFiles) | **Get** /api/admin/v1/Reports | Returns a list of files
[**AdminReportsGetPermissions**](AdminReportsApi.md#AdminReportsGetPermissions) | **Get** /api/admin/v1/Reports/{id}/permissions | Get all file permissions
[**AdminReportsUpdateFile**](AdminReportsApi.md#AdminReportsUpdateFile) | **Put** /api/admin/v1/Reports/{id} | Update a file
[**AdminReportsUpdatePermissions**](AdminReportsApi.md#AdminReportsUpdatePermissions) | **Post** /api/admin/v1/Reports/{id}/permissions | Update permissions to file
[**AdminReportsUploadFile**](AdminReportsApi.md#AdminReportsUploadFile) | **Post** /api/admin/v1/Reports | Upload a file to the specified folder



## AdminReportFoldersDeleteFolder

> AdminReportFoldersDeleteFolder(ctx, id).Recursive(recursive).Execute()

Delete specified folder



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
    id := "id_example" // string | folder id
    recursive := true // bool | delete folder's content (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersDeleteFolder(context.Background(), id).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersDeleteFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | delete folder&#39;s content | [default to true]

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


## AdminReportFoldersGetFolder

> FileVM AdminReportFoldersGetFolder(ctx, id).Execute()

Returns a folder by id

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
    id := "id_example" // string | Identifier of a folder

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersGetFolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersGetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportFoldersGetFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportFoldersGetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportFoldersGetFolders

> FilesVM AdminReportFoldersGetFolders(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

Returns a list of folders



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
    subscriptionId := "subscriptionId_example" // string | Allows to filter by subscriptions ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersGetFolders(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersGetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportFoldersGetFolders`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportFoldersGetFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersGetFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]
 **subscriptionId** | **string** | Allows to filter by subscriptions ID | 

### Return type

[**FilesVM**](FilesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportFoldersGetPermissions

> FilePermissionsVM AdminReportFoldersGetPermissions(ctx, id).Execute()

Get all folder permissions

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
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportFoldersGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportFoldersGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FilePermissionsVM**](FilePermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportFoldersPostFolder

> FileVM AdminReportFoldersPostFolder(ctx).FolderVm(folderVm).Execute()

Create a folder



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
    folderVm := *openapiclient.NewAdminReportFolderCreateVM() // AdminReportFolderCreateVM | folder create vm (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersPostFolder(context.Background()).FolderVm(folderVm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersPostFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportFoldersPostFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportFoldersPostFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersPostFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderVm** | [**AdminReportFolderCreateVM**](AdminReportFolderCreateVM.md) | folder create vm | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportFoldersUpdateFolder

> FileVM AdminReportFoldersUpdateFolder(ctx, id).FolderVM(folderVM).Execute()

Update a folder



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
    id := "id_example" // string | folder id
    folderVM := *openapiclient.NewFileUpdateVM() // FileUpdateVM | folder's view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersUpdateFolder(context.Background(), id).FolderVM(folderVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersUpdateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportFoldersUpdateFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportFoldersUpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersUpdateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderVM** | [**FileUpdateVM**](FileUpdateVM.md) | folder&#39;s view model | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportFoldersUpdatePermissions

> AdminReportFoldersUpdatePermissions(ctx, id).NewPermissions(newPermissions).Execute()

Revoke permission

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
    id := "id_example" // string | folder id
    newPermissions := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissions(), int32(123)) // UpdateFilePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportFoldersUpdatePermissions(context.Background(), id).NewPermissions(newPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportFoldersUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportFoldersUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newPermissions** | [**UpdateFilePermissionsVM**](UpdateFilePermissionsVM.md) |  | 

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


## AdminReportsDeleteFile

> AdminReportsDeleteFile(ctx, id).Execute()

Delete specified file



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
    id := "id_example" // string | file id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsDeleteFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsDeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportsDeleteFileRequest struct via the builder pattern


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


## AdminReportsGetFile

> ReportVM AdminReportsGetFile(ctx, id).Execute()

Returns a file by id

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
    id := "id_example" // string | Identifier of file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsGetFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportsGetFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportsGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of file | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportsGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportsGetFiles

> ReportsVM AdminReportsGetFiles(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

Returns a list of files



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
    subscriptionId := "subscriptionId_example" // string | Subscription Id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsGetFiles(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsGetFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportsGetFiles`: ReportsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportsGetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportsGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]
 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**ReportsVM**](ReportsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportsGetPermissions

> FilePermissionsVM AdminReportsGetPermissions(ctx, id).Execute()

Get all file permissions

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportsGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FilePermissionsVM**](FilePermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportsUpdateFile

> ReportVM AdminReportsUpdateFile(ctx, id).FileVM(fileVM).Execute()

Update a file



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
    id := "id_example" // string | file id
    fileVM := *openapiclient.NewFileUpdateVM() // FileUpdateVM | file's view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsUpdateFile(context.Background(), id).FileVM(fileVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsUpdateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportsUpdateFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportsUpdateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportsUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileVM** | [**FileUpdateVM**](FileUpdateVM.md) | file&#39;s view model | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminReportsUpdatePermissions

> AdminReportsUpdatePermissions(ctx, id).NewPermissions(newPermissions).Execute()

Update permissions to file

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
    newPermissions := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissions(), int32(123)) // UpdateFilePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsUpdatePermissions(context.Background(), id).NewPermissions(newPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminReportsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newPermissions** | [**UpdateFilePermissionsVM**](UpdateFilePermissionsVM.md) |  | 

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


## AdminReportsUploadFile

> ReportVM AdminReportsUploadFile(ctx).FileVM(fileVM).Execute()

Upload a file to the specified folder



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
    fileVM := *openapiclient.NewReportCreateAdminVM() // ReportCreateAdminVM | file's view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminReportsApi.AdminReportsUploadFile(context.Background()).FileVM(fileVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminReportsApi.AdminReportsUploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminReportsUploadFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `AdminReportsApi.AdminReportsUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminReportsUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileVM** | [**ReportCreateAdminVM**](ReportCreateAdminVM.md) | file&#39;s view model | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

