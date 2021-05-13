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

// SubscriptionPermission struct for SubscriptionPermission
type SubscriptionPermission struct {
	Create *int32 `json:"create,omitempty"`
	Delete *int32 `json:"delete,omitempty"`
	Execute *int32 `json:"execute,omitempty"`
	Get *int32 `json:"get,omitempty"`
	Update *int32 `json:"update,omitempty"`
	Administrate *int32 `json:"administrate,omitempty"`
}

// NewSubscriptionPermission instantiates a new SubscriptionPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPermission() *SubscriptionPermission {
	this := SubscriptionPermission{}
	return &this
}

// NewSubscriptionPermissionWithDefaults instantiates a new SubscriptionPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPermissionWithDefaults() *SubscriptionPermission {
	this := SubscriptionPermission{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *SubscriptionPermission) GetCreate() int32 {
	if o == nil || o.Create == nil {
		var ret int32
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermission) GetCreateOk() (*int32, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *SubscriptionPermission) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given int32 and assigns it to the Create field.
func (o *SubscriptionPermission) SetCreate(v int32) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *SubscriptionPermission) GetDelete() int32 {
	if o == nil || o.Delete == nil {
		var ret int32
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermission) GetDeleteOk() (*int32, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *SubscriptionPermission) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given int32 and assigns it to the Delete field.
func (o *SubscriptionPermission) SetDelete(v int32) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *SubscriptionPermission) GetExecute() int32 {
	if o == nil || o.Execute == nil {
		var ret int32
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermission) GetExecuteOk() (*int32, bool) {
	if o == nil || o.Execute == nil {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *SubscriptionPermission) HasExecute() bool {
	if o != nil && o.Execute != nil {
		return true
	}

	return false
}

// SetExecute gets a reference to the given int32 and assigns it to the Execute field.
func (o *SubscriptionPermission) SetExecute(v int32) {
	o.Execute = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *SubscriptionPermission) GetGet() int32 {
	if o == nil || o.Get == nil {
		var ret int32
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermission) GetGetOk() (*int32, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *SubscriptionPermission) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given int32 and assigns it to the Get field.
func (o *SubscriptionPermission) SetGet(v int32) {
	o.Get = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *SubscriptionPermission) GetUpdate() int32 {
	if o == nil || o.Update == nil {
		var ret int32
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermission) GetUpdateOk() (*int32, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *SubscriptionPermission) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given int32 and assigns it to the Update field.
func (o *SubscriptionPermission) SetUpdate(v int32) {
	o.Update = &v
}

// GetAdministrate returns the Administrate field value if set, zero value otherwise.
func (o *SubscriptionPermission) GetAdministrate() int32 {
	if o == nil || o.Administrate == nil {
		var ret int32
		return ret
	}
	return *o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPermission) GetAdministrateOk() (*int32, bool) {
	if o == nil || o.Administrate == nil {
		return nil, false
	}
	return o.Administrate, true
}

// HasAdministrate returns a boolean if a field has been set.
func (o *SubscriptionPermission) HasAdministrate() bool {
	if o != nil && o.Administrate != nil {
		return true
	}

	return false
}

// SetAdministrate gets a reference to the given int32 and assigns it to the Administrate field.
func (o *SubscriptionPermission) SetAdministrate(v int32) {
	o.Administrate = &v
}

func (o SubscriptionPermission) MarshalJSON() ([]byte, error) {
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

type NullableSubscriptionPermission struct {
	value *SubscriptionPermission
	isSet bool
}

func (v NullableSubscriptionPermission) Get() *SubscriptionPermission {
	return v.value
}

func (v *NullableSubscriptionPermission) Set(val *SubscriptionPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPermission(val *SubscriptionPermission) *NullableSubscriptionPermission {
	return &NullableSubscriptionPermission{value: val, isSet: true}
}

func (v NullableSubscriptionPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


