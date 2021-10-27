# CreatePrepareTemplateTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exports** | Pointer to [**[]CreateExportReportTaskVM**](CreateExportReportTaskVM.md) |  | [optional] 
**PagesCount** | Pointer to **NullableInt32** |  | [optional] 
**ReportParameters** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewCreatePrepareTemplateTaskVM

`func NewCreatePrepareTemplateTaskVM() *CreatePrepareTemplateTaskVM`

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

### SetPagesCountNil

`func (o *CreatePrepareTemplateTaskVM) SetPagesCountNil(b bool)`

 SetPagesCountNil sets the value for PagesCount to be an explicit nil

### UnsetPagesCount
`func (o *CreatePrepareTemplateTaskVM) UnsetPagesCount()`

UnsetPagesCount ensures that no value is present for PagesCount, not even an explicit nil
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
### GetName

`func (o *CreatePrepareTemplateTaskVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePrepareTemplateTaskVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePrepareTemplateTaskVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePrepareTemplateTaskVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatePrepareTemplateTaskVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatePrepareTemplateTaskVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *CreatePrepareTemplateTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreatePrepareTemplateTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreatePrepareTemplateTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreatePrepareTemplateTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreatePrepareTemplateTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreatePrepareTemplateTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *CreatePrepareTemplateTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatePrepareTemplateTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatePrepareTemplateTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *CreatePrepareTemplateTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


