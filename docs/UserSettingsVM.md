# UserSettingsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileVisibility** | Pointer to [**ProfileVisibility**](ProfileVisibility.md) |  | [optional] 
**DefaultSubscription** | Pointer to **NullableString** |  | [optional] 
**ShowHiddenFilesAndFolders** | Pointer to **bool** |  | [optional] 
**SlaAcceptedDateTime** | Pointer to **NullableTime** |  | [optional] 
**SubscribedNotifications** | Pointer to [**[]AuditType**](AuditType.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUserSettingsVM

`func NewUserSettingsVM(t string, ) *UserSettingsVM`

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
### GetShowHiddenFilesAndFolders

`func (o *UserSettingsVM) GetShowHiddenFilesAndFolders() bool`

GetShowHiddenFilesAndFolders returns the ShowHiddenFilesAndFolders field if non-nil, zero value otherwise.

### GetShowHiddenFilesAndFoldersOk

`func (o *UserSettingsVM) GetShowHiddenFilesAndFoldersOk() (*bool, bool)`

GetShowHiddenFilesAndFoldersOk returns a tuple with the ShowHiddenFilesAndFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHiddenFilesAndFolders

`func (o *UserSettingsVM) SetShowHiddenFilesAndFolders(v bool)`

SetShowHiddenFilesAndFolders sets ShowHiddenFilesAndFolders field to given value.

### HasShowHiddenFilesAndFolders

`func (o *UserSettingsVM) HasShowHiddenFilesAndFolders() bool`

HasShowHiddenFilesAndFolders returns a boolean if a field has been set.

### GetSlaAcceptedDateTime

`func (o *UserSettingsVM) GetSlaAcceptedDateTime() time.Time`

GetSlaAcceptedDateTime returns the SlaAcceptedDateTime field if non-nil, zero value otherwise.

### GetSlaAcceptedDateTimeOk

`func (o *UserSettingsVM) GetSlaAcceptedDateTimeOk() (*time.Time, bool)`

GetSlaAcceptedDateTimeOk returns a tuple with the SlaAcceptedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaAcceptedDateTime

`func (o *UserSettingsVM) SetSlaAcceptedDateTime(v time.Time)`

SetSlaAcceptedDateTime sets SlaAcceptedDateTime field to given value.

### HasSlaAcceptedDateTime

`func (o *UserSettingsVM) HasSlaAcceptedDateTime() bool`

HasSlaAcceptedDateTime returns a boolean if a field has been set.

### SetSlaAcceptedDateTimeNil

`func (o *UserSettingsVM) SetSlaAcceptedDateTimeNil(b bool)`

 SetSlaAcceptedDateTimeNil sets the value for SlaAcceptedDateTime to be an explicit nil

### UnsetSlaAcceptedDateTime
`func (o *UserSettingsVM) UnsetSlaAcceptedDateTime()`

UnsetSlaAcceptedDateTime ensures that no value is present for SlaAcceptedDateTime, not even an explicit nil
### GetSubscribedNotifications

`func (o *UserSettingsVM) GetSubscribedNotifications() []AuditType`

GetSubscribedNotifications returns the SubscribedNotifications field if non-nil, zero value otherwise.

### GetSubscribedNotificationsOk

`func (o *UserSettingsVM) GetSubscribedNotificationsOk() (*[]AuditType, bool)`

GetSubscribedNotificationsOk returns a tuple with the SubscribedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedNotifications

`func (o *UserSettingsVM) SetSubscribedNotifications(v []AuditType)`

SetSubscribedNotifications sets SubscribedNotifications field to given value.

### HasSubscribedNotifications

`func (o *UserSettingsVM) HasSubscribedNotifications() bool`

HasSubscribedNotifications returns a boolean if a field has been set.

### SetSubscribedNotificationsNil

`func (o *UserSettingsVM) SetSubscribedNotificationsNil(b bool)`

 SetSubscribedNotificationsNil sets the value for SubscribedNotifications to be an explicit nil

### UnsetSubscribedNotifications
`func (o *UserSettingsVM) UnsetSubscribedNotifications()`

UnsetSubscribedNotifications ensures that no value is present for SubscribedNotifications, not even an explicit nil
### GetT

`func (o *UserSettingsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UserSettingsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UserSettingsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


