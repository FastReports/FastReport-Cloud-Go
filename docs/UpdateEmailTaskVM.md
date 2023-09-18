# UpdateEmailTaskVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** |  | [optional] 
**EnableSsl** | Pointer to **NullableBool** |  | [optional] 
**From** | Pointer to **NullableString** |  | [optional] 
**IsBodyHtml** | Pointer to **NullableBool** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Server** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**To** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateEmailTaskVM

`func NewUpdateEmailTaskVM() *UpdateEmailTaskVM`

NewUpdateEmailTaskVM instantiates a new UpdateEmailTaskVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailTaskVMWithDefaults

`func NewUpdateEmailTaskVMWithDefaults() *UpdateEmailTaskVM`

NewUpdateEmailTaskVMWithDefaults instantiates a new UpdateEmailTaskVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *UpdateEmailTaskVM) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateEmailTaskVM) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateEmailTaskVM) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateEmailTaskVM) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *UpdateEmailTaskVM) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *UpdateEmailTaskVM) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetEnableSsl

`func (o *UpdateEmailTaskVM) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *UpdateEmailTaskVM) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *UpdateEmailTaskVM) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *UpdateEmailTaskVM) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### SetEnableSslNil

`func (o *UpdateEmailTaskVM) SetEnableSslNil(b bool)`

 SetEnableSslNil sets the value for EnableSsl to be an explicit nil

### UnsetEnableSsl
`func (o *UpdateEmailTaskVM) UnsetEnableSsl()`

UnsetEnableSsl ensures that no value is present for EnableSsl, not even an explicit nil
### GetFrom

`func (o *UpdateEmailTaskVM) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UpdateEmailTaskVM) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UpdateEmailTaskVM) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *UpdateEmailTaskVM) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *UpdateEmailTaskVM) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *UpdateEmailTaskVM) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetIsBodyHtml

`func (o *UpdateEmailTaskVM) GetIsBodyHtml() bool`

GetIsBodyHtml returns the IsBodyHtml field if non-nil, zero value otherwise.

### GetIsBodyHtmlOk

`func (o *UpdateEmailTaskVM) GetIsBodyHtmlOk() (*bool, bool)`

GetIsBodyHtmlOk returns a tuple with the IsBodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBodyHtml

`func (o *UpdateEmailTaskVM) SetIsBodyHtml(v bool)`

SetIsBodyHtml sets IsBodyHtml field to given value.

### HasIsBodyHtml

`func (o *UpdateEmailTaskVM) HasIsBodyHtml() bool`

HasIsBodyHtml returns a boolean if a field has been set.

### SetIsBodyHtmlNil

`func (o *UpdateEmailTaskVM) SetIsBodyHtmlNil(b bool)`

 SetIsBodyHtmlNil sets the value for IsBodyHtml to be an explicit nil

### UnsetIsBodyHtml
`func (o *UpdateEmailTaskVM) UnsetIsBodyHtml()`

UnsetIsBodyHtml ensures that no value is present for IsBodyHtml, not even an explicit nil
### GetPassword

`func (o *UpdateEmailTaskVM) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateEmailTaskVM) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateEmailTaskVM) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateEmailTaskVM) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *UpdateEmailTaskVM) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *UpdateEmailTaskVM) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetPort

`func (o *UpdateEmailTaskVM) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateEmailTaskVM) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateEmailTaskVM) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateEmailTaskVM) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *UpdateEmailTaskVM) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *UpdateEmailTaskVM) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetServer

`func (o *UpdateEmailTaskVM) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateEmailTaskVM) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateEmailTaskVM) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateEmailTaskVM) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *UpdateEmailTaskVM) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *UpdateEmailTaskVM) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetSubject

`func (o *UpdateEmailTaskVM) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpdateEmailTaskVM) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpdateEmailTaskVM) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpdateEmailTaskVM) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *UpdateEmailTaskVM) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *UpdateEmailTaskVM) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetTo

`func (o *UpdateEmailTaskVM) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UpdateEmailTaskVM) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UpdateEmailTaskVM) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *UpdateEmailTaskVM) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *UpdateEmailTaskVM) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *UpdateEmailTaskVM) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetUsername

`func (o *UpdateEmailTaskVM) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateEmailTaskVM) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateEmailTaskVM) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateEmailTaskVM) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UpdateEmailTaskVM) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UpdateEmailTaskVM) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


