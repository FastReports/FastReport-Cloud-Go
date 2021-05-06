# \AdminTemplatesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminTemplateFoldersDeleteFolder**](AdminTemplatesApi.md#AdminTemplateFoldersDeleteFolder) | **Delete** /api/admin/v1/TemplateFolders/{id} | Delete specified folder
[**AdminTemplateFoldersGetFolder**](AdminTemplatesApi.md#AdminTemplateFoldersGetFolder) | **Get** /api/admin/v1/TemplateFolders/{id} | Returns a folder by id
[**AdminTemplateFoldersGetFolders**](AdminTemplatesApi.md#AdminTemplateFoldersGetFolders) | **Get** /api/admin/v1/TemplateFolders | Returns a list of folders
[**AdminTemplateFoldersGetPermissions**](AdminTemplatesApi.md#AdminTemplateFoldersGetPermissions) | **Get** /api/admin/v1/TemplateFolders/{id}/permissions | Get all folder permissions
[**AdminTemplateFoldersPostFolder**](AdminTemplatesApi.md#AdminTemplateFoldersPostFolder) | **Post** /api/admin/v1/TemplateFolders | Create a folder
[**AdminTemplateFoldersUpdateFolder**](AdminTemplatesApi.md#AdminTemplateFoldersUpdateFolder) | **Put** /api/admin/v1/TemplateFolders/{id} | Update a folder
[**AdminTemplateFoldersUpdatePermissions**](AdminTemplatesApi.md#AdminTemplateFoldersUpdatePermissions) | **Post** /api/admin/v1/TemplateFolders/{id}/permissions | Revoke permission
[**AdminTemplatesDeleteFile**](AdminTemplatesApi.md#AdminTemplatesDeleteFile) | **Delete** /api/admin/v1/Templates/{id} | Delete specified file
[**AdminTemplatesGetFile**](AdminTemplatesApi.md#AdminTemplatesGetFile) | **Get** /api/admin/v1/Templates/{id} | Returns a file by id
[**AdminTemplatesGetFiles**](AdminTemplatesApi.md#AdminTemplatesGetFiles) | **Get** /api/admin/v1/Templates | Returns a list of files
[**AdminTemplatesGetPermissions**](AdminTemplatesApi.md#AdminTemplatesGetPermissions) | **Get** /api/admin/v1/Templates/{id}/permissions | Get all file permissions
[**AdminTemplatesUpdateFile**](AdminTemplatesApi.md#AdminTemplatesUpdateFile) | **Put** /api/admin/v1/Templates/{id} | Update a file
[**AdminTemplatesUpdatePermissions**](AdminTemplatesApi.md#AdminTemplatesUpdatePermissions) | **Post** /api/admin/v1/Templates/{id}/permissions | Update permissions to file
[**AdminTemplatesUploadFile**](AdminTemplatesApi.md#AdminTemplatesUploadFile) | **Post** /api/admin/v1/Templates | Upload a file to the specified folder



## AdminTemplateFoldersDeleteFolder

> AdminTemplateFoldersDeleteFolder(ctx, id).Recursive(recursive).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersDeleteFolder(context.Background(), id).Recursive(recursive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersDeleteFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminTemplateFoldersDeleteFolderRequest struct via the builder pattern


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


## AdminTemplateFoldersGetFolder

> FileVM AdminTemplateFoldersGetFolder(ctx, id).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersGetFolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersGetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplateFoldersGetFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplateFoldersGetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a folder | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplateFoldersGetFolderRequest struct via the builder pattern


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


## AdminTemplateFoldersGetFolders

> FilesVM AdminTemplateFoldersGetFolders(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersGetFolders(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersGetFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplateFoldersGetFolders`: FilesVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplateFoldersGetFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplateFoldersGetFoldersRequest struct via the builder pattern


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


## AdminTemplateFoldersGetPermissions

> FilePermissionsVM AdminTemplateFoldersGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplateFoldersGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplateFoldersGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplateFoldersGetPermissionsRequest struct via the builder pattern


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


## AdminTemplateFoldersPostFolder

> FileVM AdminTemplateFoldersPostFolder(ctx).FolderVm(folderVm).Execute()

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
    folderVm := *openapiclient.NewAdminTemplateFolderCreateVM() // AdminTemplateFolderCreateVM | folder create vm (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersPostFolder(context.Background()).FolderVm(folderVm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersPostFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplateFoldersPostFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplateFoldersPostFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplateFoldersPostFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderVm** | [**AdminTemplateFolderCreateVM**](AdminTemplateFolderCreateVM.md) | folder create vm | 

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


## AdminTemplateFoldersUpdateFolder

> FileVM AdminTemplateFoldersUpdateFolder(ctx, id).FolderVM(folderVM).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersUpdateFolder(context.Background(), id).FolderVM(folderVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersUpdateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplateFoldersUpdateFolder`: FileVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplateFoldersUpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | folder id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplateFoldersUpdateFolderRequest struct via the builder pattern


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


## AdminTemplateFoldersUpdatePermissions

> AdminTemplateFoldersUpdatePermissions(ctx, id).NewPermissions(newPermissions).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplateFoldersUpdatePermissions(context.Background(), id).NewPermissions(newPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplateFoldersUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminTemplateFoldersUpdatePermissionsRequest struct via the builder pattern


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


## AdminTemplatesDeleteFile

> AdminTemplatesDeleteFile(ctx, id).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesDeleteFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesDeleteFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminTemplatesDeleteFileRequest struct via the builder pattern


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


## AdminTemplatesGetFile

> TemplateVM AdminTemplatesGetFile(ctx, id).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesGetFile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesGetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplatesGetFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplatesGetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of file | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplatesGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTemplatesGetFiles

> TemplatesVM AdminTemplatesGetFiles(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesGetFiles(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesGetFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplatesGetFiles`: TemplatesVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplatesGetFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplatesGetFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Variable for pagination, defautl value is 0 | [default to 0]
 **take** | **int32** | Variable for pagination, default value is 10 | [default to 10]
 **subscriptionId** | **string** | Subscription Id | 

### Return type

[**TemplatesVM**](TemplatesVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTemplatesGetPermissions

> FilePermissionsVM AdminTemplatesGetPermissions(ctx, id).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplatesGetPermissions`: FilePermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplatesGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplatesGetPermissionsRequest struct via the builder pattern


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


## AdminTemplatesUpdateFile

> TemplateVM AdminTemplatesUpdateFile(ctx, id).FileVM(fileVM).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesUpdateFile(context.Background(), id).FileVM(fileVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesUpdateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplatesUpdateFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplatesUpdateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | file id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplatesUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileVM** | [**FileUpdateVM**](FileUpdateVM.md) | file&#39;s view model | 

### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminTemplatesUpdatePermissions

> AdminTemplatesUpdatePermissions(ctx, id).NewPermissions(newPermissions).Execute()

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
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesUpdatePermissions(context.Background(), id).NewPermissions(newPermissions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesUpdatePermissions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdminTemplatesUpdatePermissionsRequest struct via the builder pattern


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


## AdminTemplatesUploadFile

> TemplateVM AdminTemplatesUploadFile(ctx).FileVM(fileVM).Execute()

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
    fileVM := *openapiclient.NewTemplateCreateAdminVM() // TemplateCreateAdminVM | file's view model (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminTemplatesApi.AdminTemplatesUploadFile(context.Background()).FileVM(fileVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminTemplatesApi.AdminTemplatesUploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminTemplatesUploadFile`: TemplateVM
    fmt.Fprintf(os.Stdout, "Response from `AdminTemplatesApi.AdminTemplatesUploadFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminTemplatesUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileVM** | [**TemplateCreateAdminVM**](TemplateCreateAdminVM.md) | file&#39;s view model | 

### Return type

[**TemplateVM**](TemplateVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/_*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

