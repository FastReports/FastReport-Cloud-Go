# AdminSubscriptionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultPermissions** | Pointer to [**DefaultPermissions**](DefaultPermissions.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Current** | Pointer to [**SubscriptionPeriodVM**](SubscriptionPeriodVM.md) |  | [optional] 
**Old** | Pointer to [**[]SubscriptionPeriodVM**](SubscriptionPeriodVM.md) |  | [optional] 
**TemplatesFolder** | Pointer to [**SubscriptionFolder**](SubscriptionFolder.md) |  | [optional] 
**ReportsFolder** | Pointer to [**SubscriptionFolder**](SubscriptionFolder.md) |  | [optional] 
**ExportsFolder** | Pointer to [**SubscriptionFolder**](SubscriptionFolder.md) |  | [optional] 

## Methods

### NewAdminSubscriptionVM

`func NewAdminSubscriptionVM() *AdminSubscriptionVM`

NewAdminSubscriptionVM instantiates a new AdminSubscriptionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminSubscriptionVMWithDefaults

`func NewAdminSubscriptionVMWithDefaults() *AdminSubscriptionVM`

NewAdminSubscriptionVMWithDefaults instantiates a new AdminSubscriptionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultPermissions

`func (o *AdminSubscriptionVM) GetDefaultPermissions() DefaultPermissions`

GetDefaultPermissions returns the DefaultPermissions field if non-nil, zero value otherwise.

### GetDefaultPermissionsOk

`func (o *AdminSubscriptionVM) GetDefaultPermissionsOk() (*DefaultPermissions, bool)`

GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissions

`func (o *AdminSubscriptionVM) SetDefaultPermissions(v DefaultPermissions)`

SetDefaultPermissions sets DefaultPermissions field to given value.

### HasDefaultPermissions

`func (o *AdminSubscriptionVM) HasDefaultPermissions() bool`

HasDefaultPermissions returns a boolean if a field has been set.

### GetId

`func (o *AdminSubscriptionVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminSubscriptionVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminSubscriptionVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminSubscriptionVM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AdminSubscriptionVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminSubscriptionVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminSubscriptionVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminSubscriptionVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocale

`func (o *AdminSubscriptionVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *AdminSubscriptionVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *AdminSubscriptionVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *AdminSubscriptionVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetCurrent

`func (o *AdminSubscriptionVM) GetCurrent() SubscriptionPeriodVM`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AdminSubscriptionVM) GetCurrentOk() (*SubscriptionPeriodVM, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AdminSubscriptionVM) SetCurrent(v SubscriptionPeriodVM)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *AdminSubscriptionVM) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetOld

`func (o *AdminSubscriptionVM) GetOld() []SubscriptionPeriodVM`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *AdminSubscriptionVM) GetOldOk() (*[]SubscriptionPeriodVM, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *AdminSubscriptionVM) SetOld(v []SubscriptionPeriodVM)`

SetOld sets Old field to given value.

### HasOld

`func (o *AdminSubscriptionVM) HasOld() bool`

HasOld returns a boolean if a field has been set.

### GetTemplatesFolder

`func (o *AdminSubscriptionVM) GetTemplatesFolder() SubscriptionFolder`

GetTemplatesFolder returns the TemplatesFolder field if non-nil, zero value otherwise.

### GetTemplatesFolderOk

`func (o *AdminSubscriptionVM) GetTemplatesFolderOk() (*SubscriptionFolder, bool)`

GetTemplatesFolderOk returns a tuple with the TemplatesFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesFolder

`func (o *AdminSubscriptionVM) SetTemplatesFolder(v SubscriptionFolder)`

SetTemplatesFolder sets TemplatesFolder field to given value.

### HasTemplatesFolder

`func (o *AdminSubscriptionVM) HasTemplatesFolder() bool`

HasTemplatesFolder returns a boolean if a field has been set.

### GetReportsFolder

`func (o *AdminSubscriptionVM) GetReportsFolder() SubscriptionFolder`

GetReportsFolder returns the ReportsFolder field if non-nil, zero value otherwise.

### GetReportsFolderOk

`func (o *AdminSubscriptionVM) GetReportsFolderOk() (*SubscriptionFolder, bool)`

GetReportsFolderOk returns a tuple with the ReportsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsFolder

`func (o *AdminSubscriptionVM) SetReportsFolder(v SubscriptionFolder)`

SetReportsFolder sets ReportsFolder field to given value.

### HasReportsFolder

`func (o *AdminSubscriptionVM) HasReportsFolder() bool`

HasReportsFolder returns a boolean if a field has been set.

### GetExportsFolder

`func (o *AdminSubscriptionVM) GetExportsFolder() SubscriptionFolder`

GetExportsFolder returns the ExportsFolder field if non-nil, zero value otherwise.

### GetExportsFolderOk

`func (o *AdminSubscriptionVM) GetExportsFolderOk() (*SubscriptionFolder, bool)`

GetExportsFolderOk returns a tuple with the ExportsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportsFolder

`func (o *AdminSubscriptionVM) SetExportsFolder(v SubscriptionFolder)`

SetExportsFolder sets ExportsFolder field to given value.

### HasExportsFolder

`func (o *AdminSubscriptionVM) HasExportsFolder() bool`

HasExportsFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


