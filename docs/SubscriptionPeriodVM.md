# SubscriptionPeriodVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Plan** | Pointer to [**SubscriptionPlanVM**](SubscriptionPlanVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewSubscriptionPeriodVM

`func NewSubscriptionPeriodVM(t string, ) *SubscriptionPeriodVM`

NewSubscriptionPeriodVM instantiates a new SubscriptionPeriodVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPeriodVMWithDefaults

`func NewSubscriptionPeriodVMWithDefaults() *SubscriptionPeriodVM`

NewSubscriptionPeriodVMWithDefaults instantiates a new SubscriptionPeriodVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SubscriptionPeriodVM) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionPeriodVM) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionPeriodVM) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SubscriptionPeriodVM) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *SubscriptionPeriodVM) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionPeriodVM) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionPeriodVM) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *SubscriptionPeriodVM) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetPlan

`func (o *SubscriptionPeriodVM) GetPlan() SubscriptionPlanVM`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *SubscriptionPeriodVM) GetPlanOk() (*SubscriptionPlanVM, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *SubscriptionPeriodVM) SetPlan(v SubscriptionPlanVM)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *SubscriptionPeriodVM) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetT

`func (o *SubscriptionPeriodVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *SubscriptionPeriodVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *SubscriptionPeriodVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


