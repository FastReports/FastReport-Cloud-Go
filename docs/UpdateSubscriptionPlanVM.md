# UpdateSubscriptionPlanVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**TimePeriodType** | Pointer to [**NullableTimePeriodType**](TimePeriodType.md) |  | [optional] 
**TimePeriod** | Pointer to **NullableInt32** |  | [optional] 
**TemplatesSpaceLimit** | Pointer to **NullableInt64** |  | [optional] 
**ReportsSpaceLimit** | Pointer to **NullableInt64** |  | [optional] 
**ExportsSpaceLimit** | Pointer to **NullableInt64** |  | [optional] 
**FileUploadSizeLimit** | Pointer to **NullableInt64** |  | [optional] 
**DataSourceLimit** | Pointer to **NullableInt32** |  | [optional] 
**MaxUsersCount** | Pointer to **NullableInt32** |  | [optional] 
**GroupLimit** | Pointer to **NullableInt32** |  | [optional] 
**OnlineDesigner** | Pointer to **NullableBool** |  | [optional] 
**IsDemo** | Pointer to **NullableBool** |  | [optional] 
**UrlToBuy** | Pointer to **NullableString** |  | [optional] 
**UnlimitedPage** | Pointer to **NullableBool** |  | [optional] 
**PageLimit** | Pointer to **NullableInt32** |  | [optional] 
**ReadonlyTimeLimitType** | Pointer to [**NullableTimePeriodType**](TimePeriodType.md) |  | [optional] 
**ReadonlyTimeLimit** | Pointer to **NullableInt32** |  | [optional] 
**Tasks** | Pointer to [**TaskSettingsVM**](TaskSettingsVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateSubscriptionPlanVM

`func NewUpdateSubscriptionPlanVM(t string, ) *UpdateSubscriptionPlanVM`

NewUpdateSubscriptionPlanVM instantiates a new UpdateSubscriptionPlanVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubscriptionPlanVMWithDefaults

`func NewUpdateSubscriptionPlanVMWithDefaults() *UpdateSubscriptionPlanVM`

NewUpdateSubscriptionPlanVMWithDefaults instantiates a new UpdateSubscriptionPlanVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsActive

`func (o *UpdateSubscriptionPlanVM) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateSubscriptionPlanVM) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateSubscriptionPlanVM) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateSubscriptionPlanVM) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UpdateSubscriptionPlanVM) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UpdateSubscriptionPlanVM) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetDisplayName

`func (o *UpdateSubscriptionPlanVM) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateSubscriptionPlanVM) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateSubscriptionPlanVM) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateSubscriptionPlanVM) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UpdateSubscriptionPlanVM) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UpdateSubscriptionPlanVM) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTimePeriodType

`func (o *UpdateSubscriptionPlanVM) GetTimePeriodType() TimePeriodType`

GetTimePeriodType returns the TimePeriodType field if non-nil, zero value otherwise.

### GetTimePeriodTypeOk

`func (o *UpdateSubscriptionPlanVM) GetTimePeriodTypeOk() (*TimePeriodType, bool)`

GetTimePeriodTypeOk returns a tuple with the TimePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodType

`func (o *UpdateSubscriptionPlanVM) SetTimePeriodType(v TimePeriodType)`

SetTimePeriodType sets TimePeriodType field to given value.

### HasTimePeriodType

`func (o *UpdateSubscriptionPlanVM) HasTimePeriodType() bool`

HasTimePeriodType returns a boolean if a field has been set.

### SetTimePeriodTypeNil

`func (o *UpdateSubscriptionPlanVM) SetTimePeriodTypeNil(b bool)`

 SetTimePeriodTypeNil sets the value for TimePeriodType to be an explicit nil

### UnsetTimePeriodType
`func (o *UpdateSubscriptionPlanVM) UnsetTimePeriodType()`

UnsetTimePeriodType ensures that no value is present for TimePeriodType, not even an explicit nil
### GetTimePeriod

`func (o *UpdateSubscriptionPlanVM) GetTimePeriod() int32`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *UpdateSubscriptionPlanVM) GetTimePeriodOk() (*int32, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *UpdateSubscriptionPlanVM) SetTimePeriod(v int32)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *UpdateSubscriptionPlanVM) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### SetTimePeriodNil

`func (o *UpdateSubscriptionPlanVM) SetTimePeriodNil(b bool)`

 SetTimePeriodNil sets the value for TimePeriod to be an explicit nil

