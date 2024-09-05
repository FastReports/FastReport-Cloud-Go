# AdminUpdateCurrentSubscriptionPlanVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [**UpdateSubscriptionPlanVM**](UpdateSubscriptionPlanVM.md) |  | 
**Start** | Pointer to **NullableTime** |  | [optional] 
**End** | Pointer to **NullableTime** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAdminUpdateCurrentSubscriptionPlanVM

`func NewAdminUpdateCurrentSubscriptionPlanVM(plan UpdateSubscriptionPlanVM, t string, ) *AdminUpdateCurrentSubscriptionPlanVM`

NewAdminUpdateCurrentSubscriptionPlanVM instantiates a new AdminUpdateCurrentSubscriptionPlanVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUpdateCurrentSubscriptionPlanVMWithDefaults

`func NewAdminUpdateCurrentSubscriptionPlanVMWithDefaults() *AdminUpdateCurrentSubscriptionPlanVM`

NewAdminUpdateCurrentSubscriptionPlanVMWithDefaults instantiates a new AdminUpdateCurrentSubscriptionPlanVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlan

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetPlan() UpdateSubscriptionPlanVM`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetPlanOk() (*UpdateSubscriptionPlanVM, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AdminUpdateCurrentSubscriptionPlanVM) SetPlan(v UpdateSubscriptionPlanVM)`

SetPlan sets Plan field to given value.


### GetStart

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AdminUpdateCurrentSubscriptionPlanVM) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *AdminUpdateCurrentSubscriptionPlanVM) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *AdminUpdateCurrentSubscriptionPlanVM) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *AdminUpdateCurrentSubscriptionPlanVM) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AdminUpdateCurrentSubscriptionPlanVM) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AdminUpdateCurrentSubscriptionPlanVM) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *AdminUpdateCurrentSubscriptionPlanVM) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *AdminUpdateCurrentSubscriptionPlanVM) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetT

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AdminUpdateCurrentSubscriptionPlanVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AdminUpdateCurrentSubscriptionPlanVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


