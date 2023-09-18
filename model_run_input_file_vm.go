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

// checks if the RunInputFileVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunInputFileVM{}

// RunInputFileVM struct for RunInputFileVM
type RunInputFileVM struct {
	Content NullableString `json:"content,omitempty"`
	FileName NullableString `json:"fileName,omitempty"`
}

// NewRunInputFileVM instantiates a new RunInputFileVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunInputFileVM() *RunInputFileVM {
	this := RunInputFileVM{}
	return &this
}

// NewRunInputFileVMWithDefaults instantiates a new RunInputFileVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunInputFileVMWithDefaults() *RunInputFileVM {
	this := RunInputFileVM{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunInputFileVM) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunInputFileVM) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *RunInputFileVM) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *RunInputFileVM) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *RunInputFileVM) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *RunInputFileVM) UnsetContent() {
	o.Content.Unset()
}

// GetFileName returns the FileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunInputFileVM) GetFileName() string {
	if o == nil || IsNil(o.FileName.Get()) {
		var ret string
		return ret
	}
	return *o.FileName.Get()
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunInputFileVM) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileName.Get(), o.FileName.IsSet()
}

// HasFileName returns a boolean if a field has been set.
func (o *RunInputFileVM) HasFileName() bool {
	if o != nil && o.FileName.IsSet() {
		return true
	}

	return false
}

// SetFileName gets a reference to the given NullableString and assigns it to the FileName field.
func (o *RunInputFileVM) SetFileName(v string) {
	o.FileName.Set(&v)
}
// SetFileNameNil sets the value for FileName to be an explicit nil
func (o *RunInputFileVM) SetFileNameNil() {
	o.FileName.Set(nil)
}

// UnsetFileName ensures that no value is present for FileName, not even an explicit nil
func (o *RunInputFileVM) UnsetFileName() {
	o.FileName.Unset()
}

func (o RunInputFileVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunInputFileVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.FileName.IsSet() {
		toSerialize["fileName"] = o.FileName.Get()
	}
	return toSerialize, nil
}

type NullableRunInputFileVM struct {
	value *RunInputFileVM
	isSet bool
}

func (v NullableRunInputFileVM) Get() *RunInputFileVM {
	return v.value
}

func (v *NullableRunInputFileVM) Set(val *RunInputFileVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRunInputFileVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRunInputFileVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunInputFileVM(val *RunInputFileVM) *NullableRunInputFileVM {
	return &NullableRunInputFileVM{value: val, isSet: true}
}

func (v NullableRunInputFileVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunInputFileVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


