# ReportVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | Pointer to **NullableString** |  | [optional] 
**ReportInfo** | Pointer to [**ReportInfo**](ReportInfo.md) |  | [optional] 

## Methods

### NewReportVM

`func NewReportVM() *ReportVM`

NewReportVM instantiates a new ReportVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportVMWithDefaults

`func NewReportVMWithDefaults() *ReportVM`

NewReportVMWithDefaults instantiates a new ReportVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *ReportVM) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *ReportVM) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *ReportVM) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *ReportVM) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *ReportVM) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *ReportVM) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetReportInfo

`func (o *ReportVM) GetReportInfo() ReportInfo`

GetReportInfo returns the ReportInfo field if non-nil, zero value otherwise.

### GetReportInfoOk

`func (o *ReportVM) GetReportInfoOk() (*ReportInfo, bool)`

GetReportInfoOk returns a tuple with the ReportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInfo

`func (o *ReportVM) SetReportInfo(v ReportInfo)`

SetReportInfo sets ReportInfo field to given value.

### HasReportInfo

`func (o *ReportVM) HasReportInfo() bool`

HasReportInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


