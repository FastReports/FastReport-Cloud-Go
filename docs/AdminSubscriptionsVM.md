# AdminSubscriptionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriptions** | Pointer to [**[]AdminSubscriptionVM**](AdminSubscriptionVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminSubscriptionsVM

`func NewAdminSubscriptionsVM() *AdminSubscriptionsVM`

NewAdminSubscriptionsVM instantiates a new AdminSubscriptionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminSubscriptionsVMWithDefaults

`func NewAdminSubscriptionsVMWithDefaults() *AdminSubscriptionsVM`

NewAdminSubscriptionsVMWithDefaults instantiates a new AdminSubscriptionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptions

`func (o *AdminSubscriptionsVM) GetSubscriptions() []AdminSubscriptionVM`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AdminSubscriptionsVM) GetSubscriptionsOk() (*[]AdminSubscriptionVM, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AdminSubscriptionsVM) SetSubscriptions(v []AdminSubscriptionVM)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AdminSubscriptionsVM) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetCount

`func (o *AdminSubscriptionsVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AdminSubscriptionsVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AdminSubscriptionsVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *AdminSubscriptionsVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *AdminSubscriptionsVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *AdminSubscriptionsVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *AdminSubscriptionsVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *AdminSubscriptionsVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *AdminSubscriptionsVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *AdminSubscriptionsVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *AdminSubscriptionsVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *AdminSubscriptionsVM) HasTake() bool`

HasTake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


