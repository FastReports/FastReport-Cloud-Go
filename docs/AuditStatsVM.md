# AuditStatsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**map[string][]AuditStatVM**](array.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditStatsVM

`func NewAuditStatsVM(t string, ) *AuditStatsVM`

NewAuditStatsVM instantiates a new AuditStatsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditStatsVMWithDefaults

`func NewAuditStatsVMWithDefaults() *AuditStatsVM`

NewAuditStatsVMWithDefaults instantiates a new AuditStatsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuditStatsVM) GetItems() map[string][]AuditStatVM`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditStatsVM) GetItemsOk() (*map[string][]AuditStatVM, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditStatsVM) SetItems(v map[string][]AuditStatVM)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuditStatsVM) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *AuditStatsVM) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AuditStatsVM) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetT

`func (o *AuditStatsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditStatsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditStatsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


