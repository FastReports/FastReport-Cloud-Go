# FilePermissionCRUDVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**FileCreate**](FileCreate.md) |  | [optional] 
**Delete** | Pointer to [**FileDelete**](FileDelete.md) |  | [optional] 
**Execute** | Pointer to [**FileExecute**](FileExecute.md) |  | [optional] 
**Get** | Pointer to [**FileGet**](FileGet.md) |  | [optional] 
**Update** | Pointer to [**FileUpdate**](FileUpdate.md) |  | [optional] 
**Administrate** | Pointer to [**FileAdministrate**](FileAdministrate.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFilePermissionCRUDVM

`func NewFilePermissionCRUDVM(t string, ) *FilePermissionCRUDVM`

NewFilePermissionCRUDVM instantiates a new FilePermissionCRUDVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePermissionCRUDVMWithDefaults

`func NewFilePermissionCRUDVMWithDefaults() *FilePermissionCRUDVM`

NewFilePermissionCRUDVMWithDefaults instantiates a new FilePermissionCRUDVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *FilePermissionCRUDVM) GetCreate() FileCreate`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *FilePermissionCRUDVM) GetCreateOk() (*FileCreate, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *FilePermissionCRUDVM) SetCreate(v FileCreate)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *FilePermissionCRUDVM) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelete

`func (o *FilePermissionCRUDVM) GetDelete() FileDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FilePermissionCRUDVM) GetDeleteOk() (*FileDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FilePermissionCRUDVM) SetDelete(v FileDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *FilePermissionCRUDVM) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetExecute

`func (o *FilePermissionCRUDVM) GetExecute() FileExecute`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *FilePermissionCRUDVM) GetExecuteOk() (*FileExecute, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *FilePermissionCRUDVM) SetExecute(v FileExecute)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *FilePermissionCRUDVM) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetGet

`func (o *FilePermissionCRUDVM) GetGet() FileGet`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *FilePermissionCRUDVM) GetGetOk() (*FileGet, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *FilePermissionCRUDVM) SetGet(v FileGet)`

SetGet sets Get field to given value.

### HasGet

`func (o *FilePermissionCRUDVM) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetUpdate

`func (o *FilePermissionCRUDVM) GetUpdate() FileUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *FilePermissionCRUDVM) GetUpdateOk() (*FileUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *FilePermissionCRUDVM) SetUpdate(v FileUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *FilePermissionCRUDVM) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAdministrate

`func (o *FilePermissionCRUDVM) GetAdministrate() FileAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *FilePermissionCRUDVM) GetAdministrateOk() (*FileAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *FilePermissionCRUDVM) SetAdministrate(v FileAdministrate)`

SetAdministrate sets Administrate field to given value.

### HasAdministrate

`func (o *FilePermissionCRUDVM) HasAdministrate() bool`

HasAdministrate returns a boolean if a field has been set.

### GetT

`func (o *FilePermissionCRUDVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FilePermissionCRUDVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FilePermissionCRUDVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


