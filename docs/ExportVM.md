# ExportVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to [**ExportFormat**](ExportFormat.md) |  | [optional] 
**ReportId** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewExportVM

`func NewExportVM(t string, ) *ExportVM`

NewExportVM instantiates a new ExportVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportVMWithDefaults

`func NewExportVMWithDefaults() *ExportVM`

NewExportVMWithDefaults instantiates a new ExportVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *ExportVM) GetFormat() ExportFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportVM) GetFormatOk() (*ExportFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportVM) SetFormat(v ExportFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ExportVM) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetReportId

`func (o *ExportVM) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ExportVM) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ExportVM) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *ExportVM) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### SetReportIdNil

`func (o *ExportVM) SetReportIdNil(b bool)`

 SetReportIdNil sets the value for ReportId to be an explicit nil

### UnsetReportId
`func (o *ExportVM) UnsetReportId()`

UnsetReportId ensures that no value is present for ReportId, not even an explicit nil
### GetTemplateId

`func (o *ExportVM) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ExportVM) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ExportVM) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ExportVM) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ExportVM) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ExportVM) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetT

`func (o *ExportVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ExportVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ExportVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


