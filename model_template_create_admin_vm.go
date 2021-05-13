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

// TemplateCreateAdminVM struct for TemplateCreateAdminVM
type TemplateCreateAdminVM struct {
	OwnerId *string `json:"ownerId,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	Name *string `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Icon *string `json:"icon,omitempty"`
	Content *string `json:"content,omitempty"`
}

// NewTemplateCreateAdminVM instantiates a new TemplateCreateAdminVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCreateAdminVM() *TemplateCreateAdminVM {
	this := TemplateCreateAdminVM{}
	return &this
}

// NewTemplateCreateAdminVMWithDefaults instantiates a new TemplateCreateAdminVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCreateAdminVMWithDefaults() *TemplateCreateAdminVM {
	this := TemplateCreateAdminVM{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *TemplateCreateAdminVM) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateAdminVM) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *TemplateCreateAdminVM) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *TemplateCreateAdminVM) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *TemplateCreateAdminVM) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateAdminVM) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *TemplateCreateAdminVM) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *TemplateCreateAdminVM) SetParentId(v string) {
	o.ParentId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateCreateAdminVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateAdminVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateCreateAdminVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateCreateAdminVM) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TemplateCreateAdminVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateAdminVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TemplateCreateAdminVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TemplateCreateAdminVM) SetTags(v []string) {
	o.Tags = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *TemplateCreateAdminVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateAdminVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *TemplateCreateAdminVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *TemplateCreateAdminVM) SetIcon(v string) {
	o.Icon = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TemplateCreateAdminVM) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateAdminVM) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TemplateCreateAdminVM) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *TemplateCreateAdminVM) SetContent(v string) {
	o.Content = &v
}

func (o TemplateCreateAdminVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
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

type NullableTemplateCreateAdminVM struct {
	value *TemplateCreateAdminVM
	isSet bool
}

func (v NullableTemplateCreateAdminVM) Get() *TemplateCreateAdminVM {
	return v.value
}

func (v *NullableTemplateCreateAdminVM) Set(val *TemplateCreateAdminVM) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCreateAdminVM) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCreateAdminVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCreateAdminVM(val *TemplateCreateAdminVM) *NullableTemplateCreateAdminVM {
	return &NullableTemplateCreateAdminVM{value: val, isSet: true}
}

func (v NullableTemplateCreateAdminVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCreateAdminVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


