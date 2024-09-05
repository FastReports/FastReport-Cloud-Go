# GroupPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**GroupPermissionsCRUDVM**](GroupPermissionsCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewGroupPermissionsVM

`func NewGroupPermissionsVM(t string, ) *GroupPermissionsVM`

NewGroupPermissionsVM instantiates a new GroupPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPermissionsVMWithDefaults

`func NewGroupPermissionsVMWithDefaults() *GroupPermissionsVM`

NewGroupPermissionsVMWithDefaults instantiates a new GroupPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *GroupPermissionsVM) GetPermissions() GroupPermissionsCRUDVM`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GroupPermissionsVM) GetPermissionsOk() (*GroupPermissionsCRUDVM, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GroupPermissionsVM) SetPermissions(v GroupPermissionsCRUDVM)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GroupPermissionsVM) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetT

`func (o *GroupPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GroupPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GroupPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


