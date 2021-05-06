# UpdateFilePermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPermissions** | [**FilePermissions**](FilePermissions.md) |  | 
**Administrate** | **int32** |  | 

## Methods

### NewUpdateFilePermissionsVM

`func NewUpdateFilePermissionsVM(newPermissions FilePermissions, administrate int32, ) *UpdateFilePermissionsVM`

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

`func (o *UpdateFilePermissionsVM) GetNewPermissions() FilePermissions`

GetNewPermissions returns the NewPermissions field if non-nil, zero value otherwise.

### GetNewPermissionsOk

`func (o *UpdateFilePermissionsVM) GetNewPermissionsOk() (*FilePermissions, bool)`

GetNewPermissionsOk returns a tuple with the NewPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPermissions

`func (o *UpdateFilePermissionsVM) SetNewPermissions(v FilePermissions)`

SetNewPermissions sets NewPermissions field to given value.


### GetAdministrate

`func (o *UpdateFilePermissionsVM) GetAdministrate() int32`

GetAdministrate returns the Administrate field if non-nil, zero value otherwise.

### GetAdministrateOk

`func (o *UpdateFilePermissionsVM) GetAdministrateOk() (*int32, bool)`

GetAdministrateOk returns a tuple with the Administrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrate

`func (o *UpdateFilePermissionsVM) SetAdministrate(v int32)`

SetAdministrate sets Administrate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


