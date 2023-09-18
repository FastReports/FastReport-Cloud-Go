# ContactsVM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]ContactVM**](ContactVM.md) |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**Take** | Pointer to **int32** |  | [optional] 
**Count** | Pointer to **int64** |  | [optional] 

## Methods

### NewContactsVM

`func NewContactsVM() *ContactsVM`

NewContactsVM instantiates a new ContactsVM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsVMWithDefaults

`func NewContactsVMWithDefaults() *ContactsVM`

NewContactsVMWithDefaults instantiates a new ContactsVM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *ContactsVM) GetContacts() []ContactVM`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ContactsVM) GetContactsOk() (*[]ContactVM, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ContactsVM) SetContacts(v []ContactVM)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ContactsVM) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *ContactsVM) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *ContactsVM) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetSkip

`func (o *ContactsVM) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ContactsVM) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ContactsVM) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ContactsVM) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetTake

`func (o *ContactsVM) GetTake() int32`

GetTake returns the Take field if non-nil, zero value otherwise.

### GetTakeOk

`func (o *ContactsVM) GetTakeOk() (*int32, bool)`

GetTakeOk returns a tuple with the Take field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTake

`func (o *ContactsVM) SetTake(v int32)`

SetTake sets Take field to given value.

### HasTake

`func (o *ContactsVM) HasTake() bool`

HasTake returns a boolean if a field has been set.

### GetCount

`func (o *ContactsVM) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ContactsVM) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ContactsVM) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ContactsVM) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


