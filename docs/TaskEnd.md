# TaskEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **NullableInt32** |  | [optional] 
**On** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewTaskEnd

`func NewTaskEnd() *TaskEnd`

NewTaskEnd instantiates a new TaskEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskEndWithDefaults

`func NewTaskEndWithDefaults() *TaskEnd`

NewTaskEndWithDefaults instantiates a new TaskEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *TaskEnd) GetAfter() int32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *TaskEnd) GetAfterOk() (*int32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *TaskEnd) SetAfter(v int32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *TaskEnd) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *TaskEnd) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *TaskEnd) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetOn

`func (o *TaskEnd) GetOn() time.Time`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *TaskEnd) GetOnOk() (*time.Time, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *TaskEnd) SetOn(v time.Time)`

SetOn sets On field to given value.

### HasOn

`func (o *TaskEnd) HasOn() bool`

HasOn returns a boolean if a field has been set.

### SetOnNil

`func (o *TaskEnd) SetOnNil(b bool)`

 SetOnNil sets the value for On to be an explicit nil

### UnsetOn
`func (o *TaskEnd) UnsetOn()`

UnsetOn ensures that no value is present for On, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


