# MyPermissionsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | Pointer to [**SubscriptionPermissionCRUDVM**](SubscriptionPermissionCRUDVM.md) |  | [optional] 
**Files** | Pointer to [**FilePermissionCRUDVM**](FilePermissionCRUDVM.md) |  | [optional] 
**Datasources** | Pointer to [**DataSourcePermissionCRUDVM**](DataSourcePermissionCRUDVM.md) |  | [optional] 
**Groups** | Pointer to [**GroupPermissionCRUDVM**](GroupPermissionCRUDVM.md) |  | [optional] 
**Tasks** | Pointer to [**TaskPermissionCRUDVM**](TaskPermissionCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewMyPermissionsVM

`func NewMyPermissionsVM(t string, ) *MyPermissionsVM`

NewMyPermissionsVM instantiates a new MyPermissionsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyPermissionsVMWithDefaults

`func NewMyPermissionsVMWithDefaults() *MyPermissionsVM`

NewMyPermissionsVMWithDefaults instantiates a new MyPermissionsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *MyPermissionsVM) GetSubscription() SubscriptionPermissionCRUDVM`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *MyPermissionsVM) GetSubscriptionOk() (*SubscriptionPermissionCRUDVM, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *MyPermissionsVM) SetSubscription(v SubscriptionPermissionCRUDVM)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *MyPermissionsVM) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetFiles

`func (o *MyPermissionsVM) GetFiles() FilePermissionCRUDVM`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *MyPermissionsVM) GetFilesOk() (*FilePermissionCRUDVM, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *MyPermissionsVM) SetFiles(v FilePermissionCRUDVM)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *MyPermissionsVM) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetDatasources

`func (o *MyPermissionsVM) GetDatasources() DataSourcePermissionCRUDVM`

GetDatasources returns the Datasources field if non-nil, zero value otherwise.

### GetDatasourcesOk

`func (o *MyPermissionsVM) GetDatasourcesOk() (*DataSourcePermissionCRUDVM, bool)`

GetDatasourcesOk returns a tuple with the Datasources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasources

`func (o *MyPermissionsVM) SetDatasources(v DataSourcePermissionCRUDVM)`

SetDatasources sets Datasources field to given value.

### HasDatasources

`func (o *MyPermissionsVM) HasDatasources() bool`

HasDatasources returns a boolean if a field has been set.

### GetGroups

`func (o *MyPermissionsVM) GetGroups() GroupPermissionCRUDVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *MyPermissionsVM) GetGroupsOk() (*GroupPermissionCRUDVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *MyPermissionsVM) SetGroups(v GroupPermissionCRUDVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *MyPermissionsVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetTasks

`func (o *MyPermissionsVM) GetTasks() TaskPermissionCRUDVM`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *MyPermissionsVM) GetTasksOk() (*TaskPermissionCRUDVM, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *MyPermissionsVM) SetTasks(v TaskPermissionCRUDVM)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *MyPermissionsVM) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetT

`func (o *MyPermissionsVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *MyPermissionsVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *MyPermissionsVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


