# TemplatesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]TemplateVM**](TemplateVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 

## Methods

### NewTemplatesVM

`func NewTemplatesVM() *TemplatesVM`

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
### GetCount

`func (o *TemplatesVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TemplatesVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TemplatesVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *TemplatesVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *TemplatesVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *TemplatesVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *TemplatesVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *TemplatesVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *TemplatesVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *TemplatesVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *TemplatesVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *TemplatesVM) HasTake() bool`

HasTake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


