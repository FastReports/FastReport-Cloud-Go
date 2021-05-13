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

// AdminTemplateFolderCreateVM struct for AdminTemplateFolderCreateVM
type AdminTemplateFolderCreateVM struct {
	ParentId *string `json:"parentId,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Name *string `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Icon *string `json:"icon,omitempty"`
}

// NewAdminTemplateFolderCreateVM instantiates a new AdminTemplateFolderCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminTemplateFolderCreateVM() *AdminTemplateFolderCreateVM {
	this := AdminTemplateFolderCreateVM{}
	return &this
}

// NewAdminTemplateFolderCreateVMWithDefaults instantiates a new AdminTemplateFolderCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminTemplateFolderCreateVMWithDefaults() *AdminTemplateFolderCreateVM {
	this := AdminTemplateFolderCreateVM{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *AdminTemplateFolderCreateVM) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminTemplateFolderCreateVM) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *AdminTemplateFolderCreateVM) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *AdminTemplateFolderCreateVM) SetParentId(v string) {
	o.ParentId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AdminTemplateFolderCreateVM) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminTemplateFolderCreateVM) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AdminTemplateFolderCreateVM) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AdminTemplateFolderCreateVM) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdminTemplateFolderCreateVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminTemplateFolderCreateVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdminTemplateFolderCreateVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdminTemplateFolderCreateVM) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdminTemplateFolderCreateVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminTemplateFolderCreateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdminTemplateFolderCreateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AdminTemplateFolderCreateVM) SetTags(v []string) {
	o.Tags = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *AdminTemplateFolderCreateVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminTemplateFolderCreateVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *AdminTemplateFolderCreateVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *AdminTemplateFolderCreateVM) SetIcon(v string) {
	o.Icon = &v
}

func (o AdminTemplateFolderCreateVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
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
	return json.Marshal(toSerialize)
}

type NullableAdminTemplateFolderCreateVM struct {
	value *AdminTemplateFolderCreateVM
	isSet bool
}

func (v NullableAdminTemplateFolderCreateVM) Get() *AdminTemplateFolderCreateVM {
	return v.value
}

func (v *NullableAdminTemplateFolderCreateVM) Set(val *AdminTemplateFolderCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminTemplateFolderCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminTemplateFolderCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminTemplateFolderCreateVM(val *AdminTemplateFolderCreateVM) *NullableAdminTemplateFolderCreateVM {
	return &NullableAdminTemplateFolderCreateVM{value: val, isSet: true}
}

func (v NullableAdminTemplateFolderCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminTemplateFolderCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


