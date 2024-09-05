# DataSourceSelectCommandVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Command** | **string** |  | 
**Parameters** | Pointer to [**[]DataSourceSelectCommandParameterVM**](DataSourceSelectCommandParameterVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewDataSourceSelectCommandVM

`func NewDataSourceSelectCommandVM(name string, command string, t string, ) *DataSourceSelectCommandVM`

NewDataSourceSelectCommandVM instantiates a new DataSourceSelectCommandVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceSelectCommandVMWithDefaults

`func NewDataSourceSelectCommandVMWithDefaults() *DataSourceSelectCommandVM`

NewDataSourceSelectCommandVMWithDefaults instantiates a new DataSourceSelectCommandVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DataSourceSelectCommandVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceSelectCommandVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceSelectCommandVM) SetName(v string)`

SetName sets Name field to given value.


### GetCommand

`func (o *DataSourceSelectCommandVM) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *DataSourceSelectCommandVM) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *DataSourceSelectCommandVM) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetParameters

`func (o *DataSourceSelectCommandVM) GetParameters() []DataSourceSelectCommandParameterVM`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DataSourceSelectCommandVM) GetParametersOk() (*[]DataSourceSelectCommandParameterVM, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DataSourceSelectCommandVM) SetParameters(v []DataSourceSelectCommandParameterVM)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DataSourceSelectCommandVM) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *DataSourceSelectCommandVM) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *DataSourceSelectCommandVM) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
### GetT

`func (o *DataSourceSelectCommandVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *DataSourceSelectCommandVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *DataSourceSelectCommandVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


