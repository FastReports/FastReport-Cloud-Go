# CreateSubscriptionInviteVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usages** | Pointer to **int64** |  | [optional] 
**Durable** | Pointer to **bool** |  | [optional] 
**ExpiredDate** | Pointer to **time.Time** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateSubscriptionInviteVM

`func NewCreateSubscriptionInviteVM(t string, ) *CreateSubscriptionInviteVM`

NewCreateSubscriptionInviteVM instantiates a new CreateSubscriptionInviteVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionInviteVMWithDefaults

`func NewCreateSubscriptionInviteVMWithDefaults() *CreateSubscriptionInviteVM`

NewCreateSubscriptionInviteVMWithDefaults instantiates a new CreateSubscriptionInviteVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsages

`func (o *CreateSubscriptionInviteVM) GetUsages() int64`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *CreateSubscriptionInviteVM) GetUsagesOk() (*int64, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *CreateSubscriptionInviteVM) SetUsages(v int64)`

SetUsages sets Usages field to given value.

### HasUsages

`func (o *CreateSubscriptionInviteVM) HasUsages() bool`

HasUsages returns a boolean if a field has been set.

### GetDurable

`func (o *CreateSubscriptionInviteVM) GetDurable() bool`

GetDurable returns the Durable field if non-nil, zero value otherwise.

### GetDurableOk

`func (o *CreateSubscriptionInviteVM) GetDurableOk() (*bool, bool)`

GetDurableOk returns a tuple with the Durable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurable

`func (o *CreateSubscriptionInviteVM) SetDurable(v bool)`

SetDurable sets Durable field to given value.

### HasDurable

`func (o *CreateSubscriptionInviteVM) HasDurable() bool`

HasDurable returns a boolean if a field has been set.

### GetExpiredDate

`func (o *CreateSubscriptionInviteVM) GetExpiredDate() time.Time`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *CreateSubscriptionInviteVM) GetExpiredDateOk() (*time.Time, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *CreateSubscriptionInviteVM) SetExpiredDate(v time.Time)`

SetExpiredDate sets ExpiredDate field to given value.

### HasExpiredDate

`func (o *CreateSubscriptionInviteVM) HasExpiredDate() bool`

HasExpiredDate returns a boolean if a field has been set.

### GetT

`func (o *CreateSubscriptionInviteVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateSubscriptionInviteVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateSubscriptionInviteVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


