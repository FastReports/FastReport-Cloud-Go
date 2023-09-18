# RunPrepareTemplateTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exports** | Pointer to [**[]RunExportReportTaskVM**](RunExportReportTaskVM.md) |  | [optional] 
**PagesCount** | Pointer to **int32** |  | [optional] 
**ReportParameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRunPrepareTemplateTaskVM

`func NewRunPrepareTemplateTaskVM() *RunPrepareTemplateTaskVM`

NewRunPrepareTemplateTaskVM instantiates a new RunPrepareTemplateTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunPrepareTemplateTaskVMWithDefaults

`func NewRunPrepareTemplateTaskVMWithDefaults() *RunPrepareTemplateTaskVM`

NewRunPrepareTemplateTaskVMWithDefaults instantiates a new RunPrepareTemplateTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExports

`func (o *RunPrepareTemplateTaskVM) GetExports() []RunExportReportTaskVM`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *RunPrepareTemplateTaskVM) GetExportsOk() (*[]RunExportReportTaskVM, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *RunPrepareTemplateTaskVM) SetExports(v []RunExportReportTaskVM)`

SetExports sets Exports field to given value.

### HasExports

`func (o *RunPrepareTemplateTaskVM) HasExports() bool`

HasExports returns a boolean if a field has been set.

### SetExportsNil

`func (o *RunPrepareTemplateTaskVM) SetExportsNil(b bool)`

 SetExportsNil sets the value for Exports to be an explicit nil

### UnsetExports
`func (o *RunPrepareTemplateTaskVM) UnsetExports()`

UnsetExports ensures that no value is present for Exports, not even an explicit nil
### GetPagesCount

`func (o *RunPrepareTemplateTaskVM) GetPagesCount() int32`

GetPagesCount returns the PagesCount field if non-nil, zero value otherwise.

### GetPagesCountOk

`func (o *RunPrepareTemplateTaskVM) GetPagesCountOk() (*int32, bool)`

GetPagesCountOk returns a tuple with the PagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCount

`func (o *RunPrepareTemplateTaskVM) SetPagesCount(v int32)`

SetPagesCount sets PagesCount field to given value.

### HasPagesCount

`func (o *RunPrepareTemplateTaskVM) HasPagesCount() bool`

HasPagesCount returns a boolean if a field has been set.

### GetReportParameters

`func (o *RunPrepareTemplateTaskVM) GetReportParameters() map[string]string`

GetReportParameters returns the ReportParameters field if non-nil, zero value otherwise.

### GetReportParametersOk

`func (o *RunPrepareTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool)`

GetReportParametersOk returns a tuple with the ReportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportParameters

`func (o *RunPrepareTemplateTaskVM) SetReportParameters(v map[string]string)`

SetReportParameters sets ReportParameters field to given value.

### HasReportParameters

`func (o *RunPrepareTemplateTaskVM) HasReportParameters() bool`

HasReportParameters returns a boolean if a field has been set.

### SetReportParametersNil

`func (o *RunPrepareTemplateTaskVM) SetReportParametersNil(b bool)`

 SetReportParametersNil sets the value for ReportParameters to be an explicit nil

### UnsetReportParameters
`func (o *RunPrepareTemplateTaskVM) UnsetReportParameters()`

UnsetReportParameters ensures that no value is present for ReportParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


