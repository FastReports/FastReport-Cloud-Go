# RegisterUserVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**AdminPermission** | Pointer to [**AdminPermission**](AdminPermission.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewRegisterUserVM

`func NewRegisterUserVM(t string, ) *RegisterUserVM`

NewRegisterUserVM instantiates a new RegisterUserVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserVMWithDefaults

`func NewRegisterUserVMWithDefaults() *RegisterUserVM`

NewRegisterUserVMWithDefaults instantiates a new RegisterUserVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegisterUserVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisterUserVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisterUserVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegisterUserVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RegisterUserVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RegisterUserVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubscriptions

`func (o *RegisterUserVM) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *RegisterUserVM) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *RegisterUserVM) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *RegisterUserVM) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *RegisterUserVM) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *RegisterUserVM) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetGroups

`func (o *RegisterUserVM) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RegisterUserVM) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RegisterUserVM) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RegisterUserVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *RegisterUserVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *RegisterUserVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetAdminPermission

`func (o *RegisterUserVM) GetAdminPermission() AdminPermission`

GetAdminPermission returns the AdminPermission field if non-nil, zero value otherwise.

### GetAdminPermissionOk

`func (o *RegisterUserVM) GetAdminPermissionOk() (*AdminPermission, bool)`

GetAdminPermissionOk returns a tuple with the AdminPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermission

`func (o *RegisterUserVM) SetAdminPermission(v AdminPermission)`

SetAdminPermission sets AdminPermission field to given value.

### HasAdminPermission

`func (o *RegisterUserVM) HasAdminPermission() bool`

HasAdminPermission returns a boolean if a field has been set.

### GetName

`func (o *RegisterUserVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterUserVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterUserVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterUserVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RegisterUserVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RegisterUserVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUsername

`func (o *RegisterUserVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterUserVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterUserVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RegisterUserVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RegisterUserVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RegisterUserVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *RegisterUserVM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterUserVM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterUserVM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RegisterUserVM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *RegisterUserVM) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RegisterUserVM) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPassword

`func (o *RegisterUserVM) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterUserVM) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterUserVM) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RegisterUserVM) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *RegisterUserVM) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *RegisterUserVM) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetIsAdmin

`func (o *RegisterUserVM) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *RegisterUserVM) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *RegisterUserVM) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *RegisterUserVM) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetProvider

`func (o *RegisterUserVM) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *RegisterUserVM) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *RegisterUserVM) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *RegisterUserVM) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *RegisterUserVM) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *RegisterUserVM) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetT

`func (o *RegisterUserVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RegisterUserVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RegisterUserVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


