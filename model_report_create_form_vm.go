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

// checks if the ReportCreateFormVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportCreateFormVM{}

// ReportCreateFormVM struct for ReportCreateFormVM
type ReportCreateFormVM struct {
	FileCreateFormVM
	TemplateId NullableString `json:"templateId,omitempty"`
	T string `json:"$t"`
}

type _ReportCreateFormVM ReportCreateFormVM

// NewReportCreateFormVM instantiates a new ReportCreateFormVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportCreateFormVM(t string, fileContent *os.File) *ReportCreateFormVM {
	this := ReportCreateFormVM{}
	this.FileContent = fileContent
	this.T = t
	return &this
}

// NewReportCreateFormVMWithDefaults instantiates a new ReportCreateFormVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportCreateFormVMWithDefaults() *ReportCreateFormVM {
	this := ReportCreateFormVM{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportCreateFormVM) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId.Get()) {
		var ret string
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportCreateFormVM) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *ReportCreateFormVM) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableString and assigns it to the TemplateId field.
func (o *ReportCreateFormVM) SetTemplateId(v string) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *ReportCreateFormVM) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *ReportCreateFormVM) UnsetTemplateId() {
	o.TemplateId.Unset()
}

// GetT returns the T field value
func (o *ReportCreateFormVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *ReportCreateFormVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *ReportCreateFormVM) SetT(v string) {
	o.T = v
}

func (o ReportCreateFormVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportCreateFormVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedFileCreateFormVM, errFileCreateFormVM := json.Marshal(o.FileCreateFormVM)
	if errFileCreateFormVM != nil {
		return map[string]interface{}{}, errFileCreateFormVM
	}
	errFileCreateFormVM = json.Unmarshal([]byte(serializedFileCreateFormVM), &toSerialize)
	if errFileCreateFormVM != nil {
		return map[string]interface{}{}, errFileCreateFormVM
	}
	if o.TemplateId.IsSet() {
		toSerialize["templateId"] = o.TemplateId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *ReportCreateFormVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
		"fileContent",
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

	varReportCreateFormVM := _ReportCreateFormVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportCreateFormVM)

	if err != nil {
		return err
	}

	*o = ReportCreateFormVM(varReportCreateFormVM)

	return err
}

type NullableReportCreateFormVM struct {
	value *ReportCreateFormVM
	isSet bool
}

func (v NullableReportCreateFormVM) Get() *ReportCreateFormVM {
	return v.value
}

func (v *NullableReportCreateFormVM) Set(val *ReportCreateFormVM) {
	v.value = val
	v.isSet = true
}

func (v NullableReportCreateFormVM) IsSet() bool {
	return v.isSet
}

func (v *NullableReportCreateFormVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportCreateFormVM(val *ReportCreateFormVM) *NullableReportCreateFormVM {
	return &NullableReportCreateFormVM{value: val, isSet: true}
}

func (v NullableReportCreateFormVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportCreateFormVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

