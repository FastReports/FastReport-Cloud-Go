# TaskPermissionsCRUDVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**TaskPermissionCRUDVM**](TaskPermissionCRUDVM.md) |  | [optional] 
**Groups** | Pointer to [**map[string]TaskPermissionCRUDVM**](TaskPermissionCRUDVM.md) |  | [optional] 
**Other** | Pointer to [**TaskPermissionCRUDVM**](TaskPermissionCRUDVM.md) |  | [optional] 
**Anon** | Pointer to [**TaskPermissionCRUDVM**](TaskPermissionCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTaskPermissionsCRUDVM

`func NewTaskPermissionsCRUDVM(t string, ) *TaskPermissionsCRUDVM`

NewTaskPermissionsCRUDVM instantiates a new TaskPermissionsCRUDVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPermissionsCRUDVMWithDefaults

`func NewTaskPermissionsCRUDVMWithDefaults() *TaskPermissionsCRUDVM`

NewTaskPermissionsCRUDVMWithDefaults instantiates a new TaskPermissionsCRUDVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *TaskPermissionsCRUDVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *TaskPermissionsCRUDVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *TaskPermissionsCRUDVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *TaskPermissionsCRUDVM) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *TaskPermissionsCRUDVM) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *TaskPermissionsCRUDVM) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *TaskPermissionsCRUDVM) GetOwner() TaskPermissionCRUDVM`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TaskPermissionsCRUDVM) GetOwnerOk() (*TaskPermissionCRUDVM, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TaskPermissionsCRUDVM) SetOwner(v TaskPermissionCRUDVM)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TaskPermissionsCRUDVM) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *TaskPermissionsCRUDVM) GetGroups() map[string]TaskPermissionCRUDVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TaskPermissionsCRUDVM) GetGroupsOk() (*map[string]TaskPermissionCRUDVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TaskPermissionsCRUDVM) SetGroups(v map[string]TaskPermissionCRUDVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TaskPermissionsCRUDVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *TaskPermissionsCRUDVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *TaskPermissionsCRUDVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *TaskPermissionsCRUDVM) GetOther() TaskPermissionCRUDVM`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *TaskPermissionsCRUDVM) GetOtherOk() (*TaskPermissionCRUDVM, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *TaskPermissionsCRUDVM) SetOther(v TaskPermissionCRUDVM)`

SetOther sets Other field to given value.

### HasOther

`func (o *TaskPermissionsCRUDVM) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *TaskPermissionsCRUDVM) GetAnon() TaskPermissionCRUDVM`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *TaskPermissionsCRUDVM) GetAnonOk() (*TaskPermissionCRUDVM, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *TaskPermissionsCRUDVM) SetAnon(v TaskPermissionCRUDVM)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *TaskPermissionsCRUDVM) HasAnon() bool`

HasAnon returns a boolean if a field has been set.

### GetT

`func (o *TaskPermissionsCRUDVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TaskPermissionsCRUDVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TaskPermissionsCRUDVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


