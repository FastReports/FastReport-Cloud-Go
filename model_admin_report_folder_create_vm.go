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

// AdminReportFolderCreateVM struct for AdminReportFolderCreateVM
type AdminReportFolderCreateVM struct {
	ParentId *string `json:"parentId,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Name *string `json:"name,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	Icon *string `json:"icon,omitempty"`
}

// NewAdminReportFolderCreateVM instantiates a new AdminReportFolderCreateVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminReportFolderCreateVM() *AdminReportFolderCreateVM {
	this := AdminReportFolderCreateVM{}
	return &this
}

// NewAdminReportFolderCreateVMWithDefaults instantiates a new AdminReportFolderCreateVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminReportFolderCreateVMWithDefaults() *AdminReportFolderCreateVM {
	this := AdminReportFolderCreateVM{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *AdminReportFolderCreateVM) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReportFolderCreateVM) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *AdminReportFolderCreateVM) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *AdminReportFolderCreateVM) SetParentId(v string) {
	o.ParentId = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AdminReportFolderCreateVM) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReportFolderCreateVM) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AdminReportFolderCreateVM) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *AdminReportFolderCreateVM) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdminReportFolderCreateVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReportFolderCreateVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdminReportFolderCreateVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdminReportFolderCreateVM) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AdminReportFolderCreateVM) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReportFolderCreateVM) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AdminReportFolderCreateVM) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AdminReportFolderCreateVM) SetTags(v []string) {
	o.Tags = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *AdminReportFolderCreateVM) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminReportFolderCreateVM) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *AdminReportFolderCreateVM) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *AdminReportFolderCreateVM) SetIcon(v string) {
	o.Icon = &v
}

func (o AdminReportFolderCreateVM) MarshalJSON() ([]byte, error) {
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

type NullableAdminReportFolderCreateVM struct {
	value *AdminReportFolderCreateVM
	isSet bool
}

func (v NullableAdminReportFolderCreateVM) Get() *AdminReportFolderCreateVM {
	return v.value
}

func (v *NullableAdminReportFolderCreateVM) Set(val *AdminReportFolderCreateVM) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminReportFolderCreateVM) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminReportFolderCreateVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminReportFolderCreateVM(val *AdminReportFolderCreateVM) *NullableAdminReportFolderCreateVM {
	return &NullableAdminReportFolderCreateVM{value: val, isSet: true}
}

func (v NullableAdminReportFolderCreateVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminReportFolderCreateVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


