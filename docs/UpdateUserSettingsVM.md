# UpdateUserSettingsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileVisibility** | Pointer to **int32** |  | [optional] 
**DefaultSubscription** | Pointer to **string** |  | [optional] 

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

`func (o *UpdateUserSettingsVM) GetProfileVisibility() int32`

GetProfileVisibility returns the ProfileVisibility field if non-nil, zero value otherwise.

### GetProfileVisibilityOk

`func (o *UpdateUserSettingsVM) GetProfileVisibilityOk() (*int32, bool)`

GetProfileVisibilityOk returns a tuple with the ProfileVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileVisibility

`func (o *UpdateUserSettingsVM) SetProfileVisibility(v int32)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


