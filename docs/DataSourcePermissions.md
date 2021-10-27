# DataSourcePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**DataSourcePermission**](DataSourcePermission.md) |  | [optional] 
**Groups** | Pointer to [**map[string]DataSourcePermission**](DataSourcePermission.md) |  | [optional] 
**Other** | Pointer to [**DataSourcePermission**](DataSourcePermission.md) |  | [optional] 
**Anon** | Pointer to [**DataSourcePermission**](DataSourcePermission.md) |  | [optional] 

## Methods

### NewDataSourcePermissions

`func NewDataSourcePermissions() *DataSourcePermissions`

NewDataSourcePermissions instantiates a new DataSourcePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourcePermissionsWithDefaults

`func NewDataSourcePermissionsWithDefaults() *DataSourcePermissions`

NewDataSourcePermissionsWithDefaults instantiates a new DataSourcePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *DataSourcePermissions) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DataSourcePermissions) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DataSourcePermissions) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DataSourcePermissions) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *DataSourcePermissions) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *DataSourcePermissions) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *DataSourcePermissions) GetOwner() DataSourcePermission`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DataSourcePermissions) GetOwnerOk() (*DataSourcePermission, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DataSourcePermissions) SetOwner(v DataSourcePermission)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DataSourcePermissions) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *DataSourcePermissions) GetGroups() map[string]DataSourcePermission`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DataSourcePermissions) GetGroupsOk() (*map[string]DataSourcePermission, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DataSourcePermissions) SetGroups(v map[string]DataSourcePermission)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DataSourcePermissions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *DataSourcePermissions) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *DataSourcePermissions) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *DataSourcePermissions) GetOther() DataSourcePermission`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *DataSourcePermissions) GetOtherOk() (*DataSourcePermission, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *DataSourcePermissions) SetOther(v DataSourcePermission)`

SetOther sets Other field to given value.

### HasOther

`func (o *DataSourcePermissions) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *DataSourcePermissions) GetAnon() DataSourcePermission`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *DataSourcePermissions) GetAnonOk() (*DataSourcePermission, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *DataSourcePermissions) SetAnon(v DataSourcePermission)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *DataSourcePermissions) HasAnon() bool`

HasAnon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


