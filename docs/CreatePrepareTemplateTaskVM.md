# CreatePrepareTemplateTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exports** | Pointer to [**[]CreateExportReportTaskVM**](CreateExportReportTaskVM.md) |  | [optional] 
**PagesCount** | Pointer to **int32** |  | [optional] 
**ReportParameters** | Pointer to **map[string]string** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreatePrepareTemplateTaskVM

`func NewCreatePrepareTemplateTaskVM(t string, ) *CreatePrepareTemplateTaskVM`

NewCreatePrepareTemplateTaskVM instantiates a new CreatePrepareTemplateTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePrepareTemplateTaskVMWithDefaults

`func NewCreatePrepareTemplateTaskVMWithDefaults() *CreatePrepareTemplateTaskVM`

NewCreatePrepareTemplateTaskVMWithDefaults instantiates a new CreatePrepareTemplateTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExports

`func (o *CreatePrepareTemplateTaskVM) GetExports() []CreateExportReportTaskVM`

GetExports returns the Exports field if non-nil, zero value otherwise.

### GetExportsOk

`func (o *CreatePrepareTemplateTaskVM) GetExportsOk() (*[]CreateExportReportTaskVM, bool)`

GetExportsOk returns a tuple with the Exports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExports

`func (o *CreatePrepareTemplateTaskVM) SetExports(v []CreateExportReportTaskVM)`

SetExports sets Exports field to given value.

### HasExports

`func (o *CreatePrepareTemplateTaskVM) HasExports() bool`

HasExports returns a boolean if a field has been set.

### SetExportsNil

`func (o *CreatePrepareTemplateTaskVM) SetExportsNil(b bool)`

 SetExportsNil sets the value for Exports to be an explicit nil

### UnsetExports
`func (o *CreatePrepareTemplateTaskVM) UnsetExports()`

UnsetExports ensures that no value is present for Exports, not even an explicit nil
### GetPagesCount

`func (o *CreatePrepareTemplateTaskVM) GetPagesCount() int32`

GetPagesCount returns the PagesCount field if non-nil, zero value otherwise.

### GetPagesCountOk

`func (o *CreatePrepareTemplateTaskVM) GetPagesCountOk() (*int32, bool)`

GetPagesCountOk returns a tuple with the PagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCount

`func (o *CreatePrepareTemplateTaskVM) SetPagesCount(v int32)`

SetPagesCount sets PagesCount field to given value.

### HasPagesCount

`func (o *CreatePrepareTemplateTaskVM) HasPagesCount() bool`

HasPagesCount returns a boolean if a field has been set.

### GetReportParameters

`func (o *CreatePrepareTemplateTaskVM) GetReportParameters() map[string]string`

GetReportParameters returns the ReportParameters field if non-nil, zero value otherwise.

### GetReportParametersOk

`func (o *CreatePrepareTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool)`

GetReportParametersOk returns a tuple with the ReportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportParameters

`func (o *CreatePrepareTemplateTaskVM) SetReportParameters(v map[string]string)`

SetReportParameters sets ReportParameters field to given value.

### HasReportParameters

`func (o *CreatePrepareTemplateTaskVM) HasReportParameters() bool`

HasReportParameters returns a boolean if a field has been set.

### SetReportParametersNil

`func (o *CreatePrepareTemplateTaskVM) SetReportParametersNil(b bool)`

 SetReportParametersNil sets the value for ReportParameters to be an explicit nil

### UnsetReportParameters
`func (o *CreatePrepareTemplateTaskVM) UnsetReportParameters()`

UnsetReportParameters ensures that no value is present for ReportParameters, not even an explicit nil
### GetT

`func (o *CreatePrepareTemplateTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreatePrepareTemplateTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreatePrepareTemplateTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


