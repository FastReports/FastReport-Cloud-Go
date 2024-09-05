# DataSourceSelectCommandParameterVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DataType** | **string** |  | 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewDataSourceSelectCommandParameterVM

`func NewDataSourceSelectCommandParameterVM(name string, dataType string, t string, ) *DataSourceSelectCommandParameterVM`

NewDataSourceSelectCommandParameterVM instantiates a new DataSourceSelectCommandParameterVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceSelectCommandParameterVMWithDefaults

`func NewDataSourceSelectCommandParameterVMWithDefaults() *DataSourceSelectCommandParameterVM`

NewDataSourceSelectCommandParameterVMWithDefaults instantiates a new DataSourceSelectCommandParameterVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataSourceSelectCommandParameterVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceSelectCommandParameterVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceSelectCommandParameterVM) SetName(v string)`

SetName sets Name field to given value.


### GetDataType

`func (o *DataSourceSelectCommandParameterVM) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DataSourceSelectCommandParameterVM) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DataSourceSelectCommandParameterVM) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetDefaultValue

`func (o *DataSourceSelectCommandParameterVM) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DataSourceSelectCommandParameterVM) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DataSourceSelectCommandParameterVM) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DataSourceSelectCommandParameterVM) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *DataSourceSelectCommandParameterVM) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *DataSourceSelectCommandParameterVM) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetSize

`func (o *DataSourceSelectCommandParameterVM) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DataSourceSelectCommandParameterVM) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DataSourceSelectCommandParameterVM) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DataSourceSelectCommandParameterVM) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetT

`func (o *DataSourceSelectCommandParameterVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *DataSourceSelectCommandParameterVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *DataSourceSelectCommandParameterVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


