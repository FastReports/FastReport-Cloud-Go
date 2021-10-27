# UserSettingsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileVisibility** | Pointer to [**ProfileVisibility**](ProfileVisibility.md) |  | [optional] 
**DefaultSubscription** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserSettingsVM

`func NewUserSettingsVM() *UserSettingsVM`

NewUserSettingsVM instantiates a new UserSettingsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsVMWithDefaults

`func NewUserSettingsVMWithDefaults() *UserSettingsVM`

NewUserSettingsVMWithDefaults instantiates a new UserSettingsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileVisibility

`func (o *UserSettingsVM) GetProfileVisibility() ProfileVisibility`

GetProfileVisibility returns the ProfileVisibility field if non-nil, zero value otherwise.

### GetProfileVisibilityOk

`func (o *UserSettingsVM) GetProfileVisibilityOk() (*ProfileVisibility, bool)`

GetProfileVisibilityOk returns a tuple with the ProfileVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileVisibility

`func (o *UserSettingsVM) SetProfileVisibility(v ProfileVisibility)`

SetProfileVisibility sets ProfileVisibility field to given value.

### HasProfileVisibility

`func (o *UserSettingsVM) HasProfileVisibility() bool`

HasProfileVisibility returns a boolean if a field has been set.

### GetDefaultSubscription

`func (o *UserSettingsVM) GetDefaultSubscription() string`

GetDefaultSubscription returns the DefaultSubscription field if non-nil, zero value otherwise.

### GetDefaultSubscriptionOk

`func (o *UserSettingsVM) GetDefaultSubscriptionOk() (*string, bool)`

GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubscription

`func (o *UserSettingsVM) SetDefaultSubscription(v string)`

SetDefaultSubscription sets DefaultSubscription field to given value.

### HasDefaultSubscription

`func (o *UserSettingsVM) HasDefaultSubscription() bool`

HasDefaultSubscription returns a boolean if a field has been set.

### SetDefaultSubscriptionNil

`func (o *UserSettingsVM) SetDefaultSubscriptionNil(b bool)`

 SetDefaultSubscriptionNil sets the value for DefaultSubscription to be an explicit nil

### UnsetDefaultSubscription
`func (o *UserSettingsVM) UnsetDefaultSubscription()`

UnsetDefaultSubscription ensures that no value is present for DefaultSubscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


