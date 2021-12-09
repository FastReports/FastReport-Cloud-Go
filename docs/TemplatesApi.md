# \TemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateFolderAndFileGetCount**](TemplatesApi.md#TemplateFolderAndFileGetCount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
[**TemplateFolderAndFileGetFoldersAndFiles**](TemplatesApi.md#TemplateFolderAndFileGetFoldersAndFiles) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
[**TemplateFoldersCopyFolder**](TemplatesApi.md#TemplateFoldersCopyFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
[**TemplateFoldersDeleteFolder**](TemplatesApi.md#TemplateFoldersDeleteFolder) | **Delete** /api/rp/v1/Templates/Folder/{id} | Delete specified folder
[**TemplateFoldersGetBreadcrumbs**](TemplatesApi.md#TemplateFoldersGetBreadcrumbs) | **Get** /api/rp/v1/Templates/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
[**TemplateFoldersGetFolder**](TemplatesApi.md#TemplateFoldersGetFolder) | **Get** /api/rp/v1/Templates/Folder/{id} | Get specified folder
[**TemplateFoldersGetFolders**](TemplatesApi.md#TemplateFoldersGetFolders) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFolders | Get all folders from specified folder
[**TemplateFoldersGetFoldersCount**](TemplatesApi.md#TemplateFoldersGetFoldersCount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
[**TemplateFoldersGetPermissions**](TemplatesApi.md#TemplateFoldersGetPermissions) | **Get** /api/rp/v1/Templates/Folder/{id}/permissions | Get all folder permissions
[**TemplateFoldersGetRootFolder**](TemplatesApi.md#TemplateFoldersGetRootFolder) | **Get** /api/rp/v1/Templates/Root | Get user&#39;s root folder (without parents)
[**TemplateFoldersMoveFolder**](TemplatesApi.md#TemplateFoldersMoveFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Move/{folderId} | Move folder to a specified folder
[**TemplateFoldersPostFolder**](TemplatesApi.md#TemplateFoldersPostFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Folder | Create folder
[**TemplateFoldersRenameFolder**](TemplatesApi.md#TemplateFoldersRenameFolder) | **Put** /api/rp/v1/Templates/Folder/{id}/Rename | Rename a folder
[**TemplateFoldersUpdateIcon**](TemplatesApi.md#TemplateFoldersUpdateIcon) | **Put** /api/rp/v1/Templates/Folder/{id}/Icon | Update a folder&#39;s icon
[**TemplateFoldersUpdatePermissions**](TemplatesApi.md#TemplateFoldersUpdatePermissions) | **Post** /api/rp/v1/Templates/{id}/permissions | Update permissions
[**TemplateFoldersUpdateTags**](TemplatesApi.md#TemplateFoldersUpdateTags) | **Put** /api/rp/v1/Templates/Folder/{id}/UpdateTags | Update tags
[**TemplatesCopyFile**](TemplatesApi.md#TemplatesCopyFile) | **Post** /api/rp/v1/Templates/File/{id}/Copy/{folderId} | Copy file to a specified folder
[**TemplatesDeleteFile**](TemplatesApi.md#TemplatesDeleteFile) | **Delete** /api/rp/v1/Templates/File/{id} | Delete specified file
[**TemplatesExport**](TemplatesApi.md#TemplatesExport) | **Post** /api/rp/v1/Templates/File/{id}/Export | Export specified report template to a specified format
[**TemplatesGetFile**](TemplatesApi.md#TemplatesGetFile) | **Get** /api/rp/v1/Templates/File/{id} | Get specified file
[**TemplatesGetFilesCount**](TemplatesApi.md#TemplatesGetFilesCount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
[**TemplatesGetFilesList**](TemplatesApi.md#TemplatesGetFilesList) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFiles | Get all files from specified folder. &lt;br /&gt;  User with Get Entity permission can access this method. &lt;br /&gt;  The method will returns minimal infomration about the file: &lt;br /&gt;  id, name, size, editedTime, createdTime, tags, status, statusReason.
[**TemplatesGetPermissions**](TemplatesApi.md#TemplatesGetPermissions) | **Get** /api/rp/v1/Templates/File/{id}/permissions | Get all file permissions
[**TemplatesMoveFile**](TemplatesApi.md#TemplatesMoveFile) | **Post** /api/rp/v1/Templates/File/{id}/Move/{folderId} | Move file to a specified folder
[**TemplatesPrepare**](TemplatesApi.md#TemplatesPrepare) | **Post** /api/rp/v1/Templates/File/{id}/Prepare | Prepare specified template to report
[**TemplatesRenameFile**](TemplatesApi.md#TemplatesRenameFile) | **Put** /api/rp/v1/Templates/File/{id}/Rename | Rename a file
[**TemplatesUpdateIcon**](TemplatesApi.md#TemplatesUpdateIcon) | **Put** /api/rp/v1/Templates/File/{id}/Icon | Update a files&#39;s icon
[**TemplatesUpdatePermissions**](TemplatesApi.md#TemplatesUpdatePermissions) | **Post** /api/rp/v1/Templates/File/{id}/permissions | Update permissions
[**TemplatesUpdateTags**](TemplatesApi.md#TemplatesUpdateTags) | **Put** /api/rp/v1/Templates/File/{id}/UpdateTags | Update tags
[**TemplatesUploadFile**](TemplatesApi.md#TemplatesUploadFile) | **Post** /api/rp/v1/Templates/Folder/{id}/File | Upload a file to the specified folder  !



## TemplateFolderAndFileGetCount

> CountVM TemplateFolderAndFileGetCount(ctx, id).SearchPattern(searchPattern).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFolderAndFileGetCount(context.Background(), id).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFolderAndFileGetCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFolderAndFileGetCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFolderAndFileGetCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFolderAndFileGetCountRequest struct via the builder pattern


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


## TemplateFolderAndFileGetFoldersAndFiles

> FilesVM TemplateFolderAndFileGetFoldersAndFiles(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFolderAndFileGetFoldersAndFiles(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFolderAndFileGetFoldersAndFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFolderAndFileGetFoldersAndFiles`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFolderAndFileGetFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFolderAndFileGetFoldersAndFilesRequest struct via the builder pattern


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


## TemplateFoldersCopyFolder

> FileVM TemplateFoldersCopyFolder(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersCopyFolder(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersCopyFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersCopyFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersCopyFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | moving folder id | 
**folderId** | **string** | destination folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersCopyFolderRequest struct via the builder pattern


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


## TemplateFoldersDeleteFolder

> TemplateFoldersDeleteFolder(ctx, id).Recursive(recursive).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersDeleteFolder(context.Background(), id).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersDeleteFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplateFoldersDeleteFolderRequest struct via the builder pattern


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


## TemplateFoldersGetBreadcrumbs

> BreadcrumbsVM TemplateFoldersGetBreadcrumbs(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersGetBreadcrumbs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetBreadcrumbs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetBreadcrumbs`: BreadcrumbsVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetBreadcrumbs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetBreadcrumbsRequest struct via the builder pattern


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


## TemplateFoldersGetFolder

> FileVM TemplateFoldersGetFolder(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersGetFolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetFolderRequest struct via the builder pattern


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


## TemplateFoldersGetFolders

> FilesVM TemplateFoldersGetFolders(ctx, id).Skip(skip).Take(take).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersGetFolders(context.Background(), id).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetFolders`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetFolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetFoldersRequest struct via the builder pattern


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


## TemplateFoldersGetFoldersCount

> CountVM TemplateFoldersGetFoldersCount(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersGetFoldersCount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetFoldersCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetFoldersCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetFoldersCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetFoldersCountRequest struct via the builder pattern


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


## TemplateFoldersGetPermissions

> FilePermissionsVM TemplateFoldersGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetPermissionsRequest struct via the builder pattern


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


## TemplateFoldersGetRootFolder

> FileVM TemplateFoldersGetRootFolder(ctx).SubscriptionId(subscriptionId).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersGetRootFolder(context.Background()).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetRootFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetRootFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetRootFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetRootFolderRequest struct via the builder pattern


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


## TemplateFoldersMoveFolder

> FileVM TemplateFoldersMoveFolder(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersMoveFolder(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersMoveFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersMoveFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersMoveFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | moving folder id | 
**folderId** | **string** | destination folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersMoveFolderRequest struct via the builder pattern


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


## TemplateFoldersPostFolder

> FileVM TemplateFoldersPostFolder(ctx, id).TemplateFolderCreateVM(templateFolderCreateVM).Execute()

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
    templateFolderCreateVM := *openapiclient.NewTemplateFolderCreateVM() // TemplateFolderCreateVM | create VM (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.TemplateFoldersPostFolder(context.Background(), id).TemplateFolderCreateVM(templateFolderCreateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersPostFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersPostFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersPostFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of parent folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersPostFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateFolderCreateVM** | [**TemplateFolderCreateVM**](TemplateFolderCreateVM.md) | create VM | 

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


## TemplateFoldersRenameFolder

> FileVM TemplateFoldersRenameFolder(ctx, id).FolderRenameVM(folderRenameVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersRenameFolder(context.Background(), id).FolderRenameVM(folderRenameVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersRenameFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersRenameFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersRenameFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersRenameFolderRequest struct via the builder pattern


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


## TemplateFoldersUpdateIcon

> FileVM TemplateFoldersUpdateIcon(ctx, id).FolderIconVM(folderIconVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersUpdateIcon(context.Background(), id).FolderIconVM(folderIconVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersUpdateIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersUpdateIcon`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersUpdateIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersUpdateIconRequest struct via the builder pattern


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


## TemplateFoldersUpdatePermissions

> TemplateFoldersUpdatePermissions(ctx, id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplateFoldersUpdatePermissionsRequest struct via the builder pattern


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


## TemplateFoldersUpdateTags

> FileVM TemplateFoldersUpdateTags(ctx, id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplateFoldersUpdateTags(context.Background(), id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersUpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersUpdateTags`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersUpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersUpdateTagsRequest struct via the builder pattern


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


## TemplatesCopyFile

> TemplateVM TemplatesCopyFile(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesCopyFile(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesCopyFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesCopyFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesCopyFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**folderId** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesCopyFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesDeleteFile

> TemplatesDeleteFile(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesDeleteFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesDeleteFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplatesDeleteFileRequest struct via the builder pattern


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


## TemplatesExport

> ExportVM TemplatesExport(ctx, id).ExportTemplateVM(exportTemplateVM).Execute()

Export specified report template to a specified format



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
    exportTemplateVM := *openapiclient.NewExportTemplateVM() // ExportTemplateVM | export parameters (string only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.TemplatesExport(context.Background(), id).ExportTemplateVM(exportTemplateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesExport`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | report id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportTemplateVM** | [**ExportTemplateVM**](ExportTemplateVM.md) | export parameters (string only) | 

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


## TemplatesGetFile

> TemplateVM TemplatesGetFile(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesGetFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesGetFilesCount

> CountVM TemplatesGetFilesCount(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesGetFilesCount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetFilesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetFilesCount`: CountVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetFilesCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetFilesCountRequest struct via the builder pattern


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


## TemplatesGetFilesList

> TemplatesVM TemplatesGetFilesList(ctx, id).Skip(skip).Take(take).SearchPattern(searchPattern).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesGetFilesList(context.Background(), id).Skip(skip).Take(take).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetFilesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetFilesList`: TemplatesVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetFilesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetFilesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | number of files, that have to be skipped | [default to 0]
 **take** | **int32** | number of files, that have to be returned | [default to 10]
 **searchPattern** | **string** |  | 

### Return type

[**TemplatesVM**](TemplatesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesGetPermissions

> FilePermissionsVM TemplatesGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetPermissionsRequest struct via the builder pattern


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


## TemplatesMoveFile

> TemplateVM TemplatesMoveFile(ctx, id, folderId).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesMoveFile(context.Background(), id, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesMoveFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesMoveFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesMoveFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**folderId** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesMoveFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesPrepare

> ReportVM TemplatesPrepare(ctx, id).PrepareTemplateVM(prepareTemplateVM).Execute()

Prepare specified template to report



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
    id := "id_example" // string | template id
    prepareTemplateVM := *openapiclient.NewPrepareTemplateVM() // PrepareTemplateVM | Template prepare view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.TemplatesPrepare(context.Background(), id).PrepareTemplateVM(prepareTemplateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesPrepare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesPrepare`: ReportVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prepareTemplateVM** | [**PrepareTemplateVM**](PrepareTemplateVM.md) | Template prepare view model | 

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


## TemplatesRenameFile

> TemplateVM TemplatesRenameFile(ctx, id).FileRenameVM(fileRenameVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesRenameFile(context.Background(), id).FileRenameVM(fileRenameVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesRenameFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesRenameFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesRenameFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesRenameFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileRenameVM** | [**FileRenameVM**](FileRenameVM.md) |  | 

### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUpdateIcon

> TemplateVM TemplatesUpdateIcon(ctx, id).FileIconVM(fileIconVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesUpdateIcon(context.Background(), id).FileIconVM(fileIconVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesUpdateIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesUpdateIcon`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesUpdateIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUpdateIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileIconVM** | [**FileIconVM**](FileIconVM.md) |  | 

### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUpdatePermissions

> TemplatesUpdatePermissions(ctx, id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplatesUpdatePermissionsRequest struct via the builder pattern


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


## TemplatesUpdateTags

> TemplateVM TemplatesUpdateTags(ctx, id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()

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
    resp, r, err := api_client.TemplatesApi.TemplatesUpdateTags(context.Background(), id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesUpdateTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesUpdateTags`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesUpdateTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUpdateTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileTagsUpdateVM** | [**FileTagsUpdateVM**](FileTagsUpdateVM.md) |  | 

### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUploadFile

> TemplateVM TemplatesUploadFile(ctx, id).TemplateCreateVM(templateCreateVM).Execute()

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
    templateCreateVM := *openapiclient.NewTemplateCreateVM() // TemplateCreateVM | file's view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TemplatesApi.TemplatesUploadFile(context.Background(), id).TemplateCreateVM(templateCreateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesUploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesUploadFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesUploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **templateCreateVM** | [**TemplateCreateVM**](TemplateCreateVM.md) | file&#39;s view model | 

### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/_*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

