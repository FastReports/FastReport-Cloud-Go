# UserProfileVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserProfileVM

`func NewUserProfileVM() *UserProfileVM`

NewUserProfileVM instantiates a new UserProfileVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileVMWithDefaults

`func NewUserProfileVMWithDefaults() *UserProfileVM`

NewUserProfileVMWithDefaults instantiates a new UserProfileVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserProfileVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserProfileVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserProfileVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserProfileVM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserProfileVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserProfileVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserProfileVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserProfileVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *UserProfileVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserProfileVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserProfileVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserProfileVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserProfileVM) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProfileVM) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProfileVM) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProfileVM) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *UserProfileVM) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *UserProfileVM) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *UserProfileVM) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *UserProfileVM) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


