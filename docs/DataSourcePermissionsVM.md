# DataSourcePermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**DataSourcePermissionsCRUDVM**](DataSourcePermissionsCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewDataSourcePermissionsVM

`func NewDataSourcePermissionsVM(t string, ) *DataSourcePermissionsVM`

NewDataSourcePermissionsVM instantiates a new DataSourcePermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourcePermissionsVMWithDefaults

`func NewDataSourcePermissionsVMWithDefaults() *DataSourcePermissionsVM`

NewDataSourcePermissionsVMWithDefaults instantiates a new DataSourcePermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *DataSourcePermissionsVM) GetPermissions() DataSourcePermissionsCRUDVM`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DataSourcePermissionsVM) GetPermissionsOk() (*DataSourcePermissionsCRUDVM, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DataSourcePermissionsVM) SetPermissions(v DataSourcePermissionsCRUDVM)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DataSourcePermissionsVM) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetT

`func (o *DataSourcePermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *DataSourcePermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *DataSourcePermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


