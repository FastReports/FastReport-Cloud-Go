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

// checks if the FolderIconVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderIconVM{}

// FolderIconVM struct for FolderIconVM
type FolderIconVM struct {
	Icon string `json:"icon"`
}

// NewFolderIconVM instantiates a new FolderIconVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderIconVM(icon string) *FolderIconVM {
	this := FolderIconVM{}
	this.Icon = icon
	return &this
}

// NewFolderIconVMWithDefaults instantiates a new FolderIconVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderIconVMWithDefaults() *FolderIconVM {
	this := FolderIconVM{}
	return &this
}

// GetIcon returns the Icon field value
func (o *FolderIconVM) GetIcon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Icon
}

// GetIconOk returns a tuple with the Icon field value
// and a boolean to check if the value has been set.
func (o *FolderIconVM) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icon, true
}

// SetIcon sets field value
func (o *FolderIconVM) SetIcon(v string) {
	o.Icon = v
}

func (o FolderIconVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderIconVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["icon"] = o.Icon
	return toSerialize, nil
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


