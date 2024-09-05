# AuditTaskActionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAuditTaskActionVM

`func NewAuditTaskActionVM(t string, ) *AuditTaskActionVM`

NewAuditTaskActionVM instantiates a new AuditTaskActionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditTaskActionVMWithDefaults

`func NewAuditTaskActionVMWithDefaults() *AuditTaskActionVM`

NewAuditTaskActionVMWithDefaults instantiates a new AuditTaskActionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *AuditTaskActionVM) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *AuditTaskActionVM) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *AuditTaskActionVM) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *AuditTaskActionVM) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### SetMessageIdNil

`func (o *AuditTaskActionVM) SetMessageIdNil(b bool)`

 SetMessageIdNil sets the value for MessageId to be an explicit nil

### UnsetMessageId
`func (o *AuditTaskActionVM) UnsetMessageId()`

UnsetMessageId ensures that no value is present for MessageId, not even an explicit nil
### GetFileName

`func (o *AuditTaskActionVM) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *AuditTaskActionVM) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *AuditTaskActionVM) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *AuditTaskActionVM) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *AuditTaskActionVM) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *AuditTaskActionVM) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetT

`func (o *AuditTaskActionVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AuditTaskActionVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AuditTaskActionVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


