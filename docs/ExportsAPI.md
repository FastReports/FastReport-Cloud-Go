# \ExportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportFolderAndFileClearRecycleBin**](ExportsAPI.md#ExportFolderAndFileClearRecycleBin) | **Delete** /api/rp/v1/Exports/{subscriptionId}/ClearRecycleBin | Delete all folders and files from recycle bin
[**ExportFolderAndFileCopyFiles**](ExportsAPI.md#ExportFolderAndFileCopyFiles) | **Post** /api/rp/v1/Exports/{subscriptionId}/CopyFiles | Copy folders and files to a specified folder
[**ExportFolderAndFileCountRecycleBinFoldersAndFiles**](ExportsAPI.md#ExportFolderAndFileCountRecycleBinFoldersAndFiles) | **Get** /api/rp/v1/Exports/{subscriptionId}/CountRecycleBinFolderAndFiles | Count all folders and files from recycle bin
[**ExportFolderAndFileDeleteFiles**](ExportsAPI.md#ExportFolderAndFileDeleteFiles) | **Post** /api/rp/v1/Exports/{subscriptionId}/DeleteFiles | Delete folders and files
[**ExportFolderAndFileGetCount**](ExportsAPI.md#ExportFolderAndFileGetCount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
[**ExportFolderAndFileGetFoldersAndFiles**](ExportsAPI.md#ExportFolderAndFileGetFoldersAndFiles) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
[**ExportFolderAndFileGetRecycleBinFoldersAndFiles**](ExportsAPI.md#ExportFolderAndFileGetRecycleBinFoldersAndFiles) | **Get** /api/rp/v1/Exports/{subscriptionId}/ListRecycleBinFolderAndFiles | Get all folders and files from recycle bin
[**ExportFolderAndFileMoveFiles**](ExportsAPI.md#ExportFolderAndFileMoveFiles) | **Post** /api/rp/v1/Exports/{subscriptionId}/MoveFiles | Move folders and files to a specified folder
[**ExportFolderAndFileMoveFilesToBin**](ExportsAPI.md#ExportFolderAndFileMoveFilesToBin) | **Post** /api/rp/v1/Exports/{subscriptionId}/ToBin | Move folders and files to bin
[**ExportFolderAndFileRecoverAllFromRecycleBin**](ExportsAPI.md#ExportFolderAndFileRecoverAllFromRecycleBin) | **Post** /api/rp/v1/Exports/{subscriptionId}/RecoverRecycleBin | Recover all folders and files from recycle bin
[**ExportFolderAndFileRecoverFiles**](ExportsAPI.md#ExportFolderAndFileRecoverFiles) | **Post** /api/rp/v1/Exports/{subscriptionId}/RecoverFiles | Recover folders and files from bin
[**ExportFoldersCalculateFolderSize**](ExportsAPI.md#ExportFoldersCalculateFolderSize) | **Get** /api/rp/v1/Exports/Folder/{id}/size | Get specified folder, calculate it&#39;s size
[**ExportFoldersCopyFolder**](ExportsAPI.md#ExportFoldersCopyFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
[**ExportFoldersDeleteFolder**](ExportsAPI.md#ExportFoldersDeleteFolder) | **Delete** /api/rp/v1/Exports/Folder/{id} | Delete specified folder
[**ExportFoldersGetBreadcrumbs**](ExportsAPI.md#ExportFoldersGetBreadcrumbs) | **Get** /api/rp/v1/Exports/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
[**ExportFoldersGetFolder**](ExportsAPI.md#ExportFoldersGetFolder) | **Get** /api/rp/v1/Exports/Folder/{id} | Get specified folder
[**ExportFoldersGetFolders**](ExportsAPI.md#ExportFoldersGetFolders) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFolders | Get all folders from specified folder
[**ExportFoldersGetFoldersCount**](ExportsAPI.md#ExportFoldersGetFoldersCount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
[**ExportFoldersGetOrCreate**](ExportsAPI.md#ExportFoldersGetOrCreate) | **Get** /api/rp/v1/Exports/Folder/getOrCreate | Get specified folder
[**ExportFoldersGetPermissions**](ExportsAPI.md#ExportFoldersGetPermissions) | **Get** /api/rp/v1/Exports/Folder/{id}/permissions | Get all folder permissions
[**ExportFoldersGetRootFolder**](ExportsAPI.md#ExportFoldersGetRootFolder) | **Get** /api/rp/v1/Exports/Root | Get user&#39;s root folder (without parents)
[**ExportFoldersMoveFolder**](ExportsAPI.md#ExportFoldersMoveFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Move/{folderId} | Move folder to a specified folder
[**ExportFoldersMoveFolderToBin**](ExportsAPI.md#ExportFoldersMoveFolderToBin) | **Delete** /api/rp/v1/Exports/Folder/{id}/ToBin | Move specified folder to recycle bin
[**ExportFoldersPostFolder**](ExportsAPI.md#ExportFoldersPostFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Folder | Create folder
[**ExportFoldersRecoverFolder**](ExportsAPI.md#ExportFoldersRecoverFolder) | **Post** /api/rp/v1/Exports/Folder/{id}/Recover | Recover specified folder
[**ExportFoldersRenameFolder**](ExportsAPI.md#ExportFoldersRenameFolder) | **Put** /api/rp/v1/Exports/Folder/{id}/Rename | Rename a folder
[**ExportFoldersUpdateIcon**](ExportsAPI.md#ExportFoldersUpdateIcon) | **Put** /api/rp/v1/Exports/Folder/{id}/Icon | Update a folder&#39;s icon
[**ExportFoldersUpdatePermissions**](ExportsAPI.md#ExportFoldersUpdatePermissions) | **Post** /api/rp/v1/Exports/{id}/permissions | Update permissions
[**ExportFoldersUpdateTags**](ExportsAPI.md#ExportFoldersUpdateTags) | **Put** /api/rp/v1/Exports/Folder/{id}/UpdateTags | Update tags
[**ExportsCopyFile**](ExportsAPI.md#ExportsCopyFile) | **Post** /api/rp/v1/Exports/File/{id}/Copy/{folderId} | Copy file to a specified folder
[**ExportsCreateSharingKey**](ExportsAPI.md#ExportsCreateSharingKey) | **Post** /api/rp/v1/Exports/File/{id}/sharingKey | Create a new key, that can be used to share access to a file  (You need Administrate.Anon permission to create a new key)
[**ExportsDeleteFile**](ExportsAPI.md#ExportsDeleteFile) | **Delete** /api/rp/v1/Exports/File/{id} | Delete specified file
[**ExportsDeleteSharingKey**](ExportsAPI.md#ExportsDeleteSharingKey) | **Delete** /api/rp/v1/Exports/File/{id}/sharingKey | Deletes a sharing key, making links, that utilizing it no longer work
[**ExportsGetFile**](ExportsAPI.md#ExportsGetFile) | **Get** /api/rp/v1/Exports/File/{id} | Get export by specified id
[**ExportsGetFileHistory**](ExportsAPI.md#ExportsGetFileHistory) | **Get** /api/rp/v1/Exports/File/{id}/History | Returns list of actions, performed on this file
[**ExportsGetFilesCount**](ExportsAPI.md#ExportsGetFilesCount) | **Get** /api/rp/v1/Exports/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
[**ExportsGetFilesList**](ExportsAPI.md#ExportsGetFilesList) | **Get** /api/rp/v1/Exports/Folder/{id}/ListFiles | Get all files from specified folder. &lt;br /&gt;  User with Get Entity permission can access this method. &lt;br /&gt;  The method will returns minimal infomration about the file: &lt;br /&gt;  id, name, size, editedTime, createdTime, tags, status, statusReason.
[**ExportsGetPermissions**](ExportsAPI.md#ExportsGetPermissions) | **Get** /api/rp/v1/Exports/File/{id}/permissions | 
[**ExportsGetSharingKeys**](ExportsAPI.md#ExportsGetSharingKeys) | **Get** /api/rp/v1/Exports/File/{id}/sharingKeys | Returns all sharing keys, associated with the file
[**ExportsMoveFile**](ExportsAPI.md#ExportsMoveFile) | **Post** /api/rp/v1/Exports/File/{id}/Move/{folderId} | Move file to a specified folder
[**ExportsMoveFileToBin**](ExportsAPI.md#ExportsMoveFileToBin) | **Delete** /api/rp/v1/Exports/File/{id}/ToBin | Move specified file to recycle bin
[**ExportsRecoverFile**](ExportsAPI.md#ExportsRecoverFile) | **Post** /api/rp/v1/Exports/File/{id}/Recover | Recover specified file from bin
[**ExportsRenameFile**](ExportsAPI.md#ExportsRenameFile) | **Put** /api/rp/v1/Exports/File/{id}/Rename | Rename a file
[**ExportsUpdateIcon**](ExportsAPI.md#ExportsUpdateIcon) | **Put** /api/rp/v1/Exports/File/{id}/Icon | Update a files&#39;s icon
[**ExportsUpdatePermissions**](ExportsAPI.md#ExportsUpdatePermissions) | **Post** /api/rp/v1/Exports/File/{id}/permissions | Update permissions
[**ExportsUpdateTags**](ExportsAPI.md#ExportsUpdateTags) | **Put** /api/rp/v1/Exports/File/{id}/UpdateTags | Update tags



## ExportFolderAndFileClearRecycleBin

> ExportFolderAndFileClearRecycleBin(ctx, subscriptionId).Execute()

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
	r, err := apiClient.ExportsAPI.ExportFolderAndFileClearRecycleBin(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileClearRecycleBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileClearRecycleBinRequest struct via the builder pattern


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


## ExportFolderAndFileCopyFiles

> ExportFolderAndFileCopyFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

Copy folders and files to a specified folder



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
	selectedFilesVM := *openapiclient.NewSelectedFilesVM("T_example") // SelectedFilesVM | VM with files' ids and params of their destination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportFolderAndFileCopyFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileCopyFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileCopyFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedFilesVM** | [**SelectedFilesVM**](SelectedFilesVM.md) | VM with files&#39; ids and params of their destination | 

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


## ExportFolderAndFileCountRecycleBinFoldersAndFiles

> CountVM ExportFolderAndFileCountRecycleBinFoldersAndFiles(ctx, subscriptionId).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

Count all folders and files from recycle bin



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
	searchPattern := "searchPattern_example" // string |  (optional) (default to "")
	useRegex := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFolderAndFileCountRecycleBinFoldersAndFiles(context.Background(), subscriptionId).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileCountRecycleBinFoldersAndFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFolderAndFileCountRecycleBinFoldersAndFiles`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFolderAndFileCountRecycleBinFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFolderAndFileCountRecycleBinFoldersAndFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchPattern** | **string** |  | [default to &quot;&quot;]
 **useRegex** | **bool** |  | [default to false]

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


## ExportFolderAndFileDeleteFiles

> ExportFolderAndFileDeleteFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

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
	selectedFilesVM := *openapiclient.NewSelectedFilesVM("T_example") // SelectedFilesVM | VM with files' ids and params of their destination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportFolderAndFileDeleteFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileDeleteFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileDeleteFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedFilesVM** | [**SelectedFilesVM**](SelectedFilesVM.md) | VM with files&#39; ids and params of their destination | 

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


## ExportFolderAndFileGetCount

> CountVM ExportFolderAndFileGetCount(ctx, id).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ExportsAPI.ExportFolderAndFileGetCount(context.Background(), id).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileGetCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFolderAndFileGetCount`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFolderAndFileGetCount`: %v\n", resp)
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


## ExportFolderAndFileGetFoldersAndFiles

> FilesVM ExportFolderAndFileGetFoldersAndFiles(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	orderBy := openapiclient.FileSorting("CreatedTime") // FileSorting | indicates a field to sort by (optional)
	desc := true // bool | indicates if sorting is descending (optional) (default to false)
	searchPattern := "searchPattern_example" // string |  (optional) (default to "")
	useRegex := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFolderAndFileGetFoldersAndFiles(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileGetFoldersAndFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFolderAndFileGetFoldersAndFiles`: FilesVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFolderAndFileGetFoldersAndFiles`: %v\n", resp)
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


## ExportFolderAndFileGetRecycleBinFoldersAndFiles

> FilesVM ExportFolderAndFileGetRecycleBinFoldersAndFiles(ctx, subscriptionId).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	orderBy := openapiclient.FileSorting("CreatedTime") // FileSorting | indicates a field to sort by (optional)
	desc := true // bool | indicates if sorting is descending (optional) (default to false)
	searchPattern := "searchPattern_example" // string |  (optional) (default to "")
	useRegex := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFolderAndFileGetRecycleBinFoldersAndFiles(context.Background(), subscriptionId).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileGetRecycleBinFoldersAndFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFolderAndFileGetRecycleBinFoldersAndFiles`: FilesVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFolderAndFileGetRecycleBinFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFolderAndFileGetRecycleBinFoldersAndFilesRequest struct via the builder pattern


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


## ExportFolderAndFileMoveFiles

> ExportFolderAndFileMoveFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

Move folders and files to a specified folder



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
	selectedFilesVM := *openapiclient.NewSelectedFilesVM("T_example") // SelectedFilesVM | VM with files' ids and params of their destination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportFolderAndFileMoveFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileMoveFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileMoveFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedFilesVM** | [**SelectedFilesVM**](SelectedFilesVM.md) | VM with files&#39; ids and params of their destination | 

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


## ExportFolderAndFileMoveFilesToBin

> ExportFolderAndFileMoveFilesToBin(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

Move folders and files to bin



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
	selectedFilesVM := *openapiclient.NewSelectedFilesVM("T_example") // SelectedFilesVM | VM with files' ids and params of their destination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportFolderAndFileMoveFilesToBin(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileMoveFilesToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileMoveFilesToBinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedFilesVM** | [**SelectedFilesVM**](SelectedFilesVM.md) | VM with files&#39; ids and params of their destination | 

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


## ExportFolderAndFileRecoverAllFromRecycleBin

> ExportFolderAndFileRecoverAllFromRecycleBin(ctx, subscriptionId).Execute()

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
	r, err := apiClient.ExportsAPI.ExportFolderAndFileRecoverAllFromRecycleBin(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileRecoverAllFromRecycleBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileRecoverAllFromRecycleBinRequest struct via the builder pattern


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


## ExportFolderAndFileRecoverFiles

> ExportFolderAndFileRecoverFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

Recover folders and files from bin



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
	selectedFilesVM := *openapiclient.NewSelectedFilesVM("T_example") // SelectedFilesVM | VM with files' ids and params of their destination (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportFolderAndFileRecoverFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFolderAndFileRecoverFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFolderAndFileRecoverFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **selectedFilesVM** | [**SelectedFilesVM**](SelectedFilesVM.md) | VM with files&#39; ids and params of their destination | 

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


## ExportFoldersCalculateFolderSize

> FolderSizeVM ExportFoldersCalculateFolderSize(ctx, id).Execute()

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
	resp, r, err := apiClient.ExportsAPI.ExportFoldersCalculateFolderSize(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersCalculateFolderSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersCalculateFolderSize`: FolderSizeVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersCalculateFolderSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersCalculateFolderSizeRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | moving folder id
	folderId := "folderId_example" // string | destination folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersCopyFolder(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersCopyFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersCopyFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersCopyFolder`: %v\n", resp)
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

> ExportFoldersDeleteFolder(ctx, id).Execute()

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
	r, err := apiClient.ExportsAPI.ExportFoldersDeleteFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersDeleteFolder``: %v\n", err)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetBreadcrumbs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetBreadcrumbs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetBreadcrumbs`: BreadcrumbsVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetBreadcrumbs`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetFolder`: %v\n", resp)
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

> FilesVM ExportFoldersGetFolders(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	orderBy := openapiclient.FileSorting("CreatedTime") // FileSorting |  (optional)
	desc := true // bool |  (optional) (default to false)
	searchPattern := "searchPattern_example" // string |  (optional) (default to "")
	useRegex := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetFolders(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetFolders`: FilesVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetFolders`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetFoldersCount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetFoldersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetFoldersCount`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetFoldersCount`: %v\n", resp)
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


## ExportFoldersGetOrCreate

> FileVM ExportFoldersGetOrCreate(ctx).Name(name).SubscriptionId(subscriptionId).ParentId(parentId).Execute()

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
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetOrCreate(context.Background()).Name(name).SubscriptionId(subscriptionId).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetOrCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetOrCreate`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetOrCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFoldersGetOrCreateRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetPermissions`: FilePermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetPermissions`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersGetRootFolder(context.Background()).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersGetRootFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersGetRootFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersGetRootFolder`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | moving folder id
	folderId := "folderId_example" // string | destination folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersMoveFolder(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersMoveFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersMoveFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersMoveFolder`: %v\n", resp)
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


## ExportFoldersMoveFolderToBin

> ExportFoldersMoveFolderToBin(ctx, id).Execute()

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
	r, err := apiClient.ExportsAPI.ExportFoldersMoveFolderToBin(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersMoveFolderToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFoldersMoveFolderToBinRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | Identifier of parent folder id
	exportFolderCreateVM := *openapiclient.NewExportFolderCreateVM("T_example") // ExportFolderCreateVM | create VM (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersPostFolder(context.Background(), id).ExportFolderCreateVM(exportFolderCreateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersPostFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersPostFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersPostFolder`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFoldersRecoverFolder

> ExportFoldersRecoverFolder(ctx, id).RecoveryPath(recoveryPath).Execute()

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
	r, err := apiClient.ExportsAPI.ExportFoldersRecoverFolder(context.Background(), id).RecoveryPath(recoveryPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersRecoverFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportFoldersRecoverFolderRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	folderRenameVM := *openapiclient.NewFolderRenameVM("Name_example", "T_example") // FolderRenameVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersRenameFolder(context.Background(), id).FolderRenameVM(folderRenameVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersRenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersRenameFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersRenameFolder`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | Identifier of folder
	folderIconVM := *openapiclient.NewFolderIconVM(string(123), "T_example") // FolderIconVM | Update icon model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersUpdateIcon(context.Background(), id).FolderIconVM(folderIconVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersUpdateIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersUpdateIcon`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersUpdateIcon`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissionsCRUDVM("T_example"), openapiclient.FileAdministrate(0), "T_example") // UpdateFilePermissionsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportFoldersUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersUpdatePermissions``: %v\n", err)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	folderTagsUpdateVM := *openapiclient.NewFolderTagsUpdateVM("T_example") // FolderTagsUpdateVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportFoldersUpdateTags(context.Background(), id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportFoldersUpdateTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFoldersUpdateTags`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportFoldersUpdateTags`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id
	folderId := "folderId_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsCopyFile(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsCopyFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsCopyFile`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsCopyFile`: %v\n", resp)
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


## ExportsCreateSharingKey

> FileSharingKeysVM ExportsCreateSharingKey(ctx, id).CreateFileShareVM(createFileShareVM).Execute()

Create a new key, that can be used to share access to a file  (You need Administrate.Anon permission to create a new key)

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
	createFileShareVM := *openapiclient.NewCreateFileShareVM() // CreateFileShareVM | parameters for sharing key creation (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsCreateSharingKey(context.Background(), id).CreateFileShareVM(createFileShareVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsCreateSharingKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsCreateSharingKey`: FileSharingKeysVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsCreateSharingKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsCreateSharingKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFileShareVM** | [**CreateFileShareVM**](CreateFileShareVM.md) | parameters for sharing key creation | 

### Return type

[**FileSharingKeysVM**](FileSharingKeysVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportsDeleteFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsDeleteFile``: %v\n", err)
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


## ExportsDeleteSharingKey

> ExportsDeleteSharingKey(ctx, id, key).Execute()

Deletes a sharing key, making links, that utilizing it no longer work

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
	key := "key_example" // string | key to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportsDeleteSharingKey(context.Background(), id, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsDeleteSharingKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 
**key** | **string** | key to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsDeleteSharingKeyRequest struct via the builder pattern


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

Get export by specified id

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
	id := "id_example" // string | id of export

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsGetFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsGetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsGetFile`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id of export | 

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


## ExportsGetFileHistory

> AuditActionsVM ExportsGetFileHistory(ctx, id).Skip(skip).Take(take).Execute()

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
	resp, r, err := apiClient.ExportsAPI.ExportsGetFileHistory(context.Background(), id).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsGetFileHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsGetFileHistory`: AuditActionsVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsGetFileHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsGetFileHistoryRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsGetFilesCount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsGetFilesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsGetFilesCount`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsGetFilesCount`: %v\n", resp)
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

> ExportsVM ExportsGetFilesList(ctx, id).Skip(skip).Take(take).SearchPattern(searchPattern).OrderBy(orderBy).Desc(desc).UseRegex(useRegex).Execute()

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
	orderBy := openapiclient.FileSorting("CreatedTime") // FileSorting |  (optional)
	desc := true // bool |  (optional) (default to false)
	useRegex := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsGetFilesList(context.Background(), id).Skip(skip).Take(take).SearchPattern(searchPattern).OrderBy(orderBy).Desc(desc).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsGetFilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsGetFilesList`: ExportsVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsGetFilesList`: %v\n", resp)
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
 **orderBy** | [**FileSorting**](FileSorting.md) |  | 
 **desc** | **bool** |  | [default to false]
 **useRegex** | **bool** |  | [default to false]

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
	resp, r, err := apiClient.ExportsAPI.ExportsGetPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsGetPermissions`: FilePermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsGetPermissions`: %v\n", resp)
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


## ExportsGetSharingKeys

> FileSharingKeysVM ExportsGetSharingKeys(ctx, id).Execute()

Returns all sharing keys, associated with the file

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
	resp, r, err := apiClient.ExportsAPI.ExportsGetSharingKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsGetSharingKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsGetSharingKeys`: FileSharingKeysVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsGetSharingKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportsGetSharingKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileSharingKeysVM**](FileSharingKeysVM.md)

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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id
	folderId := "folderId_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsMoveFile(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsMoveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsMoveFile`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsMoveFile`: %v\n", resp)
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


## ExportsMoveFileToBin

> ExportsMoveFileToBin(ctx, id).Execute()

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
	r, err := apiClient.ExportsAPI.ExportsMoveFileToBin(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsMoveFileToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportsMoveFileToBinRequest struct via the builder pattern


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


## ExportsRecoverFile

> ExportsRecoverFile(ctx, id).RecoveryPath(recoveryPath).Execute()

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
	r, err := apiClient.ExportsAPI.ExportsRecoverFile(context.Background(), id).RecoveryPath(recoveryPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsRecoverFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExportsRecoverFileRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	fileRenameVM := *openapiclient.NewFileRenameVM("T_example") // FileRenameVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsRenameFile(context.Background(), id).FileRenameVM(fileRenameVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsRenameFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsRenameFile`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsRenameFile`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	fileIconVM := *openapiclient.NewFileIconVM("T_example") // FileIconVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsUpdateIcon(context.Background(), id).FileIconVM(fileIconVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsUpdateIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsUpdateIcon`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsUpdateIcon`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissionsCRUDVM("T_example"), openapiclient.FileAdministrate(0), "T_example") // UpdateFilePermissionsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportsAPI.ExportsUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsUpdatePermissions``: %v\n", err)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	fileTagsUpdateVM := *openapiclient.NewFileTagsUpdateVM("T_example") // FileTagsUpdateVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportsAPI.ExportsUpdateTags(context.Background(), id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.ExportsUpdateTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportsUpdateTags`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.ExportsUpdateTags`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

