# GlobalEventCaptureSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEventCapturedAsUserInput** | **string** | Additional events to be captured globally as user input.   For example, DragStart or DragEnd. | 
**Change** | Pointer to **bool** | Change enabled/disabled. | [optional] 
**Click** | **bool** | Click enabled/disabled. | 
**DoubleClick** | **bool** | DoubleClick enabled/disabled. | 
**KeyDown** | **bool** | KeyDown enabled/disabled. | 
**KeyUp** | **bool** | KeyUp enabled/disabled. | 
**MouseDown** | **bool** | MouseDown enabled/disabled. | 
**MouseUp** | **bool** | MouseUp enabled/disabled. | 
**Scroll** | **bool** | Scroll enabled/disabled. | 
**TouchEnd** | Pointer to **bool** | TouchEnd enabled/disabled. | [optional] 
**TouchStart** | Pointer to **bool** | TouchStart enabled/disabled. | [optional] 

## Methods

### NewGlobalEventCaptureSettings

`func NewGlobalEventCaptureSettings(additionalEventCapturedAsUserInput string, click bool, doubleClick bool, keyDown bool, keyUp bool, mouseDown bool, mouseUp bool, scroll bool, ) *GlobalEventCaptureSettings`

NewGlobalEventCaptureSettings instantiates a new GlobalEventCaptureSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalEventCaptureSettingsWithDefaults

`func NewGlobalEventCaptureSettingsWithDefaults() *GlobalEventCaptureSettings`

NewGlobalEventCaptureSettingsWithDefaults instantiates a new GlobalEventCaptureSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalEventCapturedAsUserInput

`func (o *GlobalEventCaptureSettings) GetAdditionalEventCapturedAsUserInput() string`

GetAdditionalEventCapturedAsUserInput returns the AdditionalEventCapturedAsUserInput field if non-nil, zero value otherwise.

### GetAdditionalEventCapturedAsUserInputOk

`func (o *GlobalEventCaptureSettings) GetAdditionalEventCapturedAsUserInputOk() (*string, bool)`

GetAdditionalEventCapturedAsUserInputOk returns a tuple with the AdditionalEventCapturedAsUserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEventCapturedAsUserInput

`func (o *GlobalEventCaptureSettings) SetAdditionalEventCapturedAsUserInput(v string)`

SetAdditionalEventCapturedAsUserInput sets AdditionalEventCapturedAsUserInput field to given value.


### GetChange

`func (o *GlobalEventCaptureSettings) GetChange() bool`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *GlobalEventCaptureSettings) GetChangeOk() (*bool, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *GlobalEventCaptureSettings) SetChange(v bool)`

SetChange sets Change field to given value.

### HasChange

`func (o *GlobalEventCaptureSettings) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetClick

`func (o *GlobalEventCaptureSettings) GetClick() bool`

GetClick returns the Click field if non-nil, zero value otherwise.

### GetClickOk

`func (o *GlobalEventCaptureSettings) GetClickOk() (*bool, bool)`

GetClickOk returns a tuple with the Click field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClick

`func (o *GlobalEventCaptureSettings) SetClick(v bool)`

SetClick sets Click field to given value.


### GetDoubleClick

`func (o *GlobalEventCaptureSettings) GetDoubleClick() bool`

GetDoubleClick returns the DoubleClick field if non-nil, zero value otherwise.

### GetDoubleClickOk

`func (o *GlobalEventCaptureSettings) GetDoubleClickOk() (*bool, bool)`

GetDoubleClickOk returns a tuple with the DoubleClick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleClick

`func (o *GlobalEventCaptureSettings) SetDoubleClick(v bool)`

SetDoubleClick sets DoubleClick field to given value.


### GetKeyDown

`func (o *GlobalEventCaptureSettings) GetKeyDown() bool`

GetKeyDown returns the KeyDown field if non-nil, zero value otherwise.

### GetKeyDownOk

`func (o *GlobalEventCaptureSettings) GetKeyDownOk() (*bool, bool)`

GetKeyDownOk returns a tuple with the KeyDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDown

`func (o *GlobalEventCaptureSettings) SetKeyDown(v bool)`

SetKeyDown sets KeyDown field to given value.


### GetKeyUp

`func (o *GlobalEventCaptureSettings) GetKeyUp() bool`

GetKeyUp returns the KeyUp field if non-nil, zero value otherwise.

### GetKeyUpOk

`func (o *GlobalEventCaptureSettings) GetKeyUpOk() (*bool, bool)`

GetKeyUpOk returns a tuple with the KeyUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUp

`func (o *GlobalEventCaptureSettings) SetKeyUp(v bool)`

SetKeyUp sets KeyUp field to given value.


### GetMouseDown

`func (o *GlobalEventCaptureSettings) GetMouseDown() bool`

GetMouseDown returns the MouseDown field if non-nil, zero value otherwise.

### GetMouseDownOk

`func (o *GlobalEventCaptureSettings) GetMouseDownOk() (*bool, bool)`

GetMouseDownOk returns a tuple with the MouseDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseDown

`func (o *GlobalEventCaptureSettings) SetMouseDown(v bool)`

SetMouseDown sets MouseDown field to given value.


### GetMouseUp

`func (o *GlobalEventCaptureSettings) GetMouseUp() bool`

GetMouseUp returns the MouseUp field if non-nil, zero value otherwise.

### GetMouseUpOk

`func (o *GlobalEventCaptureSettings) GetMouseUpOk() (*bool, bool)`

GetMouseUpOk returns a tuple with the MouseUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMouseUp

`func (o *GlobalEventCaptureSettings) SetMouseUp(v bool)`

SetMouseUp sets MouseUp field to given value.


### GetScroll

`func (o *GlobalEventCaptureSettings) GetScroll() bool`

GetScroll returns the Scroll field if non-nil, zero value otherwise.

### GetScrollOk

`func (o *GlobalEventCaptureSettings) GetScrollOk() (*bool, bool)`

GetScrollOk returns a tuple with the Scroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScroll

`func (o *GlobalEventCaptureSettings) SetScroll(v bool)`

SetScroll sets Scroll field to given value.


### GetTouchEnd

`func (o *GlobalEventCaptureSettings) GetTouchEnd() bool`

GetTouchEnd returns the TouchEnd field if non-nil, zero value otherwise.

### GetTouchEndOk

`func (o *GlobalEventCaptureSettings) GetTouchEndOk() (*bool, bool)`

GetTouchEndOk returns a tuple with the TouchEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTouchEnd

`func (o *GlobalEventCaptureSettings) SetTouchEnd(v bool)`

SetTouchEnd sets TouchEnd field to given value.

### HasTouchEnd

`func (o *GlobalEventCaptureSettings) HasTouchEnd() bool`

HasTouchEnd returns a boolean if a field has been set.

### GetTouchStart

`func (o *GlobalEventCaptureSettings) GetTouchStart() bool`

GetTouchStart returns the TouchStart field if non-nil, zero value otherwise.

### GetTouchStartOk

`func (o *GlobalEventCaptureSettings) GetTouchStartOk() (*bool, bool)`

GetTouchStartOk returns a tuple with the TouchStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTouchStart

`func (o *GlobalEventCaptureSettings) SetTouchStart(v bool)`

SetTouchStart sets TouchStart field to given value.

### HasTouchStart

`func (o *GlobalEventCaptureSettings) HasTouchStart() bool`

HasTouchStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


