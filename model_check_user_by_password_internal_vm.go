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

// checks if the CheckUserByPasswordInternalVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckUserByPasswordInternalVM{}

// CheckUserByPasswordInternalVM struct for CheckUserByPasswordInternalVM
type CheckUserByPasswordInternalVM struct {
	CloudBaseVM
	Email NullableString `json:"email,omitempty"`
	Password NullableString `json:"password,omitempty"`
	T string `json:"$t"`
}

type _CheckUserByPasswordInternalVM CheckUserByPasswordInternalVM

// NewCheckUserByPasswordInternalVM instantiates a new CheckUserByPasswordInternalVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckUserByPasswordInternalVM(t string) *CheckUserByPasswordInternalVM {
	this := CheckUserByPasswordInternalVM{}
	this.T = t
	return &this
}

// NewCheckUserByPasswordInternalVMWithDefaults instantiates a new CheckUserByPasswordInternalVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckUserByPasswordInternalVMWithDefaults() *CheckUserByPasswordInternalVM {
	this := CheckUserByPasswordInternalVM{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckUserByPasswordInternalVM) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckUserByPasswordInternalVM) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *CheckUserByPasswordInternalVM) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *CheckUserByPasswordInternalVM) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *CheckUserByPasswordInternalVM) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *CheckUserByPasswordInternalVM) UnsetEmail() {
	o.Email.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CheckUserByPasswordInternalVM) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CheckUserByPasswordInternalVM) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CheckUserByPasswordInternalVM) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CheckUserByPasswordInternalVM) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CheckUserByPasswordInternalVM) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CheckUserByPasswordInternalVM) UnsetPassword() {
	o.Password.Unset()
}

// GetT returns the T field value
func (o *CheckUserByPasswordInternalVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CheckUserByPasswordInternalVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CheckUserByPasswordInternalVM) SetT(v string) {
	o.T = v
}

func (o CheckUserByPasswordInternalVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckUserByPasswordInternalVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CheckUserByPasswordInternalVM) UnmarshalJSON(data []byte) (err error) {
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

	varCheckUserByPasswordInternalVM := _CheckUserByPasswordInternalVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCheckUserByPasswordInternalVM)

	if err != nil {
		return err
	}

	*o = CheckUserByPasswordInternalVM(varCheckUserByPasswordInternalVM)

	return err
}

type NullableCheckUserByPasswordInternalVM struct {
	value *CheckUserByPasswordInternalVM
	isSet bool
}

func (v NullableCheckUserByPasswordInternalVM) Get() *CheckUserByPasswordInternalVM {
	return v.value
}

func (v *NullableCheckUserByPasswordInternalVM) Set(val *CheckUserByPasswordInternalVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckUserByPasswordInternalVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckUserByPasswordInternalVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckUserByPasswordInternalVM(val *CheckUserByPasswordInternalVM) *NullableCheckUserByPasswordInternalVM {
	return &NullableCheckUserByPasswordInternalVM{value: val, isSet: true}
}

func (v NullableCheckUserByPasswordInternalVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckUserByPasswordInternalVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


