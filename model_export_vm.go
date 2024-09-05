/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExportVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportVM{}

// ExportVM struct for ExportVM
type ExportVM struct {
	FileVM
	Format *ExportFormat `json:"format,omitempty"`
	ReportId NullableString `json:"reportId,omitempty"`
	TemplateId NullableString `json:"templateId,omitempty"`
	T string `json:"$t"`
}

type _ExportVM ExportVM

// NewExportVM instantiates a new ExportVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportVM(t string) *ExportVM {
	this := ExportVM{}
	this.T = t
	return &this
}

// NewExportVMWithDefaults instantiates a new ExportVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportVMWithDefaults() *ExportVM {
	this := ExportVM{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ExportVM) GetFormat() ExportFormat {
	if o == nil || IsNil(o.Format) {
		var ret ExportFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportVM) GetFormatOk() (*ExportFormat, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ExportVM) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given ExportFormat and assigns it to the Format field.
func (o *ExportVM) SetFormat(v ExportFormat) {
	o.Format = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportVM) GetReportId() string {
	if o == nil || IsNil(o.ReportId.Get()) {
		var ret string
		return ret
	}
	return *o.ReportId.Get()
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportVM) GetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportId.Get(), o.ReportId.IsSet()
}

// HasReportId returns a boolean if a field has been set.
func (o *ExportVM) HasReportId() bool {
	if o != nil && o.ReportId.IsSet() {
		return true
	}

	return false
}

// SetReportId gets a reference to the given NullableString and assigns it to the ReportId field.
func (o *ExportVM) SetReportId(v string) {
	o.ReportId.Set(&v)
}
// SetReportIdNil sets the value for ReportId to be an explicit nil
func (o *ExportVM) SetReportIdNil() {
	o.ReportId.Set(nil)
}

// UnsetReportId ensures that no value is present for ReportId, not even an explicit nil
func (o *ExportVM) UnsetReportId() {
	o.ReportId.Unset()
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportVM) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportVM) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ExportVM) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *ExportVM) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *ExportVM) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *ExportVM) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetT returns the T field value
func (o *ExportVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *ExportVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *ExportVM) SetT(v string) {
	o.T = v
}

func (o ExportVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFileVM, errFileVM := json.Marshal(o.FileVM)
	if errFileVM != nil {
		return map[string]interface{}{}, errFileVM
	}
	errFileVM = json.Unmarshal([]byte(serializedFileVM), &toSerialize)
	if errFileVM != nil {
		return map[string]interface{}{}, errFileVM
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if o.ReportId.IsSet() {
		toSerialize["reportId"] = o.ReportId.Get()
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *ExportVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExportVM := _ExportVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExportVM)

	if err != nil {
		return err
	}

	*o = ExportVM(varExportVM)

	return err
}

type NullableExportVM struct {
	value *ExportVM
	isSet bool
}

func (v NullableExportVM) Get() *ExportVM {
	return v.value
}

func (v *NullableExportVM) Set(val *ExportVM) {
	v.value = val
	v.isSet = true
}

func (v NullableExportVM) IsSet() bool {
	return v.isSet
}

func (v *NullableExportVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportVM(val *ExportVM) *NullableExportVM {
	return &NullableExportVM{value: val, isSet: true}
}

func (v NullableExportVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


