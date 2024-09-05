# CreateAuditActionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**EntityId** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**AuditType**](AuditType.md) |  | [optional] 
**MessageId** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateAuditActionVM

`func NewCreateAuditActionVM(t string, ) *CreateAuditActionVM`

NewCreateAuditActionVM instantiates a new CreateAuditActionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuditActionVMWithDefaults

`func NewCreateAuditActionVMWithDefaults() *CreateAuditActionVM`

NewCreateAuditActionVMWithDefaults instantiates a new CreateAuditActionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateAuditActionVM) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateAuditActionVM) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateAuditActionVM) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateAuditActionVM) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreateAuditActionVM) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreateAuditActionVM) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetEntityId

`func (o *CreateAuditActionVM) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *CreateAuditActionVM) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *CreateAuditActionVM) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *CreateAuditActionVM) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *CreateAuditActionVM) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *CreateAuditActionVM) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetSubscriptionId

`func (o *CreateAuditActionVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateAuditActionVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateAuditActionVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateAuditActionVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateAuditActionVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateAuditActionVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *CreateAuditActionVM) GetType() AuditType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAuditActionVM) GetTypeOk() (*AuditType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAuditActionVM) SetType(v AuditType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateAuditActionVM) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessageId

`func (o *CreateAuditActionVM) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *CreateAuditActionVM) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *CreateAuditActionVM) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *CreateAuditActionVM) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### SetMessageIdNil

`func (o *CreateAuditActionVM) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *CreateAuditActionVM) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetFileName

`func (o *CreateAuditActionVM) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CreateAuditActionVM) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CreateAuditActionVM) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CreateAuditActionVM) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *CreateAuditActionVM) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *CreateAuditActionVM) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetT

`func (o *CreateAuditActionVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateAuditActionVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateAuditActionVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


