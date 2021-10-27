# FilesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]FileVM**](FileVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 

## Methods

### NewFilesVM

`func NewFilesVM() *FilesVM`

NewFilesVM instantiates a new FilesVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesVMWithDefaults

`func NewFilesVMWithDefaults() *FilesVM`

NewFilesVMWithDefaults instantiates a new FilesVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *FilesVM) GetFiles() []FileVM`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *FilesVM) GetFilesOk() (*[]FileVM, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *FilesVM) SetFiles(v []FileVM)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *FilesVM) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *FilesVM) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *FilesVM) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetCount

`func (o *FilesVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FilesVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FilesVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *FilesVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *FilesVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *FilesVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *FilesVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *FilesVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *FilesVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *FilesVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *FilesVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *FilesVM) HasTake() bool`

HasTake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


