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

// checks if the SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission{}

// SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission struct for SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission
type SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission struct {
	Create *SubscriptionCreate `json:"create,omitempty"`
	Delete *SubscriptionDelete `json:"delete,omitempty"`
	Execute *SubscriptionExecute `json:"execute,omitempty"`
	Get *SubscriptionGet `json:"get,omitempty"`
	Update *SubscriptionUpdate `json:"update,omitempty"`
	Administrate *SubscriptionAdministrate `json:"administrate,omitempty"`
}

// NewSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission instantiates a new SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission() *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission {
	this := SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission{}
	return &this
}

// NewSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermissionWithDefaults instantiates a new SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermissionWithDefaults() *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission {
	this := SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetCreate() SubscriptionCreate {
	if o == nil || IsNil(o.Create) {
		var ret SubscriptionCreate
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetCreateOk() (*SubscriptionCreate, bool) {
	if o == nil || IsNil(o.Create) {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) HasCreate() bool {
	if o != nil && !IsNil(o.Create) {
		return true
	}

	return false
}

// SetCreate gets a reference to the given SubscriptionCreate and assigns it to the Create field.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) SetCreate(v SubscriptionCreate) {
	o.Create = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetDelete() SubscriptionDelete {
	if o == nil || IsNil(o.Delete) {
		var ret SubscriptionDelete
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetDeleteOk() (*SubscriptionDelete, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given SubscriptionDelete and assigns it to the Delete field.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) SetDelete(v SubscriptionDelete) {
	o.Delete = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetExecute() SubscriptionExecute {
	if o == nil || IsNil(o.Execute) {
		var ret SubscriptionExecute
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetExecuteOk() (*SubscriptionExecute, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given SubscriptionExecute and assigns it to the Execute field.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) SetExecute(v SubscriptionExecute) {
	o.Execute = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetGet() SubscriptionGet {
	if o == nil || IsNil(o.Get) {
		var ret SubscriptionGet
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetGetOk() (*SubscriptionGet, bool) {
	if o == nil || IsNil(o.Get) {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) HasGet() bool {
	if o != nil && !IsNil(o.Get) {
		return true
	}

	return false
}

// SetGet gets a reference to the given SubscriptionGet and assigns it to the Get field.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) SetGet(v SubscriptionGet) {
	o.Get = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetUpdate() SubscriptionUpdate {
	if o == nil || IsNil(o.Update) {
		var ret SubscriptionUpdate
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetUpdateOk() (*SubscriptionUpdate, bool) {
	if o == nil || IsNil(o.Update) {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) HasUpdate() bool {
	if o != nil && !IsNil(o.Update) {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given SubscriptionUpdate and assigns it to the Update field.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) SetUpdate(v SubscriptionUpdate) {
	o.Update = &v
}

// GetAdministrate returns the Administrate field value if set, zero value otherwise.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetAdministrate() SubscriptionAdministrate {
	if o == nil || IsNil(o.Administrate) {
		var ret SubscriptionAdministrate
		return ret
	}
	return *o.Administrate
}

// GetAdministrateOk returns a tuple with the Administrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) GetAdministrateOk() (*SubscriptionAdministrate, bool) {
	if o == nil || IsNil(o.Administrate) {
		return nil, false
	}
	return o.Administrate, true
}

// HasAdministrate returns a boolean if a field has been set.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) HasAdministrate() bool {
	if o != nil && !IsNil(o.Administrate) {
		return true
	}

	return false
}

// SetAdministrate gets a reference to the given SubscriptionAdministrate and assigns it to the Administrate field.
func (o *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) SetAdministrate(v SubscriptionAdministrate) {
	o.Administrate = &v
}

func (o SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Create) {
		toSerialize["create"] = o.Create
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}
	if !IsNil(o.Get) {
		toSerialize["get"] = o.Get
	}
	if !IsNil(o.Update) {
		toSerialize["update"] = o.Update
	}
	if !IsNil(o.Administrate) {
		toSerialize["administrate"] = o.Administrate
	}
	return toSerialize, nil
}

type NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission struct {
	value *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission
	isSet bool
}

func (v NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) Get() *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission {
	return v.value
}

func (v *NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) Set(val *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission(val *SubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) *NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission {
	return &NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission{value: val, isSet: true}
}

func (v NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionCreateSubscriptionGetSubscriptionUpdateSubscriptionDeleteSubscriptionExecuteSubscriptionAdministratePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


