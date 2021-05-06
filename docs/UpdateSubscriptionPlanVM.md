# UpdateSubscriptionPlanVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsActive** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**TimePeriodType** | Pointer to **string** |  | [optional] 
**TimePeriod** | Pointer to **int32** |  | [optional] 
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

### NewUpdateSubscriptionPlanVM

`func NewUpdateSubscriptionPlanVM() *UpdateSubscriptionPlanVM`

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

### GetTimePeriodType

`func (o *UpdateSubscriptionPlanVM) GetTimePeriodType() string`

GetTimePeriodType returns the TimePeriodType field if non-nil, zero value otherwise.

### GetTimePeriodTypeOk

`func (o *UpdateSubscriptionPlanVM) GetTimePeriodTypeOk() (*string, bool)`

GetTimePeriodTypeOk returns a tuple with the TimePeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodType

`func (o *UpdateSubscriptionPlanVM) SetTimePeriodType(v string)`

SetTimePeriodType sets TimePeriodType field to given value.

### HasTimePeriodType

`func (o *UpdateSubscriptionPlanVM) HasTimePeriodType() bool`

HasTimePeriodType returns a boolean if a field has been set.

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

### GetHasSpaceOverdraft

`func (o *UpdateSubscriptionPlanVM) GetHasSpaceOverdraft() bool`

GetHasSpaceOverdraft returns the HasSpaceOverdraft field if non-nil, zero value otherwise.

### GetHasSpaceOverdraftOk

`func (o *UpdateSubscriptionPlanVM) GetHasSpaceOverdraftOk() (*bool, bool)`

GetHasSpaceOverdraftOk returns a tuple with the HasSpaceOverdraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSpaceOverdraft

`func (o *UpdateSubscriptionPlanVM) SetHasSpaceOverdraft(v bool)`

SetHasSpaceOverdraft sets HasSpaceOverdraft field to given value.

### HasHasSpaceOverdraft

`func (o *UpdateSubscriptionPlanVM) HasHasSpaceOverdraft() bool`

HasHasSpaceOverdraft returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


