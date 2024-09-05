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

// checks if the SubscriptionInviteVM type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionInviteVM{}

// SubscriptionInviteVM struct for SubscriptionInviteVM
type SubscriptionInviteVM struct {
	CloudBaseVM
	Usages *int64 `json:"usages,omitempty"`
	Durable *bool `json:"durable,omitempty"`
	AccessToken NullableString `json:"accessToken,omitempty"`
	ExpiredDate *time.Time `json:"expiredDate,omitempty"`
	AddedUsers []InvitedUser `json:"addedUsers,omitempty"`
	CreatorUserId NullableString `json:"creatorUserId,omitempty"`
	T string `json:"$t"`
}

type _SubscriptionInviteVM SubscriptionInviteVM

// NewSubscriptionInviteVM instantiates a new SubscriptionInviteVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionInviteVM(t string) *SubscriptionInviteVM {
	this := SubscriptionInviteVM{}
	this.T = t
	return &this
}

// NewSubscriptionInviteVMWithDefaults instantiates a new SubscriptionInviteVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionInviteVMWithDefaults() *SubscriptionInviteVM {
	this := SubscriptionInviteVM{}
	return &this
}

// GetUsages returns the Usages field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetUsages() int64 {
	if o == nil || IsNil(o.Usages) {
		var ret int64
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetUsagesOk() (*int64, bool) {
	if o == nil || IsNil(o.Usages) {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasUsages() bool {
	if o != nil && !IsNil(o.Usages) {
		return true
	}

	return false
}

// SetUsages gets a reference to the given int64 and assigns it to the Usages field.
func (o *SubscriptionInviteVM) SetUsages(v int64) {
	o.Usages = &v
}

// GetDurable returns the Durable field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetDurable() bool {
	if o == nil || IsNil(o.Durable) {
		var ret bool
		return ret
	}
	return *o.Durable
}

// GetDurableOk returns a tuple with the Durable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetDurableOk() (*bool, bool) {
	if o == nil || IsNil(o.Durable) {
		return nil, false
	}
	return o.Durable, true
}

// HasDurable returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasDurable() bool {
	if o != nil && !IsNil(o.Durable) {
		return true
	}

	return false
}

// SetDurable gets a reference to the given bool and assigns it to the Durable field.
func (o *SubscriptionInviteVM) SetDurable(v bool) {
	o.Durable = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionInviteVM) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionInviteVM) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *SubscriptionInviteVM) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *SubscriptionInviteVM) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *SubscriptionInviteVM) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetExpiredDate returns the ExpiredDate field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetExpiredDate() time.Time {
	if o == nil || IsNil(o.ExpiredDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredDate
}

// GetExpiredDateOk returns a tuple with the ExpiredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetExpiredDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiredDate) {
		return nil, false
	}
	return o.ExpiredDate, true
}

// HasExpiredDate returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasExpiredDate() bool {
	if o != nil && !IsNil(o.ExpiredDate) {
		return true
	}

	return false
}

// SetExpiredDate gets a reference to the given time.Time and assigns it to the ExpiredDate field.
func (o *SubscriptionInviteVM) SetExpiredDate(v time.Time) {
	o.ExpiredDate = &v
}

// GetAddedUsers returns the AddedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionInviteVM) GetAddedUsers() []InvitedUser {
	if o == nil {
		var ret []InvitedUser
		return ret
	}
	return o.AddedUsers
}

// GetAddedUsersOk returns a tuple with the AddedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionInviteVM) GetAddedUsersOk() ([]InvitedUser, bool) {
	if o == nil || IsNil(o.AddedUsers) {
		return nil, false
	}
	return o.AddedUsers, true
}

// HasAddedUsers returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasAddedUsers() bool {
	if o != nil && IsNil(o.AddedUsers) {
		return true
	}

	return false
}

// SetAddedUsers gets a reference to the given []InvitedUser and assigns it to the AddedUsers field.
func (o *SubscriptionInviteVM) SetAddedUsers(v []InvitedUser) {
	o.AddedUsers = v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubscriptionInviteVM) GetCreatorUserId() string {
	if o == nil || IsNil(o.CreatorUserId.Get()) {
		var ret string
		return ret
	}
	return *o.CreatorUserId.Get()
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionInviteVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorUserId.Get(), o.CreatorUserId.IsSet()
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId.IsSet() {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given NullableString and assigns it to the CreatorUserId field.
func (o *SubscriptionInviteVM) SetCreatorUserId(v string) {
	o.CreatorUserId.Set(&v)
}
// SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil
func (o *SubscriptionInviteVM) SetCreatorUserIdNil() {
	o.CreatorUserId.Set(nil)
}

// UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
func (o *SubscriptionInviteVM) UnsetCreatorUserId() {
	o.CreatorUserId.Unset()
}

// GetT returns the T field value
func (o *SubscriptionInviteVM) GetT() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.T
}

// GetTOk returns a tuple with the T field value
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetTOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.T, true
}

// SetT sets field value
func (o *SubscriptionInviteVM) SetT(v string) {
	o.T = v
}

func (o SubscriptionInviteVM) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionInviteVM) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCloudBaseVM, errCloudBaseVM := json.Marshal(o.CloudBaseVM)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	errCloudBaseVM = json.Unmarshal([]byte(serializedCloudBaseVM), &toSerialize)
	if errCloudBaseVM != nil {
		return map[string]interface{}{}, errCloudBaseVM
	}
	if !IsNil(o.Usages) {
		toSerialize["usages"] = o.Usages
	}
	if !IsNil(o.Durable) {
		toSerialize["durable"] = o.Durable
	}
	if o.AccessToken.IsSet() {
		toSerialize["accessToken"] = o.AccessToken.Get()
	}
	if !IsNil(o.ExpiredDate) {
		toSerialize["expiredDate"] = o.ExpiredDate
	}
	if o.AddedUsers != nil {
		toSerialize["addedUsers"] = o.AddedUsers
	}
	if o.CreatorUserId.IsSet() {
		toSerialize["creatorUserId"] = o.CreatorUserId.Get()
	}
	toSerialize["$t"] = o.T
	return toSerialize, nil
}

func (o *SubscriptionInviteVM) UnmarshalJSON(data []byte) (err error) {
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

	varSubscriptionInviteVM := _SubscriptionInviteVM{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionInviteVM)

	if err != nil {
		return err
	}

	*o = SubscriptionInviteVM(varSubscriptionInviteVM)

	return err
}

type NullableSubscriptionInviteVM struct {
	value *SubscriptionInviteVM
	isSet bool
}

func (v NullableSubscriptionInviteVM) Get() *SubscriptionInviteVM {
	return v.value
}

func (v *NullableSubscriptionInviteVM) Set(val *SubscriptionInviteVM) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionInviteVM) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionInviteVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionInviteVM(val *SubscriptionInviteVM) *NullableSubscriptionInviteVM {
	return &NullableSubscriptionInviteVM{value: val, isSet: true}
}

func (v NullableSubscriptionInviteVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionInviteVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


