# AdminSubscriptionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPermissions** | Pointer to [**DefaultPermissionsVM**](DefaultPermissionsVM.md) |  | [optional] 
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAdminSubscriptionVM

`func NewAdminSubscriptionVM(t string, ) *AdminSubscriptionVM`

NewAdminSubscriptionVM instantiates a new AdminSubscriptionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminSubscriptionVMWithDefaults

`func NewAdminSubscriptionVMWithDefaults() *AdminSubscriptionVM`

NewAdminSubscriptionVMWithDefaults instantiates a new AdminSubscriptionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPermissions

`func (o *AdminSubscriptionVM) GetDefaultPermissions() DefaultPermissionsVM`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *AdminSubscriptionVM) GetDefaultPermissionsOk() (*DefaultPermissionsVM, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *AdminSubscriptionVM) SetDefaultPermissions(v DefaultPermissionsVM)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *AdminSubscriptionVM) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetOwnerId

`func (o *AdminSubscriptionVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AdminSubscriptionVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AdminSubscriptionVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AdminSubscriptionVM) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *AdminSubscriptionVM) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *AdminSubscriptionVM) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetT

`func (o *AdminSubscriptionVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AdminSubscriptionVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AdminSubscriptionVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


