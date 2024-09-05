# \DownloadAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadGetExport**](DownloadAPI.md#DownloadGetExport) | **Get** /download/e/{id} | Returns a export file with specified id
[**DownloadGetExportThumbnail**](DownloadAPI.md#DownloadGetExportThumbnail) | **Get** /download/e/{id}/thumbnail | Returns export&#39;s thumbnail
[**DownloadGetExports**](DownloadAPI.md#DownloadGetExports) | **Get** /download/es/{archiveName} | Returns a zip archive with selected ids
[**DownloadGetLastSVGExport**](DownloadAPI.md#DownloadGetLastSVGExport) | **Get** /download/lastPreview/{reportId} | returns export, that was created from report with specified id.  INTERNAL USAGE ONLY!
[**DownloadGetReport**](DownloadAPI.md#DownloadGetReport) | **Get** /download/r/{id} | Returns a prepared file with specified id
[**DownloadGetReportThumbnail**](DownloadAPI.md#DownloadGetReportThumbnail) | **Get** /download/r/{id}/thumbnail | Returns report&#39;s thumbnail
[**DownloadGetReports**](DownloadAPI.md#DownloadGetReports) | **Get** /download/rs/{archiveName} | Returns a zip archive with selected files
[**DownloadGetTemplate**](DownloadAPI.md#DownloadGetTemplate) | **Get** /download/t/{id} | Returns a Template file with specified id
[**DownloadGetTemplateThumbnail**](DownloadAPI.md#DownloadGetTemplateThumbnail) | **Get** /download/t/{id}/thumbnail | Returns template&#39;s thumbnail
[**DownloadGetTemplates**](DownloadAPI.md#DownloadGetTemplates) | **Get** /download/ts/{archiveName} | Returns a zip archive with selected files



## DownloadGetExport

> *os.File DownloadGetExport(ctx, id).Preview(preview).Execute()

Returns a export file with specified id

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
	preview := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadAPI.DownloadGetExport(context.Background(), id).Preview(preview).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preview** | **bool** |  | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/octet-stream, application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetExportThumbnail

> *os.File DownloadGetExportThumbnail(ctx, id).Execute()

Returns export's thumbnail

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
	resp, r, err := apiClient.DownloadAPI.DownloadGetExportThumbnail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetExportThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetExportThumbnail`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetExportThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetExportThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetExports

> *os.File DownloadGetExports(ctx, archiveName).FileIds(fileIds).FolderIds(folderIds).Execute()

Returns a zip archive with selected ids

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
	archiveName := "archiveName_example" // string | name of the created archive
	fileIds := "fileIds_example" // string | ids separated with a ',' sign (optional)
	folderIds := "folderIds_example" // string | ids separated with a ',' sign (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadAPI.DownloadGetExports(context.Background(), archiveName).FileIds(fileIds).FolderIds(folderIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetExports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetExports`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveName** | **string** | name of the created archive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileIds** | **string** | ids separated with a &#39;,&#39; sign | 
 **folderIds** | **string** | ids separated with a &#39;,&#39; sign | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetLastSVGExport

> *os.File DownloadGetLastSVGExport(ctx, reportId).Execute()

returns export, that was created from report with specified id.  INTERNAL USAGE ONLY!

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
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadAPI.DownloadGetLastSVGExport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetLastSVGExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetLastSVGExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetLastSVGExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetLastSVGExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetReport

> *os.File DownloadGetReport(ctx, id).Execute()

Returns a prepared file with specified id

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
	resp, r, err := apiClient.DownloadAPI.DownloadGetReport(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetReport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetReportThumbnail

> *os.File DownloadGetReportThumbnail(ctx, id).Execute()

Returns report's thumbnail

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
	resp, r, err := apiClient.DownloadAPI.DownloadGetReportThumbnail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetReportThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetReportThumbnail`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetReportThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetReportThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetReports

> *os.File DownloadGetReports(ctx, archiveName).FileIds(fileIds).FolderIds(folderIds).Execute()

Returns a zip archive with selected files

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
	archiveName := "archiveName_example" // string | name of the created archive
	fileIds := "fileIds_example" // string | ids separated with a ',' sign (optional)
	folderIds := "folderIds_example" // string | ids separated with a ',' sign (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadAPI.DownloadGetReports(context.Background(), archiveName).FileIds(fileIds).FolderIds(folderIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetReports`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveName** | **string** | name of the created archive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileIds** | **string** | ids separated with a &#39;,&#39; sign | 
 **folderIds** | **string** | ids separated with a &#39;,&#39; sign | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetTemplate

> *os.File DownloadGetTemplate(ctx, id).Execute()

Returns a Template file with specified id

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadAPI.DownloadGetTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetTemplate`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | template id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetTemplateThumbnail

> *os.File DownloadGetTemplateThumbnail(ctx, id).Execute()

Returns template's thumbnail

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
	resp, r, err := apiClient.DownloadAPI.DownloadGetTemplateThumbnail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetTemplateThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetTemplateThumbnail`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetTemplateThumbnail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetTemplateThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, image/png, image/jpeg

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGetTemplates

> *os.File DownloadGetTemplates(ctx, archiveName).FileIds(fileIds).FolderIds(folderIds).Execute()

Returns a zip archive with selected files

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
	archiveName := "archiveName_example" // string | name of the created archive
	fileIds := "fileIds_example" // string | ids separated with a ',' sign (optional)
	folderIds := "folderIds_example" // string | ids separated with a ',' sign (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DownloadAPI.DownloadGetTemplates(context.Background(), archiveName).FileIds(fileIds).FolderIds(folderIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DownloadAPI.DownloadGetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadGetTemplates`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DownloadAPI.DownloadGetTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**archiveName** | **string** | name of the created archive | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGetTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileIds** | **string** | ids separated with a &#39;,&#39; sign | 
 **folderIds** | **string** | ids separated with a &#39;,&#39; sign | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKey](../README.md#ApiKey), [JWT](../README.md#JWT)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

