# CreateExportTemplateTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportParameters** | Pointer to **map[string]string** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewCreateExportTemplateTaskVM

`func NewCreateExportTemplateTaskVM() *CreateExportTemplateTaskVM`

NewCreateExportTemplateTaskVM instantiates a new CreateExportTemplateTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExportTemplateTaskVMWithDefaults

`func NewCreateExportTemplateTaskVMWithDefaults() *CreateExportTemplateTaskVM`

NewCreateExportTemplateTaskVMWithDefaults instantiates a new CreateExportTemplateTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportParameters

`func (o *CreateExportTemplateTaskVM) GetReportParameters() map[string]string`

GetReportParameters returns the ReportParameters field if non-nil, zero value otherwise.

### GetReportParametersOk

`func (o *CreateExportTemplateTaskVM) GetReportParametersOk() (*map[string]string, bool)`

GetReportParametersOk returns a tuple with the ReportParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportParameters

`func (o *CreateExportTemplateTaskVM) SetReportParameters(v map[string]string)`

SetReportParameters sets ReportParameters field to given value.

### HasReportParameters

`func (o *CreateExportTemplateTaskVM) HasReportParameters() bool`

HasReportParameters returns a boolean if a field has been set.

### SetReportParametersNil

`func (o *CreateExportTemplateTaskVM) SetReportParametersNil(b bool)`

 SetReportParametersNil sets the value for ReportParameters to be an explicit nil

### UnsetReportParameters
`func (o *CreateExportTemplateTaskVM) UnsetReportParameters()`

UnsetReportParameters ensures that no value is present for ReportParameters, not even an explicit nil
### GetName

`func (o *CreateExportTemplateTaskVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateExportTemplateTaskVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateExportTemplateTaskVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateExportTemplateTaskVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateExportTemplateTaskVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateExportTemplateTaskVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *CreateExportTemplateTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateExportTemplateTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateExportTemplateTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateExportTemplateTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateExportTemplateTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateExportTemplateTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *CreateExportTemplateTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateExportTemplateTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateExportTemplateTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateExportTemplateTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


