# UpdateTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **NullableString** |  | [optional] 
**StartsOn** | Pointer to **NullableTime** |  | [optional] 
**Ends** | Pointer to [**CreateTaskEndVM**](CreateTaskEndVM.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateTaskBaseVM

`func NewUpdateTaskBaseVM(t string, ) *UpdateTaskBaseVM`

NewUpdateTaskBaseVM instantiates a new UpdateTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTaskBaseVMWithDefaults

`func NewUpdateTaskBaseVMWithDefaults() *UpdateTaskBaseVM`

NewUpdateTaskBaseVMWithDefaults instantiates a new UpdateTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *UpdateTaskBaseVM) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *UpdateTaskBaseVM) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *UpdateTaskBaseVM) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *UpdateTaskBaseVM) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### SetCronExpressionNil

`func (o *UpdateTaskBaseVM) SetCronExpressionNil(b bool)`

 SetCronExpressionNil sets the value for CronExpression to be an explicit nil

### UnsetCronExpression
`func (o *UpdateTaskBaseVM) UnsetCronExpression()`

UnsetCronExpression ensures that no value is present for CronExpression, not even an explicit nil
### GetStartsOn

`func (o *UpdateTaskBaseVM) GetStartsOn() time.Time`

GetStartsOn returns the StartsOn field if non-nil, zero value otherwise.

### GetStartsOnOk

`func (o *UpdateTaskBaseVM) GetStartsOnOk() (*time.Time, bool)`

GetStartsOnOk returns a tuple with the StartsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsOn

`func (o *UpdateTaskBaseVM) SetStartsOn(v time.Time)`

SetStartsOn sets StartsOn field to given value.

### HasStartsOn

`func (o *UpdateTaskBaseVM) HasStartsOn() bool`

HasStartsOn returns a boolean if a field has been set.

### SetStartsOnNil

`func (o *UpdateTaskBaseVM) SetStartsOnNil(b bool)`

 SetStartsOnNil sets the value for StartsOn to be an explicit nil

### UnsetStartsOn
`func (o *UpdateTaskBaseVM) UnsetStartsOn()`

UnsetStartsOn ensures that no value is present for StartsOn, not even an explicit nil
### GetEnds

`func (o *UpdateTaskBaseVM) GetEnds() CreateTaskEndVM`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *UpdateTaskBaseVM) GetEndsOk() (*CreateTaskEndVM, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *UpdateTaskBaseVM) SetEnds(v CreateTaskEndVM)`

SetEnds sets Ends field to given value.

### HasEnds

`func (o *UpdateTaskBaseVM) HasEnds() bool`

HasEnds returns a boolean if a field has been set.

### GetName

`func (o *UpdateTaskBaseVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTaskBaseVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTaskBaseVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTaskBaseVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTaskBaseVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTaskBaseVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetT

`func (o *UpdateTaskBaseVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateTaskBaseVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateTaskBaseVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


