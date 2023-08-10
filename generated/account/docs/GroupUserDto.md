# GroupUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | The UUID of the user. | 
**Email** | **string** | The email address of the user. | 
**Name** | Pointer to **string** | The first name of the user. | [optional] 
**Surname** | Pointer to **string** | The last name of the user. | [optional] 
**UserStatus** | Pointer to **string** | The status of this user in Dynatrace:   * &#x60;ACTIVE&#x60;: The user is active. * &#x60;INACTIVE&#x60;: The user is deactivated and cannot sign in to Dynatrace.  * &#x60;PENDING&#x60;: The user received an invitation, but hasn&#39;t completed sign-up yet.  * &#x60;DELETED&#x60;: The user was deleted and cannot sign in to Dynatrace anymore.  * &#x60;ECUSTOMS_MANUALLY_BLOCKED&#x60;: The user is blocked due to to a trade and export compliance violation.   | [optional] 
**EmergencyContact** | Pointer to **bool** | The user is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) an emergency contact for the account. | [optional] 
**Groups** | [**[]AccountGroupDto**](AccountGroupDto.md) | A list of groups of which the user is a member. | 

## Methods

### NewGroupUserDto

`func NewGroupUserDto(uid string, email string, groups []AccountGroupDto, ) *GroupUserDto`

NewGroupUserDto instantiates a new GroupUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserDtoWithDefaults

`func NewGroupUserDtoWithDefaults() *GroupUserDto`

NewGroupUserDtoWithDefaults instantiates a new GroupUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *GroupUserDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GroupUserDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GroupUserDto) SetUid(v string)`

SetUid sets Uid field to given value.


### GetEmail

`func (o *GroupUserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GroupUserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GroupUserDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *GroupUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupUserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *GroupUserDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *GroupUserDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *GroupUserDto) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *GroupUserDto) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetUserStatus

`func (o *GroupUserDto) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *GroupUserDto) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *GroupUserDto) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *GroupUserDto) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEmergencyContact

`func (o *GroupUserDto) GetEmergencyContact() bool`

GetEmergencyContact returns the EmergencyContact field if non-nil, zero value otherwise.

### GetEmergencyContactOk

`func (o *GroupUserDto) GetEmergencyContactOk() (*bool, bool)`

GetEmergencyContactOk returns a tuple with the EmergencyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContact

`func (o *GroupUserDto) SetEmergencyContact(v bool)`

SetEmergencyContact sets EmergencyContact field to given value.

### HasEmergencyContact

`func (o *GroupUserDto) HasEmergencyContact() bool`

HasEmergencyContact returns a boolean if a field has been set.

### GetGroups

`func (o *GroupUserDto) GetGroups() []AccountGroupDto`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupUserDto) GetGroupsOk() (*[]AccountGroupDto, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupUserDto) SetGroups(v []AccountGroupDto)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


