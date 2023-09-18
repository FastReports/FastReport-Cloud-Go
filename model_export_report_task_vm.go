/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
)

// checks if the ExportReportTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportReportTaskVM{}

// ExportReportTaskVM struct for ExportReportTaskVM
type ExportReportTaskVM struct {
	ExportParameters map[string]string `json:"exportParameters,omitempty"`
	Format *ExportFormat `json:"format,omitempty"`
	PagesCount *int32 `json:"pagesCount,omitempty"`
	T string `json:"$t"`
}

// NewExportReportTaskVM instantiates a new ExportReportTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportReportTaskVM(t string) *ExportReportTaskVM {
	this := ExportReportTaskVM{}
	this.T = t
	return &this
}

// NewExportReportTaskVMWithDefaults instantiates a new ExportReportTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportReportTaskVMWithDefaults() *ExportReportTaskVM {
	this := ExportReportTaskVM{}
	return &this
}

// GetExportParameters returns the ExportParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportReportTaskVM) GetExportParameters() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ExportParameters
}

// GetExportParametersOk returns a tuple with the ExportParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportReportTaskVM) GetExportParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ExportParameters) {
		return nil, false
	}
	return &o.ExportParameters, true
}

// HasExportParameters returns a boolean if a field has been set.
func (o *ExportReportTaskVM) HasExportParameters() bool {
	if o != nil && IsNil(o.ExportParameters) {
		return true
	}

	return false
}

// SetExportParameters gets a reference to the given map[string]string and assigns it to the ExportParameters field.
func (o *ExportReportTaskVM) SetExportParameters(v map[string]string) {
	o.ExportParameters = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ExportReportTaskVM) GetFormat() ExportFormat {
	if o == nil || IsNil(o.Format) {
		var ret ExportFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportReportTaskVM) GetFormatOk() (*ExportFormat, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ExportReportTaskVM) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given ExportFormat and assigns it to the Format field.
func (o *ExportReportTaskVM) SetFormat(v ExportFormat) {
	o.Format = &v
}

// GetPagesCount returns the PagesCount field value if set, zero value otherwise.
func (o *ExportReportTaskVM) GetPagesCount() int32 {
	if o == nil || IsNil(o.PagesCount) {
		var ret int32
		return ret
	}
	return *o.PagesCount
}

// GetPagesCountOk returns a tuple with the PagesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportReportTaskVM) GetPagesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PagesCount) {
		return nil, false
	}
	return o.PagesCount, true
}

// HasPagesCount returns a boolean if a field has been set.
func (o *ExportReportTaskVM) HasPagesCount() bool {
	if o != nil && !IsNil(o.PagesCount) {
		return true
	}

	return false
}

// SetPagesCount gets a reference to the given int32 and assigns it to the PagesCount field.
func (o *ExportReportTaskVM) SetPagesCount(v int32) {
	o.PagesCount = &v
}

// GetT returns the T field value
func (o *ExportReportTaskVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *ExportReportTaskVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *ExportReportTaskVM) SetT(v string) {
	o.T = v
}

func (o ExportReportTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportReportTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExportParameters != nil {
		toSerialize["exportParameters"] = o.ExportParameters
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.PagesCount) {
		toSerialize["pagesCount"] = o.PagesCount
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

type NullableExportReportTaskVM struct {
	value *ExportReportTaskVM
	isSet bool
}

func (v NullableExportReportTaskVM) Get() *ExportReportTaskVM {
	return v.value
}

func (v *NullableExportReportTaskVM) Set(val *ExportReportTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableExportReportTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableExportReportTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportReportTaskVM(val *ExportReportTaskVM) *NullableExportReportTaskVM {
	return &NullableExportReportTaskVM{value: val, isSet: true}
}

func (v NullableExportReportTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportReportTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


