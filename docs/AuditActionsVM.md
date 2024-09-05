# AuditActionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AuditActionVM**](AuditActionVM.md) |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditActionsVM

`func NewAuditActionsVM(t string, ) *AuditActionsVM`

NewAuditActionsVM instantiates a new AuditActionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditActionsVMWithDefaults

`func NewAuditActionsVMWithDefaults() *AuditActionsVM`

NewAuditActionsVMWithDefaults instantiates a new AuditActionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuditActionsVM) GetItems() []AuditActionVM`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditActionsVM) GetItemsOk() (*[]AuditActionVM, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditActionsVM) SetItems(v []AuditActionVM)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuditActionsVM) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *AuditActionsVM) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AuditActionsVM) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetHasMore

`func (o *AuditActionsVM) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *AuditActionsVM) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *AuditActionsVM) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *AuditActionsVM) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetSkip

`func (o *AuditActionsVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *AuditActionsVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *AuditActionsVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *AuditActionsVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *AuditActionsVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *AuditActionsVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *AuditActionsVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *AuditActionsVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetT

`func (o *AuditActionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditActionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditActionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


