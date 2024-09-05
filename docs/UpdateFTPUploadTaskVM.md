# UpdateFTPUploadTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archive** | Pointer to **NullableBool** |  | [optional] 
**ArchiveName** | Pointer to **NullableString** |  | [optional] 
**DestinationFolder** | Pointer to **NullableString** |  | [optional] 
**FtpHost** | Pointer to **NullableString** |  | [optional] 
**FtpPassword** | Pointer to **NullableString** |  | [optional] 
**FtpPort** | Pointer to **NullableInt32** |  | [optional] 
**FtpUsername** | Pointer to **NullableString** |  | [optional] 
**UseSFTP** | Pointer to **NullableBool** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewUpdateFTPUploadTaskVM

`func NewUpdateFTPUploadTaskVM(t string, ) *UpdateFTPUploadTaskVM`

NewUpdateFTPUploadTaskVM instantiates a new UpdateFTPUploadTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFTPUploadTaskVMWithDefaults

`func NewUpdateFTPUploadTaskVMWithDefaults() *UpdateFTPUploadTaskVM`

NewUpdateFTPUploadTaskVMWithDefaults instantiates a new UpdateFTPUploadTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchive

`func (o *UpdateFTPUploadTaskVM) GetArchive() bool`

GetArchive returns the Archive field if non-nil, zero value otherwise.

### GetArchiveOk

`func (o *UpdateFTPUploadTaskVM) GetArchiveOk() (*bool, bool)`

GetArchiveOk returns a tuple with the Archive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchive

`func (o *UpdateFTPUploadTaskVM) SetArchive(v bool)`

SetArchive sets Archive field to given value.

### HasArchive

`func (o *UpdateFTPUploadTaskVM) HasArchive() bool`

HasArchive returns a boolean if a field has been set.

### SetArchiveNil

`func (o *UpdateFTPUploadTaskVM) SetArchiveNil(b bool)`

 SetArchiveNil sets the value for Archive to be an explicit nil

### UnsetArchive
`func (o *UpdateFTPUploadTaskVM) UnsetArchive()`

UnsetArchive ensures that no value is present for Archive, not even an explicit nil
### GetArchiveName

`func (o *UpdateFTPUploadTaskVM) GetArchiveName() string`

GetArchiveName returns the ArchiveName field if non-nil, zero value otherwise.

### GetArchiveNameOk

`func (o *UpdateFTPUploadTaskVM) GetArchiveNameOk() (*string, bool)`

GetArchiveNameOk returns a tuple with the ArchiveName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveName

`func (o *UpdateFTPUploadTaskVM) SetArchiveName(v string)`

SetArchiveName sets ArchiveName field to given value.

### HasArchiveName

`func (o *UpdateFTPUploadTaskVM) HasArchiveName() bool`

HasArchiveName returns a boolean if a field has been set.

### SetArchiveNameNil

`func (o *UpdateFTPUploadTaskVM) SetArchiveNameNil(b bool)`

 SetArchiveNameNil sets the value for ArchiveName to be an explicit nil

### UnsetArchiveName
`func (o *UpdateFTPUploadTaskVM) UnsetArchiveName()`

UnsetArchiveName ensures that no value is present for ArchiveName, not even an explicit nil
### GetDestinationFolder

`func (o *UpdateFTPUploadTaskVM) GetDestinationFolder() string`

GetDestinationFolder returns the DestinationFolder field if non-nil, zero value otherwise.

### GetDestinationFolderOk

`func (o *UpdateFTPUploadTaskVM) GetDestinationFolderOk() (*string, bool)`

GetDestinationFolderOk returns a tuple with the DestinationFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFolder

`func (o *UpdateFTPUploadTaskVM) SetDestinationFolder(v string)`

SetDestinationFolder sets DestinationFolder field to given value.

### HasDestinationFolder

`func (o *UpdateFTPUploadTaskVM) HasDestinationFolder() bool`

HasDestinationFolder returns a boolean if a field has been set.

### SetDestinationFolderNil

`func (o *UpdateFTPUploadTaskVM) SetDestinationFolderNil(b bool)`

 SetDestinationFolderNil sets the value for DestinationFolder to be an explicit nil

### UnsetDestinationFolder
`func (o *UpdateFTPUploadTaskVM) UnsetDestinationFolder()`

UnsetDestinationFolder ensures that no value is present for DestinationFolder, not even an explicit nil
### GetFtpHost

`func (o *UpdateFTPUploadTaskVM) GetFtpHost() string`

GetFtpHost returns the FtpHost field if non-nil, zero value otherwise.

### GetFtpHostOk

`func (o *UpdateFTPUploadTaskVM) GetFtpHostOk() (*string, bool)`

GetFtpHostOk returns a tuple with the FtpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpHost

`func (o *UpdateFTPUploadTaskVM) SetFtpHost(v string)`

