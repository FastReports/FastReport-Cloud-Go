# FileCreateFormVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to **Nullable*os.File** |  | [optional] 
**FileContent** | ***os.File** |  | 
**T** | **string** |  | 

## Methods

### NewFileCreateFormVM

`func NewFileCreateFormVM(fileContent *os.File, t string, ) *FileCreateFormVM`

NewFileCreateFormVM instantiates a new FileCreateFormVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCreateFormVMWithDefaults

`func NewFileCreateFormVMWithDefaults() *FileCreateFormVM`

NewFileCreateFormVMWithDefaults instantiates a new FileCreateFormVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *FileCreateFormVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FileCreateFormVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FileCreateFormVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FileCreateFormVM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *FileCreateFormVM) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *FileCreateFormVM) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIcon

`func (o *FileCreateFormVM) GetIcon() *os.File`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FileCreateFormVM) GetIconOk() (**os.File, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FileCreateFormVM) SetIcon(v *os.File)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FileCreateFormVM) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *FileCreateFormVM) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *FileCreateFormVM) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetFileContent

`func (o *FileCreateFormVM) GetFileContent() *os.File`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *FileCreateFormVM) GetFileContentOk() (**os.File, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *FileCreateFormVM) SetFileContent(v *os.File)`

SetFileContent sets FileContent field to given value.


### GetT

`func (o *FileCreateFormVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FileCreateFormVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FileCreateFormVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


