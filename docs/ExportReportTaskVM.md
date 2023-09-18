# ExportReportTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportParameters** | Pointer to **map[string]string** |  | [optional] 
**Format** | Pointer to [**ExportFormat**](ExportFormat.md) |  | [optional] 
**PagesCount** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewExportReportTaskVM

`func NewExportReportTaskVM(t string, ) *ExportReportTaskVM`

NewExportReportTaskVM instantiates a new ExportReportTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportReportTaskVMWithDefaults

`func NewExportReportTaskVMWithDefaults() *ExportReportTaskVM`

NewExportReportTaskVMWithDefaults instantiates a new ExportReportTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportParameters

`func (o *ExportReportTaskVM) GetExportParameters() map[string]string`

GetExportParameters returns the ExportParameters field if non-nil, zero value otherwise.

### GetExportParametersOk

`func (o *ExportReportTaskVM) GetExportParametersOk() (*map[string]string, bool)`

GetExportParametersOk returns a tuple with the ExportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportParameters

`func (o *ExportReportTaskVM) SetExportParameters(v map[string]string)`

SetExportParameters sets ExportParameters field to given value.

### HasExportParameters

`func (o *ExportReportTaskVM) HasExportParameters() bool`

HasExportParameters returns a boolean if a field has been set.

### SetExportParametersNil

`func (o *ExportReportTaskVM) SetExportParametersNil(b bool)`

 SetExportParametersNil sets the value for ExportParameters to be an explicit nil

### UnsetExportParameters
`func (o *ExportReportTaskVM) UnsetExportParameters()`

UnsetExportParameters ensures that no value is present for ExportParameters, not even an explicit nil
### GetFormat

`func (o *ExportReportTaskVM) GetFormat() ExportFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportReportTaskVM) GetFormatOk() (*ExportFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportReportTaskVM) SetFormat(v ExportFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ExportReportTaskVM) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPagesCount

`func (o *ExportReportTaskVM) GetPagesCount() int32`

GetPagesCount returns the PagesCount field if non-nil, zero value otherwise.

### GetPagesCountOk

`func (o *ExportReportTaskVM) GetPagesCountOk() (*int32, bool)`

GetPagesCountOk returns a tuple with the PagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCount

`func (o *ExportReportTaskVM) SetPagesCount(v int32)`

SetPagesCount sets PagesCount field to given value.

### HasPagesCount

`func (o *ExportReportTaskVM) HasPagesCount() bool`

HasPagesCount returns a boolean if a field has been set.

### GetT

`func (o *ExportReportTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ExportReportTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ExportReportTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


