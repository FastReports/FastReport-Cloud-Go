# CreateTransportTaskBaseVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | Pointer to [**[]InputFileVM**](InputFileVM.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewCreateTransportTaskBaseVM

`func NewCreateTransportTaskBaseVM() *CreateTransportTaskBaseVM`

NewCreateTransportTaskBaseVM instantiates a new CreateTransportTaskBaseVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransportTaskBaseVMWithDefaults

`func NewCreateTransportTaskBaseVMWithDefaults() *CreateTransportTaskBaseVM`

NewCreateTransportTaskBaseVMWithDefaults instantiates a new CreateTransportTaskBaseVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiles

`func (o *CreateTransportTaskBaseVM) GetFiles() []InputFileVM`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CreateTransportTaskBaseVM) GetFilesOk() (*[]InputFileVM, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CreateTransportTaskBaseVM) SetFiles(v []InputFileVM)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CreateTransportTaskBaseVM) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### SetFilesNil

`func (o *CreateTransportTaskBaseVM) SetFilesNil(b bool)`

 SetFilesNil sets the value for Files to be an explicit nil

### UnsetFiles
`func (o *CreateTransportTaskBaseVM) UnsetFiles()`

UnsetFiles ensures that no value is present for Files, not even an explicit nil
### GetName

`func (o *CreateTransportTaskBaseVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTransportTaskBaseVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTransportTaskBaseVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTransportTaskBaseVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTransportTaskBaseVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTransportTaskBaseVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *CreateTransportTaskBaseVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateTransportTaskBaseVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateTransportTaskBaseVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateTransportTaskBaseVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *CreateTransportTaskBaseVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *CreateTransportTaskBaseVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *CreateTransportTaskBaseVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateTransportTaskBaseVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateTransportTaskBaseVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateTransportTaskBaseVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


