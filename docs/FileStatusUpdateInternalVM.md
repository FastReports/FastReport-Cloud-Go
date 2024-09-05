# FileStatusUpdateInternalVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**FileStatus**](FileStatus.md) |  | [optional] 
**Reason** | Pointer to [**FileStatusReason**](FileStatusReason.md) |  | [optional] 
**EditorId** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewFileStatusUpdateInternalVM

`func NewFileStatusUpdateInternalVM(t string, ) *FileStatusUpdateInternalVM`

NewFileStatusUpdateInternalVM instantiates a new FileStatusUpdateInternalVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileStatusUpdateInternalVMWithDefaults

`func NewFileStatusUpdateInternalVMWithDefaults() *FileStatusUpdateInternalVM`

NewFileStatusUpdateInternalVMWithDefaults instantiates a new FileStatusUpdateInternalVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FileStatusUpdateInternalVM) GetStatus() FileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileStatusUpdateInternalVM) GetStatusOk() (*FileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileStatusUpdateInternalVM) SetStatus(v FileStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileStatusUpdateInternalVM) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *FileStatusUpdateInternalVM) GetReason() FileStatusReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *FileStatusUpdateInternalVM) GetReasonOk() (*FileStatusReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *FileStatusUpdateInternalVM) SetReason(v FileStatusReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *FileStatusUpdateInternalVM) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetEditorId

`func (o *FileStatusUpdateInternalVM) GetEditorId() string`

GetEditorId returns the EditorId field if non-nil, zero value otherwise.

### GetEditorIdOk

`func (o *FileStatusUpdateInternalVM) GetEditorIdOk() (*string, bool)`

GetEditorIdOk returns a tuple with the EditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorId

`func (o *FileStatusUpdateInternalVM) SetEditorId(v string)`

SetEditorId sets EditorId field to given value.

### HasEditorId

`func (o *FileStatusUpdateInternalVM) HasEditorId() bool`

HasEditorId returns a boolean if a field has been set.

### SetEditorIdNil

`func (o *FileStatusUpdateInternalVM) SetEditorIdNil(b bool)`

 SetEditorIdNil sets the value for EditorId to be an explicit nil

### UnsetEditorId
`func (o *FileStatusUpdateInternalVM) UnsetEditorId()`

UnsetEditorId ensures that no value is present for EditorId, not even an explicit nil
### GetErrorMessage

`func (o *FileStatusUpdateInternalVM) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FileStatusUpdateInternalVM) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FileStatusUpdateInternalVM) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FileStatusUpdateInternalVM) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *FileStatusUpdateInternalVM) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *FileStatusUpdateInternalVM) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetT

`func (o *FileStatusUpdateInternalVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *FileStatusUpdateInternalVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *FileStatusUpdateInternalVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