### UnsetTimePeriod
`func (o *UpdateSubscriptionPlanVM) UnsetTimePeriod()`

UnsetTimePeriod ensures that no value is present for TimePeriod, not even an explicit nil
### GetTemplatesSpaceLimit

`func (o *UpdateSubscriptionPlanVM) GetTemplatesSpaceLimit() int64`

GetTemplatesSpaceLimit returns the TemplatesSpaceLimit field if non-nil, zero value otherwise.

### GetTemplatesSpaceLimitOk

`func (o *UpdateSubscriptionPlanVM) GetTemplatesSpaceLimitOk() (*int64, bool)`

GetTemplatesSpaceLimitOk returns a tuple with the TemplatesSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesSpaceLimit

`func (o *UpdateSubscriptionPlanVM) SetTemplatesSpaceLimit(v int64)`

SetTemplatesSpaceLimit sets TemplatesSpaceLimit field to given value.

### HasTemplatesSpaceLimit

`func (o *UpdateSubscriptionPlanVM) HasTemplatesSpaceLimit() bool`

HasTemplatesSpaceLimit returns a boolean if a field has been set.

### SetTemplatesSpaceLimitNil

`func (o *UpdateSubscriptionPlanVM) SetTemplatesSpaceLimitNil(b bool)`

 SetTemplatesSpaceLimitNil sets the value for TemplatesSpaceLimit to be an explicit nil

### UnsetTemplatesSpaceLimit
`func (o *UpdateSubscriptionPlanVM) UnsetTemplatesSpaceLimit()`

UnsetTemplatesSpaceLimit ensures that no value is present for TemplatesSpaceLimit, not even an explicit nil
### GetReportsSpaceLimit

`func (o *UpdateSubscriptionPlanVM) GetReportsSpaceLimit() int64`

GetReportsSpaceLimit returns the ReportsSpaceLimit field if non-nil, zero value otherwise.

### GetReportsSpaceLimitOk

`func (o *UpdateSubscriptionPlanVM) GetReportsSpaceLimitOk() (*int64, bool)`

GetReportsSpaceLimitOk returns a tuple with the ReportsSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsSpaceLimit

`func (o *UpdateSubscriptionPlanVM) SetReportsSpaceLimit(v int64)`

SetReportsSpaceLimit sets ReportsSpaceLimit field to given value.

### HasReportsSpaceLimit

`func (o *UpdateSubscriptionPlanVM) HasReportsSpaceLimit() bool`

HasReportsSpaceLimit returns a boolean if a field has been set.

### SetReportsSpaceLimitNil

`func (o *UpdateSubscriptionPlanVM) SetReportsSpaceLimitNil(b bool)`

 SetReportsSpaceLimitNil sets the value for ReportsSpaceLimit to be an explicit nil

### UnsetReportsSpaceLimit
`func (o *UpdateSubscriptionPlanVM) UnsetReportsSpaceLimit()`

UnsetReportsSpaceLimit ensures that no value is present for ReportsSpaceLimit, not even an explicit nil
### GetExportsSpaceLimit

`func (o *UpdateSubscriptionPlanVM) GetExportsSpaceLimit() int64`

GetExportsSpaceLimit returns the ExportsSpaceLimit field if non-nil, zero value otherwise.

### GetExportsSpaceLimitOk

`func (o *UpdateSubscriptionPlanVM) GetExportsSpaceLimitOk() (*int64, bool)`

GetExportsSpaceLimitOk returns a tuple with the ExportsSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportsSpaceLimit

`func (o *UpdateSubscriptionPlanVM) SetExportsSpaceLimit(v int64)`

SetExportsSpaceLimit sets ExportsSpaceLimit field to given value.

### HasExportsSpaceLimit

`func (o *UpdateSubscriptionPlanVM) HasExportsSpaceLimit() bool`

HasExportsSpaceLimit returns a boolean if a field has been set.

### SetExportsSpaceLimitNil

`func (o *UpdateSubscriptionPlanVM) SetExportsSpaceLimitNil(b bool)`

 SetExportsSpaceLimitNil sets the value for ExportsSpaceLimit to be an explicit nil

### UnsetExportsSpaceLimit
`func (o *UpdateSubscriptionPlanVM) UnsetExportsSpaceLimit()`

UnsetExportsSpaceLimit ensures that no value is present for ExportsSpaceLimit, not even an explicit nil
### GetFileUploadSizeLimit

