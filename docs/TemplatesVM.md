# TemplatesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TemplateVM**](TemplateVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTemplatesVM

`func NewTemplatesVM(t string, ) *TemplatesVM`

NewTemplatesVM instantiates a new TemplatesVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesVMWithDefaults

`func NewTemplatesVMWithDefaults() *TemplatesVM`

NewTemplatesVMWithDefaults instantiates a new TemplatesVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *TemplatesVM) GetFiles() []TemplateVM`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *TemplatesVM) GetFilesOk() (*[]TemplateVM, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *TemplatesVM) SetFiles(v []TemplateVM)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *TemplatesVM) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *TemplatesVM) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *TemplatesVM) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetT

`func (o *TemplatesVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TemplatesVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TemplatesVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


