/*
Uptrends API v4

This document describes Uptrends API version 4. This Swagger environment also lets you execute API methods directly.  Please note that this is not a sandbox environment: these API methods operate directly on your actual Uptrends account.  For more information, please visit https://www.uptrends.com/api.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uptrends

import (
	"encoding/json"
)

// EscalationLevel struct for EscalationLevel
type EscalationLevel struct {
	EscalationMode *EscalationMode `json:"EscalationMode,omitempty"`
	Id *int32 `json:"Id,omitempty"`
	ThresholdErrorCount *int32 `json:"ThresholdErrorCount,omitempty"`
	ThresholdMinutes *int32 `json:"ThresholdMinutes,omitempty"`
	IsActive *bool `json:"IsActive,omitempty"`
	Message *string `json:"Message,omitempty"`
	NumberOfReminders *int32 `json:"NumberOfReminders,omitempty"`
	ReminderDelay *int32 `json:"ReminderDelay,omitempty"`
	IncludeTraceRoute *bool `json:"IncludeTraceRoute,omitempty"`
	// Hash corresponding with this escalation level.
	Hash *string `json:"Hash,omitempty"`
}

// NewEscalationLevel instantiates a new EscalationLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEscalationLevel() *EscalationLevel {
	this := EscalationLevel{}
	return &this
}

// NewEscalationLevelWithDefaults instantiates a new EscalationLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEscalationLevelWithDefaults() *EscalationLevel {
	this := EscalationLevel{}
	return &this
}

// GetEscalationMode returns the EscalationMode field value if set, zero value otherwise.
func (o *EscalationLevel) GetEscalationMode() EscalationMode {
	if o == nil || isNil(o.EscalationMode) {
		var ret EscalationMode
		return ret
	}
	return *o.EscalationMode
}

// GetEscalationModeOk returns a tuple with the EscalationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetEscalationModeOk() (*EscalationMode, bool) {
	if o == nil || isNil(o.EscalationMode) {
    return nil, false
	}
	return o.EscalationMode, true
}

// HasEscalationMode returns a boolean if a field has been set.
func (o *EscalationLevel) HasEscalationMode() bool {
	if o != nil && !isNil(o.EscalationMode) {
		return true
	}

	return false
}

// SetEscalationMode gets a reference to the given EscalationMode and assigns it to the EscalationMode field.
func (o *EscalationLevel) SetEscalationMode(v EscalationMode) {
	o.EscalationMode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EscalationLevel) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EscalationLevel) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *EscalationLevel) SetId(v int32) {
	o.Id = &v
}

// GetThresholdErrorCount returns the ThresholdErrorCount field value if set, zero value otherwise.
func (o *EscalationLevel) GetThresholdErrorCount() int32 {
	if o == nil || isNil(o.ThresholdErrorCount) {
		var ret int32
		return ret
	}
	return *o.ThresholdErrorCount
}

// GetThresholdErrorCountOk returns a tuple with the ThresholdErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetThresholdErrorCountOk() (*int32, bool) {
	if o == nil || isNil(o.ThresholdErrorCount) {
    return nil, false
	}
	return o.ThresholdErrorCount, true
}

// HasThresholdErrorCount returns a boolean if a field has been set.
func (o *EscalationLevel) HasThresholdErrorCount() bool {
	if o != nil && !isNil(o.ThresholdErrorCount) {
		return true
	}

	return false
}

// SetThresholdErrorCount gets a reference to the given int32 and assigns it to the ThresholdErrorCount field.
func (o *EscalationLevel) SetThresholdErrorCount(v int32) {
	o.ThresholdErrorCount = &v
}

// GetThresholdMinutes returns the ThresholdMinutes field value if set, zero value otherwise.
func (o *EscalationLevel) GetThresholdMinutes() int32 {
	if o == nil || isNil(o.ThresholdMinutes) {
		var ret int32
		return ret
	}
	return *o.ThresholdMinutes
}

// GetThresholdMinutesOk returns a tuple with the ThresholdMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetThresholdMinutesOk() (*int32, bool) {
	if o == nil || isNil(o.ThresholdMinutes) {
    return nil, false
	}
	return o.ThresholdMinutes, true
}

// HasThresholdMinutes returns a boolean if a field has been set.
func (o *EscalationLevel) HasThresholdMinutes() bool {
	if o != nil && !isNil(o.ThresholdMinutes) {
		return true
	}

	return false
}

// SetThresholdMinutes gets a reference to the given int32 and assigns it to the ThresholdMinutes field.
func (o *EscalationLevel) SetThresholdMinutes(v int32) {
	o.ThresholdMinutes = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *EscalationLevel) GetIsActive() bool {
	if o == nil || isNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetIsActiveOk() (*bool, bool) {
	if o == nil || isNil(o.IsActive) {
    return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *EscalationLevel) HasIsActive() bool {
	if o != nil && !isNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *EscalationLevel) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EscalationLevel) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EscalationLevel) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EscalationLevel) SetMessage(v string) {
	o.Message = &v
}

// GetNumberOfReminders returns the NumberOfReminders field value if set, zero value otherwise.
func (o *EscalationLevel) GetNumberOfReminders() int32 {
	if o == nil || isNil(o.NumberOfReminders) {
		var ret int32
		return ret
	}
	return *o.NumberOfReminders
}

// GetNumberOfRemindersOk returns a tuple with the NumberOfReminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetNumberOfRemindersOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfReminders) {
    return nil, false
	}
	return o.NumberOfReminders, true
}

// HasNumberOfReminders returns a boolean if a field has been set.
func (o *EscalationLevel) HasNumberOfReminders() bool {
	if o != nil && !isNil(o.NumberOfReminders) {
		return true
	}

	return false
}

// SetNumberOfReminders gets a reference to the given int32 and assigns it to the NumberOfReminders field.
func (o *EscalationLevel) SetNumberOfReminders(v int32) {
	o.NumberOfReminders = &v
}

// GetReminderDelay returns the ReminderDelay field value if set, zero value otherwise.
func (o *EscalationLevel) GetReminderDelay() int32 {
	if o == nil || isNil(o.ReminderDelay) {
		var ret int32
		return ret
	}
	return *o.ReminderDelay
}

// GetReminderDelayOk returns a tuple with the ReminderDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetReminderDelayOk() (*int32, bool) {
	if o == nil || isNil(o.ReminderDelay) {
    return nil, false
	}
	return o.ReminderDelay, true
}

// HasReminderDelay returns a boolean if a field has been set.
func (o *EscalationLevel) HasReminderDelay() bool {
	if o != nil && !isNil(o.ReminderDelay) {
		return true
	}

	return false
}

// SetReminderDelay gets a reference to the given int32 and assigns it to the ReminderDelay field.
func (o *EscalationLevel) SetReminderDelay(v int32) {
	o.ReminderDelay = &v
}

// GetIncludeTraceRoute returns the IncludeTraceRoute field value if set, zero value otherwise.
func (o *EscalationLevel) GetIncludeTraceRoute() bool {
	if o == nil || isNil(o.IncludeTraceRoute) {
		var ret bool
		return ret
	}
	return *o.IncludeTraceRoute
}

// GetIncludeTraceRouteOk returns a tuple with the IncludeTraceRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetIncludeTraceRouteOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeTraceRoute) {
    return nil, false
	}
	return o.IncludeTraceRoute, true
}

// HasIncludeTraceRoute returns a boolean if a field has been set.
func (o *EscalationLevel) HasIncludeTraceRoute() bool {
	if o != nil && !isNil(o.IncludeTraceRoute) {
		return true
	}

	return false
}

// SetIncludeTraceRoute gets a reference to the given bool and assigns it to the IncludeTraceRoute field.
func (o *EscalationLevel) SetIncludeTraceRoute(v bool) {
	o.IncludeTraceRoute = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *EscalationLevel) GetHash() string {
	if o == nil || isNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationLevel) GetHashOk() (*string, bool) {
	if o == nil || isNil(o.Hash) {
    return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *EscalationLevel) HasHash() bool {
	if o != nil && !isNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *EscalationLevel) SetHash(v string) {
	o.Hash = &v
}

func (o EscalationLevel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EscalationMode) {
		toSerialize["EscalationMode"] = o.EscalationMode
	}
	if !isNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !isNil(o.ThresholdErrorCount) {
		toSerialize["ThresholdErrorCount"] = o.ThresholdErrorCount
	}
	if !isNil(o.ThresholdMinutes) {
		toSerialize["ThresholdMinutes"] = o.ThresholdMinutes
	}
	if !isNil(o.IsActive) {
		toSerialize["IsActive"] = o.IsActive
	}
	if !isNil(o.Message) {
		toSerialize["Message"] = o.Message
	}
	if !isNil(o.NumberOfReminders) {
		toSerialize["NumberOfReminders"] = o.NumberOfReminders
	}
	if !isNil(o.ReminderDelay) {
		toSerialize["ReminderDelay"] = o.ReminderDelay
	}
	if !isNil(o.IncludeTraceRoute) {
		toSerialize["IncludeTraceRoute"] = o.IncludeTraceRoute
	}
	if !isNil(o.Hash) {
		toSerialize["Hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullableEscalationLevel struct {
	value *EscalationLevel
	isSet bool
}

func (v NullableEscalationLevel) Get() *EscalationLevel {
	return v.value
}

func (v *NullableEscalationLevel) Set(val *EscalationLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableEscalationLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableEscalationLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEscalationLevel(val *EscalationLevel) *NullableEscalationLevel {
	return &NullableEscalationLevel{value: val, isSet: true}
}

func (v NullableEscalationLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEscalationLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


