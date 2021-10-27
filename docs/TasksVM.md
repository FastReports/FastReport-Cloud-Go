# TasksVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**Tasks** | Pointer to [**[]TaskBaseVM**](TaskBaseVM.md) |  | [optional] 

## Methods

### NewTasksVM

`func NewTasksVM() *TasksVM`

NewTasksVM instantiates a new TasksVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksVMWithDefaults

`func NewTasksVMWithDefaults() *TasksVM`

NewTasksVMWithDefaults instantiates a new TasksVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TasksVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TasksVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TasksVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TasksVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *TasksVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *TasksVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *TasksVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *TasksVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *TasksVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *TasksVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *TasksVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *TasksVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetTasks

`func (o *TasksVM) GetTasks() []TaskBaseVM`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *TasksVM) GetTasksOk() (*[]TaskBaseVM, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *TasksVM) SetTasks(v []TaskBaseVM)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *TasksVM) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *TasksVM) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *TasksVM) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


