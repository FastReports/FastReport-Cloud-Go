/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// SubscriptionPlanVM struct for SubscriptionPlanVM
type SubscriptionPlanVM struct {
	Id *string `json:"id,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	TimePeriodType *string `json:"timePeriodType,omitempty"`
	TimePeriod *int32 `json:"timePeriod,omitempty"`
	ReadonlyTimeLimitType *string `json:"readonlyTimeLimitType,omitempty"`
	ReadonlyTimeLimit *int32 `json:"readonlyTimeLimit,omitempty"`
	TemplatesSpaceLimit *int64 `json:"templatesSpaceLimit,omitempty"`
	ReportsSpaceLimit *int64 `json:"reportsSpaceLimit,omitempty"`
	ExportsSpaceLimit *int64 `json:"exportsSpaceLimit,omitempty"`
	FileUploadSizeLimit *int64 `json:"fileUploadSizeLimit,omitempty"`
	DataSourceLimit *int32 `json:"dataSourceLimit,omitempty"`
	MaxUsersCount *int32 `json:"maxUsersCount,omitempty"`
	HasSpaceOverdraft *bool `json:"hasSpaceOverdraft,omitempty"`
	GroupLimit *int32 `json:"groupLimit,omitempty"`
	OnlineDesigner *bool `json:"onlineDesigner,omitempty"`
	IsDemo *bool `json:"isDemo,omitempty"`
	UrlToBuy *string `json:"urlToBuy,omitempty"`
	UnlimitedPage *bool `json:"unlimitedPage,omitempty"`
	PageLimit *int32 `json:"pageLimit,omitempty"`
}

// NewSubscriptionPlanVM instantiates a new SubscriptionPlanVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPlanVM() *SubscriptionPlanVM {
	this := SubscriptionPlanVM{}
	return &this
}

// NewSubscriptionPlanVMWithDefaults instantiates a new SubscriptionPlanVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPlanVMWithDefaults() *SubscriptionPlanVM {
	this := SubscriptionPlanVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionPlanVM) SetId(v string) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *SubscriptionPlanVM) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SubscriptionPlanVM) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTimePeriodType returns the TimePeriodType field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetTimePeriodType() string {
	if o == nil || o.TimePeriodType == nil {
		var ret string
		return ret
	}
	return *o.TimePeriodType
}

// GetTimePeriodTypeOk returns a tuple with the TimePeriodType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetTimePeriodTypeOk() (*string, bool) {
	if o == nil || o.TimePeriodType == nil {
		return nil, false
	}
	return o.TimePeriodType, true
}

// HasTimePeriodType returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTimePeriodType() bool {
	if o != nil && o.TimePeriodType != nil {
		return true
	}

	return false
}

// SetTimePeriodType gets a reference to the given string and assigns it to the TimePeriodType field.
func (o *SubscriptionPlanVM) SetTimePeriodType(v string) {
	o.TimePeriodType = &v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetTimePeriod() int32 {
	if o == nil || o.TimePeriod == nil {
		var ret int32
		return ret
	}
	return *o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetTimePeriodOk() (*int32, bool) {
	if o == nil || o.TimePeriod == nil {
		return nil, false
	}
	return o.TimePeriod, true
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTimePeriod() bool {
	if o != nil && o.TimePeriod != nil {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given int32 and assigns it to the TimePeriod field.
func (o *SubscriptionPlanVM) SetTimePeriod(v int32) {
	o.TimePeriod = &v
}

// GetReadonlyTimeLimitType returns the ReadonlyTimeLimitType field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimitType() string {
	if o == nil || o.ReadonlyTimeLimitType == nil {
		var ret string
		return ret
	}
	return *o.ReadonlyTimeLimitType
}

// GetReadonlyTimeLimitTypeOk returns a tuple with the ReadonlyTimeLimitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimitTypeOk() (*string, bool) {
	if o == nil || o.ReadonlyTimeLimitType == nil {
		return nil, false
	}
	return o.ReadonlyTimeLimitType, true
}

// HasReadonlyTimeLimitType returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasReadonlyTimeLimitType() bool {
	if o != nil && o.ReadonlyTimeLimitType != nil {
		return true
	}

	return false
}

// SetReadonlyTimeLimitType gets a reference to the given string and assigns it to the ReadonlyTimeLimitType field.
func (o *SubscriptionPlanVM) SetReadonlyTimeLimitType(v string) {
	o.ReadonlyTimeLimitType = &v
}

// GetReadonlyTimeLimit returns the ReadonlyTimeLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimit() int32 {
	if o == nil || o.ReadonlyTimeLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReadonlyTimeLimit
}

// GetReadonlyTimeLimitOk returns a tuple with the ReadonlyTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetReadonlyTimeLimitOk() (*int32, bool) {
	if o == nil || o.ReadonlyTimeLimit == nil {
		return nil, false
	}
	return o.ReadonlyTimeLimit, true
}

// HasReadonlyTimeLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasReadonlyTimeLimit() bool {
	if o != nil && o.ReadonlyTimeLimit != nil {
		return true
	}

	return false
}

// SetReadonlyTimeLimit gets a reference to the given int32 and assigns it to the ReadonlyTimeLimit field.
func (o *SubscriptionPlanVM) SetReadonlyTimeLimit(v int32) {
	o.ReadonlyTimeLimit = &v
}

// GetTemplatesSpaceLimit returns the TemplatesSpaceLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetTemplatesSpaceLimit() int64 {
	if o == nil || o.TemplatesSpaceLimit == nil {
		var ret int64
		return ret
	}
	return *o.TemplatesSpaceLimit
}

// GetTemplatesSpaceLimitOk returns a tuple with the TemplatesSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetTemplatesSpaceLimitOk() (*int64, bool) {
	if o == nil || o.TemplatesSpaceLimit == nil {
		return nil, false
	}
	return o.TemplatesSpaceLimit, true
}

// HasTemplatesSpaceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasTemplatesSpaceLimit() bool {
	if o != nil && o.TemplatesSpaceLimit != nil {
		return true
	}

	return false
}

// SetTemplatesSpaceLimit gets a reference to the given int64 and assigns it to the TemplatesSpaceLimit field.
func (o *SubscriptionPlanVM) SetTemplatesSpaceLimit(v int64) {
	o.TemplatesSpaceLimit = &v
}

// GetReportsSpaceLimit returns the ReportsSpaceLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetReportsSpaceLimit() int64 {
	if o == nil || o.ReportsSpaceLimit == nil {
		var ret int64
		return ret
	}
	return *o.ReportsSpaceLimit
}

// GetReportsSpaceLimitOk returns a tuple with the ReportsSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetReportsSpaceLimitOk() (*int64, bool) {
	if o == nil || o.ReportsSpaceLimit == nil {
		return nil, false
	}
	return o.ReportsSpaceLimit, true
}

// HasReportsSpaceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasReportsSpaceLimit() bool {
	if o != nil && o.ReportsSpaceLimit != nil {
		return true
	}

	return false
}

// SetReportsSpaceLimit gets a reference to the given int64 and assigns it to the ReportsSpaceLimit field.
func (o *SubscriptionPlanVM) SetReportsSpaceLimit(v int64) {
	o.ReportsSpaceLimit = &v
}

// GetExportsSpaceLimit returns the ExportsSpaceLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetExportsSpaceLimit() int64 {
	if o == nil || o.ExportsSpaceLimit == nil {
		var ret int64
		return ret
	}
	return *o.ExportsSpaceLimit
}

// GetExportsSpaceLimitOk returns a tuple with the ExportsSpaceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetExportsSpaceLimitOk() (*int64, bool) {
	if o == nil || o.ExportsSpaceLimit == nil {
		return nil, false
	}
	return o.ExportsSpaceLimit, true
}

// HasExportsSpaceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasExportsSpaceLimit() bool {
	if o != nil && o.ExportsSpaceLimit != nil {
		return true
	}

	return false
}

// SetExportsSpaceLimit gets a reference to the given int64 and assigns it to the ExportsSpaceLimit field.
func (o *SubscriptionPlanVM) SetExportsSpaceLimit(v int64) {
	o.ExportsSpaceLimit = &v
}

// GetFileUploadSizeLimit returns the FileUploadSizeLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetFileUploadSizeLimit() int64 {
	if o == nil || o.FileUploadSizeLimit == nil {
		var ret int64
		return ret
	}
	return *o.FileUploadSizeLimit
}

// GetFileUploadSizeLimitOk returns a tuple with the FileUploadSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetFileUploadSizeLimitOk() (*int64, bool) {
	if o == nil || o.FileUploadSizeLimit == nil {
		return nil, false
	}
	return o.FileUploadSizeLimit, true
}

// HasFileUploadSizeLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasFileUploadSizeLimit() bool {
	if o != nil && o.FileUploadSizeLimit != nil {
		return true
	}

	return false
}

// SetFileUploadSizeLimit gets a reference to the given int64 and assigns it to the FileUploadSizeLimit field.
func (o *SubscriptionPlanVM) SetFileUploadSizeLimit(v int64) {
	o.FileUploadSizeLimit = &v
}

// GetDataSourceLimit returns the DataSourceLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetDataSourceLimit() int32 {
	if o == nil || o.DataSourceLimit == nil {
		var ret int32
		return ret
	}
	return *o.DataSourceLimit
}

// GetDataSourceLimitOk returns a tuple with the DataSourceLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetDataSourceLimitOk() (*int32, bool) {
	if o == nil || o.DataSourceLimit == nil {
		return nil, false
	}
	return o.DataSourceLimit, true
}

// HasDataSourceLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasDataSourceLimit() bool {
	if o != nil && o.DataSourceLimit != nil {
		return true
	}

	return false
}

// SetDataSourceLimit gets a reference to the given int32 and assigns it to the DataSourceLimit field.
func (o *SubscriptionPlanVM) SetDataSourceLimit(v int32) {
	o.DataSourceLimit = &v
}

// GetMaxUsersCount returns the MaxUsersCount field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetMaxUsersCount() int32 {
	if o == nil || o.MaxUsersCount == nil {
		var ret int32
		return ret
	}
	return *o.MaxUsersCount
}

// GetMaxUsersCountOk returns a tuple with the MaxUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetMaxUsersCountOk() (*int32, bool) {
	if o == nil || o.MaxUsersCount == nil {
		return nil, false
	}
	return o.MaxUsersCount, true
}

// HasMaxUsersCount returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasMaxUsersCount() bool {
	if o != nil && o.MaxUsersCount != nil {
		return true
	}

	return false
}

// SetMaxUsersCount gets a reference to the given int32 and assigns it to the MaxUsersCount field.
func (o *SubscriptionPlanVM) SetMaxUsersCount(v int32) {
	o.MaxUsersCount = &v
}

// GetHasSpaceOverdraft returns the HasSpaceOverdraft field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetHasSpaceOverdraft() bool {
	if o == nil || o.HasSpaceOverdraft == nil {
		var ret bool
		return ret
	}
	return *o.HasSpaceOverdraft
}

// GetHasSpaceOverdraftOk returns a tuple with the HasSpaceOverdraft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetHasSpaceOverdraftOk() (*bool, bool) {
	if o == nil || o.HasSpaceOverdraft == nil {
		return nil, false
	}
	return o.HasSpaceOverdraft, true
}

// HasHasSpaceOverdraft returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasHasSpaceOverdraft() bool {
	if o != nil && o.HasSpaceOverdraft != nil {
		return true
	}

	return false
}

// SetHasSpaceOverdraft gets a reference to the given bool and assigns it to the HasSpaceOverdraft field.
func (o *SubscriptionPlanVM) SetHasSpaceOverdraft(v bool) {
	o.HasSpaceOverdraft = &v
}

// GetGroupLimit returns the GroupLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetGroupLimit() int32 {
	if o == nil || o.GroupLimit == nil {
		var ret int32
		return ret
	}
	return *o.GroupLimit
}

// GetGroupLimitOk returns a tuple with the GroupLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetGroupLimitOk() (*int32, bool) {
	if o == nil || o.GroupLimit == nil {
		return nil, false
	}
	return o.GroupLimit, true
}

// HasGroupLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasGroupLimit() bool {
	if o != nil && o.GroupLimit != nil {
		return true
	}

	return false
}

// SetGroupLimit gets a reference to the given int32 and assigns it to the GroupLimit field.
func (o *SubscriptionPlanVM) SetGroupLimit(v int32) {
	o.GroupLimit = &v
}

// GetOnlineDesigner returns the OnlineDesigner field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetOnlineDesigner() bool {
	if o == nil || o.OnlineDesigner == nil {
		var ret bool
		return ret
	}
	return *o.OnlineDesigner
}

// GetOnlineDesignerOk returns a tuple with the OnlineDesigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetOnlineDesignerOk() (*bool, bool) {
	if o == nil || o.OnlineDesigner == nil {
		return nil, false
	}
	return o.OnlineDesigner, true
}

// HasOnlineDesigner returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasOnlineDesigner() bool {
	if o != nil && o.OnlineDesigner != nil {
		return true
	}

	return false
}

// SetOnlineDesigner gets a reference to the given bool and assigns it to the OnlineDesigner field.
func (o *SubscriptionPlanVM) SetOnlineDesigner(v bool) {
	o.OnlineDesigner = &v
}

// GetIsDemo returns the IsDemo field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetIsDemo() bool {
	if o == nil || o.IsDemo == nil {
		var ret bool
		return ret
	}
	return *o.IsDemo
}

// GetIsDemoOk returns a tuple with the IsDemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetIsDemoOk() (*bool, bool) {
	if o == nil || o.IsDemo == nil {
		return nil, false
	}
	return o.IsDemo, true
}

// HasIsDemo returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasIsDemo() bool {
	if o != nil && o.IsDemo != nil {
		return true
	}

	return false
}

// SetIsDemo gets a reference to the given bool and assigns it to the IsDemo field.
func (o *SubscriptionPlanVM) SetIsDemo(v bool) {
	o.IsDemo = &v
}

// GetUrlToBuy returns the UrlToBuy field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetUrlToBuy() string {
	if o == nil || o.UrlToBuy == nil {
		var ret string
		return ret
	}
	return *o.UrlToBuy
}

// GetUrlToBuyOk returns a tuple with the UrlToBuy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetUrlToBuyOk() (*string, bool) {
	if o == nil || o.UrlToBuy == nil {
		return nil, false
	}
	return o.UrlToBuy, true
}

// HasUrlToBuy returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasUrlToBuy() bool {
	if o != nil && o.UrlToBuy != nil {
		return true
	}

	return false
}

// SetUrlToBuy gets a reference to the given string and assigns it to the UrlToBuy field.
func (o *SubscriptionPlanVM) SetUrlToBuy(v string) {
	o.UrlToBuy = &v
}

// GetUnlimitedPage returns the UnlimitedPage field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetUnlimitedPage() bool {
	if o == nil || o.UnlimitedPage == nil {
		var ret bool
		return ret
	}
	return *o.UnlimitedPage
}

// GetUnlimitedPageOk returns a tuple with the UnlimitedPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetUnlimitedPageOk() (*bool, bool) {
	if o == nil || o.UnlimitedPage == nil {
		return nil, false
	}
	return o.UnlimitedPage, true
}

// HasUnlimitedPage returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasUnlimitedPage() bool {
	if o != nil && o.UnlimitedPage != nil {
		return true
	}

	return false
}

// SetUnlimitedPage gets a reference to the given bool and assigns it to the UnlimitedPage field.
func (o *SubscriptionPlanVM) SetUnlimitedPage(v bool) {
	o.UnlimitedPage = &v
}

// GetPageLimit returns the PageLimit field value if set, zero value otherwise.
func (o *SubscriptionPlanVM) GetPageLimit() int32 {
	if o == nil || o.PageLimit == nil {
		var ret int32
		return ret
	}
	return *o.PageLimit
}

// GetPageLimitOk returns a tuple with the PageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPlanVM) GetPageLimitOk() (*int32, bool) {
	if o == nil || o.PageLimit == nil {
		return nil, false
	}
	return o.PageLimit, true
}

// HasPageLimit returns a boolean if a field has been set.
func (o *SubscriptionPlanVM) HasPageLimit() bool {
	if o != nil && o.PageLimit != nil {
		return true
	}

	return false
}

// SetPageLimit gets a reference to the given int32 and assigns it to the PageLimit field.
func (o *SubscriptionPlanVM) SetPageLimit(v int32) {
	o.PageLimit = &v
}

func (o SubscriptionPlanVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.TimePeriodType != nil {
		toSerialize["timePeriodType"] = o.TimePeriodType
	}
	if o.TimePeriod != nil {
		toSerialize["timePeriod"] = o.TimePeriod
	}
	if o.ReadonlyTimeLimitType != nil {
		toSerialize["readonlyTimeLimitType"] = o.ReadonlyTimeLimitType
	}
	if o.ReadonlyTimeLimit != nil {
		toSerialize["readonlyTimeLimit"] = o.ReadonlyTimeLimit
	}
	if o.TemplatesSpaceLimit != nil {
		toSerialize["templatesSpaceLimit"] = o.TemplatesSpaceLimit
	}
	if o.ReportsSpaceLimit != nil {
		toSerialize["reportsSpaceLimit"] = o.ReportsSpaceLimit
	}
	if o.ExportsSpaceLimit != nil {
		toSerialize["exportsSpaceLimit"] = o.ExportsSpaceLimit
	}
	if o.FileUploadSizeLimit != nil {
		toSerialize["fileUploadSizeLimit"] = o.FileUploadSizeLimit
	}
	if o.DataSourceLimit != nil {
		toSerialize["dataSourceLimit"] = o.DataSourceLimit
	}
	if o.MaxUsersCount != nil {
		toSerialize["maxUsersCount"] = o.MaxUsersCount
	}
	if o.HasSpaceOverdraft != nil {
		toSerialize["hasSpaceOverdraft"] = o.HasSpaceOverdraft
	}
	if o.GroupLimit != nil {
		toSerialize["groupLimit"] = o.GroupLimit
	}
	if o.OnlineDesigner != nil {
		toSerialize["onlineDesigner"] = o.OnlineDesigner
	}
	if o.IsDemo != nil {
		toSerialize["isDemo"] = o.IsDemo
	}
	if o.UrlToBuy != nil {
		toSerialize["urlToBuy"] = o.UrlToBuy
	}
	if o.UnlimitedPage != nil {
		toSerialize["unlimitedPage"] = o.UnlimitedPage
	}
	if o.PageLimit != nil {
		toSerialize["pageLimit"] = o.PageLimit
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionPlanVM struct {
	value *SubscriptionPlanVM
	isSet bool
}

func (v NullableSubscriptionPlanVM) Get() *SubscriptionPlanVM {
	return v.value
}

func (v *NullableSubscriptionPlanVM) Set(val *SubscriptionPlanVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPlanVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPlanVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPlanVM(val *SubscriptionPlanVM) *NullableSubscriptionPlanVM {
	return &NullableSubscriptionPlanVM{value: val, isSet: true}
}

func (v NullableSubscriptionPlanVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPlanVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


