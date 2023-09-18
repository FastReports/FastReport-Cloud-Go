# ContactGroupVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **NullableString** |  | [optional] 
**EditedTime** | Pointer to **time.Time** |  | [optional] 
**EditorUserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactGroupVM

`func NewContactGroupVM() *ContactGroupVM`

NewContactGroupVM instantiates a new ContactGroupVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactGroupVMWithDefaults

`func NewContactGroupVMWithDefaults() *ContactGroupVM`

NewContactGroupVMWithDefaults instantiates a new ContactGroupVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactGroupVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactGroupVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactGroupVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactGroupVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContactGroupVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContactGroupVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ContactGroupVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactGroupVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactGroupVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactGroupVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ContactGroupVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContactGroupVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *ContactGroupVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ContactGroupVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ContactGroupVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ContactGroupVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *ContactGroupVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *ContactGroupVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetCreatedTime

`func (o *ContactGroupVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ContactGroupVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ContactGroupVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ContactGroupVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *ContactGroupVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *ContactGroupVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *ContactGroupVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *ContactGroupVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### SetCreatorUserIdNil

`func (o *ContactGroupVM) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *ContactGroupVM) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetEditedTime

`func (o *ContactGroupVM) GetEditedTime() time.Time`

GetEditedTime returns the EditedTime field if non-nil, zero value otherwise.

### GetEditedTimeOk

`func (o *ContactGroupVM) GetEditedTimeOk() (*time.Time, bool)`

GetEditedTimeOk returns a tuple with the EditedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTime

`func (o *ContactGroupVM) SetEditedTime(v time.Time)`

SetEditedTime sets EditedTime field to given value.

### HasEditedTime

`func (o *ContactGroupVM) HasEditedTime() bool`

HasEditedTime returns a boolean if a field has been set.

### GetEditorUserId

`func (o *ContactGroupVM) GetEditorUserId() string`

GetEditorUserId returns the EditorUserId field if non-nil, zero value otherwise.

### GetEditorUserIdOk

`func (o *ContactGroupVM) GetEditorUserIdOk() (*string, bool)`

GetEditorUserIdOk returns a tuple with the EditorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUserId

`func (o *ContactGroupVM) SetEditorUserId(v string)`

SetEditorUserId sets EditorUserId field to given value.

### HasEditorUserId

`func (o *ContactGroupVM) HasEditorUserId() bool`

HasEditorUserId returns a boolean if a field has been set.

### SetEditorUserIdNil

`func (o *ContactGroupVM) SetEditorUserIdNil(b bool)`

 SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil

### UnsetEditorUserId
`func (o *ContactGroupVM) UnsetEditorUserId()`

UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


