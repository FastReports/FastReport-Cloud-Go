# AuditStatVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AuditType**](AuditType.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditStatVM

`func NewAuditStatVM(t string, ) *AuditStatVM`

NewAuditStatVM instantiates a new AuditStatVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditStatVMWithDefaults

`func NewAuditStatVMWithDefaults() *AuditStatVM`

NewAuditStatVMWithDefaults instantiates a new AuditStatVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *AuditStatVM) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AuditStatVM) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AuditStatVM) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AuditStatVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetType

`func (o *AuditStatVM) GetType() AuditType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditStatVM) GetTypeOk() (*AuditType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditStatVM) SetType(v AuditType)`

SetType sets Type field to given value.

### HasType

`func (o *AuditStatVM) HasType() bool`

HasType returns a boolean if a field has been set.

### GetT

`func (o *AuditStatVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditStatVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditStatVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


