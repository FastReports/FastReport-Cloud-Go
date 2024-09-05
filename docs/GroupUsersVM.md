# GroupUsersVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]GroupUserVM**](GroupUserVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewGroupUsersVM

`func NewGroupUsersVM(t string, ) *GroupUsersVM`

NewGroupUsersVM instantiates a new GroupUsersVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUsersVMWithDefaults

`func NewGroupUsersVMWithDefaults() *GroupUsersVM`

NewGroupUsersVMWithDefaults instantiates a new GroupUsersVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GroupUsersVM) GetUsers() []GroupUserVM`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GroupUsersVM) GetUsersOk() (*[]GroupUserVM, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GroupUsersVM) SetUsers(v []GroupUserVM)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GroupUsersVM) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *GroupUsersVM) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *GroupUsersVM) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetCount

`func (o *GroupUsersVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupUsersVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupUsersVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *GroupUsersVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTake

`func (o *GroupUsersVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *GroupUsersVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *GroupUsersVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *GroupUsersVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetSkip

`func (o *GroupUsersVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *GroupUsersVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *GroupUsersVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *GroupUsersVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetT

`func (o *GroupUsersVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GroupUsersVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GroupUsersVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


