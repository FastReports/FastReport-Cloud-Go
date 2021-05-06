# SubscriptionPlanVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**TimePeriodType** | Pointer to **string** |  | [optional] 
**TimePeriod** | Pointer to **int32** |  | [optional] 
**ReadonlyTimeLimitType** | Pointer to **string** |  | [optional] 
**ReadonlyTimeLimit** | Pointer to **int32** |  | [optional] 
**TemplatesSpaceLimit** | Pointer to **int64** |  | [optional] 
**ReportsSpaceLimit** | Pointer to **int64** |  | [optional] 
**ExportsSpaceLimit** | Pointer to **int64** |  | [optional] 
**FileUploadSizeLimit** | Pointer to **int64** |  | [optional] 
**DataSourceLimit** | Pointer to **int32** |  | [optional] 
**MaxUsersCount** | Pointer to **int32** |  | [optional] 
**HasSpaceOverdraft** | Pointer to **bool** |  | [optional] 
**GroupLimit** | Pointer to **int32** |  | [optional] 
**OnlineDesigner** | Pointer to **bool** |  | [optional] 
**IsDemo** | Pointer to **bool** |  | [optional] 
**UrlToBuy** | Pointer to **string** |  | [optional] 
**UnlimitedPage** | Pointer to **bool** |  | [optional] 
**PageLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewSubscriptionPlanVM

`func NewSubscriptionPlanVM() *SubscriptionPlanVM`

NewSubscriptionPlanVM instantiates a new SubscriptionPlanVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPlanVMWithDefaults

`func NewSubscriptionPlanVMWithDefaults() *SubscriptionPlanVM`

NewSubscriptionPlanVMWithDefaults instantiates a new SubscriptionPlanVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionPlanVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionPlanVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionPlanVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionPlanVM) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsActive

`func (o *SubscriptionPlanVM) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SubscriptionPlanVM) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SubscriptionPlanVM) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SubscriptionPlanVM) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDisplayName

`func (o *SubscriptionPlanVM) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SubscriptionPlanVM) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SubscriptionPlanVM) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SubscriptionPlanVM) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTimePeriodType

`func (o *SubscriptionPlanVM) GetTimePeriodType() string`

GetTimePeriodType returns the TimePeriodType field if non-nil, zero value otherwise.

### GetTimePeriodTypeOk

`func (o *SubscriptionPlanVM) GetTimePeriodTypeOk() (*string, bool)`

GetTimePeriodTypeOk returns a tuple with the TimePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodType

`func (o *SubscriptionPlanVM) SetTimePeriodType(v string)`

SetTimePeriodType sets TimePeriodType field to given value.

### HasTimePeriodType

`func (o *SubscriptionPlanVM) HasTimePeriodType() bool`

HasTimePeriodType returns a boolean if a field has been set.

### GetTimePeriod

`func (o *SubscriptionPlanVM) GetTimePeriod() int32`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *SubscriptionPlanVM) GetTimePeriodOk() (*int32, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *SubscriptionPlanVM) SetTimePeriod(v int32)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *SubscriptionPlanVM) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetReadonlyTimeLimitType

`func (o *SubscriptionPlanVM) GetReadonlyTimeLimitType() string`

GetReadonlyTimeLimitType returns the ReadonlyTimeLimitType field if non-nil, zero value otherwise.

### GetReadonlyTimeLimitTypeOk

`func (o *SubscriptionPlanVM) GetReadonlyTimeLimitTypeOk() (*string, bool)`

GetReadonlyTimeLimitTypeOk returns a tuple with the ReadonlyTimeLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyTimeLimitType

`func (o *SubscriptionPlanVM) SetReadonlyTimeLimitType(v string)`

SetReadonlyTimeLimitType sets ReadonlyTimeLimitType field to given value.

### HasReadonlyTimeLimitType

`func (o *SubscriptionPlanVM) HasReadonlyTimeLimitType() bool`

HasReadonlyTimeLimitType returns a boolean if a field has been set.

### GetReadonlyTimeLimit

`func (o *SubscriptionPlanVM) GetReadonlyTimeLimit() int32`

GetReadonlyTimeLimit returns the ReadonlyTimeLimit field if non-nil, zero value otherwise.

### GetReadonlyTimeLimitOk

`func (o *SubscriptionPlanVM) GetReadonlyTimeLimitOk() (*int32, bool)`

GetReadonlyTimeLimitOk returns a tuple with the ReadonlyTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonlyTimeLimit

`func (o *SubscriptionPlanVM) SetReadonlyTimeLimit(v int32)`

SetReadonlyTimeLimit sets ReadonlyTimeLimit field to given value.

### HasReadonlyTimeLimit

`func (o *SubscriptionPlanVM) HasReadonlyTimeLimit() bool`

