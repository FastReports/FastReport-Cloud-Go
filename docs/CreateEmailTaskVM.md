# CreateEmailTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** |  | [optional] 
**EnableSsl** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **NullableString** |  | [optional] 
**IsBodyHtml** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Server** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**T** | **string** |  | 

## Methods

### NewCreateEmailTaskVM

`func NewCreateEmailTaskVM(t string, ) *CreateEmailTaskVM`

NewCreateEmailTaskVM instantiates a new CreateEmailTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailTaskVMWithDefaults

`func NewCreateEmailTaskVMWithDefaults() *CreateEmailTaskVM`

NewCreateEmailTaskVMWithDefaults instantiates a new CreateEmailTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *CreateEmailTaskVM) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateEmailTaskVM) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateEmailTaskVM) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateEmailTaskVM) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *CreateEmailTaskVM) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *CreateEmailTaskVM) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetEnableSsl

`func (o *CreateEmailTaskVM) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *CreateEmailTaskVM) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *CreateEmailTaskVM) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *CreateEmailTaskVM) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### GetFrom

`func (o *CreateEmailTaskVM) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CreateEmailTaskVM) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CreateEmailTaskVM) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CreateEmailTaskVM) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *CreateEmailTaskVM) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *CreateEmailTaskVM) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetIsBodyHtml

`func (o *CreateEmailTaskVM) GetIsBodyHtml() bool`

GetIsBodyHtml returns the IsBodyHtml field if non-nil, zero value otherwise.

### GetIsBodyHtmlOk

`func (o *CreateEmailTaskVM) GetIsBodyHtmlOk() (*bool, bool)`

GetIsBodyHtmlOk returns a tuple with the IsBodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBodyHtml

`func (o *CreateEmailTaskVM) SetIsBodyHtml(v bool)`

SetIsBodyHtml sets IsBodyHtml field to given value.

### HasIsBodyHtml

`func (o *CreateEmailTaskVM) HasIsBodyHtml() bool`

HasIsBodyHtml returns a boolean if a field has been set.

### GetPassword

`func (o *CreateEmailTaskVM) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateEmailTaskVM) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateEmailTaskVM) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateEmailTaskVM) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateEmailTaskVM) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateEmailTaskVM) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPort

`func (o *CreateEmailTaskVM) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateEmailTaskVM) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateEmailTaskVM) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateEmailTaskVM) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServer

`func (o *CreateEmailTaskVM) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateEmailTaskVM) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateEmailTaskVM) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *CreateEmailTaskVM) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *CreateEmailTaskVM) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *CreateEmailTaskVM) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetSubject

`func (o *CreateEmailTaskVM) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateEmailTaskVM) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateEmailTaskVM) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CreateEmailTaskVM) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CreateEmailTaskVM) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CreateEmailTaskVM) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTo

`func (o *CreateEmailTaskVM) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CreateEmailTaskVM) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CreateEmailTaskVM) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *CreateEmailTaskVM) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *CreateEmailTaskVM) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *CreateEmailTaskVM) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUsername

`func (o *CreateEmailTaskVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateEmailTaskVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateEmailTaskVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateEmailTaskVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CreateEmailTaskVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CreateEmailTaskVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetT

`func (o *CreateEmailTaskVM) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *CreateEmailTaskVM) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *CreateEmailTaskVM) SetT(v string)`

SetT sets T field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


