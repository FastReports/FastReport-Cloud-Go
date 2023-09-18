# \TemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateFolderAndFileClearRecycleBin**](TemplatesApi.md#TemplateFolderAndFileClearRecycleBin) | **Delete** /api/rp/v1/Templates/{subscriptionId}/ClearRecycleBin | Delete all folders and files from recycle bin
[**TemplateFolderAndFileDeleteFiles**](TemplatesApi.md#TemplateFolderAndFileDeleteFiles) | **Post** /api/rp/v1/Templates/{subscriptionId}/DeleteFiles | Delete folders and files
[**TemplateFolderAndFileGetCount**](TemplatesApi.md#TemplateFolderAndFileGetCount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
[**TemplateFolderAndFileGetFoldersAndFiles**](TemplatesApi.md#TemplateFolderAndFileGetFoldersAndFiles) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
[**TemplateFolderAndFileGetRecycleBinFoldersAndFiles**](TemplatesApi.md#TemplateFolderAndFileGetRecycleBinFoldersAndFiles) | **Get** /api/rp/v1/Templates/{subscriptionId}/ListRecycleBinFolderAndFiles | Get all folders and files from recycle bin
[**TemplateFolderAndFileRecoverAllFromRecycleBin**](TemplatesApi.md#TemplateFolderAndFileRecoverAllFromRecycleBin) | **Post** /api/rp/v1/Templates/{subscriptionId}/RecoverRecycleBin | Recover all folders and files from recycle bin
[**TemplateFoldersCalculateFolderSize**](TemplatesApi.md#TemplateFoldersCalculateFolderSize) | **Get** /api/rp/v1/Templates/Folder/{id}/size | Get specified folder, calculate it&#39;s size
[**TemplateFoldersCopyFolder**](TemplatesApi.md#TemplateFoldersCopyFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
[**TemplateFoldersDeleteFolder**](TemplatesApi.md#TemplateFoldersDeleteFolder) | **Delete** /api/rp/v1/Templates/Folder/{id} | Delete specified folder
[**TemplateFoldersExport**](TemplatesApi.md#TemplateFoldersExport) | **Post** /api/rp/v1/Templates/Folder/{id}/Export | Export specified template folder to a specified format
[**TemplateFoldersGetBreadcrumbs**](TemplatesApi.md#TemplateFoldersGetBreadcrumbs) | **Get** /api/rp/v1/Templates/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
[**TemplateFoldersGetFolder**](TemplatesApi.md#TemplateFoldersGetFolder) | **Get** /api/rp/v1/Templates/Folder/{id} | Get specified folder
[**TemplateFoldersGetFolders**](TemplatesApi.md#TemplateFoldersGetFolders) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFolders | Get all folders from specified folder
[**TemplateFoldersGetFoldersCount**](TemplatesApi.md#TemplateFoldersGetFoldersCount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
[**TemplateFoldersGetOrCreate**](TemplatesApi.md#TemplateFoldersGetOrCreate) | **Get** /api/rp/v1/Templates/Folder/getOrCreate | Get specified folder
[**TemplateFoldersGetPermissions**](TemplatesApi.md#TemplateFoldersGetPermissions) | **Get** /api/rp/v1/Templates/Folder/{id}/permissions | Get all folder permissions
[**TemplateFoldersGetRootFolder**](TemplatesApi.md#TemplateFoldersGetRootFolder) | **Get** /api/rp/v1/Templates/Root | Get user&#39;s root folder (without parents)
[**TemplateFoldersMoveFolder**](TemplatesApi.md#TemplateFoldersMoveFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Move/{folderId} | Move folder to a specified folder
[**TemplateFoldersMoveFolderToBin**](TemplatesApi.md#TemplateFoldersMoveFolderToBin) | **Delete** /api/rp/v1/Templates/Folder/{id}/ToBin | Move specified folder to recycle bin
[**TemplateFoldersPostFolder**](TemplatesApi.md#TemplateFoldersPostFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Folder | Create folder
[**TemplateFoldersPrepare**](TemplatesApi.md#TemplateFoldersPrepare) | **Post** /api/rp/v1/Templates/Folder/{id}/Prepare | Prepare specified template folder to report folder
[**TemplateFoldersRecoverFolder**](TemplatesApi.md#TemplateFoldersRecoverFolder) | **Post** /api/rp/v1/Templates/Folder/{id}/Recover | Recover specified folder
[**TemplateFoldersRenameFolder**](TemplatesApi.md#TemplateFoldersRenameFolder) | **Put** /api/rp/v1/Templates/Folder/{id}/Rename | Rename a folder
[**TemplateFoldersUpdateIcon**](TemplatesApi.md#TemplateFoldersUpdateIcon) | **Put** /api/rp/v1/Templates/Folder/{id}/Icon | Update a folder&#39;s icon
[**TemplateFoldersUpdatePermissions**](TemplatesApi.md#TemplateFoldersUpdatePermissions) | **Post** /api/rp/v1/Templates/{id}/permissions | Update permissions
[**TemplateFoldersUpdateTags**](TemplatesApi.md#TemplateFoldersUpdateTags) | **Put** /api/rp/v1/Templates/Folder/{id}/UpdateTags | Update tags
[**TemplatesCopyFile**](TemplatesApi.md#TemplatesCopyFile) | **Post** /api/rp/v1/Templates/File/{id}/Copy/{folderId} | Copy file to a specified folder
[**TemplatesDeleteFile**](TemplatesApi.md#TemplatesDeleteFile) | **Delete** /api/rp/v1/Templates/File/{id} | Delete specified file
[**TemplatesExport**](TemplatesApi.md#TemplatesExport) | **Post** /api/rp/v1/Templates/File/{id}/Export | Export specified report template to a specified format
[**TemplatesGetFile**](TemplatesApi.md#TemplatesGetFile) | **Get** /api/rp/v1/Templates/File/{id} | Get specified file
[**TemplatesGetFileHistory**](TemplatesApi.md#TemplatesGetFileHistory) | **Get** /api/rp/v1/Templates/File/{id}/History | Returns list of actions, performed on this file
[**TemplatesGetFilesCount**](TemplatesApi.md#TemplatesGetFilesCount) | **Get** /api/rp/v1/Templates/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
[**TemplatesGetFilesList**](TemplatesApi.md#TemplatesGetFilesList) | **Get** /api/rp/v1/Templates/Folder/{id}/ListFiles | Get all files from specified folder. &lt;br /&gt;  User with Get Entity permission can access this method. &lt;br /&gt;  The method will returns minimal infomration about the file: &lt;br /&gt;  id, name, size, editedTime, createdTime, tags, status, statusReason.
[**TemplatesGetPermissions**](TemplatesApi.md#TemplatesGetPermissions) | **Get** /api/rp/v1/Templates/File/{id}/permissions | Get all file permissions
[**TemplatesMoveFile**](TemplatesApi.md#TemplatesMoveFile) | **Post** /api/rp/v1/Templates/File/{id}/Move/{folderId} | Move file to a specified folder
[**TemplatesMoveFileToBin**](TemplatesApi.md#TemplatesMoveFileToBin) | **Delete** /api/rp/v1/Templates/File/{id}/ToBin | Move specified file to recycle bin
[**TemplatesPrepare**](TemplatesApi.md#TemplatesPrepare) | **Post** /api/rp/v1/Templates/File/{id}/Prepare | Prepare specified template to report
[**TemplatesRecoverFile**](TemplatesApi.md#TemplatesRecoverFile) | **Post** /api/rp/v1/Templates/File/{id}/Recover | Recover specified file from bin
[**TemplatesRenameFile**](TemplatesApi.md#TemplatesRenameFile) | **Put** /api/rp/v1/Templates/File/{id}/Rename | Rename a file
[**TemplatesStaticPreview**](TemplatesApi.md#TemplatesStaticPreview) | **Post** /api/rp/v1/Templates/File/{id}/StaticPreview | Make preview for the report.  Generate a new or return exist prepared svg files.  If template was changed will be returned a new.  Pass the &#x60;&#x60; parameter to check prepared timestamp
[**TemplatesUpdateContent**](TemplatesApi.md#TemplatesUpdateContent) | **Put** /api/rp/v1/Templates/File/{id}/Content | Updates contnet of the template
[**TemplatesUpdateIcon**](TemplatesApi.md#TemplatesUpdateIcon) | **Put** /api/rp/v1/Templates/File/{id}/Icon | Update a files&#39;s icon
[**TemplatesUpdatePermissions**](TemplatesApi.md#TemplatesUpdatePermissions) | **Post** /api/rp/v1/Templates/File/{id}/permissions | Update permissions
[**TemplatesUpdateTags**](TemplatesApi.md#TemplatesUpdateTags) | **Put** /api/rp/v1/Templates/File/{id}/UpdateTags | Update tags
[**TemplatesUploadFile**](TemplatesApi.md#TemplatesUploadFile) | **Post** /api/rp/v1/Templates/Folder/{id}/File | Upload a file to the specified folder  !



## TemplateFolderAndFileClearRecycleBin

> TemplateFolderAndFileClearRecycleBin(ctx, subscriptionId).Execute()

Delete all folders and files from recycle bin



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
    subscriptionId := "subscriptionId_example" // string | subscription id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFolderAndFileClearRecycleBin(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFolderAndFileClearRecycleBin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFolderAndFileClearRecycleBinRequest struct via the builder pattern


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


## TemplateFolderAndFileDeleteFiles

> TemplateFolderAndFileDeleteFiles(ctx, subscriptionId).SelectedFilesForDeletingVM(selectedFilesForDeletingVM).Execute()

Delete folders and files



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
    subscriptionId := "subscriptionId_example" // string | id of current subscription
    selectedFilesForDeletingVM := *openapiclient.NewSelectedFilesForDeletingVM() // SelectedFilesForDeletingVM | VM with files' ids and params of their destination (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFolderAndFileDeleteFiles(context.Background(), subscriptionId).SelectedFilesForDeletingVM(selectedFilesForDeletingVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFolderAndFileDeleteFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | id of current subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFolderAndFileDeleteFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedFilesForDeletingVM** | [**SelectedFilesForDeletingVM**](SelectedFilesForDeletingVM.md) | VM with files&#39; ids and params of their destination | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateFolderAndFileGetCount

> CountVM TemplateFolderAndFileGetCount(ctx, id).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

Get count of files and folders what contains in a specified folder



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
    id := "id_example" // string | folder id
    searchPattern := "searchPattern_example" // string | string, that must be incuded in file or folder name to be counted <br />              (leave undefined to count all files and folders) (optional)
    useRegex := true // bool | set this to true if you want to use regular expression to search (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFolderAndFileGetCount(context.Background(), id).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
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
 **useRegex** | **bool** | set this to true if you want to use regular expression to search | [default to false]

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

> FilesVM TemplateFolderAndFileGetFoldersAndFiles(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

Get all folders and files from specified folder



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
    id := "id_example" // string | folder id
    skip := int32(56) // int32 | number of folder and files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of folder and files, that have to be returned (optional) (default to 10)
    orderBy := openapiclient.FileSorting("None") // FileSorting | indicates a field to sort by (optional)
    desc := true // bool | indicates if sorting is descending (optional) (default to false)
    searchPattern := "searchPattern_example" // string |  (optional) (default to "")
    useRegex := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFolderAndFileGetFoldersAndFiles(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
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
 **useRegex** | **bool** |  | [default to false]

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


## TemplateFolderAndFileGetRecycleBinFoldersAndFiles

> FilesVM TemplateFolderAndFileGetRecycleBinFoldersAndFiles(ctx, subscriptionId).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

Get all folders and files from recycle bin



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
    subscriptionId := "subscriptionId_example" // string | subscription id
    skip := int32(56) // int32 | number of folder and files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of folder and files, that have to be returned (optional) (default to 10)
    orderBy := openapiclient.FileSorting("None") // FileSorting | indicates a field to sort by (optional)
    desc := true // bool | indicates if sorting is descending (optional) (default to false)
    searchPattern := "searchPattern_example" // string |  (optional) (default to "")
    useRegex := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFolderAndFileGetRecycleBinFoldersAndFiles(context.Background(), subscriptionId).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFolderAndFileGetRecycleBinFoldersAndFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFolderAndFileGetRecycleBinFoldersAndFiles`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFolderAndFileGetRecycleBinFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFolderAndFileGetRecycleBinFoldersAndFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | number of folder and files, that have to be skipped | [default to 0]
 **take** | **int32** | number of folder and files, that have to be returned | [default to 10]
 **orderBy** | [**FileSorting**](FileSorting.md) | indicates a field to sort by | 
 **desc** | **bool** | indicates if sorting is descending | [default to false]
 **searchPattern** | **string** |  | [default to &quot;&quot;]
 **useRegex** | **bool** |  | [default to false]

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


## TemplateFolderAndFileRecoverAllFromRecycleBin

> TemplateFolderAndFileRecoverAllFromRecycleBin(ctx, subscriptionId).Execute()

Recover all folders and files from recycle bin



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
    subscriptionId := "subscriptionId_example" // string | subscription id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFolderAndFileRecoverAllFromRecycleBin(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFolderAndFileRecoverAllFromRecycleBin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFolderAndFileRecoverAllFromRecycleBinRequest struct via the builder pattern


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


## TemplateFoldersCalculateFolderSize

> FolderSizeVM TemplateFoldersCalculateFolderSize(ctx, id).Execute()

Get specified folder, calculate it's size



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
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersCalculateFolderSize(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersCalculateFolderSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersCalculateFolderSize`: FolderSizeVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersCalculateFolderSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersCalculateFolderSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderSizeVM**](FolderSizeVM.md)

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | moving folder id
    folderId := "folderId_example" // string | destination folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersCopyFolder(context.Background(), id, folderId).Execute()
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

> TemplateFoldersDeleteFolder(ctx, id).Execute()

Delete specified folder



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
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFoldersDeleteFolder(context.Background(), id).Execute()
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


## TemplateFoldersExport

> FileVM TemplateFoldersExport(ctx, id).ExportTemplateVM(exportTemplateVM).Execute()

Export specified template folder to a specified format



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
    id := "id_example" // string | template folder id
    exportTemplateVM := *openapiclient.NewExportTemplateVM() // ExportTemplateVM | export parameters (string only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersExport(context.Background(), id).ExportTemplateVM(exportTemplateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersExport`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportTemplateVM** | [**ExportTemplateVM**](ExportTemplateVM.md) | export parameters (string only) | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetBreadcrumbs(context.Background(), id).Execute()
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetFolder(context.Background(), id).Execute()
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

> FilesVM TemplateFoldersGetFolders(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

Get all folders from specified folder



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
    id := "id_example" // string | folder id
    skip := int32(56) // int32 | number of files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of files, that have to be returned (optional) (default to 10)
    orderBy := openapiclient.FileSorting("None") // FileSorting |  (optional)
    desc := true // bool |  (optional) (default to false)
    searchPattern := "searchPattern_example" // string |  (optional) (default to "")
    useRegex := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetFolders(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
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
 **orderBy** | [**FileSorting**](FileSorting.md) |  | 
 **desc** | **bool** |  | [default to false]
 **searchPattern** | **string** |  | [default to &quot;&quot;]
 **useRegex** | **bool** |  | [default to false]

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetFoldersCount(context.Background(), id).Execute()
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


## TemplateFoldersGetOrCreate

> FileVM TemplateFoldersGetOrCreate(ctx).Name(name).SubscriptionId(subscriptionId).ParentId(parentId).Execute()

Get specified folder



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
    name := "name_example" // string | folder name (optional)
    subscriptionId := "subscriptionId_example" // string | subscriptionId (optional)
    parentId := "parentId_example" // string | parent folder id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetOrCreate(context.Background()).Name(name).SubscriptionId(subscriptionId).ParentId(parentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersGetOrCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersGetOrCreate`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersGetOrCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersGetOrCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | folder name | 
 **subscriptionId** | **string** | subscriptionId | 
 **parentId** | **string** | parent folder id | 

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetPermissions(context.Background(), id).Execute()
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    subscriptionId := "subscriptionId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersGetRootFolder(context.Background()).SubscriptionId(subscriptionId).Execute()
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | moving folder id
    folderId := "folderId_example" // string | destination folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersMoveFolder(context.Background(), id, folderId).Execute()
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


## TemplateFoldersMoveFolderToBin

> TemplateFoldersMoveFolderToBin(ctx, id).Execute()

Move specified folder to recycle bin



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
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFoldersMoveFolderToBin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersMoveFolderToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplateFoldersMoveFolderToBinRequest struct via the builder pattern


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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of parent folder id
    templateFolderCreateVM := *openapiclient.NewTemplateFolderCreateVM() // TemplateFolderCreateVM | create VM (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersPostFolder(context.Background(), id).TemplateFolderCreateVM(templateFolderCreateVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateFoldersPrepare

> FileVM TemplateFoldersPrepare(ctx, id).PrepareTemplateVM(prepareTemplateVM).Execute()

Prepare specified template folder to report folder



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
    id := "id_example" // string | template id
    prepareTemplateVM := *openapiclient.NewPrepareTemplateVM() // PrepareTemplateVM | Template folder prepare view model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersPrepare(context.Background(), id).PrepareTemplateVM(prepareTemplateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersPrepare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplateFoldersPrepare`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplateFoldersPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplateFoldersPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prepareTemplateVM** | [**PrepareTemplateVM**](PrepareTemplateVM.md) | Template folder prepare view model | 

### Return type

[**FileVM**](FileVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplateFoldersRecoverFolder

> TemplateFoldersRecoverFolder(ctx, id).RecoveryPath(recoveryPath).Execute()

Recover specified folder



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
    id := "id_example" // string | folder id
    recoveryPath := "recoveryPath_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFoldersRecoverFolder(context.Background(), id).RecoveryPath(recoveryPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplateFoldersRecoverFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplateFoldersRecoverFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recoveryPath** | **string** |  | 

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    folderRenameVM := *openapiclient.NewFolderRenameVM("Name_example") // FolderRenameVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersRenameFolder(context.Background(), id).FolderRenameVM(folderRenameVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of folder
    folderIconVM := *openapiclient.NewFolderIconVM(string(123)) // FolderIconVM | Update icon model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersUpdateIcon(context.Background(), id).FolderIconVM(folderIconVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissions(), openapiclient.FileAdministrate(0)) // UpdateFilePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplateFoldersUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    folderTagsUpdateVM := *openapiclient.NewFolderTagsUpdateVM() // FolderTagsUpdateVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplateFoldersUpdateTags(context.Background(), id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | file id
    folderId := "folderId_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesCopyFile(context.Background(), id, folderId).Execute()
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | file id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplatesDeleteFile(context.Background(), id).Execute()
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | report id
    exportTemplateVM := *openapiclient.NewExportTemplateVM() // ExportTemplateVM | export parameters (string only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesExport(context.Background(), id).ExportTemplateVM(exportTemplateVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | file id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetFile(context.Background(), id).Execute()
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


## TemplatesGetFileHistory

> AuditActionsVM TemplatesGetFileHistory(ctx, id).Skip(skip).Take(take).Execute()

Returns list of actions, performed on this file

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
    skip := int32(56) // int32 |  (optional) (default to 0)
    take := int32(56) // int32 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetFileHistory(context.Background(), id).Skip(skip).Take(take).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGetFileHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesGetFileHistory`: AuditActionsVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGetFileHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesGetFileHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** |  | [default to 0]
 **take** | **int32** |  | [default to 10]

### Return type

[**AuditActionsVM**](AuditActionsVM.md)

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetFilesCount(context.Background(), id).Execute()
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

> TemplatesVM TemplatesGetFilesList(ctx, id).Skip(skip).Take(take).SearchPattern(searchPattern).OrderBy(orderBy).Desc(desc).UseRegex(useRegex).Execute()

Get all files from specified folder. <br />  User with Get Entity permission can access this method. <br />  The method will returns minimal infomration about the file: <br />  id, name, size, editedTime, createdTime, tags, status, statusReason.

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
    id := "id_example" // string | folder id
    skip := int32(56) // int32 | number of files, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of files, that have to be returned (optional) (default to 10)
    searchPattern := "searchPattern_example" // string |  (optional)
    orderBy := openapiclient.FileSorting("None") // FileSorting |  (optional)
    desc := true // bool |  (optional) (default to false)
    useRegex := true // bool |  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetFilesList(context.Background(), id).Skip(skip).Take(take).SearchPattern(searchPattern).OrderBy(orderBy).Desc(desc).UseRegex(useRegex).Execute()
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
 **orderBy** | [**FileSorting**](FileSorting.md) |  | 
 **desc** | **bool** |  | [default to false]
 **useRegex** | **bool** |  | [default to false]

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesGetPermissions(context.Background(), id).Execute()
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | file id
    folderId := "folderId_example" // string | folder id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesMoveFile(context.Background(), id, folderId).Execute()
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


## TemplatesMoveFileToBin

> TemplatesMoveFileToBin(ctx, id).Execute()

Move specified file to recycle bin



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
    id := "id_example" // string | file id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplatesMoveFileToBin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesMoveFileToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplatesMoveFileToBinRequest struct via the builder pattern


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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | template id
    prepareTemplateVM := *openapiclient.NewPrepareTemplateVM() // PrepareTemplateVM | Template prepare view model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesPrepare(context.Background(), id).PrepareTemplateVM(prepareTemplateVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesRecoverFile

> TemplatesRecoverFile(ctx, id).RecoveryPath(recoveryPath).Execute()

Recover specified file from bin



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
    id := "id_example" // string | file id
    recoveryPath := "recoveryPath_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplatesRecoverFile(context.Background(), id).RecoveryPath(recoveryPath).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesRecoverFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTemplatesRecoverFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recoveryPath** | **string** |  | 

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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    fileRenameVM := *openapiclient.NewFileRenameVM() // FileRenameVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesRenameFile(context.Background(), id).FileRenameVM(fileRenameVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesStaticPreview

> ExportVM TemplatesStaticPreview(ctx, id).PreviewTemplateVM(previewTemplateVM).Execute()

Make preview for the report.  Generate a new or return exist prepared svg files.  If template was changed will be returned a new.  Pass the `` parameter to check prepared timestamp

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
    id := "id_example" // string | template id
    previewTemplateVM := *openapiclient.NewPreviewTemplateVM() // PreviewTemplateVM | Model with parameters (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesStaticPreview(context.Background(), id).PreviewTemplateVM(previewTemplateVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesStaticPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TemplatesStaticPreview`: ExportVM
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesStaticPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesStaticPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewTemplateVM** | [**PreviewTemplateVM**](PreviewTemplateVM.md) | Model with parameters | 

### Return type

[**ExportVM**](ExportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TemplatesUpdateContent

> TemplatesUpdateContent(ctx, id).UpdateFileContentVM(updateFileContentVM).Execute()

Updates contnet of the template

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
    id := "id_example" // string | template id
    updateFileContentVM := *openapiclient.NewUpdateFileContentVM(string(123)) // UpdateFileContentVM | VM with only byte[] with new content (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplatesUpdateContent(context.Background(), id).UpdateFileContentVM(updateFileContentVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesUpdateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTemplatesUpdateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFileContentVM** | [**UpdateFileContentVM**](UpdateFileContentVM.md) | VM with only byte[] with new content | 

### Return type

 (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    fileIconVM := *openapiclient.NewFileIconVM() // FileIconVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesUpdateIcon(context.Background(), id).FileIconVM(fileIconVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissions(), openapiclient.FileAdministrate(0)) // UpdateFilePermissionsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TemplatesApi.TemplatesUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | 
    fileTagsUpdateVM := *openapiclient.NewFileTagsUpdateVM() // FileTagsUpdateVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesUpdateTags(context.Background(), id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
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
    openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
    id := "id_example" // string | Identifier of folder
    templateCreateVM := *openapiclient.NewTemplateCreateVM() // TemplateCreateVM | file's view model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TemplatesApi.TemplatesUploadFile(context.Background(), id).TemplateCreateVM(templateCreateVM).Execute()
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

