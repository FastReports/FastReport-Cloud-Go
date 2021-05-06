# DataSourcesVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSources** | Pointer to [**[]DataSourceVM**](DataSourceVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 

## Methods

### NewDataSourcesVM

`func NewDataSourcesVM() *DataSourcesVM`

NewDataSourcesVM instantiates a new DataSourcesVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourcesVMWithDefaults

`func NewDataSourcesVMWithDefaults() *DataSourcesVM`

NewDataSourcesVMWithDefaults instantiates a new DataSourcesVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSources

`func (o *DataSourcesVM) GetDataSources() []DataSourceVM`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *DataSourcesVM) GetDataSourcesOk() (*[]DataSourceVM, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *DataSourcesVM) SetDataSources(v []DataSourceVM)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *DataSourcesVM) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### GetCount

`func (o *DataSourcesVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DataSourcesVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DataSourcesVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *DataSourcesVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *DataSourcesVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *DataSourcesVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *DataSourcesVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *DataSourcesVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *DataSourcesVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *DataSourcesVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *DataSourcesVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *DataSourcesVM) HasTake() bool`

HasTake returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


