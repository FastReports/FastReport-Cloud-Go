# DefaultPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePermissions** | Pointer to [**FilePermissions**](FilePermissions.md) |  | [optional] 
**DataSourcePermissions** | Pointer to [**DataSourcePermissions**](DataSourcePermissions.md) |  | [optional] 
**GroupPermissions** | Pointer to [**GroupPermissions**](GroupPermissions.md) |  | [optional] 

## Methods

### NewDefaultPermissionsVM

`func NewDefaultPermissionsVM() *DefaultPermissionsVM`

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

`func (o *DefaultPermissionsVM) GetFilePermissions() FilePermissions`

GetFilePermissions returns the FilePermissions field if non-nil, zero value otherwise.

### GetFilePermissionsOk

`func (o *DefaultPermissionsVM) GetFilePermissionsOk() (*FilePermissions, bool)`

GetFilePermissionsOk returns a tuple with the FilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePermissions

`func (o *DefaultPermissionsVM) SetFilePermissions(v FilePermissions)`

SetFilePermissions sets FilePermissions field to given value.

### HasFilePermissions

`func (o *DefaultPermissionsVM) HasFilePermissions() bool`

HasFilePermissions returns a boolean if a field has been set.

### GetDataSourcePermissions

`func (o *DefaultPermissionsVM) GetDataSourcePermissions() DataSourcePermissions`

GetDataSourcePermissions returns the DataSourcePermissions field if non-nil, zero value otherwise.

### GetDataSourcePermissionsOk

`func (o *DefaultPermissionsVM) GetDataSourcePermissionsOk() (*DataSourcePermissions, bool)`

GetDataSourcePermissionsOk returns a tuple with the DataSourcePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourcePermissions

`func (o *DefaultPermissionsVM) SetDataSourcePermissions(v DataSourcePermissions)`

SetDataSourcePermissions sets DataSourcePermissions field to given value.

### HasDataSourcePermissions

`func (o *DefaultPermissionsVM) HasDataSourcePermissions() bool`

HasDataSourcePermissions returns a boolean if a field has been set.

### GetGroupPermissions

`func (o *DefaultPermissionsVM) GetGroupPermissions() GroupPermissions`

GetGroupPermissions returns the GroupPermissions field if non-nil, zero value otherwise.

### GetGroupPermissionsOk

`func (o *DefaultPermissionsVM) GetGroupPermissionsOk() (*GroupPermissions, bool)`

GetGroupPermissionsOk returns a tuple with the GroupPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPermissions

`func (o *DefaultPermissionsVM) SetGroupPermissions(v GroupPermissions)`

SetGroupPermissions sets GroupPermissions field to given value.

### HasGroupPermissions

`func (o *DefaultPermissionsVM) HasGroupPermissions() bool`

HasGroupPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


