# CreateExportReportTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportParameters** | Pointer to **map[string]string** |  | [optional] 
**Format** | Pointer to [**ExportFormat**](ExportFormat.md) |  | [optional] 
**PagesCount** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateExportReportTaskVM

`func NewCreateExportReportTaskVM(t string, ) *CreateExportReportTaskVM`

NewCreateExportReportTaskVM instantiates a new CreateExportReportTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExportReportTaskVMWithDefaults

`func NewCreateExportReportTaskVMWithDefaults() *CreateExportReportTaskVM`

NewCreateExportReportTaskVMWithDefaults instantiates a new CreateExportReportTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportParameters

`func (o *CreateExportReportTaskVM) GetExportParameters() map[string]string`

GetExportParameters returns the ExportParameters field if non-nil, zero value otherwise.

### GetExportParametersOk

`func (o *CreateExportReportTaskVM) GetExportParametersOk() (*map[string]string, bool)`

GetExportParametersOk returns a tuple with the ExportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportParameters

`func (o *CreateExportReportTaskVM) SetExportParameters(v map[string]string)`

SetExportParameters sets ExportParameters field to given value.

### HasExportParameters

`func (o *CreateExportReportTaskVM) HasExportParameters() bool`

HasExportParameters returns a boolean if a field has been set.

### SetExportParametersNil

`func (o *CreateExportReportTaskVM) SetExportParametersNil(b bool)`

 SetExportParametersNil sets the value for ExportParameters to be an explicit nil

### UnsetExportParameters
`func (o *CreateExportReportTaskVM) UnsetExportParameters()`

UnsetExportParameters ensures that no value is present for ExportParameters, not even an explicit nil
### GetFormat

`func (o *CreateExportReportTaskVM) GetFormat() ExportFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CreateExportReportTaskVM) GetFormatOk() (*ExportFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CreateExportReportTaskVM) SetFormat(v ExportFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CreateExportReportTaskVM) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPagesCount

`func (o *CreateExportReportTaskVM) GetPagesCount() int32`

GetPagesCount returns the PagesCount field if non-nil, zero value otherwise.

### GetPagesCountOk

`func (o *CreateExportReportTaskVM) GetPagesCountOk() (*int32, bool)`

GetPagesCountOk returns a tuple with the PagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCount

`func (o *CreateExportReportTaskVM) SetPagesCount(v int32)`

SetPagesCount sets PagesCount field to given value.

### HasPagesCount

`func (o *CreateExportReportTaskVM) HasPagesCount() bool`

HasPagesCount returns a boolean if a field has been set.

### GetT

`func (o *CreateExportReportTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateExportReportTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateExportReportTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


