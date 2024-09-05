# GroupUserVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewGroupUserVM

`func NewGroupUserVM(t string, ) *GroupUserVM`

NewGroupUserVM instantiates a new GroupUserVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserVMWithDefaults

`func NewGroupUserVMWithDefaults() *GroupUserVM`

NewGroupUserVMWithDefaults instantiates a new GroupUserVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GroupUserVM) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupUserVM) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupUserVM) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GroupUserVM) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GroupUserVM) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GroupUserVM) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetT

`func (o *GroupUserVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *GroupUserVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *GroupUserVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


