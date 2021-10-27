# SubscriptionPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**SubscriptionCreate**](SubscriptionCreate.md) |  | [optional] 
**Delete** | Pointer to [**SubscriptionDelete**](SubscriptionDelete.md) |  | [optional] 
**Execute** | Pointer to [**SubscriptionExecute**](SubscriptionExecute.md) |  | [optional] 
**Get** | Pointer to [**SubscriptionGet**](SubscriptionGet.md) |  | [optional] 
**Update** | Pointer to [**SubscriptionUpdate**](SubscriptionUpdate.md) |  | [optional] 
**Administrate** | Pointer to [**SubscriptionAdministrate**](SubscriptionAdministrate.md) |  | [optional] 

## Methods

### NewSubscriptionPermission

`func NewSubscriptionPermission() *SubscriptionPermission`

NewSubscriptionPermission instantiates a new SubscriptionPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPermissionWithDefaults

`func NewSubscriptionPermissionWithDefaults() *SubscriptionPermission`

NewSubscriptionPermissionWithDefaults instantiates a new SubscriptionPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *SubscriptionPermission) GetCreate() SubscriptionCreate`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *SubscriptionPermission) GetCreateOk() (*SubscriptionCreate, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *SubscriptionPermission) SetCreate(v SubscriptionCreate)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *SubscriptionPermission) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelete

`func (o *SubscriptionPermission) GetDelete() SubscriptionDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *SubscriptionPermission) GetDeleteOk() (*SubscriptionDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *SubscriptionPermission) SetDelete(v SubscriptionDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *SubscriptionPermission) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetExecute

`func (o *SubscriptionPermission) GetExecute() SubscriptionExecute`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *SubscriptionPermission) GetExecuteOk() (*SubscriptionExecute, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *SubscriptionPermission) SetExecute(v SubscriptionExecute)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *SubscriptionPermission) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetGet

`func (o *SubscriptionPermission) GetGet() SubscriptionGet`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *SubscriptionPermission) GetGetOk() (*SubscriptionGet, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *SubscriptionPermission) SetGet(v SubscriptionGet)`

SetGet sets Get field to given value.

### HasGet

`func (o *SubscriptionPermission) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetUpdate

`func (o *SubscriptionPermission) GetUpdate() SubscriptionUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *SubscriptionPermission) GetUpdateOk() (*SubscriptionUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *SubscriptionPermission) SetUpdate(v SubscriptionUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *SubscriptionPermission) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAdministrate

`func (o *SubscriptionPermission) GetAdministrate() SubscriptionAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *SubscriptionPermission) GetAdministrateOk() (*SubscriptionAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *SubscriptionPermission) SetAdministrate(v SubscriptionAdministrate)`

SetAdministrate sets Administrate field to given value.

### HasAdministrate

`func (o *SubscriptionPermission) HasAdministrate() bool`

HasAdministrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


