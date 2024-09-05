# AuditFilePropertyChangedVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **NullableString** |  | [optional] 
**OldValue** | Pointer to **interface{}** |  | [optional] 
**NewValue** | Pointer to **interface{}** |  | [optional] 
**EntityType** | Pointer to [**EntityType**](EntityType.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditFilePropertyChangedVM

`func NewAuditFilePropertyChangedVM(t string, ) *AuditFilePropertyChangedVM`

NewAuditFilePropertyChangedVM instantiates a new AuditFilePropertyChangedVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditFilePropertyChangedVMWithDefaults

`func NewAuditFilePropertyChangedVMWithDefaults() *AuditFilePropertyChangedVM`

NewAuditFilePropertyChangedVMWithDefaults instantiates a new AuditFilePropertyChangedVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *AuditFilePropertyChangedVM) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *AuditFilePropertyChangedVM) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *AuditFilePropertyChangedVM) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *AuditFilePropertyChangedVM) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### SetPropertyNameNil

`func (o *AuditFilePropertyChangedVM) SetPropertyNameNil(b bool)`

 SetPropertyNameNil sets the value for PropertyName to be an explicit nil

### UnsetPropertyName
`func (o *AuditFilePropertyChangedVM) UnsetPropertyName()`

UnsetPropertyName ensures that no value is present for PropertyName, not even an explicit nil
### GetOldValue

`func (o *AuditFilePropertyChangedVM) GetOldValue() interface{}`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *AuditFilePropertyChangedVM) GetOldValueOk() (*interface{}, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *AuditFilePropertyChangedVM) SetOldValue(v interface{})`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *AuditFilePropertyChangedVM) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### SetOldValueNil

`func (o *AuditFilePropertyChangedVM) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *AuditFilePropertyChangedVM) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
### GetNewValue

`func (o *AuditFilePropertyChangedVM) GetNewValue() interface{}`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *AuditFilePropertyChangedVM) GetNewValueOk() (*interface{}, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *AuditFilePropertyChangedVM) SetNewValue(v interface{})`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *AuditFilePropertyChangedVM) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *AuditFilePropertyChangedVM) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *AuditFilePropertyChangedVM) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetEntityType

`func (o *AuditFilePropertyChangedVM) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *AuditFilePropertyChangedVM) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *AuditFilePropertyChangedVM) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *AuditFilePropertyChangedVM) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetT

`func (o *AuditFilePropertyChangedVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditFilePropertyChangedVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditFilePropertyChangedVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


