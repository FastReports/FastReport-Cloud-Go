# UpdateFilePermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPermissions** | [**FilePermissionsCRUDVM**](FilePermissionsCRUDVM.md) |  | 
**Administrate** | [**FileAdministrate**](FileAdministrate.md) |  | 
**T** | **string** |  | 

## Methods

### NewUpdateFilePermissionsVM

`func NewUpdateFilePermissionsVM(newPermissions FilePermissionsCRUDVM, administrate FileAdministrate, t string, ) *UpdateFilePermissionsVM`

NewUpdateFilePermissionsVM instantiates a new UpdateFilePermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFilePermissionsVMWithDefaults

`func NewUpdateFilePermissionsVMWithDefaults() *UpdateFilePermissionsVM`

NewUpdateFilePermissionsVMWithDefaults instantiates a new UpdateFilePermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPermissions

`func (o *UpdateFilePermissionsVM) GetNewPermissions() FilePermissionsCRUDVM`

GetNewPermissions returns the NewPermissions field if non-nil, zero value otherwise.

### GetNewPermissionsOk

`func (o *UpdateFilePermissionsVM) GetNewPermissionsOk() (*FilePermissionsCRUDVM, bool)`

GetNewPermissionsOk returns a tuple with the NewPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissions

`func (o *UpdateFilePermissionsVM) SetNewPermissions(v FilePermissionsCRUDVM)`

SetNewPermissions sets NewPermissions field to given value.


### GetAdministrate

`func (o *UpdateFilePermissionsVM) GetAdministrate() FileAdministrate`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *UpdateFilePermissionsVM) GetAdministrateOk() (*FileAdministrate, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *UpdateFilePermissionsVM) SetAdministrate(v FileAdministrate)`

SetAdministrate sets Administrate field to given value.


### GetT

`func (o *UpdateFilePermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateFilePermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateFilePermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


