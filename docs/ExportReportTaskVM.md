# ExportReportTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportParameters** | Pointer to **map[string]string** |  | [optional] 
**Format** | Pointer to [**ExportFormat**](ExportFormat.md) |  | [optional] 
**PagesCount** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewExportReportTaskVM

`func NewExportReportTaskVM() *ExportReportTaskVM`

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

### SetPagesCountNil

`func (o *ExportReportTaskVM) SetPagesCountNil(b bool)`

 SetPagesCountNil sets the value for PagesCount to be an explicit nil

### UnsetPagesCount
`func (o *ExportReportTaskVM) UnsetPagesCount()`

UnsetPagesCount ensures that no value is present for PagesCount, not even an explicit nil
### GetName

`func (o *ExportReportTaskVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportReportTaskVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportReportTaskVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExportReportTaskVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExportReportTaskVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExportReportTaskVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *ExportReportTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ExportReportTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ExportReportTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ExportReportTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *ExportReportTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *ExportReportTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *ExportReportTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportReportTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportReportTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *ExportReportTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


