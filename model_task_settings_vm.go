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

// TaskSettingsVM struct for TaskSettingsVM
type TaskSettingsVM struct {
	Prepare NullableBool `json:"prepare,omitempty"`
	ExportTemplate NullableBool `json:"exportTemplate,omitempty"`
	ExportReport NullableBool `json:"exportReport,omitempty"`
	SendViaEmail NullableBool `json:"sendViaEmail,omitempty"`
	SendViaWebhook NullableBool `json:"sendViaWebhook,omitempty"`
	FetchData NullableBool `json:"fetchData,omitempty"`
	ThumbnailReport NullableBool `json:"thumbnailReport,omitempty"`
	ThumbnailTemplate NullableBool `json:"thumbnailTemplate,omitempty"`
}

// NewTaskSettingsVM instantiates a new TaskSettingsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskSettingsVM() *TaskSettingsVM {
	this := TaskSettingsVM{}
	return &this
}

// NewTaskSettingsVMWithDefaults instantiates a new TaskSettingsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskSettingsVMWithDefaults() *TaskSettingsVM {
	this := TaskSettingsVM{}
	return &this
}

// GetPrepare returns the Prepare field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetPrepare() bool {
	if o == nil || o.Prepare.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Prepare.Get()
}

// GetPrepareOk returns a tuple with the Prepare field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetPrepareOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Prepare.Get(), o.Prepare.IsSet()
}

// HasPrepare returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasPrepare() bool {
	if o != nil && o.Prepare.IsSet() {
		return true
	}

	return false
}

// SetPrepare gets a reference to the given NullableBool and assigns it to the Prepare field.
func (o *TaskSettingsVM) SetPrepare(v bool) {
	o.Prepare.Set(&v)
}
// SetPrepareNil sets the value for Prepare to be an explicit nil
func (o *TaskSettingsVM) SetPrepareNil() {
	o.Prepare.Set(nil)
}

// UnsetPrepare ensures that no value is present for Prepare, not even an explicit nil
func (o *TaskSettingsVM) UnsetPrepare() {
	o.Prepare.Unset()
}

// GetExportTemplate returns the ExportTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetExportTemplate() bool {
	if o == nil || o.ExportTemplate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ExportTemplate.Get()
}

// GetExportTemplateOk returns a tuple with the ExportTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetExportTemplateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExportTemplate.Get(), o.ExportTemplate.IsSet()
}

// HasExportTemplate returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasExportTemplate() bool {
	if o != nil && o.ExportTemplate.IsSet() {
		return true
	}

	return false
}

// SetExportTemplate gets a reference to the given NullableBool and assigns it to the ExportTemplate field.
func (o *TaskSettingsVM) SetExportTemplate(v bool) {
	o.ExportTemplate.Set(&v)
}
// SetExportTemplateNil sets the value for ExportTemplate to be an explicit nil
func (o *TaskSettingsVM) SetExportTemplateNil() {
	o.ExportTemplate.Set(nil)
}

// UnsetExportTemplate ensures that no value is present for ExportTemplate, not even an explicit nil
func (o *TaskSettingsVM) UnsetExportTemplate() {
	o.ExportTemplate.Unset()
}

// GetExportReport returns the ExportReport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetExportReport() bool {
	if o == nil || o.ExportReport.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ExportReport.Get()
}

// GetExportReportOk returns a tuple with the ExportReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetExportReportOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExportReport.Get(), o.ExportReport.IsSet()
}

// HasExportReport returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasExportReport() bool {
	if o != nil && o.ExportReport.IsSet() {
		return true
	}

	return false
}

// SetExportReport gets a reference to the given NullableBool and assigns it to the ExportReport field.
func (o *TaskSettingsVM) SetExportReport(v bool) {
	o.ExportReport.Set(&v)
}
// SetExportReportNil sets the value for ExportReport to be an explicit nil
func (o *TaskSettingsVM) SetExportReportNil() {
	o.ExportReport.Set(nil)
}

// UnsetExportReport ensures that no value is present for ExportReport, not even an explicit nil
func (o *TaskSettingsVM) UnsetExportReport() {
	o.ExportReport.Unset()
}

// GetSendViaEmail returns the SendViaEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetSendViaEmail() bool {
	if o == nil || o.SendViaEmail.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SendViaEmail.Get()
}

// GetSendViaEmailOk returns a tuple with the SendViaEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetSendViaEmailOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SendViaEmail.Get(), o.SendViaEmail.IsSet()
}

// HasSendViaEmail returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasSendViaEmail() bool {
	if o != nil && o.SendViaEmail.IsSet() {
		return true
	}

	return false
}

// SetSendViaEmail gets a reference to the given NullableBool and assigns it to the SendViaEmail field.
func (o *TaskSettingsVM) SetSendViaEmail(v bool) {
	o.SendViaEmail.Set(&v)
}
// SetSendViaEmailNil sets the value for SendViaEmail to be an explicit nil
func (o *TaskSettingsVM) SetSendViaEmailNil() {
	o.SendViaEmail.Set(nil)
}

// UnsetSendViaEmail ensures that no value is present for SendViaEmail, not even an explicit nil
func (o *TaskSettingsVM) UnsetSendViaEmail() {
	o.SendViaEmail.Unset()
}

// GetSendViaWebhook returns the SendViaWebhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetSendViaWebhook() bool {
	if o == nil || o.SendViaWebhook.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SendViaWebhook.Get()
}

// GetSendViaWebhookOk returns a tuple with the SendViaWebhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetSendViaWebhookOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SendViaWebhook.Get(), o.SendViaWebhook.IsSet()
}

// HasSendViaWebhook returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasSendViaWebhook() bool {
	if o != nil && o.SendViaWebhook.IsSet() {
		return true
	}

	return false
}

