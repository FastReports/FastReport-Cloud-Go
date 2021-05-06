# FilePermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**FilePermissions**](FilePermissions.md) |  | [optional] 

## Methods

### NewFilePermissionsVM

`func NewFilePermissionsVM() *FilePermissionsVM`

NewFilePermissionsVM instantiates a new FilePermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePermissionsVMWithDefaults

`func NewFilePermissionsVMWithDefaults() *FilePermissionsVM`

NewFilePermissionsVMWithDefaults instantiates a new FilePermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *FilePermissionsVM) GetPermissions() FilePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *FilePermissionsVM) GetPermissionsOk() (*FilePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *FilePermissionsVM) SetPermissions(v FilePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *FilePermissionsVM) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


