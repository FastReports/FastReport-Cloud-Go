# FileVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**FileType**](FileType.md) |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**FileStatus**](FileStatus.md) |  | [optional] 
**StatusReason** | Pointer to [**FileStatusReason**](FileStatusReason.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **NullableString** |  | [optional] 
**EditedTime** | Pointer to **time.Time** |  | [optional] 
**EditorUserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFileVM

`func NewFileVM() *FileVM`

NewFileVM instantiates a new FileVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileVMWithDefaults

`func NewFileVMWithDefaults() *FileVM`

NewFileVMWithDefaults instantiates a new FileVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentId

`func (o *FileVM) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FileVM) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FileVM) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *FileVM) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *FileVM) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *FileVM) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetTags

`func (o *FileVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FileVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FileVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FileVM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FileVM) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FileVM) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIcon

`func (o *FileVM) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FileVM) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FileVM) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FileVM) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *FileVM) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *FileVM) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetType

`func (o *FileVM) GetType() FileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileVM) GetTypeOk() (*FileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileVM) SetType(v FileType)`

SetType sets Type field to given value.

### HasType

`func (o *FileVM) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *FileVM) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FileVM) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FileVM) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FileVM) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *FileVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FileVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FileVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FileVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *FileVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *FileVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetStatus

`func (o *FileVM) GetStatus() FileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileVM) GetStatusOk() (*FileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileVM) SetStatus(v FileStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileVM) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *FileVM) GetStatusReason() FileStatusReason`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *FileVM) GetStatusReasonOk() (*FileStatusReason, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *FileVM) SetStatusReason(v FileStatusReason)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *FileVM) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetId

`func (o *FileVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FileVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *FileVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FileVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedTime

`func (o *FileVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *FileVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *FileVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *FileVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *FileVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *FileVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *FileVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *FileVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### SetCreatorUserIdNil

`func (o *FileVM) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *FileVM) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetEditedTime

`func (o *FileVM) GetEditedTime() time.Time`

GetEditedTime returns the EditedTime field if non-nil, zero value otherwise.

### GetEditedTimeOk

`func (o *FileVM) GetEditedTimeOk() (*time.Time, bool)`

GetEditedTimeOk returns a tuple with the EditedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTime

`func (o *FileVM) SetEditedTime(v time.Time)`

SetEditedTime sets EditedTime field to given value.

### HasEditedTime

`func (o *FileVM) HasEditedTime() bool`

HasEditedTime returns a boolean if a field has been set.

### GetEditorUserId

`func (o *FileVM) GetEditorUserId() string`

GetEditorUserId returns the EditorUserId field if non-nil, zero value otherwise.

### GetEditorUserIdOk

`func (o *FileVM) GetEditorUserIdOk() (*string, bool)`

GetEditorUserIdOk returns a tuple with the EditorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUserId

`func (o *FileVM) SetEditorUserId(v string)`

SetEditorUserId sets EditorUserId field to given value.

### HasEditorUserId

`func (o *FileVM) HasEditorUserId() bool`

HasEditorUserId returns a boolean if a field has been set.

### SetEditorUserIdNil

`func (o *FileVM) SetEditorUserIdNil(b bool)`

 SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil

### UnsetEditorUserId
`func (o *FileVM) UnsetEditorUserId()`

UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


