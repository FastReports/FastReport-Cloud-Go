# CreateFileShareVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Permission** | Pointer to [**FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 

## Methods

### NewCreateFileShareVM

`func NewCreateFileShareVM() *CreateFileShareVM`

NewCreateFileShareVM instantiates a new CreateFileShareVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFileShareVMWithDefaults

`func NewCreateFileShareVMWithDefaults() *CreateFileShareVM`

NewCreateFileShareVMWithDefaults instantiates a new CreateFileShareVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *CreateFileShareVM) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateFileShareVM) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CreateFileShareVM) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CreateFileShareVM) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *CreateFileShareVM) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *CreateFileShareVM) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *CreateFileShareVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFileShareVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFileShareVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateFileShareVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateFileShareVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateFileShareVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermission

`func (o *CreateFileShareVM) GetPermission() FilePermissionCRUDVM`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreateFileShareVM) GetPermissionOk() (*FilePermissionCRUDVM, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreateFileShareVM) SetPermission(v FilePermissionCRUDVM)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CreateFileShareVM) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


