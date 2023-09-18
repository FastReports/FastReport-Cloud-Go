# CreateTransformTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFile** | Pointer to [**InputFileVM**](InputFileVM.md) |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**OutputFile** | Pointer to [**OutputFileVM**](OutputFileVM.md) |  | [optional] 
**Transports** | Pointer to [**[]CreateTransportTaskBaseVM**](CreateTransportTaskBaseVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateTransformTaskBaseVM

`func NewCreateTransformTaskBaseVM(t string, ) *CreateTransformTaskBaseVM`

NewCreateTransformTaskBaseVM instantiates a new CreateTransformTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransformTaskBaseVMWithDefaults

`func NewCreateTransformTaskBaseVMWithDefaults() *CreateTransformTaskBaseVM`

NewCreateTransformTaskBaseVMWithDefaults instantiates a new CreateTransformTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFile

`func (o *CreateTransformTaskBaseVM) GetInputFile() InputFileVM`

GetInputFile returns the InputFile field if non-nil, zero value otherwise.

### GetInputFileOk

`func (o *CreateTransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool)`

GetInputFileOk returns a tuple with the InputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFile

`func (o *CreateTransformTaskBaseVM) SetInputFile(v InputFileVM)`

SetInputFile sets InputFile field to given value.

### HasInputFile

`func (o *CreateTransformTaskBaseVM) HasInputFile() bool`

HasInputFile returns a boolean if a field has been set.

### GetLocale

`func (o *CreateTransformTaskBaseVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateTransformTaskBaseVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateTransformTaskBaseVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateTransformTaskBaseVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *CreateTransformTaskBaseVM) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *CreateTransformTaskBaseVM) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetOutputFile

`func (o *CreateTransformTaskBaseVM) GetOutputFile() OutputFileVM`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *CreateTransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *CreateTransformTaskBaseVM) SetOutputFile(v OutputFileVM)`

SetOutputFile sets OutputFile field to given value.

### HasOutputFile

`func (o *CreateTransformTaskBaseVM) HasOutputFile() bool`

HasOutputFile returns a boolean if a field has been set.

### GetTransports

`func (o *CreateTransformTaskBaseVM) GetTransports() []CreateTransportTaskBaseVM`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *CreateTransformTaskBaseVM) GetTransportsOk() (*[]CreateTransportTaskBaseVM, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *CreateTransformTaskBaseVM) SetTransports(v []CreateTransportTaskBaseVM)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *CreateTransformTaskBaseVM) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### SetTransportsNil

`func (o *CreateTransformTaskBaseVM) SetTransportsNil(b bool)`

 SetTransportsNil sets the value for Transports to be an explicit nil

### UnsetTransports
`func (o *CreateTransformTaskBaseVM) UnsetTransports()`

UnsetTransports ensures that no value is present for Transports, not even an explicit nil
### GetT

`func (o *CreateTransformTaskBaseVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateTransformTaskBaseVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateTransformTaskBaseVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


