# SubscriptionPlansVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionPlans** | Pointer to [**[]SubscriptionPlanVM**](SubscriptionPlanVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewSubscriptionPlansVM

`func NewSubscriptionPlansVM(t string, ) *SubscriptionPlansVM`

NewSubscriptionPlansVM instantiates a new SubscriptionPlansVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlansVMWithDefaults

`func NewSubscriptionPlansVMWithDefaults() *SubscriptionPlansVM`

NewSubscriptionPlansVMWithDefaults instantiates a new SubscriptionPlansVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionPlans

`func (o *SubscriptionPlansVM) GetSubscriptionPlans() []SubscriptionPlanVM`

GetSubscriptionPlans returns the SubscriptionPlans field if non-nil, zero value otherwise.

### GetSubscriptionPlansOk

`func (o *SubscriptionPlansVM) GetSubscriptionPlansOk() (*[]SubscriptionPlanVM, bool)`

GetSubscriptionPlansOk returns a tuple with the SubscriptionPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlans

`func (o *SubscriptionPlansVM) SetSubscriptionPlans(v []SubscriptionPlanVM)`

SetSubscriptionPlans sets SubscriptionPlans field to given value.

### HasSubscriptionPlans

`func (o *SubscriptionPlansVM) HasSubscriptionPlans() bool`

HasSubscriptionPlans returns a boolean if a field has been set.

### SetSubscriptionPlansNil

`func (o *SubscriptionPlansVM) SetSubscriptionPlansNil(b bool)`

 SetSubscriptionPlansNil sets the value for SubscriptionPlans to be an explicit nil

### UnsetSubscriptionPlans
`func (o *SubscriptionPlansVM) UnsetSubscriptionPlans()`

UnsetSubscriptionPlans ensures that no value is present for SubscriptionPlans, not even an explicit nil
### GetCount

`func (o *SubscriptionPlansVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SubscriptionPlansVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SubscriptionPlansVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *SubscriptionPlansVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *SubscriptionPlansVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *SubscriptionPlansVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *SubscriptionPlansVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *SubscriptionPlansVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *SubscriptionPlansVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *SubscriptionPlansVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *SubscriptionPlansVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *SubscriptionPlansVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetT

`func (o *SubscriptionPlansVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SubscriptionPlansVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SubscriptionPlansVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


