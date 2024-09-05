# FetchTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSourceId** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFetchTaskVM

`func NewFetchTaskVM(t string, ) *FetchTaskVM`

NewFetchTaskVM instantiates a new FetchTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchTaskVMWithDefaults

`func NewFetchTaskVMWithDefaults() *FetchTaskVM`

NewFetchTaskVMWithDefaults instantiates a new FetchTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSourceId

`func (o *FetchTaskVM) GetDataSourceId() string`

GetDataSourceId returns the DataSourceId field if non-nil, zero value otherwise.

### GetDataSourceIdOk

`func (o *FetchTaskVM) GetDataSourceIdOk() (*string, bool)`

GetDataSourceIdOk returns a tuple with the DataSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceId

`func (o *FetchTaskVM) SetDataSourceId(v string)`

SetDataSourceId sets DataSourceId field to given value.

### HasDataSourceId

`func (o *FetchTaskVM) HasDataSourceId() bool`

HasDataSourceId returns a boolean if a field has been set.

### SetDataSourceIdNil

`func (o *FetchTaskVM) SetDataSourceIdNil(b bool)`

 SetDataSourceIdNil sets the value for DataSourceId to be an explicit nil

### UnsetDataSourceId
`func (o *FetchTaskVM) UnsetDataSourceId()`

UnsetDataSourceId ensures that no value is present for DataSourceId, not even an explicit nil
### GetT

`func (o *FetchTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FetchTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FetchTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


