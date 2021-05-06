# FilePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to **int32** |  | [optional] 
**Delete** | Pointer to **int32** |  | [optional] 
**Execute** | Pointer to **int32** |  | [optional] 
**Get** | Pointer to **int32** |  | [optional] 
**Update** | Pointer to **int32** |  | [optional] 
**Administrate** | Pointer to **int32** |  | [optional] 

## Methods

### NewFilePermission

`func NewFilePermission() *FilePermission`

NewFilePermission instantiates a new FilePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePermissionWithDefaults

`func NewFilePermissionWithDefaults() *FilePermission`

NewFilePermissionWithDefaults instantiates a new FilePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *FilePermission) GetCreate() int32`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FilePermission) GetCreateOk() (*int32, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FilePermission) SetCreate(v int32)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FilePermission) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelete

`func (o *FilePermission) GetDelete() int32`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FilePermission) GetDeleteOk() (*int32, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FilePermission) SetDelete(v int32)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *FilePermission) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetExecute

`func (o *FilePermission) GetExecute() int32`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *FilePermission) GetExecuteOk() (*int32, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *FilePermission) SetExecute(v int32)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *FilePermission) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetGet

`func (o *FilePermission) GetGet() int32`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *FilePermission) GetGetOk() (*int32, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *FilePermission) SetGet(v int32)`

SetGet sets Get field to given value.

### HasGet

`func (o *FilePermission) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetUpdate

`func (o *FilePermission) GetUpdate() int32`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FilePermission) GetUpdateOk() (*int32, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FilePermission) SetUpdate(v int32)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FilePermission) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAdministrate

`func (o *FilePermission) GetAdministrate() int32`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *FilePermission) GetAdministrateOk() (*int32, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *FilePermission) SetAdministrate(v int32)`

SetAdministrate sets Administrate field to given value.

### HasAdministrate

`func (o *FilePermission) HasAdministrate() bool`

HasAdministrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


