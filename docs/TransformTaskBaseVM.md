# TransformTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFile** | Pointer to [**InputFileVM**](InputFileVM.md) |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**OutputFile** | Pointer to [**OutputFileVM**](OutputFileVM.md) |  | [optional] 
**TransportIds** | Pointer to **[]string** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewTransformTaskBaseVM

`func NewTransformTaskBaseVM(t string, ) *TransformTaskBaseVM`

NewTransformTaskBaseVM instantiates a new TransformTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformTaskBaseVMWithDefaults

`func NewTransformTaskBaseVMWithDefaults() *TransformTaskBaseVM`

NewTransformTaskBaseVMWithDefaults instantiates a new TransformTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFile

`func (o *TransformTaskBaseVM) GetInputFile() InputFileVM`

GetInputFile returns the InputFile field if non-nil, zero value otherwise.

### GetInputFileOk

`func (o *TransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool)`

GetInputFileOk returns a tuple with the InputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFile

`func (o *TransformTaskBaseVM) SetInputFile(v InputFileVM)`

SetInputFile sets InputFile field to given value.

### HasInputFile

`func (o *TransformTaskBaseVM) HasInputFile() bool`

HasInputFile returns a boolean if a field has been set.

### GetLocale

`func (o *TransformTaskBaseVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TransformTaskBaseVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TransformTaskBaseVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TransformTaskBaseVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *TransformTaskBaseVM) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *TransformTaskBaseVM) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetOutputFile

`func (o *TransformTaskBaseVM) GetOutputFile() OutputFileVM`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *TransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *TransformTaskBaseVM) SetOutputFile(v OutputFileVM)`

SetOutputFile sets OutputFile field to given value.

### HasOutputFile

`func (o *TransformTaskBaseVM) HasOutputFile() bool`

HasOutputFile returns a boolean if a field has been set.

### GetTransportIds

`func (o *TransformTaskBaseVM) GetTransportIds() []string`

GetTransportIds returns the TransportIds field if non-nil, zero value otherwise.

### GetTransportIdsOk

`func (o *TransformTaskBaseVM) GetTransportIdsOk() (*[]string, bool)`

GetTransportIdsOk returns a tuple with the TransportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportIds

`func (o *TransformTaskBaseVM) SetTransportIds(v []string)`

SetTransportIds sets TransportIds field to given value.

### HasTransportIds

`func (o *TransformTaskBaseVM) HasTransportIds() bool`

HasTransportIds returns a boolean if a field has been set.

### SetTransportIdsNil

`func (o *TransformTaskBaseVM) SetTransportIdsNil(b bool)`

 SetTransportIdsNil sets the value for TransportIds to be an explicit nil

### UnsetTransportIds
`func (o *TransformTaskBaseVM) UnsetTransportIds()`

UnsetTransportIds ensures that no value is present for TransportIds, not even an explicit nil
### GetT

`func (o *TransformTaskBaseVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TransformTaskBaseVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TransformTaskBaseVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