HasReadonlyTimeLimit returns a boolean if a field has been set.

### GetTemplatesSpaceLimit

`func (o *SubscriptionPlanVM) GetTemplatesSpaceLimit() int64`

GetTemplatesSpaceLimit returns the TemplatesSpaceLimit field if non-nil, zero value otherwise.

### GetTemplatesSpaceLimitOk

`func (o *SubscriptionPlanVM) GetTemplatesSpaceLimitOk() (*int64, bool)`

GetTemplatesSpaceLimitOk returns a tuple with the TemplatesSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatesSpaceLimit

`func (o *SubscriptionPlanVM) SetTemplatesSpaceLimit(v int64)`

SetTemplatesSpaceLimit sets TemplatesSpaceLimit field to given value.

### HasTemplatesSpaceLimit

`func (o *SubscriptionPlanVM) HasTemplatesSpaceLimit() bool`

HasTemplatesSpaceLimit returns a boolean if a field has been set.

### GetReportsSpaceLimit

`func (o *SubscriptionPlanVM) GetReportsSpaceLimit() int64`

GetReportsSpaceLimit returns the ReportsSpaceLimit field if non-nil, zero value otherwise.

### GetReportsSpaceLimitOk

`func (o *SubscriptionPlanVM) GetReportsSpaceLimitOk() (*int64, bool)`

GetReportsSpaceLimitOk returns a tuple with the ReportsSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsSpaceLimit

`func (o *SubscriptionPlanVM) SetReportsSpaceLimit(v int64)`

SetReportsSpaceLimit sets ReportsSpaceLimit field to given value.

### HasReportsSpaceLimit

`func (o *SubscriptionPlanVM) HasReportsSpaceLimit() bool`

HasReportsSpaceLimit returns a boolean if a field has been set.

### GetExportsSpaceLimit

`func (o *SubscriptionPlanVM) GetExportsSpaceLimit() int64`

GetExportsSpaceLimit returns the ExportsSpaceLimit field if non-nil, zero value otherwise.

### GetExportsSpaceLimitOk

`func (o *SubscriptionPlanVM) GetExportsSpaceLimitOk() (*int64, bool)`

GetExportsSpaceLimitOk returns a tuple with the ExportsSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportsSpaceLimit

`func (o *SubscriptionPlanVM) SetExportsSpaceLimit(v int64)`

SetExportsSpaceLimit sets ExportsSpaceLimit field to given value.

### HasExportsSpaceLimit

`func (o *SubscriptionPlanVM) HasExportsSpaceLimit() bool`

HasExportsSpaceLimit returns a boolean if a field has been set.

### GetFileUploadSizeLimit

`func (o *SubscriptionPlanVM) GetFileUploadSizeLimit() int64`

GetFileUploadSizeLimit returns the FileUploadSizeLimit field if non-nil, zero value otherwise.

### GetFileUploadSizeLimitOk

`func (o *SubscriptionPlanVM) GetFileUploadSizeLimitOk() (*int64, bool)`

GetFileUploadSizeLimitOk returns a tuple with the FileUploadSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadSizeLimit

`func (o *SubscriptionPlanVM) SetFileUploadSizeLimit(v int64)`

SetFileUploadSizeLimit sets FileUploadSizeLimit field to given value.

### HasFileUploadSizeLimit

`func (o *SubscriptionPlanVM) HasFileUploadSizeLimit() bool`

HasFileUploadSizeLimit returns a boolean if a field has been set.

### GetDataSourceLimit

`func (o *SubscriptionPlanVM) GetDataSourceLimit() int32`

GetDataSourceLimit returns the DataSourceLimit field if non-nil, zero value otherwise.

### GetDataSourceLimitOk

`func (o *SubscriptionPlanVM) GetDataSourceLimitOk() (*int32, bool)`

GetDataSourceLimitOk returns a tuple with the DataSourceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceLimit

`func (o *SubscriptionPlanVM) SetDataSourceLimit(v int32)`

SetDataSourceLimit sets DataSourceLimit field to given value.

### HasDataSourceLimit

`func (o *SubscriptionPlanVM) HasDataSourceLimit() bool`

HasDataSourceLimit returns a boolean if a field has been set.

### GetMaxUsersCount

`func (o *SubscriptionPlanVM) GetMaxUsersCount() int32`

GetMaxUsersCount returns the MaxUsersCount field if non-nil, zero value otherwise.

### GetMaxUsersCountOk

`func (o *SubscriptionPlanVM) GetMaxUsersCountOk() (*int32, bool)`

GetMaxUsersCountOk returns a tuple with the MaxUsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsersCount

`func (o *SubscriptionPlanVM) SetMaxUsersCount(v int32)`

