# SubscriptionPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**SubscriptionPermission**](SubscriptionPermission.md) |  | [optional] 
**Groups** | Pointer to [**map[string]SubscriptionPermission**](SubscriptionPermission.md) |  | [optional] 
**Other** | Pointer to [**SubscriptionPermission**](SubscriptionPermission.md) |  | [optional] 
**Anon** | Pointer to [**SubscriptionPermission**](SubscriptionPermission.md) |  | [optional] 

## Methods

### NewSubscriptionPermissions

`func NewSubscriptionPermissions() *SubscriptionPermissions`

NewSubscriptionPermissions instantiates a new SubscriptionPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPermissionsWithDefaults

`func NewSubscriptionPermissionsWithDefaults() *SubscriptionPermissions`

NewSubscriptionPermissionsWithDefaults instantiates a new SubscriptionPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *SubscriptionPermissions) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SubscriptionPermissions) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SubscriptionPermissions) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SubscriptionPermissions) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *SubscriptionPermissions) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *SubscriptionPermissions) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *SubscriptionPermissions) GetOwner() SubscriptionPermission`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SubscriptionPermissions) GetOwnerOk() (*SubscriptionPermission, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SubscriptionPermissions) SetOwner(v SubscriptionPermission)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SubscriptionPermissions) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *SubscriptionPermissions) GetGroups() map[string]SubscriptionPermission`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SubscriptionPermissions) GetGroupsOk() (*map[string]SubscriptionPermission, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SubscriptionPermissions) SetGroups(v map[string]SubscriptionPermission)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SubscriptionPermissions) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *SubscriptionPermissions) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *SubscriptionPermissions) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *SubscriptionPermissions) GetOther() SubscriptionPermission`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *SubscriptionPermissions) GetOtherOk() (*SubscriptionPermission, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *SubscriptionPermissions) SetOther(v SubscriptionPermission)`

SetOther sets Other field to given value.

### HasOther

`func (o *SubscriptionPermissions) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *SubscriptionPermissions) GetAnon() SubscriptionPermission`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *SubscriptionPermissions) GetAnonOk() (*SubscriptionPermission, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *SubscriptionPermissions) SetAnon(v SubscriptionPermission)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *SubscriptionPermissions) HasAnon() bool`

HasAnon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


