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

// checks if the ContactsVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactsVM{}

// ContactsVM struct for ContactsVM
type ContactsVM struct {
	Contacts []ContactVM `json:"contacts,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	Take *int32 `json:"take,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewContactsVM instantiates a new ContactsVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactsVM() *ContactsVM {
	this := ContactsVM{}
	return &this
}

// NewContactsVMWithDefaults instantiates a new ContactsVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactsVMWithDefaults() *ContactsVM {
	this := ContactsVM{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContactsVM) GetContacts() []ContactVM {
	if o == nil {
		var ret []ContactVM
		return ret
	}
	return o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContactsVM) GetContactsOk() ([]ContactVM, bool) {
	if o == nil || IsNil(o.Contacts) {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *ContactsVM) HasContacts() bool {
	if o != nil && IsNil(o.Contacts) {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []ContactVM and assigns it to the Contacts field.
func (o *ContactsVM) SetContacts(v []ContactVM) {
	o.Contacts = v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ContactsVM) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsVM) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ContactsVM) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ContactsVM) SetSkip(v int32) {
	o.Skip = &v
}

// GetTake returns the Take field value if set, zero value otherwise.
func (o *ContactsVM) GetTake() int32 {
	if o == nil || IsNil(o.Take) {
		var ret int32
		return ret
	}
	return *o.Take
}

// GetTakeOk returns a tuple with the Take field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsVM) GetTakeOk() (*int32, bool) {
	if o == nil || IsNil(o.Take) {
		return nil, false
	}
	return o.Take, true
}

// HasTake returns a boolean if a field has been set.
func (o *ContactsVM) HasTake() bool {
	if o != nil && !IsNil(o.Take) {
		return true
	}

	return false
}

// SetTake gets a reference to the given int32 and assigns it to the Take field.
func (o *ContactsVM) SetTake(v int32) {
	o.Take = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ContactsVM) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactsVM) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ContactsVM) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ContactsVM) SetCount(v int64) {
	o.Count = &v
}

func (o ContactsVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactsVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.Take) {
		toSerialize["take"] = o.Take
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableContactsVM struct {
	value *ContactsVM
	isSet bool
}

func (v NullableContactsVM) Get() *ContactsVM {
	return v.value
}

func (v *NullableContactsVM) Set(val *ContactsVM) {
	v.value = val
	v.isSet = true
}

func (v NullableContactsVM) IsSet() bool {
	return v.isSet
}

func (v *NullableContactsVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactsVM(val *ContactsVM) *NullableContactsVM {
	return &NullableContactsVM{value: val, isSet: true}
}

func (v NullableContactsVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactsVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


