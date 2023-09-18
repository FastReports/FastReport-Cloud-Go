# ContactVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **NullableString** |  | [optional] 
**EditedTime** | Pointer to **time.Time** |  | [optional] 
**EditorUserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContactVM

`func NewContactVM() *ContactVM`

NewContactVM instantiates a new ContactVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactVMWithDefaults

`func NewContactVMWithDefaults() *ContactVM`

NewContactVMWithDefaults instantiates a new ContactVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContactVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContactVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContactVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ContactVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ContactVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ContactVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *ContactVM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactVM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactVM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactVM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ContactVM) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ContactVM) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroups

`func (o *ContactVM) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ContactVM) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ContactVM) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ContactVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *ContactVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ContactVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetSubscriptionId

`func (o *ContactVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ContactVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ContactVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ContactVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *ContactVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *ContactVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetCreatedTime

`func (o *ContactVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ContactVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ContactVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ContactVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *ContactVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *ContactVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *ContactVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *ContactVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### SetCreatorUserIdNil

`func (o *ContactVM) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *ContactVM) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetEditedTime

`func (o *ContactVM) GetEditedTime() time.Time`

GetEditedTime returns the EditedTime field if non-nil, zero value otherwise.

### GetEditedTimeOk

`func (o *ContactVM) GetEditedTimeOk() (*time.Time, bool)`

GetEditedTimeOk returns a tuple with the EditedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTime

`func (o *ContactVM) SetEditedTime(v time.Time)`

SetEditedTime sets EditedTime field to given value.

### HasEditedTime

`func (o *ContactVM) HasEditedTime() bool`

HasEditedTime returns a boolean if a field has been set.

### GetEditorUserId

`func (o *ContactVM) GetEditorUserId() string`

GetEditorUserId returns the EditorUserId field if non-nil, zero value otherwise.

### GetEditorUserIdOk

`func (o *ContactVM) GetEditorUserIdOk() (*string, bool)`

GetEditorUserIdOk returns a tuple with the EditorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUserId

`func (o *ContactVM) SetEditorUserId(v string)`

SetEditorUserId sets EditorUserId field to given value.

### HasEditorUserId

`func (o *ContactVM) HasEditorUserId() bool`

HasEditorUserId returns a boolean if a field has been set.

### SetEditorUserIdNil

`func (o *ContactVM) SetEditorUserIdNil(b bool)`

 SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil

### UnsetEditorUserId
`func (o *ContactVM) UnsetEditorUserId()`

UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