SetFtpHost sets FtpHost field to given value.

### HasFtpHost

`func (o *UpdateFTPUploadTaskVM) HasFtpHost() bool`

HasFtpHost returns a boolean if a field has been set.

### SetFtpHostNil

`func (o *UpdateFTPUploadTaskVM) SetFtpHostNil(b bool)`

 SetFtpHostNil sets the value for FtpHost to be an explicit nil

### UnsetFtpHost
`func (o *UpdateFTPUploadTaskVM) UnsetFtpHost()`

UnsetFtpHost ensures that no value is present for FtpHost, not even an explicit nil
### GetFtpPassword

`func (o *UpdateFTPUploadTaskVM) GetFtpPassword() string`

GetFtpPassword returns the FtpPassword field if non-nil, zero value otherwise.

### GetFtpPasswordOk

`func (o *UpdateFTPUploadTaskVM) GetFtpPasswordOk() (*string, bool)`

GetFtpPasswordOk returns a tuple with the FtpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPassword

`func (o *UpdateFTPUploadTaskVM) SetFtpPassword(v string)`

SetFtpPassword sets FtpPassword field to given value.

### HasFtpPassword

`func (o *UpdateFTPUploadTaskVM) HasFtpPassword() bool`

HasFtpPassword returns a boolean if a field has been set.

### SetFtpPasswordNil

`func (o *UpdateFTPUploadTaskVM) SetFtpPasswordNil(b bool)`

 SetFtpPasswordNil sets the value for FtpPassword to be an explicit nil

### UnsetFtpPassword
`func (o *UpdateFTPUploadTaskVM) UnsetFtpPassword()`

UnsetFtpPassword ensures that no value is present for FtpPassword, not even an explicit nil
### GetFtpPort

`func (o *UpdateFTPUploadTaskVM) GetFtpPort() int32`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *UpdateFTPUploadTaskVM) GetFtpPortOk() (*int32, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *UpdateFTPUploadTaskVM) SetFtpPort(v int32)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *UpdateFTPUploadTaskVM) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### SetFtpPortNil

`func (o *UpdateFTPUploadTaskVM) SetFtpPortNil(b bool)`

 SetFtpPortNil sets the value for FtpPort to be an explicit nil

### UnsetFtpPort
`func (o *UpdateFTPUploadTaskVM) UnsetFtpPort()`

UnsetFtpPort ensures that no value is present for FtpPort, not even an explicit nil
### GetFtpUsername

`func (o *UpdateFTPUploadTaskVM) GetFtpUsername() string`

GetFtpUsername returns the FtpUsername field if non-nil, zero value otherwise.

### GetFtpUsernameOk

`func (o *UpdateFTPUploadTaskVM) GetFtpUsernameOk() (*string, bool)`

GetFtpUsernameOk returns a tuple with the FtpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpUsername

`func (o *UpdateFTPUploadTaskVM) SetFtpUsername(v string)`

SetFtpUsername sets FtpUsername field to given value.

### HasFtpUsername

`func (o *UpdateFTPUploadTaskVM) HasFtpUsername() bool`

HasFtpUsername returns a boolean if a field has been set.

### SetFtpUsernameNil

`func (o *UpdateFTPUploadTaskVM) SetFtpUsernameNil(b bool)`

 SetFtpUsernameNil sets the value for FtpUsername to be an explicit nil

### UnsetFtpUsername
`func (o *UpdateFTPUploadTaskVM) UnsetFtpUsername()`

UnsetFtpUsername ensures that no value is present for FtpUsername, not even an explicit nil
### GetUseSFTP

`func (o *UpdateFTPUploadTaskVM) GetUseSFTP() bool`

GetUseSFTP returns the UseSFTP field if non-nil, zero value otherwise.

### GetUseSFTPOk

`func (o *UpdateFTPUploadTaskVM) GetUseSFTPOk() (*bool, bool)`

GetUseSFTPOk returns a tuple with the UseSFTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSFTP

`func (o *UpdateFTPUploadTaskVM) SetUseSFTP(v bool)`

SetUseSFTP sets UseSFTP field to given value.

### HasUseSFTP

`func (o *UpdateFTPUploadTaskVM) HasUseSFTP() bool`

HasUseSFTP returns a boolean if a field has been set.

### SetUseSFTPNil

`func (o *UpdateFTPUploadTaskVM) SetUseSFTPNil(b bool)`

 SetUseSFTPNil sets the value for UseSFTP to be an explicit nil

### UnsetUseSFTP
`func (o *UpdateFTPUploadTaskVM) UnsetUseSFTP()`

UnsetUseSFTP ensures that no value is present for UseSFTP, not even an explicit nil
### GetT

`func (o *UpdateFTPUploadTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *UpdateFTPUploadTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *UpdateFTPUploadTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


