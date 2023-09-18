# CreateTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **NullableString** |  | [optional] 
**DelayedRunTime** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateTaskBaseVM

`func NewCreateTaskBaseVM(t string, ) *CreateTaskBaseVM`

NewCreateTaskBaseVM instantiates a new CreateTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskBaseVMWithDefaults

`func NewCreateTaskBaseVMWithDefaults() *CreateTaskBaseVM`

NewCreateTaskBaseVMWithDefaults instantiates a new CreateTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *CreateTaskBaseVM) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *CreateTaskBaseVM) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *CreateTaskBaseVM) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *CreateTaskBaseVM) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *CreateTaskBaseVM) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *CreateTaskBaseVM) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetDelayedRunTime

`func (o *CreateTaskBaseVM) GetDelayedRunTime() time.Time`

GetDelayedRunTime returns the DelayedRunTime field if non-nil, zero value otherwise.

### GetDelayedRunTimeOk

`func (o *CreateTaskBaseVM) GetDelayedRunTimeOk() (*time.Time, bool)`

GetDelayedRunTimeOk returns a tuple with the DelayedRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayedRunTime

`func (o *CreateTaskBaseVM) SetDelayedRunTime(v time.Time)`

SetDelayedRunTime sets DelayedRunTime field to given value.

### HasDelayedRunTime

`func (o *CreateTaskBaseVM) HasDelayedRunTime() bool`

HasDelayedRunTime returns a boolean if a field has been set.

### SetDelayedRunTimeNil

`func (o *CreateTaskBaseVM) SetDelayedRunTimeNil(b bool)`

 SetDelayedRunTimeNil sets the value for DelayedRunTime to be an explicit nil

### UnsetDelayedRunTime
`func (o *CreateTaskBaseVM) UnsetDelayedRunTime()`

UnsetDelayedRunTime ensures that no value is present for DelayedRunTime, not even an explicit nil
### GetName

`func (o *CreateTaskBaseVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTaskBaseVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTaskBaseVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTaskBaseVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTaskBaseVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTaskBaseVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *CreateTaskBaseVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateTaskBaseVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateTaskBaseVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateTaskBaseVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateTaskBaseVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateTaskBaseVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetT

`func (o *CreateTaskBaseVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateTaskBaseVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateTaskBaseVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


