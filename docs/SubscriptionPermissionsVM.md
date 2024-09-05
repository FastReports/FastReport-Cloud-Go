# SubscriptionPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to [**SubscriptionPermissionsCRUDVM**](SubscriptionPermissionsCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewSubscriptionPermissionsVM

`func NewSubscriptionPermissionsVM(t string, ) *SubscriptionPermissionsVM`

NewSubscriptionPermissionsVM instantiates a new SubscriptionPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPermissionsVMWithDefaults

`func NewSubscriptionPermissionsVMWithDefaults() *SubscriptionPermissionsVM`

NewSubscriptionPermissionsVMWithDefaults instantiates a new SubscriptionPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *SubscriptionPermissionsVM) GetPermissions() SubscriptionPermissionsCRUDVM`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SubscriptionPermissionsVM) GetPermissionsOk() (*SubscriptionPermissionsCRUDVM, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SubscriptionPermissionsVM) SetPermissions(v SubscriptionPermissionsCRUDVM)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SubscriptionPermissionsVM) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetT

`func (o *SubscriptionPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SubscriptionPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SubscriptionPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


