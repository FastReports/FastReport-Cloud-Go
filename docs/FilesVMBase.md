# FilesVMBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]FileVM**](FileVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFilesVMBase

`func NewFilesVMBase(t string, ) *FilesVMBase`

NewFilesVMBase instantiates a new FilesVMBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesVMBaseWithDefaults

`func NewFilesVMBaseWithDefaults() *FilesVMBase`

NewFilesVMBaseWithDefaults instantiates a new FilesVMBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FilesVMBase) GetFiles() []FileVM`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FilesVMBase) GetFilesOk() (*[]FileVM, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FilesVMBase) SetFiles(v []FileVM)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FilesVMBase) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *FilesVMBase) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *FilesVMBase) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetCount

`func (o *FilesVMBase) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FilesVMBase) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FilesVMBase) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *FilesVMBase) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *FilesVMBase) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *FilesVMBase) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *FilesVMBase) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *FilesVMBase) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *FilesVMBase) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *FilesVMBase) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *FilesVMBase) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *FilesVMBase) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetT

`func (o *FilesVMBase) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FilesVMBase) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FilesVMBase) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


