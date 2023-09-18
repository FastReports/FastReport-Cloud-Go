# UpdateExportReportTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportParameters** | Pointer to **map[string]string** |  | [optional] 
**Format** | Pointer to [**NullableExportFormat**](ExportFormat.md) |  | [optional] 
**PagesCount** | Pointer to **NullableInt32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateExportReportTaskVM

`func NewUpdateExportReportTaskVM(t string, ) *UpdateExportReportTaskVM`

NewUpdateExportReportTaskVM instantiates a new UpdateExportReportTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExportReportTaskVMWithDefaults

`func NewUpdateExportReportTaskVMWithDefaults() *UpdateExportReportTaskVM`

NewUpdateExportReportTaskVMWithDefaults instantiates a new UpdateExportReportTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportParameters

`func (o *UpdateExportReportTaskVM) GetExportParameters() map[string]string`

GetExportParameters returns the ExportParameters field if non-nil, zero value otherwise.

### GetExportParametersOk

`func (o *UpdateExportReportTaskVM) GetExportParametersOk() (*map[string]string, bool)`

GetExportParametersOk returns a tuple with the ExportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportParameters

`func (o *UpdateExportReportTaskVM) SetExportParameters(v map[string]string)`

SetExportParameters sets ExportParameters field to given value.

### HasExportParameters

`func (o *UpdateExportReportTaskVM) HasExportParameters() bool`

HasExportParameters returns a boolean if a field has been set.

### SetExportParametersNil

`func (o *UpdateExportReportTaskVM) SetExportParametersNil(b bool)`

 SetExportParametersNil sets the value for ExportParameters to be an explicit nil

### UnsetExportParameters
`func (o *UpdateExportReportTaskVM) UnsetExportParameters()`

UnsetExportParameters ensures that no value is present for ExportParameters, not even an explicit nil
### GetFormat

`func (o *UpdateExportReportTaskVM) GetFormat() ExportFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UpdateExportReportTaskVM) GetFormatOk() (*ExportFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UpdateExportReportTaskVM) SetFormat(v ExportFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UpdateExportReportTaskVM) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *UpdateExportReportTaskVM) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *UpdateExportReportTaskVM) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetPagesCount

`func (o *UpdateExportReportTaskVM) GetPagesCount() int32`

GetPagesCount returns the PagesCount field if non-nil, zero value otherwise.

### GetPagesCountOk

`func (o *UpdateExportReportTaskVM) GetPagesCountOk() (*int32, bool)`

GetPagesCountOk returns a tuple with the PagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagesCount

`func (o *UpdateExportReportTaskVM) SetPagesCount(v int32)`

SetPagesCount sets PagesCount field to given value.

### HasPagesCount

`func (o *UpdateExportReportTaskVM) HasPagesCount() bool`

HasPagesCount returns a boolean if a field has been set.

### SetPagesCountNil

`func (o *UpdateExportReportTaskVM) SetPagesCountNil(b bool)`

 SetPagesCountNil sets the value for PagesCount to be an explicit nil

### UnsetPagesCount
`func (o *UpdateExportReportTaskVM) UnsetPagesCount()`

UnsetPagesCount ensures that no value is present for PagesCount, not even an explicit nil
### GetT

`func (o *UpdateExportReportTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateExportReportTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateExportReportTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


