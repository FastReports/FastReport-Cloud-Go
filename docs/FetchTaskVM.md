# FetchTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to [**DataSourceConnectionType**](DataSourceConnectionType.md) |  | [optional] 
**ConnectionString** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewFetchTaskVM

`func NewFetchTaskVM(connectionString string, ) *FetchTaskVM`

NewFetchTaskVM instantiates a new FetchTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchTaskVMWithDefaults

`func NewFetchTaskVMWithDefaults() *FetchTaskVM`

NewFetchTaskVMWithDefaults instantiates a new FetchTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *FetchTaskVM) GetConnectionType() DataSourceConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *FetchTaskVM) GetConnectionTypeOk() (*DataSourceConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *FetchTaskVM) SetConnectionType(v DataSourceConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *FetchTaskVM) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetConnectionString

`func (o *FetchTaskVM) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *FetchTaskVM) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *FetchTaskVM) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetName

`func (o *FetchTaskVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FetchTaskVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FetchTaskVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FetchTaskVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FetchTaskVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FetchTaskVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubscriptionId

`func (o *FetchTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FetchTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FetchTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FetchTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *FetchTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *FetchTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *FetchTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FetchTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FetchTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *FetchTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


