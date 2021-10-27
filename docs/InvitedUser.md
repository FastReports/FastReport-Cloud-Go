# InvitedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**InvitedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInvitedUser

`func NewInvitedUser() *InvitedUser`

NewInvitedUser instantiates a new InvitedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitedUserWithDefaults

`func NewInvitedUserWithDefaults() *InvitedUser`

NewInvitedUserWithDefaults instantiates a new InvitedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *InvitedUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InvitedUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InvitedUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *InvitedUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *InvitedUser) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *InvitedUser) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetInvitedAt

`func (o *InvitedUser) GetInvitedAt() time.Time`

GetInvitedAt returns the InvitedAt field if non-nil, zero value otherwise.

### GetInvitedAtOk

`func (o *InvitedUser) GetInvitedAtOk() (*time.Time, bool)`

GetInvitedAtOk returns a tuple with the InvitedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedAt

`func (o *InvitedUser) SetInvitedAt(v time.Time)`

SetInvitedAt sets InvitedAt field to given value.

### HasInvitedAt

`func (o *InvitedUser) HasInvitedAt() bool`

HasInvitedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


