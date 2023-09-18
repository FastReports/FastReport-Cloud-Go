# TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**TaskPermission**](TaskPermission.md) |  | [optional] 
**Groups** | Pointer to [**map[string]TaskPermission**](TaskPermission.md) |  | [optional] 
**Other** | Pointer to [**TaskPermission**](TaskPermission.md) |  | [optional] 
**Anon** | Pointer to [**TaskPermission**](TaskPermission.md) |  | [optional] 

## Methods

### NewTaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions

`func NewTaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions() *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions`

NewTaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions instantiates a new TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissionsWithDefaults

`func NewTaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissionsWithDefaults() *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions`

NewTaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissionsWithDefaults instantiates a new TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetOwner() TaskPermission`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetOwnerOk() (*TaskPermission, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetOwner(v TaskPermission)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetGroups() map[string]TaskPermission`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetGroupsOk() (*map[string]TaskPermission, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetGroups(v map[string]TaskPermission)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetOther() TaskPermission`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetOtherOk() (*TaskPermission, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetOther(v TaskPermission)`

SetOther sets Other field to given value.

### HasOther

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetAnon() TaskPermission`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) GetAnonOk() (*TaskPermission, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) SetAnon(v TaskPermission)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *TaskPermissionTaskCreateTaskGetTaskUpdateTaskDeleteTaskExecuteTaskAdministratePermissions) HasAnon() bool`

HasAnon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


