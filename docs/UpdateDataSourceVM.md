# UpdateDataSourceVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**ConnectionString** | Pointer to **NullableString** |  | [optional] 
**SelectCommands** | Pointer to [**[]DataSourceSelectCommandVM**](DataSourceSelectCommandVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateDataSourceVM

`func NewUpdateDataSourceVM(t string, ) *UpdateDataSourceVM`

NewUpdateDataSourceVM instantiates a new UpdateDataSourceVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDataSourceVMWithDefaults

`func NewUpdateDataSourceVMWithDefaults() *UpdateDataSourceVM`

NewUpdateDataSourceVMWithDefaults instantiates a new UpdateDataSourceVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDataSourceVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDataSourceVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDataSourceVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDataSourceVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateDataSourceVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateDataSourceVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *UpdateDataSourceVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UpdateDataSourceVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UpdateDataSourceVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UpdateDataSourceVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *UpdateDataSourceVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *UpdateDataSourceVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetConnectionString

`func (o *UpdateDataSourceVM) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *UpdateDataSourceVM) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *UpdateDataSourceVM) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *UpdateDataSourceVM) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### SetConnectionStringNil

`func (o *UpdateDataSourceVM) SetConnectionStringNil(b bool)`

 SetConnectionStringNil sets the value for ConnectionString to be an explicit nil

### UnsetConnectionString
`func (o *UpdateDataSourceVM) UnsetConnectionString()`

UnsetConnectionString ensures that no value is present for ConnectionString, not even an explicit nil
### GetSelectCommands

`func (o *UpdateDataSourceVM) GetSelectCommands() []DataSourceSelectCommandVM`

GetSelectCommands returns the SelectCommands field if non-nil, zero value otherwise.

### GetSelectCommandsOk

`func (o *UpdateDataSourceVM) GetSelectCommandsOk() (*[]DataSourceSelectCommandVM, bool)`

GetSelectCommandsOk returns a tuple with the SelectCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectCommands

`func (o *UpdateDataSourceVM) SetSelectCommands(v []DataSourceSelectCommandVM)`

SetSelectCommands sets SelectCommands field to given value.

### HasSelectCommands

`func (o *UpdateDataSourceVM) HasSelectCommands() bool`

HasSelectCommands returns a boolean if a field has been set.

### SetSelectCommandsNil

`func (o *UpdateDataSourceVM) SetSelectCommandsNil(b bool)`

 SetSelectCommandsNil sets the value for SelectCommands to be an explicit nil

### UnsetSelectCommands
`func (o *UpdateDataSourceVM) UnsetSelectCommands()`

UnsetSelectCommands ensures that no value is present for SelectCommands, not even an explicit nil
### GetT

`func (o *UpdateDataSourceVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateDataSourceVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateDataSourceVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


