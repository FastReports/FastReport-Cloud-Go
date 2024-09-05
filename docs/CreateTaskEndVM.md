# CreateTaskEndVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **NullableInt32** |  | [optional] 
**On** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCreateTaskEndVM

`func NewCreateTaskEndVM() *CreateTaskEndVM`

NewCreateTaskEndVM instantiates a new CreateTaskEndVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskEndVMWithDefaults

`func NewCreateTaskEndVMWithDefaults() *CreateTaskEndVM`

NewCreateTaskEndVMWithDefaults instantiates a new CreateTaskEndVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *CreateTaskEndVM) GetAfter() int32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CreateTaskEndVM) GetAfterOk() (*int32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CreateTaskEndVM) SetAfter(v int32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *CreateTaskEndVM) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### SetAfterNil

`func (o *CreateTaskEndVM) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *CreateTaskEndVM) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetOn

`func (o *CreateTaskEndVM) GetOn() time.Time`

GetOn returns the On field if non-nil, zero value otherwise.

### GetOnOk

`func (o *CreateTaskEndVM) GetOnOk() (*time.Time, bool)`

GetOnOk returns a tuple with the On field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOn

`func (o *CreateTaskEndVM) SetOn(v time.Time)`

SetOn sets On field to given value.

### HasOn

`func (o *CreateTaskEndVM) HasOn() bool`

HasOn returns a boolean if a field has been set.

### SetOnNil

`func (o *CreateTaskEndVM) SetOnNil(b bool)`

 SetOnNil sets the value for On to be an explicit nil

### UnsetOn
`func (o *CreateTaskEndVM) UnsetOn()`

UnsetOn ensures that no value is present for On, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


