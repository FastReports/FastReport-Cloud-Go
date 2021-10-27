# DataSourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**DataSourceCreate**](DataSourceCreate.md) |  | [optional] 
**Delete** | Pointer to [**DataSourceDelete**](DataSourceDelete.md) |  | [optional] 
**Execute** | Pointer to [**DataSourceExecute**](DataSourceExecute.md) |  | [optional] 
**Get** | Pointer to [**DataSourceGet**](DataSourceGet.md) |  | [optional] 
**Update** | Pointer to [**DataSourceUpdate**](DataSourceUpdate.md) |  | [optional] 
**Administrate** | Pointer to [**DataSourceAdministrate**](DataSourceAdministrate.md) |  | [optional] 

## Methods

### NewDataSourcePermission

`func NewDataSourcePermission() *DataSourcePermission`

NewDataSourcePermission instantiates a new DataSourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourcePermissionWithDefaults

`func NewDataSourcePermissionWithDefaults() *DataSourcePermission`

NewDataSourcePermissionWithDefaults instantiates a new DataSourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *DataSourcePermission) GetCreate() DataSourceCreate`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *DataSourcePermission) GetCreateOk() (*DataSourceCreate, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *DataSourcePermission) SetCreate(v DataSourceCreate)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *DataSourcePermission) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetDelete

`func (o *DataSourcePermission) GetDelete() DataSourceDelete`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *DataSourcePermission) GetDeleteOk() (*DataSourceDelete, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *DataSourcePermission) SetDelete(v DataSourceDelete)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *DataSourcePermission) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetExecute

`func (o *DataSourcePermission) GetExecute() DataSourceExecute`

GetExecute returns the Execute field if non-nil, zero value otherwise.

### GetExecuteOk

`func (o *DataSourcePermission) GetExecuteOk() (*DataSourceExecute, bool)`

GetExecuteOk returns a tuple with the Execute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecute

`func (o *DataSourcePermission) SetExecute(v DataSourceExecute)`

SetExecute sets Execute field to given value.

### HasExecute

`func (o *DataSourcePermission) HasExecute() bool`

HasExecute returns a boolean if a field has been set.

### GetGet

`func (o *DataSourcePermission) GetGet() DataSourceGet`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *DataSourcePermission) GetGetOk() (*DataSourceGet, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *DataSourcePermission) SetGet(v DataSourceGet)`

SetGet sets Get field to given value.

### HasGet

`func (o *DataSourcePermission) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetUpdate

`func (o *DataSourcePermission) GetUpdate() DataSourceUpdate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *DataSourcePermission) GetUpdateOk() (*DataSourceUpdate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *DataSourcePermission) SetUpdate(v DataSourceUpdate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *DataSourcePermission) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetAdministrate

`func (o *DataSourcePermission) GetAdministrate() DataSourceAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *DataSourcePermission) GetAdministrateOk() (*DataSourceAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *DataSourcePermission) SetAdministrate(v DataSourceAdministrate)`

SetAdministrate sets Administrate field to given value.

### HasAdministrate

`func (o *DataSourcePermission) HasAdministrate() bool`

HasAdministrate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


