/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SubscriptionUserVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionUserVM{}

// SubscriptionUserVM struct for SubscriptionUserVM
type SubscriptionUserVM struct {
	CloudBaseVM
	UserId NullableString `json:"userId,omitempty"`
	Groups []GroupVM `json:"groups,omitempty"`
	T string `json:"$t"`
}

type _SubscriptionUserVM SubscriptionUserVM

// NewSubscriptionUserVM instantiates a new SubscriptionUserVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionUserVM(t string) *SubscriptionUserVM {
	this := SubscriptionUserVM{}
	this.T = t
	return &this
}

// NewSubscriptionUserVMWithDefaults instantiates a new SubscriptionUserVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionUserVMWithDefaults() *SubscriptionUserVM {
	this := SubscriptionUserVM{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionUserVM) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionUserVM) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *SubscriptionUserVM) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *SubscriptionUserVM) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *SubscriptionUserVM) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *SubscriptionUserVM) UnsetUserId() {
	o.UserId.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionUserVM) GetGroups() []GroupVM {
	if o == nil {
		var ret []GroupVM
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionUserVM) GetGroupsOk() ([]GroupVM, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *SubscriptionUserVM) HasGroups() bool {
	if o != nil && IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupVM and assigns it to the Groups field.
func (o *SubscriptionUserVM) SetGroups(v []GroupVM) {
	o.Groups = v
}

// GetT returns the T field value
func (o *SubscriptionUserVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *SubscriptionUserVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *SubscriptionUserVM) SetT(v string) {
	o.T = v
}

func (o SubscriptionUserVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionUserVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *SubscriptionUserVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"$t",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubscriptionUserVM := _SubscriptionUserVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionUserVM)

	if err != nil {
		return err
	}

	*o = SubscriptionUserVM(varSubscriptionUserVM)

	return err
}

type NullableSubscriptionUserVM struct {
	value *SubscriptionUserVM
	isSet bool
}

func (v NullableSubscriptionUserVM) Get() *SubscriptionUserVM {
	return v.value
}

func (v *NullableSubscriptionUserVM) Set(val *SubscriptionUserVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionUserVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionUserVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionUserVM(val *SubscriptionUserVM) *NullableSubscriptionUserVM {
	return &NullableSubscriptionUserVM{value: val, isSet: true}
}

func (v NullableSubscriptionUserVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionUserVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


