# FileShareVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **NullableTime** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Permission** | Pointer to [**FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFileShareVM

`func NewFileShareVM() *FileShareVM`

NewFileShareVM instantiates a new FileShareVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareVMWithDefaults

`func NewFileShareVMWithDefaults() *FileShareVM`

NewFileShareVMWithDefaults instantiates a new FileShareVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *FileShareVM) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *FileShareVM) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *FileShareVM) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *FileShareVM) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *FileShareVM) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *FileShareVM) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetName

`func (o *FileShareVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileShareVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileShareVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileShareVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FileShareVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FileShareVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermission

`func (o *FileShareVM) GetPermission() FilePermissionCRUDVM`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *FileShareVM) GetPermissionOk() (*FilePermissionCRUDVM, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *FileShareVM) SetPermission(v FilePermissionCRUDVM)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *FileShareVM) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetKey

`func (o *FileShareVM) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FileShareVM) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FileShareVM) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FileShareVM) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *FileShareVM) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *FileShareVM) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


