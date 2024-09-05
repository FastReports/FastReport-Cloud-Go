# ExportsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]ExportVM**](ExportVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewExportsVM

`func NewExportsVM(t string, ) *ExportsVM`

NewExportsVM instantiates a new ExportsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportsVMWithDefaults

`func NewExportsVMWithDefaults() *ExportsVM`

NewExportsVMWithDefaults instantiates a new ExportsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *ExportsVM) GetFiles() []ExportVM`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ExportsVM) GetFilesOk() (*[]ExportVM, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ExportsVM) SetFiles(v []ExportVM)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ExportsVM) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *ExportsVM) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *ExportsVM) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetT

`func (o *ExportsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *ExportsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *ExportsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


