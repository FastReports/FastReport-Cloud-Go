# \ExportsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportFolderAndFileGetCount**](ExportsApi.md#ExportFolderAndFileGetCount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
[**ExportFolderAndFileGetFoldersAndFiles**](ExportsApi.md#ExportFolderAndFileGetFoldersAndFiles) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
[**ExportFoldersCopyFolder**](ExportsApi.md#ExportFoldersCopyFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
[**ExportFoldersDeleteFolder**](ExportsApi.md#ExportFoldersDeleteFolder) | **Delete** /api/rp/v1/Exports/Folder/{id} | Delete specified folder
[**ExportFoldersGetBreadcrumbs**](ExportsApi.md#ExportFoldersGetBreadcrumbs) | **Get** /api/rp/v1/Exports/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
[**ExportFoldersGetFolder**](ExportsApi.md#ExportFoldersGetFolder) | **Get** /api/rp/v1/Exports/Folder/{id} | Get specified folder
[**ExportFoldersGetFolders**](ExportsApi.md#ExportFoldersGetFolders) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFolders | Get all folders from specified folder
[**ExportFoldersGetFoldersCount**](ExportsApi.md#ExportFoldersGetFoldersCount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
[**ExportFoldersGetPermissions**](ExportsApi.md#ExportFoldersGetPermissions) | **Get** /api/rp/v1/Exports/Folder/{id}/permissions | Get all folder permissions
[**ExportFoldersGetRootFolder**](ExportsApi.md#ExportFoldersGetRootFolder) | **Get** /api/rp/v1/Exports/Root | Get user&#39;s root folder (without parents)
[**ExportFoldersMoveFolder**](ExportsApi.md#ExportFoldersMoveFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Move/{folderId} | Move folder to a specified folder
[**ExportFoldersPostFolder**](ExportsApi.md#ExportFoldersPostFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Folder | Create folder
[**ExportFoldersRenameFolder**](ExportsApi.md#ExportFoldersRenameFolder) | **Put** /api/rp/v1/Exports/Folder/{id}/Rename | Rename a folder
[**ExportFoldersUpdateIcon**](ExportsApi.md#ExportFoldersUpdateIcon) | **Put** /api/rp/v1/Exports/Folder/{id}/Icon | Update a folder&#39;s icon
[**ExportFoldersUpdatePermissions**](ExportsApi.md#ExportFoldersUpdatePermissions) | **Post** /api/rp/v1/Exports/{id}/permissions | Update permissions
[**ExportFoldersUpdateTags**](ExportsApi.md#ExportFoldersUpdateTags) | **Put** /api/rp/v1/Exports/Folder/{id}/UpdateTags | Update tags
[**ExportsCopyFile**](ExportsApi.md#ExportsCopyFile) | **Post** /api/rp/v1/Exports/File/{id}/Copy/{folderId} | Copy file to a specified folder
[**ExportsDeleteFile**](ExportsApi.md#ExportsDeleteFile) | **Delete** /api/rp/v1/Exports/File/{id} | Delete specified file
[**ExportsGetFile**](ExportsApi.md#ExportsGetFile) | **Get** /api/rp/v1/Exports/File/{id} | Get specified file
[**ExportsGetFilesCount**](ExportsApi.md#ExportsGetFilesCount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
[**ExportsGetFilesList**](ExportsApi.md#ExportsGetFilesList) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFiles | Get all files from specified folder. &lt;br /&gt;  User with Get Entity permission can access this method. &lt;br /&gt;  The method will returns minimal infomration about the file: &lt;br /&gt;  id, name, size, editedTime, createdTime, tags, status, statusReason.
[**ExportsGetPermissions**](ExportsApi.md#ExportsGetPermissions) | **Get** /api/rp/v1/Exports/File/{id}/permissions | Get all file permissions
[**ExportsMoveFile**](ExportsApi.md#ExportsMoveFile) | **Post** /api/rp/v1/Exports/File/{id}/Move/{folderId} | Move file to a specified folder
[**ExportsRenameFile**](ExportsApi.md#ExportsRenameFile) | **Put** /api/rp/v1/Exports/File/{id}/Rename | Rename a file
[**ExportsUpdateIcon**](ExportsApi.md#ExportsUpdateIcon) | **Put** /api/rp/v1/Exports/File/{id}/Icon | Update a files&#39;s icon
[**ExportsUpdatePermissions**](ExportsApi.md#ExportsUpdatePermissions) | **Post** /api/rp/v1/Exports/File/{id}/permissions | Update permissions
[**ExportsUpdateTags**](ExportsApi.md#ExportsUpdateTags) | **Put** /api/rp/v1/Exports/File/{id}/UpdateTags | Update tags



## ExportFolderAndFileGetCount

> CountVM ExportFolderAndFileGetCount(ctx, id).SearchPattern(searchPattern).Execute()

Get count of files and folders what contains in a specified folder



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
    searchPattern := "searchPattern_example" // string | string, that must be incuded in file or folder name to be counted <br />              (leave undefined to count all files and folders) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFolderAndFileGetCount(context.Background(), id).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFolderAndFileGetCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFolderAndFileGetCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFolderAndFileGetCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFolderAndFileGetCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchPattern** | **string** | string, that must be incuded in file or folder name to be counted &lt;br /&gt;              (leave undefined to count all files and folders) | 

### Return type

[**CountVM**](CountVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFolderAndFileGetFoldersAndFiles

> FilesVM ExportFolderAndFileGetFoldersAndFiles(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).Execute()

Get all folders and files from specified folder



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
    skip := int32(56) // int32 | number of folder and files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of folder and files, that have to be returned (optional) (default to 10)
    orderBy := openapiclient.FileSorting("None") // FileSorting | indicates a field to sort by (optional)
    desc := true // bool | indicates if sorting is descending (optional) (default to false)
    searchPattern := "searchPattern_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFolderAndFileGetFoldersAndFiles(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFolderAndFileGetFoldersAndFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFolderAndFileGetFoldersAndFiles`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFolderAndFileGetFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFolderAndFileGetFoldersAndFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | number of folder and files, that have to be skipped | [default to 0]
 **take** | **int32** | number of folder and files, that have to be returned | [default to 10]
 **orderBy** | [**FileSorting**](FileSorting.md) | indicates a field to sort by | 
 **desc** | **bool** | indicates if sorting is descending | [default to false]
 **searchPattern** | **string** |  | [default to &quot;&quot;]

### Return type

[**FilesVM**](FilesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersCopyFolder

> FileVM ExportFoldersCopyFolder(ctx, id, folderId).Execute()

Move folder to a specified folder



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
    id := "id_example" // string | moving folder id
    folderId := "folderId_example" // string | destination folder id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersCopyFolder(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersCopyFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersCopyFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersCopyFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | moving folder id | 
**folderId** | **string** | destination folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersCopyFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersDeleteFolder

> ExportFoldersDeleteFolder(ctx, id).Recursive(recursive).Execute()

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
    recursive := true // bool | delete all childs (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersDeleteFolder(context.Background(), id).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersDeleteFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFoldersDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | delete all childs | [default to true]

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


## ExportFoldersGetBreadcrumbs

> BreadcrumbsVM ExportFoldersGetBreadcrumbs(ctx, id).Execute()

Get specified folder breadcrumbs



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
    resp, r, err := api_client.ExportsApi.ExportFoldersGetBreadcrumbs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersGetBreadcrumbs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersGetBreadcrumbs`: BreadcrumbsVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersGetBreadcrumbs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetBreadcrumbsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BreadcrumbsVM**](BreadcrumbsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersGetFolder

> FileVM ExportFoldersGetFolder(ctx, id).Execute()

Get specified folder



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
    resp, r, err := api_client.ExportsApi.ExportFoldersGetFolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersGetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersGetFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersGetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersGetFolders

> FilesVM ExportFoldersGetFolders(ctx, id).Skip(skip).Take(take).Execute()

Get all folders from specified folder



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
    skip := int32(56) // int32 | number of files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of files, that have to be returned (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersGetFolders(context.Background(), id).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersGetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersGetFolders`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersGetFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | number of files, that have to be skipped | [default to 0]
 **take** | **int32** | number of files, that have to be returned | [default to 10]

### Return type

[**FilesVM**](FilesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersGetFoldersCount

> CountVM ExportFoldersGetFoldersCount(ctx, id).Execute()

Get count of folders what contains in a specified folder



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
    resp, r, err := api_client.ExportsApi.ExportFoldersGetFoldersCount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersGetFoldersCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersGetFoldersCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersGetFoldersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetFoldersCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountVM**](CountVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersGetPermissions

> FilePermissionsVM ExportFoldersGetPermissions(ctx, id).Execute()

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FilePermissionsVM**](FilePermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersGetRootFolder

> FileVM ExportFoldersGetRootFolder(ctx).SubscriptionId(subscriptionId).Execute()

Get user's root folder (without parents)



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
    subscriptionId := "subscriptionId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersGetRootFolder(context.Background()).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersGetRootFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersGetRootFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersGetRootFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetRootFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** |  | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersMoveFolder

> FileVM ExportFoldersMoveFolder(ctx, id, folderId).Execute()

Move folder to a specified folder



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
    id := "id_example" // string | moving folder id
    folderId := "folderId_example" // string | destination folder id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersMoveFolder(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersMoveFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersMoveFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersMoveFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | moving folder id | 
**folderId** | **string** | destination folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersMoveFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersPostFolder

> FileVM ExportFoldersPostFolder(ctx, id).ExportFolderCreateVM(exportFolderCreateVM).Execute()

Create folder



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
    id := "id_example" // string | Identifier of parent folder id
    exportFolderCreateVM := *openapiclient.NewExportFolderCreateVM() // ExportFolderCreateVM | create VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersPostFolder(context.Background(), id).ExportFolderCreateVM(exportFolderCreateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersPostFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersPostFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersPostFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of parent folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersPostFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportFolderCreateVM** | [**ExportFolderCreateVM**](ExportFolderCreateVM.md) | create VM | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersRenameFolder

> FileVM ExportFoldersRenameFolder(ctx, id).FolderRenameVM(folderRenameVM).Execute()

Rename a folder



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
    folderRenameVM := *openapiclient.NewFolderRenameVM("Name_example") // FolderRenameVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersRenameFolder(context.Background(), id).FolderRenameVM(folderRenameVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersRenameFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersRenameFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersRenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersRenameFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderRenameVM** | [**FolderRenameVM**](FolderRenameVM.md) |  | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersUpdateIcon

> FileVM ExportFoldersUpdateIcon(ctx, id).FolderIconVM(folderIconVM).Execute()

Update a folder's icon



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
    id := "id_example" // string | Identifier of folder
    folderIconVM := *openapiclient.NewFolderIconVM(string(123)) // FolderIconVM | Update icon model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersUpdateIcon(context.Background(), id).FolderIconVM(folderIconVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersUpdateIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersUpdateIcon`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersUpdateIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersUpdateIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderIconVM** | [**FolderIconVM**](FolderIconVM.md) | Update icon model | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersUpdatePermissions

> ExportFoldersUpdatePermissions(ctx, id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()

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
    updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissions(), openapiclient.FileAdministrate(0)) // UpdateFilePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFoldersUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFilePermissionsVM** | [**UpdateFilePermissionsVM**](UpdateFilePermissionsVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersUpdateTags

> FileVM ExportFoldersUpdateTags(ctx, id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()

Update tags



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
    folderTagsUpdateVM := *openapiclient.NewFolderTagsUpdateVM() // FolderTagsUpdateVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportFoldersUpdateTags(context.Background(), id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportFoldersUpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFoldersUpdateTags`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportFoldersUpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderTagsUpdateVM** | [**FolderTagsUpdateVM**](FolderTagsUpdateVM.md) |  | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsCopyFile

> ExportVM ExportsCopyFile(ctx, id, folderId).Execute()

Copy file to a specified folder

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
    folderId := "folderId_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsCopyFile(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsCopyFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsCopyFile`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsCopyFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**folderId** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsCopyFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsDeleteFile

> ExportsDeleteFile(ctx, id).Execute()

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
    resp, r, err := api_client.ExportsApi.ExportsDeleteFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsDeleteFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportsDeleteFileRequest struct via the builder pattern


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


## ExportsGetFile

> ExportVM ExportsGetFile(ctx, id).Execute()

Get specified file



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
    resp, r, err := api_client.ExportsApi.ExportsGetFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsGetFile`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsGetFilesCount

> CountVM ExportsGetFilesCount(ctx, id).Execute()

Get count of files what contains in a specified folder



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
    resp, r, err := api_client.ExportsApi.ExportsGetFilesCount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsGetFilesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsGetFilesCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsGetFilesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsGetFilesCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountVM**](CountVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsGetFilesList

> ExportsVM ExportsGetFilesList(ctx, id).Skip(skip).Take(take).SearchPattern(searchPattern).Execute()

Get all files from specified folder. <br />  User with Get Entity permission can access this method. <br />  The method will returns minimal infomration about the file: <br />  id, name, size, editedTime, createdTime, tags, status, statusReason.

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
    skip := int32(56) // int32 | number of files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of files, that have to be returned (optional) (default to 10)
    searchPattern := "searchPattern_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsGetFilesList(context.Background(), id).Skip(skip).Take(take).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsGetFilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsGetFilesList`: ExportsVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsGetFilesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsGetFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | number of files, that have to be skipped | [default to 0]
 **take** | **int32** | number of files, that have to be returned | [default to 10]
 **searchPattern** | **string** |  | 

### Return type

[**ExportsVM**](ExportsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsGetPermissions

> FilePermissionsVM ExportsGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.ExportsApi.ExportsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FilePermissionsVM**](FilePermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsMoveFile

> ExportVM ExportsMoveFile(ctx, id, folderId).Execute()

Move file to a specified folder



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
    folderId := "folderId_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsMoveFile(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsMoveFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsMoveFile`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsMoveFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**folderId** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsMoveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsRenameFile

> ExportVM ExportsRenameFile(ctx, id).FileRenameVM(fileRenameVM).Execute()

Rename a file



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
    fileRenameVM := *openapiclient.NewFileRenameVM() // FileRenameVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsRenameFile(context.Background(), id).FileRenameVM(fileRenameVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsRenameFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsRenameFile`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsRenameFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsRenameFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileRenameVM** | [**FileRenameVM**](FileRenameVM.md) |  | 

### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsUpdateIcon

> ExportVM ExportsUpdateIcon(ctx, id).FileIconVM(fileIconVM).Execute()

Update a files's icon



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
    fileIconVM := *openapiclient.NewFileIconVM() // FileIconVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsUpdateIcon(context.Background(), id).FileIconVM(fileIconVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsUpdateIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsUpdateIcon`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsUpdateIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsUpdateIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileIconVM** | [**FileIconVM**](FileIconVM.md) |  | 

### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsUpdatePermissions

> ExportsUpdatePermissions(ctx, id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()

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
    updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissions(), openapiclient.FileAdministrate(0)) // UpdateFilePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportsUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFilePermissionsVM** | [**UpdateFilePermissionsVM**](UpdateFilePermissionsVM.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportsUpdateTags

> ExportVM ExportsUpdateTags(ctx, id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()

Update tags



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
    fileTagsUpdateVM := *openapiclient.NewFileTagsUpdateVM() // FileTagsUpdateVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ExportsApi.ExportsUpdateTags(context.Background(), id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExportsApi.ExportsUpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportsUpdateTags`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ExportsApi.ExportsUpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileTagsUpdateVM** | [**FileTagsUpdateVM**](FileTagsUpdateVM.md) |  | 

### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