SetMaxUsersCount sets MaxUsersCount field to given value.

### HasMaxUsersCount

`func (o *SubscriptionPlanVM) HasMaxUsersCount() bool`

HasMaxUsersCount returns a boolean if a field has been set.

### GetHasSpaceOverdraft

`func (o *SubscriptionPlanVM) GetHasSpaceOverdraft() bool`

GetHasSpaceOverdraft returns the HasSpaceOverdraft field if non-nil, zero value otherwise.

### GetHasSpaceOverdraftOk

`func (o *SubscriptionPlanVM) GetHasSpaceOverdraftOk() (*bool, bool)`

GetHasSpaceOverdraftOk returns a tuple with the HasSpaceOverdraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpaceOverdraft

`func (o *SubscriptionPlanVM) SetHasSpaceOverdraft(v bool)`

SetHasSpaceOverdraft sets HasSpaceOverdraft field to given value.

### HasHasSpaceOverdraft

`func (o *SubscriptionPlanVM) HasHasSpaceOverdraft() bool`

HasHasSpaceOverdraft returns a boolean if a field has been set.

### GetGroupLimit

`func (o *SubscriptionPlanVM) GetGroupLimit() int32`

GetGroupLimit returns the GroupLimit field if non-nil, zero value otherwise.

### GetGroupLimitOk

`func (o *SubscriptionPlanVM) GetGroupLimitOk() (*int32, bool)`

GetGroupLimitOk returns a tuple with the GroupLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLimit

`func (o *SubscriptionPlanVM) SetGroupLimit(v int32)`

SetGroupLimit sets GroupLimit field to given value.

### HasGroupLimit

`func (o *SubscriptionPlanVM) HasGroupLimit() bool`

HasGroupLimit returns a boolean if a field has been set.

### GetOnlineDesigner

`func (o *SubscriptionPlanVM) GetOnlineDesigner() bool`

GetOnlineDesigner returns the OnlineDesigner field if non-nil, zero value otherwise.

### GetOnlineDesignerOk

`func (o *SubscriptionPlanVM) GetOnlineDesignerOk() (*bool, bool)`

GetOnlineDesignerOk returns a tuple with the OnlineDesigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineDesigner

`func (o *SubscriptionPlanVM) SetOnlineDesigner(v bool)`

SetOnlineDesigner sets OnlineDesigner field to given value.

### HasOnlineDesigner

`func (o *SubscriptionPlanVM) HasOnlineDesigner() bool`

HasOnlineDesigner returns a boolean if a field has been set.

### GetIsDemo

`func (o *SubscriptionPlanVM) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *SubscriptionPlanVM) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *SubscriptionPlanVM) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *SubscriptionPlanVM) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### GetUrlToBuy

`func (o *SubscriptionPlanVM) GetUrlToBuy() string`

GetUrlToBuy returns the UrlToBuy field if non-nil, zero value otherwise.

### GetUrlToBuyOk

`func (o *SubscriptionPlanVM) GetUrlToBuyOk() (*string, bool)`

GetUrlToBuyOk returns a tuple with the UrlToBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlToBuy

`func (o *SubscriptionPlanVM) SetUrlToBuy(v string)`

SetUrlToBuy sets UrlToBuy field to given value.

### HasUrlToBuy

`func (o *SubscriptionPlanVM) HasUrlToBuy() bool`

HasUrlToBuy returns a boolean if a field has been set.

### GetUnlimitedPage

`func (o *SubscriptionPlanVM) GetUnlimitedPage() bool`

GetUnlimitedPage returns the UnlimitedPage field if non-nil, zero value otherwise.

### GetUnlimitedPageOk

`func (o *SubscriptionPlanVM) GetUnlimitedPageOk() (*bool, bool)`

GetUnlimitedPageOk returns a tuple with the UnlimitedPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimitedPage

`func (o *SubscriptionPlanVM) SetUnlimitedPage(v bool)`

SetUnlimitedPage sets UnlimitedPage field to given value.

### HasUnlimitedPage

`func (o *SubscriptionPlanVM) HasUnlimitedPage() bool`

HasUnlimitedPage returns a boolean if a field has been set.

### GetPageLimit

`func (o *SubscriptionPlanVM) GetPageLimit() int32`

GetPageLimit returns the PageLimit field if non-nil, zero value otherwise.

### GetPageLimitOk

`func (o *SubscriptionPlanVM) GetPageLimitOk() (*int32, bool)`

GetPageLimitOk returns a tuple with the PageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageLimit

`func (o *SubscriptionPlanVM) SetPageLimit(v int32)`

SetPageLimit sets PageLimit field to given value.

### HasPageLimit

`func (o *SubscriptionPlanVM) HasPageLimit() bool`

HasPageLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


