# RunTransformTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **NullableString** |  | [optional] 
**InputFile** | Pointer to [**RunInputFileVM**](RunInputFileVM.md) |  | [optional] 
**OutputFile** | Pointer to [**OutputFileVM**](OutputFileVM.md) |  | [optional] 
**Transports** | Pointer to [**[]RunTransportTaskBaseVM**](RunTransportTaskBaseVM.md) |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewRunTransformTaskBaseVM

`func NewRunTransformTaskBaseVM() *RunTransformTaskBaseVM`

NewRunTransformTaskBaseVM instantiates a new RunTransformTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunTransformTaskBaseVMWithDefaults

`func NewRunTransformTaskBaseVMWithDefaults() *RunTransformTaskBaseVM`

NewRunTransformTaskBaseVMWithDefaults instantiates a new RunTransformTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocale

`func (o *RunTransformTaskBaseVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *RunTransformTaskBaseVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *RunTransformTaskBaseVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *RunTransformTaskBaseVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *RunTransformTaskBaseVM) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *RunTransformTaskBaseVM) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetInputFile

`func (o *RunTransformTaskBaseVM) GetInputFile() RunInputFileVM`

GetInputFile returns the InputFile field if non-nil, zero value otherwise.

### GetInputFileOk

`func (o *RunTransformTaskBaseVM) GetInputFileOk() (*RunInputFileVM, bool)`

GetInputFileOk returns a tuple with the InputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFile

`func (o *RunTransformTaskBaseVM) SetInputFile(v RunInputFileVM)`

SetInputFile sets InputFile field to given value.

### HasInputFile

`func (o *RunTransformTaskBaseVM) HasInputFile() bool`

HasInputFile returns a boolean if a field has been set.

### GetOutputFile

`func (o *RunTransformTaskBaseVM) GetOutputFile() OutputFileVM`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *RunTransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *RunTransformTaskBaseVM) SetOutputFile(v OutputFileVM)`

SetOutputFile sets OutputFile field to given value.

### HasOutputFile

`func (o *RunTransformTaskBaseVM) HasOutputFile() bool`

HasOutputFile returns a boolean if a field has been set.

### GetTransports

`func (o *RunTransformTaskBaseVM) GetTransports() []RunTransportTaskBaseVM`

GetTransports returns the Transports field if non-nil, zero value otherwise.

### GetTransportsOk

`func (o *RunTransformTaskBaseVM) GetTransportsOk() (*[]RunTransportTaskBaseVM, bool)`

GetTransportsOk returns a tuple with the Transports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransports

`func (o *RunTransformTaskBaseVM) SetTransports(v []RunTransportTaskBaseVM)`

SetTransports sets Transports field to given value.

### HasTransports

`func (o *RunTransformTaskBaseVM) HasTransports() bool`

HasTransports returns a boolean if a field has been set.

### SetTransportsNil

`func (o *RunTransformTaskBaseVM) SetTransportsNil(b bool)`

 SetTransportsNil sets the value for Transports to be an explicit nil

### UnsetTransports
`func (o *RunTransformTaskBaseVM) UnsetTransports()`

UnsetTransports ensures that no value is present for Transports, not even an explicit nil
### GetSubscriptionId

`func (o *RunTransformTaskBaseVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RunTransformTaskBaseVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RunTransformTaskBaseVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RunTransformTaskBaseVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *RunTransformTaskBaseVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *RunTransformTaskBaseVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *RunTransformTaskBaseVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunTransformTaskBaseVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunTransformTaskBaseVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *RunTransformTaskBaseVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


