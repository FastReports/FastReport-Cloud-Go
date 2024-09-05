/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the CreateSubscriptionPeriodVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSubscriptionPeriodVM{}

// CreateSubscriptionPeriodVM struct for CreateSubscriptionPeriodVM
type CreateSubscriptionPeriodVM struct {
	CloudBaseVM
	PlanId string `json:"planId"`
	Start NullableTime `json:"start,omitempty"`
	End NullableTime `json:"end,omitempty"`
	T string `json:"$t"`
}

type _CreateSubscriptionPeriodVM CreateSubscriptionPeriodVM

// NewCreateSubscriptionPeriodVM instantiates a new CreateSubscriptionPeriodVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSubscriptionPeriodVM(planId string, t string) *CreateSubscriptionPeriodVM {
	this := CreateSubscriptionPeriodVM{}
	this.T = t
	return &this
}

// NewCreateSubscriptionPeriodVMWithDefaults instantiates a new CreateSubscriptionPeriodVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSubscriptionPeriodVMWithDefaults() *CreateSubscriptionPeriodVM {
	this := CreateSubscriptionPeriodVM{}
	return &this
}

// GetPlanId returns the PlanId field value
func (o *CreateSubscriptionPeriodVM) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionPeriodVM) GetPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *CreateSubscriptionPeriodVM) SetPlanId(v string) {
	o.PlanId = v
}

// GetStart returns the Start field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSubscriptionPeriodVM) GetStart() time.Time {
	if o == nil || IsNil(o.Start.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Start.Get()
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSubscriptionPeriodVM) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Start.Get(), o.Start.IsSet()
}

// HasStart returns a boolean if a field has been set.
func (o *CreateSubscriptionPeriodVM) HasStart() bool {
	if o != nil && o.Start.IsSet() {
		return true
	}

	return false
}

// SetStart gets a reference to the given NullableTime and assigns it to the Start field.
func (o *CreateSubscriptionPeriodVM) SetStart(v time.Time) {
	o.Start.Set(&v)
}
// SetStartNil sets the value for Start to be an explicit nil
func (o *CreateSubscriptionPeriodVM) SetStartNil() {
	o.Start.Set(nil)
}

// UnsetStart ensures that no value is present for Start, not even an explicit nil
func (o *CreateSubscriptionPeriodVM) UnsetStart() {
	o.Start.Unset()
}

// GetEnd returns the End field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateSubscriptionPeriodVM) GetEnd() time.Time {
	if o == nil || IsNil(o.End.Get()) {
		var ret time.Time
		return ret
	}
	return *o.End.Get()
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateSubscriptionPeriodVM) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.End.Get(), o.End.IsSet()
}

// HasEnd returns a boolean if a field has been set.
func (o *CreateSubscriptionPeriodVM) HasEnd() bool {
	if o != nil && o.End.IsSet() {
		return true
	}

	return false
}

// SetEnd gets a reference to the given NullableTime and assigns it to the End field.
func (o *CreateSubscriptionPeriodVM) SetEnd(v time.Time) {
	o.End.Set(&v)
}
// SetEndNil sets the value for End to be an explicit nil
func (o *CreateSubscriptionPeriodVM) SetEndNil() {
	o.End.Set(nil)
}

// UnsetEnd ensures that no value is present for End, not even an explicit nil
func (o *CreateSubscriptionPeriodVM) UnsetEnd() {
	o.End.Unset()
}

// GetT returns the T field value
func (o *CreateSubscriptionPeriodVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *CreateSubscriptionPeriodVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *CreateSubscriptionPeriodVM) SetT(v string) {
	o.T = v
}

func (o CreateSubscriptionPeriodVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSubscriptionPeriodVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	toSerialize["planId"] = o.PlanId
	if o.Start.IsSet() {
		toSerialize["start"] = o.Start.Get()
	}
	if o.End.IsSet() {
		toSerialize["end"] = o.End.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *CreateSubscriptionPeriodVM) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"planId",
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

	varCreateSubscriptionPeriodVM := _CreateSubscriptionPeriodVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSubscriptionPeriodVM)

	if err != nil {
		return err
	}

	*o = CreateSubscriptionPeriodVM(varCreateSubscriptionPeriodVM)

	return err
}

type NullableCreateSubscriptionPeriodVM struct {
	value *CreateSubscriptionPeriodVM
	isSet bool
}

func (v NullableCreateSubscriptionPeriodVM) Get() *CreateSubscriptionPeriodVM {
	return v.value
}

func (v *NullableCreateSubscriptionPeriodVM) Set(val *CreateSubscriptionPeriodVM) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSubscriptionPeriodVM) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSubscriptionPeriodVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSubscriptionPeriodVM(val *CreateSubscriptionPeriodVM) *NullableCreateSubscriptionPeriodVM {
	return &NullableCreateSubscriptionPeriodVM{value: val, isSet: true}
}

func (v NullableCreateSubscriptionPeriodVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSubscriptionPeriodVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


