# UserResultVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAuthenticated** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**AdminPermission** | Pointer to [**AdminPermission**](AdminPermission.md) |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**Provider** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Settings** | Pointer to [**UserSettingsVM**](UserSettingsVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUserResultVM

`func NewUserResultVM(t string, ) *UserResultVM`

NewUserResultVM instantiates a new UserResultVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResultVMWithDefaults

`func NewUserResultVMWithDefaults() *UserResultVM`

NewUserResultVMWithDefaults instantiates a new UserResultVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAuthenticated

`func (o *UserResultVM) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *UserResultVM) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *UserResultVM) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *UserResultVM) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### GetId

`func (o *UserResultVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResultVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResultVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserResultVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UserResultVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UserResultVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsAdmin

`func (o *UserResultVM) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UserResultVM) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *UserResultVM) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *UserResultVM) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAdminPermission

`func (o *UserResultVM) GetAdminPermission() AdminPermission`

GetAdminPermission returns the AdminPermission field if non-nil, zero value otherwise.

### GetAdminPermissionOk

`func (o *UserResultVM) GetAdminPermissionOk() (*AdminPermission, bool)`

GetAdminPermissionOk returns a tuple with the AdminPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermission

`func (o *UserResultVM) SetAdminPermission(v AdminPermission)`

SetAdminPermission sets AdminPermission field to given value.

### HasAdminPermission

`func (o *UserResultVM) HasAdminPermission() bool`

HasAdminPermission returns a boolean if a field has been set.

### GetSubscriptions

`func (o *UserResultVM) GetSubscriptions() []string`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *UserResultVM) GetSubscriptionsOk() (*[]string, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *UserResultVM) SetSubscriptions(v []string)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *UserResultVM) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### SetSubscriptionsNil

`func (o *UserResultVM) SetSubscriptionsNil(b bool)`

 SetSubscriptionsNil sets the value for Subscriptions to be an explicit nil

### UnsetSubscriptions
`func (o *UserResultVM) UnsetSubscriptions()`

UnsetSubscriptions ensures that no value is present for Subscriptions, not even an explicit nil
### GetGroups

`func (o *UserResultVM) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserResultVM) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserResultVM) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserResultVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *UserResultVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *UserResultVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetProvider

`func (o *UserResultVM) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserResultVM) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserResultVM) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserResultVM) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *UserResultVM) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UserResultVM) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetEmail

`func (o *UserResultVM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserResultVM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserResultVM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserResultVM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UserResultVM) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserResultVM) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *UserResultVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserResultVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserResultVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserResultVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserResultVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserResultVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUsername

`func (o *UserResultVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserResultVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserResultVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserResultVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserResultVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserResultVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetSettings

`func (o *UserResultVM) GetSettings() UserSettingsVM`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *UserResultVM) GetSettingsOk() (*UserSettingsVM, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *UserResultVM) SetSettings(v UserSettingsVM)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *UserResultVM) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetT

`func (o *UserResultVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UserResultVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UserResultVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


