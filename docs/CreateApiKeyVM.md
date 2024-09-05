# CreateApiKeyVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Expired** | **time.Time** |  | 
**T** | **string** |  | 

## Methods

### NewCreateApiKeyVM

`func NewCreateApiKeyVM(expired time.Time, t string, ) *CreateApiKeyVM`

NewCreateApiKeyVM instantiates a new CreateApiKeyVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyVMWithDefaults

`func NewCreateApiKeyVMWithDefaults() *CreateApiKeyVM`

NewCreateApiKeyVMWithDefaults instantiates a new CreateApiKeyVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateApiKeyVM) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApiKeyVM) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApiKeyVM) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateApiKeyVM) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateApiKeyVM) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateApiKeyVM) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExpired

`func (o *CreateApiKeyVM) GetExpired() time.Time`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *CreateApiKeyVM) GetExpiredOk() (*time.Time, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *CreateApiKeyVM) SetExpired(v time.Time)`

SetExpired sets Expired field to given value.


### GetT

`func (o *CreateApiKeyVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateApiKeyVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateApiKeyVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


