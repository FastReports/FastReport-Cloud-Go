# FileRenameVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFileRenameVM

`func NewFileRenameVM(t string, ) *FileRenameVM`

NewFileRenameVM instantiates a new FileRenameVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRenameVMWithDefaults

`func NewFileRenameVMWithDefaults() *FileRenameVM`

NewFileRenameVMWithDefaults instantiates a new FileRenameVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileRenameVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileRenameVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileRenameVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileRenameVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileRenameVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileRenameVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetT

`func (o *FileRenameVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FileRenameVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FileRenameVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


