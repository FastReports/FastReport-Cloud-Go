# UpdateSubscriptionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DefaultPermissions** | Pointer to [**DefaultPermissionsVM**](DefaultPermissionsVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateSubscriptionVM

`func NewUpdateSubscriptionVM(t string, ) *UpdateSubscriptionVM`

NewUpdateSubscriptionVM instantiates a new UpdateSubscriptionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionVMWithDefaults

`func NewUpdateSubscriptionVMWithDefaults() *UpdateSubscriptionVM`

NewUpdateSubscriptionVMWithDefaults instantiates a new UpdateSubscriptionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSubscriptionVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSubscriptionVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSubscriptionVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSubscriptionVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateSubscriptionVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateSubscriptionVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLocale

`func (o *UpdateSubscriptionVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UpdateSubscriptionVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UpdateSubscriptionVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UpdateSubscriptionVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *UpdateSubscriptionVM) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *UpdateSubscriptionVM) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetTags

`func (o *UpdateSubscriptionVM) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateSubscriptionVM) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateSubscriptionVM) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateSubscriptionVM) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateSubscriptionVM) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateSubscriptionVM) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDefaultPermissions

`func (o *UpdateSubscriptionVM) GetDefaultPermissions() DefaultPermissionsVM`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *UpdateSubscriptionVM) GetDefaultPermissionsOk() (*DefaultPermissionsVM, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *UpdateSubscriptionVM) SetDefaultPermissions(v DefaultPermissionsVM)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *UpdateSubscriptionVM) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetT

`func (o *UpdateSubscriptionVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateSubscriptionVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateSubscriptionVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


