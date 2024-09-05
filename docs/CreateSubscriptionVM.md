# CreateSubscriptionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanId** | **string** |  | 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**Name** | **string** |  | 
**UserId** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateSubscriptionVM

`func NewCreateSubscriptionVM(planId string, name string, userId string, t string, ) *CreateSubscriptionVM`

NewCreateSubscriptionVM instantiates a new CreateSubscriptionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionVMWithDefaults

`func NewCreateSubscriptionVMWithDefaults() *CreateSubscriptionVM`

NewCreateSubscriptionVMWithDefaults instantiates a new CreateSubscriptionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanId

`func (o *CreateSubscriptionVM) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *CreateSubscriptionVM) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *CreateSubscriptionVM) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetStart

`func (o *CreateSubscriptionVM) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CreateSubscriptionVM) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CreateSubscriptionVM) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *CreateSubscriptionVM) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *CreateSubscriptionVM) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *CreateSubscriptionVM) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *CreateSubscriptionVM) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CreateSubscriptionVM) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CreateSubscriptionVM) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *CreateSubscriptionVM) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *CreateSubscriptionVM) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *CreateSubscriptionVM) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetName

`func (o *CreateSubscriptionVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubscriptionVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubscriptionVM) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *CreateSubscriptionVM) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateSubscriptionVM) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateSubscriptionVM) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTags

`func (o *CreateSubscriptionVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateSubscriptionVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateSubscriptionVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateSubscriptionVM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateSubscriptionVM) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateSubscriptionVM) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetT

`func (o *CreateSubscriptionVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateSubscriptionVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateSubscriptionVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


