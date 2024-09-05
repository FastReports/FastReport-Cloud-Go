# CreateSubscriptionPeriodVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** |  | 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateSubscriptionPeriodVM

`func NewCreateSubscriptionPeriodVM(planId string, t string, ) *CreateSubscriptionPeriodVM`

NewCreateSubscriptionPeriodVM instantiates a new CreateSubscriptionPeriodVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionPeriodVMWithDefaults

`func NewCreateSubscriptionPeriodVMWithDefaults() *CreateSubscriptionPeriodVM`

NewCreateSubscriptionPeriodVMWithDefaults instantiates a new CreateSubscriptionPeriodVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CreateSubscriptionPeriodVM) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateSubscriptionPeriodVM) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateSubscriptionPeriodVM) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetStart

`func (o *CreateSubscriptionPeriodVM) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CreateSubscriptionPeriodVM) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CreateSubscriptionPeriodVM) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *CreateSubscriptionPeriodVM) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *CreateSubscriptionPeriodVM) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *CreateSubscriptionPeriodVM) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *CreateSubscriptionPeriodVM) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CreateSubscriptionPeriodVM) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CreateSubscriptionPeriodVM) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CreateSubscriptionPeriodVM) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *CreateSubscriptionPeriodVM) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *CreateSubscriptionPeriodVM) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetT

`func (o *CreateSubscriptionPeriodVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateSubscriptionPeriodVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateSubscriptionPeriodVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


