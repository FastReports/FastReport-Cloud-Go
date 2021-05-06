# SubscriptionUsersVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]SubscriptionUserVM**](SubscriptionUserVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubscriptionUsersVM

`func NewSubscriptionUsersVM() *SubscriptionUsersVM`

NewSubscriptionUsersVM instantiates a new SubscriptionUsersVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUsersVMWithDefaults

`func NewSubscriptionUsersVMWithDefaults() *SubscriptionUsersVM`

NewSubscriptionUsersVMWithDefaults instantiates a new SubscriptionUsersVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *SubscriptionUsersVM) GetUsers() []SubscriptionUserVM`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *SubscriptionUsersVM) GetUsersOk() (*[]SubscriptionUserVM, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *SubscriptionUsersVM) SetUsers(v []SubscriptionUserVM)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *SubscriptionUsersVM) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCount

`func (o *SubscriptionUsersVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SubscriptionUsersVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SubscriptionUsersVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SubscriptionUsersVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTake

`func (o *SubscriptionUsersVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *SubscriptionUsersVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *SubscriptionUsersVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *SubscriptionUsersVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetSkip

`func (o *SubscriptionUsersVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *SubscriptionUsersVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *SubscriptionUsersVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *SubscriptionUsersVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


