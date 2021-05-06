# \AdminSubscriptionProblemSolvingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubscriptionProblemSolvingSolveProblems**](AdminSubscriptionProblemSolvingApi.md#AdminSubscriptionProblemSolvingSolveProblems) | **Post** /api/admin/v1/Analytics/Solve | Solve problems provided by FastReport.Cloud.Admin.Controllers.SubscriptionAnalyticsController



## AdminSubscriptionProblemSolvingSolveProblems

> AdminSubscriptionProblemSolvingSolveProblems(ctx).AnalysisResults(analysisResults).Execute()

Solve problems provided by FastReport.Cloud.Admin.Controllers.SubscriptionAnalyticsController

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
    analysisResults := *openapiclient.NewAnalysisResultsVM() // AnalysisResultsVM |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AdminSubscriptionProblemSolvingApi.AdminSubscriptionProblemSolvingSolveProblems(context.Background()).AnalysisResults(analysisResults).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionProblemSolvingApi.AdminSubscriptionProblemSolvingSolveProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminSubscriptionProblemSolvingSolveProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **analysisResults** | [**AnalysisResultsVM**](AnalysisResultsVM.md) |  | 

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

