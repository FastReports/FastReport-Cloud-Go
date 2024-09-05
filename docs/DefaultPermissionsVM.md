# DefaultPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePermissions** | Pointer to [**FilePermissionsCRUDVM**](FilePermissionsCRUDVM.md) |  | [optional] 
**DataSourcePermissions** | Pointer to [**DataSourcePermissionsCRUDVM**](DataSourcePermissionsCRUDVM.md) |  | [optional] 
**GroupPermissions** | Pointer to [**GroupPermissionsCRUDVM**](GroupPermissionsCRUDVM.md) |  | [optional] 
**TaskPermissions** | Pointer to [**TaskPermissionsCRUDVM**](TaskPermissionsCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewDefaultPermissionsVM

`func NewDefaultPermissionsVM(t string, ) *DefaultPermissionsVM`

NewDefaultPermissionsVM instantiates a new DefaultPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultPermissionsVMWithDefaults

`func NewDefaultPermissionsVMWithDefaults() *DefaultPermissionsVM`

NewDefaultPermissionsVMWithDefaults instantiates a new DefaultPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePermissions

`func (o *DefaultPermissionsVM) GetFilePermissions() FilePermissionsCRUDVM`

GetFilePermissions returns the FilePermissions field if non-nil, zero value otherwise.

### GetFilePermissionsOk

`func (o *DefaultPermissionsVM) GetFilePermissionsOk() (*FilePermissionsCRUDVM, bool)`

GetFilePermissionsOk returns a tuple with the FilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePermissions

`func (o *DefaultPermissionsVM) SetFilePermissions(v FilePermissionsCRUDVM)`

SetFilePermissions sets FilePermissions field to given value.

### HasFilePermissions

`func (o *DefaultPermissionsVM) HasFilePermissions() bool`

HasFilePermissions returns a boolean if a field has been set.

### GetDataSourcePermissions

`func (o *DefaultPermissionsVM) GetDataSourcePermissions() DataSourcePermissionsCRUDVM`

GetDataSourcePermissions returns the DataSourcePermissions field if non-nil, zero value otherwise.

### GetDataSourcePermissionsOk

`func (o *DefaultPermissionsVM) GetDataSourcePermissionsOk() (*DataSourcePermissionsCRUDVM, bool)`

GetDataSourcePermissionsOk returns a tuple with the DataSourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourcePermissions

`func (o *DefaultPermissionsVM) SetDataSourcePermissions(v DataSourcePermissionsCRUDVM)`

SetDataSourcePermissions sets DataSourcePermissions field to given value.

### HasDataSourcePermissions

`func (o *DefaultPermissionsVM) HasDataSourcePermissions() bool`

HasDataSourcePermissions returns a boolean if a field has been set.

### GetGroupPermissions

`func (o *DefaultPermissionsVM) GetGroupPermissions() GroupPermissionsCRUDVM`

GetGroupPermissions returns the GroupPermissions field if non-nil, zero value otherwise.

### GetGroupPermissionsOk

`func (o *DefaultPermissionsVM) GetGroupPermissionsOk() (*GroupPermissionsCRUDVM, bool)`

GetGroupPermissionsOk returns a tuple with the GroupPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPermissions

`func (o *DefaultPermissionsVM) SetGroupPermissions(v GroupPermissionsCRUDVM)`

SetGroupPermissions sets GroupPermissions field to given value.

### HasGroupPermissions

`func (o *DefaultPermissionsVM) HasGroupPermissions() bool`

HasGroupPermissions returns a boolean if a field has been set.

### GetTaskPermissions

`func (o *DefaultPermissionsVM) GetTaskPermissions() TaskPermissionsCRUDVM`

GetTaskPermissions returns the TaskPermissions field if non-nil, zero value otherwise.

### GetTaskPermissionsOk

`func (o *DefaultPermissionsVM) GetTaskPermissionsOk() (*TaskPermissionsCRUDVM, bool)`

GetTaskPermissionsOk returns a tuple with the TaskPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskPermissions

`func (o *DefaultPermissionsVM) SetTaskPermissions(v TaskPermissionsCRUDVM)`

SetTaskPermissions sets TaskPermissions field to given value.

### HasTaskPermissions

`func (o *DefaultPermissionsVM) HasTaskPermissions() bool`

HasTaskPermissions returns a boolean if a field has been set.

### GetT

`func (o *DefaultPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *DefaultPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *DefaultPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


