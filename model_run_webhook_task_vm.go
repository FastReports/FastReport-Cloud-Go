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

// checks if the RunWebhookTaskVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunWebhookTaskVM{}

// RunWebhookTaskVM struct for RunWebhookTaskVM
type RunWebhookTaskVM struct {
	Headers map[string]string `json:"headers,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewRunWebhookTaskVM instantiates a new RunWebhookTaskVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunWebhookTaskVM(t string) *RunWebhookTaskVM {
	this := RunWebhookTaskVM{}
	this.T = t
	return &this
}

// NewRunWebhookTaskVMWithDefaults instantiates a new RunWebhookTaskVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWebhookTaskVMWithDefaults() *RunWebhookTaskVM {
	this := RunWebhookTaskVM{}
	return &this
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunWebhookTaskVM) GetHeaders() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunWebhookTaskVM) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *RunWebhookTaskVM) HasHeaders() bool {
	if o != nil && IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *RunWebhookTaskVM) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunWebhookTaskVM) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunWebhookTaskVM) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *RunWebhookTaskVM) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *RunWebhookTaskVM) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *RunWebhookTaskVM) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *RunWebhookTaskVM) UnsetUrl() {
	o.Url.Unset()
}

func (o RunWebhookTaskVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunWebhookTaskVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullableRunWebhookTaskVM struct {
	value *RunWebhookTaskVM
	isSet bool
}

func (v NullableRunWebhookTaskVM) Get() *RunWebhookTaskVM {
	return v.value
}

func (v *NullableRunWebhookTaskVM) Set(val *RunWebhookTaskVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRunWebhookTaskVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRunWebhookTaskVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunWebhookTaskVM(val *RunWebhookTaskVM) *NullableRunWebhookTaskVM {
	return &NullableRunWebhookTaskVM{value: val, isSet: true}
}

func (v NullableRunWebhookTaskVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunWebhookTaskVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


