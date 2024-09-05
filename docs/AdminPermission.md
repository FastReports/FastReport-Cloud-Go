# AdminPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**AdminCreate**](AdminCreate.md) |  | [optional] 
**Delete** | Pointer to [**AdminDelete**](AdminDelete.md) |  | [optional] 
**Execute** | Pointer to [**AdminExecute**](AdminExecute.md) |  | [optional] 
**Get** | Pointer to [**AdminGet**](AdminGet.md) |  | [optional] 
**Update** | Pointer to [**AdminUpdate**](AdminUpdate.md) |  | [optional] 
**Administrate** | Pointer to [**AdminAdministrate**](AdminAdministrate.md) |  | [optional] 

## Methods

### NewAdminPermission

`func NewAdminPermission() *AdminPermission`

NewAdminPermission instantiates a new AdminPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminPermissionWithDefaults

`func NewAdminPermissionWithDefaults() *AdminPermission`

NewAdminPermissionWithDefaults instantiates a new AdminPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *AdminPermission) GetCreate() AdminCreate`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *AdminPermission) GetCreateOk() (*AdminCreate, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *AdminPermission) SetCreate(v AdminCreate)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *AdminPermission) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelete

`func (o *AdminPermission) GetDelete() AdminDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AdminPermission) GetDeleteOk() (*AdminDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AdminPermission) SetDelete(v AdminDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AdminPermission) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetExecute

`func (o *AdminPermission) GetExecute() AdminExecute`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *AdminPermission) GetExecuteOk() (*AdminExecute, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *AdminPermission) SetExecute(v AdminExecute)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *AdminPermission) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetGet

`func (o *AdminPermission) GetGet() AdminGet`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *AdminPermission) GetGetOk() (*AdminGet, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *AdminPermission) SetGet(v AdminGet)`

SetGet sets Get field to given value.

### HasGet

`func (o *AdminPermission) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetUpdate

`func (o *AdminPermission) GetUpdate() AdminUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *AdminPermission) GetUpdateOk() (*AdminUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *AdminPermission) SetUpdate(v AdminUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *AdminPermission) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAdministrate

`func (o *AdminPermission) GetAdministrate() AdminAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *AdminPermission) GetAdministrateOk() (*AdminAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *AdminPermission) SetAdministrate(v AdminAdministrate)`

SetAdministrate sets Administrate field to given value.

### HasAdministrate

`func (o *AdminPermission) HasAdministrate() bool`

HasAdministrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


