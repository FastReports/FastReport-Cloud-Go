# RunFetchTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | Pointer to [**DataSourceConnectionType**](DataSourceConnectionType.md) |  | [optional] 
**ConnectionString** | **string** |  | 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**TaskType**](TaskType.md) |  | [optional] 

## Methods

### NewRunFetchTaskVM

`func NewRunFetchTaskVM(connectionString string, ) *RunFetchTaskVM`

NewRunFetchTaskVM instantiates a new RunFetchTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFetchTaskVMWithDefaults

`func NewRunFetchTaskVMWithDefaults() *RunFetchTaskVM`

NewRunFetchTaskVMWithDefaults instantiates a new RunFetchTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *RunFetchTaskVM) GetConnectionType() DataSourceConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *RunFetchTaskVM) GetConnectionTypeOk() (*DataSourceConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *RunFetchTaskVM) SetConnectionType(v DataSourceConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *RunFetchTaskVM) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetConnectionString

`func (o *RunFetchTaskVM) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *RunFetchTaskVM) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *RunFetchTaskVM) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetSubscriptionId

`func (o *RunFetchTaskVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RunFetchTaskVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RunFetchTaskVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RunFetchTaskVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *RunFetchTaskVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *RunFetchTaskVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetType

`func (o *RunFetchTaskVM) GetType() TaskType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunFetchTaskVM) GetTypeOk() (*TaskType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunFetchTaskVM) SetType(v TaskType)`

SetType sets Type field to given value.

### HasType

`func (o *RunFetchTaskVM) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


