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

// FolderIconVM struct for FolderIconVM
type FolderIconVM struct {
	Icon *string `json:"icon,omitempty"`
}

// NewFolderIconVM instantiates a new FolderIconVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderIconVM() *FolderIconVM {
	this := FolderIconVM{}
	return &this
}

// NewFolderIconVMWithDefaults instantiates a new FolderIconVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderIconVMWithDefaults() *FolderIconVM {
	this := FolderIconVM{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *FolderIconVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderIconVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *FolderIconVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *FolderIconVM) SetIcon(v string) {
	o.Icon = &v
}

func (o FolderIconVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableFolderIconVM struct {
	value *FolderIconVM
	isSet bool
}

func (v NullableFolderIconVM) Get() *FolderIconVM {
	return v.value
}

func (v *NullableFolderIconVM) Set(val *FolderIconVM) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderIconVM) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderIconVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderIconVM(val *FolderIconVM) *NullableFolderIconVM {
	return &NullableFolderIconVM{value: val, isSet: true}
}

func (v NullableFolderIconVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderIconVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


