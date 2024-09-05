# UserVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**AdminPermission** | Pointer to [**AdminPermission**](AdminPermission.md) |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUserVM

`func NewUserVM(t string, ) *UserVM`

NewUserVM instantiates a new UserVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserVMWithDefaults

`func NewUserVMWithDefaults() *UserVM`

NewUserVMWithDefaults instantiates a new UserVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubscriptions

`func (o *UserVM) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UserVM) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UserVM) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UserVM) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *UserVM) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *UserVM) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetGroups

`func (o *UserVM) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserVM) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserVM) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *UserVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *UserVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetAdminPermission

`func (o *UserVM) GetAdminPermission() AdminPermission`

GetAdminPermission returns the AdminPermission field if non-nil, zero value otherwise.

### GetAdminPermissionOk

`func (o *UserVM) GetAdminPermissionOk() (*AdminPermission, bool)`

GetAdminPermissionOk returns a tuple with the AdminPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermission

`func (o *UserVM) SetAdminPermission(v AdminPermission)`

SetAdminPermission sets AdminPermission field to given value.

### HasAdminPermission

`func (o *UserVM) HasAdminPermission() bool`

HasAdminPermission returns a boolean if a field has been set.

### GetIsAdmin

`func (o *UserVM) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UserVM) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UserVM) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UserVM) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetName

`func (o *UserVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUsername

`func (o *UserVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *UserVM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserVM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserVM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserVM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserVM) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserVM) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetCreatedTime

`func (o *UserVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *UserVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *UserVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *UserVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetT

`func (o *UserVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UserVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UserVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


