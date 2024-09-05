# DataSourcePermissionsCRUDVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**DataSourcePermissionCRUDVM**](DataSourcePermissionCRUDVM.md) |  | [optional] 
**Groups** | Pointer to [**map[string]DataSourcePermissionCRUDVM**](DataSourcePermissionCRUDVM.md) |  | [optional] 
**Other** | Pointer to [**DataSourcePermissionCRUDVM**](DataSourcePermissionCRUDVM.md) |  | [optional] 
**Anon** | Pointer to [**DataSourcePermissionCRUDVM**](DataSourcePermissionCRUDVM.md) |  | [optional] 
**T** | **string** |  | 

## Methods

### NewDataSourcePermissionsCRUDVM

`func NewDataSourcePermissionsCRUDVM(t string, ) *DataSourcePermissionsCRUDVM`

NewDataSourcePermissionsCRUDVM instantiates a new DataSourcePermissionsCRUDVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourcePermissionsCRUDVMWithDefaults

`func NewDataSourcePermissionsCRUDVMWithDefaults() *DataSourcePermissionsCRUDVM`

NewDataSourcePermissionsCRUDVMWithDefaults instantiates a new DataSourcePermissionsCRUDVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnerId

`func (o *DataSourcePermissionsCRUDVM) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *DataSourcePermissionsCRUDVM) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *DataSourcePermissionsCRUDVM) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *DataSourcePermissionsCRUDVM) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *DataSourcePermissionsCRUDVM) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *DataSourcePermissionsCRUDVM) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwner

`func (o *DataSourcePermissionsCRUDVM) GetOwner() DataSourcePermissionCRUDVM`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DataSourcePermissionsCRUDVM) GetOwnerOk() (*DataSourcePermissionCRUDVM, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DataSourcePermissionsCRUDVM) SetOwner(v DataSourcePermissionCRUDVM)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DataSourcePermissionsCRUDVM) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetGroups

`func (o *DataSourcePermissionsCRUDVM) GetGroups() map[string]DataSourcePermissionCRUDVM`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *DataSourcePermissionsCRUDVM) GetGroupsOk() (*map[string]DataSourcePermissionCRUDVM, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *DataSourcePermissionsCRUDVM) SetGroups(v map[string]DataSourcePermissionCRUDVM)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *DataSourcePermissionsCRUDVM) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *DataSourcePermissionsCRUDVM) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *DataSourcePermissionsCRUDVM) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetOther

`func (o *DataSourcePermissionsCRUDVM) GetOther() DataSourcePermissionCRUDVM`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *DataSourcePermissionsCRUDVM) GetOtherOk() (*DataSourcePermissionCRUDVM, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *DataSourcePermissionsCRUDVM) SetOther(v DataSourcePermissionCRUDVM)`

SetOther sets Other field to given value.

### HasOther

`func (o *DataSourcePermissionsCRUDVM) HasOther() bool`

HasOther returns a boolean if a field has been set.

### GetAnon

`func (o *DataSourcePermissionsCRUDVM) GetAnon() DataSourcePermissionCRUDVM`

GetAnon returns the Anon field if non-nil, zero value otherwise.

### GetAnonOk

`func (o *DataSourcePermissionsCRUDVM) GetAnonOk() (*DataSourcePermissionCRUDVM, bool)`

GetAnonOk returns a tuple with the Anon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnon

`func (o *DataSourcePermissionsCRUDVM) SetAnon(v DataSourcePermissionCRUDVM)`

SetAnon sets Anon field to given value.

### HasAnon

`func (o *DataSourcePermissionsCRUDVM) HasAnon() bool`

HasAnon returns a boolean if a field has been set.

### GetT

`func (o *DataSourcePermissionsCRUDVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *DataSourcePermissionsCRUDVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *DataSourcePermissionsCRUDVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


