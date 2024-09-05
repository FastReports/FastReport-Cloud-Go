# GroupsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]GroupVM**](GroupVM.md) |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewGroupsVM

`func NewGroupsVM(t string, ) *GroupsVM`

NewGroupsVM instantiates a new GroupsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsVMWithDefaults

`func NewGroupsVMWithDefaults() *GroupsVM`

NewGroupsVMWithDefaults instantiates a new GroupsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GroupsVM) GetGroups() []GroupVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupsVM) GetGroupsOk() (*[]GroupVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupsVM) SetGroups(v []GroupVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GroupsVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *GroupsVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *GroupsVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetCount

`func (o *GroupsVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupsVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupsVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *GroupsVM) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSkip

`func (o *GroupsVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *GroupsVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *GroupsVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *GroupsVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *GroupsVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *GroupsVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *GroupsVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *GroupsVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetT

`func (o *GroupsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GroupsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GroupsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


