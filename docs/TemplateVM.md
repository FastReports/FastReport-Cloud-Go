# TemplateVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportInfo** | Pointer to [**ReportInfo**](ReportInfo.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **string** |  | [optional] 
**EditedTime** | Pointer to **time.Time** |  | [optional] 
**EditorUserId** | Pointer to **string** |  | [optional] 

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

### GetName

`func (o *TemplateVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *TemplateVM) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TemplateVM) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TemplateVM) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TemplateVM) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTags

`func (o *TemplateVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TemplateVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TemplateVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TemplateVM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIcon

`func (o *TemplateVM) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *TemplateVM) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *TemplateVM) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *TemplateVM) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetType

`func (o *TemplateVM) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateVM) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateVM) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateVM) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *TemplateVM) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TemplateVM) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TemplateVM) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *TemplateVM) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *TemplateVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *TemplateVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *TemplateVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *TemplateVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetStatus

`func (o *TemplateVM) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TemplateVM) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TemplateVM) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TemplateVM) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *TemplateVM) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *TemplateVM) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *TemplateVM) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *TemplateVM) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


