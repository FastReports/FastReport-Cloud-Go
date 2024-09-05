# AnalysisResultVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**ProblemLevel**](ProblemLevel.md) |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**FileId** | Pointer to **NullableString** |  | [optional] 
**CollectionName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**ProblemType**](ProblemType.md) |  | [optional] 
**Signature** | Pointer to **NullableString** |  | [optional] 
**NewSize** | Pointer to **int64** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewAnalysisResultVM

`func NewAnalysisResultVM(t string, ) *AnalysisResultVM`

NewAnalysisResultVM instantiates a new AnalysisResultVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisResultVMWithDefaults

`func NewAnalysisResultVMWithDefaults() *AnalysisResultVM`

NewAnalysisResultVMWithDefaults instantiates a new AnalysisResultVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *AnalysisResultVM) GetLevel() ProblemLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnalysisResultVM) GetLevelOk() (*ProblemLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnalysisResultVM) SetLevel(v ProblemLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnalysisResultVM) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetDetail

`func (o *AnalysisResultVM) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AnalysisResultVM) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AnalysisResultVM) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AnalysisResultVM) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *AnalysisResultVM) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *AnalysisResultVM) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetId

`func (o *AnalysisResultVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalysisResultVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalysisResultVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalysisResultVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AnalysisResultVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AnalysisResultVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSubscriptionId

`func (o *AnalysisResultVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AnalysisResultVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AnalysisResultVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AnalysisResultVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *AnalysisResultVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *AnalysisResultVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetFileId

`func (o *AnalysisResultVM) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *AnalysisResultVM) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *AnalysisResultVM) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *AnalysisResultVM) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *AnalysisResultVM) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *AnalysisResultVM) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetCollectionName

`func (o *AnalysisResultVM) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *AnalysisResultVM) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *AnalysisResultVM) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *AnalysisResultVM) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### SetCollectionNameNil

`func (o *AnalysisResultVM) SetCollectionNameNil(b bool)`

 SetCollectionNameNil sets the value for CollectionName to be an explicit nil

### UnsetCollectionName
`func (o *AnalysisResultVM) UnsetCollectionName()`

UnsetCollectionName ensures that no value is present for CollectionName, not even an explicit nil
### GetType

`func (o *AnalysisResultVM) GetType() ProblemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalysisResultVM) GetTypeOk() (*ProblemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalysisResultVM) SetType(v ProblemType)`

SetType sets Type field to given value.

### HasType

`func (o *AnalysisResultVM) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSignature

`func (o *AnalysisResultVM) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *AnalysisResultVM) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *AnalysisResultVM) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *AnalysisResultVM) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *AnalysisResultVM) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *AnalysisResultVM) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetNewSize

`func (o *AnalysisResultVM) GetNewSize() int64`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *AnalysisResultVM) GetNewSizeOk() (*int64, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *AnalysisResultVM) SetNewSize(v int64)`

SetNewSize sets NewSize field to given value.

### HasNewSize

`func (o *AnalysisResultVM) HasNewSize() bool`

HasNewSize returns a boolean if a field has been set.

### GetT

`func (o *AnalysisResultVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *AnalysisResultVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *AnalysisResultVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


