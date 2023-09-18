# ContactGroupsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]ContactGroupVM**](ContactGroupVM.md) |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewContactGroupsVM

`func NewContactGroupsVM() *ContactGroupsVM`

NewContactGroupsVM instantiates a new ContactGroupsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactGroupsVMWithDefaults

`func NewContactGroupsVMWithDefaults() *ContactGroupsVM`

NewContactGroupsVMWithDefaults instantiates a new ContactGroupsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ContactGroupsVM) GetGroups() []ContactGroupVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ContactGroupsVM) GetGroupsOk() (*[]ContactGroupVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ContactGroupsVM) SetGroups(v []ContactGroupVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ContactGroupsVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *ContactGroupsVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ContactGroupsVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetSkip

`func (o *ContactGroupsVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ContactGroupsVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ContactGroupsVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ContactGroupsVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *ContactGroupsVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *ContactGroupsVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *ContactGroupsVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *ContactGroupsVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetCount

`func (o *ContactGroupsVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ContactGroupsVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ContactGroupsVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ContactGroupsVM) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


