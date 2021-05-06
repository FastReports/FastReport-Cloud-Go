# ApiKeysVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | Pointer to [**[]ApiKeyVM**](ApiKeyVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiKeysVM

`func NewApiKeysVM() *ApiKeysVM`

NewApiKeysVM instantiates a new ApiKeysVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeysVMWithDefaults

`func NewApiKeysVMWithDefaults() *ApiKeysVM`

NewApiKeysVMWithDefaults instantiates a new ApiKeysVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeys

`func (o *ApiKeysVM) GetApiKeys() []ApiKeyVM`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *ApiKeysVM) GetApiKeysOk() (*[]ApiKeyVM, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *ApiKeysVM) SetApiKeys(v []ApiKeyVM)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *ApiKeysVM) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetCount

`func (o *ApiKeysVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiKeysVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiKeysVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ApiKeysVM) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