`func (o *UpdateSubscriptionPlanVM) GetFileUploadSizeLimit() int64`

GetFileUploadSizeLimit returns the FileUploadSizeLimit field if non-nil, zero value otherwise.

### GetFileUploadSizeLimitOk

`func (o *UpdateSubscriptionPlanVM) GetFileUploadSizeLimitOk() (*int64, bool)`

GetFileUploadSizeLimitOk returns a tuple with the FileUploadSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSizeLimit

`func (o *UpdateSubscriptionPlanVM) SetFileUploadSizeLimit(v int64)`

SetFileUploadSizeLimit sets FileUploadSizeLimit field to given value.

### HasFileUploadSizeLimit

`func (o *UpdateSubscriptionPlanVM) HasFileUploadSizeLimit() bool`

HasFileUploadSizeLimit returns a boolean if a field has been set.

### SetFileUploadSizeLimitNil

`func (o *UpdateSubscriptionPlanVM) SetFileUploadSizeLimitNil(b bool)`

 SetFileUploadSizeLimitNil sets the value for FileUploadSizeLimit to be an explicit nil

### UnsetFileUploadSizeLimit
`func (o *UpdateSubscriptionPlanVM) UnsetFileUploadSizeLimit()`

UnsetFileUploadSizeLimit ensures that no value is present for FileUploadSizeLimit, not even an explicit nil
### GetDataSourceLimit

`func (o *UpdateSubscriptionPlanVM) GetDataSourceLimit() int32`

GetDataSourceLimit returns the DataSourceLimit field if non-nil, zero value otherwise.

### GetDataSourceLimitOk

`func (o *UpdateSubscriptionPlanVM) GetDataSourceLimitOk() (*int32, bool)`

GetDataSourceLimitOk returns a tuple with the DataSourceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceLimit

`func (o *UpdateSubscriptionPlanVM) SetDataSourceLimit(v int32)`

SetDataSourceLimit sets DataSourceLimit field to given value.

### HasDataSourceLimit

`func (o *UpdateSubscriptionPlanVM) HasDataSourceLimit() bool`

HasDataSourceLimit returns a boolean if a field has been set.

### SetDataSourceLimitNil

`func (o *UpdateSubscriptionPlanVM) SetDataSourceLimitNil(b bool)`

 SetDataSourceLimitNil sets the value for DataSourceLimit to be an explicit nil

### UnsetDataSourceLimit
`func (o *UpdateSubscriptionPlanVM) UnsetDataSourceLimit()`

UnsetDataSourceLimit ensures that no value is present for DataSourceLimit, not even an explicit nil
### GetMaxUsersCount

`func (o *UpdateSubscriptionPlanVM) GetMaxUsersCount() int32`

GetMaxUsersCount returns the MaxUsersCount field if non-nil, zero value otherwise.

### GetMaxUsersCountOk

`func (o *UpdateSubscriptionPlanVM) GetMaxUsersCountOk() (*int32, bool)`

GetMaxUsersCountOk returns a tuple with the MaxUsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsersCount

`func (o *UpdateSubscriptionPlanVM) SetMaxUsersCount(v int32)`

SetMaxUsersCount sets MaxUsersCount field to given value.

### HasMaxUsersCount

`func (o *UpdateSubscriptionPlanVM) HasMaxUsersCount() bool`

HasMaxUsersCount returns a boolean if a field has been set.

### SetMaxUsersCountNil

`func (o *UpdateSubscriptionPlanVM) SetMaxUsersCountNil(b bool)`

 SetMaxUsersCountNil sets the value for MaxUsersCount to be an explicit nil

### UnsetMaxUsersCount
`func (o *UpdateSubscriptionPlanVM) UnsetMaxUsersCount()`

UnsetMaxUsersCount ensures that no value is present for MaxUsersCount, not even an explicit nil
### GetGroupLimit

`func (o *UpdateSubscriptionPlanVM) GetGroupLimit() int32`

GetGroupLimit returns the GroupLimit field if non-nil, zero value otherwise.

### GetGroupLimitOk

`func (o *UpdateSubscriptionPlanVM) GetGroupLimitOk() (*int32, bool)`

GetGroupLimitOk returns a tuple with the GroupLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLimit

`func (o *UpdateSubscriptionPlanVM) SetGroupLimit(v int32)`

SetGroupLimit sets GroupLimit field to given value.

### HasGroupLimit

`func (o *UpdateSubscriptionPlanVM) HasGroupLimit() bool`

