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
)

// RegisterUserVM struct for RegisterUserVM
type RegisterUserVM struct {
	Id *string `json:"id,omitempty"`
	Subscriptions *[]string `json:"subscriptions,omitempty"`
	Groups *[]string `json:"groups,omitempty"`
	AdminPermission *AdminPermission `json:"adminPermission,omitempty"`
	Name *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	IsAdmin *bool `json:"isAdmin,omitempty"`
	Provider *string `json:"provider,omitempty"`
}

// NewRegisterUserVM instantiates a new RegisterUserVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterUserVM() *RegisterUserVM {
	this := RegisterUserVM{}
	return &this
}

// NewRegisterUserVMWithDefaults instantiates a new RegisterUserVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterUserVMWithDefaults() *RegisterUserVM {
	this := RegisterUserVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegisterUserVM) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegisterUserVM) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegisterUserVM) SetId(v string) {
	o.Id = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *RegisterUserVM) GetSubscriptions() []string {
	if o == nil || o.Subscriptions == nil {
		var ret []string
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetSubscriptionsOk() (*[]string, bool) {
	if o == nil || o.Subscriptions == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *RegisterUserVM) HasSubscriptions() bool {
	if o != nil && o.Subscriptions != nil {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given []string and assigns it to the Subscriptions field.
func (o *RegisterUserVM) SetSubscriptions(v []string) {
	o.Subscriptions = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *RegisterUserVM) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *RegisterUserVM) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *RegisterUserVM) SetGroups(v []string) {
	o.Groups = &v
}

// GetAdminPermission returns the AdminPermission field value if set, zero value otherwise.
func (o *RegisterUserVM) GetAdminPermission() AdminPermission {
	if o == nil || o.AdminPermission == nil {
		var ret AdminPermission
		return ret
	}
	return *o.AdminPermission
}

// GetAdminPermissionOk returns a tuple with the AdminPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetAdminPermissionOk() (*AdminPermission, bool) {
	if o == nil || o.AdminPermission == nil {
		return nil, false
	}
	return o.AdminPermission, true
}

// HasAdminPermission returns a boolean if a field has been set.
func (o *RegisterUserVM) HasAdminPermission() bool {
	if o != nil && o.AdminPermission != nil {
		return true
	}

	return false
}

// SetAdminPermission gets a reference to the given AdminPermission and assigns it to the AdminPermission field.
func (o *RegisterUserVM) SetAdminPermission(v AdminPermission) {
	o.AdminPermission = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegisterUserVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegisterUserVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegisterUserVM) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *RegisterUserVM) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *RegisterUserVM) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *RegisterUserVM) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RegisterUserVM) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RegisterUserVM) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RegisterUserVM) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *RegisterUserVM) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *RegisterUserVM) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *RegisterUserVM) SetPassword(v string) {
	o.Password = &v
}

// GetIsAdmin returns the IsAdmin field value if set, zero value otherwise.
func (o *RegisterUserVM) GetIsAdmin() bool {
	if o == nil || o.IsAdmin == nil {
		var ret bool
		return ret
	}
	return *o.IsAdmin
}

// GetIsAdminOk returns a tuple with the IsAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetIsAdminOk() (*bool, bool) {
	if o == nil || o.IsAdmin == nil {
		return nil, false
	}
	return o.IsAdmin, true
}

// HasIsAdmin returns a boolean if a field has been set.
func (o *RegisterUserVM) HasIsAdmin() bool {
	if o != nil && o.IsAdmin != nil {
		return true
	}

	return false
}

// SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.
func (o *RegisterUserVM) SetIsAdmin(v bool) {
	o.IsAdmin = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *RegisterUserVM) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserVM) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *RegisterUserVM) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *RegisterUserVM) SetProvider(v string) {
	o.Provider = &v
}

func (o RegisterUserVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Subscriptions != nil {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.AdminPermission != nil {
		toSerialize["adminPermission"] = o.AdminPermission
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.IsAdmin != nil {
		toSerialize["isAdmin"] = o.IsAdmin
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterUserVM struct {
	value *RegisterUserVM
	isSet bool
}

func (v NullableRegisterUserVM) Get() *RegisterUserVM {
	return v.value
}

func (v *NullableRegisterUserVM) Set(val *RegisterUserVM) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterUserVM) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterUserVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterUserVM(val *RegisterUserVM) *NullableRegisterUserVM {
	return &NullableRegisterUserVM{value: val, isSet: true}
}

func (v NullableRegisterUserVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterUserVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


