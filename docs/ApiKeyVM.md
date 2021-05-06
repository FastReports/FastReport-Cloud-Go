# ApiKeyVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Expired** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApiKeyVM

`func NewApiKeyVM() *ApiKeyVM`

NewApiKeyVM instantiates a new ApiKeyVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyVMWithDefaults

`func NewApiKeyVMWithDefaults() *ApiKeyVM`

NewApiKeyVMWithDefaults instantiates a new ApiKeyVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ApiKeyVM) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiKeyVM) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiKeyVM) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiKeyVM) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyVM) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyVM) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyVM) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyVM) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpired

`func (o *ApiKeyVM) GetExpired() time.Time`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ApiKeyVM) GetExpiredOk() (*time.Time, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ApiKeyVM) SetExpired(v time.Time)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ApiKeyVM) HasExpired() bool`

HasExpired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


