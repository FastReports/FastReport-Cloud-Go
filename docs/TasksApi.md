# \TasksApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksCreateTask**](TasksApi.md#TasksCreateTask) | **Post** /api/tasks/v1/Tasks | Create a new task
[**TasksDeleteTask**](TasksApi.md#TasksDeleteTask) | **Delete** /api/tasks/v1/Tasks/{taskId} | Delete a task from a storage
[**TasksGet**](TasksApi.md#TasksGet) | **Get** /api/tasks/v1/Tasks/{taskId} | Get a task by a specified id
[**TasksGetList**](TasksApi.md#TasksGetList) | **Get** /api/tasks/v1/Tasks | Get tasks list
[**TasksGetPermissions**](TasksApi.md#TasksGetPermissions) | **Get** /api/tasks/v1/Tasks/{id}/permissions | Get all Task permissions
[**TasksRenameTask**](TasksApi.md#TasksRenameTask) | **Put** /api/tasks/v1/Tasks/{taskId}/rename | Rename a task
[**TasksRunTask**](TasksApi.md#TasksRunTask) | **Post** /api/tasks/v1/Tasks/run | Run a task from request body
[**TasksRunTaskById**](TasksApi.md#TasksRunTaskById) | **Post** /api/tasks/v1/Tasks/{taskId}/run | Run a task by id
[**TasksUpdatePermissions**](TasksApi.md#TasksUpdatePermissions) | **Post** /api/tasks/v1/Tasks/{id}/permissions | Update permissions
[**TasksUpdateTask**](TasksApi.md#TasksUpdateTask) | **Put** /api/tasks/v1/Tasks/{taskId} | Update a task



## TasksCreateTask

> TaskBaseVM TasksCreateTask(ctx).CreateTaskBaseVM(createTaskBaseVM).Execute()

Create a new task

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
    createTaskBaseVM := *openapiclient.NewCreateTaskBaseVM("T_example") // CreateTaskBaseVM | task's view model. You have to specify task type (type: \"ExportTemplate\") (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksCreateTask(context.Background()).CreateTaskBaseVM(createTaskBaseVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksCreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksCreateTask`: TaskBaseVM
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksCreateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTaskBaseVM** | [**CreateTaskBaseVM**](CreateTaskBaseVM.md) | task&#39;s view model. You have to specify task type (type: \&quot;ExportTemplate\&quot;) | 

### Return type

[**TaskBaseVM**](TaskBaseVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksDeleteTask

> TasksDeleteTask(ctx, taskId).Execute()

Delete a task from a storage

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
    taskId := "taskId_example" // string | deleting task id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TasksApi.TasksDeleteTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksDeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | deleting task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksDeleteTaskRequest struct via the builder pattern


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


## TasksGet

> TaskBaseVM TasksGet(ctx, taskId).Execute()

Get a task by a specified id

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
    taskId := "taskId_example" // string | a task id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksGet(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksGet`: TaskBaseVM
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | a task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskBaseVM**](TaskBaseVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksGetList

> TasksVM TasksGetList(ctx).Skip(skip).Take(take).SubscriptionId(subscriptionId).SearchPattern(searchPattern).Execute()

Get tasks list

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
    skip := int32(56) // int32 | number of tasks, that have to be skipped (optional) (default to 0)
    take := int32(56) // int32 | number of tasks, that have to be returned (optional) (default to 10)
    subscriptionId := "subscriptionId_example" // string | subscription id (optional)
    searchPattern := "searchPattern_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksGetList(context.Background()).Skip(skip).Take(take).SubscriptionId(subscriptionId).SearchPattern(searchPattern).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksGetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksGetList`: TasksVM
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksGetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | number of tasks, that have to be skipped | [default to 0]
 **take** | **int32** | number of tasks, that have to be returned | [default to 10]
 **subscriptionId** | **string** | subscription id | 
 **searchPattern** | **string** |  | [default to &quot;&quot;]

### Return type

[**TasksVM**](TasksVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksGetPermissions

> TaskPermissionsVM TasksGetPermissions(ctx, id).Execute()

Get all Task permissions

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
    id := "id_example" // string | task id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksGetPermissions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksGetPermissions`: TaskPermissionsVM
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksGetPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskPermissionsVM**](TaskPermissionsVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRenameTask

> TaskBaseVM TasksRenameTask(ctx, taskId).NewName(newName).Execute()

Rename a task

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
    taskId := "taskId_example" // string | renaming task id
    newName := "newName_example" // string | task's new Name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksRenameTask(context.Background(), taskId).NewName(newName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksRenameTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksRenameTask`: TaskBaseVM
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksRenameTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | renaming task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRenameTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newName** | **string** | task&#39;s new Name | 

### Return type

[**TaskBaseVM**](TaskBaseVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksRunTask

> TasksRunTask(ctx).RunTaskBaseVM(runTaskBaseVM).Execute()

Run a task from request body

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
    runTaskBaseVM := *openapiclient.NewRunTaskBaseVM("T_example") // RunTaskBaseVM | task's view model (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TasksApi.TasksRunTask(context.Background()).RunTaskBaseVM(runTaskBaseVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksRunTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runTaskBaseVM** | [**RunTaskBaseVM**](RunTaskBaseVM.md) | task&#39;s view model | 

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


## TasksRunTaskById

> TasksRunTaskById(ctx, taskId).Execute()

Run a task by id

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
    taskId := "taskId_example" // string | task id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TasksApi.TasksRunTaskById(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksRunTaskById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksRunTaskByIdRequest struct via the builder pattern


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


## TasksUpdatePermissions

> TasksUpdatePermissions(ctx, id).UpdateTaskPermissionsVM(updateTaskPermissionsVM).Execute()

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
    id := "id_example" // string | task id
    updateTaskPermissionsVM := *openapiclient.NewUpdateTaskPermissionsVM(openapiclient.TaskAdministrate(0), *openapiclient.NewTaskPermissions()) // UpdateTaskPermissionsVM | new permissions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TasksApi.TasksUpdatePermissions(context.Background(), id).UpdateTaskPermissionsVM(updateTaskPermissionsVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksUpdatePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksUpdatePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskPermissionsVM** | [**UpdateTaskPermissionsVM**](UpdateTaskPermissionsVM.md) | new permissions | 

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


## TasksUpdateTask

> TaskBaseVM TasksUpdateTask(ctx, taskId).UpdateTaskBaseVM(updateTaskBaseVM).Execute()

Update a task

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
    taskId := "taskId_example" // string | updating task id
    updateTaskBaseVM := *openapiclient.NewUpdateTaskBaseVM("T_example") // UpdateTaskBaseVM | task's view model. You have to specify task type (type: \"ExportTemplate\") (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TasksApi.TasksUpdateTask(context.Background(), taskId).UpdateTaskBaseVM(updateTaskBaseVM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TasksApi.TasksUpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TasksUpdateTask`: TaskBaseVM
    fmt.Fprintf(os.Stdout, "Response from `TasksApi.TasksUpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | updating task id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTaskBaseVM** | [**UpdateTaskBaseVM**](UpdateTaskBaseVM.md) | task&#39;s view model. You have to specify task type (type: \&quot;ExportTemplate\&quot;) | 

### Return type

[**TaskBaseVM**](TaskBaseVM.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

