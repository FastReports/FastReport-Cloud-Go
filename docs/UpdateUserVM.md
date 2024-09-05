# UpdateUserVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminPermission** | Pointer to [**AdminPermission**](AdminPermission.md) |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**IsAdmin** | Pointer to **NullableBool** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateUserVM

`func NewUpdateUserVM(t string, ) *UpdateUserVM`

NewUpdateUserVM instantiates a new UpdateUserVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserVMWithDefaults

`func NewUpdateUserVMWithDefaults() *UpdateUserVM`

NewUpdateUserVMWithDefaults instantiates a new UpdateUserVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminPermission

`func (o *UpdateUserVM) GetAdminPermission() AdminPermission`

GetAdminPermission returns the AdminPermission field if non-nil, zero value otherwise.

### GetAdminPermissionOk

`func (o *UpdateUserVM) GetAdminPermissionOk() (*AdminPermission, bool)`

GetAdminPermissionOk returns a tuple with the AdminPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermission

`func (o *UpdateUserVM) SetAdminPermission(v AdminPermission)`

SetAdminPermission sets AdminPermission field to given value.

### HasAdminPermission

`func (o *UpdateUserVM) HasAdminPermission() bool`

HasAdminPermission returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUserVM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserVM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserVM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserVM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateUserVM) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateUserVM) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetGroups

`func (o *UpdateUserVM) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UpdateUserVM) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UpdateUserVM) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UpdateUserVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *UpdateUserVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *UpdateUserVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetIsAdmin

`func (o *UpdateUserVM) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UpdateUserVM) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UpdateUserVM) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UpdateUserVM) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdminNil

`func (o *UpdateUserVM) SetIsAdminNil(b bool)`

 SetIsAdminNil sets the value for IsAdmin to be an explicit nil

### UnsetIsAdmin
`func (o *UpdateUserVM) UnsetIsAdmin()`

UnsetIsAdmin ensures that no value is present for IsAdmin, not even an explicit nil
### GetName

`func (o *UpdateUserVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateUserVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateUserVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPassword

`func (o *UpdateUserVM) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUserVM) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUserVM) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUserVM) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateUserVM) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateUserVM) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetProvider

`func (o *UpdateUserVM) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UpdateUserVM) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UpdateUserVM) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UpdateUserVM) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *UpdateUserVM) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UpdateUserVM) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetSubscriptions

`func (o *UpdateUserVM) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UpdateUserVM) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UpdateUserVM) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UpdateUserVM) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *UpdateUserVM) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *UpdateUserVM) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetUsername

`func (o *UpdateUserVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateUserVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateUserVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetT

`func (o *UpdateUserVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateUserVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateUserVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


