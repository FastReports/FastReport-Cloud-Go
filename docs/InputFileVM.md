# InputFileVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**FileKind**](FileKind.md) |  | [optional] 

## Methods

### NewInputFileVM

`func NewInputFileVM() *InputFileVM`

NewInputFileVM instantiates a new InputFileVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFileVMWithDefaults

`func NewInputFileVMWithDefaults() *InputFileVM`

NewInputFileVMWithDefaults instantiates a new InputFileVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *InputFileVM) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *InputFileVM) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *InputFileVM) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *InputFileVM) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *InputFileVM) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *InputFileVM) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetFileName

`func (o *InputFileVM) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *InputFileVM) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *InputFileVM) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *InputFileVM) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *InputFileVM) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *InputFileVM) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetType

`func (o *InputFileVM) GetType() FileKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputFileVM) GetTypeOk() (*FileKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputFileVM) SetType(v FileKind)`

SetType sets Type field to given value.

### HasType

`func (o *InputFileVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


