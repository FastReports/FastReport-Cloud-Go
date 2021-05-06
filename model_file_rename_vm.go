/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FastReport.Cloud.SDK

import (
	"encoding/json"
)

// FileRenameVM struct for FileRenameVM
type FileRenameVM struct {
	Name *string `json:"name,omitempty"`
}

// NewFileRenameVM instantiates a new FileRenameVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileRenameVM() *FileRenameVM {
	this := FileRenameVM{}
	return &this
}

// NewFileRenameVMWithDefaults instantiates a new FileRenameVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileRenameVMWithDefaults() *FileRenameVM {
	this := FileRenameVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FileRenameVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRenameVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FileRenameVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FileRenameVM) SetName(v string) {
	o.Name = &v
}

func (o FileRenameVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableFileRenameVM struct {
	value *FileRenameVM
	isSet bool
}

func (v NullableFileRenameVM) Get() *FileRenameVM {
	return v.value
}

func (v *NullableFileRenameVM) Set(val *FileRenameVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFileRenameVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFileRenameVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileRenameVM(val *FileRenameVM) *NullableFileRenameVM {
	return &NullableFileRenameVM{value: val, isSet: true}
}

func (v NullableFileRenameVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileRenameVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

