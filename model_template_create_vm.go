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

// TemplateCreateVM struct for TemplateCreateVM
type TemplateCreateVM struct {
	Name *string `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Content *string `json:"content,omitempty"`
}

// NewTemplateCreateVM instantiates a new TemplateCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCreateVM() *TemplateCreateVM {
	this := TemplateCreateVM{}
	return &this
}

// NewTemplateCreateVMWithDefaults instantiates a new TemplateCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCreateVMWithDefaults() *TemplateCreateVM {
	this := TemplateCreateVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateCreateVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateCreateVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateCreateVM) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TemplateCreateVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TemplateCreateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TemplateCreateVM) SetTags(v []string) {
	o.Tags = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *TemplateCreateVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *TemplateCreateVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *TemplateCreateVM) SetIcon(v string) {
	o.Icon = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TemplateCreateVM) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateVM) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TemplateCreateVM) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *TemplateCreateVM) SetContent(v string) {
	o.Content = &v
}

func (o TemplateCreateVM) MarshalJSON() ([]byte, error) {
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
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateCreateVM struct {
	value *TemplateCreateVM
	isSet bool
}

func (v NullableTemplateCreateVM) Get() *TemplateCreateVM {
	return v.value
}

func (v *NullableTemplateCreateVM) Set(val *TemplateCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCreateVM(val *TemplateCreateVM) *NullableTemplateCreateVM {
	return &NullableTemplateCreateVM{value: val, isSet: true}
}

func (v NullableTemplateCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

