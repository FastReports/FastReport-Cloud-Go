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

// checks if the CreateTransformTaskBaseVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTransformTaskBaseVM{}

// CreateTransformTaskBaseVM struct for CreateTransformTaskBaseVM
type CreateTransformTaskBaseVM struct {
	InputFile *InputFileVM `json:"inputFile,omitempty"`
	Locale NullableString `json:"locale,omitempty"`
	OutputFile *OutputFileVM `json:"outputFile,omitempty"`
	Transports []CreateTransportTaskBaseVM `json:"transports,omitempty"`
	T string `json:"$t"`
}

// NewCreateTransformTaskBaseVM instantiates a new CreateTransformTaskBaseVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransformTaskBaseVM(t string) *CreateTransformTaskBaseVM {
	this := CreateTransformTaskBaseVM{}
	this.T = t
	return &this
}

// NewCreateTransformTaskBaseVMWithDefaults instantiates a new CreateTransformTaskBaseVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransformTaskBaseVMWithDefaults() *CreateTransformTaskBaseVM {
	this := CreateTransformTaskBaseVM{}
	return &this
}

// GetInputFile returns the InputFile field value if set, zero value otherwise.
func (o *CreateTransformTaskBaseVM) GetInputFile() InputFileVM {
	if o == nil || IsNil(o.InputFile) {
		var ret InputFileVM
		return ret
	}
	return *o.InputFile
}

// GetInputFileOk returns a tuple with the InputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformTaskBaseVM) GetInputFileOk() (*InputFileVM, bool) {
	if o == nil || IsNil(o.InputFile) {
		return nil, false
	}
	return o.InputFile, true
}

// HasInputFile returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasInputFile() bool {
	if o != nil && !IsNil(o.InputFile) {
		return true
	}

	return false
}

// SetInputFile gets a reference to the given InputFileVM and assigns it to the InputFile field.
func (o *CreateTransformTaskBaseVM) SetInputFile(v InputFileVM) {
	o.InputFile = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTransformTaskBaseVM) GetLocale() string {
	if o == nil || IsNil(o.Locale.Get()) {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTransformTaskBaseVM) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *CreateTransformTaskBaseVM) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *CreateTransformTaskBaseVM) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *CreateTransformTaskBaseVM) UnsetLocale() {
	o.Locale.Unset()
}

// GetOutputFile returns the OutputFile field value if set, zero value otherwise.
func (o *CreateTransformTaskBaseVM) GetOutputFile() OutputFileVM {
	if o == nil || IsNil(o.OutputFile) {
		var ret OutputFileVM
		return ret
	}
	return *o.OutputFile
}

// GetOutputFileOk returns a tuple with the OutputFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTransformTaskBaseVM) GetOutputFileOk() (*OutputFileVM, bool) {
	if o == nil || IsNil(o.OutputFile) {
		return nil, false
	}
	return o.OutputFile, true
}

// HasOutputFile returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasOutputFile() bool {
	if o != nil && !IsNil(o.OutputFile) {
		return true
	}

	return false
}

// SetOutputFile gets a reference to the given OutputFileVM and assigns it to the OutputFile field.
func (o *CreateTransformTaskBaseVM) SetOutputFile(v OutputFileVM) {
	o.OutputFile = &v
}

// GetTransports returns the Transports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateTransformTaskBaseVM) GetTransports() []CreateTransportTaskBaseVM {
	if o == nil {
		var ret []CreateTransportTaskBaseVM
		return ret
	}
	return o.Transports
}

// GetTransportsOk returns a tuple with the Transports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateTransformTaskBaseVM) GetTransportsOk() ([]CreateTransportTaskBaseVM, bool) {
	if o == nil || IsNil(o.Transports) {
		return nil, false
	}
	return o.Transports, true
}

// HasTransports returns a boolean if a field has been set.
func (o *CreateTransformTaskBaseVM) HasTransports() bool {
	if o != nil && IsNil(o.Transports) {
		return true
	}

	return false
}

// SetTransports gets a reference to the given []CreateTransportTaskBaseVM and assigns it to the Transports field.
func (o *CreateTransformTaskBaseVM) SetTransports(v []CreateTransportTaskBaseVM) {
	o.Transports = v
}

// GetT returns the T field value
func (o *CreateTransformTaskBaseVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateTransformTaskBaseVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateTransformTaskBaseVM) SetT(v string) {
	o.T = v
}

func (o CreateTransformTaskBaseVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTransformTaskBaseVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InputFile) {
		toSerialize["inputFile"] = o.InputFile
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	if !IsNil(o.OutputFile) {
		toSerialize["outputFile"] = o.OutputFile
	}
	if o.Transports != nil {
		toSerialize["transports"] = o.Transports
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

type NullableCreateTransformTaskBaseVM struct {
	value *CreateTransformTaskBaseVM
	isSet bool
}

func (v NullableCreateTransformTaskBaseVM) Get() *CreateTransformTaskBaseVM {
	return v.value
}

func (v *NullableCreateTransformTaskBaseVM) Set(val *CreateTransformTaskBaseVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransformTaskBaseVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransformTaskBaseVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransformTaskBaseVM(val *CreateTransformTaskBaseVM) *NullableCreateTransformTaskBaseVM {
	return &NullableCreateTransformTaskBaseVM{value: val, isSet: true}
}

func (v NullableCreateTransformTaskBaseVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransformTaskBaseVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


