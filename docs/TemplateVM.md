# TemplateVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportInfo** | Pointer to [**ReportInfo**](ReportInfo.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **NullableString** |  | [optional] 
**EditedTime** | Pointer to **time.Time** |  | [optional] 
**EditorUserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTemplateVM

`func NewTemplateVM() *TemplateVM`

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

### GetId

`func (o *TemplateVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TemplateVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TemplateVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedTime

`func (o *TemplateVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TemplateVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TemplateVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *TemplateVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *TemplateVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *TemplateVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *TemplateVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *TemplateVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### SetCreatorUserIdNil

`func (o *TemplateVM) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *TemplateVM) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetEditedTime

`func (o *TemplateVM) GetEditedTime() time.Time`

GetEditedTime returns the EditedTime field if non-nil, zero value otherwise.

### GetEditedTimeOk

`func (o *TemplateVM) GetEditedTimeOk() (*time.Time, bool)`

GetEditedTimeOk returns a tuple with the EditedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTime

`func (o *TemplateVM) SetEditedTime(v time.Time)`

SetEditedTime sets EditedTime field to given value.

### HasEditedTime

`func (o *TemplateVM) HasEditedTime() bool`

HasEditedTime returns a boolean if a field has been set.

### GetEditorUserId

`func (o *TemplateVM) GetEditorUserId() string`

GetEditorUserId returns the EditorUserId field if non-nil, zero value otherwise.

### GetEditorUserIdOk

`func (o *TemplateVM) GetEditorUserIdOk() (*string, bool)`

GetEditorUserIdOk returns a tuple with the EditorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUserId

`func (o *TemplateVM) SetEditorUserId(v string)`

SetEditorUserId sets EditorUserId field to given value.

### HasEditorUserId

`func (o *TemplateVM) HasEditorUserId() bool`

HasEditorUserId returns a boolean if a field has been set.

### SetEditorUserIdNil

`func (o *TemplateVM) SetEditorUserIdNil(b bool)`

 SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil

### UnsetEditorUserId
`func (o *TemplateVM) UnsetEditorUserId()`

UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


