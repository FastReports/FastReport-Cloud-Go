# DataSourceVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**ConnectionType** | Pointer to [**DataSourceConnectionType**](DataSourceConnectionType.md) |  | [optional] 
**ConnectionString** | Pointer to **NullableString** |  | [optional] 
**DataStructure** | Pointer to **NullableString** |  | [optional] 
**SubscriptionId** | Pointer to **NullableString** |  | [optional] 
**EditedTime** | Pointer to **time.Time** |  | [optional] 
**EditorUserId** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CreatorUserId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**DataSourceStatus**](DataSourceStatus.md) |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDataSourceVM

`func NewDataSourceVM() *DataSourceVM`

NewDataSourceVM instantiates a new DataSourceVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceVMWithDefaults

`func NewDataSourceVMWithDefaults() *DataSourceVM`

NewDataSourceVMWithDefaults instantiates a new DataSourceVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DataSourceVM) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataSourceVM) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataSourceVM) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DataSourceVM) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DataSourceVM) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DataSourceVM) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DataSourceVM) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSourceVM) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSourceVM) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataSourceVM) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DataSourceVM) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DataSourceVM) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetConnectionType

`func (o *DataSourceVM) GetConnectionType() DataSourceConnectionType`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *DataSourceVM) GetConnectionTypeOk() (*DataSourceConnectionType, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *DataSourceVM) SetConnectionType(v DataSourceConnectionType)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *DataSourceVM) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetConnectionString

`func (o *DataSourceVM) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *DataSourceVM) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *DataSourceVM) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *DataSourceVM) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### SetConnectionStringNil

`func (o *DataSourceVM) SetConnectionStringNil(b bool)`

 SetConnectionStringNil sets the value for ConnectionString to be an explicit nil

### UnsetConnectionString
`func (o *DataSourceVM) UnsetConnectionString()`

UnsetConnectionString ensures that no value is present for ConnectionString, not even an explicit nil
### GetDataStructure

`func (o *DataSourceVM) GetDataStructure() string`

GetDataStructure returns the DataStructure field if non-nil, zero value otherwise.

### GetDataStructureOk

`func (o *DataSourceVM) GetDataStructureOk() (*string, bool)`

GetDataStructureOk returns a tuple with the DataStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStructure

`func (o *DataSourceVM) SetDataStructure(v string)`

SetDataStructure sets DataStructure field to given value.

### HasDataStructure

`func (o *DataSourceVM) HasDataStructure() bool`

HasDataStructure returns a boolean if a field has been set.

### SetDataStructureNil

`func (o *DataSourceVM) SetDataStructureNil(b bool)`

 SetDataStructureNil sets the value for DataStructure to be an explicit nil

### UnsetDataStructure
`func (o *DataSourceVM) UnsetDataStructure()`

UnsetDataStructure ensures that no value is present for DataStructure, not even an explicit nil
### GetSubscriptionId

`func (o *DataSourceVM) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DataSourceVM) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DataSourceVM) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DataSourceVM) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### SetSubscriptionIdNil

`func (o *DataSourceVM) SetSubscriptionIdNil(b bool)`

 SetSubscriptionIdNil sets the value for SubscriptionId to be an explicit nil

### UnsetSubscriptionId
`func (o *DataSourceVM) UnsetSubscriptionId()`

UnsetSubscriptionId ensures that no value is present for SubscriptionId, not even an explicit nil
### GetEditedTime

`func (o *DataSourceVM) GetEditedTime() time.Time`

GetEditedTime returns the EditedTime field if non-nil, zero value otherwise.

### GetEditedTimeOk

`func (o *DataSourceVM) GetEditedTimeOk() (*time.Time, bool)`

GetEditedTimeOk returns a tuple with the EditedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedTime

`func (o *DataSourceVM) SetEditedTime(v time.Time)`

SetEditedTime sets EditedTime field to given value.

### HasEditedTime

`func (o *DataSourceVM) HasEditedTime() bool`

HasEditedTime returns a boolean if a field has been set.

### GetEditorUserId

`func (o *DataSourceVM) GetEditorUserId() string`

GetEditorUserId returns the EditorUserId field if non-nil, zero value otherwise.

### GetEditorUserIdOk

`func (o *DataSourceVM) GetEditorUserIdOk() (*string, bool)`

GetEditorUserIdOk returns a tuple with the EditorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorUserId

`func (o *DataSourceVM) SetEditorUserId(v string)`

SetEditorUserId sets EditorUserId field to given value.

### HasEditorUserId

`func (o *DataSourceVM) HasEditorUserId() bool`

HasEditorUserId returns a boolean if a field has been set.

### SetEditorUserIdNil

`func (o *DataSourceVM) SetEditorUserIdNil(b bool)`

 SetEditorUserIdNil sets the value for EditorUserId to be an explicit nil

### UnsetEditorUserId
`func (o *DataSourceVM) UnsetEditorUserId()`

UnsetEditorUserId ensures that no value is present for EditorUserId, not even an explicit nil
### GetCreatedTime

`func (o *DataSourceVM) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *DataSourceVM) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *DataSourceVM) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *DataSourceVM) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCreatorUserId

`func (o *DataSourceVM) GetCreatorUserId() string`

GetCreatorUserId returns the CreatorUserId field if non-nil, zero value otherwise.

### GetCreatorUserIdOk

`func (o *DataSourceVM) GetCreatorUserIdOk() (*string, bool)`

GetCreatorUserIdOk returns a tuple with the CreatorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUserId

`func (o *DataSourceVM) SetCreatorUserId(v string)`

SetCreatorUserId sets CreatorUserId field to given value.

### HasCreatorUserId

`func (o *DataSourceVM) HasCreatorUserId() bool`

HasCreatorUserId returns a boolean if a field has been set.

### SetCreatorUserIdNil

`func (o *DataSourceVM) SetCreatorUserIdNil(b bool)`

 SetCreatorUserIdNil sets the value for CreatorUserId to be an explicit nil

### UnsetCreatorUserId
`func (o *DataSourceVM) UnsetCreatorUserId()`

UnsetCreatorUserId ensures that no value is present for CreatorUserId, not even an explicit nil
### GetStatus

`func (o *DataSourceVM) GetStatus() DataSourceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DataSourceVM) GetStatusOk() (*DataSourceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DataSourceVM) SetStatus(v DataSourceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DataSourceVM) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetErrorMessage

`func (o *DataSourceVM) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *DataSourceVM) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *DataSourceVM) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *DataSourceVM) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *DataSourceVM) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *DataSourceVM) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


