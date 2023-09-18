# TaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **NullableString** |  | [optional] 
**DelayedRunTime** | Pointer to **NullableTime** |  | [optional] 
**DelayedWasRunTime** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**RecurrentRunTime** | Pointer to **NullableTime** |  | [optional] [readonly] 
**RecurrentWasRunTime** | Pointer to **NullableTime** |  | [optional] [readonly] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTaskBaseVM

`func NewTaskBaseVM(t string, ) *TaskBaseVM`

NewTaskBaseVM instantiates a new TaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskBaseVMWithDefaults

`func NewTaskBaseVMWithDefaults() *TaskBaseVM`

NewTaskBaseVMWithDefaults instantiates a new TaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *TaskBaseVM) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *TaskBaseVM) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *TaskBaseVM) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *TaskBaseVM) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *TaskBaseVM) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *TaskBaseVM) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetDelayedRunTime

`func (o *TaskBaseVM) GetDelayedRunTime() time.Time`

GetDelayedRunTime returns the DelayedRunTime field if non-nil, zero value otherwise.

### GetDelayedRunTimeOk

`func (o *TaskBaseVM) GetDelayedRunTimeOk() (*time.Time, bool)`

GetDelayedRunTimeOk returns a tuple with the DelayedRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedRunTime

`func (o *TaskBaseVM) SetDelayedRunTime(v time.Time)`

SetDelayedRunTime sets DelayedRunTime field to given value.

### HasDelayedRunTime

`func (o *TaskBaseVM) HasDelayedRunTime() bool`

HasDelayedRunTime returns a boolean if a field has been set.

### SetDelayedRunTimeNil

`func (o *TaskBaseVM) SetDelayedRunTimeNil(b bool)`

 SetDelayedRunTimeNil sets the value for DelayedRunTime to be an explicit nil

### UnsetDelayedRunTime
`func (o *TaskBaseVM) UnsetDelayedRunTime()`

UnsetDelayedRunTime ensures that no value is present for DelayedRunTime, not even an explicit nil
### GetDelayedWasRunTime

`func (o *TaskBaseVM) GetDelayedWasRunTime() time.Time`

GetDelayedWasRunTime returns the DelayedWasRunTime field if non-nil, zero value otherwise.

### GetDelayedWasRunTimeOk

`func (o *TaskBaseVM) GetDelayedWasRunTimeOk() (*time.Time, bool)`

GetDelayedWasRunTimeOk returns a tuple with the DelayedWasRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedWasRunTime

`func (o *TaskBaseVM) SetDelayedWasRunTime(v time.Time)`

SetDelayedWasRunTime sets DelayedWasRunTime field to given value.

### HasDelayedWasRunTime

`func (o *TaskBaseVM) HasDelayedWasRunTime() bool`

HasDelayedWasRunTime returns a boolean if a field has been set.

### SetDelayedWasRunTimeNil

`func (o *TaskBaseVM) SetDelayedWasRunTimeNil(b bool)`

 SetDelayedWasRunTimeNil sets the value for DelayedWasRunTime to be an explicit nil

### UnsetDelayedWasRunTime
`func (o *TaskBaseVM) UnsetDelayedWasRunTime()`

UnsetDelayedWasRunTime ensures that no value is present for DelayedWasRunTime, not even an explicit nil
### GetId

`func (o *TaskBaseVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskBaseVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskBaseVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskBaseVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskBaseVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskBaseVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TaskBaseVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskBaseVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskBaseVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskBaseVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskBaseVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskBaseVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRecurrentRunTime

`func (o *TaskBaseVM) GetRecurrentRunTime() time.Time`

GetRecurrentRunTime returns the RecurrentRunTime field if non-nil, zero value otherwise.

### GetRecurrentRunTimeOk

`func (o *TaskBaseVM) GetRecurrentRunTimeOk() (*time.Time, bool)`

GetRecurrentRunTimeOk returns a tuple with the RecurrentRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrentRunTime

`func (o *TaskBaseVM) SetRecurrentRunTime(v time.Time)`

SetRecurrentRunTime sets RecurrentRunTime field to given value.

### HasRecurrentRunTime

`func (o *TaskBaseVM) HasRecurrentRunTime() bool`

HasRecurrentRunTime returns a boolean if a field has been set.

### SetRecurrentRunTimeNil

`func (o *TaskBaseVM) SetRecurrentRunTimeNil(b bool)`

 SetRecurrentRunTimeNil sets the value for RecurrentRunTime to be an explicit nil

### UnsetRecurrentRunTime
`func (o *TaskBaseVM) UnsetRecurrentRunTime()`

UnsetRecurrentRunTime ensures that no value is present for RecurrentRunTime, not even an explicit nil
### GetRecurrentWasRunTime

`func (o *TaskBaseVM) GetRecurrentWasRunTime() time.Time`

GetRecurrentWasRunTime returns the RecurrentWasRunTime field if non-nil, zero value otherwise.

### GetRecurrentWasRunTimeOk

`func (o *TaskBaseVM) GetRecurrentWasRunTimeOk() (*time.Time, bool)`

GetRecurrentWasRunTimeOk returns a tuple with the RecurrentWasRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrentWasRunTime

`func (o *TaskBaseVM) SetRecurrentWasRunTime(v time.Time)`

SetRecurrentWasRunTime sets RecurrentWasRunTime field to given value.

### HasRecurrentWasRunTime

`func (o *TaskBaseVM) HasRecurrentWasRunTime() bool`

HasRecurrentWasRunTime returns a boolean if a field has been set.

### SetRecurrentWasRunTimeNil

`func (o *TaskBaseVM) SetRecurrentWasRunTimeNil(b bool)`

 SetRecurrentWasRunTimeNil sets the value for RecurrentWasRunTime to be an explicit nil

### UnsetRecurrentWasRunTime
`func (o *TaskBaseVM) UnsetRecurrentWasRunTime()`

UnsetRecurrentWasRunTime ensures that no value is present for RecurrentWasRunTime, not even an explicit nil
### GetSubscriptionId

`func (o *TaskBaseVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *TaskBaseVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *TaskBaseVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *TaskBaseVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *TaskBaseVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *TaskBaseVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetT

`func (o *TaskBaseVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TaskBaseVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TaskBaseVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