HasGroupLimit returns a boolean if a field has been set.

### SetGroupLimitNil

`func (o *UpdateSubscriptionPlanVM) SetGroupLimitNil(b bool)`

 SetGroupLimitNil sets the value for GroupLimit to be an explicit nil

### UnsetGroupLimit
`func (o *UpdateSubscriptionPlanVM) UnsetGroupLimit()`

UnsetGroupLimit ensures that no value is present for GroupLimit, not even an explicit nil
### GetOnlineDesigner

`func (o *UpdateSubscriptionPlanVM) GetOnlineDesigner() bool`

GetOnlineDesigner returns the OnlineDesigner field if non-nil, zero value otherwise.

### GetOnlineDesignerOk

`func (o *UpdateSubscriptionPlanVM) GetOnlineDesignerOk() (*bool, bool)`

GetOnlineDesignerOk returns a tuple with the OnlineDesigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDesigner

`func (o *UpdateSubscriptionPlanVM) SetOnlineDesigner(v bool)`

SetOnlineDesigner sets OnlineDesigner field to given value.

### HasOnlineDesigner

`func (o *UpdateSubscriptionPlanVM) HasOnlineDesigner() bool`

HasOnlineDesigner returns a boolean if a field has been set.

### SetOnlineDesignerNil

`func (o *UpdateSubscriptionPlanVM) SetOnlineDesignerNil(b bool)`

 SetOnlineDesignerNil sets the value for OnlineDesigner to be an explicit nil

### UnsetOnlineDesigner
`func (o *UpdateSubscriptionPlanVM) UnsetOnlineDesigner()`

UnsetOnlineDesigner ensures that no value is present for OnlineDesigner, not even an explicit nil
### GetIsDemo

