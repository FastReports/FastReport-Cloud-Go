# \ReportsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportFolderAndFileGetCount**](ReportsApi.md#ReportFolderAndFileGetCount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
[**ReportFolderAndFileGetFoldersAndFiles**](ReportsApi.md#ReportFolderAndFileGetFoldersAndFiles) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
[**ReportFoldersCopyFolder**](ReportsApi.md#ReportFoldersCopyFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
[**ReportFoldersDeleteFolder**](ReportsApi.md#ReportFoldersDeleteFolder) | **Delete** /api/rp/v1/Reports/Folder/{id} | Delete specified folder
[**ReportFoldersGetBreadcrumbs**](ReportsApi.md#ReportFoldersGetBreadcrumbs) | **Get** /api/rp/v1/Reports/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
[**ReportFoldersGetFolder**](ReportsApi.md#ReportFoldersGetFolder) | **Get** /api/rp/v1/Reports/Folder/{id} | Get specified folder
[**ReportFoldersGetFolders**](ReportsApi.md#ReportFoldersGetFolders) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFolders | Get all folders from specified folder
[**ReportFoldersGetFoldersCount**](ReportsApi.md#ReportFoldersGetFoldersCount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
[**ReportFoldersGetPermissions**](ReportsApi.md#ReportFoldersGetPermissions) | **Get** /api/rp/v1/Reports/Folder/{id}/permissions | Get all folder permissions
[**ReportFoldersGetRootFolder**](ReportsApi.md#ReportFoldersGetRootFolder) | **Get** /api/rp/v1/Reports/Root | Get user&#39;s root folder (without parents)
[**ReportFoldersMoveFolder**](ReportsApi.md#ReportFoldersMoveFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Move/{folderId} | Move folder to a specified folder
[**ReportFoldersPostFolder**](ReportsApi.md#ReportFoldersPostFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Folder | Create folder
[**ReportFoldersRenameFolder**](ReportsApi.md#ReportFoldersRenameFolder) | **Put** /api/rp/v1/Reports/Folder/{id}/Rename | Rename a folder
[**ReportFoldersUpdateIcon**](ReportsApi.md#ReportFoldersUpdateIcon) | **Put** /api/rp/v1/Reports/Folder/{id}/Icon | Update a folder&#39;s icon
[**ReportFoldersUpdatePermissions**](ReportsApi.md#ReportFoldersUpdatePermissions) | **Post** /api/rp/v1/Reports/{id}/permissions | Update permissions
[**ReportFoldersUpdateTags**](ReportsApi.md#ReportFoldersUpdateTags) | **Put** /api/rp/v1/Reports/Folder/{id}/UpdateTags | Update tags
[**ReportsCopyFile**](ReportsApi.md#ReportsCopyFile) | **Post** /api/rp/v1/Reports/File/{id}/Copy/{folderId} | Copy file to a specified folder
[**ReportsDeleteFile**](ReportsApi.md#ReportsDeleteFile) | **Delete** /api/rp/v1/Reports/File/{id} | Delete specified file
[**ReportsExport**](ReportsApi.md#ReportsExport) | **Post** /api/rp/v1/Reports/File/{id}/Export | Export specified report to a specified format
[**ReportsGetFile**](ReportsApi.md#ReportsGetFile) | **Get** /api/rp/v1/Reports/File/{id} | Get specified file
[**ReportsGetFilesCount**](ReportsApi.md#ReportsGetFilesCount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
[**ReportsGetFilesList**](ReportsApi.md#ReportsGetFilesList) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFiles | Get all files from specified folder. &lt;br /&gt;  User with Get Entity permission can access this method. &lt;br /&gt;  The method will returns minimal infomration about the file: &lt;br /&gt;  id, name, size, editedTime, createdTime, tags, status, statusReason.
[**ReportsGetPermissions**](ReportsApi.md#ReportsGetPermissions) | **Get** /api/rp/v1/Reports/File/{id}/permissions | Get all file permissions
[**ReportsMoveFile**](ReportsApi.md#ReportsMoveFile) | **Post** /api/rp/v1/Reports/File/{id}/Move/{folderId} | Move file to a specified folder
[**ReportsRenameFile**](ReportsApi.md#ReportsRenameFile) | **Put** /api/rp/v1/Reports/File/{id}/Rename | Rename a file
[**ReportsUpdateIcon**](ReportsApi.md#ReportsUpdateIcon) | **Put** /api/rp/v1/Reports/File/{id}/Icon | Update a files&#39;s icon
[**ReportsUpdatePermissions**](ReportsApi.md#ReportsUpdatePermissions) | **Post** /api/rp/v1/Reports/File/{id}/permissions | Update permissions
[**ReportsUpdateTags**](ReportsApi.md#ReportsUpdateTags) | **Put** /api/rp/v1/Reports/File/{id}/UpdateTags | Update tags
[**ReportsUploadFile**](ReportsApi.md#ReportsUploadFile) | **Post** /api/rp/v1/Reports/Folder/{id}/File | Upload a file to the specified folder  !



## ReportFolderAndFileGetCount

> CountVM ReportFolderAndFileGetCount(ctx, id).SearchPattern(searchPattern).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFolderAndFileGetCount(context.Background(), id).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFolderAndFileGetCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFolderAndFileGetCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFolderAndFileGetCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFolderAndFileGetCountRequest struct via the builder pattern


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


## ReportFolderAndFileGetFoldersAndFiles

> FilesVM ReportFolderAndFileGetFoldersAndFiles(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFolderAndFileGetFoldersAndFiles(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFolderAndFileGetFoldersAndFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFolderAndFileGetFoldersAndFiles`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFolderAndFileGetFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFolderAndFileGetFoldersAndFilesRequest struct via the builder pattern


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


## ReportFoldersCopyFolder

> FileVM ReportFoldersCopyFolder(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersCopyFolder(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersCopyFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersCopyFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersCopyFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | moving folder id | 
**folderId** | **string** | destination folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersCopyFolderRequest struct via the builder pattern


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


## ReportFoldersDeleteFolder

> ReportFoldersDeleteFolder(ctx, id).Recursive(recursive).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersDeleteFolder(context.Background(), id).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersDeleteFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFoldersDeleteFolderRequest struct via the builder pattern


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


## ReportFoldersGetBreadcrumbs

> BreadcrumbsVM ReportFoldersGetBreadcrumbs(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersGetBreadcrumbs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersGetBreadcrumbs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersGetBreadcrumbs`: BreadcrumbsVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersGetBreadcrumbs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetBreadcrumbsRequest struct via the builder pattern


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


## ReportFoldersGetFolder

> FileVM ReportFoldersGetFolder(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersGetFolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersGetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersGetFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersGetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetFolderRequest struct via the builder pattern


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


## ReportFoldersGetFolders

> FilesVM ReportFoldersGetFolders(ctx, id).Skip(skip).Take(take).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersGetFolders(context.Background(), id).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersGetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersGetFolders`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersGetFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetFoldersRequest struct via the builder pattern


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


## ReportFoldersGetFoldersCount

> CountVM ReportFoldersGetFoldersCount(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersGetFoldersCount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersGetFoldersCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersGetFoldersCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersGetFoldersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetFoldersCountRequest struct via the builder pattern


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


## ReportFoldersGetPermissions

> FilePermissionsVM ReportFoldersGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetPermissionsRequest struct via the builder pattern


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


## ReportFoldersGetRootFolder

> FileVM ReportFoldersGetRootFolder(ctx).SubscriptionId(subscriptionId).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersGetRootFolder(context.Background()).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersGetRootFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersGetRootFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersGetRootFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetRootFolderRequest struct via the builder pattern


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


## ReportFoldersMoveFolder

> FileVM ReportFoldersMoveFolder(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersMoveFolder(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersMoveFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersMoveFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersMoveFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | moving folder id | 
**folderId** | **string** | destination folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersMoveFolderRequest struct via the builder pattern


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


## ReportFoldersPostFolder

> FileVM ReportFoldersPostFolder(ctx, id).ReportFolderCreateVM(reportFolderCreateVM).Execute()

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
    reportFolderCreateVM := *openapiclient.NewReportFolderCreateVM() // ReportFolderCreateVM | create VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ReportFoldersPostFolder(context.Background(), id).ReportFolderCreateVM(reportFolderCreateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersPostFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersPostFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersPostFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of parent folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersPostFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportFolderCreateVM** | [**ReportFolderCreateVM**](ReportFolderCreateVM.md) | create VM | 

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


## ReportFoldersRenameFolder

> FileVM ReportFoldersRenameFolder(ctx, id).FolderRenameVM(folderRenameVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersRenameFolder(context.Background(), id).FolderRenameVM(folderRenameVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersRenameFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersRenameFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersRenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersRenameFolderRequest struct via the builder pattern


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


## ReportFoldersUpdateIcon

> FileVM ReportFoldersUpdateIcon(ctx, id).FolderIconVM(folderIconVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersUpdateIcon(context.Background(), id).FolderIconVM(folderIconVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersUpdateIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersUpdateIcon`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersUpdateIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersUpdateIconRequest struct via the builder pattern


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


## ReportFoldersUpdatePermissions

> ReportFoldersUpdatePermissions(ctx, id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFoldersUpdatePermissionsRequest struct via the builder pattern


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


## ReportFoldersUpdateTags

> FileVM ReportFoldersUpdateTags(ctx, id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportFoldersUpdateTags(context.Background(), id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportFoldersUpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportFoldersUpdateTags`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportFoldersUpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersUpdateTagsRequest struct via the builder pattern


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


## ReportsCopyFile

> ReportVM ReportsCopyFile(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsCopyFile(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsCopyFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsCopyFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsCopyFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**folderId** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsCopyFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsDeleteFile

> ReportsDeleteFile(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsDeleteFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsDeleteFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportsDeleteFileRequest struct via the builder pattern


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


## ReportsExport

> ExportVM ReportsExport(ctx, id).ExportReportVM(exportReportVM).Execute()

Export specified report to a specified format



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
    id := "id_example" // string | report id
    exportReportVM := *openapiclient.NewExportReportVM() // ExportReportVM | export parameters (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ReportsExport(context.Background(), id).ExportReportVM(exportReportVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsExport`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | report id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportReportVM** | [**ExportReportVM**](ExportReportVM.md) | export parameters | 

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


## ReportsGetFile

> ReportVM ReportsGetFile(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsGetFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGetFilesCount

> CountVM ReportsGetFilesCount(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsGetFilesCount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsGetFilesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetFilesCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsGetFilesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetFilesCountRequest struct via the builder pattern


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


## ReportsGetFilesList

> ReportsVM ReportsGetFilesList(ctx, id).Skip(skip).Take(take).SearchPattern(searchPattern).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsGetFilesList(context.Background(), id).Skip(skip).Take(take).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsGetFilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetFilesList`: ReportsVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsGetFilesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | number of files, that have to be skipped | [default to 0]
 **take** | **int32** | number of files, that have to be returned | [default to 10]
 **searchPattern** | **string** |  | 

### Return type

[**ReportsVM**](ReportsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsGetPermissions

> FilePermissionsVM ReportsGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetPermissionsRequest struct via the builder pattern


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


## ReportsMoveFile

> ReportVM ReportsMoveFile(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsMoveFile(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsMoveFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsMoveFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsMoveFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**folderId** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsMoveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsRenameFile

> ReportVM ReportsRenameFile(ctx, id).FileRenameVM(fileRenameVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsRenameFile(context.Background(), id).FileRenameVM(fileRenameVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsRenameFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsRenameFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsRenameFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsRenameFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileRenameVM** | [**FileRenameVM**](FileRenameVM.md) |  | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUpdateIcon

> ReportVM ReportsUpdateIcon(ctx, id).FileIconVM(fileIconVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsUpdateIcon(context.Background(), id).FileIconVM(fileIconVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsUpdateIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsUpdateIcon`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsUpdateIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUpdateIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileIconVM** | [**FileIconVM**](FileIconVM.md) |  | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUpdatePermissions

> ReportsUpdatePermissions(ctx, id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportsUpdatePermissionsRequest struct via the builder pattern


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


## ReportsUpdateTags

> ReportVM ReportsUpdateTags(ctx, id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()

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
    resp, r, err := api_client.ReportsApi.ReportsUpdateTags(context.Background(), id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsUpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsUpdateTags`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsUpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileTagsUpdateVM** | [**FileTagsUpdateVM**](FileTagsUpdateVM.md) |  | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUploadFile

> ReportVM ReportsUploadFile(ctx, id).ReportCreateVM(reportCreateVM).Execute()

Upload a file to the specified folder  !



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
    reportCreateVM := *openapiclient.NewReportCreateVM() // ReportCreateVM | file's view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ReportsApi.ReportsUploadFile(context.Background(), id).ReportCreateVM(reportCreateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsApi.ReportsUploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsUploadFile`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `ReportsApi.ReportsUploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportCreateVM** | [**ReportCreateVM**](ReportCreateVM.md) | file&#39;s view model | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

