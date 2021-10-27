# ServerConfigurationVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** |  | [optional] 
**CorporateServerMode** | Pointer to **bool** |  | [optional] 
**AppMixins** | Pointer to [**AppMixins**](AppMixins.md) |  | [optional] 

## Methods

### NewServerConfigurationVM

`func NewServerConfigurationVM() *ServerConfigurationVM`

NewServerConfigurationVM instantiates a new ServerConfigurationVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigurationVMWithDefaults

`func NewServerConfigurationVMWithDefaults() *ServerConfigurationVM`

NewServerConfigurationVMWithDefaults instantiates a new ServerConfigurationVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ServerConfigurationVM) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ServerConfigurationVM) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ServerConfigurationVM) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ServerConfigurationVM) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ServerConfigurationVM) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ServerConfigurationVM) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetCorporateServerMode

`func (o *ServerConfigurationVM) GetCorporateServerMode() bool`

GetCorporateServerMode returns the CorporateServerMode field if non-nil, zero value otherwise.

### GetCorporateServerModeOk

`func (o *ServerConfigurationVM) GetCorporateServerModeOk() (*bool, bool)`

GetCorporateServerModeOk returns a tuple with the CorporateServerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateServerMode

`func (o *ServerConfigurationVM) SetCorporateServerMode(v bool)`

SetCorporateServerMode sets CorporateServerMode field to given value.

### HasCorporateServerMode

`func (o *ServerConfigurationVM) HasCorporateServerMode() bool`

HasCorporateServerMode returns a boolean if a field has been set.

### GetAppMixins

`func (o *ServerConfigurationVM) GetAppMixins() AppMixins`

GetAppMixins returns the AppMixins field if non-nil, zero value otherwise.

### GetAppMixinsOk

`func (o *ServerConfigurationVM) GetAppMixinsOk() (*AppMixins, bool)`

GetAppMixinsOk returns a tuple with the AppMixins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppMixins

`func (o *ServerConfigurationVM) SetAppMixins(v AppMixins)`

SetAppMixins sets AppMixins field to given value.

### HasAppMixins

`func (o *ServerConfigurationVM) HasAppMixins() bool`

HasAppMixins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


