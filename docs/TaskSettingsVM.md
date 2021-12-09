# TaskSettingsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prepare** | Pointer to **NullableBool** |  | [optional] 
**ExportTemplate** | Pointer to **NullableBool** |  | [optional] 
**ExportReport** | Pointer to **NullableBool** |  | [optional] 
**SendViaEmail** | Pointer to **NullableBool** |  | [optional] 
**SendViaWebhook** | Pointer to **NullableBool** |  | [optional] 
**FetchData** | Pointer to **NullableBool** |  | [optional] 
**ThumbnailReport** | Pointer to **NullableBool** |  | [optional] 
**ThumbnailTemplate** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewTaskSettingsVM

`func NewTaskSettingsVM() *TaskSettingsVM`

NewTaskSettingsVM instantiates a new TaskSettingsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskSettingsVMWithDefaults

`func NewTaskSettingsVMWithDefaults() *TaskSettingsVM`

NewTaskSettingsVMWithDefaults instantiates a new TaskSettingsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrepare

`func (o *TaskSettingsVM) GetPrepare() bool`

GetPrepare returns the Prepare field if non-nil, zero value otherwise.

### GetPrepareOk

`func (o *TaskSettingsVM) GetPrepareOk() (*bool, bool)`

GetPrepareOk returns a tuple with the Prepare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepare

`func (o *TaskSettingsVM) SetPrepare(v bool)`

SetPrepare sets Prepare field to given value.

### HasPrepare

`func (o *TaskSettingsVM) HasPrepare() bool`

HasPrepare returns a boolean if a field has been set.

### SetPrepareNil

`func (o *TaskSettingsVM) SetPrepareNil(b bool)`

 SetPrepareNil sets the value for Prepare to be an explicit nil

### UnsetPrepare
`func (o *TaskSettingsVM) UnsetPrepare()`

UnsetPrepare ensures that no value is present for Prepare, not even an explicit nil
### GetExportTemplate

`func (o *TaskSettingsVM) GetExportTemplate() bool`

GetExportTemplate returns the ExportTemplate field if non-nil, zero value otherwise.

### GetExportTemplateOk

`func (o *TaskSettingsVM) GetExportTemplateOk() (*bool, bool)`

GetExportTemplateOk returns a tuple with the ExportTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTemplate

`func (o *TaskSettingsVM) SetExportTemplate(v bool)`

SetExportTemplate sets ExportTemplate field to given value.

### HasExportTemplate

`func (o *TaskSettingsVM) HasExportTemplate() bool`

HasExportTemplate returns a boolean if a field has been set.

### SetExportTemplateNil

`func (o *TaskSettingsVM) SetExportTemplateNil(b bool)`

 SetExportTemplateNil sets the value for ExportTemplate to be an explicit nil

### UnsetExportTemplate
`func (o *TaskSettingsVM) UnsetExportTemplate()`

UnsetExportTemplate ensures that no value is present for ExportTemplate, not even an explicit nil
### GetExportReport

`func (o *TaskSettingsVM) GetExportReport() bool`

GetExportReport returns the ExportReport field if non-nil, zero value otherwise.

### GetExportReportOk

`func (o *TaskSettingsVM) GetExportReportOk() (*bool, bool)`

GetExportReportOk returns a tuple with the ExportReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportReport

`func (o *TaskSettingsVM) SetExportReport(v bool)`

SetExportReport sets ExportReport field to given value.

### HasExportReport

`func (o *TaskSettingsVM) HasExportReport() bool`

HasExportReport returns a boolean if a field has been set.

### SetExportReportNil

`func (o *TaskSettingsVM) SetExportReportNil(b bool)`

 SetExportReportNil sets the value for ExportReport to be an explicit nil

### UnsetExportReport
`func (o *TaskSettingsVM) UnsetExportReport()`

UnsetExportReport ensures that no value is present for ExportReport, not even an explicit nil
### GetSendViaEmail

`func (o *TaskSettingsVM) GetSendViaEmail() bool`

GetSendViaEmail returns the SendViaEmail field if non-nil, zero value otherwise.

### GetSendViaEmailOk

`func (o *TaskSettingsVM) GetSendViaEmailOk() (*bool, bool)`

GetSendViaEmailOk returns a tuple with the SendViaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendViaEmail

`func (o *TaskSettingsVM) SetSendViaEmail(v bool)`

SetSendViaEmail sets SendViaEmail field to given value.

### HasSendViaEmail

`func (o *TaskSettingsVM) HasSendViaEmail() bool`

HasSendViaEmail returns a boolean if a field has been set.

### SetSendViaEmailNil

