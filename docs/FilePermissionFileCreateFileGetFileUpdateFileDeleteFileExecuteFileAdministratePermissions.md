# FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**FilePermission**](FilePermission.md) |  | [optional] 
**Groups** | Pointer to [**map[string]FilePermission**](FilePermission.md) |  | [optional] 
**Other** | Pointer to [**FilePermission**](FilePermission.md) |  | [optional] 
**Anon** | Pointer to [**FilePermission**](FilePermission.md) |  | [optional] 

## Methods

### NewFilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions

`func NewFilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions() *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions`

NewFilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions instantiates a new FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissionsWithDefaults

`func NewFilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissionsWithDefaults() *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions`

NewFilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissionsWithDefaults instantiates a new FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetOwner() FilePermission`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetOwnerOk() (*FilePermission, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetOwner(v FilePermission)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetGroups() map[string]FilePermission`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetGroupsOk() (*map[string]FilePermission, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetGroups(v map[string]FilePermission)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetOther() FilePermission`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetOtherOk() (*FilePermission, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetOther(v FilePermission)`

SetOther sets Other field to given value.

### HasOther

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetAnon() FilePermission`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) GetAnonOk() (*FilePermission, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) SetAnon(v FilePermission)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *FilePermissionFileCreateFileGetFileUpdateFileDeleteFileExecuteFileAdministratePermissions) HasAnon() bool`

HasAnon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


