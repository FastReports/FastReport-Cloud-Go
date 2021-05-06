# SubscriptionInviteVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usages** | Pointer to **int64** |  | [optional] 
**Durable** | Pointer to **bool** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**ExpiredDate** | Pointer to **time.Time** |  | [optional] 
**AddedUsers** | Pointer to [**[]InvitedUser**](InvitedUser.md) |  | [optional] 
**CreatorUserId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscriptionInviteVM

`func NewSubscriptionInviteVM() *SubscriptionInviteVM`

NewSubscriptionInviteVM instantiates a new SubscriptionInviteVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionInviteVMWithDefaults

`func NewSubscriptionInviteVMWithDefaults() *SubscriptionInviteVM`

NewSubscriptionInviteVMWithDefaults instantiates a new SubscriptionInviteVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsages

`func (o *SubscriptionInviteVM) GetUsages() int64`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *SubscriptionInviteVM) GetUsagesOk() (*int64, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *SubscriptionInviteVM) SetUsages(v int64)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *SubscriptionInviteVM) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetDurable

`func (o *SubscriptionInviteVM) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *SubscriptionInviteVM) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *SubscriptionInviteVM) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *SubscriptionInviteVM) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetAccessToken

`func (o *SubscriptionInviteVM) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SubscriptionInviteVM) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SubscriptionInviteVM) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *SubscriptionInviteVM) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetExpiredDate

`func (o *SubscriptionInviteVM) GetExpiredDate() time.Time`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *SubscriptionInviteVM) GetExpiredDateOk() (*time.Time, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *SubscriptionInviteVM) SetExpiredDate(v time.Time)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *SubscriptionInviteVM) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.

### GetAddedUsers

`func (o *SubscriptionInviteVM) GetAddedUsers() []InvitedUser`

GetAddedUsers returns the AddedUsers field if non-nil, zero value otherwise.

### GetAddedUsersOk

`func (o *SubscriptionInviteVM) GetAddedUsersOk() (*[]InvitedUser, bool)`

GetAddedUsersOk returns a tuple with the AddedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedUsers

`func (o *SubscriptionInviteVM) SetAddedUsers(v []InvitedUser)`

SetAddedUsers sets AddedUsers field to given value.

### HasAddedUsers

`func (o *SubscriptionInviteVM) HasAddedUsers() bool`

HasAddedUsers returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *SubscriptionInviteVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *SubscriptionInviteVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *SubscriptionInviteVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *SubscriptionInviteVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


