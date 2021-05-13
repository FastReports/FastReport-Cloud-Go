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

// SubscriptionVM struct for SubscriptionVM
type SubscriptionVM struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Current *SubscriptionPeriodVM `json:"current,omitempty"`
	Old *[]SubscriptionPeriodVM `json:"old,omitempty"`
	TemplatesFolder *SubscriptionFolder `json:"templatesFolder,omitempty"`
	ReportsFolder *SubscriptionFolder `json:"reportsFolder,omitempty"`
	ExportsFolder *SubscriptionFolder `json:"exportsFolder,omitempty"`
}

// NewSubscriptionVM instantiates a new SubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionVM() *SubscriptionVM {
	this := SubscriptionVM{}
	return &this
}

// NewSubscriptionVMWithDefaults instantiates a new SubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionVMWithDefaults() *SubscriptionVM {
	this := SubscriptionVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionVM) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionVM) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionVM) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubscriptionVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubscriptionVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubscriptionVM) SetName(v string) {
	o.Name = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *SubscriptionVM) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *SubscriptionVM) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *SubscriptionVM) SetLocale(v string) {
	o.Locale = &v
}

// GetCurrent returns the Current field value if set, zero value otherwise.
func (o *SubscriptionVM) GetCurrent() SubscriptionPeriodVM {
	if o == nil || o.Current == nil {
		var ret SubscriptionPeriodVM
		return ret
	}
	return *o.Current
}

// GetCurrentOk returns a tuple with the Current field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetCurrentOk() (*SubscriptionPeriodVM, bool) {
	if o == nil || o.Current == nil {
		return nil, false
	}
	return o.Current, true
}

// HasCurrent returns a boolean if a field has been set.
func (o *SubscriptionVM) HasCurrent() bool {
	if o != nil && o.Current != nil {
		return true
	}

	return false
}

// SetCurrent gets a reference to the given SubscriptionPeriodVM and assigns it to the Current field.
func (o *SubscriptionVM) SetCurrent(v SubscriptionPeriodVM) {
	o.Current = &v
}

// GetOld returns the Old field value if set, zero value otherwise.
func (o *SubscriptionVM) GetOld() []SubscriptionPeriodVM {
	if o == nil || o.Old == nil {
		var ret []SubscriptionPeriodVM
		return ret
	}
	return *o.Old
}

// GetOldOk returns a tuple with the Old field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetOldOk() (*[]SubscriptionPeriodVM, bool) {
	if o == nil || o.Old == nil {
		return nil, false
	}
	return o.Old, true
}

// HasOld returns a boolean if a field has been set.
func (o *SubscriptionVM) HasOld() bool {
	if o != nil && o.Old != nil {
		return true
	}

	return false
}

// SetOld gets a reference to the given []SubscriptionPeriodVM and assigns it to the Old field.
func (o *SubscriptionVM) SetOld(v []SubscriptionPeriodVM) {
	o.Old = &v
}

// GetTemplatesFolder returns the TemplatesFolder field value if set, zero value otherwise.
func (o *SubscriptionVM) GetTemplatesFolder() SubscriptionFolder {
	if o == nil || o.TemplatesFolder == nil {
		var ret SubscriptionFolder
		return ret
	}
	return *o.TemplatesFolder
}

// GetTemplatesFolderOk returns a tuple with the TemplatesFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetTemplatesFolderOk() (*SubscriptionFolder, bool) {
	if o == nil || o.TemplatesFolder == nil {
		return nil, false
	}
	return o.TemplatesFolder, true
}

// HasTemplatesFolder returns a boolean if a field has been set.
func (o *SubscriptionVM) HasTemplatesFolder() bool {
	if o != nil && o.TemplatesFolder != nil {
		return true
	}

	return false
}

// SetTemplatesFolder gets a reference to the given SubscriptionFolder and assigns it to the TemplatesFolder field.
func (o *SubscriptionVM) SetTemplatesFolder(v SubscriptionFolder) {
	o.TemplatesFolder = &v
}

// GetReportsFolder returns the ReportsFolder field value if set, zero value otherwise.
func (o *SubscriptionVM) GetReportsFolder() SubscriptionFolder {
	if o == nil || o.ReportsFolder == nil {
		var ret SubscriptionFolder
		return ret
	}
	return *o.ReportsFolder
}

// GetReportsFolderOk returns a tuple with the ReportsFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetReportsFolderOk() (*SubscriptionFolder, bool) {
	if o == nil || o.ReportsFolder == nil {
		return nil, false
	}
	return o.ReportsFolder, true
}

// HasReportsFolder returns a boolean if a field has been set.
func (o *SubscriptionVM) HasReportsFolder() bool {
	if o != nil && o.ReportsFolder != nil {
		return true
	}

	return false
}

// SetReportsFolder gets a reference to the given SubscriptionFolder and assigns it to the ReportsFolder field.
func (o *SubscriptionVM) SetReportsFolder(v SubscriptionFolder) {
	o.ReportsFolder = &v
}

// GetExportsFolder returns the ExportsFolder field value if set, zero value otherwise.
func (o *SubscriptionVM) GetExportsFolder() SubscriptionFolder {
	if o == nil || o.ExportsFolder == nil {
		var ret SubscriptionFolder
		return ret
	}
	return *o.ExportsFolder
}

// GetExportsFolderOk returns a tuple with the ExportsFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionVM) GetExportsFolderOk() (*SubscriptionFolder, bool) {
	if o == nil || o.ExportsFolder == nil {
		return nil, false
	}
	return o.ExportsFolder, true
}

// HasExportsFolder returns a boolean if a field has been set.
func (o *SubscriptionVM) HasExportsFolder() bool {
	if o != nil && o.ExportsFolder != nil {
		return true
	}

	return false
}

// SetExportsFolder gets a reference to the given SubscriptionFolder and assigns it to the ExportsFolder field.
func (o *SubscriptionVM) SetExportsFolder(v SubscriptionFolder) {
	o.ExportsFolder = &v
}

func (o SubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Current != nil {
		toSerialize["current"] = o.Current
	}
	if o.Old != nil {
		toSerialize["old"] = o.Old
	}
	if o.TemplatesFolder != nil {
		toSerialize["templatesFolder"] = o.TemplatesFolder
	}
	if o.ReportsFolder != nil {
		toSerialize["reportsFolder"] = o.ReportsFolder
	}
	if o.ExportsFolder != nil {
		toSerialize["exportsFolder"] = o.ExportsFolder
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionVM struct {
	value *SubscriptionVM
	isSet bool
}

func (v NullableSubscriptionVM) Get() *SubscriptionVM {
	return v.value
}

func (v *NullableSubscriptionVM) Set(val *SubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionVM(val *SubscriptionVM) *NullableSubscriptionVM {
	return &NullableSubscriptionVM{value: val, isSet: true}
}

func (v NullableSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