`func (o *TaskSettingsVM) SetSendViaEmailNil(b bool)`

 SetSendViaEmailNil sets the value for SendViaEmail to be an explicit nil

### UnsetSendViaEmail
`func (o *TaskSettingsVM) UnsetSendViaEmail()`

UnsetSendViaEmail ensures that no value is present for SendViaEmail, not even an explicit nil
### GetSendViaWebhook

`func (o *TaskSettingsVM) GetSendViaWebhook() bool`

GetSendViaWebhook returns the SendViaWebhook field if non-nil, zero value otherwise.

### GetSendViaWebhookOk

`func (o *TaskSettingsVM) GetSendViaWebhookOk() (*bool, bool)`

GetSendViaWebhookOk returns a tuple with the SendViaWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendViaWebhook

`func (o *TaskSettingsVM) SetSendViaWebhook(v bool)`

SetSendViaWebhook sets SendViaWebhook field to given value.

### HasSendViaWebhook

`func (o *TaskSettingsVM) HasSendViaWebhook() bool`

HasSendViaWebhook returns a boolean if a field has been set.

### SetSendViaWebhookNil

`func (o *TaskSettingsVM) SetSendViaWebhookNil(b bool)`

 SetSendViaWebhookNil sets the value for SendViaWebhook to be an explicit nil

### UnsetSendViaWebhook
`func (o *TaskSettingsVM) UnsetSendViaWebhook()`

UnsetSendViaWebhook ensures that no value is present for SendViaWebhook, not even an explicit nil
### GetFetchData

`func (o *TaskSettingsVM) GetFetchData() bool`

GetFetchData returns the FetchData field if non-nil, zero value otherwise.

### GetFetchDataOk

`func (o *TaskSettingsVM) GetFetchDataOk() (*bool, bool)`

GetFetchDataOk returns a tuple with the FetchData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchData

`func (o *TaskSettingsVM) SetFetchData(v bool)`

SetFetchData sets FetchData field to given value.

### HasFetchData

`func (o *TaskSettingsVM) HasFetchData() bool`

HasFetchData returns a boolean if a field has been set.

### SetFetchDataNil

`func (o *TaskSettingsVM) SetFetchDataNil(b bool)`

 SetFetchDataNil sets the value for FetchData to be an explicit nil

### UnsetFetchData
`func (o *TaskSettingsVM) UnsetFetchData()`

UnsetFetchData ensures that no value is present for FetchData, not even an explicit nil
### GetThumbnailReport

`func (o *TaskSettingsVM) GetThumbnailReport() bool`

GetThumbnailReport returns the ThumbnailReport field if non-nil, zero value otherwise.

### GetThumbnailReportOk

`func (o *TaskSettingsVM) GetThumbnailReportOk() (*bool, bool)`

GetThumbnailReportOk returns a tuple with the ThumbnailReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailReport

`func (o *TaskSettingsVM) SetThumbnailReport(v bool)`

SetThumbnailReport sets ThumbnailReport field to given value.

### HasThumbnailReport

`func (o *TaskSettingsVM) HasThumbnailReport() bool`

HasThumbnailReport returns a boolean if a field has been set.

### SetThumbnailReportNil

`func (o *TaskSettingsVM) SetThumbnailReportNil(b bool)`

 SetThumbnailReportNil sets the value for ThumbnailReport to be an explicit nil

### UnsetThumbnailReport
`func (o *TaskSettingsVM) UnsetThumbnailReport()`

UnsetThumbnailReport ensures that no value is present for ThumbnailReport, not even an explicit nil
### GetThumbnailTemplate

`func (o *TaskSettingsVM) GetThumbnailTemplate() bool`

GetThumbnailTemplate returns the ThumbnailTemplate field if non-nil, zero value otherwise.

### GetThumbnailTemplateOk

`func (o *TaskSettingsVM) GetThumbnailTemplateOk() (*bool, bool)`

GetThumbnailTemplateOk returns a tuple with the ThumbnailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailTemplate

`func (o *TaskSettingsVM) SetThumbnailTemplate(v bool)`

SetThumbnailTemplate sets ThumbnailTemplate field to given value.

### HasThumbnailTemplate

`func (o *TaskSettingsVM) HasThumbnailTemplate() bool`

HasThumbnailTemplate returns a boolean if a field has been set.

### SetThumbnailTemplateNil

`func (o *TaskSettingsVM) SetThumbnailTemplateNil(b bool)`

 SetThumbnailTemplateNil sets the value for ThumbnailTemplate to be an explicit nil

### UnsetThumbnailTemplate
`func (o *TaskSettingsVM) UnsetThumbnailTemplate()`

UnsetThumbnailTemplate ensures that no value is present for ThumbnailTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


