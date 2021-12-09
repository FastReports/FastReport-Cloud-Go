# UpdateUserSettingsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileVisibility** | Pointer to [**ProfileVisibility**](ProfileVisibility.md) |  | [optional] 
**DefaultSubscription** | Pointer to **NullableString** |  | [optional] 
**ShowHiddenFilesAndFolders** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewUpdateUserSettingsVM

`func NewUpdateUserSettingsVM() *UpdateUserSettingsVM`

NewUpdateUserSettingsVM instantiates a new UpdateUserSettingsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserSettingsVMWithDefaults

`func NewUpdateUserSettingsVMWithDefaults() *UpdateUserSettingsVM`

NewUpdateUserSettingsVMWithDefaults instantiates a new UpdateUserSettingsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileVisibility

`func (o *UpdateUserSettingsVM) GetProfileVisibility() ProfileVisibility`

GetProfileVisibility returns the ProfileVisibility field if non-nil, zero value otherwise.

### GetProfileVisibilityOk

`func (o *UpdateUserSettingsVM) GetProfileVisibilityOk() (*ProfileVisibility, bool)`

GetProfileVisibilityOk returns a tuple with the ProfileVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileVisibility

`func (o *UpdateUserSettingsVM) SetProfileVisibility(v ProfileVisibility)`

SetProfileVisibility sets ProfileVisibility field to given value.

### HasProfileVisibility

`func (o *UpdateUserSettingsVM) HasProfileVisibility() bool`

HasProfileVisibility returns a boolean if a field has been set.

### GetDefaultSubscription

`func (o *UpdateUserSettingsVM) GetDefaultSubscription() string`

GetDefaultSubscription returns the DefaultSubscription field if non-nil, zero value otherwise.

### GetDefaultSubscriptionOk

`func (o *UpdateUserSettingsVM) GetDefaultSubscriptionOk() (*string, bool)`

GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubscription

`func (o *UpdateUserSettingsVM) SetDefaultSubscription(v string)`

SetDefaultSubscription sets DefaultSubscription field to given value.

### HasDefaultSubscription

`func (o *UpdateUserSettingsVM) HasDefaultSubscription() bool`

HasDefaultSubscription returns a boolean if a field has been set.

### SetDefaultSubscriptionNil

`func (o *UpdateUserSettingsVM) SetDefaultSubscriptionNil(b bool)`

 SetDefaultSubscriptionNil sets the value for DefaultSubscription to be an explicit nil

### UnsetDefaultSubscription
`func (o *UpdateUserSettingsVM) UnsetDefaultSubscription()`

UnsetDefaultSubscription ensures that no value is present for DefaultSubscription, not even an explicit nil
### GetShowHiddenFilesAndFolders

`func (o *UpdateUserSettingsVM) GetShowHiddenFilesAndFolders() bool`

GetShowHiddenFilesAndFolders returns the ShowHiddenFilesAndFolders field if non-nil, zero value otherwise.

### GetShowHiddenFilesAndFoldersOk

`func (o *UpdateUserSettingsVM) GetShowHiddenFilesAndFoldersOk() (*bool, bool)`

GetShowHiddenFilesAndFoldersOk returns a tuple with the ShowHiddenFilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHiddenFilesAndFolders

`func (o *UpdateUserSettingsVM) SetShowHiddenFilesAndFolders(v bool)`

SetShowHiddenFilesAndFolders sets ShowHiddenFilesAndFolders field to given value.

### HasShowHiddenFilesAndFolders

`func (o *UpdateUserSettingsVM) HasShowHiddenFilesAndFolders() bool`

HasShowHiddenFilesAndFolders returns a boolean if a field has been set.

### SetShowHiddenFilesAndFoldersNil

`func (o *UpdateUserSettingsVM) SetShowHiddenFilesAndFoldersNil(b bool)`

 SetShowHiddenFilesAndFoldersNil sets the value for ShowHiddenFilesAndFolders to be an explicit nil

### UnsetShowHiddenFilesAndFolders
`func (o *UpdateUserSettingsVM) UnsetShowHiddenFilesAndFolders()`

UnsetShowHiddenFilesAndFolders ensures that no value is present for ShowHiddenFilesAndFolders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


