# UserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | The UUID of the user. | 
**Email** | **string** | The email address of the user. | 
**Name** | Pointer to **string** | The first name of the user. | [optional] 
**Surname** | Pointer to **string** | The last name of the user. | [optional] 
**UserStatus** | Pointer to **string** | The status of this user in Dynatrace:   * &#x60;ACTIVE&#x60;: The user is active. * &#x60;INACTIVE&#x60;: The user is deactivated and cannot sign in to Dynatrace.  * &#x60;PENDING&#x60;: The user received an invitation, but hasn&#39;t completed sign-up yet.  * &#x60;DELETED&#x60;: The user was deleted and cannot sign in to Dynatrace anymore.  * &#x60;ECUSTOMS_MANUALLY_BLOCKED&#x60;: The user is blocked due to to a trade and export compliance violation.   | [optional] 
**EmergencyContact** | Pointer to **bool** | The user is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) an emergency contact for the account. | [optional] 

## Methods

### NewUserDto

`func NewUserDto(uid string, email string, ) *UserDto`

NewUserDto instantiates a new UserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoWithDefaults

`func NewUserDtoWithDefaults() *UserDto`

NewUserDtoWithDefaults instantiates a new UserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUid

`func (o *UserDto) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserDto) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserDto) SetUid(v string)`

SetUid sets Uid field to given value.


### GetEmail

`func (o *UserDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSurname

`func (o *UserDto) GetSurname() string`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *UserDto) GetSurnameOk() (*string, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *UserDto) SetSurname(v string)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *UserDto) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetUserStatus

`func (o *UserDto) GetUserStatus() string`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *UserDto) GetUserStatusOk() (*string, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *UserDto) SetUserStatus(v string)`

SetUserStatus sets UserStatus field to given value.

### HasUserStatus

`func (o *UserDto) HasUserStatus() bool`

HasUserStatus returns a boolean if a field has been set.

### GetEmergencyContact

`func (o *UserDto) GetEmergencyContact() bool`

GetEmergencyContact returns the EmergencyContact field if non-nil, zero value otherwise.

### GetEmergencyContactOk

`func (o *UserDto) GetEmergencyContactOk() (*bool, bool)`

GetEmergencyContactOk returns a tuple with the EmergencyContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyContact

`func (o *UserDto) SetEmergencyContact(v bool)`

SetEmergencyContact sets EmergencyContact field to given value.

### HasEmergencyContact

`func (o *UserDto) HasEmergencyContact() bool`

HasEmergencyContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