`func (o *UpdateSubscriptionPlanVM) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *UpdateSubscriptionPlanVM) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *UpdateSubscriptionPlanVM) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *UpdateSubscriptionPlanVM) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### SetIsDemoNil

`func (o *UpdateSubscriptionPlanVM) SetIsDemoNil(b bool)`

 SetIsDemoNil sets the value for IsDemo to be an explicit nil

### UnsetIsDemo
`func (o *UpdateSubscriptionPlanVM) UnsetIsDemo()`

UnsetIsDemo ensures that no value is present for IsDemo, not even an explicit nil
### GetUrlToBuy

`func (o *UpdateSubscriptionPlanVM) GetUrlToBuy() string`

GetUrlToBuy returns the UrlToBuy field if non-nil, zero value otherwise.

### GetUrlToBuyOk

`func (o *UpdateSubscriptionPlanVM) GetUrlToBuyOk() (*string, bool)`

GetUrlToBuyOk returns a tuple with the UrlToBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlToBuy

`func (o *UpdateSubscriptionPlanVM) SetUrlToBuy(v string)`

SetUrlToBuy sets UrlToBuy field to given value.

### HasUrlToBuy

`func (o *UpdateSubscriptionPlanVM) HasUrlToBuy() bool`

HasUrlToBuy returns a boolean if a field has been set.

### SetUrlToBuyNil

`func (o *UpdateSubscriptionPlanVM) SetUrlToBuyNil(b bool)`

 SetUrlToBuyNil sets the value for UrlToBuy to be an explicit nil

### UnsetUrlToBuy
`func (o *UpdateSubscriptionPlanVM) UnsetUrlToBuy()`

UnsetUrlToBuy ensures that no value is present for UrlToBuy, not even an explicit nil
### GetUnlimitedPage

`func (o *UpdateSubscriptionPlanVM) GetUnlimitedPage() bool`

GetUnlimitedPage returns the UnlimitedPage field if non-nil, zero value otherwise.

### GetUnlimitedPageOk

`func (o *UpdateSubscriptionPlanVM) GetUnlimitedPageOk() (*bool, bool)`

GetUnlimitedPageOk returns a tuple with the UnlimitedPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimitedPage

`func (o *UpdateSubscriptionPlanVM) SetUnlimitedPage(v bool)`

SetUnlimitedPage sets UnlimitedPage field to given value.

### HasUnlimitedPage

`func (o *UpdateSubscriptionPlanVM) HasUnlimitedPage() bool`

HasUnlimitedPage returns a boolean if a field has been set.

### SetUnlimitedPageNil

`func (o *UpdateSubscriptionPlanVM) SetUnlimitedPageNil(b bool)`

 SetUnlimitedPageNil sets the value for UnlimitedPage to be an explicit nil

### UnsetUnlimitedPage
`func (o *UpdateSubscriptionPlanVM) UnsetUnlimitedPage()`

UnsetUnlimitedPage ensures that no value is present for UnlimitedPage, not even an explicit nil
### GetPageLimit

`func (o *UpdateSubscriptionPlanVM) GetPageLimit() int32`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *UpdateSubscriptionPlanVM) GetPageLimitOk() (*int32, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *UpdateSubscriptionPlanVM) SetPageLimit(v int32)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *UpdateSubscriptionPlanVM) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.

### SetPageLimitNil

`func (o *UpdateSubscriptionPlanVM) SetPageLimitNil(b bool)`

 SetPageLimitNil sets the value for PageLimit to be an explicit nil

### UnsetPageLimit
`func (o *UpdateSubscriptionPlanVM) UnsetPageLimit()`

UnsetPageLimit ensures that no value is present for PageLimit, not even an explicit nil
### GetReadonlyTimeLimitType

`func (o *UpdateSubscriptionPlanVM) GetReadonlyTimeLimitType() TimePeriodType`

GetReadonlyTimeLimitType returns the ReadonlyTimeLimitType field if non-nil, zero value otherwise.

### GetReadonlyTimeLimitTypeOk

`func (o *UpdateSubscriptionPlanVM) GetReadonlyTimeLimitTypeOk() (*TimePeriodType, bool)`

GetReadonlyTimeLimitTypeOk returns a tuple with the ReadonlyTimeLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyTimeLimitType

`func (o *UpdateSubscriptionPlanVM) SetReadonlyTimeLimitType(v TimePeriodType)`

SetReadonlyTimeLimitType sets ReadonlyTimeLimitType field to given value.

### HasReadonlyTimeLimitType

`func (o *UpdateSubscriptionPlanVM) HasReadonlyTimeLimitType() bool`

HasReadonlyTimeLimitType returns a boolean if a field has been set.

### SetReadonlyTimeLimitTypeNil

`func (o *UpdateSubscriptionPlanVM) SetReadonlyTimeLimitTypeNil(b bool)`

 SetReadonlyTimeLimitTypeNil sets the value for ReadonlyTimeLimitType to be an explicit nil

### UnsetReadonlyTimeLimitType
`func (o *UpdateSubscriptionPlanVM) UnsetReadonlyTimeLimitType()`

UnsetReadonlyTimeLimitType ensures that no value is present for ReadonlyTimeLimitType, not even an explicit nil
### GetReadonlyTimeLimit

`func (o *UpdateSubscriptionPlanVM) GetReadonlyTimeLimit() int32`

GetReadonlyTimeLimit returns the ReadonlyTimeLimit field if non-nil, zero value otherwise.

### GetReadonlyTimeLimitOk

`func (o *UpdateSubscriptionPlanVM) GetReadonlyTimeLimitOk() (*int32, bool)`

GetReadonlyTimeLimitOk returns a tuple with the ReadonlyTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyTimeLimit

`func (o *UpdateSubscriptionPlanVM) SetReadonlyTimeLimit(v int32)`

SetReadonlyTimeLimit sets ReadonlyTimeLimit field to given value.

### HasReadonlyTimeLimit

`func (o *UpdateSubscriptionPlanVM) HasReadonlyTimeLimit() bool`

HasReadonlyTimeLimit returns a boolean if a field has been set.

### SetReadonlyTimeLimitNil

`func (o *UpdateSubscriptionPlanVM) SetReadonlyTimeLimitNil(b bool)`

 SetReadonlyTimeLimitNil sets the value for ReadonlyTimeLimit to be an explicit nil

### UnsetReadonlyTimeLimit
`func (o *UpdateSubscriptionPlanVM) UnsetReadonlyTimeLimit()`

UnsetReadonlyTimeLimit ensures that no value is present for ReadonlyTimeLimit, not even an explicit nil
### GetTasks

`func (o *UpdateSubscriptionPlanVM) GetTasks() TaskSettingsVM`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *UpdateSubscriptionPlanVM) GetTasksOk() (*TaskSettingsVM, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *UpdateSubscriptionPlanVM) SetTasks(v TaskSettingsVM)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *UpdateSubscriptionPlanVM) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetT

`func (o *UpdateSubscriptionPlanVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateSubscriptionPlanVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateSubscriptionPlanVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


