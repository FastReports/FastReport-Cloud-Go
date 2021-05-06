/*
 * FastReport Cloud
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package FastReport.Cloud.SDK

import (
	"encoding/json"
	"time"
)

// DataSourceVM struct for DataSourceVM
type DataSourceVM struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ConnectionType *string `json:"connectionType,omitempty"`
	ConnectionString *string `json:"connectionString,omitempty"`
	DataStructure *string `json:"dataStructure,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	EditedTime *time.Time `json:"editedTime,omitempty"`
	EditorUserId *string `json:"editorUserId,omitempty"`
	CreatedTime *time.Time `json:"createdTime,omitempty"`
	CreatorUserId *string `json:"creatorUserId,omitempty"`
	IsConnected *bool `json:"isConnected,omitempty"`
}

// NewDataSourceVM instantiates a new DataSourceVM object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataSourceVM() *DataSourceVM {
	this := DataSourceVM{}
	return &this
}

// NewDataSourceVMWithDefaults instantiates a new DataSourceVM object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceVMWithDefaults() *DataSourceVM {
	this := DataSourceVM{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataSourceVM) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataSourceVM) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataSourceVM) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DataSourceVM) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DataSourceVM) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DataSourceVM) SetName(v string) {
	o.Name = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *DataSourceVM) GetConnectionType() string {
	if o == nil || o.ConnectionType == nil {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetConnectionTypeOk() (*string, bool) {
	if o == nil || o.ConnectionType == nil {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *DataSourceVM) HasConnectionType() bool {
	if o != nil && o.ConnectionType != nil {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *DataSourceVM) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetConnectionString returns the ConnectionString field value if set, zero value otherwise.
func (o *DataSourceVM) GetConnectionString() string {
	if o == nil || o.ConnectionString == nil {
		var ret string
		return ret
	}
	return *o.ConnectionString
}

// GetConnectionStringOk returns a tuple with the ConnectionString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetConnectionStringOk() (*string, bool) {
	if o == nil || o.ConnectionString == nil {
		return nil, false
	}
	return o.ConnectionString, true
}

// HasConnectionString returns a boolean if a field has been set.
func (o *DataSourceVM) HasConnectionString() bool {
	if o != nil && o.ConnectionString != nil {
		return true
	}

	return false
}

// SetConnectionString gets a reference to the given string and assigns it to the ConnectionString field.
func (o *DataSourceVM) SetConnectionString(v string) {
	o.ConnectionString = &v
}

// GetDataStructure returns the DataStructure field value if set, zero value otherwise.
func (o *DataSourceVM) GetDataStructure() string {
	if o == nil || o.DataStructure == nil {
		var ret string
		return ret
	}
	return *o.DataStructure
}

// GetDataStructureOk returns a tuple with the DataStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetDataStructureOk() (*string, bool) {
	if o == nil || o.DataStructure == nil {
		return nil, false
	}
	return o.DataStructure, true
}

// HasDataStructure returns a boolean if a field has been set.
func (o *DataSourceVM) HasDataStructure() bool {
	if o != nil && o.DataStructure != nil {
		return true
	}

	return false
}

// SetDataStructure gets a reference to the given string and assigns it to the DataStructure field.
func (o *DataSourceVM) SetDataStructure(v string) {
	o.DataStructure = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *DataSourceVM) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *DataSourceVM) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *DataSourceVM) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetEditedTime returns the EditedTime field value if set, zero value otherwise.
func (o *DataSourceVM) GetEditedTime() time.Time {
	if o == nil || o.EditedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EditedTime
}

// GetEditedTimeOk returns a tuple with the EditedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetEditedTimeOk() (*time.Time, bool) {
	if o == nil || o.EditedTime == nil {
		return nil, false
	}
	return o.EditedTime, true
}

// HasEditedTime returns a boolean if a field has been set.
func (o *DataSourceVM) HasEditedTime() bool {
	if o != nil && o.EditedTime != nil {
		return true
	}

	return false
}

// SetEditedTime gets a reference to the given time.Time and assigns it to the EditedTime field.
func (o *DataSourceVM) SetEditedTime(v time.Time) {
	o.EditedTime = &v
}

// GetEditorUserId returns the EditorUserId field value if set, zero value otherwise.
func (o *DataSourceVM) GetEditorUserId() string {
	if o == nil || o.EditorUserId == nil {
		var ret string
		return ret
	}
	return *o.EditorUserId
}

// GetEditorUserIdOk returns a tuple with the EditorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetEditorUserIdOk() (*string, bool) {
	if o == nil || o.EditorUserId == nil {
		return nil, false
	}
	return o.EditorUserId, true
}

// HasEditorUserId returns a boolean if a field has been set.
func (o *DataSourceVM) HasEditorUserId() bool {
	if o != nil && o.EditorUserId != nil {
		return true
	}

	return false
}

// SetEditorUserId gets a reference to the given string and assigns it to the EditorUserId field.
func (o *DataSourceVM) SetEditorUserId(v string) {
	o.EditorUserId = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *DataSourceVM) GetCreatedTime() time.Time {
	if o == nil || o.CreatedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedTime == nil {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *DataSourceVM) HasCreatedTime() bool {
	if o != nil && o.CreatedTime != nil {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *DataSourceVM) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCreatorUserId returns the CreatorUserId field value if set, zero value otherwise.
func (o *DataSourceVM) GetCreatorUserId() string {
	if o == nil || o.CreatorUserId == nil {
		var ret string
		return ret
	}
	return *o.CreatorUserId
}

// GetCreatorUserIdOk returns a tuple with the CreatorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetCreatorUserIdOk() (*string, bool) {
	if o == nil || o.CreatorUserId == nil {
		return nil, false
	}
	return o.CreatorUserId, true
}

// HasCreatorUserId returns a boolean if a field has been set.
func (o *DataSourceVM) HasCreatorUserId() bool {
	if o != nil && o.CreatorUserId != nil {
		return true
	}

	return false
}

// SetCreatorUserId gets a reference to the given string and assigns it to the CreatorUserId field.
func (o *DataSourceVM) SetCreatorUserId(v string) {
	o.CreatorUserId = &v
}

// GetIsConnected returns the IsConnected field value if set, zero value otherwise.
func (o *DataSourceVM) GetIsConnected() bool {
	if o == nil || o.IsConnected == nil {
		var ret bool
		return ret
	}
	return *o.IsConnected
}

// GetIsConnectedOk returns a tuple with the IsConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataSourceVM) GetIsConnectedOk() (*bool, bool) {
	if o == nil || o.IsConnected == nil {
		return nil, false
	}
	return o.IsConnected, true
}

// HasIsConnected returns a boolean if a field has been set.
func (o *DataSourceVM) HasIsConnected() bool {
	if o != nil && o.IsConnected != nil {
		return true
	}

	return false
}

// SetIsConnected gets a reference to the given bool and assigns it to the IsConnected field.
func (o *DataSourceVM) SetIsConnected(v bool) {
	o.IsConnected = &v
}

func (o DataSourceVM) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ConnectionType != nil {
		toSerialize["connectionType"] = o.ConnectionType
	}
	if o.ConnectionString != nil {
		toSerialize["connectionString"] = o.ConnectionString
	}
	if o.DataStructure != nil {
		toSerialize["dataStructure"] = o.DataStructure
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.EditedTime != nil {
		toSerialize["editedTime"] = o.EditedTime
	}
	if o.EditorUserId != nil {
		toSerialize["editorUserId"] = o.EditorUserId
	}
	if o.CreatedTime != nil {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.CreatorUserId != nil {
		toSerialize["creatorUserId"] = o.CreatorUserId
	}
	if o.IsConnected != nil {
		toSerialize["isConnected"] = o.IsConnected
	}
	return json.Marshal(toSerialize)
}

type NullableDataSourceVM struct {
	value *DataSourceVM
	isSet bool
}

func (v NullableDataSourceVM) Get() *DataSourceVM {
	return v.value
}

func (v *NullableDataSourceVM) Set(val *DataSourceVM) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSourceVM) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSourceVM) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSourceVM(val *DataSourceVM) *NullableDataSourceVM {
	return &NullableDataSourceVM{value: val, isSet: true}
}

func (v NullableDataSourceVM) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSourceVM) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

