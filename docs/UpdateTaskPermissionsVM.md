# UpdateTaskPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Administrate** | [**TaskAdministrate**](TaskAdministrate.md) |  | 
**NewPermissions** | [**TaskPermissionsCRUDVM**](TaskPermissionsCRUDVM.md) |  | 
**T** | **string** |  | 

## Methods

### NewUpdateTaskPermissionsVM

`func NewUpdateTaskPermissionsVM(administrate TaskAdministrate, newPermissions TaskPermissionsCRUDVM, t string, ) *UpdateTaskPermissionsVM`

NewUpdateTaskPermissionsVM instantiates a new UpdateTaskPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskPermissionsVMWithDefaults

`func NewUpdateTaskPermissionsVMWithDefaults() *UpdateTaskPermissionsVM`

NewUpdateTaskPermissionsVMWithDefaults instantiates a new UpdateTaskPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministrate

`func (o *UpdateTaskPermissionsVM) GetAdministrate() TaskAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *UpdateTaskPermissionsVM) GetAdministrateOk() (*TaskAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *UpdateTaskPermissionsVM) SetAdministrate(v TaskAdministrate)`

SetAdministrate sets Administrate field to given value.


### GetNewPermissions

`func (o *UpdateTaskPermissionsVM) GetNewPermissions() TaskPermissionsCRUDVM`

GetNewPermissions returns the NewPermissions field if non-nil, zero value otherwise.

### GetNewPermissionsOk

`func (o *UpdateTaskPermissionsVM) GetNewPermissionsOk() (*TaskPermissionsCRUDVM, bool)`

GetNewPermissionsOk returns a tuple with the NewPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissions

`func (o *UpdateTaskPermissionsVM) SetNewPermissions(v TaskPermissionsCRUDVM)`

SetNewPermissions sets NewPermissions field to given value.


### GetT

`func (o *UpdateTaskPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateTaskPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateTaskPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


