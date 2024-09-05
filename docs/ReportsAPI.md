# \ReportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportFolderAndFileClearRecycleBin**](ReportsAPI.md#ReportFolderAndFileClearRecycleBin) | **Delete** /api/rp/v1/Reports/{subscriptionId}/ClearRecycleBin | Delete all folders and files from recycle bin
[**ReportFolderAndFileCopyFiles**](ReportsAPI.md#ReportFolderAndFileCopyFiles) | **Post** /api/rp/v1/Reports/{subscriptionId}/CopyFiles | Copy folders and files to a specified folder
[**ReportFolderAndFileCountRecycleBinFoldersAndFiles**](ReportsAPI.md#ReportFolderAndFileCountRecycleBinFoldersAndFiles) | **Get** /api/rp/v1/Reports/{subscriptionId}/CountRecycleBinFolderAndFiles | Count all folders and files from recycle bin
[**ReportFolderAndFileDeleteFiles**](ReportsAPI.md#ReportFolderAndFileDeleteFiles) | **Post** /api/rp/v1/Reports/{subscriptionId}/DeleteFiles | Delete folders and files
[**ReportFolderAndFileGetCount**](ReportsAPI.md#ReportFolderAndFileGetCount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFolderAndFiles | Get count of files and folders what contains in a specified folder
[**ReportFolderAndFileGetFoldersAndFiles**](ReportsAPI.md#ReportFolderAndFileGetFoldersAndFiles) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFolderAndFiles | Get all folders and files from specified folder
[**ReportFolderAndFileGetRecycleBinFoldersAndFiles**](ReportsAPI.md#ReportFolderAndFileGetRecycleBinFoldersAndFiles) | **Get** /api/rp/v1/Reports/{subscriptionId}/ListRecycleBinFolderAndFiles | Get all folders and files from recycle bin
[**ReportFolderAndFileMoveFiles**](ReportsAPI.md#ReportFolderAndFileMoveFiles) | **Post** /api/rp/v1/Reports/{subscriptionId}/MoveFiles | Move folders and files to a specified folder
[**ReportFolderAndFileMoveFilesToBin**](ReportsAPI.md#ReportFolderAndFileMoveFilesToBin) | **Post** /api/rp/v1/Reports/{subscriptionId}/ToBin | Move folders and files to bin
[**ReportFolderAndFileRecoverAllFromRecycleBin**](ReportsAPI.md#ReportFolderAndFileRecoverAllFromRecycleBin) | **Post** /api/rp/v1/Reports/{subscriptionId}/RecoverRecycleBin | Recover all folders and files from recycle bin
[**ReportFolderAndFileRecoverFiles**](ReportsAPI.md#ReportFolderAndFileRecoverFiles) | **Post** /api/rp/v1/Reports/{subscriptionId}/RecoverFiles | Recover folders and files from bin
[**ReportFoldersCalculateFolderSize**](ReportsAPI.md#ReportFoldersCalculateFolderSize) | **Get** /api/rp/v1/Reports/Folder/{id}/size | Get specified folder, calculate it&#39;s size
[**ReportFoldersCopyFolder**](ReportsAPI.md#ReportFoldersCopyFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Copy/{folderId} | Move folder to a specified folder
[**ReportFoldersDeleteFolder**](ReportsAPI.md#ReportFoldersDeleteFolder) | **Delete** /api/rp/v1/Reports/Folder/{id} | Delete specified folder
[**ReportFoldersExport**](ReportsAPI.md#ReportFoldersExport) | **Post** /api/rp/v1/Reports/Folder/{id}/Export | Export specified report folder to a specified format
[**ReportFoldersGetBreadcrumbs**](ReportsAPI.md#ReportFoldersGetBreadcrumbs) | **Get** /api/rp/v1/Reports/Folder/{id}/Breadcrumbs | Get specified folder breadcrumbs
[**ReportFoldersGetFolder**](ReportsAPI.md#ReportFoldersGetFolder) | **Get** /api/rp/v1/Reports/Folder/{id} | Get specified folder
[**ReportFoldersGetFolders**](ReportsAPI.md#ReportFoldersGetFolders) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFolders | Get all folders from specified folder
[**ReportFoldersGetFoldersCount**](ReportsAPI.md#ReportFoldersGetFoldersCount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFolders | Get count of folders what contains in a specified folder
[**ReportFoldersGetOrCreate**](ReportsAPI.md#ReportFoldersGetOrCreate) | **Get** /api/rp/v1/Reports/Folder/getOrCreate | Get specified folder
[**ReportFoldersGetPermissions**](ReportsAPI.md#ReportFoldersGetPermissions) | **Get** /api/rp/v1/Reports/Folder/{id}/permissions | Get all folder permissions
[**ReportFoldersGetRootFolder**](ReportsAPI.md#ReportFoldersGetRootFolder) | **Get** /api/rp/v1/Reports/Root | Get user&#39;s root folder (without parents)
[**ReportFoldersMoveFolder**](ReportsAPI.md#ReportFoldersMoveFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Move/{folderId} | Move folder to a specified folder
[**ReportFoldersMoveFolderToBin**](ReportsAPI.md#ReportFoldersMoveFolderToBin) | **Delete** /api/rp/v1/Reports/Folder/{id}/ToBin | Move specified folder to recycle bin
[**ReportFoldersPostFolder**](ReportsAPI.md#ReportFoldersPostFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Folder | Create folder
[**ReportFoldersRecoverFolder**](ReportsAPI.md#ReportFoldersRecoverFolder) | **Post** /api/rp/v1/Reports/Folder/{id}/Recover | Recover specified folder
[**ReportFoldersRenameFolder**](ReportsAPI.md#ReportFoldersRenameFolder) | **Put** /api/rp/v1/Reports/Folder/{id}/Rename | Rename a folder
[**ReportFoldersUpdateIcon**](ReportsAPI.md#ReportFoldersUpdateIcon) | **Put** /api/rp/v1/Reports/Folder/{id}/Icon | Update a folder&#39;s icon
[**ReportFoldersUpdatePermissions**](ReportsAPI.md#ReportFoldersUpdatePermissions) | **Post** /api/rp/v1/Reports/{id}/permissions | Update permissions
[**ReportFoldersUpdateTags**](ReportsAPI.md#ReportFoldersUpdateTags) | **Put** /api/rp/v1/Reports/Folder/{id}/UpdateTags | Update tags
[**ReportsCopyFile**](ReportsAPI.md#ReportsCopyFile) | **Post** /api/rp/v1/Reports/File/{id}/Copy/{folderId} | Copy file to a specified folder
[**ReportsCreateSharingKey**](ReportsAPI.md#ReportsCreateSharingKey) | **Post** /api/rp/v1/Reports/File/{id}/sharingKey | Create a new key, that can be used to share access to a file  (You need Administrate.Anon permission to create a new key)
[**ReportsDeleteFile**](ReportsAPI.md#ReportsDeleteFile) | **Delete** /api/rp/v1/Reports/File/{id} | Delete specified file
[**ReportsDeleteSharingKey**](ReportsAPI.md#ReportsDeleteSharingKey) | **Delete** /api/rp/v1/Reports/File/{id}/sharingKey | Deletes a sharing key, making links, that utilizing it no longer work
[**ReportsExport**](ReportsAPI.md#ReportsExport) | **Post** /api/rp/v1/Reports/File/{id}/Export | Export specified report to a specified format
[**ReportsGetFile**](ReportsAPI.md#ReportsGetFile) | **Get** /api/rp/v1/Reports/File/{id} | Get specified file
[**ReportsGetFileHistory**](ReportsAPI.md#ReportsGetFileHistory) | **Get** /api/rp/v1/Reports/File/{id}/History | Returns list of actions, performed on this file
[**ReportsGetFilesCount**](ReportsAPI.md#ReportsGetFilesCount) | **Get** /api/rp/v1/Reports/Folder/{id}/CountFiles | Get count of files what contains in a specified folder
[**ReportsGetFilesList**](ReportsAPI.md#ReportsGetFilesList) | **Get** /api/rp/v1/Reports/Folder/{id}/ListFiles | Get all files from specified folder. &lt;br /&gt;  User with Get Entity permission can access this method. &lt;br /&gt;  The method will returns minimal infomration about the file: &lt;br /&gt;  id, name, size, editedTime, createdTime, tags, status, statusReason.
[**ReportsGetPermissions**](ReportsAPI.md#ReportsGetPermissions) | **Get** /api/rp/v1/Reports/File/{id}/permissions | 
[**ReportsGetSharingKeys**](ReportsAPI.md#ReportsGetSharingKeys) | **Get** /api/rp/v1/Reports/File/{id}/sharingKeys | Returns all sharing keys, associated with the file
[**ReportsMoveFile**](ReportsAPI.md#ReportsMoveFile) | **Post** /api/rp/v1/Reports/File/{id}/Move/{folderId} | Move file to a specified folder
[**ReportsMoveFileToBin**](ReportsAPI.md#ReportsMoveFileToBin) | **Delete** /api/rp/v1/Reports/File/{id}/ToBin | Move specified file to recycle bin
[**ReportsRecoverFile**](ReportsAPI.md#ReportsRecoverFile) | **Post** /api/rp/v1/Reports/File/{id}/Recover | Recover specified file from bin
[**ReportsRenameFile**](ReportsAPI.md#ReportsRenameFile) | **Put** /api/rp/v1/Reports/File/{id}/Rename | Rename a file
[**ReportsStaticPreview**](ReportsAPI.md#ReportsStaticPreview) | **Post** /api/rp/v1/Reports/File/{id}/StaticPreview | Make preview for the report.  Generate a new or return exist prepared svg files.  If template was changed will be returned a new.  Pass the &#x60;&#x60; parameter to check prepared timestamp
[**ReportsUpdateIcon**](ReportsAPI.md#ReportsUpdateIcon) | **Put** /api/rp/v1/Reports/File/{id}/Icon | Update a files&#39;s icon
[**ReportsUpdatePermissions**](ReportsAPI.md#ReportsUpdatePermissions) | **Post** /api/rp/v1/Reports/File/{id}/permissions | Update permissions
[**ReportsUpdateTags**](ReportsAPI.md#ReportsUpdateTags) | **Put** /api/rp/v1/Reports/File/{id}/UpdateTags | Update tags
[**ReportsUploadFile**](ReportsAPI.md#ReportsUploadFile) | **Post** /api/rp/v1/Reports/Folder/{id}/File | Upload a file to the specified folder. The method is deprecated, use the UploadFileV2 method instead!
[**ReportsUploadFileV2**](ReportsAPI.md#ReportsUploadFileV2) | **Post** /api/rp/v2/Reports/Folder/{id}/File | Alternative api for upload a file to the specified folder!



## ReportFolderAndFileClearRecycleBin

> ReportFolderAndFileClearRecycleBin(ctx, subscriptionId).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileClearRecycleBin(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileClearRecycleBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileClearRecycleBinRequest struct via the builder pattern


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


## ReportFolderAndFileCopyFiles

> ReportFolderAndFileCopyFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileCopyFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileCopyFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileCopyFilesRequest struct via the builder pattern


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


## ReportFolderAndFileCountRecycleBinFoldersAndFiles

> CountVM ReportFolderAndFileCountRecycleBinFoldersAndFiles(ctx, subscriptionId).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFolderAndFileCountRecycleBinFoldersAndFiles(context.Background(), subscriptionId).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileCountRecycleBinFoldersAndFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFolderAndFileCountRecycleBinFoldersAndFiles`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFolderAndFileCountRecycleBinFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFolderAndFileCountRecycleBinFoldersAndFilesRequest struct via the builder pattern


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


## ReportFolderAndFileDeleteFiles

> ReportFolderAndFileDeleteFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileDeleteFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileDeleteFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileDeleteFilesRequest struct via the builder pattern


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


## ReportFolderAndFileGetCount

> CountVM ReportFolderAndFileGetCount(ctx, id).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFolderAndFileGetCount(context.Background(), id).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileGetCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFolderAndFileGetCount`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFolderAndFileGetCount`: %v\n", resp)
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


## ReportFolderAndFileGetFoldersAndFiles

> FilesVM ReportFolderAndFileGetFoldersAndFiles(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFolderAndFileGetFoldersAndFiles(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileGetFoldersAndFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFolderAndFileGetFoldersAndFiles`: FilesVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFolderAndFileGetFoldersAndFiles`: %v\n", resp)
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


## ReportFolderAndFileGetRecycleBinFoldersAndFiles

> FilesVM ReportFolderAndFileGetRecycleBinFoldersAndFiles(ctx, subscriptionId).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFolderAndFileGetRecycleBinFoldersAndFiles(context.Background(), subscriptionId).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileGetRecycleBinFoldersAndFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFolderAndFileGetRecycleBinFoldersAndFiles`: FilesVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFolderAndFileGetRecycleBinFoldersAndFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | subscription id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFolderAndFileGetRecycleBinFoldersAndFilesRequest struct via the builder pattern


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


## ReportFolderAndFileMoveFiles

> ReportFolderAndFileMoveFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileMoveFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileMoveFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileMoveFilesRequest struct via the builder pattern


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


## ReportFolderAndFileMoveFilesToBin

> ReportFolderAndFileMoveFilesToBin(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileMoveFilesToBin(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileMoveFilesToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileMoveFilesToBinRequest struct via the builder pattern


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


## ReportFolderAndFileRecoverAllFromRecycleBin

> ReportFolderAndFileRecoverAllFromRecycleBin(ctx, subscriptionId).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileRecoverAllFromRecycleBin(context.Background(), subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileRecoverAllFromRecycleBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileRecoverAllFromRecycleBinRequest struct via the builder pattern


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


## ReportFolderAndFileRecoverFiles

> ReportFolderAndFileRecoverFiles(ctx, subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFolderAndFileRecoverFiles(context.Background(), subscriptionId).SelectedFilesVM(selectedFilesVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFolderAndFileRecoverFiles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFolderAndFileRecoverFilesRequest struct via the builder pattern


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


## ReportFoldersCalculateFolderSize

> FolderSizeVM ReportFoldersCalculateFolderSize(ctx, id).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFoldersCalculateFolderSize(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersCalculateFolderSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersCalculateFolderSize`: FolderSizeVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersCalculateFolderSize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersCalculateFolderSizeRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | moving folder id
	folderId := "folderId_example" // string | destination folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersCopyFolder(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersCopyFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersCopyFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersCopyFolder`: %v\n", resp)
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

> ReportFoldersDeleteFolder(ctx, id).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFoldersDeleteFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersDeleteFolder``: %v\n", err)
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


## ReportFoldersExport

> FileVM ReportFoldersExport(ctx, id).ExportReportVM(exportReportVM).Execute()

Export specified report folder to a specified format



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
	id := "id_example" // string | report folder id
	exportReportVM := *openapiclient.NewExportReportVM("T_example") // ExportReportVM | export parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersExport(context.Background(), id).ExportReportVM(exportReportVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersExport`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | report folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportReportVM** | [**ExportReportVM**](ExportReportVM.md) | export parameters | 

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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetBreadcrumbs(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetBreadcrumbs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetBreadcrumbs`: BreadcrumbsVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetBreadcrumbs`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetFolder(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetFolder`: %v\n", resp)
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

> FilesVM ReportFoldersGetFolders(ctx, id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetFolders(context.Background(), id).Skip(skip).Take(take).OrderBy(orderBy).Desc(desc).SearchPattern(searchPattern).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetFolders`: FilesVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetFolders`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetFoldersCount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetFoldersCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetFoldersCount`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetFoldersCount`: %v\n", resp)
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


## ReportFoldersGetOrCreate

> FileVM ReportFoldersGetOrCreate(ctx).Name(name).SubscriptionId(subscriptionId).ParentId(parentId).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetOrCreate(context.Background()).Name(name).SubscriptionId(subscriptionId).ParentId(parentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetOrCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetOrCreate`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetOrCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportFoldersGetOrCreateRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetPermissions`: FilePermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetPermissions`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	subscriptionId := "subscriptionId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersGetRootFolder(context.Background()).SubscriptionId(subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersGetRootFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersGetRootFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersGetRootFolder`: %v\n", resp)
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | moving folder id
	folderId := "folderId_example" // string | destination folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersMoveFolder(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersMoveFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersMoveFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersMoveFolder`: %v\n", resp)
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


## ReportFoldersMoveFolderToBin

> ReportFoldersMoveFolderToBin(ctx, id).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFoldersMoveFolderToBin(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersMoveFolderToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFoldersMoveFolderToBinRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | Identifier of parent folder id
	reportFolderCreateVM := *openapiclient.NewReportFolderCreateVM("T_example") // ReportFolderCreateVM | create VM (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersPostFolder(context.Background(), id).ReportFolderCreateVM(reportFolderCreateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersPostFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersPostFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersPostFolder`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportFoldersRecoverFolder

> ReportFoldersRecoverFolder(ctx, id).RecoveryPath(recoveryPath).Execute()

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
	r, err := apiClient.ReportsAPI.ReportFoldersRecoverFolder(context.Background(), id).RecoveryPath(recoveryPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersRecoverFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportFoldersRecoverFolderRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	folderRenameVM := *openapiclient.NewFolderRenameVM("Name_example", "T_example") // FolderRenameVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersRenameFolder(context.Background(), id).FolderRenameVM(folderRenameVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersRenameFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersRenameFolder`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersRenameFolder`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | Identifier of folder
	folderIconVM := *openapiclient.NewFolderIconVM(string(123), "T_example") // FolderIconVM | Update icon model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersUpdateIcon(context.Background(), id).FolderIconVM(folderIconVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersUpdateIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersUpdateIcon`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersUpdateIcon`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissionsCRUDVM("T_example"), openapiclient.FileAdministrate(0), "T_example") // UpdateFilePermissionsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.ReportFoldersUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersUpdatePermissions``: %v\n", err)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	folderTagsUpdateVM := *openapiclient.NewFolderTagsUpdateVM("T_example") // FolderTagsUpdateVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportFoldersUpdateTags(context.Background(), id).FolderTagsUpdateVM(folderTagsUpdateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportFoldersUpdateTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportFoldersUpdateTags`: FileVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportFoldersUpdateTags`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id
	folderId := "folderId_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsCopyFile(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsCopyFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsCopyFile`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsCopyFile`: %v\n", resp)
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


## ReportsCreateSharingKey

> FileSharingKeysVM ReportsCreateSharingKey(ctx, id).CreateFileShareVM(createFileShareVM).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportsCreateSharingKey(context.Background(), id).CreateFileShareVM(createFileShareVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsCreateSharingKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsCreateSharingKey`: FileSharingKeysVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsCreateSharingKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsCreateSharingKeyRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.ReportsDeleteFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsDeleteFile``: %v\n", err)
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


## ReportsDeleteSharingKey

> ReportsDeleteSharingKey(ctx, id, key).Execute()

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
	r, err := apiClient.ReportsAPI.ReportsDeleteSharingKey(context.Background(), id, key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsDeleteSharingKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportsDeleteSharingKeyRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | report id
	exportReportVM := *openapiclient.NewExportReportVM("T_example") // ExportReportVM | export parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsExport(context.Background(), id).ExportReportVM(exportReportVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsExport`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsExport`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGetFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGetFile`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGetFile`: %v\n", resp)
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


## ReportsGetFileHistory

> AuditActionsVM ReportsGetFileHistory(ctx, id).Skip(skip).Take(take).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportsGetFileHistory(context.Background(), id).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGetFileHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGetFileHistory`: AuditActionsVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGetFileHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetFileHistoryRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsGetFilesCount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGetFilesCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGetFilesCount`: CountVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGetFilesCount`: %v\n", resp)
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

> ReportsVM ReportsGetFilesList(ctx, id).Skip(skip).Take(take).SearchPattern(searchPattern).OrderBy(orderBy).Desc(desc).UseRegex(useRegex).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportsGetFilesList(context.Background(), id).Skip(skip).Take(take).SearchPattern(searchPattern).OrderBy(orderBy).Desc(desc).UseRegex(useRegex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGetFilesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGetFilesList`: ReportsVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGetFilesList`: %v\n", resp)
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
 **orderBy** | [**FileSorting**](FileSorting.md) |  | 
 **desc** | **bool** |  | [default to false]
 **useRegex** | **bool** |  | [default to false]

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
	resp, r, err := apiClient.ReportsAPI.ReportsGetPermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGetPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGetPermissions`: FilePermissionsVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGetPermissions`: %v\n", resp)
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


## ReportsGetSharingKeys

> FileSharingKeysVM ReportsGetSharingKeys(ctx, id).Execute()

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
	resp, r, err := apiClient.ReportsAPI.ReportsGetSharingKeys(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsGetSharingKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsGetSharingKeys`: FileSharingKeysVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsGetSharingKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsGetSharingKeysRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | file id
	folderId := "folderId_example" // string | folder id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsMoveFile(context.Background(), id, folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsMoveFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsMoveFile`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsMoveFile`: %v\n", resp)
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


## ReportsMoveFileToBin

> ReportsMoveFileToBin(ctx, id).Execute()

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
	r, err := apiClient.ReportsAPI.ReportsMoveFileToBin(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsMoveFileToBin``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportsMoveFileToBinRequest struct via the builder pattern


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


## ReportsRecoverFile

> ReportsRecoverFile(ctx, id).RecoveryPath(recoveryPath).Execute()

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
	r, err := apiClient.ReportsAPI.ReportsRecoverFile(context.Background(), id).RecoveryPath(recoveryPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsRecoverFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiReportsRecoverFileRequest struct via the builder pattern


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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	fileRenameVM := *openapiclient.NewFileRenameVM("T_example") // FileRenameVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsRenameFile(context.Background(), id).FileRenameVM(fileRenameVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsRenameFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsRenameFile`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsRenameFile`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsStaticPreview

> ExportVM ReportsStaticPreview(ctx, id).PreviewReportVM(previewReportVM).Execute()

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
	previewReportVM := *openapiclient.NewPreviewReportVM("T_example") // PreviewReportVM | Model with parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsStaticPreview(context.Background(), id).PreviewReportVM(previewReportVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsStaticPreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsStaticPreview`: ExportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsStaticPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsStaticPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewReportVM** | [**PreviewReportVM**](PreviewReportVM.md) | Model with parameters | 

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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	fileIconVM := *openapiclient.NewFileIconVM("T_example") // FileIconVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsUpdateIcon(context.Background(), id).FileIconVM(fileIconVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsUpdateIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsUpdateIcon`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsUpdateIcon`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	updateFilePermissionsVM := *openapiclient.NewUpdateFilePermissionsVM(*openapiclient.NewFilePermissionsCRUDVM("T_example"), openapiclient.FileAdministrate(0), "T_example") // UpdateFilePermissionsVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.ReportsUpdatePermissions(context.Background(), id).UpdateFilePermissionsVM(updateFilePermissionsVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsUpdatePermissions``: %v\n", err)
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

- **Content-Type**: application/json, text/json, application/*+json
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
	openapiclient "github.com/fastreports/gofrcloud"
)

func main() {
	id := "id_example" // string | 
	fileTagsUpdateVM := *openapiclient.NewFileTagsUpdateVM("T_example") // FileTagsUpdateVM |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsUpdateTags(context.Background(), id).FileTagsUpdateVM(fileTagsUpdateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsUpdateTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsUpdateTags`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsUpdateTags`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUploadFile

> ReportVM ReportsUploadFile(ctx, id).ReportCreateVM(reportCreateVM).Execute()

Upload a file to the specified folder. The method is deprecated, use the UploadFileV2 method instead!



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
	reportCreateVM := *openapiclient.NewReportCreateVM("T_example") // ReportCreateVM | file's view model (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsUploadFile(context.Background(), id).ReportCreateVM(reportCreateVM).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsUploadFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsUploadFile`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsUploadFile`: %v\n", resp)
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

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsUploadFileV2

> ReportVM ReportsUploadFileV2(ctx, id).FileContent(fileContent).TemplateId(templateId).Tags(tags).Icon(icon).Execute()

Alternative api for upload a file to the specified folder!



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
	fileContent := os.NewFile(1234, "some_file") // *os.File | 
	templateId := "templateId_example" // string |  (optional)
	tags := []string{"Inner_example"} // []string |  (optional)
	icon := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.ReportsUploadFileV2(context.Background(), id).FileContent(fileContent).TemplateId(templateId).Tags(tags).Icon(icon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsUploadFileV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReportsUploadFileV2`: ReportVM
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsUploadFileV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsUploadFileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileContent** | ***os.File** |  | 
 **templateId** | **string** |  | 
 **tags** | **[]string** |  | 
 **icon** | ***os.File** |  | 

### Return type

[**ReportVM**](ReportVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

