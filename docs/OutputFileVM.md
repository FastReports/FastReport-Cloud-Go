# OutputFileVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **NullableString** |  | [optional] 
**FolderId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**FileKind**](FileKind.md) |  | [optional] 

## Methods

### NewOutputFileVM

`func NewOutputFileVM() *OutputFileVM`

NewOutputFileVM instantiates a new OutputFileVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputFileVMWithDefaults

`func NewOutputFileVMWithDefaults() *OutputFileVM`

NewOutputFileVMWithDefaults instantiates a new OutputFileVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *OutputFileVM) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *OutputFileVM) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *OutputFileVM) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *OutputFileVM) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *OutputFileVM) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *OutputFileVM) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetFolderId

`func (o *OutputFileVM) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *OutputFileVM) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *OutputFileVM) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *OutputFileVM) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *OutputFileVM) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *OutputFileVM) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetType

`func (o *OutputFileVM) GetType() FileKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputFileVM) GetTypeOk() (*FileKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputFileVM) SetType(v FileKind)`

SetType sets Type field to given value.

### HasType

`func (o *OutputFileVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


