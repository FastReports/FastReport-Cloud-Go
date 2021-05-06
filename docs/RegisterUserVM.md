# RegisterUserVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Subscriptions** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 
**AdminPermission** | Pointer to [**AdminPermission**](AdminPermission.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**IsAdmin** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewRegisterUserVM

`func NewRegisterUserVM() *RegisterUserVM`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


