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
	"time"
)

// SubscriptionInviteVM struct for SubscriptionInviteVM
type SubscriptionInviteVM struct {
	Usages *int64 `json:"usages,omitempty"`
	Durable *bool `json:"durable,omitempty"`
	AccessToken *string `json:"accessToken,omitempty"`
	ExpiredDate *time.Time `json:"expiredDate,omitempty"`
	AddedUsers *[]InvitedUser `json:"addedUsers,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty"`
}

// NewSubscriptionInviteVM instantiates a new SubscriptionInviteVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionInviteVM() *SubscriptionInviteVM {
	this := SubscriptionInviteVM{}
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
	if o == nil || o.Usages == nil {
		var ret int64
		return ret
	}
	return *o.Usages
}

// GetUsagesOk returns a tuple with the Usages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetUsagesOk() (*int64, bool) {
	if o == nil || o.Usages == nil {
		return nil, false
	}
	return o.Usages, true
}

// HasUsages returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasUsages() bool {
	if o != nil && o.Usages != nil {
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
	if o == nil || o.Durable == nil {
		var ret bool
		return ret
	}
	return *o.Durable
}

// GetDurableOk returns a tuple with the Durable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetDurableOk() (*bool, bool) {
	if o == nil || o.Durable == nil {
		return nil, false
	}
	return o.Durable, true
}

// HasDurable returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasDurable() bool {
	if o != nil && o.Durable != nil {
		return true
	}

	return false
}

// SetDurable gets a reference to the given bool and assigns it to the Durable field.
func (o *SubscriptionInviteVM) SetDurable(v bool) {
	o.Durable = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *SubscriptionInviteVM) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiredDate returns the ExpiredDate field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetExpiredDate() time.Time {
	if o == nil || o.ExpiredDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiredDate
}

// GetExpiredDateOk returns a tuple with the ExpiredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetExpiredDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiredDate == nil {
		return nil, false
	}
	return o.ExpiredDate, true
}

// HasExpiredDate returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasExpiredDate() bool {
	if o != nil && o.ExpiredDate != nil {
		return true
	}

	return false
}

// SetExpiredDate gets a reference to the given time.Time and assigns it to the ExpiredDate field.
func (o *SubscriptionInviteVM) SetExpiredDate(v time.Time) {
	o.ExpiredDate = &v
}

// GetAddedUsers returns the AddedUsers field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetAddedUsers() []InvitedUser {
	if o == nil || o.AddedUsers == nil {
		var ret []InvitedUser
		return ret
	}
	return *o.AddedUsers
}

// GetAddedUsersOk returns a tuple with the AddedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetAddedUsersOk() (*[]InvitedUser, bool) {
	if o == nil || o.AddedUsers == nil {
		return nil, false
	}
	return o.AddedUsers, true
}

// HasAddedUsers returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasAddedUsers() bool {
	if o != nil && o.AddedUsers != nil {
		return true
	}

	return false
}

// SetAddedUsers gets a reference to the given []InvitedUser and assigns it to the AddedUsers field.
func (o *SubscriptionInviteVM) SetAddedUsers(v []InvitedUser) {
	o.AddedUsers = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise.
func (o *SubscriptionInviteVM) GetCreatorUserId() string {
	if o == nil || o.CreatorUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatorUserId
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionInviteVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil || o.CreatorUserId == nil {
		return nil, false
	}
	return o.CreatorUserId, true
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *SubscriptionInviteVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId != nil {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given string and assigns it to the CreatorUserId field.
func (o *SubscriptionInviteVM) SetCreatorUserId(v string) {
	o.CreatorUserId = &v
}

func (o SubscriptionInviteVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usages != nil {
		toSerialize["usages"] = o.Usages
	}
	if o.Durable != nil {
		toSerialize["durable"] = o.Durable
	}
	if o.AccessToken != nil {
		toSerialize["accessToken"] = o.AccessToken
	}
	if o.ExpiredDate != nil {
		toSerialize["expiredDate"] = o.ExpiredDate
	}
	if o.AddedUsers != nil {
		toSerialize["addedUsers"] = o.AddedUsers
	}
	if o.CreatorUserId != nil {
		toSerialize["creatorUserId"] = o.CreatorUserId
	}
	return json.Marshal(toSerialize)
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


