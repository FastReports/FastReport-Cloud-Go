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

// FileTagsUpdateVM struct for FileTagsUpdateVM
type FileTagsUpdateVM struct {
	Tags *[]string `json:"tags,omitempty"`
}

// NewFileTagsUpdateVM instantiates a new FileTagsUpdateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileTagsUpdateVM() *FileTagsUpdateVM {
	this := FileTagsUpdateVM{}
	return &this
}

// NewFileTagsUpdateVMWithDefaults instantiates a new FileTagsUpdateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileTagsUpdateVMWithDefaults() *FileTagsUpdateVM {
	this := FileTagsUpdateVM{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FileTagsUpdateVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTagsUpdateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FileTagsUpdateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FileTagsUpdateVM) SetTags(v []string) {
	o.Tags = &v
}

func (o FileTagsUpdateVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableFileTagsUpdateVM struct {
	value *FileTagsUpdateVM
	isSet bool
}

func (v NullableFileTagsUpdateVM) Get() *FileTagsUpdateVM {
	return v.value
}

func (v *NullableFileTagsUpdateVM) Set(val *FileTagsUpdateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFileTagsUpdateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFileTagsUpdateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileTagsUpdateVM(val *FileTagsUpdateVM) *NullableFileTagsUpdateVM {
	return &NullableFileTagsUpdateVM{value: val, isSet: true}
}

func (v NullableFileTagsUpdateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileTagsUpdateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


