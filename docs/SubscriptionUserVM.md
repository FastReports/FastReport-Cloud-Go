# SubscriptionUserVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**Groups** | Pointer to [**[]GroupVM**](GroupVM.md) |  | [optional] 

## Methods

### NewSubscriptionUserVM

`func NewSubscriptionUserVM() *SubscriptionUserVM`

NewSubscriptionUserVM instantiates a new SubscriptionUserVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUserVMWithDefaults

`func NewSubscriptionUserVMWithDefaults() *SubscriptionUserVM`

NewSubscriptionUserVMWithDefaults instantiates a new SubscriptionUserVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SubscriptionUserVM) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubscriptionUserVM) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubscriptionUserVM) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SubscriptionUserVM) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *SubscriptionUserVM) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *SubscriptionUserVM) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetGroups

`func (o *SubscriptionUserVM) GetGroups() []GroupVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SubscriptionUserVM) GetGroupsOk() (*[]GroupVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SubscriptionUserVM) SetGroups(v []GroupVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SubscriptionUserVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *SubscriptionUserVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *SubscriptionUserVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


