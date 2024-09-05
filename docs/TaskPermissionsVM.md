# TaskPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**TaskPermissionsCRUDVM**](TaskPermissionsCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTaskPermissionsVM

`func NewTaskPermissionsVM(t string, ) *TaskPermissionsVM`

NewTaskPermissionsVM instantiates a new TaskPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPermissionsVMWithDefaults

`func NewTaskPermissionsVMWithDefaults() *TaskPermissionsVM`

NewTaskPermissionsVMWithDefaults instantiates a new TaskPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *TaskPermissionsVM) GetPermissions() TaskPermissionsCRUDVM`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *TaskPermissionsVM) GetPermissionsOk() (*TaskPermissionsCRUDVM, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *TaskPermissionsVM) SetPermissions(v TaskPermissionsCRUDVM)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *TaskPermissionsVM) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetT

`func (o *TaskPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TaskPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TaskPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


