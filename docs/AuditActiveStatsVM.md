# AuditActiveStatsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **map[string]int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditActiveStatsVM

`func NewAuditActiveStatsVM(t string, ) *AuditActiveStatsVM`

NewAuditActiveStatsVM instantiates a new AuditActiveStatsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditActiveStatsVMWithDefaults

`func NewAuditActiveStatsVMWithDefaults() *AuditActiveStatsVM`

NewAuditActiveStatsVMWithDefaults instantiates a new AuditActiveStatsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuditActiveStatsVM) GetItems() map[string]int32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditActiveStatsVM) GetItemsOk() (*map[string]int32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditActiveStatsVM) SetItems(v map[string]int32)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuditActiveStatsVM) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *AuditActiveStatsVM) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *AuditActiveStatsVM) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetT

`func (o *AuditActiveStatsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditActiveStatsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditActiveStatsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


