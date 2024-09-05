# FrontendAppVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mixins** | Pointer to [**AppMixinsVM**](AppMixinsVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFrontendAppVM

`func NewFrontendAppVM(t string, ) *FrontendAppVM`

NewFrontendAppVM instantiates a new FrontendAppVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrontendAppVMWithDefaults

`func NewFrontendAppVMWithDefaults() *FrontendAppVM`

NewFrontendAppVMWithDefaults instantiates a new FrontendAppVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMixins

`func (o *FrontendAppVM) GetMixins() AppMixinsVM`

GetMixins returns the Mixins field if non-nil, zero value otherwise.

### GetMixinsOk

`func (o *FrontendAppVM) GetMixinsOk() (*AppMixinsVM, bool)`

GetMixinsOk returns a tuple with the Mixins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMixins

`func (o *FrontendAppVM) SetMixins(v AppMixinsVM)`

SetMixins sets Mixins field to given value.

### HasMixins

`func (o *FrontendAppVM) HasMixins() bool`

HasMixins returns a boolean if a field has been set.

### GetT

`func (o *FrontendAppVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FrontendAppVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FrontendAppVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


