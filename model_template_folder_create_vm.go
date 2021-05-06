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

// TemplateFolderCreateVM struct for TemplateFolderCreateVM
type TemplateFolderCreateVM struct {
	Name *string `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Icon *string `json:"icon,omitempty"`
}

// NewTemplateFolderCreateVM instantiates a new TemplateFolderCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateFolderCreateVM() *TemplateFolderCreateVM {
	this := TemplateFolderCreateVM{}
	return &this
}

// NewTemplateFolderCreateVMWithDefaults instantiates a new TemplateFolderCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateFolderCreateVMWithDefaults() *TemplateFolderCreateVM {
	this := TemplateFolderCreateVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateFolderCreateVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFolderCreateVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateFolderCreateVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateFolderCreateVM) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TemplateFolderCreateVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFolderCreateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TemplateFolderCreateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TemplateFolderCreateVM) SetTags(v []string) {
	o.Tags = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *TemplateFolderCreateVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateFolderCreateVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *TemplateFolderCreateVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *TemplateFolderCreateVM) SetIcon(v string) {
	o.Icon = &v
}

func (o TemplateFolderCreateVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateFolderCreateVM struct {
	value *TemplateFolderCreateVM
	isSet bool
}

func (v NullableTemplateFolderCreateVM) Get() *TemplateFolderCreateVM {
	return v.value
}

func (v *NullableTemplateFolderCreateVM) Set(val *TemplateFolderCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateFolderCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateFolderCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateFolderCreateVM(val *TemplateFolderCreateVM) *NullableTemplateFolderCreateVM {
	return &NullableTemplateFolderCreateVM{value: val, isSet: true}
}

func (v NullableTemplateFolderCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateFolderCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

