# RestOfSpaceVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestOfSpace** | Pointer to **int64** |  | [optional] 
**SubscriptionPlan** | Pointer to [**SubscriptionPlanVM**](SubscriptionPlanVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewRestOfSpaceVM

`func NewRestOfSpaceVM(t string, ) *RestOfSpaceVM`

NewRestOfSpaceVM instantiates a new RestOfSpaceVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestOfSpaceVMWithDefaults

`func NewRestOfSpaceVMWithDefaults() *RestOfSpaceVM`

NewRestOfSpaceVMWithDefaults instantiates a new RestOfSpaceVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestOfSpace

`func (o *RestOfSpaceVM) GetRestOfSpace() int64`

GetRestOfSpace returns the RestOfSpace field if non-nil, zero value otherwise.

### GetRestOfSpaceOk

`func (o *RestOfSpaceVM) GetRestOfSpaceOk() (*int64, bool)`

GetRestOfSpaceOk returns a tuple with the RestOfSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestOfSpace

`func (o *RestOfSpaceVM) SetRestOfSpace(v int64)`

SetRestOfSpace sets RestOfSpace field to given value.

### HasRestOfSpace

`func (o *RestOfSpaceVM) HasRestOfSpace() bool`

HasRestOfSpace returns a boolean if a field has been set.

### GetSubscriptionPlan

`func (o *RestOfSpaceVM) GetSubscriptionPlan() SubscriptionPlanVM`

GetSubscriptionPlan returns the SubscriptionPlan field if non-nil, zero value otherwise.

### GetSubscriptionPlanOk

`func (o *RestOfSpaceVM) GetSubscriptionPlanOk() (*SubscriptionPlanVM, bool)`

GetSubscriptionPlanOk returns a tuple with the SubscriptionPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPlan

`func (o *RestOfSpaceVM) SetSubscriptionPlan(v SubscriptionPlanVM)`

SetSubscriptionPlan sets SubscriptionPlan field to given value.

### HasSubscriptionPlan

`func (o *RestOfSpaceVM) HasSubscriptionPlan() bool`

HasSubscriptionPlan returns a boolean if a field has been set.

### GetT

`func (o *RestOfSpaceVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RestOfSpaceVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RestOfSpaceVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


