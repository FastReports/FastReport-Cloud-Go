# GroupPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**GroupPermission**](GroupPermission.md) |  | [optional] 
**Groups** | Pointer to [**map[string]GroupPermission**](GroupPermission.md) |  | [optional] 
**Other** | Pointer to [**GroupPermission**](GroupPermission.md) |  | [optional] 
**Anon** | Pointer to [**GroupPermission**](GroupPermission.md) |  | [optional] 

## Methods

### NewGroupPermissions

`func NewGroupPermissions() *GroupPermissions`

NewGroupPermissions instantiates a new GroupPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPermissionsWithDefaults

`func NewGroupPermissionsWithDefaults() *GroupPermissions`

NewGroupPermissionsWithDefaults instantiates a new GroupPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *GroupPermissions) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GroupPermissions) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GroupPermissions) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GroupPermissions) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *GroupPermissions) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *GroupPermissions) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *GroupPermissions) GetOwner() GroupPermission`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GroupPermissions) GetOwnerOk() (*GroupPermission, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GroupPermissions) SetOwner(v GroupPermission)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GroupPermissions) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *GroupPermissions) GetGroups() map[string]GroupPermission`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupPermissions) GetGroupsOk() (*map[string]GroupPermission, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupPermissions) SetGroups(v map[string]GroupPermission)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GroupPermissions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *GroupPermissions) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *GroupPermissions) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *GroupPermissions) GetOther() GroupPermission`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *GroupPermissions) GetOtherOk() (*GroupPermission, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *GroupPermissions) SetOther(v GroupPermission)`

SetOther sets Other field to given value.

### HasOther

`func (o *GroupPermissions) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *GroupPermissions) GetAnon() GroupPermission`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *GroupPermissions) GetAnonOk() (*GroupPermission, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *GroupPermissions) SetAnon(v GroupPermission)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *GroupPermissions) HasAnon() bool`

HasAnon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


