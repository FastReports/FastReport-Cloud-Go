# RunInputFileVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**EntityId** | Pointer to **NullableString** |  | [optional] 
**FileName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**FileKind**](FileKind.md) |  | [optional] 

## Methods

### NewRunInputFileVM

`func NewRunInputFileVM() *RunInputFileVM`

NewRunInputFileVM instantiates a new RunInputFileVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunInputFileVMWithDefaults

`func NewRunInputFileVMWithDefaults() *RunInputFileVM`

NewRunInputFileVMWithDefaults instantiates a new RunInputFileVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *RunInputFileVM) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RunInputFileVM) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RunInputFileVM) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RunInputFileVM) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *RunInputFileVM) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *RunInputFileVM) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEntityId

`func (o *RunInputFileVM) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *RunInputFileVM) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *RunInputFileVM) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *RunInputFileVM) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *RunInputFileVM) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *RunInputFileVM) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetFileName

`func (o *RunInputFileVM) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *RunInputFileVM) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *RunInputFileVM) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *RunInputFileVM) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *RunInputFileVM) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *RunInputFileVM) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetType

`func (o *RunInputFileVM) GetType() FileKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunInputFileVM) GetTypeOk() (*FileKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunInputFileVM) SetType(v FileKind)`

SetType sets Type field to given value.

### HasType

`func (o *RunInputFileVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


