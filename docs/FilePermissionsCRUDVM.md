# FilePermissionsCRUDVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 
**Groups** | Pointer to [**map[string]FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 
**Other** | Pointer to [**FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 
**Anon** | Pointer to [**FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFilePermissionsCRUDVM

`func NewFilePermissionsCRUDVM(t string, ) *FilePermissionsCRUDVM`

NewFilePermissionsCRUDVM instantiates a new FilePermissionsCRUDVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePermissionsCRUDVMWithDefaults

`func NewFilePermissionsCRUDVMWithDefaults() *FilePermissionsCRUDVM`

NewFilePermissionsCRUDVMWithDefaults instantiates a new FilePermissionsCRUDVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *FilePermissionsCRUDVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *FilePermissionsCRUDVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *FilePermissionsCRUDVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *FilePermissionsCRUDVM) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *FilePermissionsCRUDVM) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *FilePermissionsCRUDVM) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *FilePermissionsCRUDVM) GetOwner() FilePermissionCRUDVM`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *FilePermissionsCRUDVM) GetOwnerOk() (*FilePermissionCRUDVM, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *FilePermissionsCRUDVM) SetOwner(v FilePermissionCRUDVM)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *FilePermissionsCRUDVM) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *FilePermissionsCRUDVM) GetGroups() map[string]FilePermissionCRUDVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FilePermissionsCRUDVM) GetGroupsOk() (*map[string]FilePermissionCRUDVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FilePermissionsCRUDVM) SetGroups(v map[string]FilePermissionCRUDVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FilePermissionsCRUDVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *FilePermissionsCRUDVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *FilePermissionsCRUDVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *FilePermissionsCRUDVM) GetOther() FilePermissionCRUDVM`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *FilePermissionsCRUDVM) GetOtherOk() (*FilePermissionCRUDVM, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *FilePermissionsCRUDVM) SetOther(v FilePermissionCRUDVM)`

SetOther sets Other field to given value.

### HasOther

`func (o *FilePermissionsCRUDVM) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *FilePermissionsCRUDVM) GetAnon() FilePermissionCRUDVM`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *FilePermissionsCRUDVM) GetAnonOk() (*FilePermissionCRUDVM, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *FilePermissionsCRUDVM) SetAnon(v FilePermissionCRUDVM)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *FilePermissionsCRUDVM) HasAnon() bool`

HasAnon returns a boolean if a field has been set.

### GetT

`func (o *FilePermissionsCRUDVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FilePermissionsCRUDVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FilePermissionsCRUDVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


