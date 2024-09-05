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

// checks if the SubscriptionPeriodVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPeriodVM{}

// SubscriptionPeriodVM struct for SubscriptionPeriodVM
type SubscriptionPeriodVM struct {
	CloudBaseVM
	StartTime *time.Time `json:"startTime,omitempty"`
	EndTime *time.Time `json:"endTime,omitempty"`
	Plan *SubscriptionPlanVM `json:"plan,omitempty"`
	T string `json:"$t"`
}

type _SubscriptionPeriodVM SubscriptionPeriodVM

// NewSubscriptionPeriodVM instantiates a new SubscriptionPeriodVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPeriodVM(t string) *SubscriptionPeriodVM {
	this := SubscriptionPeriodVM{}
	this.T = t
	return &this
}

// NewSubscriptionPeriodVMWithDefaults instantiates a new SubscriptionPeriodVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPeriodVMWithDefaults() *SubscriptionPeriodVM {
	this := SubscriptionPeriodVM{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *SubscriptionPeriodVM) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPeriodVM) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *SubscriptionPeriodVM) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *SubscriptionPeriodVM) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *SubscriptionPeriodVM) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPeriodVM) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *SubscriptionPeriodVM) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *SubscriptionPeriodVM) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetPlan returns the Plan field value if set, zero value otherwise.
func (o *SubscriptionPeriodVM) GetPlan() SubscriptionPlanVM {
	if o == nil || IsNil(o.Plan) {
		var ret SubscriptionPlanVM
		return ret
	}
	return *o.Plan
}

// GetPlanOk returns a tuple with the Plan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPeriodVM) GetPlanOk() (*SubscriptionPlanVM, bool) {
	if o == nil || IsNil(o.Plan) {
		return nil, false
	}
	return o.Plan, true
}

// HasPlan returns a boolean if a field has been set.
func (o *SubscriptionPeriodVM) HasPlan() bool {
	if o != nil && !IsNil(o.Plan) {
		return true
	}

	return false
}

// SetPlan gets a reference to the given SubscriptionPlanVM and assigns it to the Plan field.
func (o *SubscriptionPeriodVM) SetPlan(v SubscriptionPlanVM) {
	o.Plan = &v
}

// GetT returns the T field value
func (o *SubscriptionPeriodVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPeriodVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *SubscriptionPeriodVM) SetT(v string) {
	o.T = v
}

func (o SubscriptionPeriodVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPeriodVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.Plan) {
		toSerialize["plan"] = o.Plan
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *SubscriptionPeriodVM) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionPeriodVM := _SubscriptionPeriodVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionPeriodVM)

	if err != nil {
		return err
	}

	*o = SubscriptionPeriodVM(varSubscriptionPeriodVM)

	return err
}

type NullableSubscriptionPeriodVM struct {
	value *SubscriptionPeriodVM
	isSet bool
}

func (v NullableSubscriptionPeriodVM) Get() *SubscriptionPeriodVM {
	return v.value
}

func (v *NullableSubscriptionPeriodVM) Set(val *SubscriptionPeriodVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPeriodVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPeriodVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPeriodVM(val *SubscriptionPeriodVM) *NullableSubscriptionPeriodVM {
	return &NullableSubscriptionPeriodVM{value: val, isSet: true}
}

func (v NullableSubscriptionPeriodVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPeriodVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


