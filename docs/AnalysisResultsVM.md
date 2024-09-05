# AnalysisResultsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]AnalysisResultVM**](AnalysisResultVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAnalysisResultsVM

`func NewAnalysisResultsVM(t string, ) *AnalysisResultsVM`

NewAnalysisResultsVM instantiates a new AnalysisResultsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisResultsVMWithDefaults

`func NewAnalysisResultsVMWithDefaults() *AnalysisResultsVM`

NewAnalysisResultsVMWithDefaults instantiates a new AnalysisResultsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *AnalysisResultsVM) GetResults() []AnalysisResultVM`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AnalysisResultsVM) GetResultsOk() (*[]AnalysisResultVM, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AnalysisResultsVM) SetResults(v []AnalysisResultVM)`

SetResults sets Results field to given value.

### HasResults

`func (o *AnalysisResultsVM) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *AnalysisResultsVM) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *AnalysisResultsVM) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetT

`func (o *AnalysisResultsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AnalysisResultsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AnalysisResultsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


