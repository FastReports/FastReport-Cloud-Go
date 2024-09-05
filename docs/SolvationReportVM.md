# SolvationReportVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SolvedProblems** | Pointer to **map[string]int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewSolvationReportVM

`func NewSolvationReportVM(t string, ) *SolvationReportVM`

NewSolvationReportVM instantiates a new SolvationReportVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolvationReportVMWithDefaults

`func NewSolvationReportVMWithDefaults() *SolvationReportVM`

NewSolvationReportVMWithDefaults instantiates a new SolvationReportVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSolvedProblems

`func (o *SolvationReportVM) GetSolvedProblems() map[string]int32`

GetSolvedProblems returns the SolvedProblems field if non-nil, zero value otherwise.

### GetSolvedProblemsOk

`func (o *SolvationReportVM) GetSolvedProblemsOk() (*map[string]int32, bool)`

GetSolvedProblemsOk returns a tuple with the SolvedProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolvedProblems

`func (o *SolvationReportVM) SetSolvedProblems(v map[string]int32)`

SetSolvedProblems sets SolvedProblems field to given value.

### HasSolvedProblems

`func (o *SolvationReportVM) HasSolvedProblems() bool`

HasSolvedProblems returns a boolean if a field has been set.

### SetSolvedProblemsNil

`func (o *SolvationReportVM) SetSolvedProblemsNil(b bool)`

 SetSolvedProblemsNil sets the value for SolvedProblems to be an explicit nil

### UnsetSolvedProblems
`func (o *SolvationReportVM) UnsetSolvedProblems()`

UnsetSolvedProblems ensures that no value is present for SolvedProblems, not even an explicit nil
### GetT

`func (o *SolvationReportVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SolvationReportVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SolvationReportVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


