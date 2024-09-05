# UpdateGroupPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPermissions** | [**GroupPermissionsCRUDVM**](GroupPermissionsCRUDVM.md) |  | 
**Administrate** | [**GroupAdministrate**](GroupAdministrate.md) |  | 
**T** | **string** |  | 

## Methods

### NewUpdateGroupPermissionsVM

`func NewUpdateGroupPermissionsVM(newPermissions GroupPermissionsCRUDVM, administrate GroupAdministrate, t string, ) *UpdateGroupPermissionsVM`

NewUpdateGroupPermissionsVM instantiates a new UpdateGroupPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGroupPermissionsVMWithDefaults

`func NewUpdateGroupPermissionsVMWithDefaults() *UpdateGroupPermissionsVM`

NewUpdateGroupPermissionsVMWithDefaults instantiates a new UpdateGroupPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPermissions

`func (o *UpdateGroupPermissionsVM) GetNewPermissions() GroupPermissionsCRUDVM`

GetNewPermissions returns the NewPermissions field if non-nil, zero value otherwise.

### GetNewPermissionsOk

`func (o *UpdateGroupPermissionsVM) GetNewPermissionsOk() (*GroupPermissionsCRUDVM, bool)`

GetNewPermissionsOk returns a tuple with the NewPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissions

`func (o *UpdateGroupPermissionsVM) SetNewPermissions(v GroupPermissionsCRUDVM)`

SetNewPermissions sets NewPermissions field to given value.


### GetAdministrate

`func (o *UpdateGroupPermissionsVM) GetAdministrate() GroupAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *UpdateGroupPermissionsVM) GetAdministrateOk() (*GroupAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *UpdateGroupPermissionsVM) SetAdministrate(v GroupAdministrate)`

SetAdministrate sets Administrate field to given value.


### GetT

`func (o *UpdateGroupPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateGroupPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateGroupPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


