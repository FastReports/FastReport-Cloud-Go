# TaskIdsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Ids** | Pointer to **[]string** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTaskIdsVM

`func NewTaskIdsVM(t string, ) *TaskIdsVM`

NewTaskIdsVM instantiates a new TaskIdsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskIdsVMWithDefaults

`func NewTaskIdsVMWithDefaults() *TaskIdsVM`

NewTaskIdsVMWithDefaults instantiates a new TaskIdsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TaskIdsVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TaskIdsVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TaskIdsVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TaskIdsVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIds

`func (o *TaskIdsVM) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *TaskIdsVM) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *TaskIdsVM) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *TaskIdsVM) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *TaskIdsVM) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *TaskIdsVM) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetSkip

`func (o *TaskIdsVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *TaskIdsVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *TaskIdsVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *TaskIdsVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *TaskIdsVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *TaskIdsVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *TaskIdsVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *TaskIdsVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetT

`func (o *TaskIdsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TaskIdsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TaskIdsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


