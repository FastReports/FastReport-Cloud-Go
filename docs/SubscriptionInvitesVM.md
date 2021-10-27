# SubscriptionInvitesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | Pointer to [**[]SubscriptionInviteVM**](SubscriptionInviteVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewSubscriptionInvitesVM

`func NewSubscriptionInvitesVM() *SubscriptionInvitesVM`

NewSubscriptionInvitesVM instantiates a new SubscriptionInvitesVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionInvitesVMWithDefaults

`func NewSubscriptionInvitesVMWithDefaults() *SubscriptionInvitesVM`

NewSubscriptionInvitesVMWithDefaults instantiates a new SubscriptionInvitesVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *SubscriptionInvitesVM) GetInvites() []SubscriptionInviteVM`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *SubscriptionInvitesVM) GetInvitesOk() (*[]SubscriptionInviteVM, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *SubscriptionInvitesVM) SetInvites(v []SubscriptionInviteVM)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *SubscriptionInvitesVM) HasInvites() bool`

HasInvites returns a boolean if a field has been set.

### SetInvitesNil

`func (o *SubscriptionInvitesVM) SetInvitesNil(b bool)`

 SetInvitesNil sets the value for Invites to be an explicit nil

### UnsetInvites
`func (o *SubscriptionInvitesVM) UnsetInvites()`

UnsetInvites ensures that no value is present for Invites, not even an explicit nil
### GetCount

`func (o *SubscriptionInvitesVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SubscriptionInvitesVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SubscriptionInvitesVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SubscriptionInvitesVM) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


