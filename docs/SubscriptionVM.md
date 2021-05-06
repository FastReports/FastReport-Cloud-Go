# SubscriptionVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**Current** | Pointer to [**SubscriptionPeriodVM**](SubscriptionPeriodVM.md) |  | [optional] 
**Old** | Pointer to [**[]SubscriptionPeriodVM**](SubscriptionPeriodVM.md) |  | [optional] 
**TemplatesFolder** | Pointer to [**SubscriptionFolder**](SubscriptionFolder.md) |  | [optional] 
**ReportsFolder** | Pointer to [**SubscriptionFolder**](SubscriptionFolder.md) |  | [optional] 
**ExportsFolder** | Pointer to [**SubscriptionFolder**](SubscriptionFolder.md) |  | [optional] 

## Methods

### NewSubscriptionVM

`func NewSubscriptionVM() *SubscriptionVM`

NewSubscriptionVM instantiates a new SubscriptionVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionVMWithDefaults

`func NewSubscriptionVMWithDefaults() *SubscriptionVM`

NewSubscriptionVMWithDefaults instantiates a new SubscriptionVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionVM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SubscriptionVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubscriptionVM) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocale

`func (o *SubscriptionVM) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *SubscriptionVM) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *SubscriptionVM) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *SubscriptionVM) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetCurrent

`func (o *SubscriptionVM) GetCurrent() SubscriptionPeriodVM`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *SubscriptionVM) GetCurrentOk() (*SubscriptionPeriodVM, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *SubscriptionVM) SetCurrent(v SubscriptionPeriodVM)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *SubscriptionVM) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetOld

`func (o *SubscriptionVM) GetOld() []SubscriptionPeriodVM`

GetOld returns the Old field if non-nil, zero value otherwise.

### GetOldOk

`func (o *SubscriptionVM) GetOldOk() (*[]SubscriptionPeriodVM, bool)`

GetOldOk returns a tuple with the Old field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOld

`func (o *SubscriptionVM) SetOld(v []SubscriptionPeriodVM)`

SetOld sets Old field to given value.

### HasOld

`func (o *SubscriptionVM) HasOld() bool`

HasOld returns a boolean if a field has been set.

### GetTemplatesFolder

`func (o *SubscriptionVM) GetTemplatesFolder() SubscriptionFolder`

GetTemplatesFolder returns the TemplatesFolder field if non-nil, zero value otherwise.

### GetTemplatesFolderOk

`func (o *SubscriptionVM) GetTemplatesFolderOk() (*SubscriptionFolder, bool)`

GetTemplatesFolderOk returns a tuple with the TemplatesFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesFolder

`func (o *SubscriptionVM) SetTemplatesFolder(v SubscriptionFolder)`

SetTemplatesFolder sets TemplatesFolder field to given value.

### HasTemplatesFolder

`func (o *SubscriptionVM) HasTemplatesFolder() bool`

HasTemplatesFolder returns a boolean if a field has been set.

### GetReportsFolder

`func (o *SubscriptionVM) GetReportsFolder() SubscriptionFolder`

GetReportsFolder returns the ReportsFolder field if non-nil, zero value otherwise.

### GetReportsFolderOk

`func (o *SubscriptionVM) GetReportsFolderOk() (*SubscriptionFolder, bool)`

GetReportsFolderOk returns a tuple with the ReportsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsFolder

`func (o *SubscriptionVM) SetReportsFolder(v SubscriptionFolder)`

SetReportsFolder sets ReportsFolder field to given value.

### HasReportsFolder

`func (o *SubscriptionVM) HasReportsFolder() bool`

HasReportsFolder returns a boolean if a field has been set.

### GetExportsFolder

`func (o *SubscriptionVM) GetExportsFolder() SubscriptionFolder`

GetExportsFolder returns the ExportsFolder field if non-nil, zero value otherwise.

### GetExportsFolderOk

`func (o *SubscriptionVM) GetExportsFolderOk() (*SubscriptionFolder, bool)`

GetExportsFolderOk returns a tuple with the ExportsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportsFolder

`func (o *SubscriptionVM) SetExportsFolder(v SubscriptionFolder)`

SetExportsFolder sets ExportsFolder field to given value.

### HasExportsFolder

`func (o *SubscriptionVM) HasExportsFolder() bool`

HasExportsFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


