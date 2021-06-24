# CreateDataSourceVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ConnectionString** | **string** |  | 
**SubscriptionId** | **string** |  | 
**ConnectionType** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateDataSourceVM

`func NewCreateDataSourceVM(connectionString string, subscriptionId string, ) *CreateDataSourceVM`

NewCreateDataSourceVM instantiates a new CreateDataSourceVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDataSourceVMWithDefaults

`func NewCreateDataSourceVMWithDefaults() *CreateDataSourceVM`

NewCreateDataSourceVMWithDefaults instantiates a new CreateDataSourceVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDataSourceVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDataSourceVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDataSourceVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDataSourceVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnectionString

`func (o *CreateDataSourceVM) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *CreateDataSourceVM) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *CreateDataSourceVM) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.


### GetSubscriptionId

`func (o *CreateDataSourceVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateDataSourceVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateDataSourceVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetConnectionType

`func (o *CreateDataSourceVM) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *CreateDataSourceVM) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *CreateDataSourceVM) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *CreateDataSourceVM) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


