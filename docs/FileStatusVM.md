# FileStatusVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileStatus** | Pointer to [**FileStatus**](FileStatus.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFileStatusVM

`func NewFileStatusVM(t string, ) *FileStatusVM`

NewFileStatusVM instantiates a new FileStatusVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStatusVMWithDefaults

`func NewFileStatusVMWithDefaults() *FileStatusVM`

NewFileStatusVMWithDefaults instantiates a new FileStatusVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileStatus

`func (o *FileStatusVM) GetFileStatus() FileStatus`

GetFileStatus returns the FileStatus field if non-nil, zero value otherwise.

### GetFileStatusOk

`func (o *FileStatusVM) GetFileStatusOk() (*FileStatus, bool)`

GetFileStatusOk returns a tuple with the FileStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileStatus

`func (o *FileStatusVM) SetFileStatus(v FileStatus)`

SetFileStatus sets FileStatus field to given value.

### HasFileStatus

`func (o *FileStatusVM) HasFileStatus() bool`

HasFileStatus returns a boolean if a field has been set.

### GetT

`func (o *FileStatusVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FileStatusVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FileStatusVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


