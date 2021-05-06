# FilePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**FilePermission**](FilePermission.md) |  | [optional] 
**Groups** | Pointer to [**map[string]FilePermission**](FilePermission.md) |  | [optional] 
**Other** | Pointer to [**FilePermission**](FilePermission.md) |  | [optional] 
**Anon** | Pointer to [**FilePermission**](FilePermission.md) |  | [optional] 

## Methods

### NewFilePermissions

`func NewFilePermissions() *FilePermissions`

NewFilePermissions instantiates a new FilePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePermissionsWithDefaults

`func NewFilePermissionsWithDefaults() *FilePermissions`

NewFilePermissionsWithDefaults instantiates a new FilePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *FilePermissions) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *FilePermissions) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *FilePermissions) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *FilePermissions) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwner

`func (o *FilePermissions) GetOwner() FilePermission`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FilePermissions) GetOwnerOk() (*FilePermission, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FilePermissions) SetOwner(v FilePermission)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FilePermissions) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *FilePermissions) GetGroups() map[string]FilePermission`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FilePermissions) GetGroupsOk() (*map[string]FilePermission, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FilePermissions) SetGroups(v map[string]FilePermission)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FilePermissions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetOther

`func (o *FilePermissions) GetOther() FilePermission`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *FilePermissions) GetOtherOk() (*FilePermission, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *FilePermissions) SetOther(v FilePermission)`

SetOther sets Other field to given value.

### HasOther

`func (o *FilePermissions) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *FilePermissions) GetAnon() FilePermission`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *FilePermissions) GetAnonOk() (*FilePermission, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *FilePermissions) SetAnon(v FilePermission)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *FilePermissions) HasAnon() bool`

HasAnon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


