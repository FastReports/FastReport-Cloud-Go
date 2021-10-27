# CreateTransformTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | Pointer to **NullableString** |  | [optional] 
**InputFile** | Pointer to [**InputFileVM**](InputFileVM.md) |  | [optional] 
**OutputFile** | Pointer to [**OutputFileVM**](OutputFileVM.md) |  | [optional] 
**Transports** | Pointer to [**[]CreateTransportTaskBaseVM**](CreateTransportTaskBaseVM.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewCreateTransformTaskBaseVM

`func NewCreateTransformTaskBaseVM() *CreateTransformTaskBaseVM`

NewCreateTransformTaskBaseVM instantiates a new CreateTransformTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransformTaskBaseVMWithDefaults

`func NewCreateTransformTaskBaseVMWithDefaults() *CreateTransformTaskBaseVM`

NewCreateTransformTaskBaseVMWithDefaults instantiates a new CreateTransformTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetName

`func (o *CreateTransformTaskBaseVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransformTaskBaseVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransformTaskBaseVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTransformTaskBaseVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTransformTaskBaseVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTransformTaskBaseVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *CreateTransformTaskBaseVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateTransformTaskBaseVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateTransformTaskBaseVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateTransformTaskBaseVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateTransformTaskBaseVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateTransformTaskBaseVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *CreateTransformTaskBaseVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTransformTaskBaseVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTransformTaskBaseVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTransformTaskBaseVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


