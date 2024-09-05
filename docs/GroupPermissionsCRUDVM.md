# GroupPermissionsCRUDVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**GroupPermissionCRUDVM**](GroupPermissionCRUDVM.md) |  | [optional] 
**Groups** | Pointer to [**map[string]GroupPermissionCRUDVM**](GroupPermissionCRUDVM.md) |  | [optional] 
**Other** | Pointer to [**GroupPermissionCRUDVM**](GroupPermissionCRUDVM.md) |  | [optional] 
**Anon** | Pointer to [**GroupPermissionCRUDVM**](GroupPermissionCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewGroupPermissionsCRUDVM

`func NewGroupPermissionsCRUDVM(t string, ) *GroupPermissionsCRUDVM`

NewGroupPermissionsCRUDVM instantiates a new GroupPermissionsCRUDVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPermissionsCRUDVMWithDefaults

`func NewGroupPermissionsCRUDVMWithDefaults() *GroupPermissionsCRUDVM`

NewGroupPermissionsCRUDVMWithDefaults instantiates a new GroupPermissionsCRUDVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *GroupPermissionsCRUDVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GroupPermissionsCRUDVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GroupPermissionsCRUDVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GroupPermissionsCRUDVM) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *GroupPermissionsCRUDVM) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *GroupPermissionsCRUDVM) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *GroupPermissionsCRUDVM) GetOwner() GroupPermissionCRUDVM`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GroupPermissionsCRUDVM) GetOwnerOk() (*GroupPermissionCRUDVM, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GroupPermissionsCRUDVM) SetOwner(v GroupPermissionCRUDVM)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GroupPermissionsCRUDVM) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *GroupPermissionsCRUDVM) GetGroups() map[string]GroupPermissionCRUDVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupPermissionsCRUDVM) GetGroupsOk() (*map[string]GroupPermissionCRUDVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupPermissionsCRUDVM) SetGroups(v map[string]GroupPermissionCRUDVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GroupPermissionsCRUDVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *GroupPermissionsCRUDVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *GroupPermissionsCRUDVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *GroupPermissionsCRUDVM) GetOther() GroupPermissionCRUDVM`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *GroupPermissionsCRUDVM) GetOtherOk() (*GroupPermissionCRUDVM, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *GroupPermissionsCRUDVM) SetOther(v GroupPermissionCRUDVM)`

SetOther sets Other field to given value.

### HasOther

`func (o *GroupPermissionsCRUDVM) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *GroupPermissionsCRUDVM) GetAnon() GroupPermissionCRUDVM`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *GroupPermissionsCRUDVM) GetAnonOk() (*GroupPermissionCRUDVM, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *GroupPermissionsCRUDVM) SetAnon(v GroupPermissionCRUDVM)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *GroupPermissionsCRUDVM) HasAnon() bool`

HasAnon returns a boolean if a field has been set.

### GetT

`func (o *GroupPermissionsCRUDVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GroupPermissionsCRUDVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GroupPermissionsCRUDVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


