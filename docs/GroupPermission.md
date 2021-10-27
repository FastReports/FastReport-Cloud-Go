# GroupPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**GroupCreate**](GroupCreate.md) |  | [optional] 
**Delete** | Pointer to [**GroupDelete**](GroupDelete.md) |  | [optional] 
**Execute** | Pointer to [**GroupExecute**](GroupExecute.md) |  | [optional] 
**Get** | Pointer to [**GroupGet**](GroupGet.md) |  | [optional] 
**Update** | Pointer to [**GroupUpdate**](GroupUpdate.md) |  | [optional] 
**Administrate** | Pointer to [**GroupAdministrate**](GroupAdministrate.md) |  | [optional] 

## Methods

### NewGroupPermission

`func NewGroupPermission() *GroupPermission`

NewGroupPermission instantiates a new GroupPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPermissionWithDefaults

`func NewGroupPermissionWithDefaults() *GroupPermission`

NewGroupPermissionWithDefaults instantiates a new GroupPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *GroupPermission) GetCreate() GroupCreate`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *GroupPermission) GetCreateOk() (*GroupCreate, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *GroupPermission) SetCreate(v GroupCreate)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *GroupPermission) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelete

`func (o *GroupPermission) GetDelete() GroupDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *GroupPermission) GetDeleteOk() (*GroupDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *GroupPermission) SetDelete(v GroupDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *GroupPermission) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetExecute

`func (o *GroupPermission) GetExecute() GroupExecute`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *GroupPermission) GetExecuteOk() (*GroupExecute, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *GroupPermission) SetExecute(v GroupExecute)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *GroupPermission) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetGet

`func (o *GroupPermission) GetGet() GroupGet`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *GroupPermission) GetGetOk() (*GroupGet, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *GroupPermission) SetGet(v GroupGet)`

SetGet sets Get field to given value.

### HasGet

`func (o *GroupPermission) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetUpdate

`func (o *GroupPermission) GetUpdate() GroupUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *GroupPermission) GetUpdateOk() (*GroupUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *GroupPermission) SetUpdate(v GroupUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *GroupPermission) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAdministrate

`func (o *GroupPermission) GetAdministrate() GroupAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *GroupPermission) GetAdministrateOk() (*GroupAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *GroupPermission) SetAdministrate(v GroupAdministrate)`

SetAdministrate sets Administrate field to given value.

### HasAdministrate

`func (o *GroupPermission) HasAdministrate() bool`

HasAdministrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