// SetSendViaWebhook gets a reference to the given NullableBool and assigns it to the SendViaWebhook field.
func (o *TaskSettingsVM) SetSendViaWebhook(v bool) {
	o.SendViaWebhook.Set(&v)
}
// SetSendViaWebhookNil sets the value for SendViaWebhook to be an explicit nil
func (o *TaskSettingsVM) SetSendViaWebhookNil() {
	o.SendViaWebhook.Set(nil)
}

// UnsetSendViaWebhook ensures that no value is present for SendViaWebhook, not even an explicit nil
func (o *TaskSettingsVM) UnsetSendViaWebhook() {
	o.SendViaWebhook.Unset()
}

// GetFetchData returns the FetchData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetFetchData() bool {
	if o == nil || o.FetchData.Get() == nil {
		var ret bool
		return ret
	}
	return *o.FetchData.Get()
}

// GetFetchDataOk returns a tuple with the FetchData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetFetchDataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FetchData.Get(), o.FetchData.IsSet()
}

// HasFetchData returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasFetchData() bool {
	if o != nil && o.FetchData.IsSet() {
		return true
	}

	return false
}

// SetFetchData gets a reference to the given NullableBool and assigns it to the FetchData field.
func (o *TaskSettingsVM) SetFetchData(v bool) {
	o.FetchData.Set(&v)
}
// SetFetchDataNil sets the value for FetchData to be an explicit nil
func (o *TaskSettingsVM) SetFetchDataNil() {
	o.FetchData.Set(nil)
}

// UnsetFetchData ensures that no value is present for FetchData, not even an explicit nil
func (o *TaskSettingsVM) UnsetFetchData() {
	o.FetchData.Unset()
}

// GetThumbnailReport returns the ThumbnailReport field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetThumbnailReport() bool {
	if o == nil || o.ThumbnailReport.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ThumbnailReport.Get()
}

// GetThumbnailReportOk returns a tuple with the ThumbnailReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetThumbnailReportOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThumbnailReport.Get(), o.ThumbnailReport.IsSet()
}

// HasThumbnailReport returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasThumbnailReport() bool {
	if o != nil && o.ThumbnailReport.IsSet() {
		return true
	}

	return false
}

// SetThumbnailReport gets a reference to the given NullableBool and assigns it to the ThumbnailReport field.
func (o *TaskSettingsVM) SetThumbnailReport(v bool) {
	o.ThumbnailReport.Set(&v)
}
// SetThumbnailReportNil sets the value for ThumbnailReport to be an explicit nil
func (o *TaskSettingsVM) SetThumbnailReportNil() {
	o.ThumbnailReport.Set(nil)
}

// UnsetThumbnailReport ensures that no value is present for ThumbnailReport, not even an explicit nil
func (o *TaskSettingsVM) UnsetThumbnailReport() {
	o.ThumbnailReport.Unset()
}

// GetThumbnailTemplate returns the ThumbnailTemplate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskSettingsVM) GetThumbnailTemplate() bool {
	if o == nil || o.ThumbnailTemplate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ThumbnailTemplate.Get()
}

// GetThumbnailTemplateOk returns a tuple with the ThumbnailTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskSettingsVM) GetThumbnailTemplateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ThumbnailTemplate.Get(), o.ThumbnailTemplate.IsSet()
}

// HasThumbnailTemplate returns a boolean if a field has been set.
func (o *TaskSettingsVM) HasThumbnailTemplate() bool {
	if o != nil && o.ThumbnailTemplate.IsSet() {
		return true
	}

	return false
}

// SetThumbnailTemplate gets a reference to the given NullableBool and assigns it to the ThumbnailTemplate field.
func (o *TaskSettingsVM) SetThumbnailTemplate(v bool) {
	o.ThumbnailTemplate.Set(&v)
}
// SetThumbnailTemplateNil sets the value for ThumbnailTemplate to be an explicit nil
func (o *TaskSettingsVM) SetThumbnailTemplateNil() {
	o.ThumbnailTemplate.Set(nil)
}

// UnsetThumbnailTemplate ensures that no value is present for ThumbnailTemplate, not even an explicit nil
func (o *TaskSettingsVM) UnsetThumbnailTemplate() {
	o.ThumbnailTemplate.Unset()
}

func (o TaskSettingsVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prepare.IsSet() {
		toSerialize["prepare"] = o.Prepare.Get()
	}
	if o.ExportTemplate.IsSet() {
		toSerialize["exportTemplate"] = o.ExportTemplate.Get()
	}
	if o.ExportReport.IsSet() {
		toSerialize["exportReport"] = o.ExportReport.Get()
	}
	if o.SendViaEmail.IsSet() {
		toSerialize["sendViaEmail"] = o.SendViaEmail.Get()
	}
	if o.SendViaWebhook.IsSet() {
		toSerialize["sendViaWebhook"] = o.SendViaWebhook.Get()
	}
	if o.FetchData.IsSet() {
		toSerialize["fetchData"] = o.FetchData.Get()
	}
	if o.ThumbnailReport.IsSet() {
		toSerialize["thumbnailReport"] = o.ThumbnailReport.Get()
	}
	if o.ThumbnailTemplate.IsSet() {
		toSerialize["thumbnailTemplate"] = o.ThumbnailTemplate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableTaskSettingsVM struct {
	value *TaskSettingsVM
	isSet bool
}

func (v NullableTaskSettingsVM) Get() *TaskSettingsVM {
	return v.value
}

func (v *NullableTaskSettingsVM) Set(val *TaskSettingsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskSettingsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskSettingsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskSettingsVM(val *TaskSettingsVM) *NullableTaskSettingsVM {
	return &NullableTaskSettingsVM{value: val, isSet: true}
}

func (v NullableTaskSettingsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskSettingsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


