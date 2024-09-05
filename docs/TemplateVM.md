# TemplateVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportInfo** | Pointer to [**ReportInfo**](ReportInfo.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTemplateVM

`func NewTemplateVM(t string, ) *TemplateVM`

NewTemplateVM instantiates a new TemplateVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateVMWithDefaults

`func NewTemplateVMWithDefaults() *TemplateVM`

NewTemplateVMWithDefaults instantiates a new TemplateVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportInfo

`func (o *TemplateVM) GetReportInfo() ReportInfo`

GetReportInfo returns the ReportInfo field if non-nil, zero value otherwise.

### GetReportInfoOk

`func (o *TemplateVM) GetReportInfoOk() (*ReportInfo, bool)`

GetReportInfoOk returns a tuple with the ReportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInfo

`func (o *TemplateVM) SetReportInfo(v ReportInfo)`

SetReportInfo sets ReportInfo field to given value.

### HasReportInfo

`func (o *TemplateVM) HasReportInfo() bool`

HasReportInfo returns a boolean if a field has been set.

### GetT

`func (o *TemplateVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TemplateVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TemplateVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


