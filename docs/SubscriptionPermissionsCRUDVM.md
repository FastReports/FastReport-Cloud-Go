# SubscriptionPermissionsCRUDVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**SubscriptionPermissionCRUDVM**](SubscriptionPermissionCRUDVM.md) |  | [optional] 
**Groups** | Pointer to [**map[string]SubscriptionPermissionCRUDVM**](SubscriptionPermissionCRUDVM.md) |  | [optional] 
**Other** | Pointer to [**SubscriptionPermissionCRUDVM**](SubscriptionPermissionCRUDVM.md) |  | [optional] 
**Anon** | Pointer to [**SubscriptionPermissionCRUDVM**](SubscriptionPermissionCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewSubscriptionPermissionsCRUDVM

`func NewSubscriptionPermissionsCRUDVM(t string, ) *SubscriptionPermissionsCRUDVM`

NewSubscriptionPermissionsCRUDVM instantiates a new SubscriptionPermissionsCRUDVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPermissionsCRUDVMWithDefaults

`func NewSubscriptionPermissionsCRUDVMWithDefaults() *SubscriptionPermissionsCRUDVM`

NewSubscriptionPermissionsCRUDVMWithDefaults instantiates a new SubscriptionPermissionsCRUDVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *SubscriptionPermissionsCRUDVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *SubscriptionPermissionsCRUDVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *SubscriptionPermissionsCRUDVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *SubscriptionPermissionsCRUDVM) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *SubscriptionPermissionsCRUDVM) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *SubscriptionPermissionsCRUDVM) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *SubscriptionPermissionsCRUDVM) GetOwner() SubscriptionPermissionCRUDVM`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SubscriptionPermissionsCRUDVM) GetOwnerOk() (*SubscriptionPermissionCRUDVM, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SubscriptionPermissionsCRUDVM) SetOwner(v SubscriptionPermissionCRUDVM)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *SubscriptionPermissionsCRUDVM) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *SubscriptionPermissionsCRUDVM) GetGroups() map[string]SubscriptionPermissionCRUDVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SubscriptionPermissionsCRUDVM) GetGroupsOk() (*map[string]SubscriptionPermissionCRUDVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SubscriptionPermissionsCRUDVM) SetGroups(v map[string]SubscriptionPermissionCRUDVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SubscriptionPermissionsCRUDVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *SubscriptionPermissionsCRUDVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *SubscriptionPermissionsCRUDVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *SubscriptionPermissionsCRUDVM) GetOther() SubscriptionPermissionCRUDVM`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *SubscriptionPermissionsCRUDVM) GetOtherOk() (*SubscriptionPermissionCRUDVM, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *SubscriptionPermissionsCRUDVM) SetOther(v SubscriptionPermissionCRUDVM)`

SetOther sets Other field to given value.

### HasOther

`func (o *SubscriptionPermissionsCRUDVM) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *SubscriptionPermissionsCRUDVM) GetAnon() SubscriptionPermissionCRUDVM`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *SubscriptionPermissionsCRUDVM) GetAnonOk() (*SubscriptionPermissionCRUDVM, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *SubscriptionPermissionsCRUDVM) SetAnon(v SubscriptionPermissionCRUDVM)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *SubscriptionPermissionsCRUDVM) HasAnon() bool`

HasAnon returns a boolean if a field has been set.

### GetT

`func (o *SubscriptionPermissionsCRUDVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SubscriptionPermissionsCRUDVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SubscriptionPermissionsCRUDVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


