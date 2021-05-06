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

// GroupPermission struct for GroupPermission
type GroupPermission struct {
	Create *int32 `json:"create,omitempty"`
	Delete *int32 `json:"delete,omitempty"`
	Execute *int32 `json:"execute,omitempty"`
	Get *int32 `json:"get,omitempty"`
	Update *int32 `json:"update,omitempty"`
	Administrate *int32 `json:"administrate,omitempty"`
}

// NewGroupPermission instantiates a new GroupPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupPermission() *GroupPermission {
	this := GroupPermission{}
	return &this
}

// NewGroupPermissionWithDefaults instantiates a new GroupPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupPermissionWithDefaults() *GroupPermission {
	this := GroupPermission{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *GroupPermission) GetCreate() int32 {
	if o == nil || o.Create == nil {
		var ret int32
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermission) GetCreateOk() (*int32, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *GroupPermission) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given int32 and assigns it to the Create field.
func (o *GroupPermission) SetCreate(v int32) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *GroupPermission) GetDelete() int32 {
	if o == nil || o.Delete == nil {
		var ret int32
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermission) GetDeleteOk() (*int32, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *GroupPermission) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given int32 and assigns it to the Delete field.
func (o *GroupPermission) SetDelete(v int32) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *GroupPermission) GetExecute() int32 {
	if o == nil || o.Execute == nil {
		var ret int32
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermission) GetExecuteOk() (*int32, bool) {
	if o == nil || o.Execute == nil {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *GroupPermission) HasExecute() bool {
	if o != nil && o.Execute != nil {
		return true
	}

	return false
}

// SetExecute gets a reference to the given int32 and assigns it to the Execute field.
func (o *GroupPermission) SetExecute(v int32) {
	o.Execute = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *GroupPermission) GetGet() int32 {
	if o == nil || o.Get == nil {
		var ret int32
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermission) GetGetOk() (*int32, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *GroupPermission) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given int32 and assigns it to the Get field.
func (o *GroupPermission) SetGet(v int32) {
	o.Get = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *GroupPermission) GetUpdate() int32 {
	if o == nil || o.Update == nil {
		var ret int32
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermission) GetUpdateOk() (*int32, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *GroupPermission) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given int32 and assigns it to the Update field.
func (o *GroupPermission) SetUpdate(v int32) {
	o.Update = &v
}

// GetAdministrate returns the Administrate field value if set, zero value otherwise.
func (o *GroupPermission) GetAdministrate() int32 {
	if o == nil || o.Administrate == nil {
		var ret int32
		return ret
	}
	return *o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupPermission) GetAdministrateOk() (*int32, bool) {
	if o == nil || o.Administrate == nil {
		return nil, false
	}
	return o.Administrate, true
}

// HasAdministrate returns a boolean if a field has been set.
func (o *GroupPermission) HasAdministrate() bool {
	if o != nil && o.Administrate != nil {
		return true
	}

	return false
}

// SetAdministrate gets a reference to the given int32 and assigns it to the Administrate field.
func (o *GroupPermission) SetAdministrate(v int32) {
	o.Administrate = &v
}

func (o GroupPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}
	if o.Execute != nil {
		toSerialize["execute"] = o.Execute
	}
	if o.Get != nil {
		toSerialize["get"] = o.Get
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	if o.Administrate != nil {
		toSerialize["administrate"] = o.Administrate
	}
	return json.Marshal(toSerialize)
}

type NullableGroupPermission struct {
	value *GroupPermission
	isSet bool
}

func (v NullableGroupPermission) Get() *GroupPermission {
	return v.value
}

func (v *NullableGroupPermission) Set(val *GroupPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupPermission(val *GroupPermission) *NullableGroupPermission {
	return &NullableGroupPermission{value: val, isSet: true}
}

func (v NullableGroupPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


