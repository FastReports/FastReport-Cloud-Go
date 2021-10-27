# AdminFolderCreateVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | **string** |  | 
**OwnerId** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAdminFolderCreateVM

`func NewAdminFolderCreateVM(parentId string, ownerId string, ) *AdminFolderCreateVM`

NewAdminFolderCreateVM instantiates a new AdminFolderCreateVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminFolderCreateVMWithDefaults

`func NewAdminFolderCreateVMWithDefaults() *AdminFolderCreateVM`

NewAdminFolderCreateVMWithDefaults instantiates a new AdminFolderCreateVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *AdminFolderCreateVM) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AdminFolderCreateVM) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AdminFolderCreateVM) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetOwnerId

`func (o *AdminFolderCreateVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AdminFolderCreateVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AdminFolderCreateVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.


### GetName

`func (o *AdminFolderCreateVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminFolderCreateVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminFolderCreateVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminFolderCreateVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminFolderCreateVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminFolderCreateVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTags

`func (o *AdminFolderCreateVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AdminFolderCreateVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AdminFolderCreateVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AdminFolderCreateVM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AdminFolderCreateVM) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AdminFolderCreateVM) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetIcon

`func (o *AdminFolderCreateVM) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AdminFolderCreateVM) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AdminFolderCreateVM) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AdminFolderCreateVM) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *AdminFolderCreateVM) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *AdminFolderCreateVM) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


