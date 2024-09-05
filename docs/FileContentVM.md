# FileContentVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFileContentVM

`func NewFileContentVM(t string, ) *FileContentVM`

NewFileContentVM instantiates a new FileContentVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileContentVMWithDefaults

`func NewFileContentVMWithDefaults() *FileContentVM`

NewFileContentVMWithDefaults instantiates a new FileContentVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *FileContentVM) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FileContentVM) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FileContentVM) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FileContentVM) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *FileContentVM) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *FileContentVM) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetT

`func (o *FileContentVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FileContentVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FileContentVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


