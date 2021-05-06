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

// UpdateSubscriptionVM struct for UpdateSubscriptionVM
type UpdateSubscriptionVM struct {
	Name *string `json:"name,omitempty"`
	Locale *string `json:"locale,omitempty"`
	DefaultPermissions *DefaultPermissions `json:"defaultPermissions,omitempty"`
}

// NewUpdateSubscriptionVM instantiates a new UpdateSubscriptionVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionVM() *UpdateSubscriptionVM {
	this := UpdateSubscriptionVM{}
	return &this
}

// NewUpdateSubscriptionVMWithDefaults instantiates a new UpdateSubscriptionVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionVMWithDefaults() *UpdateSubscriptionVM {
	this := UpdateSubscriptionVM{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSubscriptionVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSubscriptionVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSubscriptionVM) SetName(v string) {
	o.Name = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *UpdateSubscriptionVM) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionVM) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UpdateSubscriptionVM) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *UpdateSubscriptionVM) SetLocale(v string) {
	o.Locale = &v
}

// GetDefaultPermissions returns the DefaultPermissions field value if set, zero value otherwise.
func (o *UpdateSubscriptionVM) GetDefaultPermissions() DefaultPermissions {
	if o == nil || o.DefaultPermissions == nil {
		var ret DefaultPermissions
		return ret
	}
	return *o.DefaultPermissions
}

// GetDefaultPermissionsOk returns a tuple with the DefaultPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionVM) GetDefaultPermissionsOk() (*DefaultPermissions, bool) {
	if o == nil || o.DefaultPermissions == nil {
		return nil, false
	}
	return o.DefaultPermissions, true
}

// HasDefaultPermissions returns a boolean if a field has been set.
func (o *UpdateSubscriptionVM) HasDefaultPermissions() bool {
	if o != nil && o.DefaultPermissions != nil {
		return true
	}

	return false
}

// SetDefaultPermissions gets a reference to the given DefaultPermissions and assigns it to the DefaultPermissions field.
func (o *UpdateSubscriptionVM) SetDefaultPermissions(v DefaultPermissions) {
	o.DefaultPermissions = &v
}

func (o UpdateSubscriptionVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.DefaultPermissions != nil {
		toSerialize["defaultPermissions"] = o.DefaultPermissions
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateSubscriptionVM struct {
	value *UpdateSubscriptionVM
	isSet bool
}

func (v NullableUpdateSubscriptionVM) Get() *UpdateSubscriptionVM {
	return v.value
}

func (v *NullableUpdateSubscriptionVM) Set(val *UpdateSubscriptionVM) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionVM) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionVM(val *UpdateSubscriptionVM) *NullableUpdateSubscriptionVM {
	return &NullableUpdateSubscriptionVM{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

