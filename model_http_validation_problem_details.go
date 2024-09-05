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

// checks if the HttpValidationProblemDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpValidationProblemDetails{}

// HttpValidationProblemDetails struct for HttpValidationProblemDetails
type HttpValidationProblemDetails struct {
	Errors *map[string][]string `json:"errors,omitempty"`
	Type NullableString `json:"type,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Status NullableInt32 `json:"status,omitempty"`
	Detail NullableString `json:"detail,omitempty"`
	Instance NullableString `json:"instance,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HttpValidationProblemDetails HttpValidationProblemDetails

// NewHttpValidationProblemDetails instantiates a new HttpValidationProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpValidationProblemDetails() *HttpValidationProblemDetails {
	this := HttpValidationProblemDetails{}
	return &this
}

// NewHttpValidationProblemDetailsWithDefaults instantiates a new HttpValidationProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpValidationProblemDetailsWithDefaults() *HttpValidationProblemDetails {
	this := HttpValidationProblemDetails{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *HttpValidationProblemDetails) GetErrors() map[string][]string {
	if o == nil || IsNil(o.Errors) {
		var ret map[string][]string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpValidationProblemDetails) GetErrorsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *HttpValidationProblemDetails) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string][]string and assigns it to the Errors field.
func (o *HttpValidationProblemDetails) SetErrors(v map[string][]string) {
	o.Errors = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpValidationProblemDetails) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpValidationProblemDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *HttpValidationProblemDetails) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *HttpValidationProblemDetails) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *HttpValidationProblemDetails) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *HttpValidationProblemDetails) UnsetType() {
	o.Type.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpValidationProblemDetails) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpValidationProblemDetails) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *HttpValidationProblemDetails) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *HttpValidationProblemDetails) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *HttpValidationProblemDetails) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *HttpValidationProblemDetails) UnsetTitle() {
	o.Title.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpValidationProblemDetails) GetStatus() int32 {
	if o == nil || IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpValidationProblemDetails) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *HttpValidationProblemDetails) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableInt32 and assigns it to the Status field.
func (o *HttpValidationProblemDetails) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *HttpValidationProblemDetails) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *HttpValidationProblemDetails) UnsetStatus() {
	o.Status.Unset()
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpValidationProblemDetails) GetDetail() string {
	if o == nil || IsNil(o.Detail.Get()) {
		var ret string
		return ret
	}
	return *o.Detail.Get()
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpValidationProblemDetails) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Detail.Get(), o.Detail.IsSet()
}

// HasDetail returns a boolean if a field has been set.
func (o *HttpValidationProblemDetails) HasDetail() bool {
	if o != nil && o.Detail.IsSet() {
		return true
	}

	return false
}

// SetDetail gets a reference to the given NullableString and assigns it to the Detail field.
func (o *HttpValidationProblemDetails) SetDetail(v string) {
	o.Detail.Set(&v)
}
// SetDetailNil sets the value for Detail to be an explicit nil
func (o *HttpValidationProblemDetails) SetDetailNil() {
	o.Detail.Set(nil)
}

// UnsetDetail ensures that no value is present for Detail, not even an explicit nil
func (o *HttpValidationProblemDetails) UnsetDetail() {
	o.Detail.Unset()
}

// GetInstance returns the Instance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HttpValidationProblemDetails) GetInstance() string {
	if o == nil || IsNil(o.Instance.Get()) {
		var ret string
		return ret
	}
	return *o.Instance.Get()
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HttpValidationProblemDetails) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Instance.Get(), o.Instance.IsSet()
}

// HasInstance returns a boolean if a field has been set.
func (o *HttpValidationProblemDetails) HasInstance() bool {
	if o != nil && o.Instance.IsSet() {
		return true
	}

	return false
}

// SetInstance gets a reference to the given NullableString and assigns it to the Instance field.
func (o *HttpValidationProblemDetails) SetInstance(v string) {
	o.Instance.Set(&v)
}
// SetInstanceNil sets the value for Instance to be an explicit nil
func (o *HttpValidationProblemDetails) SetInstanceNil() {
	o.Instance.Set(nil)
}

// UnsetInstance ensures that no value is present for Instance, not even an explicit nil
func (o *HttpValidationProblemDetails) UnsetInstance() {
	o.Instance.Unset()
}

func (o HttpValidationProblemDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpValidationProblemDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Detail.IsSet() {
		toSerialize["detail"] = o.Detail.Get()
	}
	if o.Instance.IsSet() {
		toSerialize["instance"] = o.Instance.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *HttpValidationProblemDetails) UnmarshalJSON(data []byte) (err error) {
	varHttpValidationProblemDetails := _HttpValidationProblemDetails{}

	err = json.Unmarshal(data, &varHttpValidationProblemDetails)

	if err != nil {
		return err
	}

	*o = HttpValidationProblemDetails(varHttpValidationProblemDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errors")
		delete(additionalProperties, "type")
		delete(additionalProperties, "title")
		delete(additionalProperties, "status")
		delete(additionalProperties, "detail")
		delete(additionalProperties, "instance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHttpValidationProblemDetails struct {
	value *HttpValidationProblemDetails
	isSet bool
}

func (v NullableHttpValidationProblemDetails) Get() *HttpValidationProblemDetails {
	return v.value
}

func (v *NullableHttpValidationProblemDetails) Set(val *HttpValidationProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpValidationProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpValidationProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpValidationProblemDetails(val *HttpValidationProblemDetails) *NullableHttpValidationProblemDetails {
	return &NullableHttpValidationProblemDetails{value: val, isSet: true}
}

func (v NullableHttpValidationProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpValidationProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


