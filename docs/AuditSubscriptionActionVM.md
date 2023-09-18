# AuditSubscriptionActionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodStart** | Pointer to **time.Time** |  | [optional] 
**PeriodEnd** | Pointer to **time.Time** |  | [optional] 
**PlanId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuditSubscriptionActionVM

`func NewAuditSubscriptionActionVM() *AuditSubscriptionActionVM`

NewAuditSubscriptionActionVM instantiates a new AuditSubscriptionActionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditSubscriptionActionVMWithDefaults

`func NewAuditSubscriptionActionVMWithDefaults() *AuditSubscriptionActionVM`

NewAuditSubscriptionActionVMWithDefaults instantiates a new AuditSubscriptionActionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodStart

`func (o *AuditSubscriptionActionVM) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *AuditSubscriptionActionVM) GetPeriodStartOk() (*time.Time, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *AuditSubscriptionActionVM) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### HasPeriodStart

`func (o *AuditSubscriptionActionVM) HasPeriodStart() bool`

HasPeriodStart returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *AuditSubscriptionActionVM) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *AuditSubscriptionActionVM) GetPeriodEndOk() (*time.Time, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *AuditSubscriptionActionVM) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### HasPeriodEnd

`func (o *AuditSubscriptionActionVM) HasPeriodEnd() bool`

HasPeriodEnd returns a boolean if a field has been set.

### GetPlanId

`func (o *AuditSubscriptionActionVM) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *AuditSubscriptionActionVM) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *AuditSubscriptionActionVM) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *AuditSubscriptionActionVM) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### SetPlanIdNil

`func (o *AuditSubscriptionActionVM) SetPlanIdNil(b bool)`

 SetPlanIdNil sets the value for PlanId to be an explicit nil

### UnsetPlanId
`func (o *AuditSubscriptionActionVM) UnsetPlanId()`

UnsetPlanId ensures that no value is present for PlanId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


