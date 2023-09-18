# UpdateTransformTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFile** | Pointer to [**InputFileVM**](InputFileVM.md) |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**OutputFile** | Pointer to [**OutputFileVM**](OutputFileVM.md) |  | [optional] 
**TransportIds** | Pointer to **[]string** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateTransformTaskBaseVM

`func NewUpdateTransformTaskBaseVM(t string, ) *UpdateTransformTaskBaseVM`

NewUpdateTransformTaskBaseVM instantiates a new UpdateTransformTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTransformTaskBaseVMWithDefaults

`func NewUpdateTransformTaskBaseVMWithDefaults() *UpdateTransformTaskBaseVM`

NewUpdateTransformTaskBaseVMWithDefaults instantiates a new UpdateTransformTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFile

`func (o *UpdateTransformTaskBaseVM) GetInputFile() InputFileVM`

GetInputFile returns the InputFile field if non-nil, zero value otherwise.

### GetInputFileOk

`func (o *UpdateTransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool)`

GetInputFileOk returns a tuple with the InputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFile

`func (o *UpdateTransformTaskBaseVM) SetInputFile(v InputFileVM)`

SetInputFile sets InputFile field to given value.

### HasInputFile

`func (o *UpdateTransformTaskBaseVM) HasInputFile() bool`

HasInputFile returns a boolean if a field has been set.

### GetLocale

`func (o *UpdateTransformTaskBaseVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateTransformTaskBaseVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateTransformTaskBaseVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateTransformTaskBaseVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UpdateTransformTaskBaseVM) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UpdateTransformTaskBaseVM) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetOutputFile

`func (o *UpdateTransformTaskBaseVM) GetOutputFile() OutputFileVM`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *UpdateTransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *UpdateTransformTaskBaseVM) SetOutputFile(v OutputFileVM)`

SetOutputFile sets OutputFile field to given value.

### HasOutputFile

`func (o *UpdateTransformTaskBaseVM) HasOutputFile() bool`

HasOutputFile returns a boolean if a field has been set.

### GetTransportIds

`func (o *UpdateTransformTaskBaseVM) GetTransportIds() []string`

GetTransportIds returns the TransportIds field if non-nil, zero value otherwise.

### GetTransportIdsOk

`func (o *UpdateTransformTaskBaseVM) GetTransportIdsOk() (*[]string, bool)`

GetTransportIdsOk returns a tuple with the TransportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportIds

`func (o *UpdateTransformTaskBaseVM) SetTransportIds(v []string)`

SetTransportIds sets TransportIds field to given value.

### HasTransportIds

`func (o *UpdateTransformTaskBaseVM) HasTransportIds() bool`

HasTransportIds returns a boolean if a field has been set.

### SetTransportIdsNil

`func (o *UpdateTransformTaskBaseVM) SetTransportIdsNil(b bool)`

 SetTransportIdsNil sets the value for TransportIds to be an explicit nil

### UnsetTransportIds
`func (o *UpdateTransformTaskBaseVM) UnsetTransportIds()`

UnsetTransportIds ensures that no value is present for TransportIds, not even an explicit nil
### GetT

`func (o *UpdateTransformTaskBaseVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateTransformTaskBaseVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateTransformTaskBaseVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


