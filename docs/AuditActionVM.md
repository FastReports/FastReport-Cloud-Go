# AuditActionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**EntityId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AuditType**](AuditType.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**AdminAction** | Pointer to **bool** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditActionVM

`func NewAuditActionVM(t string, ) *AuditActionVM`

NewAuditActionVM instantiates a new AuditActionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditActionVMWithDefaults

`func NewAuditActionVMWithDefaults() *AuditActionVM`

NewAuditActionVMWithDefaults instantiates a new AuditActionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AuditActionVM) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditActionVM) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditActionVM) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuditActionVM) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AuditActionVM) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AuditActionVM) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEntityId

`func (o *AuditActionVM) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AuditActionVM) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AuditActionVM) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AuditActionVM) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *AuditActionVM) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *AuditActionVM) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetSubscriptionId

`func (o *AuditActionVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AuditActionVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AuditActionVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AuditActionVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AuditActionVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AuditActionVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *AuditActionVM) GetType() AuditType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditActionVM) GetTypeOk() (*AuditType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditActionVM) SetType(v AuditType)`

SetType sets Type field to given value.

### HasType

`func (o *AuditActionVM) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AuditActionVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditActionVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditActionVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuditActionVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AuditActionVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AuditActionVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCreatedTime

`func (o *AuditActionVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AuditActionVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AuditActionVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AuditActionVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *AuditActionVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *AuditActionVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *AuditActionVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *AuditActionVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### SetCreatorUserIdNil

`func (o *AuditActionVM) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *AuditActionVM) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetName

`func (o *AuditActionVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuditActionVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuditActionVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuditActionVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AuditActionVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuditActionVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAdminAction

`func (o *AuditActionVM) GetAdminAction() bool`

GetAdminAction returns the AdminAction field if non-nil, zero value otherwise.

### GetAdminActionOk

`func (o *AuditActionVM) GetAdminActionOk() (*bool, bool)`

GetAdminActionOk returns a tuple with the AdminAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminAction

`func (o *AuditActionVM) SetAdminAction(v bool)`

SetAdminAction sets AdminAction field to given value.

### HasAdminAction

`func (o *AuditActionVM) HasAdminAction() bool`

HasAdminAction returns a boolean if a field has been set.

### GetT

`func (o *AuditActionVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditActionVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditActionVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


